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

// ZonePageShieldConnectionService contains methods and other services that help
// with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZonePageShieldConnectionService] method instead.
type ZonePageShieldConnectionService struct {
	Options []option.RequestOption
}

// NewZonePageShieldConnectionService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZonePageShieldConnectionService(opts ...option.RequestOption) (r *ZonePageShieldConnectionService) {
	r = &ZonePageShieldConnectionService{}
	r.Options = opts
	return
}

// Fetches a connection detected by Page Shield by connection ID.
func (r *ZonePageShieldConnectionService) Get(ctx context.Context, zoneID string, connectionID string, opts ...option.RequestOption) (res *ZonePageShieldConnectionGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if connectionID == "" {
		err = errors.New("missing required connection_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/page_shield/connections/%s", zoneID, connectionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Lists all connections detected by Page Shield.
func (r *ZonePageShieldConnectionService) List(ctx context.Context, zoneID string, query ZonePageShieldConnectionListParams, opts ...option.RequestOption) (res *ZonePageShieldConnectionListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/page_shield/connections", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type Connection struct {
	// Identifier
	ID                        string         `json:"id,required"`
	AddedAt                   time.Time      `json:"added_at,required" format:"date-time"`
	FirstSeenAt               time.Time      `json:"first_seen_at,required" format:"date-time"`
	Host                      string         `json:"host,required"`
	LastSeenAt                time.Time      `json:"last_seen_at,required" format:"date-time"`
	URL                       string         `json:"url,required"`
	URLContainsCdnCgiPath     bool           `json:"url_contains_cdn_cgi_path,required"`
	DomainReportedMalicious   bool           `json:"domain_reported_malicious"`
	FirstPageURL              string         `json:"first_page_url"`
	MaliciousDomainCategories []string       `json:"malicious_domain_categories"`
	MaliciousURLCategories    []string       `json:"malicious_url_categories"`
	PageURLs                  []string       `json:"page_urls"`
	URLReportedMalicious      bool           `json:"url_reported_malicious"`
	JSON                      connectionJSON `json:"-"`
}

// connectionJSON contains the JSON metadata for the struct [Connection]
type connectionJSON struct {
	ID                        apijson.Field
	AddedAt                   apijson.Field
	FirstSeenAt               apijson.Field
	Host                      apijson.Field
	LastSeenAt                apijson.Field
	URL                       apijson.Field
	URLContainsCdnCgiPath     apijson.Field
	DomainReportedMalicious   apijson.Field
	FirstPageURL              apijson.Field
	MaliciousDomainCategories apijson.Field
	MaliciousURLCategories    apijson.Field
	PageURLs                  apijson.Field
	URLReportedMalicious      apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *Connection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectionJSON) RawJSON() string {
	return r.raw
}

type ListResponseCollectionPageShield struct {
	ResultInfo ListResponseCollectionPageShieldResultInfo `json:"result_info,required"`
	JSON       listResponseCollectionPageShieldJSON       `json:"-"`
	ResponseCommonShield
}

// listResponseCollectionPageShieldJSON contains the JSON metadata for the struct
// [ListResponseCollectionPageShield]
type listResponseCollectionPageShieldJSON struct {
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ListResponseCollectionPageShield) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listResponseCollectionPageShieldJSON) RawJSON() string {
	return r.raw
}

type ListResponseCollectionPageShieldResultInfo struct {
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
	JSON       listResponseCollectionPageShieldResultInfoJSON `json:"-"`
}

// listResponseCollectionPageShieldResultInfoJSON contains the JSON metadata for
// the struct [ListResponseCollectionPageShieldResultInfo]
type listResponseCollectionPageShieldResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	TotalPages  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ListResponseCollectionPageShieldResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listResponseCollectionPageShieldResultInfoJSON) RawJSON() string {
	return r.raw
}

type ZonePageShieldConnectionGetResponse struct {
	Result Connection                              `json:"result,required"`
	JSON   zonePageShieldConnectionGetResponseJSON `json:"-"`
	GetResponseCollection
}

// zonePageShieldConnectionGetResponseJSON contains the JSON metadata for the
// struct [ZonePageShieldConnectionGetResponse]
type zonePageShieldConnectionGetResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonePageShieldConnectionGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zonePageShieldConnectionGetResponseJSON) RawJSON() string {
	return r.raw
}

type ZonePageShieldConnectionListResponse struct {
	Result []Connection                             `json:"result"`
	JSON   zonePageShieldConnectionListResponseJSON `json:"-"`
	ListResponseCollectionPageShield
}

// zonePageShieldConnectionListResponseJSON contains the JSON metadata for the
// struct [ZonePageShieldConnectionListResponse]
type zonePageShieldConnectionListResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonePageShieldConnectionListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zonePageShieldConnectionListResponseJSON) RawJSON() string {
	return r.raw
}

type ZonePageShieldConnectionListParams struct {
	// The direction used to sort returned connections.
	Direction param.Field[ZonePageShieldConnectionListParamsDirection] `query:"direction"`
	// When true, excludes connections seen in a `/cdn-cgi` path from the returned
	// connections. The default value is true.
	ExcludeCdnCgi param.Field[bool] `query:"exclude_cdn_cgi"`
	// Excludes connections whose URL contains one of the URL-encoded URLs separated by
	// commas.
	ExcludeURLs param.Field[string] `query:"exclude_urls"`
	// Export the list of connections as a file.
	Export param.Field[ZonePageShieldConnectionListParamsExport] `query:"export"`
	// Includes connections that match one or more URL-encoded hostnames separated by
	// commas.
	//
	// Wildcards are supported at the start and end of each hostname to support starts
	// with, ends with and contains. If no wildcards are used, results will be filtered
	// by exact match
	Hosts param.Field[string] `query:"hosts"`
	// The field used to sort returned connections.
	OrderBy param.Field[ZonePageShieldConnectionListParamsOrderBy] `query:"order_by"`
	// The current page number of the paginated results.
	//
	// We additionally support a special value "all". When "all" is used, the API will
	// return all the connections with the applied filters in a single page. This
	// feature is best-effort and it may only work for zones with a low number of
	// connections
	Page param.Field[string] `query:"page"`
	// Includes connections that match one or more page URLs (separated by commas)
	// where they were last seen
	//
	// Wildcards are supported at the start and end of each page URL to support starts
	// with, ends with and contains. If no wildcards are used, results will be filtered
	// by exact match
	PageURL param.Field[string] `query:"page_url"`
	// The number of results per page.
	PerPage param.Field[float64] `query:"per_page"`
	// When true, malicious connections appear first in the returned connections.
	PrioritizeMalicious param.Field[bool] `query:"prioritize_malicious"`
	// Filters the returned connections using a comma-separated list of connection
	// statuses. Accepted values: `active`, `infrequent`, and `inactive`. The default
	// value is `active`.
	Status param.Field[string] `query:"status"`
	// Includes connections whose URL contain one or more URL-encoded URLs separated by
	// commas.
	URLs param.Field[string] `query:"urls"`
}

// URLQuery serializes [ZonePageShieldConnectionListParams]'s query parameters as
// `url.Values`.
func (r ZonePageShieldConnectionListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The direction used to sort returned connections.
type ZonePageShieldConnectionListParamsDirection string

const (
	ZonePageShieldConnectionListParamsDirectionAsc  ZonePageShieldConnectionListParamsDirection = "asc"
	ZonePageShieldConnectionListParamsDirectionDesc ZonePageShieldConnectionListParamsDirection = "desc"
)

func (r ZonePageShieldConnectionListParamsDirection) IsKnown() bool {
	switch r {
	case ZonePageShieldConnectionListParamsDirectionAsc, ZonePageShieldConnectionListParamsDirectionDesc:
		return true
	}
	return false
}

// Export the list of connections as a file.
type ZonePageShieldConnectionListParamsExport string

const (
	ZonePageShieldConnectionListParamsExportCsv ZonePageShieldConnectionListParamsExport = "csv"
)

func (r ZonePageShieldConnectionListParamsExport) IsKnown() bool {
	switch r {
	case ZonePageShieldConnectionListParamsExportCsv:
		return true
	}
	return false
}

// The field used to sort returned connections.
type ZonePageShieldConnectionListParamsOrderBy string

const (
	ZonePageShieldConnectionListParamsOrderByFirstSeenAt ZonePageShieldConnectionListParamsOrderBy = "first_seen_at"
	ZonePageShieldConnectionListParamsOrderByLastSeenAt  ZonePageShieldConnectionListParamsOrderBy = "last_seen_at"
)

func (r ZonePageShieldConnectionListParamsOrderBy) IsKnown() bool {
	switch r {
	case ZonePageShieldConnectionListParamsOrderByFirstSeenAt, ZonePageShieldConnectionListParamsOrderByLastSeenAt:
		return true
	}
	return false
}
