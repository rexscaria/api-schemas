// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
)

// ZoneCacheRegionalTieredCacheService contains methods and other services that
// help with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneCacheRegionalTieredCacheService] method instead.
type ZoneCacheRegionalTieredCacheService struct {
	Options []option.RequestOption
}

// NewZoneCacheRegionalTieredCacheService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZoneCacheRegionalTieredCacheService(opts ...option.RequestOption) (r *ZoneCacheRegionalTieredCacheService) {
	r = &ZoneCacheRegionalTieredCacheService{}
	r.Options = opts
	return
}

// Instructs Cloudflare to check a regional hub data center on the way to your
// upper tier. This can help improve performance for smart and custom tiered cache
// topologies.
func (r *ZoneCacheRegionalTieredCacheService) Get(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *ZoneCacheRegionalTieredCacheGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/cache/regional_tiered_cache", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Instructs Cloudflare to check a regional hub data center on the way to your
// upper tier. This can help improve performance for smart and custom tiered cache
// topologies.
func (r *ZoneCacheRegionalTieredCacheService) Update(ctx context.Context, zoneID string, body ZoneCacheRegionalTieredCacheUpdateParams, opts ...option.RequestOption) (res *ZoneCacheRegionalTieredCacheUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/cache/regional_tiered_cache", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Value of the Regional Tiered Cache zone setting.
type RegionalTieredCacheValue string

const (
	RegionalTieredCacheValueOn  RegionalTieredCacheValue = "on"
	RegionalTieredCacheValueOff RegionalTieredCacheValue = "off"
)

func (r RegionalTieredCacheValue) IsKnown() bool {
	switch r {
	case RegionalTieredCacheValueOn, RegionalTieredCacheValueOff:
		return true
	}
	return false
}

type ResponseValueTiered struct {
	Result ResponseValueTieredResult `json:"result"`
	JSON   responseValueTieredJSON   `json:"-"`
}

// responseValueTieredJSON contains the JSON metadata for the struct
// [ResponseValueTiered]
type responseValueTieredJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResponseValueTiered) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r responseValueTieredJSON) RawJSON() string {
	return r.raw
}

type ResponseValueTieredResult struct {
	// Value of the Regional Tiered Cache zone setting.
	Value RegionalTieredCacheValue `json:"value,required"`
	// ID of the zone setting.
	ID   ResponseValueTieredResultID   `json:"id"`
	JSON responseValueTieredResultJSON `json:"-"`
	BaseCacheRule
}

// responseValueTieredResultJSON contains the JSON metadata for the struct
// [ResponseValueTieredResult]
type responseValueTieredResultJSON struct {
	Value       apijson.Field
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResponseValueTieredResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r responseValueTieredResultJSON) RawJSON() string {
	return r.raw
}

// ID of the zone setting.
type ResponseValueTieredResultID string

const (
	ResponseValueTieredResultIDTcRegional ResponseValueTieredResultID = "tc_regional"
)

func (r ResponseValueTieredResultID) IsKnown() bool {
	switch r {
	case ResponseValueTieredResultIDTcRegional:
		return true
	}
	return false
}

type ZoneCacheRegionalTieredCacheGetResponse struct {
	JSON zoneCacheRegionalTieredCacheGetResponseJSON `json:"-"`
	CacheRulesZoneCacheSettingsResponse
	ResponseValueTiered
}

// zoneCacheRegionalTieredCacheGetResponseJSON contains the JSON metadata for the
// struct [ZoneCacheRegionalTieredCacheGetResponse]
type zoneCacheRegionalTieredCacheGetResponseJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCacheRegionalTieredCacheGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneCacheRegionalTieredCacheGetResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneCacheRegionalTieredCacheUpdateResponse struct {
	JSON zoneCacheRegionalTieredCacheUpdateResponseJSON `json:"-"`
	CacheRulesZoneCacheSettingsResponse
	ResponseValueTiered
}

// zoneCacheRegionalTieredCacheUpdateResponseJSON contains the JSON metadata for
// the struct [ZoneCacheRegionalTieredCacheUpdateResponse]
type zoneCacheRegionalTieredCacheUpdateResponseJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCacheRegionalTieredCacheUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneCacheRegionalTieredCacheUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneCacheRegionalTieredCacheUpdateParams struct {
	// Value of the Regional Tiered Cache zone setting.
	Value param.Field[RegionalTieredCacheValue] `json:"value,required"`
}

func (r ZoneCacheRegionalTieredCacheUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
