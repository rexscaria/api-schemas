// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/apiquery"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
)

// AccountUrlscannerScanService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountUrlscannerScanService] method instead.
type AccountUrlscannerScanService struct {
	Options []option.RequestOption
}

// NewAccountUrlscannerScanService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountUrlscannerScanService(opts ...option.RequestOption) (r *AccountUrlscannerScanService) {
	r = &AccountUrlscannerScanService{}
	r.Options = opts
	return
}

// Submit a URL to scan. You can also set some options, like the visibility level
// and custom headers. Check limits at
// https://developers.cloudflare.com/security-center/investigate/scan-limits/.
//
// Deprecated: Use
// [V2](https://developers.cloudflare.com/api/resources/url_scanner/subresources/scans/methods/create/)
// instead.
func (r *AccountUrlscannerScanService) New(ctx context.Context, accountID string, body AccountUrlscannerScanNewParams, opts ...option.RequestOption) (res *AccountUrlscannerScanNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/urlscanner/scan", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Get URL scan by uuid
//
// Deprecated: Use
// [V2](https://developers.cloudflare.com/api/resources/url_scanner/subresources/scans/methods/get/)
// instead.
func (r *AccountUrlscannerScanService) Get(ctx context.Context, accountID string, scanID string, query AccountUrlscannerScanGetParams, opts ...option.RequestOption) (res *AccountUrlscannerScanGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if scanID == "" {
		err = errors.New("missing required scan_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/urlscanner/scan/%s", accountID, scanID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Search scans by date and webpages' requests, including full URL (after
// redirects), hostname, and path. <br/> A successful scan will appear in search
// results a few minutes after finishing but may take much longer if the system in
// under load. By default, only successfully completed scans will appear in search
// results, unless searching by `scanId`. Please take into account that older scans
// may be removed from the search index at an unspecified time.
//
// Deprecated: Use
// [V2](https://developers.cloudflare.com/api/resources/url_scanner/subresources/scans/methods/list/)
// instead.
func (r *AccountUrlscannerScanService) List(ctx context.Context, accountID string, query AccountUrlscannerScanListParams, opts ...option.RequestOption) (res *AccountUrlscannerScanListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/urlscanner/scan", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Get a URL scan's HAR file. See HAR spec at
// http://www.softwareishard.com/blog/har-12-spec/.
//
// Deprecated: Use
// [V2](https://developers.cloudflare.com/api/resources/url_scanner/subresources/scans/methods/har/)
// instead.
func (r *AccountUrlscannerScanService) GetHar(ctx context.Context, accountID string, scanID string, opts ...option.RequestOption) (res *AccountUrlscannerScanGetHarResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if scanID == "" {
		err = errors.New("missing required scan_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/urlscanner/scan/%s/har", accountID, scanID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Get scan's screenshot by resolution (desktop/mobile/tablet).
//
// Deprecated: Use
// [V2](https://developers.cloudflare.com/api/resources/url_scanner/subresources/scans/methods/screenshot/)
// instead.
func (r *AccountUrlscannerScanService) GetScreenshot(ctx context.Context, accountID string, scanID string, query AccountUrlscannerScanGetScreenshotParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "image/png")}, opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if scanID == "" {
		err = errors.New("missing required scan_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/urlscanner/scan/%s/screenshot", accountID, scanID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type AccountUrlscannerScanNewResponse struct {
	Errors   []AccountUrlscannerScanNewResponseError   `json:"errors" api:"required"`
	Messages []AccountUrlscannerScanNewResponseMessage `json:"messages" api:"required"`
	Result   AccountUrlscannerScanNewResponseResult    `json:"result" api:"required"`
	Success  bool                                      `json:"success" api:"required"`
	JSON     accountUrlscannerScanNewResponseJSON      `json:"-"`
}

// accountUrlscannerScanNewResponseJSON contains the JSON metadata for the struct
// [AccountUrlscannerScanNewResponse]
type accountUrlscannerScanNewResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUrlscannerScanNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerScanNewResponseJSON) RawJSON() string {
	return r.raw
}

type AccountUrlscannerScanNewResponseError struct {
	Message string                                    `json:"message" api:"required"`
	JSON    accountUrlscannerScanNewResponseErrorJSON `json:"-"`
}

// accountUrlscannerScanNewResponseErrorJSON contains the JSON metadata for the
// struct [AccountUrlscannerScanNewResponseError]
type accountUrlscannerScanNewResponseErrorJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUrlscannerScanNewResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerScanNewResponseErrorJSON) RawJSON() string {
	return r.raw
}

type AccountUrlscannerScanNewResponseMessage struct {
	Message string                                      `json:"message" api:"required"`
	JSON    accountUrlscannerScanNewResponseMessageJSON `json:"-"`
}

// accountUrlscannerScanNewResponseMessageJSON contains the JSON metadata for the
// struct [AccountUrlscannerScanNewResponseMessage]
type accountUrlscannerScanNewResponseMessageJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUrlscannerScanNewResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerScanNewResponseMessageJSON) RawJSON() string {
	return r.raw
}

type AccountUrlscannerScanNewResponseResult struct {
	// Time when url was submitted for scanning.
	Time time.Time `json:"time" api:"required" format:"date-time"`
	// Canonical form of submitted URL. Use this if you want to later search by URL.
	URL string `json:"url" api:"required"`
	// Scan ID.
	Uuid string `json:"uuid" api:"required" format:"uuid"`
	// Submitted visibility status.
	Visibility AccountUrlscannerScanNewResponseResultVisibility `json:"visibility" api:"required"`
	JSON       accountUrlscannerScanNewResponseResultJSON       `json:"-"`
}

// accountUrlscannerScanNewResponseResultJSON contains the JSON metadata for the
// struct [AccountUrlscannerScanNewResponseResult]
type accountUrlscannerScanNewResponseResultJSON struct {
	Time        apijson.Field
	URL         apijson.Field
	Uuid        apijson.Field
	Visibility  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUrlscannerScanNewResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerScanNewResponseResultJSON) RawJSON() string {
	return r.raw
}

// Submitted visibility status.
type AccountUrlscannerScanNewResponseResultVisibility string

const (
	AccountUrlscannerScanNewResponseResultVisibilityPublic   AccountUrlscannerScanNewResponseResultVisibility = "public"
	AccountUrlscannerScanNewResponseResultVisibilityUnlisted AccountUrlscannerScanNewResponseResultVisibility = "unlisted"
)

func (r AccountUrlscannerScanNewResponseResultVisibility) IsKnown() bool {
	switch r {
	case AccountUrlscannerScanNewResponseResultVisibilityPublic, AccountUrlscannerScanNewResponseResultVisibilityUnlisted:
		return true
	}
	return false
}

type AccountUrlscannerScanGetResponse struct {
	Errors   []AccountUrlscannerScanGetResponseError   `json:"errors" api:"required"`
	Messages []AccountUrlscannerScanGetResponseMessage `json:"messages" api:"required"`
	Result   AccountUrlscannerScanGetResponseResult    `json:"result" api:"required"`
	// Whether request was successful or not
	Success bool                                 `json:"success" api:"required"`
	JSON    accountUrlscannerScanGetResponseJSON `json:"-"`
}

// accountUrlscannerScanGetResponseJSON contains the JSON metadata for the struct
// [AccountUrlscannerScanGetResponse]
type accountUrlscannerScanGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUrlscannerScanGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerScanGetResponseJSON) RawJSON() string {
	return r.raw
}

type AccountUrlscannerScanGetResponseError struct {
	Message string                                    `json:"message" api:"required"`
	JSON    accountUrlscannerScanGetResponseErrorJSON `json:"-"`
}

// accountUrlscannerScanGetResponseErrorJSON contains the JSON metadata for the
// struct [AccountUrlscannerScanGetResponseError]
type accountUrlscannerScanGetResponseErrorJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUrlscannerScanGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerScanGetResponseErrorJSON) RawJSON() string {
	return r.raw
}

type AccountUrlscannerScanGetResponseMessage struct {
	Message string                                      `json:"message" api:"required"`
	JSON    accountUrlscannerScanGetResponseMessageJSON `json:"-"`
}

// accountUrlscannerScanGetResponseMessageJSON contains the JSON metadata for the
// struct [AccountUrlscannerScanGetResponseMessage]
type accountUrlscannerScanGetResponseMessageJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUrlscannerScanGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerScanGetResponseMessageJSON) RawJSON() string {
	return r.raw
}

type AccountUrlscannerScanGetResponseResult struct {
	Scan AccountUrlscannerScanGetResponseResultScan `json:"scan" api:"required"`
	JSON accountUrlscannerScanGetResponseResultJSON `json:"-"`
}

// accountUrlscannerScanGetResponseResultJSON contains the JSON metadata for the
// struct [AccountUrlscannerScanGetResponseResult]
type accountUrlscannerScanGetResponseResultJSON struct {
	Scan        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUrlscannerScanGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerScanGetResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountUrlscannerScanGetResponseResultScan struct {
	Certificates []AccountUrlscannerScanGetResponseResultScanCertificate `json:"certificates" api:"required"`
	Geo          AccountUrlscannerScanGetResponseResultScanGeo           `json:"geo" api:"required"`
	Meta         AccountUrlscannerScanGetResponseResultScanMeta          `json:"meta" api:"required"`
	Page         AccountUrlscannerScanGetResponseResultScanPage          `json:"page" api:"required"`
	Performance  []AccountUrlscannerScanGetResponseResultScanPerformance `json:"performance" api:"required"`
	Task         AccountUrlscannerScanGetResponseResultScanTask          `json:"task" api:"required"`
	Verdicts     AccountUrlscannerScanGetResponseResultScanVerdicts      `json:"verdicts" api:"required"`
	// Dictionary of Autonomous System Numbers where ASN's are the keys
	Asns    AccountUrlscannerScanGetResponseResultScanAsns    `json:"asns"`
	Domains AccountUrlscannerScanGetResponseResultScanDomains `json:"domains"`
	IPs     AccountUrlscannerScanGetResponseResultScanIPs     `json:"ips"`
	Links   AccountUrlscannerScanGetResponseResultScanLinks   `json:"links"`
	JSON    accountUrlscannerScanGetResponseResultScanJSON    `json:"-"`
}

// accountUrlscannerScanGetResponseResultScanJSON contains the JSON metadata for
// the struct [AccountUrlscannerScanGetResponseResultScan]
type accountUrlscannerScanGetResponseResultScanJSON struct {
	Certificates apijson.Field
	Geo          apijson.Field
	Meta         apijson.Field
	Page         apijson.Field
	Performance  apijson.Field
	Task         apijson.Field
	Verdicts     apijson.Field
	Asns         apijson.Field
	Domains      apijson.Field
	IPs          apijson.Field
	Links        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountUrlscannerScanGetResponseResultScan) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerScanGetResponseResultScanJSON) RawJSON() string {
	return r.raw
}

type AccountUrlscannerScanGetResponseResultScanCertificate struct {
	Issuer      string                                                    `json:"issuer" api:"required"`
	SubjectName string                                                    `json:"subjectName" api:"required"`
	ValidFrom   float64                                                   `json:"validFrom" api:"required"`
	ValidTo     float64                                                   `json:"validTo" api:"required"`
	JSON        accountUrlscannerScanGetResponseResultScanCertificateJSON `json:"-"`
}

// accountUrlscannerScanGetResponseResultScanCertificateJSON contains the JSON
// metadata for the struct [AccountUrlscannerScanGetResponseResultScanCertificate]
type accountUrlscannerScanGetResponseResultScanCertificateJSON struct {
	Issuer      apijson.Field
	SubjectName apijson.Field
	ValidFrom   apijson.Field
	ValidTo     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUrlscannerScanGetResponseResultScanCertificate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerScanGetResponseResultScanCertificateJSON) RawJSON() string {
	return r.raw
}

type AccountUrlscannerScanGetResponseResultScanGeo struct {
	Continents []string                                          `json:"continents" api:"required"`
	Locations  []string                                          `json:"locations" api:"required"`
	JSON       accountUrlscannerScanGetResponseResultScanGeoJSON `json:"-"`
}

// accountUrlscannerScanGetResponseResultScanGeoJSON contains the JSON metadata for
// the struct [AccountUrlscannerScanGetResponseResultScanGeo]
type accountUrlscannerScanGetResponseResultScanGeoJSON struct {
	Continents  apijson.Field
	Locations   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUrlscannerScanGetResponseResultScanGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerScanGetResponseResultScanGeoJSON) RawJSON() string {
	return r.raw
}

type AccountUrlscannerScanGetResponseResultScanMeta struct {
	Processors AccountUrlscannerScanGetResponseResultScanMetaProcessors `json:"processors" api:"required"`
	JSON       accountUrlscannerScanGetResponseResultScanMetaJSON       `json:"-"`
}

// accountUrlscannerScanGetResponseResultScanMetaJSON contains the JSON metadata
// for the struct [AccountUrlscannerScanGetResponseResultScanMeta]
type accountUrlscannerScanGetResponseResultScanMetaJSON struct {
	Processors  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUrlscannerScanGetResponseResultScanMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerScanGetResponseResultScanMetaJSON) RawJSON() string {
	return r.raw
}

type AccountUrlscannerScanGetResponseResultScanMetaProcessors struct {
	Categories AccountUrlscannerScanGetResponseResultScanMetaProcessorsCategories `json:"categories" api:"required"`
	Phishing   []string                                                           `json:"phishing" api:"required"`
	Rank       AccountUrlscannerScanGetResponseResultScanMetaProcessorsRank       `json:"rank" api:"required"`
	Tech       []AccountUrlscannerScanGetResponseResultScanMetaProcessorsTech     `json:"tech" api:"required"`
	JSON       accountUrlscannerScanGetResponseResultScanMetaProcessorsJSON       `json:"-"`
}

// accountUrlscannerScanGetResponseResultScanMetaProcessorsJSON contains the JSON
// metadata for the struct
// [AccountUrlscannerScanGetResponseResultScanMetaProcessors]
type accountUrlscannerScanGetResponseResultScanMetaProcessorsJSON struct {
	Categories  apijson.Field
	Phishing    apijson.Field
	Rank        apijson.Field
	Tech        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUrlscannerScanGetResponseResultScanMetaProcessors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerScanGetResponseResultScanMetaProcessorsJSON) RawJSON() string {
	return r.raw
}

type AccountUrlscannerScanGetResponseResultScanMetaProcessorsCategories struct {
	Content []AccountUrlscannerScanGetResponseResultScanMetaProcessorsCategoriesContent `json:"content" api:"required"`
	Risks   []AccountUrlscannerScanGetResponseResultScanMetaProcessorsCategoriesRisk    `json:"risks" api:"required"`
	JSON    accountUrlscannerScanGetResponseResultScanMetaProcessorsCategoriesJSON      `json:"-"`
}

// accountUrlscannerScanGetResponseResultScanMetaProcessorsCategoriesJSON contains
// the JSON metadata for the struct
// [AccountUrlscannerScanGetResponseResultScanMetaProcessorsCategories]
type accountUrlscannerScanGetResponseResultScanMetaProcessorsCategoriesJSON struct {
	Content     apijson.Field
	Risks       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUrlscannerScanGetResponseResultScanMetaProcessorsCategories) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerScanGetResponseResultScanMetaProcessorsCategoriesJSON) RawJSON() string {
	return r.raw
}

type AccountUrlscannerScanGetResponseResultScanMetaProcessorsCategoriesContent struct {
	ID              int64                                                                         `json:"id" api:"required"`
	Name            string                                                                        `json:"name" api:"required"`
	SuperCategoryID int64                                                                         `json:"super_category_id"`
	JSON            accountUrlscannerScanGetResponseResultScanMetaProcessorsCategoriesContentJSON `json:"-"`
}

// accountUrlscannerScanGetResponseResultScanMetaProcessorsCategoriesContentJSON
// contains the JSON metadata for the struct
// [AccountUrlscannerScanGetResponseResultScanMetaProcessorsCategoriesContent]
type accountUrlscannerScanGetResponseResultScanMetaProcessorsCategoriesContentJSON struct {
	ID              apijson.Field
	Name            apijson.Field
	SuperCategoryID apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccountUrlscannerScanGetResponseResultScanMetaProcessorsCategoriesContent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerScanGetResponseResultScanMetaProcessorsCategoriesContentJSON) RawJSON() string {
	return r.raw
}

type AccountUrlscannerScanGetResponseResultScanMetaProcessorsCategoriesRisk struct {
	ID              int64                                                                      `json:"id" api:"required"`
	Name            string                                                                     `json:"name" api:"required"`
	SuperCategoryID int64                                                                      `json:"super_category_id" api:"required"`
	JSON            accountUrlscannerScanGetResponseResultScanMetaProcessorsCategoriesRiskJSON `json:"-"`
}

// accountUrlscannerScanGetResponseResultScanMetaProcessorsCategoriesRiskJSON
// contains the JSON metadata for the struct
// [AccountUrlscannerScanGetResponseResultScanMetaProcessorsCategoriesRisk]
type accountUrlscannerScanGetResponseResultScanMetaProcessorsCategoriesRiskJSON struct {
	ID              apijson.Field
	Name            apijson.Field
	SuperCategoryID apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccountUrlscannerScanGetResponseResultScanMetaProcessorsCategoriesRisk) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerScanGetResponseResultScanMetaProcessorsCategoriesRiskJSON) RawJSON() string {
	return r.raw
}

type AccountUrlscannerScanGetResponseResultScanMetaProcessorsRank struct {
	Bucket string `json:"bucket" api:"required"`
	Name   string `json:"name" api:"required"`
	// Rank in the Global Radar Rank, if set. See more at
	// https://blog.cloudflare.com/radar-domain-rankings/
	Rank int64                                                            `json:"rank"`
	JSON accountUrlscannerScanGetResponseResultScanMetaProcessorsRankJSON `json:"-"`
}

// accountUrlscannerScanGetResponseResultScanMetaProcessorsRankJSON contains the
// JSON metadata for the struct
// [AccountUrlscannerScanGetResponseResultScanMetaProcessorsRank]
type accountUrlscannerScanGetResponseResultScanMetaProcessorsRankJSON struct {
	Bucket      apijson.Field
	Name        apijson.Field
	Rank        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUrlscannerScanGetResponseResultScanMetaProcessorsRank) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerScanGetResponseResultScanMetaProcessorsRankJSON) RawJSON() string {
	return r.raw
}

type AccountUrlscannerScanGetResponseResultScanMetaProcessorsTech struct {
	Categories  []AccountUrlscannerScanGetResponseResultScanMetaProcessorsTechCategory `json:"categories" api:"required"`
	Confidence  int64                                                                  `json:"confidence" api:"required"`
	Evidence    AccountUrlscannerScanGetResponseResultScanMetaProcessorsTechEvidence   `json:"evidence" api:"required"`
	Icon        string                                                                 `json:"icon" api:"required"`
	Name        string                                                                 `json:"name" api:"required"`
	Slug        string                                                                 `json:"slug" api:"required"`
	Website     string                                                                 `json:"website" api:"required"`
	Description string                                                                 `json:"description"`
	JSON        accountUrlscannerScanGetResponseResultScanMetaProcessorsTechJSON       `json:"-"`
}

// accountUrlscannerScanGetResponseResultScanMetaProcessorsTechJSON contains the
// JSON metadata for the struct
// [AccountUrlscannerScanGetResponseResultScanMetaProcessorsTech]
type accountUrlscannerScanGetResponseResultScanMetaProcessorsTechJSON struct {
	Categories  apijson.Field
	Confidence  apijson.Field
	Evidence    apijson.Field
	Icon        apijson.Field
	Name        apijson.Field
	Slug        apijson.Field
	Website     apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUrlscannerScanGetResponseResultScanMetaProcessorsTech) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerScanGetResponseResultScanMetaProcessorsTechJSON) RawJSON() string {
	return r.raw
}

type AccountUrlscannerScanGetResponseResultScanMetaProcessorsTechCategory struct {
	ID       int64                                                                    `json:"id" api:"required"`
	Groups   []int64                                                                  `json:"groups" api:"required"`
	Name     string                                                                   `json:"name" api:"required"`
	Priority int64                                                                    `json:"priority" api:"required"`
	Slug     string                                                                   `json:"slug" api:"required"`
	JSON     accountUrlscannerScanGetResponseResultScanMetaProcessorsTechCategoryJSON `json:"-"`
}

// accountUrlscannerScanGetResponseResultScanMetaProcessorsTechCategoryJSON
// contains the JSON metadata for the struct
// [AccountUrlscannerScanGetResponseResultScanMetaProcessorsTechCategory]
type accountUrlscannerScanGetResponseResultScanMetaProcessorsTechCategoryJSON struct {
	ID          apijson.Field
	Groups      apijson.Field
	Name        apijson.Field
	Priority    apijson.Field
	Slug        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUrlscannerScanGetResponseResultScanMetaProcessorsTechCategory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerScanGetResponseResultScanMetaProcessorsTechCategoryJSON) RawJSON() string {
	return r.raw
}

type AccountUrlscannerScanGetResponseResultScanMetaProcessorsTechEvidence struct {
	ImpliedBy []string                                                                      `json:"impliedBy" api:"required"`
	Patterns  []AccountUrlscannerScanGetResponseResultScanMetaProcessorsTechEvidencePattern `json:"patterns" api:"required"`
	JSON      accountUrlscannerScanGetResponseResultScanMetaProcessorsTechEvidenceJSON      `json:"-"`
}

// accountUrlscannerScanGetResponseResultScanMetaProcessorsTechEvidenceJSON
// contains the JSON metadata for the struct
// [AccountUrlscannerScanGetResponseResultScanMetaProcessorsTechEvidence]
type accountUrlscannerScanGetResponseResultScanMetaProcessorsTechEvidenceJSON struct {
	ImpliedBy   apijson.Field
	Patterns    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUrlscannerScanGetResponseResultScanMetaProcessorsTechEvidence) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerScanGetResponseResultScanMetaProcessorsTechEvidenceJSON) RawJSON() string {
	return r.raw
}

type AccountUrlscannerScanGetResponseResultScanMetaProcessorsTechEvidencePattern struct {
	Confidence int64    `json:"confidence" api:"required"`
	Excludes   []string `json:"excludes" api:"required"`
	Implies    []string `json:"implies" api:"required"`
	Match      string   `json:"match" api:"required"`
	// Header or Cookie name when set
	Name    string                                                                          `json:"name" api:"required"`
	Regex   string                                                                          `json:"regex" api:"required"`
	Type    string                                                                          `json:"type" api:"required"`
	Value   string                                                                          `json:"value" api:"required"`
	Version string                                                                          `json:"version" api:"required"`
	JSON    accountUrlscannerScanGetResponseResultScanMetaProcessorsTechEvidencePatternJSON `json:"-"`
}

// accountUrlscannerScanGetResponseResultScanMetaProcessorsTechEvidencePatternJSON
// contains the JSON metadata for the struct
// [AccountUrlscannerScanGetResponseResultScanMetaProcessorsTechEvidencePattern]
type accountUrlscannerScanGetResponseResultScanMetaProcessorsTechEvidencePatternJSON struct {
	Confidence  apijson.Field
	Excludes    apijson.Field
	Implies     apijson.Field
	Match       apijson.Field
	Name        apijson.Field
	Regex       apijson.Field
	Type        apijson.Field
	Value       apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUrlscannerScanGetResponseResultScanMetaProcessorsTechEvidencePattern) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerScanGetResponseResultScanMetaProcessorsTechEvidencePatternJSON) RawJSON() string {
	return r.raw
}

type AccountUrlscannerScanGetResponseResultScanPage struct {
	Asn                   string                                                            `json:"asn" api:"required"`
	AsnLocationAlpha2     string                                                            `json:"asnLocationAlpha2" api:"required"`
	Asnname               string                                                            `json:"asnname" api:"required"`
	Console               []AccountUrlscannerScanGetResponseResultScanPageConsole           `json:"console" api:"required"`
	Cookies               []AccountUrlscannerScanGetResponseResultScanPageCookie            `json:"cookies" api:"required"`
	Country               string                                                            `json:"country" api:"required"`
	CountryLocationAlpha2 string                                                            `json:"countryLocationAlpha2" api:"required"`
	Domain                string                                                            `json:"domain" api:"required"`
	Headers               []AccountUrlscannerScanGetResponseResultScanPageHeader            `json:"headers" api:"required"`
	IP                    string                                                            `json:"ip" api:"required"`
	Js                    AccountUrlscannerScanGetResponseResultScanPageJs                  `json:"js" api:"required"`
	SecurityViolations    []AccountUrlscannerScanGetResponseResultScanPageSecurityViolation `json:"securityViolations" api:"required"`
	Status                float64                                                           `json:"status" api:"required"`
	Subdivision1Name      string                                                            `json:"subdivision1Name" api:"required"`
	Subdivision2name      string                                                            `json:"subdivision2name" api:"required"`
	URL                   string                                                            `json:"url" api:"required"`
	JSON                  accountUrlscannerScanGetResponseResultScanPageJSON                `json:"-"`
}

// accountUrlscannerScanGetResponseResultScanPageJSON contains the JSON metadata
// for the struct [AccountUrlscannerScanGetResponseResultScanPage]
type accountUrlscannerScanGetResponseResultScanPageJSON struct {
	Asn                   apijson.Field
	AsnLocationAlpha2     apijson.Field
	Asnname               apijson.Field
	Console               apijson.Field
	Cookies               apijson.Field
	Country               apijson.Field
	CountryLocationAlpha2 apijson.Field
	Domain                apijson.Field
	Headers               apijson.Field
	IP                    apijson.Field
	Js                    apijson.Field
	SecurityViolations    apijson.Field
	Status                apijson.Field
	Subdivision1Name      apijson.Field
	Subdivision2name      apijson.Field
	URL                   apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *AccountUrlscannerScanGetResponseResultScanPage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerScanGetResponseResultScanPageJSON) RawJSON() string {
	return r.raw
}

type AccountUrlscannerScanGetResponseResultScanPageConsole struct {
	Category string                                                    `json:"category" api:"required"`
	Text     string                                                    `json:"text" api:"required"`
	Type     string                                                    `json:"type" api:"required"`
	URL      string                                                    `json:"url"`
	JSON     accountUrlscannerScanGetResponseResultScanPageConsoleJSON `json:"-"`
}

// accountUrlscannerScanGetResponseResultScanPageConsoleJSON contains the JSON
// metadata for the struct [AccountUrlscannerScanGetResponseResultScanPageConsole]
type accountUrlscannerScanGetResponseResultScanPageConsoleJSON struct {
	Category    apijson.Field
	Text        apijson.Field
	Type        apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUrlscannerScanGetResponseResultScanPageConsole) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerScanGetResponseResultScanPageConsoleJSON) RawJSON() string {
	return r.raw
}

type AccountUrlscannerScanGetResponseResultScanPageCookie struct {
	Domain       string                                                   `json:"domain" api:"required"`
	Expires      float64                                                  `json:"expires" api:"required"`
	HTTPOnly     bool                                                     `json:"httpOnly" api:"required"`
	Name         string                                                   `json:"name" api:"required"`
	Path         string                                                   `json:"path" api:"required"`
	SameParty    bool                                                     `json:"sameParty" api:"required"`
	Secure       bool                                                     `json:"secure" api:"required"`
	Session      bool                                                     `json:"session" api:"required"`
	Size         float64                                                  `json:"size" api:"required"`
	SourcePort   float64                                                  `json:"sourcePort" api:"required"`
	SourceScheme string                                                   `json:"sourceScheme" api:"required"`
	Value        string                                                   `json:"value" api:"required"`
	Priority     string                                                   `json:"priority"`
	JSON         accountUrlscannerScanGetResponseResultScanPageCookieJSON `json:"-"`
}

// accountUrlscannerScanGetResponseResultScanPageCookieJSON contains the JSON
// metadata for the struct [AccountUrlscannerScanGetResponseResultScanPageCookie]
type accountUrlscannerScanGetResponseResultScanPageCookieJSON struct {
	Domain       apijson.Field
	Expires      apijson.Field
	HTTPOnly     apijson.Field
	Name         apijson.Field
	Path         apijson.Field
	SameParty    apijson.Field
	Secure       apijson.Field
	Session      apijson.Field
	Size         apijson.Field
	SourcePort   apijson.Field
	SourceScheme apijson.Field
	Value        apijson.Field
	Priority     apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountUrlscannerScanGetResponseResultScanPageCookie) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerScanGetResponseResultScanPageCookieJSON) RawJSON() string {
	return r.raw
}

type AccountUrlscannerScanGetResponseResultScanPageHeader struct {
	Name  string                                                   `json:"name" api:"required"`
	Value string                                                   `json:"value" api:"required"`
	JSON  accountUrlscannerScanGetResponseResultScanPageHeaderJSON `json:"-"`
}

// accountUrlscannerScanGetResponseResultScanPageHeaderJSON contains the JSON
// metadata for the struct [AccountUrlscannerScanGetResponseResultScanPageHeader]
type accountUrlscannerScanGetResponseResultScanPageHeaderJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUrlscannerScanGetResponseResultScanPageHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerScanGetResponseResultScanPageHeaderJSON) RawJSON() string {
	return r.raw
}

type AccountUrlscannerScanGetResponseResultScanPageJs struct {
	Variables []AccountUrlscannerScanGetResponseResultScanPageJsVariable `json:"variables" api:"required"`
	JSON      accountUrlscannerScanGetResponseResultScanPageJsJSON       `json:"-"`
}

// accountUrlscannerScanGetResponseResultScanPageJsJSON contains the JSON metadata
// for the struct [AccountUrlscannerScanGetResponseResultScanPageJs]
type accountUrlscannerScanGetResponseResultScanPageJsJSON struct {
	Variables   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUrlscannerScanGetResponseResultScanPageJs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerScanGetResponseResultScanPageJsJSON) RawJSON() string {
	return r.raw
}

type AccountUrlscannerScanGetResponseResultScanPageJsVariable struct {
	Name string                                                       `json:"name" api:"required"`
	Type string                                                       `json:"type" api:"required"`
	JSON accountUrlscannerScanGetResponseResultScanPageJsVariableJSON `json:"-"`
}

// accountUrlscannerScanGetResponseResultScanPageJsVariableJSON contains the JSON
// metadata for the struct
// [AccountUrlscannerScanGetResponseResultScanPageJsVariable]
type accountUrlscannerScanGetResponseResultScanPageJsVariableJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUrlscannerScanGetResponseResultScanPageJsVariable) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerScanGetResponseResultScanPageJsVariableJSON) RawJSON() string {
	return r.raw
}

type AccountUrlscannerScanGetResponseResultScanPageSecurityViolation struct {
	Category string                                                              `json:"category" api:"required"`
	Text     string                                                              `json:"text" api:"required"`
	URL      string                                                              `json:"url" api:"required"`
	JSON     accountUrlscannerScanGetResponseResultScanPageSecurityViolationJSON `json:"-"`
}

// accountUrlscannerScanGetResponseResultScanPageSecurityViolationJSON contains the
// JSON metadata for the struct
// [AccountUrlscannerScanGetResponseResultScanPageSecurityViolation]
type accountUrlscannerScanGetResponseResultScanPageSecurityViolationJSON struct {
	Category    apijson.Field
	Text        apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUrlscannerScanGetResponseResultScanPageSecurityViolation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerScanGetResponseResultScanPageSecurityViolationJSON) RawJSON() string {
	return r.raw
}

type AccountUrlscannerScanGetResponseResultScanPerformance struct {
	ConnectEnd                 float64                                                   `json:"connectEnd" api:"required"`
	ConnectStart               float64                                                   `json:"connectStart" api:"required"`
	DecodedBodySize            float64                                                   `json:"decodedBodySize" api:"required"`
	DomainLookupEnd            float64                                                   `json:"domainLookupEnd" api:"required"`
	DomainLookupStart          float64                                                   `json:"domainLookupStart" api:"required"`
	DomComplete                float64                                                   `json:"domComplete" api:"required"`
	DomContentLoadedEventEnd   float64                                                   `json:"domContentLoadedEventEnd" api:"required"`
	DomContentLoadedEventStart float64                                                   `json:"domContentLoadedEventStart" api:"required"`
	DomInteractive             float64                                                   `json:"domInteractive" api:"required"`
	Duration                   float64                                                   `json:"duration" api:"required"`
	EncodedBodySize            float64                                                   `json:"encodedBodySize" api:"required"`
	EntryType                  string                                                    `json:"entryType" api:"required"`
	FetchStart                 float64                                                   `json:"fetchStart" api:"required"`
	InitiatorType              string                                                    `json:"initiatorType" api:"required"`
	LoadEventEnd               float64                                                   `json:"loadEventEnd" api:"required"`
	LoadEventStart             float64                                                   `json:"loadEventStart" api:"required"`
	Name                       string                                                    `json:"name" api:"required"`
	NextHopProtocol            string                                                    `json:"nextHopProtocol" api:"required"`
	RedirectCount              float64                                                   `json:"redirectCount" api:"required"`
	RedirectEnd                float64                                                   `json:"redirectEnd" api:"required"`
	RedirectStart              float64                                                   `json:"redirectStart" api:"required"`
	RequestStart               float64                                                   `json:"requestStart" api:"required"`
	ResponseEnd                float64                                                   `json:"responseEnd" api:"required"`
	ResponseStart              float64                                                   `json:"responseStart" api:"required"`
	SecureConnectionStart      float64                                                   `json:"secureConnectionStart" api:"required"`
	StartTime                  float64                                                   `json:"startTime" api:"required"`
	TransferSize               float64                                                   `json:"transferSize" api:"required"`
	Type                       string                                                    `json:"type" api:"required"`
	UnloadEventEnd             float64                                                   `json:"unloadEventEnd" api:"required"`
	UnloadEventStart           float64                                                   `json:"unloadEventStart" api:"required"`
	WorkerStart                float64                                                   `json:"workerStart" api:"required"`
	JSON                       accountUrlscannerScanGetResponseResultScanPerformanceJSON `json:"-"`
}

// accountUrlscannerScanGetResponseResultScanPerformanceJSON contains the JSON
// metadata for the struct [AccountUrlscannerScanGetResponseResultScanPerformance]
type accountUrlscannerScanGetResponseResultScanPerformanceJSON struct {
	ConnectEnd                 apijson.Field
	ConnectStart               apijson.Field
	DecodedBodySize            apijson.Field
	DomainLookupEnd            apijson.Field
	DomainLookupStart          apijson.Field
	DomComplete                apijson.Field
	DomContentLoadedEventEnd   apijson.Field
	DomContentLoadedEventStart apijson.Field
	DomInteractive             apijson.Field
	Duration                   apijson.Field
	EncodedBodySize            apijson.Field
	EntryType                  apijson.Field
	FetchStart                 apijson.Field
	InitiatorType              apijson.Field
	LoadEventEnd               apijson.Field
	LoadEventStart             apijson.Field
	Name                       apijson.Field
	NextHopProtocol            apijson.Field
	RedirectCount              apijson.Field
	RedirectEnd                apijson.Field
	RedirectStart              apijson.Field
	RequestStart               apijson.Field
	ResponseEnd                apijson.Field
	ResponseStart              apijson.Field
	SecureConnectionStart      apijson.Field
	StartTime                  apijson.Field
	TransferSize               apijson.Field
	Type                       apijson.Field
	UnloadEventEnd             apijson.Field
	UnloadEventStart           apijson.Field
	WorkerStart                apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *AccountUrlscannerScanGetResponseResultScanPerformance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerScanGetResponseResultScanPerformanceJSON) RawJSON() string {
	return r.raw
}

type AccountUrlscannerScanGetResponseResultScanTask struct {
	// Submitter location
	ClientLocation string                                                   `json:"clientLocation" api:"required"`
	ClientType     AccountUrlscannerScanGetResponseResultScanTaskClientType `json:"clientType" api:"required"`
	// URL of the primary request, after all HTTP redirects
	EffectiveURL string                                                    `json:"effectiveUrl" api:"required"`
	Errors       []AccountUrlscannerScanGetResponseResultScanTaskError     `json:"errors" api:"required"`
	ScannedFrom  AccountUrlscannerScanGetResponseResultScanTaskScannedFrom `json:"scannedFrom" api:"required"`
	Status       AccountUrlscannerScanGetResponseResultScanTaskStatus      `json:"status" api:"required"`
	Success      bool                                                      `json:"success" api:"required"`
	Time         string                                                    `json:"time" api:"required"`
	TimeEnd      string                                                    `json:"timeEnd" api:"required"`
	// Submitted URL
	URL string `json:"url" api:"required"`
	// Scan ID
	Uuid       string                                                   `json:"uuid" api:"required"`
	Visibility AccountUrlscannerScanGetResponseResultScanTaskVisibility `json:"visibility" api:"required"`
	JSON       accountUrlscannerScanGetResponseResultScanTaskJSON       `json:"-"`
}

// accountUrlscannerScanGetResponseResultScanTaskJSON contains the JSON metadata
// for the struct [AccountUrlscannerScanGetResponseResultScanTask]
type accountUrlscannerScanGetResponseResultScanTaskJSON struct {
	ClientLocation apijson.Field
	ClientType     apijson.Field
	EffectiveURL   apijson.Field
	Errors         apijson.Field
	ScannedFrom    apijson.Field
	Status         apijson.Field
	Success        apijson.Field
	Time           apijson.Field
	TimeEnd        apijson.Field
	URL            apijson.Field
	Uuid           apijson.Field
	Visibility     apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccountUrlscannerScanGetResponseResultScanTask) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerScanGetResponseResultScanTaskJSON) RawJSON() string {
	return r.raw
}

type AccountUrlscannerScanGetResponseResultScanTaskClientType string

const (
	AccountUrlscannerScanGetResponseResultScanTaskClientTypeSite      AccountUrlscannerScanGetResponseResultScanTaskClientType = "Site"
	AccountUrlscannerScanGetResponseResultScanTaskClientTypeAutomatic AccountUrlscannerScanGetResponseResultScanTaskClientType = "Automatic"
	AccountUrlscannerScanGetResponseResultScanTaskClientTypeAPI       AccountUrlscannerScanGetResponseResultScanTaskClientType = "Api"
)

func (r AccountUrlscannerScanGetResponseResultScanTaskClientType) IsKnown() bool {
	switch r {
	case AccountUrlscannerScanGetResponseResultScanTaskClientTypeSite, AccountUrlscannerScanGetResponseResultScanTaskClientTypeAutomatic, AccountUrlscannerScanGetResponseResultScanTaskClientTypeAPI:
		return true
	}
	return false
}

type AccountUrlscannerScanGetResponseResultScanTaskError struct {
	Message string                                                  `json:"message" api:"required"`
	JSON    accountUrlscannerScanGetResponseResultScanTaskErrorJSON `json:"-"`
}

// accountUrlscannerScanGetResponseResultScanTaskErrorJSON contains the JSON
// metadata for the struct [AccountUrlscannerScanGetResponseResultScanTaskError]
type accountUrlscannerScanGetResponseResultScanTaskErrorJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUrlscannerScanGetResponseResultScanTaskError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerScanGetResponseResultScanTaskErrorJSON) RawJSON() string {
	return r.raw
}

type AccountUrlscannerScanGetResponseResultScanTaskScannedFrom struct {
	// IATA code of Cloudflare datacenter
	Colo string                                                        `json:"colo" api:"required"`
	JSON accountUrlscannerScanGetResponseResultScanTaskScannedFromJSON `json:"-"`
}

// accountUrlscannerScanGetResponseResultScanTaskScannedFromJSON contains the JSON
// metadata for the struct
// [AccountUrlscannerScanGetResponseResultScanTaskScannedFrom]
type accountUrlscannerScanGetResponseResultScanTaskScannedFromJSON struct {
	Colo        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUrlscannerScanGetResponseResultScanTaskScannedFrom) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerScanGetResponseResultScanTaskScannedFromJSON) RawJSON() string {
	return r.raw
}

type AccountUrlscannerScanGetResponseResultScanTaskStatus string

const (
	AccountUrlscannerScanGetResponseResultScanTaskStatusQueued           AccountUrlscannerScanGetResponseResultScanTaskStatus = "Queued"
	AccountUrlscannerScanGetResponseResultScanTaskStatusInProgress       AccountUrlscannerScanGetResponseResultScanTaskStatus = "InProgress"
	AccountUrlscannerScanGetResponseResultScanTaskStatusInPostProcessing AccountUrlscannerScanGetResponseResultScanTaskStatus = "InPostProcessing"
	AccountUrlscannerScanGetResponseResultScanTaskStatusFinished         AccountUrlscannerScanGetResponseResultScanTaskStatus = "Finished"
)

func (r AccountUrlscannerScanGetResponseResultScanTaskStatus) IsKnown() bool {
	switch r {
	case AccountUrlscannerScanGetResponseResultScanTaskStatusQueued, AccountUrlscannerScanGetResponseResultScanTaskStatusInProgress, AccountUrlscannerScanGetResponseResultScanTaskStatusInPostProcessing, AccountUrlscannerScanGetResponseResultScanTaskStatusFinished:
		return true
	}
	return false
}

type AccountUrlscannerScanGetResponseResultScanTaskVisibility string

const (
	AccountUrlscannerScanGetResponseResultScanTaskVisibilityPublic   AccountUrlscannerScanGetResponseResultScanTaskVisibility = "Public"
	AccountUrlscannerScanGetResponseResultScanTaskVisibilityUnlisted AccountUrlscannerScanGetResponseResultScanTaskVisibility = "Unlisted"
)

func (r AccountUrlscannerScanGetResponseResultScanTaskVisibility) IsKnown() bool {
	switch r {
	case AccountUrlscannerScanGetResponseResultScanTaskVisibilityPublic, AccountUrlscannerScanGetResponseResultScanTaskVisibilityUnlisted:
		return true
	}
	return false
}

type AccountUrlscannerScanGetResponseResultScanVerdicts struct {
	Overall AccountUrlscannerScanGetResponseResultScanVerdictsOverall `json:"overall" api:"required"`
	JSON    accountUrlscannerScanGetResponseResultScanVerdictsJSON    `json:"-"`
}

// accountUrlscannerScanGetResponseResultScanVerdictsJSON contains the JSON
// metadata for the struct [AccountUrlscannerScanGetResponseResultScanVerdicts]
type accountUrlscannerScanGetResponseResultScanVerdictsJSON struct {
	Overall     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUrlscannerScanGetResponseResultScanVerdicts) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerScanGetResponseResultScanVerdictsJSON) RawJSON() string {
	return r.raw
}

type AccountUrlscannerScanGetResponseResultScanVerdictsOverall struct {
	Categories []AccountUrlscannerScanGetResponseResultScanVerdictsOverallCategory `json:"categories" api:"required"`
	// At least one of our subsystems marked the site as potentially malicious at the
	// time of the scan.
	Malicious bool                                                          `json:"malicious" api:"required"`
	Phishing  []string                                                      `json:"phishing" api:"required"`
	JSON      accountUrlscannerScanGetResponseResultScanVerdictsOverallJSON `json:"-"`
}

// accountUrlscannerScanGetResponseResultScanVerdictsOverallJSON contains the JSON
// metadata for the struct
// [AccountUrlscannerScanGetResponseResultScanVerdictsOverall]
type accountUrlscannerScanGetResponseResultScanVerdictsOverallJSON struct {
	Categories  apijson.Field
	Malicious   apijson.Field
	Phishing    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUrlscannerScanGetResponseResultScanVerdictsOverall) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerScanGetResponseResultScanVerdictsOverallJSON) RawJSON() string {
	return r.raw
}

type AccountUrlscannerScanGetResponseResultScanVerdictsOverallCategory struct {
	ID              float64                                                               `json:"id" api:"required"`
	Name            string                                                                `json:"name" api:"required"`
	SuperCategoryID float64                                                               `json:"super_category_id" api:"required"`
	JSON            accountUrlscannerScanGetResponseResultScanVerdictsOverallCategoryJSON `json:"-"`
}

// accountUrlscannerScanGetResponseResultScanVerdictsOverallCategoryJSON contains
// the JSON metadata for the struct
// [AccountUrlscannerScanGetResponseResultScanVerdictsOverallCategory]
type accountUrlscannerScanGetResponseResultScanVerdictsOverallCategoryJSON struct {
	ID              apijson.Field
	Name            apijson.Field
	SuperCategoryID apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccountUrlscannerScanGetResponseResultScanVerdictsOverallCategory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerScanGetResponseResultScanVerdictsOverallCategoryJSON) RawJSON() string {
	return r.raw
}

// Dictionary of Autonomous System Numbers where ASN's are the keys
type AccountUrlscannerScanGetResponseResultScanAsns struct {
	// ASN's contacted
	Asn  AccountUrlscannerScanGetResponseResultScanAsnsAsn  `json:"asn"`
	JSON accountUrlscannerScanGetResponseResultScanAsnsJSON `json:"-"`
}

// accountUrlscannerScanGetResponseResultScanAsnsJSON contains the JSON metadata
// for the struct [AccountUrlscannerScanGetResponseResultScanAsns]
type accountUrlscannerScanGetResponseResultScanAsnsJSON struct {
	Asn         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUrlscannerScanGetResponseResultScanAsns) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerScanGetResponseResultScanAsnsJSON) RawJSON() string {
	return r.raw
}

// ASN's contacted
type AccountUrlscannerScanGetResponseResultScanAsnsAsn struct {
	Asn            string                                                `json:"asn" api:"required"`
	Description    string                                                `json:"description" api:"required"`
	LocationAlpha2 string                                                `json:"location_alpha2" api:"required"`
	Name           string                                                `json:"name" api:"required"`
	OrgName        string                                                `json:"org_name" api:"required"`
	JSON           accountUrlscannerScanGetResponseResultScanAsnsAsnJSON `json:"-"`
}

// accountUrlscannerScanGetResponseResultScanAsnsAsnJSON contains the JSON metadata
// for the struct [AccountUrlscannerScanGetResponseResultScanAsnsAsn]
type accountUrlscannerScanGetResponseResultScanAsnsAsnJSON struct {
	Asn            apijson.Field
	Description    apijson.Field
	LocationAlpha2 apijson.Field
	Name           apijson.Field
	OrgName        apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccountUrlscannerScanGetResponseResultScanAsnsAsn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerScanGetResponseResultScanAsnsAsnJSON) RawJSON() string {
	return r.raw
}

type AccountUrlscannerScanGetResponseResultScanDomains struct {
	ExampleCom AccountUrlscannerScanGetResponseResultScanDomainsExampleCom `json:"example.com"`
	JSON       accountUrlscannerScanGetResponseResultScanDomainsJSON       `json:"-"`
}

// accountUrlscannerScanGetResponseResultScanDomainsJSON contains the JSON metadata
// for the struct [AccountUrlscannerScanGetResponseResultScanDomains]
type accountUrlscannerScanGetResponseResultScanDomainsJSON struct {
	ExampleCom  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUrlscannerScanGetResponseResultScanDomains) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerScanGetResponseResultScanDomainsJSON) RawJSON() string {
	return r.raw
}

type AccountUrlscannerScanGetResponseResultScanDomainsExampleCom struct {
	Categories AccountUrlscannerScanGetResponseResultScanDomainsExampleComCategories `json:"categories" api:"required"`
	DNS        []AccountUrlscannerScanGetResponseResultScanDomainsExampleComDNS      `json:"dns" api:"required"`
	Name       string                                                                `json:"name" api:"required"`
	Rank       AccountUrlscannerScanGetResponseResultScanDomainsExampleComRank       `json:"rank" api:"required"`
	Type       string                                                                `json:"type" api:"required"`
	JSON       accountUrlscannerScanGetResponseResultScanDomainsExampleComJSON       `json:"-"`
}

// accountUrlscannerScanGetResponseResultScanDomainsExampleComJSON contains the
// JSON metadata for the struct
// [AccountUrlscannerScanGetResponseResultScanDomainsExampleCom]
type accountUrlscannerScanGetResponseResultScanDomainsExampleComJSON struct {
	Categories  apijson.Field
	DNS         apijson.Field
	Name        apijson.Field
	Rank        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUrlscannerScanGetResponseResultScanDomainsExampleCom) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerScanGetResponseResultScanDomainsExampleComJSON) RawJSON() string {
	return r.raw
}

type AccountUrlscannerScanGetResponseResultScanDomainsExampleComCategories struct {
	Inherited AccountUrlscannerScanGetResponseResultScanDomainsExampleComCategoriesInherited `json:"inherited" api:"required"`
	Content   []AccountUrlscannerScanGetResponseResultScanDomainsExampleComCategoriesContent `json:"content"`
	Risks     []AccountUrlscannerScanGetResponseResultScanDomainsExampleComCategoriesRisk    `json:"risks"`
	JSON      accountUrlscannerScanGetResponseResultScanDomainsExampleComCategoriesJSON      `json:"-"`
}

// accountUrlscannerScanGetResponseResultScanDomainsExampleComCategoriesJSON
// contains the JSON metadata for the struct
// [AccountUrlscannerScanGetResponseResultScanDomainsExampleComCategories]
type accountUrlscannerScanGetResponseResultScanDomainsExampleComCategoriesJSON struct {
	Inherited   apijson.Field
	Content     apijson.Field
	Risks       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUrlscannerScanGetResponseResultScanDomainsExampleComCategories) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerScanGetResponseResultScanDomainsExampleComCategoriesJSON) RawJSON() string {
	return r.raw
}

type AccountUrlscannerScanGetResponseResultScanDomainsExampleComCategoriesInherited struct {
	Content []AccountUrlscannerScanGetResponseResultScanDomainsExampleComCategoriesInheritedContent `json:"content"`
	From    string                                                                                  `json:"from"`
	Risks   []AccountUrlscannerScanGetResponseResultScanDomainsExampleComCategoriesInheritedRisk    `json:"risks"`
	JSON    accountUrlscannerScanGetResponseResultScanDomainsExampleComCategoriesInheritedJSON      `json:"-"`
}

// accountUrlscannerScanGetResponseResultScanDomainsExampleComCategoriesInheritedJSON
// contains the JSON metadata for the struct
// [AccountUrlscannerScanGetResponseResultScanDomainsExampleComCategoriesInherited]
type accountUrlscannerScanGetResponseResultScanDomainsExampleComCategoriesInheritedJSON struct {
	Content     apijson.Field
	From        apijson.Field
	Risks       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUrlscannerScanGetResponseResultScanDomainsExampleComCategoriesInherited) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerScanGetResponseResultScanDomainsExampleComCategoriesInheritedJSON) RawJSON() string {
	return r.raw
}

type AccountUrlscannerScanGetResponseResultScanDomainsExampleComCategoriesInheritedContent struct {
	ID              int64                                                                                     `json:"id" api:"required"`
	Name            string                                                                                    `json:"name" api:"required"`
	SuperCategoryID int64                                                                                     `json:"super_category_id"`
	JSON            accountUrlscannerScanGetResponseResultScanDomainsExampleComCategoriesInheritedContentJSON `json:"-"`
}

// accountUrlscannerScanGetResponseResultScanDomainsExampleComCategoriesInheritedContentJSON
// contains the JSON metadata for the struct
// [AccountUrlscannerScanGetResponseResultScanDomainsExampleComCategoriesInheritedContent]
type accountUrlscannerScanGetResponseResultScanDomainsExampleComCategoriesInheritedContentJSON struct {
	ID              apijson.Field
	Name            apijson.Field
	SuperCategoryID apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccountUrlscannerScanGetResponseResultScanDomainsExampleComCategoriesInheritedContent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerScanGetResponseResultScanDomainsExampleComCategoriesInheritedContentJSON) RawJSON() string {
	return r.raw
}

type AccountUrlscannerScanGetResponseResultScanDomainsExampleComCategoriesInheritedRisk struct {
	ID              int64                                                                                  `json:"id" api:"required"`
	Name            string                                                                                 `json:"name" api:"required"`
	SuperCategoryID int64                                                                                  `json:"super_category_id"`
	JSON            accountUrlscannerScanGetResponseResultScanDomainsExampleComCategoriesInheritedRiskJSON `json:"-"`
}

// accountUrlscannerScanGetResponseResultScanDomainsExampleComCategoriesInheritedRiskJSON
// contains the JSON metadata for the struct
// [AccountUrlscannerScanGetResponseResultScanDomainsExampleComCategoriesInheritedRisk]
type accountUrlscannerScanGetResponseResultScanDomainsExampleComCategoriesInheritedRiskJSON struct {
	ID              apijson.Field
	Name            apijson.Field
	SuperCategoryID apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccountUrlscannerScanGetResponseResultScanDomainsExampleComCategoriesInheritedRisk) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerScanGetResponseResultScanDomainsExampleComCategoriesInheritedRiskJSON) RawJSON() string {
	return r.raw
}

type AccountUrlscannerScanGetResponseResultScanDomainsExampleComCategoriesContent struct {
	ID              int64                                                                            `json:"id" api:"required"`
	Name            string                                                                           `json:"name" api:"required"`
	SuperCategoryID int64                                                                            `json:"super_category_id"`
	JSON            accountUrlscannerScanGetResponseResultScanDomainsExampleComCategoriesContentJSON `json:"-"`
}

// accountUrlscannerScanGetResponseResultScanDomainsExampleComCategoriesContentJSON
// contains the JSON metadata for the struct
// [AccountUrlscannerScanGetResponseResultScanDomainsExampleComCategoriesContent]
type accountUrlscannerScanGetResponseResultScanDomainsExampleComCategoriesContentJSON struct {
	ID              apijson.Field
	Name            apijson.Field
	SuperCategoryID apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccountUrlscannerScanGetResponseResultScanDomainsExampleComCategoriesContent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerScanGetResponseResultScanDomainsExampleComCategoriesContentJSON) RawJSON() string {
	return r.raw
}

type AccountUrlscannerScanGetResponseResultScanDomainsExampleComCategoriesRisk struct {
	ID              int64                                                                         `json:"id" api:"required"`
	Name            string                                                                        `json:"name" api:"required"`
	SuperCategoryID int64                                                                         `json:"super_category_id"`
	JSON            accountUrlscannerScanGetResponseResultScanDomainsExampleComCategoriesRiskJSON `json:"-"`
}

// accountUrlscannerScanGetResponseResultScanDomainsExampleComCategoriesRiskJSON
// contains the JSON metadata for the struct
// [AccountUrlscannerScanGetResponseResultScanDomainsExampleComCategoriesRisk]
type accountUrlscannerScanGetResponseResultScanDomainsExampleComCategoriesRiskJSON struct {
	ID              apijson.Field
	Name            apijson.Field
	SuperCategoryID apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccountUrlscannerScanGetResponseResultScanDomainsExampleComCategoriesRisk) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerScanGetResponseResultScanDomainsExampleComCategoriesRiskJSON) RawJSON() string {
	return r.raw
}

type AccountUrlscannerScanGetResponseResultScanDomainsExampleComDNS struct {
	Address     string                                                             `json:"address" api:"required"`
	DnssecValid bool                                                               `json:"dnssec_valid" api:"required"`
	Name        string                                                             `json:"name" api:"required"`
	Type        string                                                             `json:"type" api:"required"`
	JSON        accountUrlscannerScanGetResponseResultScanDomainsExampleComDNSJSON `json:"-"`
}

// accountUrlscannerScanGetResponseResultScanDomainsExampleComDNSJSON contains the
// JSON metadata for the struct
// [AccountUrlscannerScanGetResponseResultScanDomainsExampleComDNS]
type accountUrlscannerScanGetResponseResultScanDomainsExampleComDNSJSON struct {
	Address     apijson.Field
	DnssecValid apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUrlscannerScanGetResponseResultScanDomainsExampleComDNS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerScanGetResponseResultScanDomainsExampleComDNSJSON) RawJSON() string {
	return r.raw
}

type AccountUrlscannerScanGetResponseResultScanDomainsExampleComRank struct {
	Bucket string `json:"bucket" api:"required"`
	Name   string `json:"name" api:"required"`
	// Rank in the Global Radar Rank, if set. See more at
	// https://blog.cloudflare.com/radar-domain-rankings/
	Rank int64                                                               `json:"rank"`
	JSON accountUrlscannerScanGetResponseResultScanDomainsExampleComRankJSON `json:"-"`
}

// accountUrlscannerScanGetResponseResultScanDomainsExampleComRankJSON contains the
// JSON metadata for the struct
// [AccountUrlscannerScanGetResponseResultScanDomainsExampleComRank]
type accountUrlscannerScanGetResponseResultScanDomainsExampleComRankJSON struct {
	Bucket      apijson.Field
	Name        apijson.Field
	Rank        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUrlscannerScanGetResponseResultScanDomainsExampleComRank) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerScanGetResponseResultScanDomainsExampleComRankJSON) RawJSON() string {
	return r.raw
}

type AccountUrlscannerScanGetResponseResultScanIPs struct {
	IP   AccountUrlscannerScanGetResponseResultScanIPsIP   `json:"ip"`
	JSON accountUrlscannerScanGetResponseResultScanIPsJSON `json:"-"`
}

// accountUrlscannerScanGetResponseResultScanIPsJSON contains the JSON metadata for
// the struct [AccountUrlscannerScanGetResponseResultScanIPs]
type accountUrlscannerScanGetResponseResultScanIPsJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUrlscannerScanGetResponseResultScanIPs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerScanGetResponseResultScanIPsJSON) RawJSON() string {
	return r.raw
}

type AccountUrlscannerScanGetResponseResultScanIPsIP struct {
	Asn               string                                              `json:"asn" api:"required"`
	AsnDescription    string                                              `json:"asnDescription" api:"required"`
	AsnLocationAlpha2 string                                              `json:"asnLocationAlpha2" api:"required"`
	AsnName           string                                              `json:"asnName" api:"required"`
	AsnOrgName        string                                              `json:"asnOrgName" api:"required"`
	Continent         string                                              `json:"continent" api:"required"`
	GeonameID         string                                              `json:"geonameId" api:"required"`
	IP                string                                              `json:"ip" api:"required"`
	IPVersion         string                                              `json:"ipVersion" api:"required"`
	Latitude          string                                              `json:"latitude" api:"required"`
	LocationAlpha2    string                                              `json:"locationAlpha2" api:"required"`
	LocationName      string                                              `json:"locationName" api:"required"`
	Longitude         string                                              `json:"longitude" api:"required"`
	Subdivision1Name  string                                              `json:"subdivision1Name" api:"required"`
	Subdivision2Name  string                                              `json:"subdivision2Name" api:"required"`
	JSON              accountUrlscannerScanGetResponseResultScanIPsIPJSON `json:"-"`
}

// accountUrlscannerScanGetResponseResultScanIPsIPJSON contains the JSON metadata
// for the struct [AccountUrlscannerScanGetResponseResultScanIPsIP]
type accountUrlscannerScanGetResponseResultScanIPsIPJSON struct {
	Asn               apijson.Field
	AsnDescription    apijson.Field
	AsnLocationAlpha2 apijson.Field
	AsnName           apijson.Field
	AsnOrgName        apijson.Field
	Continent         apijson.Field
	GeonameID         apijson.Field
	IP                apijson.Field
	IPVersion         apijson.Field
	Latitude          apijson.Field
	LocationAlpha2    apijson.Field
	LocationName      apijson.Field
	Longitude         apijson.Field
	Subdivision1Name  apijson.Field
	Subdivision2Name  apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *AccountUrlscannerScanGetResponseResultScanIPsIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerScanGetResponseResultScanIPsIPJSON) RawJSON() string {
	return r.raw
}

type AccountUrlscannerScanGetResponseResultScanLinks struct {
	Link AccountUrlscannerScanGetResponseResultScanLinksLink `json:"link"`
	JSON accountUrlscannerScanGetResponseResultScanLinksJSON `json:"-"`
}

// accountUrlscannerScanGetResponseResultScanLinksJSON contains the JSON metadata
// for the struct [AccountUrlscannerScanGetResponseResultScanLinks]
type accountUrlscannerScanGetResponseResultScanLinksJSON struct {
	Link        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUrlscannerScanGetResponseResultScanLinks) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerScanGetResponseResultScanLinksJSON) RawJSON() string {
	return r.raw
}

type AccountUrlscannerScanGetResponseResultScanLinksLink struct {
	// Outgoing link detected in the DOM
	Href string                                                  `json:"href" api:"required"`
	Text string                                                  `json:"text" api:"required"`
	JSON accountUrlscannerScanGetResponseResultScanLinksLinkJSON `json:"-"`
}

// accountUrlscannerScanGetResponseResultScanLinksLinkJSON contains the JSON
// metadata for the struct [AccountUrlscannerScanGetResponseResultScanLinksLink]
type accountUrlscannerScanGetResponseResultScanLinksLinkJSON struct {
	Href        apijson.Field
	Text        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUrlscannerScanGetResponseResultScanLinksLink) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerScanGetResponseResultScanLinksLinkJSON) RawJSON() string {
	return r.raw
}

type AccountUrlscannerScanListResponse struct {
	Errors   []AccountUrlscannerScanListResponseError   `json:"errors" api:"required"`
	Messages []AccountUrlscannerScanListResponseMessage `json:"messages" api:"required"`
	Result   AccountUrlscannerScanListResponseResult    `json:"result" api:"required"`
	// Whether search request was successful or not
	Success bool                                  `json:"success" api:"required"`
	JSON    accountUrlscannerScanListResponseJSON `json:"-"`
}

// accountUrlscannerScanListResponseJSON contains the JSON metadata for the struct
// [AccountUrlscannerScanListResponse]
type accountUrlscannerScanListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUrlscannerScanListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerScanListResponseJSON) RawJSON() string {
	return r.raw
}

type AccountUrlscannerScanListResponseError struct {
	Message string                                     `json:"message" api:"required"`
	JSON    accountUrlscannerScanListResponseErrorJSON `json:"-"`
}

// accountUrlscannerScanListResponseErrorJSON contains the JSON metadata for the
// struct [AccountUrlscannerScanListResponseError]
type accountUrlscannerScanListResponseErrorJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUrlscannerScanListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerScanListResponseErrorJSON) RawJSON() string {
	return r.raw
}

type AccountUrlscannerScanListResponseMessage struct {
	Message string                                       `json:"message" api:"required"`
	JSON    accountUrlscannerScanListResponseMessageJSON `json:"-"`
}

// accountUrlscannerScanListResponseMessageJSON contains the JSON metadata for the
// struct [AccountUrlscannerScanListResponseMessage]
type accountUrlscannerScanListResponseMessageJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUrlscannerScanListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerScanListResponseMessageJSON) RawJSON() string {
	return r.raw
}

type AccountUrlscannerScanListResponseResult struct {
	Tasks []AccountUrlscannerScanListResponseResultTask `json:"tasks" api:"required"`
	JSON  accountUrlscannerScanListResponseResultJSON   `json:"-"`
}

// accountUrlscannerScanListResponseResultJSON contains the JSON metadata for the
// struct [AccountUrlscannerScanListResponseResult]
type accountUrlscannerScanListResponseResultJSON struct {
	Tasks       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUrlscannerScanListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerScanListResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountUrlscannerScanListResponseResultTask struct {
	// Alpha-2 country code
	Country string `json:"country" api:"required"`
	// Whether scan was successful or not
	Success bool `json:"success" api:"required"`
	// When scan was submitted (UTC)
	Time time.Time `json:"time" api:"required" format:"date-time"`
	// Scan url (after redirects)
	URL string `json:"url" api:"required"`
	// Scan id
	Uuid string `json:"uuid" api:"required" format:"uuid"`
	// Submitted visibility status.
	Visibility AccountUrlscannerScanListResponseResultTasksVisibility `json:"visibility" api:"required"`
	JSON       accountUrlscannerScanListResponseResultTaskJSON        `json:"-"`
}

// accountUrlscannerScanListResponseResultTaskJSON contains the JSON metadata for
// the struct [AccountUrlscannerScanListResponseResultTask]
type accountUrlscannerScanListResponseResultTaskJSON struct {
	Country     apijson.Field
	Success     apijson.Field
	Time        apijson.Field
	URL         apijson.Field
	Uuid        apijson.Field
	Visibility  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUrlscannerScanListResponseResultTask) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerScanListResponseResultTaskJSON) RawJSON() string {
	return r.raw
}

// Submitted visibility status.
type AccountUrlscannerScanListResponseResultTasksVisibility string

const (
	AccountUrlscannerScanListResponseResultTasksVisibilityPublic   AccountUrlscannerScanListResponseResultTasksVisibility = "public"
	AccountUrlscannerScanListResponseResultTasksVisibilityUnlisted AccountUrlscannerScanListResponseResultTasksVisibility = "unlisted"
)

func (r AccountUrlscannerScanListResponseResultTasksVisibility) IsKnown() bool {
	switch r {
	case AccountUrlscannerScanListResponseResultTasksVisibilityPublic, AccountUrlscannerScanListResponseResultTasksVisibilityUnlisted:
		return true
	}
	return false
}

type AccountUrlscannerScanGetHarResponse struct {
	Errors   []AccountUrlscannerScanGetHarResponseError   `json:"errors" api:"required"`
	Messages []AccountUrlscannerScanGetHarResponseMessage `json:"messages" api:"required"`
	Result   AccountUrlscannerScanGetHarResponseResult    `json:"result" api:"required"`
	// Whether search request was successful or not
	Success bool                                    `json:"success" api:"required"`
	JSON    accountUrlscannerScanGetHarResponseJSON `json:"-"`
}

// accountUrlscannerScanGetHarResponseJSON contains the JSON metadata for the
// struct [AccountUrlscannerScanGetHarResponse]
type accountUrlscannerScanGetHarResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUrlscannerScanGetHarResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerScanGetHarResponseJSON) RawJSON() string {
	return r.raw
}

type AccountUrlscannerScanGetHarResponseError struct {
	Message string                                       `json:"message" api:"required"`
	JSON    accountUrlscannerScanGetHarResponseErrorJSON `json:"-"`
}

// accountUrlscannerScanGetHarResponseErrorJSON contains the JSON metadata for the
// struct [AccountUrlscannerScanGetHarResponseError]
type accountUrlscannerScanGetHarResponseErrorJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUrlscannerScanGetHarResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerScanGetHarResponseErrorJSON) RawJSON() string {
	return r.raw
}

type AccountUrlscannerScanGetHarResponseMessage struct {
	Message string                                         `json:"message" api:"required"`
	JSON    accountUrlscannerScanGetHarResponseMessageJSON `json:"-"`
}

// accountUrlscannerScanGetHarResponseMessageJSON contains the JSON metadata for
// the struct [AccountUrlscannerScanGetHarResponseMessage]
type accountUrlscannerScanGetHarResponseMessageJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUrlscannerScanGetHarResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerScanGetHarResponseMessageJSON) RawJSON() string {
	return r.raw
}

type AccountUrlscannerScanGetHarResponseResult struct {
	Har  AccountUrlscannerScanGetHarResponseResultHar  `json:"har" api:"required"`
	JSON accountUrlscannerScanGetHarResponseResultJSON `json:"-"`
}

// accountUrlscannerScanGetHarResponseResultJSON contains the JSON metadata for the
// struct [AccountUrlscannerScanGetHarResponseResult]
type accountUrlscannerScanGetHarResponseResultJSON struct {
	Har         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUrlscannerScanGetHarResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerScanGetHarResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountUrlscannerScanGetHarResponseResultHar struct {
	Log  AccountUrlscannerScanGetHarResponseResultHarLog  `json:"log" api:"required"`
	JSON accountUrlscannerScanGetHarResponseResultHarJSON `json:"-"`
}

// accountUrlscannerScanGetHarResponseResultHarJSON contains the JSON metadata for
// the struct [AccountUrlscannerScanGetHarResponseResultHar]
type accountUrlscannerScanGetHarResponseResultHarJSON struct {
	Log         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUrlscannerScanGetHarResponseResultHar) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerScanGetHarResponseResultHarJSON) RawJSON() string {
	return r.raw
}

type AccountUrlscannerScanGetHarResponseResultHarLog struct {
	Creator AccountUrlscannerScanGetHarResponseResultHarLogCreator `json:"creator" api:"required"`
	Entries []AccountUrlscannerScanGetHarResponseResultHarLogEntry `json:"entries" api:"required"`
	Pages   []AccountUrlscannerScanGetHarResponseResultHarLogPage  `json:"pages" api:"required"`
	Version string                                                 `json:"version" api:"required"`
	JSON    accountUrlscannerScanGetHarResponseResultHarLogJSON    `json:"-"`
}

// accountUrlscannerScanGetHarResponseResultHarLogJSON contains the JSON metadata
// for the struct [AccountUrlscannerScanGetHarResponseResultHarLog]
type accountUrlscannerScanGetHarResponseResultHarLogJSON struct {
	Creator     apijson.Field
	Entries     apijson.Field
	Pages       apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUrlscannerScanGetHarResponseResultHarLog) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerScanGetHarResponseResultHarLogJSON) RawJSON() string {
	return r.raw
}

type AccountUrlscannerScanGetHarResponseResultHarLogCreator struct {
	Comment string                                                     `json:"comment" api:"required"`
	Name    string                                                     `json:"name" api:"required"`
	Version string                                                     `json:"version" api:"required"`
	JSON    accountUrlscannerScanGetHarResponseResultHarLogCreatorJSON `json:"-"`
}

// accountUrlscannerScanGetHarResponseResultHarLogCreatorJSON contains the JSON
// metadata for the struct [AccountUrlscannerScanGetHarResponseResultHarLogCreator]
type accountUrlscannerScanGetHarResponseResultHarLogCreatorJSON struct {
	Comment     apijson.Field
	Name        apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUrlscannerScanGetHarResponseResultHarLogCreator) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerScanGetHarResponseResultHarLogCreatorJSON) RawJSON() string {
	return r.raw
}

type AccountUrlscannerScanGetHarResponseResultHarLogEntry struct {
	InitialPriority string                                                         `json:"_initialPriority" api:"required"`
	InitiatorType   string                                                         `json:"_initiator_type" api:"required"`
	Priority        string                                                         `json:"_priority" api:"required"`
	RequestID       string                                                         `json:"_requestId" api:"required"`
	RequestTime     float64                                                        `json:"_requestTime" api:"required"`
	ResourceType    string                                                         `json:"_resourceType" api:"required"`
	Cache           interface{}                                                    `json:"cache" api:"required"`
	Connection      string                                                         `json:"connection" api:"required"`
	Pageref         string                                                         `json:"pageref" api:"required"`
	Request         AccountUrlscannerScanGetHarResponseResultHarLogEntriesRequest  `json:"request" api:"required"`
	Response        AccountUrlscannerScanGetHarResponseResultHarLogEntriesResponse `json:"response" api:"required"`
	ServerIPAddress string                                                         `json:"serverIPAddress" api:"required"`
	StartedDateTime string                                                         `json:"startedDateTime" api:"required"`
	Time            float64                                                        `json:"time" api:"required"`
	JSON            accountUrlscannerScanGetHarResponseResultHarLogEntryJSON       `json:"-"`
}

// accountUrlscannerScanGetHarResponseResultHarLogEntryJSON contains the JSON
// metadata for the struct [AccountUrlscannerScanGetHarResponseResultHarLogEntry]
type accountUrlscannerScanGetHarResponseResultHarLogEntryJSON struct {
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

func (r *AccountUrlscannerScanGetHarResponseResultHarLogEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerScanGetHarResponseResultHarLogEntryJSON) RawJSON() string {
	return r.raw
}

type AccountUrlscannerScanGetHarResponseResultHarLogEntriesRequest struct {
	BodySize    float64                                                               `json:"bodySize" api:"required"`
	Headers     []AccountUrlscannerScanGetHarResponseResultHarLogEntriesRequestHeader `json:"headers" api:"required"`
	HeadersSize float64                                                               `json:"headersSize" api:"required"`
	HTTPVersion string                                                                `json:"httpVersion" api:"required"`
	Method      string                                                                `json:"method" api:"required"`
	URL         string                                                                `json:"url" api:"required"`
	JSON        accountUrlscannerScanGetHarResponseResultHarLogEntriesRequestJSON     `json:"-"`
}

// accountUrlscannerScanGetHarResponseResultHarLogEntriesRequestJSON contains the
// JSON metadata for the struct
// [AccountUrlscannerScanGetHarResponseResultHarLogEntriesRequest]
type accountUrlscannerScanGetHarResponseResultHarLogEntriesRequestJSON struct {
	BodySize    apijson.Field
	Headers     apijson.Field
	HeadersSize apijson.Field
	HTTPVersion apijson.Field
	Method      apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUrlscannerScanGetHarResponseResultHarLogEntriesRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerScanGetHarResponseResultHarLogEntriesRequestJSON) RawJSON() string {
	return r.raw
}

type AccountUrlscannerScanGetHarResponseResultHarLogEntriesRequestHeader struct {
	Name  string                                                                  `json:"name" api:"required"`
	Value string                                                                  `json:"value" api:"required"`
	JSON  accountUrlscannerScanGetHarResponseResultHarLogEntriesRequestHeaderJSON `json:"-"`
}

// accountUrlscannerScanGetHarResponseResultHarLogEntriesRequestHeaderJSON contains
// the JSON metadata for the struct
// [AccountUrlscannerScanGetHarResponseResultHarLogEntriesRequestHeader]
type accountUrlscannerScanGetHarResponseResultHarLogEntriesRequestHeaderJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUrlscannerScanGetHarResponseResultHarLogEntriesRequestHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerScanGetHarResponseResultHarLogEntriesRequestHeaderJSON) RawJSON() string {
	return r.raw
}

type AccountUrlscannerScanGetHarResponseResultHarLogEntriesResponse struct {
	TransferSize float64                                                                `json:"_transferSize" api:"required"`
	BodySize     float64                                                                `json:"bodySize" api:"required"`
	Content      AccountUrlscannerScanGetHarResponseResultHarLogEntriesResponseContent  `json:"content" api:"required"`
	Headers      []AccountUrlscannerScanGetHarResponseResultHarLogEntriesResponseHeader `json:"headers" api:"required"`
	HeadersSize  float64                                                                `json:"headersSize" api:"required"`
	HTTPVersion  string                                                                 `json:"httpVersion" api:"required"`
	RedirectURL  string                                                                 `json:"redirectURL" api:"required"`
	Status       float64                                                                `json:"status" api:"required"`
	StatusText   string                                                                 `json:"statusText" api:"required"`
	JSON         accountUrlscannerScanGetHarResponseResultHarLogEntriesResponseJSON     `json:"-"`
}

// accountUrlscannerScanGetHarResponseResultHarLogEntriesResponseJSON contains the
// JSON metadata for the struct
// [AccountUrlscannerScanGetHarResponseResultHarLogEntriesResponse]
type accountUrlscannerScanGetHarResponseResultHarLogEntriesResponseJSON struct {
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

func (r *AccountUrlscannerScanGetHarResponseResultHarLogEntriesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerScanGetHarResponseResultHarLogEntriesResponseJSON) RawJSON() string {
	return r.raw
}

type AccountUrlscannerScanGetHarResponseResultHarLogEntriesResponseContent struct {
	MimeType    string                                                                    `json:"mimeType" api:"required"`
	Size        float64                                                                   `json:"size" api:"required"`
	Compression int64                                                                     `json:"compression"`
	JSON        accountUrlscannerScanGetHarResponseResultHarLogEntriesResponseContentJSON `json:"-"`
}

// accountUrlscannerScanGetHarResponseResultHarLogEntriesResponseContentJSON
// contains the JSON metadata for the struct
// [AccountUrlscannerScanGetHarResponseResultHarLogEntriesResponseContent]
type accountUrlscannerScanGetHarResponseResultHarLogEntriesResponseContentJSON struct {
	MimeType    apijson.Field
	Size        apijson.Field
	Compression apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUrlscannerScanGetHarResponseResultHarLogEntriesResponseContent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerScanGetHarResponseResultHarLogEntriesResponseContentJSON) RawJSON() string {
	return r.raw
}

type AccountUrlscannerScanGetHarResponseResultHarLogEntriesResponseHeader struct {
	Name  string                                                                   `json:"name" api:"required"`
	Value string                                                                   `json:"value" api:"required"`
	JSON  accountUrlscannerScanGetHarResponseResultHarLogEntriesResponseHeaderJSON `json:"-"`
}

// accountUrlscannerScanGetHarResponseResultHarLogEntriesResponseHeaderJSON
// contains the JSON metadata for the struct
// [AccountUrlscannerScanGetHarResponseResultHarLogEntriesResponseHeader]
type accountUrlscannerScanGetHarResponseResultHarLogEntriesResponseHeaderJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUrlscannerScanGetHarResponseResultHarLogEntriesResponseHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerScanGetHarResponseResultHarLogEntriesResponseHeaderJSON) RawJSON() string {
	return r.raw
}

type AccountUrlscannerScanGetHarResponseResultHarLogPage struct {
	ID              string                                                          `json:"id" api:"required"`
	PageTimings     AccountUrlscannerScanGetHarResponseResultHarLogPagesPageTimings `json:"pageTimings" api:"required"`
	StartedDateTime string                                                          `json:"startedDateTime" api:"required"`
	Title           string                                                          `json:"title" api:"required"`
	JSON            accountUrlscannerScanGetHarResponseResultHarLogPageJSON         `json:"-"`
}

// accountUrlscannerScanGetHarResponseResultHarLogPageJSON contains the JSON
// metadata for the struct [AccountUrlscannerScanGetHarResponseResultHarLogPage]
type accountUrlscannerScanGetHarResponseResultHarLogPageJSON struct {
	ID              apijson.Field
	PageTimings     apijson.Field
	StartedDateTime apijson.Field
	Title           apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccountUrlscannerScanGetHarResponseResultHarLogPage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerScanGetHarResponseResultHarLogPageJSON) RawJSON() string {
	return r.raw
}

type AccountUrlscannerScanGetHarResponseResultHarLogPagesPageTimings struct {
	OnContentLoad float64                                                             `json:"onContentLoad" api:"required"`
	OnLoad        float64                                                             `json:"onLoad" api:"required"`
	JSON          accountUrlscannerScanGetHarResponseResultHarLogPagesPageTimingsJSON `json:"-"`
}

// accountUrlscannerScanGetHarResponseResultHarLogPagesPageTimingsJSON contains the
// JSON metadata for the struct
// [AccountUrlscannerScanGetHarResponseResultHarLogPagesPageTimings]
type accountUrlscannerScanGetHarResponseResultHarLogPagesPageTimingsJSON struct {
	OnContentLoad apijson.Field
	OnLoad        apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccountUrlscannerScanGetHarResponseResultHarLogPagesPageTimings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUrlscannerScanGetHarResponseResultHarLogPagesPageTimingsJSON) RawJSON() string {
	return r.raw
}

type AccountUrlscannerScanNewParams struct {
	URL param.Field[string] `json:"url" api:"required"`
	// Country to geo egress from
	Country param.Field[AccountUrlscannerScanNewParamsCountry] `json:"country"`
	// Set custom headers.
	CustomHeaders param.Field[map[string]string] `json:"customHeaders"`
	// Take multiple screenshots targeting different device types.
	ScreenshotsResolutions param.Field[[]AccountUrlscannerScanNewParamsScreenshotsResolution] `json:"screenshotsResolutions"`
	// The option `Public` means it will be included in listings like recent scans and
	// search results. `Unlisted` means it will not be included in the aforementioned
	// listings, users will need to have the scan's ID to access it. A a scan will be
	// automatically marked as unlisted if it fails, if it contains potential PII or
	// other sensitive material.
	Visibility param.Field[AccountUrlscannerScanNewParamsVisibility] `json:"visibility"`
}

func (r AccountUrlscannerScanNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Country to geo egress from
type AccountUrlscannerScanNewParamsCountry string

const (
	AccountUrlscannerScanNewParamsCountryAf AccountUrlscannerScanNewParamsCountry = "AF"
	AccountUrlscannerScanNewParamsCountryAl AccountUrlscannerScanNewParamsCountry = "AL"
	AccountUrlscannerScanNewParamsCountryDz AccountUrlscannerScanNewParamsCountry = "DZ"
	AccountUrlscannerScanNewParamsCountryAd AccountUrlscannerScanNewParamsCountry = "AD"
	AccountUrlscannerScanNewParamsCountryAo AccountUrlscannerScanNewParamsCountry = "AO"
	AccountUrlscannerScanNewParamsCountryAg AccountUrlscannerScanNewParamsCountry = "AG"
	AccountUrlscannerScanNewParamsCountryAr AccountUrlscannerScanNewParamsCountry = "AR"
	AccountUrlscannerScanNewParamsCountryAm AccountUrlscannerScanNewParamsCountry = "AM"
	AccountUrlscannerScanNewParamsCountryAu AccountUrlscannerScanNewParamsCountry = "AU"
	AccountUrlscannerScanNewParamsCountryAt AccountUrlscannerScanNewParamsCountry = "AT"
	AccountUrlscannerScanNewParamsCountryAz AccountUrlscannerScanNewParamsCountry = "AZ"
	AccountUrlscannerScanNewParamsCountryBh AccountUrlscannerScanNewParamsCountry = "BH"
	AccountUrlscannerScanNewParamsCountryBd AccountUrlscannerScanNewParamsCountry = "BD"
	AccountUrlscannerScanNewParamsCountryBb AccountUrlscannerScanNewParamsCountry = "BB"
	AccountUrlscannerScanNewParamsCountryBy AccountUrlscannerScanNewParamsCountry = "BY"
	AccountUrlscannerScanNewParamsCountryBe AccountUrlscannerScanNewParamsCountry = "BE"
	AccountUrlscannerScanNewParamsCountryBz AccountUrlscannerScanNewParamsCountry = "BZ"
	AccountUrlscannerScanNewParamsCountryBj AccountUrlscannerScanNewParamsCountry = "BJ"
	AccountUrlscannerScanNewParamsCountryBm AccountUrlscannerScanNewParamsCountry = "BM"
	AccountUrlscannerScanNewParamsCountryBt AccountUrlscannerScanNewParamsCountry = "BT"
	AccountUrlscannerScanNewParamsCountryBo AccountUrlscannerScanNewParamsCountry = "BO"
	AccountUrlscannerScanNewParamsCountryBa AccountUrlscannerScanNewParamsCountry = "BA"
	AccountUrlscannerScanNewParamsCountryBw AccountUrlscannerScanNewParamsCountry = "BW"
	AccountUrlscannerScanNewParamsCountryBr AccountUrlscannerScanNewParamsCountry = "BR"
	AccountUrlscannerScanNewParamsCountryBn AccountUrlscannerScanNewParamsCountry = "BN"
	AccountUrlscannerScanNewParamsCountryBg AccountUrlscannerScanNewParamsCountry = "BG"
	AccountUrlscannerScanNewParamsCountryBf AccountUrlscannerScanNewParamsCountry = "BF"
	AccountUrlscannerScanNewParamsCountryBi AccountUrlscannerScanNewParamsCountry = "BI"
	AccountUrlscannerScanNewParamsCountryKh AccountUrlscannerScanNewParamsCountry = "KH"
	AccountUrlscannerScanNewParamsCountryCm AccountUrlscannerScanNewParamsCountry = "CM"
	AccountUrlscannerScanNewParamsCountryCa AccountUrlscannerScanNewParamsCountry = "CA"
	AccountUrlscannerScanNewParamsCountryCv AccountUrlscannerScanNewParamsCountry = "CV"
	AccountUrlscannerScanNewParamsCountryKy AccountUrlscannerScanNewParamsCountry = "KY"
	AccountUrlscannerScanNewParamsCountryCf AccountUrlscannerScanNewParamsCountry = "CF"
	AccountUrlscannerScanNewParamsCountryTd AccountUrlscannerScanNewParamsCountry = "TD"
	AccountUrlscannerScanNewParamsCountryCl AccountUrlscannerScanNewParamsCountry = "CL"
	AccountUrlscannerScanNewParamsCountryCn AccountUrlscannerScanNewParamsCountry = "CN"
	AccountUrlscannerScanNewParamsCountryCo AccountUrlscannerScanNewParamsCountry = "CO"
	AccountUrlscannerScanNewParamsCountryKm AccountUrlscannerScanNewParamsCountry = "KM"
	AccountUrlscannerScanNewParamsCountryCg AccountUrlscannerScanNewParamsCountry = "CG"
	AccountUrlscannerScanNewParamsCountryCr AccountUrlscannerScanNewParamsCountry = "CR"
	AccountUrlscannerScanNewParamsCountryCi AccountUrlscannerScanNewParamsCountry = "CI"
	AccountUrlscannerScanNewParamsCountryHr AccountUrlscannerScanNewParamsCountry = "HR"
	AccountUrlscannerScanNewParamsCountryCu AccountUrlscannerScanNewParamsCountry = "CU"
	AccountUrlscannerScanNewParamsCountryCy AccountUrlscannerScanNewParamsCountry = "CY"
	AccountUrlscannerScanNewParamsCountryCz AccountUrlscannerScanNewParamsCountry = "CZ"
	AccountUrlscannerScanNewParamsCountryCd AccountUrlscannerScanNewParamsCountry = "CD"
	AccountUrlscannerScanNewParamsCountryDk AccountUrlscannerScanNewParamsCountry = "DK"
	AccountUrlscannerScanNewParamsCountryDj AccountUrlscannerScanNewParamsCountry = "DJ"
	AccountUrlscannerScanNewParamsCountryDm AccountUrlscannerScanNewParamsCountry = "DM"
	AccountUrlscannerScanNewParamsCountryDo AccountUrlscannerScanNewParamsCountry = "DO"
	AccountUrlscannerScanNewParamsCountryEc AccountUrlscannerScanNewParamsCountry = "EC"
	AccountUrlscannerScanNewParamsCountryEg AccountUrlscannerScanNewParamsCountry = "EG"
	AccountUrlscannerScanNewParamsCountrySv AccountUrlscannerScanNewParamsCountry = "SV"
	AccountUrlscannerScanNewParamsCountryGq AccountUrlscannerScanNewParamsCountry = "GQ"
	AccountUrlscannerScanNewParamsCountryEr AccountUrlscannerScanNewParamsCountry = "ER"
	AccountUrlscannerScanNewParamsCountryEe AccountUrlscannerScanNewParamsCountry = "EE"
	AccountUrlscannerScanNewParamsCountrySz AccountUrlscannerScanNewParamsCountry = "SZ"
	AccountUrlscannerScanNewParamsCountryEt AccountUrlscannerScanNewParamsCountry = "ET"
	AccountUrlscannerScanNewParamsCountryFj AccountUrlscannerScanNewParamsCountry = "FJ"
	AccountUrlscannerScanNewParamsCountryFi AccountUrlscannerScanNewParamsCountry = "FI"
	AccountUrlscannerScanNewParamsCountryFr AccountUrlscannerScanNewParamsCountry = "FR"
	AccountUrlscannerScanNewParamsCountryGa AccountUrlscannerScanNewParamsCountry = "GA"
	AccountUrlscannerScanNewParamsCountryGe AccountUrlscannerScanNewParamsCountry = "GE"
	AccountUrlscannerScanNewParamsCountryDe AccountUrlscannerScanNewParamsCountry = "DE"
	AccountUrlscannerScanNewParamsCountryGh AccountUrlscannerScanNewParamsCountry = "GH"
	AccountUrlscannerScanNewParamsCountryGr AccountUrlscannerScanNewParamsCountry = "GR"
	AccountUrlscannerScanNewParamsCountryGl AccountUrlscannerScanNewParamsCountry = "GL"
	AccountUrlscannerScanNewParamsCountryGd AccountUrlscannerScanNewParamsCountry = "GD"
	AccountUrlscannerScanNewParamsCountryGt AccountUrlscannerScanNewParamsCountry = "GT"
	AccountUrlscannerScanNewParamsCountryGn AccountUrlscannerScanNewParamsCountry = "GN"
	AccountUrlscannerScanNewParamsCountryGw AccountUrlscannerScanNewParamsCountry = "GW"
	AccountUrlscannerScanNewParamsCountryGy AccountUrlscannerScanNewParamsCountry = "GY"
	AccountUrlscannerScanNewParamsCountryHt AccountUrlscannerScanNewParamsCountry = "HT"
	AccountUrlscannerScanNewParamsCountryHn AccountUrlscannerScanNewParamsCountry = "HN"
	AccountUrlscannerScanNewParamsCountryHu AccountUrlscannerScanNewParamsCountry = "HU"
	AccountUrlscannerScanNewParamsCountryIs AccountUrlscannerScanNewParamsCountry = "IS"
	AccountUrlscannerScanNewParamsCountryIn AccountUrlscannerScanNewParamsCountry = "IN"
	AccountUrlscannerScanNewParamsCountryID AccountUrlscannerScanNewParamsCountry = "ID"
	AccountUrlscannerScanNewParamsCountryIr AccountUrlscannerScanNewParamsCountry = "IR"
	AccountUrlscannerScanNewParamsCountryIq AccountUrlscannerScanNewParamsCountry = "IQ"
	AccountUrlscannerScanNewParamsCountryIe AccountUrlscannerScanNewParamsCountry = "IE"
	AccountUrlscannerScanNewParamsCountryIl AccountUrlscannerScanNewParamsCountry = "IL"
	AccountUrlscannerScanNewParamsCountryIt AccountUrlscannerScanNewParamsCountry = "IT"
	AccountUrlscannerScanNewParamsCountryJm AccountUrlscannerScanNewParamsCountry = "JM"
	AccountUrlscannerScanNewParamsCountryJp AccountUrlscannerScanNewParamsCountry = "JP"
	AccountUrlscannerScanNewParamsCountryJo AccountUrlscannerScanNewParamsCountry = "JO"
	AccountUrlscannerScanNewParamsCountryKz AccountUrlscannerScanNewParamsCountry = "KZ"
	AccountUrlscannerScanNewParamsCountryKe AccountUrlscannerScanNewParamsCountry = "KE"
	AccountUrlscannerScanNewParamsCountryKi AccountUrlscannerScanNewParamsCountry = "KI"
	AccountUrlscannerScanNewParamsCountryKw AccountUrlscannerScanNewParamsCountry = "KW"
	AccountUrlscannerScanNewParamsCountryKg AccountUrlscannerScanNewParamsCountry = "KG"
	AccountUrlscannerScanNewParamsCountryLa AccountUrlscannerScanNewParamsCountry = "LA"
	AccountUrlscannerScanNewParamsCountryLv AccountUrlscannerScanNewParamsCountry = "LV"
	AccountUrlscannerScanNewParamsCountryLb AccountUrlscannerScanNewParamsCountry = "LB"
	AccountUrlscannerScanNewParamsCountryLs AccountUrlscannerScanNewParamsCountry = "LS"
	AccountUrlscannerScanNewParamsCountryLr AccountUrlscannerScanNewParamsCountry = "LR"
	AccountUrlscannerScanNewParamsCountryLy AccountUrlscannerScanNewParamsCountry = "LY"
	AccountUrlscannerScanNewParamsCountryLi AccountUrlscannerScanNewParamsCountry = "LI"
	AccountUrlscannerScanNewParamsCountryLt AccountUrlscannerScanNewParamsCountry = "LT"
	AccountUrlscannerScanNewParamsCountryLu AccountUrlscannerScanNewParamsCountry = "LU"
	AccountUrlscannerScanNewParamsCountryMo AccountUrlscannerScanNewParamsCountry = "MO"
	AccountUrlscannerScanNewParamsCountryMg AccountUrlscannerScanNewParamsCountry = "MG"
	AccountUrlscannerScanNewParamsCountryMw AccountUrlscannerScanNewParamsCountry = "MW"
	AccountUrlscannerScanNewParamsCountryMy AccountUrlscannerScanNewParamsCountry = "MY"
	AccountUrlscannerScanNewParamsCountryMv AccountUrlscannerScanNewParamsCountry = "MV"
	AccountUrlscannerScanNewParamsCountryMl AccountUrlscannerScanNewParamsCountry = "ML"
	AccountUrlscannerScanNewParamsCountryMr AccountUrlscannerScanNewParamsCountry = "MR"
	AccountUrlscannerScanNewParamsCountryMu AccountUrlscannerScanNewParamsCountry = "MU"
	AccountUrlscannerScanNewParamsCountryMx AccountUrlscannerScanNewParamsCountry = "MX"
	AccountUrlscannerScanNewParamsCountryFm AccountUrlscannerScanNewParamsCountry = "FM"
	AccountUrlscannerScanNewParamsCountryMd AccountUrlscannerScanNewParamsCountry = "MD"
	AccountUrlscannerScanNewParamsCountryMc AccountUrlscannerScanNewParamsCountry = "MC"
	AccountUrlscannerScanNewParamsCountryMn AccountUrlscannerScanNewParamsCountry = "MN"
	AccountUrlscannerScanNewParamsCountryMs AccountUrlscannerScanNewParamsCountry = "MS"
	AccountUrlscannerScanNewParamsCountryMa AccountUrlscannerScanNewParamsCountry = "MA"
	AccountUrlscannerScanNewParamsCountryMz AccountUrlscannerScanNewParamsCountry = "MZ"
	AccountUrlscannerScanNewParamsCountryMm AccountUrlscannerScanNewParamsCountry = "MM"
	AccountUrlscannerScanNewParamsCountryNa AccountUrlscannerScanNewParamsCountry = "NA"
	AccountUrlscannerScanNewParamsCountryNr AccountUrlscannerScanNewParamsCountry = "NR"
	AccountUrlscannerScanNewParamsCountryNp AccountUrlscannerScanNewParamsCountry = "NP"
	AccountUrlscannerScanNewParamsCountryNl AccountUrlscannerScanNewParamsCountry = "NL"
	AccountUrlscannerScanNewParamsCountryNz AccountUrlscannerScanNewParamsCountry = "NZ"
	AccountUrlscannerScanNewParamsCountryNi AccountUrlscannerScanNewParamsCountry = "NI"
	AccountUrlscannerScanNewParamsCountryNe AccountUrlscannerScanNewParamsCountry = "NE"
	AccountUrlscannerScanNewParamsCountryNg AccountUrlscannerScanNewParamsCountry = "NG"
	AccountUrlscannerScanNewParamsCountryKp AccountUrlscannerScanNewParamsCountry = "KP"
	AccountUrlscannerScanNewParamsCountryMk AccountUrlscannerScanNewParamsCountry = "MK"
	AccountUrlscannerScanNewParamsCountryNo AccountUrlscannerScanNewParamsCountry = "NO"
	AccountUrlscannerScanNewParamsCountryOm AccountUrlscannerScanNewParamsCountry = "OM"
	AccountUrlscannerScanNewParamsCountryPk AccountUrlscannerScanNewParamsCountry = "PK"
	AccountUrlscannerScanNewParamsCountryPs AccountUrlscannerScanNewParamsCountry = "PS"
	AccountUrlscannerScanNewParamsCountryPa AccountUrlscannerScanNewParamsCountry = "PA"
	AccountUrlscannerScanNewParamsCountryPg AccountUrlscannerScanNewParamsCountry = "PG"
	AccountUrlscannerScanNewParamsCountryPy AccountUrlscannerScanNewParamsCountry = "PY"
	AccountUrlscannerScanNewParamsCountryPe AccountUrlscannerScanNewParamsCountry = "PE"
	AccountUrlscannerScanNewParamsCountryPh AccountUrlscannerScanNewParamsCountry = "PH"
	AccountUrlscannerScanNewParamsCountryPl AccountUrlscannerScanNewParamsCountry = "PL"
	AccountUrlscannerScanNewParamsCountryPt AccountUrlscannerScanNewParamsCountry = "PT"
	AccountUrlscannerScanNewParamsCountryQa AccountUrlscannerScanNewParamsCountry = "QA"
	AccountUrlscannerScanNewParamsCountryRo AccountUrlscannerScanNewParamsCountry = "RO"
	AccountUrlscannerScanNewParamsCountryRu AccountUrlscannerScanNewParamsCountry = "RU"
	AccountUrlscannerScanNewParamsCountryRw AccountUrlscannerScanNewParamsCountry = "RW"
	AccountUrlscannerScanNewParamsCountrySh AccountUrlscannerScanNewParamsCountry = "SH"
	AccountUrlscannerScanNewParamsCountryKn AccountUrlscannerScanNewParamsCountry = "KN"
	AccountUrlscannerScanNewParamsCountryLc AccountUrlscannerScanNewParamsCountry = "LC"
	AccountUrlscannerScanNewParamsCountryVc AccountUrlscannerScanNewParamsCountry = "VC"
	AccountUrlscannerScanNewParamsCountryWs AccountUrlscannerScanNewParamsCountry = "WS"
	AccountUrlscannerScanNewParamsCountrySm AccountUrlscannerScanNewParamsCountry = "SM"
	AccountUrlscannerScanNewParamsCountrySt AccountUrlscannerScanNewParamsCountry = "ST"
	AccountUrlscannerScanNewParamsCountrySa AccountUrlscannerScanNewParamsCountry = "SA"
	AccountUrlscannerScanNewParamsCountrySn AccountUrlscannerScanNewParamsCountry = "SN"
	AccountUrlscannerScanNewParamsCountryRs AccountUrlscannerScanNewParamsCountry = "RS"
	AccountUrlscannerScanNewParamsCountrySc AccountUrlscannerScanNewParamsCountry = "SC"
	AccountUrlscannerScanNewParamsCountrySl AccountUrlscannerScanNewParamsCountry = "SL"
	AccountUrlscannerScanNewParamsCountrySk AccountUrlscannerScanNewParamsCountry = "SK"
	AccountUrlscannerScanNewParamsCountrySi AccountUrlscannerScanNewParamsCountry = "SI"
	AccountUrlscannerScanNewParamsCountrySb AccountUrlscannerScanNewParamsCountry = "SB"
	AccountUrlscannerScanNewParamsCountrySo AccountUrlscannerScanNewParamsCountry = "SO"
	AccountUrlscannerScanNewParamsCountryZa AccountUrlscannerScanNewParamsCountry = "ZA"
	AccountUrlscannerScanNewParamsCountryKr AccountUrlscannerScanNewParamsCountry = "KR"
	AccountUrlscannerScanNewParamsCountrySS AccountUrlscannerScanNewParamsCountry = "SS"
	AccountUrlscannerScanNewParamsCountryEs AccountUrlscannerScanNewParamsCountry = "ES"
	AccountUrlscannerScanNewParamsCountryLk AccountUrlscannerScanNewParamsCountry = "LK"
	AccountUrlscannerScanNewParamsCountrySd AccountUrlscannerScanNewParamsCountry = "SD"
	AccountUrlscannerScanNewParamsCountrySr AccountUrlscannerScanNewParamsCountry = "SR"
	AccountUrlscannerScanNewParamsCountrySe AccountUrlscannerScanNewParamsCountry = "SE"
	AccountUrlscannerScanNewParamsCountryCh AccountUrlscannerScanNewParamsCountry = "CH"
	AccountUrlscannerScanNewParamsCountrySy AccountUrlscannerScanNewParamsCountry = "SY"
	AccountUrlscannerScanNewParamsCountryTw AccountUrlscannerScanNewParamsCountry = "TW"
	AccountUrlscannerScanNewParamsCountryTj AccountUrlscannerScanNewParamsCountry = "TJ"
	AccountUrlscannerScanNewParamsCountryTz AccountUrlscannerScanNewParamsCountry = "TZ"
	AccountUrlscannerScanNewParamsCountryTh AccountUrlscannerScanNewParamsCountry = "TH"
	AccountUrlscannerScanNewParamsCountryBs AccountUrlscannerScanNewParamsCountry = "BS"
	AccountUrlscannerScanNewParamsCountryGm AccountUrlscannerScanNewParamsCountry = "GM"
	AccountUrlscannerScanNewParamsCountryTl AccountUrlscannerScanNewParamsCountry = "TL"
	AccountUrlscannerScanNewParamsCountryTg AccountUrlscannerScanNewParamsCountry = "TG"
	AccountUrlscannerScanNewParamsCountryTo AccountUrlscannerScanNewParamsCountry = "TO"
	AccountUrlscannerScanNewParamsCountryTt AccountUrlscannerScanNewParamsCountry = "TT"
	AccountUrlscannerScanNewParamsCountryTn AccountUrlscannerScanNewParamsCountry = "TN"
	AccountUrlscannerScanNewParamsCountryTr AccountUrlscannerScanNewParamsCountry = "TR"
	AccountUrlscannerScanNewParamsCountryTm AccountUrlscannerScanNewParamsCountry = "TM"
	AccountUrlscannerScanNewParamsCountryUg AccountUrlscannerScanNewParamsCountry = "UG"
	AccountUrlscannerScanNewParamsCountryUa AccountUrlscannerScanNewParamsCountry = "UA"
	AccountUrlscannerScanNewParamsCountryAe AccountUrlscannerScanNewParamsCountry = "AE"
	AccountUrlscannerScanNewParamsCountryGB AccountUrlscannerScanNewParamsCountry = "GB"
	AccountUrlscannerScanNewParamsCountryUs AccountUrlscannerScanNewParamsCountry = "US"
	AccountUrlscannerScanNewParamsCountryUy AccountUrlscannerScanNewParamsCountry = "UY"
	AccountUrlscannerScanNewParamsCountryUz AccountUrlscannerScanNewParamsCountry = "UZ"
	AccountUrlscannerScanNewParamsCountryVu AccountUrlscannerScanNewParamsCountry = "VU"
	AccountUrlscannerScanNewParamsCountryVe AccountUrlscannerScanNewParamsCountry = "VE"
	AccountUrlscannerScanNewParamsCountryVn AccountUrlscannerScanNewParamsCountry = "VN"
	AccountUrlscannerScanNewParamsCountryYe AccountUrlscannerScanNewParamsCountry = "YE"
	AccountUrlscannerScanNewParamsCountryZm AccountUrlscannerScanNewParamsCountry = "ZM"
	AccountUrlscannerScanNewParamsCountryZw AccountUrlscannerScanNewParamsCountry = "ZW"
)

func (r AccountUrlscannerScanNewParamsCountry) IsKnown() bool {
	switch r {
	case AccountUrlscannerScanNewParamsCountryAf, AccountUrlscannerScanNewParamsCountryAl, AccountUrlscannerScanNewParamsCountryDz, AccountUrlscannerScanNewParamsCountryAd, AccountUrlscannerScanNewParamsCountryAo, AccountUrlscannerScanNewParamsCountryAg, AccountUrlscannerScanNewParamsCountryAr, AccountUrlscannerScanNewParamsCountryAm, AccountUrlscannerScanNewParamsCountryAu, AccountUrlscannerScanNewParamsCountryAt, AccountUrlscannerScanNewParamsCountryAz, AccountUrlscannerScanNewParamsCountryBh, AccountUrlscannerScanNewParamsCountryBd, AccountUrlscannerScanNewParamsCountryBb, AccountUrlscannerScanNewParamsCountryBy, AccountUrlscannerScanNewParamsCountryBe, AccountUrlscannerScanNewParamsCountryBz, AccountUrlscannerScanNewParamsCountryBj, AccountUrlscannerScanNewParamsCountryBm, AccountUrlscannerScanNewParamsCountryBt, AccountUrlscannerScanNewParamsCountryBo, AccountUrlscannerScanNewParamsCountryBa, AccountUrlscannerScanNewParamsCountryBw, AccountUrlscannerScanNewParamsCountryBr, AccountUrlscannerScanNewParamsCountryBn, AccountUrlscannerScanNewParamsCountryBg, AccountUrlscannerScanNewParamsCountryBf, AccountUrlscannerScanNewParamsCountryBi, AccountUrlscannerScanNewParamsCountryKh, AccountUrlscannerScanNewParamsCountryCm, AccountUrlscannerScanNewParamsCountryCa, AccountUrlscannerScanNewParamsCountryCv, AccountUrlscannerScanNewParamsCountryKy, AccountUrlscannerScanNewParamsCountryCf, AccountUrlscannerScanNewParamsCountryTd, AccountUrlscannerScanNewParamsCountryCl, AccountUrlscannerScanNewParamsCountryCn, AccountUrlscannerScanNewParamsCountryCo, AccountUrlscannerScanNewParamsCountryKm, AccountUrlscannerScanNewParamsCountryCg, AccountUrlscannerScanNewParamsCountryCr, AccountUrlscannerScanNewParamsCountryCi, AccountUrlscannerScanNewParamsCountryHr, AccountUrlscannerScanNewParamsCountryCu, AccountUrlscannerScanNewParamsCountryCy, AccountUrlscannerScanNewParamsCountryCz, AccountUrlscannerScanNewParamsCountryCd, AccountUrlscannerScanNewParamsCountryDk, AccountUrlscannerScanNewParamsCountryDj, AccountUrlscannerScanNewParamsCountryDm, AccountUrlscannerScanNewParamsCountryDo, AccountUrlscannerScanNewParamsCountryEc, AccountUrlscannerScanNewParamsCountryEg, AccountUrlscannerScanNewParamsCountrySv, AccountUrlscannerScanNewParamsCountryGq, AccountUrlscannerScanNewParamsCountryEr, AccountUrlscannerScanNewParamsCountryEe, AccountUrlscannerScanNewParamsCountrySz, AccountUrlscannerScanNewParamsCountryEt, AccountUrlscannerScanNewParamsCountryFj, AccountUrlscannerScanNewParamsCountryFi, AccountUrlscannerScanNewParamsCountryFr, AccountUrlscannerScanNewParamsCountryGa, AccountUrlscannerScanNewParamsCountryGe, AccountUrlscannerScanNewParamsCountryDe, AccountUrlscannerScanNewParamsCountryGh, AccountUrlscannerScanNewParamsCountryGr, AccountUrlscannerScanNewParamsCountryGl, AccountUrlscannerScanNewParamsCountryGd, AccountUrlscannerScanNewParamsCountryGt, AccountUrlscannerScanNewParamsCountryGn, AccountUrlscannerScanNewParamsCountryGw, AccountUrlscannerScanNewParamsCountryGy, AccountUrlscannerScanNewParamsCountryHt, AccountUrlscannerScanNewParamsCountryHn, AccountUrlscannerScanNewParamsCountryHu, AccountUrlscannerScanNewParamsCountryIs, AccountUrlscannerScanNewParamsCountryIn, AccountUrlscannerScanNewParamsCountryID, AccountUrlscannerScanNewParamsCountryIr, AccountUrlscannerScanNewParamsCountryIq, AccountUrlscannerScanNewParamsCountryIe, AccountUrlscannerScanNewParamsCountryIl, AccountUrlscannerScanNewParamsCountryIt, AccountUrlscannerScanNewParamsCountryJm, AccountUrlscannerScanNewParamsCountryJp, AccountUrlscannerScanNewParamsCountryJo, AccountUrlscannerScanNewParamsCountryKz, AccountUrlscannerScanNewParamsCountryKe, AccountUrlscannerScanNewParamsCountryKi, AccountUrlscannerScanNewParamsCountryKw, AccountUrlscannerScanNewParamsCountryKg, AccountUrlscannerScanNewParamsCountryLa, AccountUrlscannerScanNewParamsCountryLv, AccountUrlscannerScanNewParamsCountryLb, AccountUrlscannerScanNewParamsCountryLs, AccountUrlscannerScanNewParamsCountryLr, AccountUrlscannerScanNewParamsCountryLy, AccountUrlscannerScanNewParamsCountryLi, AccountUrlscannerScanNewParamsCountryLt, AccountUrlscannerScanNewParamsCountryLu, AccountUrlscannerScanNewParamsCountryMo, AccountUrlscannerScanNewParamsCountryMg, AccountUrlscannerScanNewParamsCountryMw, AccountUrlscannerScanNewParamsCountryMy, AccountUrlscannerScanNewParamsCountryMv, AccountUrlscannerScanNewParamsCountryMl, AccountUrlscannerScanNewParamsCountryMr, AccountUrlscannerScanNewParamsCountryMu, AccountUrlscannerScanNewParamsCountryMx, AccountUrlscannerScanNewParamsCountryFm, AccountUrlscannerScanNewParamsCountryMd, AccountUrlscannerScanNewParamsCountryMc, AccountUrlscannerScanNewParamsCountryMn, AccountUrlscannerScanNewParamsCountryMs, AccountUrlscannerScanNewParamsCountryMa, AccountUrlscannerScanNewParamsCountryMz, AccountUrlscannerScanNewParamsCountryMm, AccountUrlscannerScanNewParamsCountryNa, AccountUrlscannerScanNewParamsCountryNr, AccountUrlscannerScanNewParamsCountryNp, AccountUrlscannerScanNewParamsCountryNl, AccountUrlscannerScanNewParamsCountryNz, AccountUrlscannerScanNewParamsCountryNi, AccountUrlscannerScanNewParamsCountryNe, AccountUrlscannerScanNewParamsCountryNg, AccountUrlscannerScanNewParamsCountryKp, AccountUrlscannerScanNewParamsCountryMk, AccountUrlscannerScanNewParamsCountryNo, AccountUrlscannerScanNewParamsCountryOm, AccountUrlscannerScanNewParamsCountryPk, AccountUrlscannerScanNewParamsCountryPs, AccountUrlscannerScanNewParamsCountryPa, AccountUrlscannerScanNewParamsCountryPg, AccountUrlscannerScanNewParamsCountryPy, AccountUrlscannerScanNewParamsCountryPe, AccountUrlscannerScanNewParamsCountryPh, AccountUrlscannerScanNewParamsCountryPl, AccountUrlscannerScanNewParamsCountryPt, AccountUrlscannerScanNewParamsCountryQa, AccountUrlscannerScanNewParamsCountryRo, AccountUrlscannerScanNewParamsCountryRu, AccountUrlscannerScanNewParamsCountryRw, AccountUrlscannerScanNewParamsCountrySh, AccountUrlscannerScanNewParamsCountryKn, AccountUrlscannerScanNewParamsCountryLc, AccountUrlscannerScanNewParamsCountryVc, AccountUrlscannerScanNewParamsCountryWs, AccountUrlscannerScanNewParamsCountrySm, AccountUrlscannerScanNewParamsCountrySt, AccountUrlscannerScanNewParamsCountrySa, AccountUrlscannerScanNewParamsCountrySn, AccountUrlscannerScanNewParamsCountryRs, AccountUrlscannerScanNewParamsCountrySc, AccountUrlscannerScanNewParamsCountrySl, AccountUrlscannerScanNewParamsCountrySk, AccountUrlscannerScanNewParamsCountrySi, AccountUrlscannerScanNewParamsCountrySb, AccountUrlscannerScanNewParamsCountrySo, AccountUrlscannerScanNewParamsCountryZa, AccountUrlscannerScanNewParamsCountryKr, AccountUrlscannerScanNewParamsCountrySS, AccountUrlscannerScanNewParamsCountryEs, AccountUrlscannerScanNewParamsCountryLk, AccountUrlscannerScanNewParamsCountrySd, AccountUrlscannerScanNewParamsCountrySr, AccountUrlscannerScanNewParamsCountrySe, AccountUrlscannerScanNewParamsCountryCh, AccountUrlscannerScanNewParamsCountrySy, AccountUrlscannerScanNewParamsCountryTw, AccountUrlscannerScanNewParamsCountryTj, AccountUrlscannerScanNewParamsCountryTz, AccountUrlscannerScanNewParamsCountryTh, AccountUrlscannerScanNewParamsCountryBs, AccountUrlscannerScanNewParamsCountryGm, AccountUrlscannerScanNewParamsCountryTl, AccountUrlscannerScanNewParamsCountryTg, AccountUrlscannerScanNewParamsCountryTo, AccountUrlscannerScanNewParamsCountryTt, AccountUrlscannerScanNewParamsCountryTn, AccountUrlscannerScanNewParamsCountryTr, AccountUrlscannerScanNewParamsCountryTm, AccountUrlscannerScanNewParamsCountryUg, AccountUrlscannerScanNewParamsCountryUa, AccountUrlscannerScanNewParamsCountryAe, AccountUrlscannerScanNewParamsCountryGB, AccountUrlscannerScanNewParamsCountryUs, AccountUrlscannerScanNewParamsCountryUy, AccountUrlscannerScanNewParamsCountryUz, AccountUrlscannerScanNewParamsCountryVu, AccountUrlscannerScanNewParamsCountryVe, AccountUrlscannerScanNewParamsCountryVn, AccountUrlscannerScanNewParamsCountryYe, AccountUrlscannerScanNewParamsCountryZm, AccountUrlscannerScanNewParamsCountryZw:
		return true
	}
	return false
}

// Device resolutions.
type AccountUrlscannerScanNewParamsScreenshotsResolution string

const (
	AccountUrlscannerScanNewParamsScreenshotsResolutionDesktop AccountUrlscannerScanNewParamsScreenshotsResolution = "desktop"
	AccountUrlscannerScanNewParamsScreenshotsResolutionMobile  AccountUrlscannerScanNewParamsScreenshotsResolution = "mobile"
	AccountUrlscannerScanNewParamsScreenshotsResolutionTablet  AccountUrlscannerScanNewParamsScreenshotsResolution = "tablet"
)

func (r AccountUrlscannerScanNewParamsScreenshotsResolution) IsKnown() bool {
	switch r {
	case AccountUrlscannerScanNewParamsScreenshotsResolutionDesktop, AccountUrlscannerScanNewParamsScreenshotsResolutionMobile, AccountUrlscannerScanNewParamsScreenshotsResolutionTablet:
		return true
	}
	return false
}

// The option `Public` means it will be included in listings like recent scans and
// search results. `Unlisted` means it will not be included in the aforementioned
// listings, users will need to have the scan's ID to access it. A a scan will be
// automatically marked as unlisted if it fails, if it contains potential PII or
// other sensitive material.
type AccountUrlscannerScanNewParamsVisibility string

const (
	AccountUrlscannerScanNewParamsVisibilityPublic   AccountUrlscannerScanNewParamsVisibility = "Public"
	AccountUrlscannerScanNewParamsVisibilityUnlisted AccountUrlscannerScanNewParamsVisibility = "Unlisted"
)

func (r AccountUrlscannerScanNewParamsVisibility) IsKnown() bool {
	switch r {
	case AccountUrlscannerScanNewParamsVisibilityPublic, AccountUrlscannerScanNewParamsVisibilityUnlisted:
		return true
	}
	return false
}

type AccountUrlscannerScanGetParams struct {
	// Whether to return full report (scan summary and network log).
	Full param.Field[bool] `query:"full"`
}

// URLQuery serializes [AccountUrlscannerScanGetParams]'s query parameters as
// `url.Values`.
func (r AccountUrlscannerScanGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountUrlscannerScanListParams struct {
	// Return only scans created by account.
	AccountScans param.Field[bool] `query:"account_scans"`
	// Filter scans requested before date (inclusive).
	DateEnd param.Field[time.Time] `query:"date_end" format:"date-time"`
	// Filter scans requested after date (inclusive).
	DateStart param.Field[time.Time] `query:"date_start" format:"date-time"`
	// Filter scans by hash of any html/js/css request made by the webpage.
	Hash param.Field[string] `query:"hash"`
	// Filter scans by hostname of _any_ request made by the webpage.
	Hostname param.Field[string] `query:"hostname"`
	// Filter scans by IP address (IPv4 or IPv6) of _any_ request made by the webpage.
	IP param.Field[string] `query:"ip"`
	// Filter scans by malicious verdict.
	IsMalicious param.Field[bool] `query:"is_malicious"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Pagination cursor to get the next set of results.
	NextCursor param.Field[string] `query:"next_cursor"`
	// Filter scans by main page Autonomous System Number (ASN).
	PageAsn param.Field[string] `query:"page_asn"`
	// Filter scans by main page hostname (domain of effective URL).
	PageHostname param.Field[string] `query:"page_hostname"`
	// Filter scans by main page IP address (IPv4 or IPv6).
	PageIP param.Field[string] `query:"page_ip"`
	// Filter scans by exact match of effective URL path (also supports suffix search).
	PagePath param.Field[string] `query:"page_path"`
	// Filter scans by submitted or scanned URL
	PageURL param.Field[string] `query:"page_url"`
	// Filter scans by url path of _any_ request made by the webpage.
	Path param.Field[string] `query:"path"`
	// Scan UUID.
	ScanID param.Field[string] `query:"scan_id" format:"uuid"`
	// Filter scans by URL of _any_ request made by the webpage
	URL param.Field[string] `query:"url"`
}

// URLQuery serializes [AccountUrlscannerScanListParams]'s query parameters as
// `url.Values`.
func (r AccountUrlscannerScanListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountUrlscannerScanGetScreenshotParams struct {
	// Target device type.
	Resolution param.Field[AccountUrlscannerScanGetScreenshotParamsResolution] `query:"resolution"`
}

// URLQuery serializes [AccountUrlscannerScanGetScreenshotParams]'s query
// parameters as `url.Values`.
func (r AccountUrlscannerScanGetScreenshotParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Target device type.
type AccountUrlscannerScanGetScreenshotParamsResolution string

const (
	AccountUrlscannerScanGetScreenshotParamsResolutionDesktop AccountUrlscannerScanGetScreenshotParamsResolution = "desktop"
	AccountUrlscannerScanGetScreenshotParamsResolutionMobile  AccountUrlscannerScanGetScreenshotParamsResolution = "mobile"
	AccountUrlscannerScanGetScreenshotParamsResolutionTablet  AccountUrlscannerScanGetScreenshotParamsResolution = "tablet"
)

func (r AccountUrlscannerScanGetScreenshotParamsResolution) IsKnown() bool {
	switch r {
	case AccountUrlscannerScanGetScreenshotParamsResolutionDesktop, AccountUrlscannerScanGetScreenshotParamsResolutionMobile, AccountUrlscannerScanGetScreenshotParamsResolutionTablet:
		return true
	}
	return false
}
