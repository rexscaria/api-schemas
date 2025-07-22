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

type CacheRulesAegisResponseValue struct {
	// Aegis provides dedicated egress IPs (from Cloudflare to your origin) for your
	// layer 7 WAF and CDN services. The egress IPs are reserved exclusively for your
	// account so that you can increase your origin security by only allowing traffic
	// from a small list of IP addresses.
	Result CacheRulesAegisResponseValueResult `json:"result"`
	JSON   cacheRulesAegisResponseValueJSON   `json:"-"`
}

// cacheRulesAegisResponseValueJSON contains the JSON metadata for the struct
// [CacheRulesAegisResponseValue]
type cacheRulesAegisResponseValueJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CacheRulesAegisResponseValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cacheRulesAegisResponseValueJSON) RawJSON() string {
	return r.raw
}

// Aegis provides dedicated egress IPs (from Cloudflare to your origin) for your
// layer 7 WAF and CDN services. The egress IPs are reserved exclusively for your
// account so that you can increase your origin security by only allowing traffic
// from a small list of IP addresses.
type CacheRulesAegisResponseValueResult struct {
	// ID of the zone setting.
	ID CacheRulesAegisResponseValueResultID `json:"id"`
	// Value of the zone setting.
	Value CacheRulesAegisValue                   `json:"value"`
	JSON  cacheRulesAegisResponseValueResultJSON `json:"-"`
	BaseCacheRule
}

// cacheRulesAegisResponseValueResultJSON contains the JSON metadata for the struct
// [CacheRulesAegisResponseValueResult]
type cacheRulesAegisResponseValueResultJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CacheRulesAegisResponseValueResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cacheRulesAegisResponseValueResultJSON) RawJSON() string {
	return r.raw
}

// ID of the zone setting.
type CacheRulesAegisResponseValueResultID string

const (
	CacheRulesAegisResponseValueResultIDAegis CacheRulesAegisResponseValueResultID = "aegis"
)

func (r CacheRulesAegisResponseValueResultID) IsKnown() bool {
	switch r {
	case CacheRulesAegisResponseValueResultIDAegis:
		return true
	}
	return false
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

type CacheRulesZoneComplexCacheSettingsResponseSingle struct {
	Result CacheRulesZoneComplexCacheSettingsResponseSingleResult `json:"result"`
	JSON   cacheRulesZoneComplexCacheSettingsResponseSingleJSON   `json:"-"`
	APIResponseCacheRules
}

// cacheRulesZoneComplexCacheSettingsResponseSingleJSON contains the JSON metadata
// for the struct [CacheRulesZoneComplexCacheSettingsResponseSingle]
type cacheRulesZoneComplexCacheSettingsResponseSingleJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CacheRulesZoneComplexCacheSettingsResponseSingle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cacheRulesZoneComplexCacheSettingsResponseSingleJSON) RawJSON() string {
	return r.raw
}

type CacheRulesZoneComplexCacheSettingsResponseSingleResult struct {
	// The identifier of the caching setting
	ID string `json:"id,required"`
	// Whether the setting is editable
	Editable bool `json:"editable,required"`
	// The value of the feature
	Value interface{} `json:"value,required"`
	// The time when the setting was last modified
	ModifiedOn time.Time                                                  `json:"modified_on" format:"date-time"`
	JSON       cacheRulesZoneComplexCacheSettingsResponseSingleResultJSON `json:"-"`
}

// cacheRulesZoneComplexCacheSettingsResponseSingleResultJSON contains the JSON
// metadata for the struct [CacheRulesZoneComplexCacheSettingsResponseSingleResult]
type cacheRulesZoneComplexCacheSettingsResponseSingleResultJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	Value       apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CacheRulesZoneComplexCacheSettingsResponseSingleResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cacheRulesZoneComplexCacheSettingsResponseSingleResultJSON) RawJSON() string {
	return r.raw
}

type ZoneSettingAegiGetResponse struct {
	JSON zoneSettingAegiGetResponseJSON `json:"-"`
	CacheRulesZoneComplexCacheSettingsResponseSingle
	CacheRulesAegisResponseValue
}

// zoneSettingAegiGetResponseJSON contains the JSON metadata for the struct
// [ZoneSettingAegiGetResponse]
type zoneSettingAegiGetResponseJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingAegiGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingAegiGetResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneSettingAegiUpdateResponse struct {
	JSON zoneSettingAegiUpdateResponseJSON `json:"-"`
	CacheRulesZoneComplexCacheSettingsResponseSingle
	CacheRulesAegisResponseValue
}

// zoneSettingAegiUpdateResponseJSON contains the JSON metadata for the struct
// [ZoneSettingAegiUpdateResponse]
type zoneSettingAegiUpdateResponseJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingAegiUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingAegiUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneSettingAegiUpdateParams struct {
	// Value of the zone setting.
	Value param.Field[CacheRulesAegisValueParam] `json:"value,required"`
}

func (r ZoneSettingAegiUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
