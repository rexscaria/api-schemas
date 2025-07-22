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

// ZoneSpeedAPIScheduleService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneSpeedAPIScheduleService] method instead.
type ZoneSpeedAPIScheduleService struct {
	Options []option.RequestOption
}

// NewZoneSpeedAPIScheduleService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneSpeedAPIScheduleService(opts ...option.RequestOption) (r *ZoneSpeedAPIScheduleService) {
	r = &ZoneSpeedAPIScheduleService{}
	r.Options = opts
	return
}

// Creates a scheduled test for a page.
func (r *ZoneSpeedAPIScheduleService) New(ctx context.Context, zoneID string, url string, body ZoneSpeedAPIScheduleNewParams, opts ...option.RequestOption) (res *ZoneSpeedAPIScheduleNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if url == "" {
		err = errors.New("missing required url parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/speed_api/schedule/%s", zoneID, url)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Deletes a scheduled test for a page.
func (r *ZoneSpeedAPIScheduleService) Delete(ctx context.Context, zoneID string, url string, body ZoneSpeedAPIScheduleDeleteParams, opts ...option.RequestOption) (res *ObservatoryCountResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if url == "" {
		err = errors.New("missing required url parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/speed_api/schedule/%s", zoneID, url)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

// Retrieves the test schedule for a page in a specific region.
func (r *ZoneSpeedAPIScheduleService) Get(ctx context.Context, zoneID string, url string, query ZoneSpeedAPIScheduleGetParams, opts ...option.RequestOption) (res *ZoneSpeedAPIScheduleGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if url == "" {
		err = errors.New("missing required url parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/speed_api/schedule/%s", zoneID, url)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// The test schedule.
type ObservatorySchedule struct {
	// The frequency of the test.
	Frequency ObservatoryScheduleFrequency `json:"frequency"`
	// A test region.
	Region ObservatoryRegion `json:"region"`
	// A URL.
	URL  string                  `json:"url"`
	JSON observatoryScheduleJSON `json:"-"`
}

// observatoryScheduleJSON contains the JSON metadata for the struct
// [ObservatorySchedule]
type observatoryScheduleJSON struct {
	Frequency   apijson.Field
	Region      apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservatorySchedule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observatoryScheduleJSON) RawJSON() string {
	return r.raw
}

type ZoneSpeedAPIScheduleNewResponse struct {
	Result ZoneSpeedAPIScheduleNewResponseResult `json:"result"`
	JSON   zoneSpeedAPIScheduleNewResponseJSON   `json:"-"`
	ObservatoryAPIResponseSingle
}

// zoneSpeedAPIScheduleNewResponseJSON contains the JSON metadata for the struct
// [ZoneSpeedAPIScheduleNewResponse]
type zoneSpeedAPIScheduleNewResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSpeedAPIScheduleNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSpeedAPIScheduleNewResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneSpeedAPIScheduleNewResponseResult struct {
	// The test schedule.
	Schedule ObservatorySchedule                       `json:"schedule"`
	Test     ObservatoryPageTest                       `json:"test"`
	JSON     zoneSpeedAPIScheduleNewResponseResultJSON `json:"-"`
}

// zoneSpeedAPIScheduleNewResponseResultJSON contains the JSON metadata for the
// struct [ZoneSpeedAPIScheduleNewResponseResult]
type zoneSpeedAPIScheduleNewResponseResultJSON struct {
	Schedule    apijson.Field
	Test        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSpeedAPIScheduleNewResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSpeedAPIScheduleNewResponseResultJSON) RawJSON() string {
	return r.raw
}

type ZoneSpeedAPIScheduleGetResponse struct {
	// The test schedule.
	Result ObservatorySchedule                 `json:"result"`
	JSON   zoneSpeedAPIScheduleGetResponseJSON `json:"-"`
	ObservatoryAPIResponseSingle
}

// zoneSpeedAPIScheduleGetResponseJSON contains the JSON metadata for the struct
// [ZoneSpeedAPIScheduleGetResponse]
type zoneSpeedAPIScheduleGetResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSpeedAPIScheduleGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSpeedAPIScheduleGetResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneSpeedAPIScheduleNewParams struct {
	// A test region.
	Region param.Field[ObservatoryRegion] `query:"region"`
}

// URLQuery serializes [ZoneSpeedAPIScheduleNewParams]'s query parameters as
// `url.Values`.
func (r ZoneSpeedAPIScheduleNewParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ZoneSpeedAPIScheduleDeleteParams struct {
	// A test region.
	Region param.Field[ObservatoryRegion] `query:"region"`
}

// URLQuery serializes [ZoneSpeedAPIScheduleDeleteParams]'s query parameters as
// `url.Values`.
func (r ZoneSpeedAPIScheduleDeleteParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ZoneSpeedAPIScheduleGetParams struct {
	// A test region.
	Region param.Field[ObservatoryRegion] `query:"region"`
}

// URLQuery serializes [ZoneSpeedAPIScheduleGetParams]'s query parameters as
// `url.Values`.
func (r ZoneSpeedAPIScheduleGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
