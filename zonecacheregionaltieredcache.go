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

type ZoneCacheRegionalTieredCacheGetResponse struct {
	Errors   []MessagesCacheRulesItem `json:"errors,required"`
	Messages []MessagesCacheRulesItem `json:"messages,required"`
	// Whether the API call was successful
	Success ZoneCacheRegionalTieredCacheGetResponseSuccess `json:"success,required"`
	Result  ZoneCacheRegionalTieredCacheGetResponseResult  `json:"result"`
	JSON    zoneCacheRegionalTieredCacheGetResponseJSON    `json:"-"`
}

// zoneCacheRegionalTieredCacheGetResponseJSON contains the JSON metadata for the
// struct [ZoneCacheRegionalTieredCacheGetResponse]
type zoneCacheRegionalTieredCacheGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCacheRegionalTieredCacheGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneCacheRegionalTieredCacheGetResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ZoneCacheRegionalTieredCacheGetResponseSuccess bool

const (
	ZoneCacheRegionalTieredCacheGetResponseSuccessTrue ZoneCacheRegionalTieredCacheGetResponseSuccess = true
)

func (r ZoneCacheRegionalTieredCacheGetResponseSuccess) IsKnown() bool {
	switch r {
	case ZoneCacheRegionalTieredCacheGetResponseSuccessTrue:
		return true
	}
	return false
}

type ZoneCacheRegionalTieredCacheGetResponseResult struct {
	// ID of the zone setting.
	ID ZoneCacheRegionalTieredCacheGetResponseResultID `json:"id,required"`
	// Whether the setting is editable
	Editable bool `json:"editable,required"`
	// The value of the feature
	Value RegionalTieredCacheValue `json:"value,required"`
	// Last time this setting was modified.
	ModifiedOn time.Time                                         `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneCacheRegionalTieredCacheGetResponseResultJSON `json:"-"`
}

// zoneCacheRegionalTieredCacheGetResponseResultJSON contains the JSON metadata for
// the struct [ZoneCacheRegionalTieredCacheGetResponseResult]
type zoneCacheRegionalTieredCacheGetResponseResultJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	Value       apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCacheRegionalTieredCacheGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneCacheRegionalTieredCacheGetResponseResultJSON) RawJSON() string {
	return r.raw
}

// ID of the zone setting.
type ZoneCacheRegionalTieredCacheGetResponseResultID string

const (
	ZoneCacheRegionalTieredCacheGetResponseResultIDTcRegional ZoneCacheRegionalTieredCacheGetResponseResultID = "tc_regional"
)

func (r ZoneCacheRegionalTieredCacheGetResponseResultID) IsKnown() bool {
	switch r {
	case ZoneCacheRegionalTieredCacheGetResponseResultIDTcRegional:
		return true
	}
	return false
}

type ZoneCacheRegionalTieredCacheUpdateResponse struct {
	Errors   []MessagesCacheRulesItem `json:"errors,required"`
	Messages []MessagesCacheRulesItem `json:"messages,required"`
	// Whether the API call was successful
	Success ZoneCacheRegionalTieredCacheUpdateResponseSuccess `json:"success,required"`
	Result  ZoneCacheRegionalTieredCacheUpdateResponseResult  `json:"result"`
	JSON    zoneCacheRegionalTieredCacheUpdateResponseJSON    `json:"-"`
}

// zoneCacheRegionalTieredCacheUpdateResponseJSON contains the JSON metadata for
// the struct [ZoneCacheRegionalTieredCacheUpdateResponse]
type zoneCacheRegionalTieredCacheUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCacheRegionalTieredCacheUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneCacheRegionalTieredCacheUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ZoneCacheRegionalTieredCacheUpdateResponseSuccess bool

const (
	ZoneCacheRegionalTieredCacheUpdateResponseSuccessTrue ZoneCacheRegionalTieredCacheUpdateResponseSuccess = true
)

func (r ZoneCacheRegionalTieredCacheUpdateResponseSuccess) IsKnown() bool {
	switch r {
	case ZoneCacheRegionalTieredCacheUpdateResponseSuccessTrue:
		return true
	}
	return false
}

type ZoneCacheRegionalTieredCacheUpdateResponseResult struct {
	// ID of the zone setting.
	ID ZoneCacheRegionalTieredCacheUpdateResponseResultID `json:"id,required"`
	// Whether the setting is editable
	Editable bool `json:"editable,required"`
	// The value of the feature
	Value RegionalTieredCacheValue `json:"value,required"`
	// Last time this setting was modified.
	ModifiedOn time.Time                                            `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneCacheRegionalTieredCacheUpdateResponseResultJSON `json:"-"`
}

// zoneCacheRegionalTieredCacheUpdateResponseResultJSON contains the JSON metadata
// for the struct [ZoneCacheRegionalTieredCacheUpdateResponseResult]
type zoneCacheRegionalTieredCacheUpdateResponseResultJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	Value       apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCacheRegionalTieredCacheUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneCacheRegionalTieredCacheUpdateResponseResultJSON) RawJSON() string {
	return r.raw
}

// ID of the zone setting.
type ZoneCacheRegionalTieredCacheUpdateResponseResultID string

const (
	ZoneCacheRegionalTieredCacheUpdateResponseResultIDTcRegional ZoneCacheRegionalTieredCacheUpdateResponseResultID = "tc_regional"
)

func (r ZoneCacheRegionalTieredCacheUpdateResponseResultID) IsKnown() bool {
	switch r {
	case ZoneCacheRegionalTieredCacheUpdateResponseResultIDTcRegional:
		return true
	}
	return false
}

type ZoneCacheRegionalTieredCacheUpdateParams struct {
	// Value of the Regional Tiered Cache zone setting.
	Value param.Field[RegionalTieredCacheValue] `json:"value,required"`
}

func (r ZoneCacheRegionalTieredCacheUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
