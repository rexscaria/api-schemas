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

// RadarEmailRoutingTimeseriesGroupService contains methods and other services that
// help with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRadarEmailRoutingTimeseriesGroupService] method instead.
type RadarEmailRoutingTimeseriesGroupService struct {
	Options []option.RequestOption
}

// NewRadarEmailRoutingTimeseriesGroupService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarEmailRoutingTimeseriesGroupService(opts ...option.RequestOption) (r *RadarEmailRoutingTimeseriesGroupService) {
	r = &RadarEmailRoutingTimeseriesGroupService{}
	r.Options = opts
	return
}

// Retrieves the distribution of emails by ARC (Authenticated Received Chain)
// validation over time.
func (r *RadarEmailRoutingTimeseriesGroupService) GetArc(ctx context.Context, query RadarEmailRoutingTimeseriesGroupGetArcParams, opts ...option.RequestOption) (res *RadarEmailRoutingTimeseriesGroupGetArcResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/email/routing/timeseries_groups/arc"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves the distribution of emails by DKIM (DomainKeys Identified Mail)
// validation over time.
func (r *RadarEmailRoutingTimeseriesGroupService) GetDkim(ctx context.Context, query RadarEmailRoutingTimeseriesGroupGetDkimParams, opts ...option.RequestOption) (res *RadarEmailRoutingTimeseriesGroupGetDkimResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/email/routing/timeseries_groups/dkim"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves the distribution of emails by DMARC (Domain-based Message
// Authentication, Reporting and Conformance) validation over time.
func (r *RadarEmailRoutingTimeseriesGroupService) GetDmarc(ctx context.Context, query RadarEmailRoutingTimeseriesGroupGetDmarcParams, opts ...option.RequestOption) (res *RadarEmailRoutingTimeseriesGroupGetDmarcResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/email/routing/timeseries_groups/dmarc"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves the distribution of emails by encryption status (encrypted vs.
// not-encrypted) over time.
func (r *RadarEmailRoutingTimeseriesGroupService) GetEncrypted(ctx context.Context, query RadarEmailRoutingTimeseriesGroupGetEncryptedParams, opts ...option.RequestOption) (res *RadarEmailRoutingTimeseriesGroupGetEncryptedResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/email/routing/timeseries_groups/encrypted"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves the distribution of emails by IP version over time.
func (r *RadarEmailRoutingTimeseriesGroupService) GetIPVersion(ctx context.Context, query RadarEmailRoutingTimeseriesGroupGetIPVersionParams, opts ...option.RequestOption) (res *RadarEmailRoutingTimeseriesGroupGetIPVersionResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/email/routing/timeseries_groups/ip_version"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves the distribution of emails by SPF (Sender Policy Framework) validation
// over time.
func (r *RadarEmailRoutingTimeseriesGroupService) GetSpf(ctx context.Context, query RadarEmailRoutingTimeseriesGroupGetSpfParams, opts ...option.RequestOption) (res *RadarEmailRoutingTimeseriesGroupGetSpfResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/email/routing/timeseries_groups/spf"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarEmailRoutingTimeseriesGroupGetArcResponse struct {
	Result  RadarEmailRoutingTimeseriesGroupGetArcResponseResult `json:"result,required"`
	Success bool                                                 `json:"success,required"`
	JSON    radarEmailRoutingTimeseriesGroupGetArcResponseJSON   `json:"-"`
}

// radarEmailRoutingTimeseriesGroupGetArcResponseJSON contains the JSON metadata
// for the struct [RadarEmailRoutingTimeseriesGroupGetArcResponse]
type radarEmailRoutingTimeseriesGroupGetArcResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailRoutingTimeseriesGroupGetArcResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailRoutingTimeseriesGroupGetArcResponseJSON) RawJSON() string {
	return r.raw
}

type RadarEmailRoutingTimeseriesGroupGetArcResponseResult struct {
	Meta   interface{}                                                `json:"meta,required"`
	Serie0 RadarEmailRoutingTimeseriesGroupGetArcResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarEmailRoutingTimeseriesGroupGetArcResponseResultJSON   `json:"-"`
}

// radarEmailRoutingTimeseriesGroupGetArcResponseResultJSON contains the JSON
// metadata for the struct [RadarEmailRoutingTimeseriesGroupGetArcResponseResult]
type radarEmailRoutingTimeseriesGroupGetArcResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailRoutingTimeseriesGroupGetArcResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailRoutingTimeseriesGroupGetArcResponseResultJSON) RawJSON() string {
	return r.raw
}

type RadarEmailRoutingTimeseriesGroupGetArcResponseResultSerie0 struct {
	Fail []string                                                       `json:"FAIL,required"`
	None []string                                                       `json:"NONE,required"`
	Pass []string                                                       `json:"PASS,required"`
	JSON radarEmailRoutingTimeseriesGroupGetArcResponseResultSerie0JSON `json:"-"`
}

// radarEmailRoutingTimeseriesGroupGetArcResponseResultSerie0JSON contains the JSON
// metadata for the struct
// [RadarEmailRoutingTimeseriesGroupGetArcResponseResultSerie0]
type radarEmailRoutingTimeseriesGroupGetArcResponseResultSerie0JSON struct {
	Fail        apijson.Field
	None        apijson.Field
	Pass        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailRoutingTimeseriesGroupGetArcResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailRoutingTimeseriesGroupGetArcResponseResultSerie0JSON) RawJSON() string {
	return r.raw
}

type RadarEmailRoutingTimeseriesGroupGetDkimResponse struct {
	Result  RadarEmailRoutingTimeseriesGroupGetDkimResponseResult `json:"result,required"`
	Success bool                                                  `json:"success,required"`
	JSON    radarEmailRoutingTimeseriesGroupGetDkimResponseJSON   `json:"-"`
}

// radarEmailRoutingTimeseriesGroupGetDkimResponseJSON contains the JSON metadata
// for the struct [RadarEmailRoutingTimeseriesGroupGetDkimResponse]
type radarEmailRoutingTimeseriesGroupGetDkimResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailRoutingTimeseriesGroupGetDkimResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailRoutingTimeseriesGroupGetDkimResponseJSON) RawJSON() string {
	return r.raw
}

type RadarEmailRoutingTimeseriesGroupGetDkimResponseResult struct {
	Meta   interface{}                                                 `json:"meta,required"`
	Serie0 RadarEmailRoutingTimeseriesGroupGetDkimResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarEmailRoutingTimeseriesGroupGetDkimResponseResultJSON   `json:"-"`
}

// radarEmailRoutingTimeseriesGroupGetDkimResponseResultJSON contains the JSON
// metadata for the struct [RadarEmailRoutingTimeseriesGroupGetDkimResponseResult]
type radarEmailRoutingTimeseriesGroupGetDkimResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailRoutingTimeseriesGroupGetDkimResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailRoutingTimeseriesGroupGetDkimResponseResultJSON) RawJSON() string {
	return r.raw
}

type RadarEmailRoutingTimeseriesGroupGetDkimResponseResultSerie0 struct {
	Fail []string                                                        `json:"FAIL,required"`
	None []string                                                        `json:"NONE,required"`
	Pass []string                                                        `json:"PASS,required"`
	JSON radarEmailRoutingTimeseriesGroupGetDkimResponseResultSerie0JSON `json:"-"`
}

// radarEmailRoutingTimeseriesGroupGetDkimResponseResultSerie0JSON contains the
// JSON metadata for the struct
// [RadarEmailRoutingTimeseriesGroupGetDkimResponseResultSerie0]
type radarEmailRoutingTimeseriesGroupGetDkimResponseResultSerie0JSON struct {
	Fail        apijson.Field
	None        apijson.Field
	Pass        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailRoutingTimeseriesGroupGetDkimResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailRoutingTimeseriesGroupGetDkimResponseResultSerie0JSON) RawJSON() string {
	return r.raw
}

type RadarEmailRoutingTimeseriesGroupGetDmarcResponse struct {
	Result  RadarEmailRoutingTimeseriesGroupGetDmarcResponseResult `json:"result,required"`
	Success bool                                                   `json:"success,required"`
	JSON    radarEmailRoutingTimeseriesGroupGetDmarcResponseJSON   `json:"-"`
}

// radarEmailRoutingTimeseriesGroupGetDmarcResponseJSON contains the JSON metadata
// for the struct [RadarEmailRoutingTimeseriesGroupGetDmarcResponse]
type radarEmailRoutingTimeseriesGroupGetDmarcResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailRoutingTimeseriesGroupGetDmarcResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailRoutingTimeseriesGroupGetDmarcResponseJSON) RawJSON() string {
	return r.raw
}

type RadarEmailRoutingTimeseriesGroupGetDmarcResponseResult struct {
	Meta   interface{}                                                  `json:"meta,required"`
	Serie0 RadarEmailRoutingTimeseriesGroupGetDmarcResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarEmailRoutingTimeseriesGroupGetDmarcResponseResultJSON   `json:"-"`
}

// radarEmailRoutingTimeseriesGroupGetDmarcResponseResultJSON contains the JSON
// metadata for the struct [RadarEmailRoutingTimeseriesGroupGetDmarcResponseResult]
type radarEmailRoutingTimeseriesGroupGetDmarcResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailRoutingTimeseriesGroupGetDmarcResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailRoutingTimeseriesGroupGetDmarcResponseResultJSON) RawJSON() string {
	return r.raw
}

type RadarEmailRoutingTimeseriesGroupGetDmarcResponseResultSerie0 struct {
	Fail []string                                                         `json:"FAIL,required"`
	None []string                                                         `json:"NONE,required"`
	Pass []string                                                         `json:"PASS,required"`
	JSON radarEmailRoutingTimeseriesGroupGetDmarcResponseResultSerie0JSON `json:"-"`
}

// radarEmailRoutingTimeseriesGroupGetDmarcResponseResultSerie0JSON contains the
// JSON metadata for the struct
// [RadarEmailRoutingTimeseriesGroupGetDmarcResponseResultSerie0]
type radarEmailRoutingTimeseriesGroupGetDmarcResponseResultSerie0JSON struct {
	Fail        apijson.Field
	None        apijson.Field
	Pass        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailRoutingTimeseriesGroupGetDmarcResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailRoutingTimeseriesGroupGetDmarcResponseResultSerie0JSON) RawJSON() string {
	return r.raw
}

type RadarEmailRoutingTimeseriesGroupGetEncryptedResponse struct {
	Result  RadarEmailRoutingTimeseriesGroupGetEncryptedResponseResult `json:"result,required"`
	Success bool                                                       `json:"success,required"`
	JSON    radarEmailRoutingTimeseriesGroupGetEncryptedResponseJSON   `json:"-"`
}

// radarEmailRoutingTimeseriesGroupGetEncryptedResponseJSON contains the JSON
// metadata for the struct [RadarEmailRoutingTimeseriesGroupGetEncryptedResponse]
type radarEmailRoutingTimeseriesGroupGetEncryptedResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailRoutingTimeseriesGroupGetEncryptedResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailRoutingTimeseriesGroupGetEncryptedResponseJSON) RawJSON() string {
	return r.raw
}

type RadarEmailRoutingTimeseriesGroupGetEncryptedResponseResult struct {
	Meta   interface{}                                                      `json:"meta,required"`
	Serie0 RadarEmailRoutingTimeseriesGroupGetEncryptedResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarEmailRoutingTimeseriesGroupGetEncryptedResponseResultJSON   `json:"-"`
}

// radarEmailRoutingTimeseriesGroupGetEncryptedResponseResultJSON contains the JSON
// metadata for the struct
// [RadarEmailRoutingTimeseriesGroupGetEncryptedResponseResult]
type radarEmailRoutingTimeseriesGroupGetEncryptedResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailRoutingTimeseriesGroupGetEncryptedResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailRoutingTimeseriesGroupGetEncryptedResponseResultJSON) RawJSON() string {
	return r.raw
}

type RadarEmailRoutingTimeseriesGroupGetEncryptedResponseResultSerie0 struct {
	Encrypted    []string                                                             `json:"ENCRYPTED,required"`
	NotEncrypted []string                                                             `json:"NOT_ENCRYPTED,required"`
	JSON         radarEmailRoutingTimeseriesGroupGetEncryptedResponseResultSerie0JSON `json:"-"`
}

// radarEmailRoutingTimeseriesGroupGetEncryptedResponseResultSerie0JSON contains
// the JSON metadata for the struct
// [RadarEmailRoutingTimeseriesGroupGetEncryptedResponseResultSerie0]
type radarEmailRoutingTimeseriesGroupGetEncryptedResponseResultSerie0JSON struct {
	Encrypted    apijson.Field
	NotEncrypted apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RadarEmailRoutingTimeseriesGroupGetEncryptedResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailRoutingTimeseriesGroupGetEncryptedResponseResultSerie0JSON) RawJSON() string {
	return r.raw
}

type RadarEmailRoutingTimeseriesGroupGetIPVersionResponse struct {
	Result  RadarEmailRoutingTimeseriesGroupGetIPVersionResponseResult `json:"result,required"`
	Success bool                                                       `json:"success,required"`
	JSON    radarEmailRoutingTimeseriesGroupGetIPVersionResponseJSON   `json:"-"`
}

// radarEmailRoutingTimeseriesGroupGetIPVersionResponseJSON contains the JSON
// metadata for the struct [RadarEmailRoutingTimeseriesGroupGetIPVersionResponse]
type radarEmailRoutingTimeseriesGroupGetIPVersionResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailRoutingTimeseriesGroupGetIPVersionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailRoutingTimeseriesGroupGetIPVersionResponseJSON) RawJSON() string {
	return r.raw
}

type RadarEmailRoutingTimeseriesGroupGetIPVersionResponseResult struct {
	Meta   interface{}                                                      `json:"meta,required"`
	Serie0 RadarEmailRoutingTimeseriesGroupGetIPVersionResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarEmailRoutingTimeseriesGroupGetIPVersionResponseResultJSON   `json:"-"`
}

// radarEmailRoutingTimeseriesGroupGetIPVersionResponseResultJSON contains the JSON
// metadata for the struct
// [RadarEmailRoutingTimeseriesGroupGetIPVersionResponseResult]
type radarEmailRoutingTimeseriesGroupGetIPVersionResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailRoutingTimeseriesGroupGetIPVersionResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailRoutingTimeseriesGroupGetIPVersionResponseResultJSON) RawJSON() string {
	return r.raw
}

type RadarEmailRoutingTimeseriesGroupGetIPVersionResponseResultSerie0 struct {
	IPv4 []string                                                             `json:"IPv4,required"`
	IPv6 []string                                                             `json:"IPv6,required"`
	JSON radarEmailRoutingTimeseriesGroupGetIPVersionResponseResultSerie0JSON `json:"-"`
}

// radarEmailRoutingTimeseriesGroupGetIPVersionResponseResultSerie0JSON contains
// the JSON metadata for the struct
// [RadarEmailRoutingTimeseriesGroupGetIPVersionResponseResultSerie0]
type radarEmailRoutingTimeseriesGroupGetIPVersionResponseResultSerie0JSON struct {
	IPv4        apijson.Field
	IPv6        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailRoutingTimeseriesGroupGetIPVersionResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailRoutingTimeseriesGroupGetIPVersionResponseResultSerie0JSON) RawJSON() string {
	return r.raw
}

type RadarEmailRoutingTimeseriesGroupGetSpfResponse struct {
	Result  RadarEmailRoutingTimeseriesGroupGetSpfResponseResult `json:"result,required"`
	Success bool                                                 `json:"success,required"`
	JSON    radarEmailRoutingTimeseriesGroupGetSpfResponseJSON   `json:"-"`
}

// radarEmailRoutingTimeseriesGroupGetSpfResponseJSON contains the JSON metadata
// for the struct [RadarEmailRoutingTimeseriesGroupGetSpfResponse]
type radarEmailRoutingTimeseriesGroupGetSpfResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailRoutingTimeseriesGroupGetSpfResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailRoutingTimeseriesGroupGetSpfResponseJSON) RawJSON() string {
	return r.raw
}

type RadarEmailRoutingTimeseriesGroupGetSpfResponseResult struct {
	Meta   interface{}                                                `json:"meta,required"`
	Serie0 RadarEmailRoutingTimeseriesGroupGetSpfResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarEmailRoutingTimeseriesGroupGetSpfResponseResultJSON   `json:"-"`
}

// radarEmailRoutingTimeseriesGroupGetSpfResponseResultJSON contains the JSON
// metadata for the struct [RadarEmailRoutingTimeseriesGroupGetSpfResponseResult]
type radarEmailRoutingTimeseriesGroupGetSpfResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailRoutingTimeseriesGroupGetSpfResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailRoutingTimeseriesGroupGetSpfResponseResultJSON) RawJSON() string {
	return r.raw
}

type RadarEmailRoutingTimeseriesGroupGetSpfResponseResultSerie0 struct {
	Fail []string                                                       `json:"FAIL,required"`
	None []string                                                       `json:"NONE,required"`
	Pass []string                                                       `json:"PASS,required"`
	JSON radarEmailRoutingTimeseriesGroupGetSpfResponseResultSerie0JSON `json:"-"`
}

// radarEmailRoutingTimeseriesGroupGetSpfResponseResultSerie0JSON contains the JSON
// metadata for the struct
// [RadarEmailRoutingTimeseriesGroupGetSpfResponseResultSerie0]
type radarEmailRoutingTimeseriesGroupGetSpfResponseResultSerie0JSON struct {
	Fail        apijson.Field
	None        apijson.Field
	Pass        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailRoutingTimeseriesGroupGetSpfResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailRoutingTimeseriesGroupGetSpfResponseResultSerie0JSON) RawJSON() string {
	return r.raw
}

type RadarEmailRoutingTimeseriesGroupGetArcParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarEmailRoutingTimeseriesGroupGetArcParamsAggInterval] `query:"aggInterval"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by the specified date range. For example, use `7d` and
	// `7dcontrol` to compare this week with the previous week. Use this parameter or
	// set specific start and end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filters results by DKIM (DomainKeys Identified Mail) validation status.
	Dkim param.Field[[]RadarEmailRoutingTimeseriesGroupGetArcParamsDkim] `query:"dkim"`
	// Filters results by DMARC (Domain-based Message Authentication, Reporting and
	// Conformance) validation status.
	Dmarc param.Field[[]RadarEmailRoutingTimeseriesGroupGetArcParamsDmarc] `query:"dmarc"`
	// Filters results by encryption status (encrypted vs. not-encrypted).
	Encrypted param.Field[[]RadarEmailRoutingTimeseriesGroupGetArcParamsEncrypted] `query:"encrypted"`
	// Format in which results will be returned.
	Format param.Field[RadarEmailRoutingTimeseriesGroupGetArcParamsFormat] `query:"format"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]RadarEmailRoutingTimeseriesGroupGetArcParamsIPVersion] `query:"ipVersion"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by SPF (Sender Policy Framework) validation status.
	Spf param.Field[[]RadarEmailRoutingTimeseriesGroupGetArcParamsSpf] `query:"spf"`
}

// URLQuery serializes [RadarEmailRoutingTimeseriesGroupGetArcParams]'s query
// parameters as `url.Values`.
func (r RadarEmailRoutingTimeseriesGroupGetArcParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarEmailRoutingTimeseriesGroupGetArcParamsAggInterval string

const (
	RadarEmailRoutingTimeseriesGroupGetArcParamsAggInterval15m RadarEmailRoutingTimeseriesGroupGetArcParamsAggInterval = "15m"
	RadarEmailRoutingTimeseriesGroupGetArcParamsAggInterval1h  RadarEmailRoutingTimeseriesGroupGetArcParamsAggInterval = "1h"
	RadarEmailRoutingTimeseriesGroupGetArcParamsAggInterval1d  RadarEmailRoutingTimeseriesGroupGetArcParamsAggInterval = "1d"
	RadarEmailRoutingTimeseriesGroupGetArcParamsAggInterval1w  RadarEmailRoutingTimeseriesGroupGetArcParamsAggInterval = "1w"
)

func (r RadarEmailRoutingTimeseriesGroupGetArcParamsAggInterval) IsKnown() bool {
	switch r {
	case RadarEmailRoutingTimeseriesGroupGetArcParamsAggInterval15m, RadarEmailRoutingTimeseriesGroupGetArcParamsAggInterval1h, RadarEmailRoutingTimeseriesGroupGetArcParamsAggInterval1d, RadarEmailRoutingTimeseriesGroupGetArcParamsAggInterval1w:
		return true
	}
	return false
}

type RadarEmailRoutingTimeseriesGroupGetArcParamsDkim string

const (
	RadarEmailRoutingTimeseriesGroupGetArcParamsDkimPass RadarEmailRoutingTimeseriesGroupGetArcParamsDkim = "PASS"
	RadarEmailRoutingTimeseriesGroupGetArcParamsDkimNone RadarEmailRoutingTimeseriesGroupGetArcParamsDkim = "NONE"
	RadarEmailRoutingTimeseriesGroupGetArcParamsDkimFail RadarEmailRoutingTimeseriesGroupGetArcParamsDkim = "FAIL"
)

func (r RadarEmailRoutingTimeseriesGroupGetArcParamsDkim) IsKnown() bool {
	switch r {
	case RadarEmailRoutingTimeseriesGroupGetArcParamsDkimPass, RadarEmailRoutingTimeseriesGroupGetArcParamsDkimNone, RadarEmailRoutingTimeseriesGroupGetArcParamsDkimFail:
		return true
	}
	return false
}

type RadarEmailRoutingTimeseriesGroupGetArcParamsDmarc string

const (
	RadarEmailRoutingTimeseriesGroupGetArcParamsDmarcPass RadarEmailRoutingTimeseriesGroupGetArcParamsDmarc = "PASS"
	RadarEmailRoutingTimeseriesGroupGetArcParamsDmarcNone RadarEmailRoutingTimeseriesGroupGetArcParamsDmarc = "NONE"
	RadarEmailRoutingTimeseriesGroupGetArcParamsDmarcFail RadarEmailRoutingTimeseriesGroupGetArcParamsDmarc = "FAIL"
)

func (r RadarEmailRoutingTimeseriesGroupGetArcParamsDmarc) IsKnown() bool {
	switch r {
	case RadarEmailRoutingTimeseriesGroupGetArcParamsDmarcPass, RadarEmailRoutingTimeseriesGroupGetArcParamsDmarcNone, RadarEmailRoutingTimeseriesGroupGetArcParamsDmarcFail:
		return true
	}
	return false
}

type RadarEmailRoutingTimeseriesGroupGetArcParamsEncrypted string

const (
	RadarEmailRoutingTimeseriesGroupGetArcParamsEncryptedEncrypted    RadarEmailRoutingTimeseriesGroupGetArcParamsEncrypted = "ENCRYPTED"
	RadarEmailRoutingTimeseriesGroupGetArcParamsEncryptedNotEncrypted RadarEmailRoutingTimeseriesGroupGetArcParamsEncrypted = "NOT_ENCRYPTED"
)

func (r RadarEmailRoutingTimeseriesGroupGetArcParamsEncrypted) IsKnown() bool {
	switch r {
	case RadarEmailRoutingTimeseriesGroupGetArcParamsEncryptedEncrypted, RadarEmailRoutingTimeseriesGroupGetArcParamsEncryptedNotEncrypted:
		return true
	}
	return false
}

// Format in which results will be returned.
type RadarEmailRoutingTimeseriesGroupGetArcParamsFormat string

const (
	RadarEmailRoutingTimeseriesGroupGetArcParamsFormatJson RadarEmailRoutingTimeseriesGroupGetArcParamsFormat = "JSON"
	RadarEmailRoutingTimeseriesGroupGetArcParamsFormatCsv  RadarEmailRoutingTimeseriesGroupGetArcParamsFormat = "CSV"
)

func (r RadarEmailRoutingTimeseriesGroupGetArcParamsFormat) IsKnown() bool {
	switch r {
	case RadarEmailRoutingTimeseriesGroupGetArcParamsFormatJson, RadarEmailRoutingTimeseriesGroupGetArcParamsFormatCsv:
		return true
	}
	return false
}

type RadarEmailRoutingTimeseriesGroupGetArcParamsIPVersion string

const (
	RadarEmailRoutingTimeseriesGroupGetArcParamsIPVersionIPv4 RadarEmailRoutingTimeseriesGroupGetArcParamsIPVersion = "IPv4"
	RadarEmailRoutingTimeseriesGroupGetArcParamsIPVersionIPv6 RadarEmailRoutingTimeseriesGroupGetArcParamsIPVersion = "IPv6"
)

func (r RadarEmailRoutingTimeseriesGroupGetArcParamsIPVersion) IsKnown() bool {
	switch r {
	case RadarEmailRoutingTimeseriesGroupGetArcParamsIPVersionIPv4, RadarEmailRoutingTimeseriesGroupGetArcParamsIPVersionIPv6:
		return true
	}
	return false
}

type RadarEmailRoutingTimeseriesGroupGetArcParamsSpf string

const (
	RadarEmailRoutingTimeseriesGroupGetArcParamsSpfPass RadarEmailRoutingTimeseriesGroupGetArcParamsSpf = "PASS"
	RadarEmailRoutingTimeseriesGroupGetArcParamsSpfNone RadarEmailRoutingTimeseriesGroupGetArcParamsSpf = "NONE"
	RadarEmailRoutingTimeseriesGroupGetArcParamsSpfFail RadarEmailRoutingTimeseriesGroupGetArcParamsSpf = "FAIL"
)

func (r RadarEmailRoutingTimeseriesGroupGetArcParamsSpf) IsKnown() bool {
	switch r {
	case RadarEmailRoutingTimeseriesGroupGetArcParamsSpfPass, RadarEmailRoutingTimeseriesGroupGetArcParamsSpfNone, RadarEmailRoutingTimeseriesGroupGetArcParamsSpfFail:
		return true
	}
	return false
}

type RadarEmailRoutingTimeseriesGroupGetDkimParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarEmailRoutingTimeseriesGroupGetDkimParamsAggInterval] `query:"aggInterval"`
	// Filters results by ARC (Authenticated Received Chain) validation.
	Arc param.Field[[]RadarEmailRoutingTimeseriesGroupGetDkimParamsArc] `query:"arc"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by the specified date range. For example, use `7d` and
	// `7dcontrol` to compare this week with the previous week. Use this parameter or
	// set specific start and end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filters results by DMARC (Domain-based Message Authentication, Reporting and
	// Conformance) validation status.
	Dmarc param.Field[[]RadarEmailRoutingTimeseriesGroupGetDkimParamsDmarc] `query:"dmarc"`
	// Filters results by encryption status (encrypted vs. not-encrypted).
	Encrypted param.Field[[]RadarEmailRoutingTimeseriesGroupGetDkimParamsEncrypted] `query:"encrypted"`
	// Format in which results will be returned.
	Format param.Field[RadarEmailRoutingTimeseriesGroupGetDkimParamsFormat] `query:"format"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]RadarEmailRoutingTimeseriesGroupGetDkimParamsIPVersion] `query:"ipVersion"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by SPF (Sender Policy Framework) validation status.
	Spf param.Field[[]RadarEmailRoutingTimeseriesGroupGetDkimParamsSpf] `query:"spf"`
}

// URLQuery serializes [RadarEmailRoutingTimeseriesGroupGetDkimParams]'s query
// parameters as `url.Values`.
func (r RadarEmailRoutingTimeseriesGroupGetDkimParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarEmailRoutingTimeseriesGroupGetDkimParamsAggInterval string

const (
	RadarEmailRoutingTimeseriesGroupGetDkimParamsAggInterval15m RadarEmailRoutingTimeseriesGroupGetDkimParamsAggInterval = "15m"
	RadarEmailRoutingTimeseriesGroupGetDkimParamsAggInterval1h  RadarEmailRoutingTimeseriesGroupGetDkimParamsAggInterval = "1h"
	RadarEmailRoutingTimeseriesGroupGetDkimParamsAggInterval1d  RadarEmailRoutingTimeseriesGroupGetDkimParamsAggInterval = "1d"
	RadarEmailRoutingTimeseriesGroupGetDkimParamsAggInterval1w  RadarEmailRoutingTimeseriesGroupGetDkimParamsAggInterval = "1w"
)

func (r RadarEmailRoutingTimeseriesGroupGetDkimParamsAggInterval) IsKnown() bool {
	switch r {
	case RadarEmailRoutingTimeseriesGroupGetDkimParamsAggInterval15m, RadarEmailRoutingTimeseriesGroupGetDkimParamsAggInterval1h, RadarEmailRoutingTimeseriesGroupGetDkimParamsAggInterval1d, RadarEmailRoutingTimeseriesGroupGetDkimParamsAggInterval1w:
		return true
	}
	return false
}

type RadarEmailRoutingTimeseriesGroupGetDkimParamsArc string

const (
	RadarEmailRoutingTimeseriesGroupGetDkimParamsArcPass RadarEmailRoutingTimeseriesGroupGetDkimParamsArc = "PASS"
	RadarEmailRoutingTimeseriesGroupGetDkimParamsArcNone RadarEmailRoutingTimeseriesGroupGetDkimParamsArc = "NONE"
	RadarEmailRoutingTimeseriesGroupGetDkimParamsArcFail RadarEmailRoutingTimeseriesGroupGetDkimParamsArc = "FAIL"
)

func (r RadarEmailRoutingTimeseriesGroupGetDkimParamsArc) IsKnown() bool {
	switch r {
	case RadarEmailRoutingTimeseriesGroupGetDkimParamsArcPass, RadarEmailRoutingTimeseriesGroupGetDkimParamsArcNone, RadarEmailRoutingTimeseriesGroupGetDkimParamsArcFail:
		return true
	}
	return false
}

type RadarEmailRoutingTimeseriesGroupGetDkimParamsDmarc string

const (
	RadarEmailRoutingTimeseriesGroupGetDkimParamsDmarcPass RadarEmailRoutingTimeseriesGroupGetDkimParamsDmarc = "PASS"
	RadarEmailRoutingTimeseriesGroupGetDkimParamsDmarcNone RadarEmailRoutingTimeseriesGroupGetDkimParamsDmarc = "NONE"
	RadarEmailRoutingTimeseriesGroupGetDkimParamsDmarcFail RadarEmailRoutingTimeseriesGroupGetDkimParamsDmarc = "FAIL"
)

func (r RadarEmailRoutingTimeseriesGroupGetDkimParamsDmarc) IsKnown() bool {
	switch r {
	case RadarEmailRoutingTimeseriesGroupGetDkimParamsDmarcPass, RadarEmailRoutingTimeseriesGroupGetDkimParamsDmarcNone, RadarEmailRoutingTimeseriesGroupGetDkimParamsDmarcFail:
		return true
	}
	return false
}

type RadarEmailRoutingTimeseriesGroupGetDkimParamsEncrypted string

const (
	RadarEmailRoutingTimeseriesGroupGetDkimParamsEncryptedEncrypted    RadarEmailRoutingTimeseriesGroupGetDkimParamsEncrypted = "ENCRYPTED"
	RadarEmailRoutingTimeseriesGroupGetDkimParamsEncryptedNotEncrypted RadarEmailRoutingTimeseriesGroupGetDkimParamsEncrypted = "NOT_ENCRYPTED"
)

func (r RadarEmailRoutingTimeseriesGroupGetDkimParamsEncrypted) IsKnown() bool {
	switch r {
	case RadarEmailRoutingTimeseriesGroupGetDkimParamsEncryptedEncrypted, RadarEmailRoutingTimeseriesGroupGetDkimParamsEncryptedNotEncrypted:
		return true
	}
	return false
}

// Format in which results will be returned.
type RadarEmailRoutingTimeseriesGroupGetDkimParamsFormat string

const (
	RadarEmailRoutingTimeseriesGroupGetDkimParamsFormatJson RadarEmailRoutingTimeseriesGroupGetDkimParamsFormat = "JSON"
	RadarEmailRoutingTimeseriesGroupGetDkimParamsFormatCsv  RadarEmailRoutingTimeseriesGroupGetDkimParamsFormat = "CSV"
)

func (r RadarEmailRoutingTimeseriesGroupGetDkimParamsFormat) IsKnown() bool {
	switch r {
	case RadarEmailRoutingTimeseriesGroupGetDkimParamsFormatJson, RadarEmailRoutingTimeseriesGroupGetDkimParamsFormatCsv:
		return true
	}
	return false
}

type RadarEmailRoutingTimeseriesGroupGetDkimParamsIPVersion string

const (
	RadarEmailRoutingTimeseriesGroupGetDkimParamsIPVersionIPv4 RadarEmailRoutingTimeseriesGroupGetDkimParamsIPVersion = "IPv4"
	RadarEmailRoutingTimeseriesGroupGetDkimParamsIPVersionIPv6 RadarEmailRoutingTimeseriesGroupGetDkimParamsIPVersion = "IPv6"
)

func (r RadarEmailRoutingTimeseriesGroupGetDkimParamsIPVersion) IsKnown() bool {
	switch r {
	case RadarEmailRoutingTimeseriesGroupGetDkimParamsIPVersionIPv4, RadarEmailRoutingTimeseriesGroupGetDkimParamsIPVersionIPv6:
		return true
	}
	return false
}

type RadarEmailRoutingTimeseriesGroupGetDkimParamsSpf string

const (
	RadarEmailRoutingTimeseriesGroupGetDkimParamsSpfPass RadarEmailRoutingTimeseriesGroupGetDkimParamsSpf = "PASS"
	RadarEmailRoutingTimeseriesGroupGetDkimParamsSpfNone RadarEmailRoutingTimeseriesGroupGetDkimParamsSpf = "NONE"
	RadarEmailRoutingTimeseriesGroupGetDkimParamsSpfFail RadarEmailRoutingTimeseriesGroupGetDkimParamsSpf = "FAIL"
)

func (r RadarEmailRoutingTimeseriesGroupGetDkimParamsSpf) IsKnown() bool {
	switch r {
	case RadarEmailRoutingTimeseriesGroupGetDkimParamsSpfPass, RadarEmailRoutingTimeseriesGroupGetDkimParamsSpfNone, RadarEmailRoutingTimeseriesGroupGetDkimParamsSpfFail:
		return true
	}
	return false
}

type RadarEmailRoutingTimeseriesGroupGetDmarcParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarEmailRoutingTimeseriesGroupGetDmarcParamsAggInterval] `query:"aggInterval"`
	// Filters results by ARC (Authenticated Received Chain) validation.
	Arc param.Field[[]RadarEmailRoutingTimeseriesGroupGetDmarcParamsArc] `query:"arc"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by the specified date range. For example, use `7d` and
	// `7dcontrol` to compare this week with the previous week. Use this parameter or
	// set specific start and end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filters results by DKIM (DomainKeys Identified Mail) validation status.
	Dkim param.Field[[]RadarEmailRoutingTimeseriesGroupGetDmarcParamsDkim] `query:"dkim"`
	// Filters results by encryption status (encrypted vs. not-encrypted).
	Encrypted param.Field[[]RadarEmailRoutingTimeseriesGroupGetDmarcParamsEncrypted] `query:"encrypted"`
	// Format in which results will be returned.
	Format param.Field[RadarEmailRoutingTimeseriesGroupGetDmarcParamsFormat] `query:"format"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]RadarEmailRoutingTimeseriesGroupGetDmarcParamsIPVersion] `query:"ipVersion"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by SPF (Sender Policy Framework) validation status.
	Spf param.Field[[]RadarEmailRoutingTimeseriesGroupGetDmarcParamsSpf] `query:"spf"`
}

// URLQuery serializes [RadarEmailRoutingTimeseriesGroupGetDmarcParams]'s query
// parameters as `url.Values`.
func (r RadarEmailRoutingTimeseriesGroupGetDmarcParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarEmailRoutingTimeseriesGroupGetDmarcParamsAggInterval string

const (
	RadarEmailRoutingTimeseriesGroupGetDmarcParamsAggInterval15m RadarEmailRoutingTimeseriesGroupGetDmarcParamsAggInterval = "15m"
	RadarEmailRoutingTimeseriesGroupGetDmarcParamsAggInterval1h  RadarEmailRoutingTimeseriesGroupGetDmarcParamsAggInterval = "1h"
	RadarEmailRoutingTimeseriesGroupGetDmarcParamsAggInterval1d  RadarEmailRoutingTimeseriesGroupGetDmarcParamsAggInterval = "1d"
	RadarEmailRoutingTimeseriesGroupGetDmarcParamsAggInterval1w  RadarEmailRoutingTimeseriesGroupGetDmarcParamsAggInterval = "1w"
)

func (r RadarEmailRoutingTimeseriesGroupGetDmarcParamsAggInterval) IsKnown() bool {
	switch r {
	case RadarEmailRoutingTimeseriesGroupGetDmarcParamsAggInterval15m, RadarEmailRoutingTimeseriesGroupGetDmarcParamsAggInterval1h, RadarEmailRoutingTimeseriesGroupGetDmarcParamsAggInterval1d, RadarEmailRoutingTimeseriesGroupGetDmarcParamsAggInterval1w:
		return true
	}
	return false
}

type RadarEmailRoutingTimeseriesGroupGetDmarcParamsArc string

const (
	RadarEmailRoutingTimeseriesGroupGetDmarcParamsArcPass RadarEmailRoutingTimeseriesGroupGetDmarcParamsArc = "PASS"
	RadarEmailRoutingTimeseriesGroupGetDmarcParamsArcNone RadarEmailRoutingTimeseriesGroupGetDmarcParamsArc = "NONE"
	RadarEmailRoutingTimeseriesGroupGetDmarcParamsArcFail RadarEmailRoutingTimeseriesGroupGetDmarcParamsArc = "FAIL"
)

func (r RadarEmailRoutingTimeseriesGroupGetDmarcParamsArc) IsKnown() bool {
	switch r {
	case RadarEmailRoutingTimeseriesGroupGetDmarcParamsArcPass, RadarEmailRoutingTimeseriesGroupGetDmarcParamsArcNone, RadarEmailRoutingTimeseriesGroupGetDmarcParamsArcFail:
		return true
	}
	return false
}

type RadarEmailRoutingTimeseriesGroupGetDmarcParamsDkim string

const (
	RadarEmailRoutingTimeseriesGroupGetDmarcParamsDkimPass RadarEmailRoutingTimeseriesGroupGetDmarcParamsDkim = "PASS"
	RadarEmailRoutingTimeseriesGroupGetDmarcParamsDkimNone RadarEmailRoutingTimeseriesGroupGetDmarcParamsDkim = "NONE"
	RadarEmailRoutingTimeseriesGroupGetDmarcParamsDkimFail RadarEmailRoutingTimeseriesGroupGetDmarcParamsDkim = "FAIL"
)

func (r RadarEmailRoutingTimeseriesGroupGetDmarcParamsDkim) IsKnown() bool {
	switch r {
	case RadarEmailRoutingTimeseriesGroupGetDmarcParamsDkimPass, RadarEmailRoutingTimeseriesGroupGetDmarcParamsDkimNone, RadarEmailRoutingTimeseriesGroupGetDmarcParamsDkimFail:
		return true
	}
	return false
}

type RadarEmailRoutingTimeseriesGroupGetDmarcParamsEncrypted string

const (
	RadarEmailRoutingTimeseriesGroupGetDmarcParamsEncryptedEncrypted    RadarEmailRoutingTimeseriesGroupGetDmarcParamsEncrypted = "ENCRYPTED"
	RadarEmailRoutingTimeseriesGroupGetDmarcParamsEncryptedNotEncrypted RadarEmailRoutingTimeseriesGroupGetDmarcParamsEncrypted = "NOT_ENCRYPTED"
)

func (r RadarEmailRoutingTimeseriesGroupGetDmarcParamsEncrypted) IsKnown() bool {
	switch r {
	case RadarEmailRoutingTimeseriesGroupGetDmarcParamsEncryptedEncrypted, RadarEmailRoutingTimeseriesGroupGetDmarcParamsEncryptedNotEncrypted:
		return true
	}
	return false
}

// Format in which results will be returned.
type RadarEmailRoutingTimeseriesGroupGetDmarcParamsFormat string

const (
	RadarEmailRoutingTimeseriesGroupGetDmarcParamsFormatJson RadarEmailRoutingTimeseriesGroupGetDmarcParamsFormat = "JSON"
	RadarEmailRoutingTimeseriesGroupGetDmarcParamsFormatCsv  RadarEmailRoutingTimeseriesGroupGetDmarcParamsFormat = "CSV"
)

func (r RadarEmailRoutingTimeseriesGroupGetDmarcParamsFormat) IsKnown() bool {
	switch r {
	case RadarEmailRoutingTimeseriesGroupGetDmarcParamsFormatJson, RadarEmailRoutingTimeseriesGroupGetDmarcParamsFormatCsv:
		return true
	}
	return false
}

type RadarEmailRoutingTimeseriesGroupGetDmarcParamsIPVersion string

const (
	RadarEmailRoutingTimeseriesGroupGetDmarcParamsIPVersionIPv4 RadarEmailRoutingTimeseriesGroupGetDmarcParamsIPVersion = "IPv4"
	RadarEmailRoutingTimeseriesGroupGetDmarcParamsIPVersionIPv6 RadarEmailRoutingTimeseriesGroupGetDmarcParamsIPVersion = "IPv6"
)

func (r RadarEmailRoutingTimeseriesGroupGetDmarcParamsIPVersion) IsKnown() bool {
	switch r {
	case RadarEmailRoutingTimeseriesGroupGetDmarcParamsIPVersionIPv4, RadarEmailRoutingTimeseriesGroupGetDmarcParamsIPVersionIPv6:
		return true
	}
	return false
}

type RadarEmailRoutingTimeseriesGroupGetDmarcParamsSpf string

const (
	RadarEmailRoutingTimeseriesGroupGetDmarcParamsSpfPass RadarEmailRoutingTimeseriesGroupGetDmarcParamsSpf = "PASS"
	RadarEmailRoutingTimeseriesGroupGetDmarcParamsSpfNone RadarEmailRoutingTimeseriesGroupGetDmarcParamsSpf = "NONE"
	RadarEmailRoutingTimeseriesGroupGetDmarcParamsSpfFail RadarEmailRoutingTimeseriesGroupGetDmarcParamsSpf = "FAIL"
)

func (r RadarEmailRoutingTimeseriesGroupGetDmarcParamsSpf) IsKnown() bool {
	switch r {
	case RadarEmailRoutingTimeseriesGroupGetDmarcParamsSpfPass, RadarEmailRoutingTimeseriesGroupGetDmarcParamsSpfNone, RadarEmailRoutingTimeseriesGroupGetDmarcParamsSpfFail:
		return true
	}
	return false
}

type RadarEmailRoutingTimeseriesGroupGetEncryptedParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarEmailRoutingTimeseriesGroupGetEncryptedParamsAggInterval] `query:"aggInterval"`
	// Filters results by ARC (Authenticated Received Chain) validation.
	Arc param.Field[[]RadarEmailRoutingTimeseriesGroupGetEncryptedParamsArc] `query:"arc"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by the specified date range. For example, use `7d` and
	// `7dcontrol` to compare this week with the previous week. Use this parameter or
	// set specific start and end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filters results by DKIM (DomainKeys Identified Mail) validation status.
	Dkim param.Field[[]RadarEmailRoutingTimeseriesGroupGetEncryptedParamsDkim] `query:"dkim"`
	// Filters results by DMARC (Domain-based Message Authentication, Reporting and
	// Conformance) validation status.
	Dmarc param.Field[[]RadarEmailRoutingTimeseriesGroupGetEncryptedParamsDmarc] `query:"dmarc"`
	// Format in which results will be returned.
	Format param.Field[RadarEmailRoutingTimeseriesGroupGetEncryptedParamsFormat] `query:"format"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]RadarEmailRoutingTimeseriesGroupGetEncryptedParamsIPVersion] `query:"ipVersion"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by SPF (Sender Policy Framework) validation status.
	Spf param.Field[[]RadarEmailRoutingTimeseriesGroupGetEncryptedParamsSpf] `query:"spf"`
}

// URLQuery serializes [RadarEmailRoutingTimeseriesGroupGetEncryptedParams]'s query
// parameters as `url.Values`.
func (r RadarEmailRoutingTimeseriesGroupGetEncryptedParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarEmailRoutingTimeseriesGroupGetEncryptedParamsAggInterval string

const (
	RadarEmailRoutingTimeseriesGroupGetEncryptedParamsAggInterval15m RadarEmailRoutingTimeseriesGroupGetEncryptedParamsAggInterval = "15m"
	RadarEmailRoutingTimeseriesGroupGetEncryptedParamsAggInterval1h  RadarEmailRoutingTimeseriesGroupGetEncryptedParamsAggInterval = "1h"
	RadarEmailRoutingTimeseriesGroupGetEncryptedParamsAggInterval1d  RadarEmailRoutingTimeseriesGroupGetEncryptedParamsAggInterval = "1d"
	RadarEmailRoutingTimeseriesGroupGetEncryptedParamsAggInterval1w  RadarEmailRoutingTimeseriesGroupGetEncryptedParamsAggInterval = "1w"
)

func (r RadarEmailRoutingTimeseriesGroupGetEncryptedParamsAggInterval) IsKnown() bool {
	switch r {
	case RadarEmailRoutingTimeseriesGroupGetEncryptedParamsAggInterval15m, RadarEmailRoutingTimeseriesGroupGetEncryptedParamsAggInterval1h, RadarEmailRoutingTimeseriesGroupGetEncryptedParamsAggInterval1d, RadarEmailRoutingTimeseriesGroupGetEncryptedParamsAggInterval1w:
		return true
	}
	return false
}

type RadarEmailRoutingTimeseriesGroupGetEncryptedParamsArc string

const (
	RadarEmailRoutingTimeseriesGroupGetEncryptedParamsArcPass RadarEmailRoutingTimeseriesGroupGetEncryptedParamsArc = "PASS"
	RadarEmailRoutingTimeseriesGroupGetEncryptedParamsArcNone RadarEmailRoutingTimeseriesGroupGetEncryptedParamsArc = "NONE"
	RadarEmailRoutingTimeseriesGroupGetEncryptedParamsArcFail RadarEmailRoutingTimeseriesGroupGetEncryptedParamsArc = "FAIL"
)

func (r RadarEmailRoutingTimeseriesGroupGetEncryptedParamsArc) IsKnown() bool {
	switch r {
	case RadarEmailRoutingTimeseriesGroupGetEncryptedParamsArcPass, RadarEmailRoutingTimeseriesGroupGetEncryptedParamsArcNone, RadarEmailRoutingTimeseriesGroupGetEncryptedParamsArcFail:
		return true
	}
	return false
}

type RadarEmailRoutingTimeseriesGroupGetEncryptedParamsDkim string

const (
	RadarEmailRoutingTimeseriesGroupGetEncryptedParamsDkimPass RadarEmailRoutingTimeseriesGroupGetEncryptedParamsDkim = "PASS"
	RadarEmailRoutingTimeseriesGroupGetEncryptedParamsDkimNone RadarEmailRoutingTimeseriesGroupGetEncryptedParamsDkim = "NONE"
	RadarEmailRoutingTimeseriesGroupGetEncryptedParamsDkimFail RadarEmailRoutingTimeseriesGroupGetEncryptedParamsDkim = "FAIL"
)

func (r RadarEmailRoutingTimeseriesGroupGetEncryptedParamsDkim) IsKnown() bool {
	switch r {
	case RadarEmailRoutingTimeseriesGroupGetEncryptedParamsDkimPass, RadarEmailRoutingTimeseriesGroupGetEncryptedParamsDkimNone, RadarEmailRoutingTimeseriesGroupGetEncryptedParamsDkimFail:
		return true
	}
	return false
}

type RadarEmailRoutingTimeseriesGroupGetEncryptedParamsDmarc string

const (
	RadarEmailRoutingTimeseriesGroupGetEncryptedParamsDmarcPass RadarEmailRoutingTimeseriesGroupGetEncryptedParamsDmarc = "PASS"
	RadarEmailRoutingTimeseriesGroupGetEncryptedParamsDmarcNone RadarEmailRoutingTimeseriesGroupGetEncryptedParamsDmarc = "NONE"
	RadarEmailRoutingTimeseriesGroupGetEncryptedParamsDmarcFail RadarEmailRoutingTimeseriesGroupGetEncryptedParamsDmarc = "FAIL"
)

func (r RadarEmailRoutingTimeseriesGroupGetEncryptedParamsDmarc) IsKnown() bool {
	switch r {
	case RadarEmailRoutingTimeseriesGroupGetEncryptedParamsDmarcPass, RadarEmailRoutingTimeseriesGroupGetEncryptedParamsDmarcNone, RadarEmailRoutingTimeseriesGroupGetEncryptedParamsDmarcFail:
		return true
	}
	return false
}

// Format in which results will be returned.
type RadarEmailRoutingTimeseriesGroupGetEncryptedParamsFormat string

const (
	RadarEmailRoutingTimeseriesGroupGetEncryptedParamsFormatJson RadarEmailRoutingTimeseriesGroupGetEncryptedParamsFormat = "JSON"
	RadarEmailRoutingTimeseriesGroupGetEncryptedParamsFormatCsv  RadarEmailRoutingTimeseriesGroupGetEncryptedParamsFormat = "CSV"
)

func (r RadarEmailRoutingTimeseriesGroupGetEncryptedParamsFormat) IsKnown() bool {
	switch r {
	case RadarEmailRoutingTimeseriesGroupGetEncryptedParamsFormatJson, RadarEmailRoutingTimeseriesGroupGetEncryptedParamsFormatCsv:
		return true
	}
	return false
}

type RadarEmailRoutingTimeseriesGroupGetEncryptedParamsIPVersion string

const (
	RadarEmailRoutingTimeseriesGroupGetEncryptedParamsIPVersionIPv4 RadarEmailRoutingTimeseriesGroupGetEncryptedParamsIPVersion = "IPv4"
	RadarEmailRoutingTimeseriesGroupGetEncryptedParamsIPVersionIPv6 RadarEmailRoutingTimeseriesGroupGetEncryptedParamsIPVersion = "IPv6"
)

func (r RadarEmailRoutingTimeseriesGroupGetEncryptedParamsIPVersion) IsKnown() bool {
	switch r {
	case RadarEmailRoutingTimeseriesGroupGetEncryptedParamsIPVersionIPv4, RadarEmailRoutingTimeseriesGroupGetEncryptedParamsIPVersionIPv6:
		return true
	}
	return false
}

type RadarEmailRoutingTimeseriesGroupGetEncryptedParamsSpf string

const (
	RadarEmailRoutingTimeseriesGroupGetEncryptedParamsSpfPass RadarEmailRoutingTimeseriesGroupGetEncryptedParamsSpf = "PASS"
	RadarEmailRoutingTimeseriesGroupGetEncryptedParamsSpfNone RadarEmailRoutingTimeseriesGroupGetEncryptedParamsSpf = "NONE"
	RadarEmailRoutingTimeseriesGroupGetEncryptedParamsSpfFail RadarEmailRoutingTimeseriesGroupGetEncryptedParamsSpf = "FAIL"
)

func (r RadarEmailRoutingTimeseriesGroupGetEncryptedParamsSpf) IsKnown() bool {
	switch r {
	case RadarEmailRoutingTimeseriesGroupGetEncryptedParamsSpfPass, RadarEmailRoutingTimeseriesGroupGetEncryptedParamsSpfNone, RadarEmailRoutingTimeseriesGroupGetEncryptedParamsSpfFail:
		return true
	}
	return false
}

type RadarEmailRoutingTimeseriesGroupGetIPVersionParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarEmailRoutingTimeseriesGroupGetIPVersionParamsAggInterval] `query:"aggInterval"`
	// Filters results by ARC (Authenticated Received Chain) validation.
	Arc param.Field[[]RadarEmailRoutingTimeseriesGroupGetIPVersionParamsArc] `query:"arc"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by the specified date range. For example, use `7d` and
	// `7dcontrol` to compare this week with the previous week. Use this parameter or
	// set specific start and end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filters results by DKIM (DomainKeys Identified Mail) validation status.
	Dkim param.Field[[]RadarEmailRoutingTimeseriesGroupGetIPVersionParamsDkim] `query:"dkim"`
	// Filters results by DMARC (Domain-based Message Authentication, Reporting and
	// Conformance) validation status.
	Dmarc param.Field[[]RadarEmailRoutingTimeseriesGroupGetIPVersionParamsDmarc] `query:"dmarc"`
	// Filters results by encryption status (encrypted vs. not-encrypted).
	Encrypted param.Field[[]RadarEmailRoutingTimeseriesGroupGetIPVersionParamsEncrypted] `query:"encrypted"`
	// Format in which results will be returned.
	Format param.Field[RadarEmailRoutingTimeseriesGroupGetIPVersionParamsFormat] `query:"format"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by SPF (Sender Policy Framework) validation status.
	Spf param.Field[[]RadarEmailRoutingTimeseriesGroupGetIPVersionParamsSpf] `query:"spf"`
}

// URLQuery serializes [RadarEmailRoutingTimeseriesGroupGetIPVersionParams]'s query
// parameters as `url.Values`.
func (r RadarEmailRoutingTimeseriesGroupGetIPVersionParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarEmailRoutingTimeseriesGroupGetIPVersionParamsAggInterval string

const (
	RadarEmailRoutingTimeseriesGroupGetIPVersionParamsAggInterval15m RadarEmailRoutingTimeseriesGroupGetIPVersionParamsAggInterval = "15m"
	RadarEmailRoutingTimeseriesGroupGetIPVersionParamsAggInterval1h  RadarEmailRoutingTimeseriesGroupGetIPVersionParamsAggInterval = "1h"
	RadarEmailRoutingTimeseriesGroupGetIPVersionParamsAggInterval1d  RadarEmailRoutingTimeseriesGroupGetIPVersionParamsAggInterval = "1d"
	RadarEmailRoutingTimeseriesGroupGetIPVersionParamsAggInterval1w  RadarEmailRoutingTimeseriesGroupGetIPVersionParamsAggInterval = "1w"
)

func (r RadarEmailRoutingTimeseriesGroupGetIPVersionParamsAggInterval) IsKnown() bool {
	switch r {
	case RadarEmailRoutingTimeseriesGroupGetIPVersionParamsAggInterval15m, RadarEmailRoutingTimeseriesGroupGetIPVersionParamsAggInterval1h, RadarEmailRoutingTimeseriesGroupGetIPVersionParamsAggInterval1d, RadarEmailRoutingTimeseriesGroupGetIPVersionParamsAggInterval1w:
		return true
	}
	return false
}

type RadarEmailRoutingTimeseriesGroupGetIPVersionParamsArc string

const (
	RadarEmailRoutingTimeseriesGroupGetIPVersionParamsArcPass RadarEmailRoutingTimeseriesGroupGetIPVersionParamsArc = "PASS"
	RadarEmailRoutingTimeseriesGroupGetIPVersionParamsArcNone RadarEmailRoutingTimeseriesGroupGetIPVersionParamsArc = "NONE"
	RadarEmailRoutingTimeseriesGroupGetIPVersionParamsArcFail RadarEmailRoutingTimeseriesGroupGetIPVersionParamsArc = "FAIL"
)

func (r RadarEmailRoutingTimeseriesGroupGetIPVersionParamsArc) IsKnown() bool {
	switch r {
	case RadarEmailRoutingTimeseriesGroupGetIPVersionParamsArcPass, RadarEmailRoutingTimeseriesGroupGetIPVersionParamsArcNone, RadarEmailRoutingTimeseriesGroupGetIPVersionParamsArcFail:
		return true
	}
	return false
}

type RadarEmailRoutingTimeseriesGroupGetIPVersionParamsDkim string

const (
	RadarEmailRoutingTimeseriesGroupGetIPVersionParamsDkimPass RadarEmailRoutingTimeseriesGroupGetIPVersionParamsDkim = "PASS"
	RadarEmailRoutingTimeseriesGroupGetIPVersionParamsDkimNone RadarEmailRoutingTimeseriesGroupGetIPVersionParamsDkim = "NONE"
	RadarEmailRoutingTimeseriesGroupGetIPVersionParamsDkimFail RadarEmailRoutingTimeseriesGroupGetIPVersionParamsDkim = "FAIL"
)

func (r RadarEmailRoutingTimeseriesGroupGetIPVersionParamsDkim) IsKnown() bool {
	switch r {
	case RadarEmailRoutingTimeseriesGroupGetIPVersionParamsDkimPass, RadarEmailRoutingTimeseriesGroupGetIPVersionParamsDkimNone, RadarEmailRoutingTimeseriesGroupGetIPVersionParamsDkimFail:
		return true
	}
	return false
}

type RadarEmailRoutingTimeseriesGroupGetIPVersionParamsDmarc string

const (
	RadarEmailRoutingTimeseriesGroupGetIPVersionParamsDmarcPass RadarEmailRoutingTimeseriesGroupGetIPVersionParamsDmarc = "PASS"
	RadarEmailRoutingTimeseriesGroupGetIPVersionParamsDmarcNone RadarEmailRoutingTimeseriesGroupGetIPVersionParamsDmarc = "NONE"
	RadarEmailRoutingTimeseriesGroupGetIPVersionParamsDmarcFail RadarEmailRoutingTimeseriesGroupGetIPVersionParamsDmarc = "FAIL"
)

func (r RadarEmailRoutingTimeseriesGroupGetIPVersionParamsDmarc) IsKnown() bool {
	switch r {
	case RadarEmailRoutingTimeseriesGroupGetIPVersionParamsDmarcPass, RadarEmailRoutingTimeseriesGroupGetIPVersionParamsDmarcNone, RadarEmailRoutingTimeseriesGroupGetIPVersionParamsDmarcFail:
		return true
	}
	return false
}

type RadarEmailRoutingTimeseriesGroupGetIPVersionParamsEncrypted string

const (
	RadarEmailRoutingTimeseriesGroupGetIPVersionParamsEncryptedEncrypted    RadarEmailRoutingTimeseriesGroupGetIPVersionParamsEncrypted = "ENCRYPTED"
	RadarEmailRoutingTimeseriesGroupGetIPVersionParamsEncryptedNotEncrypted RadarEmailRoutingTimeseriesGroupGetIPVersionParamsEncrypted = "NOT_ENCRYPTED"
)

func (r RadarEmailRoutingTimeseriesGroupGetIPVersionParamsEncrypted) IsKnown() bool {
	switch r {
	case RadarEmailRoutingTimeseriesGroupGetIPVersionParamsEncryptedEncrypted, RadarEmailRoutingTimeseriesGroupGetIPVersionParamsEncryptedNotEncrypted:
		return true
	}
	return false
}

// Format in which results will be returned.
type RadarEmailRoutingTimeseriesGroupGetIPVersionParamsFormat string

const (
	RadarEmailRoutingTimeseriesGroupGetIPVersionParamsFormatJson RadarEmailRoutingTimeseriesGroupGetIPVersionParamsFormat = "JSON"
	RadarEmailRoutingTimeseriesGroupGetIPVersionParamsFormatCsv  RadarEmailRoutingTimeseriesGroupGetIPVersionParamsFormat = "CSV"
)

func (r RadarEmailRoutingTimeseriesGroupGetIPVersionParamsFormat) IsKnown() bool {
	switch r {
	case RadarEmailRoutingTimeseriesGroupGetIPVersionParamsFormatJson, RadarEmailRoutingTimeseriesGroupGetIPVersionParamsFormatCsv:
		return true
	}
	return false
}

type RadarEmailRoutingTimeseriesGroupGetIPVersionParamsSpf string

const (
	RadarEmailRoutingTimeseriesGroupGetIPVersionParamsSpfPass RadarEmailRoutingTimeseriesGroupGetIPVersionParamsSpf = "PASS"
	RadarEmailRoutingTimeseriesGroupGetIPVersionParamsSpfNone RadarEmailRoutingTimeseriesGroupGetIPVersionParamsSpf = "NONE"
	RadarEmailRoutingTimeseriesGroupGetIPVersionParamsSpfFail RadarEmailRoutingTimeseriesGroupGetIPVersionParamsSpf = "FAIL"
)

func (r RadarEmailRoutingTimeseriesGroupGetIPVersionParamsSpf) IsKnown() bool {
	switch r {
	case RadarEmailRoutingTimeseriesGroupGetIPVersionParamsSpfPass, RadarEmailRoutingTimeseriesGroupGetIPVersionParamsSpfNone, RadarEmailRoutingTimeseriesGroupGetIPVersionParamsSpfFail:
		return true
	}
	return false
}

type RadarEmailRoutingTimeseriesGroupGetSpfParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarEmailRoutingTimeseriesGroupGetSpfParamsAggInterval] `query:"aggInterval"`
	// Filters results by ARC (Authenticated Received Chain) validation.
	Arc param.Field[[]RadarEmailRoutingTimeseriesGroupGetSpfParamsArc] `query:"arc"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by the specified date range. For example, use `7d` and
	// `7dcontrol` to compare this week with the previous week. Use this parameter or
	// set specific start and end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filters results by DKIM (DomainKeys Identified Mail) validation status.
	Dkim param.Field[[]RadarEmailRoutingTimeseriesGroupGetSpfParamsDkim] `query:"dkim"`
	// Filters results by DMARC (Domain-based Message Authentication, Reporting and
	// Conformance) validation status.
	Dmarc param.Field[[]RadarEmailRoutingTimeseriesGroupGetSpfParamsDmarc] `query:"dmarc"`
	// Filters results by encryption status (encrypted vs. not-encrypted).
	Encrypted param.Field[[]RadarEmailRoutingTimeseriesGroupGetSpfParamsEncrypted] `query:"encrypted"`
	// Format in which results will be returned.
	Format param.Field[RadarEmailRoutingTimeseriesGroupGetSpfParamsFormat] `query:"format"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]RadarEmailRoutingTimeseriesGroupGetSpfParamsIPVersion] `query:"ipVersion"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarEmailRoutingTimeseriesGroupGetSpfParams]'s query
// parameters as `url.Values`.
func (r RadarEmailRoutingTimeseriesGroupGetSpfParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarEmailRoutingTimeseriesGroupGetSpfParamsAggInterval string

const (
	RadarEmailRoutingTimeseriesGroupGetSpfParamsAggInterval15m RadarEmailRoutingTimeseriesGroupGetSpfParamsAggInterval = "15m"
	RadarEmailRoutingTimeseriesGroupGetSpfParamsAggInterval1h  RadarEmailRoutingTimeseriesGroupGetSpfParamsAggInterval = "1h"
	RadarEmailRoutingTimeseriesGroupGetSpfParamsAggInterval1d  RadarEmailRoutingTimeseriesGroupGetSpfParamsAggInterval = "1d"
	RadarEmailRoutingTimeseriesGroupGetSpfParamsAggInterval1w  RadarEmailRoutingTimeseriesGroupGetSpfParamsAggInterval = "1w"
)

func (r RadarEmailRoutingTimeseriesGroupGetSpfParamsAggInterval) IsKnown() bool {
	switch r {
	case RadarEmailRoutingTimeseriesGroupGetSpfParamsAggInterval15m, RadarEmailRoutingTimeseriesGroupGetSpfParamsAggInterval1h, RadarEmailRoutingTimeseriesGroupGetSpfParamsAggInterval1d, RadarEmailRoutingTimeseriesGroupGetSpfParamsAggInterval1w:
		return true
	}
	return false
}

type RadarEmailRoutingTimeseriesGroupGetSpfParamsArc string

const (
	RadarEmailRoutingTimeseriesGroupGetSpfParamsArcPass RadarEmailRoutingTimeseriesGroupGetSpfParamsArc = "PASS"
	RadarEmailRoutingTimeseriesGroupGetSpfParamsArcNone RadarEmailRoutingTimeseriesGroupGetSpfParamsArc = "NONE"
	RadarEmailRoutingTimeseriesGroupGetSpfParamsArcFail RadarEmailRoutingTimeseriesGroupGetSpfParamsArc = "FAIL"
)

func (r RadarEmailRoutingTimeseriesGroupGetSpfParamsArc) IsKnown() bool {
	switch r {
	case RadarEmailRoutingTimeseriesGroupGetSpfParamsArcPass, RadarEmailRoutingTimeseriesGroupGetSpfParamsArcNone, RadarEmailRoutingTimeseriesGroupGetSpfParamsArcFail:
		return true
	}
	return false
}

type RadarEmailRoutingTimeseriesGroupGetSpfParamsDkim string

const (
	RadarEmailRoutingTimeseriesGroupGetSpfParamsDkimPass RadarEmailRoutingTimeseriesGroupGetSpfParamsDkim = "PASS"
	RadarEmailRoutingTimeseriesGroupGetSpfParamsDkimNone RadarEmailRoutingTimeseriesGroupGetSpfParamsDkim = "NONE"
	RadarEmailRoutingTimeseriesGroupGetSpfParamsDkimFail RadarEmailRoutingTimeseriesGroupGetSpfParamsDkim = "FAIL"
)

func (r RadarEmailRoutingTimeseriesGroupGetSpfParamsDkim) IsKnown() bool {
	switch r {
	case RadarEmailRoutingTimeseriesGroupGetSpfParamsDkimPass, RadarEmailRoutingTimeseriesGroupGetSpfParamsDkimNone, RadarEmailRoutingTimeseriesGroupGetSpfParamsDkimFail:
		return true
	}
	return false
}

type RadarEmailRoutingTimeseriesGroupGetSpfParamsDmarc string

const (
	RadarEmailRoutingTimeseriesGroupGetSpfParamsDmarcPass RadarEmailRoutingTimeseriesGroupGetSpfParamsDmarc = "PASS"
	RadarEmailRoutingTimeseriesGroupGetSpfParamsDmarcNone RadarEmailRoutingTimeseriesGroupGetSpfParamsDmarc = "NONE"
	RadarEmailRoutingTimeseriesGroupGetSpfParamsDmarcFail RadarEmailRoutingTimeseriesGroupGetSpfParamsDmarc = "FAIL"
)

func (r RadarEmailRoutingTimeseriesGroupGetSpfParamsDmarc) IsKnown() bool {
	switch r {
	case RadarEmailRoutingTimeseriesGroupGetSpfParamsDmarcPass, RadarEmailRoutingTimeseriesGroupGetSpfParamsDmarcNone, RadarEmailRoutingTimeseriesGroupGetSpfParamsDmarcFail:
		return true
	}
	return false
}

type RadarEmailRoutingTimeseriesGroupGetSpfParamsEncrypted string

const (
	RadarEmailRoutingTimeseriesGroupGetSpfParamsEncryptedEncrypted    RadarEmailRoutingTimeseriesGroupGetSpfParamsEncrypted = "ENCRYPTED"
	RadarEmailRoutingTimeseriesGroupGetSpfParamsEncryptedNotEncrypted RadarEmailRoutingTimeseriesGroupGetSpfParamsEncrypted = "NOT_ENCRYPTED"
)

func (r RadarEmailRoutingTimeseriesGroupGetSpfParamsEncrypted) IsKnown() bool {
	switch r {
	case RadarEmailRoutingTimeseriesGroupGetSpfParamsEncryptedEncrypted, RadarEmailRoutingTimeseriesGroupGetSpfParamsEncryptedNotEncrypted:
		return true
	}
	return false
}

// Format in which results will be returned.
type RadarEmailRoutingTimeseriesGroupGetSpfParamsFormat string

const (
	RadarEmailRoutingTimeseriesGroupGetSpfParamsFormatJson RadarEmailRoutingTimeseriesGroupGetSpfParamsFormat = "JSON"
	RadarEmailRoutingTimeseriesGroupGetSpfParamsFormatCsv  RadarEmailRoutingTimeseriesGroupGetSpfParamsFormat = "CSV"
)

func (r RadarEmailRoutingTimeseriesGroupGetSpfParamsFormat) IsKnown() bool {
	switch r {
	case RadarEmailRoutingTimeseriesGroupGetSpfParamsFormatJson, RadarEmailRoutingTimeseriesGroupGetSpfParamsFormatCsv:
		return true
	}
	return false
}

type RadarEmailRoutingTimeseriesGroupGetSpfParamsIPVersion string

const (
	RadarEmailRoutingTimeseriesGroupGetSpfParamsIPVersionIPv4 RadarEmailRoutingTimeseriesGroupGetSpfParamsIPVersion = "IPv4"
	RadarEmailRoutingTimeseriesGroupGetSpfParamsIPVersionIPv6 RadarEmailRoutingTimeseriesGroupGetSpfParamsIPVersion = "IPv6"
)

func (r RadarEmailRoutingTimeseriesGroupGetSpfParamsIPVersion) IsKnown() bool {
	switch r {
	case RadarEmailRoutingTimeseriesGroupGetSpfParamsIPVersionIPv4, RadarEmailRoutingTimeseriesGroupGetSpfParamsIPVersionIPv6:
		return true
	}
	return false
}
