// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"mime/multipart"
	"net/http"
	"net/url"

	"github.com/rexscaria/api-schemas/internal/apiform"
	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/apiquery"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
)

// AccountPageProjectDeploymentService contains methods and other services that
// help with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountPageProjectDeploymentService] method instead.
type AccountPageProjectDeploymentService struct {
	Options []option.RequestOption
	History *AccountPageProjectDeploymentHistoryService
}

// NewAccountPageProjectDeploymentService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountPageProjectDeploymentService(opts ...option.RequestOption) (r *AccountPageProjectDeploymentService) {
	r = &AccountPageProjectDeploymentService{}
	r.Options = opts
	r.History = NewAccountPageProjectDeploymentHistoryService(opts...)
	return
}

// Start a new deployment from production. The repository and account must have
// already been authorized on the Cloudflare Pages dashboard.
func (r *AccountPageProjectDeploymentService) New(ctx context.Context, accountID string, projectName string, body AccountPageProjectDeploymentNewParams, opts ...option.RequestOption) (res *NewDeployment, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if projectName == "" {
		err = errors.New("missing required project_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/pages/projects/%s/deployments", accountID, projectName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Fetch information about a deployment.
func (r *AccountPageProjectDeploymentService) Get(ctx context.Context, accountID string, projectName string, deploymentID string, opts ...option.RequestOption) (res *ResponseDetails, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if projectName == "" {
		err = errors.New("missing required project_name parameter")
		return
	}
	if deploymentID == "" {
		err = errors.New("missing required deployment_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/pages/projects/%s/deployments/%s", accountID, projectName, deploymentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Fetch a list of project deployments.
func (r *AccountPageProjectDeploymentService) List(ctx context.Context, accountID string, projectName string, query AccountPageProjectDeploymentListParams, opts ...option.RequestOption) (res *AccountPageProjectDeploymentListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if projectName == "" {
		err = errors.New("missing required project_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/pages/projects/%s/deployments", accountID, projectName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete a deployment.
func (r *AccountPageProjectDeploymentService) Delete(ctx context.Context, accountID string, projectName string, deploymentID string, opts ...option.RequestOption) (res *AccountPageProjectDeploymentDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if projectName == "" {
		err = errors.New("missing required project_name parameter")
		return
	}
	if deploymentID == "" {
		err = errors.New("missing required deployment_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/pages/projects/%s/deployments/%s", accountID, projectName, deploymentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Retry a previous deployment.
func (r *AccountPageProjectDeploymentService) Retry(ctx context.Context, accountID string, projectName string, deploymentID string, body AccountPageProjectDeploymentRetryParams, opts ...option.RequestOption) (res *NewDeployment, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if projectName == "" {
		err = errors.New("missing required project_name parameter")
		return
	}
	if deploymentID == "" {
		err = errors.New("missing required deployment_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/pages/projects/%s/deployments/%s/retry", accountID, projectName, deploymentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Rollback the production deployment to a previous deployment. You can only
// rollback to succesful builds on production.
func (r *AccountPageProjectDeploymentService) Rollback(ctx context.Context, accountID string, projectName string, deploymentID string, body AccountPageProjectDeploymentRollbackParams, opts ...option.RequestOption) (res *ResponseDetails, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if projectName == "" {
		err = errors.New("missing required project_name parameter")
		return
	}
	if deploymentID == "" {
		err = errors.New("missing required deployment_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/pages/projects/%s/deployments/%s/rollback", accountID, projectName, deploymentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type NewDeployment struct {
	Errors   []NewDeploymentError   `json:"errors,required"`
	Messages []NewDeploymentMessage `json:"messages,required"`
	Result   Deployments            `json:"result,required"`
	// Whether the API call was successful
	Success NewDeploymentSuccess `json:"success,required"`
	JSON    newDeploymentJSON    `json:"-"`
}

// newDeploymentJSON contains the JSON metadata for the struct [NewDeployment]
type newDeploymentJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NewDeployment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r newDeploymentJSON) RawJSON() string {
	return r.raw
}

type NewDeploymentError struct {
	Code             int64                     `json:"code,required"`
	Message          string                    `json:"message,required"`
	DocumentationURL string                    `json:"documentation_url"`
	Source           NewDeploymentErrorsSource `json:"source"`
	JSON             newDeploymentErrorJSON    `json:"-"`
}

// newDeploymentErrorJSON contains the JSON metadata for the struct
// [NewDeploymentError]
type newDeploymentErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *NewDeploymentError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r newDeploymentErrorJSON) RawJSON() string {
	return r.raw
}

type NewDeploymentErrorsSource struct {
	Pointer string                        `json:"pointer"`
	JSON    newDeploymentErrorsSourceJSON `json:"-"`
}

// newDeploymentErrorsSourceJSON contains the JSON metadata for the struct
// [NewDeploymentErrorsSource]
type newDeploymentErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NewDeploymentErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r newDeploymentErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type NewDeploymentMessage struct {
	Code             int64                       `json:"code,required"`
	Message          string                      `json:"message,required"`
	DocumentationURL string                      `json:"documentation_url"`
	Source           NewDeploymentMessagesSource `json:"source"`
	JSON             newDeploymentMessageJSON    `json:"-"`
}

// newDeploymentMessageJSON contains the JSON metadata for the struct
// [NewDeploymentMessage]
type newDeploymentMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *NewDeploymentMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r newDeploymentMessageJSON) RawJSON() string {
	return r.raw
}

type NewDeploymentMessagesSource struct {
	Pointer string                          `json:"pointer"`
	JSON    newDeploymentMessagesSourceJSON `json:"-"`
}

// newDeploymentMessagesSourceJSON contains the JSON metadata for the struct
// [NewDeploymentMessagesSource]
type newDeploymentMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NewDeploymentMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r newDeploymentMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type NewDeploymentSuccess bool

const (
	NewDeploymentSuccessFalse NewDeploymentSuccess = false
	NewDeploymentSuccessTrue  NewDeploymentSuccess = true
)

func (r NewDeploymentSuccess) IsKnown() bool {
	switch r {
	case NewDeploymentSuccessFalse, NewDeploymentSuccessTrue:
		return true
	}
	return false
}

type ResponseDetails struct {
	Errors   []ResponseDetailsError   `json:"errors,required"`
	Messages []ResponseDetailsMessage `json:"messages,required"`
	Result   Deployments              `json:"result,required"`
	// Whether the API call was successful
	Success ResponseDetailsSuccess `json:"success,required"`
	JSON    responseDetailsJSON    `json:"-"`
}

// responseDetailsJSON contains the JSON metadata for the struct [ResponseDetails]
type responseDetailsJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResponseDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r responseDetailsJSON) RawJSON() string {
	return r.raw
}

type ResponseDetailsError struct {
	Code             int64                       `json:"code,required"`
	Message          string                      `json:"message,required"`
	DocumentationURL string                      `json:"documentation_url"`
	Source           ResponseDetailsErrorsSource `json:"source"`
	JSON             responseDetailsErrorJSON    `json:"-"`
}

// responseDetailsErrorJSON contains the JSON metadata for the struct
// [ResponseDetailsError]
type responseDetailsErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ResponseDetailsError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r responseDetailsErrorJSON) RawJSON() string {
	return r.raw
}

type ResponseDetailsErrorsSource struct {
	Pointer string                          `json:"pointer"`
	JSON    responseDetailsErrorsSourceJSON `json:"-"`
}

// responseDetailsErrorsSourceJSON contains the JSON metadata for the struct
// [ResponseDetailsErrorsSource]
type responseDetailsErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResponseDetailsErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r responseDetailsErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type ResponseDetailsMessage struct {
	Code             int64                         `json:"code,required"`
	Message          string                        `json:"message,required"`
	DocumentationURL string                        `json:"documentation_url"`
	Source           ResponseDetailsMessagesSource `json:"source"`
	JSON             responseDetailsMessageJSON    `json:"-"`
}

// responseDetailsMessageJSON contains the JSON metadata for the struct
// [ResponseDetailsMessage]
type responseDetailsMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ResponseDetailsMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r responseDetailsMessageJSON) RawJSON() string {
	return r.raw
}

type ResponseDetailsMessagesSource struct {
	Pointer string                            `json:"pointer"`
	JSON    responseDetailsMessagesSourceJSON `json:"-"`
}

// responseDetailsMessagesSourceJSON contains the JSON metadata for the struct
// [ResponseDetailsMessagesSource]
type responseDetailsMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResponseDetailsMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r responseDetailsMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ResponseDetailsSuccess bool

const (
	ResponseDetailsSuccessFalse ResponseDetailsSuccess = false
	ResponseDetailsSuccessTrue  ResponseDetailsSuccess = true
)

func (r ResponseDetailsSuccess) IsKnown() bool {
	switch r {
	case ResponseDetailsSuccessFalse, ResponseDetailsSuccessTrue:
		return true
	}
	return false
}

type AccountPageProjectDeploymentListResponse struct {
	Errors   []AccountPageProjectDeploymentListResponseError   `json:"errors,required"`
	Messages []AccountPageProjectDeploymentListResponseMessage `json:"messages,required"`
	Result   []Deployments                                     `json:"result,required"`
	// Whether the API call was successful
	Success    AccountPageProjectDeploymentListResponseSuccess    `json:"success,required"`
	ResultInfo AccountPageProjectDeploymentListResponseResultInfo `json:"result_info"`
	JSON       accountPageProjectDeploymentListResponseJSON       `json:"-"`
}

// accountPageProjectDeploymentListResponseJSON contains the JSON metadata for the
// struct [AccountPageProjectDeploymentListResponse]
type accountPageProjectDeploymentListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountPageProjectDeploymentListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountPageProjectDeploymentListResponseJSON) RawJSON() string {
	return r.raw
}

type AccountPageProjectDeploymentListResponseError struct {
	Code             int64                                                `json:"code,required"`
	Message          string                                               `json:"message,required"`
	DocumentationURL string                                               `json:"documentation_url"`
	Source           AccountPageProjectDeploymentListResponseErrorsSource `json:"source"`
	JSON             accountPageProjectDeploymentListResponseErrorJSON    `json:"-"`
}

// accountPageProjectDeploymentListResponseErrorJSON contains the JSON metadata for
// the struct [AccountPageProjectDeploymentListResponseError]
type accountPageProjectDeploymentListResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountPageProjectDeploymentListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountPageProjectDeploymentListResponseErrorJSON) RawJSON() string {
	return r.raw
}

type AccountPageProjectDeploymentListResponseErrorsSource struct {
	Pointer string                                                   `json:"pointer"`
	JSON    accountPageProjectDeploymentListResponseErrorsSourceJSON `json:"-"`
}

// accountPageProjectDeploymentListResponseErrorsSourceJSON contains the JSON
// metadata for the struct [AccountPageProjectDeploymentListResponseErrorsSource]
type accountPageProjectDeploymentListResponseErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountPageProjectDeploymentListResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountPageProjectDeploymentListResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AccountPageProjectDeploymentListResponseMessage struct {
	Code             int64                                                  `json:"code,required"`
	Message          string                                                 `json:"message,required"`
	DocumentationURL string                                                 `json:"documentation_url"`
	Source           AccountPageProjectDeploymentListResponseMessagesSource `json:"source"`
	JSON             accountPageProjectDeploymentListResponseMessageJSON    `json:"-"`
}

// accountPageProjectDeploymentListResponseMessageJSON contains the JSON metadata
// for the struct [AccountPageProjectDeploymentListResponseMessage]
type accountPageProjectDeploymentListResponseMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountPageProjectDeploymentListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountPageProjectDeploymentListResponseMessageJSON) RawJSON() string {
	return r.raw
}

type AccountPageProjectDeploymentListResponseMessagesSource struct {
	Pointer string                                                     `json:"pointer"`
	JSON    accountPageProjectDeploymentListResponseMessagesSourceJSON `json:"-"`
}

// accountPageProjectDeploymentListResponseMessagesSourceJSON contains the JSON
// metadata for the struct [AccountPageProjectDeploymentListResponseMessagesSource]
type accountPageProjectDeploymentListResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountPageProjectDeploymentListResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountPageProjectDeploymentListResponseMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AccountPageProjectDeploymentListResponseSuccess bool

const (
	AccountPageProjectDeploymentListResponseSuccessFalse AccountPageProjectDeploymentListResponseSuccess = false
	AccountPageProjectDeploymentListResponseSuccessTrue  AccountPageProjectDeploymentListResponseSuccess = true
)

func (r AccountPageProjectDeploymentListResponseSuccess) IsKnown() bool {
	switch r {
	case AccountPageProjectDeploymentListResponseSuccessFalse, AccountPageProjectDeploymentListResponseSuccessTrue:
		return true
	}
	return false
}

type AccountPageProjectDeploymentListResponseResultInfo struct {
	// The number of items on the current page.
	Count int64 `json:"count,required"`
	// The page currently being requested.
	Page int64 `json:"page,required"`
	// The number of items per page being returned.
	PerPage int64 `json:"per_page,required"`
	// The total count of items.
	TotalCount int64 `json:"total_count,required"`
	// The total count of pages.
	TotalPages int64                                                  `json:"total_pages"`
	JSON       accountPageProjectDeploymentListResponseResultInfoJSON `json:"-"`
}

// accountPageProjectDeploymentListResponseResultInfoJSON contains the JSON
// metadata for the struct [AccountPageProjectDeploymentListResponseResultInfo]
type accountPageProjectDeploymentListResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	TotalPages  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountPageProjectDeploymentListResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountPageProjectDeploymentListResponseResultInfoJSON) RawJSON() string {
	return r.raw
}

type AccountPageProjectDeploymentDeleteResponse struct {
	Errors   []MessagesPageItem `json:"errors,required"`
	Messages []MessagesPageItem `json:"messages,required"`
	Result   interface{}        `json:"result,required,nullable"`
	// Whether the API call was successful
	Success AccountPageProjectDeploymentDeleteResponseSuccess `json:"success,required"`
	JSON    accountPageProjectDeploymentDeleteResponseJSON    `json:"-"`
}

// accountPageProjectDeploymentDeleteResponseJSON contains the JSON metadata for
// the struct [AccountPageProjectDeploymentDeleteResponse]
type accountPageProjectDeploymentDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountPageProjectDeploymentDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountPageProjectDeploymentDeleteResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AccountPageProjectDeploymentDeleteResponseSuccess bool

const (
	AccountPageProjectDeploymentDeleteResponseSuccessFalse AccountPageProjectDeploymentDeleteResponseSuccess = false
	AccountPageProjectDeploymentDeleteResponseSuccessTrue  AccountPageProjectDeploymentDeleteResponseSuccess = true
)

func (r AccountPageProjectDeploymentDeleteResponseSuccess) IsKnown() bool {
	switch r {
	case AccountPageProjectDeploymentDeleteResponseSuccessFalse, AccountPageProjectDeploymentDeleteResponseSuccessTrue:
		return true
	}
	return false
}

type AccountPageProjectDeploymentNewParams struct {
	// The branch to build the new deployment from. The `HEAD` of the branch will be
	// used. If omitted, the production branch will be used by default.
	Branch param.Field[string] `json:"branch"`
}

func (r AccountPageProjectDeploymentNewParams) MarshalMultipart() (data []byte, contentType string, err error) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	err = apiform.MarshalRoot(r, writer)
	if err != nil {
		writer.Close()
		return nil, "", err
	}
	err = writer.Close()
	if err != nil {
		return nil, "", err
	}
	return buf.Bytes(), writer.FormDataContentType(), nil
}

type AccountPageProjectDeploymentListParams struct {
	// What type of deployments to fetch.
	Env param.Field[AccountPageProjectDeploymentListParamsEnv] `query:"env"`
}

// URLQuery serializes [AccountPageProjectDeploymentListParams]'s query parameters
// as `url.Values`.
func (r AccountPageProjectDeploymentListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// What type of deployments to fetch.
type AccountPageProjectDeploymentListParamsEnv string

const (
	AccountPageProjectDeploymentListParamsEnvProduction AccountPageProjectDeploymentListParamsEnv = "production"
	AccountPageProjectDeploymentListParamsEnvPreview    AccountPageProjectDeploymentListParamsEnv = "preview"
)

func (r AccountPageProjectDeploymentListParamsEnv) IsKnown() bool {
	switch r {
	case AccountPageProjectDeploymentListParamsEnvProduction, AccountPageProjectDeploymentListParamsEnvPreview:
		return true
	}
	return false
}

type AccountPageProjectDeploymentRetryParams struct {
	Body interface{} `json:"body,required"`
}

func (r AccountPageProjectDeploymentRetryParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type AccountPageProjectDeploymentRollbackParams struct {
	Body interface{} `json:"body,required"`
}

func (r AccountPageProjectDeploymentRollbackParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}
