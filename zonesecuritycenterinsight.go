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

// ZoneSecurityCenterInsightService contains methods and other services that help
// with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneSecurityCenterInsightService] method instead.
type ZoneSecurityCenterInsightService struct {
	Options []option.RequestOption
}

// NewZoneSecurityCenterInsightService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZoneSecurityCenterInsightService(opts ...option.RequestOption) (r *ZoneSecurityCenterInsightService) {
	r = &ZoneSecurityCenterInsightService{}
	r.Options = opts
	return
}

// Get Zone Security Center Insights
func (r *ZoneSecurityCenterInsightService) Get(ctx context.Context, zoneID string, query ZoneSecurityCenterInsightGetParams, opts ...option.RequestOption) (res *ZoneSecurityCenterInsightGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/security-center/insights", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Get Zone Security Center Insight Counts by Class
func (r *ZoneSecurityCenterInsightService) ClassCounts(ctx context.Context, zoneID string, query ZoneSecurityCenterInsightClassCountsParams, opts ...option.RequestOption) (res *ValueCountsResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/security-center/insights/class", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Archive Zone Security Center Insight
func (r *ZoneSecurityCenterInsightService) Dismiss(ctx context.Context, zoneID string, issueID string, body ZoneSecurityCenterInsightDismissParams, opts ...option.RequestOption) (res *SingleResponseReport, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if issueID == "" {
		err = errors.New("missing required issue_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/security-center/insights/%s/dismiss", zoneID, issueID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Get Zone Security Center Insight Counts by Severity
func (r *ZoneSecurityCenterInsightService) SeverityCounts(ctx context.Context, zoneID string, query ZoneSecurityCenterInsightSeverityCountsParams, opts ...option.RequestOption) (res *ValueCountsResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/security-center/insights/severity", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Get Zone Security Center Insight Counts by Type
func (r *ZoneSecurityCenterInsightService) TypeCounts(ctx context.Context, zoneID string, query ZoneSecurityCenterInsightTypeCountsParams, opts ...option.RequestOption) (res *ValueCountsResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/security-center/insights/type", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type ZoneSecurityCenterInsightGetResponse struct {
	Result ZoneSecurityCenterInsightGetResponseResult `json:"result"`
	JSON   zoneSecurityCenterInsightGetResponseJSON   `json:"-"`
	CommonResponseAttackSurfaceReport
}

// zoneSecurityCenterInsightGetResponseJSON contains the JSON metadata for the
// struct [ZoneSecurityCenterInsightGetResponse]
type zoneSecurityCenterInsightGetResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSecurityCenterInsightGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSecurityCenterInsightGetResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneSecurityCenterInsightGetResponseResult struct {
	// Total number of results
	Count  int64   `json:"count"`
	Issues []Issue `json:"issues"`
	// Current page within paginated list of results
	Page int64 `json:"page"`
	// Number of results per page of results
	PerPage int64                                          `json:"per_page"`
	JSON    zoneSecurityCenterInsightGetResponseResultJSON `json:"-"`
}

// zoneSecurityCenterInsightGetResponseResultJSON contains the JSON metadata for
// the struct [ZoneSecurityCenterInsightGetResponseResult]
type zoneSecurityCenterInsightGetResponseResultJSON struct {
	Count       apijson.Field
	Issues      apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSecurityCenterInsightGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSecurityCenterInsightGetResponseResultJSON) RawJSON() string {
	return r.raw
}

type ZoneSecurityCenterInsightGetParams struct {
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

// URLQuery serializes [ZoneSecurityCenterInsightGetParams]'s query parameters as
// `url.Values`.
func (r ZoneSecurityCenterInsightGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ZoneSecurityCenterInsightClassCountsParams struct {
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

// URLQuery serializes [ZoneSecurityCenterInsightClassCountsParams]'s query
// parameters as `url.Values`.
func (r ZoneSecurityCenterInsightClassCountsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ZoneSecurityCenterInsightDismissParams struct {
	Dismiss param.Field[bool] `json:"dismiss"`
}

func (r ZoneSecurityCenterInsightDismissParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneSecurityCenterInsightSeverityCountsParams struct {
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

// URLQuery serializes [ZoneSecurityCenterInsightSeverityCountsParams]'s query
// parameters as `url.Values`.
func (r ZoneSecurityCenterInsightSeverityCountsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ZoneSecurityCenterInsightTypeCountsParams struct {
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

// URLQuery serializes [ZoneSecurityCenterInsightTypeCountsParams]'s query
// parameters as `url.Values`.
func (r ZoneSecurityCenterInsightTypeCountsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
