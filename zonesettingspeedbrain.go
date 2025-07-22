// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/cf-rex-go/internal/apijson"
	"github.com/stainless-sdks/cf-rex-go/internal/param"
	"github.com/stainless-sdks/cf-rex-go/internal/requestconfig"
	"github.com/stainless-sdks/cf-rex-go/option"
)

// ZoneSettingSpeedBrainService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneSettingSpeedBrainService] method instead.
type ZoneSettingSpeedBrainService struct {
	Options []option.RequestOption
}

// NewZoneSettingSpeedBrainService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneSettingSpeedBrainService(opts ...option.RequestOption) (r *ZoneSettingSpeedBrainService) {
	r = &ZoneSettingSpeedBrainService{}
	r.Options = opts
	return
}

// Speed Brain lets compatible browsers speculate on content which can be
// prefetched or preloaded, making website navigation faster. Refer to the
// Cloudflare Speed Brain documentation for more information.
func (r *ZoneSettingSpeedBrainService) Get(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *ZoneSettingSpeedBrainGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/settings/speed_brain", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Speed Brain lets compatible browsers speculate on content which can be
// prefetched or preloaded, making website navigation faster. Refer to the
// Cloudflare Speed Brain documentation for more information.
func (r *ZoneSettingSpeedBrainService) Update(ctx context.Context, zoneID string, body ZoneSettingSpeedBrainUpdateParams, opts ...option.RequestOption) (res *ZoneSettingSpeedBrainUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/settings/speed_brain", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

type SpeedCloudflareSpeedBrainResponse struct {
	// Whether the feature is enabled or disabled. Defaults to "on" for Free plans,
	// otherwise defaults to "off".
	Value interface{}                           `json:"value"`
	JSON  speedCloudflareSpeedBrainResponseJSON `json:"-"`
	SpeedBase
}

// speedCloudflareSpeedBrainResponseJSON contains the JSON metadata for the struct
// [SpeedCloudflareSpeedBrainResponse]
type speedCloudflareSpeedBrainResponseJSON struct {
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpeedCloudflareSpeedBrainResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r speedCloudflareSpeedBrainResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneSettingSpeedBrainGetResponse struct {
	Result SpeedCloudflareSpeedBrainResponse    `json:"result"`
	JSON   zoneSettingSpeedBrainGetResponseJSON `json:"-"`
	SpeedAPIResponseCommon
}

// zoneSettingSpeedBrainGetResponseJSON contains the JSON metadata for the struct
// [ZoneSettingSpeedBrainGetResponse]
type zoneSettingSpeedBrainGetResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingSpeedBrainGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingSpeedBrainGetResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneSettingSpeedBrainUpdateResponse struct {
	Result SpeedCloudflareSpeedBrainResponse       `json:"result"`
	JSON   zoneSettingSpeedBrainUpdateResponseJSON `json:"-"`
	SpeedAPIResponseCommon
}

// zoneSettingSpeedBrainUpdateResponseJSON contains the JSON metadata for the
// struct [ZoneSettingSpeedBrainUpdateResponse]
type zoneSettingSpeedBrainUpdateResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingSpeedBrainUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingSpeedBrainUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneSettingSpeedBrainUpdateParams struct {
	// Whether the feature is enabled or disabled.
	Value param.Field[ZoneSettingSpeedBrainUpdateParamsValue] `json:"value,required"`
}

func (r ZoneSettingSpeedBrainUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Whether the feature is enabled or disabled.
type ZoneSettingSpeedBrainUpdateParamsValue string

const (
	ZoneSettingSpeedBrainUpdateParamsValueOn  ZoneSettingSpeedBrainUpdateParamsValue = "on"
	ZoneSettingSpeedBrainUpdateParamsValueOff ZoneSettingSpeedBrainUpdateParamsValue = "off"
)

func (r ZoneSettingSpeedBrainUpdateParamsValue) IsKnown() bool {
	switch r {
	case ZoneSettingSpeedBrainUpdateParamsValueOn, ZoneSettingSpeedBrainUpdateParamsValueOff:
		return true
	}
	return false
}
