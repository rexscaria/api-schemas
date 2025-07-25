// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
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

// RadarAs112TopLocationService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRadarAs112TopLocationService] method instead.
type RadarAs112TopLocationService struct {
	Options []option.RequestOption
}

// NewRadarAs112TopLocationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRadarAs112TopLocationService(opts ...option.RequestOption) (r *RadarAs112TopLocationService) {
	r = &RadarAs112TopLocationService{}
	r.Options = opts
	return
}

// Retrieves the top locations by AS112 DNS queries.
func (r *RadarAs112TopLocationService) Get(ctx context.Context, query RadarAs112TopLocationGetParams, opts ...option.RequestOption) (res *RadarAs112TopLocationGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/as112/top/locations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves the top locations of DNS queries to AS112 with DNSSEC (DNS Security
// Extensions) support.
func (r *RadarAs112TopLocationService) GetByDnssec(ctx context.Context, dnssec RadarAs112TopLocationGetByDnssecParamsDnssec, query RadarAs112TopLocationGetByDnssecParams, opts ...option.RequestOption) (res *RadarAs112TopLocationGetByDnssecResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("radar/as112/top/locations/dnssec/%v", dnssec)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves the top locations of DNS queries to AS112 with EDNS (Extension
// Mechanisms for DNS) support.
func (r *RadarAs112TopLocationService) GetByEdns(ctx context.Context, edns RadarAs112TopLocationGetByEdnsParamsEdns, query RadarAs112TopLocationGetByEdnsParams, opts ...option.RequestOption) (res *RadarAs112TopLocationGetByEdnsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("radar/as112/top/locations/edns/%v", edns)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves the top locations of DNS queries to AS112 for an IP version.
func (r *RadarAs112TopLocationService) GetByIPVersion(ctx context.Context, ipVersion RadarAs112TopLocationGetByIPVersionParamsIPVersion, query RadarAs112TopLocationGetByIPVersionParams, opts ...option.RequestOption) (res *RadarAs112TopLocationGetByIPVersionResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("radar/as112/top/locations/ip_version/%v", ipVersion)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarAs112TopLocationGetResponse struct {
	Result  RadarAs112TopLocationGetResponseResult `json:"result,required"`
	Success bool                                   `json:"success,required"`
	JSON    radarAs112TopLocationGetResponseJSON   `json:"-"`
}

// radarAs112TopLocationGetResponseJSON contains the JSON metadata for the struct
// [RadarAs112TopLocationGetResponse]
type radarAs112TopLocationGetResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TopLocationGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAs112TopLocationGetResponseJSON) RawJSON() string {
	return r.raw
}

type RadarAs112TopLocationGetResponseResult struct {
	// Metadata for the results.
	Meta RadarAs112TopLocationGetResponseResultMeta   `json:"meta,required"`
	Top0 []RadarAs112TopLocationGetResponseResultTop0 `json:"top_0,required"`
	JSON radarAs112TopLocationGetResponseResultJSON   `json:"-"`
}

// radarAs112TopLocationGetResponseResultJSON contains the JSON metadata for the
// struct [RadarAs112TopLocationGetResponseResult]
type radarAs112TopLocationGetResponseResultJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TopLocationGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAs112TopLocationGetResponseResultJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type RadarAs112TopLocationGetResponseResultMeta struct {
	ConfidenceInfo RadarAs112TopLocationGetResponseResultMetaConfidenceInfo `json:"confidenceInfo,required,nullable"`
	DateRange      []RadarAs112TopLocationGetResponseResultMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization RadarAs112TopLocationGetResponseResultMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []RadarAs112TopLocationGetResponseResultMetaUnit `json:"units,required"`
	JSON  radarAs112TopLocationGetResponseResultMetaJSON   `json:"-"`
}

// radarAs112TopLocationGetResponseResultMetaJSON contains the JSON metadata for
// the struct [RadarAs112TopLocationGetResponseResultMeta]
type radarAs112TopLocationGetResponseResultMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAs112TopLocationGetResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAs112TopLocationGetResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

type RadarAs112TopLocationGetResponseResultMetaConfidenceInfo struct {
	Annotations []RadarAs112TopLocationGetResponseResultMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                        `json:"level,required"`
	JSON  radarAs112TopLocationGetResponseResultMetaConfidenceInfoJSON `json:"-"`
}

// radarAs112TopLocationGetResponseResultMetaConfidenceInfoJSON contains the JSON
// metadata for the struct
// [RadarAs112TopLocationGetResponseResultMetaConfidenceInfo]
type radarAs112TopLocationGetResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TopLocationGetResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAs112TopLocationGetResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type RadarAs112TopLocationGetResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                   `json:"isInstantaneous,required"`
	LinkedURL       string                                                                 `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                              `json:"startDate,required" format:"date-time"`
	JSON            radarAs112TopLocationGetResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAs112TopLocationGetResponseResultMetaConfidenceInfoAnnotationJSON contains
// the JSON metadata for the struct
// [RadarAs112TopLocationGetResponseResultMetaConfidenceInfoAnnotation]
type radarAs112TopLocationGetResponseResultMetaConfidenceInfoAnnotationJSON struct {
	DataSource      apijson.Field
	Description     apijson.Field
	EndDate         apijson.Field
	EventType       apijson.Field
	IsInstantaneous apijson.Field
	LinkedURL       apijson.Field
	StartDate       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *RadarAs112TopLocationGetResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAs112TopLocationGetResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarAs112TopLocationGetResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                               `json:"startTime,required" format:"date-time"`
	JSON      radarAs112TopLocationGetResponseResultMetaDateRangeJSON `json:"-"`
}

// radarAs112TopLocationGetResponseResultMetaDateRangeJSON contains the JSON
// metadata for the struct [RadarAs112TopLocationGetResponseResultMetaDateRange]
type radarAs112TopLocationGetResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TopLocationGetResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAs112TopLocationGetResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarAs112TopLocationGetResponseResultMetaNormalization string

const (
	RadarAs112TopLocationGetResponseResultMetaNormalizationPercentage           RadarAs112TopLocationGetResponseResultMetaNormalization = "PERCENTAGE"
	RadarAs112TopLocationGetResponseResultMetaNormalizationMin0Max              RadarAs112TopLocationGetResponseResultMetaNormalization = "MIN0_MAX"
	RadarAs112TopLocationGetResponseResultMetaNormalizationMinMax               RadarAs112TopLocationGetResponseResultMetaNormalization = "MIN_MAX"
	RadarAs112TopLocationGetResponseResultMetaNormalizationRawValues            RadarAs112TopLocationGetResponseResultMetaNormalization = "RAW_VALUES"
	RadarAs112TopLocationGetResponseResultMetaNormalizationPercentageChange     RadarAs112TopLocationGetResponseResultMetaNormalization = "PERCENTAGE_CHANGE"
	RadarAs112TopLocationGetResponseResultMetaNormalizationRollingAverage       RadarAs112TopLocationGetResponseResultMetaNormalization = "ROLLING_AVERAGE"
	RadarAs112TopLocationGetResponseResultMetaNormalizationOverlappedPercentage RadarAs112TopLocationGetResponseResultMetaNormalization = "OVERLAPPED_PERCENTAGE"
	RadarAs112TopLocationGetResponseResultMetaNormalizationRatio                RadarAs112TopLocationGetResponseResultMetaNormalization = "RATIO"
)

func (r RadarAs112TopLocationGetResponseResultMetaNormalization) IsKnown() bool {
	switch r {
	case RadarAs112TopLocationGetResponseResultMetaNormalizationPercentage, RadarAs112TopLocationGetResponseResultMetaNormalizationMin0Max, RadarAs112TopLocationGetResponseResultMetaNormalizationMinMax, RadarAs112TopLocationGetResponseResultMetaNormalizationRawValues, RadarAs112TopLocationGetResponseResultMetaNormalizationPercentageChange, RadarAs112TopLocationGetResponseResultMetaNormalizationRollingAverage, RadarAs112TopLocationGetResponseResultMetaNormalizationOverlappedPercentage, RadarAs112TopLocationGetResponseResultMetaNormalizationRatio:
		return true
	}
	return false
}

type RadarAs112TopLocationGetResponseResultMetaUnit struct {
	Name  string                                             `json:"name,required"`
	Value string                                             `json:"value,required"`
	JSON  radarAs112TopLocationGetResponseResultMetaUnitJSON `json:"-"`
}

// radarAs112TopLocationGetResponseResultMetaUnitJSON contains the JSON metadata
// for the struct [RadarAs112TopLocationGetResponseResultMetaUnit]
type radarAs112TopLocationGetResponseResultMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TopLocationGetResponseResultMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAs112TopLocationGetResponseResultMetaUnitJSON) RawJSON() string {
	return r.raw
}

type RadarAs112TopLocationGetResponseResultTop0 struct {
	ClientCountryAlpha2 string `json:"clientCountryAlpha2,required"`
	ClientCountryName   string `json:"clientCountryName,required"`
	// A numeric string.
	Value string                                         `json:"value,required"`
	JSON  radarAs112TopLocationGetResponseResultTop0JSON `json:"-"`
}

// radarAs112TopLocationGetResponseResultTop0JSON contains the JSON metadata for
// the struct [RadarAs112TopLocationGetResponseResultTop0]
type radarAs112TopLocationGetResponseResultTop0JSON struct {
	ClientCountryAlpha2 apijson.Field
	ClientCountryName   apijson.Field
	Value               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *RadarAs112TopLocationGetResponseResultTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAs112TopLocationGetResponseResultTop0JSON) RawJSON() string {
	return r.raw
}

type RadarAs112TopLocationGetByDnssecResponse struct {
	Result  RadarAs112TopLocationGetByDnssecResponseResult `json:"result,required"`
	Success bool                                           `json:"success,required"`
	JSON    radarAs112TopLocationGetByDnssecResponseJSON   `json:"-"`
}

// radarAs112TopLocationGetByDnssecResponseJSON contains the JSON metadata for the
// struct [RadarAs112TopLocationGetByDnssecResponse]
type radarAs112TopLocationGetByDnssecResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TopLocationGetByDnssecResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAs112TopLocationGetByDnssecResponseJSON) RawJSON() string {
	return r.raw
}

type RadarAs112TopLocationGetByDnssecResponseResult struct {
	// Metadata for the results.
	Meta RadarAs112TopLocationGetByDnssecResponseResultMeta   `json:"meta,required"`
	Top0 []RadarAs112TopLocationGetByDnssecResponseResultTop0 `json:"top_0,required"`
	JSON radarAs112TopLocationGetByDnssecResponseResultJSON   `json:"-"`
}

// radarAs112TopLocationGetByDnssecResponseResultJSON contains the JSON metadata
// for the struct [RadarAs112TopLocationGetByDnssecResponseResult]
type radarAs112TopLocationGetByDnssecResponseResultJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TopLocationGetByDnssecResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAs112TopLocationGetByDnssecResponseResultJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type RadarAs112TopLocationGetByDnssecResponseResultMeta struct {
	ConfidenceInfo RadarAs112TopLocationGetByDnssecResponseResultMetaConfidenceInfo `json:"confidenceInfo,required,nullable"`
	DateRange      []RadarAs112TopLocationGetByDnssecResponseResultMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization RadarAs112TopLocationGetByDnssecResponseResultMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []RadarAs112TopLocationGetByDnssecResponseResultMetaUnit `json:"units,required"`
	JSON  radarAs112TopLocationGetByDnssecResponseResultMetaJSON   `json:"-"`
}

// radarAs112TopLocationGetByDnssecResponseResultMetaJSON contains the JSON
// metadata for the struct [RadarAs112TopLocationGetByDnssecResponseResultMeta]
type radarAs112TopLocationGetByDnssecResponseResultMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAs112TopLocationGetByDnssecResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAs112TopLocationGetByDnssecResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

type RadarAs112TopLocationGetByDnssecResponseResultMetaConfidenceInfo struct {
	Annotations []RadarAs112TopLocationGetByDnssecResponseResultMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                                `json:"level,required"`
	JSON  radarAs112TopLocationGetByDnssecResponseResultMetaConfidenceInfoJSON `json:"-"`
}

// radarAs112TopLocationGetByDnssecResponseResultMetaConfidenceInfoJSON contains
// the JSON metadata for the struct
// [RadarAs112TopLocationGetByDnssecResponseResultMetaConfidenceInfo]
type radarAs112TopLocationGetByDnssecResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TopLocationGetByDnssecResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAs112TopLocationGetByDnssecResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type RadarAs112TopLocationGetByDnssecResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                           `json:"isInstantaneous,required"`
	LinkedURL       string                                                                         `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                                      `json:"startDate,required" format:"date-time"`
	JSON            radarAs112TopLocationGetByDnssecResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAs112TopLocationGetByDnssecResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarAs112TopLocationGetByDnssecResponseResultMetaConfidenceInfoAnnotation]
type radarAs112TopLocationGetByDnssecResponseResultMetaConfidenceInfoAnnotationJSON struct {
	DataSource      apijson.Field
	Description     apijson.Field
	EndDate         apijson.Field
	EventType       apijson.Field
	IsInstantaneous apijson.Field
	LinkedURL       apijson.Field
	StartDate       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *RadarAs112TopLocationGetByDnssecResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAs112TopLocationGetByDnssecResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarAs112TopLocationGetByDnssecResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                       `json:"startTime,required" format:"date-time"`
	JSON      radarAs112TopLocationGetByDnssecResponseResultMetaDateRangeJSON `json:"-"`
}

// radarAs112TopLocationGetByDnssecResponseResultMetaDateRangeJSON contains the
// JSON metadata for the struct
// [RadarAs112TopLocationGetByDnssecResponseResultMetaDateRange]
type radarAs112TopLocationGetByDnssecResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TopLocationGetByDnssecResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAs112TopLocationGetByDnssecResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarAs112TopLocationGetByDnssecResponseResultMetaNormalization string

const (
	RadarAs112TopLocationGetByDnssecResponseResultMetaNormalizationPercentage           RadarAs112TopLocationGetByDnssecResponseResultMetaNormalization = "PERCENTAGE"
	RadarAs112TopLocationGetByDnssecResponseResultMetaNormalizationMin0Max              RadarAs112TopLocationGetByDnssecResponseResultMetaNormalization = "MIN0_MAX"
	RadarAs112TopLocationGetByDnssecResponseResultMetaNormalizationMinMax               RadarAs112TopLocationGetByDnssecResponseResultMetaNormalization = "MIN_MAX"
	RadarAs112TopLocationGetByDnssecResponseResultMetaNormalizationRawValues            RadarAs112TopLocationGetByDnssecResponseResultMetaNormalization = "RAW_VALUES"
	RadarAs112TopLocationGetByDnssecResponseResultMetaNormalizationPercentageChange     RadarAs112TopLocationGetByDnssecResponseResultMetaNormalization = "PERCENTAGE_CHANGE"
	RadarAs112TopLocationGetByDnssecResponseResultMetaNormalizationRollingAverage       RadarAs112TopLocationGetByDnssecResponseResultMetaNormalization = "ROLLING_AVERAGE"
	RadarAs112TopLocationGetByDnssecResponseResultMetaNormalizationOverlappedPercentage RadarAs112TopLocationGetByDnssecResponseResultMetaNormalization = "OVERLAPPED_PERCENTAGE"
	RadarAs112TopLocationGetByDnssecResponseResultMetaNormalizationRatio                RadarAs112TopLocationGetByDnssecResponseResultMetaNormalization = "RATIO"
)

func (r RadarAs112TopLocationGetByDnssecResponseResultMetaNormalization) IsKnown() bool {
	switch r {
	case RadarAs112TopLocationGetByDnssecResponseResultMetaNormalizationPercentage, RadarAs112TopLocationGetByDnssecResponseResultMetaNormalizationMin0Max, RadarAs112TopLocationGetByDnssecResponseResultMetaNormalizationMinMax, RadarAs112TopLocationGetByDnssecResponseResultMetaNormalizationRawValues, RadarAs112TopLocationGetByDnssecResponseResultMetaNormalizationPercentageChange, RadarAs112TopLocationGetByDnssecResponseResultMetaNormalizationRollingAverage, RadarAs112TopLocationGetByDnssecResponseResultMetaNormalizationOverlappedPercentage, RadarAs112TopLocationGetByDnssecResponseResultMetaNormalizationRatio:
		return true
	}
	return false
}

type RadarAs112TopLocationGetByDnssecResponseResultMetaUnit struct {
	Name  string                                                     `json:"name,required"`
	Value string                                                     `json:"value,required"`
	JSON  radarAs112TopLocationGetByDnssecResponseResultMetaUnitJSON `json:"-"`
}

// radarAs112TopLocationGetByDnssecResponseResultMetaUnitJSON contains the JSON
// metadata for the struct [RadarAs112TopLocationGetByDnssecResponseResultMetaUnit]
type radarAs112TopLocationGetByDnssecResponseResultMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TopLocationGetByDnssecResponseResultMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAs112TopLocationGetByDnssecResponseResultMetaUnitJSON) RawJSON() string {
	return r.raw
}

type RadarAs112TopLocationGetByDnssecResponseResultTop0 struct {
	ClientCountryAlpha2 string `json:"clientCountryAlpha2,required"`
	ClientCountryName   string `json:"clientCountryName,required"`
	// A numeric string.
	Value string                                                 `json:"value,required"`
	JSON  radarAs112TopLocationGetByDnssecResponseResultTop0JSON `json:"-"`
}

// radarAs112TopLocationGetByDnssecResponseResultTop0JSON contains the JSON
// metadata for the struct [RadarAs112TopLocationGetByDnssecResponseResultTop0]
type radarAs112TopLocationGetByDnssecResponseResultTop0JSON struct {
	ClientCountryAlpha2 apijson.Field
	ClientCountryName   apijson.Field
	Value               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *RadarAs112TopLocationGetByDnssecResponseResultTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAs112TopLocationGetByDnssecResponseResultTop0JSON) RawJSON() string {
	return r.raw
}

type RadarAs112TopLocationGetByEdnsResponse struct {
	Result  RadarAs112TopLocationGetByEdnsResponseResult `json:"result,required"`
	Success bool                                         `json:"success,required"`
	JSON    radarAs112TopLocationGetByEdnsResponseJSON   `json:"-"`
}

// radarAs112TopLocationGetByEdnsResponseJSON contains the JSON metadata for the
// struct [RadarAs112TopLocationGetByEdnsResponse]
type radarAs112TopLocationGetByEdnsResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TopLocationGetByEdnsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAs112TopLocationGetByEdnsResponseJSON) RawJSON() string {
	return r.raw
}

type RadarAs112TopLocationGetByEdnsResponseResult struct {
	// Metadata for the results.
	Meta RadarAs112TopLocationGetByEdnsResponseResultMeta   `json:"meta,required"`
	Top0 []RadarAs112TopLocationGetByEdnsResponseResultTop0 `json:"top_0,required"`
	JSON radarAs112TopLocationGetByEdnsResponseResultJSON   `json:"-"`
}

// radarAs112TopLocationGetByEdnsResponseResultJSON contains the JSON metadata for
// the struct [RadarAs112TopLocationGetByEdnsResponseResult]
type radarAs112TopLocationGetByEdnsResponseResultJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TopLocationGetByEdnsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAs112TopLocationGetByEdnsResponseResultJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type RadarAs112TopLocationGetByEdnsResponseResultMeta struct {
	ConfidenceInfo RadarAs112TopLocationGetByEdnsResponseResultMetaConfidenceInfo `json:"confidenceInfo,required,nullable"`
	DateRange      []RadarAs112TopLocationGetByEdnsResponseResultMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization RadarAs112TopLocationGetByEdnsResponseResultMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []RadarAs112TopLocationGetByEdnsResponseResultMetaUnit `json:"units,required"`
	JSON  radarAs112TopLocationGetByEdnsResponseResultMetaJSON   `json:"-"`
}

// radarAs112TopLocationGetByEdnsResponseResultMetaJSON contains the JSON metadata
// for the struct [RadarAs112TopLocationGetByEdnsResponseResultMeta]
type radarAs112TopLocationGetByEdnsResponseResultMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAs112TopLocationGetByEdnsResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAs112TopLocationGetByEdnsResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

type RadarAs112TopLocationGetByEdnsResponseResultMetaConfidenceInfo struct {
	Annotations []RadarAs112TopLocationGetByEdnsResponseResultMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                              `json:"level,required"`
	JSON  radarAs112TopLocationGetByEdnsResponseResultMetaConfidenceInfoJSON `json:"-"`
}

// radarAs112TopLocationGetByEdnsResponseResultMetaConfidenceInfoJSON contains the
// JSON metadata for the struct
// [RadarAs112TopLocationGetByEdnsResponseResultMetaConfidenceInfo]
type radarAs112TopLocationGetByEdnsResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TopLocationGetByEdnsResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAs112TopLocationGetByEdnsResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type RadarAs112TopLocationGetByEdnsResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                         `json:"isInstantaneous,required"`
	LinkedURL       string                                                                       `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                                    `json:"startDate,required" format:"date-time"`
	JSON            radarAs112TopLocationGetByEdnsResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAs112TopLocationGetByEdnsResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarAs112TopLocationGetByEdnsResponseResultMetaConfidenceInfoAnnotation]
type radarAs112TopLocationGetByEdnsResponseResultMetaConfidenceInfoAnnotationJSON struct {
	DataSource      apijson.Field
	Description     apijson.Field
	EndDate         apijson.Field
	EventType       apijson.Field
	IsInstantaneous apijson.Field
	LinkedURL       apijson.Field
	StartDate       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *RadarAs112TopLocationGetByEdnsResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAs112TopLocationGetByEdnsResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarAs112TopLocationGetByEdnsResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                     `json:"startTime,required" format:"date-time"`
	JSON      radarAs112TopLocationGetByEdnsResponseResultMetaDateRangeJSON `json:"-"`
}

// radarAs112TopLocationGetByEdnsResponseResultMetaDateRangeJSON contains the JSON
// metadata for the struct
// [RadarAs112TopLocationGetByEdnsResponseResultMetaDateRange]
type radarAs112TopLocationGetByEdnsResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TopLocationGetByEdnsResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAs112TopLocationGetByEdnsResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarAs112TopLocationGetByEdnsResponseResultMetaNormalization string

const (
	RadarAs112TopLocationGetByEdnsResponseResultMetaNormalizationPercentage           RadarAs112TopLocationGetByEdnsResponseResultMetaNormalization = "PERCENTAGE"
	RadarAs112TopLocationGetByEdnsResponseResultMetaNormalizationMin0Max              RadarAs112TopLocationGetByEdnsResponseResultMetaNormalization = "MIN0_MAX"
	RadarAs112TopLocationGetByEdnsResponseResultMetaNormalizationMinMax               RadarAs112TopLocationGetByEdnsResponseResultMetaNormalization = "MIN_MAX"
	RadarAs112TopLocationGetByEdnsResponseResultMetaNormalizationRawValues            RadarAs112TopLocationGetByEdnsResponseResultMetaNormalization = "RAW_VALUES"
	RadarAs112TopLocationGetByEdnsResponseResultMetaNormalizationPercentageChange     RadarAs112TopLocationGetByEdnsResponseResultMetaNormalization = "PERCENTAGE_CHANGE"
	RadarAs112TopLocationGetByEdnsResponseResultMetaNormalizationRollingAverage       RadarAs112TopLocationGetByEdnsResponseResultMetaNormalization = "ROLLING_AVERAGE"
	RadarAs112TopLocationGetByEdnsResponseResultMetaNormalizationOverlappedPercentage RadarAs112TopLocationGetByEdnsResponseResultMetaNormalization = "OVERLAPPED_PERCENTAGE"
	RadarAs112TopLocationGetByEdnsResponseResultMetaNormalizationRatio                RadarAs112TopLocationGetByEdnsResponseResultMetaNormalization = "RATIO"
)

func (r RadarAs112TopLocationGetByEdnsResponseResultMetaNormalization) IsKnown() bool {
	switch r {
	case RadarAs112TopLocationGetByEdnsResponseResultMetaNormalizationPercentage, RadarAs112TopLocationGetByEdnsResponseResultMetaNormalizationMin0Max, RadarAs112TopLocationGetByEdnsResponseResultMetaNormalizationMinMax, RadarAs112TopLocationGetByEdnsResponseResultMetaNormalizationRawValues, RadarAs112TopLocationGetByEdnsResponseResultMetaNormalizationPercentageChange, RadarAs112TopLocationGetByEdnsResponseResultMetaNormalizationRollingAverage, RadarAs112TopLocationGetByEdnsResponseResultMetaNormalizationOverlappedPercentage, RadarAs112TopLocationGetByEdnsResponseResultMetaNormalizationRatio:
		return true
	}
	return false
}

type RadarAs112TopLocationGetByEdnsResponseResultMetaUnit struct {
	Name  string                                                   `json:"name,required"`
	Value string                                                   `json:"value,required"`
	JSON  radarAs112TopLocationGetByEdnsResponseResultMetaUnitJSON `json:"-"`
}

// radarAs112TopLocationGetByEdnsResponseResultMetaUnitJSON contains the JSON
// metadata for the struct [RadarAs112TopLocationGetByEdnsResponseResultMetaUnit]
type radarAs112TopLocationGetByEdnsResponseResultMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TopLocationGetByEdnsResponseResultMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAs112TopLocationGetByEdnsResponseResultMetaUnitJSON) RawJSON() string {
	return r.raw
}

type RadarAs112TopLocationGetByEdnsResponseResultTop0 struct {
	ClientCountryAlpha2 string `json:"clientCountryAlpha2,required"`
	ClientCountryName   string `json:"clientCountryName,required"`
	// A numeric string.
	Value string                                               `json:"value,required"`
	JSON  radarAs112TopLocationGetByEdnsResponseResultTop0JSON `json:"-"`
}

// radarAs112TopLocationGetByEdnsResponseResultTop0JSON contains the JSON metadata
// for the struct [RadarAs112TopLocationGetByEdnsResponseResultTop0]
type radarAs112TopLocationGetByEdnsResponseResultTop0JSON struct {
	ClientCountryAlpha2 apijson.Field
	ClientCountryName   apijson.Field
	Value               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *RadarAs112TopLocationGetByEdnsResponseResultTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAs112TopLocationGetByEdnsResponseResultTop0JSON) RawJSON() string {
	return r.raw
}

type RadarAs112TopLocationGetByIPVersionResponse struct {
	Result  RadarAs112TopLocationGetByIPVersionResponseResult `json:"result,required"`
	Success bool                                              `json:"success,required"`
	JSON    radarAs112TopLocationGetByIPVersionResponseJSON   `json:"-"`
}

// radarAs112TopLocationGetByIPVersionResponseJSON contains the JSON metadata for
// the struct [RadarAs112TopLocationGetByIPVersionResponse]
type radarAs112TopLocationGetByIPVersionResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TopLocationGetByIPVersionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAs112TopLocationGetByIPVersionResponseJSON) RawJSON() string {
	return r.raw
}

type RadarAs112TopLocationGetByIPVersionResponseResult struct {
	// Metadata for the results.
	Meta RadarAs112TopLocationGetByIPVersionResponseResultMeta   `json:"meta,required"`
	Top0 []RadarAs112TopLocationGetByIPVersionResponseResultTop0 `json:"top_0,required"`
	JSON radarAs112TopLocationGetByIPVersionResponseResultJSON   `json:"-"`
}

// radarAs112TopLocationGetByIPVersionResponseResultJSON contains the JSON metadata
// for the struct [RadarAs112TopLocationGetByIPVersionResponseResult]
type radarAs112TopLocationGetByIPVersionResponseResultJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TopLocationGetByIPVersionResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAs112TopLocationGetByIPVersionResponseResultJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type RadarAs112TopLocationGetByIPVersionResponseResultMeta struct {
	ConfidenceInfo RadarAs112TopLocationGetByIPVersionResponseResultMetaConfidenceInfo `json:"confidenceInfo,required,nullable"`
	DateRange      []RadarAs112TopLocationGetByIPVersionResponseResultMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization RadarAs112TopLocationGetByIPVersionResponseResultMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []RadarAs112TopLocationGetByIPVersionResponseResultMetaUnit `json:"units,required"`
	JSON  radarAs112TopLocationGetByIPVersionResponseResultMetaJSON   `json:"-"`
}

// radarAs112TopLocationGetByIPVersionResponseResultMetaJSON contains the JSON
// metadata for the struct [RadarAs112TopLocationGetByIPVersionResponseResultMeta]
type radarAs112TopLocationGetByIPVersionResponseResultMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAs112TopLocationGetByIPVersionResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAs112TopLocationGetByIPVersionResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

type RadarAs112TopLocationGetByIPVersionResponseResultMetaConfidenceInfo struct {
	Annotations []RadarAs112TopLocationGetByIPVersionResponseResultMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                                   `json:"level,required"`
	JSON  radarAs112TopLocationGetByIPVersionResponseResultMetaConfidenceInfoJSON `json:"-"`
}

// radarAs112TopLocationGetByIPVersionResponseResultMetaConfidenceInfoJSON contains
// the JSON metadata for the struct
// [RadarAs112TopLocationGetByIPVersionResponseResultMetaConfidenceInfo]
type radarAs112TopLocationGetByIPVersionResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TopLocationGetByIPVersionResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAs112TopLocationGetByIPVersionResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type RadarAs112TopLocationGetByIPVersionResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                              `json:"isInstantaneous,required"`
	LinkedURL       string                                                                            `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                                         `json:"startDate,required" format:"date-time"`
	JSON            radarAs112TopLocationGetByIPVersionResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAs112TopLocationGetByIPVersionResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarAs112TopLocationGetByIPVersionResponseResultMetaConfidenceInfoAnnotation]
type radarAs112TopLocationGetByIPVersionResponseResultMetaConfidenceInfoAnnotationJSON struct {
	DataSource      apijson.Field
	Description     apijson.Field
	EndDate         apijson.Field
	EventType       apijson.Field
	IsInstantaneous apijson.Field
	LinkedURL       apijson.Field
	StartDate       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *RadarAs112TopLocationGetByIPVersionResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAs112TopLocationGetByIPVersionResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarAs112TopLocationGetByIPVersionResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                          `json:"startTime,required" format:"date-time"`
	JSON      radarAs112TopLocationGetByIPVersionResponseResultMetaDateRangeJSON `json:"-"`
}

// radarAs112TopLocationGetByIPVersionResponseResultMetaDateRangeJSON contains the
// JSON metadata for the struct
// [RadarAs112TopLocationGetByIPVersionResponseResultMetaDateRange]
type radarAs112TopLocationGetByIPVersionResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TopLocationGetByIPVersionResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAs112TopLocationGetByIPVersionResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarAs112TopLocationGetByIPVersionResponseResultMetaNormalization string

const (
	RadarAs112TopLocationGetByIPVersionResponseResultMetaNormalizationPercentage           RadarAs112TopLocationGetByIPVersionResponseResultMetaNormalization = "PERCENTAGE"
	RadarAs112TopLocationGetByIPVersionResponseResultMetaNormalizationMin0Max              RadarAs112TopLocationGetByIPVersionResponseResultMetaNormalization = "MIN0_MAX"
	RadarAs112TopLocationGetByIPVersionResponseResultMetaNormalizationMinMax               RadarAs112TopLocationGetByIPVersionResponseResultMetaNormalization = "MIN_MAX"
	RadarAs112TopLocationGetByIPVersionResponseResultMetaNormalizationRawValues            RadarAs112TopLocationGetByIPVersionResponseResultMetaNormalization = "RAW_VALUES"
	RadarAs112TopLocationGetByIPVersionResponseResultMetaNormalizationPercentageChange     RadarAs112TopLocationGetByIPVersionResponseResultMetaNormalization = "PERCENTAGE_CHANGE"
	RadarAs112TopLocationGetByIPVersionResponseResultMetaNormalizationRollingAverage       RadarAs112TopLocationGetByIPVersionResponseResultMetaNormalization = "ROLLING_AVERAGE"
	RadarAs112TopLocationGetByIPVersionResponseResultMetaNormalizationOverlappedPercentage RadarAs112TopLocationGetByIPVersionResponseResultMetaNormalization = "OVERLAPPED_PERCENTAGE"
	RadarAs112TopLocationGetByIPVersionResponseResultMetaNormalizationRatio                RadarAs112TopLocationGetByIPVersionResponseResultMetaNormalization = "RATIO"
)

func (r RadarAs112TopLocationGetByIPVersionResponseResultMetaNormalization) IsKnown() bool {
	switch r {
	case RadarAs112TopLocationGetByIPVersionResponseResultMetaNormalizationPercentage, RadarAs112TopLocationGetByIPVersionResponseResultMetaNormalizationMin0Max, RadarAs112TopLocationGetByIPVersionResponseResultMetaNormalizationMinMax, RadarAs112TopLocationGetByIPVersionResponseResultMetaNormalizationRawValues, RadarAs112TopLocationGetByIPVersionResponseResultMetaNormalizationPercentageChange, RadarAs112TopLocationGetByIPVersionResponseResultMetaNormalizationRollingAverage, RadarAs112TopLocationGetByIPVersionResponseResultMetaNormalizationOverlappedPercentage, RadarAs112TopLocationGetByIPVersionResponseResultMetaNormalizationRatio:
		return true
	}
	return false
}

type RadarAs112TopLocationGetByIPVersionResponseResultMetaUnit struct {
	Name  string                                                        `json:"name,required"`
	Value string                                                        `json:"value,required"`
	JSON  radarAs112TopLocationGetByIPVersionResponseResultMetaUnitJSON `json:"-"`
}

// radarAs112TopLocationGetByIPVersionResponseResultMetaUnitJSON contains the JSON
// metadata for the struct
// [RadarAs112TopLocationGetByIPVersionResponseResultMetaUnit]
type radarAs112TopLocationGetByIPVersionResponseResultMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TopLocationGetByIPVersionResponseResultMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAs112TopLocationGetByIPVersionResponseResultMetaUnitJSON) RawJSON() string {
	return r.raw
}

type RadarAs112TopLocationGetByIPVersionResponseResultTop0 struct {
	ClientCountryAlpha2 string `json:"clientCountryAlpha2,required"`
	ClientCountryName   string `json:"clientCountryName,required"`
	// A numeric string.
	Value string                                                    `json:"value,required"`
	JSON  radarAs112TopLocationGetByIPVersionResponseResultTop0JSON `json:"-"`
}

// radarAs112TopLocationGetByIPVersionResponseResultTop0JSON contains the JSON
// metadata for the struct [RadarAs112TopLocationGetByIPVersionResponseResultTop0]
type radarAs112TopLocationGetByIPVersionResponseResultTop0JSON struct {
	ClientCountryAlpha2 apijson.Field
	ClientCountryName   apijson.Field
	Value               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *RadarAs112TopLocationGetByIPVersionResponseResultTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAs112TopLocationGetByIPVersionResponseResultTop0JSON) RawJSON() string {
	return r.raw
}

type RadarAs112TopLocationGetParams struct {
	// Filters results by continent. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude continents from results. For example, `-EU,NA`
	// excludes results from EU, but includes results from NA.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by date range. For example, use `7d` and `7dcontrol` to compare
	// this week with the previous week. Use this parameter or set specific start and
	// end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format in which results will be returned.
	Format param.Field[RadarAs112TopLocationGetParamsFormat] `query:"format"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarAs112TopLocationGetParams]'s query parameters as
// `url.Values`.
func (r RadarAs112TopLocationGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Format in which results will be returned.
type RadarAs112TopLocationGetParamsFormat string

const (
	RadarAs112TopLocationGetParamsFormatJson RadarAs112TopLocationGetParamsFormat = "JSON"
	RadarAs112TopLocationGetParamsFormatCsv  RadarAs112TopLocationGetParamsFormat = "CSV"
)

func (r RadarAs112TopLocationGetParamsFormat) IsKnown() bool {
	switch r {
	case RadarAs112TopLocationGetParamsFormatJson, RadarAs112TopLocationGetParamsFormatCsv:
		return true
	}
	return false
}

type RadarAs112TopLocationGetByDnssecParams struct {
	// Filters results by continent. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude continents from results. For example, `-EU,NA`
	// excludes results from EU, but includes results from NA.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by date range. For example, use `7d` and `7dcontrol` to compare
	// this week with the previous week. Use this parameter or set specific start and
	// end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format in which results will be returned.
	Format param.Field[RadarAs112TopLocationGetByDnssecParamsFormat] `query:"format"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarAs112TopLocationGetByDnssecParams]'s query parameters
// as `url.Values`.
func (r RadarAs112TopLocationGetByDnssecParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// DNSSEC (DNS Security Extensions) status.
type RadarAs112TopLocationGetByDnssecParamsDnssec string

const (
	RadarAs112TopLocationGetByDnssecParamsDnssecSupported    RadarAs112TopLocationGetByDnssecParamsDnssec = "SUPPORTED"
	RadarAs112TopLocationGetByDnssecParamsDnssecNotSupported RadarAs112TopLocationGetByDnssecParamsDnssec = "NOT_SUPPORTED"
)

func (r RadarAs112TopLocationGetByDnssecParamsDnssec) IsKnown() bool {
	switch r {
	case RadarAs112TopLocationGetByDnssecParamsDnssecSupported, RadarAs112TopLocationGetByDnssecParamsDnssecNotSupported:
		return true
	}
	return false
}

// Format in which results will be returned.
type RadarAs112TopLocationGetByDnssecParamsFormat string

const (
	RadarAs112TopLocationGetByDnssecParamsFormatJson RadarAs112TopLocationGetByDnssecParamsFormat = "JSON"
	RadarAs112TopLocationGetByDnssecParamsFormatCsv  RadarAs112TopLocationGetByDnssecParamsFormat = "CSV"
)

func (r RadarAs112TopLocationGetByDnssecParamsFormat) IsKnown() bool {
	switch r {
	case RadarAs112TopLocationGetByDnssecParamsFormatJson, RadarAs112TopLocationGetByDnssecParamsFormatCsv:
		return true
	}
	return false
}

type RadarAs112TopLocationGetByEdnsParams struct {
	// Filters results by continent. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude continents from results. For example, `-EU,NA`
	// excludes results from EU, but includes results from NA.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by date range. For example, use `7d` and `7dcontrol` to compare
	// this week with the previous week. Use this parameter or set specific start and
	// end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format in which results will be returned.
	Format param.Field[RadarAs112TopLocationGetByEdnsParamsFormat] `query:"format"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarAs112TopLocationGetByEdnsParams]'s query parameters as
// `url.Values`.
func (r RadarAs112TopLocationGetByEdnsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// EDNS (Extension Mechanisms for DNS) status.
type RadarAs112TopLocationGetByEdnsParamsEdns string

const (
	RadarAs112TopLocationGetByEdnsParamsEdnsSupported    RadarAs112TopLocationGetByEdnsParamsEdns = "SUPPORTED"
	RadarAs112TopLocationGetByEdnsParamsEdnsNotSupported RadarAs112TopLocationGetByEdnsParamsEdns = "NOT_SUPPORTED"
)

func (r RadarAs112TopLocationGetByEdnsParamsEdns) IsKnown() bool {
	switch r {
	case RadarAs112TopLocationGetByEdnsParamsEdnsSupported, RadarAs112TopLocationGetByEdnsParamsEdnsNotSupported:
		return true
	}
	return false
}

// Format in which results will be returned.
type RadarAs112TopLocationGetByEdnsParamsFormat string

const (
	RadarAs112TopLocationGetByEdnsParamsFormatJson RadarAs112TopLocationGetByEdnsParamsFormat = "JSON"
	RadarAs112TopLocationGetByEdnsParamsFormatCsv  RadarAs112TopLocationGetByEdnsParamsFormat = "CSV"
)

func (r RadarAs112TopLocationGetByEdnsParamsFormat) IsKnown() bool {
	switch r {
	case RadarAs112TopLocationGetByEdnsParamsFormatJson, RadarAs112TopLocationGetByEdnsParamsFormatCsv:
		return true
	}
	return false
}

type RadarAs112TopLocationGetByIPVersionParams struct {
	// Filters results by continent. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude continents from results. For example, `-EU,NA`
	// excludes results from EU, but includes results from NA.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by date range. For example, use `7d` and `7dcontrol` to compare
	// this week with the previous week. Use this parameter or set specific start and
	// end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format in which results will be returned.
	Format param.Field[RadarAs112TopLocationGetByIPVersionParamsFormat] `query:"format"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarAs112TopLocationGetByIPVersionParams]'s query
// parameters as `url.Values`.
func (r RadarAs112TopLocationGetByIPVersionParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// IP version.
type RadarAs112TopLocationGetByIPVersionParamsIPVersion string

const (
	RadarAs112TopLocationGetByIPVersionParamsIPVersionIPv4 RadarAs112TopLocationGetByIPVersionParamsIPVersion = "IPv4"
	RadarAs112TopLocationGetByIPVersionParamsIPVersionIPv6 RadarAs112TopLocationGetByIPVersionParamsIPVersion = "IPv6"
)

func (r RadarAs112TopLocationGetByIPVersionParamsIPVersion) IsKnown() bool {
	switch r {
	case RadarAs112TopLocationGetByIPVersionParamsIPVersionIPv4, RadarAs112TopLocationGetByIPVersionParamsIPVersionIPv6:
		return true
	}
	return false
}

// Format in which results will be returned.
type RadarAs112TopLocationGetByIPVersionParamsFormat string

const (
	RadarAs112TopLocationGetByIPVersionParamsFormatJson RadarAs112TopLocationGetByIPVersionParamsFormat = "JSON"
	RadarAs112TopLocationGetByIPVersionParamsFormatCsv  RadarAs112TopLocationGetByIPVersionParamsFormat = "CSV"
)

func (r RadarAs112TopLocationGetByIPVersionParamsFormat) IsKnown() bool {
	switch r {
	case RadarAs112TopLocationGetByIPVersionParamsFormatJson, RadarAs112TopLocationGetByIPVersionParamsFormatCsv:
		return true
	}
	return false
}
