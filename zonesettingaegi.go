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

// ZoneSettingAegiService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneSettingAegiService] method instead.
type ZoneSettingAegiService struct {
	Options []option.RequestOption
}

// NewZoneSettingAegiService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewZoneSettingAegiService(opts ...option.RequestOption) (r *ZoneSettingAegiService) {
	r = &ZoneSettingAegiService{}
	r.Options = opts
	return
}

// Aegis provides dedicated egress IPs (from Cloudflare to your origin) for your
// layer 7 WAF and CDN services. The egress IPs are reserved exclusively for your
// account so that you can increase your origin security by only allowing traffic
// from a small list of IP addresses.
func (r *ZoneSettingAegiService) Get(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *ZoneSettingAegiGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/settings/aegis", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Aegis provides dedicated egress IPs (from Cloudflare to your origin) for your
// layer 7 WAF and CDN services. The egress IPs are reserved exclusively for your
// account so that you can increase your origin security by only allowing traffic
// from a small list of IP addresses.
func (r *ZoneSettingAegiService) Update(ctx context.Context, zoneID string, body ZoneSettingAegiUpdateParams, opts ...option.RequestOption) (res *ZoneSettingAegiUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/settings/aegis", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Value of the zone setting.
type CacheRulesAegisValue struct {
	// Whether the feature is enabled or not.
	Enabled bool `json:"enabled"`
	// Egress pool id which refers to a grouping of dedicated egress IPs through which
	// Cloudflare will connect to origin.
	PoolID string                   `json:"pool_id"`
	JSON   cacheRulesAegisValueJSON `json:"-"`
}

// cacheRulesAegisValueJSON contains the JSON metadata for the struct
// [CacheRulesAegisValue]
type cacheRulesAegisValueJSON struct {
	Enabled     apijson.Field
	PoolID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CacheRulesAegisValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cacheRulesAegisValueJSON) RawJSON() string {
	return r.raw
}

// Value of the zone setting.
type CacheRulesAegisValueParam struct {
	// Whether the feature is enabled or not.
	Enabled param.Field[bool] `json:"enabled"`
	// Egress pool id which refers to a grouping of dedicated egress IPs through which
	// Cloudflare will connect to origin.
	PoolID param.Field[string] `json:"pool_id"`
}

func (r CacheRulesAegisValueParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneSettingAegiGetResponse struct {
	Errors   []MessagesCacheRulesItem `json:"errors,required"`
	Messages []MessagesCacheRulesItem `json:"messages,required"`
	// Whether the API call was successful
	Success ZoneSettingAegiGetResponseSuccess `json:"success,required"`
	Result  ZoneSettingAegiGetResponseResult  `json:"result"`
	JSON    zoneSettingAegiGetResponseJSON    `json:"-"`
}

// zoneSettingAegiGetResponseJSON contains the JSON metadata for the struct
// [ZoneSettingAegiGetResponse]
type zoneSettingAegiGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingAegiGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingAegiGetResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ZoneSettingAegiGetResponseSuccess bool

const (
	ZoneSettingAegiGetResponseSuccessTrue ZoneSettingAegiGetResponseSuccess = true
)

func (r ZoneSettingAegiGetResponseSuccess) IsKnown() bool {
	switch r {
	case ZoneSettingAegiGetResponseSuccessTrue:
		return true
	}
	return false
}

type ZoneSettingAegiGetResponseResult struct {
	// ID of the zone setting.
	ID ZoneSettingAegiGetResponseResultID `json:"id,required"`
	// Whether the setting is editable
	Editable bool `json:"editable,required"`
	// The value of the feature
	Value CacheRulesAegisValue `json:"value,required"`
	// Last time this setting was modified.
	ModifiedOn time.Time                            `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingAegiGetResponseResultJSON `json:"-"`
}

// zoneSettingAegiGetResponseResultJSON contains the JSON metadata for the struct
// [ZoneSettingAegiGetResponseResult]
type zoneSettingAegiGetResponseResultJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	Value       apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingAegiGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingAegiGetResponseResultJSON) RawJSON() string {
	return r.raw
}

// ID of the zone setting.
type ZoneSettingAegiGetResponseResultID string

const (
	ZoneSettingAegiGetResponseResultIDAegis ZoneSettingAegiGetResponseResultID = "aegis"
)

func (r ZoneSettingAegiGetResponseResultID) IsKnown() bool {
	switch r {
	case ZoneSettingAegiGetResponseResultIDAegis:
		return true
	}
	return false
}

type ZoneSettingAegiUpdateResponse struct {
	Errors   []MessagesCacheRulesItem `json:"errors,required"`
	Messages []MessagesCacheRulesItem `json:"messages,required"`
	// Whether the API call was successful
	Success ZoneSettingAegiUpdateResponseSuccess `json:"success,required"`
	Result  ZoneSettingAegiUpdateResponseResult  `json:"result"`
	JSON    zoneSettingAegiUpdateResponseJSON    `json:"-"`
}

// zoneSettingAegiUpdateResponseJSON contains the JSON metadata for the struct
// [ZoneSettingAegiUpdateResponse]
type zoneSettingAegiUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingAegiUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingAegiUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ZoneSettingAegiUpdateResponseSuccess bool

const (
	ZoneSettingAegiUpdateResponseSuccessTrue ZoneSettingAegiUpdateResponseSuccess = true
)

func (r ZoneSettingAegiUpdateResponseSuccess) IsKnown() bool {
	switch r {
	case ZoneSettingAegiUpdateResponseSuccessTrue:
		return true
	}
	return false
}

type ZoneSettingAegiUpdateResponseResult struct {
	// ID of the zone setting.
	ID ZoneSettingAegiUpdateResponseResultID `json:"id,required"`
	// Whether the setting is editable
	Editable bool `json:"editable,required"`
	// The value of the feature
	Value CacheRulesAegisValue `json:"value,required"`
	// Last time this setting was modified.
	ModifiedOn time.Time                               `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingAegiUpdateResponseResultJSON `json:"-"`
}

// zoneSettingAegiUpdateResponseResultJSON contains the JSON metadata for the
// struct [ZoneSettingAegiUpdateResponseResult]
type zoneSettingAegiUpdateResponseResultJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	Value       apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingAegiUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingAegiUpdateResponseResultJSON) RawJSON() string {
	return r.raw
}

// ID of the zone setting.
type ZoneSettingAegiUpdateResponseResultID string

const (
	ZoneSettingAegiUpdateResponseResultIDAegis ZoneSettingAegiUpdateResponseResultID = "aegis"
)

func (r ZoneSettingAegiUpdateResponseResultID) IsKnown() bool {
	switch r {
	case ZoneSettingAegiUpdateResponseResultIDAegis:
		return true
	}
	return false
}

type ZoneSettingAegiUpdateParams struct {
	// Value of the zone setting.
	Value param.Field[CacheRulesAegisValueParam] `json:"value,required"`
}

func (r ZoneSettingAegiUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
