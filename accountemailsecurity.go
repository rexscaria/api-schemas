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
	Result     []AccountEmailSecurityGetSubmissionsResponseResult `json:"result,required"`
	ResultInfo ResultInfoEmailSecurity                            `json:"result_info,required"`
	JSON       accountEmailSecurityGetSubmissionsResponseJSON     `json:"-"`
	APIResponseEmailSecurity
}

// accountEmailSecurityGetSubmissionsResponseJSON contains the JSON metadata for
// the struct [AccountEmailSecurityGetSubmissionsResponse]
type accountEmailSecurityGetSubmissionsResponseJSON struct {
	Result      apijson.Field
	ResultInfo  apijson.Field
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
	RequestedTs          time.Time                                                            `json:"requested_ts,required" format:"date-time"`
	SubmissionID         string                                                               `json:"submission_id,required"`
	OriginalDisposition  AccountEmailSecurityGetSubmissionsResponseResultOriginalDisposition  `json:"original_disposition,nullable"`
	OriginalEdfHash      string                                                               `json:"original_edf_hash,nullable"`
	Outcome              string                                                               `json:"outcome,nullable"`
	OutcomeDisposition   AccountEmailSecurityGetSubmissionsResponseResultOutcomeDisposition   `json:"outcome_disposition,nullable"`
	RequestedBy          string                                                               `json:"requested_by,nullable"`
	RequestedDisposition AccountEmailSecurityGetSubmissionsResponseResultRequestedDisposition `json:"requested_disposition,nullable"`
	Status               string                                                               `json:"status,nullable"`
	Subject              string                                                               `json:"subject,nullable"`
	Type                 string                                                               `json:"type,nullable"`
	JSON                 accountEmailSecurityGetSubmissionsResponseResultJSON                 `json:"-"`
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

type AccountEmailSecurityGetSubmissionsResponseResultOriginalDisposition string

const (
	AccountEmailSecurityGetSubmissionsResponseResultOriginalDispositionMalicious    AccountEmailSecurityGetSubmissionsResponseResultOriginalDisposition = "MALICIOUS"
	AccountEmailSecurityGetSubmissionsResponseResultOriginalDispositionMaliciousBec AccountEmailSecurityGetSubmissionsResponseResultOriginalDisposition = "MALICIOUS-BEC"
	AccountEmailSecurityGetSubmissionsResponseResultOriginalDispositionSuspicious   AccountEmailSecurityGetSubmissionsResponseResultOriginalDisposition = "SUSPICIOUS"
	AccountEmailSecurityGetSubmissionsResponseResultOriginalDispositionSpoof        AccountEmailSecurityGetSubmissionsResponseResultOriginalDisposition = "SPOOF"
	AccountEmailSecurityGetSubmissionsResponseResultOriginalDispositionSpam         AccountEmailSecurityGetSubmissionsResponseResultOriginalDisposition = "SPAM"
	AccountEmailSecurityGetSubmissionsResponseResultOriginalDispositionBulk         AccountEmailSecurityGetSubmissionsResponseResultOriginalDisposition = "BULK"
	AccountEmailSecurityGetSubmissionsResponseResultOriginalDispositionEncrypted    AccountEmailSecurityGetSubmissionsResponseResultOriginalDisposition = "ENCRYPTED"
	AccountEmailSecurityGetSubmissionsResponseResultOriginalDispositionExternal     AccountEmailSecurityGetSubmissionsResponseResultOriginalDisposition = "EXTERNAL"
	AccountEmailSecurityGetSubmissionsResponseResultOriginalDispositionUnknown      AccountEmailSecurityGetSubmissionsResponseResultOriginalDisposition = "UNKNOWN"
	AccountEmailSecurityGetSubmissionsResponseResultOriginalDispositionNone         AccountEmailSecurityGetSubmissionsResponseResultOriginalDisposition = "NONE"
)

func (r AccountEmailSecurityGetSubmissionsResponseResultOriginalDisposition) IsKnown() bool {
	switch r {
	case AccountEmailSecurityGetSubmissionsResponseResultOriginalDispositionMalicious, AccountEmailSecurityGetSubmissionsResponseResultOriginalDispositionMaliciousBec, AccountEmailSecurityGetSubmissionsResponseResultOriginalDispositionSuspicious, AccountEmailSecurityGetSubmissionsResponseResultOriginalDispositionSpoof, AccountEmailSecurityGetSubmissionsResponseResultOriginalDispositionSpam, AccountEmailSecurityGetSubmissionsResponseResultOriginalDispositionBulk, AccountEmailSecurityGetSubmissionsResponseResultOriginalDispositionEncrypted, AccountEmailSecurityGetSubmissionsResponseResultOriginalDispositionExternal, AccountEmailSecurityGetSubmissionsResponseResultOriginalDispositionUnknown, AccountEmailSecurityGetSubmissionsResponseResultOriginalDispositionNone:
		return true
	}
	return false
}

type AccountEmailSecurityGetSubmissionsResponseResultOutcomeDisposition string

const (
	AccountEmailSecurityGetSubmissionsResponseResultOutcomeDispositionMalicious    AccountEmailSecurityGetSubmissionsResponseResultOutcomeDisposition = "MALICIOUS"
	AccountEmailSecurityGetSubmissionsResponseResultOutcomeDispositionMaliciousBec AccountEmailSecurityGetSubmissionsResponseResultOutcomeDisposition = "MALICIOUS-BEC"
	AccountEmailSecurityGetSubmissionsResponseResultOutcomeDispositionSuspicious   AccountEmailSecurityGetSubmissionsResponseResultOutcomeDisposition = "SUSPICIOUS"
	AccountEmailSecurityGetSubmissionsResponseResultOutcomeDispositionSpoof        AccountEmailSecurityGetSubmissionsResponseResultOutcomeDisposition = "SPOOF"
	AccountEmailSecurityGetSubmissionsResponseResultOutcomeDispositionSpam         AccountEmailSecurityGetSubmissionsResponseResultOutcomeDisposition = "SPAM"
	AccountEmailSecurityGetSubmissionsResponseResultOutcomeDispositionBulk         AccountEmailSecurityGetSubmissionsResponseResultOutcomeDisposition = "BULK"
	AccountEmailSecurityGetSubmissionsResponseResultOutcomeDispositionEncrypted    AccountEmailSecurityGetSubmissionsResponseResultOutcomeDisposition = "ENCRYPTED"
	AccountEmailSecurityGetSubmissionsResponseResultOutcomeDispositionExternal     AccountEmailSecurityGetSubmissionsResponseResultOutcomeDisposition = "EXTERNAL"
	AccountEmailSecurityGetSubmissionsResponseResultOutcomeDispositionUnknown      AccountEmailSecurityGetSubmissionsResponseResultOutcomeDisposition = "UNKNOWN"
	AccountEmailSecurityGetSubmissionsResponseResultOutcomeDispositionNone         AccountEmailSecurityGetSubmissionsResponseResultOutcomeDisposition = "NONE"
)

func (r AccountEmailSecurityGetSubmissionsResponseResultOutcomeDisposition) IsKnown() bool {
	switch r {
	case AccountEmailSecurityGetSubmissionsResponseResultOutcomeDispositionMalicious, AccountEmailSecurityGetSubmissionsResponseResultOutcomeDispositionMaliciousBec, AccountEmailSecurityGetSubmissionsResponseResultOutcomeDispositionSuspicious, AccountEmailSecurityGetSubmissionsResponseResultOutcomeDispositionSpoof, AccountEmailSecurityGetSubmissionsResponseResultOutcomeDispositionSpam, AccountEmailSecurityGetSubmissionsResponseResultOutcomeDispositionBulk, AccountEmailSecurityGetSubmissionsResponseResultOutcomeDispositionEncrypted, AccountEmailSecurityGetSubmissionsResponseResultOutcomeDispositionExternal, AccountEmailSecurityGetSubmissionsResponseResultOutcomeDispositionUnknown, AccountEmailSecurityGetSubmissionsResponseResultOutcomeDispositionNone:
		return true
	}
	return false
}

type AccountEmailSecurityGetSubmissionsResponseResultRequestedDisposition string

const (
	AccountEmailSecurityGetSubmissionsResponseResultRequestedDispositionMalicious    AccountEmailSecurityGetSubmissionsResponseResultRequestedDisposition = "MALICIOUS"
	AccountEmailSecurityGetSubmissionsResponseResultRequestedDispositionMaliciousBec AccountEmailSecurityGetSubmissionsResponseResultRequestedDisposition = "MALICIOUS-BEC"
	AccountEmailSecurityGetSubmissionsResponseResultRequestedDispositionSuspicious   AccountEmailSecurityGetSubmissionsResponseResultRequestedDisposition = "SUSPICIOUS"
	AccountEmailSecurityGetSubmissionsResponseResultRequestedDispositionSpoof        AccountEmailSecurityGetSubmissionsResponseResultRequestedDisposition = "SPOOF"
	AccountEmailSecurityGetSubmissionsResponseResultRequestedDispositionSpam         AccountEmailSecurityGetSubmissionsResponseResultRequestedDisposition = "SPAM"
	AccountEmailSecurityGetSubmissionsResponseResultRequestedDispositionBulk         AccountEmailSecurityGetSubmissionsResponseResultRequestedDisposition = "BULK"
	AccountEmailSecurityGetSubmissionsResponseResultRequestedDispositionEncrypted    AccountEmailSecurityGetSubmissionsResponseResultRequestedDisposition = "ENCRYPTED"
	AccountEmailSecurityGetSubmissionsResponseResultRequestedDispositionExternal     AccountEmailSecurityGetSubmissionsResponseResultRequestedDisposition = "EXTERNAL"
	AccountEmailSecurityGetSubmissionsResponseResultRequestedDispositionUnknown      AccountEmailSecurityGetSubmissionsResponseResultRequestedDisposition = "UNKNOWN"
	AccountEmailSecurityGetSubmissionsResponseResultRequestedDispositionNone         AccountEmailSecurityGetSubmissionsResponseResultRequestedDisposition = "NONE"
)

func (r AccountEmailSecurityGetSubmissionsResponseResultRequestedDisposition) IsKnown() bool {
	switch r {
	case AccountEmailSecurityGetSubmissionsResponseResultRequestedDispositionMalicious, AccountEmailSecurityGetSubmissionsResponseResultRequestedDispositionMaliciousBec, AccountEmailSecurityGetSubmissionsResponseResultRequestedDispositionSuspicious, AccountEmailSecurityGetSubmissionsResponseResultRequestedDispositionSpoof, AccountEmailSecurityGetSubmissionsResponseResultRequestedDispositionSpam, AccountEmailSecurityGetSubmissionsResponseResultRequestedDispositionBulk, AccountEmailSecurityGetSubmissionsResponseResultRequestedDispositionEncrypted, AccountEmailSecurityGetSubmissionsResponseResultRequestedDispositionExternal, AccountEmailSecurityGetSubmissionsResponseResultRequestedDispositionUnknown, AccountEmailSecurityGetSubmissionsResponseResultRequestedDispositionNone:
		return true
	}
	return false
}

type AccountEmailSecurityGetSubmissionsParams struct {
	// The end of the search date range. Defaults to `now`.
	End param.Field[time.Time] `query:"end" format:"date-time"`
	// The page number of paginated results.
	Page param.Field[int64] `query:"page"`
	// The number of results per page.
	PerPage param.Field[int64] `query:"per_page"`
	// The beginning of the search date range. Defaults to `now - 30 days`.
	Start        param.Field[time.Time]                                    `query:"start" format:"date-time"`
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
