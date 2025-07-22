// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
)

// AccountIntelAttackSurfaceReportService contains methods and other services that
// help with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountIntelAttackSurfaceReportService] method instead.
type AccountIntelAttackSurfaceReportService struct {
	Options []option.RequestOption
	Issues  *AccountIntelAttackSurfaceReportIssueService
}

// NewAccountIntelAttackSurfaceReportService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountIntelAttackSurfaceReportService(opts ...option.RequestOption) (r *AccountIntelAttackSurfaceReportService) {
	r = &AccountIntelAttackSurfaceReportService{}
	r.Options = opts
	r.Issues = NewAccountIntelAttackSurfaceReportIssueService(opts...)
	return
}

// Archive Security Center Insight
//
// Deprecated: deprecated
func (r *AccountIntelAttackSurfaceReportService) DismissIssue(ctx context.Context, accountID string, issueID string, body AccountIntelAttackSurfaceReportDismissIssueParams, opts ...option.RequestOption) (res *SingleResponseReport, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if issueID == "" {
		err = errors.New("missing required issue_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/intel/attack-surface-report/%s/dismiss", accountID, issueID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Get Security Center Issues Types
func (r *AccountIntelAttackSurfaceReportService) ListIssueTypes(ctx context.Context, accountID string, opts ...option.RequestOption) (res *AccountIntelAttackSurfaceReportListIssueTypesResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/intel/attack-surface-report/issue-types", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AttackSurfaceReportMessage struct {
	Code    int64                          `json:"code,required"`
	Message string                         `json:"message,required"`
	JSON    attackSurfaceReportMessageJSON `json:"-"`
}

// attackSurfaceReportMessageJSON contains the JSON metadata for the struct
// [AttackSurfaceReportMessage]
type attackSurfaceReportMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackSurfaceReportMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackSurfaceReportMessageJSON) RawJSON() string {
	return r.raw
}

type CommonResponseAttackSurfaceReport struct {
	Errors   []AttackSurfaceReportMessage `json:"errors,required"`
	Messages []AttackSurfaceReportMessage `json:"messages,required"`
	// Whether the API call was successful
	Success CommonResponseAttackSurfaceReportSuccess `json:"success,required"`
	JSON    commonResponseAttackSurfaceReportJSON    `json:"-"`
}

// commonResponseAttackSurfaceReportJSON contains the JSON metadata for the struct
// [CommonResponseAttackSurfaceReport]
type commonResponseAttackSurfaceReportJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CommonResponseAttackSurfaceReport) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r commonResponseAttackSurfaceReportJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type CommonResponseAttackSurfaceReportSuccess bool

const (
	CommonResponseAttackSurfaceReportSuccessTrue CommonResponseAttackSurfaceReportSuccess = true
)

func (r CommonResponseAttackSurfaceReportSuccess) IsKnown() bool {
	switch r {
	case CommonResponseAttackSurfaceReportSuccessTrue:
		return true
	}
	return false
}

type SingleResponseReport struct {
	Errors   []AttackSurfaceReportMessage `json:"errors,required"`
	Messages []AttackSurfaceReportMessage `json:"messages,required"`
	// Whether the API call was successful
	Success SingleResponseReportSuccess `json:"success,required"`
	JSON    singleResponseReportJSON    `json:"-"`
}

// singleResponseReportJSON contains the JSON metadata for the struct
// [SingleResponseReport]
type singleResponseReportJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponseReport) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r singleResponseReportJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SingleResponseReportSuccess bool

const (
	SingleResponseReportSuccessTrue SingleResponseReportSuccess = true
)

func (r SingleResponseReportSuccess) IsKnown() bool {
	switch r {
	case SingleResponseReportSuccessTrue:
		return true
	}
	return false
}

type AccountIntelAttackSurfaceReportListIssueTypesResponse struct {
	Result []string                                                  `json:"result"`
	JSON   accountIntelAttackSurfaceReportListIssueTypesResponseJSON `json:"-"`
	CommonResponseAttackSurfaceReport
}

// accountIntelAttackSurfaceReportListIssueTypesResponseJSON contains the JSON
// metadata for the struct [AccountIntelAttackSurfaceReportListIssueTypesResponse]
type accountIntelAttackSurfaceReportListIssueTypesResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountIntelAttackSurfaceReportListIssueTypesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountIntelAttackSurfaceReportListIssueTypesResponseJSON) RawJSON() string {
	return r.raw
}

type AccountIntelAttackSurfaceReportDismissIssueParams struct {
	Dismiss param.Field[bool] `json:"dismiss"`
}

func (r AccountIntelAttackSurfaceReportDismissIssueParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
