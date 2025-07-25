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

// AccountEmailSecurityService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountEmailSecurityService] method instead.
type AccountEmailSecurityService struct {
	Options     []option.RequestOption
	Investigate *AccountEmailSecurityInvestigateService
	Settings    *AccountEmailSecuritySettingService
}

// NewAccountEmailSecurityService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountEmailSecurityService(opts ...option.RequestOption) (r *AccountEmailSecurityService) {
	r = &AccountEmailSecurityService{}
	r.Options = opts
	r.Investigate = NewAccountEmailSecurityInvestigateService(opts...)
	r.Settings = NewAccountEmailSecuritySettingService(opts...)
	return
}

// This endpoint returns information for submissions to made to reclassify emails.
func (r *AccountEmailSecurityService) GetSubmissions(ctx context.Context, accountID string, query AccountEmailSecurityGetSubmissionsParams, opts ...option.RequestOption) (res *AccountEmailSecurityGetSubmissionsResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/email-security/submissions", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type AccountEmailSecurityGetSubmissionsResponse struct {
	Errors     []EmailSecurityMessage                             `json:"errors,required"`
	Messages   []EmailSecurityMessage                             `json:"messages,required"`
	Result     []AccountEmailSecurityGetSubmissionsResponseResult `json:"result,required"`
	ResultInfo ResultInfoEmailSecurity                            `json:"result_info,required"`
	Success    bool                                               `json:"success,required"`
	JSON       accountEmailSecurityGetSubmissionsResponseJSON     `json:"-"`
}

// accountEmailSecurityGetSubmissionsResponseJSON contains the JSON metadata for
// the struct [AccountEmailSecurityGetSubmissionsResponse]
type accountEmailSecurityGetSubmissionsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountEmailSecurityGetSubmissionsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountEmailSecurityGetSubmissionsResponseJSON) RawJSON() string {
	return r.raw
}

type AccountEmailSecurityGetSubmissionsResponseResult struct {
	RequestedTs          time.Time                                            `json:"requested_ts,required" format:"date-time"`
	SubmissionID         string                                               `json:"submission_id,required"`
	OriginalDisposition  DispositionLabel                                     `json:"original_disposition,nullable"`
	OriginalEdfHash      string                                               `json:"original_edf_hash,nullable"`
	Outcome              string                                               `json:"outcome,nullable"`
	OutcomeDisposition   DispositionLabel                                     `json:"outcome_disposition,nullable"`
	RequestedBy          string                                               `json:"requested_by,nullable"`
	RequestedDisposition DispositionLabel                                     `json:"requested_disposition,nullable"`
	Status               string                                               `json:"status,nullable"`
	Subject              string                                               `json:"subject,nullable"`
	Type                 string                                               `json:"type,nullable"`
	JSON                 accountEmailSecurityGetSubmissionsResponseResultJSON `json:"-"`
}

// accountEmailSecurityGetSubmissionsResponseResultJSON contains the JSON metadata
// for the struct [AccountEmailSecurityGetSubmissionsResponseResult]
type accountEmailSecurityGetSubmissionsResponseResultJSON struct {
	RequestedTs          apijson.Field
	SubmissionID         apijson.Field
	OriginalDisposition  apijson.Field
	OriginalEdfHash      apijson.Field
	Outcome              apijson.Field
	OutcomeDisposition   apijson.Field
	RequestedBy          apijson.Field
	RequestedDisposition apijson.Field
	Status               apijson.Field
	Subject              apijson.Field
	Type                 apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccountEmailSecurityGetSubmissionsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountEmailSecurityGetSubmissionsResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountEmailSecurityGetSubmissionsParams struct {
	// The end of the search date range. Defaults to `now`.
	End                 param.Field[time.Time]                                                   `query:"end" format:"date-time"`
	OriginalDisposition param.Field[AccountEmailSecurityGetSubmissionsParamsOriginalDisposition] `query:"original_disposition"`
	OutcomeDisposition  param.Field[AccountEmailSecurityGetSubmissionsParamsOutcomeDisposition]  `query:"outcome_disposition"`
	// The page number of paginated results.
	Page param.Field[int64] `query:"page"`
	// The number of results per page.
	PerPage              param.Field[int64]                                                        `query:"per_page"`
	Query                param.Field[string]                                                       `query:"query"`
	RequestedDisposition param.Field[AccountEmailSecurityGetSubmissionsParamsRequestedDisposition] `query:"requested_disposition"`
	// The beginning of the search date range. Defaults to `now - 30 days`.
	Start        param.Field[time.Time]                                    `query:"start" format:"date-time"`
	Status       param.Field[string]                                       `query:"status"`
	SubmissionID param.Field[string]                                       `query:"submission_id"`
	Type         param.Field[AccountEmailSecurityGetSubmissionsParamsType] `query:"type"`
}

// URLQuery serializes [AccountEmailSecurityGetSubmissionsParams]'s query
// parameters as `url.Values`.
func (r AccountEmailSecurityGetSubmissionsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountEmailSecurityGetSubmissionsParamsOriginalDisposition string

const (
	AccountEmailSecurityGetSubmissionsParamsOriginalDispositionMalicious  AccountEmailSecurityGetSubmissionsParamsOriginalDisposition = "MALICIOUS"
	AccountEmailSecurityGetSubmissionsParamsOriginalDispositionSuspicious AccountEmailSecurityGetSubmissionsParamsOriginalDisposition = "SUSPICIOUS"
	AccountEmailSecurityGetSubmissionsParamsOriginalDispositionSpoof      AccountEmailSecurityGetSubmissionsParamsOriginalDisposition = "SPOOF"
	AccountEmailSecurityGetSubmissionsParamsOriginalDispositionSpam       AccountEmailSecurityGetSubmissionsParamsOriginalDisposition = "SPAM"
	AccountEmailSecurityGetSubmissionsParamsOriginalDispositionBulk       AccountEmailSecurityGetSubmissionsParamsOriginalDisposition = "BULK"
	AccountEmailSecurityGetSubmissionsParamsOriginalDispositionNone       AccountEmailSecurityGetSubmissionsParamsOriginalDisposition = "NONE"
)

func (r AccountEmailSecurityGetSubmissionsParamsOriginalDisposition) IsKnown() bool {
	switch r {
	case AccountEmailSecurityGetSubmissionsParamsOriginalDispositionMalicious, AccountEmailSecurityGetSubmissionsParamsOriginalDispositionSuspicious, AccountEmailSecurityGetSubmissionsParamsOriginalDispositionSpoof, AccountEmailSecurityGetSubmissionsParamsOriginalDispositionSpam, AccountEmailSecurityGetSubmissionsParamsOriginalDispositionBulk, AccountEmailSecurityGetSubmissionsParamsOriginalDispositionNone:
		return true
	}
	return false
}

type AccountEmailSecurityGetSubmissionsParamsOutcomeDisposition string

const (
	AccountEmailSecurityGetSubmissionsParamsOutcomeDispositionMalicious  AccountEmailSecurityGetSubmissionsParamsOutcomeDisposition = "MALICIOUS"
	AccountEmailSecurityGetSubmissionsParamsOutcomeDispositionSuspicious AccountEmailSecurityGetSubmissionsParamsOutcomeDisposition = "SUSPICIOUS"
	AccountEmailSecurityGetSubmissionsParamsOutcomeDispositionSpoof      AccountEmailSecurityGetSubmissionsParamsOutcomeDisposition = "SPOOF"
	AccountEmailSecurityGetSubmissionsParamsOutcomeDispositionSpam       AccountEmailSecurityGetSubmissionsParamsOutcomeDisposition = "SPAM"
	AccountEmailSecurityGetSubmissionsParamsOutcomeDispositionBulk       AccountEmailSecurityGetSubmissionsParamsOutcomeDisposition = "BULK"
	AccountEmailSecurityGetSubmissionsParamsOutcomeDispositionNone       AccountEmailSecurityGetSubmissionsParamsOutcomeDisposition = "NONE"
)

func (r AccountEmailSecurityGetSubmissionsParamsOutcomeDisposition) IsKnown() bool {
	switch r {
	case AccountEmailSecurityGetSubmissionsParamsOutcomeDispositionMalicious, AccountEmailSecurityGetSubmissionsParamsOutcomeDispositionSuspicious, AccountEmailSecurityGetSubmissionsParamsOutcomeDispositionSpoof, AccountEmailSecurityGetSubmissionsParamsOutcomeDispositionSpam, AccountEmailSecurityGetSubmissionsParamsOutcomeDispositionBulk, AccountEmailSecurityGetSubmissionsParamsOutcomeDispositionNone:
		return true
	}
	return false
}

type AccountEmailSecurityGetSubmissionsParamsRequestedDisposition string

const (
	AccountEmailSecurityGetSubmissionsParamsRequestedDispositionMalicious  AccountEmailSecurityGetSubmissionsParamsRequestedDisposition = "MALICIOUS"
	AccountEmailSecurityGetSubmissionsParamsRequestedDispositionSuspicious AccountEmailSecurityGetSubmissionsParamsRequestedDisposition = "SUSPICIOUS"
	AccountEmailSecurityGetSubmissionsParamsRequestedDispositionSpoof      AccountEmailSecurityGetSubmissionsParamsRequestedDisposition = "SPOOF"
	AccountEmailSecurityGetSubmissionsParamsRequestedDispositionSpam       AccountEmailSecurityGetSubmissionsParamsRequestedDisposition = "SPAM"
	AccountEmailSecurityGetSubmissionsParamsRequestedDispositionBulk       AccountEmailSecurityGetSubmissionsParamsRequestedDisposition = "BULK"
	AccountEmailSecurityGetSubmissionsParamsRequestedDispositionNone       AccountEmailSecurityGetSubmissionsParamsRequestedDisposition = "NONE"
)

func (r AccountEmailSecurityGetSubmissionsParamsRequestedDisposition) IsKnown() bool {
	switch r {
	case AccountEmailSecurityGetSubmissionsParamsRequestedDispositionMalicious, AccountEmailSecurityGetSubmissionsParamsRequestedDispositionSuspicious, AccountEmailSecurityGetSubmissionsParamsRequestedDispositionSpoof, AccountEmailSecurityGetSubmissionsParamsRequestedDispositionSpam, AccountEmailSecurityGetSubmissionsParamsRequestedDispositionBulk, AccountEmailSecurityGetSubmissionsParamsRequestedDispositionNone:
		return true
	}
	return false
}

type AccountEmailSecurityGetSubmissionsParamsType string

const (
	AccountEmailSecurityGetSubmissionsParamsTypeTeam AccountEmailSecurityGetSubmissionsParamsType = "TEAM"
	AccountEmailSecurityGetSubmissionsParamsTypeUser AccountEmailSecurityGetSubmissionsParamsType = "USER"
)

func (r AccountEmailSecurityGetSubmissionsParamsType) IsKnown() bool {
	switch r {
	case AccountEmailSecurityGetSubmissionsParamsTypeTeam, AccountEmailSecurityGetSubmissionsParamsTypeUser:
		return true
	}
	return false
}
