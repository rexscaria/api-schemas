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

type ZoneCacheCacheReserveGetResponse struct {
	Errors   []MessagesCacheRulesItem `json:"errors,required"`
	Messages []MessagesCacheRulesItem `json:"messages,required"`
	// Whether the API call was successful
	Success ZoneCacheCacheReserveGetResponseSuccess `json:"success,required"`
	Result  ZoneCacheCacheReserveGetResponseResult  `json:"result"`
	JSON    zoneCacheCacheReserveGetResponseJSON    `json:"-"`
}

// zoneCacheCacheReserveGetResponseJSON contains the JSON metadata for the struct
// [ZoneCacheCacheReserveGetResponse]
type zoneCacheCacheReserveGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCacheCacheReserveGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneCacheCacheReserveGetResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ZoneCacheCacheReserveGetResponseSuccess bool

const (
	ZoneCacheCacheReserveGetResponseSuccessTrue ZoneCacheCacheReserveGetResponseSuccess = true
)

func (r ZoneCacheCacheReserveGetResponseSuccess) IsKnown() bool {
	switch r {
	case ZoneCacheCacheReserveGetResponseSuccessTrue:
		return true
	}
	return false
}

type ZoneCacheCacheReserveGetResponseResult struct {
	// ID of the zone setting.
	ID ZoneCacheCacheReserveGetResponseResultID `json:"id,required"`
	// Whether the setting is editable
	Editable bool `json:"editable,required"`
	// The value of the feature
	Value CacheReserveValue `json:"value,required"`
	// Last time this setting was modified.
	ModifiedOn time.Time                                  `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneCacheCacheReserveGetResponseResultJSON `json:"-"`
}

// zoneCacheCacheReserveGetResponseResultJSON contains the JSON metadata for the
// struct [ZoneCacheCacheReserveGetResponseResult]
type zoneCacheCacheReserveGetResponseResultJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	Value       apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCacheCacheReserveGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneCacheCacheReserveGetResponseResultJSON) RawJSON() string {
	return r.raw
}

// ID of the zone setting.
type ZoneCacheCacheReserveGetResponseResultID string

const (
	ZoneCacheCacheReserveGetResponseResultIDCacheReserve ZoneCacheCacheReserveGetResponseResultID = "cache_reserve"
)

func (r ZoneCacheCacheReserveGetResponseResultID) IsKnown() bool {
	switch r {
	case ZoneCacheCacheReserveGetResponseResultIDCacheReserve:
		return true
	}
	return false
}

type ZoneCacheCacheReserveUpdateResponse struct {
	Errors   []MessagesCacheRulesItem `json:"errors,required"`
	Messages []MessagesCacheRulesItem `json:"messages,required"`
	// Whether the API call was successful
	Success ZoneCacheCacheReserveUpdateResponseSuccess `json:"success,required"`
	Result  ZoneCacheCacheReserveUpdateResponseResult  `json:"result"`
	JSON    zoneCacheCacheReserveUpdateResponseJSON    `json:"-"`
}

// zoneCacheCacheReserveUpdateResponseJSON contains the JSON metadata for the
// struct [ZoneCacheCacheReserveUpdateResponse]
type zoneCacheCacheReserveUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCacheCacheReserveUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneCacheCacheReserveUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ZoneCacheCacheReserveUpdateResponseSuccess bool

const (
	ZoneCacheCacheReserveUpdateResponseSuccessTrue ZoneCacheCacheReserveUpdateResponseSuccess = true
)

func (r ZoneCacheCacheReserveUpdateResponseSuccess) IsKnown() bool {
	switch r {
	case ZoneCacheCacheReserveUpdateResponseSuccessTrue:
		return true
	}
	return false
}

type ZoneCacheCacheReserveUpdateResponseResult struct {
	// ID of the zone setting.
	ID ZoneCacheCacheReserveUpdateResponseResultID `json:"id,required"`
	// Whether the setting is editable
	Editable bool `json:"editable,required"`
	// The value of the feature
	Value CacheReserveValue `json:"value,required"`
	// Last time this setting was modified.
	ModifiedOn time.Time                                     `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneCacheCacheReserveUpdateResponseResultJSON `json:"-"`
}

// zoneCacheCacheReserveUpdateResponseResultJSON contains the JSON metadata for the
// struct [ZoneCacheCacheReserveUpdateResponseResult]
type zoneCacheCacheReserveUpdateResponseResultJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	Value       apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCacheCacheReserveUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneCacheCacheReserveUpdateResponseResultJSON) RawJSON() string {
	return r.raw
}

// ID of the zone setting.
type ZoneCacheCacheReserveUpdateResponseResultID string

const (
	ZoneCacheCacheReserveUpdateResponseResultIDCacheReserve ZoneCacheCacheReserveUpdateResponseResultID = "cache_reserve"
)

func (r ZoneCacheCacheReserveUpdateResponseResultID) IsKnown() bool {
	switch r {
	case ZoneCacheCacheReserveUpdateResponseResultIDCacheReserve:
		return true
	}
	return false
}

type ZoneCacheCacheReserveUpdateParams struct {
	// Value of the Cache Reserve zone setting.
	Value param.Field[CacheReserveValue] `json:"value,required"`
}

func (r ZoneCacheCacheReserveUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
