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
	Result PolicyWithID              `json:"result,required"`
	JSON   getZonePolicyResponseJSON `json:"-"`
	GetResponseCollection
}

// getZonePolicyResponseJSON contains the JSON metadata for the struct
// [GetZonePolicyResponse]
type getZonePolicyResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GetZonePolicyResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r getZonePolicyResponseJSON) RawJSON() string {
	return r.raw
}

type Policy struct {
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
	Value string     `json:"value,required"`
	JSON  policyJSON `json:"-"`
}

// policyJSON contains the JSON metadata for the struct [Policy]
type policyJSON struct {
	Action      apijson.Field
	Description apijson.Field
	Enabled     apijson.Field
	Expression  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Policy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r policyJSON) RawJSON() string {
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
	ID   string           `json:"id,required"`
	JSON policyWithIDJSON `json:"-"`
	Policy
}

// policyWithIDJSON contains the JSON metadata for the struct [PolicyWithID]
type policyWithIDJSON struct {
	ID          apijson.Field
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
	Result []PolicyWithID                       `json:"result,required"`
	JSON   zonePageShieldPolicyListResponseJSON `json:"-"`
	ListResponseCollectionPageShield
}

// zonePageShieldPolicyListResponseJSON contains the JSON metadata for the struct
// [ZonePageShieldPolicyListResponse]
type zonePageShieldPolicyListResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonePageShieldPolicyListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zonePageShieldPolicyListResponseJSON) RawJSON() string {
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
