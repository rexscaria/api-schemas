// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
)

// ZoneWaitingRoomSettingService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneWaitingRoomSettingService] method instead.
type ZoneWaitingRoomSettingService struct {
	Options []option.RequestOption
}

// NewZoneWaitingRoomSettingService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneWaitingRoomSettingService(opts ...option.RequestOption) (r *ZoneWaitingRoomSettingService) {
	r = &ZoneWaitingRoomSettingService{}
	r.Options = opts
	return
}

// Get zone-level Waiting Room settings
func (r *ZoneWaitingRoomSettingService) Get(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *ZoneSettingsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/waiting_rooms/settings", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Update zone-level Waiting Room settings
func (r *ZoneWaitingRoomSettingService) Update(ctx context.Context, zoneID string, body ZoneWaitingRoomSettingUpdateParams, opts ...option.RequestOption) (res *ZoneSettingsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/waiting_rooms/settings", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

// Patch zone-level Waiting Room settings
func (r *ZoneWaitingRoomSettingService) Patch(ctx context.Context, zoneID string, body ZoneWaitingRoomSettingPatchParams, opts ...option.RequestOption) (res *ZoneSettingsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/waiting_rooms/settings", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

type ZoneSettingsParam struct {
	// Whether to allow verified search engine crawlers to bypass all waiting rooms on
	// this zone. Verified search engine crawlers will not be tracked or counted by the
	// waiting room system, and will not appear in waiting room analytics.
	SearchEngineCrawlerBypass param.Field[bool] `json:"search_engine_crawler_bypass"`
}

func (r ZoneSettingsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneSettingsResponse struct {
	Result ZoneSettingsResponseResult `json:"result" api:"required"`
	JSON   zoneSettingsResponseJSON   `json:"-"`
}

// zoneSettingsResponseJSON contains the JSON metadata for the struct
// [ZoneSettingsResponse]
type zoneSettingsResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingsResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneSettingsResponseResult struct {
	// Whether to allow verified search engine crawlers to bypass all waiting rooms on
	// this zone. Verified search engine crawlers will not be tracked or counted by the
	// waiting room system, and will not appear in waiting room analytics.
	SearchEngineCrawlerBypass bool                           `json:"search_engine_crawler_bypass" api:"required"`
	JSON                      zoneSettingsResponseResultJSON `json:"-"`
}

// zoneSettingsResponseResultJSON contains the JSON metadata for the struct
// [ZoneSettingsResponseResult]
type zoneSettingsResponseResultJSON struct {
	SearchEngineCrawlerBypass apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *ZoneSettingsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingsResponseResultJSON) RawJSON() string {
	return r.raw
}

type ZoneWaitingRoomSettingUpdateParams struct {
	ZoneSettings ZoneSettingsParam `json:"zone_settings" api:"required"`
}

func (r ZoneWaitingRoomSettingUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.ZoneSettings)
}

type ZoneWaitingRoomSettingPatchParams struct {
	ZoneSettings ZoneSettingsParam `json:"zone_settings" api:"required"`
}

func (r ZoneWaitingRoomSettingPatchParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.ZoneSettings)
}
