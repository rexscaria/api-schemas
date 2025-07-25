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

// AccountAccessAppPolicyService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountAccessAppPolicyService] method instead.
type AccountAccessAppPolicyService struct {
	Options []option.RequestOption
}

// NewAccountAccessAppPolicyService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountAccessAppPolicyService(opts ...option.RequestOption) (r *AccountAccessAppPolicyService) {
	r = &AccountAccessAppPolicyService{}
	r.Options = opts
	return
}

// Converts an application-scoped policy to a reusable policy. The policy will no
// longer be exclusively scoped to the application. Further updates to the policy
// should go through the /accounts/{account_id}/policies/{uid} endpoint.
func (r *AccountAccessAppPolicyService) MakeReusable(ctx context.Context, accountID string, appID string, policyID string, opts ...option.RequestOption) (res *ResponseCollectionAppPolicies, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if appID == "" {
		err = errors.New("missing required app_id parameter")
		return
	}
	if policyID == "" {
		err = errors.New("missing required policy_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/access/apps/%s/policies/%s/make_reusable", accountID, appID, policyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, nil, &res, opts...)
	return
}

type BasePolicyResponse struct {
	// The UUID of the policy
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The action Access will take if a user matches this policy. Infrastructure
	// application policies can only use the Allow action.
	Decision AccessDecision `json:"decision"`
	// Rules evaluated with a NOT logical operator. To match the policy, a user cannot
	// meet any of the Exclude rules.
	Exclude []AccessRule `json:"exclude"`
	// Rules evaluated with an OR logical operator. A user needs to meet only one of
	// the Include rules.
	Include []AccessRule `json:"include"`
	// The name of the Access policy.
	Name string `json:"name"`
	// Rules evaluated with an AND logical operator. To match the policy, a user must
	// meet all of the Require rules.
	Require   []AccessRule           `json:"require"`
	UpdatedAt time.Time              `json:"updated_at" format:"date-time"`
	JSON      basePolicyResponseJSON `json:"-"`
}

// basePolicyResponseJSON contains the JSON metadata for the struct
// [BasePolicyResponse]
type basePolicyResponseJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Decision    apijson.Field
	Exclude     apijson.Field
	Include     apijson.Field
	Name        apijson.Field
	Require     apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BasePolicyResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r basePolicyResponseJSON) RawJSON() string {
	return r.raw
}

type PolicyResponseApp struct {
	// Administrators who can approve a temporary authentication request.
	ApprovalGroups []ApprovalGroupEmail `json:"approval_groups"`
	// Requires the user to request access from an administrator at the start of each
	// session.
	ApprovalRequired bool `json:"approval_required"`
	// Require this application to be served in an isolated browser for users matching
	// this policy. 'Client Web Isolation' must be on for the account in order to use
	// this feature.
	IsolationRequired bool `json:"isolation_required"`
	// The order of execution for this policy. Must be unique for each policy within an
	// app.
	Precedence int64 `json:"precedence"`
	// A custom message that will appear on the purpose justification screen.
	PurposeJustificationPrompt string `json:"purpose_justification_prompt"`
	// Require users to enter a justification when they log in to the application.
	PurposeJustificationRequired bool `json:"purpose_justification_required"`
	// The amount of time that tokens issued for the application will be valid. Must be
	// in the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s,
	// m, h.
	SessionDuration string                `json:"session_duration"`
	JSON            policyResponseAppJSON `json:"-"`
	BasePolicyResponse
}

// policyResponseAppJSON contains the JSON metadata for the struct
// [PolicyResponseApp]
type policyResponseAppJSON struct {
	ApprovalGroups               apijson.Field
	ApprovalRequired             apijson.Field
	IsolationRequired            apijson.Field
	Precedence                   apijson.Field
	PurposeJustificationPrompt   apijson.Field
	PurposeJustificationRequired apijson.Field
	SessionDuration              apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *PolicyResponseApp) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r policyResponseAppJSON) RawJSON() string {
	return r.raw
}

type ResponseCollectionAppPolicies struct {
	Errors   []MessagesAccessItem `json:"errors,required"`
	Messages []MessagesAccessItem `json:"messages,required"`
	// Whether the API call was successful.
	Success    ResponseCollectionAppPoliciesSuccess    `json:"success,required"`
	Result     []PolicyResponseApp                     `json:"result"`
	ResultInfo ResponseCollectionAppPoliciesResultInfo `json:"result_info"`
	JSON       responseCollectionAppPoliciesJSON       `json:"-"`
}

// responseCollectionAppPoliciesJSON contains the JSON metadata for the struct
// [ResponseCollectionAppPolicies]
type responseCollectionAppPoliciesJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResponseCollectionAppPolicies) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r responseCollectionAppPoliciesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ResponseCollectionAppPoliciesSuccess bool

const (
	ResponseCollectionAppPoliciesSuccessTrue ResponseCollectionAppPoliciesSuccess = true
)

func (r ResponseCollectionAppPoliciesSuccess) IsKnown() bool {
	switch r {
	case ResponseCollectionAppPoliciesSuccessTrue:
		return true
	}
	return false
}

type ResponseCollectionAppPoliciesResultInfo struct {
	// Total number of results for the requested service.
	Count float64 `json:"count"`
	// Current page within paginated list of results.
	Page float64 `json:"page"`
	// Number of results per page of results.
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters.
	TotalCount float64                                     `json:"total_count"`
	JSON       responseCollectionAppPoliciesResultInfoJSON `json:"-"`
}

// responseCollectionAppPoliciesResultInfoJSON contains the JSON metadata for the
// struct [ResponseCollectionAppPoliciesResultInfo]
type responseCollectionAppPoliciesResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResponseCollectionAppPoliciesResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r responseCollectionAppPoliciesResultInfoJSON) RawJSON() string {
	return r.raw
}
