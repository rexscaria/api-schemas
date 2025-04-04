// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/stainless-sdks/cf-rex-go/internal/apijson"
	"github.com/stainless-sdks/cf-rex-go/internal/param"
	"github.com/stainless-sdks/cf-rex-go/internal/requestconfig"
	"github.com/stainless-sdks/cf-rex-go/option"
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

type BaseCacheRule struct {
	// Identifier of the zone setting.
	ID string `json:"id,required"`
	// Last time this setting was modified.
	ModifiedOn time.Time         `json:"modified_on,nullable" format:"date-time"`
	JSON       baseCacheRuleJSON `json:"-"`
}

// baseCacheRuleJSON contains the JSON metadata for the struct [BaseCacheRule]
type baseCacheRuleJSON struct {
	ID          apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BaseCacheRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r baseCacheRuleJSON) RawJSON() string {
	return r.raw
}

type DeleteResponseSingle struct {
	Result DeleteResponseSingleResult `json:"result"`
	JSON   deleteResponseSingleJSON   `json:"-"`
	APIResponseCacheRules
}

// deleteResponseSingleJSON contains the JSON metadata for the struct
// [DeleteResponseSingle]
type deleteResponseSingleJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeleteResponseSingle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deleteResponseSingleJSON) RawJSON() string {
	return r.raw
}

type DeleteResponseSingleResult struct {
	// The identifier of the caching setting
	ID string `json:"id,required"`
	// Whether the setting is editable
	Editable bool `json:"editable,required"`
	// The time when the setting was last modified
	ModifiedOn time.Time                      `json:"modified_on" format:"date-time"`
	JSON       deleteResponseSingleResultJSON `json:"-"`
}

// deleteResponseSingleResultJSON contains the JSON metadata for the struct
// [DeleteResponseSingleResult]
type deleteResponseSingleResultJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeleteResponseSingleResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deleteResponseSingleResultJSON) RawJSON() string {
	return r.raw
}

type SmartTieredCache struct {
	// ID of the zone setting.
	ID   SmartTieredCacheID   `json:"id"`
	JSON smartTieredCacheJSON `json:"-"`
	BaseCacheRule
}

// smartTieredCacheJSON contains the JSON metadata for the struct
// [SmartTieredCache]
type smartTieredCacheJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SmartTieredCache) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r smartTieredCacheJSON) RawJSON() string {
	return r.raw
}

// ID of the zone setting.
type SmartTieredCacheID string

const (
	SmartTieredCacheIDTieredCacheSmartTopologyEnable SmartTieredCacheID = "tiered_cache_smart_topology_enable"
)

func (r SmartTieredCacheID) IsKnown() bool {
	switch r {
	case SmartTieredCacheIDTieredCacheSmartTopologyEnable:
		return true
	}
	return false
}

type SmartTieredCacheResponseValue struct {
	Result SmartTieredCacheResponseValueResult `json:"result"`
	JSON   smartTieredCacheResponseValueJSON   `json:"-"`
}

// smartTieredCacheResponseValueJSON contains the JSON metadata for the struct
// [SmartTieredCacheResponseValue]
type smartTieredCacheResponseValueJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SmartTieredCacheResponseValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r smartTieredCacheResponseValueJSON) RawJSON() string {
	return r.raw
}

type SmartTieredCacheResponseValueResult struct {
	// Value of the Smart Tiered Cache zone setting.
	Value SmartTieredCacheResponseValueResultValue `json:"value,required"`
	JSON  smartTieredCacheResponseValueResultJSON  `json:"-"`
	SmartTieredCache
}

// smartTieredCacheResponseValueResultJSON contains the JSON metadata for the
// struct [SmartTieredCacheResponseValueResult]
type smartTieredCacheResponseValueResultJSON struct {
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SmartTieredCacheResponseValueResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r smartTieredCacheResponseValueResultJSON) RawJSON() string {
	return r.raw
}

// Value of the Smart Tiered Cache zone setting.
type SmartTieredCacheResponseValueResultValue string

const (
	SmartTieredCacheResponseValueResultValueOn  SmartTieredCacheResponseValueResultValue = "on"
	SmartTieredCacheResponseValueResultValueOff SmartTieredCacheResponseValueResultValue = "off"
)

func (r SmartTieredCacheResponseValueResultValue) IsKnown() bool {
	switch r {
	case SmartTieredCacheResponseValueResultValueOn, SmartTieredCacheResponseValueResultValueOff:
		return true
	}
	return false
}

type ZoneCacheTieredCacheSmartTopologyEnableGetResponse struct {
	JSON zoneCacheTieredCacheSmartTopologyEnableGetResponseJSON `json:"-"`
	CacheRulesZoneCacheSettingsResponse
	SmartTieredCacheResponseValue
}

// zoneCacheTieredCacheSmartTopologyEnableGetResponseJSON contains the JSON
// metadata for the struct [ZoneCacheTieredCacheSmartTopologyEnableGetResponse]
type zoneCacheTieredCacheSmartTopologyEnableGetResponseJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCacheTieredCacheSmartTopologyEnableGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneCacheTieredCacheSmartTopologyEnableGetResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneCacheTieredCacheSmartTopologyEnableUpdateResponse struct {
	JSON zoneCacheTieredCacheSmartTopologyEnableUpdateResponseJSON `json:"-"`
	CacheRulesZoneCacheSettingsResponse
	SmartTieredCacheResponseValue
}

// zoneCacheTieredCacheSmartTopologyEnableUpdateResponseJSON contains the JSON
// metadata for the struct [ZoneCacheTieredCacheSmartTopologyEnableUpdateResponse]
type zoneCacheTieredCacheSmartTopologyEnableUpdateResponseJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCacheTieredCacheSmartTopologyEnableUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneCacheTieredCacheSmartTopologyEnableUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneCacheTieredCacheSmartTopologyEnableDeleteResponse struct {
	Result SmartTieredCache                                          `json:"result"`
	JSON   zoneCacheTieredCacheSmartTopologyEnableDeleteResponseJSON `json:"-"`
	DeleteResponseSingle
}

// zoneCacheTieredCacheSmartTopologyEnableDeleteResponseJSON contains the JSON
// metadata for the struct [ZoneCacheTieredCacheSmartTopologyEnableDeleteResponse]
type zoneCacheTieredCacheSmartTopologyEnableDeleteResponseJSON struct {
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
