// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/stainless-sdks/cf-rex-go/internal/apijson"
	"github.com/stainless-sdks/cf-rex-go/internal/apiquery"
	"github.com/stainless-sdks/cf-rex-go/internal/param"
	"github.com/stainless-sdks/cf-rex-go/internal/requestconfig"
	"github.com/stainless-sdks/cf-rex-go/option"
)

// RadarTrafficAnomalyService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRadarTrafficAnomalyService] method instead.
type RadarTrafficAnomalyService struct {
	Options []option.RequestOption
}

// NewRadarTrafficAnomalyService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRadarTrafficAnomalyService(opts ...option.RequestOption) (r *RadarTrafficAnomalyService) {
	r = &RadarTrafficAnomalyService{}
	r.Options = opts
	return
}

// Retrieves the latest Internet traffic anomalies, which are signals that might
// indicate an outage. These alerts are automatically detected by Radar and
// manually verified by our team.
func (r *RadarTrafficAnomalyService) List(ctx context.Context, query RadarTrafficAnomalyListParams, opts ...option.RequestOption) (res *RadarTrafficAnomalyListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/traffic_anomalies"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves the sum of Internet traffic anomalies, grouped by location. These
// anomalies are signals that might indicate an outage, automatically detected by
// Radar and manually verified by our team.
func (r *RadarTrafficAnomalyService) Locations(ctx context.Context, query RadarTrafficAnomalyLocationsParams, opts ...option.RequestOption) (res *RadarTrafficAnomalyLocationsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/traffic_anomalies/locations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarTrafficAnomalyListResponse struct {
	Result  RadarTrafficAnomalyListResponseResult `json:"result,required"`
	Success bool                                  `json:"success,required"`
	JSON    radarTrafficAnomalyListResponseJSON   `json:"-"`
}

// radarTrafficAnomalyListResponseJSON contains the JSON metadata for the struct
// [RadarTrafficAnomalyListResponse]
type radarTrafficAnomalyListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarTrafficAnomalyListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarTrafficAnomalyListResponseJSON) RawJSON() string {
	return r.raw
}

type RadarTrafficAnomalyListResponseResult struct {
	TrafficAnomalies []RadarTrafficAnomalyListResponseResultTrafficAnomaly `json:"trafficAnomalies,required"`
	JSON             radarTrafficAnomalyListResponseResultJSON             `json:"-"`
}

// radarTrafficAnomalyListResponseResultJSON contains the JSON metadata for the
// struct [RadarTrafficAnomalyListResponseResult]
type radarTrafficAnomalyListResponseResultJSON struct {
	TrafficAnomalies apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RadarTrafficAnomalyListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarTrafficAnomalyListResponseResultJSON) RawJSON() string {
	return r.raw
}

type RadarTrafficAnomalyListResponseResultTrafficAnomaly struct {
	StartDate            string                                                               `json:"startDate,required"`
	Status               string                                                               `json:"status,required"`
	Type                 string                                                               `json:"type,required"`
	Uuid                 string                                                               `json:"uuid,required"`
	AsnDetails           RadarTrafficAnomalyListResponseResultTrafficAnomaliesAsnDetails      `json:"asnDetails"`
	EndDate              string                                                               `json:"endDate"`
	LocationDetails      RadarTrafficAnomalyListResponseResultTrafficAnomaliesLocationDetails `json:"locationDetails"`
	VisibleInDataSources []string                                                             `json:"visibleInDataSources"`
	JSON                 radarTrafficAnomalyListResponseResultTrafficAnomalyJSON              `json:"-"`
}

// radarTrafficAnomalyListResponseResultTrafficAnomalyJSON contains the JSON
// metadata for the struct [RadarTrafficAnomalyListResponseResultTrafficAnomaly]
type radarTrafficAnomalyListResponseResultTrafficAnomalyJSON struct {
	StartDate            apijson.Field
	Status               apijson.Field
	Type                 apijson.Field
	Uuid                 apijson.Field
	AsnDetails           apijson.Field
	EndDate              apijson.Field
	LocationDetails      apijson.Field
	VisibleInDataSources apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *RadarTrafficAnomalyListResponseResultTrafficAnomaly) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarTrafficAnomalyListResponseResultTrafficAnomalyJSON) RawJSON() string {
	return r.raw
}

type RadarTrafficAnomalyListResponseResultTrafficAnomaliesAsnDetails struct {
	Asn       string                                                                   `json:"asn,required"`
	Name      string                                                                   `json:"name,required"`
	Locations RadarTrafficAnomalyListResponseResultTrafficAnomaliesAsnDetailsLocations `json:"locations"`
	JSON      radarTrafficAnomalyListResponseResultTrafficAnomaliesAsnDetailsJSON      `json:"-"`
}

// radarTrafficAnomalyListResponseResultTrafficAnomaliesAsnDetailsJSON contains the
// JSON metadata for the struct
// [RadarTrafficAnomalyListResponseResultTrafficAnomaliesAsnDetails]
type radarTrafficAnomalyListResponseResultTrafficAnomaliesAsnDetailsJSON struct {
	Asn         apijson.Field
	Name        apijson.Field
	Locations   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarTrafficAnomalyListResponseResultTrafficAnomaliesAsnDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarTrafficAnomalyListResponseResultTrafficAnomaliesAsnDetailsJSON) RawJSON() string {
	return r.raw
}

type RadarTrafficAnomalyListResponseResultTrafficAnomaliesAsnDetailsLocations struct {
	Code string                                                                       `json:"code,required"`
	Name string                                                                       `json:"name,required"`
	JSON radarTrafficAnomalyListResponseResultTrafficAnomaliesAsnDetailsLocationsJSON `json:"-"`
}

// radarTrafficAnomalyListResponseResultTrafficAnomaliesAsnDetailsLocationsJSON
// contains the JSON metadata for the struct
// [RadarTrafficAnomalyListResponseResultTrafficAnomaliesAsnDetailsLocations]
type radarTrafficAnomalyListResponseResultTrafficAnomaliesAsnDetailsLocationsJSON struct {
	Code        apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarTrafficAnomalyListResponseResultTrafficAnomaliesAsnDetailsLocations) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarTrafficAnomalyListResponseResultTrafficAnomaliesAsnDetailsLocationsJSON) RawJSON() string {
	return r.raw
}

type RadarTrafficAnomalyListResponseResultTrafficAnomaliesLocationDetails struct {
	Code string                                                                   `json:"code,required"`
	Name string                                                                   `json:"name,required"`
	JSON radarTrafficAnomalyListResponseResultTrafficAnomaliesLocationDetailsJSON `json:"-"`
}

// radarTrafficAnomalyListResponseResultTrafficAnomaliesLocationDetailsJSON
// contains the JSON metadata for the struct
// [RadarTrafficAnomalyListResponseResultTrafficAnomaliesLocationDetails]
type radarTrafficAnomalyListResponseResultTrafficAnomaliesLocationDetailsJSON struct {
	Code        apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarTrafficAnomalyListResponseResultTrafficAnomaliesLocationDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarTrafficAnomalyListResponseResultTrafficAnomaliesLocationDetailsJSON) RawJSON() string {
	return r.raw
}

type RadarTrafficAnomalyLocationsResponse struct {
	Result  RadarTrafficAnomalyLocationsResponseResult `json:"result,required"`
	Success bool                                       `json:"success,required"`
	JSON    radarTrafficAnomalyLocationsResponseJSON   `json:"-"`
}

// radarTrafficAnomalyLocationsResponseJSON contains the JSON metadata for the
// struct [RadarTrafficAnomalyLocationsResponse]
type radarTrafficAnomalyLocationsResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarTrafficAnomalyLocationsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarTrafficAnomalyLocationsResponseJSON) RawJSON() string {
	return r.raw
}

type RadarTrafficAnomalyLocationsResponseResult struct {
	TrafficAnomalies []RadarTrafficAnomalyLocationsResponseResultTrafficAnomaly `json:"trafficAnomalies,required"`
	JSON             radarTrafficAnomalyLocationsResponseResultJSON             `json:"-"`
}

// radarTrafficAnomalyLocationsResponseResultJSON contains the JSON metadata for
// the struct [RadarTrafficAnomalyLocationsResponseResult]
type radarTrafficAnomalyLocationsResponseResultJSON struct {
	TrafficAnomalies apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RadarTrafficAnomalyLocationsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarTrafficAnomalyLocationsResponseResultJSON) RawJSON() string {
	return r.raw
}

type RadarTrafficAnomalyLocationsResponseResultTrafficAnomaly struct {
	ClientCountryAlpha2 string                                                       `json:"clientCountryAlpha2,required"`
	ClientCountryName   string                                                       `json:"clientCountryName,required"`
	Value               string                                                       `json:"value,required"`
	JSON                radarTrafficAnomalyLocationsResponseResultTrafficAnomalyJSON `json:"-"`
}

// radarTrafficAnomalyLocationsResponseResultTrafficAnomalyJSON contains the JSON
// metadata for the struct
// [RadarTrafficAnomalyLocationsResponseResultTrafficAnomaly]
type radarTrafficAnomalyLocationsResponseResultTrafficAnomalyJSON struct {
	ClientCountryAlpha2 apijson.Field
	ClientCountryName   apijson.Field
	Value               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *RadarTrafficAnomalyLocationsResponseResultTrafficAnomaly) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarTrafficAnomalyLocationsResponseResultTrafficAnomalyJSON) RawJSON() string {
	return r.raw
}

type RadarTrafficAnomalyListParams struct {
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
	Format param.Field[RadarTrafficAnomalyListParamsFormat] `query:"format"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Location alpha-2 code.
	Location param.Field[string] `query:"location"`
	// Skips the specified number of objects before fetching the results.
	Offset param.Field[int64]                               `query:"offset"`
	Status param.Field[RadarTrafficAnomalyListParamsStatus] `query:"status"`
}

// URLQuery serializes [RadarTrafficAnomalyListParams]'s query parameters as
// `url.Values`.
func (r RadarTrafficAnomalyListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Format in which results will be returned.
type RadarTrafficAnomalyListParamsFormat string

const (
	RadarTrafficAnomalyListParamsFormatJson RadarTrafficAnomalyListParamsFormat = "JSON"
	RadarTrafficAnomalyListParamsFormatCsv  RadarTrafficAnomalyListParamsFormat = "CSV"
)

func (r RadarTrafficAnomalyListParamsFormat) IsKnown() bool {
	switch r {
	case RadarTrafficAnomalyListParamsFormatJson, RadarTrafficAnomalyListParamsFormatCsv:
		return true
	}
	return false
}

type RadarTrafficAnomalyListParamsStatus string

const (
	RadarTrafficAnomalyListParamsStatusVerified   RadarTrafficAnomalyListParamsStatus = "VERIFIED"
	RadarTrafficAnomalyListParamsStatusUnverified RadarTrafficAnomalyListParamsStatus = "UNVERIFIED"
)

func (r RadarTrafficAnomalyListParamsStatus) IsKnown() bool {
	switch r {
	case RadarTrafficAnomalyListParamsStatusVerified, RadarTrafficAnomalyListParamsStatusUnverified:
		return true
	}
	return false
}

type RadarTrafficAnomalyLocationsParams struct {
	// End of the date range (inclusive).
	DateEnd param.Field[time.Time] `query:"dateEnd" format:"date-time"`
	// Shorthand date ranges for the last X days - use when you don't need specific
	// start and end dates.
	DateRange param.Field[string] `query:"dateRange"`
	// Start of the date range (inclusive).
	DateStart param.Field[time.Time] `query:"dateStart" format:"date-time"`
	// Format in which results will be returned.
	Format param.Field[RadarTrafficAnomalyLocationsParamsFormat] `query:"format"`
	// Limits the number of objects returned in the response.
	Limit  param.Field[int64]                                    `query:"limit"`
	Status param.Field[RadarTrafficAnomalyLocationsParamsStatus] `query:"status"`
}

// URLQuery serializes [RadarTrafficAnomalyLocationsParams]'s query parameters as
// `url.Values`.
func (r RadarTrafficAnomalyLocationsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Format in which results will be returned.
type RadarTrafficAnomalyLocationsParamsFormat string

const (
	RadarTrafficAnomalyLocationsParamsFormatJson RadarTrafficAnomalyLocationsParamsFormat = "JSON"
	RadarTrafficAnomalyLocationsParamsFormatCsv  RadarTrafficAnomalyLocationsParamsFormat = "CSV"
)

func (r RadarTrafficAnomalyLocationsParamsFormat) IsKnown() bool {
	switch r {
	case RadarTrafficAnomalyLocationsParamsFormatJson, RadarTrafficAnomalyLocationsParamsFormatCsv:
		return true
	}
	return false
}

type RadarTrafficAnomalyLocationsParamsStatus string

const (
	RadarTrafficAnomalyLocationsParamsStatusVerified   RadarTrafficAnomalyLocationsParamsStatus = "VERIFIED"
	RadarTrafficAnomalyLocationsParamsStatusUnverified RadarTrafficAnomalyLocationsParamsStatus = "UNVERIFIED"
)

func (r RadarTrafficAnomalyLocationsParamsStatus) IsKnown() bool {
	switch r {
	case RadarTrafficAnomalyLocationsParamsStatusVerified, RadarTrafficAnomalyLocationsParamsStatusUnverified:
		return true
	}
	return false
}
