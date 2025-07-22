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

// ZonePageShieldCookieService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZonePageShieldCookieService] method instead.
type ZonePageShieldCookieService struct {
	Options []option.RequestOption
}

// NewZonePageShieldCookieService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZonePageShieldCookieService(opts ...option.RequestOption) (r *ZonePageShieldCookieService) {
	r = &ZonePageShieldCookieService{}
	r.Options = opts
	return
}

// Fetches a cookie collected by Page Shield by cookie ID.
func (r *ZonePageShieldCookieService) Get(ctx context.Context, zoneID string, cookieID string, opts ...option.RequestOption) (res *ZonePageShieldCookieGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if cookieID == "" {
		err = errors.New("missing required cookie_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/page_shield/cookies/%s", zoneID, cookieID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Lists all cookies collected by Page Shield.
func (r *ZonePageShieldCookieService) List(ctx context.Context, zoneID string, query ZonePageShieldCookieListParams, opts ...option.RequestOption) (res *ZonePageShieldCookieListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/page_shield/cookies", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type Cookie struct {
	// Identifier
	ID                string                  `json:"id,required"`
	FirstSeenAt       time.Time               `json:"first_seen_at,required" format:"date-time"`
	Host              string                  `json:"host,required"`
	LastSeenAt        time.Time               `json:"last_seen_at,required" format:"date-time"`
	Name              string                  `json:"name,required"`
	Type              CookieType              `json:"type,required"`
	DomainAttribute   string                  `json:"domain_attribute"`
	ExpiresAttribute  time.Time               `json:"expires_attribute" format:"date-time"`
	HTTPOnlyAttribute bool                    `json:"http_only_attribute"`
	MaxAgeAttribute   int64                   `json:"max_age_attribute"`
	PageURLs          []string                `json:"page_urls"`
	PathAttribute     string                  `json:"path_attribute"`
	SameSiteAttribute CookieSameSiteAttribute `json:"same_site_attribute"`
	SecureAttribute   bool                    `json:"secure_attribute"`
	JSON              cookieJSON              `json:"-"`
}

// cookieJSON contains the JSON metadata for the struct [Cookie]
type cookieJSON struct {
	ID                apijson.Field
	FirstSeenAt       apijson.Field
	Host              apijson.Field
	LastSeenAt        apijson.Field
	Name              apijson.Field
	Type              apijson.Field
	DomainAttribute   apijson.Field
	ExpiresAttribute  apijson.Field
	HTTPOnlyAttribute apijson.Field
	MaxAgeAttribute   apijson.Field
	PageURLs          apijson.Field
	PathAttribute     apijson.Field
	SameSiteAttribute apijson.Field
	SecureAttribute   apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *Cookie) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cookieJSON) RawJSON() string {
	return r.raw
}

type CookieType string

const (
	CookieTypeFirstParty CookieType = "first_party"
	CookieTypeUnknown    CookieType = "unknown"
)

func (r CookieType) IsKnown() bool {
	switch r {
	case CookieTypeFirstParty, CookieTypeUnknown:
		return true
	}
	return false
}

type CookieSameSiteAttribute string

const (
	CookieSameSiteAttributeLax    CookieSameSiteAttribute = "lax"
	CookieSameSiteAttributeStrict CookieSameSiteAttribute = "strict"
	CookieSameSiteAttributeNone   CookieSameSiteAttribute = "none"
)

func (r CookieSameSiteAttribute) IsKnown() bool {
	switch r {
	case CookieSameSiteAttributeLax, CookieSameSiteAttributeStrict, CookieSameSiteAttributeNone:
		return true
	}
	return false
}

type ZonePageShieldCookieGetResponse struct {
	Result Cookie                              `json:"result,required"`
	JSON   zonePageShieldCookieGetResponseJSON `json:"-"`
	GetResponseCollection
}

// zonePageShieldCookieGetResponseJSON contains the JSON metadata for the struct
// [ZonePageShieldCookieGetResponse]
type zonePageShieldCookieGetResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonePageShieldCookieGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zonePageShieldCookieGetResponseJSON) RawJSON() string {
	return r.raw
}

type ZonePageShieldCookieListResponse struct {
	Result []Cookie                             `json:"result,required"`
	JSON   zonePageShieldCookieListResponseJSON `json:"-"`
	ListResponseCollectionPageShield
}

// zonePageShieldCookieListResponseJSON contains the JSON metadata for the struct
// [ZonePageShieldCookieListResponse]
type zonePageShieldCookieListResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonePageShieldCookieListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zonePageShieldCookieListResponseJSON) RawJSON() string {
	return r.raw
}

type ZonePageShieldCookieListParams struct {
	// The direction used to sort returned cookies.'
	Direction param.Field[ZonePageShieldCookieListParamsDirection] `query:"direction"`
	// Filters the returned cookies that match the specified domain attribute
	Domain param.Field[string] `query:"domain"`
	// Export the list of cookies as a file.
	Export param.Field[ZonePageShieldCookieListParamsExport] `query:"export"`
	// Includes cookies that match one or more URL-encoded hostnames separated by
	// commas.
	//
	// Wildcards are supported at the start and end of each hostname to support starts
	// with, ends with and contains. If no wildcards are used, results will be filtered
	// by exact match
	Hosts param.Field[string] `query:"hosts"`
	// Filters the returned cookies that are set with HttpOnly
	HTTPOnly param.Field[bool] `query:"http_only"`
	// Filters the returned cookies that match the specified name. Wildcards are
	// supported at the start and end to support starts with, ends with and contains.
	// e.g. session\*
	Name param.Field[string] `query:"name"`
	// The field used to sort returned cookies.
	OrderBy param.Field[ZonePageShieldCookieListParamsOrderBy] `query:"order_by"`
	// The current page number of the paginated results.
	//
	// We additionally support a special value "all". When "all" is used, the API will
	// return all the cookies with the applied filters in a single page. This feature
	// is best-effort and it may only work for zones with a low number of cookies
	Page param.Field[string] `query:"page"`
	// Includes connections that match one or more page URLs (separated by commas)
	// where they were last seen
	//
	// Wildcards are supported at the start and end of each page URL to support starts
	// with, ends with and contains. If no wildcards are used, results will be filtered
	// by exact match
	PageURL param.Field[string] `query:"page_url"`
	// Filters the returned cookies that match the specified path attribute
	Path param.Field[string] `query:"path"`
	// The number of results per page.
	PerPage param.Field[float64] `query:"per_page"`
	// Filters the returned cookies that match the specified same_site attribute
	SameSite param.Field[ZonePageShieldCookieListParamsSameSite] `query:"same_site"`
	// Filters the returned cookies that are set with Secure
	Secure param.Field[bool] `query:"secure"`
	// Filters the returned cookies that match the specified type attribute
	Type param.Field[ZonePageShieldCookieListParamsType] `query:"type"`
}

// URLQuery serializes [ZonePageShieldCookieListParams]'s query parameters as
// `url.Values`.
func (r ZonePageShieldCookieListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The direction used to sort returned cookies.'
type ZonePageShieldCookieListParamsDirection string

const (
	ZonePageShieldCookieListParamsDirectionAsc  ZonePageShieldCookieListParamsDirection = "asc"
	ZonePageShieldCookieListParamsDirectionDesc ZonePageShieldCookieListParamsDirection = "desc"
)

func (r ZonePageShieldCookieListParamsDirection) IsKnown() bool {
	switch r {
	case ZonePageShieldCookieListParamsDirectionAsc, ZonePageShieldCookieListParamsDirectionDesc:
		return true
	}
	return false
}

// Export the list of cookies as a file.
type ZonePageShieldCookieListParamsExport string

const (
	ZonePageShieldCookieListParamsExportCsv ZonePageShieldCookieListParamsExport = "csv"
)

func (r ZonePageShieldCookieListParamsExport) IsKnown() bool {
	switch r {
	case ZonePageShieldCookieListParamsExportCsv:
		return true
	}
	return false
}

// The field used to sort returned cookies.
type ZonePageShieldCookieListParamsOrderBy string

const (
	ZonePageShieldCookieListParamsOrderByFirstSeenAt ZonePageShieldCookieListParamsOrderBy = "first_seen_at"
	ZonePageShieldCookieListParamsOrderByLastSeenAt  ZonePageShieldCookieListParamsOrderBy = "last_seen_at"
)

func (r ZonePageShieldCookieListParamsOrderBy) IsKnown() bool {
	switch r {
	case ZonePageShieldCookieListParamsOrderByFirstSeenAt, ZonePageShieldCookieListParamsOrderByLastSeenAt:
		return true
	}
	return false
}

// Filters the returned cookies that match the specified same_site attribute
type ZonePageShieldCookieListParamsSameSite string

const (
	ZonePageShieldCookieListParamsSameSiteLax    ZonePageShieldCookieListParamsSameSite = "lax"
	ZonePageShieldCookieListParamsSameSiteStrict ZonePageShieldCookieListParamsSameSite = "strict"
	ZonePageShieldCookieListParamsSameSiteNone   ZonePageShieldCookieListParamsSameSite = "none"
)

func (r ZonePageShieldCookieListParamsSameSite) IsKnown() bool {
	switch r {
	case ZonePageShieldCookieListParamsSameSiteLax, ZonePageShieldCookieListParamsSameSiteStrict, ZonePageShieldCookieListParamsSameSiteNone:
		return true
	}
	return false
}

// Filters the returned cookies that match the specified type attribute
type ZonePageShieldCookieListParamsType string

const (
	ZonePageShieldCookieListParamsTypeFirstParty ZonePageShieldCookieListParamsType = "first_party"
	ZonePageShieldCookieListParamsTypeUnknown    ZonePageShieldCookieListParamsType = "unknown"
)

func (r ZonePageShieldCookieListParamsType) IsKnown() bool {
	switch r {
	case ZonePageShieldCookieListParamsTypeFirstParty, ZonePageShieldCookieListParamsTypeUnknown:
		return true
	}
	return false
}
