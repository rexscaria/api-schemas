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

// AccountWorkflowService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountWorkflowService] method instead.
type AccountWorkflowService struct {
	Options   []option.RequestOption
	Instances *AccountWorkflowInstanceService
	Versions  *AccountWorkflowVersionService
}

// NewAccountWorkflowService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccountWorkflowService(opts ...option.RequestOption) (r *AccountWorkflowService) {
	r = &AccountWorkflowService{}
	r.Options = opts
	r.Instances = NewAccountWorkflowInstanceService(opts...)
	r.Versions = NewAccountWorkflowVersionService(opts...)
	return
}

// Get Workflow details
func (r *AccountWorkflowService) Get(ctx context.Context, accountID string, workflowName string, opts ...option.RequestOption) (res *AccountWorkflowGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if workflowName == "" {
		err = errors.New("missing required workflow_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workflows/%s", accountID, workflowName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Create/modify Workflow
func (r *AccountWorkflowService) Update(ctx context.Context, accountID string, workflowName string, body AccountWorkflowUpdateParams, opts ...option.RequestOption) (res *AccountWorkflowUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if workflowName == "" {
		err = errors.New("missing required workflow_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workflows/%s", accountID, workflowName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// List all Workflows
func (r *AccountWorkflowService) List(ctx context.Context, accountID string, query AccountWorkflowListParams, opts ...option.RequestOption) (res *AccountWorkflowListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workflows", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type AccountWorkflowGetResponse struct {
	Errors     []AccountWorkflowGetResponseError    `json:"errors,required"`
	Messages   []AccountWorkflowGetResponseMessage  `json:"messages,required"`
	Result     AccountWorkflowGetResponseResult     `json:"result,required"`
	Success    AccountWorkflowGetResponseSuccess    `json:"success,required"`
	ResultInfo AccountWorkflowGetResponseResultInfo `json:"result_info"`
	JSON       accountWorkflowGetResponseJSON       `json:"-"`
}

// accountWorkflowGetResponseJSON contains the JSON metadata for the struct
// [AccountWorkflowGetResponse]
type accountWorkflowGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkflowGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountWorkflowGetResponseJSON) RawJSON() string {
	return r.raw
}

type AccountWorkflowGetResponseError struct {
	Code    float64                             `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    accountWorkflowGetResponseErrorJSON `json:"-"`
}

// accountWorkflowGetResponseErrorJSON contains the JSON metadata for the struct
// [AccountWorkflowGetResponseError]
type accountWorkflowGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkflowGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountWorkflowGetResponseErrorJSON) RawJSON() string {
	return r.raw
}

type AccountWorkflowGetResponseMessage struct {
	Code    float64                               `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    accountWorkflowGetResponseMessageJSON `json:"-"`
}

// accountWorkflowGetResponseMessageJSON contains the JSON metadata for the struct
// [AccountWorkflowGetResponseMessage]
type accountWorkflowGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkflowGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountWorkflowGetResponseMessageJSON) RawJSON() string {
	return r.raw
}

type AccountWorkflowGetResponseResult struct {
	ID          string                                    `json:"id,required" format:"uuid"`
	ClassName   string                                    `json:"class_name,required"`
	CreatedOn   time.Time                                 `json:"created_on,required" format:"date-time"`
	Instances   AccountWorkflowGetResponseResultInstances `json:"instances,required"`
	ModifiedOn  time.Time                                 `json:"modified_on,required" format:"date-time"`
	Name        string                                    `json:"name,required"`
	ScriptName  string                                    `json:"script_name,required"`
	TriggeredOn time.Time                                 `json:"triggered_on,required,nullable" format:"date-time"`
	JSON        accountWorkflowGetResponseResultJSON      `json:"-"`
}

// accountWorkflowGetResponseResultJSON contains the JSON metadata for the struct
// [AccountWorkflowGetResponseResult]
type accountWorkflowGetResponseResultJSON struct {
	ID          apijson.Field
	ClassName   apijson.Field
	CreatedOn   apijson.Field
	Instances   apijson.Field
	ModifiedOn  apijson.Field
	Name        apijson.Field
	ScriptName  apijson.Field
	TriggeredOn apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkflowGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountWorkflowGetResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountWorkflowGetResponseResultInstances struct {
	Complete        float64                                       `json:"complete"`
	Errored         float64                                       `json:"errored"`
	Paused          float64                                       `json:"paused"`
	Queued          float64                                       `json:"queued"`
	Running         float64                                       `json:"running"`
	Terminated      float64                                       `json:"terminated"`
	Waiting         float64                                       `json:"waiting"`
	WaitingForPause float64                                       `json:"waitingForPause"`
	JSON            accountWorkflowGetResponseResultInstancesJSON `json:"-"`
}

// accountWorkflowGetResponseResultInstancesJSON contains the JSON metadata for the
// struct [AccountWorkflowGetResponseResultInstances]
type accountWorkflowGetResponseResultInstancesJSON struct {
	Complete        apijson.Field
	Errored         apijson.Field
	Paused          apijson.Field
	Queued          apijson.Field
	Running         apijson.Field
	Terminated      apijson.Field
	Waiting         apijson.Field
	WaitingForPause apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccountWorkflowGetResponseResultInstances) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountWorkflowGetResponseResultInstancesJSON) RawJSON() string {
	return r.raw
}

type AccountWorkflowGetResponseSuccess bool

const (
	AccountWorkflowGetResponseSuccessTrue AccountWorkflowGetResponseSuccess = true
)

func (r AccountWorkflowGetResponseSuccess) IsKnown() bool {
	switch r {
	case AccountWorkflowGetResponseSuccessTrue:
		return true
	}
	return false
}

type AccountWorkflowGetResponseResultInfo struct {
	Count      float64                                  `json:"count,required"`
	Page       float64                                  `json:"page,required"`
	PerPage    float64                                  `json:"per_page,required"`
	TotalCount float64                                  `json:"total_count,required"`
	JSON       accountWorkflowGetResponseResultInfoJSON `json:"-"`
}

// accountWorkflowGetResponseResultInfoJSON contains the JSON metadata for the
// struct [AccountWorkflowGetResponseResultInfo]
type accountWorkflowGetResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkflowGetResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountWorkflowGetResponseResultInfoJSON) RawJSON() string {
	return r.raw
}

type AccountWorkflowUpdateResponse struct {
	Errors     []AccountWorkflowUpdateResponseError    `json:"errors,required"`
	Messages   []AccountWorkflowUpdateResponseMessage  `json:"messages,required"`
	Result     AccountWorkflowUpdateResponseResult     `json:"result,required"`
	Success    AccountWorkflowUpdateResponseSuccess    `json:"success,required"`
	ResultInfo AccountWorkflowUpdateResponseResultInfo `json:"result_info"`
	JSON       accountWorkflowUpdateResponseJSON       `json:"-"`
}

// accountWorkflowUpdateResponseJSON contains the JSON metadata for the struct
// [AccountWorkflowUpdateResponse]
type accountWorkflowUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkflowUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountWorkflowUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type AccountWorkflowUpdateResponseError struct {
	Code    float64                                `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    accountWorkflowUpdateResponseErrorJSON `json:"-"`
}

// accountWorkflowUpdateResponseErrorJSON contains the JSON metadata for the struct
// [AccountWorkflowUpdateResponseError]
type accountWorkflowUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkflowUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountWorkflowUpdateResponseErrorJSON) RawJSON() string {
	return r.raw
}

type AccountWorkflowUpdateResponseMessage struct {
	Code    float64                                  `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    accountWorkflowUpdateResponseMessageJSON `json:"-"`
}

// accountWorkflowUpdateResponseMessageJSON contains the JSON metadata for the
// struct [AccountWorkflowUpdateResponseMessage]
type accountWorkflowUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkflowUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountWorkflowUpdateResponseMessageJSON) RawJSON() string {
	return r.raw
}

type AccountWorkflowUpdateResponseResult struct {
	ID                string                                  `json:"id,required" format:"uuid"`
	ClassName         string                                  `json:"class_name,required"`
	CreatedOn         time.Time                               `json:"created_on,required" format:"date-time"`
	IsDeleted         float64                                 `json:"is_deleted,required"`
	ModifiedOn        time.Time                               `json:"modified_on,required" format:"date-time"`
	Name              string                                  `json:"name,required"`
	ScriptName        string                                  `json:"script_name,required"`
	TerminatorRunning float64                                 `json:"terminator_running,required"`
	TriggeredOn       time.Time                               `json:"triggered_on,required,nullable" format:"date-time"`
	VersionID         string                                  `json:"version_id,required" format:"uuid"`
	JSON              accountWorkflowUpdateResponseResultJSON `json:"-"`
}

// accountWorkflowUpdateResponseResultJSON contains the JSON metadata for the
// struct [AccountWorkflowUpdateResponseResult]
type accountWorkflowUpdateResponseResultJSON struct {
	ID                apijson.Field
	ClassName         apijson.Field
	CreatedOn         apijson.Field
	IsDeleted         apijson.Field
	ModifiedOn        apijson.Field
	Name              apijson.Field
	ScriptName        apijson.Field
	TerminatorRunning apijson.Field
	TriggeredOn       apijson.Field
	VersionID         apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *AccountWorkflowUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountWorkflowUpdateResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountWorkflowUpdateResponseSuccess bool

const (
	AccountWorkflowUpdateResponseSuccessTrue AccountWorkflowUpdateResponseSuccess = true
)

func (r AccountWorkflowUpdateResponseSuccess) IsKnown() bool {
	switch r {
	case AccountWorkflowUpdateResponseSuccessTrue:
		return true
	}
	return false
}

type AccountWorkflowUpdateResponseResultInfo struct {
	Count      float64                                     `json:"count,required"`
	Page       float64                                     `json:"page,required"`
	PerPage    float64                                     `json:"per_page,required"`
	TotalCount float64                                     `json:"total_count,required"`
	JSON       accountWorkflowUpdateResponseResultInfoJSON `json:"-"`
}

// accountWorkflowUpdateResponseResultInfoJSON contains the JSON metadata for the
// struct [AccountWorkflowUpdateResponseResultInfo]
type accountWorkflowUpdateResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkflowUpdateResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountWorkflowUpdateResponseResultInfoJSON) RawJSON() string {
	return r.raw
}

type AccountWorkflowListResponse struct {
	Errors     []AccountWorkflowListResponseError    `json:"errors,required"`
	Messages   []AccountWorkflowListResponseMessage  `json:"messages,required"`
	Result     []AccountWorkflowListResponseResult   `json:"result,required"`
	Success    AccountWorkflowListResponseSuccess    `json:"success,required"`
	ResultInfo AccountWorkflowListResponseResultInfo `json:"result_info"`
	JSON       accountWorkflowListResponseJSON       `json:"-"`
}

// accountWorkflowListResponseJSON contains the JSON metadata for the struct
// [AccountWorkflowListResponse]
type accountWorkflowListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkflowListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountWorkflowListResponseJSON) RawJSON() string {
	return r.raw
}

type AccountWorkflowListResponseError struct {
	Code    float64                              `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    accountWorkflowListResponseErrorJSON `json:"-"`
}

// accountWorkflowListResponseErrorJSON contains the JSON metadata for the struct
// [AccountWorkflowListResponseError]
type accountWorkflowListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkflowListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountWorkflowListResponseErrorJSON) RawJSON() string {
	return r.raw
}

type AccountWorkflowListResponseMessage struct {
	Code    float64                                `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    accountWorkflowListResponseMessageJSON `json:"-"`
}

// accountWorkflowListResponseMessageJSON contains the JSON metadata for the struct
// [AccountWorkflowListResponseMessage]
type accountWorkflowListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkflowListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountWorkflowListResponseMessageJSON) RawJSON() string {
	return r.raw
}

type AccountWorkflowListResponseResult struct {
	ID          string                                     `json:"id,required" format:"uuid"`
	ClassName   string                                     `json:"class_name,required"`
	CreatedOn   time.Time                                  `json:"created_on,required" format:"date-time"`
	Instances   AccountWorkflowListResponseResultInstances `json:"instances,required"`
	ModifiedOn  time.Time                                  `json:"modified_on,required" format:"date-time"`
	Name        string                                     `json:"name,required"`
	ScriptName  string                                     `json:"script_name,required"`
	TriggeredOn time.Time                                  `json:"triggered_on,required,nullable" format:"date-time"`
	JSON        accountWorkflowListResponseResultJSON      `json:"-"`
}

// accountWorkflowListResponseResultJSON contains the JSON metadata for the struct
// [AccountWorkflowListResponseResult]
type accountWorkflowListResponseResultJSON struct {
	ID          apijson.Field
	ClassName   apijson.Field
	CreatedOn   apijson.Field
	Instances   apijson.Field
	ModifiedOn  apijson.Field
	Name        apijson.Field
	ScriptName  apijson.Field
	TriggeredOn apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkflowListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountWorkflowListResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountWorkflowListResponseResultInstances struct {
	Complete        float64                                        `json:"complete"`
	Errored         float64                                        `json:"errored"`
	Paused          float64                                        `json:"paused"`
	Queued          float64                                        `json:"queued"`
	Running         float64                                        `json:"running"`
	Terminated      float64                                        `json:"terminated"`
	Waiting         float64                                        `json:"waiting"`
	WaitingForPause float64                                        `json:"waitingForPause"`
	JSON            accountWorkflowListResponseResultInstancesJSON `json:"-"`
}

// accountWorkflowListResponseResultInstancesJSON contains the JSON metadata for
// the struct [AccountWorkflowListResponseResultInstances]
type accountWorkflowListResponseResultInstancesJSON struct {
	Complete        apijson.Field
	Errored         apijson.Field
	Paused          apijson.Field
	Queued          apijson.Field
	Running         apijson.Field
	Terminated      apijson.Field
	Waiting         apijson.Field
	WaitingForPause apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccountWorkflowListResponseResultInstances) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountWorkflowListResponseResultInstancesJSON) RawJSON() string {
	return r.raw
}

type AccountWorkflowListResponseSuccess bool

const (
	AccountWorkflowListResponseSuccessTrue AccountWorkflowListResponseSuccess = true
)

func (r AccountWorkflowListResponseSuccess) IsKnown() bool {
	switch r {
	case AccountWorkflowListResponseSuccessTrue:
		return true
	}
	return false
}

type AccountWorkflowListResponseResultInfo struct {
	Count      float64                                   `json:"count,required"`
	Page       float64                                   `json:"page,required"`
	PerPage    float64                                   `json:"per_page,required"`
	TotalCount float64                                   `json:"total_count,required"`
	JSON       accountWorkflowListResponseResultInfoJSON `json:"-"`
}

// accountWorkflowListResponseResultInfoJSON contains the JSON metadata for the
// struct [AccountWorkflowListResponseResultInfo]
type accountWorkflowListResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkflowListResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountWorkflowListResponseResultInfoJSON) RawJSON() string {
	return r.raw
}

type AccountWorkflowUpdateParams struct {
	ClassName  param.Field[string] `json:"class_name,required"`
	ScriptName param.Field[string] `json:"script_name,required"`
}

func (r AccountWorkflowUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountWorkflowListParams struct {
	Page    param.Field[float64] `query:"page"`
	PerPage param.Field[float64] `query:"per_page"`
	// Allows filtering workflows` name.
	Search param.Field[string] `query:"search"`
}

// URLQuery serializes [AccountWorkflowListParams]'s query parameters as
// `url.Values`.
func (r AccountWorkflowListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
