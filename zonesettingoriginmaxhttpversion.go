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

// ZoneSettingOriginMaxHTTPVersionService contains methods and other services that
// help with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneSettingOriginMaxHTTPVersionService] method instead.
type ZoneSettingOriginMaxHTTPVersionService struct {
	Options []option.RequestOption
}

// NewZoneSettingOriginMaxHTTPVersionService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZoneSettingOriginMaxHTTPVersionService(opts ...option.RequestOption) (r *ZoneSettingOriginMaxHTTPVersionService) {
	r = &ZoneSettingOriginMaxHTTPVersionService{}
	r.Options = opts
	return
}

// Origin Max HTTP Setting Version sets the highest HTTP version Cloudflare will
// attempt to use with your origin. This setting allows Cloudflare to make HTTP/2
// requests to your origin. (Refer to
// [Enable HTTP/2 to Origin](https://developers.cloudflare.com/cache/how-to/enable-http2-to-origin/),
// for more information.). The default value is "2" for all plan types except
// Enterprise where it is "1"
func (r *ZoneSettingOriginMaxHTTPVersionService) Get(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *ZoneSettingOriginMaxHTTPVersionGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/settings/origin_max_http_version", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Origin Max HTTP Setting Version sets the highest HTTP version Cloudflare will
// attempt to use with your origin. This setting allows Cloudflare to make HTTP/2
// requests to your origin. (Refer to
// [Enable HTTP/2 to Origin](https://developers.cloudflare.com/cache/how-to/enable-http2-to-origin/),
// for more information.). The default value is "2" for all plan types except
// Enterprise where it is "1"
func (r *ZoneSettingOriginMaxHTTPVersionService) Update(ctx context.Context, zoneID string, body ZoneSettingOriginMaxHTTPVersionUpdateParams, opts ...option.RequestOption) (res *ZoneSettingOriginMaxHTTPVersionUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/settings/origin_max_http_version", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

type CacheRulesOriginMaxHTTPVersionResponseValue struct {
	// Origin Max HTTP Setting Version sets the highest HTTP version Cloudflare will
	// attempt to use with your origin. This setting allows Cloudflare to make HTTP/2
	// requests to your origin. (Refer to
	// [Enable HTTP/2 to Origin](https://developers.cloudflare.com/cache/how-to/enable-http2-to-origin/),
	// for more information.). The default value is "2" for all plan types except
	// Enterprise where it is "1"
	Result CacheRulesOriginMaxHTTPVersionResponseValueResult `json:"result"`
	JSON   cacheRulesOriginMaxHTTPVersionResponseValueJSON   `json:"-"`
}

// cacheRulesOriginMaxHTTPVersionResponseValueJSON contains the JSON metadata for
// the struct [CacheRulesOriginMaxHTTPVersionResponseValue]
type cacheRulesOriginMaxHTTPVersionResponseValueJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CacheRulesOriginMaxHTTPVersionResponseValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cacheRulesOriginMaxHTTPVersionResponseValueJSON) RawJSON() string {
	return r.raw
}

// Origin Max HTTP Setting Version sets the highest HTTP version Cloudflare will
// attempt to use with your origin. This setting allows Cloudflare to make HTTP/2
// requests to your origin. (Refer to
// [Enable HTTP/2 to Origin](https://developers.cloudflare.com/cache/how-to/enable-http2-to-origin/),
// for more information.). The default value is "2" for all plan types except
// Enterprise where it is "1"
type CacheRulesOriginMaxHTTPVersionResponseValueResult struct {
	// Value of the zone setting.
	ID CacheRulesOriginMaxHTTPVersionResponseValueResultID `json:"id"`
	// Value of the Origin Max HTTP Version Setting.
	Value CacheRulesOriginMaxHTTPVersionValue                   `json:"value"`
	JSON  cacheRulesOriginMaxHTTPVersionResponseValueResultJSON `json:"-"`
	BaseCacheRule
}

// cacheRulesOriginMaxHTTPVersionResponseValueResultJSON contains the JSON metadata
// for the struct [CacheRulesOriginMaxHTTPVersionResponseValueResult]
type cacheRulesOriginMaxHTTPVersionResponseValueResultJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CacheRulesOriginMaxHTTPVersionResponseValueResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cacheRulesOriginMaxHTTPVersionResponseValueResultJSON) RawJSON() string {
	return r.raw
}

// Value of the zone setting.
type CacheRulesOriginMaxHTTPVersionResponseValueResultID string

const (
	CacheRulesOriginMaxHTTPVersionResponseValueResultIDOriginMaxHTTPVersion CacheRulesOriginMaxHTTPVersionResponseValueResultID = "origin_max_http_version"
)

func (r CacheRulesOriginMaxHTTPVersionResponseValueResultID) IsKnown() bool {
	switch r {
	case CacheRulesOriginMaxHTTPVersionResponseValueResultIDOriginMaxHTTPVersion:
		return true
	}
	return false
}

// Value of the Origin Max HTTP Version Setting.
type CacheRulesOriginMaxHTTPVersionValue string

const (
	CacheRulesOriginMaxHTTPVersionValue2 CacheRulesOriginMaxHTTPVersionValue = "2"
	CacheRulesOriginMaxHTTPVersionValue1 CacheRulesOriginMaxHTTPVersionValue = "1"
)

func (r CacheRulesOriginMaxHTTPVersionValue) IsKnown() bool {
	switch r {
	case CacheRulesOriginMaxHTTPVersionValue2, CacheRulesOriginMaxHTTPVersionValue1:
		return true
	}
	return false
}

type ZoneSettingOriginMaxHTTPVersionGetResponse struct {
	JSON zoneSettingOriginMaxHTTPVersionGetResponseJSON `json:"-"`
	CacheRulesZoneCacheSettingsResponse
	CacheRulesOriginMaxHTTPVersionResponseValue
}

// zoneSettingOriginMaxHTTPVersionGetResponseJSON contains the JSON metadata for
// the struct [ZoneSettingOriginMaxHTTPVersionGetResponse]
type zoneSettingOriginMaxHTTPVersionGetResponseJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingOriginMaxHTTPVersionGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingOriginMaxHTTPVersionGetResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneSettingOriginMaxHTTPVersionUpdateResponse struct {
	JSON zoneSettingOriginMaxHTTPVersionUpdateResponseJSON `json:"-"`
	CacheRulesZoneCacheSettingsResponse
	CacheRulesOriginMaxHTTPVersionResponseValue
}

// zoneSettingOriginMaxHTTPVersionUpdateResponseJSON contains the JSON metadata for
// the struct [ZoneSettingOriginMaxHTTPVersionUpdateResponse]
type zoneSettingOriginMaxHTTPVersionUpdateResponseJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingOriginMaxHTTPVersionUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingOriginMaxHTTPVersionUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneSettingOriginMaxHTTPVersionUpdateParams struct {
	// Value of the Origin Max HTTP Version Setting.
	Value param.Field[CacheRulesOriginMaxHTTPVersionValue] `json:"value,required"`
}

func (r ZoneSettingOriginMaxHTTPVersionUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
