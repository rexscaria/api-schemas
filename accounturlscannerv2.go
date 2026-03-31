// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/apiquery"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
)

// AccountUrlscannerV2Service contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountUrlscannerV2Service] method instead.
type AccountUrlscannerV2Service struct {
	Options     []option.RequestOption
	Screenshots *AccountUrlscannerV2ScreenshotService
}

// NewAccountUrlscannerV2Service generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountUrlscannerV2Service(opts ...option.RequestOption) (r *AccountUrlscannerV2Service) {
	r = &AccountUrlscannerV2Service{}
	r.Options = opts
	r.Screenshots = NewAccountUrlscannerV2ScreenshotService(opts...)
	return
}

// Submit URLs to scan. Check limits at
// https://developers.cloudflare.com/security-center/investigate/scan-limits/ and
// take into account scans submitted in bulk have lower priority and may take
// longer to finish.
func (r *AccountUrlscannerV2Service) BulkNewScans(ctx context.Context, accountID string, body AccountUrlscannerV2BulkNewScansParams, opts ...option.RequestOption) (res *[]AccountUrlscannerV2BulkNewScansResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/urlscanner/v2/bulk", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Submit a URL to scan. Check limits at
// https://developers.cloudflare.com/security-center/investigate/scan-limits/.
func (r *AccountUrlscannerV2Service) NewScan(ctx context.Context, accountID string, body AccountUrlscannerV2NewScanParams, opts ...option.RequestOption) (res *AccountUrlscannerV2NewScanResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/urlscanner/v2/scan", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Returns a plain text response, with the scan's DOM content as rendered by
// Chrome.
func (r *AccountUrlscannerV2Service) GetDom(ctx context.Context, accountID string, scanID string, opts ...option.RequestOption) (res *string, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if scanID == "" {
		err = errors.New("missing required scan_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/urlscanner/v2/dom/%s", accountID, scanID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Get a URL scan's HAR file. See HAR spec at
// http://www.softwareishard.com/blog/har-12-spec/.
func (r *AccountUrlscannerV2Service) GetHar(ctx context.Context, accountID string, scanID string, opts ...option.RequestOption) (res *AccountUrlscannerV2GetHarResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if scanID == "" {
		err = errors.New("missing required scan_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/urlscanner/v2/har/%s", accountID, scanID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Returns the raw response of the network request. Find the `response_id` in the
// `data.requests.response.hash`.
func (r *AccountUrlscannerV2Service) GetRawResponse(ctx context.Context, accountID string, responseID string, opts ...option.RequestOption) (res *string, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if responseID == "" {
		err = errors.New("missing required response_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/urlscanner/v2/responses/%s", accountID, responseID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Get URL scan by uuid
func (r *AccountUrlscannerV2Service) GetScan(ctx context.Context, accountID string, scanID string, opts ...option.RequestOption) (res *AccountUrlscannerV2GetScanResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if scanID == "" {
		err = errors.New("missing required scan_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/urlscanner/v2/result/%s", accountID, scanID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Use a subset of ElasticSearch Query syntax to filter scans. Some example
// queries:<br/> <br/>- 'path:"/bundles/jquery.js"': Searches for scans who
// requested resources with the given path.<br/>- 'page.asn:AS24940 AND hash:xxx':
// Websites hosted in AS24940 where a resource with the given hash was
// downloaded.<br/>- 'page.domain:microsoft\* AND verdicts.malicious:true AND NOT
// page.domain:microsoft.com': malicious scans whose hostname starts with
// "microsoft".<br/>- 'apikey:me AND date:[2025-01 TO 2025-02]': my scans from 2025
// January to 2025 February.
func (r *AccountUrlscannerV2Service) SearchScans(ctx context.Context, accountID string, query AccountUrlscannerV2SearchScansParams, opts ...option.RequestOption) (res *AccountUrlscannerV2SearchScansResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/urlscanner/v2/search", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type AccountUrlscannerV2BulkNewScansResponse struct {
	// URL to api report.
	API string `json:"api" api:"required"`
	// URL to report.
	Result string `json:"result" api:"required"`
	// Submitted URL
	URL string `json:"url" api:"required"`
	// Scan ID.
	Uuid string `json:"uuid" api:"required" format:"uuid"`
	// Submitted visibility status.
	Visibility AccountUrlscannerV2BulkNewScansResponseVisibility `json:"visibility" api:"required"`
	Options    AccountUrlscannerV2BulkNewScansResponseOptions    `json:"options"`
	JSON       accountUrlscannerV2BulkNewScansResponseJSON       `json:"-"`
}

// accountUrlscannerV2BulkNewScansResponseJSON contains the JSON metadata for the
// struct [AccountUrlscannerV2BulkNewScansResponse]
type accountUrlscannerV2BulkNewScansResponseJSON struct {
	API         apijson.Field
	Result      apijson.Field
	URL         apijson.Field
	Uuid        apijson.Field
	Visibility  apijson.Field
	Options     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUrlscannerV2BulkNewScansResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerV2BulkNewScansResponseJSON) RawJSON() string {
	return r.raw
}

// Submitted visibility status.
type AccountUrlscannerV2BulkNewScansResponseVisibility string

const (
	AccountUrlscannerV2BulkNewScansResponseVisibilityPublic   AccountUrlscannerV2BulkNewScansResponseVisibility = "public"
	AccountUrlscannerV2BulkNewScansResponseVisibilityUnlisted AccountUrlscannerV2BulkNewScansResponseVisibility = "unlisted"
)

func (r AccountUrlscannerV2BulkNewScansResponseVisibility) IsKnown() bool {
	switch r {
	case AccountUrlscannerV2BulkNewScansResponseVisibilityPublic, AccountUrlscannerV2BulkNewScansResponseVisibilityUnlisted:
		return true
	}
	return false
}

type AccountUrlscannerV2BulkNewScansResponseOptions struct {
	Useragent string                                             `json:"useragent"`
	JSON      accountUrlscannerV2BulkNewScansResponseOptionsJSON `json:"-"`
}

// accountUrlscannerV2BulkNewScansResponseOptionsJSON contains the JSON metadata
// for the struct [AccountUrlscannerV2BulkNewScansResponseOptions]
type accountUrlscannerV2BulkNewScansResponseOptionsJSON struct {
	Useragent   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUrlscannerV2BulkNewScansResponseOptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerV2BulkNewScansResponseOptionsJSON) RawJSON() string {
	return r.raw
}

type AccountUrlscannerV2NewScanResponse struct {
	// URL to api report.
	API     string `json:"api" api:"required"`
	Message string `json:"message" api:"required"`
	// Public URL to report.
	Result string `json:"result" api:"required"`
	// Canonical form of submitted URL. Use this if you want to later search by URL.
	URL string `json:"url" api:"required"`
	// Scan ID.
	Uuid string `json:"uuid" api:"required" format:"uuid"`
	// Submitted visibility status.
	Visibility AccountUrlscannerV2NewScanResponseVisibility `json:"visibility" api:"required"`
	Options    AccountUrlscannerV2NewScanResponseOptions    `json:"options"`
	JSON       accountUrlscannerV2NewScanResponseJSON       `json:"-"`
}

// accountUrlscannerV2NewScanResponseJSON contains the JSON metadata for the struct
// [AccountUrlscannerV2NewScanResponse]
type accountUrlscannerV2NewScanResponseJSON struct {
	API         apijson.Field
	Message     apijson.Field
	Result      apijson.Field
	URL         apijson.Field
	Uuid        apijson.Field
	Visibility  apijson.Field
	Options     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUrlscannerV2NewScanResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerV2NewScanResponseJSON) RawJSON() string {
	return r.raw
}

// Submitted visibility status.
type AccountUrlscannerV2NewScanResponseVisibility string

const (
	AccountUrlscannerV2NewScanResponseVisibilityPublic   AccountUrlscannerV2NewScanResponseVisibility = "public"
	AccountUrlscannerV2NewScanResponseVisibilityUnlisted AccountUrlscannerV2NewScanResponseVisibility = "unlisted"
)

func (r AccountUrlscannerV2NewScanResponseVisibility) IsKnown() bool {
	switch r {
	case AccountUrlscannerV2NewScanResponseVisibilityPublic, AccountUrlscannerV2NewScanResponseVisibilityUnlisted:
		return true
	}
	return false
}

type AccountUrlscannerV2NewScanResponseOptions struct {
	Useragent string                                        `json:"useragent"`
	JSON      accountUrlscannerV2NewScanResponseOptionsJSON `json:"-"`
}

// accountUrlscannerV2NewScanResponseOptionsJSON contains the JSON metadata for the
// struct [AccountUrlscannerV2NewScanResponseOptions]
type accountUrlscannerV2NewScanResponseOptionsJSON struct {
	Useragent   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUrlscannerV2NewScanResponseOptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerV2NewScanResponseOptionsJSON) RawJSON() string {
	return r.raw
}

type AccountUrlscannerV2GetHarResponse struct {
	Log  AccountUrlscannerV2GetHarResponseLog  `json:"log" api:"required"`
	JSON accountUrlscannerV2GetHarResponseJSON `json:"-"`
}

// accountUrlscannerV2GetHarResponseJSON contains the JSON metadata for the struct
// [AccountUrlscannerV2GetHarResponse]
type accountUrlscannerV2GetHarResponseJSON struct {
	Log         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUrlscannerV2GetHarResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerV2GetHarResponseJSON) RawJSON() string {
	return r.raw
}

type AccountUrlscannerV2GetHarResponseLog struct {
	Creator AccountUrlscannerV2GetHarResponseLogCreator `json:"creator" api:"required"`
	Entries []AccountUrlscannerV2GetHarResponseLogEntry `json:"entries" api:"required"`
	Pages   []AccountUrlscannerV2GetHarResponseLogPage  `json:"pages" api:"required"`
	Version string                                      `json:"version" api:"required"`
	JSON    accountUrlscannerV2GetHarResponseLogJSON    `json:"-"`
}

// accountUrlscannerV2GetHarResponseLogJSON contains the JSON metadata for the
// struct [AccountUrlscannerV2GetHarResponseLog]
type accountUrlscannerV2GetHarResponseLogJSON struct {
	Creator     apijson.Field
	Entries     apijson.Field
	Pages       apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUrlscannerV2GetHarResponseLog) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerV2GetHarResponseLogJSON) RawJSON() string {
	return r.raw
}

type AccountUrlscannerV2GetHarResponseLogCreator struct {
	Comment string                                          `json:"comment" api:"required"`
	Name    string                                          `json:"name" api:"required"`
	Version string                                          `json:"version" api:"required"`
	JSON    accountUrlscannerV2GetHarResponseLogCreatorJSON `json:"-"`
}

// accountUrlscannerV2GetHarResponseLogCreatorJSON contains the JSON metadata for
// the struct [AccountUrlscannerV2GetHarResponseLogCreator]
type accountUrlscannerV2GetHarResponseLogCreatorJSON struct {
	Comment     apijson.Field
	Name        apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUrlscannerV2GetHarResponseLogCreator) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerV2GetHarResponseLogCreatorJSON) RawJSON() string {
	return r.raw
}

type AccountUrlscannerV2GetHarResponseLogEntry struct {
	InitialPriority string                                              `json:"_initialPriority" api:"required"`
	InitiatorType   string                                              `json:"_initiator_type" api:"required"`
	Priority        string                                              `json:"_priority" api:"required"`
	RequestID       string                                              `json:"_requestId" api:"required"`
	RequestTime     float64                                             `json:"_requestTime" api:"required"`
	ResourceType    string                                              `json:"_resourceType" api:"required"`
	Cache           interface{}                                         `json:"cache" api:"required"`
	Connection      string                                              `json:"connection" api:"required"`
	Pageref         string                                              `json:"pageref" api:"required"`
	Request         AccountUrlscannerV2GetHarResponseLogEntriesRequest  `json:"request" api:"required"`
	Response        AccountUrlscannerV2GetHarResponseLogEntriesResponse `json:"response" api:"required"`
	ServerIPAddress string                                              `json:"serverIPAddress" api:"required"`
	StartedDateTime string                                              `json:"startedDateTime" api:"required"`
	Time            float64                                             `json:"time" api:"required"`
	JSON            accountUrlscannerV2GetHarResponseLogEntryJSON       `json:"-"`
}

// accountUrlscannerV2GetHarResponseLogEntryJSON contains the JSON metadata for the
// struct [AccountUrlscannerV2GetHarResponseLogEntry]
type accountUrlscannerV2GetHarResponseLogEntryJSON struct {
	InitialPriority apijson.Field
	InitiatorType   apijson.Field
	Priority        apijson.Field
	RequestID       apijson.Field
	RequestTime     apijson.Field
	ResourceType    apijson.Field
	Cache           apijson.Field
	Connection      apijson.Field
	Pageref         apijson.Field
	Request         apijson.Field
	Response        apijson.Field
	ServerIPAddress apijson.Field
	StartedDateTime apijson.Field
	Time            apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccountUrlscannerV2GetHarResponseLogEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerV2GetHarResponseLogEntryJSON) RawJSON() string {
	return r.raw
}

type AccountUrlscannerV2GetHarResponseLogEntriesRequest struct {
	BodySize    float64                                                    `json:"bodySize" api:"required"`
	Headers     []AccountUrlscannerV2GetHarResponseLogEntriesRequestHeader `json:"headers" api:"required"`
	HeadersSize float64                                                    `json:"headersSize" api:"required"`
	HTTPVersion string                                                     `json:"httpVersion" api:"required"`
	Method      string                                                     `json:"method" api:"required"`
	URL         string                                                     `json:"url" api:"required"`
	JSON        accountUrlscannerV2GetHarResponseLogEntriesRequestJSON     `json:"-"`
}

// accountUrlscannerV2GetHarResponseLogEntriesRequestJSON contains the JSON
// metadata for the struct [AccountUrlscannerV2GetHarResponseLogEntriesRequest]
type accountUrlscannerV2GetHarResponseLogEntriesRequestJSON struct {
	BodySize    apijson.Field
	Headers     apijson.Field
	HeadersSize apijson.Field
	HTTPVersion apijson.Field
	Method      apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUrlscannerV2GetHarResponseLogEntriesRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerV2GetHarResponseLogEntriesRequestJSON) RawJSON() string {
	return r.raw
}

type AccountUrlscannerV2GetHarResponseLogEntriesRequestHeader struct {
	Name  string                                                       `json:"name" api:"required"`
	Value string                                                       `json:"value" api:"required"`
	JSON  accountUrlscannerV2GetHarResponseLogEntriesRequestHeaderJSON `json:"-"`
}

// accountUrlscannerV2GetHarResponseLogEntriesRequestHeaderJSON contains the JSON
// metadata for the struct
// [AccountUrlscannerV2GetHarResponseLogEntriesRequestHeader]
type accountUrlscannerV2GetHarResponseLogEntriesRequestHeaderJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUrlscannerV2GetHarResponseLogEntriesRequestHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerV2GetHarResponseLogEntriesRequestHeaderJSON) RawJSON() string {
	return r.raw
}

type AccountUrlscannerV2GetHarResponseLogEntriesResponse struct {
	TransferSize float64                                                     `json:"_transferSize" api:"required"`
	BodySize     float64                                                     `json:"bodySize" api:"required"`
	Content      AccountUrlscannerV2GetHarResponseLogEntriesResponseContent  `json:"content" api:"required"`
	Headers      []AccountUrlscannerV2GetHarResponseLogEntriesResponseHeader `json:"headers" api:"required"`
	HeadersSize  float64                                                     `json:"headersSize" api:"required"`
	HTTPVersion  string                                                      `json:"httpVersion" api:"required"`
	RedirectURL  string                                                      `json:"redirectURL" api:"required"`
	Status       float64                                                     `json:"status" api:"required"`
	StatusText   string                                                      `json:"statusText" api:"required"`
	JSON         accountUrlscannerV2GetHarResponseLogEntriesResponseJSON     `json:"-"`
}

// accountUrlscannerV2GetHarResponseLogEntriesResponseJSON contains the JSON
// metadata for the struct [AccountUrlscannerV2GetHarResponseLogEntriesResponse]
type accountUrlscannerV2GetHarResponseLogEntriesResponseJSON struct {
	TransferSize apijson.Field
	BodySize     apijson.Field
	Content      apijson.Field
	Headers      apijson.Field
	HeadersSize  apijson.Field
	HTTPVersion  apijson.Field
	RedirectURL  apijson.Field
	Status       apijson.Field
	StatusText   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountUrlscannerV2GetHarResponseLogEntriesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerV2GetHarResponseLogEntriesResponseJSON) RawJSON() string {
	return r.raw
}

type AccountUrlscannerV2GetHarResponseLogEntriesResponseContent struct {
	MimeType    string                                                         `json:"mimeType" api:"required"`
	Size        float64                                                        `json:"size" api:"required"`
	Compression int64                                                          `json:"compression"`
	JSON        accountUrlscannerV2GetHarResponseLogEntriesResponseContentJSON `json:"-"`
}

// accountUrlscannerV2GetHarResponseLogEntriesResponseContentJSON contains the JSON
// metadata for the struct
// [AccountUrlscannerV2GetHarResponseLogEntriesResponseContent]
type accountUrlscannerV2GetHarResponseLogEntriesResponseContentJSON struct {
	MimeType    apijson.Field
	Size        apijson.Field
	Compression apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUrlscannerV2GetHarResponseLogEntriesResponseContent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerV2GetHarResponseLogEntriesResponseContentJSON) RawJSON() string {
	return r.raw
}

type AccountUrlscannerV2GetHarResponseLogEntriesResponseHeader struct {
	Name  string                                                        `json:"name" api:"required"`
	Value string                                                        `json:"value" api:"required"`
	JSON  accountUrlscannerV2GetHarResponseLogEntriesResponseHeaderJSON `json:"-"`
}

// accountUrlscannerV2GetHarResponseLogEntriesResponseHeaderJSON contains the JSON
// metadata for the struct
// [AccountUrlscannerV2GetHarResponseLogEntriesResponseHeader]
type accountUrlscannerV2GetHarResponseLogEntriesResponseHeaderJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUrlscannerV2GetHarResponseLogEntriesResponseHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerV2GetHarResponseLogEntriesResponseHeaderJSON) RawJSON() string {
	return r.raw
}

type AccountUrlscannerV2GetHarResponseLogPage struct {
	ID              string                                               `json:"id" api:"required"`
	PageTimings     AccountUrlscannerV2GetHarResponseLogPagesPageTimings `json:"pageTimings" api:"required"`
	StartedDateTime string                                               `json:"startedDateTime" api:"required"`
	Title           string                                               `json:"title" api:"required"`
	JSON            accountUrlscannerV2GetHarResponseLogPageJSON         `json:"-"`
}

// accountUrlscannerV2GetHarResponseLogPageJSON contains the JSON metadata for the
// struct [AccountUrlscannerV2GetHarResponseLogPage]
type accountUrlscannerV2GetHarResponseLogPageJSON struct {
	ID              apijson.Field
	PageTimings     apijson.Field
	StartedDateTime apijson.Field
	Title           apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccountUrlscannerV2GetHarResponseLogPage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerV2GetHarResponseLogPageJSON) RawJSON() string {
	return r.raw
}

type AccountUrlscannerV2GetHarResponseLogPagesPageTimings struct {
	OnContentLoad float64                                                  `json:"onContentLoad" api:"required"`
	OnLoad        float64                                                  `json:"onLoad" api:"required"`
	JSON          accountUrlscannerV2GetHarResponseLogPagesPageTimingsJSON `json:"-"`
}

// accountUrlscannerV2GetHarResponseLogPagesPageTimingsJSON contains the JSON
// metadata for the struct [AccountUrlscannerV2GetHarResponseLogPagesPageTimings]
type accountUrlscannerV2GetHarResponseLogPagesPageTimingsJSON struct {
	OnContentLoad apijson.Field
	OnLoad        apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccountUrlscannerV2GetHarResponseLogPagesPageTimings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerV2GetHarResponseLogPagesPageTimingsJSON) RawJSON() string {
	return r.raw
}

type AccountUrlscannerV2GetScanResponse struct {
	Data     AccountUrlscannerV2GetScanResponseData     `json:"data" api:"required"`
	Lists    AccountUrlscannerV2GetScanResponseLists    `json:"lists" api:"required"`
	Meta     AccountUrlscannerV2GetScanResponseMeta     `json:"meta" api:"required"`
	Page     AccountUrlscannerV2GetScanResponsePage     `json:"page" api:"required"`
	Scanner  AccountUrlscannerV2GetScanResponseScanner  `json:"scanner" api:"required"`
	Stats    AccountUrlscannerV2GetScanResponseStats    `json:"stats" api:"required"`
	Task     AccountUrlscannerV2GetScanResponseTask     `json:"task" api:"required"`
	Verdicts AccountUrlscannerV2GetScanResponseVerdicts `json:"verdicts" api:"required"`
	JSON     accountUrlscannerV2GetScanResponseJSON     `json:"-"`
}

// accountUrlscannerV2GetScanResponseJSON contains the JSON metadata for the struct
// [AccountUrlscannerV2GetScanResponse]
type accountUrlscannerV2GetScanResponseJSON struct {
	Data        apijson.Field
	Lists       apijson.Field
	Meta        apijson.Field
	Page        apijson.Field
	Scanner     apijson.Field
	Stats       apijson.Field
	Task        apijson.Field
	Verdicts    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUrlscannerV2GetScanResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerV2GetScanResponseJSON) RawJSON() string {
	return r.raw
}

type AccountUrlscannerV2GetScanResponseData struct {
	Console     []AccountUrlscannerV2GetScanResponseDataConsole     `json:"console" api:"required"`
	Cookies     []AccountUrlscannerV2GetScanResponseDataCookie      `json:"cookies" api:"required"`
	Globals     []AccountUrlscannerV2GetScanResponseDataGlobal      `json:"globals" api:"required"`
	Links       []AccountUrlscannerV2GetScanResponseDataLink        `json:"links" api:"required"`
	Performance []AccountUrlscannerV2GetScanResponseDataPerformance `json:"performance" api:"required"`
	Requests    []AccountUrlscannerV2GetScanResponseDataRequest     `json:"requests" api:"required"`
	JSON        accountUrlscannerV2GetScanResponseDataJSON          `json:"-"`
}

// accountUrlscannerV2GetScanResponseDataJSON contains the JSON metadata for the
// struct [AccountUrlscannerV2GetScanResponseData]
type accountUrlscannerV2GetScanResponseDataJSON struct {
	Console     apijson.Field
	Cookies     apijson.Field
	Globals     apijson.Field
	Links       apijson.Field
	Performance apijson.Field
	Requests    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUrlscannerV2GetScanResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerV2GetScanResponseDataJSON) RawJSON() string {
	return r.raw
}

type AccountUrlscannerV2GetScanResponseDataConsole struct {
	Message AccountUrlscannerV2GetScanResponseDataConsoleMessage `json:"message" api:"required"`
	JSON    accountUrlscannerV2GetScanResponseDataConsoleJSON    `json:"-"`
}

// accountUrlscannerV2GetScanResponseDataConsoleJSON contains the JSON metadata for
// the struct [AccountUrlscannerV2GetScanResponseDataConsole]
type accountUrlscannerV2GetScanResponseDataConsoleJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUrlscannerV2GetScanResponseDataConsole) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerV2GetScanResponseDataConsoleJSON) RawJSON() string {
	return r.raw
}

type AccountUrlscannerV2GetScanResponseDataConsoleMessage struct {
	Level  string                                                   `json:"level" api:"required"`
	Source string                                                   `json:"source" api:"required"`
	Text   string                                                   `json:"text" api:"required"`
	URL    string                                                   `json:"url" api:"required"`
	JSON   accountUrlscannerV2GetScanResponseDataConsoleMessageJSON `json:"-"`
}

// accountUrlscannerV2GetScanResponseDataConsoleMessageJSON contains the JSON
// metadata for the struct [AccountUrlscannerV2GetScanResponseDataConsoleMessage]
type accountUrlscannerV2GetScanResponseDataConsoleMessageJSON struct {
	Level       apijson.Field
	Source      apijson.Field
	Text        apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUrlscannerV2GetScanResponseDataConsoleMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerV2GetScanResponseDataConsoleMessageJSON) RawJSON() string {
	return r.raw
}

type AccountUrlscannerV2GetScanResponseDataCookie struct {
	Domain       string                                           `json:"domain" api:"required"`
	Expires      float64                                          `json:"expires" api:"required"`
	HTTPOnly     bool                                             `json:"httpOnly" api:"required"`
	Name         string                                           `json:"name" api:"required"`
	Path         string                                           `json:"path" api:"required"`
	Priority     string                                           `json:"priority" api:"required"`
	SameParty    bool                                             `json:"sameParty" api:"required"`
	Secure       bool                                             `json:"secure" api:"required"`
	Session      bool                                             `json:"session" api:"required"`
	Size         float64                                          `json:"size" api:"required"`
	SourcePort   float64                                          `json:"sourcePort" api:"required"`
	SourceScheme string                                           `json:"sourceScheme" api:"required"`
	Value        string                                           `json:"value" api:"required"`
	JSON         accountUrlscannerV2GetScanResponseDataCookieJSON `json:"-"`
}

// accountUrlscannerV2GetScanResponseDataCookieJSON contains the JSON metadata for
// the struct [AccountUrlscannerV2GetScanResponseDataCookie]
type accountUrlscannerV2GetScanResponseDataCookieJSON struct {
	Domain       apijson.Field
	Expires      apijson.Field
	HTTPOnly     apijson.Field
	Name         apijson.Field
	Path         apijson.Field
	Priority     apijson.Field
	SameParty    apijson.Field
	Secure       apijson.Field
	Session      apijson.Field
	Size         apijson.Field
	SourcePort   apijson.Field
	SourceScheme apijson.Field
	Value        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountUrlscannerV2GetScanResponseDataCookie) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerV2GetScanResponseDataCookieJSON) RawJSON() string {
	return r.raw
}

type AccountUrlscannerV2GetScanResponseDataGlobal struct {
	Prop string                                           `json:"prop" api:"required"`
	Type string                                           `json:"type" api:"required"`
	JSON accountUrlscannerV2GetScanResponseDataGlobalJSON `json:"-"`
}

// accountUrlscannerV2GetScanResponseDataGlobalJSON contains the JSON metadata for
// the struct [AccountUrlscannerV2GetScanResponseDataGlobal]
type accountUrlscannerV2GetScanResponseDataGlobalJSON struct {
	Prop        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUrlscannerV2GetScanResponseDataGlobal) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerV2GetScanResponseDataGlobalJSON) RawJSON() string {
	return r.raw
}

type AccountUrlscannerV2GetScanResponseDataLink struct {
	Href string                                         `json:"href" api:"required"`
	Text string                                         `json:"text" api:"required"`
	JSON accountUrlscannerV2GetScanResponseDataLinkJSON `json:"-"`
}

// accountUrlscannerV2GetScanResponseDataLinkJSON contains the JSON metadata for
// the struct [AccountUrlscannerV2GetScanResponseDataLink]
type accountUrlscannerV2GetScanResponseDataLinkJSON struct {
	Href        apijson.Field
	Text        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUrlscannerV2GetScanResponseDataLink) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerV2GetScanResponseDataLinkJSON) RawJSON() string {
	return r.raw
}

type AccountUrlscannerV2GetScanResponseDataPerformance struct {
	Duration  float64                                               `json:"duration" api:"required"`
	EntryType string                                                `json:"entryType" api:"required"`
	Name      string                                                `json:"name" api:"required"`
	StartTime float64                                               `json:"startTime" api:"required"`
	JSON      accountUrlscannerV2GetScanResponseDataPerformanceJSON `json:"-"`
}

// accountUrlscannerV2GetScanResponseDataPerformanceJSON contains the JSON metadata
// for the struct [AccountUrlscannerV2GetScanResponseDataPerformance]
type accountUrlscannerV2GetScanResponseDataPerformanceJSON struct {
	Duration    apijson.Field
	EntryType   apijson.Field
	Name        apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUrlscannerV2GetScanResponseDataPerformance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerV2GetScanResponseDataPerformanceJSON) RawJSON() string {
	return r.raw
}

type AccountUrlscannerV2GetScanResponseDataRequest struct {
	Request  AccountUrlscannerV2GetScanResponseDataRequestsRequest   `json:"request" api:"required"`
	Response AccountUrlscannerV2GetScanResponseDataRequestsResponse  `json:"response" api:"required"`
	Requests []AccountUrlscannerV2GetScanResponseDataRequestsRequest `json:"requests"`
	JSON     accountUrlscannerV2GetScanResponseDataRequestJSON       `json:"-"`
}

// accountUrlscannerV2GetScanResponseDataRequestJSON contains the JSON metadata for
// the struct [AccountUrlscannerV2GetScanResponseDataRequest]
type accountUrlscannerV2GetScanResponseDataRequestJSON struct {
	Request     apijson.Field
	Response    apijson.Field
	Requests    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUrlscannerV2GetScanResponseDataRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerV2GetScanResponseDataRequestJSON) RawJSON() string {
	return r.raw
}

type AccountUrlscannerV2GetScanResponseDataRequestsRequest struct {
	DocumentURL          string                                                                `json:"documentURL" api:"required"`
	HasUserGesture       bool                                                                  `json:"hasUserGesture" api:"required"`
	Initiator            AccountUrlscannerV2GetScanResponseDataRequestsRequestInitiator        `json:"initiator" api:"required"`
	RedirectHasExtraInfo bool                                                                  `json:"redirectHasExtraInfo" api:"required"`
	Request              AccountUrlscannerV2GetScanResponseDataRequestsRequestRequest          `json:"request" api:"required"`
	RequestID            string                                                                `json:"requestId" api:"required"`
	Type                 string                                                                `json:"type" api:"required"`
	WallTime             float64                                                               `json:"wallTime" api:"required"`
	FrameID              string                                                                `json:"frameId"`
	LoaderID             string                                                                `json:"loaderId"`
	PrimaryRequest       bool                                                                  `json:"primaryRequest"`
	RedirectResponse     AccountUrlscannerV2GetScanResponseDataRequestsRequestRedirectResponse `json:"redirectResponse"`
	JSON                 accountUrlscannerV2GetScanResponseDataRequestsRequestJSON             `json:"-"`
}

// accountUrlscannerV2GetScanResponseDataRequestsRequestJSON contains the JSON
// metadata for the struct [AccountUrlscannerV2GetScanResponseDataRequestsRequest]
type accountUrlscannerV2GetScanResponseDataRequestsRequestJSON struct {
	DocumentURL          apijson.Field
	HasUserGesture       apijson.Field
	Initiator            apijson.Field
	RedirectHasExtraInfo apijson.Field
	Request              apijson.Field
	RequestID            apijson.Field
	Type                 apijson.Field
	WallTime             apijson.Field
	FrameID              apijson.Field
	LoaderID             apijson.Field
	PrimaryRequest       apijson.Field
	RedirectResponse     apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccountUrlscannerV2GetScanResponseDataRequestsRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerV2GetScanResponseDataRequestsRequestJSON) RawJSON() string {
	return r.raw
}

type AccountUrlscannerV2GetScanResponseDataRequestsRequestInitiator struct {
	Host string                                                             `json:"host" api:"required"`
	Type string                                                             `json:"type" api:"required"`
	URL  string                                                             `json:"url" api:"required"`
	JSON accountUrlscannerV2GetScanResponseDataRequestsRequestInitiatorJSON `json:"-"`
}

// accountUrlscannerV2GetScanResponseDataRequestsRequestInitiatorJSON contains the
// JSON metadata for the struct
// [AccountUrlscannerV2GetScanResponseDataRequestsRequestInitiator]
type accountUrlscannerV2GetScanResponseDataRequestsRequestInitiatorJSON struct {
	Host        apijson.Field
	Type        apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUrlscannerV2GetScanResponseDataRequestsRequestInitiator) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerV2GetScanResponseDataRequestsRequestInitiatorJSON) RawJSON() string {
	return r.raw
}

type AccountUrlscannerV2GetScanResponseDataRequestsRequestRequest struct {
	InitialPriority  string                                                           `json:"initialPriority" api:"required"`
	IsSameSite       bool                                                             `json:"isSameSite" api:"required"`
	Method           string                                                           `json:"method" api:"required"`
	MixedContentType string                                                           `json:"mixedContentType" api:"required"`
	ReferrerPolicy   string                                                           `json:"referrerPolicy" api:"required"`
	URL              string                                                           `json:"url" api:"required"`
	Headers          interface{}                                                      `json:"headers"`
	JSON             accountUrlscannerV2GetScanResponseDataRequestsRequestRequestJSON `json:"-"`
}

// accountUrlscannerV2GetScanResponseDataRequestsRequestRequestJSON contains the
// JSON metadata for the struct
// [AccountUrlscannerV2GetScanResponseDataRequestsRequestRequest]
type accountUrlscannerV2GetScanResponseDataRequestsRequestRequestJSON struct {
	InitialPriority  apijson.Field
	IsSameSite       apijson.Field
	Method           apijson.Field
	MixedContentType apijson.Field
	ReferrerPolicy   apijson.Field
	URL              apijson.Field
	Headers          apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountUrlscannerV2GetScanResponseDataRequestsRequestRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerV2GetScanResponseDataRequestsRequestRequestJSON) RawJSON() string {
	return r.raw
}

type AccountUrlscannerV2GetScanResponseDataRequestsRequestRedirectResponse struct {
	Charset         string                                                                                `json:"charset" api:"required"`
	MimeType        string                                                                                `json:"mimeType" api:"required"`
	Protocol        string                                                                                `json:"protocol" api:"required"`
	RemoteIPAddress string                                                                                `json:"remoteIPAddress" api:"required"`
	RemotePort      float64                                                                               `json:"remotePort" api:"required"`
	SecurityHeaders []AccountUrlscannerV2GetScanResponseDataRequestsRequestRedirectResponseSecurityHeader `json:"securityHeaders" api:"required"`
	SecurityState   string                                                                                `json:"securityState" api:"required"`
	Status          float64                                                                               `json:"status" api:"required"`
	StatusText      string                                                                                `json:"statusText" api:"required"`
	URL             string                                                                                `json:"url" api:"required"`
	Headers         interface{}                                                                           `json:"headers"`
	JSON            accountUrlscannerV2GetScanResponseDataRequestsRequestRedirectResponseJSON             `json:"-"`
}

// accountUrlscannerV2GetScanResponseDataRequestsRequestRedirectResponseJSON
// contains the JSON metadata for the struct
// [AccountUrlscannerV2GetScanResponseDataRequestsRequestRedirectResponse]
type accountUrlscannerV2GetScanResponseDataRequestsRequestRedirectResponseJSON struct {
	Charset         apijson.Field
	MimeType        apijson.Field
	Protocol        apijson.Field
	RemoteIPAddress apijson.Field
	RemotePort      apijson.Field
	SecurityHeaders apijson.Field
	SecurityState   apijson.Field
	Status          apijson.Field
	StatusText      apijson.Field
	URL             apijson.Field
	Headers         apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccountUrlscannerV2GetScanResponseDataRequestsRequestRedirectResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerV2GetScanResponseDataRequestsRequestRedirectResponseJSON) RawJSON() string {
	return r.raw
}

type AccountUrlscannerV2GetScanResponseDataRequestsRequestRedirectResponseSecurityHeader struct {
	Name  string                                                                                  `json:"name" api:"required"`
	Value string                                                                                  `json:"value" api:"required"`
	JSON  accountUrlscannerV2GetScanResponseDataRequestsRequestRedirectResponseSecurityHeaderJSON `json:"-"`
}

// accountUrlscannerV2GetScanResponseDataRequestsRequestRedirectResponseSecurityHeaderJSON
// contains the JSON metadata for the struct
// [AccountUrlscannerV2GetScanResponseDataRequestsRequestRedirectResponseSecurityHeader]
type accountUrlscannerV2GetScanResponseDataRequestsRequestRedirectResponseSecurityHeaderJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUrlscannerV2GetScanResponseDataRequestsRequestRedirectResponseSecurityHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerV2GetScanResponseDataRequestsRequestRedirectResponseSecurityHeaderJSON) RawJSON() string {
	return r.raw
}

type AccountUrlscannerV2GetScanResponseDataRequestsResponse struct {
	Asn               AccountUrlscannerV2GetScanResponseDataRequestsResponseAsn      `json:"asn" api:"required"`
	DataLength        float64                                                        `json:"dataLength" api:"required"`
	EncodedDataLength float64                                                        `json:"encodedDataLength" api:"required"`
	Geoip             AccountUrlscannerV2GetScanResponseDataRequestsResponseGeoip    `json:"geoip" api:"required"`
	HasExtraInfo      bool                                                           `json:"hasExtraInfo" api:"required"`
	RequestID         string                                                         `json:"requestId" api:"required"`
	Response          AccountUrlscannerV2GetScanResponseDataRequestsResponseResponse `json:"response" api:"required"`
	Size              float64                                                        `json:"size" api:"required"`
	Type              string                                                         `json:"type" api:"required"`
	ContentAvailable  bool                                                           `json:"contentAvailable"`
	Hash              string                                                         `json:"hash"`
	JSON              accountUrlscannerV2GetScanResponseDataRequestsResponseJSON     `json:"-"`
}

// accountUrlscannerV2GetScanResponseDataRequestsResponseJSON contains the JSON
// metadata for the struct [AccountUrlscannerV2GetScanResponseDataRequestsResponse]
type accountUrlscannerV2GetScanResponseDataRequestsResponseJSON struct {
	Asn               apijson.Field
	DataLength        apijson.Field
	EncodedDataLength apijson.Field
	Geoip             apijson.Field
	HasExtraInfo      apijson.Field
	RequestID         apijson.Field
	Response          apijson.Field
	Size              apijson.Field
	Type              apijson.Field
	ContentAvailable  apijson.Field
	Hash              apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *AccountUrlscannerV2GetScanResponseDataRequestsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerV2GetScanResponseDataRequestsResponseJSON) RawJSON() string {
	return r.raw
}

type AccountUrlscannerV2GetScanResponseDataRequestsResponseAsn struct {
	Asn         string                                                        `json:"asn" api:"required"`
	Country     string                                                        `json:"country" api:"required"`
	Description string                                                        `json:"description" api:"required"`
	IP          string                                                        `json:"ip" api:"required"`
	Name        string                                                        `json:"name" api:"required"`
	Org         string                                                        `json:"org" api:"required"`
	JSON        accountUrlscannerV2GetScanResponseDataRequestsResponseAsnJSON `json:"-"`
}

// accountUrlscannerV2GetScanResponseDataRequestsResponseAsnJSON contains the JSON
// metadata for the struct
// [AccountUrlscannerV2GetScanResponseDataRequestsResponseAsn]
type accountUrlscannerV2GetScanResponseDataRequestsResponseAsnJSON struct {
	Asn         apijson.Field
	Country     apijson.Field
	Description apijson.Field
	IP          apijson.Field
	Name        apijson.Field
	Org         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUrlscannerV2GetScanResponseDataRequestsResponseAsn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerV2GetScanResponseDataRequestsResponseAsnJSON) RawJSON() string {
	return r.raw
}

type AccountUrlscannerV2GetScanResponseDataRequestsResponseGeoip struct {
	City        string                                                          `json:"city" api:"required"`
	Country     string                                                          `json:"country" api:"required"`
	CountryName string                                                          `json:"country_name" api:"required"`
	GeonameID   string                                                          `json:"geonameId" api:"required"`
	Ll          []float64                                                       `json:"ll" api:"required"`
	Region      string                                                          `json:"region" api:"required"`
	JSON        accountUrlscannerV2GetScanResponseDataRequestsResponseGeoipJSON `json:"-"`
}

// accountUrlscannerV2GetScanResponseDataRequestsResponseGeoipJSON contains the
// JSON metadata for the struct
// [AccountUrlscannerV2GetScanResponseDataRequestsResponseGeoip]
type accountUrlscannerV2GetScanResponseDataRequestsResponseGeoipJSON struct {
	City        apijson.Field
	Country     apijson.Field
	CountryName apijson.Field
	GeonameID   apijson.Field
	Ll          apijson.Field
	Region      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUrlscannerV2GetScanResponseDataRequestsResponseGeoip) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerV2GetScanResponseDataRequestsResponseGeoipJSON) RawJSON() string {
	return r.raw
}

type AccountUrlscannerV2GetScanResponseDataRequestsResponseResponse struct {
	Charset         string                                                                         `json:"charset" api:"required"`
	MimeType        string                                                                         `json:"mimeType" api:"required"`
	Protocol        string                                                                         `json:"protocol" api:"required"`
	RemoteIPAddress string                                                                         `json:"remoteIPAddress" api:"required"`
	RemotePort      float64                                                                        `json:"remotePort" api:"required"`
	SecurityDetails AccountUrlscannerV2GetScanResponseDataRequestsResponseResponseSecurityDetails  `json:"securityDetails" api:"required"`
	SecurityHeaders []AccountUrlscannerV2GetScanResponseDataRequestsResponseResponseSecurityHeader `json:"securityHeaders" api:"required"`
	SecurityState   string                                                                         `json:"securityState" api:"required"`
	Status          float64                                                                        `json:"status" api:"required"`
	StatusText      string                                                                         `json:"statusText" api:"required"`
	URL             string                                                                         `json:"url" api:"required"`
	Headers         interface{}                                                                    `json:"headers"`
	JSON            accountUrlscannerV2GetScanResponseDataRequestsResponseResponseJSON             `json:"-"`
}

// accountUrlscannerV2GetScanResponseDataRequestsResponseResponseJSON contains the
// JSON metadata for the struct
// [AccountUrlscannerV2GetScanResponseDataRequestsResponseResponse]
type accountUrlscannerV2GetScanResponseDataRequestsResponseResponseJSON struct {
	Charset         apijson.Field
	MimeType        apijson.Field
	Protocol        apijson.Field
	RemoteIPAddress apijson.Field
	RemotePort      apijson.Field
	SecurityDetails apijson.Field
	SecurityHeaders apijson.Field
	SecurityState   apijson.Field
	Status          apijson.Field
	StatusText      apijson.Field
	URL             apijson.Field
	Headers         apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccountUrlscannerV2GetScanResponseDataRequestsResponseResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerV2GetScanResponseDataRequestsResponseResponseJSON) RawJSON() string {
	return r.raw
}

type AccountUrlscannerV2GetScanResponseDataRequestsResponseResponseSecurityDetails struct {
	CertificateID                     float64                                                                           `json:"certificateId" api:"required"`
	CertificateTransparencyCompliance string                                                                            `json:"certificateTransparencyCompliance" api:"required"`
	Cipher                            string                                                                            `json:"cipher" api:"required"`
	EncryptedClientHello              bool                                                                              `json:"encryptedClientHello" api:"required"`
	Issuer                            string                                                                            `json:"issuer" api:"required"`
	KeyExchange                       string                                                                            `json:"keyExchange" api:"required"`
	KeyExchangeGroup                  string                                                                            `json:"keyExchangeGroup" api:"required"`
	Protocol                          string                                                                            `json:"protocol" api:"required"`
	SanList                           []string                                                                          `json:"sanList" api:"required"`
	ServerSignatureAlgorithm          float64                                                                           `json:"serverSignatureAlgorithm" api:"required"`
	SubjectName                       string                                                                            `json:"subjectName" api:"required"`
	ValidFrom                         float64                                                                           `json:"validFrom" api:"required"`
	ValidTo                           float64                                                                           `json:"validTo" api:"required"`
	JSON                              accountUrlscannerV2GetScanResponseDataRequestsResponseResponseSecurityDetailsJSON `json:"-"`
}

// accountUrlscannerV2GetScanResponseDataRequestsResponseResponseSecurityDetailsJSON
// contains the JSON metadata for the struct
// [AccountUrlscannerV2GetScanResponseDataRequestsResponseResponseSecurityDetails]
type accountUrlscannerV2GetScanResponseDataRequestsResponseResponseSecurityDetailsJSON struct {
	CertificateID                     apijson.Field
	CertificateTransparencyCompliance apijson.Field
	Cipher                            apijson.Field
	EncryptedClientHello              apijson.Field
	Issuer                            apijson.Field
	KeyExchange                       apijson.Field
	KeyExchangeGroup                  apijson.Field
	Protocol                          apijson.Field
	SanList                           apijson.Field
	ServerSignatureAlgorithm          apijson.Field
	SubjectName                       apijson.Field
	ValidFrom                         apijson.Field
	ValidTo                           apijson.Field
	raw                               string
	ExtraFields                       map[string]apijson.Field
}

func (r *AccountUrlscannerV2GetScanResponseDataRequestsResponseResponseSecurityDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerV2GetScanResponseDataRequestsResponseResponseSecurityDetailsJSON) RawJSON() string {
	return r.raw
}

type AccountUrlscannerV2GetScanResponseDataRequestsResponseResponseSecurityHeader struct {
	Name  string                                                                           `json:"name" api:"required"`
	Value string                                                                           `json:"value" api:"required"`
	JSON  accountUrlscannerV2GetScanResponseDataRequestsResponseResponseSecurityHeaderJSON `json:"-"`
}

// accountUrlscannerV2GetScanResponseDataRequestsResponseResponseSecurityHeaderJSON
// contains the JSON metadata for the struct
// [AccountUrlscannerV2GetScanResponseDataRequestsResponseResponseSecurityHeader]
type accountUrlscannerV2GetScanResponseDataRequestsResponseResponseSecurityHeaderJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUrlscannerV2GetScanResponseDataRequestsResponseResponseSecurityHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerV2GetScanResponseDataRequestsResponseResponseSecurityHeaderJSON) RawJSON() string {
	return r.raw
}

type AccountUrlscannerV2GetScanResponseLists struct {
	Asns         []string                                             `json:"asns" api:"required"`
	Certificates []AccountUrlscannerV2GetScanResponseListsCertificate `json:"certificates" api:"required"`
	Continents   []string                                             `json:"continents" api:"required"`
	Countries    []string                                             `json:"countries" api:"required"`
	Domains      []string                                             `json:"domains" api:"required"`
	Hashes       []string                                             `json:"hashes" api:"required"`
	IPs          []string                                             `json:"ips" api:"required"`
	LinkDomains  []string                                             `json:"linkDomains" api:"required"`
	Servers      []string                                             `json:"servers" api:"required"`
	URLs         []string                                             `json:"urls" api:"required"`
	JSON         accountUrlscannerV2GetScanResponseListsJSON          `json:"-"`
}

// accountUrlscannerV2GetScanResponseListsJSON contains the JSON metadata for the
// struct [AccountUrlscannerV2GetScanResponseLists]
type accountUrlscannerV2GetScanResponseListsJSON struct {
	Asns         apijson.Field
	Certificates apijson.Field
	Continents   apijson.Field
	Countries    apijson.Field
	Domains      apijson.Field
	Hashes       apijson.Field
	IPs          apijson.Field
	LinkDomains  apijson.Field
	Servers      apijson.Field
	URLs         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountUrlscannerV2GetScanResponseLists) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerV2GetScanResponseListsJSON) RawJSON() string {
	return r.raw
}

type AccountUrlscannerV2GetScanResponseListsCertificate struct {
	Issuer      string                                                 `json:"issuer" api:"required"`
	SubjectName string                                                 `json:"subjectName" api:"required"`
	ValidFrom   float64                                                `json:"validFrom" api:"required"`
	ValidTo     float64                                                `json:"validTo" api:"required"`
	JSON        accountUrlscannerV2GetScanResponseListsCertificateJSON `json:"-"`
}

// accountUrlscannerV2GetScanResponseListsCertificateJSON contains the JSON
// metadata for the struct [AccountUrlscannerV2GetScanResponseListsCertificate]
type accountUrlscannerV2GetScanResponseListsCertificateJSON struct {
	Issuer      apijson.Field
	SubjectName apijson.Field
	ValidFrom   apijson.Field
	ValidTo     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUrlscannerV2GetScanResponseListsCertificate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerV2GetScanResponseListsCertificateJSON) RawJSON() string {
	return r.raw
}

type AccountUrlscannerV2GetScanResponseMeta struct {
	Processors AccountUrlscannerV2GetScanResponseMetaProcessors `json:"processors" api:"required"`
	JSON       accountUrlscannerV2GetScanResponseMetaJSON       `json:"-"`
}

// accountUrlscannerV2GetScanResponseMetaJSON contains the JSON metadata for the
// struct [AccountUrlscannerV2GetScanResponseMeta]
type accountUrlscannerV2GetScanResponseMetaJSON struct {
	Processors  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUrlscannerV2GetScanResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerV2GetScanResponseMetaJSON) RawJSON() string {
	return r.raw
}

type AccountUrlscannerV2GetScanResponseMetaProcessors struct {
	Asn              AccountUrlscannerV2GetScanResponseMetaProcessorsAsn              `json:"asn" api:"required"`
	DNS              AccountUrlscannerV2GetScanResponseMetaProcessorsDNS              `json:"dns" api:"required"`
	DomainCategories AccountUrlscannerV2GetScanResponseMetaProcessorsDomainCategories `json:"domainCategories" api:"required"`
	Geoip            AccountUrlscannerV2GetScanResponseMetaProcessorsGeoip            `json:"geoip" api:"required"`
	Phishing         AccountUrlscannerV2GetScanResponseMetaProcessorsPhishing         `json:"phishing" api:"required"`
	RadarRank        AccountUrlscannerV2GetScanResponseMetaProcessorsRadarRank        `json:"radarRank" api:"required"`
	Wappa            AccountUrlscannerV2GetScanResponseMetaProcessorsWappa            `json:"wappa" api:"required"`
	URLCategories    AccountUrlscannerV2GetScanResponseMetaProcessorsURLCategories    `json:"urlCategories"`
	JSON             accountUrlscannerV2GetScanResponseMetaProcessorsJSON             `json:"-"`
}

// accountUrlscannerV2GetScanResponseMetaProcessorsJSON contains the JSON metadata
// for the struct [AccountUrlscannerV2GetScanResponseMetaProcessors]
type accountUrlscannerV2GetScanResponseMetaProcessorsJSON struct {
	Asn              apijson.Field
	DNS              apijson.Field
	DomainCategories apijson.Field
	Geoip            apijson.Field
	Phishing         apijson.Field
	RadarRank        apijson.Field
	Wappa            apijson.Field
	URLCategories    apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountUrlscannerV2GetScanResponseMetaProcessors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerV2GetScanResponseMetaProcessorsJSON) RawJSON() string {
	return r.raw
}

type AccountUrlscannerV2GetScanResponseMetaProcessorsAsn struct {
	Data []AccountUrlscannerV2GetScanResponseMetaProcessorsAsnData `json:"data" api:"required"`
	JSON accountUrlscannerV2GetScanResponseMetaProcessorsAsnJSON   `json:"-"`
}

// accountUrlscannerV2GetScanResponseMetaProcessorsAsnJSON contains the JSON
// metadata for the struct [AccountUrlscannerV2GetScanResponseMetaProcessorsAsn]
type accountUrlscannerV2GetScanResponseMetaProcessorsAsnJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUrlscannerV2GetScanResponseMetaProcessorsAsn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerV2GetScanResponseMetaProcessorsAsnJSON) RawJSON() string {
	return r.raw
}

type AccountUrlscannerV2GetScanResponseMetaProcessorsAsnData struct {
	Asn         string                                                      `json:"asn" api:"required"`
	Country     string                                                      `json:"country" api:"required"`
	Description string                                                      `json:"description" api:"required"`
	IP          string                                                      `json:"ip" api:"required"`
	Name        string                                                      `json:"name" api:"required"`
	JSON        accountUrlscannerV2GetScanResponseMetaProcessorsAsnDataJSON `json:"-"`
}

// accountUrlscannerV2GetScanResponseMetaProcessorsAsnDataJSON contains the JSON
// metadata for the struct
// [AccountUrlscannerV2GetScanResponseMetaProcessorsAsnData]
type accountUrlscannerV2GetScanResponseMetaProcessorsAsnDataJSON struct {
	Asn         apijson.Field
	Country     apijson.Field
	Description apijson.Field
	IP          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUrlscannerV2GetScanResponseMetaProcessorsAsnData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerV2GetScanResponseMetaProcessorsAsnDataJSON) RawJSON() string {
	return r.raw
}

type AccountUrlscannerV2GetScanResponseMetaProcessorsDNS struct {
	Data []AccountUrlscannerV2GetScanResponseMetaProcessorsDNSData `json:"data" api:"required"`
	JSON accountUrlscannerV2GetScanResponseMetaProcessorsDNSJSON   `json:"-"`
}

// accountUrlscannerV2GetScanResponseMetaProcessorsDNSJSON contains the JSON
// metadata for the struct [AccountUrlscannerV2GetScanResponseMetaProcessorsDNS]
type accountUrlscannerV2GetScanResponseMetaProcessorsDNSJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUrlscannerV2GetScanResponseMetaProcessorsDNS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerV2GetScanResponseMetaProcessorsDNSJSON) RawJSON() string {
	return r.raw
}

type AccountUrlscannerV2GetScanResponseMetaProcessorsDNSData struct {
	Address     string                                                      `json:"address" api:"required"`
	DnssecValid bool                                                        `json:"dnssec_valid" api:"required"`
	Name        string                                                      `json:"name" api:"required"`
	Type        string                                                      `json:"type" api:"required"`
	JSON        accountUrlscannerV2GetScanResponseMetaProcessorsDNSDataJSON `json:"-"`
}

// accountUrlscannerV2GetScanResponseMetaProcessorsDNSDataJSON contains the JSON
// metadata for the struct
// [AccountUrlscannerV2GetScanResponseMetaProcessorsDNSData]
type accountUrlscannerV2GetScanResponseMetaProcessorsDNSDataJSON struct {
	Address     apijson.Field
	DnssecValid apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUrlscannerV2GetScanResponseMetaProcessorsDNSData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerV2GetScanResponseMetaProcessorsDNSDataJSON) RawJSON() string {
	return r.raw
}

type AccountUrlscannerV2GetScanResponseMetaProcessorsDomainCategories struct {
	Data []AccountUrlscannerV2GetScanResponseMetaProcessorsDomainCategoriesData `json:"data" api:"required"`
	JSON accountUrlscannerV2GetScanResponseMetaProcessorsDomainCategoriesJSON   `json:"-"`
}

// accountUrlscannerV2GetScanResponseMetaProcessorsDomainCategoriesJSON contains
// the JSON metadata for the struct
// [AccountUrlscannerV2GetScanResponseMetaProcessorsDomainCategories]
type accountUrlscannerV2GetScanResponseMetaProcessorsDomainCategoriesJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUrlscannerV2GetScanResponseMetaProcessorsDomainCategories) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerV2GetScanResponseMetaProcessorsDomainCategoriesJSON) RawJSON() string {
	return r.raw
}

type AccountUrlscannerV2GetScanResponseMetaProcessorsDomainCategoriesData struct {
	Inherited interface{}                                                              `json:"inherited" api:"required"`
	IsPrimary bool                                                                     `json:"isPrimary" api:"required"`
	Name      string                                                                   `json:"name" api:"required"`
	JSON      accountUrlscannerV2GetScanResponseMetaProcessorsDomainCategoriesDataJSON `json:"-"`
}

// accountUrlscannerV2GetScanResponseMetaProcessorsDomainCategoriesDataJSON
// contains the JSON metadata for the struct
// [AccountUrlscannerV2GetScanResponseMetaProcessorsDomainCategoriesData]
type accountUrlscannerV2GetScanResponseMetaProcessorsDomainCategoriesDataJSON struct {
	Inherited   apijson.Field
	IsPrimary   apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUrlscannerV2GetScanResponseMetaProcessorsDomainCategoriesData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerV2GetScanResponseMetaProcessorsDomainCategoriesDataJSON) RawJSON() string {
	return r.raw
}

type AccountUrlscannerV2GetScanResponseMetaProcessorsGeoip struct {
	Data []AccountUrlscannerV2GetScanResponseMetaProcessorsGeoipData `json:"data" api:"required"`
	JSON accountUrlscannerV2GetScanResponseMetaProcessorsGeoipJSON   `json:"-"`
}

// accountUrlscannerV2GetScanResponseMetaProcessorsGeoipJSON contains the JSON
// metadata for the struct [AccountUrlscannerV2GetScanResponseMetaProcessorsGeoip]
type accountUrlscannerV2GetScanResponseMetaProcessorsGeoipJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUrlscannerV2GetScanResponseMetaProcessorsGeoip) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerV2GetScanResponseMetaProcessorsGeoipJSON) RawJSON() string {
	return r.raw
}

type AccountUrlscannerV2GetScanResponseMetaProcessorsGeoipData struct {
	Geoip AccountUrlscannerV2GetScanResponseMetaProcessorsGeoipDataGeoip `json:"geoip" api:"required"`
	IP    string                                                         `json:"ip" api:"required"`
	JSON  accountUrlscannerV2GetScanResponseMetaProcessorsGeoipDataJSON  `json:"-"`
}

// accountUrlscannerV2GetScanResponseMetaProcessorsGeoipDataJSON contains the JSON
// metadata for the struct
// [AccountUrlscannerV2GetScanResponseMetaProcessorsGeoipData]
type accountUrlscannerV2GetScanResponseMetaProcessorsGeoipDataJSON struct {
	Geoip       apijson.Field
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUrlscannerV2GetScanResponseMetaProcessorsGeoipData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerV2GetScanResponseMetaProcessorsGeoipDataJSON) RawJSON() string {
	return r.raw
}

type AccountUrlscannerV2GetScanResponseMetaProcessorsGeoipDataGeoip struct {
	City        string                                                             `json:"city" api:"required"`
	Country     string                                                             `json:"country" api:"required"`
	CountryName string                                                             `json:"country_name" api:"required"`
	Ll          []float64                                                          `json:"ll" api:"required"`
	Region      string                                                             `json:"region" api:"required"`
	JSON        accountUrlscannerV2GetScanResponseMetaProcessorsGeoipDataGeoipJSON `json:"-"`
}

// accountUrlscannerV2GetScanResponseMetaProcessorsGeoipDataGeoipJSON contains the
// JSON metadata for the struct
// [AccountUrlscannerV2GetScanResponseMetaProcessorsGeoipDataGeoip]
type accountUrlscannerV2GetScanResponseMetaProcessorsGeoipDataGeoipJSON struct {
	City        apijson.Field
	Country     apijson.Field
	CountryName apijson.Field
	Ll          apijson.Field
	Region      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUrlscannerV2GetScanResponseMetaProcessorsGeoipDataGeoip) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerV2GetScanResponseMetaProcessorsGeoipDataGeoipJSON) RawJSON() string {
	return r.raw
}

type AccountUrlscannerV2GetScanResponseMetaProcessorsPhishing struct {
	Data []string                                                     `json:"data" api:"required"`
	JSON accountUrlscannerV2GetScanResponseMetaProcessorsPhishingJSON `json:"-"`
}

// accountUrlscannerV2GetScanResponseMetaProcessorsPhishingJSON contains the JSON
// metadata for the struct
// [AccountUrlscannerV2GetScanResponseMetaProcessorsPhishing]
type accountUrlscannerV2GetScanResponseMetaProcessorsPhishingJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUrlscannerV2GetScanResponseMetaProcessorsPhishing) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerV2GetScanResponseMetaProcessorsPhishingJSON) RawJSON() string {
	return r.raw
}

type AccountUrlscannerV2GetScanResponseMetaProcessorsRadarRank struct {
	Data []AccountUrlscannerV2GetScanResponseMetaProcessorsRadarRankData `json:"data" api:"required"`
	JSON accountUrlscannerV2GetScanResponseMetaProcessorsRadarRankJSON   `json:"-"`
}

// accountUrlscannerV2GetScanResponseMetaProcessorsRadarRankJSON contains the JSON
// metadata for the struct
// [AccountUrlscannerV2GetScanResponseMetaProcessorsRadarRank]
type accountUrlscannerV2GetScanResponseMetaProcessorsRadarRankJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUrlscannerV2GetScanResponseMetaProcessorsRadarRank) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerV2GetScanResponseMetaProcessorsRadarRankJSON) RawJSON() string {
	return r.raw
}

type AccountUrlscannerV2GetScanResponseMetaProcessorsRadarRankData struct {
	Bucket   string                                                            `json:"bucket" api:"required"`
	Hostname string                                                            `json:"hostname" api:"required"`
	Rank     float64                                                           `json:"rank"`
	JSON     accountUrlscannerV2GetScanResponseMetaProcessorsRadarRankDataJSON `json:"-"`
}

// accountUrlscannerV2GetScanResponseMetaProcessorsRadarRankDataJSON contains the
// JSON metadata for the struct
// [AccountUrlscannerV2GetScanResponseMetaProcessorsRadarRankData]
type accountUrlscannerV2GetScanResponseMetaProcessorsRadarRankDataJSON struct {
	Bucket      apijson.Field
	Hostname    apijson.Field
	Rank        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUrlscannerV2GetScanResponseMetaProcessorsRadarRankData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerV2GetScanResponseMetaProcessorsRadarRankDataJSON) RawJSON() string {
	return r.raw
}

type AccountUrlscannerV2GetScanResponseMetaProcessorsWappa struct {
	Data []AccountUrlscannerV2GetScanResponseMetaProcessorsWappaData `json:"data" api:"required"`
	JSON accountUrlscannerV2GetScanResponseMetaProcessorsWappaJSON   `json:"-"`
}

// accountUrlscannerV2GetScanResponseMetaProcessorsWappaJSON contains the JSON
// metadata for the struct [AccountUrlscannerV2GetScanResponseMetaProcessorsWappa]
type accountUrlscannerV2GetScanResponseMetaProcessorsWappaJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUrlscannerV2GetScanResponseMetaProcessorsWappa) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerV2GetScanResponseMetaProcessorsWappaJSON) RawJSON() string {
	return r.raw
}

type AccountUrlscannerV2GetScanResponseMetaProcessorsWappaData struct {
	App             string                                                                `json:"app" api:"required"`
	Categories      []AccountUrlscannerV2GetScanResponseMetaProcessorsWappaDataCategory   `json:"categories" api:"required"`
	Confidence      []AccountUrlscannerV2GetScanResponseMetaProcessorsWappaDataConfidence `json:"confidence" api:"required"`
	ConfidenceTotal float64                                                               `json:"confidenceTotal" api:"required"`
	Icon            string                                                                `json:"icon" api:"required"`
	Website         string                                                                `json:"website" api:"required"`
	JSON            accountUrlscannerV2GetScanResponseMetaProcessorsWappaDataJSON         `json:"-"`
}

// accountUrlscannerV2GetScanResponseMetaProcessorsWappaDataJSON contains the JSON
// metadata for the struct
// [AccountUrlscannerV2GetScanResponseMetaProcessorsWappaData]
type accountUrlscannerV2GetScanResponseMetaProcessorsWappaDataJSON struct {
	App             apijson.Field
	Categories      apijson.Field
	Confidence      apijson.Field
	ConfidenceTotal apijson.Field
	Icon            apijson.Field
	Website         apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccountUrlscannerV2GetScanResponseMetaProcessorsWappaData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerV2GetScanResponseMetaProcessorsWappaDataJSON) RawJSON() string {
	return r.raw
}

type AccountUrlscannerV2GetScanResponseMetaProcessorsWappaDataCategory struct {
	Name     string                                                                `json:"name" api:"required"`
	Priority float64                                                               `json:"priority" api:"required"`
	JSON     accountUrlscannerV2GetScanResponseMetaProcessorsWappaDataCategoryJSON `json:"-"`
}

// accountUrlscannerV2GetScanResponseMetaProcessorsWappaDataCategoryJSON contains
// the JSON metadata for the struct
// [AccountUrlscannerV2GetScanResponseMetaProcessorsWappaDataCategory]
type accountUrlscannerV2GetScanResponseMetaProcessorsWappaDataCategoryJSON struct {
	Name        apijson.Field
	Priority    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUrlscannerV2GetScanResponseMetaProcessorsWappaDataCategory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerV2GetScanResponseMetaProcessorsWappaDataCategoryJSON) RawJSON() string {
	return r.raw
}

type AccountUrlscannerV2GetScanResponseMetaProcessorsWappaDataConfidence struct {
	Confidence  float64                                                                 `json:"confidence" api:"required"`
	Name        string                                                                  `json:"name" api:"required"`
	Pattern     string                                                                  `json:"pattern" api:"required"`
	PatternType string                                                                  `json:"patternType" api:"required"`
	JSON        accountUrlscannerV2GetScanResponseMetaProcessorsWappaDataConfidenceJSON `json:"-"`
}

// accountUrlscannerV2GetScanResponseMetaProcessorsWappaDataConfidenceJSON contains
// the JSON metadata for the struct
// [AccountUrlscannerV2GetScanResponseMetaProcessorsWappaDataConfidence]
type accountUrlscannerV2GetScanResponseMetaProcessorsWappaDataConfidenceJSON struct {
	Confidence  apijson.Field
	Name        apijson.Field
	Pattern     apijson.Field
	PatternType apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUrlscannerV2GetScanResponseMetaProcessorsWappaDataConfidence) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerV2GetScanResponseMetaProcessorsWappaDataConfidenceJSON) RawJSON() string {
	return r.raw
}

type AccountUrlscannerV2GetScanResponseMetaProcessorsURLCategories struct {
	Data []AccountUrlscannerV2GetScanResponseMetaProcessorsURLCategoriesData `json:"data" api:"required"`
	JSON accountUrlscannerV2GetScanResponseMetaProcessorsURLCategoriesJSON   `json:"-"`
}

// accountUrlscannerV2GetScanResponseMetaProcessorsURLCategoriesJSON contains the
// JSON metadata for the struct
// [AccountUrlscannerV2GetScanResponseMetaProcessorsURLCategories]
type accountUrlscannerV2GetScanResponseMetaProcessorsURLCategoriesJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUrlscannerV2GetScanResponseMetaProcessorsURLCategories) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerV2GetScanResponseMetaProcessorsURLCategoriesJSON) RawJSON() string {
	return r.raw
}

type AccountUrlscannerV2GetScanResponseMetaProcessorsURLCategoriesData struct {
	Content   []AccountUrlscannerV2GetScanResponseMetaProcessorsURLCategoriesDataContent `json:"content" api:"required"`
	Inherited AccountUrlscannerV2GetScanResponseMetaProcessorsURLCategoriesDataInherited `json:"inherited" api:"required"`
	Name      string                                                                     `json:"name" api:"required"`
	Risks     []AccountUrlscannerV2GetScanResponseMetaProcessorsURLCategoriesDataRisk    `json:"risks" api:"required"`
	JSON      accountUrlscannerV2GetScanResponseMetaProcessorsURLCategoriesDataJSON      `json:"-"`
}

// accountUrlscannerV2GetScanResponseMetaProcessorsURLCategoriesDataJSON contains
// the JSON metadata for the struct
// [AccountUrlscannerV2GetScanResponseMetaProcessorsURLCategoriesData]
type accountUrlscannerV2GetScanResponseMetaProcessorsURLCategoriesDataJSON struct {
	Content     apijson.Field
	Inherited   apijson.Field
	Name        apijson.Field
	Risks       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUrlscannerV2GetScanResponseMetaProcessorsURLCategoriesData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerV2GetScanResponseMetaProcessorsURLCategoriesDataJSON) RawJSON() string {
	return r.raw
}

type AccountUrlscannerV2GetScanResponseMetaProcessorsURLCategoriesDataContent struct {
	ID              float64                                                                      `json:"id" api:"required"`
	Name            string                                                                       `json:"name" api:"required"`
	SuperCategoryID float64                                                                      `json:"super_category_id" api:"required"`
	JSON            accountUrlscannerV2GetScanResponseMetaProcessorsURLCategoriesDataContentJSON `json:"-"`
}

// accountUrlscannerV2GetScanResponseMetaProcessorsURLCategoriesDataContentJSON
// contains the JSON metadata for the struct
// [AccountUrlscannerV2GetScanResponseMetaProcessorsURLCategoriesDataContent]
type accountUrlscannerV2GetScanResponseMetaProcessorsURLCategoriesDataContentJSON struct {
	ID              apijson.Field
	Name            apijson.Field
	SuperCategoryID apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccountUrlscannerV2GetScanResponseMetaProcessorsURLCategoriesDataContent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerV2GetScanResponseMetaProcessorsURLCategoriesDataContentJSON) RawJSON() string {
	return r.raw
}

type AccountUrlscannerV2GetScanResponseMetaProcessorsURLCategoriesDataInherited struct {
	Content []AccountUrlscannerV2GetScanResponseMetaProcessorsURLCategoriesDataInheritedContent `json:"content" api:"required"`
	From    string                                                                              `json:"from" api:"required"`
	Risks   []AccountUrlscannerV2GetScanResponseMetaProcessorsURLCategoriesDataInheritedRisk    `json:"risks" api:"required"`
	JSON    accountUrlscannerV2GetScanResponseMetaProcessorsURLCategoriesDataInheritedJSON      `json:"-"`
}

// accountUrlscannerV2GetScanResponseMetaProcessorsURLCategoriesDataInheritedJSON
// contains the JSON metadata for the struct
// [AccountUrlscannerV2GetScanResponseMetaProcessorsURLCategoriesDataInherited]
type accountUrlscannerV2GetScanResponseMetaProcessorsURLCategoriesDataInheritedJSON struct {
	Content     apijson.Field
	From        apijson.Field
	Risks       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUrlscannerV2GetScanResponseMetaProcessorsURLCategoriesDataInherited) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerV2GetScanResponseMetaProcessorsURLCategoriesDataInheritedJSON) RawJSON() string {
	return r.raw
}

type AccountUrlscannerV2GetScanResponseMetaProcessorsURLCategoriesDataInheritedContent struct {
	ID              float64                                                                               `json:"id" api:"required"`
	Name            string                                                                                `json:"name" api:"required"`
	SuperCategoryID float64                                                                               `json:"super_category_id" api:"required"`
	JSON            accountUrlscannerV2GetScanResponseMetaProcessorsURLCategoriesDataInheritedContentJSON `json:"-"`
}

// accountUrlscannerV2GetScanResponseMetaProcessorsURLCategoriesDataInheritedContentJSON
// contains the JSON metadata for the struct
// [AccountUrlscannerV2GetScanResponseMetaProcessorsURLCategoriesDataInheritedContent]
type accountUrlscannerV2GetScanResponseMetaProcessorsURLCategoriesDataInheritedContentJSON struct {
	ID              apijson.Field
	Name            apijson.Field
	SuperCategoryID apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccountUrlscannerV2GetScanResponseMetaProcessorsURLCategoriesDataInheritedContent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerV2GetScanResponseMetaProcessorsURLCategoriesDataInheritedContentJSON) RawJSON() string {
	return r.raw
}

type AccountUrlscannerV2GetScanResponseMetaProcessorsURLCategoriesDataInheritedRisk struct {
	ID              float64                                                                            `json:"id" api:"required"`
	Name            string                                                                             `json:"name" api:"required"`
	SuperCategoryID float64                                                                            `json:"super_category_id" api:"required"`
	JSON            accountUrlscannerV2GetScanResponseMetaProcessorsURLCategoriesDataInheritedRiskJSON `json:"-"`
}

// accountUrlscannerV2GetScanResponseMetaProcessorsURLCategoriesDataInheritedRiskJSON
// contains the JSON metadata for the struct
// [AccountUrlscannerV2GetScanResponseMetaProcessorsURLCategoriesDataInheritedRisk]
type accountUrlscannerV2GetScanResponseMetaProcessorsURLCategoriesDataInheritedRiskJSON struct {
	ID              apijson.Field
	Name            apijson.Field
	SuperCategoryID apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccountUrlscannerV2GetScanResponseMetaProcessorsURLCategoriesDataInheritedRisk) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerV2GetScanResponseMetaProcessorsURLCategoriesDataInheritedRiskJSON) RawJSON() string {
	return r.raw
}

type AccountUrlscannerV2GetScanResponseMetaProcessorsURLCategoriesDataRisk struct {
	ID              float64                                                                   `json:"id" api:"required"`
	Name            string                                                                    `json:"name" api:"required"`
	SuperCategoryID float64                                                                   `json:"super_category_id" api:"required"`
	JSON            accountUrlscannerV2GetScanResponseMetaProcessorsURLCategoriesDataRiskJSON `json:"-"`
}

// accountUrlscannerV2GetScanResponseMetaProcessorsURLCategoriesDataRiskJSON
// contains the JSON metadata for the struct
// [AccountUrlscannerV2GetScanResponseMetaProcessorsURLCategoriesDataRisk]
type accountUrlscannerV2GetScanResponseMetaProcessorsURLCategoriesDataRiskJSON struct {
	ID              apijson.Field
	Name            apijson.Field
	SuperCategoryID apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccountUrlscannerV2GetScanResponseMetaProcessorsURLCategoriesDataRisk) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerV2GetScanResponseMetaProcessorsURLCategoriesDataRiskJSON) RawJSON() string {
	return r.raw
}

type AccountUrlscannerV2GetScanResponsePage struct {
	ApexDomain   string                                           `json:"apexDomain" api:"required"`
	Asn          string                                           `json:"asn" api:"required"`
	Asnname      string                                           `json:"asnname" api:"required"`
	City         string                                           `json:"city" api:"required"`
	Country      string                                           `json:"country" api:"required"`
	Domain       string                                           `json:"domain" api:"required"`
	IP           string                                           `json:"ip" api:"required"`
	MimeType     string                                           `json:"mimeType" api:"required"`
	Server       string                                           `json:"server" api:"required"`
	Status       string                                           `json:"status" api:"required"`
	Title        string                                           `json:"title" api:"required"`
	TlsAgeDays   float64                                          `json:"tlsAgeDays" api:"required"`
	TlsIssuer    string                                           `json:"tlsIssuer" api:"required"`
	TlsValidDays float64                                          `json:"tlsValidDays" api:"required"`
	TlsValidFrom string                                           `json:"tlsValidFrom" api:"required"`
	URL          string                                           `json:"url" api:"required"`
	Screenshot   AccountUrlscannerV2GetScanResponsePageScreenshot `json:"screenshot"`
	JSON         accountUrlscannerV2GetScanResponsePageJSON       `json:"-"`
}

// accountUrlscannerV2GetScanResponsePageJSON contains the JSON metadata for the
// struct [AccountUrlscannerV2GetScanResponsePage]
type accountUrlscannerV2GetScanResponsePageJSON struct {
	ApexDomain   apijson.Field
	Asn          apijson.Field
	Asnname      apijson.Field
	City         apijson.Field
	Country      apijson.Field
	Domain       apijson.Field
	IP           apijson.Field
	MimeType     apijson.Field
	Server       apijson.Field
	Status       apijson.Field
	Title        apijson.Field
	TlsAgeDays   apijson.Field
	TlsIssuer    apijson.Field
	TlsValidDays apijson.Field
	TlsValidFrom apijson.Field
	URL          apijson.Field
	Screenshot   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountUrlscannerV2GetScanResponsePage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerV2GetScanResponsePageJSON) RawJSON() string {
	return r.raw
}

type AccountUrlscannerV2GetScanResponsePageScreenshot struct {
	Dhash   string                                               `json:"dhash" api:"required"`
	Mm3Hash float64                                              `json:"mm3Hash" api:"required"`
	Name    string                                               `json:"name" api:"required"`
	Phash   string                                               `json:"phash" api:"required"`
	JSON    accountUrlscannerV2GetScanResponsePageScreenshotJSON `json:"-"`
}

// accountUrlscannerV2GetScanResponsePageScreenshotJSON contains the JSON metadata
// for the struct [AccountUrlscannerV2GetScanResponsePageScreenshot]
type accountUrlscannerV2GetScanResponsePageScreenshotJSON struct {
	Dhash       apijson.Field
	Mm3Hash     apijson.Field
	Name        apijson.Field
	Phash       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUrlscannerV2GetScanResponsePageScreenshot) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerV2GetScanResponsePageScreenshotJSON) RawJSON() string {
	return r.raw
}

type AccountUrlscannerV2GetScanResponseScanner struct {
	Colo    string                                        `json:"colo" api:"required"`
	Country string                                        `json:"country" api:"required"`
	JSON    accountUrlscannerV2GetScanResponseScannerJSON `json:"-"`
}

// accountUrlscannerV2GetScanResponseScannerJSON contains the JSON metadata for the
// struct [AccountUrlscannerV2GetScanResponseScanner]
type accountUrlscannerV2GetScanResponseScannerJSON struct {
	Colo        apijson.Field
	Country     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUrlscannerV2GetScanResponseScanner) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerV2GetScanResponseScannerJSON) RawJSON() string {
	return r.raw
}

type AccountUrlscannerV2GetScanResponseStats struct {
	DomainStats      []AccountUrlscannerV2GetScanResponseStatsDomainStat   `json:"domainStats" api:"required"`
	IPStats          []AccountUrlscannerV2GetScanResponseStatsIPStat       `json:"ipStats" api:"required"`
	IPv6Percentage   float64                                               `json:"IPv6Percentage" api:"required"`
	Malicious        float64                                               `json:"malicious" api:"required"`
	ProtocolStats    []AccountUrlscannerV2GetScanResponseStatsProtocolStat `json:"protocolStats" api:"required"`
	ResourceStats    []AccountUrlscannerV2GetScanResponseStatsResourceStat `json:"resourceStats" api:"required"`
	SecurePercentage float64                                               `json:"securePercentage" api:"required"`
	SecureRequests   float64                                               `json:"secureRequests" api:"required"`
	ServerStats      []AccountUrlscannerV2GetScanResponseStatsServerStat   `json:"serverStats" api:"required"`
	TlsStats         []AccountUrlscannerV2GetScanResponseStatsTlsStat      `json:"tlsStats" api:"required"`
	TotalLinks       float64                                               `json:"totalLinks" api:"required"`
	UniqAsNs         float64                                               `json:"uniqASNs" api:"required"`
	UniqCountries    float64                                               `json:"uniqCountries" api:"required"`
	JSON             accountUrlscannerV2GetScanResponseStatsJSON           `json:"-"`
}

// accountUrlscannerV2GetScanResponseStatsJSON contains the JSON metadata for the
// struct [AccountUrlscannerV2GetScanResponseStats]
type accountUrlscannerV2GetScanResponseStatsJSON struct {
	DomainStats      apijson.Field
	IPStats          apijson.Field
	IPv6Percentage   apijson.Field
	Malicious        apijson.Field
	ProtocolStats    apijson.Field
	ResourceStats    apijson.Field
	SecurePercentage apijson.Field
	SecureRequests   apijson.Field
	ServerStats      apijson.Field
	TlsStats         apijson.Field
	TotalLinks       apijson.Field
	UniqAsNs         apijson.Field
	UniqCountries    apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountUrlscannerV2GetScanResponseStats) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerV2GetScanResponseStatsJSON) RawJSON() string {
	return r.raw
}

type AccountUrlscannerV2GetScanResponseStatsDomainStat struct {
	Count       float64                                               `json:"count" api:"required"`
	Countries   []string                                              `json:"countries" api:"required"`
	Domain      string                                                `json:"domain" api:"required"`
	EncodedSize float64                                               `json:"encodedSize" api:"required"`
	Index       float64                                               `json:"index" api:"required"`
	Initiators  []string                                              `json:"initiators" api:"required"`
	IPs         []string                                              `json:"ips" api:"required"`
	Redirects   float64                                               `json:"redirects" api:"required"`
	Size        float64                                               `json:"size" api:"required"`
	JSON        accountUrlscannerV2GetScanResponseStatsDomainStatJSON `json:"-"`
}

// accountUrlscannerV2GetScanResponseStatsDomainStatJSON contains the JSON metadata
// for the struct [AccountUrlscannerV2GetScanResponseStatsDomainStat]
type accountUrlscannerV2GetScanResponseStatsDomainStatJSON struct {
	Count       apijson.Field
	Countries   apijson.Field
	Domain      apijson.Field
	EncodedSize apijson.Field
	Index       apijson.Field
	Initiators  apijson.Field
	IPs         apijson.Field
	Redirects   apijson.Field
	Size        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUrlscannerV2GetScanResponseStatsDomainStat) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerV2GetScanResponseStatsDomainStatJSON) RawJSON() string {
	return r.raw
}

type AccountUrlscannerV2GetScanResponseStatsIPStat struct {
	Asn         AccountUrlscannerV2GetScanResponseStatsIPStatsAsn   `json:"asn" api:"required"`
	Countries   []string                                            `json:"countries" api:"required"`
	Domains     []string                                            `json:"domains" api:"required"`
	EncodedSize float64                                             `json:"encodedSize" api:"required"`
	Geoip       AccountUrlscannerV2GetScanResponseStatsIPStatsGeoip `json:"geoip" api:"required"`
	Index       float64                                             `json:"index" api:"required"`
	IP          string                                              `json:"ip" api:"required"`
	Ipv6        bool                                                `json:"ipv6" api:"required"`
	Redirects   float64                                             `json:"redirects" api:"required"`
	Requests    float64                                             `json:"requests" api:"required"`
	Size        float64                                             `json:"size" api:"required"`
	Count       float64                                             `json:"count"`
	JSON        accountUrlscannerV2GetScanResponseStatsIPStatJSON   `json:"-"`
}

// accountUrlscannerV2GetScanResponseStatsIPStatJSON contains the JSON metadata for
// the struct [AccountUrlscannerV2GetScanResponseStatsIPStat]
type accountUrlscannerV2GetScanResponseStatsIPStatJSON struct {
	Asn         apijson.Field
	Countries   apijson.Field
	Domains     apijson.Field
	EncodedSize apijson.Field
	Geoip       apijson.Field
	Index       apijson.Field
	IP          apijson.Field
	Ipv6        apijson.Field
	Redirects   apijson.Field
	Requests    apijson.Field
	Size        apijson.Field
	Count       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUrlscannerV2GetScanResponseStatsIPStat) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerV2GetScanResponseStatsIPStatJSON) RawJSON() string {
	return r.raw
}

type AccountUrlscannerV2GetScanResponseStatsIPStatsAsn struct {
	Asn         string                                                `json:"asn" api:"required"`
	Country     string                                                `json:"country" api:"required"`
	Description string                                                `json:"description" api:"required"`
	IP          string                                                `json:"ip" api:"required"`
	Name        string                                                `json:"name" api:"required"`
	Org         string                                                `json:"org" api:"required"`
	JSON        accountUrlscannerV2GetScanResponseStatsIPStatsAsnJSON `json:"-"`
}

// accountUrlscannerV2GetScanResponseStatsIPStatsAsnJSON contains the JSON metadata
// for the struct [AccountUrlscannerV2GetScanResponseStatsIPStatsAsn]
type accountUrlscannerV2GetScanResponseStatsIPStatsAsnJSON struct {
	Asn         apijson.Field
	Country     apijson.Field
	Description apijson.Field
	IP          apijson.Field
	Name        apijson.Field
	Org         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUrlscannerV2GetScanResponseStatsIPStatsAsn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerV2GetScanResponseStatsIPStatsAsnJSON) RawJSON() string {
	return r.raw
}

type AccountUrlscannerV2GetScanResponseStatsIPStatsGeoip struct {
	City        string                                                  `json:"city" api:"required"`
	Country     string                                                  `json:"country" api:"required"`
	CountryName string                                                  `json:"country_name" api:"required"`
	Ll          []float64                                               `json:"ll" api:"required"`
	Region      string                                                  `json:"region" api:"required"`
	JSON        accountUrlscannerV2GetScanResponseStatsIPStatsGeoipJSON `json:"-"`
}

// accountUrlscannerV2GetScanResponseStatsIPStatsGeoipJSON contains the JSON
// metadata for the struct [AccountUrlscannerV2GetScanResponseStatsIPStatsGeoip]
type accountUrlscannerV2GetScanResponseStatsIPStatsGeoipJSON struct {
	City        apijson.Field
	Country     apijson.Field
	CountryName apijson.Field
	Ll          apijson.Field
	Region      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUrlscannerV2GetScanResponseStatsIPStatsGeoip) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerV2GetScanResponseStatsIPStatsGeoipJSON) RawJSON() string {
	return r.raw
}

type AccountUrlscannerV2GetScanResponseStatsProtocolStat struct {
	Count       float64                                                 `json:"count" api:"required"`
	Countries   []string                                                `json:"countries" api:"required"`
	EncodedSize float64                                                 `json:"encodedSize" api:"required"`
	IPs         []string                                                `json:"ips" api:"required"`
	Protocol    string                                                  `json:"protocol" api:"required"`
	Size        float64                                                 `json:"size" api:"required"`
	JSON        accountUrlscannerV2GetScanResponseStatsProtocolStatJSON `json:"-"`
}

// accountUrlscannerV2GetScanResponseStatsProtocolStatJSON contains the JSON
// metadata for the struct [AccountUrlscannerV2GetScanResponseStatsProtocolStat]
type accountUrlscannerV2GetScanResponseStatsProtocolStatJSON struct {
	Count       apijson.Field
	Countries   apijson.Field
	EncodedSize apijson.Field
	IPs         apijson.Field
	Protocol    apijson.Field
	Size        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUrlscannerV2GetScanResponseStatsProtocolStat) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerV2GetScanResponseStatsProtocolStatJSON) RawJSON() string {
	return r.raw
}

type AccountUrlscannerV2GetScanResponseStatsResourceStat struct {
	Compression float64                                                 `json:"compression" api:"required"`
	Count       float64                                                 `json:"count" api:"required"`
	Countries   []string                                                `json:"countries" api:"required"`
	EncodedSize float64                                                 `json:"encodedSize" api:"required"`
	IPs         []string                                                `json:"ips" api:"required"`
	Percentage  float64                                                 `json:"percentage" api:"required"`
	Size        float64                                                 `json:"size" api:"required"`
	Type        string                                                  `json:"type" api:"required"`
	JSON        accountUrlscannerV2GetScanResponseStatsResourceStatJSON `json:"-"`
}

// accountUrlscannerV2GetScanResponseStatsResourceStatJSON contains the JSON
// metadata for the struct [AccountUrlscannerV2GetScanResponseStatsResourceStat]
type accountUrlscannerV2GetScanResponseStatsResourceStatJSON struct {
	Compression apijson.Field
	Count       apijson.Field
	Countries   apijson.Field
	EncodedSize apijson.Field
	IPs         apijson.Field
	Percentage  apijson.Field
	Size        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUrlscannerV2GetScanResponseStatsResourceStat) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerV2GetScanResponseStatsResourceStatJSON) RawJSON() string {
	return r.raw
}

type AccountUrlscannerV2GetScanResponseStatsServerStat struct {
	Count       float64                                               `json:"count" api:"required"`
	Countries   []string                                              `json:"countries" api:"required"`
	EncodedSize float64                                               `json:"encodedSize" api:"required"`
	IPs         []string                                              `json:"ips" api:"required"`
	Server      string                                                `json:"server" api:"required"`
	Size        float64                                               `json:"size" api:"required"`
	JSON        accountUrlscannerV2GetScanResponseStatsServerStatJSON `json:"-"`
}

// accountUrlscannerV2GetScanResponseStatsServerStatJSON contains the JSON metadata
// for the struct [AccountUrlscannerV2GetScanResponseStatsServerStat]
type accountUrlscannerV2GetScanResponseStatsServerStatJSON struct {
	Count       apijson.Field
	Countries   apijson.Field
	EncodedSize apijson.Field
	IPs         apijson.Field
	Server      apijson.Field
	Size        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUrlscannerV2GetScanResponseStatsServerStat) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerV2GetScanResponseStatsServerStatJSON) RawJSON() string {
	return r.raw
}

type AccountUrlscannerV2GetScanResponseStatsTlsStat struct {
	Count         float64                                                  `json:"count" api:"required"`
	Countries     []string                                                 `json:"countries" api:"required"`
	EncodedSize   float64                                                  `json:"encodedSize" api:"required"`
	IPs           []string                                                 `json:"ips" api:"required"`
	Protocols     AccountUrlscannerV2GetScanResponseStatsTlsStatsProtocols `json:"protocols" api:"required"`
	SecurityState string                                                   `json:"securityState" api:"required"`
	Size          float64                                                  `json:"size" api:"required"`
	JSON          accountUrlscannerV2GetScanResponseStatsTlsStatJSON       `json:"-"`
}

// accountUrlscannerV2GetScanResponseStatsTlsStatJSON contains the JSON metadata
// for the struct [AccountUrlscannerV2GetScanResponseStatsTlsStat]
type accountUrlscannerV2GetScanResponseStatsTlsStatJSON struct {
	Count         apijson.Field
	Countries     apijson.Field
	EncodedSize   apijson.Field
	IPs           apijson.Field
	Protocols     apijson.Field
	SecurityState apijson.Field
	Size          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccountUrlscannerV2GetScanResponseStatsTlsStat) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerV2GetScanResponseStatsTlsStatJSON) RawJSON() string {
	return r.raw
}

type AccountUrlscannerV2GetScanResponseStatsTlsStatsProtocols struct {
	Tls1_3Aes128Gcm float64                                                      `json:"TLS 1.3 / AES_128_GCM" api:"required"`
	JSON            accountUrlscannerV2GetScanResponseStatsTlsStatsProtocolsJSON `json:"-"`
}

// accountUrlscannerV2GetScanResponseStatsTlsStatsProtocolsJSON contains the JSON
// metadata for the struct
// [AccountUrlscannerV2GetScanResponseStatsTlsStatsProtocols]
type accountUrlscannerV2GetScanResponseStatsTlsStatsProtocolsJSON struct {
	Tls1_3Aes128Gcm apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccountUrlscannerV2GetScanResponseStatsTlsStatsProtocols) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerV2GetScanResponseStatsTlsStatsProtocolsJSON) RawJSON() string {
	return r.raw
}

type AccountUrlscannerV2GetScanResponseTask struct {
	ApexDomain    string                                        `json:"apexDomain" api:"required"`
	Domain        string                                        `json:"domain" api:"required"`
	DomURL        string                                        `json:"domURL" api:"required"`
	Method        string                                        `json:"method" api:"required"`
	Options       AccountUrlscannerV2GetScanResponseTaskOptions `json:"options" api:"required"`
	ReportURL     string                                        `json:"reportURL" api:"required"`
	ScreenshotURL string                                        `json:"screenshotURL" api:"required"`
	Source        string                                        `json:"source" api:"required"`
	Success       bool                                          `json:"success" api:"required"`
	Time          string                                        `json:"time" api:"required"`
	URL           string                                        `json:"url" api:"required"`
	Uuid          string                                        `json:"uuid" api:"required"`
	Visibility    string                                        `json:"visibility" api:"required"`
	JSON          accountUrlscannerV2GetScanResponseTaskJSON    `json:"-"`
}

// accountUrlscannerV2GetScanResponseTaskJSON contains the JSON metadata for the
// struct [AccountUrlscannerV2GetScanResponseTask]
type accountUrlscannerV2GetScanResponseTaskJSON struct {
	ApexDomain    apijson.Field
	Domain        apijson.Field
	DomURL        apijson.Field
	Method        apijson.Field
	Options       apijson.Field
	ReportURL     apijson.Field
	ScreenshotURL apijson.Field
	Source        apijson.Field
	Success       apijson.Field
	Time          apijson.Field
	URL           apijson.Field
	Uuid          apijson.Field
	Visibility    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccountUrlscannerV2GetScanResponseTask) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerV2GetScanResponseTaskJSON) RawJSON() string {
	return r.raw
}

type AccountUrlscannerV2GetScanResponseTaskOptions struct {
	// Custom headers set.
	CustomHeaders          interface{}                                       `json:"customHeaders"`
	ScreenshotsResolutions []string                                          `json:"screenshotsResolutions"`
	JSON                   accountUrlscannerV2GetScanResponseTaskOptionsJSON `json:"-"`
}

// accountUrlscannerV2GetScanResponseTaskOptionsJSON contains the JSON metadata for
// the struct [AccountUrlscannerV2GetScanResponseTaskOptions]
type accountUrlscannerV2GetScanResponseTaskOptionsJSON struct {
	CustomHeaders          apijson.Field
	ScreenshotsResolutions apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccountUrlscannerV2GetScanResponseTaskOptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerV2GetScanResponseTaskOptionsJSON) RawJSON() string {
	return r.raw
}

type AccountUrlscannerV2GetScanResponseVerdicts struct {
	Overall AccountUrlscannerV2GetScanResponseVerdictsOverall `json:"overall" api:"required"`
	JSON    accountUrlscannerV2GetScanResponseVerdictsJSON    `json:"-"`
}

// accountUrlscannerV2GetScanResponseVerdictsJSON contains the JSON metadata for
// the struct [AccountUrlscannerV2GetScanResponseVerdicts]
type accountUrlscannerV2GetScanResponseVerdictsJSON struct {
	Overall     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUrlscannerV2GetScanResponseVerdicts) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerV2GetScanResponseVerdictsJSON) RawJSON() string {
	return r.raw
}

type AccountUrlscannerV2GetScanResponseVerdictsOverall struct {
	Categories  []string                                              `json:"categories" api:"required"`
	HasVerdicts bool                                                  `json:"hasVerdicts" api:"required"`
	Malicious   bool                                                  `json:"malicious" api:"required"`
	Tags        []string                                              `json:"tags" api:"required"`
	JSON        accountUrlscannerV2GetScanResponseVerdictsOverallJSON `json:"-"`
}

// accountUrlscannerV2GetScanResponseVerdictsOverallJSON contains the JSON metadata
// for the struct [AccountUrlscannerV2GetScanResponseVerdictsOverall]
type accountUrlscannerV2GetScanResponseVerdictsOverallJSON struct {
	Categories  apijson.Field
	HasVerdicts apijson.Field
	Malicious   apijson.Field
	Tags        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUrlscannerV2GetScanResponseVerdictsOverall) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerV2GetScanResponseVerdictsOverallJSON) RawJSON() string {
	return r.raw
}

type AccountUrlscannerV2SearchScansResponse struct {
	Results []AccountUrlscannerV2SearchScansResponseResult `json:"results" api:"required"`
	JSON    accountUrlscannerV2SearchScansResponseJSON     `json:"-"`
}

// accountUrlscannerV2SearchScansResponseJSON contains the JSON metadata for the
// struct [AccountUrlscannerV2SearchScansResponse]
type accountUrlscannerV2SearchScansResponseJSON struct {
	Results     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUrlscannerV2SearchScansResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerV2SearchScansResponseJSON) RawJSON() string {
	return r.raw
}

type AccountUrlscannerV2SearchScansResponseResult struct {
	ID       string                                                `json:"_id" api:"required"`
	Page     AccountUrlscannerV2SearchScansResponseResultsPage     `json:"page" api:"required"`
	Result   string                                                `json:"result" api:"required"`
	Stats    AccountUrlscannerV2SearchScansResponseResultsStats    `json:"stats" api:"required"`
	Task     AccountUrlscannerV2SearchScansResponseResultsTask     `json:"task" api:"required"`
	Verdicts AccountUrlscannerV2SearchScansResponseResultsVerdicts `json:"verdicts" api:"required"`
	JSON     accountUrlscannerV2SearchScansResponseResultJSON      `json:"-"`
}

// accountUrlscannerV2SearchScansResponseResultJSON contains the JSON metadata for
// the struct [AccountUrlscannerV2SearchScansResponseResult]
type accountUrlscannerV2SearchScansResponseResultJSON struct {
	ID          apijson.Field
	Page        apijson.Field
	Result      apijson.Field
	Stats       apijson.Field
	Task        apijson.Field
	Verdicts    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUrlscannerV2SearchScansResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerV2SearchScansResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountUrlscannerV2SearchScansResponseResultsPage struct {
	Asn     string                                                `json:"asn" api:"required"`
	Country string                                                `json:"country" api:"required"`
	IP      string                                                `json:"ip" api:"required"`
	URL     string                                                `json:"url" api:"required"`
	JSON    accountUrlscannerV2SearchScansResponseResultsPageJSON `json:"-"`
}

// accountUrlscannerV2SearchScansResponseResultsPageJSON contains the JSON metadata
// for the struct [AccountUrlscannerV2SearchScansResponseResultsPage]
type accountUrlscannerV2SearchScansResponseResultsPageJSON struct {
	Asn         apijson.Field
	Country     apijson.Field
	IP          apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUrlscannerV2SearchScansResponseResultsPage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerV2SearchScansResponseResultsPageJSON) RawJSON() string {
	return r.raw
}

type AccountUrlscannerV2SearchScansResponseResultsStats struct {
	DataLength    float64                                                `json:"dataLength" api:"required"`
	Requests      float64                                                `json:"requests" api:"required"`
	UniqCountries float64                                                `json:"uniqCountries" api:"required"`
	UniqIPs       float64                                                `json:"uniqIPs" api:"required"`
	JSON          accountUrlscannerV2SearchScansResponseResultsStatsJSON `json:"-"`
}

// accountUrlscannerV2SearchScansResponseResultsStatsJSON contains the JSON
// metadata for the struct [AccountUrlscannerV2SearchScansResponseResultsStats]
type accountUrlscannerV2SearchScansResponseResultsStatsJSON struct {
	DataLength    apijson.Field
	Requests      apijson.Field
	UniqCountries apijson.Field
	UniqIPs       apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccountUrlscannerV2SearchScansResponseResultsStats) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerV2SearchScansResponseResultsStatsJSON) RawJSON() string {
	return r.raw
}

type AccountUrlscannerV2SearchScansResponseResultsTask struct {
	Time       string                                                `json:"time" api:"required"`
	URL        string                                                `json:"url" api:"required"`
	Uuid       string                                                `json:"uuid" api:"required"`
	Visibility string                                                `json:"visibility" api:"required"`
	JSON       accountUrlscannerV2SearchScansResponseResultsTaskJSON `json:"-"`
}

// accountUrlscannerV2SearchScansResponseResultsTaskJSON contains the JSON metadata
// for the struct [AccountUrlscannerV2SearchScansResponseResultsTask]
type accountUrlscannerV2SearchScansResponseResultsTaskJSON struct {
	Time        apijson.Field
	URL         apijson.Field
	Uuid        apijson.Field
	Visibility  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUrlscannerV2SearchScansResponseResultsTask) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerV2SearchScansResponseResultsTaskJSON) RawJSON() string {
	return r.raw
}

type AccountUrlscannerV2SearchScansResponseResultsVerdicts struct {
	Malicious bool                                                      `json:"malicious" api:"required"`
	JSON      accountUrlscannerV2SearchScansResponseResultsVerdictsJSON `json:"-"`
}

// accountUrlscannerV2SearchScansResponseResultsVerdictsJSON contains the JSON
// metadata for the struct [AccountUrlscannerV2SearchScansResponseResultsVerdicts]
type accountUrlscannerV2SearchScansResponseResultsVerdictsJSON struct {
	Malicious   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUrlscannerV2SearchScansResponseResultsVerdicts) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerV2SearchScansResponseResultsVerdictsJSON) RawJSON() string {
	return r.raw
}

type AccountUrlscannerV2BulkNewScansParams struct {
	// List of urls to scan (up to a 100).
	Body []AccountUrlscannerV2BulkNewScansParamsBody `json:"body"`
}

func (r AccountUrlscannerV2BulkNewScansParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type AccountUrlscannerV2BulkNewScansParamsBody struct {
	URL         param.Field[string] `json:"url" api:"required"`
	Customagent param.Field[string] `json:"customagent"`
	// Set custom headers.
	CustomHeaders param.Field[map[string]string] `json:"customHeaders"`
	Referer       param.Field[string]            `json:"referer"`
	// Take multiple screenshots targeting different device types.
	ScreenshotsResolutions param.Field[[]AccountUrlscannerV2BulkNewScansParamsBodyScreenshotsResolution] `json:"screenshotsResolutions"`
	// The option `Public` means it will be included in listings like recent scans and
	// search results. `Unlisted` means it will not be included in the aforementioned
	// listings, users will need to have the scan's ID to access it. A a scan will be
	// automatically marked as unlisted if it fails, if it contains potential PII or
	// other sensitive material.
	Visibility param.Field[AccountUrlscannerV2BulkNewScansParamsBodyVisibility] `json:"visibility"`
}

func (r AccountUrlscannerV2BulkNewScansParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Device resolutions.
type AccountUrlscannerV2BulkNewScansParamsBodyScreenshotsResolution string

const (
	AccountUrlscannerV2BulkNewScansParamsBodyScreenshotsResolutionDesktop AccountUrlscannerV2BulkNewScansParamsBodyScreenshotsResolution = "desktop"
	AccountUrlscannerV2BulkNewScansParamsBodyScreenshotsResolutionMobile  AccountUrlscannerV2BulkNewScansParamsBodyScreenshotsResolution = "mobile"
	AccountUrlscannerV2BulkNewScansParamsBodyScreenshotsResolutionTablet  AccountUrlscannerV2BulkNewScansParamsBodyScreenshotsResolution = "tablet"
)

func (r AccountUrlscannerV2BulkNewScansParamsBodyScreenshotsResolution) IsKnown() bool {
	switch r {
	case AccountUrlscannerV2BulkNewScansParamsBodyScreenshotsResolutionDesktop, AccountUrlscannerV2BulkNewScansParamsBodyScreenshotsResolutionMobile, AccountUrlscannerV2BulkNewScansParamsBodyScreenshotsResolutionTablet:
		return true
	}
	return false
}

// The option `Public` means it will be included in listings like recent scans and
// search results. `Unlisted` means it will not be included in the aforementioned
// listings, users will need to have the scan's ID to access it. A a scan will be
// automatically marked as unlisted if it fails, if it contains potential PII or
// other sensitive material.
type AccountUrlscannerV2BulkNewScansParamsBodyVisibility string

const (
	AccountUrlscannerV2BulkNewScansParamsBodyVisibilityPublic   AccountUrlscannerV2BulkNewScansParamsBodyVisibility = "Public"
	AccountUrlscannerV2BulkNewScansParamsBodyVisibilityUnlisted AccountUrlscannerV2BulkNewScansParamsBodyVisibility = "Unlisted"
)

func (r AccountUrlscannerV2BulkNewScansParamsBodyVisibility) IsKnown() bool {
	switch r {
	case AccountUrlscannerV2BulkNewScansParamsBodyVisibilityPublic, AccountUrlscannerV2BulkNewScansParamsBodyVisibilityUnlisted:
		return true
	}
	return false
}

type AccountUrlscannerV2NewScanParams struct {
	URL param.Field[string] `json:"url" api:"required"`
	// Country to geo egress from
	Country     param.Field[AccountUrlscannerV2NewScanParamsCountry] `json:"country"`
	Customagent param.Field[string]                                  `json:"customagent"`
	// Set custom headers.
	CustomHeaders param.Field[map[string]string] `json:"customHeaders"`
	Referer       param.Field[string]            `json:"referer"`
	// Take multiple screenshots targeting different device types.
	ScreenshotsResolutions param.Field[[]AccountUrlscannerV2NewScanParamsScreenshotsResolution] `json:"screenshotsResolutions"`
	// The option `Public` means it will be included in listings like recent scans and
	// search results. `Unlisted` means it will not be included in the aforementioned
	// listings, users will need to have the scan's ID to access it. A a scan will be
	// automatically marked as unlisted if it fails, if it contains potential PII or
	// other sensitive material.
	Visibility param.Field[AccountUrlscannerV2NewScanParamsVisibility] `json:"visibility"`
}

func (r AccountUrlscannerV2NewScanParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Country to geo egress from
type AccountUrlscannerV2NewScanParamsCountry string

const (
	AccountUrlscannerV2NewScanParamsCountryAf AccountUrlscannerV2NewScanParamsCountry = "AF"
	AccountUrlscannerV2NewScanParamsCountryAl AccountUrlscannerV2NewScanParamsCountry = "AL"
	AccountUrlscannerV2NewScanParamsCountryDz AccountUrlscannerV2NewScanParamsCountry = "DZ"
	AccountUrlscannerV2NewScanParamsCountryAd AccountUrlscannerV2NewScanParamsCountry = "AD"
	AccountUrlscannerV2NewScanParamsCountryAo AccountUrlscannerV2NewScanParamsCountry = "AO"
	AccountUrlscannerV2NewScanParamsCountryAg AccountUrlscannerV2NewScanParamsCountry = "AG"
	AccountUrlscannerV2NewScanParamsCountryAr AccountUrlscannerV2NewScanParamsCountry = "AR"
	AccountUrlscannerV2NewScanParamsCountryAm AccountUrlscannerV2NewScanParamsCountry = "AM"
	AccountUrlscannerV2NewScanParamsCountryAu AccountUrlscannerV2NewScanParamsCountry = "AU"
	AccountUrlscannerV2NewScanParamsCountryAt AccountUrlscannerV2NewScanParamsCountry = "AT"
	AccountUrlscannerV2NewScanParamsCountryAz AccountUrlscannerV2NewScanParamsCountry = "AZ"
	AccountUrlscannerV2NewScanParamsCountryBh AccountUrlscannerV2NewScanParamsCountry = "BH"
	AccountUrlscannerV2NewScanParamsCountryBd AccountUrlscannerV2NewScanParamsCountry = "BD"
	AccountUrlscannerV2NewScanParamsCountryBb AccountUrlscannerV2NewScanParamsCountry = "BB"
	AccountUrlscannerV2NewScanParamsCountryBy AccountUrlscannerV2NewScanParamsCountry = "BY"
	AccountUrlscannerV2NewScanParamsCountryBe AccountUrlscannerV2NewScanParamsCountry = "BE"
	AccountUrlscannerV2NewScanParamsCountryBz AccountUrlscannerV2NewScanParamsCountry = "BZ"
	AccountUrlscannerV2NewScanParamsCountryBj AccountUrlscannerV2NewScanParamsCountry = "BJ"
	AccountUrlscannerV2NewScanParamsCountryBm AccountUrlscannerV2NewScanParamsCountry = "BM"
	AccountUrlscannerV2NewScanParamsCountryBt AccountUrlscannerV2NewScanParamsCountry = "BT"
	AccountUrlscannerV2NewScanParamsCountryBo AccountUrlscannerV2NewScanParamsCountry = "BO"
	AccountUrlscannerV2NewScanParamsCountryBa AccountUrlscannerV2NewScanParamsCountry = "BA"
	AccountUrlscannerV2NewScanParamsCountryBw AccountUrlscannerV2NewScanParamsCountry = "BW"
	AccountUrlscannerV2NewScanParamsCountryBr AccountUrlscannerV2NewScanParamsCountry = "BR"
	AccountUrlscannerV2NewScanParamsCountryBn AccountUrlscannerV2NewScanParamsCountry = "BN"
	AccountUrlscannerV2NewScanParamsCountryBg AccountUrlscannerV2NewScanParamsCountry = "BG"
	AccountUrlscannerV2NewScanParamsCountryBf AccountUrlscannerV2NewScanParamsCountry = "BF"
	AccountUrlscannerV2NewScanParamsCountryBi AccountUrlscannerV2NewScanParamsCountry = "BI"
	AccountUrlscannerV2NewScanParamsCountryKh AccountUrlscannerV2NewScanParamsCountry = "KH"
	AccountUrlscannerV2NewScanParamsCountryCm AccountUrlscannerV2NewScanParamsCountry = "CM"
	AccountUrlscannerV2NewScanParamsCountryCa AccountUrlscannerV2NewScanParamsCountry = "CA"
	AccountUrlscannerV2NewScanParamsCountryCv AccountUrlscannerV2NewScanParamsCountry = "CV"
	AccountUrlscannerV2NewScanParamsCountryKy AccountUrlscannerV2NewScanParamsCountry = "KY"
	AccountUrlscannerV2NewScanParamsCountryCf AccountUrlscannerV2NewScanParamsCountry = "CF"
	AccountUrlscannerV2NewScanParamsCountryTd AccountUrlscannerV2NewScanParamsCountry = "TD"
	AccountUrlscannerV2NewScanParamsCountryCl AccountUrlscannerV2NewScanParamsCountry = "CL"
	AccountUrlscannerV2NewScanParamsCountryCn AccountUrlscannerV2NewScanParamsCountry = "CN"
	AccountUrlscannerV2NewScanParamsCountryCo AccountUrlscannerV2NewScanParamsCountry = "CO"
	AccountUrlscannerV2NewScanParamsCountryKm AccountUrlscannerV2NewScanParamsCountry = "KM"
	AccountUrlscannerV2NewScanParamsCountryCg AccountUrlscannerV2NewScanParamsCountry = "CG"
	AccountUrlscannerV2NewScanParamsCountryCr AccountUrlscannerV2NewScanParamsCountry = "CR"
	AccountUrlscannerV2NewScanParamsCountryCi AccountUrlscannerV2NewScanParamsCountry = "CI"
	AccountUrlscannerV2NewScanParamsCountryHr AccountUrlscannerV2NewScanParamsCountry = "HR"
	AccountUrlscannerV2NewScanParamsCountryCu AccountUrlscannerV2NewScanParamsCountry = "CU"
	AccountUrlscannerV2NewScanParamsCountryCy AccountUrlscannerV2NewScanParamsCountry = "CY"
	AccountUrlscannerV2NewScanParamsCountryCz AccountUrlscannerV2NewScanParamsCountry = "CZ"
	AccountUrlscannerV2NewScanParamsCountryCd AccountUrlscannerV2NewScanParamsCountry = "CD"
	AccountUrlscannerV2NewScanParamsCountryDk AccountUrlscannerV2NewScanParamsCountry = "DK"
	AccountUrlscannerV2NewScanParamsCountryDj AccountUrlscannerV2NewScanParamsCountry = "DJ"
	AccountUrlscannerV2NewScanParamsCountryDm AccountUrlscannerV2NewScanParamsCountry = "DM"
	AccountUrlscannerV2NewScanParamsCountryDo AccountUrlscannerV2NewScanParamsCountry = "DO"
	AccountUrlscannerV2NewScanParamsCountryEc AccountUrlscannerV2NewScanParamsCountry = "EC"
	AccountUrlscannerV2NewScanParamsCountryEg AccountUrlscannerV2NewScanParamsCountry = "EG"
	AccountUrlscannerV2NewScanParamsCountrySv AccountUrlscannerV2NewScanParamsCountry = "SV"
	AccountUrlscannerV2NewScanParamsCountryGq AccountUrlscannerV2NewScanParamsCountry = "GQ"
	AccountUrlscannerV2NewScanParamsCountryEr AccountUrlscannerV2NewScanParamsCountry = "ER"
	AccountUrlscannerV2NewScanParamsCountryEe AccountUrlscannerV2NewScanParamsCountry = "EE"
	AccountUrlscannerV2NewScanParamsCountrySz AccountUrlscannerV2NewScanParamsCountry = "SZ"
	AccountUrlscannerV2NewScanParamsCountryEt AccountUrlscannerV2NewScanParamsCountry = "ET"
	AccountUrlscannerV2NewScanParamsCountryFj AccountUrlscannerV2NewScanParamsCountry = "FJ"
	AccountUrlscannerV2NewScanParamsCountryFi AccountUrlscannerV2NewScanParamsCountry = "FI"
	AccountUrlscannerV2NewScanParamsCountryFr AccountUrlscannerV2NewScanParamsCountry = "FR"
	AccountUrlscannerV2NewScanParamsCountryGa AccountUrlscannerV2NewScanParamsCountry = "GA"
	AccountUrlscannerV2NewScanParamsCountryGe AccountUrlscannerV2NewScanParamsCountry = "GE"
	AccountUrlscannerV2NewScanParamsCountryDe AccountUrlscannerV2NewScanParamsCountry = "DE"
	AccountUrlscannerV2NewScanParamsCountryGh AccountUrlscannerV2NewScanParamsCountry = "GH"
	AccountUrlscannerV2NewScanParamsCountryGr AccountUrlscannerV2NewScanParamsCountry = "GR"
	AccountUrlscannerV2NewScanParamsCountryGl AccountUrlscannerV2NewScanParamsCountry = "GL"
	AccountUrlscannerV2NewScanParamsCountryGd AccountUrlscannerV2NewScanParamsCountry = "GD"
	AccountUrlscannerV2NewScanParamsCountryGt AccountUrlscannerV2NewScanParamsCountry = "GT"
	AccountUrlscannerV2NewScanParamsCountryGn AccountUrlscannerV2NewScanParamsCountry = "GN"
	AccountUrlscannerV2NewScanParamsCountryGw AccountUrlscannerV2NewScanParamsCountry = "GW"
	AccountUrlscannerV2NewScanParamsCountryGy AccountUrlscannerV2NewScanParamsCountry = "GY"
	AccountUrlscannerV2NewScanParamsCountryHt AccountUrlscannerV2NewScanParamsCountry = "HT"
	AccountUrlscannerV2NewScanParamsCountryHn AccountUrlscannerV2NewScanParamsCountry = "HN"
	AccountUrlscannerV2NewScanParamsCountryHu AccountUrlscannerV2NewScanParamsCountry = "HU"
	AccountUrlscannerV2NewScanParamsCountryIs AccountUrlscannerV2NewScanParamsCountry = "IS"
	AccountUrlscannerV2NewScanParamsCountryIn AccountUrlscannerV2NewScanParamsCountry = "IN"
	AccountUrlscannerV2NewScanParamsCountryID AccountUrlscannerV2NewScanParamsCountry = "ID"
	AccountUrlscannerV2NewScanParamsCountryIr AccountUrlscannerV2NewScanParamsCountry = "IR"
	AccountUrlscannerV2NewScanParamsCountryIq AccountUrlscannerV2NewScanParamsCountry = "IQ"
	AccountUrlscannerV2NewScanParamsCountryIe AccountUrlscannerV2NewScanParamsCountry = "IE"
	AccountUrlscannerV2NewScanParamsCountryIl AccountUrlscannerV2NewScanParamsCountry = "IL"
	AccountUrlscannerV2NewScanParamsCountryIt AccountUrlscannerV2NewScanParamsCountry = "IT"
	AccountUrlscannerV2NewScanParamsCountryJm AccountUrlscannerV2NewScanParamsCountry = "JM"
	AccountUrlscannerV2NewScanParamsCountryJp AccountUrlscannerV2NewScanParamsCountry = "JP"
	AccountUrlscannerV2NewScanParamsCountryJo AccountUrlscannerV2NewScanParamsCountry = "JO"
	AccountUrlscannerV2NewScanParamsCountryKz AccountUrlscannerV2NewScanParamsCountry = "KZ"
	AccountUrlscannerV2NewScanParamsCountryKe AccountUrlscannerV2NewScanParamsCountry = "KE"
	AccountUrlscannerV2NewScanParamsCountryKi AccountUrlscannerV2NewScanParamsCountry = "KI"
	AccountUrlscannerV2NewScanParamsCountryKw AccountUrlscannerV2NewScanParamsCountry = "KW"
	AccountUrlscannerV2NewScanParamsCountryKg AccountUrlscannerV2NewScanParamsCountry = "KG"
	AccountUrlscannerV2NewScanParamsCountryLa AccountUrlscannerV2NewScanParamsCountry = "LA"
	AccountUrlscannerV2NewScanParamsCountryLv AccountUrlscannerV2NewScanParamsCountry = "LV"
	AccountUrlscannerV2NewScanParamsCountryLb AccountUrlscannerV2NewScanParamsCountry = "LB"
	AccountUrlscannerV2NewScanParamsCountryLs AccountUrlscannerV2NewScanParamsCountry = "LS"
	AccountUrlscannerV2NewScanParamsCountryLr AccountUrlscannerV2NewScanParamsCountry = "LR"
	AccountUrlscannerV2NewScanParamsCountryLy AccountUrlscannerV2NewScanParamsCountry = "LY"
	AccountUrlscannerV2NewScanParamsCountryLi AccountUrlscannerV2NewScanParamsCountry = "LI"
	AccountUrlscannerV2NewScanParamsCountryLt AccountUrlscannerV2NewScanParamsCountry = "LT"
	AccountUrlscannerV2NewScanParamsCountryLu AccountUrlscannerV2NewScanParamsCountry = "LU"
	AccountUrlscannerV2NewScanParamsCountryMo AccountUrlscannerV2NewScanParamsCountry = "MO"
	AccountUrlscannerV2NewScanParamsCountryMg AccountUrlscannerV2NewScanParamsCountry = "MG"
	AccountUrlscannerV2NewScanParamsCountryMw AccountUrlscannerV2NewScanParamsCountry = "MW"
	AccountUrlscannerV2NewScanParamsCountryMy AccountUrlscannerV2NewScanParamsCountry = "MY"
	AccountUrlscannerV2NewScanParamsCountryMv AccountUrlscannerV2NewScanParamsCountry = "MV"
	AccountUrlscannerV2NewScanParamsCountryMl AccountUrlscannerV2NewScanParamsCountry = "ML"
	AccountUrlscannerV2NewScanParamsCountryMr AccountUrlscannerV2NewScanParamsCountry = "MR"
	AccountUrlscannerV2NewScanParamsCountryMu AccountUrlscannerV2NewScanParamsCountry = "MU"
	AccountUrlscannerV2NewScanParamsCountryMx AccountUrlscannerV2NewScanParamsCountry = "MX"
	AccountUrlscannerV2NewScanParamsCountryFm AccountUrlscannerV2NewScanParamsCountry = "FM"
	AccountUrlscannerV2NewScanParamsCountryMd AccountUrlscannerV2NewScanParamsCountry = "MD"
	AccountUrlscannerV2NewScanParamsCountryMc AccountUrlscannerV2NewScanParamsCountry = "MC"
	AccountUrlscannerV2NewScanParamsCountryMn AccountUrlscannerV2NewScanParamsCountry = "MN"
	AccountUrlscannerV2NewScanParamsCountryMs AccountUrlscannerV2NewScanParamsCountry = "MS"
	AccountUrlscannerV2NewScanParamsCountryMa AccountUrlscannerV2NewScanParamsCountry = "MA"
	AccountUrlscannerV2NewScanParamsCountryMz AccountUrlscannerV2NewScanParamsCountry = "MZ"
	AccountUrlscannerV2NewScanParamsCountryMm AccountUrlscannerV2NewScanParamsCountry = "MM"
	AccountUrlscannerV2NewScanParamsCountryNa AccountUrlscannerV2NewScanParamsCountry = "NA"
	AccountUrlscannerV2NewScanParamsCountryNr AccountUrlscannerV2NewScanParamsCountry = "NR"
	AccountUrlscannerV2NewScanParamsCountryNp AccountUrlscannerV2NewScanParamsCountry = "NP"
	AccountUrlscannerV2NewScanParamsCountryNl AccountUrlscannerV2NewScanParamsCountry = "NL"
	AccountUrlscannerV2NewScanParamsCountryNz AccountUrlscannerV2NewScanParamsCountry = "NZ"
	AccountUrlscannerV2NewScanParamsCountryNi AccountUrlscannerV2NewScanParamsCountry = "NI"
	AccountUrlscannerV2NewScanParamsCountryNe AccountUrlscannerV2NewScanParamsCountry = "NE"
	AccountUrlscannerV2NewScanParamsCountryNg AccountUrlscannerV2NewScanParamsCountry = "NG"
	AccountUrlscannerV2NewScanParamsCountryKp AccountUrlscannerV2NewScanParamsCountry = "KP"
	AccountUrlscannerV2NewScanParamsCountryMk AccountUrlscannerV2NewScanParamsCountry = "MK"
	AccountUrlscannerV2NewScanParamsCountryNo AccountUrlscannerV2NewScanParamsCountry = "NO"
	AccountUrlscannerV2NewScanParamsCountryOm AccountUrlscannerV2NewScanParamsCountry = "OM"
	AccountUrlscannerV2NewScanParamsCountryPk AccountUrlscannerV2NewScanParamsCountry = "PK"
	AccountUrlscannerV2NewScanParamsCountryPs AccountUrlscannerV2NewScanParamsCountry = "PS"
	AccountUrlscannerV2NewScanParamsCountryPa AccountUrlscannerV2NewScanParamsCountry = "PA"
	AccountUrlscannerV2NewScanParamsCountryPg AccountUrlscannerV2NewScanParamsCountry = "PG"
	AccountUrlscannerV2NewScanParamsCountryPy AccountUrlscannerV2NewScanParamsCountry = "PY"
	AccountUrlscannerV2NewScanParamsCountryPe AccountUrlscannerV2NewScanParamsCountry = "PE"
	AccountUrlscannerV2NewScanParamsCountryPh AccountUrlscannerV2NewScanParamsCountry = "PH"
	AccountUrlscannerV2NewScanParamsCountryPl AccountUrlscannerV2NewScanParamsCountry = "PL"
	AccountUrlscannerV2NewScanParamsCountryPt AccountUrlscannerV2NewScanParamsCountry = "PT"
	AccountUrlscannerV2NewScanParamsCountryQa AccountUrlscannerV2NewScanParamsCountry = "QA"
	AccountUrlscannerV2NewScanParamsCountryRo AccountUrlscannerV2NewScanParamsCountry = "RO"
	AccountUrlscannerV2NewScanParamsCountryRu AccountUrlscannerV2NewScanParamsCountry = "RU"
	AccountUrlscannerV2NewScanParamsCountryRw AccountUrlscannerV2NewScanParamsCountry = "RW"
	AccountUrlscannerV2NewScanParamsCountrySh AccountUrlscannerV2NewScanParamsCountry = "SH"
	AccountUrlscannerV2NewScanParamsCountryKn AccountUrlscannerV2NewScanParamsCountry = "KN"
	AccountUrlscannerV2NewScanParamsCountryLc AccountUrlscannerV2NewScanParamsCountry = "LC"
	AccountUrlscannerV2NewScanParamsCountryVc AccountUrlscannerV2NewScanParamsCountry = "VC"
	AccountUrlscannerV2NewScanParamsCountryWs AccountUrlscannerV2NewScanParamsCountry = "WS"
	AccountUrlscannerV2NewScanParamsCountrySm AccountUrlscannerV2NewScanParamsCountry = "SM"
	AccountUrlscannerV2NewScanParamsCountrySt AccountUrlscannerV2NewScanParamsCountry = "ST"
	AccountUrlscannerV2NewScanParamsCountrySa AccountUrlscannerV2NewScanParamsCountry = "SA"
	AccountUrlscannerV2NewScanParamsCountrySn AccountUrlscannerV2NewScanParamsCountry = "SN"
	AccountUrlscannerV2NewScanParamsCountryRs AccountUrlscannerV2NewScanParamsCountry = "RS"
	AccountUrlscannerV2NewScanParamsCountrySc AccountUrlscannerV2NewScanParamsCountry = "SC"
	AccountUrlscannerV2NewScanParamsCountrySl AccountUrlscannerV2NewScanParamsCountry = "SL"
	AccountUrlscannerV2NewScanParamsCountrySk AccountUrlscannerV2NewScanParamsCountry = "SK"
	AccountUrlscannerV2NewScanParamsCountrySi AccountUrlscannerV2NewScanParamsCountry = "SI"
	AccountUrlscannerV2NewScanParamsCountrySb AccountUrlscannerV2NewScanParamsCountry = "SB"
	AccountUrlscannerV2NewScanParamsCountrySo AccountUrlscannerV2NewScanParamsCountry = "SO"
	AccountUrlscannerV2NewScanParamsCountryZa AccountUrlscannerV2NewScanParamsCountry = "ZA"
	AccountUrlscannerV2NewScanParamsCountryKr AccountUrlscannerV2NewScanParamsCountry = "KR"
	AccountUrlscannerV2NewScanParamsCountrySS AccountUrlscannerV2NewScanParamsCountry = "SS"
	AccountUrlscannerV2NewScanParamsCountryEs AccountUrlscannerV2NewScanParamsCountry = "ES"
	AccountUrlscannerV2NewScanParamsCountryLk AccountUrlscannerV2NewScanParamsCountry = "LK"
	AccountUrlscannerV2NewScanParamsCountrySd AccountUrlscannerV2NewScanParamsCountry = "SD"
	AccountUrlscannerV2NewScanParamsCountrySr AccountUrlscannerV2NewScanParamsCountry = "SR"
	AccountUrlscannerV2NewScanParamsCountrySe AccountUrlscannerV2NewScanParamsCountry = "SE"
	AccountUrlscannerV2NewScanParamsCountryCh AccountUrlscannerV2NewScanParamsCountry = "CH"
	AccountUrlscannerV2NewScanParamsCountrySy AccountUrlscannerV2NewScanParamsCountry = "SY"
	AccountUrlscannerV2NewScanParamsCountryTw AccountUrlscannerV2NewScanParamsCountry = "TW"
	AccountUrlscannerV2NewScanParamsCountryTj AccountUrlscannerV2NewScanParamsCountry = "TJ"
	AccountUrlscannerV2NewScanParamsCountryTz AccountUrlscannerV2NewScanParamsCountry = "TZ"
	AccountUrlscannerV2NewScanParamsCountryTh AccountUrlscannerV2NewScanParamsCountry = "TH"
	AccountUrlscannerV2NewScanParamsCountryBs AccountUrlscannerV2NewScanParamsCountry = "BS"
	AccountUrlscannerV2NewScanParamsCountryGm AccountUrlscannerV2NewScanParamsCountry = "GM"
	AccountUrlscannerV2NewScanParamsCountryTl AccountUrlscannerV2NewScanParamsCountry = "TL"
	AccountUrlscannerV2NewScanParamsCountryTg AccountUrlscannerV2NewScanParamsCountry = "TG"
	AccountUrlscannerV2NewScanParamsCountryTo AccountUrlscannerV2NewScanParamsCountry = "TO"
	AccountUrlscannerV2NewScanParamsCountryTt AccountUrlscannerV2NewScanParamsCountry = "TT"
	AccountUrlscannerV2NewScanParamsCountryTn AccountUrlscannerV2NewScanParamsCountry = "TN"
	AccountUrlscannerV2NewScanParamsCountryTr AccountUrlscannerV2NewScanParamsCountry = "TR"
	AccountUrlscannerV2NewScanParamsCountryTm AccountUrlscannerV2NewScanParamsCountry = "TM"
	AccountUrlscannerV2NewScanParamsCountryUg AccountUrlscannerV2NewScanParamsCountry = "UG"
	AccountUrlscannerV2NewScanParamsCountryUa AccountUrlscannerV2NewScanParamsCountry = "UA"
	AccountUrlscannerV2NewScanParamsCountryAe AccountUrlscannerV2NewScanParamsCountry = "AE"
	AccountUrlscannerV2NewScanParamsCountryGB AccountUrlscannerV2NewScanParamsCountry = "GB"
	AccountUrlscannerV2NewScanParamsCountryUs AccountUrlscannerV2NewScanParamsCountry = "US"
	AccountUrlscannerV2NewScanParamsCountryUy AccountUrlscannerV2NewScanParamsCountry = "UY"
	AccountUrlscannerV2NewScanParamsCountryUz AccountUrlscannerV2NewScanParamsCountry = "UZ"
	AccountUrlscannerV2NewScanParamsCountryVu AccountUrlscannerV2NewScanParamsCountry = "VU"
	AccountUrlscannerV2NewScanParamsCountryVe AccountUrlscannerV2NewScanParamsCountry = "VE"
	AccountUrlscannerV2NewScanParamsCountryVn AccountUrlscannerV2NewScanParamsCountry = "VN"
	AccountUrlscannerV2NewScanParamsCountryYe AccountUrlscannerV2NewScanParamsCountry = "YE"
	AccountUrlscannerV2NewScanParamsCountryZm AccountUrlscannerV2NewScanParamsCountry = "ZM"
	AccountUrlscannerV2NewScanParamsCountryZw AccountUrlscannerV2NewScanParamsCountry = "ZW"
)

func (r AccountUrlscannerV2NewScanParamsCountry) IsKnown() bool {
	switch r {
	case AccountUrlscannerV2NewScanParamsCountryAf, AccountUrlscannerV2NewScanParamsCountryAl, AccountUrlscannerV2NewScanParamsCountryDz, AccountUrlscannerV2NewScanParamsCountryAd, AccountUrlscannerV2NewScanParamsCountryAo, AccountUrlscannerV2NewScanParamsCountryAg, AccountUrlscannerV2NewScanParamsCountryAr, AccountUrlscannerV2NewScanParamsCountryAm, AccountUrlscannerV2NewScanParamsCountryAu, AccountUrlscannerV2NewScanParamsCountryAt, AccountUrlscannerV2NewScanParamsCountryAz, AccountUrlscannerV2NewScanParamsCountryBh, AccountUrlscannerV2NewScanParamsCountryBd, AccountUrlscannerV2NewScanParamsCountryBb, AccountUrlscannerV2NewScanParamsCountryBy, AccountUrlscannerV2NewScanParamsCountryBe, AccountUrlscannerV2NewScanParamsCountryBz, AccountUrlscannerV2NewScanParamsCountryBj, AccountUrlscannerV2NewScanParamsCountryBm, AccountUrlscannerV2NewScanParamsCountryBt, AccountUrlscannerV2NewScanParamsCountryBo, AccountUrlscannerV2NewScanParamsCountryBa, AccountUrlscannerV2NewScanParamsCountryBw, AccountUrlscannerV2NewScanParamsCountryBr, AccountUrlscannerV2NewScanParamsCountryBn, AccountUrlscannerV2NewScanParamsCountryBg, AccountUrlscannerV2NewScanParamsCountryBf, AccountUrlscannerV2NewScanParamsCountryBi, AccountUrlscannerV2NewScanParamsCountryKh, AccountUrlscannerV2NewScanParamsCountryCm, AccountUrlscannerV2NewScanParamsCountryCa, AccountUrlscannerV2NewScanParamsCountryCv, AccountUrlscannerV2NewScanParamsCountryKy, AccountUrlscannerV2NewScanParamsCountryCf, AccountUrlscannerV2NewScanParamsCountryTd, AccountUrlscannerV2NewScanParamsCountryCl, AccountUrlscannerV2NewScanParamsCountryCn, AccountUrlscannerV2NewScanParamsCountryCo, AccountUrlscannerV2NewScanParamsCountryKm, AccountUrlscannerV2NewScanParamsCountryCg, AccountUrlscannerV2NewScanParamsCountryCr, AccountUrlscannerV2NewScanParamsCountryCi, AccountUrlscannerV2NewScanParamsCountryHr, AccountUrlscannerV2NewScanParamsCountryCu, AccountUrlscannerV2NewScanParamsCountryCy, AccountUrlscannerV2NewScanParamsCountryCz, AccountUrlscannerV2NewScanParamsCountryCd, AccountUrlscannerV2NewScanParamsCountryDk, AccountUrlscannerV2NewScanParamsCountryDj, AccountUrlscannerV2NewScanParamsCountryDm, AccountUrlscannerV2NewScanParamsCountryDo, AccountUrlscannerV2NewScanParamsCountryEc, AccountUrlscannerV2NewScanParamsCountryEg, AccountUrlscannerV2NewScanParamsCountrySv, AccountUrlscannerV2NewScanParamsCountryGq, AccountUrlscannerV2NewScanParamsCountryEr, AccountUrlscannerV2NewScanParamsCountryEe, AccountUrlscannerV2NewScanParamsCountrySz, AccountUrlscannerV2NewScanParamsCountryEt, AccountUrlscannerV2NewScanParamsCountryFj, AccountUrlscannerV2NewScanParamsCountryFi, AccountUrlscannerV2NewScanParamsCountryFr, AccountUrlscannerV2NewScanParamsCountryGa, AccountUrlscannerV2NewScanParamsCountryGe, AccountUrlscannerV2NewScanParamsCountryDe, AccountUrlscannerV2NewScanParamsCountryGh, AccountUrlscannerV2NewScanParamsCountryGr, AccountUrlscannerV2NewScanParamsCountryGl, AccountUrlscannerV2NewScanParamsCountryGd, AccountUrlscannerV2NewScanParamsCountryGt, AccountUrlscannerV2NewScanParamsCountryGn, AccountUrlscannerV2NewScanParamsCountryGw, AccountUrlscannerV2NewScanParamsCountryGy, AccountUrlscannerV2NewScanParamsCountryHt, AccountUrlscannerV2NewScanParamsCountryHn, AccountUrlscannerV2NewScanParamsCountryHu, AccountUrlscannerV2NewScanParamsCountryIs, AccountUrlscannerV2NewScanParamsCountryIn, AccountUrlscannerV2NewScanParamsCountryID, AccountUrlscannerV2NewScanParamsCountryIr, AccountUrlscannerV2NewScanParamsCountryIq, AccountUrlscannerV2NewScanParamsCountryIe, AccountUrlscannerV2NewScanParamsCountryIl, AccountUrlscannerV2NewScanParamsCountryIt, AccountUrlscannerV2NewScanParamsCountryJm, AccountUrlscannerV2NewScanParamsCountryJp, AccountUrlscannerV2NewScanParamsCountryJo, AccountUrlscannerV2NewScanParamsCountryKz, AccountUrlscannerV2NewScanParamsCountryKe, AccountUrlscannerV2NewScanParamsCountryKi, AccountUrlscannerV2NewScanParamsCountryKw, AccountUrlscannerV2NewScanParamsCountryKg, AccountUrlscannerV2NewScanParamsCountryLa, AccountUrlscannerV2NewScanParamsCountryLv, AccountUrlscannerV2NewScanParamsCountryLb, AccountUrlscannerV2NewScanParamsCountryLs, AccountUrlscannerV2NewScanParamsCountryLr, AccountUrlscannerV2NewScanParamsCountryLy, AccountUrlscannerV2NewScanParamsCountryLi, AccountUrlscannerV2NewScanParamsCountryLt, AccountUrlscannerV2NewScanParamsCountryLu, AccountUrlscannerV2NewScanParamsCountryMo, AccountUrlscannerV2NewScanParamsCountryMg, AccountUrlscannerV2NewScanParamsCountryMw, AccountUrlscannerV2NewScanParamsCountryMy, AccountUrlscannerV2NewScanParamsCountryMv, AccountUrlscannerV2NewScanParamsCountryMl, AccountUrlscannerV2NewScanParamsCountryMr, AccountUrlscannerV2NewScanParamsCountryMu, AccountUrlscannerV2NewScanParamsCountryMx, AccountUrlscannerV2NewScanParamsCountryFm, AccountUrlscannerV2NewScanParamsCountryMd, AccountUrlscannerV2NewScanParamsCountryMc, AccountUrlscannerV2NewScanParamsCountryMn, AccountUrlscannerV2NewScanParamsCountryMs, AccountUrlscannerV2NewScanParamsCountryMa, AccountUrlscannerV2NewScanParamsCountryMz, AccountUrlscannerV2NewScanParamsCountryMm, AccountUrlscannerV2NewScanParamsCountryNa, AccountUrlscannerV2NewScanParamsCountryNr, AccountUrlscannerV2NewScanParamsCountryNp, AccountUrlscannerV2NewScanParamsCountryNl, AccountUrlscannerV2NewScanParamsCountryNz, AccountUrlscannerV2NewScanParamsCountryNi, AccountUrlscannerV2NewScanParamsCountryNe, AccountUrlscannerV2NewScanParamsCountryNg, AccountUrlscannerV2NewScanParamsCountryKp, AccountUrlscannerV2NewScanParamsCountryMk, AccountUrlscannerV2NewScanParamsCountryNo, AccountUrlscannerV2NewScanParamsCountryOm, AccountUrlscannerV2NewScanParamsCountryPk, AccountUrlscannerV2NewScanParamsCountryPs, AccountUrlscannerV2NewScanParamsCountryPa, AccountUrlscannerV2NewScanParamsCountryPg, AccountUrlscannerV2NewScanParamsCountryPy, AccountUrlscannerV2NewScanParamsCountryPe, AccountUrlscannerV2NewScanParamsCountryPh, AccountUrlscannerV2NewScanParamsCountryPl, AccountUrlscannerV2NewScanParamsCountryPt, AccountUrlscannerV2NewScanParamsCountryQa, AccountUrlscannerV2NewScanParamsCountryRo, AccountUrlscannerV2NewScanParamsCountryRu, AccountUrlscannerV2NewScanParamsCountryRw, AccountUrlscannerV2NewScanParamsCountrySh, AccountUrlscannerV2NewScanParamsCountryKn, AccountUrlscannerV2NewScanParamsCountryLc, AccountUrlscannerV2NewScanParamsCountryVc, AccountUrlscannerV2NewScanParamsCountryWs, AccountUrlscannerV2NewScanParamsCountrySm, AccountUrlscannerV2NewScanParamsCountrySt, AccountUrlscannerV2NewScanParamsCountrySa, AccountUrlscannerV2NewScanParamsCountrySn, AccountUrlscannerV2NewScanParamsCountryRs, AccountUrlscannerV2NewScanParamsCountrySc, AccountUrlscannerV2NewScanParamsCountrySl, AccountUrlscannerV2NewScanParamsCountrySk, AccountUrlscannerV2NewScanParamsCountrySi, AccountUrlscannerV2NewScanParamsCountrySb, AccountUrlscannerV2NewScanParamsCountrySo, AccountUrlscannerV2NewScanParamsCountryZa, AccountUrlscannerV2NewScanParamsCountryKr, AccountUrlscannerV2NewScanParamsCountrySS, AccountUrlscannerV2NewScanParamsCountryEs, AccountUrlscannerV2NewScanParamsCountryLk, AccountUrlscannerV2NewScanParamsCountrySd, AccountUrlscannerV2NewScanParamsCountrySr, AccountUrlscannerV2NewScanParamsCountrySe, AccountUrlscannerV2NewScanParamsCountryCh, AccountUrlscannerV2NewScanParamsCountrySy, AccountUrlscannerV2NewScanParamsCountryTw, AccountUrlscannerV2NewScanParamsCountryTj, AccountUrlscannerV2NewScanParamsCountryTz, AccountUrlscannerV2NewScanParamsCountryTh, AccountUrlscannerV2NewScanParamsCountryBs, AccountUrlscannerV2NewScanParamsCountryGm, AccountUrlscannerV2NewScanParamsCountryTl, AccountUrlscannerV2NewScanParamsCountryTg, AccountUrlscannerV2NewScanParamsCountryTo, AccountUrlscannerV2NewScanParamsCountryTt, AccountUrlscannerV2NewScanParamsCountryTn, AccountUrlscannerV2NewScanParamsCountryTr, AccountUrlscannerV2NewScanParamsCountryTm, AccountUrlscannerV2NewScanParamsCountryUg, AccountUrlscannerV2NewScanParamsCountryUa, AccountUrlscannerV2NewScanParamsCountryAe, AccountUrlscannerV2NewScanParamsCountryGB, AccountUrlscannerV2NewScanParamsCountryUs, AccountUrlscannerV2NewScanParamsCountryUy, AccountUrlscannerV2NewScanParamsCountryUz, AccountUrlscannerV2NewScanParamsCountryVu, AccountUrlscannerV2NewScanParamsCountryVe, AccountUrlscannerV2NewScanParamsCountryVn, AccountUrlscannerV2NewScanParamsCountryYe, AccountUrlscannerV2NewScanParamsCountryZm, AccountUrlscannerV2NewScanParamsCountryZw:
		return true
	}
	return false
}

// Device resolutions.
type AccountUrlscannerV2NewScanParamsScreenshotsResolution string

const (
	AccountUrlscannerV2NewScanParamsScreenshotsResolutionDesktop AccountUrlscannerV2NewScanParamsScreenshotsResolution = "desktop"
	AccountUrlscannerV2NewScanParamsScreenshotsResolutionMobile  AccountUrlscannerV2NewScanParamsScreenshotsResolution = "mobile"
	AccountUrlscannerV2NewScanParamsScreenshotsResolutionTablet  AccountUrlscannerV2NewScanParamsScreenshotsResolution = "tablet"
)

func (r AccountUrlscannerV2NewScanParamsScreenshotsResolution) IsKnown() bool {
	switch r {
	case AccountUrlscannerV2NewScanParamsScreenshotsResolutionDesktop, AccountUrlscannerV2NewScanParamsScreenshotsResolutionMobile, AccountUrlscannerV2NewScanParamsScreenshotsResolutionTablet:
		return true
	}
	return false
}

// The option `Public` means it will be included in listings like recent scans and
// search results. `Unlisted` means it will not be included in the aforementioned
// listings, users will need to have the scan's ID to access it. A a scan will be
// automatically marked as unlisted if it fails, if it contains potential PII or
// other sensitive material.
type AccountUrlscannerV2NewScanParamsVisibility string

const (
	AccountUrlscannerV2NewScanParamsVisibilityPublic   AccountUrlscannerV2NewScanParamsVisibility = "Public"
	AccountUrlscannerV2NewScanParamsVisibilityUnlisted AccountUrlscannerV2NewScanParamsVisibility = "Unlisted"
)

func (r AccountUrlscannerV2NewScanParamsVisibility) IsKnown() bool {
	switch r {
	case AccountUrlscannerV2NewScanParamsVisibilityPublic, AccountUrlscannerV2NewScanParamsVisibilityUnlisted:
		return true
	}
	return false
}

type AccountUrlscannerV2SearchScansParams struct {
	// Filter scans
	Q param.Field[string] `query:"q"`
	// Limit the number of objects in the response.
	Size param.Field[int64] `query:"size"`
}

// URLQuery serializes [AccountUrlscannerV2SearchScansParams]'s query parameters as
// `url.Values`.
func (r AccountUrlscannerV2SearchScansParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
