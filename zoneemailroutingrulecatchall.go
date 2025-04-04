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

// ZoneEmailRoutingRuleCatchAllService contains methods and other services that
// help with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneEmailRoutingRuleCatchAllService] method instead.
type ZoneEmailRoutingRuleCatchAllService struct {
	Options []option.RequestOption
}

// NewZoneEmailRoutingRuleCatchAllService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZoneEmailRoutingRuleCatchAllService(opts ...option.RequestOption) (r *ZoneEmailRoutingRuleCatchAllService) {
	r = &ZoneEmailRoutingRuleCatchAllService{}
	r.Options = opts
	return
}

// Get information on the default catch-all routing rule.
func (r *ZoneEmailRoutingRuleCatchAllService) Get(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *EmailCatchAllRuleResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/email/routing/rules/catch_all", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Enable or disable catch-all routing rule, or change action to forward to
// specific destination address.
func (r *ZoneEmailRoutingRuleCatchAllService) Update(ctx context.Context, zoneID string, body ZoneEmailRoutingRuleCatchAllUpdateParams, opts ...option.RequestOption) (res *EmailCatchAllRuleResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/email/routing/rules/catch_all", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type EmailCatchAllRuleResponseSingle struct {
	Result EmailCatchAllRuleResponseSingleResult `json:"result"`
	JSON   emailCatchAllRuleResponseSingleJSON   `json:"-"`
	EmailAPIResponseSingle
}

// emailCatchAllRuleResponseSingleJSON contains the JSON metadata for the struct
// [EmailCatchAllRuleResponseSingle]
type emailCatchAllRuleResponseSingleJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailCatchAllRuleResponseSingle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailCatchAllRuleResponseSingleJSON) RawJSON() string {
	return r.raw
}

type EmailCatchAllRuleResponseSingleResult struct {
	// Routing rule identifier.
	ID string `json:"id"`
	// List actions for the catch-all routing rule.
	Actions []EmailRuleCatchallAction `json:"actions"`
	// Routing rule status.
	Enabled EmailRuleEnabled `json:"enabled"`
	// List of matchers for the catch-all routing rule.
	Matchers []EmailRuleCatchallMatcher `json:"matchers"`
	// Routing rule name.
	Name string `json:"name"`
	// Routing rule tag. (Deprecated, replaced by routing rule identifier)
	//
	// Deprecated: deprecated
	Tag  string                                    `json:"tag"`
	JSON emailCatchAllRuleResponseSingleResultJSON `json:"-"`
}

// emailCatchAllRuleResponseSingleResultJSON contains the JSON metadata for the
// struct [EmailCatchAllRuleResponseSingleResult]
type emailCatchAllRuleResponseSingleResultJSON struct {
	ID          apijson.Field
	Actions     apijson.Field
	Enabled     apijson.Field
	Matchers    apijson.Field
	Name        apijson.Field
	Tag         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailCatchAllRuleResponseSingleResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailCatchAllRuleResponseSingleResultJSON) RawJSON() string {
	return r.raw
}

// Action for the catch-all routing rule.
type EmailRuleCatchallAction struct {
	// Type of action for catch-all rule.
	Type  EmailRuleCatchallActionType `json:"type,required"`
	Value []string                    `json:"value"`
	JSON  emailRuleCatchallActionJSON `json:"-"`
}

// emailRuleCatchallActionJSON contains the JSON metadata for the struct
// [EmailRuleCatchallAction]
type emailRuleCatchallActionJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRuleCatchallAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRuleCatchallActionJSON) RawJSON() string {
	return r.raw
}

// Type of action for catch-all rule.
type EmailRuleCatchallActionType string

const (
	EmailRuleCatchallActionTypeDrop    EmailRuleCatchallActionType = "drop"
	EmailRuleCatchallActionTypeForward EmailRuleCatchallActionType = "forward"
	EmailRuleCatchallActionTypeWorker  EmailRuleCatchallActionType = "worker"
)

func (r EmailRuleCatchallActionType) IsKnown() bool {
	switch r {
	case EmailRuleCatchallActionTypeDrop, EmailRuleCatchallActionTypeForward, EmailRuleCatchallActionTypeWorker:
		return true
	}
	return false
}

// Action for the catch-all routing rule.
type EmailRuleCatchallActionParam struct {
	// Type of action for catch-all rule.
	Type  param.Field[EmailRuleCatchallActionType] `json:"type,required"`
	Value param.Field[[]string]                    `json:"value"`
}

func (r EmailRuleCatchallActionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Matcher for catch-all routing rule.
type EmailRuleCatchallMatcher struct {
	// Type of matcher. Default is 'all'.
	Type EmailRuleCatchallMatcherType `json:"type,required"`
	JSON emailRuleCatchallMatcherJSON `json:"-"`
}

// emailRuleCatchallMatcherJSON contains the JSON metadata for the struct
// [EmailRuleCatchallMatcher]
type emailRuleCatchallMatcherJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRuleCatchallMatcher) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRuleCatchallMatcherJSON) RawJSON() string {
	return r.raw
}

// Type of matcher. Default is 'all'.
type EmailRuleCatchallMatcherType string

const (
	EmailRuleCatchallMatcherTypeAll EmailRuleCatchallMatcherType = "all"
)

func (r EmailRuleCatchallMatcherType) IsKnown() bool {
	switch r {
	case EmailRuleCatchallMatcherTypeAll:
		return true
	}
	return false
}

// Matcher for catch-all routing rule.
type EmailRuleCatchallMatcherParam struct {
	// Type of matcher. Default is 'all'.
	Type param.Field[EmailRuleCatchallMatcherType] `json:"type,required"`
}

func (r EmailRuleCatchallMatcherParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneEmailRoutingRuleCatchAllUpdateParams struct {
	// List actions for the catch-all routing rule.
	Actions param.Field[[]EmailRuleCatchallActionParam] `json:"actions,required"`
	// List of matchers for the catch-all routing rule.
	Matchers param.Field[[]EmailRuleCatchallMatcherParam] `json:"matchers,required"`
	// Routing rule status.
	Enabled param.Field[EmailRuleEnabled] `json:"enabled"`
	// Routing rule name.
	Name param.Field[string] `json:"name"`
}

func (r ZoneEmailRoutingRuleCatchAllUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
