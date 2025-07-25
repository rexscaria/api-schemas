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

// ZoneSpectrumAnalyticsAggregateService contains methods and other services that
// help with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneSpectrumAnalyticsAggregateService] method instead.
type ZoneSpectrumAnalyticsAggregateService struct {
	Options []option.RequestOption
}

// NewZoneSpectrumAnalyticsAggregateService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZoneSpectrumAnalyticsAggregateService(opts ...option.RequestOption) (r *ZoneSpectrumAnalyticsAggregateService) {
	r = &ZoneSpectrumAnalyticsAggregateService{}
	r.Options = opts
	return
}

// Retrieves analytics aggregated from the last minute of usage on Spectrum
// applications underneath a given zone.
func (r *ZoneSpectrumAnalyticsAggregateService) GetCurrent(ctx context.Context, zoneID string, query ZoneSpectrumAnalyticsAggregateGetCurrentParams, opts ...option.RequestOption) (res *ZoneSpectrumAnalyticsAggregateGetCurrentResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/spectrum/analytics/aggregate/current", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type SpectrumAnalyticsMessageItem struct {
	Code             int64                              `json:"code,required"`
	Message          string                             `json:"message,required"`
	DocumentationURL string                             `json:"documentation_url"`
	Source           SpectrumAnalyticsMessageItemSource `json:"source"`
	JSON             spectrumAnalyticsMessageItemJSON   `json:"-"`
}

// spectrumAnalyticsMessageItemJSON contains the JSON metadata for the struct
// [SpectrumAnalyticsMessageItem]
type spectrumAnalyticsMessageItemJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SpectrumAnalyticsMessageItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r spectrumAnalyticsMessageItemJSON) RawJSON() string {
	return r.raw
}

type SpectrumAnalyticsMessageItemSource struct {
	Pointer string                                 `json:"pointer"`
	JSON    spectrumAnalyticsMessageItemSourceJSON `json:"-"`
}

// spectrumAnalyticsMessageItemSourceJSON contains the JSON metadata for the struct
// [SpectrumAnalyticsMessageItemSource]
type spectrumAnalyticsMessageItemSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpectrumAnalyticsMessageItemSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r spectrumAnalyticsMessageItemSourceJSON) RawJSON() string {
	return r.raw
}

type ZoneSpectrumAnalyticsAggregateGetCurrentResponse struct {
	Errors   []SpectrumAnalyticsMessageItem `json:"errors,required"`
	Messages []SpectrumAnalyticsMessageItem `json:"messages,required"`
	// Whether the API call was successful.
	Success ZoneSpectrumAnalyticsAggregateGetCurrentResponseSuccess  `json:"success,required"`
	Result  []ZoneSpectrumAnalyticsAggregateGetCurrentResponseResult `json:"result"`
	JSON    zoneSpectrumAnalyticsAggregateGetCurrentResponseJSON     `json:"-"`
}

// zoneSpectrumAnalyticsAggregateGetCurrentResponseJSON contains the JSON metadata
// for the struct [ZoneSpectrumAnalyticsAggregateGetCurrentResponse]
type zoneSpectrumAnalyticsAggregateGetCurrentResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSpectrumAnalyticsAggregateGetCurrentResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSpectrumAnalyticsAggregateGetCurrentResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ZoneSpectrumAnalyticsAggregateGetCurrentResponseSuccess bool

const (
	ZoneSpectrumAnalyticsAggregateGetCurrentResponseSuccessTrue ZoneSpectrumAnalyticsAggregateGetCurrentResponseSuccess = true
)

func (r ZoneSpectrumAnalyticsAggregateGetCurrentResponseSuccess) IsKnown() bool {
	switch r {
	case ZoneSpectrumAnalyticsAggregateGetCurrentResponseSuccessTrue:
		return true
	}
	return false
}

type ZoneSpectrumAnalyticsAggregateGetCurrentResponseResult struct {
	// Application identifier.
	AppID string `json:"appID,required"`
	// Number of bytes sent
	BytesEgress float64 `json:"bytesEgress,required"`
	// Number of bytes received
	BytesIngress float64 `json:"bytesIngress,required"`
	// Number of connections
	Connections float64 `json:"connections,required"`
	// Average duration of connections
	DurationAvg float64                                                    `json:"durationAvg,required"`
	JSON        zoneSpectrumAnalyticsAggregateGetCurrentResponseResultJSON `json:"-"`
}

// zoneSpectrumAnalyticsAggregateGetCurrentResponseResultJSON contains the JSON
// metadata for the struct [ZoneSpectrumAnalyticsAggregateGetCurrentResponseResult]
type zoneSpectrumAnalyticsAggregateGetCurrentResponseResultJSON struct {
	AppID        apijson.Field
	BytesEgress  apijson.Field
	BytesIngress apijson.Field
	Connections  apijson.Field
	DurationAvg  apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZoneSpectrumAnalyticsAggregateGetCurrentResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSpectrumAnalyticsAggregateGetCurrentResponseResultJSON) RawJSON() string {
	return r.raw
}

type ZoneSpectrumAnalyticsAggregateGetCurrentParams struct {
	// Comma-delimited list of Spectrum Application Id(s). If provided, the response
	// will be limited to Spectrum Application Id(s) that match.
	AppID param.Field[string] `query:"appID"`
	// Co-location identifier.
	ColoName param.Field[string] `query:"colo_name"`
}

// URLQuery serializes [ZoneSpectrumAnalyticsAggregateGetCurrentParams]'s query
// parameters as `url.Values`.
func (r ZoneSpectrumAnalyticsAggregateGetCurrentParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
