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
func (r *AccountUrlscannerScanService) New(ctx context.Context, accountID string, body AccountUrlscannerScanNewParams, opts ...option.RequestOption) (res *AccountUrlscannerScanNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/urlscanner/scan", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get URL scan by uuid
func (r *AccountUrlscannerScanService) Get(ctx context.Context, accountID string, scanID string, query AccountUrlscannerScanGetParams, opts ...option.RequestOption) (res *AccountUrlscannerScanGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if scanID == "" {
		err = errors.New("missing required scan_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/urlscanner/scan/%s", accountID, scanID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Search scans by date and webpages' requests, including full URL (after
// redirects), hostname, and path. <br/> A successful scan will appear in search
// results a few minutes after finishing but may take much longer if the system in
// under load. By default, only successfully completed scans will appear in search
// results, unless searching by `scanId`. Please take into account that older scans
// may be removed from the search index at an unspecified time.
func (r *AccountUrlscannerScanService) List(ctx context.Context, accountID string, query AccountUrlscannerScanListParams, opts ...option.RequestOption) (res *AccountUrlscannerScanListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/urlscanner/scan", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Get a URL scan's HAR file. See HAR spec at
// http://www.softwareishard.com/blog/har-12-spec/.
func (r *AccountUrlscannerScanService) GetHar(ctx context.Context, accountID string, scanID string, opts ...option.RequestOption) (res *AccountUrlscannerScanGetHarResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if scanID == "" {
		err = errors.New("missing required scan_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/urlscanner/scan/%s/har", accountID, scanID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get scan's screenshot by resolution (desktop/mobile/tablet).
func (r *AccountUrlscannerScanService) GetScreenshot(ctx context.Context, accountID string, scanID string, query AccountUrlscannerScanGetScreenshotParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "image/png")}, opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if scanID == "" {
		err = errors.New("missing required scan_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/urlscanner/scan/%s/screenshot", accountID, scanID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type AccountUrlscannerScanNewResponse struct {
	Errors   []AccountUrlscannerScanNewResponseError   `json:"errors,required"`
	Messages []AccountUrlscannerScanNewResponseMessage `json:"messages,required"`
	Result   AccountUrlscannerScanNewResponseResult    `json:"result,required"`
	Success  bool                                      `json:"success,required"`
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
	Message string                                    `json:"message,required"`
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
	Message string                                      `json:"message,required"`
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
	Time time.Time `json:"time,required" format:"date-time"`
	// Canonical form of submitted URL. Use this if you want to later search by URL.
	URL string `json:"url,required"`
	// Scan ID.
	Uuid string `json:"uuid,required" format:"uuid"`
	// Submitted visibility status.
	Visibility string                                     `json:"visibility,required"`
	JSON       accountUrlscannerScanNewResponseResultJSON `json:"-"`
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

type AccountUrlscannerScanGetResponse struct {
	Errors   []AccountUrlscannerScanGetResponseError   `json:"errors,required"`
	Messages []AccountUrlscannerScanGetResponseMessage `json:"messages,required"`
	Result   AccountUrlscannerScanGetResponseResult    `json:"result,required"`
	// Whether request was successful or not
	Success bool                                 `json:"success,required"`
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
	Message string                                    `json:"message,required"`
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
	Message string                                      `json:"message,required"`
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
	Scan AccountUrlscannerScanGetResponseResultScan `json:"scan,required"`
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
	Certificates []AccountUrlscannerScanGetResponseResultScanCertificate `json:"certificates,required"`
	Geo          AccountUrlscannerScanGetResponseResultScanGeo           `json:"geo,required"`
	Meta         AccountUrlscannerScanGetResponseResultScanMeta          `json:"meta,required"`
	Page         AccountUrlscannerScanGetResponseResultScanPage          `json:"page,required"`
	Performance  []AccountUrlscannerScanGetResponseResultScanPerformance `json:"performance,required"`
	Task         AccountUrlscannerScanGetResponseResultScanTask          `json:"task,required"`
	Verdicts     AccountUrlscannerScanGetResponseResultScanVerdicts      `json:"verdicts,required"`
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
	Issuer      string                                                    `json:"issuer,required"`
	SubjectName string                                                    `json:"subjectName,required"`
	ValidFrom   float64                                                   `json:"validFrom,required"`
	ValidTo     float64                                                   `json:"validTo,required"`
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
	// GeoIP continent location
	Continents []string `json:"continents,required"`
	// GeoIP country location
	Locations []string                                          `json:"locations,required"`
	JSON      accountUrlscannerScanGetResponseResultScanGeoJSON `json:"-"`
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
	Processors AccountUrlscannerScanGetResponseResultScanMetaProcessors `json:"processors,required"`
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
	Categories AccountUrlscannerScanGetResponseResultScanMetaProcessorsCategories `json:"categories,required"`
	Phishing   []string                                                           `json:"phishing,required"`
	Rank       AccountUrlscannerScanGetResponseResultScanMetaProcessorsRank       `json:"rank,required"`
	Tech       []AccountUrlscannerScanGetResponseResultScanMetaProcessorsTech     `json:"tech,required"`
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
	Content []AccountUrlscannerScanGetResponseResultScanMetaProcessorsCategoriesContent `json:"content,required"`
	Risks   []AccountUrlscannerScanGetResponseResultScanMetaProcessorsCategoriesRisk    `json:"risks,required"`
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
	ID              int64                                                                         `json:"id,required"`
	Name            string                                                                        `json:"name,required"`
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
	ID              int64                                                                      `json:"id,required"`
	Name            string                                                                     `json:"name,required"`
	SuperCategoryID int64                                                                      `json:"super_category_id,required"`
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
	Bucket string `json:"bucket,required"`
	Name   string `json:"name,required"`
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
	Categories  []AccountUrlscannerScanGetResponseResultScanMetaProcessorsTechCategory `json:"categories,required"`
	Confidence  int64                                                                  `json:"confidence,required"`
	Evidence    AccountUrlscannerScanGetResponseResultScanMetaProcessorsTechEvidence   `json:"evidence,required"`
	Icon        string                                                                 `json:"icon,required"`
	Name        string                                                                 `json:"name,required"`
	Slug        string                                                                 `json:"slug,required"`
	Website     string                                                                 `json:"website,required"`
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
	ID       int64                                                                    `json:"id,required"`
	Groups   []int64                                                                  `json:"groups,required"`
	Name     string                                                                   `json:"name,required"`
	Priority int64                                                                    `json:"priority,required"`
	Slug     string                                                                   `json:"slug,required"`
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
	ImpliedBy []string                                                                      `json:"impliedBy,required"`
	Patterns  []AccountUrlscannerScanGetResponseResultScanMetaProcessorsTechEvidencePattern `json:"patterns,required"`
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
	Confidence int64    `json:"confidence,required"`
	Excludes   []string `json:"excludes,required"`
	Implies    []string `json:"implies,required"`
	Match      string   `json:"match,required"`
	// Header or Cookie name when set
	Name    string                                                                          `json:"name,required"`
	Regex   string                                                                          `json:"regex,required"`
	Type    string                                                                          `json:"type,required"`
	Value   string                                                                          `json:"value,required"`
	Version string                                                                          `json:"version,required"`
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
	Asn                   string                                                            `json:"asn,required"`
	AsnLocationAlpha2     string                                                            `json:"asnLocationAlpha2,required"`
	Asnname               string                                                            `json:"asnname,required"`
	Console               []AccountUrlscannerScanGetResponseResultScanPageConsole           `json:"console,required"`
	Cookies               []AccountUrlscannerScanGetResponseResultScanPageCookie            `json:"cookies,required"`
	Country               string                                                            `json:"country,required"`
	CountryLocationAlpha2 string                                                            `json:"countryLocationAlpha2,required"`
	Domain                string                                                            `json:"domain,required"`
	Headers               []AccountUrlscannerScanGetResponseResultScanPageHeader            `json:"headers,required"`
	IP                    string                                                            `json:"ip,required"`
	Js                    AccountUrlscannerScanGetResponseResultScanPageJs                  `json:"js,required"`
	SecurityViolations    []AccountUrlscannerScanGetResponseResultScanPageSecurityViolation `json:"securityViolations,required"`
	Status                float64                                                           `json:"status,required"`
	Subdivision1Name      string                                                            `json:"subdivision1Name,required"`
	Subdivision2name      string                                                            `json:"subdivision2name,required"`
	URL                   string                                                            `json:"url,required"`
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
	Category string                                                    `json:"category,required"`
	Text     string                                                    `json:"text,required"`
	Type     string                                                    `json:"type,required"`
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
	Domain       string                                                   `json:"domain,required"`
	Expires      float64                                                  `json:"expires,required"`
	HTTPOnly     bool                                                     `json:"httpOnly,required"`
	Name         string                                                   `json:"name,required"`
	Path         string                                                   `json:"path,required"`
	SameParty    bool                                                     `json:"sameParty,required"`
	Secure       bool                                                     `json:"secure,required"`
	Session      bool                                                     `json:"session,required"`
	Size         float64                                                  `json:"size,required"`
	SourcePort   float64                                                  `json:"sourcePort,required"`
	SourceScheme string                                                   `json:"sourceScheme,required"`
	Value        string                                                   `json:"value,required"`
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
	Name  string                                                   `json:"name,required"`
	Value string                                                   `json:"value,required"`
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
	Variables []AccountUrlscannerScanGetResponseResultScanPageJsVariable `json:"variables,required"`
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
	Name string                                                       `json:"name,required"`
	Type string                                                       `json:"type,required"`
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
	Category string                                                              `json:"category,required"`
	Text     string                                                              `json:"text,required"`
	URL      string                                                              `json:"url,required"`
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
	ConnectEnd                 float64                                                   `json:"connectEnd,required"`
	ConnectStart               float64                                                   `json:"connectStart,required"`
	DecodedBodySize            float64                                                   `json:"decodedBodySize,required"`
	DomainLookupEnd            float64                                                   `json:"domainLookupEnd,required"`
	DomainLookupStart          float64                                                   `json:"domainLookupStart,required"`
	DomComplete                float64                                                   `json:"domComplete,required"`
	DomContentLoadedEventEnd   float64                                                   `json:"domContentLoadedEventEnd,required"`
	DomContentLoadedEventStart float64                                                   `json:"domContentLoadedEventStart,required"`
	DomInteractive             float64                                                   `json:"domInteractive,required"`
	Duration                   float64                                                   `json:"duration,required"`
	EncodedBodySize            float64                                                   `json:"encodedBodySize,required"`
	EntryType                  string                                                    `json:"entryType,required"`
	FetchStart                 float64                                                   `json:"fetchStart,required"`
	InitiatorType              string                                                    `json:"initiatorType,required"`
	LoadEventEnd               float64                                                   `json:"loadEventEnd,required"`
	LoadEventStart             float64                                                   `json:"loadEventStart,required"`
	Name                       string                                                    `json:"name,required"`
	NextHopProtocol            string                                                    `json:"nextHopProtocol,required"`
	RedirectCount              float64                                                   `json:"redirectCount,required"`
	RedirectEnd                float64                                                   `json:"redirectEnd,required"`
	RedirectStart              float64                                                   `json:"redirectStart,required"`
	RequestStart               float64                                                   `json:"requestStart,required"`
	ResponseEnd                float64                                                   `json:"responseEnd,required"`
	ResponseStart              float64                                                   `json:"responseStart,required"`
	SecureConnectionStart      float64                                                   `json:"secureConnectionStart,required"`
	StartTime                  float64                                                   `json:"startTime,required"`
	TransferSize               float64                                                   `json:"transferSize,required"`
	Type                       string                                                    `json:"type,required"`
	UnloadEventEnd             float64                                                   `json:"unloadEventEnd,required"`
	UnloadEventStart           float64                                                   `json:"unloadEventStart,required"`
	WorkerStart                float64                                                   `json:"workerStart,required"`
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
	ClientLocation string                                                   `json:"clientLocation,required"`
	ClientType     AccountUrlscannerScanGetResponseResultScanTaskClientType `json:"clientType,required"`
	// URL of the primary request, after all HTTP redirects
	EffectiveURL string                                                    `json:"effectiveUrl,required"`
	Errors       []AccountUrlscannerScanGetResponseResultScanTaskError     `json:"errors,required"`
	ScannedFrom  AccountUrlscannerScanGetResponseResultScanTaskScannedFrom `json:"scannedFrom,required"`
	Status       AccountUrlscannerScanGetResponseResultScanTaskStatus      `json:"status,required"`
	Success      bool                                                      `json:"success,required"`
	Time         string                                                    `json:"time,required"`
	TimeEnd      string                                                    `json:"timeEnd,required"`
	// Submitted URL
	URL string `json:"url,required"`
	// Scan ID
	Uuid       string                                                   `json:"uuid,required"`
	Visibility AccountUrlscannerScanGetResponseResultScanTaskVisibility `json:"visibility,required"`
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
	Message string                                                  `json:"message,required"`
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
	Colo string                                                        `json:"colo,required"`
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
	Overall AccountUrlscannerScanGetResponseResultScanVerdictsOverall `json:"overall,required"`
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
	Categories []AccountUrlscannerScanGetResponseResultScanVerdictsOverallCategory `json:"categories,required"`
	// At least one of our subsystems marked the site as potentially malicious at the
	// time of the scan.
	Malicious bool                                                          `json:"malicious,required"`
	Phishing  []string                                                      `json:"phishing,required"`
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
	ID              float64                                                               `json:"id,required"`
	Name            string                                                                `json:"name,required"`
	SuperCategoryID float64                                                               `json:"super_category_id,required"`
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
	Asn            string                                                `json:"asn,required"`
	Description    string                                                `json:"description,required"`
	LocationAlpha2 string                                                `json:"location_alpha2,required"`
	Name           string                                                `json:"name,required"`
	OrgName        string                                                `json:"org_name,required"`
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
	Categories AccountUrlscannerScanGetResponseResultScanDomainsExampleComCategories `json:"categories,required"`
	DNS        []AccountUrlscannerScanGetResponseResultScanDomainsExampleComDNS      `json:"dns,required"`
	Name       string                                                                `json:"name,required"`
	Rank       AccountUrlscannerScanGetResponseResultScanDomainsExampleComRank       `json:"rank,required"`
	Type       string                                                                `json:"type,required"`
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
	Inherited AccountUrlscannerScanGetResponseResultScanDomainsExampleComCategoriesInherited `json:"inherited,required"`
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
	ID              int64                                                                                     `json:"id,required"`
	Name            string                                                                                    `json:"name,required"`
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
	ID              int64                                                                                  `json:"id,required"`
	Name            string                                                                                 `json:"name,required"`
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
	ID              int64                                                                            `json:"id,required"`
	Name            string                                                                           `json:"name,required"`
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
	ID              int64                                                                         `json:"id,required"`
	Name            string                                                                        `json:"name,required"`
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
	Address     string                                                             `json:"address,required"`
	DnssecValid bool                                                               `json:"dnssec_valid,required"`
	Name        string                                                             `json:"name,required"`
	Type        string                                                             `json:"type,required"`
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
	Bucket string `json:"bucket,required"`
	Name   string `json:"name,required"`
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
	Asn               string                                              `json:"asn,required"`
	AsnDescription    string                                              `json:"asnDescription,required"`
	AsnLocationAlpha2 string                                              `json:"asnLocationAlpha2,required"`
	AsnName           string                                              `json:"asnName,required"`
	AsnOrgName        string                                              `json:"asnOrgName,required"`
	Continent         string                                              `json:"continent,required"`
	GeonameID         string                                              `json:"geonameId,required"`
	IP                string                                              `json:"ip,required"`
	IPVersion         string                                              `json:"ipVersion,required"`
	Latitude          string                                              `json:"latitude,required"`
	LocationAlpha2    string                                              `json:"locationAlpha2,required"`
	LocationName      string                                              `json:"locationName,required"`
	Longitude         string                                              `json:"longitude,required"`
	Subdivision1Name  string                                              `json:"subdivision1Name,required"`
	Subdivision2Name  string                                              `json:"subdivision2Name,required"`
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
	Href string                                                  `json:"href,required"`
	Text string                                                  `json:"text,required"`
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
	Errors   []AccountUrlscannerScanListResponseError   `json:"errors,required"`
	Messages []AccountUrlscannerScanListResponseMessage `json:"messages,required"`
	Result   AccountUrlscannerScanListResponseResult    `json:"result,required"`
	// Whether search request was successful or not
	Success bool                                  `json:"success,required"`
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
	Message string                                     `json:"message,required"`
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
	Message string                                       `json:"message,required"`
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
	Tasks []AccountUrlscannerScanListResponseResultTask `json:"tasks,required"`
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
	Country string `json:"country,required"`
	// Whether scan was successful or not
	Success bool `json:"success,required"`
	// When scan was submitted (UTC)
	Time time.Time `json:"time,required" format:"date-time"`
	// Scan url (after redirects)
	URL string `json:"url,required"`
	// Scan id
	Uuid string `json:"uuid,required" format:"uuid"`
	// Visibility status.
	Visibility string                                          `json:"visibility,required"`
	JSON       accountUrlscannerScanListResponseResultTaskJSON `json:"-"`
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

type AccountUrlscannerScanGetHarResponse struct {
	Errors   []AccountUrlscannerScanGetHarResponseError   `json:"errors,required"`
	Messages []AccountUrlscannerScanGetHarResponseMessage `json:"messages,required"`
	Result   AccountUrlscannerScanGetHarResponseResult    `json:"result,required"`
	// Whether search request was successful or not
	Success bool                                    `json:"success,required"`
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
	Message string                                       `json:"message,required"`
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
	Message string                                         `json:"message,required"`
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
	Har  AccountUrlscannerScanGetHarResponseResultHar  `json:"har,required"`
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
	Log  AccountUrlscannerScanGetHarResponseResultHarLog  `json:"log,required"`
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
	Creator AccountUrlscannerScanGetHarResponseResultHarLogCreator `json:"creator,required"`
	Entries []AccountUrlscannerScanGetHarResponseResultHarLogEntry `json:"entries,required"`
	Pages   []AccountUrlscannerScanGetHarResponseResultHarLogPage  `json:"pages,required"`
	Version string                                                 `json:"version,required"`
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
	Comment string                                                     `json:"comment,required"`
	Name    string                                                     `json:"name,required"`
	Version string                                                     `json:"version,required"`
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
	InitialPriority string                                                         `json:"_initialPriority,required"`
	InitiatorType   string                                                         `json:"_initiator_type,required"`
	Priority        string                                                         `json:"_priority,required"`
	RequestID       string                                                         `json:"_requestId,required"`
	RequestTime     float64                                                        `json:"_requestTime,required"`
	ResourceType    string                                                         `json:"_resourceType,required"`
	Cache           interface{}                                                    `json:"cache,required"`
	Connection      string                                                         `json:"connection,required"`
	Pageref         string                                                         `json:"pageref,required"`
	Request         AccountUrlscannerScanGetHarResponseResultHarLogEntriesRequest  `json:"request,required"`
	Response        AccountUrlscannerScanGetHarResponseResultHarLogEntriesResponse `json:"response,required"`
	ServerIPAddress string                                                         `json:"serverIPAddress,required"`
	StartedDateTime string                                                         `json:"startedDateTime,required"`
	Time            float64                                                        `json:"time,required"`
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
	BodySize    float64                                                               `json:"bodySize,required"`
	Headers     []AccountUrlscannerScanGetHarResponseResultHarLogEntriesRequestHeader `json:"headers,required"`
	HeadersSize float64                                                               `json:"headersSize,required"`
	HTTPVersion string                                                                `json:"httpVersion,required"`
	Method      string                                                                `json:"method,required"`
	URL         string                                                                `json:"url,required"`
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
	Name  string                                                                  `json:"name,required"`
	Value string                                                                  `json:"value,required"`
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
	TransferSize float64                                                                `json:"_transferSize,required"`
	BodySize     float64                                                                `json:"bodySize,required"`
	Content      AccountUrlscannerScanGetHarResponseResultHarLogEntriesResponseContent  `json:"content,required"`
	Headers      []AccountUrlscannerScanGetHarResponseResultHarLogEntriesResponseHeader `json:"headers,required"`
	HeadersSize  float64                                                                `json:"headersSize,required"`
	HTTPVersion  string                                                                 `json:"httpVersion,required"`
	RedirectURL  string                                                                 `json:"redirectURL,required"`
	Status       float64                                                                `json:"status,required"`
	StatusText   string                                                                 `json:"statusText,required"`
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
	MimeType    string                                                                    `json:"mimeType,required"`
	Size        float64                                                                   `json:"size,required"`
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
	Name  string                                                                   `json:"name,required"`
	Value string                                                                   `json:"value,required"`
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
	ID              string                                                          `json:"id,required"`
	PageTimings     AccountUrlscannerScanGetHarResponseResultHarLogPagesPageTimings `json:"pageTimings,required"`
	StartedDateTime string                                                          `json:"startedDateTime,required"`
	Title           string                                                          `json:"title,required"`
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
	OnContentLoad float64                                                             `json:"onContentLoad,required"`
	OnLoad        float64                                                             `json:"onLoad,required"`
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
	URL param.Field[string] `json:"url,required"`
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
