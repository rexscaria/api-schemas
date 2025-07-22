// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
	"github.com/tidwall/gjson"
)

// ZoneSettingService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneSettingService] method instead.
type ZoneSettingService struct {
	Options              []option.RequestOption
	Aegis                *ZoneSettingAegiService
	Fonts                *ZoneSettingFontService
	OriginH2MaxStreams   *ZoneSettingOriginH2MaxStreamService
	OriginMaxHTTPVersion *ZoneSettingOriginMaxHTTPVersionService
	Rum                  *ZoneSettingRumService
	SpeedBrain           *ZoneSettingSpeedBrainService
	Zaraz                *ZoneSettingZarazService
	SslAutomaticMode     *ZoneSettingSslAutomaticModeService
}

// NewZoneSettingService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewZoneSettingService(opts ...option.RequestOption) (r *ZoneSettingService) {
	r = &ZoneSettingService{}
	r.Options = opts
	r.Aegis = NewZoneSettingAegiService(opts...)
	r.Fonts = NewZoneSettingFontService(opts...)
	r.OriginH2MaxStreams = NewZoneSettingOriginH2MaxStreamService(opts...)
	r.OriginMaxHTTPVersion = NewZoneSettingOriginMaxHTTPVersionService(opts...)
	r.Rum = NewZoneSettingRumService(opts...)
	r.SpeedBrain = NewZoneSettingSpeedBrainService(opts...)
	r.Zaraz = NewZoneSettingZarazService(opts...)
	r.SslAutomaticMode = NewZoneSettingSslAutomaticModeService(opts...)
	return
}

// Available settings for your user in relation to a zone.
func (r *ZoneSettingService) Get(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *ZonesZoneSettingsResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/settings", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Edit settings for a zone.
func (r *ZoneSettingService) Update(ctx context.Context, zoneID string, body ZoneSettingUpdateParams, opts ...option.RequestOption) (res *ZonesZoneSettingsResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/settings", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Fetch a single zone setting by name
func (r *ZoneSettingService) GetSetting(ctx context.Context, zoneID string, settingID string, opts ...option.RequestOption) (res *ZoneSettingGetSettingResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if settingID == "" {
		err = errors.New("missing required setting_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/settings/%s", zoneID, settingID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates a single zone setting by the identifier
func (r *ZoneSettingService) UpdateSetting(ctx context.Context, zoneID string, settingID string, body ZoneSettingUpdateSettingParams, opts ...option.RequestOption) (res *ZoneSettingUpdateSettingResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if settingID == "" {
		err = errors.New("missing required setting_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/settings/%s", zoneID, settingID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// 0-RTT session resumption enabled for this zone.
type Zones0rtt struct {
	// ID of the zone setting.
	ID Zones0rttID `json:"id"`
	// Value of the 0-RTT setting.
	Value Zones0rttValue `json:"value"`
	JSON  zones0rttJSON  `json:"-"`
	ZonesBase
}

// zones0rttJSON contains the JSON metadata for the struct [Zones0rtt]
type zones0rttJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Zones0rtt) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zones0rttJSON) RawJSON() string {
	return r.raw
}

func (r Zones0rtt) implementsZonesSetting() {}

func (r Zones0rtt) implementsZonesZoneSettingsResponseCollectionResult() {}

// ID of the zone setting.
type Zones0rttID string

const (
	Zones0rttID0rtt Zones0rttID = "0rtt"
)

func (r Zones0rttID) IsKnown() bool {
	switch r {
	case Zones0rttID0rtt:
		return true
	}
	return false
}

// Value of the 0-RTT setting.
type Zones0rttValue string

const (
	Zones0rttValueOn  Zones0rttValue = "on"
	Zones0rttValueOff Zones0rttValue = "off"
)

func (r Zones0rttValue) IsKnown() bool {
	switch r {
	case Zones0rttValueOn, Zones0rttValueOff:
		return true
	}
	return false
}

// 0-RTT session resumption enabled for this zone.
type Zones0rttParam struct {
	// ID of the zone setting.
	ID param.Field[Zones0rttID] `json:"id"`
	// Value of the 0-RTT setting.
	Value param.Field[Zones0rttValue] `json:"value"`
	ZonesBaseParam
}

func (r Zones0rttParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r Zones0rttParam) implementsZonesSettingUnionParam() {}

func (r Zones0rttParam) implementsZoneSettingUpdateParamsBodyUnion() {}

// Advanced protection from Distributed Denial of Service (DDoS) attacks on your
// website. This is an uneditable value that is 'on' in the case of Business and
// Enterprise zones.
type ZonesAdvancedDdos struct {
	// ID of the zone setting.
	ID ZonesAdvancedDdosID `json:"id"`
	// Value of the zone setting. Notes: Defaults to on for Business+ plans
	Value ZonesAdvancedDdosValue `json:"value"`
	JSON  zonesAdvancedDdosJSON  `json:"-"`
	ZonesBase
}

// zonesAdvancedDdosJSON contains the JSON metadata for the struct
// [ZonesAdvancedDdos]
type zonesAdvancedDdosJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonesAdvancedDdos) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zonesAdvancedDdosJSON) RawJSON() string {
	return r.raw
}

func (r ZonesAdvancedDdos) implementsZonesSetting() {}

func (r ZonesAdvancedDdos) implementsZonesZoneSettingsResponseCollectionResult() {}

// ID of the zone setting.
type ZonesAdvancedDdosID string

const (
	ZonesAdvancedDdosIDAdvancedDdos ZonesAdvancedDdosID = "advanced_ddos"
)

func (r ZonesAdvancedDdosID) IsKnown() bool {
	switch r {
	case ZonesAdvancedDdosIDAdvancedDdos:
		return true
	}
	return false
}

// Value of the zone setting. Notes: Defaults to on for Business+ plans
type ZonesAdvancedDdosValue string

const (
	ZonesAdvancedDdosValueOn  ZonesAdvancedDdosValue = "on"
	ZonesAdvancedDdosValueOff ZonesAdvancedDdosValue = "off"
)

func (r ZonesAdvancedDdosValue) IsKnown() bool {
	switch r {
	case ZonesAdvancedDdosValueOn, ZonesAdvancedDdosValueOff:
		return true
	}
	return false
}

// Advanced protection from Distributed Denial of Service (DDoS) attacks on your
// website. This is an uneditable value that is 'on' in the case of Business and
// Enterprise zones.
type ZonesAdvancedDdosParam struct {
	// ID of the zone setting.
	ID param.Field[ZonesAdvancedDdosID] `json:"id"`
	// Value of the zone setting. Notes: Defaults to on for Business+ plans
	Value param.Field[ZonesAdvancedDdosValue] `json:"value"`
	ZonesBaseParam
}

func (r ZonesAdvancedDdosParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZonesAdvancedDdosParam) implementsZonesSettingUnionParam() {}

func (r ZonesAdvancedDdosParam) implementsZoneSettingUpdateParamsBodyUnion() {}

// When enabled, Cloudflare serves limited copies of web pages available from the
// [Internet Archive's Wayback Machine](https://archive.org/web/) if your server is
// offline. Refer to
// [Always Online](https://developers.cloudflare.com/cache/about/always-online) for
// more information.
type ZonesAlwaysOnline struct {
	// ID of the zone setting.
	ID ZonesAlwaysOnlineID `json:"id"`
	// Value of the zone setting.
	Value ZonesAlwaysOnlineValue `json:"value"`
	JSON  zonesAlwaysOnlineJSON  `json:"-"`
	ZonesBase
}

// zonesAlwaysOnlineJSON contains the JSON metadata for the struct
// [ZonesAlwaysOnline]
type zonesAlwaysOnlineJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonesAlwaysOnline) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zonesAlwaysOnlineJSON) RawJSON() string {
	return r.raw
}

func (r ZonesAlwaysOnline) implementsZonesSetting() {}

func (r ZonesAlwaysOnline) implementsZonesZoneSettingsResponseCollectionResult() {}

// ID of the zone setting.
type ZonesAlwaysOnlineID string

const (
	ZonesAlwaysOnlineIDAlwaysOnline ZonesAlwaysOnlineID = "always_online"
)

func (r ZonesAlwaysOnlineID) IsKnown() bool {
	switch r {
	case ZonesAlwaysOnlineIDAlwaysOnline:
		return true
	}
	return false
}

// Value of the zone setting.
type ZonesAlwaysOnlineValue string

const (
	ZonesAlwaysOnlineValueOn  ZonesAlwaysOnlineValue = "on"
	ZonesAlwaysOnlineValueOff ZonesAlwaysOnlineValue = "off"
)

func (r ZonesAlwaysOnlineValue) IsKnown() bool {
	switch r {
	case ZonesAlwaysOnlineValueOn, ZonesAlwaysOnlineValueOff:
		return true
	}
	return false
}

// When enabled, Cloudflare serves limited copies of web pages available from the
// [Internet Archive's Wayback Machine](https://archive.org/web/) if your server is
// offline. Refer to
// [Always Online](https://developers.cloudflare.com/cache/about/always-online) for
// more information.
type ZonesAlwaysOnlineParam struct {
	// ID of the zone setting.
	ID param.Field[ZonesAlwaysOnlineID] `json:"id"`
	// Value of the zone setting.
	Value param.Field[ZonesAlwaysOnlineValue] `json:"value"`
	ZonesBaseParam
}

func (r ZonesAlwaysOnlineParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZonesAlwaysOnlineParam) implementsZonesSettingUnionParam() {}

func (r ZonesAlwaysOnlineParam) implementsZoneSettingUpdateParamsBodyUnion() {}

type ZonesBase struct {
	// Identifier of the zone setting.
	ID string `json:"id,required"`
	// Current value of the zone setting.
	Value interface{} `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZonesBaseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time     `json:"modified_on,nullable" format:"date-time"`
	JSON       zonesBaseJSON `json:"-"`
}

// zonesBaseJSON contains the JSON metadata for the struct [ZonesBase]
type zonesBaseJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonesBase) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zonesBaseJSON) RawJSON() string {
	return r.raw
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZonesBaseEditable bool

const (
	ZonesBaseEditableTrue  ZonesBaseEditable = true
	ZonesBaseEditableFalse ZonesBaseEditable = false
)

func (r ZonesBaseEditable) IsKnown() bool {
	switch r {
	case ZonesBaseEditableTrue, ZonesBaseEditableFalse:
		return true
	}
	return false
}

type ZonesBaseParam struct {
	// Identifier of the zone setting.
	ID param.Field[string] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[interface{}] `json:"value,required"`
}

func (r ZonesBaseParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// When the client requesting an asset supports the Brotli compression algorithm,
// Cloudflare will serve a Brotli compressed version of the asset.
type ZonesBrotli struct {
	// ID of the zone setting.
	ID ZonesBrotliID `json:"id"`
	// Value of the zone setting.
	Value ZonesBrotliValue `json:"value"`
	JSON  zonesBrotliJSON  `json:"-"`
	ZonesBase
}

// zonesBrotliJSON contains the JSON metadata for the struct [ZonesBrotli]
type zonesBrotliJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonesBrotli) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zonesBrotliJSON) RawJSON() string {
	return r.raw
}

func (r ZonesBrotli) implementsZonesSetting() {}

func (r ZonesBrotli) implementsZonesZoneSettingsResponseCollectionResult() {}

// ID of the zone setting.
type ZonesBrotliID string

const (
	ZonesBrotliIDBrotli ZonesBrotliID = "brotli"
)

func (r ZonesBrotliID) IsKnown() bool {
	switch r {
	case ZonesBrotliIDBrotli:
		return true
	}
	return false
}

// Value of the zone setting.
type ZonesBrotliValue string

const (
	ZonesBrotliValueOff ZonesBrotliValue = "off"
	ZonesBrotliValueOn  ZonesBrotliValue = "on"
)

func (r ZonesBrotliValue) IsKnown() bool {
	switch r {
	case ZonesBrotliValueOff, ZonesBrotliValueOn:
		return true
	}
	return false
}

// When the client requesting an asset supports the Brotli compression algorithm,
// Cloudflare will serve a Brotli compressed version of the asset.
type ZonesBrotliParam struct {
	// ID of the zone setting.
	ID param.Field[ZonesBrotliID] `json:"id"`
	// Value of the zone setting.
	Value param.Field[ZonesBrotliValue] `json:"value"`
	ZonesBaseParam
}

func (r ZonesBrotliParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZonesBrotliParam) implementsZonesSettingUnionParam() {}

func (r ZonesBrotliParam) implementsZoneSettingUpdateParamsBodyUnion() {}

// Aegis provides dedicated egress IPs (from Cloudflare to your origin) for your
// layer 7 WAF and CDN services. The egress IPs are reserved exclusively for your
// account so that you can increase your origin security by only allowing traffic
// from a small list of IP addresses.
type ZonesCacheRulesAegis struct {
	// ID of the zone setting.
	ID ZonesCacheRulesAegisID `json:"id"`
	// Value of the zone setting.
	Value ZonesCacheRulesAegisValue `json:"value"`
	JSON  zonesCacheRulesAegisJSON  `json:"-"`
	ZonesCacheRulesBase
}

// zonesCacheRulesAegisJSON contains the JSON metadata for the struct
// [ZonesCacheRulesAegis]
type zonesCacheRulesAegisJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonesCacheRulesAegis) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zonesCacheRulesAegisJSON) RawJSON() string {
	return r.raw
}

func (r ZonesCacheRulesAegis) implementsZonesSetting() {}

func (r ZonesCacheRulesAegis) implementsZonesZoneSettingsResponseCollectionResult() {}

// ID of the zone setting.
type ZonesCacheRulesAegisID string

const (
	ZonesCacheRulesAegisIDAegis ZonesCacheRulesAegisID = "aegis"
)

func (r ZonesCacheRulesAegisID) IsKnown() bool {
	switch r {
	case ZonesCacheRulesAegisIDAegis:
		return true
	}
	return false
}

// Value of the zone setting.
type ZonesCacheRulesAegisValue struct {
	// Whether the feature is enabled or not.
	Enabled bool `json:"enabled"`
	// Egress pool id which refers to a grouping of dedicated egress IPs through which
	// Cloudflare will connect to origin.
	PoolID string                        `json:"pool_id"`
	JSON   zonesCacheRulesAegisValueJSON `json:"-"`
}

// zonesCacheRulesAegisValueJSON contains the JSON metadata for the struct
// [ZonesCacheRulesAegisValue]
type zonesCacheRulesAegisValueJSON struct {
	Enabled     apijson.Field
	PoolID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonesCacheRulesAegisValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zonesCacheRulesAegisValueJSON) RawJSON() string {
	return r.raw
}

// Aegis provides dedicated egress IPs (from Cloudflare to your origin) for your
// layer 7 WAF and CDN services. The egress IPs are reserved exclusively for your
// account so that you can increase your origin security by only allowing traffic
// from a small list of IP addresses.
type ZonesCacheRulesAegisParam struct {
	// ID of the zone setting.
	ID param.Field[ZonesCacheRulesAegisID] `json:"id"`
	// Value of the zone setting.
	Value param.Field[ZonesCacheRulesAegisValueParam] `json:"value"`
	ZonesCacheRulesBaseParam
}

func (r ZonesCacheRulesAegisParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZonesCacheRulesAegisParam) implementsZonesSettingUnionParam() {}

func (r ZonesCacheRulesAegisParam) implementsZoneSettingUpdateParamsBodyUnion() {}

// Value of the zone setting.
type ZonesCacheRulesAegisValueParam struct {
	// Whether the feature is enabled or not.
	Enabled param.Field[bool] `json:"enabled"`
	// Egress pool id which refers to a grouping of dedicated egress IPs through which
	// Cloudflare will connect to origin.
	PoolID param.Field[string] `json:"pool_id"`
}

func (r ZonesCacheRulesAegisValueParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZonesCacheRulesBase struct {
	// Identifier of the zone setting.
	ID string `json:"id,required"`
	// Last time this setting was modified.
	ModifiedOn time.Time               `json:"modified_on,nullable" format:"date-time"`
	JSON       zonesCacheRulesBaseJSON `json:"-"`
}

// zonesCacheRulesBaseJSON contains the JSON metadata for the struct
// [ZonesCacheRulesBase]
type zonesCacheRulesBaseJSON struct {
	ID          apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonesCacheRulesBase) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zonesCacheRulesBaseJSON) RawJSON() string {
	return r.raw
}

type ZonesCacheRulesBaseParam struct {
	// Identifier of the zone setting.
	ID param.Field[string] `json:"id,required"`
}

func (r ZonesCacheRulesBaseParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Origin H2 Max Streams configures the max number of concurrent requests that
// Cloudflare will send within the same connection when communicating with the
// origin server, if the origin supports it. Note that if your origin does not
// support H2 multiplexing, 5xx errors may be observed, particularly 520s. Also
// note that the default value is `100` for all plan types except Enterprise where
// it is `1`. `1` means that H2 multiplexing is disabled.
type ZonesCacheRulesOriginH2MaxStreams struct {
	// Value of the zone setting.
	ID ZonesCacheRulesOriginH2MaxStreamsID `json:"id"`
	// Value of the Origin H2 Max Streams Setting.
	Value int64                                 `json:"value"`
	JSON  zonesCacheRulesOriginH2MaxStreamsJSON `json:"-"`
	ZonesCacheRulesBase
}

// zonesCacheRulesOriginH2MaxStreamsJSON contains the JSON metadata for the struct
// [ZonesCacheRulesOriginH2MaxStreams]
type zonesCacheRulesOriginH2MaxStreamsJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonesCacheRulesOriginH2MaxStreams) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zonesCacheRulesOriginH2MaxStreamsJSON) RawJSON() string {
	return r.raw
}

func (r ZonesCacheRulesOriginH2MaxStreams) implementsZonesSetting() {}

func (r ZonesCacheRulesOriginH2MaxStreams) implementsZonesZoneSettingsResponseCollectionResult() {}

// Value of the zone setting.
type ZonesCacheRulesOriginH2MaxStreamsID string

const (
	ZonesCacheRulesOriginH2MaxStreamsIDOriginH2MaxStreams ZonesCacheRulesOriginH2MaxStreamsID = "origin_h2_max_streams"
)

func (r ZonesCacheRulesOriginH2MaxStreamsID) IsKnown() bool {
	switch r {
	case ZonesCacheRulesOriginH2MaxStreamsIDOriginH2MaxStreams:
		return true
	}
	return false
}

// Origin H2 Max Streams configures the max number of concurrent requests that
// Cloudflare will send within the same connection when communicating with the
// origin server, if the origin supports it. Note that if your origin does not
// support H2 multiplexing, 5xx errors may be observed, particularly 520s. Also
// note that the default value is `100` for all plan types except Enterprise where
// it is `1`. `1` means that H2 multiplexing is disabled.
type ZonesCacheRulesOriginH2MaxStreamsParam struct {
	// Value of the zone setting.
	ID param.Field[ZonesCacheRulesOriginH2MaxStreamsID] `json:"id"`
	// Value of the Origin H2 Max Streams Setting.
	Value param.Field[int64] `json:"value"`
	ZonesCacheRulesBaseParam
}

func (r ZonesCacheRulesOriginH2MaxStreamsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZonesCacheRulesOriginH2MaxStreamsParam) implementsZonesSettingUnionParam() {}

func (r ZonesCacheRulesOriginH2MaxStreamsParam) implementsZoneSettingUpdateParamsBodyUnion() {}

// Origin Max HTTP Setting Version sets the highest HTTP version Cloudflare will
// attempt to use with your origin. This setting allows Cloudflare to make HTTP/2
// requests to your origin. (Refer to
// [Enable HTTP/2 to Origin](https://developers.cloudflare.com/cache/how-to/enable-http2-to-origin/),
// for more information.). The default value is "2" for all plan types except
// Enterprise where it is "1"
type ZonesCacheRulesOriginMaxHTTPVersion struct {
	// Value of the zone setting.
	ID ZonesCacheRulesOriginMaxHTTPVersionID `json:"id"`
	// Value of the Origin Max HTTP Version Setting.
	Value ZonesCacheRulesOriginMaxHTTPVersionValue `json:"value"`
	JSON  zonesCacheRulesOriginMaxHTTPVersionJSON  `json:"-"`
	ZonesCacheRulesBase
}

// zonesCacheRulesOriginMaxHTTPVersionJSON contains the JSON metadata for the
// struct [ZonesCacheRulesOriginMaxHTTPVersion]
type zonesCacheRulesOriginMaxHTTPVersionJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonesCacheRulesOriginMaxHTTPVersion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zonesCacheRulesOriginMaxHTTPVersionJSON) RawJSON() string {
	return r.raw
}

func (r ZonesCacheRulesOriginMaxHTTPVersion) implementsZonesSetting() {}

func (r ZonesCacheRulesOriginMaxHTTPVersion) implementsZonesZoneSettingsResponseCollectionResult() {}

// Value of the zone setting.
type ZonesCacheRulesOriginMaxHTTPVersionID string

const (
	ZonesCacheRulesOriginMaxHTTPVersionIDOriginMaxHTTPVersion ZonesCacheRulesOriginMaxHTTPVersionID = "origin_max_http_version"
)

func (r ZonesCacheRulesOriginMaxHTTPVersionID) IsKnown() bool {
	switch r {
	case ZonesCacheRulesOriginMaxHTTPVersionIDOriginMaxHTTPVersion:
		return true
	}
	return false
}

// Value of the Origin Max HTTP Version Setting.
type ZonesCacheRulesOriginMaxHTTPVersionValue string

const (
	ZonesCacheRulesOriginMaxHTTPVersionValue2 ZonesCacheRulesOriginMaxHTTPVersionValue = "2"
	ZonesCacheRulesOriginMaxHTTPVersionValue1 ZonesCacheRulesOriginMaxHTTPVersionValue = "1"
)

func (r ZonesCacheRulesOriginMaxHTTPVersionValue) IsKnown() bool {
	switch r {
	case ZonesCacheRulesOriginMaxHTTPVersionValue2, ZonesCacheRulesOriginMaxHTTPVersionValue1:
		return true
	}
	return false
}

// Origin Max HTTP Setting Version sets the highest HTTP version Cloudflare will
// attempt to use with your origin. This setting allows Cloudflare to make HTTP/2
// requests to your origin. (Refer to
// [Enable HTTP/2 to Origin](https://developers.cloudflare.com/cache/how-to/enable-http2-to-origin/),
// for more information.). The default value is "2" for all plan types except
// Enterprise where it is "1"
type ZonesCacheRulesOriginMaxHTTPVersionParam struct {
	// Value of the zone setting.
	ID param.Field[ZonesCacheRulesOriginMaxHTTPVersionID] `json:"id"`
	// Value of the Origin Max HTTP Version Setting.
	Value param.Field[ZonesCacheRulesOriginMaxHTTPVersionValue] `json:"value"`
	ZonesCacheRulesBaseParam
}

func (r ZonesCacheRulesOriginMaxHTTPVersionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZonesCacheRulesOriginMaxHTTPVersionParam) implementsZonesSettingUnionParam() {}

func (r ZonesCacheRulesOriginMaxHTTPVersionParam) implementsZoneSettingUpdateParamsBodyUnion() {}

// Specify how long a visitor is allowed access to your site after successfully
// completing a challenge (such as a CAPTCHA). After the TTL has expired the
// visitor will have to complete a new challenge. We recommend a 15 - 45 minute
// setting and will attempt to honor any setting above 45 minutes.
// (https://support.cloudflare.com/hc/en-us/articles/200170136).
type ZonesChallengeTtl struct {
	// ID of the zone setting.
	ID ZonesChallengeTtlID `json:"id"`
	// Value of the zone setting.
	Value ZonesChallengeTtlValue `json:"value"`
	JSON  zonesChallengeTtlJSON  `json:"-"`
	ZonesBase
}

// zonesChallengeTtlJSON contains the JSON metadata for the struct
// [ZonesChallengeTtl]
type zonesChallengeTtlJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonesChallengeTtl) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zonesChallengeTtlJSON) RawJSON() string {
	return r.raw
}

func (r ZonesChallengeTtl) implementsZonesSetting() {}

func (r ZonesChallengeTtl) implementsZonesZoneSettingsResponseCollectionResult() {}

// ID of the zone setting.
type ZonesChallengeTtlID string

const (
	ZonesChallengeTtlIDChallengeTtl ZonesChallengeTtlID = "challenge_ttl"
)

func (r ZonesChallengeTtlID) IsKnown() bool {
	switch r {
	case ZonesChallengeTtlIDChallengeTtl:
		return true
	}
	return false
}

// Value of the zone setting.
type ZonesChallengeTtlValue float64

const (
	ZonesChallengeTtlValue300      ZonesChallengeTtlValue = 300
	ZonesChallengeTtlValue900      ZonesChallengeTtlValue = 900
	ZonesChallengeTtlValue1800     ZonesChallengeTtlValue = 1800
	ZonesChallengeTtlValue2700     ZonesChallengeTtlValue = 2700
	ZonesChallengeTtlValue3600     ZonesChallengeTtlValue = 3600
	ZonesChallengeTtlValue7200     ZonesChallengeTtlValue = 7200
	ZonesChallengeTtlValue10800    ZonesChallengeTtlValue = 10800
	ZonesChallengeTtlValue14400    ZonesChallengeTtlValue = 14400
	ZonesChallengeTtlValue28800    ZonesChallengeTtlValue = 28800
	ZonesChallengeTtlValue57600    ZonesChallengeTtlValue = 57600
	ZonesChallengeTtlValue86400    ZonesChallengeTtlValue = 86400
	ZonesChallengeTtlValue604800   ZonesChallengeTtlValue = 604800
	ZonesChallengeTtlValue2592000  ZonesChallengeTtlValue = 2592000
	ZonesChallengeTtlValue31536000 ZonesChallengeTtlValue = 31536000
)

func (r ZonesChallengeTtlValue) IsKnown() bool {
	switch r {
	case ZonesChallengeTtlValue300, ZonesChallengeTtlValue900, ZonesChallengeTtlValue1800, ZonesChallengeTtlValue2700, ZonesChallengeTtlValue3600, ZonesChallengeTtlValue7200, ZonesChallengeTtlValue10800, ZonesChallengeTtlValue14400, ZonesChallengeTtlValue28800, ZonesChallengeTtlValue57600, ZonesChallengeTtlValue86400, ZonesChallengeTtlValue604800, ZonesChallengeTtlValue2592000, ZonesChallengeTtlValue31536000:
		return true
	}
	return false
}

// Specify how long a visitor is allowed access to your site after successfully
// completing a challenge (such as a CAPTCHA). After the TTL has expired the
// visitor will have to complete a new challenge. We recommend a 15 - 45 minute
// setting and will attempt to honor any setting above 45 minutes.
// (https://support.cloudflare.com/hc/en-us/articles/200170136).
type ZonesChallengeTtlParam struct {
	// ID of the zone setting.
	ID param.Field[ZonesChallengeTtlID] `json:"id"`
	// Value of the zone setting.
	Value param.Field[ZonesChallengeTtlValue] `json:"value"`
	ZonesBaseParam
}

func (r ZonesChallengeTtlParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZonesChallengeTtlParam) implementsZonesSettingUnionParam() {}

func (r ZonesChallengeTtlParam) implementsZoneSettingUpdateParamsBodyUnion() {}

// An allowlist of ciphers for TLS termination. These ciphers must be in the
// BoringSSL format.
type ZonesCiphers struct {
	// ID of the zone setting.
	ID ZonesCiphersID `json:"id"`
	// Value of the zone setting.
	Value []string         `json:"value"`
	JSON  zonesCiphersJSON `json:"-"`
	ZonesBase
}

// zonesCiphersJSON contains the JSON metadata for the struct [ZonesCiphers]
type zonesCiphersJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonesCiphers) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zonesCiphersJSON) RawJSON() string {
	return r.raw
}

func (r ZonesCiphers) implementsZonesSetting() {}

func (r ZonesCiphers) implementsZonesZoneSettingsResponseCollectionResult() {}

// ID of the zone setting.
type ZonesCiphersID string

const (
	ZonesCiphersIDCiphers ZonesCiphersID = "ciphers"
)

func (r ZonesCiphersID) IsKnown() bool {
	switch r {
	case ZonesCiphersIDCiphers:
		return true
	}
	return false
}

// An allowlist of ciphers for TLS termination. These ciphers must be in the
// BoringSSL format.
type ZonesCiphersParam struct {
	// ID of the zone setting.
	ID param.Field[ZonesCiphersID] `json:"id"`
	// Value of the zone setting.
	Value param.Field[[]string] `json:"value"`
	ZonesBaseParam
}

func (r ZonesCiphersParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZonesCiphersParam) implementsZonesSettingUnionParam() {}

func (r ZonesCiphersParam) implementsZoneSettingUpdateParamsBodyUnion() {}

// Whether or not cname flattening is on.
type ZonesCnameFlattening struct {
	// How to flatten the cname destination.
	ID ZonesCnameFlatteningID `json:"id"`
	// Value of the cname flattening setting.
	Value ZonesCnameFlatteningValue `json:"value"`
	JSON  zonesCnameFlatteningJSON  `json:"-"`
	ZonesBase
}

// zonesCnameFlatteningJSON contains the JSON metadata for the struct
// [ZonesCnameFlattening]
type zonesCnameFlatteningJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonesCnameFlattening) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zonesCnameFlatteningJSON) RawJSON() string {
	return r.raw
}

func (r ZonesCnameFlattening) implementsZonesSetting() {}

func (r ZonesCnameFlattening) implementsZonesZoneSettingsResponseCollectionResult() {}

// How to flatten the cname destination.
type ZonesCnameFlatteningID string

const (
	ZonesCnameFlatteningIDCnameFlattening ZonesCnameFlatteningID = "cname_flattening"
)

func (r ZonesCnameFlatteningID) IsKnown() bool {
	switch r {
	case ZonesCnameFlatteningIDCnameFlattening:
		return true
	}
	return false
}

// Value of the cname flattening setting.
type ZonesCnameFlatteningValue string

const (
	ZonesCnameFlatteningValueFlattenAtRoot ZonesCnameFlatteningValue = "flatten_at_root"
	ZonesCnameFlatteningValueFlattenAll    ZonesCnameFlatteningValue = "flatten_all"
)

func (r ZonesCnameFlatteningValue) IsKnown() bool {
	switch r {
	case ZonesCnameFlatteningValueFlattenAtRoot, ZonesCnameFlatteningValueFlattenAll:
		return true
	}
	return false
}

// Whether or not cname flattening is on.
type ZonesCnameFlatteningParam struct {
	// How to flatten the cname destination.
	ID param.Field[ZonesCnameFlatteningID] `json:"id"`
	// Value of the cname flattening setting.
	Value param.Field[ZonesCnameFlatteningValue] `json:"value"`
	ZonesBaseParam
}

func (r ZonesCnameFlatteningParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZonesCnameFlatteningParam) implementsZonesSettingUnionParam() {}

func (r ZonesCnameFlatteningParam) implementsZoneSettingUpdateParamsBodyUnion() {}

// Development Mode temporarily allows you to enter development mode for your
// websites if you need to make changes to your site. This will bypass Cloudflare's
// accelerated cache and slow down your site, but is useful if you are making
// changes to cacheable content (like images, css, or JavaScript) and would like to
// see those changes right away. Once entered, development mode will last for 3
// hours and then automatically toggle off.
type ZonesDevelopmentMode struct {
	// ID of the zone setting.
	ID ZonesDevelopmentModeID `json:"id"`
	// Value of the zone setting. Notes: The interval (in seconds) from when
	// development mode expires (positive integer) or last expired (negative integer)
	// for the domain. If development mode has never been enabled, this value is false.
	TimeRemaining float64 `json:"time_remaining"`
	// Value of the zone setting.
	Value ZonesDevelopmentModeValue `json:"value"`
	JSON  zonesDevelopmentModeJSON  `json:"-"`
	ZonesBase
}

// zonesDevelopmentModeJSON contains the JSON metadata for the struct
// [ZonesDevelopmentMode]
type zonesDevelopmentModeJSON struct {
	ID            apijson.Field
	TimeRemaining apijson.Field
	Value         apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ZonesDevelopmentMode) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zonesDevelopmentModeJSON) RawJSON() string {
	return r.raw
}

func (r ZonesDevelopmentMode) implementsZonesSetting() {}

func (r ZonesDevelopmentMode) implementsZonesZoneSettingsResponseCollectionResult() {}

// ID of the zone setting.
type ZonesDevelopmentModeID string

const (
	ZonesDevelopmentModeIDDevelopmentMode ZonesDevelopmentModeID = "development_mode"
)

func (r ZonesDevelopmentModeID) IsKnown() bool {
	switch r {
	case ZonesDevelopmentModeIDDevelopmentMode:
		return true
	}
	return false
}

// Value of the zone setting.
type ZonesDevelopmentModeValue string

const (
	ZonesDevelopmentModeValueOn  ZonesDevelopmentModeValue = "on"
	ZonesDevelopmentModeValueOff ZonesDevelopmentModeValue = "off"
)

func (r ZonesDevelopmentModeValue) IsKnown() bool {
	switch r {
	case ZonesDevelopmentModeValueOn, ZonesDevelopmentModeValueOff:
		return true
	}
	return false
}

// Development Mode temporarily allows you to enter development mode for your
// websites if you need to make changes to your site. This will bypass Cloudflare's
// accelerated cache and slow down your site, but is useful if you are making
// changes to cacheable content (like images, css, or JavaScript) and would like to
// see those changes right away. Once entered, development mode will last for 3
// hours and then automatically toggle off.
type ZonesDevelopmentModeParam struct {
	// ID of the zone setting.
	ID param.Field[ZonesDevelopmentModeID] `json:"id"`
	// Value of the zone setting.
	Value param.Field[ZonesDevelopmentModeValue] `json:"value"`
	ZonesBaseParam
}

func (r ZonesDevelopmentModeParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZonesDevelopmentModeParam) implementsZonesSettingUnionParam() {}

func (r ZonesDevelopmentModeParam) implementsZoneSettingUpdateParamsBodyUnion() {}

// When enabled, Cloudflare will attempt to speed up overall page loads by serving
// `103` responses with `Link` headers from the final response. Refer to
// [Early Hints](https://developers.cloudflare.com/cache/about/early-hints) for
// more information.
type ZonesEarlyHints struct {
	// ID of the zone setting.
	ID ZonesEarlyHintsID `json:"id"`
	// Value of the zone setting.
	Value ZonesEarlyHintsValue `json:"value"`
	JSON  zonesEarlyHintsJSON  `json:"-"`
	ZonesBase
}

// zonesEarlyHintsJSON contains the JSON metadata for the struct [ZonesEarlyHints]
type zonesEarlyHintsJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonesEarlyHints) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zonesEarlyHintsJSON) RawJSON() string {
	return r.raw
}

func (r ZonesEarlyHints) implementsZonesSetting() {}

func (r ZonesEarlyHints) implementsZonesZoneSettingsResponseCollectionResult() {}

// ID of the zone setting.
type ZonesEarlyHintsID string

const (
	ZonesEarlyHintsIDEarlyHints ZonesEarlyHintsID = "early_hints"
)

func (r ZonesEarlyHintsID) IsKnown() bool {
	switch r {
	case ZonesEarlyHintsIDEarlyHints:
		return true
	}
	return false
}

// Value of the zone setting.
type ZonesEarlyHintsValue string

const (
	ZonesEarlyHintsValueOn  ZonesEarlyHintsValue = "on"
	ZonesEarlyHintsValueOff ZonesEarlyHintsValue = "off"
)

func (r ZonesEarlyHintsValue) IsKnown() bool {
	switch r {
	case ZonesEarlyHintsValueOn, ZonesEarlyHintsValueOff:
		return true
	}
	return false
}

// When enabled, Cloudflare will attempt to speed up overall page loads by serving
// `103` responses with `Link` headers from the final response. Refer to
// [Early Hints](https://developers.cloudflare.com/cache/about/early-hints) for
// more information.
type ZonesEarlyHintsParam struct {
	// ID of the zone setting.
	ID param.Field[ZonesEarlyHintsID] `json:"id"`
	// Value of the zone setting.
	Value param.Field[ZonesEarlyHintsValue] `json:"value"`
	ZonesBaseParam
}

func (r ZonesEarlyHintsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZonesEarlyHintsParam) implementsZonesSettingUnionParam() {}

func (r ZonesEarlyHintsParam) implementsZoneSettingUpdateParamsBodyUnion() {}

// HTTP/2 Edge Prioritization optimises the delivery of resources served through
// HTTP/2 to improve page load performance. It also supports fine control of
// content delivery when used in conjunction with Workers.
type ZonesH2Prioritization struct {
	// ID of the zone setting.
	ID ZonesH2PrioritizationID `json:"id"`
	// Value of the zone setting.
	Value ZonesH2PrioritizationValue `json:"value"`
	JSON  zonesH2PrioritizationJSON  `json:"-"`
	ZonesBase
}

// zonesH2PrioritizationJSON contains the JSON metadata for the struct
// [ZonesH2Prioritization]
type zonesH2PrioritizationJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonesH2Prioritization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zonesH2PrioritizationJSON) RawJSON() string {
	return r.raw
}

func (r ZonesH2Prioritization) implementsZonesSetting() {}

func (r ZonesH2Prioritization) implementsZonesZoneSettingsResponseCollectionResult() {}

// ID of the zone setting.
type ZonesH2PrioritizationID string

const (
	ZonesH2PrioritizationIDH2Prioritization ZonesH2PrioritizationID = "h2_prioritization"
)

func (r ZonesH2PrioritizationID) IsKnown() bool {
	switch r {
	case ZonesH2PrioritizationIDH2Prioritization:
		return true
	}
	return false
}

// Value of the zone setting.
type ZonesH2PrioritizationValue string

const (
	ZonesH2PrioritizationValueOn     ZonesH2PrioritizationValue = "on"
	ZonesH2PrioritizationValueOff    ZonesH2PrioritizationValue = "off"
	ZonesH2PrioritizationValueCustom ZonesH2PrioritizationValue = "custom"
)

func (r ZonesH2PrioritizationValue) IsKnown() bool {
	switch r {
	case ZonesH2PrioritizationValueOn, ZonesH2PrioritizationValueOff, ZonesH2PrioritizationValueCustom:
		return true
	}
	return false
}

// HTTP/2 Edge Prioritization optimises the delivery of resources served through
// HTTP/2 to improve page load performance. It also supports fine control of
// content delivery when used in conjunction with Workers.
type ZonesH2PrioritizationParam struct {
	// ID of the zone setting.
	ID param.Field[ZonesH2PrioritizationID] `json:"id"`
	// Value of the zone setting.
	Value param.Field[ZonesH2PrioritizationValue] `json:"value"`
	ZonesBaseParam
}

func (r ZonesH2PrioritizationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZonesH2PrioritizationParam) implementsZonesSettingUnionParam() {}

func (r ZonesH2PrioritizationParam) implementsZoneSettingUpdateParamsBodyUnion() {}

// When enabled, the Hotlink Protection option ensures that other sites cannot suck
// up your bandwidth by building pages that use images hosted on your site. Anytime
// a request for an image on your site hits Cloudflare, we check to ensure that
// it's not another site requesting them. People will still be able to download and
// view images from your page, but other sites won't be able to steal them for use
// on their own pages.
// (https://support.cloudflare.com/hc/en-us/articles/200170026).
type ZonesHotlinkProtection struct {
	// ID of the zone setting.
	ID ZonesHotlinkProtectionID `json:"id"`
	// Value of the zone setting.
	Value ZonesHotlinkProtectionValue `json:"value"`
	JSON  zonesHotlinkProtectionJSON  `json:"-"`
	ZonesBase
}

// zonesHotlinkProtectionJSON contains the JSON metadata for the struct
// [ZonesHotlinkProtection]
type zonesHotlinkProtectionJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonesHotlinkProtection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zonesHotlinkProtectionJSON) RawJSON() string {
	return r.raw
}

func (r ZonesHotlinkProtection) implementsZonesSetting() {}

func (r ZonesHotlinkProtection) implementsZonesZoneSettingsResponseCollectionResult() {}

// ID of the zone setting.
type ZonesHotlinkProtectionID string

const (
	ZonesHotlinkProtectionIDHotlinkProtection ZonesHotlinkProtectionID = "hotlink_protection"
)

func (r ZonesHotlinkProtectionID) IsKnown() bool {
	switch r {
	case ZonesHotlinkProtectionIDHotlinkProtection:
		return true
	}
	return false
}

// Value of the zone setting.
type ZonesHotlinkProtectionValue string

const (
	ZonesHotlinkProtectionValueOn  ZonesHotlinkProtectionValue = "on"
	ZonesHotlinkProtectionValueOff ZonesHotlinkProtectionValue = "off"
)

func (r ZonesHotlinkProtectionValue) IsKnown() bool {
	switch r {
	case ZonesHotlinkProtectionValueOn, ZonesHotlinkProtectionValueOff:
		return true
	}
	return false
}

// When enabled, the Hotlink Protection option ensures that other sites cannot suck
// up your bandwidth by building pages that use images hosted on your site. Anytime
// a request for an image on your site hits Cloudflare, we check to ensure that
// it's not another site requesting them. People will still be able to download and
// view images from your page, but other sites won't be able to steal them for use
// on their own pages.
// (https://support.cloudflare.com/hc/en-us/articles/200170026).
type ZonesHotlinkProtectionParam struct {
	// ID of the zone setting.
	ID param.Field[ZonesHotlinkProtectionID] `json:"id"`
	// Value of the zone setting.
	Value param.Field[ZonesHotlinkProtectionValue] `json:"value"`
	ZonesBaseParam
}

func (r ZonesHotlinkProtectionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZonesHotlinkProtectionParam) implementsZonesSettingUnionParam() {}

func (r ZonesHotlinkProtectionParam) implementsZoneSettingUpdateParamsBodyUnion() {}

// HTTP2 enabled for this zone.
type ZonesHttp2 struct {
	// ID of the zone setting.
	ID ZonesHttp2ID `json:"id"`
	// Value of the HTTP2 setting.
	Value ZonesHttp2Value `json:"value"`
	JSON  zonesHttp2JSON  `json:"-"`
	ZonesBase
}

// zonesHttp2JSON contains the JSON metadata for the struct [ZonesHttp2]
type zonesHttp2JSON struct {
	ID          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonesHttp2) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zonesHttp2JSON) RawJSON() string {
	return r.raw
}

func (r ZonesHttp2) implementsZonesSetting() {}

func (r ZonesHttp2) implementsZonesZoneSettingsResponseCollectionResult() {}

// ID of the zone setting.
type ZonesHttp2ID string

const (
	ZonesHttp2IDHttp2 ZonesHttp2ID = "http2"
)

func (r ZonesHttp2ID) IsKnown() bool {
	switch r {
	case ZonesHttp2IDHttp2:
		return true
	}
	return false
}

// Value of the HTTP2 setting.
type ZonesHttp2Value string

const (
	ZonesHttp2ValueOn  ZonesHttp2Value = "on"
	ZonesHttp2ValueOff ZonesHttp2Value = "off"
)

func (r ZonesHttp2Value) IsKnown() bool {
	switch r {
	case ZonesHttp2ValueOn, ZonesHttp2ValueOff:
		return true
	}
	return false
}

// HTTP2 enabled for this zone.
type ZonesHttp2Param struct {
	// ID of the zone setting.
	ID param.Field[ZonesHttp2ID] `json:"id"`
	// Value of the HTTP2 setting.
	Value param.Field[ZonesHttp2Value] `json:"value"`
	ZonesBaseParam
}

func (r ZonesHttp2Param) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZonesHttp2Param) implementsZonesSettingUnionParam() {}

func (r ZonesHttp2Param) implementsZoneSettingUpdateParamsBodyUnion() {}

// HTTP3 enabled for this zone.
type ZonesHttp3 struct {
	// ID of the zone setting.
	ID ZonesHttp3ID `json:"id"`
	// Value of the HTTP3 setting.
	Value ZonesHttp3Value `json:"value"`
	JSON  zonesHttp3JSON  `json:"-"`
	ZonesBase
}

// zonesHttp3JSON contains the JSON metadata for the struct [ZonesHttp3]
type zonesHttp3JSON struct {
	ID          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonesHttp3) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zonesHttp3JSON) RawJSON() string {
	return r.raw
}

func (r ZonesHttp3) implementsZonesSetting() {}

func (r ZonesHttp3) implementsZonesZoneSettingsResponseCollectionResult() {}

// ID of the zone setting.
type ZonesHttp3ID string

const (
	ZonesHttp3IDHttp3 ZonesHttp3ID = "http3"
)

func (r ZonesHttp3ID) IsKnown() bool {
	switch r {
	case ZonesHttp3IDHttp3:
		return true
	}
	return false
}

// Value of the HTTP3 setting.
type ZonesHttp3Value string

const (
	ZonesHttp3ValueOn  ZonesHttp3Value = "on"
	ZonesHttp3ValueOff ZonesHttp3Value = "off"
)

func (r ZonesHttp3Value) IsKnown() bool {
	switch r {
	case ZonesHttp3ValueOn, ZonesHttp3ValueOff:
		return true
	}
	return false
}

// HTTP3 enabled for this zone.
type ZonesHttp3Param struct {
	// ID of the zone setting.
	ID param.Field[ZonesHttp3ID] `json:"id"`
	// Value of the HTTP3 setting.
	Value param.Field[ZonesHttp3Value] `json:"value"`
	ZonesBaseParam
}

func (r ZonesHttp3Param) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZonesHttp3Param) implementsZonesSettingUnionParam() {}

func (r ZonesHttp3Param) implementsZoneSettingUpdateParamsBodyUnion() {}

// Image Resizing provides on-demand resizing, conversion and optimisation for
// images served through Cloudflare's network. Refer to the
// [Image Resizing documentation](https://developers.cloudflare.com/images/) for
// more information.
type ZonesImageResizing struct {
	// ID of the zone setting.
	ID ZonesImageResizingID `json:"id"`
	// Whether the feature is enabled, disabled, or enabled in `open proxy` mode.
	Value ZonesImageResizingValue `json:"value"`
	JSON  zonesImageResizingJSON  `json:"-"`
	ZonesBase
}

// zonesImageResizingJSON contains the JSON metadata for the struct
// [ZonesImageResizing]
type zonesImageResizingJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonesImageResizing) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zonesImageResizingJSON) RawJSON() string {
	return r.raw
}

func (r ZonesImageResizing) implementsZonesSetting() {}

func (r ZonesImageResizing) implementsZonesZoneSettingsResponseCollectionResult() {}

// ID of the zone setting.
type ZonesImageResizingID string

const (
	ZonesImageResizingIDImageResizing ZonesImageResizingID = "image_resizing"
)

func (r ZonesImageResizingID) IsKnown() bool {
	switch r {
	case ZonesImageResizingIDImageResizing:
		return true
	}
	return false
}

// Whether the feature is enabled, disabled, or enabled in `open proxy` mode.
type ZonesImageResizingValue string

const (
	ZonesImageResizingValueOn   ZonesImageResizingValue = "on"
	ZonesImageResizingValueOff  ZonesImageResizingValue = "off"
	ZonesImageResizingValueOpen ZonesImageResizingValue = "open"
)

func (r ZonesImageResizingValue) IsKnown() bool {
	switch r {
	case ZonesImageResizingValueOn, ZonesImageResizingValueOff, ZonesImageResizingValueOpen:
		return true
	}
	return false
}

// Image Resizing provides on-demand resizing, conversion and optimisation for
// images served through Cloudflare's network. Refer to the
// [Image Resizing documentation](https://developers.cloudflare.com/images/) for
// more information.
type ZonesImageResizingParam struct {
	// ID of the zone setting.
	ID param.Field[ZonesImageResizingID] `json:"id"`
	// Whether the feature is enabled, disabled, or enabled in `open proxy` mode.
	Value param.Field[ZonesImageResizingValue] `json:"value"`
	ZonesBaseParam
}

func (r ZonesImageResizingParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZonesImageResizingParam) implementsZonesSettingUnionParam() {}

// Enable IPv6 on all subdomains that are Cloudflare enabled.
// (https://support.cloudflare.com/hc/en-us/articles/200168586).
type ZonesIpv6 struct {
	// ID of the zone setting.
	ID ZonesIpv6ID `json:"id"`
	// Value of the zone setting.
	Value ZonesIpv6Value `json:"value"`
	JSON  zonesIpv6JSON  `json:"-"`
	ZonesBase
}

// zonesIpv6JSON contains the JSON metadata for the struct [ZonesIpv6]
type zonesIpv6JSON struct {
	ID          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonesIpv6) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zonesIpv6JSON) RawJSON() string {
	return r.raw
}

func (r ZonesIpv6) implementsZonesSetting() {}

func (r ZonesIpv6) implementsZonesZoneSettingsResponseCollectionResult() {}

// ID of the zone setting.
type ZonesIpv6ID string

const (
	ZonesIpv6IDIpv6 ZonesIpv6ID = "ipv6"
)

func (r ZonesIpv6ID) IsKnown() bool {
	switch r {
	case ZonesIpv6IDIpv6:
		return true
	}
	return false
}

// Value of the zone setting.
type ZonesIpv6Value string

const (
	ZonesIpv6ValueOff ZonesIpv6Value = "off"
	ZonesIpv6ValueOn  ZonesIpv6Value = "on"
)

func (r ZonesIpv6Value) IsKnown() bool {
	switch r {
	case ZonesIpv6ValueOff, ZonesIpv6ValueOn:
		return true
	}
	return false
}

// Enable IPv6 on all subdomains that are Cloudflare enabled.
// (https://support.cloudflare.com/hc/en-us/articles/200168586).
type ZonesIpv6Param struct {
	// ID of the zone setting.
	ID param.Field[ZonesIpv6ID] `json:"id"`
	// Value of the zone setting.
	Value param.Field[ZonesIpv6Value] `json:"value"`
	ZonesBaseParam
}

func (r ZonesIpv6Param) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZonesIpv6Param) implementsZonesSettingUnionParam() {}

func (r ZonesIpv6Param) implementsZoneSettingUpdateParamsBodyUnion() {}

// Maximum size of an allowable upload.
type ZonesMaxUpload struct {
	// identifier of the zone setting.
	ID ZonesMaxUploadID `json:"id"`
	// Value of the zone setting. Notes: The size depends on the plan level of the
	// zone. (Enterprise = 500, Business = 200, Pro = 100, Free = 100)
	Value ZonesMaxUploadValue `json:"value"`
	JSON  zonesMaxUploadJSON  `json:"-"`
	ZonesBase
}

// zonesMaxUploadJSON contains the JSON metadata for the struct [ZonesMaxUpload]
type zonesMaxUploadJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonesMaxUpload) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zonesMaxUploadJSON) RawJSON() string {
	return r.raw
}

func (r ZonesMaxUpload) implementsZonesSetting() {}

func (r ZonesMaxUpload) implementsZonesZoneSettingsResponseCollectionResult() {}

// identifier of the zone setting.
type ZonesMaxUploadID string

const (
	ZonesMaxUploadIDMaxUpload ZonesMaxUploadID = "max_upload"
)

func (r ZonesMaxUploadID) IsKnown() bool {
	switch r {
	case ZonesMaxUploadIDMaxUpload:
		return true
	}
	return false
}

// Value of the zone setting. Notes: The size depends on the plan level of the
// zone. (Enterprise = 500, Business = 200, Pro = 100, Free = 100)
type ZonesMaxUploadValue float64

const (
	ZonesMaxUploadValue100 ZonesMaxUploadValue = 100
	ZonesMaxUploadValue200 ZonesMaxUploadValue = 200
	ZonesMaxUploadValue500 ZonesMaxUploadValue = 500
)

func (r ZonesMaxUploadValue) IsKnown() bool {
	switch r {
	case ZonesMaxUploadValue100, ZonesMaxUploadValue200, ZonesMaxUploadValue500:
		return true
	}
	return false
}

// Maximum size of an allowable upload.
type ZonesMaxUploadParam struct {
	// identifier of the zone setting.
	ID param.Field[ZonesMaxUploadID] `json:"id"`
	// Value of the zone setting. Notes: The size depends on the plan level of the
	// zone. (Enterprise = 500, Business = 200, Pro = 100, Free = 100)
	Value param.Field[ZonesMaxUploadValue] `json:"value"`
	ZonesBaseParam
}

func (r ZonesMaxUploadParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZonesMaxUploadParam) implementsZonesSettingUnionParam() {}

func (r ZonesMaxUploadParam) implementsZoneSettingUpdateParamsBodyUnion() {}

// Only accepts HTTPS requests that use at least the TLS protocol version
// specified. For example, if TLS 1.1 is selected, TLS 1.0 connections will be
// rejected, while 1.1, 1.2, and 1.3 (if enabled) will be permitted.
type ZonesMinTlsVersion struct {
	// ID of the zone setting.
	ID ZonesMinTlsVersionID `json:"id"`
	// Value of the zone setting.
	Value ZonesMinTlsVersionValue `json:"value"`
	JSON  zonesMinTlsVersionJSON  `json:"-"`
	ZonesBase
}

// zonesMinTlsVersionJSON contains the JSON metadata for the struct
// [ZonesMinTlsVersion]
type zonesMinTlsVersionJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonesMinTlsVersion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zonesMinTlsVersionJSON) RawJSON() string {
	return r.raw
}

func (r ZonesMinTlsVersion) implementsZonesSetting() {}

func (r ZonesMinTlsVersion) implementsZonesZoneSettingsResponseCollectionResult() {}

// ID of the zone setting.
type ZonesMinTlsVersionID string

const (
	ZonesMinTlsVersionIDMinTlsVersion ZonesMinTlsVersionID = "min_tls_version"
)

func (r ZonesMinTlsVersionID) IsKnown() bool {
	switch r {
	case ZonesMinTlsVersionIDMinTlsVersion:
		return true
	}
	return false
}

// Value of the zone setting.
type ZonesMinTlsVersionValue string

const (
	ZonesMinTlsVersionValue1_0 ZonesMinTlsVersionValue = "1.0"
	ZonesMinTlsVersionValue1_1 ZonesMinTlsVersionValue = "1.1"
	ZonesMinTlsVersionValue1_2 ZonesMinTlsVersionValue = "1.2"
	ZonesMinTlsVersionValue1_3 ZonesMinTlsVersionValue = "1.3"
)

func (r ZonesMinTlsVersionValue) IsKnown() bool {
	switch r {
	case ZonesMinTlsVersionValue1_0, ZonesMinTlsVersionValue1_1, ZonesMinTlsVersionValue1_2, ZonesMinTlsVersionValue1_3:
		return true
	}
	return false
}

// Only accepts HTTPS requests that use at least the TLS protocol version
// specified. For example, if TLS 1.1 is selected, TLS 1.0 connections will be
// rejected, while 1.1, 1.2, and 1.3 (if enabled) will be permitted.
type ZonesMinTlsVersionParam struct {
	// ID of the zone setting.
	ID param.Field[ZonesMinTlsVersionID] `json:"id"`
	// Value of the zone setting.
	Value param.Field[ZonesMinTlsVersionValue] `json:"value"`
	ZonesBaseParam
}

func (r ZonesMinTlsVersionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZonesMinTlsVersionParam) implementsZonesSettingUnionParam() {}

func (r ZonesMinTlsVersionParam) implementsZoneSettingUpdateParamsBodyUnion() {}

// Enable Network Error Logging reporting on your zone. (Beta)
type ZonesNel struct {
	// Zone setting identifier.
	ID ZonesNelID `json:"id"`
	// Value of the zone setting.
	Value ZonesNelValue `json:"value"`
	JSON  zonesNelJSON  `json:"-"`
	ZonesBase
}

// zonesNelJSON contains the JSON metadata for the struct [ZonesNel]
type zonesNelJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonesNel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zonesNelJSON) RawJSON() string {
	return r.raw
}

func (r ZonesNel) implementsZonesSetting() {}

func (r ZonesNel) implementsZonesZoneSettingsResponseCollectionResult() {}

// Zone setting identifier.
type ZonesNelID string

const (
	ZonesNelIDNel ZonesNelID = "nel"
)

func (r ZonesNelID) IsKnown() bool {
	switch r {
	case ZonesNelIDNel:
		return true
	}
	return false
}

// Value of the zone setting.
type ZonesNelValue struct {
	Enabled bool              `json:"enabled"`
	JSON    zonesNelValueJSON `json:"-"`
}

// zonesNelValueJSON contains the JSON metadata for the struct [ZonesNelValue]
type zonesNelValueJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonesNelValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zonesNelValueJSON) RawJSON() string {
	return r.raw
}

// Enable Network Error Logging reporting on your zone. (Beta)
type ZonesNelParam struct {
	// Zone setting identifier.
	ID param.Field[ZonesNelID] `json:"id"`
	// Value of the zone setting.
	Value param.Field[ZonesNelValueParam] `json:"value"`
	ZonesBaseParam
}

func (r ZonesNelParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZonesNelParam) implementsZonesSettingUnionParam() {}

func (r ZonesNelParam) implementsZoneSettingUpdateParamsBodyUnion() {}

// Value of the zone setting.
type ZonesNelValueParam struct {
	Enabled param.Field[bool] `json:"enabled"`
}

func (r ZonesNelValueParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Add an Alt-Svc header to all legitimate requests from Tor, allowing the
// connection to use our onion services instead of exit nodes.
type ZonesOpportunisticOnion struct {
	// ID of the zone setting.
	ID ZonesOpportunisticOnionID `json:"id"`
	// Value of the zone setting. Notes: Default value depends on the zone's plan
	// level.
	Value ZonesOpportunisticOnionValue `json:"value"`
	JSON  zonesOpportunisticOnionJSON  `json:"-"`
	ZonesBase
}

// zonesOpportunisticOnionJSON contains the JSON metadata for the struct
// [ZonesOpportunisticOnion]
type zonesOpportunisticOnionJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonesOpportunisticOnion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zonesOpportunisticOnionJSON) RawJSON() string {
	return r.raw
}

func (r ZonesOpportunisticOnion) implementsZonesSetting() {}

func (r ZonesOpportunisticOnion) implementsZonesZoneSettingsResponseCollectionResult() {}

// ID of the zone setting.
type ZonesOpportunisticOnionID string

const (
	ZonesOpportunisticOnionIDOpportunisticOnion ZonesOpportunisticOnionID = "opportunistic_onion"
)

func (r ZonesOpportunisticOnionID) IsKnown() bool {
	switch r {
	case ZonesOpportunisticOnionIDOpportunisticOnion:
		return true
	}
	return false
}

// Value of the zone setting. Notes: Default value depends on the zone's plan
// level.
type ZonesOpportunisticOnionValue string

const (
	ZonesOpportunisticOnionValueOn  ZonesOpportunisticOnionValue = "on"
	ZonesOpportunisticOnionValueOff ZonesOpportunisticOnionValue = "off"
)

func (r ZonesOpportunisticOnionValue) IsKnown() bool {
	switch r {
	case ZonesOpportunisticOnionValueOn, ZonesOpportunisticOnionValueOff:
		return true
	}
	return false
}

// Add an Alt-Svc header to all legitimate requests from Tor, allowing the
// connection to use our onion services instead of exit nodes.
type ZonesOpportunisticOnionParam struct {
	// ID of the zone setting.
	ID param.Field[ZonesOpportunisticOnionID] `json:"id"`
	// Value of the zone setting. Notes: Default value depends on the zone's plan
	// level.
	Value param.Field[ZonesOpportunisticOnionValue] `json:"value"`
	ZonesBaseParam
}

func (r ZonesOpportunisticOnionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZonesOpportunisticOnionParam) implementsZonesSettingUnionParam() {}

func (r ZonesOpportunisticOnionParam) implementsZoneSettingUpdateParamsBodyUnion() {}

// Orange to Orange (O2O) allows zones on Cloudflare to CNAME to other zones also
// on Cloudflare.
type ZonesOrangeToOrange struct {
	// ID of the zone setting.
	ID ZonesOrangeToOrangeID `json:"id"`
	// Value of the zone setting.
	Value ZonesOrangeToOrangeValue `json:"value"`
	JSON  zonesOrangeToOrangeJSON  `json:"-"`
	ZonesBase
}

// zonesOrangeToOrangeJSON contains the JSON metadata for the struct
// [ZonesOrangeToOrange]
type zonesOrangeToOrangeJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonesOrangeToOrange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zonesOrangeToOrangeJSON) RawJSON() string {
	return r.raw
}

func (r ZonesOrangeToOrange) implementsZonesSetting() {}

func (r ZonesOrangeToOrange) implementsZonesZoneSettingsResponseCollectionResult() {}

// ID of the zone setting.
type ZonesOrangeToOrangeID string

const (
	ZonesOrangeToOrangeIDOrangeToOrange ZonesOrangeToOrangeID = "orange_to_orange"
)

func (r ZonesOrangeToOrangeID) IsKnown() bool {
	switch r {
	case ZonesOrangeToOrangeIDOrangeToOrange:
		return true
	}
	return false
}

// Value of the zone setting.
type ZonesOrangeToOrangeValue string

const (
	ZonesOrangeToOrangeValueOn  ZonesOrangeToOrangeValue = "on"
	ZonesOrangeToOrangeValueOff ZonesOrangeToOrangeValue = "off"
)

func (r ZonesOrangeToOrangeValue) IsKnown() bool {
	switch r {
	case ZonesOrangeToOrangeValueOn, ZonesOrangeToOrangeValueOff:
		return true
	}
	return false
}

// Orange to Orange (O2O) allows zones on Cloudflare to CNAME to other zones also
// on Cloudflare.
type ZonesOrangeToOrangeParam struct {
	// ID of the zone setting.
	ID param.Field[ZonesOrangeToOrangeID] `json:"id"`
	// Value of the zone setting.
	Value param.Field[ZonesOrangeToOrangeValue] `json:"value"`
	ZonesBaseParam
}

func (r ZonesOrangeToOrangeParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZonesOrangeToOrangeParam) implementsZonesSettingUnionParam() {}

func (r ZonesOrangeToOrangeParam) implementsZoneSettingUpdateParamsBodyUnion() {}

// Cloudflare will prefetch any URLs that are included in the response headers.
// This is limited to Enterprise Zones.
type ZonesPrefetchPreload struct {
	// ID of the zone setting.
	ID ZonesPrefetchPreloadID `json:"id"`
	// Value of the zone setting.
	Value ZonesPrefetchPreloadValue `json:"value"`
	JSON  zonesPrefetchPreloadJSON  `json:"-"`
	ZonesBase
}

// zonesPrefetchPreloadJSON contains the JSON metadata for the struct
// [ZonesPrefetchPreload]
type zonesPrefetchPreloadJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonesPrefetchPreload) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zonesPrefetchPreloadJSON) RawJSON() string {
	return r.raw
}

func (r ZonesPrefetchPreload) implementsZonesSetting() {}

func (r ZonesPrefetchPreload) implementsZonesZoneSettingsResponseCollectionResult() {}

// ID of the zone setting.
type ZonesPrefetchPreloadID string

const (
	ZonesPrefetchPreloadIDPrefetchPreload ZonesPrefetchPreloadID = "prefetch_preload"
)

func (r ZonesPrefetchPreloadID) IsKnown() bool {
	switch r {
	case ZonesPrefetchPreloadIDPrefetchPreload:
		return true
	}
	return false
}

// Value of the zone setting.
type ZonesPrefetchPreloadValue string

const (
	ZonesPrefetchPreloadValueOn  ZonesPrefetchPreloadValue = "on"
	ZonesPrefetchPreloadValueOff ZonesPrefetchPreloadValue = "off"
)

func (r ZonesPrefetchPreloadValue) IsKnown() bool {
	switch r {
	case ZonesPrefetchPreloadValueOn, ZonesPrefetchPreloadValueOff:
		return true
	}
	return false
}

// Cloudflare will prefetch any URLs that are included in the response headers.
// This is limited to Enterprise Zones.
type ZonesPrefetchPreloadParam struct {
	// ID of the zone setting.
	ID param.Field[ZonesPrefetchPreloadID] `json:"id"`
	// Value of the zone setting.
	Value param.Field[ZonesPrefetchPreloadValue] `json:"value"`
	ZonesBaseParam
}

func (r ZonesPrefetchPreloadParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZonesPrefetchPreloadParam) implementsZonesSettingUnionParam() {}

func (r ZonesPrefetchPreloadParam) implementsZoneSettingUpdateParamsBodyUnion() {}

// Privacy Pass v1 was a browser extension developed by the Privacy Pass Team to
// improve the browsing experience for your visitors by allowing users to reduce
// the number of CAPTCHAs shown.
// (https://support.cloudflare.com/hc/en-us/articles/115001992652-Privacy-Pass).
//
// Deprecated: Privacy Pass v1 was deprecated in 2023. (Announcement -
// https://blog.cloudflare.com/privacy-pass-standard/) and (API deprecation
// details -
// https://developers.cloudflare.com/fundamentals/api/reference/deprecations/#2024-03-31)
type ZonesPrivacyPass struct {
	// ID of the zone setting.
	ID ZonesPrivacyPassID `json:"id"`
	// Value of the Privacy Pass v1 (deprecated) zone setting
	//
	// Deprecated: Privacy Pass v1 was deprecated in 2023. (Announcement -
	// https://blog.cloudflare.com/privacy-pass-standard/) and (API deprecation
	// details -
	// https://developers.cloudflare.com/fundamentals/api/reference/deprecations/#2024-03-31)
	Value ZonesPrivacyPassValue `json:"value"`
	JSON  zonesPrivacyPassJSON  `json:"-"`
	ZonesBase
}

// zonesPrivacyPassJSON contains the JSON metadata for the struct
// [ZonesPrivacyPass]
type zonesPrivacyPassJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonesPrivacyPass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zonesPrivacyPassJSON) RawJSON() string {
	return r.raw
}

func (r ZonesPrivacyPass) implementsZonesSetting() {}

func (r ZonesPrivacyPass) implementsZonesZoneSettingsResponseCollectionResult() {}

// ID of the zone setting.
type ZonesPrivacyPassID string

const (
	ZonesPrivacyPassIDPrivacyPass ZonesPrivacyPassID = "privacy_pass"
)

func (r ZonesPrivacyPassID) IsKnown() bool {
	switch r {
	case ZonesPrivacyPassIDPrivacyPass:
		return true
	}
	return false
}

// Value of the Privacy Pass v1 (deprecated) zone setting
type ZonesPrivacyPassValue string

const (
	ZonesPrivacyPassValueOn  ZonesPrivacyPassValue = "on"
	ZonesPrivacyPassValueOff ZonesPrivacyPassValue = "off"
)

func (r ZonesPrivacyPassValue) IsKnown() bool {
	switch r {
	case ZonesPrivacyPassValueOn, ZonesPrivacyPassValueOff:
		return true
	}
	return false
}

// Privacy Pass v1 was a browser extension developed by the Privacy Pass Team to
// improve the browsing experience for your visitors by allowing users to reduce
// the number of CAPTCHAs shown.
// (https://support.cloudflare.com/hc/en-us/articles/115001992652-Privacy-Pass).
//
// Deprecated: Privacy Pass v1 was deprecated in 2023. (Announcement -
// https://blog.cloudflare.com/privacy-pass-standard/) and (API deprecation
// details -
// https://developers.cloudflare.com/fundamentals/api/reference/deprecations/#2024-03-31)
type ZonesPrivacyPassParam struct {
	// ID of the zone setting.
	ID param.Field[ZonesPrivacyPassID] `json:"id"`
	// Value of the Privacy Pass v1 (deprecated) zone setting
	Value param.Field[ZonesPrivacyPassValue] `json:"value"`
	ZonesBaseParam
}

func (r ZonesPrivacyPassParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZonesPrivacyPassParam) implementsZonesSettingUnionParam() {}

func (r ZonesPrivacyPassParam) implementsZoneSettingUpdateParamsBodyUnion() {}

// Maximum time between two read operations from origin.
type ZonesProxyReadTimeout struct {
	// ID of the zone setting.
	ID ZonesProxyReadTimeoutID `json:"id"`
	// Value of the zone setting. Notes: Value must be between 1 and 6000
	Value float64                   `json:"value"`
	JSON  zonesProxyReadTimeoutJSON `json:"-"`
	ZonesBase
}

// zonesProxyReadTimeoutJSON contains the JSON metadata for the struct
// [ZonesProxyReadTimeout]
type zonesProxyReadTimeoutJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonesProxyReadTimeout) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zonesProxyReadTimeoutJSON) RawJSON() string {
	return r.raw
}

func (r ZonesProxyReadTimeout) implementsZonesSetting() {}

func (r ZonesProxyReadTimeout) implementsZonesZoneSettingsResponseCollectionResult() {}

// ID of the zone setting.
type ZonesProxyReadTimeoutID string

const (
	ZonesProxyReadTimeoutIDProxyReadTimeout ZonesProxyReadTimeoutID = "proxy_read_timeout"
)

func (r ZonesProxyReadTimeoutID) IsKnown() bool {
	switch r {
	case ZonesProxyReadTimeoutIDProxyReadTimeout:
		return true
	}
	return false
}

// Maximum time between two read operations from origin.
type ZonesProxyReadTimeoutParam struct {
	// ID of the zone setting.
	ID param.Field[ZonesProxyReadTimeoutID] `json:"id"`
	// Value of the zone setting. Notes: Value must be between 1 and 6000
	Value param.Field[float64] `json:"value"`
	ZonesBaseParam
}

func (r ZonesProxyReadTimeoutParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZonesProxyReadTimeoutParam) implementsZonesSettingUnionParam() {}

func (r ZonesProxyReadTimeoutParam) implementsZoneSettingUpdateParamsBodyUnion() {}

// The value set for the Pseudo IPv4 setting.
type ZonesPseudoIpv4 struct {
	// Value of the Pseudo IPv4 setting.
	ID ZonesPseudoIpv4ID `json:"id"`
	// Value of the Pseudo IPv4 setting.
	Value ZonesPseudoIpv4Value `json:"value"`
	JSON  zonesPseudoIpv4JSON  `json:"-"`
	ZonesBase
}

// zonesPseudoIpv4JSON contains the JSON metadata for the struct [ZonesPseudoIpv4]
type zonesPseudoIpv4JSON struct {
	ID          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonesPseudoIpv4) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zonesPseudoIpv4JSON) RawJSON() string {
	return r.raw
}

func (r ZonesPseudoIpv4) implementsZonesSetting() {}

func (r ZonesPseudoIpv4) implementsZonesZoneSettingsResponseCollectionResult() {}

// Value of the Pseudo IPv4 setting.
type ZonesPseudoIpv4ID string

const (
	ZonesPseudoIpv4IDPseudoIpv4 ZonesPseudoIpv4ID = "pseudo_ipv4"
)

func (r ZonesPseudoIpv4ID) IsKnown() bool {
	switch r {
	case ZonesPseudoIpv4IDPseudoIpv4:
		return true
	}
	return false
}

// Value of the Pseudo IPv4 setting.
type ZonesPseudoIpv4Value string

const (
	ZonesPseudoIpv4ValueOff             ZonesPseudoIpv4Value = "off"
	ZonesPseudoIpv4ValueAddHeader       ZonesPseudoIpv4Value = "add_header"
	ZonesPseudoIpv4ValueOverwriteHeader ZonesPseudoIpv4Value = "overwrite_header"
)

func (r ZonesPseudoIpv4Value) IsKnown() bool {
	switch r {
	case ZonesPseudoIpv4ValueOff, ZonesPseudoIpv4ValueAddHeader, ZonesPseudoIpv4ValueOverwriteHeader:
		return true
	}
	return false
}

// The value set for the Pseudo IPv4 setting.
type ZonesPseudoIpv4Param struct {
	// Value of the Pseudo IPv4 setting.
	ID param.Field[ZonesPseudoIpv4ID] `json:"id"`
	// Value of the Pseudo IPv4 setting.
	Value param.Field[ZonesPseudoIpv4Value] `json:"value"`
	ZonesBaseParam
}

func (r ZonesPseudoIpv4Param) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZonesPseudoIpv4Param) implementsZonesSettingUnionParam() {}

func (r ZonesPseudoIpv4Param) implementsZoneSettingUpdateParamsBodyUnion() {}

// Automatically replace insecure JavaScript libraries with safer and faster
// alternatives provided under cdnjs and powered by Cloudflare. Currently supports
// the following libraries: Polyfill under polyfill.io.
type ZonesReplaceInsecureJs struct {
	// ID of the zone setting.
	ID ZonesReplaceInsecureJsID `json:"id"`
	// Value of the zone setting.
	Value ZonesReplaceInsecureJsValue `json:"value"`
	JSON  zonesReplaceInsecureJsJSON  `json:"-"`
	ZonesBase
}

// zonesReplaceInsecureJsJSON contains the JSON metadata for the struct
// [ZonesReplaceInsecureJs]
type zonesReplaceInsecureJsJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonesReplaceInsecureJs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zonesReplaceInsecureJsJSON) RawJSON() string {
	return r.raw
}

func (r ZonesReplaceInsecureJs) implementsZonesSetting() {}

func (r ZonesReplaceInsecureJs) implementsZonesZoneSettingsResponseCollectionResult() {}

// ID of the zone setting.
type ZonesReplaceInsecureJsID string

const (
	ZonesReplaceInsecureJsIDReplaceInsecureJs ZonesReplaceInsecureJsID = "replace_insecure_js"
)

func (r ZonesReplaceInsecureJsID) IsKnown() bool {
	switch r {
	case ZonesReplaceInsecureJsIDReplaceInsecureJs:
		return true
	}
	return false
}

// Value of the zone setting.
type ZonesReplaceInsecureJsValue string

const (
	ZonesReplaceInsecureJsValueOn  ZonesReplaceInsecureJsValue = "on"
	ZonesReplaceInsecureJsValueOff ZonesReplaceInsecureJsValue = "off"
)

func (r ZonesReplaceInsecureJsValue) IsKnown() bool {
	switch r {
	case ZonesReplaceInsecureJsValueOn, ZonesReplaceInsecureJsValueOff:
		return true
	}
	return false
}

// Automatically replace insecure JavaScript libraries with safer and faster
// alternatives provided under cdnjs and powered by Cloudflare. Currently supports
// the following libraries: Polyfill under polyfill.io.
type ZonesReplaceInsecureJsParam struct {
	// ID of the zone setting.
	ID param.Field[ZonesReplaceInsecureJsID] `json:"id"`
	// Value of the zone setting.
	Value param.Field[ZonesReplaceInsecureJsValue] `json:"value"`
	ZonesBaseParam
}

func (r ZonesReplaceInsecureJsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZonesReplaceInsecureJsParam) implementsZonesSettingUnionParam() {}

func (r ZonesReplaceInsecureJsParam) implementsZoneSettingUpdateParamsBodyUnion() {}

// Reply to all requests for URLs that use "http" with a 301 redirect to the
// equivalent "https" URL. If you only want to redirect for a subset of requests,
// consider creating an "Always use HTTPS" page rule.
type ZonesSchemasAlwaysUseHTTPS struct {
	// ID of the zone setting.
	ID ZonesSchemasAlwaysUseHTTPSID `json:"id"`
	// Value of the zone setting.
	Value ZonesSchemasAlwaysUseHTTPSValue `json:"value"`
	JSON  zonesSchemasAlwaysUseHTTPSJSON  `json:"-"`
	ZonesBase
}

// zonesSchemasAlwaysUseHTTPSJSON contains the JSON metadata for the struct
// [ZonesSchemasAlwaysUseHTTPS]
type zonesSchemasAlwaysUseHTTPSJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonesSchemasAlwaysUseHTTPS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zonesSchemasAlwaysUseHTTPSJSON) RawJSON() string {
	return r.raw
}

func (r ZonesSchemasAlwaysUseHTTPS) implementsZonesSetting() {}

func (r ZonesSchemasAlwaysUseHTTPS) implementsZonesZoneSettingsResponseCollectionResult() {}

// ID of the zone setting.
type ZonesSchemasAlwaysUseHTTPSID string

const (
	ZonesSchemasAlwaysUseHTTPSIDAlwaysUseHTTPS ZonesSchemasAlwaysUseHTTPSID = "always_use_https"
)

func (r ZonesSchemasAlwaysUseHTTPSID) IsKnown() bool {
	switch r {
	case ZonesSchemasAlwaysUseHTTPSIDAlwaysUseHTTPS:
		return true
	}
	return false
}

// Value of the zone setting.
type ZonesSchemasAlwaysUseHTTPSValue string

const (
	ZonesSchemasAlwaysUseHTTPSValueOn  ZonesSchemasAlwaysUseHTTPSValue = "on"
	ZonesSchemasAlwaysUseHTTPSValueOff ZonesSchemasAlwaysUseHTTPSValue = "off"
)

func (r ZonesSchemasAlwaysUseHTTPSValue) IsKnown() bool {
	switch r {
	case ZonesSchemasAlwaysUseHTTPSValueOn, ZonesSchemasAlwaysUseHTTPSValueOff:
		return true
	}
	return false
}

// Reply to all requests for URLs that use "http" with a 301 redirect to the
// equivalent "https" URL. If you only want to redirect for a subset of requests,
// consider creating an "Always use HTTPS" page rule.
type ZonesSchemasAlwaysUseHTTPSParam struct {
	// ID of the zone setting.
	ID param.Field[ZonesSchemasAlwaysUseHTTPSID] `json:"id"`
	// Value of the zone setting.
	Value param.Field[ZonesSchemasAlwaysUseHTTPSValue] `json:"value"`
	ZonesBaseParam
}

func (r ZonesSchemasAlwaysUseHTTPSParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZonesSchemasAlwaysUseHTTPSParam) implementsZonesSettingUnionParam() {}

func (r ZonesSchemasAlwaysUseHTTPSParam) implementsZoneSettingUpdateParamsBodyUnion() {}

// Enable the Automatic HTTPS Rewrites feature for this zone.
type ZonesSchemasAutomaticHTTPSRewrites struct {
	// ID of the zone setting.
	ID ZonesSchemasAutomaticHTTPSRewritesID `json:"id"`
	// Value of the zone setting. Notes: Default value depends on the zone's plan
	// level.
	Value ZonesSchemasAutomaticHTTPSRewritesValue `json:"value"`
	JSON  zonesSchemasAutomaticHTTPSRewritesJSON  `json:"-"`
	ZonesBase
}

// zonesSchemasAutomaticHTTPSRewritesJSON contains the JSON metadata for the struct
// [ZonesSchemasAutomaticHTTPSRewrites]
type zonesSchemasAutomaticHTTPSRewritesJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonesSchemasAutomaticHTTPSRewrites) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zonesSchemasAutomaticHTTPSRewritesJSON) RawJSON() string {
	return r.raw
}

func (r ZonesSchemasAutomaticHTTPSRewrites) implementsZonesSetting() {}

func (r ZonesSchemasAutomaticHTTPSRewrites) implementsZonesZoneSettingsResponseCollectionResult() {}

// ID of the zone setting.
type ZonesSchemasAutomaticHTTPSRewritesID string

const (
	ZonesSchemasAutomaticHTTPSRewritesIDAutomaticHTTPSRewrites ZonesSchemasAutomaticHTTPSRewritesID = "automatic_https_rewrites"
)

func (r ZonesSchemasAutomaticHTTPSRewritesID) IsKnown() bool {
	switch r {
	case ZonesSchemasAutomaticHTTPSRewritesIDAutomaticHTTPSRewrites:
		return true
	}
	return false
}

// Value of the zone setting. Notes: Default value depends on the zone's plan
// level.
type ZonesSchemasAutomaticHTTPSRewritesValue string

const (
	ZonesSchemasAutomaticHTTPSRewritesValueOn  ZonesSchemasAutomaticHTTPSRewritesValue = "on"
	ZonesSchemasAutomaticHTTPSRewritesValueOff ZonesSchemasAutomaticHTTPSRewritesValue = "off"
)

func (r ZonesSchemasAutomaticHTTPSRewritesValue) IsKnown() bool {
	switch r {
	case ZonesSchemasAutomaticHTTPSRewritesValueOn, ZonesSchemasAutomaticHTTPSRewritesValueOff:
		return true
	}
	return false
}

// Enable the Automatic HTTPS Rewrites feature for this zone.
type ZonesSchemasAutomaticHTTPSRewritesParam struct {
	// ID of the zone setting.
	ID param.Field[ZonesSchemasAutomaticHTTPSRewritesID] `json:"id"`
	// Value of the zone setting. Notes: Default value depends on the zone's plan
	// level.
	Value param.Field[ZonesSchemasAutomaticHTTPSRewritesValue] `json:"value"`
	ZonesBaseParam
}

func (r ZonesSchemasAutomaticHTTPSRewritesParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZonesSchemasAutomaticHTTPSRewritesParam) implementsZonesSettingUnionParam() {}

func (r ZonesSchemasAutomaticHTTPSRewritesParam) implementsZoneSettingUpdateParamsBodyUnion() {}

// [Automatic Platform Optimization for WordPress](https://developers.cloudflare.com/automatic-platform-optimization/)
// serves your WordPress site from Cloudflare's edge network and caches third-party
// fonts.
type ZonesSchemasAutomaticPlatformOptimization struct {
	// ID of the zone setting.
	ID    ZonesSchemasAutomaticPlatformOptimizationID    `json:"id"`
	Value ZonesSchemasAutomaticPlatformOptimizationValue `json:"value"`
	JSON  zonesSchemasAutomaticPlatformOptimizationJSON  `json:"-"`
	ZonesBase
}

// zonesSchemasAutomaticPlatformOptimizationJSON contains the JSON metadata for the
// struct [ZonesSchemasAutomaticPlatformOptimization]
type zonesSchemasAutomaticPlatformOptimizationJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonesSchemasAutomaticPlatformOptimization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zonesSchemasAutomaticPlatformOptimizationJSON) RawJSON() string {
	return r.raw
}

func (r ZonesSchemasAutomaticPlatformOptimization) implementsZonesSetting() {}

func (r ZonesSchemasAutomaticPlatformOptimization) implementsZonesZoneSettingsResponseCollectionResult() {
}

// ID of the zone setting.
type ZonesSchemasAutomaticPlatformOptimizationID string

const (
	ZonesSchemasAutomaticPlatformOptimizationIDAutomaticPlatformOptimization ZonesSchemasAutomaticPlatformOptimizationID = "automatic_platform_optimization"
)

func (r ZonesSchemasAutomaticPlatformOptimizationID) IsKnown() bool {
	switch r {
	case ZonesSchemasAutomaticPlatformOptimizationIDAutomaticPlatformOptimization:
		return true
	}
	return false
}

type ZonesSchemasAutomaticPlatformOptimizationValue struct {
	// Indicates whether or not
	// [cache by device type](https://developers.cloudflare.com/automatic-platform-optimization/reference/cache-device-type/)
	// is enabled.
	CacheByDeviceType bool `json:"cache_by_device_type,required"`
	// Indicates whether or not Cloudflare proxy is enabled.
	Cf bool `json:"cf,required"`
	// Indicates whether or not Automatic Platform Optimization is enabled.
	Enabled bool `json:"enabled,required"`
	// An array of hostnames where Automatic Platform Optimization for WordPress is
	// activated.
	Hostnames []string `json:"hostnames,required" format:"hostname"`
	// Indicates whether or not site is powered by WordPress.
	Wordpress bool `json:"wordpress,required"`
	// Indicates whether or not
	// [Cloudflare for WordPress plugin](https://wordpress.org/plugins/cloudflare/) is
	// installed.
	WpPlugin bool                                               `json:"wp_plugin,required"`
	JSON     zonesSchemasAutomaticPlatformOptimizationValueJSON `json:"-"`
}

// zonesSchemasAutomaticPlatformOptimizationValueJSON contains the JSON metadata
// for the struct [ZonesSchemasAutomaticPlatformOptimizationValue]
type zonesSchemasAutomaticPlatformOptimizationValueJSON struct {
	CacheByDeviceType apijson.Field
	Cf                apijson.Field
	Enabled           apijson.Field
	Hostnames         apijson.Field
	Wordpress         apijson.Field
	WpPlugin          apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ZonesSchemasAutomaticPlatformOptimizationValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zonesSchemasAutomaticPlatformOptimizationValueJSON) RawJSON() string {
	return r.raw
}

// [Automatic Platform Optimization for WordPress](https://developers.cloudflare.com/automatic-platform-optimization/)
// serves your WordPress site from Cloudflare's edge network and caches third-party
// fonts.
type ZonesSchemasAutomaticPlatformOptimizationParam struct {
	// ID of the zone setting.
	ID    param.Field[ZonesSchemasAutomaticPlatformOptimizationID]         `json:"id"`
	Value param.Field[ZonesSchemasAutomaticPlatformOptimizationValueParam] `json:"value"`
	ZonesBaseParam
}

func (r ZonesSchemasAutomaticPlatformOptimizationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZonesSchemasAutomaticPlatformOptimizationParam) implementsZonesSettingUnionParam() {}

func (r ZonesSchemasAutomaticPlatformOptimizationParam) implementsZoneSettingUpdateParamsBodyUnion() {
}

type ZonesSchemasAutomaticPlatformOptimizationValueParam struct {
	// Indicates whether or not
	// [cache by device type](https://developers.cloudflare.com/automatic-platform-optimization/reference/cache-device-type/)
	// is enabled.
	CacheByDeviceType param.Field[bool] `json:"cache_by_device_type,required"`
	// Indicates whether or not Cloudflare proxy is enabled.
	Cf param.Field[bool] `json:"cf,required"`
	// Indicates whether or not Automatic Platform Optimization is enabled.
	Enabled param.Field[bool] `json:"enabled,required"`
	// An array of hostnames where Automatic Platform Optimization for WordPress is
	// activated.
	Hostnames param.Field[[]string] `json:"hostnames,required" format:"hostname"`
	// Indicates whether or not site is powered by WordPress.
	Wordpress param.Field[bool] `json:"wordpress,required"`
	// Indicates whether or not
	// [Cloudflare for WordPress plugin](https://wordpress.org/plugins/cloudflare/) is
	// installed.
	WpPlugin param.Field[bool] `json:"wp_plugin,required"`
}

func (r ZonesSchemasAutomaticPlatformOptimizationValueParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Browser Cache TTL (in seconds) specifies how long Cloudflare-cached resources
// will remain on your visitors' computers. Cloudflare will honor any larger times
// specified by your server.
// (https://support.cloudflare.com/hc/en-us/articles/200168276).
type ZonesSchemasBrowserCacheTtl struct {
	// ID of the zone setting.
	ID ZonesSchemasBrowserCacheTtlID `json:"id"`
	// Value of the zone setting. Notes: Setting a TTL of 0 is equivalent to selecting
	// `Respect Existing Headers`
	Value ZonesSchemasBrowserCacheTtlValue `json:"value"`
	JSON  zonesSchemasBrowserCacheTtlJSON  `json:"-"`
	ZonesBase
}

// zonesSchemasBrowserCacheTtlJSON contains the JSON metadata for the struct
// [ZonesSchemasBrowserCacheTtl]
type zonesSchemasBrowserCacheTtlJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonesSchemasBrowserCacheTtl) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zonesSchemasBrowserCacheTtlJSON) RawJSON() string {
	return r.raw
}

func (r ZonesSchemasBrowserCacheTtl) implementsZonesSetting() {}

func (r ZonesSchemasBrowserCacheTtl) implementsZonesZoneSettingsResponseCollectionResult() {}

// ID of the zone setting.
type ZonesSchemasBrowserCacheTtlID string

const (
	ZonesSchemasBrowserCacheTtlIDBrowserCacheTtl ZonesSchemasBrowserCacheTtlID = "browser_cache_ttl"
)

func (r ZonesSchemasBrowserCacheTtlID) IsKnown() bool {
	switch r {
	case ZonesSchemasBrowserCacheTtlIDBrowserCacheTtl:
		return true
	}
	return false
}

// Value of the zone setting. Notes: Setting a TTL of 0 is equivalent to selecting
// `Respect Existing Headers`
type ZonesSchemasBrowserCacheTtlValue float64

const (
	ZonesSchemasBrowserCacheTtlValue0        ZonesSchemasBrowserCacheTtlValue = 0
	ZonesSchemasBrowserCacheTtlValue30       ZonesSchemasBrowserCacheTtlValue = 30
	ZonesSchemasBrowserCacheTtlValue60       ZonesSchemasBrowserCacheTtlValue = 60
	ZonesSchemasBrowserCacheTtlValue120      ZonesSchemasBrowserCacheTtlValue = 120
	ZonesSchemasBrowserCacheTtlValue300      ZonesSchemasBrowserCacheTtlValue = 300
	ZonesSchemasBrowserCacheTtlValue1200     ZonesSchemasBrowserCacheTtlValue = 1200
	ZonesSchemasBrowserCacheTtlValue1800     ZonesSchemasBrowserCacheTtlValue = 1800
	ZonesSchemasBrowserCacheTtlValue3600     ZonesSchemasBrowserCacheTtlValue = 3600
	ZonesSchemasBrowserCacheTtlValue7200     ZonesSchemasBrowserCacheTtlValue = 7200
	ZonesSchemasBrowserCacheTtlValue10800    ZonesSchemasBrowserCacheTtlValue = 10800
	ZonesSchemasBrowserCacheTtlValue14400    ZonesSchemasBrowserCacheTtlValue = 14400
	ZonesSchemasBrowserCacheTtlValue18000    ZonesSchemasBrowserCacheTtlValue = 18000
	ZonesSchemasBrowserCacheTtlValue28800    ZonesSchemasBrowserCacheTtlValue = 28800
	ZonesSchemasBrowserCacheTtlValue43200    ZonesSchemasBrowserCacheTtlValue = 43200
	ZonesSchemasBrowserCacheTtlValue57600    ZonesSchemasBrowserCacheTtlValue = 57600
	ZonesSchemasBrowserCacheTtlValue72000    ZonesSchemasBrowserCacheTtlValue = 72000
	ZonesSchemasBrowserCacheTtlValue86400    ZonesSchemasBrowserCacheTtlValue = 86400
	ZonesSchemasBrowserCacheTtlValue172800   ZonesSchemasBrowserCacheTtlValue = 172800
	ZonesSchemasBrowserCacheTtlValue259200   ZonesSchemasBrowserCacheTtlValue = 259200
	ZonesSchemasBrowserCacheTtlValue345600   ZonesSchemasBrowserCacheTtlValue = 345600
	ZonesSchemasBrowserCacheTtlValue432000   ZonesSchemasBrowserCacheTtlValue = 432000
	ZonesSchemasBrowserCacheTtlValue691200   ZonesSchemasBrowserCacheTtlValue = 691200
	ZonesSchemasBrowserCacheTtlValue1382400  ZonesSchemasBrowserCacheTtlValue = 1382400
	ZonesSchemasBrowserCacheTtlValue2073600  ZonesSchemasBrowserCacheTtlValue = 2073600
	ZonesSchemasBrowserCacheTtlValue2678400  ZonesSchemasBrowserCacheTtlValue = 2678400
	ZonesSchemasBrowserCacheTtlValue5356800  ZonesSchemasBrowserCacheTtlValue = 5356800
	ZonesSchemasBrowserCacheTtlValue16070400 ZonesSchemasBrowserCacheTtlValue = 16070400
	ZonesSchemasBrowserCacheTtlValue31536000 ZonesSchemasBrowserCacheTtlValue = 31536000
)

func (r ZonesSchemasBrowserCacheTtlValue) IsKnown() bool {
	switch r {
	case ZonesSchemasBrowserCacheTtlValue0, ZonesSchemasBrowserCacheTtlValue30, ZonesSchemasBrowserCacheTtlValue60, ZonesSchemasBrowserCacheTtlValue120, ZonesSchemasBrowserCacheTtlValue300, ZonesSchemasBrowserCacheTtlValue1200, ZonesSchemasBrowserCacheTtlValue1800, ZonesSchemasBrowserCacheTtlValue3600, ZonesSchemasBrowserCacheTtlValue7200, ZonesSchemasBrowserCacheTtlValue10800, ZonesSchemasBrowserCacheTtlValue14400, ZonesSchemasBrowserCacheTtlValue18000, ZonesSchemasBrowserCacheTtlValue28800, ZonesSchemasBrowserCacheTtlValue43200, ZonesSchemasBrowserCacheTtlValue57600, ZonesSchemasBrowserCacheTtlValue72000, ZonesSchemasBrowserCacheTtlValue86400, ZonesSchemasBrowserCacheTtlValue172800, ZonesSchemasBrowserCacheTtlValue259200, ZonesSchemasBrowserCacheTtlValue345600, ZonesSchemasBrowserCacheTtlValue432000, ZonesSchemasBrowserCacheTtlValue691200, ZonesSchemasBrowserCacheTtlValue1382400, ZonesSchemasBrowserCacheTtlValue2073600, ZonesSchemasBrowserCacheTtlValue2678400, ZonesSchemasBrowserCacheTtlValue5356800, ZonesSchemasBrowserCacheTtlValue16070400, ZonesSchemasBrowserCacheTtlValue31536000:
		return true
	}
	return false
}

// Browser Cache TTL (in seconds) specifies how long Cloudflare-cached resources
// will remain on your visitors' computers. Cloudflare will honor any larger times
// specified by your server.
// (https://support.cloudflare.com/hc/en-us/articles/200168276).
type ZonesSchemasBrowserCacheTtlParam struct {
	// ID of the zone setting.
	ID param.Field[ZonesSchemasBrowserCacheTtlID] `json:"id"`
	// Value of the zone setting. Notes: Setting a TTL of 0 is equivalent to selecting
	// `Respect Existing Headers`
	Value param.Field[ZonesSchemasBrowserCacheTtlValue] `json:"value"`
	ZonesBaseParam
}

func (r ZonesSchemasBrowserCacheTtlParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZonesSchemasBrowserCacheTtlParam) implementsZonesSettingUnionParam() {}

func (r ZonesSchemasBrowserCacheTtlParam) implementsZoneSettingUpdateParamsBodyUnion() {}

// Browser Integrity Check is similar to Bad Behavior and looks for common HTTP
// headers abused most commonly by spammers and denies access to your page. It will
// also challenge visitors that do not have a user agent or a non standard user
// agent (also commonly used by abuse bots, crawlers or visitors).
// (https://support.cloudflare.com/hc/en-us/articles/200170086).
type ZonesSchemasBrowserCheck struct {
	// ID of the zone setting.
	ID ZonesSchemasBrowserCheckID `json:"id"`
	// Value of the zone setting.
	Value ZonesSchemasBrowserCheckValue `json:"value"`
	JSON  zonesSchemasBrowserCheckJSON  `json:"-"`
	ZonesBase
}

// zonesSchemasBrowserCheckJSON contains the JSON metadata for the struct
// [ZonesSchemasBrowserCheck]
type zonesSchemasBrowserCheckJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonesSchemasBrowserCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zonesSchemasBrowserCheckJSON) RawJSON() string {
	return r.raw
}

func (r ZonesSchemasBrowserCheck) implementsZonesSetting() {}

func (r ZonesSchemasBrowserCheck) implementsZonesZoneSettingsResponseCollectionResult() {}

// ID of the zone setting.
type ZonesSchemasBrowserCheckID string

const (
	ZonesSchemasBrowserCheckIDBrowserCheck ZonesSchemasBrowserCheckID = "browser_check"
)

func (r ZonesSchemasBrowserCheckID) IsKnown() bool {
	switch r {
	case ZonesSchemasBrowserCheckIDBrowserCheck:
		return true
	}
	return false
}

// Value of the zone setting.
type ZonesSchemasBrowserCheckValue string

const (
	ZonesSchemasBrowserCheckValueOn  ZonesSchemasBrowserCheckValue = "on"
	ZonesSchemasBrowserCheckValueOff ZonesSchemasBrowserCheckValue = "off"
)

func (r ZonesSchemasBrowserCheckValue) IsKnown() bool {
	switch r {
	case ZonesSchemasBrowserCheckValueOn, ZonesSchemasBrowserCheckValueOff:
		return true
	}
	return false
}

// Browser Integrity Check is similar to Bad Behavior and looks for common HTTP
// headers abused most commonly by spammers and denies access to your page. It will
// also challenge visitors that do not have a user agent or a non standard user
// agent (also commonly used by abuse bots, crawlers or visitors).
// (https://support.cloudflare.com/hc/en-us/articles/200170086).
type ZonesSchemasBrowserCheckParam struct {
	// ID of the zone setting.
	ID param.Field[ZonesSchemasBrowserCheckID] `json:"id"`
	// Value of the zone setting.
	Value param.Field[ZonesSchemasBrowserCheckValue] `json:"value"`
	ZonesBaseParam
}

func (r ZonesSchemasBrowserCheckParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZonesSchemasBrowserCheckParam) implementsZonesSettingUnionParam() {}

func (r ZonesSchemasBrowserCheckParam) implementsZoneSettingUpdateParamsBodyUnion() {}

// Cache Level functions based off the setting level. The basic setting will cache
// most static resources (i.e., css, images, and JavaScript). The simplified
// setting will ignore the query string when delivering a cached resource. The
// aggressive setting will cache all static resources, including ones with a query
// string. (https://support.cloudflare.com/hc/en-us/articles/200168256).
type ZonesSchemasCacheLevel struct {
	// ID of the zone setting.
	ID ZonesSchemasCacheLevelID `json:"id"`
	// Value of the zone setting.
	Value ZonesSchemasCacheLevelValue `json:"value"`
	JSON  zonesSchemasCacheLevelJSON  `json:"-"`
	ZonesBase
}

// zonesSchemasCacheLevelJSON contains the JSON metadata for the struct
// [ZonesSchemasCacheLevel]
type zonesSchemasCacheLevelJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonesSchemasCacheLevel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zonesSchemasCacheLevelJSON) RawJSON() string {
	return r.raw
}

func (r ZonesSchemasCacheLevel) implementsZonesSetting() {}

func (r ZonesSchemasCacheLevel) implementsZonesZoneSettingsResponseCollectionResult() {}

// ID of the zone setting.
type ZonesSchemasCacheLevelID string

const (
	ZonesSchemasCacheLevelIDCacheLevel ZonesSchemasCacheLevelID = "cache_level"
)

func (r ZonesSchemasCacheLevelID) IsKnown() bool {
	switch r {
	case ZonesSchemasCacheLevelIDCacheLevel:
		return true
	}
	return false
}

// Value of the zone setting.
type ZonesSchemasCacheLevelValue string

const (
	ZonesSchemasCacheLevelValueAggressive ZonesSchemasCacheLevelValue = "aggressive"
	ZonesSchemasCacheLevelValueBasic      ZonesSchemasCacheLevelValue = "basic"
	ZonesSchemasCacheLevelValueSimplified ZonesSchemasCacheLevelValue = "simplified"
)

func (r ZonesSchemasCacheLevelValue) IsKnown() bool {
	switch r {
	case ZonesSchemasCacheLevelValueAggressive, ZonesSchemasCacheLevelValueBasic, ZonesSchemasCacheLevelValueSimplified:
		return true
	}
	return false
}

// Cache Level functions based off the setting level. The basic setting will cache
// most static resources (i.e., css, images, and JavaScript). The simplified
// setting will ignore the query string when delivering a cached resource. The
// aggressive setting will cache all static resources, including ones with a query
// string. (https://support.cloudflare.com/hc/en-us/articles/200168256).
type ZonesSchemasCacheLevelParam struct {
	// ID of the zone setting.
	ID param.Field[ZonesSchemasCacheLevelID] `json:"id"`
	// Value of the zone setting.
	Value param.Field[ZonesSchemasCacheLevelValue] `json:"value"`
	ZonesBaseParam
}

func (r ZonesSchemasCacheLevelParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZonesSchemasCacheLevelParam) implementsZonesSettingUnionParam() {}

func (r ZonesSchemasCacheLevelParam) implementsZoneSettingUpdateParamsBodyUnion() {}

// Time (in seconds) that a resource will be ensured to remain on Cloudflare's
// cache servers.
type ZonesSchemasEdgeCacheTtl struct {
	// ID of the zone setting.
	ID ZonesSchemasEdgeCacheTtlID `json:"id"`
	// Value of the zone setting. Notes: The minimum TTL available depends on the plan
	// level of the zone. (Enterprise = 30, Business = 1800, Pro = 3600, Free = 7200)
	Value ZonesSchemasEdgeCacheTtlValue `json:"value"`
	JSON  zonesSchemasEdgeCacheTtlJSON  `json:"-"`
	ZonesBase
}

// zonesSchemasEdgeCacheTtlJSON contains the JSON metadata for the struct
// [ZonesSchemasEdgeCacheTtl]
type zonesSchemasEdgeCacheTtlJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonesSchemasEdgeCacheTtl) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zonesSchemasEdgeCacheTtlJSON) RawJSON() string {
	return r.raw
}

func (r ZonesSchemasEdgeCacheTtl) implementsZonesSetting() {}

func (r ZonesSchemasEdgeCacheTtl) implementsZonesZoneSettingsResponseCollectionResult() {}

// ID of the zone setting.
type ZonesSchemasEdgeCacheTtlID string

const (
	ZonesSchemasEdgeCacheTtlIDEdgeCacheTtl ZonesSchemasEdgeCacheTtlID = "edge_cache_ttl"
)

func (r ZonesSchemasEdgeCacheTtlID) IsKnown() bool {
	switch r {
	case ZonesSchemasEdgeCacheTtlIDEdgeCacheTtl:
		return true
	}
	return false
}

// Value of the zone setting. Notes: The minimum TTL available depends on the plan
// level of the zone. (Enterprise = 30, Business = 1800, Pro = 3600, Free = 7200)
type ZonesSchemasEdgeCacheTtlValue float64

const (
	ZonesSchemasEdgeCacheTtlValue30     ZonesSchemasEdgeCacheTtlValue = 30
	ZonesSchemasEdgeCacheTtlValue60     ZonesSchemasEdgeCacheTtlValue = 60
	ZonesSchemasEdgeCacheTtlValue300    ZonesSchemasEdgeCacheTtlValue = 300
	ZonesSchemasEdgeCacheTtlValue1200   ZonesSchemasEdgeCacheTtlValue = 1200
	ZonesSchemasEdgeCacheTtlValue1800   ZonesSchemasEdgeCacheTtlValue = 1800
	ZonesSchemasEdgeCacheTtlValue3600   ZonesSchemasEdgeCacheTtlValue = 3600
	ZonesSchemasEdgeCacheTtlValue7200   ZonesSchemasEdgeCacheTtlValue = 7200
	ZonesSchemasEdgeCacheTtlValue10800  ZonesSchemasEdgeCacheTtlValue = 10800
	ZonesSchemasEdgeCacheTtlValue14400  ZonesSchemasEdgeCacheTtlValue = 14400
	ZonesSchemasEdgeCacheTtlValue18000  ZonesSchemasEdgeCacheTtlValue = 18000
	ZonesSchemasEdgeCacheTtlValue28800  ZonesSchemasEdgeCacheTtlValue = 28800
	ZonesSchemasEdgeCacheTtlValue43200  ZonesSchemasEdgeCacheTtlValue = 43200
	ZonesSchemasEdgeCacheTtlValue57600  ZonesSchemasEdgeCacheTtlValue = 57600
	ZonesSchemasEdgeCacheTtlValue72000  ZonesSchemasEdgeCacheTtlValue = 72000
	ZonesSchemasEdgeCacheTtlValue86400  ZonesSchemasEdgeCacheTtlValue = 86400
	ZonesSchemasEdgeCacheTtlValue172800 ZonesSchemasEdgeCacheTtlValue = 172800
	ZonesSchemasEdgeCacheTtlValue259200 ZonesSchemasEdgeCacheTtlValue = 259200
	ZonesSchemasEdgeCacheTtlValue345600 ZonesSchemasEdgeCacheTtlValue = 345600
	ZonesSchemasEdgeCacheTtlValue432000 ZonesSchemasEdgeCacheTtlValue = 432000
	ZonesSchemasEdgeCacheTtlValue518400 ZonesSchemasEdgeCacheTtlValue = 518400
	ZonesSchemasEdgeCacheTtlValue604800 ZonesSchemasEdgeCacheTtlValue = 604800
)

func (r ZonesSchemasEdgeCacheTtlValue) IsKnown() bool {
	switch r {
	case ZonesSchemasEdgeCacheTtlValue30, ZonesSchemasEdgeCacheTtlValue60, ZonesSchemasEdgeCacheTtlValue300, ZonesSchemasEdgeCacheTtlValue1200, ZonesSchemasEdgeCacheTtlValue1800, ZonesSchemasEdgeCacheTtlValue3600, ZonesSchemasEdgeCacheTtlValue7200, ZonesSchemasEdgeCacheTtlValue10800, ZonesSchemasEdgeCacheTtlValue14400, ZonesSchemasEdgeCacheTtlValue18000, ZonesSchemasEdgeCacheTtlValue28800, ZonesSchemasEdgeCacheTtlValue43200, ZonesSchemasEdgeCacheTtlValue57600, ZonesSchemasEdgeCacheTtlValue72000, ZonesSchemasEdgeCacheTtlValue86400, ZonesSchemasEdgeCacheTtlValue172800, ZonesSchemasEdgeCacheTtlValue259200, ZonesSchemasEdgeCacheTtlValue345600, ZonesSchemasEdgeCacheTtlValue432000, ZonesSchemasEdgeCacheTtlValue518400, ZonesSchemasEdgeCacheTtlValue604800:
		return true
	}
	return false
}

// Time (in seconds) that a resource will be ensured to remain on Cloudflare's
// cache servers.
type ZonesSchemasEdgeCacheTtlParam struct {
	// ID of the zone setting.
	ID param.Field[ZonesSchemasEdgeCacheTtlID] `json:"id"`
	// Value of the zone setting. Notes: The minimum TTL available depends on the plan
	// level of the zone. (Enterprise = 30, Business = 1800, Pro = 3600, Free = 7200)
	Value param.Field[ZonesSchemasEdgeCacheTtlValue] `json:"value"`
	ZonesBaseParam
}

func (r ZonesSchemasEdgeCacheTtlParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZonesSchemasEdgeCacheTtlParam) implementsZonesSettingUnionParam() {}

func (r ZonesSchemasEdgeCacheTtlParam) implementsZoneSettingUpdateParamsBodyUnion() {}

// Encrypt email adresses on your web page from bots, while keeping them visible to
// humans. (https://support.cloudflare.com/hc/en-us/articles/200170016).
type ZonesSchemasEmailObfuscation struct {
	// ID of the zone setting.
	ID ZonesSchemasEmailObfuscationID `json:"id"`
	// Value of the zone setting.
	Value ZonesSchemasEmailObfuscationValue `json:"value"`
	JSON  zonesSchemasEmailObfuscationJSON  `json:"-"`
	ZonesBase
}

// zonesSchemasEmailObfuscationJSON contains the JSON metadata for the struct
// [ZonesSchemasEmailObfuscation]
type zonesSchemasEmailObfuscationJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonesSchemasEmailObfuscation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zonesSchemasEmailObfuscationJSON) RawJSON() string {
	return r.raw
}

func (r ZonesSchemasEmailObfuscation) implementsZonesSetting() {}

func (r ZonesSchemasEmailObfuscation) implementsZonesZoneSettingsResponseCollectionResult() {}

// ID of the zone setting.
type ZonesSchemasEmailObfuscationID string

const (
	ZonesSchemasEmailObfuscationIDEmailObfuscation ZonesSchemasEmailObfuscationID = "email_obfuscation"
)

func (r ZonesSchemasEmailObfuscationID) IsKnown() bool {
	switch r {
	case ZonesSchemasEmailObfuscationIDEmailObfuscation:
		return true
	}
	return false
}

// Value of the zone setting.
type ZonesSchemasEmailObfuscationValue string

const (
	ZonesSchemasEmailObfuscationValueOn  ZonesSchemasEmailObfuscationValue = "on"
	ZonesSchemasEmailObfuscationValueOff ZonesSchemasEmailObfuscationValue = "off"
)

func (r ZonesSchemasEmailObfuscationValue) IsKnown() bool {
	switch r {
	case ZonesSchemasEmailObfuscationValueOn, ZonesSchemasEmailObfuscationValueOff:
		return true
	}
	return false
}

// Encrypt email adresses on your web page from bots, while keeping them visible to
// humans. (https://support.cloudflare.com/hc/en-us/articles/200170016).
type ZonesSchemasEmailObfuscationParam struct {
	// ID of the zone setting.
	ID param.Field[ZonesSchemasEmailObfuscationID] `json:"id"`
	// Value of the zone setting.
	Value param.Field[ZonesSchemasEmailObfuscationValue] `json:"value"`
	ZonesBaseParam
}

func (r ZonesSchemasEmailObfuscationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZonesSchemasEmailObfuscationParam) implementsZonesSettingUnionParam() {}

func (r ZonesSchemasEmailObfuscationParam) implementsZoneSettingUpdateParamsBodyUnion() {}

// Enable IP Geolocation to have Cloudflare geolocate visitors to your website and
// pass the country code to you.
// (https://support.cloudflare.com/hc/en-us/articles/200168236).
type ZonesSchemasIPGeolocation struct {
	// ID of the zone setting.
	ID ZonesSchemasIPGeolocationID `json:"id"`
	// Value of the zone setting.
	Value ZonesSchemasIPGeolocationValue `json:"value"`
	JSON  zonesSchemasIPGeolocationJSON  `json:"-"`
	ZonesBase
}

// zonesSchemasIPGeolocationJSON contains the JSON metadata for the struct
// [ZonesSchemasIPGeolocation]
type zonesSchemasIPGeolocationJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonesSchemasIPGeolocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zonesSchemasIPGeolocationJSON) RawJSON() string {
	return r.raw
}

func (r ZonesSchemasIPGeolocation) implementsZonesSetting() {}

func (r ZonesSchemasIPGeolocation) implementsZonesZoneSettingsResponseCollectionResult() {}

// ID of the zone setting.
type ZonesSchemasIPGeolocationID string

const (
	ZonesSchemasIPGeolocationIDIPGeolocation ZonesSchemasIPGeolocationID = "ip_geolocation"
)

func (r ZonesSchemasIPGeolocationID) IsKnown() bool {
	switch r {
	case ZonesSchemasIPGeolocationIDIPGeolocation:
		return true
	}
	return false
}

// Value of the zone setting.
type ZonesSchemasIPGeolocationValue string

const (
	ZonesSchemasIPGeolocationValueOn  ZonesSchemasIPGeolocationValue = "on"
	ZonesSchemasIPGeolocationValueOff ZonesSchemasIPGeolocationValue = "off"
)

func (r ZonesSchemasIPGeolocationValue) IsKnown() bool {
	switch r {
	case ZonesSchemasIPGeolocationValueOn, ZonesSchemasIPGeolocationValueOff:
		return true
	}
	return false
}

// Enable IP Geolocation to have Cloudflare geolocate visitors to your website and
// pass the country code to you.
// (https://support.cloudflare.com/hc/en-us/articles/200168236).
type ZonesSchemasIPGeolocationParam struct {
	// ID of the zone setting.
	ID param.Field[ZonesSchemasIPGeolocationID] `json:"id"`
	// Value of the zone setting.
	Value param.Field[ZonesSchemasIPGeolocationValue] `json:"value"`
	ZonesBaseParam
}

func (r ZonesSchemasIPGeolocationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZonesSchemasIPGeolocationParam) implementsZonesSettingUnionParam() {}

func (r ZonesSchemasIPGeolocationParam) implementsZoneSettingUpdateParamsBodyUnion() {}

// Automatically optimize image loading for website visitors on mobile devices.
// Refer to
// [our blog post](http://blog.cloudflare.com/mirage2-solving-mobile-speed) for
// more information.
type ZonesSchemasMirage struct {
	// ID of the zone setting.
	ID ZonesSchemasMirageID `json:"id"`
	// Value of the zone setting.
	Value ZonesSchemasMirageValue `json:"value"`
	JSON  zonesSchemasMirageJSON  `json:"-"`
	ZonesBase
}

// zonesSchemasMirageJSON contains the JSON metadata for the struct
// [ZonesSchemasMirage]
type zonesSchemasMirageJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonesSchemasMirage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zonesSchemasMirageJSON) RawJSON() string {
	return r.raw
}

func (r ZonesSchemasMirage) implementsZonesSetting() {}

func (r ZonesSchemasMirage) implementsZonesZoneSettingsResponseCollectionResult() {}

// ID of the zone setting.
type ZonesSchemasMirageID string

const (
	ZonesSchemasMirageIDMirage ZonesSchemasMirageID = "mirage"
)

func (r ZonesSchemasMirageID) IsKnown() bool {
	switch r {
	case ZonesSchemasMirageIDMirage:
		return true
	}
	return false
}

// Value of the zone setting.
type ZonesSchemasMirageValue string

const (
	ZonesSchemasMirageValueOn  ZonesSchemasMirageValue = "on"
	ZonesSchemasMirageValueOff ZonesSchemasMirageValue = "off"
)

func (r ZonesSchemasMirageValue) IsKnown() bool {
	switch r {
	case ZonesSchemasMirageValueOn, ZonesSchemasMirageValueOff:
		return true
	}
	return false
}

// Automatically optimize image loading for website visitors on mobile devices.
// Refer to
// [our blog post](http://blog.cloudflare.com/mirage2-solving-mobile-speed) for
// more information.
type ZonesSchemasMirageParam struct {
	// ID of the zone setting.
	ID param.Field[ZonesSchemasMirageID] `json:"id"`
	// Value of the zone setting.
	Value param.Field[ZonesSchemasMirageValue] `json:"value"`
	ZonesBaseParam
}

func (r ZonesSchemasMirageParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZonesSchemasMirageParam) implementsZonesSettingUnionParam() {}

func (r ZonesSchemasMirageParam) implementsZoneSettingUpdateParamsBodyUnion() {}

// Enables the Opportunistic Encryption feature for a zone.
type ZonesSchemasOpportunisticEncryption struct {
	// ID of the zone setting.
	ID ZonesSchemasOpportunisticEncryptionID `json:"id"`
	// Value of the zone setting. Notes: Default value depends on the zone's plan
	// level.
	Value ZonesSchemasOpportunisticEncryptionValue `json:"value"`
	JSON  zonesSchemasOpportunisticEncryptionJSON  `json:"-"`
	ZonesBase
}

// zonesSchemasOpportunisticEncryptionJSON contains the JSON metadata for the
// struct [ZonesSchemasOpportunisticEncryption]
type zonesSchemasOpportunisticEncryptionJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonesSchemasOpportunisticEncryption) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zonesSchemasOpportunisticEncryptionJSON) RawJSON() string {
	return r.raw
}

func (r ZonesSchemasOpportunisticEncryption) implementsZonesSetting() {}

func (r ZonesSchemasOpportunisticEncryption) implementsZonesZoneSettingsResponseCollectionResult() {}

// ID of the zone setting.
type ZonesSchemasOpportunisticEncryptionID string

const (
	ZonesSchemasOpportunisticEncryptionIDOpportunisticEncryption ZonesSchemasOpportunisticEncryptionID = "opportunistic_encryption"
)

func (r ZonesSchemasOpportunisticEncryptionID) IsKnown() bool {
	switch r {
	case ZonesSchemasOpportunisticEncryptionIDOpportunisticEncryption:
		return true
	}
	return false
}

// Value of the zone setting. Notes: Default value depends on the zone's plan
// level.
type ZonesSchemasOpportunisticEncryptionValue string

const (
	ZonesSchemasOpportunisticEncryptionValueOn  ZonesSchemasOpportunisticEncryptionValue = "on"
	ZonesSchemasOpportunisticEncryptionValueOff ZonesSchemasOpportunisticEncryptionValue = "off"
)

func (r ZonesSchemasOpportunisticEncryptionValue) IsKnown() bool {
	switch r {
	case ZonesSchemasOpportunisticEncryptionValueOn, ZonesSchemasOpportunisticEncryptionValueOff:
		return true
	}
	return false
}

// Enables the Opportunistic Encryption feature for a zone.
type ZonesSchemasOpportunisticEncryptionParam struct {
	// ID of the zone setting.
	ID param.Field[ZonesSchemasOpportunisticEncryptionID] `json:"id"`
	// Value of the zone setting. Notes: Default value depends on the zone's plan
	// level.
	Value param.Field[ZonesSchemasOpportunisticEncryptionValue] `json:"value"`
	ZonesBaseParam
}

func (r ZonesSchemasOpportunisticEncryptionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZonesSchemasOpportunisticEncryptionParam) implementsZonesSettingUnionParam() {}

func (r ZonesSchemasOpportunisticEncryptionParam) implementsZoneSettingUpdateParamsBodyUnion() {}

// Cloudflare will proxy customer error pages on any 502,504 errors on origin
// server instead of showing a default Cloudflare error page. This does not apply
// to 522 errors and is limited to Enterprise Zones.
type ZonesSchemasOriginErrorPagePassThru struct {
	// ID of the zone setting.
	ID ZonesSchemasOriginErrorPagePassThruID `json:"id"`
	// Value of the zone setting.
	Value ZonesSchemasOriginErrorPagePassThruValue `json:"value"`
	JSON  zonesSchemasOriginErrorPagePassThruJSON  `json:"-"`
	ZonesBase
}

// zonesSchemasOriginErrorPagePassThruJSON contains the JSON metadata for the
// struct [ZonesSchemasOriginErrorPagePassThru]
type zonesSchemasOriginErrorPagePassThruJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonesSchemasOriginErrorPagePassThru) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zonesSchemasOriginErrorPagePassThruJSON) RawJSON() string {
	return r.raw
}

func (r ZonesSchemasOriginErrorPagePassThru) implementsZonesSetting() {}

func (r ZonesSchemasOriginErrorPagePassThru) implementsZonesZoneSettingsResponseCollectionResult() {}

// ID of the zone setting.
type ZonesSchemasOriginErrorPagePassThruID string

const (
	ZonesSchemasOriginErrorPagePassThruIDOriginErrorPagePassThru ZonesSchemasOriginErrorPagePassThruID = "origin_error_page_pass_thru"
)

func (r ZonesSchemasOriginErrorPagePassThruID) IsKnown() bool {
	switch r {
	case ZonesSchemasOriginErrorPagePassThruIDOriginErrorPagePassThru:
		return true
	}
	return false
}

// Value of the zone setting.
type ZonesSchemasOriginErrorPagePassThruValue string

const (
	ZonesSchemasOriginErrorPagePassThruValueOn  ZonesSchemasOriginErrorPagePassThruValue = "on"
	ZonesSchemasOriginErrorPagePassThruValueOff ZonesSchemasOriginErrorPagePassThruValue = "off"
)

func (r ZonesSchemasOriginErrorPagePassThruValue) IsKnown() bool {
	switch r {
	case ZonesSchemasOriginErrorPagePassThruValueOn, ZonesSchemasOriginErrorPagePassThruValueOff:
		return true
	}
	return false
}

// Cloudflare will proxy customer error pages on any 502,504 errors on origin
// server instead of showing a default Cloudflare error page. This does not apply
// to 522 errors and is limited to Enterprise Zones.
type ZonesSchemasOriginErrorPagePassThruParam struct {
	// ID of the zone setting.
	ID param.Field[ZonesSchemasOriginErrorPagePassThruID] `json:"id"`
	// Value of the zone setting.
	Value param.Field[ZonesSchemasOriginErrorPagePassThruValue] `json:"value"`
	ZonesBaseParam
}

func (r ZonesSchemasOriginErrorPagePassThruParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZonesSchemasOriginErrorPagePassThruParam) implementsZonesSettingUnionParam() {}

func (r ZonesSchemasOriginErrorPagePassThruParam) implementsZoneSettingUpdateParamsBodyUnion() {}

// Removes metadata and compresses your images for faster page load times. Basic
// (Lossless): Reduce the size of PNG, JPEG, and GIF files - no impact on visual
// quality. Basic + JPEG (Lossy): Further reduce the size of JPEG files for faster
// image loading. Larger JPEGs are converted to progressive images, loading a
// lower-resolution image first and ending in a higher-resolution version. Not
// recommended for hi-res photography sites.
type ZonesSchemasPolish struct {
	// ID of the zone setting.
	ID ZonesSchemasPolishID `json:"id"`
	// Value of the zone setting.
	Value ZonesSchemasPolishValue `json:"value"`
	JSON  zonesSchemasPolishJSON  `json:"-"`
	ZonesBase
}

// zonesSchemasPolishJSON contains the JSON metadata for the struct
// [ZonesSchemasPolish]
type zonesSchemasPolishJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonesSchemasPolish) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zonesSchemasPolishJSON) RawJSON() string {
	return r.raw
}

func (r ZonesSchemasPolish) implementsZonesSetting() {}

func (r ZonesSchemasPolish) implementsZonesZoneSettingsResponseCollectionResult() {}

// ID of the zone setting.
type ZonesSchemasPolishID string

const (
	ZonesSchemasPolishIDPolish ZonesSchemasPolishID = "polish"
)

func (r ZonesSchemasPolishID) IsKnown() bool {
	switch r {
	case ZonesSchemasPolishIDPolish:
		return true
	}
	return false
}

// Value of the zone setting.
type ZonesSchemasPolishValue string

const (
	ZonesSchemasPolishValueOff      ZonesSchemasPolishValue = "off"
	ZonesSchemasPolishValueLossless ZonesSchemasPolishValue = "lossless"
	ZonesSchemasPolishValueLossy    ZonesSchemasPolishValue = "lossy"
)

func (r ZonesSchemasPolishValue) IsKnown() bool {
	switch r {
	case ZonesSchemasPolishValueOff, ZonesSchemasPolishValueLossless, ZonesSchemasPolishValueLossy:
		return true
	}
	return false
}

// Removes metadata and compresses your images for faster page load times. Basic
// (Lossless): Reduce the size of PNG, JPEG, and GIF files - no impact on visual
// quality. Basic + JPEG (Lossy): Further reduce the size of JPEG files for faster
// image loading. Larger JPEGs are converted to progressive images, loading a
// lower-resolution image first and ending in a higher-resolution version. Not
// recommended for hi-res photography sites.
type ZonesSchemasPolishParam struct {
	// ID of the zone setting.
	ID param.Field[ZonesSchemasPolishID] `json:"id"`
	// Value of the zone setting.
	Value param.Field[ZonesSchemasPolishValue] `json:"value"`
	ZonesBaseParam
}

func (r ZonesSchemasPolishParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZonesSchemasPolishParam) implementsZonesSettingUnionParam() {}

func (r ZonesSchemasPolishParam) implementsZoneSettingUpdateParamsBodyUnion() {}

// Enables or disables buffering of responses from the proxied server. Cloudflare
// may buffer the whole payload to deliver it at once to the client versus allowing
// it to be delivered in chunks. By default, the proxied server streams directly
// and is not buffered by Cloudflare. This is limited to Enterprise Zones.
type ZonesSchemasResponseBuffering struct {
	// ID of the zone setting.
	ID ZonesSchemasResponseBufferingID `json:"id"`
	// Value of the zone setting.
	Value ZonesSchemasResponseBufferingValue `json:"value"`
	JSON  zonesSchemasResponseBufferingJSON  `json:"-"`
	ZonesBase
}

// zonesSchemasResponseBufferingJSON contains the JSON metadata for the struct
// [ZonesSchemasResponseBuffering]
type zonesSchemasResponseBufferingJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonesSchemasResponseBuffering) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zonesSchemasResponseBufferingJSON) RawJSON() string {
	return r.raw
}

func (r ZonesSchemasResponseBuffering) implementsZonesSetting() {}

func (r ZonesSchemasResponseBuffering) implementsZonesZoneSettingsResponseCollectionResult() {}

// ID of the zone setting.
type ZonesSchemasResponseBufferingID string

const (
	ZonesSchemasResponseBufferingIDResponseBuffering ZonesSchemasResponseBufferingID = "response_buffering"
)

func (r ZonesSchemasResponseBufferingID) IsKnown() bool {
	switch r {
	case ZonesSchemasResponseBufferingIDResponseBuffering:
		return true
	}
	return false
}

// Value of the zone setting.
type ZonesSchemasResponseBufferingValue string

const (
	ZonesSchemasResponseBufferingValueOn  ZonesSchemasResponseBufferingValue = "on"
	ZonesSchemasResponseBufferingValueOff ZonesSchemasResponseBufferingValue = "off"
)

func (r ZonesSchemasResponseBufferingValue) IsKnown() bool {
	switch r {
	case ZonesSchemasResponseBufferingValueOn, ZonesSchemasResponseBufferingValueOff:
		return true
	}
	return false
}

// Enables or disables buffering of responses from the proxied server. Cloudflare
// may buffer the whole payload to deliver it at once to the client versus allowing
// it to be delivered in chunks. By default, the proxied server streams directly
// and is not buffered by Cloudflare. This is limited to Enterprise Zones.
type ZonesSchemasResponseBufferingParam struct {
	// ID of the zone setting.
	ID param.Field[ZonesSchemasResponseBufferingID] `json:"id"`
	// Value of the zone setting.
	Value param.Field[ZonesSchemasResponseBufferingValue] `json:"value"`
	ZonesBaseParam
}

func (r ZonesSchemasResponseBufferingParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZonesSchemasResponseBufferingParam) implementsZonesSettingUnionParam() {}

func (r ZonesSchemasResponseBufferingParam) implementsZoneSettingUpdateParamsBodyUnion() {}

// Rocket Loader is a general-purpose asynchronous JavaScript optimisation that
// prioritises rendering your content while loading your site's Javascript
// asynchronously. Turning on Rocket Loader will immediately improve a web page's
// rendering time sometimes measured as Time to First Paint (TTFP), and also the
// `window.onload` time (assuming there is JavaScript on the page). This can have a
// positive impact on your Google search ranking. When turned on, Rocket Loader
// will automatically defer the loading of all Javascript referenced in your HTML,
// with no configuration required. Refer to
// [Understanding Rocket Loader](https://support.cloudflare.com/hc/articles/200168056)
// for more information.
type ZonesSchemasRocketLoader struct {
	// ID of the zone setting.
	ID ZonesSchemasRocketLoaderID `json:"id"`
	// Value of the zone setting.
	Value ZonesSchemasRocketLoaderValue `json:"value"`
	JSON  zonesSchemasRocketLoaderJSON  `json:"-"`
	ZonesBase
}

// zonesSchemasRocketLoaderJSON contains the JSON metadata for the struct
// [ZonesSchemasRocketLoader]
type zonesSchemasRocketLoaderJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonesSchemasRocketLoader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zonesSchemasRocketLoaderJSON) RawJSON() string {
	return r.raw
}

func (r ZonesSchemasRocketLoader) implementsZonesSetting() {}

func (r ZonesSchemasRocketLoader) implementsZonesZoneSettingsResponseCollectionResult() {}

// ID of the zone setting.
type ZonesSchemasRocketLoaderID string

const (
	ZonesSchemasRocketLoaderIDRocketLoader ZonesSchemasRocketLoaderID = "rocket_loader"
)

func (r ZonesSchemasRocketLoaderID) IsKnown() bool {
	switch r {
	case ZonesSchemasRocketLoaderIDRocketLoader:
		return true
	}
	return false
}

// Value of the zone setting.
type ZonesSchemasRocketLoaderValue string

const (
	ZonesSchemasRocketLoaderValueOn  ZonesSchemasRocketLoaderValue = "on"
	ZonesSchemasRocketLoaderValueOff ZonesSchemasRocketLoaderValue = "off"
)

func (r ZonesSchemasRocketLoaderValue) IsKnown() bool {
	switch r {
	case ZonesSchemasRocketLoaderValueOn, ZonesSchemasRocketLoaderValueOff:
		return true
	}
	return false
}

// Rocket Loader is a general-purpose asynchronous JavaScript optimisation that
// prioritises rendering your content while loading your site's Javascript
// asynchronously. Turning on Rocket Loader will immediately improve a web page's
// rendering time sometimes measured as Time to First Paint (TTFP), and also the
// `window.onload` time (assuming there is JavaScript on the page). This can have a
// positive impact on your Google search ranking. When turned on, Rocket Loader
// will automatically defer the loading of all Javascript referenced in your HTML,
// with no configuration required. Refer to
// [Understanding Rocket Loader](https://support.cloudflare.com/hc/articles/200168056)
// for more information.
type ZonesSchemasRocketLoaderParam struct {
	// ID of the zone setting.
	ID param.Field[ZonesSchemasRocketLoaderID] `json:"id"`
	// Value of the zone setting.
	Value param.Field[ZonesSchemasRocketLoaderValue] `json:"value"`
	ZonesBaseParam
}

func (r ZonesSchemasRocketLoaderParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZonesSchemasRocketLoaderParam) implementsZonesSettingUnionParam() {}

func (r ZonesSchemasRocketLoaderParam) implementsZoneSettingUpdateParamsBodyUnion() {}

// Choose the appropriate security profile for your website, which will
// automatically adjust each of the security settings. If you choose to customize
// an individual security setting, the profile will become Custom.
// (https://support.cloudflare.com/hc/en-us/articles/200170056).
type ZonesSchemasSecurityLevel struct {
	// ID of the zone setting.
	ID ZonesSchemasSecurityLevelID `json:"id"`
	// Value of the zone setting.
	Value ZonesSchemasSecurityLevelValue `json:"value"`
	JSON  zonesSchemasSecurityLevelJSON  `json:"-"`
	ZonesBase
}

// zonesSchemasSecurityLevelJSON contains the JSON metadata for the struct
// [ZonesSchemasSecurityLevel]
type zonesSchemasSecurityLevelJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonesSchemasSecurityLevel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zonesSchemasSecurityLevelJSON) RawJSON() string {
	return r.raw
}

func (r ZonesSchemasSecurityLevel) implementsZonesSetting() {}

func (r ZonesSchemasSecurityLevel) implementsZonesZoneSettingsResponseCollectionResult() {}

// ID of the zone setting.
type ZonesSchemasSecurityLevelID string

const (
	ZonesSchemasSecurityLevelIDSecurityLevel ZonesSchemasSecurityLevelID = "security_level"
)

func (r ZonesSchemasSecurityLevelID) IsKnown() bool {
	switch r {
	case ZonesSchemasSecurityLevelIDSecurityLevel:
		return true
	}
	return false
}

// Value of the zone setting.
type ZonesSchemasSecurityLevelValue string

const (
	ZonesSchemasSecurityLevelValueOff            ZonesSchemasSecurityLevelValue = "off"
	ZonesSchemasSecurityLevelValueEssentiallyOff ZonesSchemasSecurityLevelValue = "essentially_off"
	ZonesSchemasSecurityLevelValueLow            ZonesSchemasSecurityLevelValue = "low"
	ZonesSchemasSecurityLevelValueMedium         ZonesSchemasSecurityLevelValue = "medium"
	ZonesSchemasSecurityLevelValueHigh           ZonesSchemasSecurityLevelValue = "high"
	ZonesSchemasSecurityLevelValueUnderAttack    ZonesSchemasSecurityLevelValue = "under_attack"
)

func (r ZonesSchemasSecurityLevelValue) IsKnown() bool {
	switch r {
	case ZonesSchemasSecurityLevelValueOff, ZonesSchemasSecurityLevelValueEssentiallyOff, ZonesSchemasSecurityLevelValueLow, ZonesSchemasSecurityLevelValueMedium, ZonesSchemasSecurityLevelValueHigh, ZonesSchemasSecurityLevelValueUnderAttack:
		return true
	}
	return false
}

// Choose the appropriate security profile for your website, which will
// automatically adjust each of the security settings. If you choose to customize
// an individual security setting, the profile will become Custom.
// (https://support.cloudflare.com/hc/en-us/articles/200170056).
type ZonesSchemasSecurityLevelParam struct {
	// ID of the zone setting.
	ID param.Field[ZonesSchemasSecurityLevelID] `json:"id"`
	// Value of the zone setting.
	Value param.Field[ZonesSchemasSecurityLevelValue] `json:"value"`
	ZonesBaseParam
}

func (r ZonesSchemasSecurityLevelParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZonesSchemasSecurityLevelParam) implementsZonesSettingUnionParam() {}

func (r ZonesSchemasSecurityLevelParam) implementsZoneSettingUpdateParamsBodyUnion() {}

// Cloudflare will treat files with the same query strings as the same file in
// cache, regardless of the order of the query strings. This is limited to
// Enterprise Zones.
type ZonesSchemasSortQueryStringForCache struct {
	// ID of the zone setting.
	ID ZonesSchemasSortQueryStringForCacheID `json:"id"`
	// Value of the zone setting.
	Value ZonesSchemasSortQueryStringForCacheValue `json:"value"`
	JSON  zonesSchemasSortQueryStringForCacheJSON  `json:"-"`
	ZonesBase
}

// zonesSchemasSortQueryStringForCacheJSON contains the JSON metadata for the
// struct [ZonesSchemasSortQueryStringForCache]
type zonesSchemasSortQueryStringForCacheJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonesSchemasSortQueryStringForCache) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zonesSchemasSortQueryStringForCacheJSON) RawJSON() string {
	return r.raw
}

func (r ZonesSchemasSortQueryStringForCache) implementsZonesSetting() {}

func (r ZonesSchemasSortQueryStringForCache) implementsZonesZoneSettingsResponseCollectionResult() {}

// ID of the zone setting.
type ZonesSchemasSortQueryStringForCacheID string

const (
	ZonesSchemasSortQueryStringForCacheIDSortQueryStringForCache ZonesSchemasSortQueryStringForCacheID = "sort_query_string_for_cache"
)

func (r ZonesSchemasSortQueryStringForCacheID) IsKnown() bool {
	switch r {
	case ZonesSchemasSortQueryStringForCacheIDSortQueryStringForCache:
		return true
	}
	return false
}

// Value of the zone setting.
type ZonesSchemasSortQueryStringForCacheValue string

const (
	ZonesSchemasSortQueryStringForCacheValueOn  ZonesSchemasSortQueryStringForCacheValue = "on"
	ZonesSchemasSortQueryStringForCacheValueOff ZonesSchemasSortQueryStringForCacheValue = "off"
)

func (r ZonesSchemasSortQueryStringForCacheValue) IsKnown() bool {
	switch r {
	case ZonesSchemasSortQueryStringForCacheValueOn, ZonesSchemasSortQueryStringForCacheValueOff:
		return true
	}
	return false
}

// Cloudflare will treat files with the same query strings as the same file in
// cache, regardless of the order of the query strings. This is limited to
// Enterprise Zones.
type ZonesSchemasSortQueryStringForCacheParam struct {
	// ID of the zone setting.
	ID param.Field[ZonesSchemasSortQueryStringForCacheID] `json:"id"`
	// Value of the zone setting.
	Value param.Field[ZonesSchemasSortQueryStringForCacheValue] `json:"value"`
	ZonesBaseParam
}

func (r ZonesSchemasSortQueryStringForCacheParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZonesSchemasSortQueryStringForCacheParam) implementsZonesSettingUnionParam() {}

func (r ZonesSchemasSortQueryStringForCacheParam) implementsZoneSettingUpdateParamsBodyUnion() {}

// SSL encrypts your visitor's connection and safeguards credit card numbers and
// other personal data to and from your website. SSL can take up to 5 minutes to
// fully activate. Requires Cloudflare active on your root domain or www domain.
// Off: no SSL between the visitor and Cloudflare, and no SSL between Cloudflare
// and your web server (all HTTP traffic). Flexible: SSL between the visitor and
// Cloudflare -- visitor sees HTTPS on your site, but no SSL between Cloudflare and
// your web server. You don't need to have an SSL cert on your web server, but your
// vistors will still see the site as being HTTPS enabled. Full: SSL between the
// visitor and Cloudflare -- visitor sees HTTPS on your site, and SSL between
// Cloudflare and your web server. You'll need to have your own SSL cert or
// self-signed cert at the very least. Full (Strict): SSL between the visitor and
// Cloudflare -- visitor sees HTTPS on your site, and SSL between Cloudflare and
// your web server. You'll need to have a valid SSL certificate installed on your
// web server. This certificate must be signed by a certificate authority, have an
// expiration date in the future, and respond for the request domain name
// (hostname). (https://support.cloudflare.com/hc/en-us/articles/200170416).
type ZonesSchemasSsl struct {
	// ID of the zone setting.
	ID ZonesSchemasSslID `json:"id"`
	// Value of the zone setting. Notes: Depends on the zone's plan level
	Value ZonesSchemasSslValue `json:"value"`
	JSON  zonesSchemasSslJSON  `json:"-"`
	ZonesBase
}

// zonesSchemasSslJSON contains the JSON metadata for the struct [ZonesSchemasSsl]
type zonesSchemasSslJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonesSchemasSsl) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zonesSchemasSslJSON) RawJSON() string {
	return r.raw
}

func (r ZonesSchemasSsl) implementsZonesSetting() {}

func (r ZonesSchemasSsl) implementsZonesZoneSettingsResponseCollectionResult() {}

// ID of the zone setting.
type ZonesSchemasSslID string

const (
	ZonesSchemasSslIDSsl ZonesSchemasSslID = "ssl"
)

func (r ZonesSchemasSslID) IsKnown() bool {
	switch r {
	case ZonesSchemasSslIDSsl:
		return true
	}
	return false
}

// Value of the zone setting. Notes: Depends on the zone's plan level
type ZonesSchemasSslValue string

const (
	ZonesSchemasSslValueOff      ZonesSchemasSslValue = "off"
	ZonesSchemasSslValueFlexible ZonesSchemasSslValue = "flexible"
	ZonesSchemasSslValueFull     ZonesSchemasSslValue = "full"
	ZonesSchemasSslValueStrict   ZonesSchemasSslValue = "strict"
)

func (r ZonesSchemasSslValue) IsKnown() bool {
	switch r {
	case ZonesSchemasSslValueOff, ZonesSchemasSslValueFlexible, ZonesSchemasSslValueFull, ZonesSchemasSslValueStrict:
		return true
	}
	return false
}

// SSL encrypts your visitor's connection and safeguards credit card numbers and
// other personal data to and from your website. SSL can take up to 5 minutes to
// fully activate. Requires Cloudflare active on your root domain or www domain.
// Off: no SSL between the visitor and Cloudflare, and no SSL between Cloudflare
// and your web server (all HTTP traffic). Flexible: SSL between the visitor and
// Cloudflare -- visitor sees HTTPS on your site, but no SSL between Cloudflare and
// your web server. You don't need to have an SSL cert on your web server, but your
// vistors will still see the site as being HTTPS enabled. Full: SSL between the
// visitor and Cloudflare -- visitor sees HTTPS on your site, and SSL between
// Cloudflare and your web server. You'll need to have your own SSL cert or
// self-signed cert at the very least. Full (Strict): SSL between the visitor and
// Cloudflare -- visitor sees HTTPS on your site, and SSL between Cloudflare and
// your web server. You'll need to have a valid SSL certificate installed on your
// web server. This certificate must be signed by a certificate authority, have an
// expiration date in the future, and respond for the request domain name
// (hostname). (https://support.cloudflare.com/hc/en-us/articles/200170416).
type ZonesSchemasSslParam struct {
	// ID of the zone setting.
	ID param.Field[ZonesSchemasSslID] `json:"id"`
	// Value of the zone setting. Notes: Depends on the zone's plan level
	Value param.Field[ZonesSchemasSslValue] `json:"value"`
	ZonesBaseParam
}

func (r ZonesSchemasSslParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZonesSchemasSslParam) implementsZonesSettingUnionParam() {}

func (r ZonesSchemasSslParam) implementsZoneSettingUpdateParamsBodyUnion() {}

// Allows customer to continue to use True Client IP (Akamai feature) in the
// headers we send to the origin. This is limited to Enterprise Zones.
type ZonesSchemasTrueClientIPHeader struct {
	// ID of the zone setting.
	ID ZonesSchemasTrueClientIPHeaderID `json:"id"`
	// Value of the zone setting.
	Value ZonesSchemasTrueClientIPHeaderValue `json:"value"`
	JSON  zonesSchemasTrueClientIPHeaderJSON  `json:"-"`
	ZonesBase
}

// zonesSchemasTrueClientIPHeaderJSON contains the JSON metadata for the struct
// [ZonesSchemasTrueClientIPHeader]
type zonesSchemasTrueClientIPHeaderJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonesSchemasTrueClientIPHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zonesSchemasTrueClientIPHeaderJSON) RawJSON() string {
	return r.raw
}

func (r ZonesSchemasTrueClientIPHeader) implementsZonesSetting() {}

func (r ZonesSchemasTrueClientIPHeader) implementsZonesZoneSettingsResponseCollectionResult() {}

// ID of the zone setting.
type ZonesSchemasTrueClientIPHeaderID string

const (
	ZonesSchemasTrueClientIPHeaderIDTrueClientIPHeader ZonesSchemasTrueClientIPHeaderID = "true_client_ip_header"
)

func (r ZonesSchemasTrueClientIPHeaderID) IsKnown() bool {
	switch r {
	case ZonesSchemasTrueClientIPHeaderIDTrueClientIPHeader:
		return true
	}
	return false
}

// Value of the zone setting.
type ZonesSchemasTrueClientIPHeaderValue string

const (
	ZonesSchemasTrueClientIPHeaderValueOn  ZonesSchemasTrueClientIPHeaderValue = "on"
	ZonesSchemasTrueClientIPHeaderValueOff ZonesSchemasTrueClientIPHeaderValue = "off"
)

func (r ZonesSchemasTrueClientIPHeaderValue) IsKnown() bool {
	switch r {
	case ZonesSchemasTrueClientIPHeaderValueOn, ZonesSchemasTrueClientIPHeaderValueOff:
		return true
	}
	return false
}

// Allows customer to continue to use True Client IP (Akamai feature) in the
// headers we send to the origin. This is limited to Enterprise Zones.
type ZonesSchemasTrueClientIPHeaderParam struct {
	// ID of the zone setting.
	ID param.Field[ZonesSchemasTrueClientIPHeaderID] `json:"id"`
	// Value of the zone setting.
	Value param.Field[ZonesSchemasTrueClientIPHeaderValue] `json:"value"`
	ZonesBaseParam
}

func (r ZonesSchemasTrueClientIPHeaderParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZonesSchemasTrueClientIPHeaderParam) implementsZonesSettingUnionParam() {}

func (r ZonesSchemasTrueClientIPHeaderParam) implementsZoneSettingUpdateParamsBodyUnion() {}

// The WAF examines HTTP requests to your website. It inspects both GET and POST
// requests and applies rules to help filter out illegitimate traffic from
// legitimate website visitors. The Cloudflare WAF inspects website addresses or
// URLs to detect anything out of the ordinary. If the Cloudflare WAF determines
// suspicious user behavior, then the WAF will 'challenge' the web visitor with a
// page that asks them to submit a CAPTCHA successfully to continue their action.
// If the challenge is failed, the action will be stopped. What this means is that
// Cloudflare's WAF will block any traffic identified as illegitimate before it
// reaches your origin web server.
// (https://support.cloudflare.com/hc/en-us/articles/200172016).
type ZonesSchemasWaf struct {
	// ID of the zone setting.
	ID ZonesSchemasWafID `json:"id"`
	// Value of the zone setting.
	Value ZonesSchemasWafValue `json:"value"`
	JSON  zonesSchemasWafJSON  `json:"-"`
	ZonesBase
}

// zonesSchemasWafJSON contains the JSON metadata for the struct [ZonesSchemasWaf]
type zonesSchemasWafJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonesSchemasWaf) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zonesSchemasWafJSON) RawJSON() string {
	return r.raw
}

func (r ZonesSchemasWaf) implementsZonesSetting() {}

func (r ZonesSchemasWaf) implementsZonesZoneSettingsResponseCollectionResult() {}

// ID of the zone setting.
type ZonesSchemasWafID string

const (
	ZonesSchemasWafIDWaf ZonesSchemasWafID = "waf"
)

func (r ZonesSchemasWafID) IsKnown() bool {
	switch r {
	case ZonesSchemasWafIDWaf:
		return true
	}
	return false
}

// Value of the zone setting.
type ZonesSchemasWafValue string

const (
	ZonesSchemasWafValueOn  ZonesSchemasWafValue = "on"
	ZonesSchemasWafValueOff ZonesSchemasWafValue = "off"
)

func (r ZonesSchemasWafValue) IsKnown() bool {
	switch r {
	case ZonesSchemasWafValueOn, ZonesSchemasWafValueOff:
		return true
	}
	return false
}

// The WAF examines HTTP requests to your website. It inspects both GET and POST
// requests and applies rules to help filter out illegitimate traffic from
// legitimate website visitors. The Cloudflare WAF inspects website addresses or
// URLs to detect anything out of the ordinary. If the Cloudflare WAF determines
// suspicious user behavior, then the WAF will 'challenge' the web visitor with a
// page that asks them to submit a CAPTCHA successfully to continue their action.
// If the challenge is failed, the action will be stopped. What this means is that
// Cloudflare's WAF will block any traffic identified as illegitimate before it
// reaches your origin web server.
// (https://support.cloudflare.com/hc/en-us/articles/200172016).
type ZonesSchemasWafParam struct {
	// ID of the zone setting.
	ID param.Field[ZonesSchemasWafID] `json:"id"`
	// Value of the zone setting.
	Value param.Field[ZonesSchemasWafValue] `json:"value"`
	ZonesBaseParam
}

func (r ZonesSchemasWafParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZonesSchemasWafParam) implementsZonesSettingUnionParam() {}

func (r ZonesSchemasWafParam) implementsZoneSettingUpdateParamsBodyUnion() {}

// Cloudflare security header for a zone.
type ZonesSecurityHeader struct {
	// ID of the zone's security header.
	ID    ZonesSecurityHeaderID    `json:"id"`
	Value ZonesSecurityHeaderValue `json:"value"`
	JSON  zonesSecurityHeaderJSON  `json:"-"`
	ZonesBase
}

// zonesSecurityHeaderJSON contains the JSON metadata for the struct
// [ZonesSecurityHeader]
type zonesSecurityHeaderJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonesSecurityHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zonesSecurityHeaderJSON) RawJSON() string {
	return r.raw
}

func (r ZonesSecurityHeader) implementsZonesSetting() {}

func (r ZonesSecurityHeader) implementsZonesZoneSettingsResponseCollectionResult() {}

// ID of the zone's security header.
type ZonesSecurityHeaderID string

const (
	ZonesSecurityHeaderIDSecurityHeader ZonesSecurityHeaderID = "security_header"
)

func (r ZonesSecurityHeaderID) IsKnown() bool {
	switch r {
	case ZonesSecurityHeaderIDSecurityHeader:
		return true
	}
	return false
}

type ZonesSecurityHeaderValue struct {
	// Strict Transport Security.
	StrictTransportSecurity ZonesSecurityHeaderValueStrictTransportSecurity `json:"strict_transport_security"`
	JSON                    zonesSecurityHeaderValueJSON                    `json:"-"`
}

// zonesSecurityHeaderValueJSON contains the JSON metadata for the struct
// [ZonesSecurityHeaderValue]
type zonesSecurityHeaderValueJSON struct {
	StrictTransportSecurity apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *ZonesSecurityHeaderValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zonesSecurityHeaderValueJSON) RawJSON() string {
	return r.raw
}

// Strict Transport Security.
type ZonesSecurityHeaderValueStrictTransportSecurity struct {
	// Whether or not strict transport security is enabled.
	Enabled bool `json:"enabled"`
	// Include all subdomains for strict transport security.
	IncludeSubdomains bool `json:"include_subdomains"`
	// Max age in seconds of the strict transport security.
	MaxAge float64 `json:"max_age"`
	// Whether or not to include 'X-Content-Type-Options: nosniff' header.
	Nosniff bool `json:"nosniff"`
	// Enable automatic preload of the HSTS configuration.
	Preload bool                                                `json:"preload"`
	JSON    zonesSecurityHeaderValueStrictTransportSecurityJSON `json:"-"`
}

// zonesSecurityHeaderValueStrictTransportSecurityJSON contains the JSON metadata
// for the struct [ZonesSecurityHeaderValueStrictTransportSecurity]
type zonesSecurityHeaderValueStrictTransportSecurityJSON struct {
	Enabled           apijson.Field
	IncludeSubdomains apijson.Field
	MaxAge            apijson.Field
	Nosniff           apijson.Field
	Preload           apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ZonesSecurityHeaderValueStrictTransportSecurity) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zonesSecurityHeaderValueStrictTransportSecurityJSON) RawJSON() string {
	return r.raw
}

// Cloudflare security header for a zone.
type ZonesSecurityHeaderParam struct {
	// ID of the zone's security header.
	ID    param.Field[ZonesSecurityHeaderID]         `json:"id"`
	Value param.Field[ZonesSecurityHeaderValueParam] `json:"value"`
	ZonesBaseParam
}

func (r ZonesSecurityHeaderParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZonesSecurityHeaderParam) implementsZonesSettingUnionParam() {}

func (r ZonesSecurityHeaderParam) implementsZoneSettingUpdateParamsBodyUnion() {}

type ZonesSecurityHeaderValueParam struct {
	// Strict Transport Security.
	StrictTransportSecurity param.Field[ZonesSecurityHeaderValueStrictTransportSecurityParam] `json:"strict_transport_security"`
}

func (r ZonesSecurityHeaderValueParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Strict Transport Security.
type ZonesSecurityHeaderValueStrictTransportSecurityParam struct {
	// Whether or not strict transport security is enabled.
	Enabled param.Field[bool] `json:"enabled"`
	// Include all subdomains for strict transport security.
	IncludeSubdomains param.Field[bool] `json:"include_subdomains"`
	// Max age in seconds of the strict transport security.
	MaxAge param.Field[float64] `json:"max_age"`
	// Whether or not to include 'X-Content-Type-Options: nosniff' header.
	Nosniff param.Field[bool] `json:"nosniff"`
	// Enable automatic preload of the HSTS configuration.
	Preload param.Field[bool] `json:"preload"`
}

func (r ZonesSecurityHeaderValueStrictTransportSecurityParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// If there is sensitive content on your website that you want visible to real
// visitors, but that you want to hide from suspicious visitors, all you have to do
// is wrap the content with Cloudflare SSE tags. Wrap any content that you want to
// be excluded from suspicious visitors in the following SSE tags:
// <!--sse--><!--/sse-->. For example: <!--sse--> Bad visitors won't see my phone
// number, 555-555-5555 <!--/sse-->. Note: SSE only will work with HTML. If you
// have HTML minification enabled, you won't see the SSE tags in your HTML source
// when it's served through Cloudflare. SSE will still function in this case, as
// Cloudflare's HTML minification and SSE functionality occur on-the-fly as the
// resource moves through our network to the visitor's computer.
// (https://support.cloudflare.com/hc/en-us/articles/200170036).
type ZonesServerSideExclude struct {
	// ID of the zone setting.
	ID ZonesServerSideExcludeID `json:"id"`
	// Value of the zone setting.
	Value ZonesServerSideExcludeValue `json:"value"`
	JSON  zonesServerSideExcludeJSON  `json:"-"`
	ZonesBase
}

// zonesServerSideExcludeJSON contains the JSON metadata for the struct
// [ZonesServerSideExclude]
type zonesServerSideExcludeJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonesServerSideExclude) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zonesServerSideExcludeJSON) RawJSON() string {
	return r.raw
}

func (r ZonesServerSideExclude) implementsZonesSetting() {}

func (r ZonesServerSideExclude) implementsZonesZoneSettingsResponseCollectionResult() {}

// ID of the zone setting.
type ZonesServerSideExcludeID string

const (
	ZonesServerSideExcludeIDServerSideExclude ZonesServerSideExcludeID = "server_side_exclude"
)

func (r ZonesServerSideExcludeID) IsKnown() bool {
	switch r {
	case ZonesServerSideExcludeIDServerSideExclude:
		return true
	}
	return false
}

// Value of the zone setting.
type ZonesServerSideExcludeValue string

const (
	ZonesServerSideExcludeValueOn  ZonesServerSideExcludeValue = "on"
	ZonesServerSideExcludeValueOff ZonesServerSideExcludeValue = "off"
)

func (r ZonesServerSideExcludeValue) IsKnown() bool {
	switch r {
	case ZonesServerSideExcludeValueOn, ZonesServerSideExcludeValueOff:
		return true
	}
	return false
}

// If there is sensitive content on your website that you want visible to real
// visitors, but that you want to hide from suspicious visitors, all you have to do
// is wrap the content with Cloudflare SSE tags. Wrap any content that you want to
// be excluded from suspicious visitors in the following SSE tags:
// <!--sse--><!--/sse-->. For example: <!--sse--> Bad visitors won't see my phone
// number, 555-555-5555 <!--/sse-->. Note: SSE only will work with HTML. If you
// have HTML minification enabled, you won't see the SSE tags in your HTML source
// when it's served through Cloudflare. SSE will still function in this case, as
// Cloudflare's HTML minification and SSE functionality occur on-the-fly as the
// resource moves through our network to the visitor's computer.
// (https://support.cloudflare.com/hc/en-us/articles/200170036).
type ZonesServerSideExcludeParam struct {
	// ID of the zone setting.
	ID param.Field[ZonesServerSideExcludeID] `json:"id"`
	// Value of the zone setting.
	Value param.Field[ZonesServerSideExcludeValue] `json:"value"`
	ZonesBaseParam
}

func (r ZonesServerSideExcludeParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZonesServerSideExcludeParam) implementsZonesSettingUnionParam() {}

func (r ZonesServerSideExcludeParam) implementsZoneSettingUpdateParamsBodyUnion() {}

// 0-RTT session resumption enabled for this zone.
type ZonesSetting struct {
	// Identifier of the zone setting.
	ID string `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZonesSettingEditable `json:"editable"`
	// ssl-recommender enrollment setting.
	Enabled bool `json:"enabled"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting. Notes: The interval (in seconds) from when
	// development mode expires (positive integer) or last expired (negative integer)
	// for the domain. If development mode has never been enabled, this value is false.
	TimeRemaining float64 `json:"time_remaining"`
	// This field can have the runtime type of [interface{}],
	// [ZonesCacheRulesAegisValue], [int64],
	// [ZonesCacheRulesOriginMaxHTTPVersionValue].
	Value interface{}      `json:"value"`
	JSON  zonesSettingJSON `json:"-"`
	union ZonesSettingUnion
}

// zonesSettingJSON contains the JSON metadata for the struct [ZonesSetting]
type zonesSettingJSON struct {
	ID            apijson.Field
	Editable      apijson.Field
	Enabled       apijson.Field
	ModifiedOn    apijson.Field
	TimeRemaining apijson.Field
	Value         apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r zonesSettingJSON) RawJSON() string {
	return r.raw
}

func (r *ZonesSetting) UnmarshalJSON(data []byte) (err error) {
	*r = ZonesSetting{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ZonesSettingUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are [Zones0rtt], [ZonesAdvancedDdos],
// [ZonesCacheRulesAegis], [ZonesAlwaysOnline], [ZonesSchemasAlwaysUseHTTPS],
// [ZonesSchemasAutomaticHTTPSRewrites], [ZonesBrotli],
// [ZonesSchemasBrowserCacheTtl], [ZonesSchemasBrowserCheck],
// [ZonesSchemasCacheLevel], [ZonesChallengeTtl], [ZonesCiphers],
// [ZonesCnameFlattening], [ZonesDevelopmentMode], [ZonesEarlyHints],
// [ZonesSchemasEdgeCacheTtl], [ZonesSchemasEmailObfuscation],
// [ZonesH2Prioritization], [ZonesHotlinkProtection], [ZonesHttp2], [ZonesHttp3],
// [ZonesImageResizing], [ZonesSchemasIPGeolocation], [ZonesIpv6],
// [ZonesMaxUpload], [ZonesMinTlsVersion], [ZonesSchemasMirage], [ZonesNel],
// [ZonesSchemasOpportunisticEncryption], [ZonesOpportunisticOnion],
// [ZonesOrangeToOrange], [ZonesSchemasOriginErrorPagePassThru],
// [ZonesCacheRulesOriginH2MaxStreams], [ZonesCacheRulesOriginMaxHTTPVersion],
// [ZonesSchemasPolish], [ZonesPrefetchPreload], [ZonesPrivacyPass],
// [ZonesProxyReadTimeout], [ZonesPseudoIpv4], [ZonesReplaceInsecureJs],
// [ZonesSchemasResponseBuffering], [ZonesSchemasRocketLoader],
// [ZonesSchemasAutomaticPlatformOptimization], [ZonesSecurityHeader],
// [ZonesSchemasSecurityLevel], [ZonesServerSideExclude], [ZonesSha1Support],
// [ZonesSchemasSortQueryStringForCache], [ZonesSchemasSsl], [ZonesSslRecommender],
// [ZonesTls1_2Only], [ZonesTls1_3], [ZonesTlsClientAuth],
// [ZonesSchemasTrueClientIPHeader], [ZonesSchemasWaf], [ZonesWebp],
// [ZonesWebsockets].
func (r ZonesSetting) AsUnion() ZonesSettingUnion {
	return r.union
}

// 0-RTT session resumption enabled for this zone.
//
// Union satisfied by [Zones0rtt], [ZonesAdvancedDdos], [ZonesCacheRulesAegis],
// [ZonesAlwaysOnline], [ZonesSchemasAlwaysUseHTTPS],
// [ZonesSchemasAutomaticHTTPSRewrites], [ZonesBrotli],
// [ZonesSchemasBrowserCacheTtl], [ZonesSchemasBrowserCheck],
// [ZonesSchemasCacheLevel], [ZonesChallengeTtl], [ZonesCiphers],
// [ZonesCnameFlattening], [ZonesDevelopmentMode], [ZonesEarlyHints],
// [ZonesSchemasEdgeCacheTtl], [ZonesSchemasEmailObfuscation],
// [ZonesH2Prioritization], [ZonesHotlinkProtection], [ZonesHttp2], [ZonesHttp3],
// [ZonesImageResizing], [ZonesSchemasIPGeolocation], [ZonesIpv6],
// [ZonesMaxUpload], [ZonesMinTlsVersion], [ZonesSchemasMirage], [ZonesNel],
// [ZonesSchemasOpportunisticEncryption], [ZonesOpportunisticOnion],
// [ZonesOrangeToOrange], [ZonesSchemasOriginErrorPagePassThru],
// [ZonesCacheRulesOriginH2MaxStreams], [ZonesCacheRulesOriginMaxHTTPVersion],
// [ZonesSchemasPolish], [ZonesPrefetchPreload], [ZonesPrivacyPass],
// [ZonesProxyReadTimeout], [ZonesPseudoIpv4], [ZonesReplaceInsecureJs],
// [ZonesSchemasResponseBuffering], [ZonesSchemasRocketLoader],
// [ZonesSchemasAutomaticPlatformOptimization], [ZonesSecurityHeader],
// [ZonesSchemasSecurityLevel], [ZonesServerSideExclude], [ZonesSha1Support],
// [ZonesSchemasSortQueryStringForCache], [ZonesSchemasSsl], [ZonesSslRecommender],
// [ZonesTls1_2Only], [ZonesTls1_3], [ZonesTlsClientAuth],
// [ZonesSchemasTrueClientIPHeader], [ZonesSchemasWaf], [ZonesWebp] or
// [ZonesWebsockets].
type ZonesSettingUnion interface {
	implementsZonesSetting()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZonesSettingUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(Zones0rtt{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesAdvancedDdos{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesCacheRulesAegis{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesAlwaysOnline{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesSchemasAlwaysUseHTTPS{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesSchemasAutomaticHTTPSRewrites{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesBrotli{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesSchemasBrowserCacheTtl{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesSchemasBrowserCheck{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesSchemasCacheLevel{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesChallengeTtl{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesCiphers{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesCnameFlattening{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesDevelopmentMode{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesEarlyHints{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesSchemasEdgeCacheTtl{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesSchemasEmailObfuscation{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesH2Prioritization{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesHotlinkProtection{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesHttp2{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesHttp3{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesImageResizing{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesSchemasIPGeolocation{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesIpv6{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesMaxUpload{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesMinTlsVersion{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesSchemasMirage{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesNel{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesSchemasOpportunisticEncryption{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesOpportunisticOnion{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesOrangeToOrange{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesSchemasOriginErrorPagePassThru{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesCacheRulesOriginH2MaxStreams{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesCacheRulesOriginMaxHTTPVersion{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesSchemasPolish{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesPrefetchPreload{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesPrivacyPass{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesProxyReadTimeout{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesPseudoIpv4{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesReplaceInsecureJs{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesSchemasResponseBuffering{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesSchemasRocketLoader{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesSchemasAutomaticPlatformOptimization{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesSecurityHeader{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesSchemasSecurityLevel{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesServerSideExclude{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesSha1Support{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesSchemasSortQueryStringForCache{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesSchemasSsl{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesSslRecommender{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesTls1_2Only{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesTls1_3{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesTlsClientAuth{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesSchemasTrueClientIPHeader{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesSchemasWaf{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesWebp{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesWebsockets{}),
		},
	)
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZonesSettingEditable bool

const (
	ZonesSettingEditableTrue  ZonesSettingEditable = true
	ZonesSettingEditableFalse ZonesSettingEditable = false
)

func (r ZonesSettingEditable) IsKnown() bool {
	switch r {
	case ZonesSettingEditableTrue, ZonesSettingEditableFalse:
		return true
	}
	return false
}

// 0-RTT session resumption enabled for this zone.
//
// Satisfied by [Zones0rttParam], [ZonesAdvancedDdosParam],
// [ZonesCacheRulesAegisParam], [ZonesAlwaysOnlineParam],
// [ZonesSchemasAlwaysUseHTTPSParam], [ZonesSchemasAutomaticHTTPSRewritesParam],
// [ZonesBrotliParam], [ZonesSchemasBrowserCacheTtlParam],
// [ZonesSchemasBrowserCheckParam], [ZonesSchemasCacheLevelParam],
// [ZonesChallengeTtlParam], [ZonesCiphersParam], [ZonesCnameFlatteningParam],
// [ZonesDevelopmentModeParam], [ZonesEarlyHintsParam],
// [ZonesSchemasEdgeCacheTtlParam], [ZonesSchemasEmailObfuscationParam],
// [ZonesH2PrioritizationParam], [ZonesHotlinkProtectionParam], [ZonesHttp2Param],
// [ZonesHttp3Param], [ZonesImageResizingParam], [ZonesSchemasIPGeolocationParam],
// [ZonesIpv6Param], [ZonesMaxUploadParam], [ZonesMinTlsVersionParam],
// [ZonesSchemasMirageParam], [ZonesNelParam],
// [ZonesSchemasOpportunisticEncryptionParam], [ZonesOpportunisticOnionParam],
// [ZonesOrangeToOrangeParam], [ZonesSchemasOriginErrorPagePassThruParam],
// [ZonesCacheRulesOriginH2MaxStreamsParam],
// [ZonesCacheRulesOriginMaxHTTPVersionParam], [ZonesSchemasPolishParam],
// [ZonesPrefetchPreloadParam], [ZonesPrivacyPassParam],
// [ZonesProxyReadTimeoutParam], [ZonesPseudoIpv4Param],
// [ZonesReplaceInsecureJsParam], [ZonesSchemasResponseBufferingParam],
// [ZonesSchemasRocketLoaderParam],
// [ZonesSchemasAutomaticPlatformOptimizationParam], [ZonesSecurityHeaderParam],
// [ZonesSchemasSecurityLevelParam], [ZonesServerSideExcludeParam],
// [ZonesSha1SupportParam], [ZonesSchemasSortQueryStringForCacheParam],
// [ZonesSchemasSslParam], [ZonesSslRecommenderParam], [ZonesTls1_2OnlyParam],
// [ZonesTls1_3Param], [ZonesTlsClientAuthParam],
// [ZonesSchemasTrueClientIPHeaderParam], [ZonesSchemasWafParam], [ZonesWebpParam],
// [ZonesWebsocketsParam].
type ZonesSettingUnionParam interface {
	implementsZonesSettingUnionParam()
}

// Allow SHA1 support.
type ZonesSha1Support struct {
	// Zone setting identifier.
	ID ZonesSha1SupportID `json:"id"`
	// Value of the zone setting.
	Value ZonesSha1SupportValue `json:"value"`
	JSON  zonesSha1SupportJSON  `json:"-"`
	ZonesBase
}

// zonesSha1SupportJSON contains the JSON metadata for the struct
// [ZonesSha1Support]
type zonesSha1SupportJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonesSha1Support) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zonesSha1SupportJSON) RawJSON() string {
	return r.raw
}

func (r ZonesSha1Support) implementsZonesSetting() {}

func (r ZonesSha1Support) implementsZonesZoneSettingsResponseCollectionResult() {}

// Zone setting identifier.
type ZonesSha1SupportID string

const (
	ZonesSha1SupportIDSha1Support ZonesSha1SupportID = "sha1_support"
)

func (r ZonesSha1SupportID) IsKnown() bool {
	switch r {
	case ZonesSha1SupportIDSha1Support:
		return true
	}
	return false
}

// Value of the zone setting.
type ZonesSha1SupportValue string

const (
	ZonesSha1SupportValueOff ZonesSha1SupportValue = "off"
	ZonesSha1SupportValueOn  ZonesSha1SupportValue = "on"
)

func (r ZonesSha1SupportValue) IsKnown() bool {
	switch r {
	case ZonesSha1SupportValueOff, ZonesSha1SupportValueOn:
		return true
	}
	return false
}

// Allow SHA1 support.
type ZonesSha1SupportParam struct {
	// Zone setting identifier.
	ID param.Field[ZonesSha1SupportID] `json:"id"`
	// Value of the zone setting.
	Value param.Field[ZonesSha1SupportValue] `json:"value"`
	ZonesBaseParam
}

func (r ZonesSha1SupportParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZonesSha1SupportParam) implementsZonesSettingUnionParam() {}

func (r ZonesSha1SupportParam) implementsZoneSettingUpdateParamsBodyUnion() {}

// Enrollment in the SSL/TLS Recommender service which tries to detect and
// recommend (by sending periodic emails) the most secure SSL/TLS setting your
// origin servers support.
type ZonesSslRecommender struct {
	// Enrollment value for SSL/TLS Recommender.
	ID ZonesSslRecommenderID `json:"id"`
	// ssl-recommender enrollment setting.
	Enabled bool                    `json:"enabled"`
	JSON    zonesSslRecommenderJSON `json:"-"`
}

// zonesSslRecommenderJSON contains the JSON metadata for the struct
// [ZonesSslRecommender]
type zonesSslRecommenderJSON struct {
	ID          apijson.Field
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonesSslRecommender) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zonesSslRecommenderJSON) RawJSON() string {
	return r.raw
}

func (r ZonesSslRecommender) implementsZonesSetting() {}

func (r ZonesSslRecommender) implementsZonesZoneSettingsResponseCollectionResult() {}

// Enrollment value for SSL/TLS Recommender.
type ZonesSslRecommenderID string

const (
	ZonesSslRecommenderIDSslRecommender ZonesSslRecommenderID = "ssl_recommender"
)

func (r ZonesSslRecommenderID) IsKnown() bool {
	switch r {
	case ZonesSslRecommenderIDSslRecommender:
		return true
	}
	return false
}

// Enrollment in the SSL/TLS Recommender service which tries to detect and
// recommend (by sending periodic emails) the most secure SSL/TLS setting your
// origin servers support.
type ZonesSslRecommenderParam struct {
	// Enrollment value for SSL/TLS Recommender.
	ID param.Field[ZonesSslRecommenderID] `json:"id"`
	// ssl-recommender enrollment setting.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r ZonesSslRecommenderParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZonesSslRecommenderParam) implementsZonesSettingUnionParam() {}

func (r ZonesSslRecommenderParam) implementsZoneSettingUpdateParamsBodyUnion() {}

// Only allows TLS1.2.
type ZonesTls1_2Only struct {
	// Zone setting identifier.
	ID ZonesTls1_2OnlyID `json:"id"`
	// Value of the zone setting.
	Value ZonesTls1_2OnlyValue `json:"value"`
	JSON  zonesTls1_2OnlyJSON  `json:"-"`
	ZonesBase
}

// zonesTls1_2OnlyJSON contains the JSON metadata for the struct [ZonesTls1_2Only]
type zonesTls1_2OnlyJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonesTls1_2Only) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zonesTls1_2OnlyJSON) RawJSON() string {
	return r.raw
}

func (r ZonesTls1_2Only) implementsZonesSetting() {}

func (r ZonesTls1_2Only) implementsZonesZoneSettingsResponseCollectionResult() {}

// Zone setting identifier.
type ZonesTls1_2OnlyID string

const (
	ZonesTls1_2OnlyIDTls1_2Only ZonesTls1_2OnlyID = "tls_1_2_only"
)

func (r ZonesTls1_2OnlyID) IsKnown() bool {
	switch r {
	case ZonesTls1_2OnlyIDTls1_2Only:
		return true
	}
	return false
}

// Value of the zone setting.
type ZonesTls1_2OnlyValue string

const (
	ZonesTls1_2OnlyValueOff ZonesTls1_2OnlyValue = "off"
	ZonesTls1_2OnlyValueOn  ZonesTls1_2OnlyValue = "on"
)

func (r ZonesTls1_2OnlyValue) IsKnown() bool {
	switch r {
	case ZonesTls1_2OnlyValueOff, ZonesTls1_2OnlyValueOn:
		return true
	}
	return false
}

// Only allows TLS1.2.
type ZonesTls1_2OnlyParam struct {
	// Zone setting identifier.
	ID param.Field[ZonesTls1_2OnlyID] `json:"id"`
	// Value of the zone setting.
	Value param.Field[ZonesTls1_2OnlyValue] `json:"value"`
	ZonesBaseParam
}

func (r ZonesTls1_2OnlyParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZonesTls1_2OnlyParam) implementsZonesSettingUnionParam() {}

func (r ZonesTls1_2OnlyParam) implementsZoneSettingUpdateParamsBodyUnion() {}

// Enables Crypto TLS 1.3 feature for a zone.
type ZonesTls1_3 struct {
	// ID of the zone setting.
	ID ZonesTls1_3ID `json:"id"`
	// Value of the zone setting. Notes: Default value depends on the zone's plan
	// level.
	Value ZonesTls1_3Value `json:"value"`
	JSON  zonesTls1_3JSON  `json:"-"`
	ZonesBase
}

// zonesTls1_3JSON contains the JSON metadata for the struct [ZonesTls1_3]
type zonesTls1_3JSON struct {
	ID          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonesTls1_3) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zonesTls1_3JSON) RawJSON() string {
	return r.raw
}

func (r ZonesTls1_3) implementsZonesSetting() {}

func (r ZonesTls1_3) implementsZonesZoneSettingsResponseCollectionResult() {}

// ID of the zone setting.
type ZonesTls1_3ID string

const (
	ZonesTls1_3IDTls1_3 ZonesTls1_3ID = "tls_1_3"
)

func (r ZonesTls1_3ID) IsKnown() bool {
	switch r {
	case ZonesTls1_3IDTls1_3:
		return true
	}
	return false
}

// Value of the zone setting. Notes: Default value depends on the zone's plan
// level.
type ZonesTls1_3Value string

const (
	ZonesTls1_3ValueOn  ZonesTls1_3Value = "on"
	ZonesTls1_3ValueOff ZonesTls1_3Value = "off"
	ZonesTls1_3ValueZrt ZonesTls1_3Value = "zrt"
)

func (r ZonesTls1_3Value) IsKnown() bool {
	switch r {
	case ZonesTls1_3ValueOn, ZonesTls1_3ValueOff, ZonesTls1_3ValueZrt:
		return true
	}
	return false
}

// Enables Crypto TLS 1.3 feature for a zone.
type ZonesTls1_3Param struct {
	// ID of the zone setting.
	ID param.Field[ZonesTls1_3ID] `json:"id"`
	// Value of the zone setting. Notes: Default value depends on the zone's plan
	// level.
	Value param.Field[ZonesTls1_3Value] `json:"value"`
	ZonesBaseParam
}

func (r ZonesTls1_3Param) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZonesTls1_3Param) implementsZonesSettingUnionParam() {}

func (r ZonesTls1_3Param) implementsZoneSettingUpdateParamsBodyUnion() {}

// TLS Client Auth requires Cloudflare to connect to your origin server using a
// client certificate (Enterprise Only).
type ZonesTlsClientAuth struct {
	// ID of the zone setting.
	ID ZonesTlsClientAuthID `json:"id"`
	// value of the zone setting.
	Value ZonesTlsClientAuthValue `json:"value"`
	JSON  zonesTlsClientAuthJSON  `json:"-"`
	ZonesBase
}

// zonesTlsClientAuthJSON contains the JSON metadata for the struct
// [ZonesTlsClientAuth]
type zonesTlsClientAuthJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonesTlsClientAuth) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zonesTlsClientAuthJSON) RawJSON() string {
	return r.raw
}

func (r ZonesTlsClientAuth) implementsZonesSetting() {}

func (r ZonesTlsClientAuth) implementsZonesZoneSettingsResponseCollectionResult() {}

// ID of the zone setting.
type ZonesTlsClientAuthID string

const (
	ZonesTlsClientAuthIDTlsClientAuth ZonesTlsClientAuthID = "tls_client_auth"
)

func (r ZonesTlsClientAuthID) IsKnown() bool {
	switch r {
	case ZonesTlsClientAuthIDTlsClientAuth:
		return true
	}
	return false
}

// value of the zone setting.
type ZonesTlsClientAuthValue string

const (
	ZonesTlsClientAuthValueOn  ZonesTlsClientAuthValue = "on"
	ZonesTlsClientAuthValueOff ZonesTlsClientAuthValue = "off"
)

func (r ZonesTlsClientAuthValue) IsKnown() bool {
	switch r {
	case ZonesTlsClientAuthValueOn, ZonesTlsClientAuthValueOff:
		return true
	}
	return false
}

// TLS Client Auth requires Cloudflare to connect to your origin server using a
// client certificate (Enterprise Only).
type ZonesTlsClientAuthParam struct {
	// ID of the zone setting.
	ID param.Field[ZonesTlsClientAuthID] `json:"id"`
	// value of the zone setting.
	Value param.Field[ZonesTlsClientAuthValue] `json:"value"`
	ZonesBaseParam
}

func (r ZonesTlsClientAuthParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZonesTlsClientAuthParam) implementsZonesSettingUnionParam() {}

func (r ZonesTlsClientAuthParam) implementsZoneSettingUpdateParamsBodyUnion() {}

// When the client requesting the image supports the WebP image codec, and WebP
// offers a performance advantage over the original image format, Cloudflare will
// serve a WebP version of the original image.
type ZonesWebp struct {
	// ID of the zone setting.
	ID ZonesWebpID `json:"id"`
	// Value of the zone setting.
	Value ZonesWebpValue `json:"value"`
	JSON  zonesWebpJSON  `json:"-"`
	ZonesBase
}

// zonesWebpJSON contains the JSON metadata for the struct [ZonesWebp]
type zonesWebpJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonesWebp) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zonesWebpJSON) RawJSON() string {
	return r.raw
}

func (r ZonesWebp) implementsZonesSetting() {}

func (r ZonesWebp) implementsZonesZoneSettingsResponseCollectionResult() {}

// ID of the zone setting.
type ZonesWebpID string

const (
	ZonesWebpIDWebp ZonesWebpID = "webp"
)

func (r ZonesWebpID) IsKnown() bool {
	switch r {
	case ZonesWebpIDWebp:
		return true
	}
	return false
}

// Value of the zone setting.
type ZonesWebpValue string

const (
	ZonesWebpValueOff ZonesWebpValue = "off"
	ZonesWebpValueOn  ZonesWebpValue = "on"
)

func (r ZonesWebpValue) IsKnown() bool {
	switch r {
	case ZonesWebpValueOff, ZonesWebpValueOn:
		return true
	}
	return false
}

// When the client requesting the image supports the WebP image codec, and WebP
// offers a performance advantage over the original image format, Cloudflare will
// serve a WebP version of the original image.
type ZonesWebpParam struct {
	// ID of the zone setting.
	ID param.Field[ZonesWebpID] `json:"id"`
	// Value of the zone setting.
	Value param.Field[ZonesWebpValue] `json:"value"`
	ZonesBaseParam
}

func (r ZonesWebpParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZonesWebpParam) implementsZonesSettingUnionParam() {}

func (r ZonesWebpParam) implementsZoneSettingUpdateParamsBodyUnion() {}

// WebSockets are open connections sustained between the client and the origin
// server. Inside a WebSockets connection, the client and the origin can pass data
// back and forth without having to reestablish sessions. This makes exchanging
// data within a WebSockets connection fast. WebSockets are often used for
// real-time applications such as live chat and gaming. For more information refer
// to
// [Can I use Cloudflare with Websockets](https://support.cloudflare.com/hc/en-us/articles/200169466-Can-I-use-Cloudflare-with-WebSockets-).
type ZonesWebsockets struct {
	// ID of the zone setting.
	ID ZonesWebsocketsID `json:"id"`
	// Value of the zone setting.
	Value ZonesWebsocketsValue `json:"value"`
	JSON  zonesWebsocketsJSON  `json:"-"`
	ZonesBase
}

// zonesWebsocketsJSON contains the JSON metadata for the struct [ZonesWebsockets]
type zonesWebsocketsJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonesWebsockets) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zonesWebsocketsJSON) RawJSON() string {
	return r.raw
}

func (r ZonesWebsockets) implementsZonesSetting() {}

func (r ZonesWebsockets) implementsZonesZoneSettingsResponseCollectionResult() {}

// ID of the zone setting.
type ZonesWebsocketsID string

const (
	ZonesWebsocketsIDWebsockets ZonesWebsocketsID = "websockets"
)

func (r ZonesWebsocketsID) IsKnown() bool {
	switch r {
	case ZonesWebsocketsIDWebsockets:
		return true
	}
	return false
}

// Value of the zone setting.
type ZonesWebsocketsValue string

const (
	ZonesWebsocketsValueOff ZonesWebsocketsValue = "off"
	ZonesWebsocketsValueOn  ZonesWebsocketsValue = "on"
)

func (r ZonesWebsocketsValue) IsKnown() bool {
	switch r {
	case ZonesWebsocketsValueOff, ZonesWebsocketsValueOn:
		return true
	}
	return false
}

// WebSockets are open connections sustained between the client and the origin
// server. Inside a WebSockets connection, the client and the origin can pass data
// back and forth without having to reestablish sessions. This makes exchanging
// data within a WebSockets connection fast. WebSockets are often used for
// real-time applications such as live chat and gaming. For more information refer
// to
// [Can I use Cloudflare with Websockets](https://support.cloudflare.com/hc/en-us/articles/200169466-Can-I-use-Cloudflare-with-WebSockets-).
type ZonesWebsocketsParam struct {
	// ID of the zone setting.
	ID param.Field[ZonesWebsocketsID] `json:"id"`
	// Value of the zone setting.
	Value param.Field[ZonesWebsocketsValue] `json:"value"`
	ZonesBaseParam
}

func (r ZonesWebsocketsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZonesWebsocketsParam) implementsZonesSettingUnionParam() {}

func (r ZonesWebsocketsParam) implementsZoneSettingUpdateParamsBodyUnion() {}

type ZonesZoneSettingsResponseCollection struct {
	Result []ZonesZoneSettingsResponseCollectionResult `json:"result"`
	JSON   zonesZoneSettingsResponseCollectionJSON     `json:"-"`
	APIResponseZones
}

// zonesZoneSettingsResponseCollectionJSON contains the JSON metadata for the
// struct [ZonesZoneSettingsResponseCollection]
type zonesZoneSettingsResponseCollectionJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonesZoneSettingsResponseCollection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zonesZoneSettingsResponseCollectionJSON) RawJSON() string {
	return r.raw
}

// 0-RTT session resumption enabled for this zone.
type ZonesZoneSettingsResponseCollectionResult struct {
	// Identifier of the zone setting.
	ID string `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZonesZoneSettingsResponseCollectionResultEditable `json:"editable"`
	// ssl-recommender enrollment setting.
	Enabled bool `json:"enabled"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting. Notes: The interval (in seconds) from when
	// development mode expires (positive integer) or last expired (negative integer)
	// for the domain. If development mode has never been enabled, this value is false.
	TimeRemaining float64 `json:"time_remaining"`
	// This field can have the runtime type of [interface{}],
	// [ZonesCacheRulesAegisValue], [int64],
	// [ZonesCacheRulesOriginMaxHTTPVersionValue].
	Value interface{}                                   `json:"value"`
	JSON  zonesZoneSettingsResponseCollectionResultJSON `json:"-"`
	union ZonesZoneSettingsResponseCollectionResultUnion
}

// zonesZoneSettingsResponseCollectionResultJSON contains the JSON metadata for the
// struct [ZonesZoneSettingsResponseCollectionResult]
type zonesZoneSettingsResponseCollectionResultJSON struct {
	ID            apijson.Field
	Editable      apijson.Field
	Enabled       apijson.Field
	ModifiedOn    apijson.Field
	TimeRemaining apijson.Field
	Value         apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r zonesZoneSettingsResponseCollectionResultJSON) RawJSON() string {
	return r.raw
}

func (r *ZonesZoneSettingsResponseCollectionResult) UnmarshalJSON(data []byte) (err error) {
	*r = ZonesZoneSettingsResponseCollectionResult{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ZonesZoneSettingsResponseCollectionResultUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are [Zones0rtt], [ZonesAdvancedDdos],
// [ZonesCacheRulesAegis], [ZonesAlwaysOnline], [ZonesSchemasAlwaysUseHTTPS],
// [ZonesSchemasAutomaticHTTPSRewrites], [ZonesBrotli],
// [ZonesSchemasBrowserCacheTtl], [ZonesSchemasBrowserCheck],
// [ZonesSchemasCacheLevel], [ZonesChallengeTtl], [ZonesCiphers],
// [ZonesCnameFlattening], [ZonesDevelopmentMode], [ZonesEarlyHints],
// [ZonesSchemasEdgeCacheTtl], [ZonesSchemasEmailObfuscation],
// [ZonesH2Prioritization], [ZonesHotlinkProtection], [ZonesHttp2], [ZonesHttp3],
// [ZonesImageResizing], [ZonesSchemasIPGeolocation], [ZonesIpv6],
// [ZonesMaxUpload], [ZonesMinTlsVersion], [ZonesSchemasMirage], [ZonesNel],
// [ZonesSchemasOpportunisticEncryption], [ZonesOpportunisticOnion],
// [ZonesOrangeToOrange], [ZonesSchemasOriginErrorPagePassThru],
// [ZonesCacheRulesOriginH2MaxStreams], [ZonesCacheRulesOriginMaxHTTPVersion],
// [ZonesSchemasPolish], [ZonesPrefetchPreload], [ZonesPrivacyPass],
// [ZonesProxyReadTimeout], [ZonesPseudoIpv4], [ZonesReplaceInsecureJs],
// [ZonesSchemasResponseBuffering], [ZonesSchemasRocketLoader],
// [ZonesSchemasAutomaticPlatformOptimization], [ZonesSecurityHeader],
// [ZonesSchemasSecurityLevel], [ZonesServerSideExclude], [ZonesSha1Support],
// [ZonesSchemasSortQueryStringForCache], [ZonesSchemasSsl], [ZonesSslRecommender],
// [ZonesTls1_2Only], [ZonesTls1_3], [ZonesTlsClientAuth],
// [ZonesSchemasTrueClientIPHeader], [ZonesSchemasWaf], [ZonesWebp],
// [ZonesWebsockets].
func (r ZonesZoneSettingsResponseCollectionResult) AsUnion() ZonesZoneSettingsResponseCollectionResultUnion {
	return r.union
}

// 0-RTT session resumption enabled for this zone.
//
// Union satisfied by [Zones0rtt], [ZonesAdvancedDdos], [ZonesCacheRulesAegis],
// [ZonesAlwaysOnline], [ZonesSchemasAlwaysUseHTTPS],
// [ZonesSchemasAutomaticHTTPSRewrites], [ZonesBrotli],
// [ZonesSchemasBrowserCacheTtl], [ZonesSchemasBrowserCheck],
// [ZonesSchemasCacheLevel], [ZonesChallengeTtl], [ZonesCiphers],
// [ZonesCnameFlattening], [ZonesDevelopmentMode], [ZonesEarlyHints],
// [ZonesSchemasEdgeCacheTtl], [ZonesSchemasEmailObfuscation],
// [ZonesH2Prioritization], [ZonesHotlinkProtection], [ZonesHttp2], [ZonesHttp3],
// [ZonesImageResizing], [ZonesSchemasIPGeolocation], [ZonesIpv6],
// [ZonesMaxUpload], [ZonesMinTlsVersion], [ZonesSchemasMirage], [ZonesNel],
// [ZonesSchemasOpportunisticEncryption], [ZonesOpportunisticOnion],
// [ZonesOrangeToOrange], [ZonesSchemasOriginErrorPagePassThru],
// [ZonesCacheRulesOriginH2MaxStreams], [ZonesCacheRulesOriginMaxHTTPVersion],
// [ZonesSchemasPolish], [ZonesPrefetchPreload], [ZonesPrivacyPass],
// [ZonesProxyReadTimeout], [ZonesPseudoIpv4], [ZonesReplaceInsecureJs],
// [ZonesSchemasResponseBuffering], [ZonesSchemasRocketLoader],
// [ZonesSchemasAutomaticPlatformOptimization], [ZonesSecurityHeader],
// [ZonesSchemasSecurityLevel], [ZonesServerSideExclude], [ZonesSha1Support],
// [ZonesSchemasSortQueryStringForCache], [ZonesSchemasSsl], [ZonesSslRecommender],
// [ZonesTls1_2Only], [ZonesTls1_3], [ZonesTlsClientAuth],
// [ZonesSchemasTrueClientIPHeader], [ZonesSchemasWaf], [ZonesWebp] or
// [ZonesWebsockets].
type ZonesZoneSettingsResponseCollectionResultUnion interface {
	implementsZonesZoneSettingsResponseCollectionResult()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZonesZoneSettingsResponseCollectionResultUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(Zones0rtt{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesAdvancedDdos{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesCacheRulesAegis{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesAlwaysOnline{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesSchemasAlwaysUseHTTPS{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesSchemasAutomaticHTTPSRewrites{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesBrotli{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesSchemasBrowserCacheTtl{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesSchemasBrowserCheck{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesSchemasCacheLevel{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesChallengeTtl{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesCiphers{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesCnameFlattening{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesDevelopmentMode{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesEarlyHints{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesSchemasEdgeCacheTtl{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesSchemasEmailObfuscation{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesH2Prioritization{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesHotlinkProtection{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesHttp2{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesHttp3{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesImageResizing{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesSchemasIPGeolocation{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesIpv6{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesMaxUpload{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesMinTlsVersion{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesSchemasMirage{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesNel{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesSchemasOpportunisticEncryption{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesOpportunisticOnion{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesOrangeToOrange{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesSchemasOriginErrorPagePassThru{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesCacheRulesOriginH2MaxStreams{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesCacheRulesOriginMaxHTTPVersion{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesSchemasPolish{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesPrefetchPreload{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesPrivacyPass{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesProxyReadTimeout{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesPseudoIpv4{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesReplaceInsecureJs{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesSchemasResponseBuffering{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesSchemasRocketLoader{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesSchemasAutomaticPlatformOptimization{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesSecurityHeader{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesSchemasSecurityLevel{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesServerSideExclude{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesSha1Support{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesSchemasSortQueryStringForCache{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesSchemasSsl{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesSslRecommender{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesTls1_2Only{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesTls1_3{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesTlsClientAuth{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesSchemasTrueClientIPHeader{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesSchemasWaf{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesWebp{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZonesWebsockets{}),
		},
	)
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZonesZoneSettingsResponseCollectionResultEditable bool

const (
	ZonesZoneSettingsResponseCollectionResultEditableTrue  ZonesZoneSettingsResponseCollectionResultEditable = true
	ZonesZoneSettingsResponseCollectionResultEditableFalse ZonesZoneSettingsResponseCollectionResultEditable = false
)

func (r ZonesZoneSettingsResponseCollectionResultEditable) IsKnown() bool {
	switch r {
	case ZonesZoneSettingsResponseCollectionResultEditableTrue, ZonesZoneSettingsResponseCollectionResultEditableFalse:
		return true
	}
	return false
}

type ZoneSettingGetSettingResponse struct {
	// 0-RTT session resumption enabled for this zone.
	Result ZonesSetting                      `json:"result"`
	JSON   zoneSettingGetSettingResponseJSON `json:"-"`
	APIResponseZones
}

// zoneSettingGetSettingResponseJSON contains the JSON metadata for the struct
// [ZoneSettingGetSettingResponse]
type zoneSettingGetSettingResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingGetSettingResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingGetSettingResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneSettingUpdateSettingResponse struct {
	// 0-RTT session resumption enabled for this zone.
	Result ZonesSetting                         `json:"result"`
	JSON   zoneSettingUpdateSettingResponseJSON `json:"-"`
	APIResponseZones
}

// zoneSettingUpdateSettingResponseJSON contains the JSON metadata for the struct
// [ZoneSettingUpdateSettingResponse]
type zoneSettingUpdateSettingResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingUpdateSettingResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingUpdateSettingResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneSettingUpdateParams struct {
	Body []ZoneSettingUpdateParamsBodyUnion `json:"body,required"`
}

func (r ZoneSettingUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

// 0-RTT session resumption enabled for this zone.
//
// Satisfied by [Zones0rttParam], [ZonesAdvancedDdosParam],
// [ZonesCacheRulesAegisParam], [ZonesAlwaysOnlineParam],
// [ZonesSchemasAlwaysUseHTTPSParam], [ZonesSchemasAutomaticHTTPSRewritesParam],
// [ZonesBrotliParam], [ZonesSchemasBrowserCacheTtlParam],
// [ZonesSchemasBrowserCheckParam], [ZonesSchemasCacheLevelParam],
// [ZonesChallengeTtlParam], [ZonesCiphersParam], [ZonesCnameFlatteningParam],
// [ZonesDevelopmentModeParam], [ZonesEarlyHintsParam],
// [ZonesSchemasEdgeCacheTtlParam], [ZonesSchemasEmailObfuscationParam],
// [ZonesH2PrioritizationParam], [ZonesHotlinkProtectionParam], [ZonesHttp2Param],
// [ZonesHttp3Param], [ZonesSchemasIPGeolocationParam], [ZonesIpv6Param],
// [ZonesMaxUploadParam], [ZonesMinTlsVersionParam], [ZonesSchemasMirageParam],
// [ZonesNelParam], [ZonesSchemasOpportunisticEncryptionParam],
// [ZonesOpportunisticOnionParam], [ZonesOrangeToOrangeParam],
// [ZonesSchemasOriginErrorPagePassThruParam],
// [ZonesCacheRulesOriginH2MaxStreamsParam],
// [ZonesCacheRulesOriginMaxHTTPVersionParam], [ZonesSchemasPolishParam],
// [ZonesPrefetchPreloadParam], [ZonesPrivacyPassParam],
// [ZonesProxyReadTimeoutParam], [ZonesPseudoIpv4Param],
// [ZonesReplaceInsecureJsParam], [ZonesSchemasResponseBufferingParam],
// [ZonesSchemasRocketLoaderParam],
// [ZonesSchemasAutomaticPlatformOptimizationParam], [ZonesSecurityHeaderParam],
// [ZonesSchemasSecurityLevelParam], [ZonesServerSideExcludeParam],
// [ZonesSha1SupportParam], [ZonesSchemasSortQueryStringForCacheParam],
// [ZonesSchemasSslParam], [ZonesSslRecommenderParam], [ZonesTls1_2OnlyParam],
// [ZonesTls1_3Param], [ZonesTlsClientAuthParam],
// [ZonesSchemasTrueClientIPHeaderParam], [ZonesSchemasWafParam], [ZonesWebpParam],
// [ZonesWebsocketsParam].
type ZoneSettingUpdateParamsBodyUnion interface {
	implementsZoneSettingUpdateParamsBodyUnion()
}

type ZoneSettingUpdateSettingParams struct {
	// 0-RTT session resumption enabled for this zone.
	ZonesSetting ZonesSettingUnionParam `json:"zones_setting,required"`
}

func (r ZoneSettingUpdateSettingParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.ZonesSetting)
}
