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

// AccountEmailSecurityInvestigateService contains methods and other services that
// help with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountEmailSecurityInvestigateService] method instead.
type AccountEmailSecurityInvestigateService struct {
	Options []option.RequestOption
	Move    *AccountEmailSecurityInvestigateMoveService
}

// NewAccountEmailSecurityInvestigateService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountEmailSecurityInvestigateService(opts ...option.RequestOption) (r *AccountEmailSecurityInvestigateService) {
	r = &AccountEmailSecurityInvestigateService{}
	r.Options = opts
	r.Move = NewAccountEmailSecurityInvestigateMoveService(opts...)
	return
}

// Get message details
func (r *AccountEmailSecurityInvestigateService) Get(ctx context.Context, accountID string, postfixID string, opts ...option.RequestOption) (res *AccountEmailSecurityInvestigateGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if postfixID == "" {
		err = errors.New("missing required postfix_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/email-security/investigate/%s", accountID, postfixID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Returns information for each email that matches the search parameter(s).
func (r *AccountEmailSecurityInvestigateService) List(ctx context.Context, accountID string, query AccountEmailSecurityInvestigateListParams, opts ...option.RequestOption) (res *AccountEmailSecurityInvestigateListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/email-security/investigate", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Returns detection details such as threat categories and sender information for
// non-benign messages.
func (r *AccountEmailSecurityInvestigateService) GetDetections(ctx context.Context, accountID string, postfixID string, opts ...option.RequestOption) (res *AccountEmailSecurityInvestigateGetDetectionsResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if postfixID == "" {
		err = errors.New("missing required postfix_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/email-security/investigate/%s/detections", accountID, postfixID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Returns the raw eml of any non-benign message.
func (r *AccountEmailSecurityInvestigateService) GetRaw(ctx context.Context, accountID string, postfixID string, opts ...option.RequestOption) (res *AccountEmailSecurityInvestigateGetRawResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if postfixID == "" {
		err = errors.New("missing required postfix_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/email-security/investigate/%s/raw", accountID, postfixID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get email trace
func (r *AccountEmailSecurityInvestigateService) GetTrace(ctx context.Context, accountID string, postfixID string, opts ...option.RequestOption) (res *AccountEmailSecurityInvestigateGetTraceResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if postfixID == "" {
		err = errors.New("missing required postfix_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/email-security/investigate/%s/trace", accountID, postfixID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Move multiple messages
func (r *AccountEmailSecurityInvestigateService) MoveMultiple(ctx context.Context, accountID string, body AccountEmailSecurityInvestigateMoveMultipleParams, opts ...option.RequestOption) (res *AccountEmailSecurityInvestigateMoveMultipleResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/email-security/investigate/move", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Returns a preview of the message body as a base64 encoded PNG image for
// non-benign messages.
func (r *AccountEmailSecurityInvestigateService) Preview(ctx context.Context, accountID string, postfixID string, opts ...option.RequestOption) (res *AccountEmailSecurityInvestigatePreviewResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if postfixID == "" {
		err = errors.New("missing required postfix_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/email-security/investigate/%s/preview", accountID, postfixID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Preview for non-detection messages
func (r *AccountEmailSecurityInvestigateService) PreviewMultiple(ctx context.Context, accountID string, body AccountEmailSecurityInvestigatePreviewMultipleParams, opts ...option.RequestOption) (res *AccountEmailSecurityInvestigatePreviewMultipleResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/email-security/investigate/preview", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Change email classfication
func (r *AccountEmailSecurityInvestigateService) Reclassify(ctx context.Context, accountID string, postfixID string, body AccountEmailSecurityInvestigateReclassifyParams, opts ...option.RequestOption) (res *AccountEmailSecurityInvestigateReclassifyResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if postfixID == "" {
		err = errors.New("missing required postfix_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/email-security/investigate/%s/reclassify", accountID, postfixID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Release messages from quarantine
func (r *AccountEmailSecurityInvestigateService) Release(ctx context.Context, accountID string, body AccountEmailSecurityInvestigateReleaseParams, opts ...option.RequestOption) (res *AccountEmailSecurityInvestigateReleaseResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/email-security/investigate/release", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type APIResponseEmailSecurity struct {
	Errors   []EmailSecurityMessage       `json:"errors,required"`
	Messages []EmailSecurityMessage       `json:"messages,required"`
	Success  bool                         `json:"success,required"`
	JSON     apiResponseEmailSecurityJSON `json:"-"`
}

// apiResponseEmailSecurityJSON contains the JSON metadata for the struct
// [APIResponseEmailSecurity]
type apiResponseEmailSecurityJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIResponseEmailSecurity) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiResponseEmailSecurityJSON) RawJSON() string {
	return r.raw
}

type EmailSecurityMessage struct {
	Code    int64                    `json:"code,required"`
	Message string                   `json:"message,required"`
	JSON    emailSecurityMessageJSON `json:"-"`
}

// emailSecurityMessageJSON contains the JSON metadata for the struct
// [EmailSecurityMessage]
type emailSecurityMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailSecurityMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailSecurityMessageJSON) RawJSON() string {
	return r.raw
}

type ResultInfoEmailSecurity struct {
	// Total number of results for the requested service
	Count int64 `json:"count,required"`
	// Current page within paginated list of results
	Page int64 `json:"page,required"`
	// Number of results per page of results
	PerPage int64 `json:"per_page,required"`
	// Total results available without any search parameters
	TotalCount int64                       `json:"total_count,required"`
	JSON       resultInfoEmailSecurityJSON `json:"-"`
}

// resultInfoEmailSecurityJSON contains the JSON metadata for the struct
// [ResultInfoEmailSecurity]
type resultInfoEmailSecurityJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResultInfoEmailSecurity) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resultInfoEmailSecurityJSON) RawJSON() string {
	return r.raw
}

type TraceLine struct {
	Lineno  int64         `json:"lineno,required"`
	Message string        `json:"message,required"`
	Ts      time.Time     `json:"ts,required" format:"date-time"`
	JSON    traceLineJSON `json:"-"`
}

// traceLineJSON contains the JSON metadata for the struct [TraceLine]
type traceLineJSON struct {
	Lineno      apijson.Field
	Message     apijson.Field
	Ts          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TraceLine) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r traceLineJSON) RawJSON() string {
	return r.raw
}

type AccountEmailSecurityInvestigateGetResponse struct {
	Result AccountEmailSecurityInvestigateGetResponseResult `json:"result,required"`
	JSON   accountEmailSecurityInvestigateGetResponseJSON   `json:"-"`
	APIResponseEmailSecurity
}

// accountEmailSecurityInvestigateGetResponseJSON contains the JSON metadata for
// the struct [AccountEmailSecurityInvestigateGetResponse]
type accountEmailSecurityInvestigateGetResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountEmailSecurityInvestigateGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountEmailSecurityInvestigateGetResponseJSON) RawJSON() string {
	return r.raw
}

type AccountEmailSecurityInvestigateGetResponseResult struct {
	ID                string      `json:"id,required"`
	ActionLog         interface{} `json:"action_log,required"`
	ClientRecipients  []string    `json:"client_recipients,required"`
	DetectionReasons  []string    `json:"detection_reasons,required"`
	IsPhishSubmission bool        `json:"is_phish_submission,required"`
	IsQuarantined     bool        `json:"is_quarantined,required"`
	// The identifier of the message.
	PostfixID        string                                                           `json:"postfix_id,required"`
	Ts               string                                                           `json:"ts,required"`
	AlertID          string                                                           `json:"alert_id,nullable"`
	DeliveryMode     AccountEmailSecurityInvestigateGetResponseResultDeliveryMode     `json:"delivery_mode,nullable"`
	EdfHash          string                                                           `json:"edf_hash,nullable"`
	FinalDisposition AccountEmailSecurityInvestigateGetResponseResultFinalDisposition `json:"final_disposition,nullable"`
	From             string                                                           `json:"from,nullable"`
	FromName         string                                                           `json:"from_name,nullable"`
	MessageID        string                                                           `json:"message_id,nullable"`
	SentDate         string                                                           `json:"sent_date,nullable"`
	Subject          string                                                           `json:"subject,nullable"`
	ThreatCategories []string                                                         `json:"threat_categories,nullable"`
	To               []string                                                         `json:"to,nullable"`
	ToName           []string                                                         `json:"to_name,nullable"`
	Validation       AccountEmailSecurityInvestigateGetResponseResultValidation       `json:"validation"`
	JSON             accountEmailSecurityInvestigateGetResponseResultJSON             `json:"-"`
}

// accountEmailSecurityInvestigateGetResponseResultJSON contains the JSON metadata
// for the struct [AccountEmailSecurityInvestigateGetResponseResult]
type accountEmailSecurityInvestigateGetResponseResultJSON struct {
	ID                apijson.Field
	ActionLog         apijson.Field
	ClientRecipients  apijson.Field
	DetectionReasons  apijson.Field
	IsPhishSubmission apijson.Field
	IsQuarantined     apijson.Field
	PostfixID         apijson.Field
	Ts                apijson.Field
	AlertID           apijson.Field
	DeliveryMode      apijson.Field
	EdfHash           apijson.Field
	FinalDisposition  apijson.Field
	From              apijson.Field
	FromName          apijson.Field
	MessageID         apijson.Field
	SentDate          apijson.Field
	Subject           apijson.Field
	ThreatCategories  apijson.Field
	To                apijson.Field
	ToName            apijson.Field
	Validation        apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *AccountEmailSecurityInvestigateGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountEmailSecurityInvestigateGetResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountEmailSecurityInvestigateGetResponseResultDeliveryMode string

const (
	AccountEmailSecurityInvestigateGetResponseResultDeliveryModeDirect                AccountEmailSecurityInvestigateGetResponseResultDeliveryMode = "DIRECT"
	AccountEmailSecurityInvestigateGetResponseResultDeliveryModeBcc                   AccountEmailSecurityInvestigateGetResponseResultDeliveryMode = "BCC"
	AccountEmailSecurityInvestigateGetResponseResultDeliveryModeJournal               AccountEmailSecurityInvestigateGetResponseResultDeliveryMode = "JOURNAL"
	AccountEmailSecurityInvestigateGetResponseResultDeliveryModeReviewSubmission      AccountEmailSecurityInvestigateGetResponseResultDeliveryMode = "REVIEW_SUBMISSION"
	AccountEmailSecurityInvestigateGetResponseResultDeliveryModeDmarcUnverified       AccountEmailSecurityInvestigateGetResponseResultDeliveryMode = "DMARC_UNVERIFIED"
	AccountEmailSecurityInvestigateGetResponseResultDeliveryModeDmarcFailureReport    AccountEmailSecurityInvestigateGetResponseResultDeliveryMode = "DMARC_FAILURE_REPORT"
	AccountEmailSecurityInvestigateGetResponseResultDeliveryModeDmarcAggregateReport  AccountEmailSecurityInvestigateGetResponseResultDeliveryMode = "DMARC_AGGREGATE_REPORT"
	AccountEmailSecurityInvestigateGetResponseResultDeliveryModeThreatIntelSubmission AccountEmailSecurityInvestigateGetResponseResultDeliveryMode = "THREAT_INTEL_SUBMISSION"
	AccountEmailSecurityInvestigateGetResponseResultDeliveryModeSimulationSubmission  AccountEmailSecurityInvestigateGetResponseResultDeliveryMode = "SIMULATION_SUBMISSION"
	AccountEmailSecurityInvestigateGetResponseResultDeliveryModeAPI                   AccountEmailSecurityInvestigateGetResponseResultDeliveryMode = "API"
	AccountEmailSecurityInvestigateGetResponseResultDeliveryModeRetroScan             AccountEmailSecurityInvestigateGetResponseResultDeliveryMode = "RETRO_SCAN"
)

func (r AccountEmailSecurityInvestigateGetResponseResultDeliveryMode) IsKnown() bool {
	switch r {
	case AccountEmailSecurityInvestigateGetResponseResultDeliveryModeDirect, AccountEmailSecurityInvestigateGetResponseResultDeliveryModeBcc, AccountEmailSecurityInvestigateGetResponseResultDeliveryModeJournal, AccountEmailSecurityInvestigateGetResponseResultDeliveryModeReviewSubmission, AccountEmailSecurityInvestigateGetResponseResultDeliveryModeDmarcUnverified, AccountEmailSecurityInvestigateGetResponseResultDeliveryModeDmarcFailureReport, AccountEmailSecurityInvestigateGetResponseResultDeliveryModeDmarcAggregateReport, AccountEmailSecurityInvestigateGetResponseResultDeliveryModeThreatIntelSubmission, AccountEmailSecurityInvestigateGetResponseResultDeliveryModeSimulationSubmission, AccountEmailSecurityInvestigateGetResponseResultDeliveryModeAPI, AccountEmailSecurityInvestigateGetResponseResultDeliveryModeRetroScan:
		return true
	}
	return false
}

type AccountEmailSecurityInvestigateGetResponseResultFinalDisposition string

const (
	AccountEmailSecurityInvestigateGetResponseResultFinalDispositionMalicious    AccountEmailSecurityInvestigateGetResponseResultFinalDisposition = "MALICIOUS"
	AccountEmailSecurityInvestigateGetResponseResultFinalDispositionMaliciousBec AccountEmailSecurityInvestigateGetResponseResultFinalDisposition = "MALICIOUS-BEC"
	AccountEmailSecurityInvestigateGetResponseResultFinalDispositionSuspicious   AccountEmailSecurityInvestigateGetResponseResultFinalDisposition = "SUSPICIOUS"
	AccountEmailSecurityInvestigateGetResponseResultFinalDispositionSpoof        AccountEmailSecurityInvestigateGetResponseResultFinalDisposition = "SPOOF"
	AccountEmailSecurityInvestigateGetResponseResultFinalDispositionSpam         AccountEmailSecurityInvestigateGetResponseResultFinalDisposition = "SPAM"
	AccountEmailSecurityInvestigateGetResponseResultFinalDispositionBulk         AccountEmailSecurityInvestigateGetResponseResultFinalDisposition = "BULK"
	AccountEmailSecurityInvestigateGetResponseResultFinalDispositionEncrypted    AccountEmailSecurityInvestigateGetResponseResultFinalDisposition = "ENCRYPTED"
	AccountEmailSecurityInvestigateGetResponseResultFinalDispositionExternal     AccountEmailSecurityInvestigateGetResponseResultFinalDisposition = "EXTERNAL"
	AccountEmailSecurityInvestigateGetResponseResultFinalDispositionUnknown      AccountEmailSecurityInvestigateGetResponseResultFinalDisposition = "UNKNOWN"
	AccountEmailSecurityInvestigateGetResponseResultFinalDispositionNone         AccountEmailSecurityInvestigateGetResponseResultFinalDisposition = "NONE"
)

func (r AccountEmailSecurityInvestigateGetResponseResultFinalDisposition) IsKnown() bool {
	switch r {
	case AccountEmailSecurityInvestigateGetResponseResultFinalDispositionMalicious, AccountEmailSecurityInvestigateGetResponseResultFinalDispositionMaliciousBec, AccountEmailSecurityInvestigateGetResponseResultFinalDispositionSuspicious, AccountEmailSecurityInvestigateGetResponseResultFinalDispositionSpoof, AccountEmailSecurityInvestigateGetResponseResultFinalDispositionSpam, AccountEmailSecurityInvestigateGetResponseResultFinalDispositionBulk, AccountEmailSecurityInvestigateGetResponseResultFinalDispositionEncrypted, AccountEmailSecurityInvestigateGetResponseResultFinalDispositionExternal, AccountEmailSecurityInvestigateGetResponseResultFinalDispositionUnknown, AccountEmailSecurityInvestigateGetResponseResultFinalDispositionNone:
		return true
	}
	return false
}

type AccountEmailSecurityInvestigateGetResponseResultValidation struct {
	Comment string                                                          `json:"comment,nullable"`
	Dkim    AccountEmailSecurityInvestigateGetResponseResultValidationDkim  `json:"dkim,nullable"`
	Dmarc   AccountEmailSecurityInvestigateGetResponseResultValidationDmarc `json:"dmarc,nullable"`
	Spf     AccountEmailSecurityInvestigateGetResponseResultValidationSpf   `json:"spf,nullable"`
	JSON    accountEmailSecurityInvestigateGetResponseResultValidationJSON  `json:"-"`
}

// accountEmailSecurityInvestigateGetResponseResultValidationJSON contains the JSON
// metadata for the struct
// [AccountEmailSecurityInvestigateGetResponseResultValidation]
type accountEmailSecurityInvestigateGetResponseResultValidationJSON struct {
	Comment     apijson.Field
	Dkim        apijson.Field
	Dmarc       apijson.Field
	Spf         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountEmailSecurityInvestigateGetResponseResultValidation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountEmailSecurityInvestigateGetResponseResultValidationJSON) RawJSON() string {
	return r.raw
}

type AccountEmailSecurityInvestigateGetResponseResultValidationDkim string

const (
	AccountEmailSecurityInvestigateGetResponseResultValidationDkimPass    AccountEmailSecurityInvestigateGetResponseResultValidationDkim = "pass"
	AccountEmailSecurityInvestigateGetResponseResultValidationDkimNeutral AccountEmailSecurityInvestigateGetResponseResultValidationDkim = "neutral"
	AccountEmailSecurityInvestigateGetResponseResultValidationDkimFail    AccountEmailSecurityInvestigateGetResponseResultValidationDkim = "fail"
	AccountEmailSecurityInvestigateGetResponseResultValidationDkimError   AccountEmailSecurityInvestigateGetResponseResultValidationDkim = "error"
	AccountEmailSecurityInvestigateGetResponseResultValidationDkimNone    AccountEmailSecurityInvestigateGetResponseResultValidationDkim = "none"
)

func (r AccountEmailSecurityInvestigateGetResponseResultValidationDkim) IsKnown() bool {
	switch r {
	case AccountEmailSecurityInvestigateGetResponseResultValidationDkimPass, AccountEmailSecurityInvestigateGetResponseResultValidationDkimNeutral, AccountEmailSecurityInvestigateGetResponseResultValidationDkimFail, AccountEmailSecurityInvestigateGetResponseResultValidationDkimError, AccountEmailSecurityInvestigateGetResponseResultValidationDkimNone:
		return true
	}
	return false
}

type AccountEmailSecurityInvestigateGetResponseResultValidationDmarc string

const (
	AccountEmailSecurityInvestigateGetResponseResultValidationDmarcPass    AccountEmailSecurityInvestigateGetResponseResultValidationDmarc = "pass"
	AccountEmailSecurityInvestigateGetResponseResultValidationDmarcNeutral AccountEmailSecurityInvestigateGetResponseResultValidationDmarc = "neutral"
	AccountEmailSecurityInvestigateGetResponseResultValidationDmarcFail    AccountEmailSecurityInvestigateGetResponseResultValidationDmarc = "fail"
	AccountEmailSecurityInvestigateGetResponseResultValidationDmarcError   AccountEmailSecurityInvestigateGetResponseResultValidationDmarc = "error"
	AccountEmailSecurityInvestigateGetResponseResultValidationDmarcNone    AccountEmailSecurityInvestigateGetResponseResultValidationDmarc = "none"
)

func (r AccountEmailSecurityInvestigateGetResponseResultValidationDmarc) IsKnown() bool {
	switch r {
	case AccountEmailSecurityInvestigateGetResponseResultValidationDmarcPass, AccountEmailSecurityInvestigateGetResponseResultValidationDmarcNeutral, AccountEmailSecurityInvestigateGetResponseResultValidationDmarcFail, AccountEmailSecurityInvestigateGetResponseResultValidationDmarcError, AccountEmailSecurityInvestigateGetResponseResultValidationDmarcNone:
		return true
	}
	return false
}

type AccountEmailSecurityInvestigateGetResponseResultValidationSpf string

const (
	AccountEmailSecurityInvestigateGetResponseResultValidationSpfPass    AccountEmailSecurityInvestigateGetResponseResultValidationSpf = "pass"
	AccountEmailSecurityInvestigateGetResponseResultValidationSpfNeutral AccountEmailSecurityInvestigateGetResponseResultValidationSpf = "neutral"
	AccountEmailSecurityInvestigateGetResponseResultValidationSpfFail    AccountEmailSecurityInvestigateGetResponseResultValidationSpf = "fail"
	AccountEmailSecurityInvestigateGetResponseResultValidationSpfError   AccountEmailSecurityInvestigateGetResponseResultValidationSpf = "error"
	AccountEmailSecurityInvestigateGetResponseResultValidationSpfNone    AccountEmailSecurityInvestigateGetResponseResultValidationSpf = "none"
)

func (r AccountEmailSecurityInvestigateGetResponseResultValidationSpf) IsKnown() bool {
	switch r {
	case AccountEmailSecurityInvestigateGetResponseResultValidationSpfPass, AccountEmailSecurityInvestigateGetResponseResultValidationSpfNeutral, AccountEmailSecurityInvestigateGetResponseResultValidationSpfFail, AccountEmailSecurityInvestigateGetResponseResultValidationSpfError, AccountEmailSecurityInvestigateGetResponseResultValidationSpfNone:
		return true
	}
	return false
}

type AccountEmailSecurityInvestigateListResponse struct {
	Result     []AccountEmailSecurityInvestigateListResponseResult `json:"result,required"`
	ResultInfo ResultInfoEmailSecurity                             `json:"result_info,required"`
	JSON       accountEmailSecurityInvestigateListResponseJSON     `json:"-"`
	APIResponseEmailSecurity
}

// accountEmailSecurityInvestigateListResponseJSON contains the JSON metadata for
// the struct [AccountEmailSecurityInvestigateListResponse]
type accountEmailSecurityInvestigateListResponseJSON struct {
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountEmailSecurityInvestigateListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountEmailSecurityInvestigateListResponseJSON) RawJSON() string {
	return r.raw
}

type AccountEmailSecurityInvestigateListResponseResult struct {
	ID                string      `json:"id,required"`
	ActionLog         interface{} `json:"action_log,required"`
	ClientRecipients  []string    `json:"client_recipients,required"`
	DetectionReasons  []string    `json:"detection_reasons,required"`
	IsPhishSubmission bool        `json:"is_phish_submission,required"`
	IsQuarantined     bool        `json:"is_quarantined,required"`
	// The identifier of the message.
	PostfixID        string                                                            `json:"postfix_id,required"`
	Ts               string                                                            `json:"ts,required"`
	AlertID          string                                                            `json:"alert_id,nullable"`
	DeliveryMode     AccountEmailSecurityInvestigateListResponseResultDeliveryMode     `json:"delivery_mode,nullable"`
	EdfHash          string                                                            `json:"edf_hash,nullable"`
	FinalDisposition AccountEmailSecurityInvestigateListResponseResultFinalDisposition `json:"final_disposition,nullable"`
	From             string                                                            `json:"from,nullable"`
	FromName         string                                                            `json:"from_name,nullable"`
	MessageID        string                                                            `json:"message_id,nullable"`
	SentDate         string                                                            `json:"sent_date,nullable"`
	Subject          string                                                            `json:"subject,nullable"`
	ThreatCategories []string                                                          `json:"threat_categories,nullable"`
	To               []string                                                          `json:"to,nullable"`
	ToName           []string                                                          `json:"to_name,nullable"`
	Validation       AccountEmailSecurityInvestigateListResponseResultValidation       `json:"validation"`
	JSON             accountEmailSecurityInvestigateListResponseResultJSON             `json:"-"`
}

// accountEmailSecurityInvestigateListResponseResultJSON contains the JSON metadata
// for the struct [AccountEmailSecurityInvestigateListResponseResult]
type accountEmailSecurityInvestigateListResponseResultJSON struct {
	ID                apijson.Field
	ActionLog         apijson.Field
	ClientRecipients  apijson.Field
	DetectionReasons  apijson.Field
	IsPhishSubmission apijson.Field
	IsQuarantined     apijson.Field
	PostfixID         apijson.Field
	Ts                apijson.Field
	AlertID           apijson.Field
	DeliveryMode      apijson.Field
	EdfHash           apijson.Field
	FinalDisposition  apijson.Field
	From              apijson.Field
	FromName          apijson.Field
	MessageID         apijson.Field
	SentDate          apijson.Field
	Subject           apijson.Field
	ThreatCategories  apijson.Field
	To                apijson.Field
	ToName            apijson.Field
	Validation        apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *AccountEmailSecurityInvestigateListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountEmailSecurityInvestigateListResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountEmailSecurityInvestigateListResponseResultDeliveryMode string

const (
	AccountEmailSecurityInvestigateListResponseResultDeliveryModeDirect                AccountEmailSecurityInvestigateListResponseResultDeliveryMode = "DIRECT"
	AccountEmailSecurityInvestigateListResponseResultDeliveryModeBcc                   AccountEmailSecurityInvestigateListResponseResultDeliveryMode = "BCC"
	AccountEmailSecurityInvestigateListResponseResultDeliveryModeJournal               AccountEmailSecurityInvestigateListResponseResultDeliveryMode = "JOURNAL"
	AccountEmailSecurityInvestigateListResponseResultDeliveryModeReviewSubmission      AccountEmailSecurityInvestigateListResponseResultDeliveryMode = "REVIEW_SUBMISSION"
	AccountEmailSecurityInvestigateListResponseResultDeliveryModeDmarcUnverified       AccountEmailSecurityInvestigateListResponseResultDeliveryMode = "DMARC_UNVERIFIED"
	AccountEmailSecurityInvestigateListResponseResultDeliveryModeDmarcFailureReport    AccountEmailSecurityInvestigateListResponseResultDeliveryMode = "DMARC_FAILURE_REPORT"
	AccountEmailSecurityInvestigateListResponseResultDeliveryModeDmarcAggregateReport  AccountEmailSecurityInvestigateListResponseResultDeliveryMode = "DMARC_AGGREGATE_REPORT"
	AccountEmailSecurityInvestigateListResponseResultDeliveryModeThreatIntelSubmission AccountEmailSecurityInvestigateListResponseResultDeliveryMode = "THREAT_INTEL_SUBMISSION"
	AccountEmailSecurityInvestigateListResponseResultDeliveryModeSimulationSubmission  AccountEmailSecurityInvestigateListResponseResultDeliveryMode = "SIMULATION_SUBMISSION"
	AccountEmailSecurityInvestigateListResponseResultDeliveryModeAPI                   AccountEmailSecurityInvestigateListResponseResultDeliveryMode = "API"
	AccountEmailSecurityInvestigateListResponseResultDeliveryModeRetroScan             AccountEmailSecurityInvestigateListResponseResultDeliveryMode = "RETRO_SCAN"
)

func (r AccountEmailSecurityInvestigateListResponseResultDeliveryMode) IsKnown() bool {
	switch r {
	case AccountEmailSecurityInvestigateListResponseResultDeliveryModeDirect, AccountEmailSecurityInvestigateListResponseResultDeliveryModeBcc, AccountEmailSecurityInvestigateListResponseResultDeliveryModeJournal, AccountEmailSecurityInvestigateListResponseResultDeliveryModeReviewSubmission, AccountEmailSecurityInvestigateListResponseResultDeliveryModeDmarcUnverified, AccountEmailSecurityInvestigateListResponseResultDeliveryModeDmarcFailureReport, AccountEmailSecurityInvestigateListResponseResultDeliveryModeDmarcAggregateReport, AccountEmailSecurityInvestigateListResponseResultDeliveryModeThreatIntelSubmission, AccountEmailSecurityInvestigateListResponseResultDeliveryModeSimulationSubmission, AccountEmailSecurityInvestigateListResponseResultDeliveryModeAPI, AccountEmailSecurityInvestigateListResponseResultDeliveryModeRetroScan:
		return true
	}
	return false
}

type AccountEmailSecurityInvestigateListResponseResultFinalDisposition string

const (
	AccountEmailSecurityInvestigateListResponseResultFinalDispositionMalicious    AccountEmailSecurityInvestigateListResponseResultFinalDisposition = "MALICIOUS"
	AccountEmailSecurityInvestigateListResponseResultFinalDispositionMaliciousBec AccountEmailSecurityInvestigateListResponseResultFinalDisposition = "MALICIOUS-BEC"
	AccountEmailSecurityInvestigateListResponseResultFinalDispositionSuspicious   AccountEmailSecurityInvestigateListResponseResultFinalDisposition = "SUSPICIOUS"
	AccountEmailSecurityInvestigateListResponseResultFinalDispositionSpoof        AccountEmailSecurityInvestigateListResponseResultFinalDisposition = "SPOOF"
	AccountEmailSecurityInvestigateListResponseResultFinalDispositionSpam         AccountEmailSecurityInvestigateListResponseResultFinalDisposition = "SPAM"
	AccountEmailSecurityInvestigateListResponseResultFinalDispositionBulk         AccountEmailSecurityInvestigateListResponseResultFinalDisposition = "BULK"
	AccountEmailSecurityInvestigateListResponseResultFinalDispositionEncrypted    AccountEmailSecurityInvestigateListResponseResultFinalDisposition = "ENCRYPTED"
	AccountEmailSecurityInvestigateListResponseResultFinalDispositionExternal     AccountEmailSecurityInvestigateListResponseResultFinalDisposition = "EXTERNAL"
	AccountEmailSecurityInvestigateListResponseResultFinalDispositionUnknown      AccountEmailSecurityInvestigateListResponseResultFinalDisposition = "UNKNOWN"
	AccountEmailSecurityInvestigateListResponseResultFinalDispositionNone         AccountEmailSecurityInvestigateListResponseResultFinalDisposition = "NONE"
)

func (r AccountEmailSecurityInvestigateListResponseResultFinalDisposition) IsKnown() bool {
	switch r {
	case AccountEmailSecurityInvestigateListResponseResultFinalDispositionMalicious, AccountEmailSecurityInvestigateListResponseResultFinalDispositionMaliciousBec, AccountEmailSecurityInvestigateListResponseResultFinalDispositionSuspicious, AccountEmailSecurityInvestigateListResponseResultFinalDispositionSpoof, AccountEmailSecurityInvestigateListResponseResultFinalDispositionSpam, AccountEmailSecurityInvestigateListResponseResultFinalDispositionBulk, AccountEmailSecurityInvestigateListResponseResultFinalDispositionEncrypted, AccountEmailSecurityInvestigateListResponseResultFinalDispositionExternal, AccountEmailSecurityInvestigateListResponseResultFinalDispositionUnknown, AccountEmailSecurityInvestigateListResponseResultFinalDispositionNone:
		return true
	}
	return false
}

type AccountEmailSecurityInvestigateListResponseResultValidation struct {
	Comment string                                                           `json:"comment,nullable"`
	Dkim    AccountEmailSecurityInvestigateListResponseResultValidationDkim  `json:"dkim,nullable"`
	Dmarc   AccountEmailSecurityInvestigateListResponseResultValidationDmarc `json:"dmarc,nullable"`
	Spf     AccountEmailSecurityInvestigateListResponseResultValidationSpf   `json:"spf,nullable"`
	JSON    accountEmailSecurityInvestigateListResponseResultValidationJSON  `json:"-"`
}

// accountEmailSecurityInvestigateListResponseResultValidationJSON contains the
// JSON metadata for the struct
// [AccountEmailSecurityInvestigateListResponseResultValidation]
type accountEmailSecurityInvestigateListResponseResultValidationJSON struct {
	Comment     apijson.Field
	Dkim        apijson.Field
	Dmarc       apijson.Field
	Spf         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountEmailSecurityInvestigateListResponseResultValidation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountEmailSecurityInvestigateListResponseResultValidationJSON) RawJSON() string {
	return r.raw
}

type AccountEmailSecurityInvestigateListResponseResultValidationDkim string

const (
	AccountEmailSecurityInvestigateListResponseResultValidationDkimPass    AccountEmailSecurityInvestigateListResponseResultValidationDkim = "pass"
	AccountEmailSecurityInvestigateListResponseResultValidationDkimNeutral AccountEmailSecurityInvestigateListResponseResultValidationDkim = "neutral"
	AccountEmailSecurityInvestigateListResponseResultValidationDkimFail    AccountEmailSecurityInvestigateListResponseResultValidationDkim = "fail"
	AccountEmailSecurityInvestigateListResponseResultValidationDkimError   AccountEmailSecurityInvestigateListResponseResultValidationDkim = "error"
	AccountEmailSecurityInvestigateListResponseResultValidationDkimNone    AccountEmailSecurityInvestigateListResponseResultValidationDkim = "none"
)

func (r AccountEmailSecurityInvestigateListResponseResultValidationDkim) IsKnown() bool {
	switch r {
	case AccountEmailSecurityInvestigateListResponseResultValidationDkimPass, AccountEmailSecurityInvestigateListResponseResultValidationDkimNeutral, AccountEmailSecurityInvestigateListResponseResultValidationDkimFail, AccountEmailSecurityInvestigateListResponseResultValidationDkimError, AccountEmailSecurityInvestigateListResponseResultValidationDkimNone:
		return true
	}
	return false
}

type AccountEmailSecurityInvestigateListResponseResultValidationDmarc string

const (
	AccountEmailSecurityInvestigateListResponseResultValidationDmarcPass    AccountEmailSecurityInvestigateListResponseResultValidationDmarc = "pass"
	AccountEmailSecurityInvestigateListResponseResultValidationDmarcNeutral AccountEmailSecurityInvestigateListResponseResultValidationDmarc = "neutral"
	AccountEmailSecurityInvestigateListResponseResultValidationDmarcFail    AccountEmailSecurityInvestigateListResponseResultValidationDmarc = "fail"
	AccountEmailSecurityInvestigateListResponseResultValidationDmarcError   AccountEmailSecurityInvestigateListResponseResultValidationDmarc = "error"
	AccountEmailSecurityInvestigateListResponseResultValidationDmarcNone    AccountEmailSecurityInvestigateListResponseResultValidationDmarc = "none"
)

func (r AccountEmailSecurityInvestigateListResponseResultValidationDmarc) IsKnown() bool {
	switch r {
	case AccountEmailSecurityInvestigateListResponseResultValidationDmarcPass, AccountEmailSecurityInvestigateListResponseResultValidationDmarcNeutral, AccountEmailSecurityInvestigateListResponseResultValidationDmarcFail, AccountEmailSecurityInvestigateListResponseResultValidationDmarcError, AccountEmailSecurityInvestigateListResponseResultValidationDmarcNone:
		return true
	}
	return false
}

type AccountEmailSecurityInvestigateListResponseResultValidationSpf string

const (
	AccountEmailSecurityInvestigateListResponseResultValidationSpfPass    AccountEmailSecurityInvestigateListResponseResultValidationSpf = "pass"
	AccountEmailSecurityInvestigateListResponseResultValidationSpfNeutral AccountEmailSecurityInvestigateListResponseResultValidationSpf = "neutral"
	AccountEmailSecurityInvestigateListResponseResultValidationSpfFail    AccountEmailSecurityInvestigateListResponseResultValidationSpf = "fail"
	AccountEmailSecurityInvestigateListResponseResultValidationSpfError   AccountEmailSecurityInvestigateListResponseResultValidationSpf = "error"
	AccountEmailSecurityInvestigateListResponseResultValidationSpfNone    AccountEmailSecurityInvestigateListResponseResultValidationSpf = "none"
)

func (r AccountEmailSecurityInvestigateListResponseResultValidationSpf) IsKnown() bool {
	switch r {
	case AccountEmailSecurityInvestigateListResponseResultValidationSpfPass, AccountEmailSecurityInvestigateListResponseResultValidationSpfNeutral, AccountEmailSecurityInvestigateListResponseResultValidationSpfFail, AccountEmailSecurityInvestigateListResponseResultValidationSpfError, AccountEmailSecurityInvestigateListResponseResultValidationSpfNone:
		return true
	}
	return false
}

type AccountEmailSecurityInvestigateGetDetectionsResponse struct {
	Result AccountEmailSecurityInvestigateGetDetectionsResponseResult `json:"result,required"`
	JSON   accountEmailSecurityInvestigateGetDetectionsResponseJSON   `json:"-"`
	APIResponseEmailSecurity
}

// accountEmailSecurityInvestigateGetDetectionsResponseJSON contains the JSON
// metadata for the struct [AccountEmailSecurityInvestigateGetDetectionsResponse]
type accountEmailSecurityInvestigateGetDetectionsResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountEmailSecurityInvestigateGetDetectionsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountEmailSecurityInvestigateGetDetectionsResponseJSON) RawJSON() string {
	return r.raw
}

type AccountEmailSecurityInvestigateGetDetectionsResponseResult struct {
	Action           string                                                                     `json:"action,required"`
	Attachments      []AccountEmailSecurityInvestigateGetDetectionsResponseResultAttachment     `json:"attachments,required"`
	Headers          []AccountEmailSecurityInvestigateGetDetectionsResponseResultHeader         `json:"headers,required"`
	Links            []AccountEmailSecurityInvestigateGetDetectionsResponseResultLink           `json:"links,required"`
	SenderInfo       AccountEmailSecurityInvestigateGetDetectionsResponseResultSenderInfo       `json:"sender_info,required"`
	ThreatCategories []AccountEmailSecurityInvestigateGetDetectionsResponseResultThreatCategory `json:"threat_categories,required"`
	Validation       AccountEmailSecurityInvestigateGetDetectionsResponseResultValidation       `json:"validation,required"`
	FinalDisposition AccountEmailSecurityInvestigateGetDetectionsResponseResultFinalDisposition `json:"final_disposition,nullable"`
	JSON             accountEmailSecurityInvestigateGetDetectionsResponseResultJSON             `json:"-"`
}

// accountEmailSecurityInvestigateGetDetectionsResponseResultJSON contains the JSON
// metadata for the struct
// [AccountEmailSecurityInvestigateGetDetectionsResponseResult]
type accountEmailSecurityInvestigateGetDetectionsResponseResultJSON struct {
	Action           apijson.Field
	Attachments      apijson.Field
	Headers          apijson.Field
	Links            apijson.Field
	SenderInfo       apijson.Field
	ThreatCategories apijson.Field
	Validation       apijson.Field
	FinalDisposition apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountEmailSecurityInvestigateGetDetectionsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountEmailSecurityInvestigateGetDetectionsResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountEmailSecurityInvestigateGetDetectionsResponseResultAttachment struct {
	Size        int64                                                                          `json:"size,required"`
	ContentType string                                                                         `json:"content_type,nullable"`
	Detection   AccountEmailSecurityInvestigateGetDetectionsResponseResultAttachmentsDetection `json:"detection,nullable"`
	Encrypted   bool                                                                           `json:"encrypted,nullable"`
	Name        string                                                                         `json:"name,nullable"`
	JSON        accountEmailSecurityInvestigateGetDetectionsResponseResultAttachmentJSON       `json:"-"`
}

// accountEmailSecurityInvestigateGetDetectionsResponseResultAttachmentJSON
// contains the JSON metadata for the struct
// [AccountEmailSecurityInvestigateGetDetectionsResponseResultAttachment]
type accountEmailSecurityInvestigateGetDetectionsResponseResultAttachmentJSON struct {
	Size        apijson.Field
	ContentType apijson.Field
	Detection   apijson.Field
	Encrypted   apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountEmailSecurityInvestigateGetDetectionsResponseResultAttachment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountEmailSecurityInvestigateGetDetectionsResponseResultAttachmentJSON) RawJSON() string {
	return r.raw
}

type AccountEmailSecurityInvestigateGetDetectionsResponseResultAttachmentsDetection string

const (
	AccountEmailSecurityInvestigateGetDetectionsResponseResultAttachmentsDetectionMalicious    AccountEmailSecurityInvestigateGetDetectionsResponseResultAttachmentsDetection = "MALICIOUS"
	AccountEmailSecurityInvestigateGetDetectionsResponseResultAttachmentsDetectionMaliciousBec AccountEmailSecurityInvestigateGetDetectionsResponseResultAttachmentsDetection = "MALICIOUS-BEC"
	AccountEmailSecurityInvestigateGetDetectionsResponseResultAttachmentsDetectionSuspicious   AccountEmailSecurityInvestigateGetDetectionsResponseResultAttachmentsDetection = "SUSPICIOUS"
	AccountEmailSecurityInvestigateGetDetectionsResponseResultAttachmentsDetectionSpoof        AccountEmailSecurityInvestigateGetDetectionsResponseResultAttachmentsDetection = "SPOOF"
	AccountEmailSecurityInvestigateGetDetectionsResponseResultAttachmentsDetectionSpam         AccountEmailSecurityInvestigateGetDetectionsResponseResultAttachmentsDetection = "SPAM"
	AccountEmailSecurityInvestigateGetDetectionsResponseResultAttachmentsDetectionBulk         AccountEmailSecurityInvestigateGetDetectionsResponseResultAttachmentsDetection = "BULK"
	AccountEmailSecurityInvestigateGetDetectionsResponseResultAttachmentsDetectionEncrypted    AccountEmailSecurityInvestigateGetDetectionsResponseResultAttachmentsDetection = "ENCRYPTED"
	AccountEmailSecurityInvestigateGetDetectionsResponseResultAttachmentsDetectionExternal     AccountEmailSecurityInvestigateGetDetectionsResponseResultAttachmentsDetection = "EXTERNAL"
	AccountEmailSecurityInvestigateGetDetectionsResponseResultAttachmentsDetectionUnknown      AccountEmailSecurityInvestigateGetDetectionsResponseResultAttachmentsDetection = "UNKNOWN"
	AccountEmailSecurityInvestigateGetDetectionsResponseResultAttachmentsDetectionNone         AccountEmailSecurityInvestigateGetDetectionsResponseResultAttachmentsDetection = "NONE"
)

func (r AccountEmailSecurityInvestigateGetDetectionsResponseResultAttachmentsDetection) IsKnown() bool {
	switch r {
	case AccountEmailSecurityInvestigateGetDetectionsResponseResultAttachmentsDetectionMalicious, AccountEmailSecurityInvestigateGetDetectionsResponseResultAttachmentsDetectionMaliciousBec, AccountEmailSecurityInvestigateGetDetectionsResponseResultAttachmentsDetectionSuspicious, AccountEmailSecurityInvestigateGetDetectionsResponseResultAttachmentsDetectionSpoof, AccountEmailSecurityInvestigateGetDetectionsResponseResultAttachmentsDetectionSpam, AccountEmailSecurityInvestigateGetDetectionsResponseResultAttachmentsDetectionBulk, AccountEmailSecurityInvestigateGetDetectionsResponseResultAttachmentsDetectionEncrypted, AccountEmailSecurityInvestigateGetDetectionsResponseResultAttachmentsDetectionExternal, AccountEmailSecurityInvestigateGetDetectionsResponseResultAttachmentsDetectionUnknown, AccountEmailSecurityInvestigateGetDetectionsResponseResultAttachmentsDetectionNone:
		return true
	}
	return false
}

type AccountEmailSecurityInvestigateGetDetectionsResponseResultHeader struct {
	Name  string                                                               `json:"name,required"`
	Value string                                                               `json:"value,required"`
	JSON  accountEmailSecurityInvestigateGetDetectionsResponseResultHeaderJSON `json:"-"`
}

// accountEmailSecurityInvestigateGetDetectionsResponseResultHeaderJSON contains
// the JSON metadata for the struct
// [AccountEmailSecurityInvestigateGetDetectionsResponseResultHeader]
type accountEmailSecurityInvestigateGetDetectionsResponseResultHeaderJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountEmailSecurityInvestigateGetDetectionsResponseResultHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountEmailSecurityInvestigateGetDetectionsResponseResultHeaderJSON) RawJSON() string {
	return r.raw
}

type AccountEmailSecurityInvestigateGetDetectionsResponseResultLink struct {
	Href string                                                             `json:"href,required"`
	Text string                                                             `json:"text,nullable"`
	JSON accountEmailSecurityInvestigateGetDetectionsResponseResultLinkJSON `json:"-"`
}

// accountEmailSecurityInvestigateGetDetectionsResponseResultLinkJSON contains the
// JSON metadata for the struct
// [AccountEmailSecurityInvestigateGetDetectionsResponseResultLink]
type accountEmailSecurityInvestigateGetDetectionsResponseResultLinkJSON struct {
	Href        apijson.Field
	Text        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountEmailSecurityInvestigateGetDetectionsResponseResultLink) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountEmailSecurityInvestigateGetDetectionsResponseResultLinkJSON) RawJSON() string {
	return r.raw
}

type AccountEmailSecurityInvestigateGetDetectionsResponseResultSenderInfo struct {
	// The name of the autonomous system.
	AsName string `json:"as_name,nullable"`
	// The number of the autonomous system.
	AsNumber int64                                                                    `json:"as_number,nullable"`
	Geo      string                                                                   `json:"geo,nullable"`
	IP       string                                                                   `json:"ip,nullable"`
	Pld      string                                                                   `json:"pld,nullable"`
	JSON     accountEmailSecurityInvestigateGetDetectionsResponseResultSenderInfoJSON `json:"-"`
}

// accountEmailSecurityInvestigateGetDetectionsResponseResultSenderInfoJSON
// contains the JSON metadata for the struct
// [AccountEmailSecurityInvestigateGetDetectionsResponseResultSenderInfo]
type accountEmailSecurityInvestigateGetDetectionsResponseResultSenderInfoJSON struct {
	AsName      apijson.Field
	AsNumber    apijson.Field
	Geo         apijson.Field
	IP          apijson.Field
	Pld         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountEmailSecurityInvestigateGetDetectionsResponseResultSenderInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountEmailSecurityInvestigateGetDetectionsResponseResultSenderInfoJSON) RawJSON() string {
	return r.raw
}

type AccountEmailSecurityInvestigateGetDetectionsResponseResultThreatCategory struct {
	ID          int64                                                                        `json:"id,required"`
	Description string                                                                       `json:"description,nullable"`
	Name        string                                                                       `json:"name,nullable"`
	JSON        accountEmailSecurityInvestigateGetDetectionsResponseResultThreatCategoryJSON `json:"-"`
}

// accountEmailSecurityInvestigateGetDetectionsResponseResultThreatCategoryJSON
// contains the JSON metadata for the struct
// [AccountEmailSecurityInvestigateGetDetectionsResponseResultThreatCategory]
type accountEmailSecurityInvestigateGetDetectionsResponseResultThreatCategoryJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountEmailSecurityInvestigateGetDetectionsResponseResultThreatCategory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountEmailSecurityInvestigateGetDetectionsResponseResultThreatCategoryJSON) RawJSON() string {
	return r.raw
}

type AccountEmailSecurityInvestigateGetDetectionsResponseResultValidation struct {
	Comment string                                                                    `json:"comment,nullable"`
	Dkim    AccountEmailSecurityInvestigateGetDetectionsResponseResultValidationDkim  `json:"dkim,nullable"`
	Dmarc   AccountEmailSecurityInvestigateGetDetectionsResponseResultValidationDmarc `json:"dmarc,nullable"`
	Spf     AccountEmailSecurityInvestigateGetDetectionsResponseResultValidationSpf   `json:"spf,nullable"`
	JSON    accountEmailSecurityInvestigateGetDetectionsResponseResultValidationJSON  `json:"-"`
}

// accountEmailSecurityInvestigateGetDetectionsResponseResultValidationJSON
// contains the JSON metadata for the struct
// [AccountEmailSecurityInvestigateGetDetectionsResponseResultValidation]
type accountEmailSecurityInvestigateGetDetectionsResponseResultValidationJSON struct {
	Comment     apijson.Field
	Dkim        apijson.Field
	Dmarc       apijson.Field
	Spf         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountEmailSecurityInvestigateGetDetectionsResponseResultValidation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountEmailSecurityInvestigateGetDetectionsResponseResultValidationJSON) RawJSON() string {
	return r.raw
}

type AccountEmailSecurityInvestigateGetDetectionsResponseResultValidationDkim string

const (
	AccountEmailSecurityInvestigateGetDetectionsResponseResultValidationDkimPass    AccountEmailSecurityInvestigateGetDetectionsResponseResultValidationDkim = "pass"
	AccountEmailSecurityInvestigateGetDetectionsResponseResultValidationDkimNeutral AccountEmailSecurityInvestigateGetDetectionsResponseResultValidationDkim = "neutral"
	AccountEmailSecurityInvestigateGetDetectionsResponseResultValidationDkimFail    AccountEmailSecurityInvestigateGetDetectionsResponseResultValidationDkim = "fail"
	AccountEmailSecurityInvestigateGetDetectionsResponseResultValidationDkimError   AccountEmailSecurityInvestigateGetDetectionsResponseResultValidationDkim = "error"
	AccountEmailSecurityInvestigateGetDetectionsResponseResultValidationDkimNone    AccountEmailSecurityInvestigateGetDetectionsResponseResultValidationDkim = "none"
)

func (r AccountEmailSecurityInvestigateGetDetectionsResponseResultValidationDkim) IsKnown() bool {
	switch r {
	case AccountEmailSecurityInvestigateGetDetectionsResponseResultValidationDkimPass, AccountEmailSecurityInvestigateGetDetectionsResponseResultValidationDkimNeutral, AccountEmailSecurityInvestigateGetDetectionsResponseResultValidationDkimFail, AccountEmailSecurityInvestigateGetDetectionsResponseResultValidationDkimError, AccountEmailSecurityInvestigateGetDetectionsResponseResultValidationDkimNone:
		return true
	}
	return false
}

type AccountEmailSecurityInvestigateGetDetectionsResponseResultValidationDmarc string

const (
	AccountEmailSecurityInvestigateGetDetectionsResponseResultValidationDmarcPass    AccountEmailSecurityInvestigateGetDetectionsResponseResultValidationDmarc = "pass"
	AccountEmailSecurityInvestigateGetDetectionsResponseResultValidationDmarcNeutral AccountEmailSecurityInvestigateGetDetectionsResponseResultValidationDmarc = "neutral"
	AccountEmailSecurityInvestigateGetDetectionsResponseResultValidationDmarcFail    AccountEmailSecurityInvestigateGetDetectionsResponseResultValidationDmarc = "fail"
	AccountEmailSecurityInvestigateGetDetectionsResponseResultValidationDmarcError   AccountEmailSecurityInvestigateGetDetectionsResponseResultValidationDmarc = "error"
	AccountEmailSecurityInvestigateGetDetectionsResponseResultValidationDmarcNone    AccountEmailSecurityInvestigateGetDetectionsResponseResultValidationDmarc = "none"
)

func (r AccountEmailSecurityInvestigateGetDetectionsResponseResultValidationDmarc) IsKnown() bool {
	switch r {
	case AccountEmailSecurityInvestigateGetDetectionsResponseResultValidationDmarcPass, AccountEmailSecurityInvestigateGetDetectionsResponseResultValidationDmarcNeutral, AccountEmailSecurityInvestigateGetDetectionsResponseResultValidationDmarcFail, AccountEmailSecurityInvestigateGetDetectionsResponseResultValidationDmarcError, AccountEmailSecurityInvestigateGetDetectionsResponseResultValidationDmarcNone:
		return true
	}
	return false
}

type AccountEmailSecurityInvestigateGetDetectionsResponseResultValidationSpf string

const (
	AccountEmailSecurityInvestigateGetDetectionsResponseResultValidationSpfPass    AccountEmailSecurityInvestigateGetDetectionsResponseResultValidationSpf = "pass"
	AccountEmailSecurityInvestigateGetDetectionsResponseResultValidationSpfNeutral AccountEmailSecurityInvestigateGetDetectionsResponseResultValidationSpf = "neutral"
	AccountEmailSecurityInvestigateGetDetectionsResponseResultValidationSpfFail    AccountEmailSecurityInvestigateGetDetectionsResponseResultValidationSpf = "fail"
	AccountEmailSecurityInvestigateGetDetectionsResponseResultValidationSpfError   AccountEmailSecurityInvestigateGetDetectionsResponseResultValidationSpf = "error"
	AccountEmailSecurityInvestigateGetDetectionsResponseResultValidationSpfNone    AccountEmailSecurityInvestigateGetDetectionsResponseResultValidationSpf = "none"
)

func (r AccountEmailSecurityInvestigateGetDetectionsResponseResultValidationSpf) IsKnown() bool {
	switch r {
	case AccountEmailSecurityInvestigateGetDetectionsResponseResultValidationSpfPass, AccountEmailSecurityInvestigateGetDetectionsResponseResultValidationSpfNeutral, AccountEmailSecurityInvestigateGetDetectionsResponseResultValidationSpfFail, AccountEmailSecurityInvestigateGetDetectionsResponseResultValidationSpfError, AccountEmailSecurityInvestigateGetDetectionsResponseResultValidationSpfNone:
		return true
	}
	return false
}

type AccountEmailSecurityInvestigateGetDetectionsResponseResultFinalDisposition string

const (
	AccountEmailSecurityInvestigateGetDetectionsResponseResultFinalDispositionMalicious    AccountEmailSecurityInvestigateGetDetectionsResponseResultFinalDisposition = "MALICIOUS"
	AccountEmailSecurityInvestigateGetDetectionsResponseResultFinalDispositionMaliciousBec AccountEmailSecurityInvestigateGetDetectionsResponseResultFinalDisposition = "MALICIOUS-BEC"
	AccountEmailSecurityInvestigateGetDetectionsResponseResultFinalDispositionSuspicious   AccountEmailSecurityInvestigateGetDetectionsResponseResultFinalDisposition = "SUSPICIOUS"
	AccountEmailSecurityInvestigateGetDetectionsResponseResultFinalDispositionSpoof        AccountEmailSecurityInvestigateGetDetectionsResponseResultFinalDisposition = "SPOOF"
	AccountEmailSecurityInvestigateGetDetectionsResponseResultFinalDispositionSpam         AccountEmailSecurityInvestigateGetDetectionsResponseResultFinalDisposition = "SPAM"
	AccountEmailSecurityInvestigateGetDetectionsResponseResultFinalDispositionBulk         AccountEmailSecurityInvestigateGetDetectionsResponseResultFinalDisposition = "BULK"
	AccountEmailSecurityInvestigateGetDetectionsResponseResultFinalDispositionEncrypted    AccountEmailSecurityInvestigateGetDetectionsResponseResultFinalDisposition = "ENCRYPTED"
	AccountEmailSecurityInvestigateGetDetectionsResponseResultFinalDispositionExternal     AccountEmailSecurityInvestigateGetDetectionsResponseResultFinalDisposition = "EXTERNAL"
	AccountEmailSecurityInvestigateGetDetectionsResponseResultFinalDispositionUnknown      AccountEmailSecurityInvestigateGetDetectionsResponseResultFinalDisposition = "UNKNOWN"
	AccountEmailSecurityInvestigateGetDetectionsResponseResultFinalDispositionNone         AccountEmailSecurityInvestigateGetDetectionsResponseResultFinalDisposition = "NONE"
)

func (r AccountEmailSecurityInvestigateGetDetectionsResponseResultFinalDisposition) IsKnown() bool {
	switch r {
	case AccountEmailSecurityInvestigateGetDetectionsResponseResultFinalDispositionMalicious, AccountEmailSecurityInvestigateGetDetectionsResponseResultFinalDispositionMaliciousBec, AccountEmailSecurityInvestigateGetDetectionsResponseResultFinalDispositionSuspicious, AccountEmailSecurityInvestigateGetDetectionsResponseResultFinalDispositionSpoof, AccountEmailSecurityInvestigateGetDetectionsResponseResultFinalDispositionSpam, AccountEmailSecurityInvestigateGetDetectionsResponseResultFinalDispositionBulk, AccountEmailSecurityInvestigateGetDetectionsResponseResultFinalDispositionEncrypted, AccountEmailSecurityInvestigateGetDetectionsResponseResultFinalDispositionExternal, AccountEmailSecurityInvestigateGetDetectionsResponseResultFinalDispositionUnknown, AccountEmailSecurityInvestigateGetDetectionsResponseResultFinalDispositionNone:
		return true
	}
	return false
}

type AccountEmailSecurityInvestigateGetRawResponse struct {
	Result AccountEmailSecurityInvestigateGetRawResponseResult `json:"result,required"`
	JSON   accountEmailSecurityInvestigateGetRawResponseJSON   `json:"-"`
	APIResponseEmailSecurity
}

// accountEmailSecurityInvestigateGetRawResponseJSON contains the JSON metadata for
// the struct [AccountEmailSecurityInvestigateGetRawResponse]
type accountEmailSecurityInvestigateGetRawResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountEmailSecurityInvestigateGetRawResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountEmailSecurityInvestigateGetRawResponseJSON) RawJSON() string {
	return r.raw
}

type AccountEmailSecurityInvestigateGetRawResponseResult struct {
	// A UTF-8 encoded eml file of the email.
	Raw  string                                                  `json:"raw,required"`
	JSON accountEmailSecurityInvestigateGetRawResponseResultJSON `json:"-"`
}

// accountEmailSecurityInvestigateGetRawResponseResultJSON contains the JSON
// metadata for the struct [AccountEmailSecurityInvestigateGetRawResponseResult]
type accountEmailSecurityInvestigateGetRawResponseResultJSON struct {
	Raw         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountEmailSecurityInvestigateGetRawResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountEmailSecurityInvestigateGetRawResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountEmailSecurityInvestigateGetTraceResponse struct {
	Result AccountEmailSecurityInvestigateGetTraceResponseResult `json:"result,required"`
	JSON   accountEmailSecurityInvestigateGetTraceResponseJSON   `json:"-"`
	APIResponseEmailSecurity
}

// accountEmailSecurityInvestigateGetTraceResponseJSON contains the JSON metadata
// for the struct [AccountEmailSecurityInvestigateGetTraceResponse]
type accountEmailSecurityInvestigateGetTraceResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountEmailSecurityInvestigateGetTraceResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountEmailSecurityInvestigateGetTraceResponseJSON) RawJSON() string {
	return r.raw
}

type AccountEmailSecurityInvestigateGetTraceResponseResult struct {
	Inbound  AccountEmailSecurityInvestigateGetTraceResponseResultInbound  `json:"inbound,required"`
	Outbound AccountEmailSecurityInvestigateGetTraceResponseResultOutbound `json:"outbound,required"`
	JSON     accountEmailSecurityInvestigateGetTraceResponseResultJSON     `json:"-"`
}

// accountEmailSecurityInvestigateGetTraceResponseResultJSON contains the JSON
// metadata for the struct [AccountEmailSecurityInvestigateGetTraceResponseResult]
type accountEmailSecurityInvestigateGetTraceResponseResultJSON struct {
	Inbound     apijson.Field
	Outbound    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountEmailSecurityInvestigateGetTraceResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountEmailSecurityInvestigateGetTraceResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountEmailSecurityInvestigateGetTraceResponseResultInbound struct {
	Lines []TraceLine                                                      `json:"lines,nullable"`
	JSON  accountEmailSecurityInvestigateGetTraceResponseResultInboundJSON `json:"-"`
}

// accountEmailSecurityInvestigateGetTraceResponseResultInboundJSON contains the
// JSON metadata for the struct
// [AccountEmailSecurityInvestigateGetTraceResponseResultInbound]
type accountEmailSecurityInvestigateGetTraceResponseResultInboundJSON struct {
	Lines       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountEmailSecurityInvestigateGetTraceResponseResultInbound) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountEmailSecurityInvestigateGetTraceResponseResultInboundJSON) RawJSON() string {
	return r.raw
}

type AccountEmailSecurityInvestigateGetTraceResponseResultOutbound struct {
	Lines []TraceLine                                                       `json:"lines,nullable"`
	JSON  accountEmailSecurityInvestigateGetTraceResponseResultOutboundJSON `json:"-"`
}

// accountEmailSecurityInvestigateGetTraceResponseResultOutboundJSON contains the
// JSON metadata for the struct
// [AccountEmailSecurityInvestigateGetTraceResponseResultOutbound]
type accountEmailSecurityInvestigateGetTraceResponseResultOutboundJSON struct {
	Lines       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountEmailSecurityInvestigateGetTraceResponseResultOutbound) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountEmailSecurityInvestigateGetTraceResponseResultOutboundJSON) RawJSON() string {
	return r.raw
}

type AccountEmailSecurityInvestigateMoveMultipleResponse struct {
	Result []RetractionResponseItem                                `json:"result,required"`
	JSON   accountEmailSecurityInvestigateMoveMultipleResponseJSON `json:"-"`
	APIResponseEmailSecurity
}

// accountEmailSecurityInvestigateMoveMultipleResponseJSON contains the JSON
// metadata for the struct [AccountEmailSecurityInvestigateMoveMultipleResponse]
type accountEmailSecurityInvestigateMoveMultipleResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountEmailSecurityInvestigateMoveMultipleResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountEmailSecurityInvestigateMoveMultipleResponseJSON) RawJSON() string {
	return r.raw
}

type AccountEmailSecurityInvestigatePreviewResponse struct {
	Result AccountEmailSecurityInvestigatePreviewResponseResult `json:"result,required"`
	JSON   accountEmailSecurityInvestigatePreviewResponseJSON   `json:"-"`
	APIResponseEmailSecurity
}

// accountEmailSecurityInvestigatePreviewResponseJSON contains the JSON metadata
// for the struct [AccountEmailSecurityInvestigatePreviewResponse]
type accountEmailSecurityInvestigatePreviewResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountEmailSecurityInvestigatePreviewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountEmailSecurityInvestigatePreviewResponseJSON) RawJSON() string {
	return r.raw
}

type AccountEmailSecurityInvestigatePreviewResponseResult struct {
	// A base64 encoded PNG image of the email.
	Screenshot string                                                   `json:"screenshot,required"`
	JSON       accountEmailSecurityInvestigatePreviewResponseResultJSON `json:"-"`
}

// accountEmailSecurityInvestigatePreviewResponseResultJSON contains the JSON
// metadata for the struct [AccountEmailSecurityInvestigatePreviewResponseResult]
type accountEmailSecurityInvestigatePreviewResponseResultJSON struct {
	Screenshot  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountEmailSecurityInvestigatePreviewResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountEmailSecurityInvestigatePreviewResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountEmailSecurityInvestigatePreviewMultipleResponse struct {
	Result AccountEmailSecurityInvestigatePreviewMultipleResponseResult `json:"result,required"`
	JSON   accountEmailSecurityInvestigatePreviewMultipleResponseJSON   `json:"-"`
	APIResponseEmailSecurity
}

// accountEmailSecurityInvestigatePreviewMultipleResponseJSON contains the JSON
// metadata for the struct [AccountEmailSecurityInvestigatePreviewMultipleResponse]
type accountEmailSecurityInvestigatePreviewMultipleResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountEmailSecurityInvestigatePreviewMultipleResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountEmailSecurityInvestigatePreviewMultipleResponseJSON) RawJSON() string {
	return r.raw
}

type AccountEmailSecurityInvestigatePreviewMultipleResponseResult struct {
	// A base64 encoded PNG image of the email.
	Screenshot string                                                           `json:"screenshot,required"`
	JSON       accountEmailSecurityInvestigatePreviewMultipleResponseResultJSON `json:"-"`
}

// accountEmailSecurityInvestigatePreviewMultipleResponseResultJSON contains the
// JSON metadata for the struct
// [AccountEmailSecurityInvestigatePreviewMultipleResponseResult]
type accountEmailSecurityInvestigatePreviewMultipleResponseResultJSON struct {
	Screenshot  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountEmailSecurityInvestigatePreviewMultipleResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountEmailSecurityInvestigatePreviewMultipleResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountEmailSecurityInvestigateReclassifyResponse struct {
	Result interface{}                                           `json:"result,required"`
	JSON   accountEmailSecurityInvestigateReclassifyResponseJSON `json:"-"`
	APIResponseEmailSecurity
}

// accountEmailSecurityInvestigateReclassifyResponseJSON contains the JSON metadata
// for the struct [AccountEmailSecurityInvestigateReclassifyResponse]
type accountEmailSecurityInvestigateReclassifyResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountEmailSecurityInvestigateReclassifyResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountEmailSecurityInvestigateReclassifyResponseJSON) RawJSON() string {
	return r.raw
}

type AccountEmailSecurityInvestigateReleaseResponse struct {
	Result []AccountEmailSecurityInvestigateReleaseResponseResult `json:"result,required"`
	JSON   accountEmailSecurityInvestigateReleaseResponseJSON     `json:"-"`
	APIResponseEmailSecurity
}

// accountEmailSecurityInvestigateReleaseResponseJSON contains the JSON metadata
// for the struct [AccountEmailSecurityInvestigateReleaseResponse]
type accountEmailSecurityInvestigateReleaseResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountEmailSecurityInvestigateReleaseResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountEmailSecurityInvestigateReleaseResponseJSON) RawJSON() string {
	return r.raw
}

type AccountEmailSecurityInvestigateReleaseResponseResult struct {
	// The identifier of the message.
	PostfixID   string                                                   `json:"postfix_id,required"`
	Delivered   []string                                                 `json:"delivered,nullable"`
	Failed      []string                                                 `json:"failed,nullable"`
	Undelivered []string                                                 `json:"undelivered,nullable"`
	JSON        accountEmailSecurityInvestigateReleaseResponseResultJSON `json:"-"`
}

// accountEmailSecurityInvestigateReleaseResponseResultJSON contains the JSON
// metadata for the struct [AccountEmailSecurityInvestigateReleaseResponseResult]
type accountEmailSecurityInvestigateReleaseResponseResultJSON struct {
	PostfixID   apijson.Field
	Delivered   apijson.Field
	Failed      apijson.Field
	Undelivered apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountEmailSecurityInvestigateReleaseResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountEmailSecurityInvestigateReleaseResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountEmailSecurityInvestigateListParams struct {
	// Determines if the message action log is included in the response.
	ActionLog param.Field[bool]   `query:"action_log"`
	AlertID   param.Field[string] `query:"alert_id"`
	// Determines if the search results will include detections or not.
	DetectionsOnly param.Field[bool] `query:"detections_only"`
	// The sender domains the search filters by.
	Domain param.Field[string] `query:"domain"`
	// The end of the search date range. Defaults to `now`.
	End param.Field[time.Time] `query:"end" format:"date-time"`
	// The dispositions the search filters by.
	FinalDisposition param.Field[AccountEmailSecurityInvestigateListParamsFinalDisposition] `query:"final_disposition"`
	// The message actions the search filters by.
	MessageAction param.Field[AccountEmailSecurityInvestigateListParamsMessageAction] `query:"message_action"`
	MessageID     param.Field[string]                                                 `query:"message_id"`
	Metric        param.Field[string]                                                 `query:"metric"`
	// The page number of paginated results.
	Page param.Field[int64] `query:"page"`
	// The number of results per page.
	PerPage param.Field[int64] `query:"per_page"`
	// The space-delimited term used in the query. The search is case-insensitive.
	//
	// The content of the following email metadata fields are searched:
	//
	// - alert_id
	// - CC
	// - From (envelope_from)
	// - From Name
	// - final_disposition
	// - md5 hash (of any attachment)
	// - sha1 hash (of any attachment)
	// - sha256 hash (of any attachment)
	// - name (of any attachment)
	// - Reason
	// - Received DateTime (yyyy-mm-ddThh:mm:ss)
	// - Sent DateTime (yyyy-mm-ddThh:mm:ss)
	// - ReplyTo
	// - To (envelope_to)
	// - To Name
	// - Message-ID
	// - smtp_helo_server_ip
	// - smtp_previous_hop_ip
	// - x_originating_ip
	// - Subject
	Query     param.Field[string] `query:"query"`
	Recipient param.Field[string] `query:"recipient"`
	Sender    param.Field[string] `query:"sender"`
	// The beginning of the search date range. Defaults to `now - 30 days`.
	Start param.Field[time.Time] `query:"start" format:"date-time"`
}

// URLQuery serializes [AccountEmailSecurityInvestigateListParams]'s query
// parameters as `url.Values`.
func (r AccountEmailSecurityInvestigateListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The dispositions the search filters by.
type AccountEmailSecurityInvestigateListParamsFinalDisposition string

const (
	AccountEmailSecurityInvestigateListParamsFinalDispositionMalicious  AccountEmailSecurityInvestigateListParamsFinalDisposition = "MALICIOUS"
	AccountEmailSecurityInvestigateListParamsFinalDispositionSuspicious AccountEmailSecurityInvestigateListParamsFinalDisposition = "SUSPICIOUS"
	AccountEmailSecurityInvestigateListParamsFinalDispositionSpoof      AccountEmailSecurityInvestigateListParamsFinalDisposition = "SPOOF"
	AccountEmailSecurityInvestigateListParamsFinalDispositionSpam       AccountEmailSecurityInvestigateListParamsFinalDisposition = "SPAM"
	AccountEmailSecurityInvestigateListParamsFinalDispositionBulk       AccountEmailSecurityInvestigateListParamsFinalDisposition = "BULK"
	AccountEmailSecurityInvestigateListParamsFinalDispositionNone       AccountEmailSecurityInvestigateListParamsFinalDisposition = "NONE"
)

func (r AccountEmailSecurityInvestigateListParamsFinalDisposition) IsKnown() bool {
	switch r {
	case AccountEmailSecurityInvestigateListParamsFinalDispositionMalicious, AccountEmailSecurityInvestigateListParamsFinalDispositionSuspicious, AccountEmailSecurityInvestigateListParamsFinalDispositionSpoof, AccountEmailSecurityInvestigateListParamsFinalDispositionSpam, AccountEmailSecurityInvestigateListParamsFinalDispositionBulk, AccountEmailSecurityInvestigateListParamsFinalDispositionNone:
		return true
	}
	return false
}

// The message actions the search filters by.
type AccountEmailSecurityInvestigateListParamsMessageAction string

const (
	AccountEmailSecurityInvestigateListParamsMessageActionPreview            AccountEmailSecurityInvestigateListParamsMessageAction = "PREVIEW"
	AccountEmailSecurityInvestigateListParamsMessageActionQuarantineReleased AccountEmailSecurityInvestigateListParamsMessageAction = "QUARANTINE_RELEASED"
	AccountEmailSecurityInvestigateListParamsMessageActionMoved              AccountEmailSecurityInvestigateListParamsMessageAction = "MOVED"
)

func (r AccountEmailSecurityInvestigateListParamsMessageAction) IsKnown() bool {
	switch r {
	case AccountEmailSecurityInvestigateListParamsMessageActionPreview, AccountEmailSecurityInvestigateListParamsMessageActionQuarantineReleased, AccountEmailSecurityInvestigateListParamsMessageActionMoved:
		return true
	}
	return false
}

type AccountEmailSecurityInvestigateMoveMultipleParams struct {
	Destination param.Field[AccountEmailSecurityInvestigateMoveMultipleParamsDestination] `json:"destination,required"`
	PostfixIDs  param.Field[[]string]                                                     `json:"postfix_ids,required"`
}

func (r AccountEmailSecurityInvestigateMoveMultipleParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountEmailSecurityInvestigateMoveMultipleParamsDestination string

const (
	AccountEmailSecurityInvestigateMoveMultipleParamsDestinationInbox                     AccountEmailSecurityInvestigateMoveMultipleParamsDestination = "Inbox"
	AccountEmailSecurityInvestigateMoveMultipleParamsDestinationJunkEmail                 AccountEmailSecurityInvestigateMoveMultipleParamsDestination = "JunkEmail"
	AccountEmailSecurityInvestigateMoveMultipleParamsDestinationDeletedItems              AccountEmailSecurityInvestigateMoveMultipleParamsDestination = "DeletedItems"
	AccountEmailSecurityInvestigateMoveMultipleParamsDestinationRecoverableItemsDeletions AccountEmailSecurityInvestigateMoveMultipleParamsDestination = "RecoverableItemsDeletions"
	AccountEmailSecurityInvestigateMoveMultipleParamsDestinationRecoverableItemsPurges    AccountEmailSecurityInvestigateMoveMultipleParamsDestination = "RecoverableItemsPurges"
)

func (r AccountEmailSecurityInvestigateMoveMultipleParamsDestination) IsKnown() bool {
	switch r {
	case AccountEmailSecurityInvestigateMoveMultipleParamsDestinationInbox, AccountEmailSecurityInvestigateMoveMultipleParamsDestinationJunkEmail, AccountEmailSecurityInvestigateMoveMultipleParamsDestinationDeletedItems, AccountEmailSecurityInvestigateMoveMultipleParamsDestinationRecoverableItemsDeletions, AccountEmailSecurityInvestigateMoveMultipleParamsDestinationRecoverableItemsPurges:
		return true
	}
	return false
}

type AccountEmailSecurityInvestigatePreviewMultipleParams struct {
	// The identifier of the message.
	PostfixID param.Field[string] `json:"postfix_id,required"`
}

func (r AccountEmailSecurityInvestigatePreviewMultipleParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountEmailSecurityInvestigateReclassifyParams struct {
	ExpectedDisposition param.Field[AccountEmailSecurityInvestigateReclassifyParamsExpectedDisposition] `json:"expected_disposition,required"`
	// Base64 encoded content of the EML file
	EmlContent param.Field[string] `json:"eml_content"`
}

func (r AccountEmailSecurityInvestigateReclassifyParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountEmailSecurityInvestigateReclassifyParamsExpectedDisposition string

const (
	AccountEmailSecurityInvestigateReclassifyParamsExpectedDispositionNone       AccountEmailSecurityInvestigateReclassifyParamsExpectedDisposition = "NONE"
	AccountEmailSecurityInvestigateReclassifyParamsExpectedDispositionBulk       AccountEmailSecurityInvestigateReclassifyParamsExpectedDisposition = "BULK"
	AccountEmailSecurityInvestigateReclassifyParamsExpectedDispositionMalicious  AccountEmailSecurityInvestigateReclassifyParamsExpectedDisposition = "MALICIOUS"
	AccountEmailSecurityInvestigateReclassifyParamsExpectedDispositionSpam       AccountEmailSecurityInvestigateReclassifyParamsExpectedDisposition = "SPAM"
	AccountEmailSecurityInvestigateReclassifyParamsExpectedDispositionSpoof      AccountEmailSecurityInvestigateReclassifyParamsExpectedDisposition = "SPOOF"
	AccountEmailSecurityInvestigateReclassifyParamsExpectedDispositionSuspicious AccountEmailSecurityInvestigateReclassifyParamsExpectedDisposition = "SUSPICIOUS"
)

func (r AccountEmailSecurityInvestigateReclassifyParamsExpectedDisposition) IsKnown() bool {
	switch r {
	case AccountEmailSecurityInvestigateReclassifyParamsExpectedDispositionNone, AccountEmailSecurityInvestigateReclassifyParamsExpectedDispositionBulk, AccountEmailSecurityInvestigateReclassifyParamsExpectedDispositionMalicious, AccountEmailSecurityInvestigateReclassifyParamsExpectedDispositionSpam, AccountEmailSecurityInvestigateReclassifyParamsExpectedDispositionSpoof, AccountEmailSecurityInvestigateReclassifyParamsExpectedDispositionSuspicious:
		return true
	}
	return false
}

type AccountEmailSecurityInvestigateReleaseParams struct {
	// A list of messages identfied by their `postfix_id`s that should be released.
	Body []string `json:"body,required"`
}

func (r AccountEmailSecurityInvestigateReleaseParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}
