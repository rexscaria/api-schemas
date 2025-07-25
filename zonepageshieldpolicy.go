// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
)

// ZonePageShieldPolicyService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZonePageShieldPolicyService] method instead.
type ZonePageShieldPolicyService struct {
	Options []option.RequestOption
}

// NewZonePageShieldPolicyService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZonePageShieldPolicyService(opts ...option.RequestOption) (r *ZonePageShieldPolicyService) {
	r = &ZonePageShieldPolicyService{}
	r.Options = opts
	return
}

// Create a Page Shield policy.
func (r *ZonePageShieldPolicyService) New(ctx context.Context, zoneID string, body ZonePageShieldPolicyNewParams, opts ...option.RequestOption) (res *GetZonePolicyResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/page_shield/policies", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Fetches a Page Shield policy by ID.
func (r *ZonePageShieldPolicyService) Get(ctx context.Context, zoneID string, policyID string, opts ...option.RequestOption) (res *GetZonePolicyResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if policyID == "" {
		err = errors.New("missing required policy_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/page_shield/policies/%s", zoneID, policyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update a Page Shield policy by ID.
func (r *ZonePageShieldPolicyService) Update(ctx context.Context, zoneID string, policyID string, body ZonePageShieldPolicyUpdateParams, opts ...option.RequestOption) (res *GetZonePolicyResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if policyID == "" {
		err = errors.New("missing required policy_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/page_shield/policies/%s", zoneID, policyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Lists all Page Shield policies.
func (r *ZonePageShieldPolicyService) List(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *ZonePageShieldPolicyListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/page_shield/policies", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete a Page Shield policy by ID.
func (r *ZonePageShieldPolicyService) Delete(ctx context.Context, zoneID string, policyID string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if policyID == "" {
		err = errors.New("missing required policy_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/page_shield/policies/%s", zoneID, policyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type GetZonePolicyResponse struct {
	Result PolicyWithID `json:"result,required,nullable"`
	// Whether the API call was successful
	Success  GetZonePolicyResponseSuccess   `json:"success,required"`
	Errors   []GetZonePolicyResponseError   `json:"errors"`
	Messages []GetZonePolicyResponseMessage `json:"messages"`
	JSON     getZonePolicyResponseJSON      `json:"-"`
}

// getZonePolicyResponseJSON contains the JSON metadata for the struct
// [GetZonePolicyResponse]
type getZonePolicyResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	Errors      apijson.Field
	Messages    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GetZonePolicyResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r getZonePolicyResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type GetZonePolicyResponseSuccess bool

const (
	GetZonePolicyResponseSuccessTrue GetZonePolicyResponseSuccess = true
)

func (r GetZonePolicyResponseSuccess) IsKnown() bool {
	switch r {
	case GetZonePolicyResponseSuccessTrue:
		return true
	}
	return false
}

type GetZonePolicyResponseError struct {
	Code             int64                             `json:"code,required"`
	Message          string                            `json:"message,required"`
	DocumentationURL string                            `json:"documentation_url"`
	Source           GetZonePolicyResponseErrorsSource `json:"source"`
	JSON             getZonePolicyResponseErrorJSON    `json:"-"`
}

// getZonePolicyResponseErrorJSON contains the JSON metadata for the struct
// [GetZonePolicyResponseError]
type getZonePolicyResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *GetZonePolicyResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r getZonePolicyResponseErrorJSON) RawJSON() string {
	return r.raw
}

type GetZonePolicyResponseErrorsSource struct {
	Pointer string                                `json:"pointer"`
	JSON    getZonePolicyResponseErrorsSourceJSON `json:"-"`
}

// getZonePolicyResponseErrorsSourceJSON contains the JSON metadata for the struct
// [GetZonePolicyResponseErrorsSource]
type getZonePolicyResponseErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GetZonePolicyResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r getZonePolicyResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type GetZonePolicyResponseMessage struct {
	Code             int64                               `json:"code,required"`
	Message          string                              `json:"message,required"`
	DocumentationURL string                              `json:"documentation_url"`
	Source           GetZonePolicyResponseMessagesSource `json:"source"`
	JSON             getZonePolicyResponseMessageJSON    `json:"-"`
}

// getZonePolicyResponseMessageJSON contains the JSON metadata for the struct
// [GetZonePolicyResponseMessage]
type getZonePolicyResponseMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *GetZonePolicyResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r getZonePolicyResponseMessageJSON) RawJSON() string {
	return r.raw
}

type GetZonePolicyResponseMessagesSource struct {
	Pointer string                                  `json:"pointer"`
	JSON    getZonePolicyResponseMessagesSourceJSON `json:"-"`
}

// getZonePolicyResponseMessagesSourceJSON contains the JSON metadata for the
// struct [GetZonePolicyResponseMessagesSource]
type getZonePolicyResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GetZonePolicyResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r getZonePolicyResponseMessagesSourceJSON) RawJSON() string {
	return r.raw
}

type PolicyParam struct {
	// The action to take if the expression matches
	Action param.Field[PolicyAction] `json:"action,required"`
	// A description for the policy
	Description param.Field[string] `json:"description,required"`
	// Whether the policy is enabled
	Enabled param.Field[bool] `json:"enabled,required"`
	// The expression which must match for the policy to be applied, using the
	// Cloudflare Firewall rule expression syntax
	Expression param.Field[string] `json:"expression,required"`
	// The policy which will be applied
	Value param.Field[string] `json:"value,required"`
}

func (r PolicyParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The action to take if the expression matches
type PolicyAction string

const (
	PolicyActionAllow PolicyAction = "allow"
	PolicyActionLog   PolicyAction = "log"
)

func (r PolicyAction) IsKnown() bool {
	switch r {
	case PolicyActionAllow, PolicyActionLog:
		return true
	}
	return false
}

type PolicyWithID struct {
	// Identifier
	ID string `json:"id,required"`
	// The action to take if the expression matches
	Action PolicyAction `json:"action,required"`
	// A description for the policy
	Description string `json:"description,required"`
	// Whether the policy is enabled
	Enabled bool `json:"enabled,required"`
	// The expression which must match for the policy to be applied, using the
	// Cloudflare Firewall rule expression syntax
	Expression string `json:"expression,required"`
	// The policy which will be applied
	Value string           `json:"value,required"`
	JSON  policyWithIDJSON `json:"-"`
}

// policyWithIDJSON contains the JSON metadata for the struct [PolicyWithID]
type policyWithIDJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Description apijson.Field
	Enabled     apijson.Field
	Expression  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PolicyWithID) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r policyWithIDJSON) RawJSON() string {
	return r.raw
}

type ZonePageShieldPolicyListResponse struct {
	Result     []PolicyWithID                             `json:"result,required"`
	ResultInfo ZonePageShieldPolicyListResponseResultInfo `json:"result_info,required"`
	// Whether the API call was successful
	Success  ZonePageShieldPolicyListResponseSuccess   `json:"success,required"`
	Errors   []ZonePageShieldPolicyListResponseError   `json:"errors"`
	Messages []ZonePageShieldPolicyListResponseMessage `json:"messages"`
	JSON     zonePageShieldPolicyListResponseJSON      `json:"-"`
}

// zonePageShieldPolicyListResponseJSON contains the JSON metadata for the struct
// [ZonePageShieldPolicyListResponse]
type zonePageShieldPolicyListResponseJSON struct {
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	Errors      apijson.Field
	Messages    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonePageShieldPolicyListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zonePageShieldPolicyListResponseJSON) RawJSON() string {
	return r.raw
}

type ZonePageShieldPolicyListResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count,required"`
	// Current page within paginated list of results
	Page float64 `json:"page,required"`
	// Number of results per page of results
	PerPage float64 `json:"per_page,required"`
	// Total results available without any search parameters
	TotalCount float64 `json:"total_count,required"`
	// Total number of pages
	TotalPages float64                                        `json:"total_pages,required"`
	JSON       zonePageShieldPolicyListResponseResultInfoJSON `json:"-"`
}

// zonePageShieldPolicyListResponseResultInfoJSON contains the JSON metadata for
// the struct [ZonePageShieldPolicyListResponseResultInfo]
type zonePageShieldPolicyListResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	TotalPages  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonePageShieldPolicyListResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zonePageShieldPolicyListResponseResultInfoJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ZonePageShieldPolicyListResponseSuccess bool

const (
	ZonePageShieldPolicyListResponseSuccessTrue ZonePageShieldPolicyListResponseSuccess = true
)

func (r ZonePageShieldPolicyListResponseSuccess) IsKnown() bool {
	switch r {
	case ZonePageShieldPolicyListResponseSuccessTrue:
		return true
	}
	return false
}

type ZonePageShieldPolicyListResponseError struct {
	Code             int64                                        `json:"code,required"`
	Message          string                                       `json:"message,required"`
	DocumentationURL string                                       `json:"documentation_url"`
	Source           ZonePageShieldPolicyListResponseErrorsSource `json:"source"`
	JSON             zonePageShieldPolicyListResponseErrorJSON    `json:"-"`
}

// zonePageShieldPolicyListResponseErrorJSON contains the JSON metadata for the
// struct [ZonePageShieldPolicyListResponseError]
type zonePageShieldPolicyListResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZonePageShieldPolicyListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zonePageShieldPolicyListResponseErrorJSON) RawJSON() string {
	return r.raw
}

type ZonePageShieldPolicyListResponseErrorsSource struct {
	Pointer string                                           `json:"pointer"`
	JSON    zonePageShieldPolicyListResponseErrorsSourceJSON `json:"-"`
}

// zonePageShieldPolicyListResponseErrorsSourceJSON contains the JSON metadata for
// the struct [ZonePageShieldPolicyListResponseErrorsSource]
type zonePageShieldPolicyListResponseErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonePageShieldPolicyListResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zonePageShieldPolicyListResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type ZonePageShieldPolicyListResponseMessage struct {
	Code             int64                                          `json:"code,required"`
	Message          string                                         `json:"message,required"`
	DocumentationURL string                                         `json:"documentation_url"`
	Source           ZonePageShieldPolicyListResponseMessagesSource `json:"source"`
	JSON             zonePageShieldPolicyListResponseMessageJSON    `json:"-"`
}

// zonePageShieldPolicyListResponseMessageJSON contains the JSON metadata for the
// struct [ZonePageShieldPolicyListResponseMessage]
type zonePageShieldPolicyListResponseMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZonePageShieldPolicyListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zonePageShieldPolicyListResponseMessageJSON) RawJSON() string {
	return r.raw
}

type ZonePageShieldPolicyListResponseMessagesSource struct {
	Pointer string                                             `json:"pointer"`
	JSON    zonePageShieldPolicyListResponseMessagesSourceJSON `json:"-"`
}

// zonePageShieldPolicyListResponseMessagesSourceJSON contains the JSON metadata
// for the struct [ZonePageShieldPolicyListResponseMessagesSource]
type zonePageShieldPolicyListResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonePageShieldPolicyListResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zonePageShieldPolicyListResponseMessagesSourceJSON) RawJSON() string {
	return r.raw
}

type ZonePageShieldPolicyNewParams struct {
	Policy PolicyParam `json:"policy,required"`
}

func (r ZonePageShieldPolicyNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Policy)
}

type ZonePageShieldPolicyUpdateParams struct {
	// The action to take if the expression matches
	Action param.Field[PolicyAction] `json:"action"`
	// A description for the policy
	Description param.Field[string] `json:"description"`
	// Whether the policy is enabled
	Enabled param.Field[bool] `json:"enabled"`
	// The expression which must match for the policy to be applied, using the
	// Cloudflare Firewall rule expression syntax
	Expression param.Field[string] `json:"expression"`
	// The policy which will be applied
	Value param.Field[string] `json:"value"`
}

func (r ZonePageShieldPolicyUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
