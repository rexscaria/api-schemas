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

// ZoneArgoTieredCachingService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneArgoTieredCachingService] method instead.
type ZoneArgoTieredCachingService struct {
	Options []option.RequestOption
}

// NewZoneArgoTieredCachingService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneArgoTieredCachingService(opts ...option.RequestOption) (r *ZoneArgoTieredCachingService) {
	r = &ZoneArgoTieredCachingService{}
	r.Options = opts
	return
}

// Tiered Cache works by dividing Cloudflare's data centers into a hierarchy of
// lower-tiers and upper-tiers. If content is not cached in lower-tier data centers
// (generally the ones closest to a visitor), the lower-tier must ask an upper-tier
// to see if it has the content. If the upper-tier does not have the content, only
// the upper-tier can ask the origin for content. This practice improves bandwidth
// efficiency by limiting the number of data centers that can ask the origin for
// content, which reduces origin load and makes websites more cost-effective to
// operate. Additionally, Tiered Cache concentrates connections to origin servers
// so they come from a small number of data centers rather than the full set of
// network locations. This results in fewer open connections using server
// resources.
func (r *ZoneArgoTieredCachingService) Get(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *ZoneArgoTieredCachingGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/argo/tiered_caching", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Tiered Cache works by dividing Cloudflare's data centers into a hierarchy of
// lower-tiers and upper-tiers. If content is not cached in lower-tier data centers
// (generally the ones closest to a visitor), the lower-tier must ask an upper-tier
// to see if it has the content. If the upper-tier does not have the content, only
// the upper-tier can ask the origin for content. This practice improves bandwidth
// efficiency by limiting the number of data centers that can ask the origin for
// content, which reduces origin load and makes websites more cost-effective to
// operate. Additionally, Tiered Cache concentrates connections to origin servers
// so they come from a small number of data centers rather than the full set of
// network locations. This results in fewer open connections using server
// resources.
func (r *ZoneArgoTieredCachingService) Update(ctx context.Context, zoneID string, body ZoneArgoTieredCachingUpdateParams, opts ...option.RequestOption) (res *ZoneArgoTieredCachingUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/argo/tiered_caching", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

type CacheRulesTieredCacheResponse struct {
	Result CacheRulesTieredCacheResponseResult `json:"result"`
	JSON   cacheRulesTieredCacheResponseJSON   `json:"-"`
}

// cacheRulesTieredCacheResponseJSON contains the JSON metadata for the struct
// [CacheRulesTieredCacheResponse]
type cacheRulesTieredCacheResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CacheRulesTieredCacheResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cacheRulesTieredCacheResponseJSON) RawJSON() string {
	return r.raw
}

type CacheRulesTieredCacheResponseResult struct {
	// Value of the Tiered Cache zone setting.
	Value CacheRulesTieredCacheResponseResultValue `json:"value,required"`
	// ID of the zone setting.
	ID   CacheRulesTieredCacheResponseResultID   `json:"id"`
	JSON cacheRulesTieredCacheResponseResultJSON `json:"-"`
	BaseCacheRule
}

// cacheRulesTieredCacheResponseResultJSON contains the JSON metadata for the
// struct [CacheRulesTieredCacheResponseResult]
type cacheRulesTieredCacheResponseResultJSON struct {
	Value       apijson.Field
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CacheRulesTieredCacheResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cacheRulesTieredCacheResponseResultJSON) RawJSON() string {
	return r.raw
}

// Value of the Tiered Cache zone setting.
type CacheRulesTieredCacheResponseResultValue string

const (
	CacheRulesTieredCacheResponseResultValueOn  CacheRulesTieredCacheResponseResultValue = "on"
	CacheRulesTieredCacheResponseResultValueOff CacheRulesTieredCacheResponseResultValue = "off"
)

func (r CacheRulesTieredCacheResponseResultValue) IsKnown() bool {
	switch r {
	case CacheRulesTieredCacheResponseResultValueOn, CacheRulesTieredCacheResponseResultValueOff:
		return true
	}
	return false
}

// ID of the zone setting.
type CacheRulesTieredCacheResponseResultID string

const (
	CacheRulesTieredCacheResponseResultIDTieredCaching CacheRulesTieredCacheResponseResultID = "tiered_caching"
)

func (r CacheRulesTieredCacheResponseResultID) IsKnown() bool {
	switch r {
	case CacheRulesTieredCacheResponseResultIDTieredCaching:
		return true
	}
	return false
}

type CacheRulesZoneCacheSettingsResponse struct {
	Result CacheRulesZoneCacheSettingsResponseResult `json:"result"`
	JSON   cacheRulesZoneCacheSettingsResponseJSON   `json:"-"`
	APIResponseCacheRules
}

// cacheRulesZoneCacheSettingsResponseJSON contains the JSON metadata for the
// struct [CacheRulesZoneCacheSettingsResponse]
type cacheRulesZoneCacheSettingsResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CacheRulesZoneCacheSettingsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cacheRulesZoneCacheSettingsResponseJSON) RawJSON() string {
	return r.raw
}

type CacheRulesZoneCacheSettingsResponseResult struct {
	// The identifier of the caching setting
	ID string `json:"id,required"`
	// Whether the setting is editable
	Editable bool `json:"editable,required"`
	// The value of the feature
	Value string `json:"value,required"`
	// The time when the setting was last modified
	ModifiedOn time.Time                                     `json:"modified_on" format:"date-time"`
	JSON       cacheRulesZoneCacheSettingsResponseResultJSON `json:"-"`
}

// cacheRulesZoneCacheSettingsResponseResultJSON contains the JSON metadata for the
// struct [CacheRulesZoneCacheSettingsResponseResult]
type cacheRulesZoneCacheSettingsResponseResultJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	Value       apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CacheRulesZoneCacheSettingsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cacheRulesZoneCacheSettingsResponseResultJSON) RawJSON() string {
	return r.raw
}

type ZoneArgoTieredCachingGetResponse struct {
	JSON zoneArgoTieredCachingGetResponseJSON `json:"-"`
	CacheRulesZoneCacheSettingsResponse
	CacheRulesTieredCacheResponse
}

// zoneArgoTieredCachingGetResponseJSON contains the JSON metadata for the struct
// [ZoneArgoTieredCachingGetResponse]
type zoneArgoTieredCachingGetResponseJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneArgoTieredCachingGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneArgoTieredCachingGetResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneArgoTieredCachingUpdateResponse struct {
	JSON zoneArgoTieredCachingUpdateResponseJSON `json:"-"`
	CacheRulesZoneCacheSettingsResponse
	CacheRulesTieredCacheResponse
}

// zoneArgoTieredCachingUpdateResponseJSON contains the JSON metadata for the
// struct [ZoneArgoTieredCachingUpdateResponse]
type zoneArgoTieredCachingUpdateResponseJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneArgoTieredCachingUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneArgoTieredCachingUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneArgoTieredCachingUpdateParams struct {
	// Enables Tiered Caching.
	Value param.Field[ZoneArgoTieredCachingUpdateParamsValue] `json:"value,required"`
}

func (r ZoneArgoTieredCachingUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enables Tiered Caching.
type ZoneArgoTieredCachingUpdateParamsValue string

const (
	ZoneArgoTieredCachingUpdateParamsValueOn  ZoneArgoTieredCachingUpdateParamsValue = "on"
	ZoneArgoTieredCachingUpdateParamsValueOff ZoneArgoTieredCachingUpdateParamsValue = "off"
)

func (r ZoneArgoTieredCachingUpdateParamsValue) IsKnown() bool {
	switch r {
	case ZoneArgoTieredCachingUpdateParamsValueOn, ZoneArgoTieredCachingUpdateParamsValueOff:
		return true
	}
	return false
}
