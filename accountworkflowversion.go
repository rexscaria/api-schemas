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

// AccountWorkflowVersionService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountWorkflowVersionService] method instead.
type AccountWorkflowVersionService struct {
	Options []option.RequestOption
}

// NewAccountWorkflowVersionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountWorkflowVersionService(opts ...option.RequestOption) (r *AccountWorkflowVersionService) {
	r = &AccountWorkflowVersionService{}
	r.Options = opts
	return
}

// Get Workflow version details
func (r *AccountWorkflowVersionService) Get(ctx context.Context, accountID string, workflowName string, versionID string, opts ...option.RequestOption) (res *AccountWorkflowVersionGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if workflowName == "" {
		err = errors.New("missing required workflow_name parameter")
		return
	}
	if versionID == "" {
		err = errors.New("missing required version_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workflows/%s/versions/%s", accountID, workflowName, versionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List deployed Workflow versions
func (r *AccountWorkflowVersionService) List(ctx context.Context, accountID string, workflowName string, query AccountWorkflowVersionListParams, opts ...option.RequestOption) (res *AccountWorkflowVersionListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if workflowName == "" {
		err = errors.New("missing required workflow_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workflows/%s/versions", accountID, workflowName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type AccountWorkflowVersionGetResponse struct {
	Errors     []AccountWorkflowVersionGetResponseError    `json:"errors,required"`
	Messages   []AccountWorkflowVersionGetResponseMessage  `json:"messages,required"`
	Result     AccountWorkflowVersionGetResponseResult     `json:"result,required"`
	Success    AccountWorkflowVersionGetResponseSuccess    `json:"success,required"`
	ResultInfo AccountWorkflowVersionGetResponseResultInfo `json:"result_info"`
	JSON       accountWorkflowVersionGetResponseJSON       `json:"-"`
}

// accountWorkflowVersionGetResponseJSON contains the JSON metadata for the struct
// [AccountWorkflowVersionGetResponse]
type accountWorkflowVersionGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkflowVersionGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountWorkflowVersionGetResponseJSON) RawJSON() string {
	return r.raw
}

type AccountWorkflowVersionGetResponseError struct {
	Code    float64                                    `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    accountWorkflowVersionGetResponseErrorJSON `json:"-"`
}

// accountWorkflowVersionGetResponseErrorJSON contains the JSON metadata for the
// struct [AccountWorkflowVersionGetResponseError]
type accountWorkflowVersionGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkflowVersionGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountWorkflowVersionGetResponseErrorJSON) RawJSON() string {
	return r.raw
}

type AccountWorkflowVersionGetResponseMessage struct {
	Code    float64                                      `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    accountWorkflowVersionGetResponseMessageJSON `json:"-"`
}

// accountWorkflowVersionGetResponseMessageJSON contains the JSON metadata for the
// struct [AccountWorkflowVersionGetResponseMessage]
type accountWorkflowVersionGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkflowVersionGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountWorkflowVersionGetResponseMessageJSON) RawJSON() string {
	return r.raw
}

type AccountWorkflowVersionGetResponseResult struct {
	ID         string                                      `json:"id,required" format:"uuid"`
	ClassName  string                                      `json:"class_name,required"`
	CreatedOn  time.Time                                   `json:"created_on,required" format:"date-time"`
	ModifiedOn time.Time                                   `json:"modified_on,required" format:"date-time"`
	WorkflowID string                                      `json:"workflow_id,required" format:"uuid"`
	JSON       accountWorkflowVersionGetResponseResultJSON `json:"-"`
}

// accountWorkflowVersionGetResponseResultJSON contains the JSON metadata for the
// struct [AccountWorkflowVersionGetResponseResult]
type accountWorkflowVersionGetResponseResultJSON struct {
	ID          apijson.Field
	ClassName   apijson.Field
	CreatedOn   apijson.Field
	ModifiedOn  apijson.Field
	WorkflowID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkflowVersionGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountWorkflowVersionGetResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountWorkflowVersionGetResponseSuccess bool

const (
	AccountWorkflowVersionGetResponseSuccessTrue AccountWorkflowVersionGetResponseSuccess = true
)

func (r AccountWorkflowVersionGetResponseSuccess) IsKnown() bool {
	switch r {
	case AccountWorkflowVersionGetResponseSuccessTrue:
		return true
	}
	return false
}

type AccountWorkflowVersionGetResponseResultInfo struct {
	Count      float64                                         `json:"count,required"`
	Page       float64                                         `json:"page,required"`
	PerPage    float64                                         `json:"per_page,required"`
	TotalCount float64                                         `json:"total_count,required"`
	JSON       accountWorkflowVersionGetResponseResultInfoJSON `json:"-"`
}

// accountWorkflowVersionGetResponseResultInfoJSON contains the JSON metadata for
// the struct [AccountWorkflowVersionGetResponseResultInfo]
type accountWorkflowVersionGetResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkflowVersionGetResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountWorkflowVersionGetResponseResultInfoJSON) RawJSON() string {
	return r.raw
}

type AccountWorkflowVersionListResponse struct {
	Errors     []AccountWorkflowVersionListResponseError    `json:"errors,required"`
	Messages   []AccountWorkflowVersionListResponseMessage  `json:"messages,required"`
	Result     []AccountWorkflowVersionListResponseResult   `json:"result,required"`
	Success    AccountWorkflowVersionListResponseSuccess    `json:"success,required"`
	ResultInfo AccountWorkflowVersionListResponseResultInfo `json:"result_info"`
	JSON       accountWorkflowVersionListResponseJSON       `json:"-"`
}

// accountWorkflowVersionListResponseJSON contains the JSON metadata for the struct
// [AccountWorkflowVersionListResponse]
type accountWorkflowVersionListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkflowVersionListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountWorkflowVersionListResponseJSON) RawJSON() string {
	return r.raw
}

type AccountWorkflowVersionListResponseError struct {
	Code    float64                                     `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    accountWorkflowVersionListResponseErrorJSON `json:"-"`
}

// accountWorkflowVersionListResponseErrorJSON contains the JSON metadata for the
// struct [AccountWorkflowVersionListResponseError]
type accountWorkflowVersionListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkflowVersionListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountWorkflowVersionListResponseErrorJSON) RawJSON() string {
	return r.raw
}

type AccountWorkflowVersionListResponseMessage struct {
	Code    float64                                       `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    accountWorkflowVersionListResponseMessageJSON `json:"-"`
}

// accountWorkflowVersionListResponseMessageJSON contains the JSON metadata for the
// struct [AccountWorkflowVersionListResponseMessage]
type accountWorkflowVersionListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkflowVersionListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountWorkflowVersionListResponseMessageJSON) RawJSON() string {
	return r.raw
}

type AccountWorkflowVersionListResponseResult struct {
	ID         string                                       `json:"id,required" format:"uuid"`
	ClassName  string                                       `json:"class_name,required"`
	CreatedOn  time.Time                                    `json:"created_on,required" format:"date-time"`
	ModifiedOn time.Time                                    `json:"modified_on,required" format:"date-time"`
	WorkflowID string                                       `json:"workflow_id,required" format:"uuid"`
	JSON       accountWorkflowVersionListResponseResultJSON `json:"-"`
}

// accountWorkflowVersionListResponseResultJSON contains the JSON metadata for the
// struct [AccountWorkflowVersionListResponseResult]
type accountWorkflowVersionListResponseResultJSON struct {
	ID          apijson.Field
	ClassName   apijson.Field
	CreatedOn   apijson.Field
	ModifiedOn  apijson.Field
	WorkflowID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkflowVersionListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountWorkflowVersionListResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountWorkflowVersionListResponseSuccess bool

const (
	AccountWorkflowVersionListResponseSuccessTrue AccountWorkflowVersionListResponseSuccess = true
)

func (r AccountWorkflowVersionListResponseSuccess) IsKnown() bool {
	switch r {
	case AccountWorkflowVersionListResponseSuccessTrue:
		return true
	}
	return false
}

type AccountWorkflowVersionListResponseResultInfo struct {
	Count      float64                                          `json:"count,required"`
	Page       float64                                          `json:"page,required"`
	PerPage    float64                                          `json:"per_page,required"`
	TotalCount float64                                          `json:"total_count,required"`
	JSON       accountWorkflowVersionListResponseResultInfoJSON `json:"-"`
}

// accountWorkflowVersionListResponseResultInfoJSON contains the JSON metadata for
// the struct [AccountWorkflowVersionListResponseResultInfo]
type accountWorkflowVersionListResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkflowVersionListResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountWorkflowVersionListResponseResultInfoJSON) RawJSON() string {
	return r.raw
}

type AccountWorkflowVersionListParams struct {
	Page    param.Field[float64] `query:"page"`
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes [AccountWorkflowVersionListParams]'s query parameters as
// `url.Values`.
func (r AccountWorkflowVersionListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
