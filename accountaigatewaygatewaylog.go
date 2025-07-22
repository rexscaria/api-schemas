// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/apiquery"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
)

// AccountAIGatewayGatewayLogService contains methods and other services that help
// with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountAIGatewayGatewayLogService] method instead.
type AccountAIGatewayGatewayLogService struct {
	Options []option.RequestOption
}

// NewAccountAIGatewayGatewayLogService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountAIGatewayGatewayLogService(opts ...option.RequestOption) (r *AccountAIGatewayGatewayLogService) {
	r = &AccountAIGatewayGatewayLogService{}
	r.Options = opts
	return
}

// Delete Gateway Logs
func (r *AccountAIGatewayGatewayLogService) DeleteGatewayLogs(ctx context.Context, accountID string, gatewayID string, body AccountAIGatewayGatewayLogDeleteGatewayLogsParams, opts ...option.RequestOption) (res *AccountAIGatewayGatewayLogDeleteGatewayLogsResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if gatewayID == "" {
		err = errors.New("missing required gateway_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai-gateway/gateways/%s/logs", accountID, gatewayID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

// Get Gateway Log Detail
func (r *AccountAIGatewayGatewayLogService) GetGatewayLogDetail(ctx context.Context, accountID string, gatewayID string, id string, opts ...option.RequestOption) (res *AccountAIGatewayGatewayLogGetGatewayLogDetailResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if gatewayID == "" {
		err = errors.New("missing required gateway_id parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai-gateway/gateways/%s/logs/%s", accountID, gatewayID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get Gateway Log Request
func (r *AccountAIGatewayGatewayLogService) GetGatewayLogRequest(ctx context.Context, accountID string, gatewayID string, id string, opts ...option.RequestOption) (res *AccountAIGatewayGatewayLogGetGatewayLogRequestResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if gatewayID == "" {
		err = errors.New("missing required gateway_id parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai-gateway/gateways/%s/logs/%s/request", accountID, gatewayID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get Gateway Log Response
func (r *AccountAIGatewayGatewayLogService) GetGatewayLogResponse(ctx context.Context, accountID string, gatewayID string, id string, opts ...option.RequestOption) (res *AccountAIGatewayGatewayLogGetGatewayLogResponseResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if gatewayID == "" {
		err = errors.New("missing required gateway_id parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai-gateway/gateways/%s/logs/%s/response", accountID, gatewayID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List Gateway Logs
func (r *AccountAIGatewayGatewayLogService) ListGatewayLogs(ctx context.Context, accountID string, gatewayID string, query AccountAIGatewayGatewayLogListGatewayLogsParams, opts ...option.RequestOption) (res *AccountAIGatewayGatewayLogListGatewayLogsResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if gatewayID == "" {
		err = errors.New("missing required gateway_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai-gateway/gateways/%s/logs", accountID, gatewayID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Patch Gateway Log
func (r *AccountAIGatewayGatewayLogService) PatchGatewayLog(ctx context.Context, accountID string, gatewayID string, id string, body AccountAIGatewayGatewayLogPatchGatewayLogParams, opts ...option.RequestOption) (res *AccountAIGatewayGatewayLogPatchGatewayLogResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if gatewayID == "" {
		err = errors.New("missing required gateway_id parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai-gateway/gateways/%s/logs/%s", accountID, gatewayID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

type AccountAIGatewayGatewayLogDeleteGatewayLogsResponse struct {
	Success bool                                                    `json:"success,required"`
	JSON    accountAIGatewayGatewayLogDeleteGatewayLogsResponseJSON `json:"-"`
}

// accountAIGatewayGatewayLogDeleteGatewayLogsResponseJSON contains the JSON
// metadata for the struct [AccountAIGatewayGatewayLogDeleteGatewayLogsResponse]
type accountAIGatewayGatewayLogDeleteGatewayLogsResponseJSON struct {
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIGatewayGatewayLogDeleteGatewayLogsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIGatewayGatewayLogDeleteGatewayLogsResponseJSON) RawJSON() string {
	return r.raw
}

type AccountAIGatewayGatewayLogGetGatewayLogDetailResponse struct {
	Result  AccountAIGatewayGatewayLogGetGatewayLogDetailResponseResult `json:"result,required"`
	Success bool                                                        `json:"success,required"`
	JSON    accountAIGatewayGatewayLogGetGatewayLogDetailResponseJSON   `json:"-"`
}

// accountAIGatewayGatewayLogGetGatewayLogDetailResponseJSON contains the JSON
// metadata for the struct [AccountAIGatewayGatewayLogGetGatewayLogDetailResponse]
type accountAIGatewayGatewayLogGetGatewayLogDetailResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIGatewayGatewayLogGetGatewayLogDetailResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIGatewayGatewayLogGetGatewayLogDetailResponseJSON) RawJSON() string {
	return r.raw
}

type AccountAIGatewayGatewayLogGetGatewayLogDetailResponseResult struct {
	ID                   string                                                          `json:"id,required"`
	Cached               bool                                                            `json:"cached,required"`
	CreatedAt            time.Time                                                       `json:"created_at,required" format:"date-time"`
	Duration             int64                                                           `json:"duration,required"`
	Model                string                                                          `json:"model,required"`
	Path                 string                                                          `json:"path,required"`
	Provider             string                                                          `json:"provider,required"`
	Success              bool                                                            `json:"success,required"`
	TokensIn             int64                                                           `json:"tokens_in,required,nullable"`
	TokensOut            int64                                                           `json:"tokens_out,required,nullable"`
	Cost                 float64                                                         `json:"cost"`
	CustomCost           bool                                                            `json:"custom_cost"`
	Metadata             string                                                          `json:"metadata"`
	ModelType            string                                                          `json:"model_type"`
	RequestContentType   string                                                          `json:"request_content_type"`
	RequestHead          string                                                          `json:"request_head"`
	RequestHeadComplete  bool                                                            `json:"request_head_complete"`
	RequestSize          int64                                                           `json:"request_size"`
	RequestType          string                                                          `json:"request_type"`
	ResponseContentType  string                                                          `json:"response_content_type"`
	ResponseHead         string                                                          `json:"response_head"`
	ResponseHeadComplete bool                                                            `json:"response_head_complete"`
	ResponseSize         int64                                                           `json:"response_size"`
	StatusCode           int64                                                           `json:"status_code"`
	Step                 int64                                                           `json:"step"`
	JSON                 accountAIGatewayGatewayLogGetGatewayLogDetailResponseResultJSON `json:"-"`
}

// accountAIGatewayGatewayLogGetGatewayLogDetailResponseResultJSON contains the
// JSON metadata for the struct
// [AccountAIGatewayGatewayLogGetGatewayLogDetailResponseResult]
type accountAIGatewayGatewayLogGetGatewayLogDetailResponseResultJSON struct {
	ID                   apijson.Field
	Cached               apijson.Field
	CreatedAt            apijson.Field
	Duration             apijson.Field
	Model                apijson.Field
	Path                 apijson.Field
	Provider             apijson.Field
	Success              apijson.Field
	TokensIn             apijson.Field
	TokensOut            apijson.Field
	Cost                 apijson.Field
	CustomCost           apijson.Field
	Metadata             apijson.Field
	ModelType            apijson.Field
	RequestContentType   apijson.Field
	RequestHead          apijson.Field
	RequestHeadComplete  apijson.Field
	RequestSize          apijson.Field
	RequestType          apijson.Field
	ResponseContentType  apijson.Field
	ResponseHead         apijson.Field
	ResponseHeadComplete apijson.Field
	ResponseSize         apijson.Field
	StatusCode           apijson.Field
	Step                 apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccountAIGatewayGatewayLogGetGatewayLogDetailResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIGatewayGatewayLogGetGatewayLogDetailResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountAIGatewayGatewayLogGetGatewayLogRequestResponse = interface{}

type AccountAIGatewayGatewayLogGetGatewayLogResponseResponse = interface{}

type AccountAIGatewayGatewayLogListGatewayLogsResponse struct {
	Result     []AccountAIGatewayGatewayLogListGatewayLogsResponseResult   `json:"result,required"`
	ResultInfo AccountAIGatewayGatewayLogListGatewayLogsResponseResultInfo `json:"result_info,required"`
	Success    bool                                                        `json:"success,required"`
	JSON       accountAIGatewayGatewayLogListGatewayLogsResponseJSON       `json:"-"`
}

// accountAIGatewayGatewayLogListGatewayLogsResponseJSON contains the JSON metadata
// for the struct [AccountAIGatewayGatewayLogListGatewayLogsResponse]
type accountAIGatewayGatewayLogListGatewayLogsResponseJSON struct {
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIGatewayGatewayLogListGatewayLogsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIGatewayGatewayLogListGatewayLogsResponseJSON) RawJSON() string {
	return r.raw
}

type AccountAIGatewayGatewayLogListGatewayLogsResponseResult struct {
	ID                  string                                                      `json:"id,required"`
	Cached              bool                                                        `json:"cached,required"`
	CreatedAt           time.Time                                                   `json:"created_at,required" format:"date-time"`
	Duration            int64                                                       `json:"duration,required"`
	Model               string                                                      `json:"model,required"`
	Path                string                                                      `json:"path,required"`
	Provider            string                                                      `json:"provider,required"`
	Success             bool                                                        `json:"success,required"`
	TokensIn            int64                                                       `json:"tokens_in,required,nullable"`
	TokensOut           int64                                                       `json:"tokens_out,required,nullable"`
	Cost                float64                                                     `json:"cost"`
	CustomCost          bool                                                        `json:"custom_cost"`
	Metadata            string                                                      `json:"metadata"`
	ModelType           string                                                      `json:"model_type"`
	RequestContentType  string                                                      `json:"request_content_type"`
	RequestType         string                                                      `json:"request_type"`
	ResponseContentType string                                                      `json:"response_content_type"`
	StatusCode          int64                                                       `json:"status_code"`
	Step                int64                                                       `json:"step"`
	JSON                accountAIGatewayGatewayLogListGatewayLogsResponseResultJSON `json:"-"`
}

// accountAIGatewayGatewayLogListGatewayLogsResponseResultJSON contains the JSON
// metadata for the struct
// [AccountAIGatewayGatewayLogListGatewayLogsResponseResult]
type accountAIGatewayGatewayLogListGatewayLogsResponseResultJSON struct {
	ID                  apijson.Field
	Cached              apijson.Field
	CreatedAt           apijson.Field
	Duration            apijson.Field
	Model               apijson.Field
	Path                apijson.Field
	Provider            apijson.Field
	Success             apijson.Field
	TokensIn            apijson.Field
	TokensOut           apijson.Field
	Cost                apijson.Field
	CustomCost          apijson.Field
	Metadata            apijson.Field
	ModelType           apijson.Field
	RequestContentType  apijson.Field
	RequestType         apijson.Field
	ResponseContentType apijson.Field
	StatusCode          apijson.Field
	Step                apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *AccountAIGatewayGatewayLogListGatewayLogsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIGatewayGatewayLogListGatewayLogsResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountAIGatewayGatewayLogListGatewayLogsResponseResultInfo struct {
	Count          float64                                                         `json:"count"`
	MaxCost        float64                                                         `json:"max_cost"`
	MaxDuration    float64                                                         `json:"max_duration"`
	MaxTokensIn    float64                                                         `json:"max_tokens_in"`
	MaxTokensOut   float64                                                         `json:"max_tokens_out"`
	MaxTotalTokens float64                                                         `json:"max_total_tokens"`
	MinCost        float64                                                         `json:"min_cost"`
	MinDuration    float64                                                         `json:"min_duration"`
	MinTokensIn    float64                                                         `json:"min_tokens_in"`
	MinTokensOut   float64                                                         `json:"min_tokens_out"`
	MinTotalTokens float64                                                         `json:"min_total_tokens"`
	Page           float64                                                         `json:"page"`
	PerPage        float64                                                         `json:"per_page"`
	TotalCount     float64                                                         `json:"total_count"`
	JSON           accountAIGatewayGatewayLogListGatewayLogsResponseResultInfoJSON `json:"-"`
}

// accountAIGatewayGatewayLogListGatewayLogsResponseResultInfoJSON contains the
// JSON metadata for the struct
// [AccountAIGatewayGatewayLogListGatewayLogsResponseResultInfo]
type accountAIGatewayGatewayLogListGatewayLogsResponseResultInfoJSON struct {
	Count          apijson.Field
	MaxCost        apijson.Field
	MaxDuration    apijson.Field
	MaxTokensIn    apijson.Field
	MaxTokensOut   apijson.Field
	MaxTotalTokens apijson.Field
	MinCost        apijson.Field
	MinDuration    apijson.Field
	MinTokensIn    apijson.Field
	MinTokensOut   apijson.Field
	MinTotalTokens apijson.Field
	Page           apijson.Field
	PerPage        apijson.Field
	TotalCount     apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccountAIGatewayGatewayLogListGatewayLogsResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIGatewayGatewayLogListGatewayLogsResponseResultInfoJSON) RawJSON() string {
	return r.raw
}

type AccountAIGatewayGatewayLogPatchGatewayLogResponse struct {
	Result  interface{}                                           `json:"result,required"`
	Success bool                                                  `json:"success,required"`
	JSON    accountAIGatewayGatewayLogPatchGatewayLogResponseJSON `json:"-"`
}

// accountAIGatewayGatewayLogPatchGatewayLogResponseJSON contains the JSON metadata
// for the struct [AccountAIGatewayGatewayLogPatchGatewayLogResponse]
type accountAIGatewayGatewayLogPatchGatewayLogResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIGatewayGatewayLogPatchGatewayLogResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIGatewayGatewayLogPatchGatewayLogResponseJSON) RawJSON() string {
	return r.raw
}

type AccountAIGatewayGatewayLogDeleteGatewayLogsParams struct {
	Filters          param.Field[[]AccountAIGatewayGatewayLogDeleteGatewayLogsParamsFilter]         `query:"filters"`
	Limit            param.Field[int64]                                                             `query:"limit"`
	OrderBy          param.Field[AccountAIGatewayGatewayLogDeleteGatewayLogsParamsOrderBy]          `query:"order_by"`
	OrderByDirection param.Field[AccountAIGatewayGatewayLogDeleteGatewayLogsParamsOrderByDirection] `query:"order_by_direction"`
}

// URLQuery serializes [AccountAIGatewayGatewayLogDeleteGatewayLogsParams]'s query
// parameters as `url.Values`.
func (r AccountAIGatewayGatewayLogDeleteGatewayLogsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountAIGatewayGatewayLogDeleteGatewayLogsParamsFilter struct {
	Key      param.Field[AccountAIGatewayGatewayLogDeleteGatewayLogsParamsFiltersKey]          `query:"key,required"`
	Operator param.Field[AccountAIGatewayGatewayLogDeleteGatewayLogsParamsFiltersOperator]     `query:"operator,required"`
	Value    param.Field[[]AccountAIGatewayGatewayLogDeleteGatewayLogsParamsFiltersValueUnion] `query:"value,required"`
}

// URLQuery serializes [AccountAIGatewayGatewayLogDeleteGatewayLogsParamsFilter]'s
// query parameters as `url.Values`.
func (r AccountAIGatewayGatewayLogDeleteGatewayLogsParamsFilter) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountAIGatewayGatewayLogDeleteGatewayLogsParamsFiltersKey string

const (
	AccountAIGatewayGatewayLogDeleteGatewayLogsParamsFiltersKeyID                  AccountAIGatewayGatewayLogDeleteGatewayLogsParamsFiltersKey = "id"
	AccountAIGatewayGatewayLogDeleteGatewayLogsParamsFiltersKeyCreatedAt           AccountAIGatewayGatewayLogDeleteGatewayLogsParamsFiltersKey = "created_at"
	AccountAIGatewayGatewayLogDeleteGatewayLogsParamsFiltersKeyRequestContentType  AccountAIGatewayGatewayLogDeleteGatewayLogsParamsFiltersKey = "request_content_type"
	AccountAIGatewayGatewayLogDeleteGatewayLogsParamsFiltersKeyResponseContentType AccountAIGatewayGatewayLogDeleteGatewayLogsParamsFiltersKey = "response_content_type"
	AccountAIGatewayGatewayLogDeleteGatewayLogsParamsFiltersKeySuccess             AccountAIGatewayGatewayLogDeleteGatewayLogsParamsFiltersKey = "success"
	AccountAIGatewayGatewayLogDeleteGatewayLogsParamsFiltersKeyCached              AccountAIGatewayGatewayLogDeleteGatewayLogsParamsFiltersKey = "cached"
	AccountAIGatewayGatewayLogDeleteGatewayLogsParamsFiltersKeyProvider            AccountAIGatewayGatewayLogDeleteGatewayLogsParamsFiltersKey = "provider"
	AccountAIGatewayGatewayLogDeleteGatewayLogsParamsFiltersKeyModel               AccountAIGatewayGatewayLogDeleteGatewayLogsParamsFiltersKey = "model"
	AccountAIGatewayGatewayLogDeleteGatewayLogsParamsFiltersKeyModelType           AccountAIGatewayGatewayLogDeleteGatewayLogsParamsFiltersKey = "model_type"
	AccountAIGatewayGatewayLogDeleteGatewayLogsParamsFiltersKeyCost                AccountAIGatewayGatewayLogDeleteGatewayLogsParamsFiltersKey = "cost"
	AccountAIGatewayGatewayLogDeleteGatewayLogsParamsFiltersKeyTokens              AccountAIGatewayGatewayLogDeleteGatewayLogsParamsFiltersKey = "tokens"
	AccountAIGatewayGatewayLogDeleteGatewayLogsParamsFiltersKeyTokensIn            AccountAIGatewayGatewayLogDeleteGatewayLogsParamsFiltersKey = "tokens_in"
	AccountAIGatewayGatewayLogDeleteGatewayLogsParamsFiltersKeyTokensOut           AccountAIGatewayGatewayLogDeleteGatewayLogsParamsFiltersKey = "tokens_out"
	AccountAIGatewayGatewayLogDeleteGatewayLogsParamsFiltersKeyDuration            AccountAIGatewayGatewayLogDeleteGatewayLogsParamsFiltersKey = "duration"
	AccountAIGatewayGatewayLogDeleteGatewayLogsParamsFiltersKeyFeedback            AccountAIGatewayGatewayLogDeleteGatewayLogsParamsFiltersKey = "feedback"
	AccountAIGatewayGatewayLogDeleteGatewayLogsParamsFiltersKeyEventID             AccountAIGatewayGatewayLogDeleteGatewayLogsParamsFiltersKey = "event_id"
	AccountAIGatewayGatewayLogDeleteGatewayLogsParamsFiltersKeyRequestType         AccountAIGatewayGatewayLogDeleteGatewayLogsParamsFiltersKey = "request_type"
	AccountAIGatewayGatewayLogDeleteGatewayLogsParamsFiltersKeyMetadataKey         AccountAIGatewayGatewayLogDeleteGatewayLogsParamsFiltersKey = "metadata.key"
	AccountAIGatewayGatewayLogDeleteGatewayLogsParamsFiltersKeyMetadataValue       AccountAIGatewayGatewayLogDeleteGatewayLogsParamsFiltersKey = "metadata.value"
	AccountAIGatewayGatewayLogDeleteGatewayLogsParamsFiltersKeyPromptsPromptID     AccountAIGatewayGatewayLogDeleteGatewayLogsParamsFiltersKey = "prompts.prompt_id"
	AccountAIGatewayGatewayLogDeleteGatewayLogsParamsFiltersKeyPromptsVersionID    AccountAIGatewayGatewayLogDeleteGatewayLogsParamsFiltersKey = "prompts.version_id"
)

func (r AccountAIGatewayGatewayLogDeleteGatewayLogsParamsFiltersKey) IsKnown() bool {
	switch r {
	case AccountAIGatewayGatewayLogDeleteGatewayLogsParamsFiltersKeyID, AccountAIGatewayGatewayLogDeleteGatewayLogsParamsFiltersKeyCreatedAt, AccountAIGatewayGatewayLogDeleteGatewayLogsParamsFiltersKeyRequestContentType, AccountAIGatewayGatewayLogDeleteGatewayLogsParamsFiltersKeyResponseContentType, AccountAIGatewayGatewayLogDeleteGatewayLogsParamsFiltersKeySuccess, AccountAIGatewayGatewayLogDeleteGatewayLogsParamsFiltersKeyCached, AccountAIGatewayGatewayLogDeleteGatewayLogsParamsFiltersKeyProvider, AccountAIGatewayGatewayLogDeleteGatewayLogsParamsFiltersKeyModel, AccountAIGatewayGatewayLogDeleteGatewayLogsParamsFiltersKeyModelType, AccountAIGatewayGatewayLogDeleteGatewayLogsParamsFiltersKeyCost, AccountAIGatewayGatewayLogDeleteGatewayLogsParamsFiltersKeyTokens, AccountAIGatewayGatewayLogDeleteGatewayLogsParamsFiltersKeyTokensIn, AccountAIGatewayGatewayLogDeleteGatewayLogsParamsFiltersKeyTokensOut, AccountAIGatewayGatewayLogDeleteGatewayLogsParamsFiltersKeyDuration, AccountAIGatewayGatewayLogDeleteGatewayLogsParamsFiltersKeyFeedback, AccountAIGatewayGatewayLogDeleteGatewayLogsParamsFiltersKeyEventID, AccountAIGatewayGatewayLogDeleteGatewayLogsParamsFiltersKeyRequestType, AccountAIGatewayGatewayLogDeleteGatewayLogsParamsFiltersKeyMetadataKey, AccountAIGatewayGatewayLogDeleteGatewayLogsParamsFiltersKeyMetadataValue, AccountAIGatewayGatewayLogDeleteGatewayLogsParamsFiltersKeyPromptsPromptID, AccountAIGatewayGatewayLogDeleteGatewayLogsParamsFiltersKeyPromptsVersionID:
		return true
	}
	return false
}

type AccountAIGatewayGatewayLogDeleteGatewayLogsParamsFiltersOperator string

const (
	AccountAIGatewayGatewayLogDeleteGatewayLogsParamsFiltersOperatorEq       AccountAIGatewayGatewayLogDeleteGatewayLogsParamsFiltersOperator = "eq"
	AccountAIGatewayGatewayLogDeleteGatewayLogsParamsFiltersOperatorNeq      AccountAIGatewayGatewayLogDeleteGatewayLogsParamsFiltersOperator = "neq"
	AccountAIGatewayGatewayLogDeleteGatewayLogsParamsFiltersOperatorContains AccountAIGatewayGatewayLogDeleteGatewayLogsParamsFiltersOperator = "contains"
	AccountAIGatewayGatewayLogDeleteGatewayLogsParamsFiltersOperatorLt       AccountAIGatewayGatewayLogDeleteGatewayLogsParamsFiltersOperator = "lt"
	AccountAIGatewayGatewayLogDeleteGatewayLogsParamsFiltersOperatorGt       AccountAIGatewayGatewayLogDeleteGatewayLogsParamsFiltersOperator = "gt"
)

func (r AccountAIGatewayGatewayLogDeleteGatewayLogsParamsFiltersOperator) IsKnown() bool {
	switch r {
	case AccountAIGatewayGatewayLogDeleteGatewayLogsParamsFiltersOperatorEq, AccountAIGatewayGatewayLogDeleteGatewayLogsParamsFiltersOperatorNeq, AccountAIGatewayGatewayLogDeleteGatewayLogsParamsFiltersOperatorContains, AccountAIGatewayGatewayLogDeleteGatewayLogsParamsFiltersOperatorLt, AccountAIGatewayGatewayLogDeleteGatewayLogsParamsFiltersOperatorGt:
		return true
	}
	return false
}

// Satisfied by [shared.UnionString], [shared.UnionFloat], [shared.UnionBool].
type AccountAIGatewayGatewayLogDeleteGatewayLogsParamsFiltersValueUnion interface {
	ImplementsAccountAIGatewayGatewayLogDeleteGatewayLogsParamsFiltersValueUnion()
}

type AccountAIGatewayGatewayLogDeleteGatewayLogsParamsOrderBy string

const (
	AccountAIGatewayGatewayLogDeleteGatewayLogsParamsOrderByCreatedAt AccountAIGatewayGatewayLogDeleteGatewayLogsParamsOrderBy = "created_at"
	AccountAIGatewayGatewayLogDeleteGatewayLogsParamsOrderByProvider  AccountAIGatewayGatewayLogDeleteGatewayLogsParamsOrderBy = "provider"
	AccountAIGatewayGatewayLogDeleteGatewayLogsParamsOrderByModel     AccountAIGatewayGatewayLogDeleteGatewayLogsParamsOrderBy = "model"
	AccountAIGatewayGatewayLogDeleteGatewayLogsParamsOrderByModelType AccountAIGatewayGatewayLogDeleteGatewayLogsParamsOrderBy = "model_type"
	AccountAIGatewayGatewayLogDeleteGatewayLogsParamsOrderBySuccess   AccountAIGatewayGatewayLogDeleteGatewayLogsParamsOrderBy = "success"
	AccountAIGatewayGatewayLogDeleteGatewayLogsParamsOrderByCached    AccountAIGatewayGatewayLogDeleteGatewayLogsParamsOrderBy = "cached"
	AccountAIGatewayGatewayLogDeleteGatewayLogsParamsOrderByCost      AccountAIGatewayGatewayLogDeleteGatewayLogsParamsOrderBy = "cost"
	AccountAIGatewayGatewayLogDeleteGatewayLogsParamsOrderByTokensIn  AccountAIGatewayGatewayLogDeleteGatewayLogsParamsOrderBy = "tokens_in"
	AccountAIGatewayGatewayLogDeleteGatewayLogsParamsOrderByTokensOut AccountAIGatewayGatewayLogDeleteGatewayLogsParamsOrderBy = "tokens_out"
	AccountAIGatewayGatewayLogDeleteGatewayLogsParamsOrderByDuration  AccountAIGatewayGatewayLogDeleteGatewayLogsParamsOrderBy = "duration"
	AccountAIGatewayGatewayLogDeleteGatewayLogsParamsOrderByFeedback  AccountAIGatewayGatewayLogDeleteGatewayLogsParamsOrderBy = "feedback"
)

func (r AccountAIGatewayGatewayLogDeleteGatewayLogsParamsOrderBy) IsKnown() bool {
	switch r {
	case AccountAIGatewayGatewayLogDeleteGatewayLogsParamsOrderByCreatedAt, AccountAIGatewayGatewayLogDeleteGatewayLogsParamsOrderByProvider, AccountAIGatewayGatewayLogDeleteGatewayLogsParamsOrderByModel, AccountAIGatewayGatewayLogDeleteGatewayLogsParamsOrderByModelType, AccountAIGatewayGatewayLogDeleteGatewayLogsParamsOrderBySuccess, AccountAIGatewayGatewayLogDeleteGatewayLogsParamsOrderByCached, AccountAIGatewayGatewayLogDeleteGatewayLogsParamsOrderByCost, AccountAIGatewayGatewayLogDeleteGatewayLogsParamsOrderByTokensIn, AccountAIGatewayGatewayLogDeleteGatewayLogsParamsOrderByTokensOut, AccountAIGatewayGatewayLogDeleteGatewayLogsParamsOrderByDuration, AccountAIGatewayGatewayLogDeleteGatewayLogsParamsOrderByFeedback:
		return true
	}
	return false
}

type AccountAIGatewayGatewayLogDeleteGatewayLogsParamsOrderByDirection string

const (
	AccountAIGatewayGatewayLogDeleteGatewayLogsParamsOrderByDirectionAsc  AccountAIGatewayGatewayLogDeleteGatewayLogsParamsOrderByDirection = "asc"
	AccountAIGatewayGatewayLogDeleteGatewayLogsParamsOrderByDirectionDesc AccountAIGatewayGatewayLogDeleteGatewayLogsParamsOrderByDirection = "desc"
)

func (r AccountAIGatewayGatewayLogDeleteGatewayLogsParamsOrderByDirection) IsKnown() bool {
	switch r {
	case AccountAIGatewayGatewayLogDeleteGatewayLogsParamsOrderByDirectionAsc, AccountAIGatewayGatewayLogDeleteGatewayLogsParamsOrderByDirectionDesc:
		return true
	}
	return false
}

type AccountAIGatewayGatewayLogListGatewayLogsParams struct {
	Cached              param.Field[bool]                                                            `query:"cached"`
	Direction           param.Field[AccountAIGatewayGatewayLogListGatewayLogsParamsDirection]        `query:"direction"`
	EndDate             param.Field[time.Time]                                                       `query:"end_date" format:"date-time"`
	Feedback            param.Field[AccountAIGatewayGatewayLogListGatewayLogsParamsFeedback]         `query:"feedback"`
	Filters             param.Field[[]AccountAIGatewayGatewayLogListGatewayLogsParamsFilter]         `query:"filters"`
	MaxCost             param.Field[float64]                                                         `query:"max_cost"`
	MaxDuration         param.Field[float64]                                                         `query:"max_duration"`
	MaxTokensIn         param.Field[float64]                                                         `query:"max_tokens_in"`
	MaxTokensOut        param.Field[float64]                                                         `query:"max_tokens_out"`
	MaxTotalTokens      param.Field[float64]                                                         `query:"max_total_tokens"`
	MetaInfo            param.Field[bool]                                                            `query:"meta_info"`
	MinCost             param.Field[float64]                                                         `query:"min_cost"`
	MinDuration         param.Field[float64]                                                         `query:"min_duration"`
	MinTokensIn         param.Field[float64]                                                         `query:"min_tokens_in"`
	MinTokensOut        param.Field[float64]                                                         `query:"min_tokens_out"`
	MinTotalTokens      param.Field[float64]                                                         `query:"min_total_tokens"`
	Model               param.Field[string]                                                          `query:"model"`
	ModelType           param.Field[string]                                                          `query:"model_type"`
	OrderBy             param.Field[AccountAIGatewayGatewayLogListGatewayLogsParamsOrderBy]          `query:"order_by"`
	OrderByDirection    param.Field[AccountAIGatewayGatewayLogListGatewayLogsParamsOrderByDirection] `query:"order_by_direction"`
	Page                param.Field[int64]                                                           `query:"page"`
	PerPage             param.Field[int64]                                                           `query:"per_page"`
	Provider            param.Field[string]                                                          `query:"provider"`
	RequestContentType  param.Field[string]                                                          `query:"request_content_type"`
	ResponseContentType param.Field[string]                                                          `query:"response_content_type"`
	Search              param.Field[string]                                                          `query:"search"`
	StartDate           param.Field[time.Time]                                                       `query:"start_date" format:"date-time"`
	Success             param.Field[bool]                                                            `query:"success"`
}

// URLQuery serializes [AccountAIGatewayGatewayLogListGatewayLogsParams]'s query
// parameters as `url.Values`.
func (r AccountAIGatewayGatewayLogListGatewayLogsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountAIGatewayGatewayLogListGatewayLogsParamsDirection string

const (
	AccountAIGatewayGatewayLogListGatewayLogsParamsDirectionAsc  AccountAIGatewayGatewayLogListGatewayLogsParamsDirection = "asc"
	AccountAIGatewayGatewayLogListGatewayLogsParamsDirectionDesc AccountAIGatewayGatewayLogListGatewayLogsParamsDirection = "desc"
)

func (r AccountAIGatewayGatewayLogListGatewayLogsParamsDirection) IsKnown() bool {
	switch r {
	case AccountAIGatewayGatewayLogListGatewayLogsParamsDirectionAsc, AccountAIGatewayGatewayLogListGatewayLogsParamsDirectionDesc:
		return true
	}
	return false
}

type AccountAIGatewayGatewayLogListGatewayLogsParamsFeedback float64

const (
	AccountAIGatewayGatewayLogListGatewayLogsParamsFeedbackMinus1 AccountAIGatewayGatewayLogListGatewayLogsParamsFeedback = -1
	AccountAIGatewayGatewayLogListGatewayLogsParamsFeedback0      AccountAIGatewayGatewayLogListGatewayLogsParamsFeedback = 0
	AccountAIGatewayGatewayLogListGatewayLogsParamsFeedback1      AccountAIGatewayGatewayLogListGatewayLogsParamsFeedback = 1
)

func (r AccountAIGatewayGatewayLogListGatewayLogsParamsFeedback) IsKnown() bool {
	switch r {
	case AccountAIGatewayGatewayLogListGatewayLogsParamsFeedbackMinus1, AccountAIGatewayGatewayLogListGatewayLogsParamsFeedback0, AccountAIGatewayGatewayLogListGatewayLogsParamsFeedback1:
		return true
	}
	return false
}

type AccountAIGatewayGatewayLogListGatewayLogsParamsFilter struct {
	Key      param.Field[AccountAIGatewayGatewayLogListGatewayLogsParamsFiltersKey]          `query:"key,required"`
	Operator param.Field[AccountAIGatewayGatewayLogListGatewayLogsParamsFiltersOperator]     `query:"operator,required"`
	Value    param.Field[[]AccountAIGatewayGatewayLogListGatewayLogsParamsFiltersValueUnion] `query:"value,required"`
}

// URLQuery serializes [AccountAIGatewayGatewayLogListGatewayLogsParamsFilter]'s
// query parameters as `url.Values`.
func (r AccountAIGatewayGatewayLogListGatewayLogsParamsFilter) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountAIGatewayGatewayLogListGatewayLogsParamsFiltersKey string

const (
	AccountAIGatewayGatewayLogListGatewayLogsParamsFiltersKeyID                  AccountAIGatewayGatewayLogListGatewayLogsParamsFiltersKey = "id"
	AccountAIGatewayGatewayLogListGatewayLogsParamsFiltersKeyCreatedAt           AccountAIGatewayGatewayLogListGatewayLogsParamsFiltersKey = "created_at"
	AccountAIGatewayGatewayLogListGatewayLogsParamsFiltersKeyRequestContentType  AccountAIGatewayGatewayLogListGatewayLogsParamsFiltersKey = "request_content_type"
	AccountAIGatewayGatewayLogListGatewayLogsParamsFiltersKeyResponseContentType AccountAIGatewayGatewayLogListGatewayLogsParamsFiltersKey = "response_content_type"
	AccountAIGatewayGatewayLogListGatewayLogsParamsFiltersKeySuccess             AccountAIGatewayGatewayLogListGatewayLogsParamsFiltersKey = "success"
	AccountAIGatewayGatewayLogListGatewayLogsParamsFiltersKeyCached              AccountAIGatewayGatewayLogListGatewayLogsParamsFiltersKey = "cached"
	AccountAIGatewayGatewayLogListGatewayLogsParamsFiltersKeyProvider            AccountAIGatewayGatewayLogListGatewayLogsParamsFiltersKey = "provider"
	AccountAIGatewayGatewayLogListGatewayLogsParamsFiltersKeyModel               AccountAIGatewayGatewayLogListGatewayLogsParamsFiltersKey = "model"
	AccountAIGatewayGatewayLogListGatewayLogsParamsFiltersKeyModelType           AccountAIGatewayGatewayLogListGatewayLogsParamsFiltersKey = "model_type"
	AccountAIGatewayGatewayLogListGatewayLogsParamsFiltersKeyCost                AccountAIGatewayGatewayLogListGatewayLogsParamsFiltersKey = "cost"
	AccountAIGatewayGatewayLogListGatewayLogsParamsFiltersKeyTokens              AccountAIGatewayGatewayLogListGatewayLogsParamsFiltersKey = "tokens"
	AccountAIGatewayGatewayLogListGatewayLogsParamsFiltersKeyTokensIn            AccountAIGatewayGatewayLogListGatewayLogsParamsFiltersKey = "tokens_in"
	AccountAIGatewayGatewayLogListGatewayLogsParamsFiltersKeyTokensOut           AccountAIGatewayGatewayLogListGatewayLogsParamsFiltersKey = "tokens_out"
	AccountAIGatewayGatewayLogListGatewayLogsParamsFiltersKeyDuration            AccountAIGatewayGatewayLogListGatewayLogsParamsFiltersKey = "duration"
	AccountAIGatewayGatewayLogListGatewayLogsParamsFiltersKeyFeedback            AccountAIGatewayGatewayLogListGatewayLogsParamsFiltersKey = "feedback"
	AccountAIGatewayGatewayLogListGatewayLogsParamsFiltersKeyEventID             AccountAIGatewayGatewayLogListGatewayLogsParamsFiltersKey = "event_id"
	AccountAIGatewayGatewayLogListGatewayLogsParamsFiltersKeyRequestType         AccountAIGatewayGatewayLogListGatewayLogsParamsFiltersKey = "request_type"
	AccountAIGatewayGatewayLogListGatewayLogsParamsFiltersKeyMetadataKey         AccountAIGatewayGatewayLogListGatewayLogsParamsFiltersKey = "metadata.key"
	AccountAIGatewayGatewayLogListGatewayLogsParamsFiltersKeyMetadataValue       AccountAIGatewayGatewayLogListGatewayLogsParamsFiltersKey = "metadata.value"
	AccountAIGatewayGatewayLogListGatewayLogsParamsFiltersKeyPromptsPromptID     AccountAIGatewayGatewayLogListGatewayLogsParamsFiltersKey = "prompts.prompt_id"
	AccountAIGatewayGatewayLogListGatewayLogsParamsFiltersKeyPromptsVersionID    AccountAIGatewayGatewayLogListGatewayLogsParamsFiltersKey = "prompts.version_id"
)

func (r AccountAIGatewayGatewayLogListGatewayLogsParamsFiltersKey) IsKnown() bool {
	switch r {
	case AccountAIGatewayGatewayLogListGatewayLogsParamsFiltersKeyID, AccountAIGatewayGatewayLogListGatewayLogsParamsFiltersKeyCreatedAt, AccountAIGatewayGatewayLogListGatewayLogsParamsFiltersKeyRequestContentType, AccountAIGatewayGatewayLogListGatewayLogsParamsFiltersKeyResponseContentType, AccountAIGatewayGatewayLogListGatewayLogsParamsFiltersKeySuccess, AccountAIGatewayGatewayLogListGatewayLogsParamsFiltersKeyCached, AccountAIGatewayGatewayLogListGatewayLogsParamsFiltersKeyProvider, AccountAIGatewayGatewayLogListGatewayLogsParamsFiltersKeyModel, AccountAIGatewayGatewayLogListGatewayLogsParamsFiltersKeyModelType, AccountAIGatewayGatewayLogListGatewayLogsParamsFiltersKeyCost, AccountAIGatewayGatewayLogListGatewayLogsParamsFiltersKeyTokens, AccountAIGatewayGatewayLogListGatewayLogsParamsFiltersKeyTokensIn, AccountAIGatewayGatewayLogListGatewayLogsParamsFiltersKeyTokensOut, AccountAIGatewayGatewayLogListGatewayLogsParamsFiltersKeyDuration, AccountAIGatewayGatewayLogListGatewayLogsParamsFiltersKeyFeedback, AccountAIGatewayGatewayLogListGatewayLogsParamsFiltersKeyEventID, AccountAIGatewayGatewayLogListGatewayLogsParamsFiltersKeyRequestType, AccountAIGatewayGatewayLogListGatewayLogsParamsFiltersKeyMetadataKey, AccountAIGatewayGatewayLogListGatewayLogsParamsFiltersKeyMetadataValue, AccountAIGatewayGatewayLogListGatewayLogsParamsFiltersKeyPromptsPromptID, AccountAIGatewayGatewayLogListGatewayLogsParamsFiltersKeyPromptsVersionID:
		return true
	}
	return false
}

type AccountAIGatewayGatewayLogListGatewayLogsParamsFiltersOperator string

const (
	AccountAIGatewayGatewayLogListGatewayLogsParamsFiltersOperatorEq       AccountAIGatewayGatewayLogListGatewayLogsParamsFiltersOperator = "eq"
	AccountAIGatewayGatewayLogListGatewayLogsParamsFiltersOperatorNeq      AccountAIGatewayGatewayLogListGatewayLogsParamsFiltersOperator = "neq"
	AccountAIGatewayGatewayLogListGatewayLogsParamsFiltersOperatorContains AccountAIGatewayGatewayLogListGatewayLogsParamsFiltersOperator = "contains"
	AccountAIGatewayGatewayLogListGatewayLogsParamsFiltersOperatorLt       AccountAIGatewayGatewayLogListGatewayLogsParamsFiltersOperator = "lt"
	AccountAIGatewayGatewayLogListGatewayLogsParamsFiltersOperatorGt       AccountAIGatewayGatewayLogListGatewayLogsParamsFiltersOperator = "gt"
)

func (r AccountAIGatewayGatewayLogListGatewayLogsParamsFiltersOperator) IsKnown() bool {
	switch r {
	case AccountAIGatewayGatewayLogListGatewayLogsParamsFiltersOperatorEq, AccountAIGatewayGatewayLogListGatewayLogsParamsFiltersOperatorNeq, AccountAIGatewayGatewayLogListGatewayLogsParamsFiltersOperatorContains, AccountAIGatewayGatewayLogListGatewayLogsParamsFiltersOperatorLt, AccountAIGatewayGatewayLogListGatewayLogsParamsFiltersOperatorGt:
		return true
	}
	return false
}

// Satisfied by [shared.UnionString], [shared.UnionFloat], [shared.UnionBool].
type AccountAIGatewayGatewayLogListGatewayLogsParamsFiltersValueUnion interface {
	ImplementsAccountAIGatewayGatewayLogListGatewayLogsParamsFiltersValueUnion()
}

type AccountAIGatewayGatewayLogListGatewayLogsParamsOrderBy string

const (
	AccountAIGatewayGatewayLogListGatewayLogsParamsOrderByCreatedAt AccountAIGatewayGatewayLogListGatewayLogsParamsOrderBy = "created_at"
	AccountAIGatewayGatewayLogListGatewayLogsParamsOrderByProvider  AccountAIGatewayGatewayLogListGatewayLogsParamsOrderBy = "provider"
	AccountAIGatewayGatewayLogListGatewayLogsParamsOrderByModel     AccountAIGatewayGatewayLogListGatewayLogsParamsOrderBy = "model"
	AccountAIGatewayGatewayLogListGatewayLogsParamsOrderByModelType AccountAIGatewayGatewayLogListGatewayLogsParamsOrderBy = "model_type"
	AccountAIGatewayGatewayLogListGatewayLogsParamsOrderBySuccess   AccountAIGatewayGatewayLogListGatewayLogsParamsOrderBy = "success"
	AccountAIGatewayGatewayLogListGatewayLogsParamsOrderByCached    AccountAIGatewayGatewayLogListGatewayLogsParamsOrderBy = "cached"
)

func (r AccountAIGatewayGatewayLogListGatewayLogsParamsOrderBy) IsKnown() bool {
	switch r {
	case AccountAIGatewayGatewayLogListGatewayLogsParamsOrderByCreatedAt, AccountAIGatewayGatewayLogListGatewayLogsParamsOrderByProvider, AccountAIGatewayGatewayLogListGatewayLogsParamsOrderByModel, AccountAIGatewayGatewayLogListGatewayLogsParamsOrderByModelType, AccountAIGatewayGatewayLogListGatewayLogsParamsOrderBySuccess, AccountAIGatewayGatewayLogListGatewayLogsParamsOrderByCached:
		return true
	}
	return false
}

type AccountAIGatewayGatewayLogListGatewayLogsParamsOrderByDirection string

const (
	AccountAIGatewayGatewayLogListGatewayLogsParamsOrderByDirectionAsc  AccountAIGatewayGatewayLogListGatewayLogsParamsOrderByDirection = "asc"
	AccountAIGatewayGatewayLogListGatewayLogsParamsOrderByDirectionDesc AccountAIGatewayGatewayLogListGatewayLogsParamsOrderByDirection = "desc"
)

func (r AccountAIGatewayGatewayLogListGatewayLogsParamsOrderByDirection) IsKnown() bool {
	switch r {
	case AccountAIGatewayGatewayLogListGatewayLogsParamsOrderByDirectionAsc, AccountAIGatewayGatewayLogListGatewayLogsParamsOrderByDirectionDesc:
		return true
	}
	return false
}

type AccountAIGatewayGatewayLogPatchGatewayLogParams struct {
	Feedback param.Field[float64]                                                                 `json:"feedback"`
	Metadata param.Field[map[string]AccountAIGatewayGatewayLogPatchGatewayLogParamsMetadataUnion] `json:"metadata"`
	Score    param.Field[float64]                                                                 `json:"score"`
}

func (r AccountAIGatewayGatewayLogPatchGatewayLogParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by [shared.UnionString], [shared.UnionFloat], [shared.UnionBool].
type AccountAIGatewayGatewayLogPatchGatewayLogParamsMetadataUnion interface {
	ImplementsAccountAIGatewayGatewayLogPatchGatewayLogParamsMetadataUnion()
}
