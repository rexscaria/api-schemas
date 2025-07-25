// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
)

// ZoneCacheTieredCacheSmartTopologyEnableService contains methods and other
// services that help with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneCacheTieredCacheSmartTopologyEnableService] method instead.
type ZoneCacheTieredCacheSmartTopologyEnableService struct {
	Options []option.RequestOption
}

// NewZoneCacheTieredCacheSmartTopologyEnableService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewZoneCacheTieredCacheSmartTopologyEnableService(opts ...option.RequestOption) (r *ZoneCacheTieredCacheSmartTopologyEnableService) {
	r = &ZoneCacheTieredCacheSmartTopologyEnableService{}
	r.Options = opts
	return
}

// Smart Tiered Cache dynamically selects the single closest upper tier for each of
// your website’s origins with no configuration required, using our in-house
// performance and routing data. Cloudflare collects latency data for each request
// to an origin, and uses the latency data to determine how well any upper-tier
// data center is connected with an origin. As a result, Cloudflare can select the
// data center with the lowest latency to be the upper-tier for an origin.
func (r *ZoneCacheTieredCacheSmartTopologyEnableService) Get(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *ZoneCacheTieredCacheSmartTopologyEnableGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/cache/tiered_cache_smart_topology_enable", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Smart Tiered Cache dynamically selects the single closest upper tier for each of
// your website’s origins with no configuration required, using our in-house
// performance and routing data. Cloudflare collects latency data for each request
// to an origin, and uses the latency data to determine how well any upper-tier
// data center is connected with an origin. As a result, Cloudflare can select the
// data center with the lowest latency to be the upper-tier for an origin.
func (r *ZoneCacheTieredCacheSmartTopologyEnableService) Update(ctx context.Context, zoneID string, body ZoneCacheTieredCacheSmartTopologyEnableUpdateParams, opts ...option.RequestOption) (res *ZoneCacheTieredCacheSmartTopologyEnableUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/cache/tiered_cache_smart_topology_enable", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Smart Tiered Cache dynamically selects the single closest upper tier for each of
// your website’s origins with no configuration required, using our in-house
// performance and routing data. Cloudflare collects latency data for each request
// to an origin, and uses the latency data to determine how well any upper-tier
// data center is connected with an origin. As a result, Cloudflare can select the
// data center with the lowest latency to be the upper-tier for an origin.
func (r *ZoneCacheTieredCacheSmartTopologyEnableService) Delete(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *ZoneCacheTieredCacheSmartTopologyEnableDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/cache/tiered_cache_smart_topology_enable", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type ZoneCacheTieredCacheSmartTopologyEnableGetResponse struct {
	Errors   []MessagesCacheRulesItem `json:"errors,required"`
	Messages []MessagesCacheRulesItem `json:"messages,required"`
	// Whether the API call was successful
	Success ZoneCacheTieredCacheSmartTopologyEnableGetResponseSuccess `json:"success,required"`
	Result  ZoneCacheTieredCacheSmartTopologyEnableGetResponseResult  `json:"result"`
	JSON    zoneCacheTieredCacheSmartTopologyEnableGetResponseJSON    `json:"-"`
}

// zoneCacheTieredCacheSmartTopologyEnableGetResponseJSON contains the JSON
// metadata for the struct [ZoneCacheTieredCacheSmartTopologyEnableGetResponse]
type zoneCacheTieredCacheSmartTopologyEnableGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCacheTieredCacheSmartTopologyEnableGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneCacheTieredCacheSmartTopologyEnableGetResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ZoneCacheTieredCacheSmartTopologyEnableGetResponseSuccess bool

const (
	ZoneCacheTieredCacheSmartTopologyEnableGetResponseSuccessTrue ZoneCacheTieredCacheSmartTopologyEnableGetResponseSuccess = true
)

func (r ZoneCacheTieredCacheSmartTopologyEnableGetResponseSuccess) IsKnown() bool {
	switch r {
	case ZoneCacheTieredCacheSmartTopologyEnableGetResponseSuccessTrue:
		return true
	}
	return false
}

type ZoneCacheTieredCacheSmartTopologyEnableGetResponseResult struct {
	// ID of the zone setting.
	ID ZoneCacheTieredCacheSmartTopologyEnableGetResponseResultID `json:"id,required"`
	// Whether the setting is editable
	Editable bool `json:"editable,required"`
	// The value of the feature
	Value ZoneCacheTieredCacheSmartTopologyEnableGetResponseResultValue `json:"value,required"`
	// Last time this setting was modified.
	ModifiedOn time.Time                                                    `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneCacheTieredCacheSmartTopologyEnableGetResponseResultJSON `json:"-"`
}

// zoneCacheTieredCacheSmartTopologyEnableGetResponseResultJSON contains the JSON
// metadata for the struct
// [ZoneCacheTieredCacheSmartTopologyEnableGetResponseResult]
type zoneCacheTieredCacheSmartTopologyEnableGetResponseResultJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	Value       apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCacheTieredCacheSmartTopologyEnableGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneCacheTieredCacheSmartTopologyEnableGetResponseResultJSON) RawJSON() string {
	return r.raw
}

// ID of the zone setting.
type ZoneCacheTieredCacheSmartTopologyEnableGetResponseResultID string

const (
	ZoneCacheTieredCacheSmartTopologyEnableGetResponseResultIDTieredCacheSmartTopologyEnable ZoneCacheTieredCacheSmartTopologyEnableGetResponseResultID = "tiered_cache_smart_topology_enable"
)

func (r ZoneCacheTieredCacheSmartTopologyEnableGetResponseResultID) IsKnown() bool {
	switch r {
	case ZoneCacheTieredCacheSmartTopologyEnableGetResponseResultIDTieredCacheSmartTopologyEnable:
		return true
	}
	return false
}

// The value of the feature
type ZoneCacheTieredCacheSmartTopologyEnableGetResponseResultValue string

const (
	ZoneCacheTieredCacheSmartTopologyEnableGetResponseResultValueOn  ZoneCacheTieredCacheSmartTopologyEnableGetResponseResultValue = "on"
	ZoneCacheTieredCacheSmartTopologyEnableGetResponseResultValueOff ZoneCacheTieredCacheSmartTopologyEnableGetResponseResultValue = "off"
)

func (r ZoneCacheTieredCacheSmartTopologyEnableGetResponseResultValue) IsKnown() bool {
	switch r {
	case ZoneCacheTieredCacheSmartTopologyEnableGetResponseResultValueOn, ZoneCacheTieredCacheSmartTopologyEnableGetResponseResultValueOff:
		return true
	}
	return false
}

type ZoneCacheTieredCacheSmartTopologyEnableUpdateResponse struct {
	Errors   []MessagesCacheRulesItem `json:"errors,required"`
	Messages []MessagesCacheRulesItem `json:"messages,required"`
	// Whether the API call was successful
	Success ZoneCacheTieredCacheSmartTopologyEnableUpdateResponseSuccess `json:"success,required"`
	Result  ZoneCacheTieredCacheSmartTopologyEnableUpdateResponseResult  `json:"result"`
	JSON    zoneCacheTieredCacheSmartTopologyEnableUpdateResponseJSON    `json:"-"`
}

// zoneCacheTieredCacheSmartTopologyEnableUpdateResponseJSON contains the JSON
// metadata for the struct [ZoneCacheTieredCacheSmartTopologyEnableUpdateResponse]
type zoneCacheTieredCacheSmartTopologyEnableUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCacheTieredCacheSmartTopologyEnableUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneCacheTieredCacheSmartTopologyEnableUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ZoneCacheTieredCacheSmartTopologyEnableUpdateResponseSuccess bool

const (
	ZoneCacheTieredCacheSmartTopologyEnableUpdateResponseSuccessTrue ZoneCacheTieredCacheSmartTopologyEnableUpdateResponseSuccess = true
)

func (r ZoneCacheTieredCacheSmartTopologyEnableUpdateResponseSuccess) IsKnown() bool {
	switch r {
	case ZoneCacheTieredCacheSmartTopologyEnableUpdateResponseSuccessTrue:
		return true
	}
	return false
}

type ZoneCacheTieredCacheSmartTopologyEnableUpdateResponseResult struct {
	// ID of the zone setting.
	ID ZoneCacheTieredCacheSmartTopologyEnableUpdateResponseResultID `json:"id,required"`
	// Whether the setting is editable
	Editable bool `json:"editable,required"`
	// The value of the feature
	Value ZoneCacheTieredCacheSmartTopologyEnableUpdateResponseResultValue `json:"value,required"`
	// Last time this setting was modified.
	ModifiedOn time.Time                                                       `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneCacheTieredCacheSmartTopologyEnableUpdateResponseResultJSON `json:"-"`
}

// zoneCacheTieredCacheSmartTopologyEnableUpdateResponseResultJSON contains the
// JSON metadata for the struct
// [ZoneCacheTieredCacheSmartTopologyEnableUpdateResponseResult]
type zoneCacheTieredCacheSmartTopologyEnableUpdateResponseResultJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	Value       apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCacheTieredCacheSmartTopologyEnableUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneCacheTieredCacheSmartTopologyEnableUpdateResponseResultJSON) RawJSON() string {
	return r.raw
}

// ID of the zone setting.
type ZoneCacheTieredCacheSmartTopologyEnableUpdateResponseResultID string

const (
	ZoneCacheTieredCacheSmartTopologyEnableUpdateResponseResultIDTieredCacheSmartTopologyEnable ZoneCacheTieredCacheSmartTopologyEnableUpdateResponseResultID = "tiered_cache_smart_topology_enable"
)

func (r ZoneCacheTieredCacheSmartTopologyEnableUpdateResponseResultID) IsKnown() bool {
	switch r {
	case ZoneCacheTieredCacheSmartTopologyEnableUpdateResponseResultIDTieredCacheSmartTopologyEnable:
		return true
	}
	return false
}

// The value of the feature
type ZoneCacheTieredCacheSmartTopologyEnableUpdateResponseResultValue string

const (
	ZoneCacheTieredCacheSmartTopologyEnableUpdateResponseResultValueOn  ZoneCacheTieredCacheSmartTopologyEnableUpdateResponseResultValue = "on"
	ZoneCacheTieredCacheSmartTopologyEnableUpdateResponseResultValueOff ZoneCacheTieredCacheSmartTopologyEnableUpdateResponseResultValue = "off"
)

func (r ZoneCacheTieredCacheSmartTopologyEnableUpdateResponseResultValue) IsKnown() bool {
	switch r {
	case ZoneCacheTieredCacheSmartTopologyEnableUpdateResponseResultValueOn, ZoneCacheTieredCacheSmartTopologyEnableUpdateResponseResultValueOff:
		return true
	}
	return false
}

type ZoneCacheTieredCacheSmartTopologyEnableDeleteResponse struct {
	Errors   []MessagesCacheRulesItem `json:"errors,required"`
	Messages []MessagesCacheRulesItem `json:"messages,required"`
	// Whether the API call was successful
	Success ZoneCacheTieredCacheSmartTopologyEnableDeleteResponseSuccess `json:"success,required"`
	Result  ZoneCacheTieredCacheSmartTopologyEnableDeleteResponseResult  `json:"result"`
	JSON    zoneCacheTieredCacheSmartTopologyEnableDeleteResponseJSON    `json:"-"`
}

// zoneCacheTieredCacheSmartTopologyEnableDeleteResponseJSON contains the JSON
// metadata for the struct [ZoneCacheTieredCacheSmartTopologyEnableDeleteResponse]
type zoneCacheTieredCacheSmartTopologyEnableDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCacheTieredCacheSmartTopologyEnableDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneCacheTieredCacheSmartTopologyEnableDeleteResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ZoneCacheTieredCacheSmartTopologyEnableDeleteResponseSuccess bool

const (
	ZoneCacheTieredCacheSmartTopologyEnableDeleteResponseSuccessTrue ZoneCacheTieredCacheSmartTopologyEnableDeleteResponseSuccess = true
)

func (r ZoneCacheTieredCacheSmartTopologyEnableDeleteResponseSuccess) IsKnown() bool {
	switch r {
	case ZoneCacheTieredCacheSmartTopologyEnableDeleteResponseSuccessTrue:
		return true
	}
	return false
}

type ZoneCacheTieredCacheSmartTopologyEnableDeleteResponseResult struct {
	// ID of the zone setting.
	ID ZoneCacheTieredCacheSmartTopologyEnableDeleteResponseResultID `json:"id,required"`
	// Whether the setting is editable
	Editable bool `json:"editable,required"`
	// Last time this setting was modified.
	ModifiedOn time.Time                                                       `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneCacheTieredCacheSmartTopologyEnableDeleteResponseResultJSON `json:"-"`
}

// zoneCacheTieredCacheSmartTopologyEnableDeleteResponseResultJSON contains the
// JSON metadata for the struct
// [ZoneCacheTieredCacheSmartTopologyEnableDeleteResponseResult]
type zoneCacheTieredCacheSmartTopologyEnableDeleteResponseResultJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCacheTieredCacheSmartTopologyEnableDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneCacheTieredCacheSmartTopologyEnableDeleteResponseResultJSON) RawJSON() string {
	return r.raw
}

// ID of the zone setting.
type ZoneCacheTieredCacheSmartTopologyEnableDeleteResponseResultID string

const (
	ZoneCacheTieredCacheSmartTopologyEnableDeleteResponseResultIDTieredCacheSmartTopologyEnable ZoneCacheTieredCacheSmartTopologyEnableDeleteResponseResultID = "tiered_cache_smart_topology_enable"
)

func (r ZoneCacheTieredCacheSmartTopologyEnableDeleteResponseResultID) IsKnown() bool {
	switch r {
	case ZoneCacheTieredCacheSmartTopologyEnableDeleteResponseResultIDTieredCacheSmartTopologyEnable:
		return true
	}
	return false
}

type ZoneCacheTieredCacheSmartTopologyEnableUpdateParams struct {
	// Enable or disable the Smart Tiered Cache
	Value param.Field[ZoneCacheTieredCacheSmartTopologyEnableUpdateParamsValue] `json:"value,required"`
}

func (r ZoneCacheTieredCacheSmartTopologyEnableUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enable or disable the Smart Tiered Cache
type ZoneCacheTieredCacheSmartTopologyEnableUpdateParamsValue string

const (
	ZoneCacheTieredCacheSmartTopologyEnableUpdateParamsValueOn  ZoneCacheTieredCacheSmartTopologyEnableUpdateParamsValue = "on"
	ZoneCacheTieredCacheSmartTopologyEnableUpdateParamsValueOff ZoneCacheTieredCacheSmartTopologyEnableUpdateParamsValue = "off"
)

func (r ZoneCacheTieredCacheSmartTopologyEnableUpdateParamsValue) IsKnown() bool {
	switch r {
	case ZoneCacheTieredCacheSmartTopologyEnableUpdateParamsValueOn, ZoneCacheTieredCacheSmartTopologyEnableUpdateParamsValueOff:
		return true
	}
	return false
}
