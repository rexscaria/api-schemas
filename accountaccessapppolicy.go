// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/stainless-sdks/cf-rex-go/internal/apijson"
	"github.com/stainless-sdks/cf-rex-go/internal/param"
	"github.com/stainless-sdks/cf-rex-go/internal/requestconfig"
	"github.com/stainless-sdks/cf-rex-go/option"
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

// Creates a policy applying exclusive to a single application that defines the
// users or groups who can reach it. We recommend creating a reusable policy
// instead and subsequently referencing its ID in the application's 'policies'
// array.
func (r *AccountAccessAppPolicyService) New(ctx context.Context, accountID string, appID string, body AccountAccessAppPolicyNewParams, opts ...option.RequestOption) (res *SingleResponsePolicy, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if appID == "" {
		err = errors.New("missing required app_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/access/apps/%s/policies", accountID, appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Fetches a single Access policy configured for an application. Returns both
// exclusively owned and reusable policies used by the application.
func (r *AccountAccessAppPolicyService) Get(ctx context.Context, accountID string, appID string, policyID string, opts ...option.RequestOption) (res *SingleResponsePolicy, err error) {
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
	path := fmt.Sprintf("accounts/%s/access/apps/%s/policies/%s", accountID, appID, policyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates an Access policy specific to an application. To update a reusable
// policy, use the /accounts/{account_id}/policies/{uid} endpoint.
func (r *AccountAccessAppPolicyService) Update(ctx context.Context, accountID string, appID string, policyID string, body AccountAccessAppPolicyUpdateParams, opts ...option.RequestOption) (res *SingleResponsePolicy, err error) {
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
	path := fmt.Sprintf("accounts/%s/access/apps/%s/policies/%s", accountID, appID, policyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Lists Access policies configured for an application. Returns both exclusively
// scoped and reusable policies used by the application.
func (r *AccountAccessAppPolicyService) List(ctx context.Context, accountID string, appID string, opts ...option.RequestOption) (res *ResponseCollectionAppPolicies, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if appID == "" {
		err = errors.New("missing required app_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/access/apps/%s/policies", accountID, appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Deletes an Access policy specific to an application. To delete a reusable
// policy, use the /accounts/{account_id}/policies/{uid} endpoint.
func (r *AccountAccessAppPolicyService) Delete(ctx context.Context, accountID string, appID string, policyID string, opts ...option.RequestOption) (res *IDResponseApps, err error) {
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
	path := fmt.Sprintf("accounts/%s/access/apps/%s/policies/%s", accountID, appID, policyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
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

type PolicyRequestForAppParam struct {
	// The order of execution for this policy. Must be unique for each policy within an
	// app.
	Precedence param.Field[int64] `json:"precedence"`
	PolicyRequestForAccessParam
}

func (r PolicyRequestForAppParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PolicyResponseApp struct {
	// The order of execution for this policy. Must be unique for each policy within an
	// app.
	Precedence int64                 `json:"precedence"`
	JSON       policyResponseAppJSON `json:"-"`
	PolicyResponseGeneral
}

// policyResponseAppJSON contains the JSON metadata for the struct
// [PolicyResponseApp]
type policyResponseAppJSON struct {
	Precedence  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PolicyResponseApp) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r policyResponseAppJSON) RawJSON() string {
	return r.raw
}

type PolicyResponseGeneral struct {
	// Administrators who can approve a temporary authentication request.
	ApprovalGroups []ApprovalGroupEmail `json:"approval_groups"`
	// Requires the user to request access from an administrator at the start of each
	// session.
	ApprovalRequired bool `json:"approval_required"`
	// Require this application to be served in an isolated browser for users matching
	// this policy. 'Client Web Isolation' must be on for the account in order to use
	// this feature.
	IsolationRequired bool `json:"isolation_required"`
	// A custom message that will appear on the purpose justification screen.
	PurposeJustificationPrompt string `json:"purpose_justification_prompt"`
	// Require users to enter a justification when they log in to the application.
	PurposeJustificationRequired bool `json:"purpose_justification_required"`
	// The amount of time that tokens issued for the application will be valid. Must be
	// in the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s,
	// m, h.
	SessionDuration string                    `json:"session_duration"`
	JSON            policyResponseGeneralJSON `json:"-"`
	BasePolicyResponse
}

// policyResponseGeneralJSON contains the JSON metadata for the struct
// [PolicyResponseGeneral]
type policyResponseGeneralJSON struct {
	ApprovalGroups               apijson.Field
	ApprovalRequired             apijson.Field
	IsolationRequired            apijson.Field
	PurposeJustificationPrompt   apijson.Field
	PurposeJustificationRequired apijson.Field
	SessionDuration              apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *PolicyResponseGeneral) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r policyResponseGeneralJSON) RawJSON() string {
	return r.raw
}

type ResponseCollectionAppPolicies struct {
	Result []PolicyResponseApp               `json:"result"`
	JSON   responseCollectionAppPoliciesJSON `json:"-"`
	APIResponseCollectionAccess
}

// responseCollectionAppPoliciesJSON contains the JSON metadata for the struct
// [ResponseCollectionAppPolicies]
type responseCollectionAppPoliciesJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResponseCollectionAppPolicies) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r responseCollectionAppPoliciesJSON) RawJSON() string {
	return r.raw
}

type SingleResponsePolicy struct {
	Result PolicyResponseApp        `json:"result"`
	JSON   singleResponsePolicyJSON `json:"-"`
	APIResponseSingleAccess
}

// singleResponsePolicyJSON contains the JSON metadata for the struct
// [SingleResponsePolicy]
type singleResponsePolicyJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponsePolicy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r singleResponsePolicyJSON) RawJSON() string {
	return r.raw
}

type AccountAccessAppPolicyNewParams struct {
	PolicyRequestForApp PolicyRequestForAppParam `json:"policy_request_for_app,required"`
}

func (r AccountAccessAppPolicyNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.PolicyRequestForApp)
}

type AccountAccessAppPolicyUpdateParams struct {
	PolicyRequestForApp PolicyRequestForAppParam `json:"policy_request_for_app,required"`
}

func (r AccountAccessAppPolicyUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.PolicyRequestForApp)
}
