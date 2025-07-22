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

// AccountBrowserRenderingService contains methods and other services that help
// with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountBrowserRenderingService] method instead.
type AccountBrowserRenderingService struct {
	Options []option.RequestOption
}

// NewAccountBrowserRenderingService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountBrowserRenderingService(opts ...option.RequestOption) (r *AccountBrowserRenderingService) {
	r = &AccountBrowserRenderingService{}
	r.Options = opts
	return
}

// Fetches rendered HTML content from provided URL or HTML. Check available options
// like `gotoOptions` and `waitFor*` to control page load behaviour.
func (r *AccountBrowserRenderingService) GetHTMLContent(ctx context.Context, accountID string, params AccountBrowserRenderingGetHTMLContentParams, opts ...option.RequestOption) (res *AccountBrowserRenderingGetHTMLContentResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/browser-rendering/content", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Gets json from a webpage from a provided URL or HTML. Pass `prompt` or `schema`
// in the body. Control page loading with `gotoOptions` and `waitFor*` options.
func (r *AccountBrowserRenderingService) GetJson(ctx context.Context, accountID string, params AccountBrowserRenderingGetJsonParams, opts ...option.RequestOption) (res *AccountBrowserRenderingGetJsonResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/browser-rendering/json", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Get links from a web page.
func (r *AccountBrowserRenderingService) GetLinks(ctx context.Context, accountID string, params AccountBrowserRenderingGetLinksParams, opts ...option.RequestOption) (res *AccountBrowserRenderingGetLinksResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/browser-rendering/links", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Gets markdown of a webpage from provided URL or HTML. Control page loading with
// `gotoOptions` and `waitFor*` options.
func (r *AccountBrowserRenderingService) GetMarkdown(ctx context.Context, accountID string, params AccountBrowserRenderingGetMarkdownParams, opts ...option.RequestOption) (res *AccountBrowserRenderingGetMarkdownResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/browser-rendering/markdown", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Fetches rendered PDF from provided URL or HTML. Check available options like
// `gotoOptions` and `waitFor*` to control page load behaviour.
func (r *AccountBrowserRenderingService) GetPdf(ctx context.Context, accountID string, params AccountBrowserRenderingGetPdfParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/pdf")}, opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/browser-rendering/pdf", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Takes a screenshot of a webpage from provided URL or HTML. Control page loading
// with `gotoOptions` and `waitFor*` options. Customize screenshots with
// `viewport`, `fullPage`, `clip` and others.
func (r *AccountBrowserRenderingService) GetScreenshot(ctx context.Context, accountID string, params AccountBrowserRenderingGetScreenshotParams, opts ...option.RequestOption) (res *AccountBrowserRenderingGetScreenshotResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/browser-rendering/screenshot", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Returns the page's HTML content and screenshot. Control page loading with
// `gotoOptions` and `waitFor*` options. Customize screenshots with `viewport`,
// `fullPage`, `clip` and others.
func (r *AccountBrowserRenderingService) GetSnapshot(ctx context.Context, accountID string, params AccountBrowserRenderingGetSnapshotParams, opts ...option.RequestOption) (res *AccountBrowserRenderingGetSnapshotResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/browser-rendering/snapshot", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Get meta attributes like height, width, text and others of selected elements.
func (r *AccountBrowserRenderingService) ScrapeElements(ctx context.Context, accountID string, params AccountBrowserRenderingScrapeElementsParams, opts ...option.RequestOption) (res *AccountBrowserRenderingScrapeElementsResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/browser-rendering/scrape", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type AccountBrowserRenderingGetHTMLContentResponse struct {
	// Response status
	Status bool                                                 `json:"status,required"`
	Errors []AccountBrowserRenderingGetHTMLContentResponseError `json:"errors"`
	// HTML content
	Result string                                            `json:"result"`
	JSON   accountBrowserRenderingGetHTMLContentResponseJSON `json:"-"`
}

// accountBrowserRenderingGetHTMLContentResponseJSON contains the JSON metadata for
// the struct [AccountBrowserRenderingGetHTMLContentResponse]
type accountBrowserRenderingGetHTMLContentResponseJSON struct {
	Status      apijson.Field
	Errors      apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountBrowserRenderingGetHTMLContentResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountBrowserRenderingGetHTMLContentResponseJSON) RawJSON() string {
	return r.raw
}

type AccountBrowserRenderingGetHTMLContentResponseError struct {
	// Error code
	Code float64 `json:"code,required"`
	// Error Message
	Message string                                                 `json:"message,required"`
	JSON    accountBrowserRenderingGetHTMLContentResponseErrorJSON `json:"-"`
}

// accountBrowserRenderingGetHTMLContentResponseErrorJSON contains the JSON
// metadata for the struct [AccountBrowserRenderingGetHTMLContentResponseError]
type accountBrowserRenderingGetHTMLContentResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountBrowserRenderingGetHTMLContentResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountBrowserRenderingGetHTMLContentResponseErrorJSON) RawJSON() string {
	return r.raw
}

type AccountBrowserRenderingGetJsonResponse struct {
	Result map[string]interface{} `json:"result,required"`
	// Response status
	Status bool                                          `json:"status,required"`
	Errors []AccountBrowserRenderingGetJsonResponseError `json:"errors"`
	JSON   accountBrowserRenderingGetJsonResponseJSON    `json:"-"`
}

// accountBrowserRenderingGetJsonResponseJSON contains the JSON metadata for the
// struct [AccountBrowserRenderingGetJsonResponse]
type accountBrowserRenderingGetJsonResponseJSON struct {
	Result      apijson.Field
	Status      apijson.Field
	Errors      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountBrowserRenderingGetJsonResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountBrowserRenderingGetJsonResponseJSON) RawJSON() string {
	return r.raw
}

type AccountBrowserRenderingGetJsonResponseError struct {
	// Error code
	Code float64 `json:"code,required"`
	// Error Message
	Message string                                          `json:"message,required"`
	JSON    accountBrowserRenderingGetJsonResponseErrorJSON `json:"-"`
}

// accountBrowserRenderingGetJsonResponseErrorJSON contains the JSON metadata for
// the struct [AccountBrowserRenderingGetJsonResponseError]
type accountBrowserRenderingGetJsonResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountBrowserRenderingGetJsonResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountBrowserRenderingGetJsonResponseErrorJSON) RawJSON() string {
	return r.raw
}

type AccountBrowserRenderingGetLinksResponse struct {
	Result []string `json:"result,required"`
	// Response status
	Status bool                                           `json:"status,required"`
	Errors []AccountBrowserRenderingGetLinksResponseError `json:"errors"`
	JSON   accountBrowserRenderingGetLinksResponseJSON    `json:"-"`
}

// accountBrowserRenderingGetLinksResponseJSON contains the JSON metadata for the
// struct [AccountBrowserRenderingGetLinksResponse]
type accountBrowserRenderingGetLinksResponseJSON struct {
	Result      apijson.Field
	Status      apijson.Field
	Errors      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountBrowserRenderingGetLinksResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountBrowserRenderingGetLinksResponseJSON) RawJSON() string {
	return r.raw
}

type AccountBrowserRenderingGetLinksResponseError struct {
	// Error code
	Code float64 `json:"code,required"`
	// Error Message
	Message string                                           `json:"message,required"`
	JSON    accountBrowserRenderingGetLinksResponseErrorJSON `json:"-"`
}

// accountBrowserRenderingGetLinksResponseErrorJSON contains the JSON metadata for
// the struct [AccountBrowserRenderingGetLinksResponseError]
type accountBrowserRenderingGetLinksResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountBrowserRenderingGetLinksResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountBrowserRenderingGetLinksResponseErrorJSON) RawJSON() string {
	return r.raw
}

type AccountBrowserRenderingGetMarkdownResponse struct {
	// Response status
	Status bool                                              `json:"status,required"`
	Errors []AccountBrowserRenderingGetMarkdownResponseError `json:"errors"`
	// Markdown
	Result string                                         `json:"result"`
	JSON   accountBrowserRenderingGetMarkdownResponseJSON `json:"-"`
}

// accountBrowserRenderingGetMarkdownResponseJSON contains the JSON metadata for
// the struct [AccountBrowserRenderingGetMarkdownResponse]
type accountBrowserRenderingGetMarkdownResponseJSON struct {
	Status      apijson.Field
	Errors      apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountBrowserRenderingGetMarkdownResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountBrowserRenderingGetMarkdownResponseJSON) RawJSON() string {
	return r.raw
}

type AccountBrowserRenderingGetMarkdownResponseError struct {
	// Error code
	Code float64 `json:"code,required"`
	// Error Message
	Message string                                              `json:"message,required"`
	JSON    accountBrowserRenderingGetMarkdownResponseErrorJSON `json:"-"`
}

// accountBrowserRenderingGetMarkdownResponseErrorJSON contains the JSON metadata
// for the struct [AccountBrowserRenderingGetMarkdownResponseError]
type accountBrowserRenderingGetMarkdownResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountBrowserRenderingGetMarkdownResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountBrowserRenderingGetMarkdownResponseErrorJSON) RawJSON() string {
	return r.raw
}

type AccountBrowserRenderingGetScreenshotResponse struct {
	// Response status
	Status bool                                                `json:"status,required"`
	Errors []AccountBrowserRenderingGetScreenshotResponseError `json:"errors"`
	JSON   accountBrowserRenderingGetScreenshotResponseJSON    `json:"-"`
}

// accountBrowserRenderingGetScreenshotResponseJSON contains the JSON metadata for
// the struct [AccountBrowserRenderingGetScreenshotResponse]
type accountBrowserRenderingGetScreenshotResponseJSON struct {
	Status      apijson.Field
	Errors      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountBrowserRenderingGetScreenshotResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountBrowserRenderingGetScreenshotResponseJSON) RawJSON() string {
	return r.raw
}

type AccountBrowserRenderingGetScreenshotResponseError struct {
	// Error code
	Code float64 `json:"code,required"`
	// Error Message
	Message string                                                `json:"message,required"`
	JSON    accountBrowserRenderingGetScreenshotResponseErrorJSON `json:"-"`
}

// accountBrowserRenderingGetScreenshotResponseErrorJSON contains the JSON metadata
// for the struct [AccountBrowserRenderingGetScreenshotResponseError]
type accountBrowserRenderingGetScreenshotResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountBrowserRenderingGetScreenshotResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountBrowserRenderingGetScreenshotResponseErrorJSON) RawJSON() string {
	return r.raw
}

type AccountBrowserRenderingGetSnapshotResponse struct {
	// Response status
	Status bool                                              `json:"status,required"`
	Errors []AccountBrowserRenderingGetSnapshotResponseError `json:"errors"`
	Result AccountBrowserRenderingGetSnapshotResponseResult  `json:"result"`
	JSON   accountBrowserRenderingGetSnapshotResponseJSON    `json:"-"`
}

// accountBrowserRenderingGetSnapshotResponseJSON contains the JSON metadata for
// the struct [AccountBrowserRenderingGetSnapshotResponse]
type accountBrowserRenderingGetSnapshotResponseJSON struct {
	Status      apijson.Field
	Errors      apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountBrowserRenderingGetSnapshotResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountBrowserRenderingGetSnapshotResponseJSON) RawJSON() string {
	return r.raw
}

type AccountBrowserRenderingGetSnapshotResponseError struct {
	// Error code
	Code float64 `json:"code,required"`
	// Error Message
	Message string                                              `json:"message,required"`
	JSON    accountBrowserRenderingGetSnapshotResponseErrorJSON `json:"-"`
}

// accountBrowserRenderingGetSnapshotResponseErrorJSON contains the JSON metadata
// for the struct [AccountBrowserRenderingGetSnapshotResponseError]
type accountBrowserRenderingGetSnapshotResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountBrowserRenderingGetSnapshotResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountBrowserRenderingGetSnapshotResponseErrorJSON) RawJSON() string {
	return r.raw
}

type AccountBrowserRenderingGetSnapshotResponseResult struct {
	// HTML content
	Content string `json:"content,required"`
	// Base64 encoded image
	Screenshot string                                               `json:"screenshot,required"`
	JSON       accountBrowserRenderingGetSnapshotResponseResultJSON `json:"-"`
}

// accountBrowserRenderingGetSnapshotResponseResultJSON contains the JSON metadata
// for the struct [AccountBrowserRenderingGetSnapshotResponseResult]
type accountBrowserRenderingGetSnapshotResponseResultJSON struct {
	Content     apijson.Field
	Screenshot  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountBrowserRenderingGetSnapshotResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountBrowserRenderingGetSnapshotResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountBrowserRenderingScrapeElementsResponse struct {
	Result []AccountBrowserRenderingScrapeElementsResponseResult `json:"result,required"`
	// Response status
	Status bool                                                 `json:"status,required"`
	Errors []AccountBrowserRenderingScrapeElementsResponseError `json:"errors"`
	JSON   accountBrowserRenderingScrapeElementsResponseJSON    `json:"-"`
}

// accountBrowserRenderingScrapeElementsResponseJSON contains the JSON metadata for
// the struct [AccountBrowserRenderingScrapeElementsResponse]
type accountBrowserRenderingScrapeElementsResponseJSON struct {
	Result      apijson.Field
	Status      apijson.Field
	Errors      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountBrowserRenderingScrapeElementsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountBrowserRenderingScrapeElementsResponseJSON) RawJSON() string {
	return r.raw
}

type AccountBrowserRenderingScrapeElementsResponseResult struct {
	Result AccountBrowserRenderingScrapeElementsResponseResultResult `json:"result,required"`
	// Selector
	Selector string                                                  `json:"selector,required"`
	JSON     accountBrowserRenderingScrapeElementsResponseResultJSON `json:"-"`
}

// accountBrowserRenderingScrapeElementsResponseResultJSON contains the JSON
// metadata for the struct [AccountBrowserRenderingScrapeElementsResponseResult]
type accountBrowserRenderingScrapeElementsResponseResultJSON struct {
	Result      apijson.Field
	Selector    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountBrowserRenderingScrapeElementsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountBrowserRenderingScrapeElementsResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountBrowserRenderingScrapeElementsResponseResultResult struct {
	Attributes []AccountBrowserRenderingScrapeElementsResponseResultResultAttribute `json:"attributes,required"`
	// Element height
	Height float64 `json:"height,required"`
	// Html content
	HTML string `json:"html,required"`
	// Element left
	Left float64 `json:"left,required"`
	// Text content
	Text string `json:"text,required"`
	// Element top
	Top float64 `json:"top,required"`
	// Element width
	Width float64                                                       `json:"width,required"`
	JSON  accountBrowserRenderingScrapeElementsResponseResultResultJSON `json:"-"`
}

// accountBrowserRenderingScrapeElementsResponseResultResultJSON contains the JSON
// metadata for the struct
// [AccountBrowserRenderingScrapeElementsResponseResultResult]
type accountBrowserRenderingScrapeElementsResponseResultResultJSON struct {
	Attributes  apijson.Field
	Height      apijson.Field
	HTML        apijson.Field
	Left        apijson.Field
	Text        apijson.Field
	Top         apijson.Field
	Width       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountBrowserRenderingScrapeElementsResponseResultResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountBrowserRenderingScrapeElementsResponseResultResultJSON) RawJSON() string {
	return r.raw
}

type AccountBrowserRenderingScrapeElementsResponseResultResultAttribute struct {
	// Attribute name
	Name string `json:"name,required"`
	// Attribute value
	Value string                                                                 `json:"value,required"`
	JSON  accountBrowserRenderingScrapeElementsResponseResultResultAttributeJSON `json:"-"`
}

// accountBrowserRenderingScrapeElementsResponseResultResultAttributeJSON contains
// the JSON metadata for the struct
// [AccountBrowserRenderingScrapeElementsResponseResultResultAttribute]
type accountBrowserRenderingScrapeElementsResponseResultResultAttributeJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountBrowserRenderingScrapeElementsResponseResultResultAttribute) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountBrowserRenderingScrapeElementsResponseResultResultAttributeJSON) RawJSON() string {
	return r.raw
}

type AccountBrowserRenderingScrapeElementsResponseError struct {
	// Error code
	Code float64 `json:"code,required"`
	// Error Message
	Message string                                                 `json:"message,required"`
	JSON    accountBrowserRenderingScrapeElementsResponseErrorJSON `json:"-"`
}

// accountBrowserRenderingScrapeElementsResponseErrorJSON contains the JSON
// metadata for the struct [AccountBrowserRenderingScrapeElementsResponseError]
type accountBrowserRenderingScrapeElementsResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountBrowserRenderingScrapeElementsResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountBrowserRenderingScrapeElementsResponseErrorJSON) RawJSON() string {
	return r.raw
}

type AccountBrowserRenderingGetHTMLContentParams struct {
	// Cache TTL default is 5s. Set to 0 to disable.
	CacheTtl param.Field[float64] `query:"cacheTTL"`
	// Adds a `<script>` tag into the page with the desired URL or content.
	AddScriptTag param.Field[[]AccountBrowserRenderingGetHTMLContentParamsAddScriptTag] `json:"addScriptTag"`
	// Adds a `<link rel="stylesheet">` tag into the page with the desired URL or a
	// `<style type="text/css">` tag with the content.
	AddStyleTag param.Field[[]AccountBrowserRenderingGetHTMLContentParamsAddStyleTag] `json:"addStyleTag"`
	// Only allow requests that match the provided regex patterns, eg. '/^.\*\.(css)'.
	AllowRequestPattern param.Field[[]string] `json:"allowRequestPattern"`
	// Only allow requests that match the provided resource types, eg. 'image' or
	// 'script'.
	AllowResourceTypes param.Field[[]AccountBrowserRenderingGetHTMLContentParamsAllowResourceType] `json:"allowResourceTypes"`
	// Provide credentials for HTTP authentication.
	Authenticate param.Field[AccountBrowserRenderingGetHTMLContentParamsAuthenticate] `json:"authenticate"`
	// Attempt to proceed when 'awaited' events fail or timeout.
	BestAttempt param.Field[bool] `json:"bestAttempt"`
	// Check [options](https://pptr.dev/api/puppeteer.page.setcookie).
	Cookies          param.Field[[]AccountBrowserRenderingGetHTMLContentParamsCookie] `json:"cookies"`
	EmulateMediaType param.Field[string]                                              `json:"emulateMediaType"`
	// Check [options](https://pptr.dev/api/puppeteer.gotooptions).
	GotoOptions param.Field[AccountBrowserRenderingGetHTMLContentParamsGotoOptions] `json:"gotoOptions"`
	// Set the content of the page, eg: `<h1>Hello World!!</h1>`. Either `html` or
	// `url` must be set.
	HTML param.Field[string] `json:"html"`
	// Block undesired requests that match the provided regex patterns, eg.
	// '/^.\*\.(css)'.
	RejectRequestPattern param.Field[[]string] `json:"rejectRequestPattern"`
	// Block undesired requests that match the provided resource types, eg. 'image' or
	// 'script'.
	RejectResourceTypes  param.Field[[]AccountBrowserRenderingGetHTMLContentParamsRejectResourceType] `json:"rejectResourceTypes"`
	SetExtraHTTPHeaders  param.Field[map[string]string]                                               `json:"setExtraHTTPHeaders"`
	SetJavaScriptEnabled param.Field[bool]                                                            `json:"setJavaScriptEnabled"`
	// URL to navigate to, eg. `https://example.com`.
	URL       param.Field[string] `json:"url" format:"uri"`
	UserAgent param.Field[string] `json:"userAgent"`
	// Check [options](https://pptr.dev/api/puppeteer.page.setviewport).
	Viewport param.Field[AccountBrowserRenderingGetHTMLContentParamsViewport] `json:"viewport"`
	// Wait for the selector to appear in page. Check
	// [options](https://pptr.dev/api/puppeteer.page.waitforselector).
	WaitForSelector param.Field[AccountBrowserRenderingGetHTMLContentParamsWaitForSelector] `json:"waitForSelector"`
	// Waits for a specified timeout before continuing.
	WaitForTimeout param.Field[float64] `json:"waitForTimeout"`
}

func (r AccountBrowserRenderingGetHTMLContentParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [AccountBrowserRenderingGetHTMLContentParams]'s query
// parameters as `url.Values`.
func (r AccountBrowserRenderingGetHTMLContentParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountBrowserRenderingGetHTMLContentParamsAddScriptTag struct {
	ID      param.Field[string] `json:"id"`
	Content param.Field[string] `json:"content"`
	Type    param.Field[string] `json:"type"`
	URL     param.Field[string] `json:"url"`
}

func (r AccountBrowserRenderingGetHTMLContentParamsAddScriptTag) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountBrowserRenderingGetHTMLContentParamsAddStyleTag struct {
	Content param.Field[string] `json:"content"`
	URL     param.Field[string] `json:"url"`
}

func (r AccountBrowserRenderingGetHTMLContentParamsAddStyleTag) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountBrowserRenderingGetHTMLContentParamsAllowResourceType string

const (
	AccountBrowserRenderingGetHTMLContentParamsAllowResourceTypeDocument           AccountBrowserRenderingGetHTMLContentParamsAllowResourceType = "document"
	AccountBrowserRenderingGetHTMLContentParamsAllowResourceTypeStylesheet         AccountBrowserRenderingGetHTMLContentParamsAllowResourceType = "stylesheet"
	AccountBrowserRenderingGetHTMLContentParamsAllowResourceTypeImage              AccountBrowserRenderingGetHTMLContentParamsAllowResourceType = "image"
	AccountBrowserRenderingGetHTMLContentParamsAllowResourceTypeMedia              AccountBrowserRenderingGetHTMLContentParamsAllowResourceType = "media"
	AccountBrowserRenderingGetHTMLContentParamsAllowResourceTypeFont               AccountBrowserRenderingGetHTMLContentParamsAllowResourceType = "font"
	AccountBrowserRenderingGetHTMLContentParamsAllowResourceTypeScript             AccountBrowserRenderingGetHTMLContentParamsAllowResourceType = "script"
	AccountBrowserRenderingGetHTMLContentParamsAllowResourceTypeTexttrack          AccountBrowserRenderingGetHTMLContentParamsAllowResourceType = "texttrack"
	AccountBrowserRenderingGetHTMLContentParamsAllowResourceTypeXhr                AccountBrowserRenderingGetHTMLContentParamsAllowResourceType = "xhr"
	AccountBrowserRenderingGetHTMLContentParamsAllowResourceTypeFetch              AccountBrowserRenderingGetHTMLContentParamsAllowResourceType = "fetch"
	AccountBrowserRenderingGetHTMLContentParamsAllowResourceTypePrefetch           AccountBrowserRenderingGetHTMLContentParamsAllowResourceType = "prefetch"
	AccountBrowserRenderingGetHTMLContentParamsAllowResourceTypeEventsource        AccountBrowserRenderingGetHTMLContentParamsAllowResourceType = "eventsource"
	AccountBrowserRenderingGetHTMLContentParamsAllowResourceTypeWebsocket          AccountBrowserRenderingGetHTMLContentParamsAllowResourceType = "websocket"
	AccountBrowserRenderingGetHTMLContentParamsAllowResourceTypeManifest           AccountBrowserRenderingGetHTMLContentParamsAllowResourceType = "manifest"
	AccountBrowserRenderingGetHTMLContentParamsAllowResourceTypeSignedexchange     AccountBrowserRenderingGetHTMLContentParamsAllowResourceType = "signedexchange"
	AccountBrowserRenderingGetHTMLContentParamsAllowResourceTypePing               AccountBrowserRenderingGetHTMLContentParamsAllowResourceType = "ping"
	AccountBrowserRenderingGetHTMLContentParamsAllowResourceTypeCspviolationreport AccountBrowserRenderingGetHTMLContentParamsAllowResourceType = "cspviolationreport"
	AccountBrowserRenderingGetHTMLContentParamsAllowResourceTypePreflight          AccountBrowserRenderingGetHTMLContentParamsAllowResourceType = "preflight"
	AccountBrowserRenderingGetHTMLContentParamsAllowResourceTypeOther              AccountBrowserRenderingGetHTMLContentParamsAllowResourceType = "other"
)

func (r AccountBrowserRenderingGetHTMLContentParamsAllowResourceType) IsKnown() bool {
	switch r {
	case AccountBrowserRenderingGetHTMLContentParamsAllowResourceTypeDocument, AccountBrowserRenderingGetHTMLContentParamsAllowResourceTypeStylesheet, AccountBrowserRenderingGetHTMLContentParamsAllowResourceTypeImage, AccountBrowserRenderingGetHTMLContentParamsAllowResourceTypeMedia, AccountBrowserRenderingGetHTMLContentParamsAllowResourceTypeFont, AccountBrowserRenderingGetHTMLContentParamsAllowResourceTypeScript, AccountBrowserRenderingGetHTMLContentParamsAllowResourceTypeTexttrack, AccountBrowserRenderingGetHTMLContentParamsAllowResourceTypeXhr, AccountBrowserRenderingGetHTMLContentParamsAllowResourceTypeFetch, AccountBrowserRenderingGetHTMLContentParamsAllowResourceTypePrefetch, AccountBrowserRenderingGetHTMLContentParamsAllowResourceTypeEventsource, AccountBrowserRenderingGetHTMLContentParamsAllowResourceTypeWebsocket, AccountBrowserRenderingGetHTMLContentParamsAllowResourceTypeManifest, AccountBrowserRenderingGetHTMLContentParamsAllowResourceTypeSignedexchange, AccountBrowserRenderingGetHTMLContentParamsAllowResourceTypePing, AccountBrowserRenderingGetHTMLContentParamsAllowResourceTypeCspviolationreport, AccountBrowserRenderingGetHTMLContentParamsAllowResourceTypePreflight, AccountBrowserRenderingGetHTMLContentParamsAllowResourceTypeOther:
		return true
	}
	return false
}

// Provide credentials for HTTP authentication.
type AccountBrowserRenderingGetHTMLContentParamsAuthenticate struct {
	Password param.Field[string] `json:"password,required"`
	Username param.Field[string] `json:"username,required"`
}

func (r AccountBrowserRenderingGetHTMLContentParamsAuthenticate) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountBrowserRenderingGetHTMLContentParamsCookie struct {
	Name         param.Field[string]                                                         `json:"name,required"`
	Value        param.Field[string]                                                         `json:"value,required"`
	Domain       param.Field[string]                                                         `json:"domain"`
	Expires      param.Field[float64]                                                        `json:"expires"`
	HTTPOnly     param.Field[bool]                                                           `json:"httpOnly"`
	PartitionKey param.Field[string]                                                         `json:"partitionKey"`
	Path         param.Field[string]                                                         `json:"path"`
	Priority     param.Field[AccountBrowserRenderingGetHTMLContentParamsCookiesPriority]     `json:"priority"`
	SameParty    param.Field[bool]                                                           `json:"sameParty"`
	SameSite     param.Field[AccountBrowserRenderingGetHTMLContentParamsCookiesSameSite]     `json:"sameSite"`
	Secure       param.Field[bool]                                                           `json:"secure"`
	SourcePort   param.Field[float64]                                                        `json:"sourcePort"`
	SourceScheme param.Field[AccountBrowserRenderingGetHTMLContentParamsCookiesSourceScheme] `json:"sourceScheme"`
	URL          param.Field[string]                                                         `json:"url"`
}

func (r AccountBrowserRenderingGetHTMLContentParamsCookie) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountBrowserRenderingGetHTMLContentParamsCookiesPriority string

const (
	AccountBrowserRenderingGetHTMLContentParamsCookiesPriorityLow    AccountBrowserRenderingGetHTMLContentParamsCookiesPriority = "Low"
	AccountBrowserRenderingGetHTMLContentParamsCookiesPriorityMedium AccountBrowserRenderingGetHTMLContentParamsCookiesPriority = "Medium"
	AccountBrowserRenderingGetHTMLContentParamsCookiesPriorityHigh   AccountBrowserRenderingGetHTMLContentParamsCookiesPriority = "High"
)

func (r AccountBrowserRenderingGetHTMLContentParamsCookiesPriority) IsKnown() bool {
	switch r {
	case AccountBrowserRenderingGetHTMLContentParamsCookiesPriorityLow, AccountBrowserRenderingGetHTMLContentParamsCookiesPriorityMedium, AccountBrowserRenderingGetHTMLContentParamsCookiesPriorityHigh:
		return true
	}
	return false
}

type AccountBrowserRenderingGetHTMLContentParamsCookiesSameSite string

const (
	AccountBrowserRenderingGetHTMLContentParamsCookiesSameSiteStrict AccountBrowserRenderingGetHTMLContentParamsCookiesSameSite = "Strict"
	AccountBrowserRenderingGetHTMLContentParamsCookiesSameSiteLax    AccountBrowserRenderingGetHTMLContentParamsCookiesSameSite = "Lax"
	AccountBrowserRenderingGetHTMLContentParamsCookiesSameSiteNone   AccountBrowserRenderingGetHTMLContentParamsCookiesSameSite = "None"
)

func (r AccountBrowserRenderingGetHTMLContentParamsCookiesSameSite) IsKnown() bool {
	switch r {
	case AccountBrowserRenderingGetHTMLContentParamsCookiesSameSiteStrict, AccountBrowserRenderingGetHTMLContentParamsCookiesSameSiteLax, AccountBrowserRenderingGetHTMLContentParamsCookiesSameSiteNone:
		return true
	}
	return false
}

type AccountBrowserRenderingGetHTMLContentParamsCookiesSourceScheme string

const (
	AccountBrowserRenderingGetHTMLContentParamsCookiesSourceSchemeUnset     AccountBrowserRenderingGetHTMLContentParamsCookiesSourceScheme = "Unset"
	AccountBrowserRenderingGetHTMLContentParamsCookiesSourceSchemeNonSecure AccountBrowserRenderingGetHTMLContentParamsCookiesSourceScheme = "NonSecure"
	AccountBrowserRenderingGetHTMLContentParamsCookiesSourceSchemeSecure    AccountBrowserRenderingGetHTMLContentParamsCookiesSourceScheme = "Secure"
)

func (r AccountBrowserRenderingGetHTMLContentParamsCookiesSourceScheme) IsKnown() bool {
	switch r {
	case AccountBrowserRenderingGetHTMLContentParamsCookiesSourceSchemeUnset, AccountBrowserRenderingGetHTMLContentParamsCookiesSourceSchemeNonSecure, AccountBrowserRenderingGetHTMLContentParamsCookiesSourceSchemeSecure:
		return true
	}
	return false
}

// Check [options](https://pptr.dev/api/puppeteer.gotooptions).
type AccountBrowserRenderingGetHTMLContentParamsGotoOptions struct {
	Referer        param.Field[string]                                                               `json:"referer"`
	ReferrerPolicy param.Field[string]                                                               `json:"referrerPolicy"`
	Timeout        param.Field[float64]                                                              `json:"timeout"`
	WaitUntil      param.Field[AccountBrowserRenderingGetHTMLContentParamsGotoOptionsWaitUntilUnion] `json:"waitUntil"`
}

func (r AccountBrowserRenderingGetHTMLContentParamsGotoOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by
// [AccountBrowserRenderingGetHTMLContentParamsGotoOptionsWaitUntilString],
// [AccountBrowserRenderingGetHTMLContentParamsGotoOptionsWaitUntilArray].
type AccountBrowserRenderingGetHTMLContentParamsGotoOptionsWaitUntilUnion interface {
	implementsAccountBrowserRenderingGetHTMLContentParamsGotoOptionsWaitUntilUnion()
}

type AccountBrowserRenderingGetHTMLContentParamsGotoOptionsWaitUntilString string

const (
	AccountBrowserRenderingGetHTMLContentParamsGotoOptionsWaitUntilStringLoad             AccountBrowserRenderingGetHTMLContentParamsGotoOptionsWaitUntilString = "load"
	AccountBrowserRenderingGetHTMLContentParamsGotoOptionsWaitUntilStringDomcontentloaded AccountBrowserRenderingGetHTMLContentParamsGotoOptionsWaitUntilString = "domcontentloaded"
	AccountBrowserRenderingGetHTMLContentParamsGotoOptionsWaitUntilStringNetworkidle0     AccountBrowserRenderingGetHTMLContentParamsGotoOptionsWaitUntilString = "networkidle0"
	AccountBrowserRenderingGetHTMLContentParamsGotoOptionsWaitUntilStringNetworkidle2     AccountBrowserRenderingGetHTMLContentParamsGotoOptionsWaitUntilString = "networkidle2"
)

func (r AccountBrowserRenderingGetHTMLContentParamsGotoOptionsWaitUntilString) IsKnown() bool {
	switch r {
	case AccountBrowserRenderingGetHTMLContentParamsGotoOptionsWaitUntilStringLoad, AccountBrowserRenderingGetHTMLContentParamsGotoOptionsWaitUntilStringDomcontentloaded, AccountBrowserRenderingGetHTMLContentParamsGotoOptionsWaitUntilStringNetworkidle0, AccountBrowserRenderingGetHTMLContentParamsGotoOptionsWaitUntilStringNetworkidle2:
		return true
	}
	return false
}

func (r AccountBrowserRenderingGetHTMLContentParamsGotoOptionsWaitUntilString) implementsAccountBrowserRenderingGetHTMLContentParamsGotoOptionsWaitUntilUnion() {
}

type AccountBrowserRenderingGetHTMLContentParamsGotoOptionsWaitUntilArray []AccountBrowserRenderingGetHTMLContentParamsGotoOptionsWaitUntilArrayItem

func (r AccountBrowserRenderingGetHTMLContentParamsGotoOptionsWaitUntilArray) implementsAccountBrowserRenderingGetHTMLContentParamsGotoOptionsWaitUntilUnion() {
}

type AccountBrowserRenderingGetHTMLContentParamsGotoOptionsWaitUntilArrayItem string

const (
	AccountBrowserRenderingGetHTMLContentParamsGotoOptionsWaitUntilArrayItemLoad             AccountBrowserRenderingGetHTMLContentParamsGotoOptionsWaitUntilArrayItem = "load"
	AccountBrowserRenderingGetHTMLContentParamsGotoOptionsWaitUntilArrayItemDomcontentloaded AccountBrowserRenderingGetHTMLContentParamsGotoOptionsWaitUntilArrayItem = "domcontentloaded"
	AccountBrowserRenderingGetHTMLContentParamsGotoOptionsWaitUntilArrayItemNetworkidle0     AccountBrowserRenderingGetHTMLContentParamsGotoOptionsWaitUntilArrayItem = "networkidle0"
	AccountBrowserRenderingGetHTMLContentParamsGotoOptionsWaitUntilArrayItemNetworkidle2     AccountBrowserRenderingGetHTMLContentParamsGotoOptionsWaitUntilArrayItem = "networkidle2"
)

func (r AccountBrowserRenderingGetHTMLContentParamsGotoOptionsWaitUntilArrayItem) IsKnown() bool {
	switch r {
	case AccountBrowserRenderingGetHTMLContentParamsGotoOptionsWaitUntilArrayItemLoad, AccountBrowserRenderingGetHTMLContentParamsGotoOptionsWaitUntilArrayItemDomcontentloaded, AccountBrowserRenderingGetHTMLContentParamsGotoOptionsWaitUntilArrayItemNetworkidle0, AccountBrowserRenderingGetHTMLContentParamsGotoOptionsWaitUntilArrayItemNetworkidle2:
		return true
	}
	return false
}

type AccountBrowserRenderingGetHTMLContentParamsRejectResourceType string

const (
	AccountBrowserRenderingGetHTMLContentParamsRejectResourceTypeDocument           AccountBrowserRenderingGetHTMLContentParamsRejectResourceType = "document"
	AccountBrowserRenderingGetHTMLContentParamsRejectResourceTypeStylesheet         AccountBrowserRenderingGetHTMLContentParamsRejectResourceType = "stylesheet"
	AccountBrowserRenderingGetHTMLContentParamsRejectResourceTypeImage              AccountBrowserRenderingGetHTMLContentParamsRejectResourceType = "image"
	AccountBrowserRenderingGetHTMLContentParamsRejectResourceTypeMedia              AccountBrowserRenderingGetHTMLContentParamsRejectResourceType = "media"
	AccountBrowserRenderingGetHTMLContentParamsRejectResourceTypeFont               AccountBrowserRenderingGetHTMLContentParamsRejectResourceType = "font"
	AccountBrowserRenderingGetHTMLContentParamsRejectResourceTypeScript             AccountBrowserRenderingGetHTMLContentParamsRejectResourceType = "script"
	AccountBrowserRenderingGetHTMLContentParamsRejectResourceTypeTexttrack          AccountBrowserRenderingGetHTMLContentParamsRejectResourceType = "texttrack"
	AccountBrowserRenderingGetHTMLContentParamsRejectResourceTypeXhr                AccountBrowserRenderingGetHTMLContentParamsRejectResourceType = "xhr"
	AccountBrowserRenderingGetHTMLContentParamsRejectResourceTypeFetch              AccountBrowserRenderingGetHTMLContentParamsRejectResourceType = "fetch"
	AccountBrowserRenderingGetHTMLContentParamsRejectResourceTypePrefetch           AccountBrowserRenderingGetHTMLContentParamsRejectResourceType = "prefetch"
	AccountBrowserRenderingGetHTMLContentParamsRejectResourceTypeEventsource        AccountBrowserRenderingGetHTMLContentParamsRejectResourceType = "eventsource"
	AccountBrowserRenderingGetHTMLContentParamsRejectResourceTypeWebsocket          AccountBrowserRenderingGetHTMLContentParamsRejectResourceType = "websocket"
	AccountBrowserRenderingGetHTMLContentParamsRejectResourceTypeManifest           AccountBrowserRenderingGetHTMLContentParamsRejectResourceType = "manifest"
	AccountBrowserRenderingGetHTMLContentParamsRejectResourceTypeSignedexchange     AccountBrowserRenderingGetHTMLContentParamsRejectResourceType = "signedexchange"
	AccountBrowserRenderingGetHTMLContentParamsRejectResourceTypePing               AccountBrowserRenderingGetHTMLContentParamsRejectResourceType = "ping"
	AccountBrowserRenderingGetHTMLContentParamsRejectResourceTypeCspviolationreport AccountBrowserRenderingGetHTMLContentParamsRejectResourceType = "cspviolationreport"
	AccountBrowserRenderingGetHTMLContentParamsRejectResourceTypePreflight          AccountBrowserRenderingGetHTMLContentParamsRejectResourceType = "preflight"
	AccountBrowserRenderingGetHTMLContentParamsRejectResourceTypeOther              AccountBrowserRenderingGetHTMLContentParamsRejectResourceType = "other"
)

func (r AccountBrowserRenderingGetHTMLContentParamsRejectResourceType) IsKnown() bool {
	switch r {
	case AccountBrowserRenderingGetHTMLContentParamsRejectResourceTypeDocument, AccountBrowserRenderingGetHTMLContentParamsRejectResourceTypeStylesheet, AccountBrowserRenderingGetHTMLContentParamsRejectResourceTypeImage, AccountBrowserRenderingGetHTMLContentParamsRejectResourceTypeMedia, AccountBrowserRenderingGetHTMLContentParamsRejectResourceTypeFont, AccountBrowserRenderingGetHTMLContentParamsRejectResourceTypeScript, AccountBrowserRenderingGetHTMLContentParamsRejectResourceTypeTexttrack, AccountBrowserRenderingGetHTMLContentParamsRejectResourceTypeXhr, AccountBrowserRenderingGetHTMLContentParamsRejectResourceTypeFetch, AccountBrowserRenderingGetHTMLContentParamsRejectResourceTypePrefetch, AccountBrowserRenderingGetHTMLContentParamsRejectResourceTypeEventsource, AccountBrowserRenderingGetHTMLContentParamsRejectResourceTypeWebsocket, AccountBrowserRenderingGetHTMLContentParamsRejectResourceTypeManifest, AccountBrowserRenderingGetHTMLContentParamsRejectResourceTypeSignedexchange, AccountBrowserRenderingGetHTMLContentParamsRejectResourceTypePing, AccountBrowserRenderingGetHTMLContentParamsRejectResourceTypeCspviolationreport, AccountBrowserRenderingGetHTMLContentParamsRejectResourceTypePreflight, AccountBrowserRenderingGetHTMLContentParamsRejectResourceTypeOther:
		return true
	}
	return false
}

// Check [options](https://pptr.dev/api/puppeteer.page.setviewport).
type AccountBrowserRenderingGetHTMLContentParamsViewport struct {
	Height            param.Field[float64] `json:"height,required"`
	Width             param.Field[float64] `json:"width,required"`
	DeviceScaleFactor param.Field[float64] `json:"deviceScaleFactor"`
	HasTouch          param.Field[bool]    `json:"hasTouch"`
	IsLandscape       param.Field[bool]    `json:"isLandscape"`
	IsMobile          param.Field[bool]    `json:"isMobile"`
}

func (r AccountBrowserRenderingGetHTMLContentParamsViewport) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Wait for the selector to appear in page. Check
// [options](https://pptr.dev/api/puppeteer.page.waitforselector).
type AccountBrowserRenderingGetHTMLContentParamsWaitForSelector struct {
	Selector param.Field[string]                                                            `json:"selector,required"`
	Hidden   param.Field[AccountBrowserRenderingGetHTMLContentParamsWaitForSelectorHidden]  `json:"hidden"`
	Timeout  param.Field[float64]                                                           `json:"timeout"`
	Visible  param.Field[AccountBrowserRenderingGetHTMLContentParamsWaitForSelectorVisible] `json:"visible"`
}

func (r AccountBrowserRenderingGetHTMLContentParamsWaitForSelector) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountBrowserRenderingGetHTMLContentParamsWaitForSelectorHidden bool

const (
	AccountBrowserRenderingGetHTMLContentParamsWaitForSelectorHiddenTrue AccountBrowserRenderingGetHTMLContentParamsWaitForSelectorHidden = true
)

func (r AccountBrowserRenderingGetHTMLContentParamsWaitForSelectorHidden) IsKnown() bool {
	switch r {
	case AccountBrowserRenderingGetHTMLContentParamsWaitForSelectorHiddenTrue:
		return true
	}
	return false
}

type AccountBrowserRenderingGetHTMLContentParamsWaitForSelectorVisible bool

const (
	AccountBrowserRenderingGetHTMLContentParamsWaitForSelectorVisibleTrue AccountBrowserRenderingGetHTMLContentParamsWaitForSelectorVisible = true
)

func (r AccountBrowserRenderingGetHTMLContentParamsWaitForSelectorVisible) IsKnown() bool {
	switch r {
	case AccountBrowserRenderingGetHTMLContentParamsWaitForSelectorVisibleTrue:
		return true
	}
	return false
}

type AccountBrowserRenderingGetJsonParams struct {
	// Cache TTL default is 5s. Set to 0 to disable.
	CacheTtl param.Field[float64] `query:"cacheTTL"`
	// Adds a `<script>` tag into the page with the desired URL or content.
	AddScriptTag param.Field[[]AccountBrowserRenderingGetJsonParamsAddScriptTag] `json:"addScriptTag"`
	// Adds a `<link rel="stylesheet">` tag into the page with the desired URL or a
	// `<style type="text/css">` tag with the content.
	AddStyleTag param.Field[[]AccountBrowserRenderingGetJsonParamsAddStyleTag] `json:"addStyleTag"`
	// Only allow requests that match the provided regex patterns, eg. '/^.\*\.(css)'.
	AllowRequestPattern param.Field[[]string] `json:"allowRequestPattern"`
	// Only allow requests that match the provided resource types, eg. 'image' or
	// 'script'.
	AllowResourceTypes param.Field[[]AccountBrowserRenderingGetJsonParamsAllowResourceType] `json:"allowResourceTypes"`
	// Provide credentials for HTTP authentication.
	Authenticate param.Field[AccountBrowserRenderingGetJsonParamsAuthenticate] `json:"authenticate"`
	// Attempt to proceed when 'awaited' events fail or timeout.
	BestAttempt param.Field[bool] `json:"bestAttempt"`
	// Check [options](https://pptr.dev/api/puppeteer.page.setcookie).
	Cookies          param.Field[[]AccountBrowserRenderingGetJsonParamsCookie] `json:"cookies"`
	EmulateMediaType param.Field[string]                                       `json:"emulateMediaType"`
	// Check [options](https://pptr.dev/api/puppeteer.gotooptions).
	GotoOptions param.Field[AccountBrowserRenderingGetJsonParamsGotoOptions] `json:"gotoOptions"`
	// Set the content of the page, eg: `<h1>Hello World!!</h1>`. Either `html` or
	// `url` must be set.
	HTML   param.Field[string] `json:"html"`
	Prompt param.Field[string] `json:"prompt"`
	// Block undesired requests that match the provided regex patterns, eg.
	// '/^.\*\.(css)'.
	RejectRequestPattern param.Field[[]string] `json:"rejectRequestPattern"`
	// Block undesired requests that match the provided resource types, eg. 'image' or
	// 'script'.
	RejectResourceTypes  param.Field[[]AccountBrowserRenderingGetJsonParamsRejectResourceType] `json:"rejectResourceTypes"`
	ResponseFormat       param.Field[AccountBrowserRenderingGetJsonParamsResponseFormat]       `json:"response_format"`
	SetExtraHTTPHeaders  param.Field[map[string]string]                                        `json:"setExtraHTTPHeaders"`
	SetJavaScriptEnabled param.Field[bool]                                                     `json:"setJavaScriptEnabled"`
	// URL to navigate to, eg. `https://example.com`.
	URL       param.Field[string] `json:"url" format:"uri"`
	UserAgent param.Field[string] `json:"userAgent"`
	// Check [options](https://pptr.dev/api/puppeteer.page.setviewport).
	Viewport param.Field[AccountBrowserRenderingGetJsonParamsViewport] `json:"viewport"`
	// Wait for the selector to appear in page. Check
	// [options](https://pptr.dev/api/puppeteer.page.waitforselector).
	WaitForSelector param.Field[AccountBrowserRenderingGetJsonParamsWaitForSelector] `json:"waitForSelector"`
	// Waits for a specified timeout before continuing.
	WaitForTimeout param.Field[float64] `json:"waitForTimeout"`
}

func (r AccountBrowserRenderingGetJsonParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [AccountBrowserRenderingGetJsonParams]'s query parameters as
// `url.Values`.
func (r AccountBrowserRenderingGetJsonParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountBrowserRenderingGetJsonParamsAddScriptTag struct {
	ID      param.Field[string] `json:"id"`
	Content param.Field[string] `json:"content"`
	Type    param.Field[string] `json:"type"`
	URL     param.Field[string] `json:"url"`
}

func (r AccountBrowserRenderingGetJsonParamsAddScriptTag) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountBrowserRenderingGetJsonParamsAddStyleTag struct {
	Content param.Field[string] `json:"content"`
	URL     param.Field[string] `json:"url"`
}

func (r AccountBrowserRenderingGetJsonParamsAddStyleTag) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountBrowserRenderingGetJsonParamsAllowResourceType string

const (
	AccountBrowserRenderingGetJsonParamsAllowResourceTypeDocument           AccountBrowserRenderingGetJsonParamsAllowResourceType = "document"
	AccountBrowserRenderingGetJsonParamsAllowResourceTypeStylesheet         AccountBrowserRenderingGetJsonParamsAllowResourceType = "stylesheet"
	AccountBrowserRenderingGetJsonParamsAllowResourceTypeImage              AccountBrowserRenderingGetJsonParamsAllowResourceType = "image"
	AccountBrowserRenderingGetJsonParamsAllowResourceTypeMedia              AccountBrowserRenderingGetJsonParamsAllowResourceType = "media"
	AccountBrowserRenderingGetJsonParamsAllowResourceTypeFont               AccountBrowserRenderingGetJsonParamsAllowResourceType = "font"
	AccountBrowserRenderingGetJsonParamsAllowResourceTypeScript             AccountBrowserRenderingGetJsonParamsAllowResourceType = "script"
	AccountBrowserRenderingGetJsonParamsAllowResourceTypeTexttrack          AccountBrowserRenderingGetJsonParamsAllowResourceType = "texttrack"
	AccountBrowserRenderingGetJsonParamsAllowResourceTypeXhr                AccountBrowserRenderingGetJsonParamsAllowResourceType = "xhr"
	AccountBrowserRenderingGetJsonParamsAllowResourceTypeFetch              AccountBrowserRenderingGetJsonParamsAllowResourceType = "fetch"
	AccountBrowserRenderingGetJsonParamsAllowResourceTypePrefetch           AccountBrowserRenderingGetJsonParamsAllowResourceType = "prefetch"
	AccountBrowserRenderingGetJsonParamsAllowResourceTypeEventsource        AccountBrowserRenderingGetJsonParamsAllowResourceType = "eventsource"
	AccountBrowserRenderingGetJsonParamsAllowResourceTypeWebsocket          AccountBrowserRenderingGetJsonParamsAllowResourceType = "websocket"
	AccountBrowserRenderingGetJsonParamsAllowResourceTypeManifest           AccountBrowserRenderingGetJsonParamsAllowResourceType = "manifest"
	AccountBrowserRenderingGetJsonParamsAllowResourceTypeSignedexchange     AccountBrowserRenderingGetJsonParamsAllowResourceType = "signedexchange"
	AccountBrowserRenderingGetJsonParamsAllowResourceTypePing               AccountBrowserRenderingGetJsonParamsAllowResourceType = "ping"
	AccountBrowserRenderingGetJsonParamsAllowResourceTypeCspviolationreport AccountBrowserRenderingGetJsonParamsAllowResourceType = "cspviolationreport"
	AccountBrowserRenderingGetJsonParamsAllowResourceTypePreflight          AccountBrowserRenderingGetJsonParamsAllowResourceType = "preflight"
	AccountBrowserRenderingGetJsonParamsAllowResourceTypeOther              AccountBrowserRenderingGetJsonParamsAllowResourceType = "other"
)

func (r AccountBrowserRenderingGetJsonParamsAllowResourceType) IsKnown() bool {
	switch r {
	case AccountBrowserRenderingGetJsonParamsAllowResourceTypeDocument, AccountBrowserRenderingGetJsonParamsAllowResourceTypeStylesheet, AccountBrowserRenderingGetJsonParamsAllowResourceTypeImage, AccountBrowserRenderingGetJsonParamsAllowResourceTypeMedia, AccountBrowserRenderingGetJsonParamsAllowResourceTypeFont, AccountBrowserRenderingGetJsonParamsAllowResourceTypeScript, AccountBrowserRenderingGetJsonParamsAllowResourceTypeTexttrack, AccountBrowserRenderingGetJsonParamsAllowResourceTypeXhr, AccountBrowserRenderingGetJsonParamsAllowResourceTypeFetch, AccountBrowserRenderingGetJsonParamsAllowResourceTypePrefetch, AccountBrowserRenderingGetJsonParamsAllowResourceTypeEventsource, AccountBrowserRenderingGetJsonParamsAllowResourceTypeWebsocket, AccountBrowserRenderingGetJsonParamsAllowResourceTypeManifest, AccountBrowserRenderingGetJsonParamsAllowResourceTypeSignedexchange, AccountBrowserRenderingGetJsonParamsAllowResourceTypePing, AccountBrowserRenderingGetJsonParamsAllowResourceTypeCspviolationreport, AccountBrowserRenderingGetJsonParamsAllowResourceTypePreflight, AccountBrowserRenderingGetJsonParamsAllowResourceTypeOther:
		return true
	}
	return false
}

// Provide credentials for HTTP authentication.
type AccountBrowserRenderingGetJsonParamsAuthenticate struct {
	Password param.Field[string] `json:"password,required"`
	Username param.Field[string] `json:"username,required"`
}

func (r AccountBrowserRenderingGetJsonParamsAuthenticate) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountBrowserRenderingGetJsonParamsCookie struct {
	Name         param.Field[string]                                                  `json:"name,required"`
	Value        param.Field[string]                                                  `json:"value,required"`
	Domain       param.Field[string]                                                  `json:"domain"`
	Expires      param.Field[float64]                                                 `json:"expires"`
	HTTPOnly     param.Field[bool]                                                    `json:"httpOnly"`
	PartitionKey param.Field[string]                                                  `json:"partitionKey"`
	Path         param.Field[string]                                                  `json:"path"`
	Priority     param.Field[AccountBrowserRenderingGetJsonParamsCookiesPriority]     `json:"priority"`
	SameParty    param.Field[bool]                                                    `json:"sameParty"`
	SameSite     param.Field[AccountBrowserRenderingGetJsonParamsCookiesSameSite]     `json:"sameSite"`
	Secure       param.Field[bool]                                                    `json:"secure"`
	SourcePort   param.Field[float64]                                                 `json:"sourcePort"`
	SourceScheme param.Field[AccountBrowserRenderingGetJsonParamsCookiesSourceScheme] `json:"sourceScheme"`
	URL          param.Field[string]                                                  `json:"url"`
}

func (r AccountBrowserRenderingGetJsonParamsCookie) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountBrowserRenderingGetJsonParamsCookiesPriority string

const (
	AccountBrowserRenderingGetJsonParamsCookiesPriorityLow    AccountBrowserRenderingGetJsonParamsCookiesPriority = "Low"
	AccountBrowserRenderingGetJsonParamsCookiesPriorityMedium AccountBrowserRenderingGetJsonParamsCookiesPriority = "Medium"
	AccountBrowserRenderingGetJsonParamsCookiesPriorityHigh   AccountBrowserRenderingGetJsonParamsCookiesPriority = "High"
)

func (r AccountBrowserRenderingGetJsonParamsCookiesPriority) IsKnown() bool {
	switch r {
	case AccountBrowserRenderingGetJsonParamsCookiesPriorityLow, AccountBrowserRenderingGetJsonParamsCookiesPriorityMedium, AccountBrowserRenderingGetJsonParamsCookiesPriorityHigh:
		return true
	}
	return false
}

type AccountBrowserRenderingGetJsonParamsCookiesSameSite string

const (
	AccountBrowserRenderingGetJsonParamsCookiesSameSiteStrict AccountBrowserRenderingGetJsonParamsCookiesSameSite = "Strict"
	AccountBrowserRenderingGetJsonParamsCookiesSameSiteLax    AccountBrowserRenderingGetJsonParamsCookiesSameSite = "Lax"
	AccountBrowserRenderingGetJsonParamsCookiesSameSiteNone   AccountBrowserRenderingGetJsonParamsCookiesSameSite = "None"
)

func (r AccountBrowserRenderingGetJsonParamsCookiesSameSite) IsKnown() bool {
	switch r {
	case AccountBrowserRenderingGetJsonParamsCookiesSameSiteStrict, AccountBrowserRenderingGetJsonParamsCookiesSameSiteLax, AccountBrowserRenderingGetJsonParamsCookiesSameSiteNone:
		return true
	}
	return false
}

type AccountBrowserRenderingGetJsonParamsCookiesSourceScheme string

const (
	AccountBrowserRenderingGetJsonParamsCookiesSourceSchemeUnset     AccountBrowserRenderingGetJsonParamsCookiesSourceScheme = "Unset"
	AccountBrowserRenderingGetJsonParamsCookiesSourceSchemeNonSecure AccountBrowserRenderingGetJsonParamsCookiesSourceScheme = "NonSecure"
	AccountBrowserRenderingGetJsonParamsCookiesSourceSchemeSecure    AccountBrowserRenderingGetJsonParamsCookiesSourceScheme = "Secure"
)

func (r AccountBrowserRenderingGetJsonParamsCookiesSourceScheme) IsKnown() bool {
	switch r {
	case AccountBrowserRenderingGetJsonParamsCookiesSourceSchemeUnset, AccountBrowserRenderingGetJsonParamsCookiesSourceSchemeNonSecure, AccountBrowserRenderingGetJsonParamsCookiesSourceSchemeSecure:
		return true
	}
	return false
}

// Check [options](https://pptr.dev/api/puppeteer.gotooptions).
type AccountBrowserRenderingGetJsonParamsGotoOptions struct {
	Referer        param.Field[string]                                                        `json:"referer"`
	ReferrerPolicy param.Field[string]                                                        `json:"referrerPolicy"`
	Timeout        param.Field[float64]                                                       `json:"timeout"`
	WaitUntil      param.Field[AccountBrowserRenderingGetJsonParamsGotoOptionsWaitUntilUnion] `json:"waitUntil"`
}

func (r AccountBrowserRenderingGetJsonParamsGotoOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by [AccountBrowserRenderingGetJsonParamsGotoOptionsWaitUntilString],
// [AccountBrowserRenderingGetJsonParamsGotoOptionsWaitUntilArray].
type AccountBrowserRenderingGetJsonParamsGotoOptionsWaitUntilUnion interface {
	implementsAccountBrowserRenderingGetJsonParamsGotoOptionsWaitUntilUnion()
}

type AccountBrowserRenderingGetJsonParamsGotoOptionsWaitUntilString string

const (
	AccountBrowserRenderingGetJsonParamsGotoOptionsWaitUntilStringLoad             AccountBrowserRenderingGetJsonParamsGotoOptionsWaitUntilString = "load"
	AccountBrowserRenderingGetJsonParamsGotoOptionsWaitUntilStringDomcontentloaded AccountBrowserRenderingGetJsonParamsGotoOptionsWaitUntilString = "domcontentloaded"
	AccountBrowserRenderingGetJsonParamsGotoOptionsWaitUntilStringNetworkidle0     AccountBrowserRenderingGetJsonParamsGotoOptionsWaitUntilString = "networkidle0"
	AccountBrowserRenderingGetJsonParamsGotoOptionsWaitUntilStringNetworkidle2     AccountBrowserRenderingGetJsonParamsGotoOptionsWaitUntilString = "networkidle2"
)

func (r AccountBrowserRenderingGetJsonParamsGotoOptionsWaitUntilString) IsKnown() bool {
	switch r {
	case AccountBrowserRenderingGetJsonParamsGotoOptionsWaitUntilStringLoad, AccountBrowserRenderingGetJsonParamsGotoOptionsWaitUntilStringDomcontentloaded, AccountBrowserRenderingGetJsonParamsGotoOptionsWaitUntilStringNetworkidle0, AccountBrowserRenderingGetJsonParamsGotoOptionsWaitUntilStringNetworkidle2:
		return true
	}
	return false
}

func (r AccountBrowserRenderingGetJsonParamsGotoOptionsWaitUntilString) implementsAccountBrowserRenderingGetJsonParamsGotoOptionsWaitUntilUnion() {
}

type AccountBrowserRenderingGetJsonParamsGotoOptionsWaitUntilArray []AccountBrowserRenderingGetJsonParamsGotoOptionsWaitUntilArrayItem

func (r AccountBrowserRenderingGetJsonParamsGotoOptionsWaitUntilArray) implementsAccountBrowserRenderingGetJsonParamsGotoOptionsWaitUntilUnion() {
}

type AccountBrowserRenderingGetJsonParamsGotoOptionsWaitUntilArrayItem string

const (
	AccountBrowserRenderingGetJsonParamsGotoOptionsWaitUntilArrayItemLoad             AccountBrowserRenderingGetJsonParamsGotoOptionsWaitUntilArrayItem = "load"
	AccountBrowserRenderingGetJsonParamsGotoOptionsWaitUntilArrayItemDomcontentloaded AccountBrowserRenderingGetJsonParamsGotoOptionsWaitUntilArrayItem = "domcontentloaded"
	AccountBrowserRenderingGetJsonParamsGotoOptionsWaitUntilArrayItemNetworkidle0     AccountBrowserRenderingGetJsonParamsGotoOptionsWaitUntilArrayItem = "networkidle0"
	AccountBrowserRenderingGetJsonParamsGotoOptionsWaitUntilArrayItemNetworkidle2     AccountBrowserRenderingGetJsonParamsGotoOptionsWaitUntilArrayItem = "networkidle2"
)

func (r AccountBrowserRenderingGetJsonParamsGotoOptionsWaitUntilArrayItem) IsKnown() bool {
	switch r {
	case AccountBrowserRenderingGetJsonParamsGotoOptionsWaitUntilArrayItemLoad, AccountBrowserRenderingGetJsonParamsGotoOptionsWaitUntilArrayItemDomcontentloaded, AccountBrowserRenderingGetJsonParamsGotoOptionsWaitUntilArrayItemNetworkidle0, AccountBrowserRenderingGetJsonParamsGotoOptionsWaitUntilArrayItemNetworkidle2:
		return true
	}
	return false
}

type AccountBrowserRenderingGetJsonParamsRejectResourceType string

const (
	AccountBrowserRenderingGetJsonParamsRejectResourceTypeDocument           AccountBrowserRenderingGetJsonParamsRejectResourceType = "document"
	AccountBrowserRenderingGetJsonParamsRejectResourceTypeStylesheet         AccountBrowserRenderingGetJsonParamsRejectResourceType = "stylesheet"
	AccountBrowserRenderingGetJsonParamsRejectResourceTypeImage              AccountBrowserRenderingGetJsonParamsRejectResourceType = "image"
	AccountBrowserRenderingGetJsonParamsRejectResourceTypeMedia              AccountBrowserRenderingGetJsonParamsRejectResourceType = "media"
	AccountBrowserRenderingGetJsonParamsRejectResourceTypeFont               AccountBrowserRenderingGetJsonParamsRejectResourceType = "font"
	AccountBrowserRenderingGetJsonParamsRejectResourceTypeScript             AccountBrowserRenderingGetJsonParamsRejectResourceType = "script"
	AccountBrowserRenderingGetJsonParamsRejectResourceTypeTexttrack          AccountBrowserRenderingGetJsonParamsRejectResourceType = "texttrack"
	AccountBrowserRenderingGetJsonParamsRejectResourceTypeXhr                AccountBrowserRenderingGetJsonParamsRejectResourceType = "xhr"
	AccountBrowserRenderingGetJsonParamsRejectResourceTypeFetch              AccountBrowserRenderingGetJsonParamsRejectResourceType = "fetch"
	AccountBrowserRenderingGetJsonParamsRejectResourceTypePrefetch           AccountBrowserRenderingGetJsonParamsRejectResourceType = "prefetch"
	AccountBrowserRenderingGetJsonParamsRejectResourceTypeEventsource        AccountBrowserRenderingGetJsonParamsRejectResourceType = "eventsource"
	AccountBrowserRenderingGetJsonParamsRejectResourceTypeWebsocket          AccountBrowserRenderingGetJsonParamsRejectResourceType = "websocket"
	AccountBrowserRenderingGetJsonParamsRejectResourceTypeManifest           AccountBrowserRenderingGetJsonParamsRejectResourceType = "manifest"
	AccountBrowserRenderingGetJsonParamsRejectResourceTypeSignedexchange     AccountBrowserRenderingGetJsonParamsRejectResourceType = "signedexchange"
	AccountBrowserRenderingGetJsonParamsRejectResourceTypePing               AccountBrowserRenderingGetJsonParamsRejectResourceType = "ping"
	AccountBrowserRenderingGetJsonParamsRejectResourceTypeCspviolationreport AccountBrowserRenderingGetJsonParamsRejectResourceType = "cspviolationreport"
	AccountBrowserRenderingGetJsonParamsRejectResourceTypePreflight          AccountBrowserRenderingGetJsonParamsRejectResourceType = "preflight"
	AccountBrowserRenderingGetJsonParamsRejectResourceTypeOther              AccountBrowserRenderingGetJsonParamsRejectResourceType = "other"
)

func (r AccountBrowserRenderingGetJsonParamsRejectResourceType) IsKnown() bool {
	switch r {
	case AccountBrowserRenderingGetJsonParamsRejectResourceTypeDocument, AccountBrowserRenderingGetJsonParamsRejectResourceTypeStylesheet, AccountBrowserRenderingGetJsonParamsRejectResourceTypeImage, AccountBrowserRenderingGetJsonParamsRejectResourceTypeMedia, AccountBrowserRenderingGetJsonParamsRejectResourceTypeFont, AccountBrowserRenderingGetJsonParamsRejectResourceTypeScript, AccountBrowserRenderingGetJsonParamsRejectResourceTypeTexttrack, AccountBrowserRenderingGetJsonParamsRejectResourceTypeXhr, AccountBrowserRenderingGetJsonParamsRejectResourceTypeFetch, AccountBrowserRenderingGetJsonParamsRejectResourceTypePrefetch, AccountBrowserRenderingGetJsonParamsRejectResourceTypeEventsource, AccountBrowserRenderingGetJsonParamsRejectResourceTypeWebsocket, AccountBrowserRenderingGetJsonParamsRejectResourceTypeManifest, AccountBrowserRenderingGetJsonParamsRejectResourceTypeSignedexchange, AccountBrowserRenderingGetJsonParamsRejectResourceTypePing, AccountBrowserRenderingGetJsonParamsRejectResourceTypeCspviolationreport, AccountBrowserRenderingGetJsonParamsRejectResourceTypePreflight, AccountBrowserRenderingGetJsonParamsRejectResourceTypeOther:
		return true
	}
	return false
}

type AccountBrowserRenderingGetJsonParamsResponseFormat struct {
	Type param.Field[string] `json:"type,required"`
	// Schema for the response format. More information here:
	// https://developers.cloudflare.com/workers-ai/json-mode/
	Schema param.Field[map[string]interface{}] `json:"schema"`
}

func (r AccountBrowserRenderingGetJsonParamsResponseFormat) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Check [options](https://pptr.dev/api/puppeteer.page.setviewport).
type AccountBrowserRenderingGetJsonParamsViewport struct {
	Height            param.Field[float64] `json:"height,required"`
	Width             param.Field[float64] `json:"width,required"`
	DeviceScaleFactor param.Field[float64] `json:"deviceScaleFactor"`
	HasTouch          param.Field[bool]    `json:"hasTouch"`
	IsLandscape       param.Field[bool]    `json:"isLandscape"`
	IsMobile          param.Field[bool]    `json:"isMobile"`
}

func (r AccountBrowserRenderingGetJsonParamsViewport) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Wait for the selector to appear in page. Check
// [options](https://pptr.dev/api/puppeteer.page.waitforselector).
type AccountBrowserRenderingGetJsonParamsWaitForSelector struct {
	Selector param.Field[string]                                                     `json:"selector,required"`
	Hidden   param.Field[AccountBrowserRenderingGetJsonParamsWaitForSelectorHidden]  `json:"hidden"`
	Timeout  param.Field[float64]                                                    `json:"timeout"`
	Visible  param.Field[AccountBrowserRenderingGetJsonParamsWaitForSelectorVisible] `json:"visible"`
}

func (r AccountBrowserRenderingGetJsonParamsWaitForSelector) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountBrowserRenderingGetJsonParamsWaitForSelectorHidden bool

const (
	AccountBrowserRenderingGetJsonParamsWaitForSelectorHiddenTrue AccountBrowserRenderingGetJsonParamsWaitForSelectorHidden = true
)

func (r AccountBrowserRenderingGetJsonParamsWaitForSelectorHidden) IsKnown() bool {
	switch r {
	case AccountBrowserRenderingGetJsonParamsWaitForSelectorHiddenTrue:
		return true
	}
	return false
}

type AccountBrowserRenderingGetJsonParamsWaitForSelectorVisible bool

const (
	AccountBrowserRenderingGetJsonParamsWaitForSelectorVisibleTrue AccountBrowserRenderingGetJsonParamsWaitForSelectorVisible = true
)

func (r AccountBrowserRenderingGetJsonParamsWaitForSelectorVisible) IsKnown() bool {
	switch r {
	case AccountBrowserRenderingGetJsonParamsWaitForSelectorVisibleTrue:
		return true
	}
	return false
}

type AccountBrowserRenderingGetLinksParams struct {
	// Cache TTL default is 5s. Set to 0 to disable.
	CacheTtl param.Field[float64] `query:"cacheTTL"`
	// Adds a `<script>` tag into the page with the desired URL or content.
	AddScriptTag param.Field[[]AccountBrowserRenderingGetLinksParamsAddScriptTag] `json:"addScriptTag"`
	// Adds a `<link rel="stylesheet">` tag into the page with the desired URL or a
	// `<style type="text/css">` tag with the content.
	AddStyleTag param.Field[[]AccountBrowserRenderingGetLinksParamsAddStyleTag] `json:"addStyleTag"`
	// Only allow requests that match the provided regex patterns, eg. '/^.\*\.(css)'.
	AllowRequestPattern param.Field[[]string] `json:"allowRequestPattern"`
	// Only allow requests that match the provided resource types, eg. 'image' or
	// 'script'.
	AllowResourceTypes param.Field[[]AccountBrowserRenderingGetLinksParamsAllowResourceType] `json:"allowResourceTypes"`
	// Provide credentials for HTTP authentication.
	Authenticate param.Field[AccountBrowserRenderingGetLinksParamsAuthenticate] `json:"authenticate"`
	// Attempt to proceed when 'awaited' events fail or timeout.
	BestAttempt param.Field[bool] `json:"bestAttempt"`
	// Check [options](https://pptr.dev/api/puppeteer.page.setcookie).
	Cookies          param.Field[[]AccountBrowserRenderingGetLinksParamsCookie] `json:"cookies"`
	EmulateMediaType param.Field[string]                                        `json:"emulateMediaType"`
	// Check [options](https://pptr.dev/api/puppeteer.gotooptions).
	GotoOptions param.Field[AccountBrowserRenderingGetLinksParamsGotoOptions] `json:"gotoOptions"`
	// Set the content of the page, eg: `<h1>Hello World!!</h1>`. Either `html` or
	// `url` must be set.
	HTML param.Field[string] `json:"html"`
	// Block undesired requests that match the provided regex patterns, eg.
	// '/^.\*\.(css)'.
	RejectRequestPattern param.Field[[]string] `json:"rejectRequestPattern"`
	// Block undesired requests that match the provided resource types, eg. 'image' or
	// 'script'.
	RejectResourceTypes  param.Field[[]AccountBrowserRenderingGetLinksParamsRejectResourceType] `json:"rejectResourceTypes"`
	SetExtraHTTPHeaders  param.Field[map[string]string]                                         `json:"setExtraHTTPHeaders"`
	SetJavaScriptEnabled param.Field[bool]                                                      `json:"setJavaScriptEnabled"`
	// URL to navigate to, eg. `https://example.com`.
	URL       param.Field[string] `json:"url" format:"uri"`
	UserAgent param.Field[string] `json:"userAgent"`
	// Check [options](https://pptr.dev/api/puppeteer.page.setviewport).
	Viewport         param.Field[AccountBrowserRenderingGetLinksParamsViewport] `json:"viewport"`
	VisibleLinksOnly param.Field[bool]                                          `json:"visibleLinksOnly"`
	// Wait for the selector to appear in page. Check
	// [options](https://pptr.dev/api/puppeteer.page.waitforselector).
	WaitForSelector param.Field[AccountBrowserRenderingGetLinksParamsWaitForSelector] `json:"waitForSelector"`
	// Waits for a specified timeout before continuing.
	WaitForTimeout param.Field[float64] `json:"waitForTimeout"`
}

func (r AccountBrowserRenderingGetLinksParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [AccountBrowserRenderingGetLinksParams]'s query parameters
// as `url.Values`.
func (r AccountBrowserRenderingGetLinksParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountBrowserRenderingGetLinksParamsAddScriptTag struct {
	ID      param.Field[string] `json:"id"`
	Content param.Field[string] `json:"content"`
	Type    param.Field[string] `json:"type"`
	URL     param.Field[string] `json:"url"`
}

func (r AccountBrowserRenderingGetLinksParamsAddScriptTag) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountBrowserRenderingGetLinksParamsAddStyleTag struct {
	Content param.Field[string] `json:"content"`
	URL     param.Field[string] `json:"url"`
}

func (r AccountBrowserRenderingGetLinksParamsAddStyleTag) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountBrowserRenderingGetLinksParamsAllowResourceType string

const (
	AccountBrowserRenderingGetLinksParamsAllowResourceTypeDocument           AccountBrowserRenderingGetLinksParamsAllowResourceType = "document"
	AccountBrowserRenderingGetLinksParamsAllowResourceTypeStylesheet         AccountBrowserRenderingGetLinksParamsAllowResourceType = "stylesheet"
	AccountBrowserRenderingGetLinksParamsAllowResourceTypeImage              AccountBrowserRenderingGetLinksParamsAllowResourceType = "image"
	AccountBrowserRenderingGetLinksParamsAllowResourceTypeMedia              AccountBrowserRenderingGetLinksParamsAllowResourceType = "media"
	AccountBrowserRenderingGetLinksParamsAllowResourceTypeFont               AccountBrowserRenderingGetLinksParamsAllowResourceType = "font"
	AccountBrowserRenderingGetLinksParamsAllowResourceTypeScript             AccountBrowserRenderingGetLinksParamsAllowResourceType = "script"
	AccountBrowserRenderingGetLinksParamsAllowResourceTypeTexttrack          AccountBrowserRenderingGetLinksParamsAllowResourceType = "texttrack"
	AccountBrowserRenderingGetLinksParamsAllowResourceTypeXhr                AccountBrowserRenderingGetLinksParamsAllowResourceType = "xhr"
	AccountBrowserRenderingGetLinksParamsAllowResourceTypeFetch              AccountBrowserRenderingGetLinksParamsAllowResourceType = "fetch"
	AccountBrowserRenderingGetLinksParamsAllowResourceTypePrefetch           AccountBrowserRenderingGetLinksParamsAllowResourceType = "prefetch"
	AccountBrowserRenderingGetLinksParamsAllowResourceTypeEventsource        AccountBrowserRenderingGetLinksParamsAllowResourceType = "eventsource"
	AccountBrowserRenderingGetLinksParamsAllowResourceTypeWebsocket          AccountBrowserRenderingGetLinksParamsAllowResourceType = "websocket"
	AccountBrowserRenderingGetLinksParamsAllowResourceTypeManifest           AccountBrowserRenderingGetLinksParamsAllowResourceType = "manifest"
	AccountBrowserRenderingGetLinksParamsAllowResourceTypeSignedexchange     AccountBrowserRenderingGetLinksParamsAllowResourceType = "signedexchange"
	AccountBrowserRenderingGetLinksParamsAllowResourceTypePing               AccountBrowserRenderingGetLinksParamsAllowResourceType = "ping"
	AccountBrowserRenderingGetLinksParamsAllowResourceTypeCspviolationreport AccountBrowserRenderingGetLinksParamsAllowResourceType = "cspviolationreport"
	AccountBrowserRenderingGetLinksParamsAllowResourceTypePreflight          AccountBrowserRenderingGetLinksParamsAllowResourceType = "preflight"
	AccountBrowserRenderingGetLinksParamsAllowResourceTypeOther              AccountBrowserRenderingGetLinksParamsAllowResourceType = "other"
)

func (r AccountBrowserRenderingGetLinksParamsAllowResourceType) IsKnown() bool {
	switch r {
	case AccountBrowserRenderingGetLinksParamsAllowResourceTypeDocument, AccountBrowserRenderingGetLinksParamsAllowResourceTypeStylesheet, AccountBrowserRenderingGetLinksParamsAllowResourceTypeImage, AccountBrowserRenderingGetLinksParamsAllowResourceTypeMedia, AccountBrowserRenderingGetLinksParamsAllowResourceTypeFont, AccountBrowserRenderingGetLinksParamsAllowResourceTypeScript, AccountBrowserRenderingGetLinksParamsAllowResourceTypeTexttrack, AccountBrowserRenderingGetLinksParamsAllowResourceTypeXhr, AccountBrowserRenderingGetLinksParamsAllowResourceTypeFetch, AccountBrowserRenderingGetLinksParamsAllowResourceTypePrefetch, AccountBrowserRenderingGetLinksParamsAllowResourceTypeEventsource, AccountBrowserRenderingGetLinksParamsAllowResourceTypeWebsocket, AccountBrowserRenderingGetLinksParamsAllowResourceTypeManifest, AccountBrowserRenderingGetLinksParamsAllowResourceTypeSignedexchange, AccountBrowserRenderingGetLinksParamsAllowResourceTypePing, AccountBrowserRenderingGetLinksParamsAllowResourceTypeCspviolationreport, AccountBrowserRenderingGetLinksParamsAllowResourceTypePreflight, AccountBrowserRenderingGetLinksParamsAllowResourceTypeOther:
		return true
	}
	return false
}

// Provide credentials for HTTP authentication.
type AccountBrowserRenderingGetLinksParamsAuthenticate struct {
	Password param.Field[string] `json:"password,required"`
	Username param.Field[string] `json:"username,required"`
}

func (r AccountBrowserRenderingGetLinksParamsAuthenticate) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountBrowserRenderingGetLinksParamsCookie struct {
	Name         param.Field[string]                                                   `json:"name,required"`
	Value        param.Field[string]                                                   `json:"value,required"`
	Domain       param.Field[string]                                                   `json:"domain"`
	Expires      param.Field[float64]                                                  `json:"expires"`
	HTTPOnly     param.Field[bool]                                                     `json:"httpOnly"`
	PartitionKey param.Field[string]                                                   `json:"partitionKey"`
	Path         param.Field[string]                                                   `json:"path"`
	Priority     param.Field[AccountBrowserRenderingGetLinksParamsCookiesPriority]     `json:"priority"`
	SameParty    param.Field[bool]                                                     `json:"sameParty"`
	SameSite     param.Field[AccountBrowserRenderingGetLinksParamsCookiesSameSite]     `json:"sameSite"`
	Secure       param.Field[bool]                                                     `json:"secure"`
	SourcePort   param.Field[float64]                                                  `json:"sourcePort"`
	SourceScheme param.Field[AccountBrowserRenderingGetLinksParamsCookiesSourceScheme] `json:"sourceScheme"`
	URL          param.Field[string]                                                   `json:"url"`
}

func (r AccountBrowserRenderingGetLinksParamsCookie) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountBrowserRenderingGetLinksParamsCookiesPriority string

const (
	AccountBrowserRenderingGetLinksParamsCookiesPriorityLow    AccountBrowserRenderingGetLinksParamsCookiesPriority = "Low"
	AccountBrowserRenderingGetLinksParamsCookiesPriorityMedium AccountBrowserRenderingGetLinksParamsCookiesPriority = "Medium"
	AccountBrowserRenderingGetLinksParamsCookiesPriorityHigh   AccountBrowserRenderingGetLinksParamsCookiesPriority = "High"
)

func (r AccountBrowserRenderingGetLinksParamsCookiesPriority) IsKnown() bool {
	switch r {
	case AccountBrowserRenderingGetLinksParamsCookiesPriorityLow, AccountBrowserRenderingGetLinksParamsCookiesPriorityMedium, AccountBrowserRenderingGetLinksParamsCookiesPriorityHigh:
		return true
	}
	return false
}

type AccountBrowserRenderingGetLinksParamsCookiesSameSite string

const (
	AccountBrowserRenderingGetLinksParamsCookiesSameSiteStrict AccountBrowserRenderingGetLinksParamsCookiesSameSite = "Strict"
	AccountBrowserRenderingGetLinksParamsCookiesSameSiteLax    AccountBrowserRenderingGetLinksParamsCookiesSameSite = "Lax"
	AccountBrowserRenderingGetLinksParamsCookiesSameSiteNone   AccountBrowserRenderingGetLinksParamsCookiesSameSite = "None"
)

func (r AccountBrowserRenderingGetLinksParamsCookiesSameSite) IsKnown() bool {
	switch r {
	case AccountBrowserRenderingGetLinksParamsCookiesSameSiteStrict, AccountBrowserRenderingGetLinksParamsCookiesSameSiteLax, AccountBrowserRenderingGetLinksParamsCookiesSameSiteNone:
		return true
	}
	return false
}

type AccountBrowserRenderingGetLinksParamsCookiesSourceScheme string

const (
	AccountBrowserRenderingGetLinksParamsCookiesSourceSchemeUnset     AccountBrowserRenderingGetLinksParamsCookiesSourceScheme = "Unset"
	AccountBrowserRenderingGetLinksParamsCookiesSourceSchemeNonSecure AccountBrowserRenderingGetLinksParamsCookiesSourceScheme = "NonSecure"
	AccountBrowserRenderingGetLinksParamsCookiesSourceSchemeSecure    AccountBrowserRenderingGetLinksParamsCookiesSourceScheme = "Secure"
)

func (r AccountBrowserRenderingGetLinksParamsCookiesSourceScheme) IsKnown() bool {
	switch r {
	case AccountBrowserRenderingGetLinksParamsCookiesSourceSchemeUnset, AccountBrowserRenderingGetLinksParamsCookiesSourceSchemeNonSecure, AccountBrowserRenderingGetLinksParamsCookiesSourceSchemeSecure:
		return true
	}
	return false
}

// Check [options](https://pptr.dev/api/puppeteer.gotooptions).
type AccountBrowserRenderingGetLinksParamsGotoOptions struct {
	Referer        param.Field[string]                                                         `json:"referer"`
	ReferrerPolicy param.Field[string]                                                         `json:"referrerPolicy"`
	Timeout        param.Field[float64]                                                        `json:"timeout"`
	WaitUntil      param.Field[AccountBrowserRenderingGetLinksParamsGotoOptionsWaitUntilUnion] `json:"waitUntil"`
}

func (r AccountBrowserRenderingGetLinksParamsGotoOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by [AccountBrowserRenderingGetLinksParamsGotoOptionsWaitUntilString],
// [AccountBrowserRenderingGetLinksParamsGotoOptionsWaitUntilArray].
type AccountBrowserRenderingGetLinksParamsGotoOptionsWaitUntilUnion interface {
	implementsAccountBrowserRenderingGetLinksParamsGotoOptionsWaitUntilUnion()
}

type AccountBrowserRenderingGetLinksParamsGotoOptionsWaitUntilString string

const (
	AccountBrowserRenderingGetLinksParamsGotoOptionsWaitUntilStringLoad             AccountBrowserRenderingGetLinksParamsGotoOptionsWaitUntilString = "load"
	AccountBrowserRenderingGetLinksParamsGotoOptionsWaitUntilStringDomcontentloaded AccountBrowserRenderingGetLinksParamsGotoOptionsWaitUntilString = "domcontentloaded"
	AccountBrowserRenderingGetLinksParamsGotoOptionsWaitUntilStringNetworkidle0     AccountBrowserRenderingGetLinksParamsGotoOptionsWaitUntilString = "networkidle0"
	AccountBrowserRenderingGetLinksParamsGotoOptionsWaitUntilStringNetworkidle2     AccountBrowserRenderingGetLinksParamsGotoOptionsWaitUntilString = "networkidle2"
)

func (r AccountBrowserRenderingGetLinksParamsGotoOptionsWaitUntilString) IsKnown() bool {
	switch r {
	case AccountBrowserRenderingGetLinksParamsGotoOptionsWaitUntilStringLoad, AccountBrowserRenderingGetLinksParamsGotoOptionsWaitUntilStringDomcontentloaded, AccountBrowserRenderingGetLinksParamsGotoOptionsWaitUntilStringNetworkidle0, AccountBrowserRenderingGetLinksParamsGotoOptionsWaitUntilStringNetworkidle2:
		return true
	}
	return false
}

func (r AccountBrowserRenderingGetLinksParamsGotoOptionsWaitUntilString) implementsAccountBrowserRenderingGetLinksParamsGotoOptionsWaitUntilUnion() {
}

type AccountBrowserRenderingGetLinksParamsGotoOptionsWaitUntilArray []AccountBrowserRenderingGetLinksParamsGotoOptionsWaitUntilArrayItem

func (r AccountBrowserRenderingGetLinksParamsGotoOptionsWaitUntilArray) implementsAccountBrowserRenderingGetLinksParamsGotoOptionsWaitUntilUnion() {
}

type AccountBrowserRenderingGetLinksParamsGotoOptionsWaitUntilArrayItem string

const (
	AccountBrowserRenderingGetLinksParamsGotoOptionsWaitUntilArrayItemLoad             AccountBrowserRenderingGetLinksParamsGotoOptionsWaitUntilArrayItem = "load"
	AccountBrowserRenderingGetLinksParamsGotoOptionsWaitUntilArrayItemDomcontentloaded AccountBrowserRenderingGetLinksParamsGotoOptionsWaitUntilArrayItem = "domcontentloaded"
	AccountBrowserRenderingGetLinksParamsGotoOptionsWaitUntilArrayItemNetworkidle0     AccountBrowserRenderingGetLinksParamsGotoOptionsWaitUntilArrayItem = "networkidle0"
	AccountBrowserRenderingGetLinksParamsGotoOptionsWaitUntilArrayItemNetworkidle2     AccountBrowserRenderingGetLinksParamsGotoOptionsWaitUntilArrayItem = "networkidle2"
)

func (r AccountBrowserRenderingGetLinksParamsGotoOptionsWaitUntilArrayItem) IsKnown() bool {
	switch r {
	case AccountBrowserRenderingGetLinksParamsGotoOptionsWaitUntilArrayItemLoad, AccountBrowserRenderingGetLinksParamsGotoOptionsWaitUntilArrayItemDomcontentloaded, AccountBrowserRenderingGetLinksParamsGotoOptionsWaitUntilArrayItemNetworkidle0, AccountBrowserRenderingGetLinksParamsGotoOptionsWaitUntilArrayItemNetworkidle2:
		return true
	}
	return false
}

type AccountBrowserRenderingGetLinksParamsRejectResourceType string

const (
	AccountBrowserRenderingGetLinksParamsRejectResourceTypeDocument           AccountBrowserRenderingGetLinksParamsRejectResourceType = "document"
	AccountBrowserRenderingGetLinksParamsRejectResourceTypeStylesheet         AccountBrowserRenderingGetLinksParamsRejectResourceType = "stylesheet"
	AccountBrowserRenderingGetLinksParamsRejectResourceTypeImage              AccountBrowserRenderingGetLinksParamsRejectResourceType = "image"
	AccountBrowserRenderingGetLinksParamsRejectResourceTypeMedia              AccountBrowserRenderingGetLinksParamsRejectResourceType = "media"
	AccountBrowserRenderingGetLinksParamsRejectResourceTypeFont               AccountBrowserRenderingGetLinksParamsRejectResourceType = "font"
	AccountBrowserRenderingGetLinksParamsRejectResourceTypeScript             AccountBrowserRenderingGetLinksParamsRejectResourceType = "script"
	AccountBrowserRenderingGetLinksParamsRejectResourceTypeTexttrack          AccountBrowserRenderingGetLinksParamsRejectResourceType = "texttrack"
	AccountBrowserRenderingGetLinksParamsRejectResourceTypeXhr                AccountBrowserRenderingGetLinksParamsRejectResourceType = "xhr"
	AccountBrowserRenderingGetLinksParamsRejectResourceTypeFetch              AccountBrowserRenderingGetLinksParamsRejectResourceType = "fetch"
	AccountBrowserRenderingGetLinksParamsRejectResourceTypePrefetch           AccountBrowserRenderingGetLinksParamsRejectResourceType = "prefetch"
	AccountBrowserRenderingGetLinksParamsRejectResourceTypeEventsource        AccountBrowserRenderingGetLinksParamsRejectResourceType = "eventsource"
	AccountBrowserRenderingGetLinksParamsRejectResourceTypeWebsocket          AccountBrowserRenderingGetLinksParamsRejectResourceType = "websocket"
	AccountBrowserRenderingGetLinksParamsRejectResourceTypeManifest           AccountBrowserRenderingGetLinksParamsRejectResourceType = "manifest"
	AccountBrowserRenderingGetLinksParamsRejectResourceTypeSignedexchange     AccountBrowserRenderingGetLinksParamsRejectResourceType = "signedexchange"
	AccountBrowserRenderingGetLinksParamsRejectResourceTypePing               AccountBrowserRenderingGetLinksParamsRejectResourceType = "ping"
	AccountBrowserRenderingGetLinksParamsRejectResourceTypeCspviolationreport AccountBrowserRenderingGetLinksParamsRejectResourceType = "cspviolationreport"
	AccountBrowserRenderingGetLinksParamsRejectResourceTypePreflight          AccountBrowserRenderingGetLinksParamsRejectResourceType = "preflight"
	AccountBrowserRenderingGetLinksParamsRejectResourceTypeOther              AccountBrowserRenderingGetLinksParamsRejectResourceType = "other"
)

func (r AccountBrowserRenderingGetLinksParamsRejectResourceType) IsKnown() bool {
	switch r {
	case AccountBrowserRenderingGetLinksParamsRejectResourceTypeDocument, AccountBrowserRenderingGetLinksParamsRejectResourceTypeStylesheet, AccountBrowserRenderingGetLinksParamsRejectResourceTypeImage, AccountBrowserRenderingGetLinksParamsRejectResourceTypeMedia, AccountBrowserRenderingGetLinksParamsRejectResourceTypeFont, AccountBrowserRenderingGetLinksParamsRejectResourceTypeScript, AccountBrowserRenderingGetLinksParamsRejectResourceTypeTexttrack, AccountBrowserRenderingGetLinksParamsRejectResourceTypeXhr, AccountBrowserRenderingGetLinksParamsRejectResourceTypeFetch, AccountBrowserRenderingGetLinksParamsRejectResourceTypePrefetch, AccountBrowserRenderingGetLinksParamsRejectResourceTypeEventsource, AccountBrowserRenderingGetLinksParamsRejectResourceTypeWebsocket, AccountBrowserRenderingGetLinksParamsRejectResourceTypeManifest, AccountBrowserRenderingGetLinksParamsRejectResourceTypeSignedexchange, AccountBrowserRenderingGetLinksParamsRejectResourceTypePing, AccountBrowserRenderingGetLinksParamsRejectResourceTypeCspviolationreport, AccountBrowserRenderingGetLinksParamsRejectResourceTypePreflight, AccountBrowserRenderingGetLinksParamsRejectResourceTypeOther:
		return true
	}
	return false
}

// Check [options](https://pptr.dev/api/puppeteer.page.setviewport).
type AccountBrowserRenderingGetLinksParamsViewport struct {
	Height            param.Field[float64] `json:"height,required"`
	Width             param.Field[float64] `json:"width,required"`
	DeviceScaleFactor param.Field[float64] `json:"deviceScaleFactor"`
	HasTouch          param.Field[bool]    `json:"hasTouch"`
	IsLandscape       param.Field[bool]    `json:"isLandscape"`
	IsMobile          param.Field[bool]    `json:"isMobile"`
}

func (r AccountBrowserRenderingGetLinksParamsViewport) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Wait for the selector to appear in page. Check
// [options](https://pptr.dev/api/puppeteer.page.waitforselector).
type AccountBrowserRenderingGetLinksParamsWaitForSelector struct {
	Selector param.Field[string]                                                      `json:"selector,required"`
	Hidden   param.Field[AccountBrowserRenderingGetLinksParamsWaitForSelectorHidden]  `json:"hidden"`
	Timeout  param.Field[float64]                                                     `json:"timeout"`
	Visible  param.Field[AccountBrowserRenderingGetLinksParamsWaitForSelectorVisible] `json:"visible"`
}

func (r AccountBrowserRenderingGetLinksParamsWaitForSelector) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountBrowserRenderingGetLinksParamsWaitForSelectorHidden bool

const (
	AccountBrowserRenderingGetLinksParamsWaitForSelectorHiddenTrue AccountBrowserRenderingGetLinksParamsWaitForSelectorHidden = true
)

func (r AccountBrowserRenderingGetLinksParamsWaitForSelectorHidden) IsKnown() bool {
	switch r {
	case AccountBrowserRenderingGetLinksParamsWaitForSelectorHiddenTrue:
		return true
	}
	return false
}

type AccountBrowserRenderingGetLinksParamsWaitForSelectorVisible bool

const (
	AccountBrowserRenderingGetLinksParamsWaitForSelectorVisibleTrue AccountBrowserRenderingGetLinksParamsWaitForSelectorVisible = true
)

func (r AccountBrowserRenderingGetLinksParamsWaitForSelectorVisible) IsKnown() bool {
	switch r {
	case AccountBrowserRenderingGetLinksParamsWaitForSelectorVisibleTrue:
		return true
	}
	return false
}

type AccountBrowserRenderingGetMarkdownParams struct {
	// Cache TTL default is 5s. Set to 0 to disable.
	CacheTtl param.Field[float64] `query:"cacheTTL"`
	// Adds a `<script>` tag into the page with the desired URL or content.
	AddScriptTag param.Field[[]AccountBrowserRenderingGetMarkdownParamsAddScriptTag] `json:"addScriptTag"`
	// Adds a `<link rel="stylesheet">` tag into the page with the desired URL or a
	// `<style type="text/css">` tag with the content.
	AddStyleTag param.Field[[]AccountBrowserRenderingGetMarkdownParamsAddStyleTag] `json:"addStyleTag"`
	// Only allow requests that match the provided regex patterns, eg. '/^.\*\.(css)'.
	AllowRequestPattern param.Field[[]string] `json:"allowRequestPattern"`
	// Only allow requests that match the provided resource types, eg. 'image' or
	// 'script'.
	AllowResourceTypes param.Field[[]AccountBrowserRenderingGetMarkdownParamsAllowResourceType] `json:"allowResourceTypes"`
	// Provide credentials for HTTP authentication.
	Authenticate param.Field[AccountBrowserRenderingGetMarkdownParamsAuthenticate] `json:"authenticate"`
	// Attempt to proceed when 'awaited' events fail or timeout.
	BestAttempt param.Field[bool] `json:"bestAttempt"`
	// Check [options](https://pptr.dev/api/puppeteer.page.setcookie).
	Cookies          param.Field[[]AccountBrowserRenderingGetMarkdownParamsCookie] `json:"cookies"`
	EmulateMediaType param.Field[string]                                           `json:"emulateMediaType"`
	// Check [options](https://pptr.dev/api/puppeteer.gotooptions).
	GotoOptions param.Field[AccountBrowserRenderingGetMarkdownParamsGotoOptions] `json:"gotoOptions"`
	// Set the content of the page, eg: `<h1>Hello World!!</h1>`. Either `html` or
	// `url` must be set.
	HTML param.Field[string] `json:"html"`
	// Block undesired requests that match the provided regex patterns, eg.
	// '/^.\*\.(css)'.
	RejectRequestPattern param.Field[[]string] `json:"rejectRequestPattern"`
	// Block undesired requests that match the provided resource types, eg. 'image' or
	// 'script'.
	RejectResourceTypes  param.Field[[]AccountBrowserRenderingGetMarkdownParamsRejectResourceType] `json:"rejectResourceTypes"`
	SetExtraHTTPHeaders  param.Field[map[string]string]                                            `json:"setExtraHTTPHeaders"`
	SetJavaScriptEnabled param.Field[bool]                                                         `json:"setJavaScriptEnabled"`
	// URL to navigate to, eg. `https://example.com`.
	URL       param.Field[string] `json:"url" format:"uri"`
	UserAgent param.Field[string] `json:"userAgent"`
	// Check [options](https://pptr.dev/api/puppeteer.page.setviewport).
	Viewport param.Field[AccountBrowserRenderingGetMarkdownParamsViewport] `json:"viewport"`
	// Wait for the selector to appear in page. Check
	// [options](https://pptr.dev/api/puppeteer.page.waitforselector).
	WaitForSelector param.Field[AccountBrowserRenderingGetMarkdownParamsWaitForSelector] `json:"waitForSelector"`
	// Waits for a specified timeout before continuing.
	WaitForTimeout param.Field[float64] `json:"waitForTimeout"`
}

func (r AccountBrowserRenderingGetMarkdownParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [AccountBrowserRenderingGetMarkdownParams]'s query
// parameters as `url.Values`.
func (r AccountBrowserRenderingGetMarkdownParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountBrowserRenderingGetMarkdownParamsAddScriptTag struct {
	ID      param.Field[string] `json:"id"`
	Content param.Field[string] `json:"content"`
	Type    param.Field[string] `json:"type"`
	URL     param.Field[string] `json:"url"`
}

func (r AccountBrowserRenderingGetMarkdownParamsAddScriptTag) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountBrowserRenderingGetMarkdownParamsAddStyleTag struct {
	Content param.Field[string] `json:"content"`
	URL     param.Field[string] `json:"url"`
}

func (r AccountBrowserRenderingGetMarkdownParamsAddStyleTag) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountBrowserRenderingGetMarkdownParamsAllowResourceType string

const (
	AccountBrowserRenderingGetMarkdownParamsAllowResourceTypeDocument           AccountBrowserRenderingGetMarkdownParamsAllowResourceType = "document"
	AccountBrowserRenderingGetMarkdownParamsAllowResourceTypeStylesheet         AccountBrowserRenderingGetMarkdownParamsAllowResourceType = "stylesheet"
	AccountBrowserRenderingGetMarkdownParamsAllowResourceTypeImage              AccountBrowserRenderingGetMarkdownParamsAllowResourceType = "image"
	AccountBrowserRenderingGetMarkdownParamsAllowResourceTypeMedia              AccountBrowserRenderingGetMarkdownParamsAllowResourceType = "media"
	AccountBrowserRenderingGetMarkdownParamsAllowResourceTypeFont               AccountBrowserRenderingGetMarkdownParamsAllowResourceType = "font"
	AccountBrowserRenderingGetMarkdownParamsAllowResourceTypeScript             AccountBrowserRenderingGetMarkdownParamsAllowResourceType = "script"
	AccountBrowserRenderingGetMarkdownParamsAllowResourceTypeTexttrack          AccountBrowserRenderingGetMarkdownParamsAllowResourceType = "texttrack"
	AccountBrowserRenderingGetMarkdownParamsAllowResourceTypeXhr                AccountBrowserRenderingGetMarkdownParamsAllowResourceType = "xhr"
	AccountBrowserRenderingGetMarkdownParamsAllowResourceTypeFetch              AccountBrowserRenderingGetMarkdownParamsAllowResourceType = "fetch"
	AccountBrowserRenderingGetMarkdownParamsAllowResourceTypePrefetch           AccountBrowserRenderingGetMarkdownParamsAllowResourceType = "prefetch"
	AccountBrowserRenderingGetMarkdownParamsAllowResourceTypeEventsource        AccountBrowserRenderingGetMarkdownParamsAllowResourceType = "eventsource"
	AccountBrowserRenderingGetMarkdownParamsAllowResourceTypeWebsocket          AccountBrowserRenderingGetMarkdownParamsAllowResourceType = "websocket"
	AccountBrowserRenderingGetMarkdownParamsAllowResourceTypeManifest           AccountBrowserRenderingGetMarkdownParamsAllowResourceType = "manifest"
	AccountBrowserRenderingGetMarkdownParamsAllowResourceTypeSignedexchange     AccountBrowserRenderingGetMarkdownParamsAllowResourceType = "signedexchange"
	AccountBrowserRenderingGetMarkdownParamsAllowResourceTypePing               AccountBrowserRenderingGetMarkdownParamsAllowResourceType = "ping"
	AccountBrowserRenderingGetMarkdownParamsAllowResourceTypeCspviolationreport AccountBrowserRenderingGetMarkdownParamsAllowResourceType = "cspviolationreport"
	AccountBrowserRenderingGetMarkdownParamsAllowResourceTypePreflight          AccountBrowserRenderingGetMarkdownParamsAllowResourceType = "preflight"
	AccountBrowserRenderingGetMarkdownParamsAllowResourceTypeOther              AccountBrowserRenderingGetMarkdownParamsAllowResourceType = "other"
)

func (r AccountBrowserRenderingGetMarkdownParamsAllowResourceType) IsKnown() bool {
	switch r {
	case AccountBrowserRenderingGetMarkdownParamsAllowResourceTypeDocument, AccountBrowserRenderingGetMarkdownParamsAllowResourceTypeStylesheet, AccountBrowserRenderingGetMarkdownParamsAllowResourceTypeImage, AccountBrowserRenderingGetMarkdownParamsAllowResourceTypeMedia, AccountBrowserRenderingGetMarkdownParamsAllowResourceTypeFont, AccountBrowserRenderingGetMarkdownParamsAllowResourceTypeScript, AccountBrowserRenderingGetMarkdownParamsAllowResourceTypeTexttrack, AccountBrowserRenderingGetMarkdownParamsAllowResourceTypeXhr, AccountBrowserRenderingGetMarkdownParamsAllowResourceTypeFetch, AccountBrowserRenderingGetMarkdownParamsAllowResourceTypePrefetch, AccountBrowserRenderingGetMarkdownParamsAllowResourceTypeEventsource, AccountBrowserRenderingGetMarkdownParamsAllowResourceTypeWebsocket, AccountBrowserRenderingGetMarkdownParamsAllowResourceTypeManifest, AccountBrowserRenderingGetMarkdownParamsAllowResourceTypeSignedexchange, AccountBrowserRenderingGetMarkdownParamsAllowResourceTypePing, AccountBrowserRenderingGetMarkdownParamsAllowResourceTypeCspviolationreport, AccountBrowserRenderingGetMarkdownParamsAllowResourceTypePreflight, AccountBrowserRenderingGetMarkdownParamsAllowResourceTypeOther:
		return true
	}
	return false
}

// Provide credentials for HTTP authentication.
type AccountBrowserRenderingGetMarkdownParamsAuthenticate struct {
	Password param.Field[string] `json:"password,required"`
	Username param.Field[string] `json:"username,required"`
}

func (r AccountBrowserRenderingGetMarkdownParamsAuthenticate) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountBrowserRenderingGetMarkdownParamsCookie struct {
	Name         param.Field[string]                                                      `json:"name,required"`
	Value        param.Field[string]                                                      `json:"value,required"`
	Domain       param.Field[string]                                                      `json:"domain"`
	Expires      param.Field[float64]                                                     `json:"expires"`
	HTTPOnly     param.Field[bool]                                                        `json:"httpOnly"`
	PartitionKey param.Field[string]                                                      `json:"partitionKey"`
	Path         param.Field[string]                                                      `json:"path"`
	Priority     param.Field[AccountBrowserRenderingGetMarkdownParamsCookiesPriority]     `json:"priority"`
	SameParty    param.Field[bool]                                                        `json:"sameParty"`
	SameSite     param.Field[AccountBrowserRenderingGetMarkdownParamsCookiesSameSite]     `json:"sameSite"`
	Secure       param.Field[bool]                                                        `json:"secure"`
	SourcePort   param.Field[float64]                                                     `json:"sourcePort"`
	SourceScheme param.Field[AccountBrowserRenderingGetMarkdownParamsCookiesSourceScheme] `json:"sourceScheme"`
	URL          param.Field[string]                                                      `json:"url"`
}

func (r AccountBrowserRenderingGetMarkdownParamsCookie) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountBrowserRenderingGetMarkdownParamsCookiesPriority string

const (
	AccountBrowserRenderingGetMarkdownParamsCookiesPriorityLow    AccountBrowserRenderingGetMarkdownParamsCookiesPriority = "Low"
	AccountBrowserRenderingGetMarkdownParamsCookiesPriorityMedium AccountBrowserRenderingGetMarkdownParamsCookiesPriority = "Medium"
	AccountBrowserRenderingGetMarkdownParamsCookiesPriorityHigh   AccountBrowserRenderingGetMarkdownParamsCookiesPriority = "High"
)

func (r AccountBrowserRenderingGetMarkdownParamsCookiesPriority) IsKnown() bool {
	switch r {
	case AccountBrowserRenderingGetMarkdownParamsCookiesPriorityLow, AccountBrowserRenderingGetMarkdownParamsCookiesPriorityMedium, AccountBrowserRenderingGetMarkdownParamsCookiesPriorityHigh:
		return true
	}
	return false
}

type AccountBrowserRenderingGetMarkdownParamsCookiesSameSite string

const (
	AccountBrowserRenderingGetMarkdownParamsCookiesSameSiteStrict AccountBrowserRenderingGetMarkdownParamsCookiesSameSite = "Strict"
	AccountBrowserRenderingGetMarkdownParamsCookiesSameSiteLax    AccountBrowserRenderingGetMarkdownParamsCookiesSameSite = "Lax"
	AccountBrowserRenderingGetMarkdownParamsCookiesSameSiteNone   AccountBrowserRenderingGetMarkdownParamsCookiesSameSite = "None"
)

func (r AccountBrowserRenderingGetMarkdownParamsCookiesSameSite) IsKnown() bool {
	switch r {
	case AccountBrowserRenderingGetMarkdownParamsCookiesSameSiteStrict, AccountBrowserRenderingGetMarkdownParamsCookiesSameSiteLax, AccountBrowserRenderingGetMarkdownParamsCookiesSameSiteNone:
		return true
	}
	return false
}

type AccountBrowserRenderingGetMarkdownParamsCookiesSourceScheme string

const (
	AccountBrowserRenderingGetMarkdownParamsCookiesSourceSchemeUnset     AccountBrowserRenderingGetMarkdownParamsCookiesSourceScheme = "Unset"
	AccountBrowserRenderingGetMarkdownParamsCookiesSourceSchemeNonSecure AccountBrowserRenderingGetMarkdownParamsCookiesSourceScheme = "NonSecure"
	AccountBrowserRenderingGetMarkdownParamsCookiesSourceSchemeSecure    AccountBrowserRenderingGetMarkdownParamsCookiesSourceScheme = "Secure"
)

func (r AccountBrowserRenderingGetMarkdownParamsCookiesSourceScheme) IsKnown() bool {
	switch r {
	case AccountBrowserRenderingGetMarkdownParamsCookiesSourceSchemeUnset, AccountBrowserRenderingGetMarkdownParamsCookiesSourceSchemeNonSecure, AccountBrowserRenderingGetMarkdownParamsCookiesSourceSchemeSecure:
		return true
	}
	return false
}

// Check [options](https://pptr.dev/api/puppeteer.gotooptions).
type AccountBrowserRenderingGetMarkdownParamsGotoOptions struct {
	Referer        param.Field[string]                                                            `json:"referer"`
	ReferrerPolicy param.Field[string]                                                            `json:"referrerPolicy"`
	Timeout        param.Field[float64]                                                           `json:"timeout"`
	WaitUntil      param.Field[AccountBrowserRenderingGetMarkdownParamsGotoOptionsWaitUntilUnion] `json:"waitUntil"`
}

func (r AccountBrowserRenderingGetMarkdownParamsGotoOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by
// [AccountBrowserRenderingGetMarkdownParamsGotoOptionsWaitUntilString],
// [AccountBrowserRenderingGetMarkdownParamsGotoOptionsWaitUntilArray].
type AccountBrowserRenderingGetMarkdownParamsGotoOptionsWaitUntilUnion interface {
	implementsAccountBrowserRenderingGetMarkdownParamsGotoOptionsWaitUntilUnion()
}

type AccountBrowserRenderingGetMarkdownParamsGotoOptionsWaitUntilString string

const (
	AccountBrowserRenderingGetMarkdownParamsGotoOptionsWaitUntilStringLoad             AccountBrowserRenderingGetMarkdownParamsGotoOptionsWaitUntilString = "load"
	AccountBrowserRenderingGetMarkdownParamsGotoOptionsWaitUntilStringDomcontentloaded AccountBrowserRenderingGetMarkdownParamsGotoOptionsWaitUntilString = "domcontentloaded"
	AccountBrowserRenderingGetMarkdownParamsGotoOptionsWaitUntilStringNetworkidle0     AccountBrowserRenderingGetMarkdownParamsGotoOptionsWaitUntilString = "networkidle0"
	AccountBrowserRenderingGetMarkdownParamsGotoOptionsWaitUntilStringNetworkidle2     AccountBrowserRenderingGetMarkdownParamsGotoOptionsWaitUntilString = "networkidle2"
)

func (r AccountBrowserRenderingGetMarkdownParamsGotoOptionsWaitUntilString) IsKnown() bool {
	switch r {
	case AccountBrowserRenderingGetMarkdownParamsGotoOptionsWaitUntilStringLoad, AccountBrowserRenderingGetMarkdownParamsGotoOptionsWaitUntilStringDomcontentloaded, AccountBrowserRenderingGetMarkdownParamsGotoOptionsWaitUntilStringNetworkidle0, AccountBrowserRenderingGetMarkdownParamsGotoOptionsWaitUntilStringNetworkidle2:
		return true
	}
	return false
}

func (r AccountBrowserRenderingGetMarkdownParamsGotoOptionsWaitUntilString) implementsAccountBrowserRenderingGetMarkdownParamsGotoOptionsWaitUntilUnion() {
}

type AccountBrowserRenderingGetMarkdownParamsGotoOptionsWaitUntilArray []AccountBrowserRenderingGetMarkdownParamsGotoOptionsWaitUntilArrayItem

func (r AccountBrowserRenderingGetMarkdownParamsGotoOptionsWaitUntilArray) implementsAccountBrowserRenderingGetMarkdownParamsGotoOptionsWaitUntilUnion() {
}

type AccountBrowserRenderingGetMarkdownParamsGotoOptionsWaitUntilArrayItem string

const (
	AccountBrowserRenderingGetMarkdownParamsGotoOptionsWaitUntilArrayItemLoad             AccountBrowserRenderingGetMarkdownParamsGotoOptionsWaitUntilArrayItem = "load"
	AccountBrowserRenderingGetMarkdownParamsGotoOptionsWaitUntilArrayItemDomcontentloaded AccountBrowserRenderingGetMarkdownParamsGotoOptionsWaitUntilArrayItem = "domcontentloaded"
	AccountBrowserRenderingGetMarkdownParamsGotoOptionsWaitUntilArrayItemNetworkidle0     AccountBrowserRenderingGetMarkdownParamsGotoOptionsWaitUntilArrayItem = "networkidle0"
	AccountBrowserRenderingGetMarkdownParamsGotoOptionsWaitUntilArrayItemNetworkidle2     AccountBrowserRenderingGetMarkdownParamsGotoOptionsWaitUntilArrayItem = "networkidle2"
)

func (r AccountBrowserRenderingGetMarkdownParamsGotoOptionsWaitUntilArrayItem) IsKnown() bool {
	switch r {
	case AccountBrowserRenderingGetMarkdownParamsGotoOptionsWaitUntilArrayItemLoad, AccountBrowserRenderingGetMarkdownParamsGotoOptionsWaitUntilArrayItemDomcontentloaded, AccountBrowserRenderingGetMarkdownParamsGotoOptionsWaitUntilArrayItemNetworkidle0, AccountBrowserRenderingGetMarkdownParamsGotoOptionsWaitUntilArrayItemNetworkidle2:
		return true
	}
	return false
}

type AccountBrowserRenderingGetMarkdownParamsRejectResourceType string

const (
	AccountBrowserRenderingGetMarkdownParamsRejectResourceTypeDocument           AccountBrowserRenderingGetMarkdownParamsRejectResourceType = "document"
	AccountBrowserRenderingGetMarkdownParamsRejectResourceTypeStylesheet         AccountBrowserRenderingGetMarkdownParamsRejectResourceType = "stylesheet"
	AccountBrowserRenderingGetMarkdownParamsRejectResourceTypeImage              AccountBrowserRenderingGetMarkdownParamsRejectResourceType = "image"
	AccountBrowserRenderingGetMarkdownParamsRejectResourceTypeMedia              AccountBrowserRenderingGetMarkdownParamsRejectResourceType = "media"
	AccountBrowserRenderingGetMarkdownParamsRejectResourceTypeFont               AccountBrowserRenderingGetMarkdownParamsRejectResourceType = "font"
	AccountBrowserRenderingGetMarkdownParamsRejectResourceTypeScript             AccountBrowserRenderingGetMarkdownParamsRejectResourceType = "script"
	AccountBrowserRenderingGetMarkdownParamsRejectResourceTypeTexttrack          AccountBrowserRenderingGetMarkdownParamsRejectResourceType = "texttrack"
	AccountBrowserRenderingGetMarkdownParamsRejectResourceTypeXhr                AccountBrowserRenderingGetMarkdownParamsRejectResourceType = "xhr"
	AccountBrowserRenderingGetMarkdownParamsRejectResourceTypeFetch              AccountBrowserRenderingGetMarkdownParamsRejectResourceType = "fetch"
	AccountBrowserRenderingGetMarkdownParamsRejectResourceTypePrefetch           AccountBrowserRenderingGetMarkdownParamsRejectResourceType = "prefetch"
	AccountBrowserRenderingGetMarkdownParamsRejectResourceTypeEventsource        AccountBrowserRenderingGetMarkdownParamsRejectResourceType = "eventsource"
	AccountBrowserRenderingGetMarkdownParamsRejectResourceTypeWebsocket          AccountBrowserRenderingGetMarkdownParamsRejectResourceType = "websocket"
	AccountBrowserRenderingGetMarkdownParamsRejectResourceTypeManifest           AccountBrowserRenderingGetMarkdownParamsRejectResourceType = "manifest"
	AccountBrowserRenderingGetMarkdownParamsRejectResourceTypeSignedexchange     AccountBrowserRenderingGetMarkdownParamsRejectResourceType = "signedexchange"
	AccountBrowserRenderingGetMarkdownParamsRejectResourceTypePing               AccountBrowserRenderingGetMarkdownParamsRejectResourceType = "ping"
	AccountBrowserRenderingGetMarkdownParamsRejectResourceTypeCspviolationreport AccountBrowserRenderingGetMarkdownParamsRejectResourceType = "cspviolationreport"
	AccountBrowserRenderingGetMarkdownParamsRejectResourceTypePreflight          AccountBrowserRenderingGetMarkdownParamsRejectResourceType = "preflight"
	AccountBrowserRenderingGetMarkdownParamsRejectResourceTypeOther              AccountBrowserRenderingGetMarkdownParamsRejectResourceType = "other"
)

func (r AccountBrowserRenderingGetMarkdownParamsRejectResourceType) IsKnown() bool {
	switch r {
	case AccountBrowserRenderingGetMarkdownParamsRejectResourceTypeDocument, AccountBrowserRenderingGetMarkdownParamsRejectResourceTypeStylesheet, AccountBrowserRenderingGetMarkdownParamsRejectResourceTypeImage, AccountBrowserRenderingGetMarkdownParamsRejectResourceTypeMedia, AccountBrowserRenderingGetMarkdownParamsRejectResourceTypeFont, AccountBrowserRenderingGetMarkdownParamsRejectResourceTypeScript, AccountBrowserRenderingGetMarkdownParamsRejectResourceTypeTexttrack, AccountBrowserRenderingGetMarkdownParamsRejectResourceTypeXhr, AccountBrowserRenderingGetMarkdownParamsRejectResourceTypeFetch, AccountBrowserRenderingGetMarkdownParamsRejectResourceTypePrefetch, AccountBrowserRenderingGetMarkdownParamsRejectResourceTypeEventsource, AccountBrowserRenderingGetMarkdownParamsRejectResourceTypeWebsocket, AccountBrowserRenderingGetMarkdownParamsRejectResourceTypeManifest, AccountBrowserRenderingGetMarkdownParamsRejectResourceTypeSignedexchange, AccountBrowserRenderingGetMarkdownParamsRejectResourceTypePing, AccountBrowserRenderingGetMarkdownParamsRejectResourceTypeCspviolationreport, AccountBrowserRenderingGetMarkdownParamsRejectResourceTypePreflight, AccountBrowserRenderingGetMarkdownParamsRejectResourceTypeOther:
		return true
	}
	return false
}

// Check [options](https://pptr.dev/api/puppeteer.page.setviewport).
type AccountBrowserRenderingGetMarkdownParamsViewport struct {
	Height            param.Field[float64] `json:"height,required"`
	Width             param.Field[float64] `json:"width,required"`
	DeviceScaleFactor param.Field[float64] `json:"deviceScaleFactor"`
	HasTouch          param.Field[bool]    `json:"hasTouch"`
	IsLandscape       param.Field[bool]    `json:"isLandscape"`
	IsMobile          param.Field[bool]    `json:"isMobile"`
}

func (r AccountBrowserRenderingGetMarkdownParamsViewport) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Wait for the selector to appear in page. Check
// [options](https://pptr.dev/api/puppeteer.page.waitforselector).
type AccountBrowserRenderingGetMarkdownParamsWaitForSelector struct {
	Selector param.Field[string]                                                         `json:"selector,required"`
	Hidden   param.Field[AccountBrowserRenderingGetMarkdownParamsWaitForSelectorHidden]  `json:"hidden"`
	Timeout  param.Field[float64]                                                        `json:"timeout"`
	Visible  param.Field[AccountBrowserRenderingGetMarkdownParamsWaitForSelectorVisible] `json:"visible"`
}

func (r AccountBrowserRenderingGetMarkdownParamsWaitForSelector) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountBrowserRenderingGetMarkdownParamsWaitForSelectorHidden bool

const (
	AccountBrowserRenderingGetMarkdownParamsWaitForSelectorHiddenTrue AccountBrowserRenderingGetMarkdownParamsWaitForSelectorHidden = true
)

func (r AccountBrowserRenderingGetMarkdownParamsWaitForSelectorHidden) IsKnown() bool {
	switch r {
	case AccountBrowserRenderingGetMarkdownParamsWaitForSelectorHiddenTrue:
		return true
	}
	return false
}

type AccountBrowserRenderingGetMarkdownParamsWaitForSelectorVisible bool

const (
	AccountBrowserRenderingGetMarkdownParamsWaitForSelectorVisibleTrue AccountBrowserRenderingGetMarkdownParamsWaitForSelectorVisible = true
)

func (r AccountBrowserRenderingGetMarkdownParamsWaitForSelectorVisible) IsKnown() bool {
	switch r {
	case AccountBrowserRenderingGetMarkdownParamsWaitForSelectorVisibleTrue:
		return true
	}
	return false
}

type AccountBrowserRenderingGetPdfParams struct {
	// Cache TTL default is 5s. Set to 0 to disable.
	CacheTtl param.Field[float64] `query:"cacheTTL"`
	// Adds a `<script>` tag into the page with the desired URL or content.
	AddScriptTag param.Field[[]AccountBrowserRenderingGetPdfParamsAddScriptTag] `json:"addScriptTag"`
	// Adds a `<link rel="stylesheet">` tag into the page with the desired URL or a
	// `<style type="text/css">` tag with the content.
	AddStyleTag param.Field[[]AccountBrowserRenderingGetPdfParamsAddStyleTag] `json:"addStyleTag"`
	// Only allow requests that match the provided regex patterns, eg. '/^.\*\.(css)'.
	AllowRequestPattern param.Field[[]string] `json:"allowRequestPattern"`
	// Only allow requests that match the provided resource types, eg. 'image' or
	// 'script'.
	AllowResourceTypes param.Field[[]AccountBrowserRenderingGetPdfParamsAllowResourceType] `json:"allowResourceTypes"`
	// Provide credentials for HTTP authentication.
	Authenticate param.Field[AccountBrowserRenderingGetPdfParamsAuthenticate] `json:"authenticate"`
	// Attempt to proceed when 'awaited' events fail or timeout.
	BestAttempt param.Field[bool] `json:"bestAttempt"`
	// Check [options](https://pptr.dev/api/puppeteer.page.setcookie).
	Cookies          param.Field[[]AccountBrowserRenderingGetPdfParamsCookie] `json:"cookies"`
	EmulateMediaType param.Field[string]                                      `json:"emulateMediaType"`
	// Check [options](https://pptr.dev/api/puppeteer.gotooptions).
	GotoOptions param.Field[AccountBrowserRenderingGetPdfParamsGotoOptions] `json:"gotoOptions"`
	// Set the content of the page, eg: `<h1>Hello World!!</h1>`. Either `html` or
	// `url` must be set.
	HTML param.Field[string] `json:"html"`
	// Block undesired requests that match the provided regex patterns, eg.
	// '/^.\*\.(css)'.
	RejectRequestPattern param.Field[[]string] `json:"rejectRequestPattern"`
	// Block undesired requests that match the provided resource types, eg. 'image' or
	// 'script'.
	RejectResourceTypes  param.Field[[]AccountBrowserRenderingGetPdfParamsRejectResourceType] `json:"rejectResourceTypes"`
	SetExtraHTTPHeaders  param.Field[map[string]string]                                       `json:"setExtraHTTPHeaders"`
	SetJavaScriptEnabled param.Field[bool]                                                    `json:"setJavaScriptEnabled"`
	// URL to navigate to, eg. `https://example.com`.
	URL       param.Field[string] `json:"url" format:"uri"`
	UserAgent param.Field[string] `json:"userAgent"`
	// Check [options](https://pptr.dev/api/puppeteer.page.setviewport).
	Viewport param.Field[AccountBrowserRenderingGetPdfParamsViewport] `json:"viewport"`
	// Wait for the selector to appear in page. Check
	// [options](https://pptr.dev/api/puppeteer.page.waitforselector).
	WaitForSelector param.Field[AccountBrowserRenderingGetPdfParamsWaitForSelector] `json:"waitForSelector"`
	// Waits for a specified timeout before continuing.
	WaitForTimeout param.Field[float64] `json:"waitForTimeout"`
}

func (r AccountBrowserRenderingGetPdfParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [AccountBrowserRenderingGetPdfParams]'s query parameters as
// `url.Values`.
func (r AccountBrowserRenderingGetPdfParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountBrowserRenderingGetPdfParamsAddScriptTag struct {
	ID      param.Field[string] `json:"id"`
	Content param.Field[string] `json:"content"`
	Type    param.Field[string] `json:"type"`
	URL     param.Field[string] `json:"url"`
}

func (r AccountBrowserRenderingGetPdfParamsAddScriptTag) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountBrowserRenderingGetPdfParamsAddStyleTag struct {
	Content param.Field[string] `json:"content"`
	URL     param.Field[string] `json:"url"`
}

func (r AccountBrowserRenderingGetPdfParamsAddStyleTag) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountBrowserRenderingGetPdfParamsAllowResourceType string

const (
	AccountBrowserRenderingGetPdfParamsAllowResourceTypeDocument           AccountBrowserRenderingGetPdfParamsAllowResourceType = "document"
	AccountBrowserRenderingGetPdfParamsAllowResourceTypeStylesheet         AccountBrowserRenderingGetPdfParamsAllowResourceType = "stylesheet"
	AccountBrowserRenderingGetPdfParamsAllowResourceTypeImage              AccountBrowserRenderingGetPdfParamsAllowResourceType = "image"
	AccountBrowserRenderingGetPdfParamsAllowResourceTypeMedia              AccountBrowserRenderingGetPdfParamsAllowResourceType = "media"
	AccountBrowserRenderingGetPdfParamsAllowResourceTypeFont               AccountBrowserRenderingGetPdfParamsAllowResourceType = "font"
	AccountBrowserRenderingGetPdfParamsAllowResourceTypeScript             AccountBrowserRenderingGetPdfParamsAllowResourceType = "script"
	AccountBrowserRenderingGetPdfParamsAllowResourceTypeTexttrack          AccountBrowserRenderingGetPdfParamsAllowResourceType = "texttrack"
	AccountBrowserRenderingGetPdfParamsAllowResourceTypeXhr                AccountBrowserRenderingGetPdfParamsAllowResourceType = "xhr"
	AccountBrowserRenderingGetPdfParamsAllowResourceTypeFetch              AccountBrowserRenderingGetPdfParamsAllowResourceType = "fetch"
	AccountBrowserRenderingGetPdfParamsAllowResourceTypePrefetch           AccountBrowserRenderingGetPdfParamsAllowResourceType = "prefetch"
	AccountBrowserRenderingGetPdfParamsAllowResourceTypeEventsource        AccountBrowserRenderingGetPdfParamsAllowResourceType = "eventsource"
	AccountBrowserRenderingGetPdfParamsAllowResourceTypeWebsocket          AccountBrowserRenderingGetPdfParamsAllowResourceType = "websocket"
	AccountBrowserRenderingGetPdfParamsAllowResourceTypeManifest           AccountBrowserRenderingGetPdfParamsAllowResourceType = "manifest"
	AccountBrowserRenderingGetPdfParamsAllowResourceTypeSignedexchange     AccountBrowserRenderingGetPdfParamsAllowResourceType = "signedexchange"
	AccountBrowserRenderingGetPdfParamsAllowResourceTypePing               AccountBrowserRenderingGetPdfParamsAllowResourceType = "ping"
	AccountBrowserRenderingGetPdfParamsAllowResourceTypeCspviolationreport AccountBrowserRenderingGetPdfParamsAllowResourceType = "cspviolationreport"
	AccountBrowserRenderingGetPdfParamsAllowResourceTypePreflight          AccountBrowserRenderingGetPdfParamsAllowResourceType = "preflight"
	AccountBrowserRenderingGetPdfParamsAllowResourceTypeOther              AccountBrowserRenderingGetPdfParamsAllowResourceType = "other"
)

func (r AccountBrowserRenderingGetPdfParamsAllowResourceType) IsKnown() bool {
	switch r {
	case AccountBrowserRenderingGetPdfParamsAllowResourceTypeDocument, AccountBrowserRenderingGetPdfParamsAllowResourceTypeStylesheet, AccountBrowserRenderingGetPdfParamsAllowResourceTypeImage, AccountBrowserRenderingGetPdfParamsAllowResourceTypeMedia, AccountBrowserRenderingGetPdfParamsAllowResourceTypeFont, AccountBrowserRenderingGetPdfParamsAllowResourceTypeScript, AccountBrowserRenderingGetPdfParamsAllowResourceTypeTexttrack, AccountBrowserRenderingGetPdfParamsAllowResourceTypeXhr, AccountBrowserRenderingGetPdfParamsAllowResourceTypeFetch, AccountBrowserRenderingGetPdfParamsAllowResourceTypePrefetch, AccountBrowserRenderingGetPdfParamsAllowResourceTypeEventsource, AccountBrowserRenderingGetPdfParamsAllowResourceTypeWebsocket, AccountBrowserRenderingGetPdfParamsAllowResourceTypeManifest, AccountBrowserRenderingGetPdfParamsAllowResourceTypeSignedexchange, AccountBrowserRenderingGetPdfParamsAllowResourceTypePing, AccountBrowserRenderingGetPdfParamsAllowResourceTypeCspviolationreport, AccountBrowserRenderingGetPdfParamsAllowResourceTypePreflight, AccountBrowserRenderingGetPdfParamsAllowResourceTypeOther:
		return true
	}
	return false
}

// Provide credentials for HTTP authentication.
type AccountBrowserRenderingGetPdfParamsAuthenticate struct {
	Password param.Field[string] `json:"password,required"`
	Username param.Field[string] `json:"username,required"`
}

func (r AccountBrowserRenderingGetPdfParamsAuthenticate) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountBrowserRenderingGetPdfParamsCookie struct {
	Name         param.Field[string]                                                 `json:"name,required"`
	Value        param.Field[string]                                                 `json:"value,required"`
	Domain       param.Field[string]                                                 `json:"domain"`
	Expires      param.Field[float64]                                                `json:"expires"`
	HTTPOnly     param.Field[bool]                                                   `json:"httpOnly"`
	PartitionKey param.Field[string]                                                 `json:"partitionKey"`
	Path         param.Field[string]                                                 `json:"path"`
	Priority     param.Field[AccountBrowserRenderingGetPdfParamsCookiesPriority]     `json:"priority"`
	SameParty    param.Field[bool]                                                   `json:"sameParty"`
	SameSite     param.Field[AccountBrowserRenderingGetPdfParamsCookiesSameSite]     `json:"sameSite"`
	Secure       param.Field[bool]                                                   `json:"secure"`
	SourcePort   param.Field[float64]                                                `json:"sourcePort"`
	SourceScheme param.Field[AccountBrowserRenderingGetPdfParamsCookiesSourceScheme] `json:"sourceScheme"`
	URL          param.Field[string]                                                 `json:"url"`
}

func (r AccountBrowserRenderingGetPdfParamsCookie) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountBrowserRenderingGetPdfParamsCookiesPriority string

const (
	AccountBrowserRenderingGetPdfParamsCookiesPriorityLow    AccountBrowserRenderingGetPdfParamsCookiesPriority = "Low"
	AccountBrowserRenderingGetPdfParamsCookiesPriorityMedium AccountBrowserRenderingGetPdfParamsCookiesPriority = "Medium"
	AccountBrowserRenderingGetPdfParamsCookiesPriorityHigh   AccountBrowserRenderingGetPdfParamsCookiesPriority = "High"
)

func (r AccountBrowserRenderingGetPdfParamsCookiesPriority) IsKnown() bool {
	switch r {
	case AccountBrowserRenderingGetPdfParamsCookiesPriorityLow, AccountBrowserRenderingGetPdfParamsCookiesPriorityMedium, AccountBrowserRenderingGetPdfParamsCookiesPriorityHigh:
		return true
	}
	return false
}

type AccountBrowserRenderingGetPdfParamsCookiesSameSite string

const (
	AccountBrowserRenderingGetPdfParamsCookiesSameSiteStrict AccountBrowserRenderingGetPdfParamsCookiesSameSite = "Strict"
	AccountBrowserRenderingGetPdfParamsCookiesSameSiteLax    AccountBrowserRenderingGetPdfParamsCookiesSameSite = "Lax"
	AccountBrowserRenderingGetPdfParamsCookiesSameSiteNone   AccountBrowserRenderingGetPdfParamsCookiesSameSite = "None"
)

func (r AccountBrowserRenderingGetPdfParamsCookiesSameSite) IsKnown() bool {
	switch r {
	case AccountBrowserRenderingGetPdfParamsCookiesSameSiteStrict, AccountBrowserRenderingGetPdfParamsCookiesSameSiteLax, AccountBrowserRenderingGetPdfParamsCookiesSameSiteNone:
		return true
	}
	return false
}

type AccountBrowserRenderingGetPdfParamsCookiesSourceScheme string

const (
	AccountBrowserRenderingGetPdfParamsCookiesSourceSchemeUnset     AccountBrowserRenderingGetPdfParamsCookiesSourceScheme = "Unset"
	AccountBrowserRenderingGetPdfParamsCookiesSourceSchemeNonSecure AccountBrowserRenderingGetPdfParamsCookiesSourceScheme = "NonSecure"
	AccountBrowserRenderingGetPdfParamsCookiesSourceSchemeSecure    AccountBrowserRenderingGetPdfParamsCookiesSourceScheme = "Secure"
)

func (r AccountBrowserRenderingGetPdfParamsCookiesSourceScheme) IsKnown() bool {
	switch r {
	case AccountBrowserRenderingGetPdfParamsCookiesSourceSchemeUnset, AccountBrowserRenderingGetPdfParamsCookiesSourceSchemeNonSecure, AccountBrowserRenderingGetPdfParamsCookiesSourceSchemeSecure:
		return true
	}
	return false
}

// Check [options](https://pptr.dev/api/puppeteer.gotooptions).
type AccountBrowserRenderingGetPdfParamsGotoOptions struct {
	Referer        param.Field[string]                                                       `json:"referer"`
	ReferrerPolicy param.Field[string]                                                       `json:"referrerPolicy"`
	Timeout        param.Field[float64]                                                      `json:"timeout"`
	WaitUntil      param.Field[AccountBrowserRenderingGetPdfParamsGotoOptionsWaitUntilUnion] `json:"waitUntil"`
}

func (r AccountBrowserRenderingGetPdfParamsGotoOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by [AccountBrowserRenderingGetPdfParamsGotoOptionsWaitUntilString],
// [AccountBrowserRenderingGetPdfParamsGotoOptionsWaitUntilArray].
type AccountBrowserRenderingGetPdfParamsGotoOptionsWaitUntilUnion interface {
	implementsAccountBrowserRenderingGetPdfParamsGotoOptionsWaitUntilUnion()
}

type AccountBrowserRenderingGetPdfParamsGotoOptionsWaitUntilString string

const (
	AccountBrowserRenderingGetPdfParamsGotoOptionsWaitUntilStringLoad             AccountBrowserRenderingGetPdfParamsGotoOptionsWaitUntilString = "load"
	AccountBrowserRenderingGetPdfParamsGotoOptionsWaitUntilStringDomcontentloaded AccountBrowserRenderingGetPdfParamsGotoOptionsWaitUntilString = "domcontentloaded"
	AccountBrowserRenderingGetPdfParamsGotoOptionsWaitUntilStringNetworkidle0     AccountBrowserRenderingGetPdfParamsGotoOptionsWaitUntilString = "networkidle0"
	AccountBrowserRenderingGetPdfParamsGotoOptionsWaitUntilStringNetworkidle2     AccountBrowserRenderingGetPdfParamsGotoOptionsWaitUntilString = "networkidle2"
)

func (r AccountBrowserRenderingGetPdfParamsGotoOptionsWaitUntilString) IsKnown() bool {
	switch r {
	case AccountBrowserRenderingGetPdfParamsGotoOptionsWaitUntilStringLoad, AccountBrowserRenderingGetPdfParamsGotoOptionsWaitUntilStringDomcontentloaded, AccountBrowserRenderingGetPdfParamsGotoOptionsWaitUntilStringNetworkidle0, AccountBrowserRenderingGetPdfParamsGotoOptionsWaitUntilStringNetworkidle2:
		return true
	}
	return false
}

func (r AccountBrowserRenderingGetPdfParamsGotoOptionsWaitUntilString) implementsAccountBrowserRenderingGetPdfParamsGotoOptionsWaitUntilUnion() {
}

type AccountBrowserRenderingGetPdfParamsGotoOptionsWaitUntilArray []AccountBrowserRenderingGetPdfParamsGotoOptionsWaitUntilArrayItem

func (r AccountBrowserRenderingGetPdfParamsGotoOptionsWaitUntilArray) implementsAccountBrowserRenderingGetPdfParamsGotoOptionsWaitUntilUnion() {
}

type AccountBrowserRenderingGetPdfParamsGotoOptionsWaitUntilArrayItem string

const (
	AccountBrowserRenderingGetPdfParamsGotoOptionsWaitUntilArrayItemLoad             AccountBrowserRenderingGetPdfParamsGotoOptionsWaitUntilArrayItem = "load"
	AccountBrowserRenderingGetPdfParamsGotoOptionsWaitUntilArrayItemDomcontentloaded AccountBrowserRenderingGetPdfParamsGotoOptionsWaitUntilArrayItem = "domcontentloaded"
	AccountBrowserRenderingGetPdfParamsGotoOptionsWaitUntilArrayItemNetworkidle0     AccountBrowserRenderingGetPdfParamsGotoOptionsWaitUntilArrayItem = "networkidle0"
	AccountBrowserRenderingGetPdfParamsGotoOptionsWaitUntilArrayItemNetworkidle2     AccountBrowserRenderingGetPdfParamsGotoOptionsWaitUntilArrayItem = "networkidle2"
)

func (r AccountBrowserRenderingGetPdfParamsGotoOptionsWaitUntilArrayItem) IsKnown() bool {
	switch r {
	case AccountBrowserRenderingGetPdfParamsGotoOptionsWaitUntilArrayItemLoad, AccountBrowserRenderingGetPdfParamsGotoOptionsWaitUntilArrayItemDomcontentloaded, AccountBrowserRenderingGetPdfParamsGotoOptionsWaitUntilArrayItemNetworkidle0, AccountBrowserRenderingGetPdfParamsGotoOptionsWaitUntilArrayItemNetworkidle2:
		return true
	}
	return false
}

type AccountBrowserRenderingGetPdfParamsRejectResourceType string

const (
	AccountBrowserRenderingGetPdfParamsRejectResourceTypeDocument           AccountBrowserRenderingGetPdfParamsRejectResourceType = "document"
	AccountBrowserRenderingGetPdfParamsRejectResourceTypeStylesheet         AccountBrowserRenderingGetPdfParamsRejectResourceType = "stylesheet"
	AccountBrowserRenderingGetPdfParamsRejectResourceTypeImage              AccountBrowserRenderingGetPdfParamsRejectResourceType = "image"
	AccountBrowserRenderingGetPdfParamsRejectResourceTypeMedia              AccountBrowserRenderingGetPdfParamsRejectResourceType = "media"
	AccountBrowserRenderingGetPdfParamsRejectResourceTypeFont               AccountBrowserRenderingGetPdfParamsRejectResourceType = "font"
	AccountBrowserRenderingGetPdfParamsRejectResourceTypeScript             AccountBrowserRenderingGetPdfParamsRejectResourceType = "script"
	AccountBrowserRenderingGetPdfParamsRejectResourceTypeTexttrack          AccountBrowserRenderingGetPdfParamsRejectResourceType = "texttrack"
	AccountBrowserRenderingGetPdfParamsRejectResourceTypeXhr                AccountBrowserRenderingGetPdfParamsRejectResourceType = "xhr"
	AccountBrowserRenderingGetPdfParamsRejectResourceTypeFetch              AccountBrowserRenderingGetPdfParamsRejectResourceType = "fetch"
	AccountBrowserRenderingGetPdfParamsRejectResourceTypePrefetch           AccountBrowserRenderingGetPdfParamsRejectResourceType = "prefetch"
	AccountBrowserRenderingGetPdfParamsRejectResourceTypeEventsource        AccountBrowserRenderingGetPdfParamsRejectResourceType = "eventsource"
	AccountBrowserRenderingGetPdfParamsRejectResourceTypeWebsocket          AccountBrowserRenderingGetPdfParamsRejectResourceType = "websocket"
	AccountBrowserRenderingGetPdfParamsRejectResourceTypeManifest           AccountBrowserRenderingGetPdfParamsRejectResourceType = "manifest"
	AccountBrowserRenderingGetPdfParamsRejectResourceTypeSignedexchange     AccountBrowserRenderingGetPdfParamsRejectResourceType = "signedexchange"
	AccountBrowserRenderingGetPdfParamsRejectResourceTypePing               AccountBrowserRenderingGetPdfParamsRejectResourceType = "ping"
	AccountBrowserRenderingGetPdfParamsRejectResourceTypeCspviolationreport AccountBrowserRenderingGetPdfParamsRejectResourceType = "cspviolationreport"
	AccountBrowserRenderingGetPdfParamsRejectResourceTypePreflight          AccountBrowserRenderingGetPdfParamsRejectResourceType = "preflight"
	AccountBrowserRenderingGetPdfParamsRejectResourceTypeOther              AccountBrowserRenderingGetPdfParamsRejectResourceType = "other"
)

func (r AccountBrowserRenderingGetPdfParamsRejectResourceType) IsKnown() bool {
	switch r {
	case AccountBrowserRenderingGetPdfParamsRejectResourceTypeDocument, AccountBrowserRenderingGetPdfParamsRejectResourceTypeStylesheet, AccountBrowserRenderingGetPdfParamsRejectResourceTypeImage, AccountBrowserRenderingGetPdfParamsRejectResourceTypeMedia, AccountBrowserRenderingGetPdfParamsRejectResourceTypeFont, AccountBrowserRenderingGetPdfParamsRejectResourceTypeScript, AccountBrowserRenderingGetPdfParamsRejectResourceTypeTexttrack, AccountBrowserRenderingGetPdfParamsRejectResourceTypeXhr, AccountBrowserRenderingGetPdfParamsRejectResourceTypeFetch, AccountBrowserRenderingGetPdfParamsRejectResourceTypePrefetch, AccountBrowserRenderingGetPdfParamsRejectResourceTypeEventsource, AccountBrowserRenderingGetPdfParamsRejectResourceTypeWebsocket, AccountBrowserRenderingGetPdfParamsRejectResourceTypeManifest, AccountBrowserRenderingGetPdfParamsRejectResourceTypeSignedexchange, AccountBrowserRenderingGetPdfParamsRejectResourceTypePing, AccountBrowserRenderingGetPdfParamsRejectResourceTypeCspviolationreport, AccountBrowserRenderingGetPdfParamsRejectResourceTypePreflight, AccountBrowserRenderingGetPdfParamsRejectResourceTypeOther:
		return true
	}
	return false
}

// Check [options](https://pptr.dev/api/puppeteer.page.setviewport).
type AccountBrowserRenderingGetPdfParamsViewport struct {
	Height            param.Field[float64] `json:"height,required"`
	Width             param.Field[float64] `json:"width,required"`
	DeviceScaleFactor param.Field[float64] `json:"deviceScaleFactor"`
	HasTouch          param.Field[bool]    `json:"hasTouch"`
	IsLandscape       param.Field[bool]    `json:"isLandscape"`
	IsMobile          param.Field[bool]    `json:"isMobile"`
}

func (r AccountBrowserRenderingGetPdfParamsViewport) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Wait for the selector to appear in page. Check
// [options](https://pptr.dev/api/puppeteer.page.waitforselector).
type AccountBrowserRenderingGetPdfParamsWaitForSelector struct {
	Selector param.Field[string]                                                    `json:"selector,required"`
	Hidden   param.Field[AccountBrowserRenderingGetPdfParamsWaitForSelectorHidden]  `json:"hidden"`
	Timeout  param.Field[float64]                                                   `json:"timeout"`
	Visible  param.Field[AccountBrowserRenderingGetPdfParamsWaitForSelectorVisible] `json:"visible"`
}

func (r AccountBrowserRenderingGetPdfParamsWaitForSelector) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountBrowserRenderingGetPdfParamsWaitForSelectorHidden bool

const (
	AccountBrowserRenderingGetPdfParamsWaitForSelectorHiddenTrue AccountBrowserRenderingGetPdfParamsWaitForSelectorHidden = true
)

func (r AccountBrowserRenderingGetPdfParamsWaitForSelectorHidden) IsKnown() bool {
	switch r {
	case AccountBrowserRenderingGetPdfParamsWaitForSelectorHiddenTrue:
		return true
	}
	return false
}

type AccountBrowserRenderingGetPdfParamsWaitForSelectorVisible bool

const (
	AccountBrowserRenderingGetPdfParamsWaitForSelectorVisibleTrue AccountBrowserRenderingGetPdfParamsWaitForSelectorVisible = true
)

func (r AccountBrowserRenderingGetPdfParamsWaitForSelectorVisible) IsKnown() bool {
	switch r {
	case AccountBrowserRenderingGetPdfParamsWaitForSelectorVisibleTrue:
		return true
	}
	return false
}

type AccountBrowserRenderingGetScreenshotParams struct {
	// Cache TTL default is 5s. Set to 0 to disable.
	CacheTtl param.Field[float64] `query:"cacheTTL"`
	// Adds a `<script>` tag into the page with the desired URL or content.
	AddScriptTag param.Field[[]AccountBrowserRenderingGetScreenshotParamsAddScriptTag] `json:"addScriptTag"`
	// Adds a `<link rel="stylesheet">` tag into the page with the desired URL or a
	// `<style type="text/css">` tag with the content.
	AddStyleTag param.Field[[]AccountBrowserRenderingGetScreenshotParamsAddStyleTag] `json:"addStyleTag"`
	// Only allow requests that match the provided regex patterns, eg. '/^.\*\.(css)'.
	AllowRequestPattern param.Field[[]string] `json:"allowRequestPattern"`
	// Only allow requests that match the provided resource types, eg. 'image' or
	// 'script'.
	AllowResourceTypes param.Field[[]AccountBrowserRenderingGetScreenshotParamsAllowResourceType] `json:"allowResourceTypes"`
	// Provide credentials for HTTP authentication.
	Authenticate param.Field[AccountBrowserRenderingGetScreenshotParamsAuthenticate] `json:"authenticate"`
	// Attempt to proceed when 'awaited' events fail or timeout.
	BestAttempt param.Field[bool] `json:"bestAttempt"`
	// Check [options](https://pptr.dev/api/puppeteer.page.setcookie).
	Cookies          param.Field[[]AccountBrowserRenderingGetScreenshotParamsCookie] `json:"cookies"`
	EmulateMediaType param.Field[string]                                             `json:"emulateMediaType"`
	// Check [options](https://pptr.dev/api/puppeteer.gotooptions).
	GotoOptions param.Field[AccountBrowserRenderingGetScreenshotParamsGotoOptions] `json:"gotoOptions"`
	// Set the content of the page, eg: `<h1>Hello World!!</h1>`. Either `html` or
	// `url` must be set.
	HTML param.Field[string] `json:"html"`
	// Block undesired requests that match the provided regex patterns, eg.
	// '/^.\*\.(css)'.
	RejectRequestPattern param.Field[[]string] `json:"rejectRequestPattern"`
	// Block undesired requests that match the provided resource types, eg. 'image' or
	// 'script'.
	RejectResourceTypes param.Field[[]AccountBrowserRenderingGetScreenshotParamsRejectResourceType] `json:"rejectResourceTypes"`
	// Check [options](https://pptr.dev/api/puppeteer.screenshotoptions).
	ScreenshotOptions    param.Field[AccountBrowserRenderingGetScreenshotParamsScreenshotOptions] `json:"screenshotOptions"`
	ScrollPage           param.Field[bool]                                                        `json:"scrollPage"`
	Selector             param.Field[string]                                                      `json:"selector"`
	SetExtraHTTPHeaders  param.Field[map[string]string]                                           `json:"setExtraHTTPHeaders"`
	SetJavaScriptEnabled param.Field[bool]                                                        `json:"setJavaScriptEnabled"`
	// URL to navigate to, eg. `https://example.com`.
	URL       param.Field[string] `json:"url" format:"uri"`
	UserAgent param.Field[string] `json:"userAgent"`
	// Check [options](https://pptr.dev/api/puppeteer.page.setviewport).
	Viewport param.Field[AccountBrowserRenderingGetScreenshotParamsViewport] `json:"viewport"`
	// Wait for the selector to appear in page. Check
	// [options](https://pptr.dev/api/puppeteer.page.waitforselector).
	WaitForSelector param.Field[AccountBrowserRenderingGetScreenshotParamsWaitForSelector] `json:"waitForSelector"`
	// Waits for a specified timeout before continuing.
	WaitForTimeout param.Field[float64] `json:"waitForTimeout"`
}

func (r AccountBrowserRenderingGetScreenshotParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [AccountBrowserRenderingGetScreenshotParams]'s query
// parameters as `url.Values`.
func (r AccountBrowserRenderingGetScreenshotParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountBrowserRenderingGetScreenshotParamsAddScriptTag struct {
	ID      param.Field[string] `json:"id"`
	Content param.Field[string] `json:"content"`
	Type    param.Field[string] `json:"type"`
	URL     param.Field[string] `json:"url"`
}

func (r AccountBrowserRenderingGetScreenshotParamsAddScriptTag) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountBrowserRenderingGetScreenshotParamsAddStyleTag struct {
	Content param.Field[string] `json:"content"`
	URL     param.Field[string] `json:"url"`
}

func (r AccountBrowserRenderingGetScreenshotParamsAddStyleTag) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountBrowserRenderingGetScreenshotParamsAllowResourceType string

const (
	AccountBrowserRenderingGetScreenshotParamsAllowResourceTypeDocument           AccountBrowserRenderingGetScreenshotParamsAllowResourceType = "document"
	AccountBrowserRenderingGetScreenshotParamsAllowResourceTypeStylesheet         AccountBrowserRenderingGetScreenshotParamsAllowResourceType = "stylesheet"
	AccountBrowserRenderingGetScreenshotParamsAllowResourceTypeImage              AccountBrowserRenderingGetScreenshotParamsAllowResourceType = "image"
	AccountBrowserRenderingGetScreenshotParamsAllowResourceTypeMedia              AccountBrowserRenderingGetScreenshotParamsAllowResourceType = "media"
	AccountBrowserRenderingGetScreenshotParamsAllowResourceTypeFont               AccountBrowserRenderingGetScreenshotParamsAllowResourceType = "font"
	AccountBrowserRenderingGetScreenshotParamsAllowResourceTypeScript             AccountBrowserRenderingGetScreenshotParamsAllowResourceType = "script"
	AccountBrowserRenderingGetScreenshotParamsAllowResourceTypeTexttrack          AccountBrowserRenderingGetScreenshotParamsAllowResourceType = "texttrack"
	AccountBrowserRenderingGetScreenshotParamsAllowResourceTypeXhr                AccountBrowserRenderingGetScreenshotParamsAllowResourceType = "xhr"
	AccountBrowserRenderingGetScreenshotParamsAllowResourceTypeFetch              AccountBrowserRenderingGetScreenshotParamsAllowResourceType = "fetch"
	AccountBrowserRenderingGetScreenshotParamsAllowResourceTypePrefetch           AccountBrowserRenderingGetScreenshotParamsAllowResourceType = "prefetch"
	AccountBrowserRenderingGetScreenshotParamsAllowResourceTypeEventsource        AccountBrowserRenderingGetScreenshotParamsAllowResourceType = "eventsource"
	AccountBrowserRenderingGetScreenshotParamsAllowResourceTypeWebsocket          AccountBrowserRenderingGetScreenshotParamsAllowResourceType = "websocket"
	AccountBrowserRenderingGetScreenshotParamsAllowResourceTypeManifest           AccountBrowserRenderingGetScreenshotParamsAllowResourceType = "manifest"
	AccountBrowserRenderingGetScreenshotParamsAllowResourceTypeSignedexchange     AccountBrowserRenderingGetScreenshotParamsAllowResourceType = "signedexchange"
	AccountBrowserRenderingGetScreenshotParamsAllowResourceTypePing               AccountBrowserRenderingGetScreenshotParamsAllowResourceType = "ping"
	AccountBrowserRenderingGetScreenshotParamsAllowResourceTypeCspviolationreport AccountBrowserRenderingGetScreenshotParamsAllowResourceType = "cspviolationreport"
	AccountBrowserRenderingGetScreenshotParamsAllowResourceTypePreflight          AccountBrowserRenderingGetScreenshotParamsAllowResourceType = "preflight"
	AccountBrowserRenderingGetScreenshotParamsAllowResourceTypeOther              AccountBrowserRenderingGetScreenshotParamsAllowResourceType = "other"
)

func (r AccountBrowserRenderingGetScreenshotParamsAllowResourceType) IsKnown() bool {
	switch r {
	case AccountBrowserRenderingGetScreenshotParamsAllowResourceTypeDocument, AccountBrowserRenderingGetScreenshotParamsAllowResourceTypeStylesheet, AccountBrowserRenderingGetScreenshotParamsAllowResourceTypeImage, AccountBrowserRenderingGetScreenshotParamsAllowResourceTypeMedia, AccountBrowserRenderingGetScreenshotParamsAllowResourceTypeFont, AccountBrowserRenderingGetScreenshotParamsAllowResourceTypeScript, AccountBrowserRenderingGetScreenshotParamsAllowResourceTypeTexttrack, AccountBrowserRenderingGetScreenshotParamsAllowResourceTypeXhr, AccountBrowserRenderingGetScreenshotParamsAllowResourceTypeFetch, AccountBrowserRenderingGetScreenshotParamsAllowResourceTypePrefetch, AccountBrowserRenderingGetScreenshotParamsAllowResourceTypeEventsource, AccountBrowserRenderingGetScreenshotParamsAllowResourceTypeWebsocket, AccountBrowserRenderingGetScreenshotParamsAllowResourceTypeManifest, AccountBrowserRenderingGetScreenshotParamsAllowResourceTypeSignedexchange, AccountBrowserRenderingGetScreenshotParamsAllowResourceTypePing, AccountBrowserRenderingGetScreenshotParamsAllowResourceTypeCspviolationreport, AccountBrowserRenderingGetScreenshotParamsAllowResourceTypePreflight, AccountBrowserRenderingGetScreenshotParamsAllowResourceTypeOther:
		return true
	}
	return false
}

// Provide credentials for HTTP authentication.
type AccountBrowserRenderingGetScreenshotParamsAuthenticate struct {
	Password param.Field[string] `json:"password,required"`
	Username param.Field[string] `json:"username,required"`
}

func (r AccountBrowserRenderingGetScreenshotParamsAuthenticate) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountBrowserRenderingGetScreenshotParamsCookie struct {
	Name         param.Field[string]                                                        `json:"name,required"`
	Value        param.Field[string]                                                        `json:"value,required"`
	Domain       param.Field[string]                                                        `json:"domain"`
	Expires      param.Field[float64]                                                       `json:"expires"`
	HTTPOnly     param.Field[bool]                                                          `json:"httpOnly"`
	PartitionKey param.Field[string]                                                        `json:"partitionKey"`
	Path         param.Field[string]                                                        `json:"path"`
	Priority     param.Field[AccountBrowserRenderingGetScreenshotParamsCookiesPriority]     `json:"priority"`
	SameParty    param.Field[bool]                                                          `json:"sameParty"`
	SameSite     param.Field[AccountBrowserRenderingGetScreenshotParamsCookiesSameSite]     `json:"sameSite"`
	Secure       param.Field[bool]                                                          `json:"secure"`
	SourcePort   param.Field[float64]                                                       `json:"sourcePort"`
	SourceScheme param.Field[AccountBrowserRenderingGetScreenshotParamsCookiesSourceScheme] `json:"sourceScheme"`
	URL          param.Field[string]                                                        `json:"url"`
}

func (r AccountBrowserRenderingGetScreenshotParamsCookie) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountBrowserRenderingGetScreenshotParamsCookiesPriority string

const (
	AccountBrowserRenderingGetScreenshotParamsCookiesPriorityLow    AccountBrowserRenderingGetScreenshotParamsCookiesPriority = "Low"
	AccountBrowserRenderingGetScreenshotParamsCookiesPriorityMedium AccountBrowserRenderingGetScreenshotParamsCookiesPriority = "Medium"
	AccountBrowserRenderingGetScreenshotParamsCookiesPriorityHigh   AccountBrowserRenderingGetScreenshotParamsCookiesPriority = "High"
)

func (r AccountBrowserRenderingGetScreenshotParamsCookiesPriority) IsKnown() bool {
	switch r {
	case AccountBrowserRenderingGetScreenshotParamsCookiesPriorityLow, AccountBrowserRenderingGetScreenshotParamsCookiesPriorityMedium, AccountBrowserRenderingGetScreenshotParamsCookiesPriorityHigh:
		return true
	}
	return false
}

type AccountBrowserRenderingGetScreenshotParamsCookiesSameSite string

const (
	AccountBrowserRenderingGetScreenshotParamsCookiesSameSiteStrict AccountBrowserRenderingGetScreenshotParamsCookiesSameSite = "Strict"
	AccountBrowserRenderingGetScreenshotParamsCookiesSameSiteLax    AccountBrowserRenderingGetScreenshotParamsCookiesSameSite = "Lax"
	AccountBrowserRenderingGetScreenshotParamsCookiesSameSiteNone   AccountBrowserRenderingGetScreenshotParamsCookiesSameSite = "None"
)

func (r AccountBrowserRenderingGetScreenshotParamsCookiesSameSite) IsKnown() bool {
	switch r {
	case AccountBrowserRenderingGetScreenshotParamsCookiesSameSiteStrict, AccountBrowserRenderingGetScreenshotParamsCookiesSameSiteLax, AccountBrowserRenderingGetScreenshotParamsCookiesSameSiteNone:
		return true
	}
	return false
}

type AccountBrowserRenderingGetScreenshotParamsCookiesSourceScheme string

const (
	AccountBrowserRenderingGetScreenshotParamsCookiesSourceSchemeUnset     AccountBrowserRenderingGetScreenshotParamsCookiesSourceScheme = "Unset"
	AccountBrowserRenderingGetScreenshotParamsCookiesSourceSchemeNonSecure AccountBrowserRenderingGetScreenshotParamsCookiesSourceScheme = "NonSecure"
	AccountBrowserRenderingGetScreenshotParamsCookiesSourceSchemeSecure    AccountBrowserRenderingGetScreenshotParamsCookiesSourceScheme = "Secure"
)

func (r AccountBrowserRenderingGetScreenshotParamsCookiesSourceScheme) IsKnown() bool {
	switch r {
	case AccountBrowserRenderingGetScreenshotParamsCookiesSourceSchemeUnset, AccountBrowserRenderingGetScreenshotParamsCookiesSourceSchemeNonSecure, AccountBrowserRenderingGetScreenshotParamsCookiesSourceSchemeSecure:
		return true
	}
	return false
}

// Check [options](https://pptr.dev/api/puppeteer.gotooptions).
type AccountBrowserRenderingGetScreenshotParamsGotoOptions struct {
	Referer        param.Field[string]                                                              `json:"referer"`
	ReferrerPolicy param.Field[string]                                                              `json:"referrerPolicy"`
	Timeout        param.Field[float64]                                                             `json:"timeout"`
	WaitUntil      param.Field[AccountBrowserRenderingGetScreenshotParamsGotoOptionsWaitUntilUnion] `json:"waitUntil"`
}

func (r AccountBrowserRenderingGetScreenshotParamsGotoOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by
// [AccountBrowserRenderingGetScreenshotParamsGotoOptionsWaitUntilString],
// [AccountBrowserRenderingGetScreenshotParamsGotoOptionsWaitUntilArray].
type AccountBrowserRenderingGetScreenshotParamsGotoOptionsWaitUntilUnion interface {
	implementsAccountBrowserRenderingGetScreenshotParamsGotoOptionsWaitUntilUnion()
}

type AccountBrowserRenderingGetScreenshotParamsGotoOptionsWaitUntilString string

const (
	AccountBrowserRenderingGetScreenshotParamsGotoOptionsWaitUntilStringLoad             AccountBrowserRenderingGetScreenshotParamsGotoOptionsWaitUntilString = "load"
	AccountBrowserRenderingGetScreenshotParamsGotoOptionsWaitUntilStringDomcontentloaded AccountBrowserRenderingGetScreenshotParamsGotoOptionsWaitUntilString = "domcontentloaded"
	AccountBrowserRenderingGetScreenshotParamsGotoOptionsWaitUntilStringNetworkidle0     AccountBrowserRenderingGetScreenshotParamsGotoOptionsWaitUntilString = "networkidle0"
	AccountBrowserRenderingGetScreenshotParamsGotoOptionsWaitUntilStringNetworkidle2     AccountBrowserRenderingGetScreenshotParamsGotoOptionsWaitUntilString = "networkidle2"
)

func (r AccountBrowserRenderingGetScreenshotParamsGotoOptionsWaitUntilString) IsKnown() bool {
	switch r {
	case AccountBrowserRenderingGetScreenshotParamsGotoOptionsWaitUntilStringLoad, AccountBrowserRenderingGetScreenshotParamsGotoOptionsWaitUntilStringDomcontentloaded, AccountBrowserRenderingGetScreenshotParamsGotoOptionsWaitUntilStringNetworkidle0, AccountBrowserRenderingGetScreenshotParamsGotoOptionsWaitUntilStringNetworkidle2:
		return true
	}
	return false
}

func (r AccountBrowserRenderingGetScreenshotParamsGotoOptionsWaitUntilString) implementsAccountBrowserRenderingGetScreenshotParamsGotoOptionsWaitUntilUnion() {
}

type AccountBrowserRenderingGetScreenshotParamsGotoOptionsWaitUntilArray []AccountBrowserRenderingGetScreenshotParamsGotoOptionsWaitUntilArrayItem

func (r AccountBrowserRenderingGetScreenshotParamsGotoOptionsWaitUntilArray) implementsAccountBrowserRenderingGetScreenshotParamsGotoOptionsWaitUntilUnion() {
}

type AccountBrowserRenderingGetScreenshotParamsGotoOptionsWaitUntilArrayItem string

const (
	AccountBrowserRenderingGetScreenshotParamsGotoOptionsWaitUntilArrayItemLoad             AccountBrowserRenderingGetScreenshotParamsGotoOptionsWaitUntilArrayItem = "load"
	AccountBrowserRenderingGetScreenshotParamsGotoOptionsWaitUntilArrayItemDomcontentloaded AccountBrowserRenderingGetScreenshotParamsGotoOptionsWaitUntilArrayItem = "domcontentloaded"
	AccountBrowserRenderingGetScreenshotParamsGotoOptionsWaitUntilArrayItemNetworkidle0     AccountBrowserRenderingGetScreenshotParamsGotoOptionsWaitUntilArrayItem = "networkidle0"
	AccountBrowserRenderingGetScreenshotParamsGotoOptionsWaitUntilArrayItemNetworkidle2     AccountBrowserRenderingGetScreenshotParamsGotoOptionsWaitUntilArrayItem = "networkidle2"
)

func (r AccountBrowserRenderingGetScreenshotParamsGotoOptionsWaitUntilArrayItem) IsKnown() bool {
	switch r {
	case AccountBrowserRenderingGetScreenshotParamsGotoOptionsWaitUntilArrayItemLoad, AccountBrowserRenderingGetScreenshotParamsGotoOptionsWaitUntilArrayItemDomcontentloaded, AccountBrowserRenderingGetScreenshotParamsGotoOptionsWaitUntilArrayItemNetworkidle0, AccountBrowserRenderingGetScreenshotParamsGotoOptionsWaitUntilArrayItemNetworkidle2:
		return true
	}
	return false
}

type AccountBrowserRenderingGetScreenshotParamsRejectResourceType string

const (
	AccountBrowserRenderingGetScreenshotParamsRejectResourceTypeDocument           AccountBrowserRenderingGetScreenshotParamsRejectResourceType = "document"
	AccountBrowserRenderingGetScreenshotParamsRejectResourceTypeStylesheet         AccountBrowserRenderingGetScreenshotParamsRejectResourceType = "stylesheet"
	AccountBrowserRenderingGetScreenshotParamsRejectResourceTypeImage              AccountBrowserRenderingGetScreenshotParamsRejectResourceType = "image"
	AccountBrowserRenderingGetScreenshotParamsRejectResourceTypeMedia              AccountBrowserRenderingGetScreenshotParamsRejectResourceType = "media"
	AccountBrowserRenderingGetScreenshotParamsRejectResourceTypeFont               AccountBrowserRenderingGetScreenshotParamsRejectResourceType = "font"
	AccountBrowserRenderingGetScreenshotParamsRejectResourceTypeScript             AccountBrowserRenderingGetScreenshotParamsRejectResourceType = "script"
	AccountBrowserRenderingGetScreenshotParamsRejectResourceTypeTexttrack          AccountBrowserRenderingGetScreenshotParamsRejectResourceType = "texttrack"
	AccountBrowserRenderingGetScreenshotParamsRejectResourceTypeXhr                AccountBrowserRenderingGetScreenshotParamsRejectResourceType = "xhr"
	AccountBrowserRenderingGetScreenshotParamsRejectResourceTypeFetch              AccountBrowserRenderingGetScreenshotParamsRejectResourceType = "fetch"
	AccountBrowserRenderingGetScreenshotParamsRejectResourceTypePrefetch           AccountBrowserRenderingGetScreenshotParamsRejectResourceType = "prefetch"
	AccountBrowserRenderingGetScreenshotParamsRejectResourceTypeEventsource        AccountBrowserRenderingGetScreenshotParamsRejectResourceType = "eventsource"
	AccountBrowserRenderingGetScreenshotParamsRejectResourceTypeWebsocket          AccountBrowserRenderingGetScreenshotParamsRejectResourceType = "websocket"
	AccountBrowserRenderingGetScreenshotParamsRejectResourceTypeManifest           AccountBrowserRenderingGetScreenshotParamsRejectResourceType = "manifest"
	AccountBrowserRenderingGetScreenshotParamsRejectResourceTypeSignedexchange     AccountBrowserRenderingGetScreenshotParamsRejectResourceType = "signedexchange"
	AccountBrowserRenderingGetScreenshotParamsRejectResourceTypePing               AccountBrowserRenderingGetScreenshotParamsRejectResourceType = "ping"
	AccountBrowserRenderingGetScreenshotParamsRejectResourceTypeCspviolationreport AccountBrowserRenderingGetScreenshotParamsRejectResourceType = "cspviolationreport"
	AccountBrowserRenderingGetScreenshotParamsRejectResourceTypePreflight          AccountBrowserRenderingGetScreenshotParamsRejectResourceType = "preflight"
	AccountBrowserRenderingGetScreenshotParamsRejectResourceTypeOther              AccountBrowserRenderingGetScreenshotParamsRejectResourceType = "other"
)

func (r AccountBrowserRenderingGetScreenshotParamsRejectResourceType) IsKnown() bool {
	switch r {
	case AccountBrowserRenderingGetScreenshotParamsRejectResourceTypeDocument, AccountBrowserRenderingGetScreenshotParamsRejectResourceTypeStylesheet, AccountBrowserRenderingGetScreenshotParamsRejectResourceTypeImage, AccountBrowserRenderingGetScreenshotParamsRejectResourceTypeMedia, AccountBrowserRenderingGetScreenshotParamsRejectResourceTypeFont, AccountBrowserRenderingGetScreenshotParamsRejectResourceTypeScript, AccountBrowserRenderingGetScreenshotParamsRejectResourceTypeTexttrack, AccountBrowserRenderingGetScreenshotParamsRejectResourceTypeXhr, AccountBrowserRenderingGetScreenshotParamsRejectResourceTypeFetch, AccountBrowserRenderingGetScreenshotParamsRejectResourceTypePrefetch, AccountBrowserRenderingGetScreenshotParamsRejectResourceTypeEventsource, AccountBrowserRenderingGetScreenshotParamsRejectResourceTypeWebsocket, AccountBrowserRenderingGetScreenshotParamsRejectResourceTypeManifest, AccountBrowserRenderingGetScreenshotParamsRejectResourceTypeSignedexchange, AccountBrowserRenderingGetScreenshotParamsRejectResourceTypePing, AccountBrowserRenderingGetScreenshotParamsRejectResourceTypeCspviolationreport, AccountBrowserRenderingGetScreenshotParamsRejectResourceTypePreflight, AccountBrowserRenderingGetScreenshotParamsRejectResourceTypeOther:
		return true
	}
	return false
}

// Check [options](https://pptr.dev/api/puppeteer.screenshotoptions).
type AccountBrowserRenderingGetScreenshotParamsScreenshotOptions struct {
	CaptureBeyondViewport param.Field[bool]                                                                `json:"captureBeyondViewport"`
	Clip                  param.Field[AccountBrowserRenderingGetScreenshotParamsScreenshotOptionsClip]     `json:"clip"`
	Encoding              param.Field[AccountBrowserRenderingGetScreenshotParamsScreenshotOptionsEncoding] `json:"encoding"`
	FromSurface           param.Field[bool]                                                                `json:"fromSurface"`
	FullPage              param.Field[bool]                                                                `json:"fullPage"`
	OmitBackground        param.Field[bool]                                                                `json:"omitBackground"`
	OptimizeForSpeed      param.Field[bool]                                                                `json:"optimizeForSpeed"`
	Quality               param.Field[float64]                                                             `json:"quality"`
	Type                  param.Field[AccountBrowserRenderingGetScreenshotParamsScreenshotOptionsType]     `json:"type"`
}

func (r AccountBrowserRenderingGetScreenshotParamsScreenshotOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountBrowserRenderingGetScreenshotParamsScreenshotOptionsClip struct {
	Height param.Field[float64] `json:"height,required"`
	Width  param.Field[float64] `json:"width,required"`
	X      param.Field[float64] `json:"x,required"`
	Y      param.Field[float64] `json:"y,required"`
	Scale  param.Field[float64] `json:"scale"`
}

func (r AccountBrowserRenderingGetScreenshotParamsScreenshotOptionsClip) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountBrowserRenderingGetScreenshotParamsScreenshotOptionsEncoding string

const (
	AccountBrowserRenderingGetScreenshotParamsScreenshotOptionsEncodingBinary AccountBrowserRenderingGetScreenshotParamsScreenshotOptionsEncoding = "binary"
	AccountBrowserRenderingGetScreenshotParamsScreenshotOptionsEncodingBase64 AccountBrowserRenderingGetScreenshotParamsScreenshotOptionsEncoding = "base64"
)

func (r AccountBrowserRenderingGetScreenshotParamsScreenshotOptionsEncoding) IsKnown() bool {
	switch r {
	case AccountBrowserRenderingGetScreenshotParamsScreenshotOptionsEncodingBinary, AccountBrowserRenderingGetScreenshotParamsScreenshotOptionsEncodingBase64:
		return true
	}
	return false
}

type AccountBrowserRenderingGetScreenshotParamsScreenshotOptionsType string

const (
	AccountBrowserRenderingGetScreenshotParamsScreenshotOptionsTypePng  AccountBrowserRenderingGetScreenshotParamsScreenshotOptionsType = "png"
	AccountBrowserRenderingGetScreenshotParamsScreenshotOptionsTypeJpeg AccountBrowserRenderingGetScreenshotParamsScreenshotOptionsType = "jpeg"
	AccountBrowserRenderingGetScreenshotParamsScreenshotOptionsTypeWebp AccountBrowserRenderingGetScreenshotParamsScreenshotOptionsType = "webp"
)

func (r AccountBrowserRenderingGetScreenshotParamsScreenshotOptionsType) IsKnown() bool {
	switch r {
	case AccountBrowserRenderingGetScreenshotParamsScreenshotOptionsTypePng, AccountBrowserRenderingGetScreenshotParamsScreenshotOptionsTypeJpeg, AccountBrowserRenderingGetScreenshotParamsScreenshotOptionsTypeWebp:
		return true
	}
	return false
}

// Check [options](https://pptr.dev/api/puppeteer.page.setviewport).
type AccountBrowserRenderingGetScreenshotParamsViewport struct {
	Height            param.Field[float64] `json:"height,required"`
	Width             param.Field[float64] `json:"width,required"`
	DeviceScaleFactor param.Field[float64] `json:"deviceScaleFactor"`
	HasTouch          param.Field[bool]    `json:"hasTouch"`
	IsLandscape       param.Field[bool]    `json:"isLandscape"`
	IsMobile          param.Field[bool]    `json:"isMobile"`
}

func (r AccountBrowserRenderingGetScreenshotParamsViewport) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Wait for the selector to appear in page. Check
// [options](https://pptr.dev/api/puppeteer.page.waitforselector).
type AccountBrowserRenderingGetScreenshotParamsWaitForSelector struct {
	Selector param.Field[string]                                                           `json:"selector,required"`
	Hidden   param.Field[AccountBrowserRenderingGetScreenshotParamsWaitForSelectorHidden]  `json:"hidden"`
	Timeout  param.Field[float64]                                                          `json:"timeout"`
	Visible  param.Field[AccountBrowserRenderingGetScreenshotParamsWaitForSelectorVisible] `json:"visible"`
}

func (r AccountBrowserRenderingGetScreenshotParamsWaitForSelector) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountBrowserRenderingGetScreenshotParamsWaitForSelectorHidden bool

const (
	AccountBrowserRenderingGetScreenshotParamsWaitForSelectorHiddenTrue AccountBrowserRenderingGetScreenshotParamsWaitForSelectorHidden = true
)

func (r AccountBrowserRenderingGetScreenshotParamsWaitForSelectorHidden) IsKnown() bool {
	switch r {
	case AccountBrowserRenderingGetScreenshotParamsWaitForSelectorHiddenTrue:
		return true
	}
	return false
}

type AccountBrowserRenderingGetScreenshotParamsWaitForSelectorVisible bool

const (
	AccountBrowserRenderingGetScreenshotParamsWaitForSelectorVisibleTrue AccountBrowserRenderingGetScreenshotParamsWaitForSelectorVisible = true
)

func (r AccountBrowserRenderingGetScreenshotParamsWaitForSelectorVisible) IsKnown() bool {
	switch r {
	case AccountBrowserRenderingGetScreenshotParamsWaitForSelectorVisibleTrue:
		return true
	}
	return false
}

type AccountBrowserRenderingGetSnapshotParams struct {
	// Cache TTL default is 5s. Set to 0 to disable.
	CacheTtl param.Field[float64] `query:"cacheTTL"`
	// Adds a `<script>` tag into the page with the desired URL or content.
	AddScriptTag param.Field[[]AccountBrowserRenderingGetSnapshotParamsAddScriptTag] `json:"addScriptTag"`
	// Adds a `<link rel="stylesheet">` tag into the page with the desired URL or a
	// `<style type="text/css">` tag with the content.
	AddStyleTag param.Field[[]AccountBrowserRenderingGetSnapshotParamsAddStyleTag] `json:"addStyleTag"`
	// Only allow requests that match the provided regex patterns, eg. '/^.\*\.(css)'.
	AllowRequestPattern param.Field[[]string] `json:"allowRequestPattern"`
	// Only allow requests that match the provided resource types, eg. 'image' or
	// 'script'.
	AllowResourceTypes param.Field[[]AccountBrowserRenderingGetSnapshotParamsAllowResourceType] `json:"allowResourceTypes"`
	// Provide credentials for HTTP authentication.
	Authenticate param.Field[AccountBrowserRenderingGetSnapshotParamsAuthenticate] `json:"authenticate"`
	// Attempt to proceed when 'awaited' events fail or timeout.
	BestAttempt param.Field[bool] `json:"bestAttempt"`
	// Check [options](https://pptr.dev/api/puppeteer.page.setcookie).
	Cookies          param.Field[[]AccountBrowserRenderingGetSnapshotParamsCookie] `json:"cookies"`
	EmulateMediaType param.Field[string]                                           `json:"emulateMediaType"`
	// Check [options](https://pptr.dev/api/puppeteer.gotooptions).
	GotoOptions param.Field[AccountBrowserRenderingGetSnapshotParamsGotoOptions] `json:"gotoOptions"`
	// Set the content of the page, eg: `<h1>Hello World!!</h1>`. Either `html` or
	// `url` must be set.
	HTML param.Field[string] `json:"html"`
	// Block undesired requests that match the provided regex patterns, eg.
	// '/^.\*\.(css)'.
	RejectRequestPattern param.Field[[]string] `json:"rejectRequestPattern"`
	// Block undesired requests that match the provided resource types, eg. 'image' or
	// 'script'.
	RejectResourceTypes  param.Field[[]AccountBrowserRenderingGetSnapshotParamsRejectResourceType] `json:"rejectResourceTypes"`
	ScreenshotOptions    param.Field[AccountBrowserRenderingGetSnapshotParamsScreenshotOptions]    `json:"screenshotOptions"`
	SetExtraHTTPHeaders  param.Field[map[string]string]                                            `json:"setExtraHTTPHeaders"`
	SetJavaScriptEnabled param.Field[bool]                                                         `json:"setJavaScriptEnabled"`
	// URL to navigate to, eg. `https://example.com`.
	URL       param.Field[string] `json:"url" format:"uri"`
	UserAgent param.Field[string] `json:"userAgent"`
	// Check [options](https://pptr.dev/api/puppeteer.page.setviewport).
	Viewport param.Field[AccountBrowserRenderingGetSnapshotParamsViewport] `json:"viewport"`
	// Wait for the selector to appear in page. Check
	// [options](https://pptr.dev/api/puppeteer.page.waitforselector).
	WaitForSelector param.Field[AccountBrowserRenderingGetSnapshotParamsWaitForSelector] `json:"waitForSelector"`
	// Waits for a specified timeout before continuing.
	WaitForTimeout param.Field[float64] `json:"waitForTimeout"`
}

func (r AccountBrowserRenderingGetSnapshotParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [AccountBrowserRenderingGetSnapshotParams]'s query
// parameters as `url.Values`.
func (r AccountBrowserRenderingGetSnapshotParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountBrowserRenderingGetSnapshotParamsAddScriptTag struct {
	ID      param.Field[string] `json:"id"`
	Content param.Field[string] `json:"content"`
	Type    param.Field[string] `json:"type"`
	URL     param.Field[string] `json:"url"`
}

func (r AccountBrowserRenderingGetSnapshotParamsAddScriptTag) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountBrowserRenderingGetSnapshotParamsAddStyleTag struct {
	Content param.Field[string] `json:"content"`
	URL     param.Field[string] `json:"url"`
}

func (r AccountBrowserRenderingGetSnapshotParamsAddStyleTag) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountBrowserRenderingGetSnapshotParamsAllowResourceType string

const (
	AccountBrowserRenderingGetSnapshotParamsAllowResourceTypeDocument           AccountBrowserRenderingGetSnapshotParamsAllowResourceType = "document"
	AccountBrowserRenderingGetSnapshotParamsAllowResourceTypeStylesheet         AccountBrowserRenderingGetSnapshotParamsAllowResourceType = "stylesheet"
	AccountBrowserRenderingGetSnapshotParamsAllowResourceTypeImage              AccountBrowserRenderingGetSnapshotParamsAllowResourceType = "image"
	AccountBrowserRenderingGetSnapshotParamsAllowResourceTypeMedia              AccountBrowserRenderingGetSnapshotParamsAllowResourceType = "media"
	AccountBrowserRenderingGetSnapshotParamsAllowResourceTypeFont               AccountBrowserRenderingGetSnapshotParamsAllowResourceType = "font"
	AccountBrowserRenderingGetSnapshotParamsAllowResourceTypeScript             AccountBrowserRenderingGetSnapshotParamsAllowResourceType = "script"
	AccountBrowserRenderingGetSnapshotParamsAllowResourceTypeTexttrack          AccountBrowserRenderingGetSnapshotParamsAllowResourceType = "texttrack"
	AccountBrowserRenderingGetSnapshotParamsAllowResourceTypeXhr                AccountBrowserRenderingGetSnapshotParamsAllowResourceType = "xhr"
	AccountBrowserRenderingGetSnapshotParamsAllowResourceTypeFetch              AccountBrowserRenderingGetSnapshotParamsAllowResourceType = "fetch"
	AccountBrowserRenderingGetSnapshotParamsAllowResourceTypePrefetch           AccountBrowserRenderingGetSnapshotParamsAllowResourceType = "prefetch"
	AccountBrowserRenderingGetSnapshotParamsAllowResourceTypeEventsource        AccountBrowserRenderingGetSnapshotParamsAllowResourceType = "eventsource"
	AccountBrowserRenderingGetSnapshotParamsAllowResourceTypeWebsocket          AccountBrowserRenderingGetSnapshotParamsAllowResourceType = "websocket"
	AccountBrowserRenderingGetSnapshotParamsAllowResourceTypeManifest           AccountBrowserRenderingGetSnapshotParamsAllowResourceType = "manifest"
	AccountBrowserRenderingGetSnapshotParamsAllowResourceTypeSignedexchange     AccountBrowserRenderingGetSnapshotParamsAllowResourceType = "signedexchange"
	AccountBrowserRenderingGetSnapshotParamsAllowResourceTypePing               AccountBrowserRenderingGetSnapshotParamsAllowResourceType = "ping"
	AccountBrowserRenderingGetSnapshotParamsAllowResourceTypeCspviolationreport AccountBrowserRenderingGetSnapshotParamsAllowResourceType = "cspviolationreport"
	AccountBrowserRenderingGetSnapshotParamsAllowResourceTypePreflight          AccountBrowserRenderingGetSnapshotParamsAllowResourceType = "preflight"
	AccountBrowserRenderingGetSnapshotParamsAllowResourceTypeOther              AccountBrowserRenderingGetSnapshotParamsAllowResourceType = "other"
)

func (r AccountBrowserRenderingGetSnapshotParamsAllowResourceType) IsKnown() bool {
	switch r {
	case AccountBrowserRenderingGetSnapshotParamsAllowResourceTypeDocument, AccountBrowserRenderingGetSnapshotParamsAllowResourceTypeStylesheet, AccountBrowserRenderingGetSnapshotParamsAllowResourceTypeImage, AccountBrowserRenderingGetSnapshotParamsAllowResourceTypeMedia, AccountBrowserRenderingGetSnapshotParamsAllowResourceTypeFont, AccountBrowserRenderingGetSnapshotParamsAllowResourceTypeScript, AccountBrowserRenderingGetSnapshotParamsAllowResourceTypeTexttrack, AccountBrowserRenderingGetSnapshotParamsAllowResourceTypeXhr, AccountBrowserRenderingGetSnapshotParamsAllowResourceTypeFetch, AccountBrowserRenderingGetSnapshotParamsAllowResourceTypePrefetch, AccountBrowserRenderingGetSnapshotParamsAllowResourceTypeEventsource, AccountBrowserRenderingGetSnapshotParamsAllowResourceTypeWebsocket, AccountBrowserRenderingGetSnapshotParamsAllowResourceTypeManifest, AccountBrowserRenderingGetSnapshotParamsAllowResourceTypeSignedexchange, AccountBrowserRenderingGetSnapshotParamsAllowResourceTypePing, AccountBrowserRenderingGetSnapshotParamsAllowResourceTypeCspviolationreport, AccountBrowserRenderingGetSnapshotParamsAllowResourceTypePreflight, AccountBrowserRenderingGetSnapshotParamsAllowResourceTypeOther:
		return true
	}
	return false
}

// Provide credentials for HTTP authentication.
type AccountBrowserRenderingGetSnapshotParamsAuthenticate struct {
	Password param.Field[string] `json:"password,required"`
	Username param.Field[string] `json:"username,required"`
}

func (r AccountBrowserRenderingGetSnapshotParamsAuthenticate) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountBrowserRenderingGetSnapshotParamsCookie struct {
	Name         param.Field[string]                                                      `json:"name,required"`
	Value        param.Field[string]                                                      `json:"value,required"`
	Domain       param.Field[string]                                                      `json:"domain"`
	Expires      param.Field[float64]                                                     `json:"expires"`
	HTTPOnly     param.Field[bool]                                                        `json:"httpOnly"`
	PartitionKey param.Field[string]                                                      `json:"partitionKey"`
	Path         param.Field[string]                                                      `json:"path"`
	Priority     param.Field[AccountBrowserRenderingGetSnapshotParamsCookiesPriority]     `json:"priority"`
	SameParty    param.Field[bool]                                                        `json:"sameParty"`
	SameSite     param.Field[AccountBrowserRenderingGetSnapshotParamsCookiesSameSite]     `json:"sameSite"`
	Secure       param.Field[bool]                                                        `json:"secure"`
	SourcePort   param.Field[float64]                                                     `json:"sourcePort"`
	SourceScheme param.Field[AccountBrowserRenderingGetSnapshotParamsCookiesSourceScheme] `json:"sourceScheme"`
	URL          param.Field[string]                                                      `json:"url"`
}

func (r AccountBrowserRenderingGetSnapshotParamsCookie) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountBrowserRenderingGetSnapshotParamsCookiesPriority string

const (
	AccountBrowserRenderingGetSnapshotParamsCookiesPriorityLow    AccountBrowserRenderingGetSnapshotParamsCookiesPriority = "Low"
	AccountBrowserRenderingGetSnapshotParamsCookiesPriorityMedium AccountBrowserRenderingGetSnapshotParamsCookiesPriority = "Medium"
	AccountBrowserRenderingGetSnapshotParamsCookiesPriorityHigh   AccountBrowserRenderingGetSnapshotParamsCookiesPriority = "High"
)

func (r AccountBrowserRenderingGetSnapshotParamsCookiesPriority) IsKnown() bool {
	switch r {
	case AccountBrowserRenderingGetSnapshotParamsCookiesPriorityLow, AccountBrowserRenderingGetSnapshotParamsCookiesPriorityMedium, AccountBrowserRenderingGetSnapshotParamsCookiesPriorityHigh:
		return true
	}
	return false
}

type AccountBrowserRenderingGetSnapshotParamsCookiesSameSite string

const (
	AccountBrowserRenderingGetSnapshotParamsCookiesSameSiteStrict AccountBrowserRenderingGetSnapshotParamsCookiesSameSite = "Strict"
	AccountBrowserRenderingGetSnapshotParamsCookiesSameSiteLax    AccountBrowserRenderingGetSnapshotParamsCookiesSameSite = "Lax"
	AccountBrowserRenderingGetSnapshotParamsCookiesSameSiteNone   AccountBrowserRenderingGetSnapshotParamsCookiesSameSite = "None"
)

func (r AccountBrowserRenderingGetSnapshotParamsCookiesSameSite) IsKnown() bool {
	switch r {
	case AccountBrowserRenderingGetSnapshotParamsCookiesSameSiteStrict, AccountBrowserRenderingGetSnapshotParamsCookiesSameSiteLax, AccountBrowserRenderingGetSnapshotParamsCookiesSameSiteNone:
		return true
	}
	return false
}

type AccountBrowserRenderingGetSnapshotParamsCookiesSourceScheme string

const (
	AccountBrowserRenderingGetSnapshotParamsCookiesSourceSchemeUnset     AccountBrowserRenderingGetSnapshotParamsCookiesSourceScheme = "Unset"
	AccountBrowserRenderingGetSnapshotParamsCookiesSourceSchemeNonSecure AccountBrowserRenderingGetSnapshotParamsCookiesSourceScheme = "NonSecure"
	AccountBrowserRenderingGetSnapshotParamsCookiesSourceSchemeSecure    AccountBrowserRenderingGetSnapshotParamsCookiesSourceScheme = "Secure"
)

func (r AccountBrowserRenderingGetSnapshotParamsCookiesSourceScheme) IsKnown() bool {
	switch r {
	case AccountBrowserRenderingGetSnapshotParamsCookiesSourceSchemeUnset, AccountBrowserRenderingGetSnapshotParamsCookiesSourceSchemeNonSecure, AccountBrowserRenderingGetSnapshotParamsCookiesSourceSchemeSecure:
		return true
	}
	return false
}

// Check [options](https://pptr.dev/api/puppeteer.gotooptions).
type AccountBrowserRenderingGetSnapshotParamsGotoOptions struct {
	Referer        param.Field[string]                                                            `json:"referer"`
	ReferrerPolicy param.Field[string]                                                            `json:"referrerPolicy"`
	Timeout        param.Field[float64]                                                           `json:"timeout"`
	WaitUntil      param.Field[AccountBrowserRenderingGetSnapshotParamsGotoOptionsWaitUntilUnion] `json:"waitUntil"`
}

func (r AccountBrowserRenderingGetSnapshotParamsGotoOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by
// [AccountBrowserRenderingGetSnapshotParamsGotoOptionsWaitUntilString],
// [AccountBrowserRenderingGetSnapshotParamsGotoOptionsWaitUntilArray].
type AccountBrowserRenderingGetSnapshotParamsGotoOptionsWaitUntilUnion interface {
	implementsAccountBrowserRenderingGetSnapshotParamsGotoOptionsWaitUntilUnion()
}

type AccountBrowserRenderingGetSnapshotParamsGotoOptionsWaitUntilString string

const (
	AccountBrowserRenderingGetSnapshotParamsGotoOptionsWaitUntilStringLoad             AccountBrowserRenderingGetSnapshotParamsGotoOptionsWaitUntilString = "load"
	AccountBrowserRenderingGetSnapshotParamsGotoOptionsWaitUntilStringDomcontentloaded AccountBrowserRenderingGetSnapshotParamsGotoOptionsWaitUntilString = "domcontentloaded"
	AccountBrowserRenderingGetSnapshotParamsGotoOptionsWaitUntilStringNetworkidle0     AccountBrowserRenderingGetSnapshotParamsGotoOptionsWaitUntilString = "networkidle0"
	AccountBrowserRenderingGetSnapshotParamsGotoOptionsWaitUntilStringNetworkidle2     AccountBrowserRenderingGetSnapshotParamsGotoOptionsWaitUntilString = "networkidle2"
)

func (r AccountBrowserRenderingGetSnapshotParamsGotoOptionsWaitUntilString) IsKnown() bool {
	switch r {
	case AccountBrowserRenderingGetSnapshotParamsGotoOptionsWaitUntilStringLoad, AccountBrowserRenderingGetSnapshotParamsGotoOptionsWaitUntilStringDomcontentloaded, AccountBrowserRenderingGetSnapshotParamsGotoOptionsWaitUntilStringNetworkidle0, AccountBrowserRenderingGetSnapshotParamsGotoOptionsWaitUntilStringNetworkidle2:
		return true
	}
	return false
}

func (r AccountBrowserRenderingGetSnapshotParamsGotoOptionsWaitUntilString) implementsAccountBrowserRenderingGetSnapshotParamsGotoOptionsWaitUntilUnion() {
}

type AccountBrowserRenderingGetSnapshotParamsGotoOptionsWaitUntilArray []AccountBrowserRenderingGetSnapshotParamsGotoOptionsWaitUntilArrayItem

func (r AccountBrowserRenderingGetSnapshotParamsGotoOptionsWaitUntilArray) implementsAccountBrowserRenderingGetSnapshotParamsGotoOptionsWaitUntilUnion() {
}

type AccountBrowserRenderingGetSnapshotParamsGotoOptionsWaitUntilArrayItem string

const (
	AccountBrowserRenderingGetSnapshotParamsGotoOptionsWaitUntilArrayItemLoad             AccountBrowserRenderingGetSnapshotParamsGotoOptionsWaitUntilArrayItem = "load"
	AccountBrowserRenderingGetSnapshotParamsGotoOptionsWaitUntilArrayItemDomcontentloaded AccountBrowserRenderingGetSnapshotParamsGotoOptionsWaitUntilArrayItem = "domcontentloaded"
	AccountBrowserRenderingGetSnapshotParamsGotoOptionsWaitUntilArrayItemNetworkidle0     AccountBrowserRenderingGetSnapshotParamsGotoOptionsWaitUntilArrayItem = "networkidle0"
	AccountBrowserRenderingGetSnapshotParamsGotoOptionsWaitUntilArrayItemNetworkidle2     AccountBrowserRenderingGetSnapshotParamsGotoOptionsWaitUntilArrayItem = "networkidle2"
)

func (r AccountBrowserRenderingGetSnapshotParamsGotoOptionsWaitUntilArrayItem) IsKnown() bool {
	switch r {
	case AccountBrowserRenderingGetSnapshotParamsGotoOptionsWaitUntilArrayItemLoad, AccountBrowserRenderingGetSnapshotParamsGotoOptionsWaitUntilArrayItemDomcontentloaded, AccountBrowserRenderingGetSnapshotParamsGotoOptionsWaitUntilArrayItemNetworkidle0, AccountBrowserRenderingGetSnapshotParamsGotoOptionsWaitUntilArrayItemNetworkidle2:
		return true
	}
	return false
}

type AccountBrowserRenderingGetSnapshotParamsRejectResourceType string

const (
	AccountBrowserRenderingGetSnapshotParamsRejectResourceTypeDocument           AccountBrowserRenderingGetSnapshotParamsRejectResourceType = "document"
	AccountBrowserRenderingGetSnapshotParamsRejectResourceTypeStylesheet         AccountBrowserRenderingGetSnapshotParamsRejectResourceType = "stylesheet"
	AccountBrowserRenderingGetSnapshotParamsRejectResourceTypeImage              AccountBrowserRenderingGetSnapshotParamsRejectResourceType = "image"
	AccountBrowserRenderingGetSnapshotParamsRejectResourceTypeMedia              AccountBrowserRenderingGetSnapshotParamsRejectResourceType = "media"
	AccountBrowserRenderingGetSnapshotParamsRejectResourceTypeFont               AccountBrowserRenderingGetSnapshotParamsRejectResourceType = "font"
	AccountBrowserRenderingGetSnapshotParamsRejectResourceTypeScript             AccountBrowserRenderingGetSnapshotParamsRejectResourceType = "script"
	AccountBrowserRenderingGetSnapshotParamsRejectResourceTypeTexttrack          AccountBrowserRenderingGetSnapshotParamsRejectResourceType = "texttrack"
	AccountBrowserRenderingGetSnapshotParamsRejectResourceTypeXhr                AccountBrowserRenderingGetSnapshotParamsRejectResourceType = "xhr"
	AccountBrowserRenderingGetSnapshotParamsRejectResourceTypeFetch              AccountBrowserRenderingGetSnapshotParamsRejectResourceType = "fetch"
	AccountBrowserRenderingGetSnapshotParamsRejectResourceTypePrefetch           AccountBrowserRenderingGetSnapshotParamsRejectResourceType = "prefetch"
	AccountBrowserRenderingGetSnapshotParamsRejectResourceTypeEventsource        AccountBrowserRenderingGetSnapshotParamsRejectResourceType = "eventsource"
	AccountBrowserRenderingGetSnapshotParamsRejectResourceTypeWebsocket          AccountBrowserRenderingGetSnapshotParamsRejectResourceType = "websocket"
	AccountBrowserRenderingGetSnapshotParamsRejectResourceTypeManifest           AccountBrowserRenderingGetSnapshotParamsRejectResourceType = "manifest"
	AccountBrowserRenderingGetSnapshotParamsRejectResourceTypeSignedexchange     AccountBrowserRenderingGetSnapshotParamsRejectResourceType = "signedexchange"
	AccountBrowserRenderingGetSnapshotParamsRejectResourceTypePing               AccountBrowserRenderingGetSnapshotParamsRejectResourceType = "ping"
	AccountBrowserRenderingGetSnapshotParamsRejectResourceTypeCspviolationreport AccountBrowserRenderingGetSnapshotParamsRejectResourceType = "cspviolationreport"
	AccountBrowserRenderingGetSnapshotParamsRejectResourceTypePreflight          AccountBrowserRenderingGetSnapshotParamsRejectResourceType = "preflight"
	AccountBrowserRenderingGetSnapshotParamsRejectResourceTypeOther              AccountBrowserRenderingGetSnapshotParamsRejectResourceType = "other"
)

func (r AccountBrowserRenderingGetSnapshotParamsRejectResourceType) IsKnown() bool {
	switch r {
	case AccountBrowserRenderingGetSnapshotParamsRejectResourceTypeDocument, AccountBrowserRenderingGetSnapshotParamsRejectResourceTypeStylesheet, AccountBrowserRenderingGetSnapshotParamsRejectResourceTypeImage, AccountBrowserRenderingGetSnapshotParamsRejectResourceTypeMedia, AccountBrowserRenderingGetSnapshotParamsRejectResourceTypeFont, AccountBrowserRenderingGetSnapshotParamsRejectResourceTypeScript, AccountBrowserRenderingGetSnapshotParamsRejectResourceTypeTexttrack, AccountBrowserRenderingGetSnapshotParamsRejectResourceTypeXhr, AccountBrowserRenderingGetSnapshotParamsRejectResourceTypeFetch, AccountBrowserRenderingGetSnapshotParamsRejectResourceTypePrefetch, AccountBrowserRenderingGetSnapshotParamsRejectResourceTypeEventsource, AccountBrowserRenderingGetSnapshotParamsRejectResourceTypeWebsocket, AccountBrowserRenderingGetSnapshotParamsRejectResourceTypeManifest, AccountBrowserRenderingGetSnapshotParamsRejectResourceTypeSignedexchange, AccountBrowserRenderingGetSnapshotParamsRejectResourceTypePing, AccountBrowserRenderingGetSnapshotParamsRejectResourceTypeCspviolationreport, AccountBrowserRenderingGetSnapshotParamsRejectResourceTypePreflight, AccountBrowserRenderingGetSnapshotParamsRejectResourceTypeOther:
		return true
	}
	return false
}

type AccountBrowserRenderingGetSnapshotParamsScreenshotOptions struct {
	CaptureBeyondViewport param.Field[bool]                                                          `json:"captureBeyondViewport"`
	Clip                  param.Field[AccountBrowserRenderingGetSnapshotParamsScreenshotOptionsClip] `json:"clip"`
	FromSurface           param.Field[bool]                                                          `json:"fromSurface"`
	FullPage              param.Field[bool]                                                          `json:"fullPage"`
	OmitBackground        param.Field[bool]                                                          `json:"omitBackground"`
	OptimizeForSpeed      param.Field[bool]                                                          `json:"optimizeForSpeed"`
	Quality               param.Field[float64]                                                       `json:"quality"`
	Type                  param.Field[AccountBrowserRenderingGetSnapshotParamsScreenshotOptionsType] `json:"type"`
}

func (r AccountBrowserRenderingGetSnapshotParamsScreenshotOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountBrowserRenderingGetSnapshotParamsScreenshotOptionsClip struct {
	Height param.Field[float64] `json:"height,required"`
	Width  param.Field[float64] `json:"width,required"`
	X      param.Field[float64] `json:"x,required"`
	Y      param.Field[float64] `json:"y,required"`
	Scale  param.Field[float64] `json:"scale"`
}

func (r AccountBrowserRenderingGetSnapshotParamsScreenshotOptionsClip) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountBrowserRenderingGetSnapshotParamsScreenshotOptionsType string

const (
	AccountBrowserRenderingGetSnapshotParamsScreenshotOptionsTypePng  AccountBrowserRenderingGetSnapshotParamsScreenshotOptionsType = "png"
	AccountBrowserRenderingGetSnapshotParamsScreenshotOptionsTypeJpeg AccountBrowserRenderingGetSnapshotParamsScreenshotOptionsType = "jpeg"
	AccountBrowserRenderingGetSnapshotParamsScreenshotOptionsTypeWebp AccountBrowserRenderingGetSnapshotParamsScreenshotOptionsType = "webp"
)

func (r AccountBrowserRenderingGetSnapshotParamsScreenshotOptionsType) IsKnown() bool {
	switch r {
	case AccountBrowserRenderingGetSnapshotParamsScreenshotOptionsTypePng, AccountBrowserRenderingGetSnapshotParamsScreenshotOptionsTypeJpeg, AccountBrowserRenderingGetSnapshotParamsScreenshotOptionsTypeWebp:
		return true
	}
	return false
}

// Check [options](https://pptr.dev/api/puppeteer.page.setviewport).
type AccountBrowserRenderingGetSnapshotParamsViewport struct {
	Height            param.Field[float64] `json:"height,required"`
	Width             param.Field[float64] `json:"width,required"`
	DeviceScaleFactor param.Field[float64] `json:"deviceScaleFactor"`
	HasTouch          param.Field[bool]    `json:"hasTouch"`
	IsLandscape       param.Field[bool]    `json:"isLandscape"`
	IsMobile          param.Field[bool]    `json:"isMobile"`
}

func (r AccountBrowserRenderingGetSnapshotParamsViewport) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Wait for the selector to appear in page. Check
// [options](https://pptr.dev/api/puppeteer.page.waitforselector).
type AccountBrowserRenderingGetSnapshotParamsWaitForSelector struct {
	Selector param.Field[string]                                                         `json:"selector,required"`
	Hidden   param.Field[AccountBrowserRenderingGetSnapshotParamsWaitForSelectorHidden]  `json:"hidden"`
	Timeout  param.Field[float64]                                                        `json:"timeout"`
	Visible  param.Field[AccountBrowserRenderingGetSnapshotParamsWaitForSelectorVisible] `json:"visible"`
}

func (r AccountBrowserRenderingGetSnapshotParamsWaitForSelector) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountBrowserRenderingGetSnapshotParamsWaitForSelectorHidden bool

const (
	AccountBrowserRenderingGetSnapshotParamsWaitForSelectorHiddenTrue AccountBrowserRenderingGetSnapshotParamsWaitForSelectorHidden = true
)

func (r AccountBrowserRenderingGetSnapshotParamsWaitForSelectorHidden) IsKnown() bool {
	switch r {
	case AccountBrowserRenderingGetSnapshotParamsWaitForSelectorHiddenTrue:
		return true
	}
	return false
}

type AccountBrowserRenderingGetSnapshotParamsWaitForSelectorVisible bool

const (
	AccountBrowserRenderingGetSnapshotParamsWaitForSelectorVisibleTrue AccountBrowserRenderingGetSnapshotParamsWaitForSelectorVisible = true
)

func (r AccountBrowserRenderingGetSnapshotParamsWaitForSelectorVisible) IsKnown() bool {
	switch r {
	case AccountBrowserRenderingGetSnapshotParamsWaitForSelectorVisibleTrue:
		return true
	}
	return false
}

type AccountBrowserRenderingScrapeElementsParams struct {
	Elements param.Field[[]AccountBrowserRenderingScrapeElementsParamsElement] `json:"elements,required"`
	// Cache TTL default is 5s. Set to 0 to disable.
	CacheTtl param.Field[float64] `query:"cacheTTL"`
	// Adds a `<script>` tag into the page with the desired URL or content.
	AddScriptTag param.Field[[]AccountBrowserRenderingScrapeElementsParamsAddScriptTag] `json:"addScriptTag"`
	// Adds a `<link rel="stylesheet">` tag into the page with the desired URL or a
	// `<style type="text/css">` tag with the content.
	AddStyleTag param.Field[[]AccountBrowserRenderingScrapeElementsParamsAddStyleTag] `json:"addStyleTag"`
	// Only allow requests that match the provided regex patterns, eg. '/^.\*\.(css)'.
	AllowRequestPattern param.Field[[]string] `json:"allowRequestPattern"`
	// Only allow requests that match the provided resource types, eg. 'image' or
	// 'script'.
	AllowResourceTypes param.Field[[]AccountBrowserRenderingScrapeElementsParamsAllowResourceType] `json:"allowResourceTypes"`
	// Provide credentials for HTTP authentication.
	Authenticate param.Field[AccountBrowserRenderingScrapeElementsParamsAuthenticate] `json:"authenticate"`
	// Attempt to proceed when 'awaited' events fail or timeout.
	BestAttempt param.Field[bool] `json:"bestAttempt"`
	// Check [options](https://pptr.dev/api/puppeteer.page.setcookie).
	Cookies          param.Field[[]AccountBrowserRenderingScrapeElementsParamsCookie] `json:"cookies"`
	EmulateMediaType param.Field[string]                                              `json:"emulateMediaType"`
	// Check [options](https://pptr.dev/api/puppeteer.gotooptions).
	GotoOptions param.Field[AccountBrowserRenderingScrapeElementsParamsGotoOptions] `json:"gotoOptions"`
	// Set the content of the page, eg: `<h1>Hello World!!</h1>`. Either `html` or
	// `url` must be set.
	HTML param.Field[string] `json:"html"`
	// Block undesired requests that match the provided regex patterns, eg.
	// '/^.\*\.(css)'.
	RejectRequestPattern param.Field[[]string] `json:"rejectRequestPattern"`
	// Block undesired requests that match the provided resource types, eg. 'image' or
	// 'script'.
	RejectResourceTypes  param.Field[[]AccountBrowserRenderingScrapeElementsParamsRejectResourceType] `json:"rejectResourceTypes"`
	SetExtraHTTPHeaders  param.Field[map[string]string]                                               `json:"setExtraHTTPHeaders"`
	SetJavaScriptEnabled param.Field[bool]                                                            `json:"setJavaScriptEnabled"`
	// URL to navigate to, eg. `https://example.com`.
	URL       param.Field[string] `json:"url" format:"uri"`
	UserAgent param.Field[string] `json:"userAgent"`
	// Check [options](https://pptr.dev/api/puppeteer.page.setviewport).
	Viewport param.Field[AccountBrowserRenderingScrapeElementsParamsViewport] `json:"viewport"`
	// Wait for the selector to appear in page. Check
	// [options](https://pptr.dev/api/puppeteer.page.waitforselector).
	WaitForSelector param.Field[AccountBrowserRenderingScrapeElementsParamsWaitForSelector] `json:"waitForSelector"`
	// Waits for a specified timeout before continuing.
	WaitForTimeout param.Field[float64] `json:"waitForTimeout"`
}

func (r AccountBrowserRenderingScrapeElementsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [AccountBrowserRenderingScrapeElementsParams]'s query
// parameters as `url.Values`.
func (r AccountBrowserRenderingScrapeElementsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountBrowserRenderingScrapeElementsParamsElement struct {
	Selector param.Field[string] `json:"selector,required"`
}

func (r AccountBrowserRenderingScrapeElementsParamsElement) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountBrowserRenderingScrapeElementsParamsAddScriptTag struct {
	ID      param.Field[string] `json:"id"`
	Content param.Field[string] `json:"content"`
	Type    param.Field[string] `json:"type"`
	URL     param.Field[string] `json:"url"`
}

func (r AccountBrowserRenderingScrapeElementsParamsAddScriptTag) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountBrowserRenderingScrapeElementsParamsAddStyleTag struct {
	Content param.Field[string] `json:"content"`
	URL     param.Field[string] `json:"url"`
}

func (r AccountBrowserRenderingScrapeElementsParamsAddStyleTag) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountBrowserRenderingScrapeElementsParamsAllowResourceType string

const (
	AccountBrowserRenderingScrapeElementsParamsAllowResourceTypeDocument           AccountBrowserRenderingScrapeElementsParamsAllowResourceType = "document"
	AccountBrowserRenderingScrapeElementsParamsAllowResourceTypeStylesheet         AccountBrowserRenderingScrapeElementsParamsAllowResourceType = "stylesheet"
	AccountBrowserRenderingScrapeElementsParamsAllowResourceTypeImage              AccountBrowserRenderingScrapeElementsParamsAllowResourceType = "image"
	AccountBrowserRenderingScrapeElementsParamsAllowResourceTypeMedia              AccountBrowserRenderingScrapeElementsParamsAllowResourceType = "media"
	AccountBrowserRenderingScrapeElementsParamsAllowResourceTypeFont               AccountBrowserRenderingScrapeElementsParamsAllowResourceType = "font"
	AccountBrowserRenderingScrapeElementsParamsAllowResourceTypeScript             AccountBrowserRenderingScrapeElementsParamsAllowResourceType = "script"
	AccountBrowserRenderingScrapeElementsParamsAllowResourceTypeTexttrack          AccountBrowserRenderingScrapeElementsParamsAllowResourceType = "texttrack"
	AccountBrowserRenderingScrapeElementsParamsAllowResourceTypeXhr                AccountBrowserRenderingScrapeElementsParamsAllowResourceType = "xhr"
	AccountBrowserRenderingScrapeElementsParamsAllowResourceTypeFetch              AccountBrowserRenderingScrapeElementsParamsAllowResourceType = "fetch"
	AccountBrowserRenderingScrapeElementsParamsAllowResourceTypePrefetch           AccountBrowserRenderingScrapeElementsParamsAllowResourceType = "prefetch"
	AccountBrowserRenderingScrapeElementsParamsAllowResourceTypeEventsource        AccountBrowserRenderingScrapeElementsParamsAllowResourceType = "eventsource"
	AccountBrowserRenderingScrapeElementsParamsAllowResourceTypeWebsocket          AccountBrowserRenderingScrapeElementsParamsAllowResourceType = "websocket"
	AccountBrowserRenderingScrapeElementsParamsAllowResourceTypeManifest           AccountBrowserRenderingScrapeElementsParamsAllowResourceType = "manifest"
	AccountBrowserRenderingScrapeElementsParamsAllowResourceTypeSignedexchange     AccountBrowserRenderingScrapeElementsParamsAllowResourceType = "signedexchange"
	AccountBrowserRenderingScrapeElementsParamsAllowResourceTypePing               AccountBrowserRenderingScrapeElementsParamsAllowResourceType = "ping"
	AccountBrowserRenderingScrapeElementsParamsAllowResourceTypeCspviolationreport AccountBrowserRenderingScrapeElementsParamsAllowResourceType = "cspviolationreport"
	AccountBrowserRenderingScrapeElementsParamsAllowResourceTypePreflight          AccountBrowserRenderingScrapeElementsParamsAllowResourceType = "preflight"
	AccountBrowserRenderingScrapeElementsParamsAllowResourceTypeOther              AccountBrowserRenderingScrapeElementsParamsAllowResourceType = "other"
)

func (r AccountBrowserRenderingScrapeElementsParamsAllowResourceType) IsKnown() bool {
	switch r {
	case AccountBrowserRenderingScrapeElementsParamsAllowResourceTypeDocument, AccountBrowserRenderingScrapeElementsParamsAllowResourceTypeStylesheet, AccountBrowserRenderingScrapeElementsParamsAllowResourceTypeImage, AccountBrowserRenderingScrapeElementsParamsAllowResourceTypeMedia, AccountBrowserRenderingScrapeElementsParamsAllowResourceTypeFont, AccountBrowserRenderingScrapeElementsParamsAllowResourceTypeScript, AccountBrowserRenderingScrapeElementsParamsAllowResourceTypeTexttrack, AccountBrowserRenderingScrapeElementsParamsAllowResourceTypeXhr, AccountBrowserRenderingScrapeElementsParamsAllowResourceTypeFetch, AccountBrowserRenderingScrapeElementsParamsAllowResourceTypePrefetch, AccountBrowserRenderingScrapeElementsParamsAllowResourceTypeEventsource, AccountBrowserRenderingScrapeElementsParamsAllowResourceTypeWebsocket, AccountBrowserRenderingScrapeElementsParamsAllowResourceTypeManifest, AccountBrowserRenderingScrapeElementsParamsAllowResourceTypeSignedexchange, AccountBrowserRenderingScrapeElementsParamsAllowResourceTypePing, AccountBrowserRenderingScrapeElementsParamsAllowResourceTypeCspviolationreport, AccountBrowserRenderingScrapeElementsParamsAllowResourceTypePreflight, AccountBrowserRenderingScrapeElementsParamsAllowResourceTypeOther:
		return true
	}
	return false
}

// Provide credentials for HTTP authentication.
type AccountBrowserRenderingScrapeElementsParamsAuthenticate struct {
	Password param.Field[string] `json:"password,required"`
	Username param.Field[string] `json:"username,required"`
}

func (r AccountBrowserRenderingScrapeElementsParamsAuthenticate) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountBrowserRenderingScrapeElementsParamsCookie struct {
	Name         param.Field[string]                                                         `json:"name,required"`
	Value        param.Field[string]                                                         `json:"value,required"`
	Domain       param.Field[string]                                                         `json:"domain"`
	Expires      param.Field[float64]                                                        `json:"expires"`
	HTTPOnly     param.Field[bool]                                                           `json:"httpOnly"`
	PartitionKey param.Field[string]                                                         `json:"partitionKey"`
	Path         param.Field[string]                                                         `json:"path"`
	Priority     param.Field[AccountBrowserRenderingScrapeElementsParamsCookiesPriority]     `json:"priority"`
	SameParty    param.Field[bool]                                                           `json:"sameParty"`
	SameSite     param.Field[AccountBrowserRenderingScrapeElementsParamsCookiesSameSite]     `json:"sameSite"`
	Secure       param.Field[bool]                                                           `json:"secure"`
	SourcePort   param.Field[float64]                                                        `json:"sourcePort"`
	SourceScheme param.Field[AccountBrowserRenderingScrapeElementsParamsCookiesSourceScheme] `json:"sourceScheme"`
	URL          param.Field[string]                                                         `json:"url"`
}

func (r AccountBrowserRenderingScrapeElementsParamsCookie) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountBrowserRenderingScrapeElementsParamsCookiesPriority string

const (
	AccountBrowserRenderingScrapeElementsParamsCookiesPriorityLow    AccountBrowserRenderingScrapeElementsParamsCookiesPriority = "Low"
	AccountBrowserRenderingScrapeElementsParamsCookiesPriorityMedium AccountBrowserRenderingScrapeElementsParamsCookiesPriority = "Medium"
	AccountBrowserRenderingScrapeElementsParamsCookiesPriorityHigh   AccountBrowserRenderingScrapeElementsParamsCookiesPriority = "High"
)

func (r AccountBrowserRenderingScrapeElementsParamsCookiesPriority) IsKnown() bool {
	switch r {
	case AccountBrowserRenderingScrapeElementsParamsCookiesPriorityLow, AccountBrowserRenderingScrapeElementsParamsCookiesPriorityMedium, AccountBrowserRenderingScrapeElementsParamsCookiesPriorityHigh:
		return true
	}
	return false
}

type AccountBrowserRenderingScrapeElementsParamsCookiesSameSite string

const (
	AccountBrowserRenderingScrapeElementsParamsCookiesSameSiteStrict AccountBrowserRenderingScrapeElementsParamsCookiesSameSite = "Strict"
	AccountBrowserRenderingScrapeElementsParamsCookiesSameSiteLax    AccountBrowserRenderingScrapeElementsParamsCookiesSameSite = "Lax"
	AccountBrowserRenderingScrapeElementsParamsCookiesSameSiteNone   AccountBrowserRenderingScrapeElementsParamsCookiesSameSite = "None"
)

func (r AccountBrowserRenderingScrapeElementsParamsCookiesSameSite) IsKnown() bool {
	switch r {
	case AccountBrowserRenderingScrapeElementsParamsCookiesSameSiteStrict, AccountBrowserRenderingScrapeElementsParamsCookiesSameSiteLax, AccountBrowserRenderingScrapeElementsParamsCookiesSameSiteNone:
		return true
	}
	return false
}

type AccountBrowserRenderingScrapeElementsParamsCookiesSourceScheme string

const (
	AccountBrowserRenderingScrapeElementsParamsCookiesSourceSchemeUnset     AccountBrowserRenderingScrapeElementsParamsCookiesSourceScheme = "Unset"
	AccountBrowserRenderingScrapeElementsParamsCookiesSourceSchemeNonSecure AccountBrowserRenderingScrapeElementsParamsCookiesSourceScheme = "NonSecure"
	AccountBrowserRenderingScrapeElementsParamsCookiesSourceSchemeSecure    AccountBrowserRenderingScrapeElementsParamsCookiesSourceScheme = "Secure"
)

func (r AccountBrowserRenderingScrapeElementsParamsCookiesSourceScheme) IsKnown() bool {
	switch r {
	case AccountBrowserRenderingScrapeElementsParamsCookiesSourceSchemeUnset, AccountBrowserRenderingScrapeElementsParamsCookiesSourceSchemeNonSecure, AccountBrowserRenderingScrapeElementsParamsCookiesSourceSchemeSecure:
		return true
	}
	return false
}

// Check [options](https://pptr.dev/api/puppeteer.gotooptions).
type AccountBrowserRenderingScrapeElementsParamsGotoOptions struct {
	Referer        param.Field[string]                                                               `json:"referer"`
	ReferrerPolicy param.Field[string]                                                               `json:"referrerPolicy"`
	Timeout        param.Field[float64]                                                              `json:"timeout"`
	WaitUntil      param.Field[AccountBrowserRenderingScrapeElementsParamsGotoOptionsWaitUntilUnion] `json:"waitUntil"`
}

func (r AccountBrowserRenderingScrapeElementsParamsGotoOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by
// [AccountBrowserRenderingScrapeElementsParamsGotoOptionsWaitUntilString],
// [AccountBrowserRenderingScrapeElementsParamsGotoOptionsWaitUntilArray].
type AccountBrowserRenderingScrapeElementsParamsGotoOptionsWaitUntilUnion interface {
	implementsAccountBrowserRenderingScrapeElementsParamsGotoOptionsWaitUntilUnion()
}

type AccountBrowserRenderingScrapeElementsParamsGotoOptionsWaitUntilString string

const (
	AccountBrowserRenderingScrapeElementsParamsGotoOptionsWaitUntilStringLoad             AccountBrowserRenderingScrapeElementsParamsGotoOptionsWaitUntilString = "load"
	AccountBrowserRenderingScrapeElementsParamsGotoOptionsWaitUntilStringDomcontentloaded AccountBrowserRenderingScrapeElementsParamsGotoOptionsWaitUntilString = "domcontentloaded"
	AccountBrowserRenderingScrapeElementsParamsGotoOptionsWaitUntilStringNetworkidle0     AccountBrowserRenderingScrapeElementsParamsGotoOptionsWaitUntilString = "networkidle0"
	AccountBrowserRenderingScrapeElementsParamsGotoOptionsWaitUntilStringNetworkidle2     AccountBrowserRenderingScrapeElementsParamsGotoOptionsWaitUntilString = "networkidle2"
)

func (r AccountBrowserRenderingScrapeElementsParamsGotoOptionsWaitUntilString) IsKnown() bool {
	switch r {
	case AccountBrowserRenderingScrapeElementsParamsGotoOptionsWaitUntilStringLoad, AccountBrowserRenderingScrapeElementsParamsGotoOptionsWaitUntilStringDomcontentloaded, AccountBrowserRenderingScrapeElementsParamsGotoOptionsWaitUntilStringNetworkidle0, AccountBrowserRenderingScrapeElementsParamsGotoOptionsWaitUntilStringNetworkidle2:
		return true
	}
	return false
}

func (r AccountBrowserRenderingScrapeElementsParamsGotoOptionsWaitUntilString) implementsAccountBrowserRenderingScrapeElementsParamsGotoOptionsWaitUntilUnion() {
}

type AccountBrowserRenderingScrapeElementsParamsGotoOptionsWaitUntilArray []AccountBrowserRenderingScrapeElementsParamsGotoOptionsWaitUntilArrayItem

func (r AccountBrowserRenderingScrapeElementsParamsGotoOptionsWaitUntilArray) implementsAccountBrowserRenderingScrapeElementsParamsGotoOptionsWaitUntilUnion() {
}

type AccountBrowserRenderingScrapeElementsParamsGotoOptionsWaitUntilArrayItem string

const (
	AccountBrowserRenderingScrapeElementsParamsGotoOptionsWaitUntilArrayItemLoad             AccountBrowserRenderingScrapeElementsParamsGotoOptionsWaitUntilArrayItem = "load"
	AccountBrowserRenderingScrapeElementsParamsGotoOptionsWaitUntilArrayItemDomcontentloaded AccountBrowserRenderingScrapeElementsParamsGotoOptionsWaitUntilArrayItem = "domcontentloaded"
	AccountBrowserRenderingScrapeElementsParamsGotoOptionsWaitUntilArrayItemNetworkidle0     AccountBrowserRenderingScrapeElementsParamsGotoOptionsWaitUntilArrayItem = "networkidle0"
	AccountBrowserRenderingScrapeElementsParamsGotoOptionsWaitUntilArrayItemNetworkidle2     AccountBrowserRenderingScrapeElementsParamsGotoOptionsWaitUntilArrayItem = "networkidle2"
)

func (r AccountBrowserRenderingScrapeElementsParamsGotoOptionsWaitUntilArrayItem) IsKnown() bool {
	switch r {
	case AccountBrowserRenderingScrapeElementsParamsGotoOptionsWaitUntilArrayItemLoad, AccountBrowserRenderingScrapeElementsParamsGotoOptionsWaitUntilArrayItemDomcontentloaded, AccountBrowserRenderingScrapeElementsParamsGotoOptionsWaitUntilArrayItemNetworkidle0, AccountBrowserRenderingScrapeElementsParamsGotoOptionsWaitUntilArrayItemNetworkidle2:
		return true
	}
	return false
}

type AccountBrowserRenderingScrapeElementsParamsRejectResourceType string

const (
	AccountBrowserRenderingScrapeElementsParamsRejectResourceTypeDocument           AccountBrowserRenderingScrapeElementsParamsRejectResourceType = "document"
	AccountBrowserRenderingScrapeElementsParamsRejectResourceTypeStylesheet         AccountBrowserRenderingScrapeElementsParamsRejectResourceType = "stylesheet"
	AccountBrowserRenderingScrapeElementsParamsRejectResourceTypeImage              AccountBrowserRenderingScrapeElementsParamsRejectResourceType = "image"
	AccountBrowserRenderingScrapeElementsParamsRejectResourceTypeMedia              AccountBrowserRenderingScrapeElementsParamsRejectResourceType = "media"
	AccountBrowserRenderingScrapeElementsParamsRejectResourceTypeFont               AccountBrowserRenderingScrapeElementsParamsRejectResourceType = "font"
	AccountBrowserRenderingScrapeElementsParamsRejectResourceTypeScript             AccountBrowserRenderingScrapeElementsParamsRejectResourceType = "script"
	AccountBrowserRenderingScrapeElementsParamsRejectResourceTypeTexttrack          AccountBrowserRenderingScrapeElementsParamsRejectResourceType = "texttrack"
	AccountBrowserRenderingScrapeElementsParamsRejectResourceTypeXhr                AccountBrowserRenderingScrapeElementsParamsRejectResourceType = "xhr"
	AccountBrowserRenderingScrapeElementsParamsRejectResourceTypeFetch              AccountBrowserRenderingScrapeElementsParamsRejectResourceType = "fetch"
	AccountBrowserRenderingScrapeElementsParamsRejectResourceTypePrefetch           AccountBrowserRenderingScrapeElementsParamsRejectResourceType = "prefetch"
	AccountBrowserRenderingScrapeElementsParamsRejectResourceTypeEventsource        AccountBrowserRenderingScrapeElementsParamsRejectResourceType = "eventsource"
	AccountBrowserRenderingScrapeElementsParamsRejectResourceTypeWebsocket          AccountBrowserRenderingScrapeElementsParamsRejectResourceType = "websocket"
	AccountBrowserRenderingScrapeElementsParamsRejectResourceTypeManifest           AccountBrowserRenderingScrapeElementsParamsRejectResourceType = "manifest"
	AccountBrowserRenderingScrapeElementsParamsRejectResourceTypeSignedexchange     AccountBrowserRenderingScrapeElementsParamsRejectResourceType = "signedexchange"
	AccountBrowserRenderingScrapeElementsParamsRejectResourceTypePing               AccountBrowserRenderingScrapeElementsParamsRejectResourceType = "ping"
	AccountBrowserRenderingScrapeElementsParamsRejectResourceTypeCspviolationreport AccountBrowserRenderingScrapeElementsParamsRejectResourceType = "cspviolationreport"
	AccountBrowserRenderingScrapeElementsParamsRejectResourceTypePreflight          AccountBrowserRenderingScrapeElementsParamsRejectResourceType = "preflight"
	AccountBrowserRenderingScrapeElementsParamsRejectResourceTypeOther              AccountBrowserRenderingScrapeElementsParamsRejectResourceType = "other"
)

func (r AccountBrowserRenderingScrapeElementsParamsRejectResourceType) IsKnown() bool {
	switch r {
	case AccountBrowserRenderingScrapeElementsParamsRejectResourceTypeDocument, AccountBrowserRenderingScrapeElementsParamsRejectResourceTypeStylesheet, AccountBrowserRenderingScrapeElementsParamsRejectResourceTypeImage, AccountBrowserRenderingScrapeElementsParamsRejectResourceTypeMedia, AccountBrowserRenderingScrapeElementsParamsRejectResourceTypeFont, AccountBrowserRenderingScrapeElementsParamsRejectResourceTypeScript, AccountBrowserRenderingScrapeElementsParamsRejectResourceTypeTexttrack, AccountBrowserRenderingScrapeElementsParamsRejectResourceTypeXhr, AccountBrowserRenderingScrapeElementsParamsRejectResourceTypeFetch, AccountBrowserRenderingScrapeElementsParamsRejectResourceTypePrefetch, AccountBrowserRenderingScrapeElementsParamsRejectResourceTypeEventsource, AccountBrowserRenderingScrapeElementsParamsRejectResourceTypeWebsocket, AccountBrowserRenderingScrapeElementsParamsRejectResourceTypeManifest, AccountBrowserRenderingScrapeElementsParamsRejectResourceTypeSignedexchange, AccountBrowserRenderingScrapeElementsParamsRejectResourceTypePing, AccountBrowserRenderingScrapeElementsParamsRejectResourceTypeCspviolationreport, AccountBrowserRenderingScrapeElementsParamsRejectResourceTypePreflight, AccountBrowserRenderingScrapeElementsParamsRejectResourceTypeOther:
		return true
	}
	return false
}

// Check [options](https://pptr.dev/api/puppeteer.page.setviewport).
type AccountBrowserRenderingScrapeElementsParamsViewport struct {
	Height            param.Field[float64] `json:"height,required"`
	Width             param.Field[float64] `json:"width,required"`
	DeviceScaleFactor param.Field[float64] `json:"deviceScaleFactor"`
	HasTouch          param.Field[bool]    `json:"hasTouch"`
	IsLandscape       param.Field[bool]    `json:"isLandscape"`
	IsMobile          param.Field[bool]    `json:"isMobile"`
}

func (r AccountBrowserRenderingScrapeElementsParamsViewport) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Wait for the selector to appear in page. Check
// [options](https://pptr.dev/api/puppeteer.page.waitforselector).
type AccountBrowserRenderingScrapeElementsParamsWaitForSelector struct {
	Selector param.Field[string]                                                            `json:"selector,required"`
	Hidden   param.Field[AccountBrowserRenderingScrapeElementsParamsWaitForSelectorHidden]  `json:"hidden"`
	Timeout  param.Field[float64]                                                           `json:"timeout"`
	Visible  param.Field[AccountBrowserRenderingScrapeElementsParamsWaitForSelectorVisible] `json:"visible"`
}

func (r AccountBrowserRenderingScrapeElementsParamsWaitForSelector) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountBrowserRenderingScrapeElementsParamsWaitForSelectorHidden bool

const (
	AccountBrowserRenderingScrapeElementsParamsWaitForSelectorHiddenTrue AccountBrowserRenderingScrapeElementsParamsWaitForSelectorHidden = true
)

func (r AccountBrowserRenderingScrapeElementsParamsWaitForSelectorHidden) IsKnown() bool {
	switch r {
	case AccountBrowserRenderingScrapeElementsParamsWaitForSelectorHiddenTrue:
		return true
	}
	return false
}

type AccountBrowserRenderingScrapeElementsParamsWaitForSelectorVisible bool

const (
	AccountBrowserRenderingScrapeElementsParamsWaitForSelectorVisibleTrue AccountBrowserRenderingScrapeElementsParamsWaitForSelectorVisible = true
)

func (r AccountBrowserRenderingScrapeElementsParamsWaitForSelectorVisible) IsKnown() bool {
	switch r {
	case AccountBrowserRenderingScrapeElementsParamsWaitForSelectorVisibleTrue:
		return true
	}
	return false
}
