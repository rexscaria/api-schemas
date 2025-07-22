// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/stainless-sdks/cf-rex-go/internal/apijson"
	"github.com/stainless-sdks/cf-rex-go/internal/apiquery"
	"github.com/stainless-sdks/cf-rex-go/internal/param"
	"github.com/stainless-sdks/cf-rex-go/internal/requestconfig"
	"github.com/stainless-sdks/cf-rex-go/option"
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

type BaseDeployment struct {
	ID          string                    `json:"id"`
	Annotations BaseDeploymentAnnotations `json:"annotations"`
	AuthorEmail string                    `json:"author_email"`
	CreatedOn   string                    `json:"created_on"`
	Source      string                    `json:"source"`
	Strategy    string                    `json:"strategy"`
	JSON        baseDeploymentJSON        `json:"-"`
}

// baseDeploymentJSON contains the JSON metadata for the struct [BaseDeployment]
type baseDeploymentJSON struct {
	ID          apijson.Field
	Annotations apijson.Field
	AuthorEmail apijson.Field
	CreatedOn   apijson.Field
	Source      apijson.Field
	Strategy    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BaseDeployment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r baseDeploymentJSON) RawJSON() string {
	return r.raw
}

type BaseDeploymentAnnotations struct {
	// Human-readable message about the deployment. Truncated to 100 bytes.
	WorkersMessage string                        `json:"workers/message"`
	JSON           baseDeploymentAnnotationsJSON `json:"-"`
}

// baseDeploymentAnnotationsJSON contains the JSON metadata for the struct
// [BaseDeploymentAnnotations]
type baseDeploymentAnnotationsJSON struct {
	WorkersMessage apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *BaseDeploymentAnnotations) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r baseDeploymentAnnotationsJSON) RawJSON() string {
	return r.raw
}

type StrategyPercentage struct {
	Strategy StrategyPercentageStrategy  `json:"strategy,required"`
	Versions []StrategyPercentageVersion `json:"versions,required"`
	JSON     strategyPercentageJSON      `json:"-"`
}

// strategyPercentageJSON contains the JSON metadata for the struct
// [StrategyPercentage]
type strategyPercentageJSON struct {
	Strategy    apijson.Field
	Versions    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StrategyPercentage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r strategyPercentageJSON) RawJSON() string {
	return r.raw
}

type StrategyPercentageStrategy string

const (
	StrategyPercentageStrategyPercentage StrategyPercentageStrategy = "percentage"
)

func (r StrategyPercentageStrategy) IsKnown() bool {
	switch r {
	case StrategyPercentageStrategyPercentage:
		return true
	}
	return false
}

type StrategyPercentageVersion struct {
	Percentage float64                       `json:"percentage,required"`
	VersionID  string                        `json:"version_id,required"`
	JSON       strategyPercentageVersionJSON `json:"-"`
}

// strategyPercentageVersionJSON contains the JSON metadata for the struct
// [StrategyPercentageVersion]
type strategyPercentageVersionJSON struct {
	Percentage  apijson.Field
	VersionID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StrategyPercentageVersion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r strategyPercentageVersionJSON) RawJSON() string {
	return r.raw
}

type AccountWorkerScriptDeploymentNewResponse struct {
	Result AccountWorkerScriptDeploymentNewResponseResult `json:"result"`
	JSON   accountWorkerScriptDeploymentNewResponseJSON   `json:"-"`
	CommonResponseWorkers
}

// accountWorkerScriptDeploymentNewResponseJSON contains the JSON metadata for the
// struct [AccountWorkerScriptDeploymentNewResponse]
type accountWorkerScriptDeploymentNewResponseJSON struct {
	Result      apijson.Field
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
	JSON accountWorkerScriptDeploymentNewResponseResultJSON `json:"-"`
	BaseDeployment
	StrategyPercentage
}

// accountWorkerScriptDeploymentNewResponseResultJSON contains the JSON metadata
// for the struct [AccountWorkerScriptDeploymentNewResponseResult]
type accountWorkerScriptDeploymentNewResponseResultJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerScriptDeploymentNewResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountWorkerScriptDeploymentNewResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountWorkerScriptDeploymentListResponse struct {
	Result AccountWorkerScriptDeploymentListResponseResult `json:"result"`
	JSON   accountWorkerScriptDeploymentListResponseJSON   `json:"-"`
	CommonResponseWorkers
}

// accountWorkerScriptDeploymentListResponseJSON contains the JSON metadata for the
// struct [AccountWorkerScriptDeploymentListResponse]
type accountWorkerScriptDeploymentListResponseJSON struct {
	Result      apijson.Field
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
	Deployments []AccountWorkerScriptDeploymentListResponseResultDeployment `json:"deployments"`
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
	JSON accountWorkerScriptDeploymentListResponseResultDeploymentJSON `json:"-"`
	BaseDeployment
	StrategyPercentage
}

// accountWorkerScriptDeploymentListResponseResultDeploymentJSON contains the JSON
// metadata for the struct
// [AccountWorkerScriptDeploymentListResponseResultDeployment]
type accountWorkerScriptDeploymentListResponseResultDeploymentJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerScriptDeploymentListResponseResultDeployment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountWorkerScriptDeploymentListResponseResultDeploymentJSON) RawJSON() string {
	return r.raw
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
	VersionID  param.Field[string]  `json:"version_id,required"`
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
