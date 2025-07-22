// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/apiquery"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
)

// AccountSecurityCenterInsightService contains methods and other services that
// help with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountSecurityCenterInsightService] method instead.
type AccountSecurityCenterInsightService struct {
	Options []option.RequestOption
}

// NewAccountSecurityCenterInsightService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountSecurityCenterInsightService(opts ...option.RequestOption) (r *AccountSecurityCenterInsightService) {
	r = &AccountSecurityCenterInsightService{}
	r.Options = opts
	return
}

// Get Security Center Insights
func (r *AccountSecurityCenterInsightService) Get(ctx context.Context, accountID string, query AccountSecurityCenterInsightGetParams, opts ...option.RequestOption) (res *AccountSecurityCenterInsightGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/security-center/insights", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Archive Security Center Insight
func (r *AccountSecurityCenterInsightService) Dismiss(ctx context.Context, accountID string, issueID string, body AccountSecurityCenterInsightDismissParams, opts ...option.RequestOption) (res *SingleResponseReport, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if issueID == "" {
		err = errors.New("missing required issue_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/security-center/insights/%s/dismiss", accountID, issueID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Get Security Center Insight Counts by Class
func (r *AccountSecurityCenterInsightService) ListByClass(ctx context.Context, accountID string, query AccountSecurityCenterInsightListByClassParams, opts ...option.RequestOption) (res *ValueCountsResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/security-center/insights/class", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Get Security Center Insight Counts by Severity
func (r *AccountSecurityCenterInsightService) ListBySeverity(ctx context.Context, accountID string, query AccountSecurityCenterInsightListBySeverityParams, opts ...option.RequestOption) (res *ValueCountsResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/security-center/insights/severity", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Get Security Center Insight Counts by Type
func (r *AccountSecurityCenterInsightService) ListByType(ctx context.Context, accountID string, query AccountSecurityCenterInsightListByTypeParams, opts ...option.RequestOption) (res *ValueCountsResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/security-center/insights/type", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type AccountSecurityCenterInsightGetResponse struct {
	Result AccountSecurityCenterInsightGetResponseResult `json:"result"`
	JSON   accountSecurityCenterInsightGetResponseJSON   `json:"-"`
	CommonResponseAttackSurfaceReport
}

// accountSecurityCenterInsightGetResponseJSON contains the JSON metadata for the
// struct [AccountSecurityCenterInsightGetResponse]
type accountSecurityCenterInsightGetResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSecurityCenterInsightGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountSecurityCenterInsightGetResponseJSON) RawJSON() string {
	return r.raw
}

type AccountSecurityCenterInsightGetResponseResult struct {
	// Total number of results
	Count  int64   `json:"count"`
	Issues []Issue `json:"issues"`
	// Current page within paginated list of results
	Page int64 `json:"page"`
	// Number of results per page of results
	PerPage int64                                             `json:"per_page"`
	JSON    accountSecurityCenterInsightGetResponseResultJSON `json:"-"`
}

// accountSecurityCenterInsightGetResponseResultJSON contains the JSON metadata for
// the struct [AccountSecurityCenterInsightGetResponseResult]
type accountSecurityCenterInsightGetResponseResultJSON struct {
	Count       apijson.Field
	Issues      apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSecurityCenterInsightGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountSecurityCenterInsightGetResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountSecurityCenterInsightGetParams struct {
	Dismissed     param.Field[bool]        `query:"dismissed"`
	IssueClass    param.Field[[]string]    `query:"issue_class"`
	IssueClassNeq param.Field[[]string]    `query:"issue_class~neq"`
	IssueType     param.Field[[]IssueType] `query:"issue_type"`
	IssueTypeNeq  param.Field[[]IssueType] `query:"issue_type~neq"`
	// Current page within paginated list of results
	Page param.Field[int64] `query:"page"`
	// Number of results per page of results
	PerPage     param.Field[int64]                `query:"per_page"`
	Product     param.Field[[]string]             `query:"product"`
	ProductNeq  param.Field[[]string]             `query:"product~neq"`
	Severity    param.Field[[]SeverityQueryParam] `query:"severity"`
	SeverityNeq param.Field[[]SeverityQueryParam] `query:"severity~neq"`
	Subject     param.Field[[]string]             `query:"subject"`
	SubjectNeq  param.Field[[]string]             `query:"subject~neq"`
}

// URLQuery serializes [AccountSecurityCenterInsightGetParams]'s query parameters
// as `url.Values`.
func (r AccountSecurityCenterInsightGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountSecurityCenterInsightDismissParams struct {
	Dismiss param.Field[bool] `json:"dismiss"`
}

func (r AccountSecurityCenterInsightDismissParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountSecurityCenterInsightListByClassParams struct {
	Dismissed     param.Field[bool]                 `query:"dismissed"`
	IssueClass    param.Field[[]string]             `query:"issue_class"`
	IssueClassNeq param.Field[[]string]             `query:"issue_class~neq"`
	IssueType     param.Field[[]IssueType]          `query:"issue_type"`
	IssueTypeNeq  param.Field[[]IssueType]          `query:"issue_type~neq"`
	Product       param.Field[[]string]             `query:"product"`
	ProductNeq    param.Field[[]string]             `query:"product~neq"`
	Severity      param.Field[[]SeverityQueryParam] `query:"severity"`
	SeverityNeq   param.Field[[]SeverityQueryParam] `query:"severity~neq"`
	Subject       param.Field[[]string]             `query:"subject"`
	SubjectNeq    param.Field[[]string]             `query:"subject~neq"`
}

// URLQuery serializes [AccountSecurityCenterInsightListByClassParams]'s query
// parameters as `url.Values`.
func (r AccountSecurityCenterInsightListByClassParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountSecurityCenterInsightListBySeverityParams struct {
	Dismissed     param.Field[bool]                 `query:"dismissed"`
	IssueClass    param.Field[[]string]             `query:"issue_class"`
	IssueClassNeq param.Field[[]string]             `query:"issue_class~neq"`
	IssueType     param.Field[[]IssueType]          `query:"issue_type"`
	IssueTypeNeq  param.Field[[]IssueType]          `query:"issue_type~neq"`
	Product       param.Field[[]string]             `query:"product"`
	ProductNeq    param.Field[[]string]             `query:"product~neq"`
	Severity      param.Field[[]SeverityQueryParam] `query:"severity"`
	SeverityNeq   param.Field[[]SeverityQueryParam] `query:"severity~neq"`
	Subject       param.Field[[]string]             `query:"subject"`
	SubjectNeq    param.Field[[]string]             `query:"subject~neq"`
}

// URLQuery serializes [AccountSecurityCenterInsightListBySeverityParams]'s query
// parameters as `url.Values`.
func (r AccountSecurityCenterInsightListBySeverityParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountSecurityCenterInsightListByTypeParams struct {
	Dismissed     param.Field[bool]                 `query:"dismissed"`
	IssueClass    param.Field[[]string]             `query:"issue_class"`
	IssueClassNeq param.Field[[]string]             `query:"issue_class~neq"`
	IssueType     param.Field[[]IssueType]          `query:"issue_type"`
	IssueTypeNeq  param.Field[[]IssueType]          `query:"issue_type~neq"`
	Product       param.Field[[]string]             `query:"product"`
	ProductNeq    param.Field[[]string]             `query:"product~neq"`
	Severity      param.Field[[]SeverityQueryParam] `query:"severity"`
	SeverityNeq   param.Field[[]SeverityQueryParam] `query:"severity~neq"`
	Subject       param.Field[[]string]             `query:"subject"`
	SubjectNeq    param.Field[[]string]             `query:"subject~neq"`
}

// URLQuery serializes [AccountSecurityCenterInsightListByTypeParams]'s query
// parameters as `url.Values`.
func (r AccountSecurityCenterInsightListByTypeParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
