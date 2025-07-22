// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/cf-rex-go/internal/apijson"
	"github.com/stainless-sdks/cf-rex-go/internal/requestconfig"
	"github.com/stainless-sdks/cf-rex-go/option"
)

// ZoneSpeedAPIService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneSpeedAPIService] method instead.
type ZoneSpeedAPIService struct {
	Options  []option.RequestOption
	Pages    *ZoneSpeedAPIPageService
	Schedule *ZoneSpeedAPIScheduleService
}

// NewZoneSpeedAPIService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewZoneSpeedAPIService(opts ...option.RequestOption) (r *ZoneSpeedAPIService) {
	r = &ZoneSpeedAPIService{}
	r.Options = opts
	r.Pages = NewZoneSpeedAPIPageService(opts...)
	r.Schedule = NewZoneSpeedAPIScheduleService(opts...)
	return
}

// Retrieves quota for all plans, as well as the current zone quota.
func (r *ZoneSpeedAPIService) GetAvailabilities(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *ZoneSpeedAPIGetAvailabilitiesResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/speed_api/availabilities", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ObservatoryAPIResponseSingle struct {
	Errors   []ObservatoryMessagesItem `json:"errors,required"`
	Messages []ObservatoryMessagesItem `json:"messages,required"`
	// Whether the API call was successful.
	Success bool                             `json:"success,required"`
	JSON    observatoryAPIResponseSingleJSON `json:"-"`
}

// observatoryAPIResponseSingleJSON contains the JSON metadata for the struct
// [ObservatoryAPIResponseSingle]
type observatoryAPIResponseSingleJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservatoryAPIResponseSingle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observatoryAPIResponseSingleJSON) RawJSON() string {
	return r.raw
}

type ObservatoryMessagesItem struct {
	Code    int64                       `json:"code,required"`
	Message string                      `json:"message,required"`
	JSON    observatoryMessagesItemJSON `json:"-"`
}

// observatoryMessagesItemJSON contains the JSON metadata for the struct
// [ObservatoryMessagesItem]
type observatoryMessagesItemJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservatoryMessagesItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observatoryMessagesItemJSON) RawJSON() string {
	return r.raw
}

// Counts per account plan.
type ObservatoryPlanPropertiesInfo struct {
	Business   int64                             `json:"business"`
	Enterprise int64                             `json:"enterprise"`
	Free       int64                             `json:"free"`
	Pro        int64                             `json:"pro"`
	JSON       observatoryPlanPropertiesInfoJSON `json:"-"`
}

// observatoryPlanPropertiesInfoJSON contains the JSON metadata for the struct
// [ObservatoryPlanPropertiesInfo]
type observatoryPlanPropertiesInfoJSON struct {
	Business    apijson.Field
	Enterprise  apijson.Field
	Free        apijson.Field
	Pro         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservatoryPlanPropertiesInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observatoryPlanPropertiesInfoJSON) RawJSON() string {
	return r.raw
}

type ZoneSpeedAPIGetAvailabilitiesResponse struct {
	Result ZoneSpeedAPIGetAvailabilitiesResponseResult `json:"result"`
	JSON   zoneSpeedAPIGetAvailabilitiesResponseJSON   `json:"-"`
	ObservatoryAPIResponseSingle
}

// zoneSpeedAPIGetAvailabilitiesResponseJSON contains the JSON metadata for the
// struct [ZoneSpeedAPIGetAvailabilitiesResponse]
type zoneSpeedAPIGetAvailabilitiesResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSpeedAPIGetAvailabilitiesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSpeedAPIGetAvailabilitiesResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneSpeedAPIGetAvailabilitiesResponseResult struct {
	Quota   ZoneSpeedAPIGetAvailabilitiesResponseResultQuota `json:"quota"`
	Regions []ObservatoryLabeledRegion                       `json:"regions"`
	// Available regions.
	RegionsPerPlan ZoneSpeedAPIGetAvailabilitiesResponseResultRegionsPerPlan `json:"regionsPerPlan"`
	JSON           zoneSpeedAPIGetAvailabilitiesResponseResultJSON           `json:"-"`
}

// zoneSpeedAPIGetAvailabilitiesResponseResultJSON contains the JSON metadata for
// the struct [ZoneSpeedAPIGetAvailabilitiesResponseResult]
type zoneSpeedAPIGetAvailabilitiesResponseResultJSON struct {
	Quota          apijson.Field
	Regions        apijson.Field
	RegionsPerPlan apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ZoneSpeedAPIGetAvailabilitiesResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSpeedAPIGetAvailabilitiesResponseResultJSON) RawJSON() string {
	return r.raw
}

type ZoneSpeedAPIGetAvailabilitiesResponseResultQuota struct {
	// Cloudflare plan.
	Plan string `json:"plan"`
	// The number of tests available per plan.
	QuotasPerPlan ZoneSpeedAPIGetAvailabilitiesResponseResultQuotaQuotasPerPlan `json:"quotasPerPlan"`
	// The number of remaining schedules available.
	RemainingSchedules float64 `json:"remainingSchedules"`
	// The number of remaining tests available.
	RemainingTests float64 `json:"remainingTests"`
	// The number of schedules available per plan.
	ScheduleQuotasPerPlan ZoneSpeedAPIGetAvailabilitiesResponseResultQuotaScheduleQuotasPerPlan `json:"scheduleQuotasPerPlan"`
	JSON                  zoneSpeedAPIGetAvailabilitiesResponseResultQuotaJSON                  `json:"-"`
}

// zoneSpeedAPIGetAvailabilitiesResponseResultQuotaJSON contains the JSON metadata
// for the struct [ZoneSpeedAPIGetAvailabilitiesResponseResultQuota]
type zoneSpeedAPIGetAvailabilitiesResponseResultQuotaJSON struct {
	Plan                  apijson.Field
	QuotasPerPlan         apijson.Field
	RemainingSchedules    apijson.Field
	RemainingTests        apijson.Field
	ScheduleQuotasPerPlan apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *ZoneSpeedAPIGetAvailabilitiesResponseResultQuota) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSpeedAPIGetAvailabilitiesResponseResultQuotaJSON) RawJSON() string {
	return r.raw
}

// The number of tests available per plan.
type ZoneSpeedAPIGetAvailabilitiesResponseResultQuotaQuotasPerPlan struct {
	// Counts per account plan.
	Value ObservatoryPlanPropertiesInfo                                     `json:"value"`
	JSON  zoneSpeedAPIGetAvailabilitiesResponseResultQuotaQuotasPerPlanJSON `json:"-"`
}

// zoneSpeedAPIGetAvailabilitiesResponseResultQuotaQuotasPerPlanJSON contains the
// JSON metadata for the struct
// [ZoneSpeedAPIGetAvailabilitiesResponseResultQuotaQuotasPerPlan]
type zoneSpeedAPIGetAvailabilitiesResponseResultQuotaQuotasPerPlanJSON struct {
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSpeedAPIGetAvailabilitiesResponseResultQuotaQuotasPerPlan) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSpeedAPIGetAvailabilitiesResponseResultQuotaQuotasPerPlanJSON) RawJSON() string {
	return r.raw
}

// The number of schedules available per plan.
type ZoneSpeedAPIGetAvailabilitiesResponseResultQuotaScheduleQuotasPerPlan struct {
	// Counts per account plan.
	Value ObservatoryPlanPropertiesInfo                                             `json:"value"`
	JSON  zoneSpeedAPIGetAvailabilitiesResponseResultQuotaScheduleQuotasPerPlanJSON `json:"-"`
}

// zoneSpeedAPIGetAvailabilitiesResponseResultQuotaScheduleQuotasPerPlanJSON
// contains the JSON metadata for the struct
// [ZoneSpeedAPIGetAvailabilitiesResponseResultQuotaScheduleQuotasPerPlan]
type zoneSpeedAPIGetAvailabilitiesResponseResultQuotaScheduleQuotasPerPlanJSON struct {
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSpeedAPIGetAvailabilitiesResponseResultQuotaScheduleQuotasPerPlan) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSpeedAPIGetAvailabilitiesResponseResultQuotaScheduleQuotasPerPlanJSON) RawJSON() string {
	return r.raw
}

// Available regions.
type ZoneSpeedAPIGetAvailabilitiesResponseResultRegionsPerPlan struct {
	Business   []ObservatoryLabeledRegion                                    `json:"business"`
	Enterprise []ObservatoryLabeledRegion                                    `json:"enterprise"`
	Free       []ObservatoryLabeledRegion                                    `json:"free"`
	Pro        []ObservatoryLabeledRegion                                    `json:"pro"`
	JSON       zoneSpeedAPIGetAvailabilitiesResponseResultRegionsPerPlanJSON `json:"-"`
}

// zoneSpeedAPIGetAvailabilitiesResponseResultRegionsPerPlanJSON contains the JSON
// metadata for the struct
// [ZoneSpeedAPIGetAvailabilitiesResponseResultRegionsPerPlan]
type zoneSpeedAPIGetAvailabilitiesResponseResultRegionsPerPlanJSON struct {
	Business    apijson.Field
	Enterprise  apijson.Field
	Free        apijson.Field
	Pro         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSpeedAPIGetAvailabilitiesResponseResultRegionsPerPlan) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSpeedAPIGetAvailabilitiesResponseResultRegionsPerPlanJSON) RawJSON() string {
	return r.raw
}
