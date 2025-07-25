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

// ZonePageShieldScriptService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZonePageShieldScriptService] method instead.
type ZonePageShieldScriptService struct {
	Options []option.RequestOption
}

// NewZonePageShieldScriptService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZonePageShieldScriptService(opts ...option.RequestOption) (r *ZonePageShieldScriptService) {
	r = &ZonePageShieldScriptService{}
	r.Options = opts
	return
}

// Fetches a script detected by Page Shield by script ID.
func (r *ZonePageShieldScriptService) Get(ctx context.Context, zoneID string, scriptID string, opts ...option.RequestOption) (res *ZonePageShieldScriptGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if scriptID == "" {
		err = errors.New("missing required script_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/page_shield/scripts/%s", zoneID, scriptID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Lists all scripts detected by Page Shield.
func (r *ZonePageShieldScriptService) List(ctx context.Context, zoneID string, query ZonePageShieldScriptListParams, opts ...option.RequestOption) (res *ZonePageShieldScriptListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/page_shield/scripts", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type Script struct {
	// Identifier
	ID                    string    `json:"id,required"`
	AddedAt               time.Time `json:"added_at,required" format:"date-time"`
	FirstSeenAt           time.Time `json:"first_seen_at,required" format:"date-time"`
	Host                  string    `json:"host,required"`
	LastSeenAt            time.Time `json:"last_seen_at,required" format:"date-time"`
	URL                   string    `json:"url,required"`
	URLContainsCdnCgiPath bool      `json:"url_contains_cdn_cgi_path,required"`
	// The cryptomining score of the JavaScript content.
	CryptominingScore int64 `json:"cryptomining_score,nullable"`
	// The dataflow score of the JavaScript content.
	DataflowScore           int64 `json:"dataflow_score,nullable"`
	DomainReportedMalicious bool  `json:"domain_reported_malicious"`
	// The timestamp of when the script was last fetched.
	FetchedAt    string `json:"fetched_at,nullable"`
	FirstPageURL string `json:"first_page_url"`
	// The computed hash of the analyzed script.
	Hash string `json:"hash,nullable"`
	// The integrity score of the JavaScript content.
	JsIntegrityScore int64 `json:"js_integrity_score,nullable"`
	// The magecart score of the JavaScript content.
	MagecartScore             int64    `json:"magecart_score,nullable"`
	MaliciousDomainCategories []string `json:"malicious_domain_categories"`
	MaliciousURLCategories    []string `json:"malicious_url_categories"`
	// The malware score of the JavaScript content.
	MalwareScore int64 `json:"malware_score,nullable"`
	// The obfuscation score of the JavaScript content.
	ObfuscationScore     int64      `json:"obfuscation_score,nullable"`
	PageURLs             []string   `json:"page_urls"`
	URLReportedMalicious bool       `json:"url_reported_malicious"`
	JSON                 scriptJSON `json:"-"`
}

// scriptJSON contains the JSON metadata for the struct [Script]
type scriptJSON struct {
	ID                        apijson.Field
	AddedAt                   apijson.Field
	FirstSeenAt               apijson.Field
	Host                      apijson.Field
	LastSeenAt                apijson.Field
	URL                       apijson.Field
	URLContainsCdnCgiPath     apijson.Field
	CryptominingScore         apijson.Field
	DataflowScore             apijson.Field
	DomainReportedMalicious   apijson.Field
	FetchedAt                 apijson.Field
	FirstPageURL              apijson.Field
	Hash                      apijson.Field
	JsIntegrityScore          apijson.Field
	MagecartScore             apijson.Field
	MaliciousDomainCategories apijson.Field
	MaliciousURLCategories    apijson.Field
	MalwareScore              apijson.Field
	ObfuscationScore          apijson.Field
	PageURLs                  apijson.Field
	URLReportedMalicious      apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *Script) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scriptJSON) RawJSON() string {
	return r.raw
}

type ZonePageShieldScriptGetResponse struct {
	Result ZonePageShieldScriptGetResponseResult `json:"result,required,nullable"`
	// Whether the API call was successful
	Success  ZonePageShieldScriptGetResponseSuccess   `json:"success,required"`
	Errors   []ZonePageShieldScriptGetResponseError   `json:"errors"`
	Messages []ZonePageShieldScriptGetResponseMessage `json:"messages"`
	JSON     zonePageShieldScriptGetResponseJSON      `json:"-"`
}

// zonePageShieldScriptGetResponseJSON contains the JSON metadata for the struct
// [ZonePageShieldScriptGetResponse]
type zonePageShieldScriptGetResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	Errors      apijson.Field
	Messages    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonePageShieldScriptGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zonePageShieldScriptGetResponseJSON) RawJSON() string {
	return r.raw
}

type ZonePageShieldScriptGetResponseResult struct {
	// Identifier
	ID                    string    `json:"id,required"`
	AddedAt               time.Time `json:"added_at,required" format:"date-time"`
	FirstSeenAt           time.Time `json:"first_seen_at,required" format:"date-time"`
	Host                  string    `json:"host,required"`
	LastSeenAt            time.Time `json:"last_seen_at,required" format:"date-time"`
	URL                   string    `json:"url,required"`
	URLContainsCdnCgiPath bool      `json:"url_contains_cdn_cgi_path,required"`
	// The cryptomining score of the JavaScript content.
	CryptominingScore int64 `json:"cryptomining_score,nullable"`
	// The dataflow score of the JavaScript content.
	DataflowScore           int64 `json:"dataflow_score,nullable"`
	DomainReportedMalicious bool  `json:"domain_reported_malicious"`
	// The timestamp of when the script was last fetched.
	FetchedAt    string `json:"fetched_at,nullable"`
	FirstPageURL string `json:"first_page_url"`
	// The computed hash of the analyzed script.
	Hash string `json:"hash,nullable"`
	// The integrity score of the JavaScript content.
	JsIntegrityScore int64 `json:"js_integrity_score,nullable"`
	// The magecart score of the JavaScript content.
	MagecartScore             int64    `json:"magecart_score,nullable"`
	MaliciousDomainCategories []string `json:"malicious_domain_categories"`
	MaliciousURLCategories    []string `json:"malicious_url_categories"`
	// The malware score of the JavaScript content.
	MalwareScore int64 `json:"malware_score,nullable"`
	// The obfuscation score of the JavaScript content.
	ObfuscationScore     int64                                          `json:"obfuscation_score,nullable"`
	PageURLs             []string                                       `json:"page_urls"`
	URLReportedMalicious bool                                           `json:"url_reported_malicious"`
	Versions             []ZonePageShieldScriptGetResponseResultVersion `json:"versions,nullable"`
	JSON                 zonePageShieldScriptGetResponseResultJSON      `json:"-"`
}

// zonePageShieldScriptGetResponseResultJSON contains the JSON metadata for the
// struct [ZonePageShieldScriptGetResponseResult]
type zonePageShieldScriptGetResponseResultJSON struct {
	ID                        apijson.Field
	AddedAt                   apijson.Field
	FirstSeenAt               apijson.Field
	Host                      apijson.Field
	LastSeenAt                apijson.Field
	URL                       apijson.Field
	URLContainsCdnCgiPath     apijson.Field
	CryptominingScore         apijson.Field
	DataflowScore             apijson.Field
	DomainReportedMalicious   apijson.Field
	FetchedAt                 apijson.Field
	FirstPageURL              apijson.Field
	Hash                      apijson.Field
	JsIntegrityScore          apijson.Field
	MagecartScore             apijson.Field
	MaliciousDomainCategories apijson.Field
	MaliciousURLCategories    apijson.Field
	MalwareScore              apijson.Field
	ObfuscationScore          apijson.Field
	PageURLs                  apijson.Field
	URLReportedMalicious      apijson.Field
	Versions                  apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *ZonePageShieldScriptGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zonePageShieldScriptGetResponseResultJSON) RawJSON() string {
	return r.raw
}

// The version of the analyzed script.
type ZonePageShieldScriptGetResponseResultVersion struct {
	// The cryptomining score of the JavaScript content.
	CryptominingScore int64 `json:"cryptomining_score,nullable"`
	// The dataflow score of the JavaScript content.
	DataflowScore int64 `json:"dataflow_score,nullable"`
	// The timestamp of when the script was last fetched.
	FetchedAt string `json:"fetched_at,nullable"`
	// The computed hash of the analyzed script.
	Hash string `json:"hash,nullable"`
	// The integrity score of the JavaScript content.
	JsIntegrityScore int64 `json:"js_integrity_score,nullable"`
	// The magecart score of the JavaScript content.
	MagecartScore int64 `json:"magecart_score,nullable"`
	// The malware score of the JavaScript content.
	MalwareScore int64 `json:"malware_score,nullable"`
	// The obfuscation score of the JavaScript content.
	ObfuscationScore int64                                            `json:"obfuscation_score,nullable"`
	JSON             zonePageShieldScriptGetResponseResultVersionJSON `json:"-"`
}

// zonePageShieldScriptGetResponseResultVersionJSON contains the JSON metadata for
// the struct [ZonePageShieldScriptGetResponseResultVersion]
type zonePageShieldScriptGetResponseResultVersionJSON struct {
	CryptominingScore apijson.Field
	DataflowScore     apijson.Field
	FetchedAt         apijson.Field
	Hash              apijson.Field
	JsIntegrityScore  apijson.Field
	MagecartScore     apijson.Field
	MalwareScore      apijson.Field
	ObfuscationScore  apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ZonePageShieldScriptGetResponseResultVersion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zonePageShieldScriptGetResponseResultVersionJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ZonePageShieldScriptGetResponseSuccess bool

const (
	ZonePageShieldScriptGetResponseSuccessTrue ZonePageShieldScriptGetResponseSuccess = true
)

func (r ZonePageShieldScriptGetResponseSuccess) IsKnown() bool {
	switch r {
	case ZonePageShieldScriptGetResponseSuccessTrue:
		return true
	}
	return false
}

type ZonePageShieldScriptGetResponseError struct {
	Code             int64                                       `json:"code,required"`
	Message          string                                      `json:"message,required"`
	DocumentationURL string                                      `json:"documentation_url"`
	Source           ZonePageShieldScriptGetResponseErrorsSource `json:"source"`
	JSON             zonePageShieldScriptGetResponseErrorJSON    `json:"-"`
}

// zonePageShieldScriptGetResponseErrorJSON contains the JSON metadata for the
// struct [ZonePageShieldScriptGetResponseError]
type zonePageShieldScriptGetResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZonePageShieldScriptGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zonePageShieldScriptGetResponseErrorJSON) RawJSON() string {
	return r.raw
}

type ZonePageShieldScriptGetResponseErrorsSource struct {
	Pointer string                                          `json:"pointer"`
	JSON    zonePageShieldScriptGetResponseErrorsSourceJSON `json:"-"`
}

// zonePageShieldScriptGetResponseErrorsSourceJSON contains the JSON metadata for
// the struct [ZonePageShieldScriptGetResponseErrorsSource]
type zonePageShieldScriptGetResponseErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonePageShieldScriptGetResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zonePageShieldScriptGetResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type ZonePageShieldScriptGetResponseMessage struct {
	Code             int64                                         `json:"code,required"`
	Message          string                                        `json:"message,required"`
	DocumentationURL string                                        `json:"documentation_url"`
	Source           ZonePageShieldScriptGetResponseMessagesSource `json:"source"`
	JSON             zonePageShieldScriptGetResponseMessageJSON    `json:"-"`
}

// zonePageShieldScriptGetResponseMessageJSON contains the JSON metadata for the
// struct [ZonePageShieldScriptGetResponseMessage]
type zonePageShieldScriptGetResponseMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZonePageShieldScriptGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zonePageShieldScriptGetResponseMessageJSON) RawJSON() string {
	return r.raw
}

type ZonePageShieldScriptGetResponseMessagesSource struct {
	Pointer string                                            `json:"pointer"`
	JSON    zonePageShieldScriptGetResponseMessagesSourceJSON `json:"-"`
}

// zonePageShieldScriptGetResponseMessagesSourceJSON contains the JSON metadata for
// the struct [ZonePageShieldScriptGetResponseMessagesSource]
type zonePageShieldScriptGetResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonePageShieldScriptGetResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zonePageShieldScriptGetResponseMessagesSourceJSON) RawJSON() string {
	return r.raw
}

type ZonePageShieldScriptListResponse struct {
	Result     []Script                                   `json:"result,required"`
	ResultInfo ZonePageShieldScriptListResponseResultInfo `json:"result_info,required"`
	// Whether the API call was successful
	Success  ZonePageShieldScriptListResponseSuccess   `json:"success,required"`
	Errors   []ZonePageShieldScriptListResponseError   `json:"errors"`
	Messages []ZonePageShieldScriptListResponseMessage `json:"messages"`
	JSON     zonePageShieldScriptListResponseJSON      `json:"-"`
}

// zonePageShieldScriptListResponseJSON contains the JSON metadata for the struct
// [ZonePageShieldScriptListResponse]
type zonePageShieldScriptListResponseJSON struct {
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	Errors      apijson.Field
	Messages    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonePageShieldScriptListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zonePageShieldScriptListResponseJSON) RawJSON() string {
	return r.raw
}

type ZonePageShieldScriptListResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count,required"`
	// Current page within paginated list of results
	Page float64 `json:"page,required"`
	// Number of results per page of results
	PerPage float64 `json:"per_page,required"`
	// Total results available without any search parameters
	TotalCount float64 `json:"total_count,required"`
	// Total number of pages
	TotalPages float64                                        `json:"total_pages,required"`
	JSON       zonePageShieldScriptListResponseResultInfoJSON `json:"-"`
}

// zonePageShieldScriptListResponseResultInfoJSON contains the JSON metadata for
// the struct [ZonePageShieldScriptListResponseResultInfo]
type zonePageShieldScriptListResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	TotalPages  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonePageShieldScriptListResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zonePageShieldScriptListResponseResultInfoJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ZonePageShieldScriptListResponseSuccess bool

const (
	ZonePageShieldScriptListResponseSuccessTrue ZonePageShieldScriptListResponseSuccess = true
)

func (r ZonePageShieldScriptListResponseSuccess) IsKnown() bool {
	switch r {
	case ZonePageShieldScriptListResponseSuccessTrue:
		return true
	}
	return false
}

type ZonePageShieldScriptListResponseError struct {
	Code             int64                                        `json:"code,required"`
	Message          string                                       `json:"message,required"`
	DocumentationURL string                                       `json:"documentation_url"`
	Source           ZonePageShieldScriptListResponseErrorsSource `json:"source"`
	JSON             zonePageShieldScriptListResponseErrorJSON    `json:"-"`
}

// zonePageShieldScriptListResponseErrorJSON contains the JSON metadata for the
// struct [ZonePageShieldScriptListResponseError]
type zonePageShieldScriptListResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZonePageShieldScriptListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zonePageShieldScriptListResponseErrorJSON) RawJSON() string {
	return r.raw
}

type ZonePageShieldScriptListResponseErrorsSource struct {
	Pointer string                                           `json:"pointer"`
	JSON    zonePageShieldScriptListResponseErrorsSourceJSON `json:"-"`
}

// zonePageShieldScriptListResponseErrorsSourceJSON contains the JSON metadata for
// the struct [ZonePageShieldScriptListResponseErrorsSource]
type zonePageShieldScriptListResponseErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonePageShieldScriptListResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zonePageShieldScriptListResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type ZonePageShieldScriptListResponseMessage struct {
	Code             int64                                          `json:"code,required"`
	Message          string                                         `json:"message,required"`
	DocumentationURL string                                         `json:"documentation_url"`
	Source           ZonePageShieldScriptListResponseMessagesSource `json:"source"`
	JSON             zonePageShieldScriptListResponseMessageJSON    `json:"-"`
}

// zonePageShieldScriptListResponseMessageJSON contains the JSON metadata for the
// struct [ZonePageShieldScriptListResponseMessage]
type zonePageShieldScriptListResponseMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZonePageShieldScriptListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zonePageShieldScriptListResponseMessageJSON) RawJSON() string {
	return r.raw
}

type ZonePageShieldScriptListResponseMessagesSource struct {
	Pointer string                                             `json:"pointer"`
	JSON    zonePageShieldScriptListResponseMessagesSourceJSON `json:"-"`
}

// zonePageShieldScriptListResponseMessagesSourceJSON contains the JSON metadata
// for the struct [ZonePageShieldScriptListResponseMessagesSource]
type zonePageShieldScriptListResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonePageShieldScriptListResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zonePageShieldScriptListResponseMessagesSourceJSON) RawJSON() string {
	return r.raw
}

type ZonePageShieldScriptListParams struct {
	// The direction used to sort returned scripts.
	Direction param.Field[ZonePageShieldScriptListParamsDirection] `query:"direction"`
	// When true, excludes scripts seen in a `/cdn-cgi` path from the returned scripts.
	// The default value is true.
	ExcludeCdnCgi param.Field[bool] `query:"exclude_cdn_cgi"`
	// When true, excludes duplicate scripts. We consider a script duplicate of another
	// if their javascript content matches and they share the same url host and zone
	// hostname. In such case, we return the most recent script for the URL host and
	// zone hostname combination.
	ExcludeDuplicates param.Field[bool] `query:"exclude_duplicates"`
	// Excludes scripts whose URL contains one of the URL-encoded URLs separated by
	// commas.
	ExcludeURLs param.Field[string] `query:"exclude_urls"`
	// Export the list of scripts as a file.
	Export param.Field[ZonePageShieldScriptListParamsExport] `query:"export"`
	// Includes scripts that match one or more URL-encoded hostnames separated by
	// commas.
	//
	// Wildcards are supported at the start and end of each hostname to support starts
	// with, ends with and contains. If no wildcards are used, results will be filtered
	// by exact match
	Hosts param.Field[string] `query:"hosts"`
	// The field used to sort returned scripts.
	OrderBy param.Field[ZonePageShieldScriptListParamsOrderBy] `query:"order_by"`
	// The current page number of the paginated results.
	//
	// We additionally support a special value "all". When "all" is used, the API will
	// return all the scripts with the applied filters in a single page. This feature
	// is best-effort and it may only work for zones with a low number of scripts
	Page param.Field[string] `query:"page"`
	// Includes scripts that match one or more page URLs (separated by commas) where
	// they were last seen
	//
	// Wildcards are supported at the start and end of each page URL to support starts
	// with, ends with and contains. If no wildcards are used, results will be filtered
	// by exact match
	PageURL param.Field[string] `query:"page_url"`
	// The number of results per page.
	PerPage param.Field[float64] `query:"per_page"`
	// When true, malicious scripts appear first in the returned scripts.
	PrioritizeMalicious param.Field[bool] `query:"prioritize_malicious"`
	// Filters the returned scripts using a comma-separated list of scripts statuses.
	// Accepted values: `active`, `infrequent`, and `inactive`. The default value is
	// `active`.
	Status param.Field[string] `query:"status"`
	// Includes scripts whose URL contain one or more URL-encoded URLs separated by
	// commas.
	URLs param.Field[string] `query:"urls"`
}

// URLQuery serializes [ZonePageShieldScriptListParams]'s query parameters as
// `url.Values`.
func (r ZonePageShieldScriptListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The direction used to sort returned scripts.
type ZonePageShieldScriptListParamsDirection string

const (
	ZonePageShieldScriptListParamsDirectionAsc  ZonePageShieldScriptListParamsDirection = "asc"
	ZonePageShieldScriptListParamsDirectionDesc ZonePageShieldScriptListParamsDirection = "desc"
)

func (r ZonePageShieldScriptListParamsDirection) IsKnown() bool {
	switch r {
	case ZonePageShieldScriptListParamsDirectionAsc, ZonePageShieldScriptListParamsDirectionDesc:
		return true
	}
	return false
}

// Export the list of scripts as a file.
type ZonePageShieldScriptListParamsExport string

const (
	ZonePageShieldScriptListParamsExportCsv ZonePageShieldScriptListParamsExport = "csv"
)

func (r ZonePageShieldScriptListParamsExport) IsKnown() bool {
	switch r {
	case ZonePageShieldScriptListParamsExportCsv:
		return true
	}
	return false
}

// The field used to sort returned scripts.
type ZonePageShieldScriptListParamsOrderBy string

const (
	ZonePageShieldScriptListParamsOrderByFirstSeenAt ZonePageShieldScriptListParamsOrderBy = "first_seen_at"
	ZonePageShieldScriptListParamsOrderByLastSeenAt  ZonePageShieldScriptListParamsOrderBy = "last_seen_at"
)

func (r ZonePageShieldScriptListParamsOrderBy) IsKnown() bool {
	switch r {
	case ZonePageShieldScriptListParamsOrderByFirstSeenAt, ZonePageShieldScriptListParamsOrderByLastSeenAt:
		return true
	}
	return false
}
