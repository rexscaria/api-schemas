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

// RadarAnnotationService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRadarAnnotationService] method instead.
type RadarAnnotationService struct {
	Options []option.RequestOption
	Outages *RadarAnnotationOutageService
}

// NewRadarAnnotationService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRadarAnnotationService(opts ...option.RequestOption) (r *RadarAnnotationService) {
	r = &RadarAnnotationService{}
	r.Options = opts
	r.Outages = NewRadarAnnotationOutageService(opts...)
	return
}

// Retrieves the latest annotations.
func (r *RadarAnnotationService) GetLatest(ctx context.Context, query RadarAnnotationGetLatestParams, opts ...option.RequestOption) (res *RadarAnnotationGetLatestResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/annotations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarAnnotationGetLatestResponse struct {
	Result  RadarAnnotationGetLatestResponseResult `json:"result,required"`
	Success bool                                   `json:"success,required"`
	JSON    radarAnnotationGetLatestResponseJSON   `json:"-"`
}

// radarAnnotationGetLatestResponseJSON contains the JSON metadata for the struct
// [RadarAnnotationGetLatestResponse]
type radarAnnotationGetLatestResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAnnotationGetLatestResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAnnotationGetLatestResponseJSON) RawJSON() string {
	return r.raw
}

type RadarAnnotationGetLatestResponseResult struct {
	Annotations []RadarAnnotationGetLatestResponseResultAnnotation `json:"annotations,required"`
	JSON        radarAnnotationGetLatestResponseResultJSON         `json:"-"`
}

// radarAnnotationGetLatestResponseResultJSON contains the JSON metadata for the
// struct [RadarAnnotationGetLatestResponseResult]
type radarAnnotationGetLatestResponseResultJSON struct {
	Annotations apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAnnotationGetLatestResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAnnotationGetLatestResponseResultJSON) RawJSON() string {
	return r.raw
}

type RadarAnnotationGetLatestResponseResultAnnotation struct {
	ID               string                                                             `json:"id,required"`
	Asns             []int64                                                            `json:"asns,required"`
	AsnsDetails      []RadarAnnotationGetLatestResponseResultAnnotationsAsnsDetail      `json:"asnsDetails,required"`
	DataSource       string                                                             `json:"dataSource,required"`
	EventType        string                                                             `json:"eventType,required"`
	Locations        []string                                                           `json:"locations,required"`
	LocationsDetails []RadarAnnotationGetLatestResponseResultAnnotationsLocationsDetail `json:"locationsDetails,required"`
	Outage           RadarAnnotationGetLatestResponseResultAnnotationsOutage            `json:"outage,required"`
	StartDate        string                                                             `json:"startDate,required"`
	Description      string                                                             `json:"description"`
	EndDate          string                                                             `json:"endDate"`
	LinkedURL        string                                                             `json:"linkedUrl"`
	Scope            string                                                             `json:"scope"`
	JSON             radarAnnotationGetLatestResponseResultAnnotationJSON               `json:"-"`
}

// radarAnnotationGetLatestResponseResultAnnotationJSON contains the JSON metadata
// for the struct [RadarAnnotationGetLatestResponseResultAnnotation]
type radarAnnotationGetLatestResponseResultAnnotationJSON struct {
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

func (r *RadarAnnotationGetLatestResponseResultAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAnnotationGetLatestResponseResultAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarAnnotationGetLatestResponseResultAnnotationsAsnsDetail struct {
	Asn       string                                                                `json:"asn,required"`
	Name      string                                                                `json:"name,required"`
	Locations RadarAnnotationGetLatestResponseResultAnnotationsAsnsDetailsLocations `json:"locations"`
	JSON      radarAnnotationGetLatestResponseResultAnnotationsAsnsDetailJSON       `json:"-"`
}

// radarAnnotationGetLatestResponseResultAnnotationsAsnsDetailJSON contains the
// JSON metadata for the struct
// [RadarAnnotationGetLatestResponseResultAnnotationsAsnsDetail]
type radarAnnotationGetLatestResponseResultAnnotationsAsnsDetailJSON struct {
	Asn         apijson.Field
	Name        apijson.Field
	Locations   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAnnotationGetLatestResponseResultAnnotationsAsnsDetail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAnnotationGetLatestResponseResultAnnotationsAsnsDetailJSON) RawJSON() string {
	return r.raw
}

type RadarAnnotationGetLatestResponseResultAnnotationsAsnsDetailsLocations struct {
	Code string                                                                    `json:"code,required"`
	Name string                                                                    `json:"name,required"`
	JSON radarAnnotationGetLatestResponseResultAnnotationsAsnsDetailsLocationsJSON `json:"-"`
}

// radarAnnotationGetLatestResponseResultAnnotationsAsnsDetailsLocationsJSON
// contains the JSON metadata for the struct
// [RadarAnnotationGetLatestResponseResultAnnotationsAsnsDetailsLocations]
type radarAnnotationGetLatestResponseResultAnnotationsAsnsDetailsLocationsJSON struct {
	Code        apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAnnotationGetLatestResponseResultAnnotationsAsnsDetailsLocations) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAnnotationGetLatestResponseResultAnnotationsAsnsDetailsLocationsJSON) RawJSON() string {
	return r.raw
}

type RadarAnnotationGetLatestResponseResultAnnotationsLocationsDetail struct {
	Code string                                                               `json:"code,required"`
	Name string                                                               `json:"name,required"`
	JSON radarAnnotationGetLatestResponseResultAnnotationsLocationsDetailJSON `json:"-"`
}

// radarAnnotationGetLatestResponseResultAnnotationsLocationsDetailJSON contains
// the JSON metadata for the struct
// [RadarAnnotationGetLatestResponseResultAnnotationsLocationsDetail]
type radarAnnotationGetLatestResponseResultAnnotationsLocationsDetailJSON struct {
	Code        apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAnnotationGetLatestResponseResultAnnotationsLocationsDetail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAnnotationGetLatestResponseResultAnnotationsLocationsDetailJSON) RawJSON() string {
	return r.raw
}

type RadarAnnotationGetLatestResponseResultAnnotationsOutage struct {
	OutageCause string                                                      `json:"outageCause,required"`
	OutageType  string                                                      `json:"outageType,required"`
	JSON        radarAnnotationGetLatestResponseResultAnnotationsOutageJSON `json:"-"`
}

// radarAnnotationGetLatestResponseResultAnnotationsOutageJSON contains the JSON
// metadata for the struct
// [RadarAnnotationGetLatestResponseResultAnnotationsOutage]
type radarAnnotationGetLatestResponseResultAnnotationsOutageJSON struct {
	OutageCause apijson.Field
	OutageType  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAnnotationGetLatestResponseResultAnnotationsOutage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAnnotationGetLatestResponseResultAnnotationsOutageJSON) RawJSON() string {
	return r.raw
}

type RadarAnnotationGetLatestParams struct {
	// Single Autonomous System Number (ASN) as integer.
	Asn param.Field[int64] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[time.Time] `query:"dateEnd" format:"date-time"`
	// Shorthand date ranges for the last X days - use when you don't need specific
	// start and end dates.
	DateRange param.Field[string] `query:"dateRange"`
	// Start of the date range (inclusive).
	DateStart param.Field[time.Time] `query:"dateStart" format:"date-time"`
	// Format in which results will be returned.
	Format param.Field[RadarAnnotationGetLatestParamsFormat] `query:"format"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Location alpha-2 code.
	Location param.Field[string] `query:"location"`
	// Skips the specified number of objects before fetching the results.
	Offset param.Field[int64] `query:"offset"`
}

// URLQuery serializes [RadarAnnotationGetLatestParams]'s query parameters as
// `url.Values`.
func (r RadarAnnotationGetLatestParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Format in which results will be returned.
type RadarAnnotationGetLatestParamsFormat string

const (
	RadarAnnotationGetLatestParamsFormatJson RadarAnnotationGetLatestParamsFormat = "JSON"
	RadarAnnotationGetLatestParamsFormatCsv  RadarAnnotationGetLatestParamsFormat = "CSV"
)

func (r RadarAnnotationGetLatestParamsFormat) IsKnown() bool {
	switch r {
	case RadarAnnotationGetLatestParamsFormatJson, RadarAnnotationGetLatestParamsFormatCsv:
		return true
	}
	return false
}
