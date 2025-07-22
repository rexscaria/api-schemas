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

	"github.com/stainless-sdks/cf-rex-go/internal/apiform"
	"github.com/stainless-sdks/cf-rex-go/internal/apijson"
	"github.com/stainless-sdks/cf-rex-go/internal/apiquery"
	"github.com/stainless-sdks/cf-rex-go/internal/param"
	"github.com/stainless-sdks/cf-rex-go/internal/requestconfig"
	"github.com/stainless-sdks/cf-rex-go/option"
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
func (r *AccountPageProjectDeploymentService) Delete(ctx context.Context, accountID string, projectName string, deploymentID string, body AccountPageProjectDeploymentDeleteParams, opts ...option.RequestOption) (res *AccountPageProjectDeploymentDeleteResponse, err error) {
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
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
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
	Result Deployments       `json:"result"`
	JSON   newDeploymentJSON `json:"-"`
	APIResponsePages
}

// newDeploymentJSON contains the JSON metadata for the struct [NewDeployment]
type newDeploymentJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NewDeployment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r newDeploymentJSON) RawJSON() string {
	return r.raw
}

type ResponseDetails struct {
	Result Deployments         `json:"result"`
	JSON   responseDetailsJSON `json:"-"`
	APIResponsePages
}

// responseDetailsJSON contains the JSON metadata for the struct [ResponseDetails]
type responseDetailsJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResponseDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r responseDetailsJSON) RawJSON() string {
	return r.raw
}

type AccountPageProjectDeploymentListResponse struct {
	Result []Deployments                                `json:"result"`
	JSON   accountPageProjectDeploymentListResponseJSON `json:"-"`
	APIResponsePages
	APIResponsePagination
}

// accountPageProjectDeploymentListResponseJSON contains the JSON metadata for the
// struct [AccountPageProjectDeploymentListResponse]
type accountPageProjectDeploymentListResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountPageProjectDeploymentListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountPageProjectDeploymentListResponseJSON) RawJSON() string {
	return r.raw
}

type AccountPageProjectDeploymentDeleteResponse struct {
	Result interface{}                                    `json:"result,nullable"`
	JSON   accountPageProjectDeploymentDeleteResponseJSON `json:"-"`
	APIResponsePages
}

// accountPageProjectDeploymentDeleteResponseJSON contains the JSON metadata for
// the struct [AccountPageProjectDeploymentDeleteResponse]
type accountPageProjectDeploymentDeleteResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountPageProjectDeploymentDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountPageProjectDeploymentDeleteResponseJSON) RawJSON() string {
	return r.raw
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

type AccountPageProjectDeploymentDeleteParams struct {
	Body interface{} `json:"body,required"`
}

func (r AccountPageProjectDeploymentDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
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
