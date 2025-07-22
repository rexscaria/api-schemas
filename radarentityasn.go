// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/stainless-sdks/cf-rex-go/internal/apijson"
	"github.com/stainless-sdks/cf-rex-go/internal/apiquery"
	"github.com/stainless-sdks/cf-rex-go/internal/param"
	"github.com/stainless-sdks/cf-rex-go/internal/requestconfig"
	"github.com/stainless-sdks/cf-rex-go/option"
)

// RadarEntityAsnService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRadarEntityAsnService] method instead.
type RadarEntityAsnService struct {
	Options []option.RequestOption
}

// NewRadarEntityAsnService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRadarEntityAsnService(opts ...option.RequestOption) (r *RadarEntityAsnService) {
	r = &RadarEntityAsnService{}
	r.Options = opts
	return
}

// Retrieves the requested autonomous system information. (A confidence level below
// `5` indicates a low level of confidence in the traffic data - normally this
// happens because Cloudflare has a small amount of traffic from/to this AS).
// Population estimates come from APNIC (refer to https://labs.apnic.net/?p=526).
func (r *RadarEntityAsnService) Get(ctx context.Context, asn int64, query RadarEntityAsnGetParams, opts ...option.RequestOption) (res *RadarEntityAsnGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("radar/entities/asns/%v", asn)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves a list of autonomous systems.
func (r *RadarEntityAsnService) List(ctx context.Context, query RadarEntityAsnListParams, opts ...option.RequestOption) (res *RadarEntityAsnListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/entities/asns"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves AS-level relationship for given networks.
func (r *RadarEntityAsnService) GetRelationships(ctx context.Context, asn int64, query RadarEntityAsnGetRelationshipsParams, opts ...option.RequestOption) (res *RadarEntityAsnGetRelationshipsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("radar/entities/asns/%v/rel", asn)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves the requested autonomous system information based on IP address.
// Population estimates come from APNIC (refer to https://labs.apnic.net/?p=526).
func (r *RadarEntityAsnService) GetByIP(ctx context.Context, query RadarEntityAsnGetByIPParams, opts ...option.RequestOption) (res *RadarEntityAsnGetByIPResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/entities/asns/ip"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarEntityAsnGetResponse struct {
	Result  RadarEntityAsnGetResponseResult `json:"result,required"`
	Success bool                            `json:"success,required"`
	JSON    radarEntityAsnGetResponseJSON   `json:"-"`
}

// radarEntityAsnGetResponseJSON contains the JSON metadata for the struct
// [RadarEntityAsnGetResponse]
type radarEntityAsnGetResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEntityAsnGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEntityAsnGetResponseJSON) RawJSON() string {
	return r.raw
}

type RadarEntityAsnGetResponseResult struct {
	Asn  RadarEntityAsnGetResponseResultAsn  `json:"asn,required"`
	JSON radarEntityAsnGetResponseResultJSON `json:"-"`
}

// radarEntityAsnGetResponseResultJSON contains the JSON metadata for the struct
// [RadarEntityAsnGetResponseResult]
type radarEntityAsnGetResponseResultJSON struct {
	Asn         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEntityAsnGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEntityAsnGetResponseResultJSON) RawJSON() string {
	return r.raw
}

type RadarEntityAsnGetResponseResultAsn struct {
	Asn             int64                                            `json:"asn,required"`
	ConfidenceLevel int64                                            `json:"confidenceLevel,required"`
	Country         string                                           `json:"country,required"`
	CountryName     string                                           `json:"countryName,required"`
	EstimatedUsers  RadarEntityAsnGetResponseResultAsnEstimatedUsers `json:"estimatedUsers,required"`
	Name            string                                           `json:"name,required"`
	OrgName         string                                           `json:"orgName,required"`
	Related         []RadarEntityAsnGetResponseResultAsnRelated      `json:"related,required"`
	// Regional Internet Registry
	Source  string `json:"source,required"`
	Website string `json:"website,required"`
	Aka     string `json:"aka"`
	// Deprecated field. Please use 'aka'.
	NameLong string                                 `json:"nameLong"`
	JSON     radarEntityAsnGetResponseResultAsnJSON `json:"-"`
}

// radarEntityAsnGetResponseResultAsnJSON contains the JSON metadata for the struct
// [RadarEntityAsnGetResponseResultAsn]
type radarEntityAsnGetResponseResultAsnJSON struct {
	Asn             apijson.Field
	ConfidenceLevel apijson.Field
	Country         apijson.Field
	CountryName     apijson.Field
	EstimatedUsers  apijson.Field
	Name            apijson.Field
	OrgName         apijson.Field
	Related         apijson.Field
	Source          apijson.Field
	Website         apijson.Field
	Aka             apijson.Field
	NameLong        apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *RadarEntityAsnGetResponseResultAsn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEntityAsnGetResponseResultAsnJSON) RawJSON() string {
	return r.raw
}

type RadarEntityAsnGetResponseResultAsnEstimatedUsers struct {
	Locations []RadarEntityAsnGetResponseResultAsnEstimatedUsersLocation `json:"locations,required"`
	// Total estimated users
	EstimatedUsers int64                                                `json:"estimatedUsers"`
	JSON           radarEntityAsnGetResponseResultAsnEstimatedUsersJSON `json:"-"`
}

// radarEntityAsnGetResponseResultAsnEstimatedUsersJSON contains the JSON metadata
// for the struct [RadarEntityAsnGetResponseResultAsnEstimatedUsers]
type radarEntityAsnGetResponseResultAsnEstimatedUsersJSON struct {
	Locations      apijson.Field
	EstimatedUsers apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarEntityAsnGetResponseResultAsnEstimatedUsers) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEntityAsnGetResponseResultAsnEstimatedUsersJSON) RawJSON() string {
	return r.raw
}

type RadarEntityAsnGetResponseResultAsnEstimatedUsersLocation struct {
	LocationAlpha2 string `json:"locationAlpha2,required"`
	LocationName   string `json:"locationName,required"`
	// Estimated users per location
	EstimatedUsers int64                                                        `json:"estimatedUsers"`
	JSON           radarEntityAsnGetResponseResultAsnEstimatedUsersLocationJSON `json:"-"`
}

// radarEntityAsnGetResponseResultAsnEstimatedUsersLocationJSON contains the JSON
// metadata for the struct
// [RadarEntityAsnGetResponseResultAsnEstimatedUsersLocation]
type radarEntityAsnGetResponseResultAsnEstimatedUsersLocationJSON struct {
	LocationAlpha2 apijson.Field
	LocationName   apijson.Field
	EstimatedUsers apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarEntityAsnGetResponseResultAsnEstimatedUsersLocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEntityAsnGetResponseResultAsnEstimatedUsersLocationJSON) RawJSON() string {
	return r.raw
}

type RadarEntityAsnGetResponseResultAsnRelated struct {
	Asn  int64  `json:"asn,required"`
	Name string `json:"name,required"`
	Aka  string `json:"aka"`
	// Total estimated users
	EstimatedUsers int64                                         `json:"estimatedUsers"`
	JSON           radarEntityAsnGetResponseResultAsnRelatedJSON `json:"-"`
}

// radarEntityAsnGetResponseResultAsnRelatedJSON contains the JSON metadata for the
// struct [RadarEntityAsnGetResponseResultAsnRelated]
type radarEntityAsnGetResponseResultAsnRelatedJSON struct {
	Asn            apijson.Field
	Name           apijson.Field
	Aka            apijson.Field
	EstimatedUsers apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarEntityAsnGetResponseResultAsnRelated) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEntityAsnGetResponseResultAsnRelatedJSON) RawJSON() string {
	return r.raw
}

type RadarEntityAsnListResponse struct {
	Result  RadarEntityAsnListResponseResult `json:"result,required"`
	Success bool                             `json:"success,required"`
	JSON    radarEntityAsnListResponseJSON   `json:"-"`
}

// radarEntityAsnListResponseJSON contains the JSON metadata for the struct
// [RadarEntityAsnListResponse]
type radarEntityAsnListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEntityAsnListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEntityAsnListResponseJSON) RawJSON() string {
	return r.raw
}

type RadarEntityAsnListResponseResult struct {
	Asns []RadarEntityAsnListResponseResultAsn `json:"asns,required"`
	JSON radarEntityAsnListResponseResultJSON  `json:"-"`
}

// radarEntityAsnListResponseResultJSON contains the JSON metadata for the struct
// [RadarEntityAsnListResponseResult]
type radarEntityAsnListResponseResultJSON struct {
	Asns        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEntityAsnListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEntityAsnListResponseResultJSON) RawJSON() string {
	return r.raw
}

type RadarEntityAsnListResponseResultAsn struct {
	Asn         int64  `json:"asn,required"`
	Country     string `json:"country,required"`
	CountryName string `json:"countryName,required"`
	Name        string `json:"name,required"`
	Aka         string `json:"aka"`
	// Deprecated field. Please use 'aka'.
	NameLong string                                  `json:"nameLong"`
	OrgName  string                                  `json:"orgName"`
	Website  string                                  `json:"website"`
	JSON     radarEntityAsnListResponseResultAsnJSON `json:"-"`
}

// radarEntityAsnListResponseResultAsnJSON contains the JSON metadata for the
// struct [RadarEntityAsnListResponseResultAsn]
type radarEntityAsnListResponseResultAsnJSON struct {
	Asn         apijson.Field
	Country     apijson.Field
	CountryName apijson.Field
	Name        apijson.Field
	Aka         apijson.Field
	NameLong    apijson.Field
	OrgName     apijson.Field
	Website     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEntityAsnListResponseResultAsn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEntityAsnListResponseResultAsnJSON) RawJSON() string {
	return r.raw
}

type RadarEntityAsnGetRelationshipsResponse struct {
	Result  RadarEntityAsnGetRelationshipsResponseResult `json:"result,required"`
	Success bool                                         `json:"success,required"`
	JSON    radarEntityAsnGetRelationshipsResponseJSON   `json:"-"`
}

// radarEntityAsnGetRelationshipsResponseJSON contains the JSON metadata for the
// struct [RadarEntityAsnGetRelationshipsResponse]
type radarEntityAsnGetRelationshipsResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEntityAsnGetRelationshipsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEntityAsnGetRelationshipsResponseJSON) RawJSON() string {
	return r.raw
}

type RadarEntityAsnGetRelationshipsResponseResult struct {
	Meta RadarEntityAsnGetRelationshipsResponseResultMeta  `json:"meta,required"`
	Rels []RadarEntityAsnGetRelationshipsResponseResultRel `json:"rels,required"`
	JSON radarEntityAsnGetRelationshipsResponseResultJSON  `json:"-"`
}

// radarEntityAsnGetRelationshipsResponseResultJSON contains the JSON metadata for
// the struct [RadarEntityAsnGetRelationshipsResponseResult]
type radarEntityAsnGetRelationshipsResponseResultJSON struct {
	Meta        apijson.Field
	Rels        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEntityAsnGetRelationshipsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEntityAsnGetRelationshipsResponseResultJSON) RawJSON() string {
	return r.raw
}

type RadarEntityAsnGetRelationshipsResponseResultMeta struct {
	DataTime   string                                               `json:"data_time,required"`
	QueryTime  string                                               `json:"query_time,required"`
	TotalPeers int64                                                `json:"total_peers,required"`
	JSON       radarEntityAsnGetRelationshipsResponseResultMetaJSON `json:"-"`
}

// radarEntityAsnGetRelationshipsResponseResultMetaJSON contains the JSON metadata
// for the struct [RadarEntityAsnGetRelationshipsResponseResultMeta]
type radarEntityAsnGetRelationshipsResponseResultMetaJSON struct {
	DataTime    apijson.Field
	QueryTime   apijson.Field
	TotalPeers  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEntityAsnGetRelationshipsResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEntityAsnGetRelationshipsResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

type RadarEntityAsnGetRelationshipsResponseResultRel struct {
	Asn1        int64                                               `json:"asn1,required"`
	Asn1Country string                                              `json:"asn1_country,required"`
	Asn1Name    string                                              `json:"asn1_name,required"`
	Asn2        int64                                               `json:"asn2,required"`
	Asn2Country string                                              `json:"asn2_country,required"`
	Asn2Name    string                                              `json:"asn2_name,required"`
	Rel         string                                              `json:"rel,required"`
	JSON        radarEntityAsnGetRelationshipsResponseResultRelJSON `json:"-"`
}

// radarEntityAsnGetRelationshipsResponseResultRelJSON contains the JSON metadata
// for the struct [RadarEntityAsnGetRelationshipsResponseResultRel]
type radarEntityAsnGetRelationshipsResponseResultRelJSON struct {
	Asn1        apijson.Field
	Asn1Country apijson.Field
	Asn1Name    apijson.Field
	Asn2        apijson.Field
	Asn2Country apijson.Field
	Asn2Name    apijson.Field
	Rel         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEntityAsnGetRelationshipsResponseResultRel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEntityAsnGetRelationshipsResponseResultRelJSON) RawJSON() string {
	return r.raw
}

type RadarEntityAsnGetByIPResponse struct {
	Result  RadarEntityAsnGetByIPResponseResult `json:"result,required"`
	Success bool                                `json:"success,required"`
	JSON    radarEntityAsnGetByIPResponseJSON   `json:"-"`
}

// radarEntityAsnGetByIPResponseJSON contains the JSON metadata for the struct
// [RadarEntityAsnGetByIPResponse]
type radarEntityAsnGetByIPResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEntityAsnGetByIPResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEntityAsnGetByIPResponseJSON) RawJSON() string {
	return r.raw
}

type RadarEntityAsnGetByIPResponseResult struct {
	Asn  RadarEntityAsnGetByIPResponseResultAsn  `json:"asn,required"`
	JSON radarEntityAsnGetByIPResponseResultJSON `json:"-"`
}

// radarEntityAsnGetByIPResponseResultJSON contains the JSON metadata for the
// struct [RadarEntityAsnGetByIPResponseResult]
type radarEntityAsnGetByIPResponseResultJSON struct {
	Asn         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEntityAsnGetByIPResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEntityAsnGetByIPResponseResultJSON) RawJSON() string {
	return r.raw
}

type RadarEntityAsnGetByIPResponseResultAsn struct {
	Asn            int64                                                `json:"asn,required"`
	Country        string                                               `json:"country,required"`
	CountryName    string                                               `json:"countryName,required"`
	EstimatedUsers RadarEntityAsnGetByIPResponseResultAsnEstimatedUsers `json:"estimatedUsers,required"`
	Name           string                                               `json:"name,required"`
	OrgName        string                                               `json:"orgName,required"`
	Related        []RadarEntityAsnGetByIPResponseResultAsnRelated      `json:"related,required"`
	// Regional Internet Registry
	Source  string `json:"source,required"`
	Website string `json:"website,required"`
	Aka     string `json:"aka"`
	// Deprecated field. Please use 'aka'.
	NameLong string                                     `json:"nameLong"`
	JSON     radarEntityAsnGetByIPResponseResultAsnJSON `json:"-"`
}

// radarEntityAsnGetByIPResponseResultAsnJSON contains the JSON metadata for the
// struct [RadarEntityAsnGetByIPResponseResultAsn]
type radarEntityAsnGetByIPResponseResultAsnJSON struct {
	Asn            apijson.Field
	Country        apijson.Field
	CountryName    apijson.Field
	EstimatedUsers apijson.Field
	Name           apijson.Field
	OrgName        apijson.Field
	Related        apijson.Field
	Source         apijson.Field
	Website        apijson.Field
	Aka            apijson.Field
	NameLong       apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarEntityAsnGetByIPResponseResultAsn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEntityAsnGetByIPResponseResultAsnJSON) RawJSON() string {
	return r.raw
}

type RadarEntityAsnGetByIPResponseResultAsnEstimatedUsers struct {
	Locations []RadarEntityAsnGetByIPResponseResultAsnEstimatedUsersLocation `json:"locations,required"`
	// Total estimated users
	EstimatedUsers int64                                                    `json:"estimatedUsers"`
	JSON           radarEntityAsnGetByIPResponseResultAsnEstimatedUsersJSON `json:"-"`
}

// radarEntityAsnGetByIPResponseResultAsnEstimatedUsersJSON contains the JSON
// metadata for the struct [RadarEntityAsnGetByIPResponseResultAsnEstimatedUsers]
type radarEntityAsnGetByIPResponseResultAsnEstimatedUsersJSON struct {
	Locations      apijson.Field
	EstimatedUsers apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarEntityAsnGetByIPResponseResultAsnEstimatedUsers) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEntityAsnGetByIPResponseResultAsnEstimatedUsersJSON) RawJSON() string {
	return r.raw
}

type RadarEntityAsnGetByIPResponseResultAsnEstimatedUsersLocation struct {
	LocationAlpha2 string `json:"locationAlpha2,required"`
	LocationName   string `json:"locationName,required"`
	// Estimated users per location
	EstimatedUsers int64                                                            `json:"estimatedUsers"`
	JSON           radarEntityAsnGetByIPResponseResultAsnEstimatedUsersLocationJSON `json:"-"`
}

// radarEntityAsnGetByIPResponseResultAsnEstimatedUsersLocationJSON contains the
// JSON metadata for the struct
// [RadarEntityAsnGetByIPResponseResultAsnEstimatedUsersLocation]
type radarEntityAsnGetByIPResponseResultAsnEstimatedUsersLocationJSON struct {
	LocationAlpha2 apijson.Field
	LocationName   apijson.Field
	EstimatedUsers apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarEntityAsnGetByIPResponseResultAsnEstimatedUsersLocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEntityAsnGetByIPResponseResultAsnEstimatedUsersLocationJSON) RawJSON() string {
	return r.raw
}

type RadarEntityAsnGetByIPResponseResultAsnRelated struct {
	Asn  int64  `json:"asn,required"`
	Name string `json:"name,required"`
	Aka  string `json:"aka"`
	// Total estimated users
	EstimatedUsers int64                                             `json:"estimatedUsers"`
	JSON           radarEntityAsnGetByIPResponseResultAsnRelatedJSON `json:"-"`
}

// radarEntityAsnGetByIPResponseResultAsnRelatedJSON contains the JSON metadata for
// the struct [RadarEntityAsnGetByIPResponseResultAsnRelated]
type radarEntityAsnGetByIPResponseResultAsnRelatedJSON struct {
	Asn            apijson.Field
	Name           apijson.Field
	Aka            apijson.Field
	EstimatedUsers apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarEntityAsnGetByIPResponseResultAsnRelated) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEntityAsnGetByIPResponseResultAsnRelatedJSON) RawJSON() string {
	return r.raw
}

type RadarEntityAsnGetParams struct {
	// Format in which results will be returned.
	Format param.Field[RadarEntityAsnGetParamsFormat] `query:"format"`
}

// URLQuery serializes [RadarEntityAsnGetParams]'s query parameters as
// `url.Values`.
func (r RadarEntityAsnGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Format in which results will be returned.
type RadarEntityAsnGetParamsFormat string

const (
	RadarEntityAsnGetParamsFormatJson RadarEntityAsnGetParamsFormat = "JSON"
	RadarEntityAsnGetParamsFormatCsv  RadarEntityAsnGetParamsFormat = "CSV"
)

func (r RadarEntityAsnGetParamsFormat) IsKnown() bool {
	switch r {
	case RadarEntityAsnGetParamsFormatJson, RadarEntityAsnGetParamsFormatCsv:
		return true
	}
	return false
}

type RadarEntityAsnListParams struct {
	// Comma-separated list of Autonomous System Numbers (ASNs).
	Asn param.Field[string] `query:"asn"`
	// Format in which results will be returned.
	Format param.Field[RadarEntityAsnListParamsFormat] `query:"format"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Location alpha-2 code.
	Location param.Field[string] `query:"location"`
	// Skips the specified number of objects before fetching the results.
	Offset param.Field[int64] `query:"offset"`
	// Metric to order the ASNs by.
	OrderBy param.Field[RadarEntityAsnListParamsOrderBy] `query:"orderBy"`
}

// URLQuery serializes [RadarEntityAsnListParams]'s query parameters as
// `url.Values`.
func (r RadarEntityAsnListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Format in which results will be returned.
type RadarEntityAsnListParamsFormat string

const (
	RadarEntityAsnListParamsFormatJson RadarEntityAsnListParamsFormat = "JSON"
	RadarEntityAsnListParamsFormatCsv  RadarEntityAsnListParamsFormat = "CSV"
)

func (r RadarEntityAsnListParamsFormat) IsKnown() bool {
	switch r {
	case RadarEntityAsnListParamsFormatJson, RadarEntityAsnListParamsFormatCsv:
		return true
	}
	return false
}

// Metric to order the ASNs by.
type RadarEntityAsnListParamsOrderBy string

const (
	RadarEntityAsnListParamsOrderByAsn        RadarEntityAsnListParamsOrderBy = "ASN"
	RadarEntityAsnListParamsOrderByPopulation RadarEntityAsnListParamsOrderBy = "POPULATION"
)

func (r RadarEntityAsnListParamsOrderBy) IsKnown() bool {
	switch r {
	case RadarEntityAsnListParamsOrderByAsn, RadarEntityAsnListParamsOrderByPopulation:
		return true
	}
	return false
}

type RadarEntityAsnGetRelationshipsParams struct {
	// Retrieves the AS relationship of ASN2 with respect to the given ASN.
	Asn2 param.Field[int64] `query:"asn2"`
	// Format in which results will be returned.
	Format param.Field[RadarEntityAsnGetRelationshipsParamsFormat] `query:"format"`
}

// URLQuery serializes [RadarEntityAsnGetRelationshipsParams]'s query parameters as
// `url.Values`.
func (r RadarEntityAsnGetRelationshipsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Format in which results will be returned.
type RadarEntityAsnGetRelationshipsParamsFormat string

const (
	RadarEntityAsnGetRelationshipsParamsFormatJson RadarEntityAsnGetRelationshipsParamsFormat = "JSON"
	RadarEntityAsnGetRelationshipsParamsFormatCsv  RadarEntityAsnGetRelationshipsParamsFormat = "CSV"
)

func (r RadarEntityAsnGetRelationshipsParamsFormat) IsKnown() bool {
	switch r {
	case RadarEntityAsnGetRelationshipsParamsFormatJson, RadarEntityAsnGetRelationshipsParamsFormatCsv:
		return true
	}
	return false
}

type RadarEntityAsnGetByIPParams struct {
	// IP address.
	IP param.Field[string] `query:"ip,required" format:"ip"`
	// Format in which results will be returned.
	Format param.Field[RadarEntityAsnGetByIPParamsFormat] `query:"format"`
}

// URLQuery serializes [RadarEntityAsnGetByIPParams]'s query parameters as
// `url.Values`.
func (r RadarEntityAsnGetByIPParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Format in which results will be returned.
type RadarEntityAsnGetByIPParamsFormat string

const (
	RadarEntityAsnGetByIPParamsFormatJson RadarEntityAsnGetByIPParamsFormat = "JSON"
	RadarEntityAsnGetByIPParamsFormatCsv  RadarEntityAsnGetByIPParamsFormat = "CSV"
)

func (r RadarEntityAsnGetByIPParamsFormat) IsKnown() bool {
	switch r {
	case RadarEntityAsnGetByIPParamsFormatJson, RadarEntityAsnGetByIPParamsFormatCsv:
		return true
	}
	return false
}
