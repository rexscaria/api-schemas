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

// AccountBrandProtectionService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountBrandProtectionService] method instead.
type AccountBrandProtectionService struct {
	Options []option.RequestOption
}

// NewAccountBrandProtectionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountBrandProtectionService(opts ...option.RequestOption) (r *AccountBrandProtectionService) {
	r = &AccountBrandProtectionService{}
	r.Options = opts
	return
}

// Gets phishing details about a URL.
func (r *AccountBrandProtectionService) GetURLInfo(ctx context.Context, accountID string, query AccountBrandProtectionGetURLInfoParams, opts ...option.RequestOption) (res *AccountBrandProtectionGetURLInfoResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/brand-protection/url-info", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Submit suspicious URL for scanning
func (r *AccountBrandProtectionService) SubmitURL(ctx context.Context, accountID string, body AccountBrandProtectionSubmitURLParams, opts ...option.RequestOption) (res *AccountBrandProtectionSubmitURLResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/brand-protection/submit", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type AccountBrandProtectionGetURLInfoResponse struct {
	Result AccountBrandProtectionGetURLInfoResponseResult `json:"result"`
	JSON   accountBrandProtectionGetURLInfoResponseJSON   `json:"-"`
	SingleResponseIntel
}

// accountBrandProtectionGetURLInfoResponseJSON contains the JSON metadata for the
// struct [AccountBrandProtectionGetURLInfoResponse]
type accountBrandProtectionGetURLInfoResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountBrandProtectionGetURLInfoResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountBrandProtectionGetURLInfoResponseJSON) RawJSON() string {
	return r.raw
}

type AccountBrandProtectionGetURLInfoResponseResult struct {
	// List of categorizations applied to this submission.
	Categorizations []AccountBrandProtectionGetURLInfoResponseResultCategorization `json:"categorizations"`
	// List of model results for completed scans.
	ModelResults []AccountBrandProtectionGetURLInfoResponseResultModelResult `json:"model_results"`
	// List of signatures that matched against site content found when crawling the
	// URL.
	RuleMatches []AccountBrandProtectionGetURLInfoResponseResultRuleMatch `json:"rule_matches"`
	// Status of the most recent scan found.
	ScanStatus AccountBrandProtectionGetURLInfoResponseResultScanStatus `json:"scan_status"`
	// For internal use.
	ScreenshotDownloadSignature string `json:"screenshot_download_signature"`
	// For internal use.
	ScreenshotPath string `json:"screenshot_path"`
	// URL that was submitted.
	URL  string                                             `json:"url"`
	JSON accountBrandProtectionGetURLInfoResponseResultJSON `json:"-"`
}

// accountBrandProtectionGetURLInfoResponseResultJSON contains the JSON metadata
// for the struct [AccountBrandProtectionGetURLInfoResponseResult]
type accountBrandProtectionGetURLInfoResponseResultJSON struct {
	Categorizations             apijson.Field
	ModelResults                apijson.Field
	RuleMatches                 apijson.Field
	ScanStatus                  apijson.Field
	ScreenshotDownloadSignature apijson.Field
	ScreenshotPath              apijson.Field
	URL                         apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *AccountBrandProtectionGetURLInfoResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountBrandProtectionGetURLInfoResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountBrandProtectionGetURLInfoResponseResultCategorization struct {
	// Name of the category applied.
	Category string `json:"category"`
	// Result of human review for this categorization.
	VerificationStatus string                                                           `json:"verification_status"`
	JSON               accountBrandProtectionGetURLInfoResponseResultCategorizationJSON `json:"-"`
}

// accountBrandProtectionGetURLInfoResponseResultCategorizationJSON contains the
// JSON metadata for the struct
// [AccountBrandProtectionGetURLInfoResponseResultCategorization]
type accountBrandProtectionGetURLInfoResponseResultCategorizationJSON struct {
	Category           apijson.Field
	VerificationStatus apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccountBrandProtectionGetURLInfoResponseResultCategorization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountBrandProtectionGetURLInfoResponseResultCategorizationJSON) RawJSON() string {
	return r.raw
}

type AccountBrandProtectionGetURLInfoResponseResultModelResult struct {
	// Name of the model.
	ModelName string `json:"model_name"`
	// Score output by the model for this submission.
	ModelScore float64                                                       `json:"model_score"`
	JSON       accountBrandProtectionGetURLInfoResponseResultModelResultJSON `json:"-"`
}

// accountBrandProtectionGetURLInfoResponseResultModelResultJSON contains the JSON
// metadata for the struct
// [AccountBrandProtectionGetURLInfoResponseResultModelResult]
type accountBrandProtectionGetURLInfoResponseResultModelResultJSON struct {
	ModelName   apijson.Field
	ModelScore  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountBrandProtectionGetURLInfoResponseResultModelResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountBrandProtectionGetURLInfoResponseResultModelResultJSON) RawJSON() string {
	return r.raw
}

type AccountBrandProtectionGetURLInfoResponseResultRuleMatch struct {
	// For internal use.
	Banning bool `json:"banning"`
	// For internal use.
	Blocking bool `json:"blocking"`
	// Description of the signature that matched.
	Description string `json:"description"`
	// Name of the signature that matched.
	Name string                                                      `json:"name"`
	JSON accountBrandProtectionGetURLInfoResponseResultRuleMatchJSON `json:"-"`
}

// accountBrandProtectionGetURLInfoResponseResultRuleMatchJSON contains the JSON
// metadata for the struct
// [AccountBrandProtectionGetURLInfoResponseResultRuleMatch]
type accountBrandProtectionGetURLInfoResponseResultRuleMatchJSON struct {
	Banning     apijson.Field
	Blocking    apijson.Field
	Description apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountBrandProtectionGetURLInfoResponseResultRuleMatch) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountBrandProtectionGetURLInfoResponseResultRuleMatchJSON) RawJSON() string {
	return r.raw
}

// Status of the most recent scan found.
type AccountBrandProtectionGetURLInfoResponseResultScanStatus struct {
	// Timestamp of when the submission was processed.
	LastProcessed string `json:"last_processed"`
	// For internal use.
	ScanComplete bool `json:"scan_complete"`
	// Status code that the crawler received when loading the submitted URL.
	StatusCode int64 `json:"status_code"`
	// ID of the most recent submission.
	SubmissionID int64                                                        `json:"submission_id"`
	JSON         accountBrandProtectionGetURLInfoResponseResultScanStatusJSON `json:"-"`
}

// accountBrandProtectionGetURLInfoResponseResultScanStatusJSON contains the JSON
// metadata for the struct
// [AccountBrandProtectionGetURLInfoResponseResultScanStatus]
type accountBrandProtectionGetURLInfoResponseResultScanStatusJSON struct {
	LastProcessed apijson.Field
	ScanComplete  apijson.Field
	StatusCode    apijson.Field
	SubmissionID  apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccountBrandProtectionGetURLInfoResponseResultScanStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountBrandProtectionGetURLInfoResponseResultScanStatusJSON) RawJSON() string {
	return r.raw
}

type AccountBrandProtectionSubmitURLResponse struct {
	Result AccountBrandProtectionSubmitURLResponseResult `json:"result"`
	JSON   accountBrandProtectionSubmitURLResponseJSON   `json:"-"`
	SingleResponseIntel
}

// accountBrandProtectionSubmitURLResponseJSON contains the JSON metadata for the
// struct [AccountBrandProtectionSubmitURLResponse]
type accountBrandProtectionSubmitURLResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountBrandProtectionSubmitURLResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountBrandProtectionSubmitURLResponseJSON) RawJSON() string {
	return r.raw
}

type AccountBrandProtectionSubmitURLResponseResult struct {
	// URLs that were excluded from scanning because their domain is in our no-scan
	// list.
	ExcludedURLs []AccountBrandProtectionSubmitURLResponseResultExcludedURL `json:"excluded_urls"`
	// URLs that were skipped because the same URL is currently being scanned
	SkippedURLs []AccountBrandProtectionSubmitURLResponseResultSkippedURL `json:"skipped_urls"`
	// URLs that were successfully submitted for scanning.
	SubmittedURLs []AccountBrandProtectionSubmitURLResponseResultSubmittedURL `json:"submitted_urls"`
	JSON          accountBrandProtectionSubmitURLResponseResultJSON           `json:"-"`
}

// accountBrandProtectionSubmitURLResponseResultJSON contains the JSON metadata for
// the struct [AccountBrandProtectionSubmitURLResponseResult]
type accountBrandProtectionSubmitURLResponseResultJSON struct {
	ExcludedURLs  apijson.Field
	SkippedURLs   apijson.Field
	SubmittedURLs apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccountBrandProtectionSubmitURLResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountBrandProtectionSubmitURLResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountBrandProtectionSubmitURLResponseResultExcludedURL struct {
	// URL that was excluded.
	URL  string                                                       `json:"url"`
	JSON accountBrandProtectionSubmitURLResponseResultExcludedURLJSON `json:"-"`
}

// accountBrandProtectionSubmitURLResponseResultExcludedURLJSON contains the JSON
// metadata for the struct
// [AccountBrandProtectionSubmitURLResponseResultExcludedURL]
type accountBrandProtectionSubmitURLResponseResultExcludedURLJSON struct {
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountBrandProtectionSubmitURLResponseResultExcludedURL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountBrandProtectionSubmitURLResponseResultExcludedURLJSON) RawJSON() string {
	return r.raw
}

type AccountBrandProtectionSubmitURLResponseResultSkippedURL struct {
	// URL that was skipped.
	URL string `json:"url"`
	// ID of the submission of that URL that is currently scanning.
	URLID int64                                                       `json:"url_id"`
	JSON  accountBrandProtectionSubmitURLResponseResultSkippedURLJSON `json:"-"`
}

// accountBrandProtectionSubmitURLResponseResultSkippedURLJSON contains the JSON
// metadata for the struct
// [AccountBrandProtectionSubmitURLResponseResultSkippedURL]
type accountBrandProtectionSubmitURLResponseResultSkippedURLJSON struct {
	URL         apijson.Field
	URLID       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountBrandProtectionSubmitURLResponseResultSkippedURL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountBrandProtectionSubmitURLResponseResultSkippedURLJSON) RawJSON() string {
	return r.raw
}

type AccountBrandProtectionSubmitURLResponseResultSubmittedURL struct {
	// URL that was submitted.
	URL string `json:"url"`
	// ID assigned to this URL submission. Used to retrieve scanning results.
	URLID int64                                                         `json:"url_id"`
	JSON  accountBrandProtectionSubmitURLResponseResultSubmittedURLJSON `json:"-"`
}

// accountBrandProtectionSubmitURLResponseResultSubmittedURLJSON contains the JSON
// metadata for the struct
// [AccountBrandProtectionSubmitURLResponseResultSubmittedURL]
type accountBrandProtectionSubmitURLResponseResultSubmittedURLJSON struct {
	URL         apijson.Field
	URLID       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountBrandProtectionSubmitURLResponseResultSubmittedURL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountBrandProtectionSubmitURLResponseResultSubmittedURLJSON) RawJSON() string {
	return r.raw
}

type AccountBrandProtectionGetURLInfoParams struct {
	// Submission URL(s) to filter submission results by.
	URL param.Field[[]string] `query:"url"`
	// Submission ID(s) to filter submission results by.
	URLID param.Field[[]int64] `query:"url_id"`
}

// URLQuery serializes [AccountBrandProtectionGetURLInfoParams]'s query parameters
// as `url.Values`.
func (r AccountBrandProtectionGetURLInfoParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountBrandProtectionSubmitURLParams struct {
	// URL(s) to filter submissions results by
	URL param.Field[string] `json:"url" format:"uri"`
}

func (r AccountBrandProtectionSubmitURLParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
