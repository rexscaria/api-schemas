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

// ZoneSettingOriginH2MaxStreamService contains methods and other services that
// help with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneSettingOriginH2MaxStreamService] method instead.
type ZoneSettingOriginH2MaxStreamService struct {
	Options []option.RequestOption
}

// NewZoneSettingOriginH2MaxStreamService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZoneSettingOriginH2MaxStreamService(opts ...option.RequestOption) (r *ZoneSettingOriginH2MaxStreamService) {
	r = &ZoneSettingOriginH2MaxStreamService{}
	r.Options = opts
	return
}

// Origin H2 Max Streams configures the max number of concurrent requests that
// Cloudflare will send within the same connection when communicating with the
// origin server, if the origin supports it. Note that if your origin does not
// support H2 multiplexing, 5xx errors may be observed, particularly 520s. Also
// note that the default value is `100` for all plan types except Enterprise where
// it is `1`. `1` means that H2 multiplexing is disabled.
func (r *ZoneSettingOriginH2MaxStreamService) Get(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *ZoneSettingOriginH2MaxStreamGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/settings/origin_h2_max_streams", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Origin H2 Max Streams configures the max number of concurrent requests that
// Cloudflare will send within the same connection when communicating with the
// origin server, if the origin supports it. Note that if your origin does not
// support H2 multiplexing, 5xx errors may be observed, particularly 520s. Also
// note that the default value is `100` for all plan types except Enterprise where
// it is `1`. `1` means that H2 multiplexing is disabled.
func (r *ZoneSettingOriginH2MaxStreamService) Update(ctx context.Context, zoneID string, body ZoneSettingOriginH2MaxStreamUpdateParams, opts ...option.RequestOption) (res *ZoneSettingOriginH2MaxStreamUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/settings/origin_h2_max_streams", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

type CacheRulesOriginH2MaxStreamsResponseValue struct {
	// Origin H2 Max Streams configures the max number of concurrent requests that
	// Cloudflare will send within the same connection when communicating with the
	// origin server, if the origin supports it. Note that if your origin does not
	// support H2 multiplexing, 5xx errors may be observed, particularly 520s. Also
	// note that the default value is `100` for all plan types except Enterprise where
	// it is `1`. `1` means that H2 multiplexing is disabled.
	Result CacheRulesOriginH2MaxStreamsResponseValueResult `json:"result"`
	JSON   cacheRulesOriginH2MaxStreamsResponseValueJSON   `json:"-"`
}

// cacheRulesOriginH2MaxStreamsResponseValueJSON contains the JSON metadata for the
// struct [CacheRulesOriginH2MaxStreamsResponseValue]
type cacheRulesOriginH2MaxStreamsResponseValueJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CacheRulesOriginH2MaxStreamsResponseValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cacheRulesOriginH2MaxStreamsResponseValueJSON) RawJSON() string {
	return r.raw
}

// Origin H2 Max Streams configures the max number of concurrent requests that
// Cloudflare will send within the same connection when communicating with the
// origin server, if the origin supports it. Note that if your origin does not
// support H2 multiplexing, 5xx errors may be observed, particularly 520s. Also
// note that the default value is `100` for all plan types except Enterprise where
// it is `1`. `1` means that H2 multiplexing is disabled.
type CacheRulesOriginH2MaxStreamsResponseValueResult struct {
	// Value of the zone setting.
	ID CacheRulesOriginH2MaxStreamsResponseValueResultID `json:"id"`
	// Value of the Origin H2 Max Streams Setting.
	Value int64                                               `json:"value"`
	JSON  cacheRulesOriginH2MaxStreamsResponseValueResultJSON `json:"-"`
	BaseCacheRule
}

// cacheRulesOriginH2MaxStreamsResponseValueResultJSON contains the JSON metadata
// for the struct [CacheRulesOriginH2MaxStreamsResponseValueResult]
type cacheRulesOriginH2MaxStreamsResponseValueResultJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CacheRulesOriginH2MaxStreamsResponseValueResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cacheRulesOriginH2MaxStreamsResponseValueResultJSON) RawJSON() string {
	return r.raw
}

// Value of the zone setting.
type CacheRulesOriginH2MaxStreamsResponseValueResultID string

const (
	CacheRulesOriginH2MaxStreamsResponseValueResultIDOriginH2MaxStreams CacheRulesOriginH2MaxStreamsResponseValueResultID = "origin_h2_max_streams"
)

func (r CacheRulesOriginH2MaxStreamsResponseValueResultID) IsKnown() bool {
	switch r {
	case CacheRulesOriginH2MaxStreamsResponseValueResultIDOriginH2MaxStreams:
		return true
	}
	return false
}

type ZoneSettingOriginH2MaxStreamGetResponse struct {
	JSON zoneSettingOriginH2MaxStreamGetResponseJSON `json:"-"`
	CacheRulesZoneCacheSettingsResponse
	CacheRulesOriginH2MaxStreamsResponseValue
}

// zoneSettingOriginH2MaxStreamGetResponseJSON contains the JSON metadata for the
// struct [ZoneSettingOriginH2MaxStreamGetResponse]
type zoneSettingOriginH2MaxStreamGetResponseJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingOriginH2MaxStreamGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingOriginH2MaxStreamGetResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneSettingOriginH2MaxStreamUpdateResponse struct {
	JSON zoneSettingOriginH2MaxStreamUpdateResponseJSON `json:"-"`
	CacheRulesZoneCacheSettingsResponse
	CacheRulesOriginH2MaxStreamsResponseValue
}

// zoneSettingOriginH2MaxStreamUpdateResponseJSON contains the JSON metadata for
// the struct [ZoneSettingOriginH2MaxStreamUpdateResponse]
type zoneSettingOriginH2MaxStreamUpdateResponseJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingOriginH2MaxStreamUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingOriginH2MaxStreamUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneSettingOriginH2MaxStreamUpdateParams struct {
	// Value of the Origin H2 Max Streams Setting.
	Value param.Field[int64] `json:"value,required"`
}

func (r ZoneSettingOriginH2MaxStreamUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
