// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"net/http"
	"net/url"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/apiquery"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
)

// RadarSearchService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRadarSearchService] method instead.
type RadarSearchService struct {
	Options []option.RequestOption
}

// NewRadarSearchService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRadarSearchService(opts ...option.RequestOption) (r *RadarSearchService) {
	r = &RadarSearchService{}
	r.Options = opts
	return
}

// Searches for locations, autonomous systems, reports, bots, certificate logs, and
// certificate authorities.
func (r *RadarSearchService) Global(ctx context.Context, query RadarSearchGlobalParams, opts ...option.RequestOption) (res *RadarSearchGlobalResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/search/global"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarSearchGlobalResponse struct {
	Result  RadarSearchGlobalResponseResult `json:"result,required"`
	Success bool                            `json:"success,required"`
	JSON    radarSearchGlobalResponseJSON   `json:"-"`
}

// radarSearchGlobalResponseJSON contains the JSON metadata for the struct
// [RadarSearchGlobalResponse]
type radarSearchGlobalResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarSearchGlobalResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarSearchGlobalResponseJSON) RawJSON() string {
	return r.raw
}

type RadarSearchGlobalResponseResult struct {
	Search []RadarSearchGlobalResponseResultSearch `json:"search,required"`
	JSON   radarSearchGlobalResponseResultJSON     `json:"-"`
}

// radarSearchGlobalResponseResultJSON contains the JSON metadata for the struct
// [RadarSearchGlobalResponseResult]
type radarSearchGlobalResponseResultJSON struct {
	Search      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarSearchGlobalResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarSearchGlobalResponseResultJSON) RawJSON() string {
	return r.raw
}

type RadarSearchGlobalResponseResultSearch struct {
	Code string                                    `json:"code,required"`
	Name string                                    `json:"name,required"`
	Type string                                    `json:"type,required"`
	JSON radarSearchGlobalResponseResultSearchJSON `json:"-"`
}

// radarSearchGlobalResponseResultSearchJSON contains the JSON metadata for the
// struct [RadarSearchGlobalResponseResultSearch]
type radarSearchGlobalResponseResultSearchJSON struct {
	Code        apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarSearchGlobalResponseResultSearch) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarSearchGlobalResponseResultSearchJSON) RawJSON() string {
	return r.raw
}

type RadarSearchGlobalParams struct {
	// String used to perform the search operation.
	Query param.Field[string] `query:"query,required"`
	// Search types excluded from results.
	Exclude param.Field[[]RadarSearchGlobalParamsExclude] `query:"exclude"`
	// Format in which results will be returned.
	Format param.Field[RadarSearchGlobalParamsFormat] `query:"format"`
	// Search types included in results.
	Include param.Field[[]RadarSearchGlobalParamsInclude] `query:"include"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Limits the number of objects per search category.
	LimitPerGroup param.Field[float64] `query:"limitPerGroup"`
}

// URLQuery serializes [RadarSearchGlobalParams]'s query parameters as
// `url.Values`.
func (r RadarSearchGlobalParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarSearchGlobalParamsExclude string

const (
	RadarSearchGlobalParamsExcludeAsns                   RadarSearchGlobalParamsExclude = "ASNS"
	RadarSearchGlobalParamsExcludeBots                   RadarSearchGlobalParamsExclude = "BOTS"
	RadarSearchGlobalParamsExcludeCertificateAuthorities RadarSearchGlobalParamsExclude = "CERTIFICATE_AUTHORITIES"
	RadarSearchGlobalParamsExcludeCertificateLogs        RadarSearchGlobalParamsExclude = "CERTIFICATE_LOGS"
	RadarSearchGlobalParamsExcludeLocations              RadarSearchGlobalParamsExclude = "LOCATIONS"
	RadarSearchGlobalParamsExcludeNotebooks              RadarSearchGlobalParamsExclude = "NOTEBOOKS"
)

func (r RadarSearchGlobalParamsExclude) IsKnown() bool {
	switch r {
	case RadarSearchGlobalParamsExcludeAsns, RadarSearchGlobalParamsExcludeBots, RadarSearchGlobalParamsExcludeCertificateAuthorities, RadarSearchGlobalParamsExcludeCertificateLogs, RadarSearchGlobalParamsExcludeLocations, RadarSearchGlobalParamsExcludeNotebooks:
		return true
	}
	return false
}

// Format in which results will be returned.
type RadarSearchGlobalParamsFormat string

const (
	RadarSearchGlobalParamsFormatJson RadarSearchGlobalParamsFormat = "JSON"
	RadarSearchGlobalParamsFormatCsv  RadarSearchGlobalParamsFormat = "CSV"
)

func (r RadarSearchGlobalParamsFormat) IsKnown() bool {
	switch r {
	case RadarSearchGlobalParamsFormatJson, RadarSearchGlobalParamsFormatCsv:
		return true
	}
	return false
}

type RadarSearchGlobalParamsInclude string

const (
	RadarSearchGlobalParamsIncludeAsns                   RadarSearchGlobalParamsInclude = "ASNS"
	RadarSearchGlobalParamsIncludeBots                   RadarSearchGlobalParamsInclude = "BOTS"
	RadarSearchGlobalParamsIncludeCertificateAuthorities RadarSearchGlobalParamsInclude = "CERTIFICATE_AUTHORITIES"
	RadarSearchGlobalParamsIncludeCertificateLogs        RadarSearchGlobalParamsInclude = "CERTIFICATE_LOGS"
	RadarSearchGlobalParamsIncludeLocations              RadarSearchGlobalParamsInclude = "LOCATIONS"
	RadarSearchGlobalParamsIncludeNotebooks              RadarSearchGlobalParamsInclude = "NOTEBOOKS"
)

func (r RadarSearchGlobalParamsInclude) IsKnown() bool {
	switch r {
	case RadarSearchGlobalParamsIncludeAsns, RadarSearchGlobalParamsIncludeBots, RadarSearchGlobalParamsIncludeCertificateAuthorities, RadarSearchGlobalParamsIncludeCertificateLogs, RadarSearchGlobalParamsIncludeLocations, RadarSearchGlobalParamsIncludeNotebooks:
		return true
	}
	return false
}
