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
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/urlscanner/v2/bulk", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Submit a URL to scan. Check limits at
// https://developers.cloudflare.com/security-center/investigate/scan-limits/.
func (r *AccountUrlscannerV2Service) NewScan(ctx context.Context, accountID string, body AccountUrlscannerV2NewScanParams, opts ...option.RequestOption) (res *AccountUrlscannerV2NewScanResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/urlscanner/v2/scan", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Returns a plain text response, with the scan's DOM content as rendered by
// Chrome.
func (r *AccountUrlscannerV2Service) GetDom(ctx context.Context, accountID string, scanID string, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if scanID == "" {
		err = errors.New("missing required scan_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/urlscanner/v2/dom/%s", accountID, scanID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get a URL scan's HAR file. See HAR spec at
// http://www.softwareishard.com/blog/har-12-spec/.
func (r *AccountUrlscannerV2Service) GetHar(ctx context.Context, accountID string, scanID string, opts ...option.RequestOption) (res *AccountUrlscannerV2GetHarResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if scanID == "" {
		err = errors.New("missing required scan_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/urlscanner/v2/har/%s", accountID, scanID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Returns the raw response of the network request. Find the `response_id` in the
// `data.requests.response.hash`.
func (r *AccountUrlscannerV2Service) GetRawResponse(ctx context.Context, accountID string, responseID string, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if responseID == "" {
		err = errors.New("missing required response_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/urlscanner/v2/responses/%s", accountID, responseID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get URL scan by uuid
func (r *AccountUrlscannerV2Service) GetScan(ctx context.Context, accountID string, scanID string, opts ...option.RequestOption) (res *AccountUrlscannerV2GetScanResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if scanID == "" {
		err = errors.New("missing required scan_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/urlscanner/v2/result/%s", accountID, scanID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
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
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/urlscanner/v2/search", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type AccountUrlscannerV2BulkNewScansResponse struct {
	// URL to api report.
	API string `json:"api,required"`
	// URL to report.
	Result string `json:"result,required"`
	// Submitted URL
	URL string `json:"url,required"`
	// Scan ID.
	Uuid string `json:"uuid,required" format:"uuid"`
	// Submitted visibility status.
	Visibility string                                         `json:"visibility,required"`
	Options    AccountUrlscannerV2BulkNewScansResponseOptions `json:"options"`
	JSON       accountUrlscannerV2BulkNewScansResponseJSON    `json:"-"`
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
	API     string `json:"api,required"`
	Message string `json:"message,required"`
	// URL to report.
	Result string `json:"result,required"`
	// Canonical form of submitted URL. Use this if you want to later search by URL.
	URL string `json:"url,required"`
	// Scan ID.
	Uuid string `json:"uuid,required" format:"uuid"`
	// Submitted visibility status.
	Visibility string                                    `json:"visibility,required"`
	Options    AccountUrlscannerV2NewScanResponseOptions `json:"options"`
	JSON       accountUrlscannerV2NewScanResponseJSON    `json:"-"`
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
	Log  AccountUrlscannerV2GetHarResponseLog  `json:"log,required"`
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
	Creator AccountUrlscannerV2GetHarResponseLogCreator `json:"creator,required"`
	Entries []AccountUrlscannerV2GetHarResponseLogEntry `json:"entries,required"`
	Pages   []AccountUrlscannerV2GetHarResponseLogPage  `json:"pages,required"`
	Version string                                      `json:"version,required"`
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
	Comment string                                          `json:"comment,required"`
	Name    string                                          `json:"name,required"`
	Version string                                          `json:"version,required"`
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
	InitialPriority string                                              `json:"_initialPriority,required"`
	InitiatorType   string                                              `json:"_initiator_type,required"`
	Priority        string                                              `json:"_priority,required"`
	RequestID       string                                              `json:"_requestId,required"`
	RequestTime     float64                                             `json:"_requestTime,required"`
	ResourceType    string                                              `json:"_resourceType,required"`
	Cache           interface{}                                         `json:"cache,required"`
	Connection      string                                              `json:"connection,required"`
	Pageref         string                                              `json:"pageref,required"`
	Request         AccountUrlscannerV2GetHarResponseLogEntriesRequest  `json:"request,required"`
	Response        AccountUrlscannerV2GetHarResponseLogEntriesResponse `json:"response,required"`
	ServerIPAddress string                                              `json:"serverIPAddress,required"`
	StartedDateTime string                                              `json:"startedDateTime,required"`
	Time            float64                                             `json:"time,required"`
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
	BodySize    float64                                                    `json:"bodySize,required"`
	Headers     []AccountUrlscannerV2GetHarResponseLogEntriesRequestHeader `json:"headers,required"`
	HeadersSize float64                                                    `json:"headersSize,required"`
	HTTPVersion string                                                     `json:"httpVersion,required"`
	Method      string                                                     `json:"method,required"`
	URL         string                                                     `json:"url,required"`
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
	Name  string                                                       `json:"name,required"`
	Value string                                                       `json:"value,required"`
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
	TransferSize float64                                                     `json:"_transferSize,required"`
	BodySize     float64                                                     `json:"bodySize,required"`
	Content      AccountUrlscannerV2GetHarResponseLogEntriesResponseContent  `json:"content,required"`
	Headers      []AccountUrlscannerV2GetHarResponseLogEntriesResponseHeader `json:"headers,required"`
	HeadersSize  float64                                                     `json:"headersSize,required"`
	HTTPVersion  string                                                      `json:"httpVersion,required"`
	RedirectURL  string                                                      `json:"redirectURL,required"`
	Status       float64                                                     `json:"status,required"`
	StatusText   string                                                      `json:"statusText,required"`
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
	MimeType    string                                                         `json:"mimeType,required"`
	Size        float64                                                        `json:"size,required"`
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
	Name  string                                                        `json:"name,required"`
	Value string                                                        `json:"value,required"`
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
	ID              string                                               `json:"id,required"`
	PageTimings     AccountUrlscannerV2GetHarResponseLogPagesPageTimings `json:"pageTimings,required"`
	StartedDateTime string                                               `json:"startedDateTime,required"`
	Title           string                                               `json:"title,required"`
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
	OnContentLoad float64                                                  `json:"onContentLoad,required"`
	OnLoad        float64                                                  `json:"onLoad,required"`
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
	Data     AccountUrlscannerV2GetScanResponseData     `json:"data,required"`
	Lists    AccountUrlscannerV2GetScanResponseLists    `json:"lists,required"`
	Meta     AccountUrlscannerV2GetScanResponseMeta     `json:"meta,required"`
	Page     AccountUrlscannerV2GetScanResponsePage     `json:"page,required"`
	Scanner  AccountUrlscannerV2GetScanResponseScanner  `json:"scanner,required"`
	Stats    AccountUrlscannerV2GetScanResponseStats    `json:"stats,required"`
	Task     AccountUrlscannerV2GetScanResponseTask     `json:"task,required"`
	Verdicts AccountUrlscannerV2GetScanResponseVerdicts `json:"verdicts,required"`
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
	Console     []AccountUrlscannerV2GetScanResponseDataConsole     `json:"console,required"`
	Cookies     []AccountUrlscannerV2GetScanResponseDataCookie      `json:"cookies,required"`
	Globals     []AccountUrlscannerV2GetScanResponseDataGlobal      `json:"globals,required"`
	Links       []AccountUrlscannerV2GetScanResponseDataLink        `json:"links,required"`
	Performance []AccountUrlscannerV2GetScanResponseDataPerformance `json:"performance,required"`
	Requests    []AccountUrlscannerV2GetScanResponseDataRequest     `json:"requests,required"`
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
	Message AccountUrlscannerV2GetScanResponseDataConsoleMessage `json:"message,required"`
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
	Level  string                                                   `json:"level,required"`
	Source string                                                   `json:"source,required"`
	Text   string                                                   `json:"text,required"`
	URL    string                                                   `json:"url,required"`
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
	Domain       string                                           `json:"domain,required"`
	Expires      float64                                          `json:"expires,required"`
	HTTPOnly     bool                                             `json:"httpOnly,required"`
	Name         string                                           `json:"name,required"`
	Path         string                                           `json:"path,required"`
	Priority     string                                           `json:"priority,required"`
	SameParty    bool                                             `json:"sameParty,required"`
	Secure       bool                                             `json:"secure,required"`
	Session      bool                                             `json:"session,required"`
	Size         float64                                          `json:"size,required"`
	SourcePort   float64                                          `json:"sourcePort,required"`
	SourceScheme string                                           `json:"sourceScheme,required"`
	Value        string                                           `json:"value,required"`
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
	Prop string                                           `json:"prop,required"`
	Type string                                           `json:"type,required"`
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
	Href string                                         `json:"href,required"`
	Text string                                         `json:"text,required"`
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
	Duration  float64                                               `json:"duration,required"`
	EntryType string                                                `json:"entryType,required"`
	Name      string                                                `json:"name,required"`
	StartTime float64                                               `json:"startTime,required"`
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
	Request  AccountUrlscannerV2GetScanResponseDataRequestsRequest   `json:"request,required"`
	Response AccountUrlscannerV2GetScanResponseDataRequestsResponse  `json:"response,required"`
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
	DocumentURL          string                                                                `json:"documentURL,required"`
	HasUserGesture       bool                                                                  `json:"hasUserGesture,required"`
	Initiator            AccountUrlscannerV2GetScanResponseDataRequestsRequestInitiator        `json:"initiator,required"`
	RedirectHasExtraInfo bool                                                                  `json:"redirectHasExtraInfo,required"`
	Request              AccountUrlscannerV2GetScanResponseDataRequestsRequestRequest          `json:"request,required"`
	RequestID            string                                                                `json:"requestId,required"`
	Type                 string                                                                `json:"type,required"`
	WallTime             float64                                                               `json:"wallTime,required"`
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
	Host string                                                             `json:"host,required"`
	Type string                                                             `json:"type,required"`
	URL  string                                                             `json:"url,required"`
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
	InitialPriority  string                                                           `json:"initialPriority,required"`
	IsSameSite       bool                                                             `json:"isSameSite,required"`
	Method           string                                                           `json:"method,required"`
	MixedContentType string                                                           `json:"mixedContentType,required"`
	ReferrerPolicy   string                                                           `json:"referrerPolicy,required"`
	URL              string                                                           `json:"url,required"`
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
	Charset         string                                                                                `json:"charset,required"`
	MimeType        string                                                                                `json:"mimeType,required"`
	Protocol        string                                                                                `json:"protocol,required"`
	RemoteIPAddress string                                                                                `json:"remoteIPAddress,required"`
	RemotePort      float64                                                                               `json:"remotePort,required"`
	SecurityHeaders []AccountUrlscannerV2GetScanResponseDataRequestsRequestRedirectResponseSecurityHeader `json:"securityHeaders,required"`
	SecurityState   string                                                                                `json:"securityState,required"`
	Status          float64                                                                               `json:"status,required"`
	StatusText      string                                                                                `json:"statusText,required"`
	URL             string                                                                                `json:"url,required"`
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
	Name  string                                                                                  `json:"name,required"`
	Value string                                                                                  `json:"value,required"`
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
	Asn               AccountUrlscannerV2GetScanResponseDataRequestsResponseAsn      `json:"asn,required"`
	DataLength        float64                                                        `json:"dataLength,required"`
	EncodedDataLength float64                                                        `json:"encodedDataLength,required"`
	Geoip             AccountUrlscannerV2GetScanResponseDataRequestsResponseGeoip    `json:"geoip,required"`
	HasExtraInfo      bool                                                           `json:"hasExtraInfo,required"`
	RequestID         string                                                         `json:"requestId,required"`
	Response          AccountUrlscannerV2GetScanResponseDataRequestsResponseResponse `json:"response,required"`
	Size              float64                                                        `json:"size,required"`
	Type              string                                                         `json:"type,required"`
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
	Asn         string                                                        `json:"asn,required"`
	Country     string                                                        `json:"country,required"`
	Description string                                                        `json:"description,required"`
	IP          string                                                        `json:"ip,required"`
	Name        string                                                        `json:"name,required"`
	Org         string                                                        `json:"org,required"`
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
	City        string                                                          `json:"city,required"`
	Country     string                                                          `json:"country,required"`
	CountryName string                                                          `json:"country_name,required"`
	GeonameID   string                                                          `json:"geonameId,required"`
	Ll          []interface{}                                                   `json:"ll,required"`
	Region      string                                                          `json:"region,required"`
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
	Charset         string                                                                         `json:"charset,required"`
	MimeType        string                                                                         `json:"mimeType,required"`
	Protocol        string                                                                         `json:"protocol,required"`
	RemoteIPAddress string                                                                         `json:"remoteIPAddress,required"`
	RemotePort      float64                                                                        `json:"remotePort,required"`
	SecurityDetails AccountUrlscannerV2GetScanResponseDataRequestsResponseResponseSecurityDetails  `json:"securityDetails,required"`
	SecurityHeaders []AccountUrlscannerV2GetScanResponseDataRequestsResponseResponseSecurityHeader `json:"securityHeaders,required"`
	SecurityState   string                                                                         `json:"securityState,required"`
	Status          float64                                                                        `json:"status,required"`
	StatusText      string                                                                         `json:"statusText,required"`
	URL             string                                                                         `json:"url,required"`
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
	CertificateID                     float64                                                                           `json:"certificateId,required"`
	CertificateTransparencyCompliance string                                                                            `json:"certificateTransparencyCompliance,required"`
	Cipher                            string                                                                            `json:"cipher,required"`
	EncryptedClientHello              bool                                                                              `json:"encryptedClientHello,required"`
	Issuer                            string                                                                            `json:"issuer,required"`
	KeyExchange                       string                                                                            `json:"keyExchange,required"`
	KeyExchangeGroup                  string                                                                            `json:"keyExchangeGroup,required"`
	Protocol                          string                                                                            `json:"protocol,required"`
	SanList                           []string                                                                          `json:"sanList,required"`
	ServerSignatureAlgorithm          float64                                                                           `json:"serverSignatureAlgorithm,required"`
	SubjectName                       string                                                                            `json:"subjectName,required"`
	ValidFrom                         float64                                                                           `json:"validFrom,required"`
	ValidTo                           float64                                                                           `json:"validTo,required"`
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
	Name  string                                                                           `json:"name,required"`
	Value string                                                                           `json:"value,required"`
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
	Asns         []string                                             `json:"asns,required"`
	Certificates []AccountUrlscannerV2GetScanResponseListsCertificate `json:"certificates,required"`
	Continents   []string                                             `json:"continents,required"`
	Countries    []string                                             `json:"countries,required"`
	Domains      []string                                             `json:"domains,required"`
	Hashes       []string                                             `json:"hashes,required"`
	IPs          []string                                             `json:"ips,required"`
	LinkDomains  []string                                             `json:"linkDomains,required"`
	Servers      []string                                             `json:"servers,required"`
	URLs         []string                                             `json:"urls,required"`
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
	Issuer      string                                                 `json:"issuer,required"`
	SubjectName string                                                 `json:"subjectName,required"`
	ValidFrom   float64                                                `json:"validFrom,required"`
	ValidTo     float64                                                `json:"validTo,required"`
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
	Processors AccountUrlscannerV2GetScanResponseMetaProcessors `json:"processors,required"`
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
	Asn              AccountUrlscannerV2GetScanResponseMetaProcessorsAsn              `json:"asn,required"`
	DNS              AccountUrlscannerV2GetScanResponseMetaProcessorsDNS              `json:"dns,required"`
	DomainCategories AccountUrlscannerV2GetScanResponseMetaProcessorsDomainCategories `json:"domainCategories,required"`
	Geoip            AccountUrlscannerV2GetScanResponseMetaProcessorsGeoip            `json:"geoip,required"`
	Phishing         AccountUrlscannerV2GetScanResponseMetaProcessorsPhishing         `json:"phishing,required"`
	RadarRank        AccountUrlscannerV2GetScanResponseMetaProcessorsRadarRank        `json:"radarRank,required"`
	Wappa            AccountUrlscannerV2GetScanResponseMetaProcessorsWappa            `json:"wappa,required"`
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
	Data []AccountUrlscannerV2GetScanResponseMetaProcessorsAsnData `json:"data,required"`
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
	Asn         string                                                      `json:"asn,required"`
	Country     string                                                      `json:"country,required"`
	Description string                                                      `json:"description,required"`
	IP          string                                                      `json:"ip,required"`
	Name        string                                                      `json:"name,required"`
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
	Data []AccountUrlscannerV2GetScanResponseMetaProcessorsDNSData `json:"data,required"`
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
	Address     string                                                      `json:"address,required"`
	DnssecValid bool                                                        `json:"dnssec_valid,required"`
	Name        string                                                      `json:"name,required"`
	Type        string                                                      `json:"type,required"`
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
	Data []AccountUrlscannerV2GetScanResponseMetaProcessorsDomainCategoriesData `json:"data,required"`
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
	Inherited interface{}                                                              `json:"inherited,required"`
	IsPrimary bool                                                                     `json:"isPrimary,required"`
	Name      string                                                                   `json:"name,required"`
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
	Data []AccountUrlscannerV2GetScanResponseMetaProcessorsGeoipData `json:"data,required"`
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
	Geoip AccountUrlscannerV2GetScanResponseMetaProcessorsGeoipDataGeoip `json:"geoip,required"`
	IP    string                                                         `json:"ip,required"`
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
	City        string                                                             `json:"city,required"`
	Country     string                                                             `json:"country,required"`
	CountryName string                                                             `json:"country_name,required"`
	Ll          []float64                                                          `json:"ll,required"`
	Region      string                                                             `json:"region,required"`
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
	Data []string                                                     `json:"data,required"`
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
	Data []AccountUrlscannerV2GetScanResponseMetaProcessorsRadarRankData `json:"data,required"`
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
	Bucket   string                                                            `json:"bucket,required"`
	Hostname string                                                            `json:"hostname,required"`
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
	Data []AccountUrlscannerV2GetScanResponseMetaProcessorsWappaData `json:"data,required"`
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
	App             string                                                                `json:"app,required"`
	Categories      []AccountUrlscannerV2GetScanResponseMetaProcessorsWappaDataCategory   `json:"categories,required"`
	Confidence      []AccountUrlscannerV2GetScanResponseMetaProcessorsWappaDataConfidence `json:"confidence,required"`
	ConfidenceTotal float64                                                               `json:"confidenceTotal,required"`
	Icon            string                                                                `json:"icon,required"`
	Website         string                                                                `json:"website,required"`
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
	Name     string                                                                `json:"name,required"`
	Priority float64                                                               `json:"priority,required"`
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
	Confidence  float64                                                                 `json:"confidence,required"`
	Name        string                                                                  `json:"name,required"`
	Pattern     string                                                                  `json:"pattern,required"`
	PatternType string                                                                  `json:"patternType,required"`
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
	Data []AccountUrlscannerV2GetScanResponseMetaProcessorsURLCategoriesData `json:"data,required"`
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
	Content   []AccountUrlscannerV2GetScanResponseMetaProcessorsURLCategoriesDataContent `json:"content,required"`
	Inherited AccountUrlscannerV2GetScanResponseMetaProcessorsURLCategoriesDataInherited `json:"inherited,required"`
	Name      string                                                                     `json:"name,required"`
	Risks     []AccountUrlscannerV2GetScanResponseMetaProcessorsURLCategoriesDataRisk    `json:"risks,required"`
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
	ID              float64                                                                      `json:"id,required"`
	Name            string                                                                       `json:"name,required"`
	SuperCategoryID float64                                                                      `json:"super_category_id,required"`
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
	Content []AccountUrlscannerV2GetScanResponseMetaProcessorsURLCategoriesDataInheritedContent `json:"content,required"`
	From    string                                                                              `json:"from,required"`
	Risks   []AccountUrlscannerV2GetScanResponseMetaProcessorsURLCategoriesDataInheritedRisk    `json:"risks,required"`
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
	ID              float64                                                                               `json:"id,required"`
	Name            string                                                                                `json:"name,required"`
	SuperCategoryID float64                                                                               `json:"super_category_id,required"`
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
	ID              float64                                                                            `json:"id,required"`
	Name            string                                                                             `json:"name,required"`
	SuperCategoryID float64                                                                            `json:"super_category_id,required"`
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
	ID              float64                                                                   `json:"id,required"`
	Name            string                                                                    `json:"name,required"`
	SuperCategoryID float64                                                                   `json:"super_category_id,required"`
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
	ApexDomain   string                                           `json:"apexDomain,required"`
	Asn          string                                           `json:"asn,required"`
	Asnname      string                                           `json:"asnname,required"`
	City         string                                           `json:"city,required"`
	Country      string                                           `json:"country,required"`
	Domain       string                                           `json:"domain,required"`
	IP           string                                           `json:"ip,required"`
	MimeType     string                                           `json:"mimeType,required"`
	Server       string                                           `json:"server,required"`
	Status       string                                           `json:"status,required"`
	Title        string                                           `json:"title,required"`
	TlsAgeDays   float64                                          `json:"tlsAgeDays,required"`
	TlsIssuer    string                                           `json:"tlsIssuer,required"`
	TlsValidDays float64                                          `json:"tlsValidDays,required"`
	TlsValidFrom string                                           `json:"tlsValidFrom,required"`
	URL          string                                           `json:"url,required"`
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
	Dhash   string                                               `json:"dhash,required"`
	Mm3Hash float64                                              `json:"mm3Hash,required"`
	Name    string                                               `json:"name,required"`
	Phash   string                                               `json:"phash,required"`
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
	Colo    string                                        `json:"colo,required"`
	Country string                                        `json:"country,required"`
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
	DomainStats      []AccountUrlscannerV2GetScanResponseStatsDomainStat   `json:"domainStats,required"`
	IPStats          []AccountUrlscannerV2GetScanResponseStatsIPStat       `json:"ipStats,required"`
	IPv6Percentage   float64                                               `json:"IPv6Percentage,required"`
	Malicious        float64                                               `json:"malicious,required"`
	ProtocolStats    []AccountUrlscannerV2GetScanResponseStatsProtocolStat `json:"protocolStats,required"`
	ResourceStats    []AccountUrlscannerV2GetScanResponseStatsResourceStat `json:"resourceStats,required"`
	SecurePercentage float64                                               `json:"securePercentage,required"`
	SecureRequests   float64                                               `json:"secureRequests,required"`
	ServerStats      []AccountUrlscannerV2GetScanResponseStatsServerStat   `json:"serverStats,required"`
	TlsStats         []AccountUrlscannerV2GetScanResponseStatsTlsStat      `json:"tlsStats,required"`
	TotalLinks       float64                                               `json:"totalLinks,required"`
	UniqAsNs         float64                                               `json:"uniqASNs,required"`
	UniqCountries    float64                                               `json:"uniqCountries,required"`
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
	Count       float64                                               `json:"count,required"`
	Countries   []string                                              `json:"countries,required"`
	Domain      string                                                `json:"domain,required"`
	EncodedSize float64                                               `json:"encodedSize,required"`
	Index       float64                                               `json:"index,required"`
	Initiators  []string                                              `json:"initiators,required"`
	IPs         []string                                              `json:"ips,required"`
	Redirects   float64                                               `json:"redirects,required"`
	Size        float64                                               `json:"size,required"`
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
	Asn         AccountUrlscannerV2GetScanResponseStatsIPStatsAsn   `json:"asn,required"`
	Countries   []string                                            `json:"countries,required"`
	Domains     []string                                            `json:"domains,required"`
	EncodedSize float64                                             `json:"encodedSize,required"`
	Geoip       AccountUrlscannerV2GetScanResponseStatsIPStatsGeoip `json:"geoip,required"`
	Index       float64                                             `json:"index,required"`
	IP          string                                              `json:"ip,required"`
	Ipv6        bool                                                `json:"ipv6,required"`
	Redirects   float64                                             `json:"redirects,required"`
	Requests    float64                                             `json:"requests,required"`
	Size        float64                                             `json:"size,required"`
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
	Asn         string                                                `json:"asn,required"`
	Country     string                                                `json:"country,required"`
	Description string                                                `json:"description,required"`
	IP          string                                                `json:"ip,required"`
	Name        string                                                `json:"name,required"`
	Org         string                                                `json:"org,required"`
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
	City        string                                                  `json:"city,required"`
	Country     string                                                  `json:"country,required"`
	CountryName string                                                  `json:"country_name,required"`
	Ll          []float64                                               `json:"ll,required"`
	Region      string                                                  `json:"region,required"`
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
	Count       float64                                                 `json:"count,required"`
	Countries   []string                                                `json:"countries,required"`
	EncodedSize float64                                                 `json:"encodedSize,required"`
	IPs         []string                                                `json:"ips,required"`
	Protocol    string                                                  `json:"protocol,required"`
	Size        float64                                                 `json:"size,required"`
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
	Compression float64                                                 `json:"compression,required"`
	Count       float64                                                 `json:"count,required"`
	Countries   []string                                                `json:"countries,required"`
	EncodedSize float64                                                 `json:"encodedSize,required"`
	IPs         []string                                                `json:"ips,required"`
	Percentage  float64                                                 `json:"percentage,required"`
	Size        float64                                                 `json:"size,required"`
	Type        string                                                  `json:"type,required"`
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
	Count       float64                                               `json:"count,required"`
	Countries   []string                                              `json:"countries,required"`
	EncodedSize float64                                               `json:"encodedSize,required"`
	IPs         []string                                              `json:"ips,required"`
	Server      string                                                `json:"server,required"`
	Size        float64                                               `json:"size,required"`
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
	Count         float64                                                  `json:"count,required"`
	Countries     []string                                                 `json:"countries,required"`
	EncodedSize   float64                                                  `json:"encodedSize,required"`
	IPs           []string                                                 `json:"ips,required"`
	Protocols     AccountUrlscannerV2GetScanResponseStatsTlsStatsProtocols `json:"protocols,required"`
	SecurityState string                                                   `json:"securityState,required"`
	Size          float64                                                  `json:"size,required"`
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
	Tls1_3Aes128Gcm float64                                                      `json:"TLS 1.3 / AES_128_GCM,required"`
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
	ApexDomain    string                                        `json:"apexDomain,required"`
	Domain        string                                        `json:"domain,required"`
	DomURL        string                                        `json:"domURL,required"`
	Method        string                                        `json:"method,required"`
	Options       AccountUrlscannerV2GetScanResponseTaskOptions `json:"options,required"`
	ReportURL     string                                        `json:"reportURL,required"`
	ScreenshotURL string                                        `json:"screenshotURL,required"`
	Source        string                                        `json:"source,required"`
	Success       bool                                          `json:"success,required"`
	Time          string                                        `json:"time,required"`
	URL           string                                        `json:"url,required"`
	Uuid          string                                        `json:"uuid,required"`
	Visibility    string                                        `json:"visibility,required"`
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
	Overall AccountUrlscannerV2GetScanResponseVerdictsOverall `json:"overall,required"`
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
	Categories  []string                                              `json:"categories,required"`
	HasVerdicts bool                                                  `json:"hasVerdicts,required"`
	Malicious   bool                                                  `json:"malicious,required"`
	Tags        []string                                              `json:"tags,required"`
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
	Results []AccountUrlscannerV2SearchScansResponseResult `json:"results,required"`
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
	ID       string                                                `json:"_id,required"`
	Page     AccountUrlscannerV2SearchScansResponseResultsPage     `json:"page,required"`
	Result   string                                                `json:"result,required"`
	Stats    AccountUrlscannerV2SearchScansResponseResultsStats    `json:"stats,required"`
	Task     AccountUrlscannerV2SearchScansResponseResultsTask     `json:"task,required"`
	Verdicts AccountUrlscannerV2SearchScansResponseResultsVerdicts `json:"verdicts,required"`
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
	Asn     string                                                `json:"asn,required"`
	Country string                                                `json:"country,required"`
	IP      string                                                `json:"ip,required"`
	URL     string                                                `json:"url,required"`
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
	DataLength    float64                                                `json:"dataLength,required"`
	Requests      float64                                                `json:"requests,required"`
	UniqCountries float64                                                `json:"uniqCountries,required"`
	UniqIPs       float64                                                `json:"uniqIPs,required"`
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
	Time       string                                                `json:"time,required"`
	URL        string                                                `json:"url,required"`
	Uuid       string                                                `json:"uuid,required"`
	Visibility string                                                `json:"visibility,required"`
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
	Malicious bool                                                      `json:"malicious,required"`
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
	URL         param.Field[string] `json:"url,required"`
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
	URL         param.Field[string] `json:"url,required"`
	Customagent param.Field[string] `json:"customagent"`
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
