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

// ZoneCacheCacheReserveService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneCacheCacheReserveService] method instead.
type ZoneCacheCacheReserveService struct {
	Options []option.RequestOption
}

// NewZoneCacheCacheReserveService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneCacheCacheReserveService(opts ...option.RequestOption) (r *ZoneCacheCacheReserveService) {
	r = &ZoneCacheCacheReserveService{}
	r.Options = opts
	return
}

// Increase cache lifetimes by automatically storing all cacheable files into
// Cloudflare's persistent object storage buckets. Requires Cache Reserve
// subscription. Note: using Tiered Cache with Cache Reserve is highly recommended
// to reduce Reserve operations costs. See the
// [developer docs](https://developers.cloudflare.com/cache/about/cache-reserve)
// for more information.
func (r *ZoneCacheCacheReserveService) Get(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *ZoneCacheCacheReserveGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/cache/cache_reserve", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Increase cache lifetimes by automatically storing all cacheable files into
// Cloudflare's persistent object storage buckets. Requires Cache Reserve
// subscription. Note: using Tiered Cache with Cache Reserve is highly recommended
// to reduce Reserve operations costs. See the
// [developer docs](https://developers.cloudflare.com/cache/about/cache-reserve)
// for more information.
func (r *ZoneCacheCacheReserveService) Update(ctx context.Context, zoneID string, body ZoneCacheCacheReserveUpdateParams, opts ...option.RequestOption) (res *ZoneCacheCacheReserveUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/cache/cache_reserve", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Value of the Cache Reserve zone setting.
type CacheReserveValue string

const (
	CacheReserveValueOn  CacheReserveValue = "on"
	CacheReserveValueOff CacheReserveValue = "off"
)

func (r CacheReserveValue) IsKnown() bool {
	switch r {
	case CacheReserveValueOn, CacheReserveValueOff:
		return true
	}
	return false
}

type ResponseValueReserve struct {
	Result ResponseValueReserveResult `json:"result"`
	JSON   responseValueReserveJSON   `json:"-"`
}

// responseValueReserveJSON contains the JSON metadata for the struct
// [ResponseValueReserve]
type responseValueReserveJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResponseValueReserve) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r responseValueReserveJSON) RawJSON() string {
	return r.raw
}

type ResponseValueReserveResult struct {
	// Value of the Cache Reserve zone setting.
	Value CacheReserveValue `json:"value,required"`
	// ID of the zone setting.
	ID   ResponseValueReserveResultID   `json:"id"`
	JSON responseValueReserveResultJSON `json:"-"`
	BaseCacheRule
}

// responseValueReserveResultJSON contains the JSON metadata for the struct
// [ResponseValueReserveResult]
type responseValueReserveResultJSON struct {
	Value       apijson.Field
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResponseValueReserveResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r responseValueReserveResultJSON) RawJSON() string {
	return r.raw
}

// ID of the zone setting.
type ResponseValueReserveResultID string

const (
	ResponseValueReserveResultIDCacheReserve ResponseValueReserveResultID = "cache_reserve"
)

func (r ResponseValueReserveResultID) IsKnown() bool {
	switch r {
	case ResponseValueReserveResultIDCacheReserve:
		return true
	}
	return false
}

type ZoneCacheCacheReserveGetResponse struct {
	JSON zoneCacheCacheReserveGetResponseJSON `json:"-"`
	CacheRulesZoneCacheSettingsResponse
	ResponseValueReserve
}

// zoneCacheCacheReserveGetResponseJSON contains the JSON metadata for the struct
// [ZoneCacheCacheReserveGetResponse]
type zoneCacheCacheReserveGetResponseJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCacheCacheReserveGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneCacheCacheReserveGetResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneCacheCacheReserveUpdateResponse struct {
	JSON zoneCacheCacheReserveUpdateResponseJSON `json:"-"`
	CacheRulesZoneCacheSettingsResponse
	ResponseValueReserve
}

// zoneCacheCacheReserveUpdateResponseJSON contains the JSON metadata for the
// struct [ZoneCacheCacheReserveUpdateResponse]
type zoneCacheCacheReserveUpdateResponseJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCacheCacheReserveUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneCacheCacheReserveUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneCacheCacheReserveUpdateParams struct {
	// Value of the Cache Reserve zone setting.
	Value param.Field[CacheReserveValue] `json:"value,required"`
}

func (r ZoneCacheCacheReserveUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
