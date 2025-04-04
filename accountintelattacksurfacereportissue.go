// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/stainless-sdks/cf-rex-go/internal/apijson"
	"github.com/stainless-sdks/cf-rex-go/internal/apiquery"
	"github.com/stainless-sdks/cf-rex-go/internal/param"
	"github.com/stainless-sdks/cf-rex-go/internal/requestconfig"
	"github.com/stainless-sdks/cf-rex-go/option"
)

// AccountIntelAttackSurfaceReportIssueService contains methods and other services
// that help with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountIntelAttackSurfaceReportIssueService] method instead.
type AccountIntelAttackSurfaceReportIssueService struct {
	Options []option.RequestOption
}

// NewAccountIntelAttackSurfaceReportIssueService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewAccountIntelAttackSurfaceReportIssueService(opts ...option.RequestOption) (r *AccountIntelAttackSurfaceReportIssueService) {
	r = &AccountIntelAttackSurfaceReportIssueService{}
	r.Options = opts
	return
}

// Get Security Center Issues
func (r *AccountIntelAttackSurfaceReportIssueService) List(ctx context.Context, accountID string, query AccountIntelAttackSurfaceReportIssueListParams, opts ...option.RequestOption) (res *AccountIntelAttackSurfaceReportIssueListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/intel/attack-surface-report/issues", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Get Security Center Issue Counts by Class
func (r *AccountIntelAttackSurfaceReportIssueService) ListByClass(ctx context.Context, accountID string, query AccountIntelAttackSurfaceReportIssueListByClassParams, opts ...option.RequestOption) (res *ValueCountsResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/intel/attack-surface-report/issues/class", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Get Security Center Issue Counts by Severity
func (r *AccountIntelAttackSurfaceReportIssueService) ListBySeverity(ctx context.Context, accountID string, query AccountIntelAttackSurfaceReportIssueListBySeverityParams, opts ...option.RequestOption) (res *ValueCountsResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/intel/attack-surface-report/issues/severity", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Get Security Center Issue Counts by Type
func (r *AccountIntelAttackSurfaceReportIssueService) ListByType(ctx context.Context, accountID string, query AccountIntelAttackSurfaceReportIssueListByTypeParams, opts ...option.RequestOption) (res *ValueCountsResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/intel/attack-surface-report/issues/type", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type Issue struct {
	ID          string        `json:"id"`
	Dismissed   bool          `json:"dismissed"`
	IssueClass  string        `json:"issue_class"`
	IssueType   IssueType     `json:"issue_type"`
	Payload     interface{}   `json:"payload"`
	ResolveLink string        `json:"resolve_link"`
	ResolveText string        `json:"resolve_text"`
	Severity    IssueSeverity `json:"severity"`
	Since       time.Time     `json:"since" format:"date-time"`
	Subject     string        `json:"subject"`
	Timestamp   time.Time     `json:"timestamp" format:"date-time"`
	JSON        issueJSON     `json:"-"`
}

// issueJSON contains the JSON metadata for the struct [Issue]
type issueJSON struct {
	ID          apijson.Field
	Dismissed   apijson.Field
	IssueClass  apijson.Field
	IssueType   apijson.Field
	Payload     apijson.Field
	ResolveLink apijson.Field
	ResolveText apijson.Field
	Severity    apijson.Field
	Since       apijson.Field
	Subject     apijson.Field
	Timestamp   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Issue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r issueJSON) RawJSON() string {
	return r.raw
}

type IssueSeverity string

const (
	IssueSeverityLow      IssueSeverity = "Low"
	IssueSeverityModerate IssueSeverity = "Moderate"
	IssueSeverityCritical IssueSeverity = "Critical"
)

func (r IssueSeverity) IsKnown() bool {
	switch r {
	case IssueSeverityLow, IssueSeverityModerate, IssueSeverityCritical:
		return true
	}
	return false
}

type IssueType string

const (
	IssueTypeComplianceViolation   IssueType = "compliance_violation"
	IssueTypeEmailSecurity         IssueType = "email_security"
	IssueTypeExposedInfrastructure IssueType = "exposed_infrastructure"
	IssueTypeInsecureConfiguration IssueType = "insecure_configuration"
	IssueTypeWeakAuthentication    IssueType = "weak_authentication"
)

func (r IssueType) IsKnown() bool {
	switch r {
	case IssueTypeComplianceViolation, IssueTypeEmailSecurity, IssueTypeExposedInfrastructure, IssueTypeInsecureConfiguration, IssueTypeWeakAuthentication:
		return true
	}
	return false
}

type SeverityQueryParam string

const (
	SeverityQueryParamLow      SeverityQueryParam = "low"
	SeverityQueryParamModerate SeverityQueryParam = "moderate"
	SeverityQueryParamCritical SeverityQueryParam = "critical"
)

func (r SeverityQueryParam) IsKnown() bool {
	switch r {
	case SeverityQueryParamLow, SeverityQueryParamModerate, SeverityQueryParamCritical:
		return true
	}
	return false
}

type ValueCountsResponse struct {
	Result []ValueCountsResponseResult `json:"result"`
	JSON   valueCountsResponseJSON     `json:"-"`
	CommonResponseAttackSurfaceReport
}

// valueCountsResponseJSON contains the JSON metadata for the struct
// [ValueCountsResponse]
type valueCountsResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ValueCountsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r valueCountsResponseJSON) RawJSON() string {
	return r.raw
}

type ValueCountsResponseResult struct {
	Count int64                         `json:"count"`
	Value string                        `json:"value"`
	JSON  valueCountsResponseResultJSON `json:"-"`
}

// valueCountsResponseResultJSON contains the JSON metadata for the struct
// [ValueCountsResponseResult]
type valueCountsResponseResultJSON struct {
	Count       apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ValueCountsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r valueCountsResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountIntelAttackSurfaceReportIssueListResponse struct {
	Result AccountIntelAttackSurfaceReportIssueListResponseResult `json:"result"`
	JSON   accountIntelAttackSurfaceReportIssueListResponseJSON   `json:"-"`
	CommonResponseAttackSurfaceReport
}

// accountIntelAttackSurfaceReportIssueListResponseJSON contains the JSON metadata
// for the struct [AccountIntelAttackSurfaceReportIssueListResponse]
type accountIntelAttackSurfaceReportIssueListResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountIntelAttackSurfaceReportIssueListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountIntelAttackSurfaceReportIssueListResponseJSON) RawJSON() string {
	return r.raw
}

type AccountIntelAttackSurfaceReportIssueListResponseResult struct {
	// Total number of results
	Count  int64   `json:"count"`
	Issues []Issue `json:"issues"`
	// Current page within paginated list of results
	Page int64 `json:"page"`
	// Number of results per page of results
	PerPage int64                                                      `json:"per_page"`
	JSON    accountIntelAttackSurfaceReportIssueListResponseResultJSON `json:"-"`
}

// accountIntelAttackSurfaceReportIssueListResponseResultJSON contains the JSON
// metadata for the struct [AccountIntelAttackSurfaceReportIssueListResponseResult]
type accountIntelAttackSurfaceReportIssueListResponseResultJSON struct {
	Count       apijson.Field
	Issues      apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountIntelAttackSurfaceReportIssueListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountIntelAttackSurfaceReportIssueListResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountIntelAttackSurfaceReportIssueListParams struct {
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

// URLQuery serializes [AccountIntelAttackSurfaceReportIssueListParams]'s query
// parameters as `url.Values`.
func (r AccountIntelAttackSurfaceReportIssueListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountIntelAttackSurfaceReportIssueListByClassParams struct {
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

// URLQuery serializes [AccountIntelAttackSurfaceReportIssueListByClassParams]'s
// query parameters as `url.Values`.
func (r AccountIntelAttackSurfaceReportIssueListByClassParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountIntelAttackSurfaceReportIssueListBySeverityParams struct {
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

// URLQuery serializes [AccountIntelAttackSurfaceReportIssueListBySeverityParams]'s
// query parameters as `url.Values`.
func (r AccountIntelAttackSurfaceReportIssueListBySeverityParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountIntelAttackSurfaceReportIssueListByTypeParams struct {
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

// URLQuery serializes [AccountIntelAttackSurfaceReportIssueListByTypeParams]'s
// query parameters as `url.Values`.
func (r AccountIntelAttackSurfaceReportIssueListByTypeParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
