// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/cf-rex-go/internal/apijson"
	"github.com/stainless-sdks/cf-rex-go/internal/param"
	"github.com/stainless-sdks/cf-rex-go/internal/requestconfig"
	"github.com/stainless-sdks/cf-rex-go/option"
)

// AccountAccessPolicyService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountAccessPolicyService] method instead.
type AccountAccessPolicyService struct {
	Options []option.RequestOption
}

// NewAccountAccessPolicyService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountAccessPolicyService(opts ...option.RequestOption) (r *AccountAccessPolicyService) {
	r = &AccountAccessPolicyService{}
	r.Options = opts
	return
}

// Creates a new Access reusable policy.
func (r *AccountAccessPolicyService) New(ctx context.Context, accountID string, body AccountAccessPolicyNewParams, opts ...option.RequestOption) (res *ReusableSingleResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/access/policies", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Fetches a single Access reusable policy.
func (r *AccountAccessPolicyService) Get(ctx context.Context, accountID string, policyID string, opts ...option.RequestOption) (res *ReusableSingleResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if policyID == "" {
		err = errors.New("missing required policy_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/access/policies/%s", accountID, policyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates a Access reusable policy.
func (r *AccountAccessPolicyService) Update(ctx context.Context, accountID string, policyID string, body AccountAccessPolicyUpdateParams, opts ...option.RequestOption) (res *ReusableSingleResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if policyID == "" {
		err = errors.New("missing required policy_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/access/policies/%s", accountID, policyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Lists Access reusable policies.
func (r *AccountAccessPolicyService) List(ctx context.Context, accountID string, opts ...option.RequestOption) (res *AccountAccessPolicyListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/access/policies", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Deletes an Access reusable policy.
func (r *AccountAccessPolicyService) Delete(ctx context.Context, accountID string, policyID string, opts ...option.RequestOption) (res *AccountAccessPolicyDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if policyID == "" {
		err = errors.New("missing required policy_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/access/policies/%s", accountID, policyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// The action Access will take if a user matches this policy. Infrastructure
// application policies can only use the Allow action.
type AccessDecision string

const (
	AccessDecisionAllow       AccessDecision = "allow"
	AccessDecisionDeny        AccessDecision = "deny"
	AccessDecisionNonIdentity AccessDecision = "non_identity"
	AccessDecisionBypass      AccessDecision = "bypass"
)

func (r AccessDecision) IsKnown() bool {
	switch r {
	case AccessDecisionAllow, AccessDecisionDeny, AccessDecisionNonIdentity, AccessDecisionBypass:
		return true
	}
	return false
}

// A group of email addresses that can approve a temporary authentication request.
type ApprovalGroupEmail struct {
	// The number of approvals needed to obtain access.
	ApprovalsNeeded float64 `json:"approvals_needed,required"`
	// A list of emails that can approve the access request.
	EmailAddresses []string `json:"email_addresses"`
	// The UUID of an re-usable email list.
	EmailListUuid string                 `json:"email_list_uuid"`
	JSON          approvalGroupEmailJSON `json:"-"`
}

// approvalGroupEmailJSON contains the JSON metadata for the struct
// [ApprovalGroupEmail]
type approvalGroupEmailJSON struct {
	ApprovalsNeeded apijson.Field
	EmailAddresses  apijson.Field
	EmailListUuid   apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ApprovalGroupEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r approvalGroupEmailJSON) RawJSON() string {
	return r.raw
}

// A group of email addresses that can approve a temporary authentication request.
type ApprovalGroupEmailParam struct {
	// The number of approvals needed to obtain access.
	ApprovalsNeeded param.Field[float64] `json:"approvals_needed,required"`
	// A list of emails that can approve the access request.
	EmailAddresses param.Field[[]string] `json:"email_addresses"`
	// The UUID of an re-usable email list.
	EmailListUuid param.Field[string] `json:"email_list_uuid"`
}

func (r ApprovalGroupEmailParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BasePolicyRequestParam struct {
	// The action Access will take if a user matches this policy. Infrastructure
	// application policies can only use the Allow action.
	Decision param.Field[AccessDecision] `json:"decision,required"`
	// Rules evaluated with an OR logical operator. A user needs to meet only one of
	// the Include rules.
	Include param.Field[[]AccessRuleUnionParam] `json:"include,required"`
	// The name of the Access policy.
	Name param.Field[string] `json:"name,required"`
	// Rules evaluated with a NOT logical operator. To match the policy, a user cannot
	// meet any of the Exclude rules.
	Exclude param.Field[[]AccessRuleUnionParam] `json:"exclude"`
	// Rules evaluated with an AND logical operator. To match the policy, a user must
	// meet all of the Require rules.
	Require param.Field[[]AccessRuleUnionParam] `json:"require"`
}

func (r BasePolicyRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PolicyRequestForAccessParam struct {
	// Administrators who can approve a temporary authentication request.
	ApprovalGroups param.Field[[]ApprovalGroupEmailParam] `json:"approval_groups"`
	// Requires the user to request access from an administrator at the start of each
	// session.
	ApprovalRequired param.Field[bool] `json:"approval_required"`
	// Require this application to be served in an isolated browser for users matching
	// this policy. 'Client Web Isolation' must be on for the account in order to use
	// this feature.
	IsolationRequired param.Field[bool] `json:"isolation_required"`
	// A custom message that will appear on the purpose justification screen.
	PurposeJustificationPrompt param.Field[string] `json:"purpose_justification_prompt"`
	// Require users to enter a justification when they log in to the application.
	PurposeJustificationRequired param.Field[bool] `json:"purpose_justification_required"`
	// The amount of time that tokens issued for the application will be valid. Must be
	// in the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s,
	// m, h.
	SessionDuration param.Field[string] `json:"session_duration"`
	BasePolicyRequestParam
}

func (r PolicyRequestForAccessParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PolicyRequestForAccessParam) ImplementsAccountAccessPolicyTestStartParamsPolicyUnion() {}

type ReusablePolicyResponse struct {
	// Number of access applications currently using this policy.
	AppCount int64                          `json:"app_count"`
	Reusable ReusablePolicyResponseReusable `json:"reusable"`
	JSON     reusablePolicyResponseJSON     `json:"-"`
	PolicyResponseGeneral
}

// reusablePolicyResponseJSON contains the JSON metadata for the struct
// [ReusablePolicyResponse]
type reusablePolicyResponseJSON struct {
	AppCount    apijson.Field
	Reusable    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ReusablePolicyResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r reusablePolicyResponseJSON) RawJSON() string {
	return r.raw
}

type ReusablePolicyResponseReusable bool

const (
	ReusablePolicyResponseReusableTrue ReusablePolicyResponseReusable = true
)

func (r ReusablePolicyResponseReusable) IsKnown() bool {
	switch r {
	case ReusablePolicyResponseReusableTrue:
		return true
	}
	return false
}

type ReusableSingleResponse struct {
	Result ReusablePolicyResponse     `json:"result"`
	JSON   reusableSingleResponseJSON `json:"-"`
	APIResponseSingleAccess
}

// reusableSingleResponseJSON contains the JSON metadata for the struct
// [ReusableSingleResponse]
type reusableSingleResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ReusableSingleResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r reusableSingleResponseJSON) RawJSON() string {
	return r.raw
}

type AccountAccessPolicyListResponse struct {
	Result []ReusablePolicyResponse            `json:"result"`
	JSON   accountAccessPolicyListResponseJSON `json:"-"`
	APIResponseCollectionAccess
}

// accountAccessPolicyListResponseJSON contains the JSON metadata for the struct
// [AccountAccessPolicyListResponse]
type accountAccessPolicyListResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessPolicyListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAccessPolicyListResponseJSON) RawJSON() string {
	return r.raw
}

type AccountAccessPolicyDeleteResponse struct {
	Result AccountAccessPolicyDeleteResponseResult `json:"result"`
	JSON   accountAccessPolicyDeleteResponseJSON   `json:"-"`
	APIResponseSingleAccess
}

// accountAccessPolicyDeleteResponseJSON contains the JSON metadata for the struct
// [AccountAccessPolicyDeleteResponse]
type accountAccessPolicyDeleteResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessPolicyDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAccessPolicyDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type AccountAccessPolicyDeleteResponseResult struct {
	// The UUID of the policy
	ID   string                                      `json:"id"`
	JSON accountAccessPolicyDeleteResponseResultJSON `json:"-"`
}

// accountAccessPolicyDeleteResponseResultJSON contains the JSON metadata for the
// struct [AccountAccessPolicyDeleteResponseResult]
type accountAccessPolicyDeleteResponseResultJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessPolicyDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAccessPolicyDeleteResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountAccessPolicyNewParams struct {
	PolicyRequestForAccess PolicyRequestForAccessParam `json:"policy_request_for_access,required"`
}

func (r AccountAccessPolicyNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.PolicyRequestForAccess)
}

type AccountAccessPolicyUpdateParams struct {
	PolicyRequestForAccess PolicyRequestForAccessParam `json:"policy_request_for_access,required"`
}

func (r AccountAccessPolicyUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.PolicyRequestForAccess)
}
