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
	Result ZonePageShieldScriptGetResponseResult `json:"result,required"`
	JSON   zonePageShieldScriptGetResponseJSON   `json:"-"`
	GetResponseCollection
}

// zonePageShieldScriptGetResponseJSON contains the JSON metadata for the struct
// [ZonePageShieldScriptGetResponse]
type zonePageShieldScriptGetResponseJSON struct {
	Result      apijson.Field
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
	Versions []ZonePageShieldScriptGetResponseResultVersion `json:"versions,nullable"`
	JSON     zonePageShieldScriptGetResponseResultJSON      `json:"-"`
	Script
}

// zonePageShieldScriptGetResponseResultJSON contains the JSON metadata for the
// struct [ZonePageShieldScriptGetResponseResult]
type zonePageShieldScriptGetResponseResultJSON struct {
	Versions    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
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

type ZonePageShieldScriptListResponse struct {
	Result []Script                             `json:"result,required"`
	JSON   zonePageShieldScriptListResponseJSON `json:"-"`
	ListResponseCollectionPageShield
}

// zonePageShieldScriptListResponseJSON contains the JSON metadata for the struct
// [ZonePageShieldScriptListResponse]
type zonePageShieldScriptListResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonePageShieldScriptListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zonePageShieldScriptListResponseJSON) RawJSON() string {
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
