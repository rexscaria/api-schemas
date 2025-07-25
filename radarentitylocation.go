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

// RadarEntityLocationService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRadarEntityLocationService] method instead.
type RadarEntityLocationService struct {
	Options []option.RequestOption
}

// NewRadarEntityLocationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRadarEntityLocationService(opts ...option.RequestOption) (r *RadarEntityLocationService) {
	r = &RadarEntityLocationService{}
	r.Options = opts
	return
}

// Retrieves the requested location information. (A confidence level below `5`
// indicates a low level of confidence in the traffic data - normally this happens
// because Cloudflare has a small amount of traffic from/to this location).
func (r *RadarEntityLocationService) Get(ctx context.Context, location string, query RadarEntityLocationGetParams, opts ...option.RequestOption) (res *RadarEntityLocationGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if location == "" {
		err = errors.New("missing required location parameter")
		return
	}
	path := fmt.Sprintf("radar/entities/locations/%s", location)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves a list of locations.
func (r *RadarEntityLocationService) List(ctx context.Context, query RadarEntityLocationListParams, opts ...option.RequestOption) (res *RadarEntityLocationListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/entities/locations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarEntityLocationGetResponse struct {
	Result  RadarEntityLocationGetResponseResult `json:"result,required"`
	Success bool                                 `json:"success,required"`
	JSON    radarEntityLocationGetResponseJSON   `json:"-"`
}

// radarEntityLocationGetResponseJSON contains the JSON metadata for the struct
// [RadarEntityLocationGetResponse]
type radarEntityLocationGetResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEntityLocationGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEntityLocationGetResponseJSON) RawJSON() string {
	return r.raw
}

type RadarEntityLocationGetResponseResult struct {
	Location RadarEntityLocationGetResponseResultLocation `json:"location,required"`
	JSON     radarEntityLocationGetResponseResultJSON     `json:"-"`
}

// radarEntityLocationGetResponseResultJSON contains the JSON metadata for the
// struct [RadarEntityLocationGetResponseResult]
type radarEntityLocationGetResponseResultJSON struct {
	Location    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEntityLocationGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEntityLocationGetResponseResultJSON) RawJSON() string {
	return r.raw
}

type RadarEntityLocationGetResponseResultLocation struct {
	Alpha2          string `json:"alpha2,required"`
	ConfidenceLevel int64  `json:"confidenceLevel,required"`
	// A numeric string.
	Latitude string `json:"latitude,required"`
	// A numeric string.
	Longitude string                                           `json:"longitude,required"`
	Name      string                                           `json:"name,required"`
	Region    string                                           `json:"region,required"`
	Subregion string                                           `json:"subregion,required"`
	JSON      radarEntityLocationGetResponseResultLocationJSON `json:"-"`
}

// radarEntityLocationGetResponseResultLocationJSON contains the JSON metadata for
// the struct [RadarEntityLocationGetResponseResultLocation]
type radarEntityLocationGetResponseResultLocationJSON struct {
	Alpha2          apijson.Field
	ConfidenceLevel apijson.Field
	Latitude        apijson.Field
	Longitude       apijson.Field
	Name            apijson.Field
	Region          apijson.Field
	Subregion       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *RadarEntityLocationGetResponseResultLocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEntityLocationGetResponseResultLocationJSON) RawJSON() string {
	return r.raw
}

type RadarEntityLocationListResponse struct {
	Result  RadarEntityLocationListResponseResult `json:"result,required"`
	Success bool                                  `json:"success,required"`
	JSON    radarEntityLocationListResponseJSON   `json:"-"`
}

// radarEntityLocationListResponseJSON contains the JSON metadata for the struct
// [RadarEntityLocationListResponse]
type radarEntityLocationListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEntityLocationListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEntityLocationListResponseJSON) RawJSON() string {
	return r.raw
}

type RadarEntityLocationListResponseResult struct {
	Locations []RadarEntityLocationListResponseResultLocation `json:"locations,required"`
	JSON      radarEntityLocationListResponseResultJSON       `json:"-"`
}

// radarEntityLocationListResponseResultJSON contains the JSON metadata for the
// struct [RadarEntityLocationListResponseResult]
type radarEntityLocationListResponseResultJSON struct {
	Locations   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEntityLocationListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEntityLocationListResponseResultJSON) RawJSON() string {
	return r.raw
}

type RadarEntityLocationListResponseResultLocation struct {
	Alpha2 string `json:"alpha2,required"`
	// A numeric string.
	Latitude string `json:"latitude,required"`
	// A numeric string.
	Longitude string                                            `json:"longitude,required"`
	Name      string                                            `json:"name,required"`
	JSON      radarEntityLocationListResponseResultLocationJSON `json:"-"`
}

// radarEntityLocationListResponseResultLocationJSON contains the JSON metadata for
// the struct [RadarEntityLocationListResponseResultLocation]
type radarEntityLocationListResponseResultLocationJSON struct {
	Alpha2      apijson.Field
	Latitude    apijson.Field
	Longitude   apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEntityLocationListResponseResultLocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEntityLocationListResponseResultLocationJSON) RawJSON() string {
	return r.raw
}

type RadarEntityLocationGetParams struct {
	// Format in which results will be returned.
	Format param.Field[RadarEntityLocationGetParamsFormat] `query:"format"`
}

// URLQuery serializes [RadarEntityLocationGetParams]'s query parameters as
// `url.Values`.
func (r RadarEntityLocationGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Format in which results will be returned.
type RadarEntityLocationGetParamsFormat string

const (
	RadarEntityLocationGetParamsFormatJson RadarEntityLocationGetParamsFormat = "JSON"
	RadarEntityLocationGetParamsFormatCsv  RadarEntityLocationGetParamsFormat = "CSV"
)

func (r RadarEntityLocationGetParamsFormat) IsKnown() bool {
	switch r {
	case RadarEntityLocationGetParamsFormatJson, RadarEntityLocationGetParamsFormatCsv:
		return true
	}
	return false
}

type RadarEntityLocationListParams struct {
	// Format in which results will be returned.
	Format param.Field[RadarEntityLocationListParamsFormat] `query:"format"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Filters results by location. Specify a comma-separated list of alpha-2 location
	// codes.
	Location param.Field[string] `query:"location"`
	// Skips the specified number of objects before fetching the results.
	Offset param.Field[int64] `query:"offset"`
}

// URLQuery serializes [RadarEntityLocationListParams]'s query parameters as
// `url.Values`.
func (r RadarEntityLocationListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Format in which results will be returned.
type RadarEntityLocationListParamsFormat string

const (
	RadarEntityLocationListParamsFormatJson RadarEntityLocationListParamsFormat = "JSON"
	RadarEntityLocationListParamsFormatCsv  RadarEntityLocationListParamsFormat = "CSV"
)

func (r RadarEntityLocationListParamsFormat) IsKnown() bool {
	switch r {
	case RadarEntityLocationListParamsFormatJson, RadarEntityLocationListParamsFormatCsv:
		return true
	}
	return false
}
