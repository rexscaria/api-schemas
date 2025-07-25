// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/apiquery"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
)

// RadarAnnotationOutageService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRadarAnnotationOutageService] method instead.
type RadarAnnotationOutageService struct {
	Options []option.RequestOption
}

// NewRadarAnnotationOutageService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRadarAnnotationOutageService(opts ...option.RequestOption) (r *RadarAnnotationOutageService) {
	r = &RadarAnnotationOutageService{}
	r.Options = opts
	return
}

// Retrieves the number of outages by location.
func (r *RadarAnnotationOutageService) GetByLocation(ctx context.Context, query RadarAnnotationOutageGetByLocationParams, opts ...option.RequestOption) (res *RadarAnnotationOutageGetByLocationResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/annotations/outages/locations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves the latest Internet outages and anomalies.
func (r *RadarAnnotationOutageService) GetLatest(ctx context.Context, query RadarAnnotationOutageGetLatestParams, opts ...option.RequestOption) (res *RadarAnnotationOutageGetLatestResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/annotations/outages"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarAnnotationOutageGetByLocationResponse struct {
	Result  RadarAnnotationOutageGetByLocationResponseResult `json:"result,required"`
	Success bool                                             `json:"success,required"`
	JSON    radarAnnotationOutageGetByLocationResponseJSON   `json:"-"`
}

// radarAnnotationOutageGetByLocationResponseJSON contains the JSON metadata for
// the struct [RadarAnnotationOutageGetByLocationResponse]
type radarAnnotationOutageGetByLocationResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAnnotationOutageGetByLocationResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAnnotationOutageGetByLocationResponseJSON) RawJSON() string {
	return r.raw
}

type RadarAnnotationOutageGetByLocationResponseResult struct {
	Annotations []RadarAnnotationOutageGetByLocationResponseResultAnnotation `json:"annotations,required"`
	JSON        radarAnnotationOutageGetByLocationResponseResultJSON         `json:"-"`
}

// radarAnnotationOutageGetByLocationResponseResultJSON contains the JSON metadata
// for the struct [RadarAnnotationOutageGetByLocationResponseResult]
type radarAnnotationOutageGetByLocationResponseResultJSON struct {
	Annotations apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAnnotationOutageGetByLocationResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAnnotationOutageGetByLocationResponseResultJSON) RawJSON() string {
	return r.raw
}

type RadarAnnotationOutageGetByLocationResponseResultAnnotation struct {
	ClientCountryAlpha2 string `json:"clientCountryAlpha2,required"`
	ClientCountryName   string `json:"clientCountryName,required"`
	// A numeric string.
	Value string                                                         `json:"value,required"`
	JSON  radarAnnotationOutageGetByLocationResponseResultAnnotationJSON `json:"-"`
}

// radarAnnotationOutageGetByLocationResponseResultAnnotationJSON contains the JSON
// metadata for the struct
// [RadarAnnotationOutageGetByLocationResponseResultAnnotation]
type radarAnnotationOutageGetByLocationResponseResultAnnotationJSON struct {
	ClientCountryAlpha2 apijson.Field
	ClientCountryName   apijson.Field
	Value               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *RadarAnnotationOutageGetByLocationResponseResultAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAnnotationOutageGetByLocationResponseResultAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarAnnotationOutageGetLatestResponse struct {
	Result  RadarAnnotationOutageGetLatestResponseResult `json:"result,required"`
	Success bool                                         `json:"success,required"`
	JSON    radarAnnotationOutageGetLatestResponseJSON   `json:"-"`
}

// radarAnnotationOutageGetLatestResponseJSON contains the JSON metadata for the
// struct [RadarAnnotationOutageGetLatestResponse]
type radarAnnotationOutageGetLatestResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAnnotationOutageGetLatestResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAnnotationOutageGetLatestResponseJSON) RawJSON() string {
	return r.raw
}

type RadarAnnotationOutageGetLatestResponseResult struct {
	Annotations []RadarAnnotationOutageGetLatestResponseResultAnnotation `json:"annotations,required"`
	JSON        radarAnnotationOutageGetLatestResponseResultJSON         `json:"-"`
}

// radarAnnotationOutageGetLatestResponseResultJSON contains the JSON metadata for
// the struct [RadarAnnotationOutageGetLatestResponseResult]
type radarAnnotationOutageGetLatestResponseResultJSON struct {
	Annotations apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAnnotationOutageGetLatestResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAnnotationOutageGetLatestResponseResultJSON) RawJSON() string {
	return r.raw
}

type RadarAnnotationOutageGetLatestResponseResultAnnotation struct {
	ID               string                                                                   `json:"id,required"`
	Asns             []int64                                                                  `json:"asns,required"`
	AsnsDetails      []RadarAnnotationOutageGetLatestResponseResultAnnotationsAsnsDetail      `json:"asnsDetails,required"`
	DataSource       string                                                                   `json:"dataSource,required"`
	EventType        string                                                                   `json:"eventType,required"`
	Locations        []string                                                                 `json:"locations,required"`
	LocationsDetails []RadarAnnotationOutageGetLatestResponseResultAnnotationsLocationsDetail `json:"locationsDetails,required"`
	Outage           RadarAnnotationOutageGetLatestResponseResultAnnotationsOutage            `json:"outage,required"`
	StartDate        time.Time                                                                `json:"startDate,required" format:"date-time"`
	Description      string                                                                   `json:"description"`
	EndDate          time.Time                                                                `json:"endDate" format:"date-time"`
	LinkedURL        string                                                                   `json:"linkedUrl"`
	Scope            string                                                                   `json:"scope"`
	JSON             radarAnnotationOutageGetLatestResponseResultAnnotationJSON               `json:"-"`
}

// radarAnnotationOutageGetLatestResponseResultAnnotationJSON contains the JSON
// metadata for the struct [RadarAnnotationOutageGetLatestResponseResultAnnotation]
type radarAnnotationOutageGetLatestResponseResultAnnotationJSON struct {
	ID               apijson.Field
	Asns             apijson.Field
	AsnsDetails      apijson.Field
	DataSource       apijson.Field
	EventType        apijson.Field
	Locations        apijson.Field
	LocationsDetails apijson.Field
	Outage           apijson.Field
	StartDate        apijson.Field
	Description      apijson.Field
	EndDate          apijson.Field
	LinkedURL        apijson.Field
	Scope            apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RadarAnnotationOutageGetLatestResponseResultAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAnnotationOutageGetLatestResponseResultAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarAnnotationOutageGetLatestResponseResultAnnotationsAsnsDetail struct {
	Asn       string                                                                      `json:"asn,required"`
	Name      string                                                                      `json:"name,required"`
	Locations RadarAnnotationOutageGetLatestResponseResultAnnotationsAsnsDetailsLocations `json:"locations"`
	JSON      radarAnnotationOutageGetLatestResponseResultAnnotationsAsnsDetailJSON       `json:"-"`
}

// radarAnnotationOutageGetLatestResponseResultAnnotationsAsnsDetailJSON contains
// the JSON metadata for the struct
// [RadarAnnotationOutageGetLatestResponseResultAnnotationsAsnsDetail]
type radarAnnotationOutageGetLatestResponseResultAnnotationsAsnsDetailJSON struct {
	Asn         apijson.Field
	Name        apijson.Field
	Locations   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAnnotationOutageGetLatestResponseResultAnnotationsAsnsDetail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAnnotationOutageGetLatestResponseResultAnnotationsAsnsDetailJSON) RawJSON() string {
	return r.raw
}

type RadarAnnotationOutageGetLatestResponseResultAnnotationsAsnsDetailsLocations struct {
	Code string                                                                          `json:"code,required"`
	Name string                                                                          `json:"name,required"`
	JSON radarAnnotationOutageGetLatestResponseResultAnnotationsAsnsDetailsLocationsJSON `json:"-"`
}

// radarAnnotationOutageGetLatestResponseResultAnnotationsAsnsDetailsLocationsJSON
// contains the JSON metadata for the struct
// [RadarAnnotationOutageGetLatestResponseResultAnnotationsAsnsDetailsLocations]
type radarAnnotationOutageGetLatestResponseResultAnnotationsAsnsDetailsLocationsJSON struct {
	Code        apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAnnotationOutageGetLatestResponseResultAnnotationsAsnsDetailsLocations) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAnnotationOutageGetLatestResponseResultAnnotationsAsnsDetailsLocationsJSON) RawJSON() string {
	return r.raw
}

type RadarAnnotationOutageGetLatestResponseResultAnnotationsLocationsDetail struct {
	Code string                                                                     `json:"code,required"`
	Name string                                                                     `json:"name,required"`
	JSON radarAnnotationOutageGetLatestResponseResultAnnotationsLocationsDetailJSON `json:"-"`
}

// radarAnnotationOutageGetLatestResponseResultAnnotationsLocationsDetailJSON
// contains the JSON metadata for the struct
// [RadarAnnotationOutageGetLatestResponseResultAnnotationsLocationsDetail]
type radarAnnotationOutageGetLatestResponseResultAnnotationsLocationsDetailJSON struct {
	Code        apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAnnotationOutageGetLatestResponseResultAnnotationsLocationsDetail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAnnotationOutageGetLatestResponseResultAnnotationsLocationsDetailJSON) RawJSON() string {
	return r.raw
}

type RadarAnnotationOutageGetLatestResponseResultAnnotationsOutage struct {
	OutageCause string                                                            `json:"outageCause,required"`
	OutageType  string                                                            `json:"outageType,required"`
	JSON        radarAnnotationOutageGetLatestResponseResultAnnotationsOutageJSON `json:"-"`
}

// radarAnnotationOutageGetLatestResponseResultAnnotationsOutageJSON contains the
// JSON metadata for the struct
// [RadarAnnotationOutageGetLatestResponseResultAnnotationsOutage]
type radarAnnotationOutageGetLatestResponseResultAnnotationsOutageJSON struct {
	OutageCause apijson.Field
	OutageType  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAnnotationOutageGetLatestResponseResultAnnotationsOutage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAnnotationOutageGetLatestResponseResultAnnotationsOutageJSON) RawJSON() string {
	return r.raw
}

type RadarAnnotationOutageGetByLocationParams struct {
	// End of the date range (inclusive).
	DateEnd param.Field[time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by date range.
	DateRange param.Field[string] `query:"dateRange"`
	// Start of the date range (inclusive).
	DateStart param.Field[time.Time] `query:"dateStart" format:"date-time"`
	// Format in which results will be returned.
	Format param.Field[RadarAnnotationOutageGetByLocationParamsFormat] `query:"format"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
}

// URLQuery serializes [RadarAnnotationOutageGetByLocationParams]'s query
// parameters as `url.Values`.
func (r RadarAnnotationOutageGetByLocationParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Format in which results will be returned.
type RadarAnnotationOutageGetByLocationParamsFormat string

const (
	RadarAnnotationOutageGetByLocationParamsFormatJson RadarAnnotationOutageGetByLocationParamsFormat = "JSON"
	RadarAnnotationOutageGetByLocationParamsFormatCsv  RadarAnnotationOutageGetByLocationParamsFormat = "CSV"
)

func (r RadarAnnotationOutageGetByLocationParamsFormat) IsKnown() bool {
	switch r {
	case RadarAnnotationOutageGetByLocationParamsFormatJson, RadarAnnotationOutageGetByLocationParamsFormatCsv:
		return true
	}
	return false
}

type RadarAnnotationOutageGetLatestParams struct {
	// Filters results by Autonomous System. Specify a single Autonomous System Number
	// (ASN) as integer.
	Asn param.Field[int64] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by date range.
	DateRange param.Field[string] `query:"dateRange"`
	// Start of the date range (inclusive).
	DateStart param.Field[time.Time] `query:"dateStart" format:"date-time"`
	// Format in which results will be returned.
	Format param.Field[RadarAnnotationOutageGetLatestParamsFormat] `query:"format"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Filters results by location. Specify an alpha-2 location code.
	Location param.Field[string] `query:"location"`
	// Skips the specified number of objects before fetching the results.
	Offset param.Field[int64] `query:"offset"`
}

// URLQuery serializes [RadarAnnotationOutageGetLatestParams]'s query parameters as
// `url.Values`.
func (r RadarAnnotationOutageGetLatestParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Format in which results will be returned.
type RadarAnnotationOutageGetLatestParamsFormat string

const (
	RadarAnnotationOutageGetLatestParamsFormatJson RadarAnnotationOutageGetLatestParamsFormat = "JSON"
	RadarAnnotationOutageGetLatestParamsFormatCsv  RadarAnnotationOutageGetLatestParamsFormat = "CSV"
)

func (r RadarAnnotationOutageGetLatestParamsFormat) IsKnown() bool {
	switch r {
	case RadarAnnotationOutageGetLatestParamsFormatJson, RadarAnnotationOutageGetLatestParamsFormatCsv:
		return true
	}
	return false
}
