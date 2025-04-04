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

// AccountZtRiskScoringBehaviorService contains methods and other services that
// help with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountZtRiskScoringBehaviorService] method instead.
type AccountZtRiskScoringBehaviorService struct {
	Options []option.RequestOption
}

// NewAccountZtRiskScoringBehaviorService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountZtRiskScoringBehaviorService(opts ...option.RequestOption) (r *AccountZtRiskScoringBehaviorService) {
	r = &AccountZtRiskScoringBehaviorService{}
	r.Options = opts
	return
}

// Get all behaviors and associated configuration
func (r *AccountZtRiskScoringBehaviorService) ListBehaviors(ctx context.Context, accountID string, opts ...option.RequestOption) (res *AccountZtRiskScoringBehaviorListBehaviorsResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/zt_risk_scoring/behaviors", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update configuration for risk behaviors
func (r *AccountZtRiskScoringBehaviorService) UpdateBehaviors(ctx context.Context, accountID string, body AccountZtRiskScoringBehaviorUpdateBehaviorsParams, opts ...option.RequestOption) (res *AccountZtRiskScoringBehaviorUpdateBehaviorsResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/zt_risk_scoring/behaviors", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type UpdateBehaviors struct {
	Behaviors map[string]UpdateBehaviorsBehavior `json:"behaviors,required"`
	JSON      updateBehaviorsJSON                `json:"-"`
}

// updateBehaviorsJSON contains the JSON metadata for the struct [UpdateBehaviors]
type updateBehaviorsJSON struct {
	Behaviors   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UpdateBehaviors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r updateBehaviorsJSON) RawJSON() string {
	return r.raw
}

type UpdateBehaviorsBehavior struct {
	Enabled   bool                        `json:"enabled,required"`
	RiskLevel RiskLevel                   `json:"risk_level,required"`
	JSON      updateBehaviorsBehaviorJSON `json:"-"`
}

// updateBehaviorsBehaviorJSON contains the JSON metadata for the struct
// [UpdateBehaviorsBehavior]
type updateBehaviorsBehaviorJSON struct {
	Enabled     apijson.Field
	RiskLevel   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UpdateBehaviorsBehavior) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r updateBehaviorsBehaviorJSON) RawJSON() string {
	return r.raw
}

type UpdateBehaviorsParam struct {
	Behaviors param.Field[map[string]UpdateBehaviorsBehaviorParam] `json:"behaviors,required"`
}

func (r UpdateBehaviorsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UpdateBehaviorsBehaviorParam struct {
	Enabled   param.Field[bool]      `json:"enabled,required"`
	RiskLevel param.Field[RiskLevel] `json:"risk_level,required"`
}

func (r UpdateBehaviorsBehaviorParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountZtRiskScoringBehaviorListBehaviorsResponse struct {
	Result AccountZtRiskScoringBehaviorListBehaviorsResponseResult `json:"result"`
	JSON   accountZtRiskScoringBehaviorListBehaviorsResponseJSON   `json:"-"`
	APIResponseSingleDlp
}

// accountZtRiskScoringBehaviorListBehaviorsResponseJSON contains the JSON metadata
// for the struct [AccountZtRiskScoringBehaviorListBehaviorsResponse]
type accountZtRiskScoringBehaviorListBehaviorsResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountZtRiskScoringBehaviorListBehaviorsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountZtRiskScoringBehaviorListBehaviorsResponseJSON) RawJSON() string {
	return r.raw
}

type AccountZtRiskScoringBehaviorListBehaviorsResponseResult struct {
	Behaviors map[string]AccountZtRiskScoringBehaviorListBehaviorsResponseResultBehavior `json:"behaviors,required"`
	JSON      accountZtRiskScoringBehaviorListBehaviorsResponseResultJSON                `json:"-"`
}

// accountZtRiskScoringBehaviorListBehaviorsResponseResultJSON contains the JSON
// metadata for the struct
// [AccountZtRiskScoringBehaviorListBehaviorsResponseResult]
type accountZtRiskScoringBehaviorListBehaviorsResponseResultJSON struct {
	Behaviors   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountZtRiskScoringBehaviorListBehaviorsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountZtRiskScoringBehaviorListBehaviorsResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountZtRiskScoringBehaviorListBehaviorsResponseResultBehavior struct {
	Description string                                                              `json:"description,required"`
	Enabled     bool                                                                `json:"enabled,required"`
	Name        string                                                              `json:"name,required"`
	RiskLevel   RiskLevel                                                           `json:"risk_level,required"`
	JSON        accountZtRiskScoringBehaviorListBehaviorsResponseResultBehaviorJSON `json:"-"`
}

// accountZtRiskScoringBehaviorListBehaviorsResponseResultBehaviorJSON contains the
// JSON metadata for the struct
// [AccountZtRiskScoringBehaviorListBehaviorsResponseResultBehavior]
type accountZtRiskScoringBehaviorListBehaviorsResponseResultBehaviorJSON struct {
	Description apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	RiskLevel   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountZtRiskScoringBehaviorListBehaviorsResponseResultBehavior) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountZtRiskScoringBehaviorListBehaviorsResponseResultBehaviorJSON) RawJSON() string {
	return r.raw
}

type AccountZtRiskScoringBehaviorUpdateBehaviorsResponse struct {
	Result UpdateBehaviors                                         `json:"result"`
	JSON   accountZtRiskScoringBehaviorUpdateBehaviorsResponseJSON `json:"-"`
	APIResponseSingleDlp
}

// accountZtRiskScoringBehaviorUpdateBehaviorsResponseJSON contains the JSON
// metadata for the struct [AccountZtRiskScoringBehaviorUpdateBehaviorsResponse]
type accountZtRiskScoringBehaviorUpdateBehaviorsResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountZtRiskScoringBehaviorUpdateBehaviorsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountZtRiskScoringBehaviorUpdateBehaviorsResponseJSON) RawJSON() string {
	return r.raw
}

type AccountZtRiskScoringBehaviorUpdateBehaviorsParams struct {
	UpdateBehaviors UpdateBehaviorsParam `json:"update_behaviors,required"`
}

func (r AccountZtRiskScoringBehaviorUpdateBehaviorsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.UpdateBehaviors)
}
