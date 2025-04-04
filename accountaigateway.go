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

// AccountAIGatewayService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountAIGatewayService] method instead.
type AccountAIGatewayService struct {
	Options  []option.RequestOption
	Gateways *AccountAIGatewayGatewayService
}

// NewAccountAIGatewayService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountAIGatewayService(opts ...option.RequestOption) (r *AccountAIGatewayService) {
	r = &AccountAIGatewayService{}
	r.Options = opts
	r.Gateways = NewAccountAIGatewayGatewayService(opts...)
	return
}

// List Evaluators
func (r *AccountAIGatewayService) ListEvaluators(ctx context.Context, accountID string, query AccountAIGatewayListEvaluatorsParams, opts ...option.RequestOption) (res *AccountAIGatewayListEvaluatorsResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai-gateway/evaluation-types", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type AccountAIGatewayListEvaluatorsResponse struct {
	Result     []AccountAIGatewayListEvaluatorsResponseResult   `json:"result,required"`
	ResultInfo AccountAIGatewayListEvaluatorsResponseResultInfo `json:"result_info,required"`
	Success    bool                                             `json:"success,required"`
	JSON       accountAIGatewayListEvaluatorsResponseJSON       `json:"-"`
}

// accountAIGatewayListEvaluatorsResponseJSON contains the JSON metadata for the
// struct [AccountAIGatewayListEvaluatorsResponse]
type accountAIGatewayListEvaluatorsResponseJSON struct {
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIGatewayListEvaluatorsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIGatewayListEvaluatorsResponseJSON) RawJSON() string {
	return r.raw
}

type AccountAIGatewayListEvaluatorsResponseResult struct {
	ID          string                                           `json:"id,required"`
	CreatedAt   time.Time                                        `json:"created_at,required" format:"date-time"`
	Description string                                           `json:"description,required"`
	Enable      bool                                             `json:"enable,required"`
	Mandatory   bool                                             `json:"mandatory,required"`
	ModifiedAt  time.Time                                        `json:"modified_at,required" format:"date-time"`
	Name        string                                           `json:"name,required"`
	Type        string                                           `json:"type,required"`
	JSON        accountAIGatewayListEvaluatorsResponseResultJSON `json:"-"`
}

// accountAIGatewayListEvaluatorsResponseResultJSON contains the JSON metadata for
// the struct [AccountAIGatewayListEvaluatorsResponseResult]
type accountAIGatewayListEvaluatorsResponseResultJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Description apijson.Field
	Enable      apijson.Field
	Mandatory   apijson.Field
	ModifiedAt  apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIGatewayListEvaluatorsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIGatewayListEvaluatorsResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountAIGatewayListEvaluatorsResponseResultInfo struct {
	Count      float64                                              `json:"count,required"`
	Page       float64                                              `json:"page,required"`
	PerPage    float64                                              `json:"per_page,required"`
	TotalCount float64                                              `json:"total_count,required"`
	JSON       accountAIGatewayListEvaluatorsResponseResultInfoJSON `json:"-"`
}

// accountAIGatewayListEvaluatorsResponseResultInfoJSON contains the JSON metadata
// for the struct [AccountAIGatewayListEvaluatorsResponseResultInfo]
type accountAIGatewayListEvaluatorsResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIGatewayListEvaluatorsResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIGatewayListEvaluatorsResponseResultInfoJSON) RawJSON() string {
	return r.raw
}

type AccountAIGatewayListEvaluatorsParams struct {
	OrderBy          param.Field[string]                                               `query:"order_by"`
	OrderByDirection param.Field[AccountAIGatewayListEvaluatorsParamsOrderByDirection] `query:"order_by_direction"`
	Page             param.Field[int64]                                                `query:"page"`
	PerPage          param.Field[int64]                                                `query:"per_page"`
}

// URLQuery serializes [AccountAIGatewayListEvaluatorsParams]'s query parameters as
// `url.Values`.
func (r AccountAIGatewayListEvaluatorsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountAIGatewayListEvaluatorsParamsOrderByDirection string

const (
	AccountAIGatewayListEvaluatorsParamsOrderByDirectionAsc  AccountAIGatewayListEvaluatorsParamsOrderByDirection = "asc"
	AccountAIGatewayListEvaluatorsParamsOrderByDirectionDesc AccountAIGatewayListEvaluatorsParamsOrderByDirection = "desc"
)

func (r AccountAIGatewayListEvaluatorsParamsOrderByDirection) IsKnown() bool {
	switch r {
	case AccountAIGatewayListEvaluatorsParamsOrderByDirectionAsc, AccountAIGatewayListEvaluatorsParamsOrderByDirectionDesc:
		return true
	}
	return false
}
