// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/apiquery"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
)

// AccountAccessPolicyTestService contains methods and other services that help
// with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountAccessPolicyTestService] method instead.
type AccountAccessPolicyTestService struct {
	Options []option.RequestOption
}

// NewAccountAccessPolicyTestService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountAccessPolicyTestService(opts ...option.RequestOption) (r *AccountAccessPolicyTestService) {
	r = &AccountAccessPolicyTestService{}
	r.Options = opts
	return
}

// Fetches the current status of a given Access policy test.
func (r *AccountAccessPolicyTestService) Get(ctx context.Context, accountID string, policyTestID string, query AccountAccessPolicyTestGetParams, opts ...option.RequestOption) (res *AccountAccessPolicyTestGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if policyTestID == "" {
		err = errors.New("missing required policy_test_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/access/policy-tests/%s", accountID, policyTestID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Starts an Access policy test.
func (r *AccountAccessPolicyTestService) Start(ctx context.Context, accountID string, body AccountAccessPolicyTestStartParams, opts ...option.RequestOption) (res *AccountAccessPolicyTestStartResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/access/policy-tests", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Fetches a single page of user results from an Access policy test.
func (r *AccountAccessPolicyTestService) Users(ctx context.Context, accountID string, policyTestID string, query AccountAccessPolicyTestUsersParams, opts ...option.RequestOption) (res *AccountAccessPolicyTestUsersResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if policyTestID == "" {
		err = errors.New("missing required policy_test_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/access/policy-tests/%s/users", accountID, policyTestID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type AccountAccessPolicyTestGetResponse struct {
	Result AccountAccessPolicyTestGetResponseResult `json:"result"`
	JSON   accountAccessPolicyTestGetResponseJSON   `json:"-"`
	APIResponseSingleAccess
}

// accountAccessPolicyTestGetResponseJSON contains the JSON metadata for the struct
// [AccountAccessPolicyTestGetResponse]
type accountAccessPolicyTestGetResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessPolicyTestGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAccessPolicyTestGetResponseJSON) RawJSON() string {
	return r.raw
}

type AccountAccessPolicyTestGetResponseResult struct {
	// The UUID of the policy test.
	ID string `json:"id"`
	// The number of pages of (processed) users.
	PagesProcessed int64 `json:"pages_processed"`
	// The percentage of (processed) users approved based on policy evaluation results.
	PercentApproved int64 `json:"percent_approved"`
	// The percentage of (processed) users blocked based on policy evaluation results.
	PercentBlocked int64 `json:"percent_blocked"`
	// The percentage of users processed so far (of the entire user base).
	PercentUsersProcessed int64 `json:"percent_users_processed"`
	// The status of the policy test.
	Status AccountAccessPolicyTestGetResponseResultStatus `json:"status"`
	// The total number of users in the user base.
	TotalUsers int64 `json:"total_users"`
	// The number of (processed) users approved based on policy evaluation results.
	UsersApproved int64 `json:"users_approved"`
	// The number of (processed) users blocked based on policy evaluation results.
	UsersBlocked int64                                        `json:"users_blocked"`
	JSON         accountAccessPolicyTestGetResponseResultJSON `json:"-"`
}

// accountAccessPolicyTestGetResponseResultJSON contains the JSON metadata for the
// struct [AccountAccessPolicyTestGetResponseResult]
type accountAccessPolicyTestGetResponseResultJSON struct {
	ID                    apijson.Field
	PagesProcessed        apijson.Field
	PercentApproved       apijson.Field
	PercentBlocked        apijson.Field
	PercentUsersProcessed apijson.Field
	Status                apijson.Field
	TotalUsers            apijson.Field
	UsersApproved         apijson.Field
	UsersBlocked          apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *AccountAccessPolicyTestGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAccessPolicyTestGetResponseResultJSON) RawJSON() string {
	return r.raw
}

// The status of the policy test.
type AccountAccessPolicyTestGetResponseResultStatus string

const (
	AccountAccessPolicyTestGetResponseResultStatusBlocked    AccountAccessPolicyTestGetResponseResultStatus = "blocked"
	AccountAccessPolicyTestGetResponseResultStatusProcessing AccountAccessPolicyTestGetResponseResultStatus = "processing"
	AccountAccessPolicyTestGetResponseResultStatusComplete   AccountAccessPolicyTestGetResponseResultStatus = "complete"
)

func (r AccountAccessPolicyTestGetResponseResultStatus) IsKnown() bool {
	switch r {
	case AccountAccessPolicyTestGetResponseResultStatusBlocked, AccountAccessPolicyTestGetResponseResultStatusProcessing, AccountAccessPolicyTestGetResponseResultStatusComplete:
		return true
	}
	return false
}

type AccountAccessPolicyTestStartResponse struct {
	Result AccountAccessPolicyTestStartResponseResult `json:"result"`
	JSON   accountAccessPolicyTestStartResponseJSON   `json:"-"`
	APIResponseSingleAccess
}

// accountAccessPolicyTestStartResponseJSON contains the JSON metadata for the
// struct [AccountAccessPolicyTestStartResponse]
type accountAccessPolicyTestStartResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessPolicyTestStartResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAccessPolicyTestStartResponseJSON) RawJSON() string {
	return r.raw
}

type AccountAccessPolicyTestStartResponseResult struct {
	// The UUID of the policy test.
	ID string `json:"id"`
	// The status of the policy test request.
	Status AccountAccessPolicyTestStartResponseResultStatus `json:"status"`
	JSON   accountAccessPolicyTestStartResponseResultJSON   `json:"-"`
}

// accountAccessPolicyTestStartResponseResultJSON contains the JSON metadata for
// the struct [AccountAccessPolicyTestStartResponseResult]
type accountAccessPolicyTestStartResponseResultJSON struct {
	ID          apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessPolicyTestStartResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAccessPolicyTestStartResponseResultJSON) RawJSON() string {
	return r.raw
}

// The status of the policy test request.
type AccountAccessPolicyTestStartResponseResultStatus string

const (
	AccountAccessPolicyTestStartResponseResultStatusSuccess AccountAccessPolicyTestStartResponseResultStatus = "success"
)

func (r AccountAccessPolicyTestStartResponseResultStatus) IsKnown() bool {
	switch r {
	case AccountAccessPolicyTestStartResponseResultStatusSuccess:
		return true
	}
	return false
}

type AccountAccessPolicyTestUsersResponse struct {
	// Page of processed users.
	Result []AccountAccessPolicyTestUsersResponseResult `json:"result"`
	JSON   accountAccessPolicyTestUsersResponseJSON     `json:"-"`
	APIResponseSingleAccess
}

// accountAccessPolicyTestUsersResponseJSON contains the JSON metadata for the
// struct [AccountAccessPolicyTestUsersResponse]
type accountAccessPolicyTestUsersResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessPolicyTestUsersResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAccessPolicyTestUsersResponseJSON) RawJSON() string {
	return r.raw
}

type AccountAccessPolicyTestUsersResponseResult struct {
	// UUID
	ID string `json:"id"`
	// The email of the user.
	Email string `json:"email" format:"email"`
	// The name of the user.
	Name string `json:"name"`
	// Policy evaluation result for an individual user.
	Status AccountAccessPolicyTestUsersResponseResultStatus `json:"status"`
	JSON   accountAccessPolicyTestUsersResponseResultJSON   `json:"-"`
}

// accountAccessPolicyTestUsersResponseResultJSON contains the JSON metadata for
// the struct [AccountAccessPolicyTestUsersResponseResult]
type accountAccessPolicyTestUsersResponseResultJSON struct {
	ID          apijson.Field
	Email       apijson.Field
	Name        apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessPolicyTestUsersResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAccessPolicyTestUsersResponseResultJSON) RawJSON() string {
	return r.raw
}

// Policy evaluation result for an individual user.
type AccountAccessPolicyTestUsersResponseResultStatus string

const (
	AccountAccessPolicyTestUsersResponseResultStatusApproved AccountAccessPolicyTestUsersResponseResultStatus = "approved"
	AccountAccessPolicyTestUsersResponseResultStatusBlocked  AccountAccessPolicyTestUsersResponseResultStatus = "blocked"
)

func (r AccountAccessPolicyTestUsersResponseResultStatus) IsKnown() bool {
	switch r {
	case AccountAccessPolicyTestUsersResponseResultStatusApproved, AccountAccessPolicyTestUsersResponseResultStatusBlocked:
		return true
	}
	return false
}

type AccountAccessPolicyTestGetParams struct {
	Page param.Field[int64] `query:"page"`
}

// URLQuery serializes [AccountAccessPolicyTestGetParams]'s query parameters as
// `url.Values`.
func (r AccountAccessPolicyTestGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountAccessPolicyTestStartParams struct {
	Policies param.Field[[]AccountAccessPolicyTestStartParamsPolicyUnion] `json:"policies"`
}

func (r AccountAccessPolicyTestStartParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The UUID of the reusable policy you wish to test
//
// Satisfied by [PolicyRequestForAccessParam], [shared.UnionString].
type AccountAccessPolicyTestStartParamsPolicyUnion interface {
	ImplementsAccountAccessPolicyTestStartParamsPolicyUnion()
}

type AccountAccessPolicyTestUsersParams struct {
	// Filter users by their policy evaluation status.
	Status param.Field[AccountAccessPolicyTestUsersParamsStatus] `query:"status"`
}

// URLQuery serializes [AccountAccessPolicyTestUsersParams]'s query parameters as
// `url.Values`.
func (r AccountAccessPolicyTestUsersParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter users by their policy evaluation status.
type AccountAccessPolicyTestUsersParamsStatus string

const (
	AccountAccessPolicyTestUsersParamsStatusSuccess AccountAccessPolicyTestUsersParamsStatus = "success"
	AccountAccessPolicyTestUsersParamsStatusFail    AccountAccessPolicyTestUsersParamsStatus = "fail"
)

func (r AccountAccessPolicyTestUsersParamsStatus) IsKnown() bool {
	switch r {
	case AccountAccessPolicyTestUsersParamsStatusSuccess, AccountAccessPolicyTestUsersParamsStatusFail:
		return true
	}
	return false
}
