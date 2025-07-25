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

// AccountWorkerScriptDeploymentService contains methods and other services that
// help with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountWorkerScriptDeploymentService] method instead.
type AccountWorkerScriptDeploymentService struct {
	Options []option.RequestOption
}

// NewAccountWorkerScriptDeploymentService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountWorkerScriptDeploymentService(opts ...option.RequestOption) (r *AccountWorkerScriptDeploymentService) {
	r = &AccountWorkerScriptDeploymentService{}
	r.Options = opts
	return
}

// Deployments configure how
// [Worker Versions](https://developers.cloudflare.com/api/operations/worker-versions-list-versions)
// are deployed to traffic. A deployment can consist of one or two versions of a
// Worker.
func (r *AccountWorkerScriptDeploymentService) New(ctx context.Context, accountID string, scriptName string, params AccountWorkerScriptDeploymentNewParams, opts ...option.RequestOption) (res *AccountWorkerScriptDeploymentNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if scriptName == "" {
		err = errors.New("missing required script_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s/deployments", accountID, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// List of Worker Deployments. The first deployment in the list is the latest
// deployment actively serving traffic.
func (r *AccountWorkerScriptDeploymentService) List(ctx context.Context, accountID string, scriptName string, opts ...option.RequestOption) (res *AccountWorkerScriptDeploymentListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if scriptName == "" {
		err = errors.New("missing required script_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s/deployments", accountID, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AccountWorkerScriptDeploymentNewResponse struct {
	Errors   []WorkersMessages                              `json:"errors,required"`
	Messages []WorkersMessages                              `json:"messages,required"`
	Result   AccountWorkerScriptDeploymentNewResponseResult `json:"result,required"`
	// Whether the API call was successful.
	Success AccountWorkerScriptDeploymentNewResponseSuccess `json:"success,required"`
	JSON    accountWorkerScriptDeploymentNewResponseJSON    `json:"-"`
}

// accountWorkerScriptDeploymentNewResponseJSON contains the JSON metadata for the
// struct [AccountWorkerScriptDeploymentNewResponse]
type accountWorkerScriptDeploymentNewResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerScriptDeploymentNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountWorkerScriptDeploymentNewResponseJSON) RawJSON() string {
	return r.raw
}

type AccountWorkerScriptDeploymentNewResponseResult struct {
	ID          string                                                    `json:"id,required" format:"uuid"`
	CreatedOn   time.Time                                                 `json:"created_on,required" format:"date-time"`
	Source      string                                                    `json:"source,required"`
	Strategy    AccountWorkerScriptDeploymentNewResponseResultStrategy    `json:"strategy,required"`
	Versions    []AccountWorkerScriptDeploymentNewResponseResultVersion   `json:"versions,required"`
	Annotations AccountWorkerScriptDeploymentNewResponseResultAnnotations `json:"annotations"`
	AuthorEmail string                                                    `json:"author_email" format:"email"`
	JSON        accountWorkerScriptDeploymentNewResponseResultJSON        `json:"-"`
}

// accountWorkerScriptDeploymentNewResponseResultJSON contains the JSON metadata
// for the struct [AccountWorkerScriptDeploymentNewResponseResult]
type accountWorkerScriptDeploymentNewResponseResultJSON struct {
	ID          apijson.Field
	CreatedOn   apijson.Field
	Source      apijson.Field
	Strategy    apijson.Field
	Versions    apijson.Field
	Annotations apijson.Field
	AuthorEmail apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerScriptDeploymentNewResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountWorkerScriptDeploymentNewResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountWorkerScriptDeploymentNewResponseResultStrategy string

const (
	AccountWorkerScriptDeploymentNewResponseResultStrategyPercentage AccountWorkerScriptDeploymentNewResponseResultStrategy = "percentage"
)

func (r AccountWorkerScriptDeploymentNewResponseResultStrategy) IsKnown() bool {
	switch r {
	case AccountWorkerScriptDeploymentNewResponseResultStrategyPercentage:
		return true
	}
	return false
}

type AccountWorkerScriptDeploymentNewResponseResultVersion struct {
	Percentage float64                                                   `json:"percentage,required"`
	VersionID  string                                                    `json:"version_id,required" format:"uuid"`
	JSON       accountWorkerScriptDeploymentNewResponseResultVersionJSON `json:"-"`
}

// accountWorkerScriptDeploymentNewResponseResultVersionJSON contains the JSON
// metadata for the struct [AccountWorkerScriptDeploymentNewResponseResultVersion]
type accountWorkerScriptDeploymentNewResponseResultVersionJSON struct {
	Percentage  apijson.Field
	VersionID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerScriptDeploymentNewResponseResultVersion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountWorkerScriptDeploymentNewResponseResultVersionJSON) RawJSON() string {
	return r.raw
}

type AccountWorkerScriptDeploymentNewResponseResultAnnotations struct {
	// Human-readable message about the deployment. Truncated to 100 bytes.
	WorkersMessage string                                                        `json:"workers/message"`
	JSON           accountWorkerScriptDeploymentNewResponseResultAnnotationsJSON `json:"-"`
}

// accountWorkerScriptDeploymentNewResponseResultAnnotationsJSON contains the JSON
// metadata for the struct
// [AccountWorkerScriptDeploymentNewResponseResultAnnotations]
type accountWorkerScriptDeploymentNewResponseResultAnnotationsJSON struct {
	WorkersMessage apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccountWorkerScriptDeploymentNewResponseResultAnnotations) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountWorkerScriptDeploymentNewResponseResultAnnotationsJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccountWorkerScriptDeploymentNewResponseSuccess bool

const (
	AccountWorkerScriptDeploymentNewResponseSuccessTrue AccountWorkerScriptDeploymentNewResponseSuccess = true
)

func (r AccountWorkerScriptDeploymentNewResponseSuccess) IsKnown() bool {
	switch r {
	case AccountWorkerScriptDeploymentNewResponseSuccessTrue:
		return true
	}
	return false
}

type AccountWorkerScriptDeploymentListResponse struct {
	Errors   []WorkersMessages                               `json:"errors,required"`
	Messages []WorkersMessages                               `json:"messages,required"`
	Result   AccountWorkerScriptDeploymentListResponseResult `json:"result,required"`
	// Whether the API call was successful.
	Success AccountWorkerScriptDeploymentListResponseSuccess `json:"success,required"`
	JSON    accountWorkerScriptDeploymentListResponseJSON    `json:"-"`
}

// accountWorkerScriptDeploymentListResponseJSON contains the JSON metadata for the
// struct [AccountWorkerScriptDeploymentListResponse]
type accountWorkerScriptDeploymentListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerScriptDeploymentListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountWorkerScriptDeploymentListResponseJSON) RawJSON() string {
	return r.raw
}

type AccountWorkerScriptDeploymentListResponseResult struct {
	Deployments []AccountWorkerScriptDeploymentListResponseResultDeployment `json:"deployments,required"`
	JSON        accountWorkerScriptDeploymentListResponseResultJSON         `json:"-"`
}

// accountWorkerScriptDeploymentListResponseResultJSON contains the JSON metadata
// for the struct [AccountWorkerScriptDeploymentListResponseResult]
type accountWorkerScriptDeploymentListResponseResultJSON struct {
	Deployments apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerScriptDeploymentListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountWorkerScriptDeploymentListResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountWorkerScriptDeploymentListResponseResultDeployment struct {
	ID          string                                                                `json:"id,required" format:"uuid"`
	CreatedOn   time.Time                                                             `json:"created_on,required" format:"date-time"`
	Source      string                                                                `json:"source,required"`
	Strategy    AccountWorkerScriptDeploymentListResponseResultDeploymentsStrategy    `json:"strategy,required"`
	Versions    []AccountWorkerScriptDeploymentListResponseResultDeploymentsVersion   `json:"versions,required"`
	Annotations AccountWorkerScriptDeploymentListResponseResultDeploymentsAnnotations `json:"annotations"`
	AuthorEmail string                                                                `json:"author_email" format:"email"`
	JSON        accountWorkerScriptDeploymentListResponseResultDeploymentJSON         `json:"-"`
}

// accountWorkerScriptDeploymentListResponseResultDeploymentJSON contains the JSON
// metadata for the struct
// [AccountWorkerScriptDeploymentListResponseResultDeployment]
type accountWorkerScriptDeploymentListResponseResultDeploymentJSON struct {
	ID          apijson.Field
	CreatedOn   apijson.Field
	Source      apijson.Field
	Strategy    apijson.Field
	Versions    apijson.Field
	Annotations apijson.Field
	AuthorEmail apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerScriptDeploymentListResponseResultDeployment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountWorkerScriptDeploymentListResponseResultDeploymentJSON) RawJSON() string {
	return r.raw
}

type AccountWorkerScriptDeploymentListResponseResultDeploymentsStrategy string

const (
	AccountWorkerScriptDeploymentListResponseResultDeploymentsStrategyPercentage AccountWorkerScriptDeploymentListResponseResultDeploymentsStrategy = "percentage"
)

func (r AccountWorkerScriptDeploymentListResponseResultDeploymentsStrategy) IsKnown() bool {
	switch r {
	case AccountWorkerScriptDeploymentListResponseResultDeploymentsStrategyPercentage:
		return true
	}
	return false
}

type AccountWorkerScriptDeploymentListResponseResultDeploymentsVersion struct {
	Percentage float64                                                               `json:"percentage,required"`
	VersionID  string                                                                `json:"version_id,required" format:"uuid"`
	JSON       accountWorkerScriptDeploymentListResponseResultDeploymentsVersionJSON `json:"-"`
}

// accountWorkerScriptDeploymentListResponseResultDeploymentsVersionJSON contains
// the JSON metadata for the struct
// [AccountWorkerScriptDeploymentListResponseResultDeploymentsVersion]
type accountWorkerScriptDeploymentListResponseResultDeploymentsVersionJSON struct {
	Percentage  apijson.Field
	VersionID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerScriptDeploymentListResponseResultDeploymentsVersion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountWorkerScriptDeploymentListResponseResultDeploymentsVersionJSON) RawJSON() string {
	return r.raw
}

type AccountWorkerScriptDeploymentListResponseResultDeploymentsAnnotations struct {
	// Human-readable message about the deployment. Truncated to 100 bytes.
	WorkersMessage string                                                                    `json:"workers/message"`
	JSON           accountWorkerScriptDeploymentListResponseResultDeploymentsAnnotationsJSON `json:"-"`
}

// accountWorkerScriptDeploymentListResponseResultDeploymentsAnnotationsJSON
// contains the JSON metadata for the struct
// [AccountWorkerScriptDeploymentListResponseResultDeploymentsAnnotations]
type accountWorkerScriptDeploymentListResponseResultDeploymentsAnnotationsJSON struct {
	WorkersMessage apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccountWorkerScriptDeploymentListResponseResultDeploymentsAnnotations) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountWorkerScriptDeploymentListResponseResultDeploymentsAnnotationsJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccountWorkerScriptDeploymentListResponseSuccess bool

const (
	AccountWorkerScriptDeploymentListResponseSuccessTrue AccountWorkerScriptDeploymentListResponseSuccess = true
)

func (r AccountWorkerScriptDeploymentListResponseSuccess) IsKnown() bool {
	switch r {
	case AccountWorkerScriptDeploymentListResponseSuccessTrue:
		return true
	}
	return false
}

type AccountWorkerScriptDeploymentNewParams struct {
	Strategy param.Field[AccountWorkerScriptDeploymentNewParamsStrategy]  `json:"strategy,required"`
	Versions param.Field[[]AccountWorkerScriptDeploymentNewParamsVersion] `json:"versions,required"`
	// If set to true, the deployment will be created even if normally blocked by
	// something such rolling back to an older version when a secret has changed.
	Force       param.Field[bool]                                              `query:"force"`
	Annotations param.Field[AccountWorkerScriptDeploymentNewParamsAnnotations] `json:"annotations"`
}

func (r AccountWorkerScriptDeploymentNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [AccountWorkerScriptDeploymentNewParams]'s query parameters
// as `url.Values`.
func (r AccountWorkerScriptDeploymentNewParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountWorkerScriptDeploymentNewParamsStrategy string

const (
	AccountWorkerScriptDeploymentNewParamsStrategyPercentage AccountWorkerScriptDeploymentNewParamsStrategy = "percentage"
)

func (r AccountWorkerScriptDeploymentNewParamsStrategy) IsKnown() bool {
	switch r {
	case AccountWorkerScriptDeploymentNewParamsStrategyPercentage:
		return true
	}
	return false
}

type AccountWorkerScriptDeploymentNewParamsVersion struct {
	Percentage param.Field[float64] `json:"percentage,required"`
	VersionID  param.Field[string]  `json:"version_id,required" format:"uuid"`
}

func (r AccountWorkerScriptDeploymentNewParamsVersion) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountWorkerScriptDeploymentNewParamsAnnotations struct {
	// Human-readable message about the deployment. Truncated to 100 bytes.
	WorkersMessage param.Field[string] `json:"workers/message"`
}

func (r AccountWorkerScriptDeploymentNewParamsAnnotations) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
