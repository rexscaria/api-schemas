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

type ZoneArgoTieredCachingGetResponse struct {
	Errors   []MessagesCacheRulesItem `json:"errors,required"`
	Messages []MessagesCacheRulesItem `json:"messages,required"`
	// Whether the API call was successful
	Success ZoneArgoTieredCachingGetResponseSuccess `json:"success,required"`
	Result  ZoneArgoTieredCachingGetResponseResult  `json:"result"`
	JSON    zoneArgoTieredCachingGetResponseJSON    `json:"-"`
}

// zoneArgoTieredCachingGetResponseJSON contains the JSON metadata for the struct
// [ZoneArgoTieredCachingGetResponse]
type zoneArgoTieredCachingGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneArgoTieredCachingGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneArgoTieredCachingGetResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ZoneArgoTieredCachingGetResponseSuccess bool

const (
	ZoneArgoTieredCachingGetResponseSuccessTrue ZoneArgoTieredCachingGetResponseSuccess = true
)

func (r ZoneArgoTieredCachingGetResponseSuccess) IsKnown() bool {
	switch r {
	case ZoneArgoTieredCachingGetResponseSuccessTrue:
		return true
	}
	return false
}

type ZoneArgoTieredCachingGetResponseResult struct {
	// ID of the zone setting.
	ID ZoneArgoTieredCachingGetResponseResultID `json:"id,required"`
	// Whether the setting is editable
	Editable bool `json:"editable,required"`
	// The value of the feature
	Value ZoneArgoTieredCachingGetResponseResultValue `json:"value,required"`
	// Last time this setting was modified.
	ModifiedOn time.Time                                  `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneArgoTieredCachingGetResponseResultJSON `json:"-"`
}

// zoneArgoTieredCachingGetResponseResultJSON contains the JSON metadata for the
// struct [ZoneArgoTieredCachingGetResponseResult]
type zoneArgoTieredCachingGetResponseResultJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	Value       apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneArgoTieredCachingGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneArgoTieredCachingGetResponseResultJSON) RawJSON() string {
	return r.raw
}

// ID of the zone setting.
type ZoneArgoTieredCachingGetResponseResultID string

const (
	ZoneArgoTieredCachingGetResponseResultIDTieredCaching ZoneArgoTieredCachingGetResponseResultID = "tiered_caching"
)

func (r ZoneArgoTieredCachingGetResponseResultID) IsKnown() bool {
	switch r {
	case ZoneArgoTieredCachingGetResponseResultIDTieredCaching:
		return true
	}
	return false
}

// The value of the feature
type ZoneArgoTieredCachingGetResponseResultValue string

const (
	ZoneArgoTieredCachingGetResponseResultValueOn  ZoneArgoTieredCachingGetResponseResultValue = "on"
	ZoneArgoTieredCachingGetResponseResultValueOff ZoneArgoTieredCachingGetResponseResultValue = "off"
)

func (r ZoneArgoTieredCachingGetResponseResultValue) IsKnown() bool {
	switch r {
	case ZoneArgoTieredCachingGetResponseResultValueOn, ZoneArgoTieredCachingGetResponseResultValueOff:
		return true
	}
	return false
}

type ZoneArgoTieredCachingUpdateResponse struct {
	Errors   []MessagesCacheRulesItem `json:"errors,required"`
	Messages []MessagesCacheRulesItem `json:"messages,required"`
	// Whether the API call was successful
	Success ZoneArgoTieredCachingUpdateResponseSuccess `json:"success,required"`
	Result  ZoneArgoTieredCachingUpdateResponseResult  `json:"result"`
	JSON    zoneArgoTieredCachingUpdateResponseJSON    `json:"-"`
}

// zoneArgoTieredCachingUpdateResponseJSON contains the JSON metadata for the
// struct [ZoneArgoTieredCachingUpdateResponse]
type zoneArgoTieredCachingUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneArgoTieredCachingUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneArgoTieredCachingUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ZoneArgoTieredCachingUpdateResponseSuccess bool

const (
	ZoneArgoTieredCachingUpdateResponseSuccessTrue ZoneArgoTieredCachingUpdateResponseSuccess = true
)

func (r ZoneArgoTieredCachingUpdateResponseSuccess) IsKnown() bool {
	switch r {
	case ZoneArgoTieredCachingUpdateResponseSuccessTrue:
		return true
	}
	return false
}

type ZoneArgoTieredCachingUpdateResponseResult struct {
	// ID of the zone setting.
	ID ZoneArgoTieredCachingUpdateResponseResultID `json:"id,required"`
	// Whether the setting is editable
	Editable bool `json:"editable,required"`
	// The value of the feature
	Value ZoneArgoTieredCachingUpdateResponseResultValue `json:"value,required"`
	// Last time this setting was modified.
	ModifiedOn time.Time                                     `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneArgoTieredCachingUpdateResponseResultJSON `json:"-"`
}

// zoneArgoTieredCachingUpdateResponseResultJSON contains the JSON metadata for the
// struct [ZoneArgoTieredCachingUpdateResponseResult]
type zoneArgoTieredCachingUpdateResponseResultJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	Value       apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneArgoTieredCachingUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneArgoTieredCachingUpdateResponseResultJSON) RawJSON() string {
	return r.raw
}

// ID of the zone setting.
type ZoneArgoTieredCachingUpdateResponseResultID string

const (
	ZoneArgoTieredCachingUpdateResponseResultIDTieredCaching ZoneArgoTieredCachingUpdateResponseResultID = "tiered_caching"
)

func (r ZoneArgoTieredCachingUpdateResponseResultID) IsKnown() bool {
	switch r {
	case ZoneArgoTieredCachingUpdateResponseResultIDTieredCaching:
		return true
	}
	return false
}

// The value of the feature
type ZoneArgoTieredCachingUpdateResponseResultValue string

const (
	ZoneArgoTieredCachingUpdateResponseResultValueOn  ZoneArgoTieredCachingUpdateResponseResultValue = "on"
	ZoneArgoTieredCachingUpdateResponseResultValueOff ZoneArgoTieredCachingUpdateResponseResultValue = "off"
)

func (r ZoneArgoTieredCachingUpdateResponseResultValue) IsKnown() bool {
	switch r {
	case ZoneArgoTieredCachingUpdateResponseResultValueOn, ZoneArgoTieredCachingUpdateResponseResultValueOff:
		return true
	}
	return false
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
