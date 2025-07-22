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
	Result AccountPageProjectDeploymentHistoryGetLogsResponseResult `json:"result"`
	JSON   accountPageProjectDeploymentHistoryGetLogsResponseJSON   `json:"-"`
	APIResponsePages
}

// accountPageProjectDeploymentHistoryGetLogsResponseJSON contains the JSON
// metadata for the struct [AccountPageProjectDeploymentHistoryGetLogsResponse]
type accountPageProjectDeploymentHistoryGetLogsResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountPageProjectDeploymentHistoryGetLogsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountPageProjectDeploymentHistoryGetLogsResponseJSON) RawJSON() string {
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
