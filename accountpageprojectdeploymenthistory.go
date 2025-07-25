// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
)

// AccountPageProjectDeploymentHistoryService contains methods and other services
// that help with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountPageProjectDeploymentHistoryService] method instead.
type AccountPageProjectDeploymentHistoryService struct {
	Options []option.RequestOption
}

// NewAccountPageProjectDeploymentHistoryService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewAccountPageProjectDeploymentHistoryService(opts ...option.RequestOption) (r *AccountPageProjectDeploymentHistoryService) {
	r = &AccountPageProjectDeploymentHistoryService{}
	r.Options = opts
	return
}

// Fetch deployment logs for a project.
func (r *AccountPageProjectDeploymentHistoryService) GetLogs(ctx context.Context, accountID string, projectName string, deploymentID string, opts ...option.RequestOption) (res *AccountPageProjectDeploymentHistoryGetLogsResponse, err error) {
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
	path := fmt.Sprintf("accounts/%s/pages/projects/%s/deployments/%s/history/logs", accountID, projectName, deploymentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AccountPageProjectDeploymentHistoryGetLogsResponse struct {
	Errors   []AccountPageProjectDeploymentHistoryGetLogsResponseError   `json:"errors,required"`
	Messages []AccountPageProjectDeploymentHistoryGetLogsResponseMessage `json:"messages,required"`
	Result   AccountPageProjectDeploymentHistoryGetLogsResponseResult    `json:"result,required"`
	// Whether the API call was successful
	Success AccountPageProjectDeploymentHistoryGetLogsResponseSuccess `json:"success,required"`
	JSON    accountPageProjectDeploymentHistoryGetLogsResponseJSON    `json:"-"`
}

// accountPageProjectDeploymentHistoryGetLogsResponseJSON contains the JSON
// metadata for the struct [AccountPageProjectDeploymentHistoryGetLogsResponse]
type accountPageProjectDeploymentHistoryGetLogsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountPageProjectDeploymentHistoryGetLogsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountPageProjectDeploymentHistoryGetLogsResponseJSON) RawJSON() string {
	return r.raw
}

type AccountPageProjectDeploymentHistoryGetLogsResponseError struct {
	Code             int64                                                          `json:"code,required"`
	Message          string                                                         `json:"message,required"`
	DocumentationURL string                                                         `json:"documentation_url"`
	Source           AccountPageProjectDeploymentHistoryGetLogsResponseErrorsSource `json:"source"`
	JSON             accountPageProjectDeploymentHistoryGetLogsResponseErrorJSON    `json:"-"`
}

// accountPageProjectDeploymentHistoryGetLogsResponseErrorJSON contains the JSON
// metadata for the struct
// [AccountPageProjectDeploymentHistoryGetLogsResponseError]
type accountPageProjectDeploymentHistoryGetLogsResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountPageProjectDeploymentHistoryGetLogsResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountPageProjectDeploymentHistoryGetLogsResponseErrorJSON) RawJSON() string {
	return r.raw
}

type AccountPageProjectDeploymentHistoryGetLogsResponseErrorsSource struct {
	Pointer string                                                             `json:"pointer"`
	JSON    accountPageProjectDeploymentHistoryGetLogsResponseErrorsSourceJSON `json:"-"`
}

// accountPageProjectDeploymentHistoryGetLogsResponseErrorsSourceJSON contains the
// JSON metadata for the struct
// [AccountPageProjectDeploymentHistoryGetLogsResponseErrorsSource]
type accountPageProjectDeploymentHistoryGetLogsResponseErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountPageProjectDeploymentHistoryGetLogsResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountPageProjectDeploymentHistoryGetLogsResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AccountPageProjectDeploymentHistoryGetLogsResponseMessage struct {
	Code             int64                                                            `json:"code,required"`
	Message          string                                                           `json:"message,required"`
	DocumentationURL string                                                           `json:"documentation_url"`
	Source           AccountPageProjectDeploymentHistoryGetLogsResponseMessagesSource `json:"source"`
	JSON             accountPageProjectDeploymentHistoryGetLogsResponseMessageJSON    `json:"-"`
}

// accountPageProjectDeploymentHistoryGetLogsResponseMessageJSON contains the JSON
// metadata for the struct
// [AccountPageProjectDeploymentHistoryGetLogsResponseMessage]
type accountPageProjectDeploymentHistoryGetLogsResponseMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountPageProjectDeploymentHistoryGetLogsResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountPageProjectDeploymentHistoryGetLogsResponseMessageJSON) RawJSON() string {
	return r.raw
}

type AccountPageProjectDeploymentHistoryGetLogsResponseMessagesSource struct {
	Pointer string                                                               `json:"pointer"`
	JSON    accountPageProjectDeploymentHistoryGetLogsResponseMessagesSourceJSON `json:"-"`
}

// accountPageProjectDeploymentHistoryGetLogsResponseMessagesSourceJSON contains
// the JSON metadata for the struct
// [AccountPageProjectDeploymentHistoryGetLogsResponseMessagesSource]
type accountPageProjectDeploymentHistoryGetLogsResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountPageProjectDeploymentHistoryGetLogsResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountPageProjectDeploymentHistoryGetLogsResponseMessagesSourceJSON) RawJSON() string {
	return r.raw
}

type AccountPageProjectDeploymentHistoryGetLogsResponseResult struct {
	Data                  []AccountPageProjectDeploymentHistoryGetLogsResponseResultData `json:"data"`
	IncludesContainerLogs bool                                                           `json:"includes_container_logs"`
	Total                 int64                                                          `json:"total"`
	JSON                  accountPageProjectDeploymentHistoryGetLogsResponseResultJSON   `json:"-"`
}

// accountPageProjectDeploymentHistoryGetLogsResponseResultJSON contains the JSON
// metadata for the struct
// [AccountPageProjectDeploymentHistoryGetLogsResponseResult]
type accountPageProjectDeploymentHistoryGetLogsResponseResultJSON struct {
	Data                  apijson.Field
	IncludesContainerLogs apijson.Field
	Total                 apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *AccountPageProjectDeploymentHistoryGetLogsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountPageProjectDeploymentHistoryGetLogsResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountPageProjectDeploymentHistoryGetLogsResponseResultData struct {
	Line string                                                           `json:"line"`
	Ts   string                                                           `json:"ts"`
	JSON accountPageProjectDeploymentHistoryGetLogsResponseResultDataJSON `json:"-"`
}

// accountPageProjectDeploymentHistoryGetLogsResponseResultDataJSON contains the
// JSON metadata for the struct
// [AccountPageProjectDeploymentHistoryGetLogsResponseResultData]
type accountPageProjectDeploymentHistoryGetLogsResponseResultDataJSON struct {
	Line        apijson.Field
	Ts          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountPageProjectDeploymentHistoryGetLogsResponseResultData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountPageProjectDeploymentHistoryGetLogsResponseResultDataJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AccountPageProjectDeploymentHistoryGetLogsResponseSuccess bool

const (
	AccountPageProjectDeploymentHistoryGetLogsResponseSuccessFalse AccountPageProjectDeploymentHistoryGetLogsResponseSuccess = false
	AccountPageProjectDeploymentHistoryGetLogsResponseSuccessTrue  AccountPageProjectDeploymentHistoryGetLogsResponseSuccess = true
)

func (r AccountPageProjectDeploymentHistoryGetLogsResponseSuccess) IsKnown() bool {
	switch r {
	case AccountPageProjectDeploymentHistoryGetLogsResponseSuccessFalse, AccountPageProjectDeploymentHistoryGetLogsResponseSuccessTrue:
		return true
	}
	return false
}
