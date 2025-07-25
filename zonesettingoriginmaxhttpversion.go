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
	Errors   []MessagesCacheRulesItem `json:"errors,required"`
	Messages []MessagesCacheRulesItem `json:"messages,required"`
	// Whether the API call was successful
	Success ZoneSettingOriginMaxHTTPVersionGetResponseSuccess `json:"success,required"`
	Result  ZoneSettingOriginMaxHTTPVersionGetResponseResult  `json:"result"`
	JSON    zoneSettingOriginMaxHTTPVersionGetResponseJSON    `json:"-"`
}

// zoneSettingOriginMaxHTTPVersionGetResponseJSON contains the JSON metadata for
// the struct [ZoneSettingOriginMaxHTTPVersionGetResponse]
type zoneSettingOriginMaxHTTPVersionGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingOriginMaxHTTPVersionGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingOriginMaxHTTPVersionGetResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ZoneSettingOriginMaxHTTPVersionGetResponseSuccess bool

const (
	ZoneSettingOriginMaxHTTPVersionGetResponseSuccessTrue ZoneSettingOriginMaxHTTPVersionGetResponseSuccess = true
)

func (r ZoneSettingOriginMaxHTTPVersionGetResponseSuccess) IsKnown() bool {
	switch r {
	case ZoneSettingOriginMaxHTTPVersionGetResponseSuccessTrue:
		return true
	}
	return false
}

type ZoneSettingOriginMaxHTTPVersionGetResponseResult struct {
	// Value of the zone setting.
	ID ZoneSettingOriginMaxHTTPVersionGetResponseResultID `json:"id,required"`
	// Whether the setting is editable
	Editable bool `json:"editable,required"`
	// The value of the feature
	Value CacheRulesOriginMaxHTTPVersionValue `json:"value,required"`
	// Last time this setting was modified.
	ModifiedOn time.Time                                            `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingOriginMaxHTTPVersionGetResponseResultJSON `json:"-"`
}

// zoneSettingOriginMaxHTTPVersionGetResponseResultJSON contains the JSON metadata
// for the struct [ZoneSettingOriginMaxHTTPVersionGetResponseResult]
type zoneSettingOriginMaxHTTPVersionGetResponseResultJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	Value       apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingOriginMaxHTTPVersionGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingOriginMaxHTTPVersionGetResponseResultJSON) RawJSON() string {
	return r.raw
}

// Value of the zone setting.
type ZoneSettingOriginMaxHTTPVersionGetResponseResultID string

const (
	ZoneSettingOriginMaxHTTPVersionGetResponseResultIDOriginMaxHTTPVersion ZoneSettingOriginMaxHTTPVersionGetResponseResultID = "origin_max_http_version"
)

func (r ZoneSettingOriginMaxHTTPVersionGetResponseResultID) IsKnown() bool {
	switch r {
	case ZoneSettingOriginMaxHTTPVersionGetResponseResultIDOriginMaxHTTPVersion:
		return true
	}
	return false
}

type ZoneSettingOriginMaxHTTPVersionUpdateResponse struct {
	Errors   []MessagesCacheRulesItem `json:"errors,required"`
	Messages []MessagesCacheRulesItem `json:"messages,required"`
	// Whether the API call was successful
	Success ZoneSettingOriginMaxHTTPVersionUpdateResponseSuccess `json:"success,required"`
	Result  ZoneSettingOriginMaxHTTPVersionUpdateResponseResult  `json:"result"`
	JSON    zoneSettingOriginMaxHTTPVersionUpdateResponseJSON    `json:"-"`
}

// zoneSettingOriginMaxHTTPVersionUpdateResponseJSON contains the JSON metadata for
// the struct [ZoneSettingOriginMaxHTTPVersionUpdateResponse]
type zoneSettingOriginMaxHTTPVersionUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingOriginMaxHTTPVersionUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingOriginMaxHTTPVersionUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ZoneSettingOriginMaxHTTPVersionUpdateResponseSuccess bool

const (
	ZoneSettingOriginMaxHTTPVersionUpdateResponseSuccessTrue ZoneSettingOriginMaxHTTPVersionUpdateResponseSuccess = true
)

func (r ZoneSettingOriginMaxHTTPVersionUpdateResponseSuccess) IsKnown() bool {
	switch r {
	case ZoneSettingOriginMaxHTTPVersionUpdateResponseSuccessTrue:
		return true
	}
	return false
}

type ZoneSettingOriginMaxHTTPVersionUpdateResponseResult struct {
	// Value of the zone setting.
	ID ZoneSettingOriginMaxHTTPVersionUpdateResponseResultID `json:"id,required"`
	// Whether the setting is editable
	Editable bool `json:"editable,required"`
	// The value of the feature
	Value CacheRulesOriginMaxHTTPVersionValue `json:"value,required"`
	// Last time this setting was modified.
	ModifiedOn time.Time                                               `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingOriginMaxHTTPVersionUpdateResponseResultJSON `json:"-"`
}

// zoneSettingOriginMaxHTTPVersionUpdateResponseResultJSON contains the JSON
// metadata for the struct [ZoneSettingOriginMaxHTTPVersionUpdateResponseResult]
type zoneSettingOriginMaxHTTPVersionUpdateResponseResultJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	Value       apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingOriginMaxHTTPVersionUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingOriginMaxHTTPVersionUpdateResponseResultJSON) RawJSON() string {
	return r.raw
}

// Value of the zone setting.
type ZoneSettingOriginMaxHTTPVersionUpdateResponseResultID string

const (
	ZoneSettingOriginMaxHTTPVersionUpdateResponseResultIDOriginMaxHTTPVersion ZoneSettingOriginMaxHTTPVersionUpdateResponseResultID = "origin_max_http_version"
)

func (r ZoneSettingOriginMaxHTTPVersionUpdateResponseResultID) IsKnown() bool {
	switch r {
	case ZoneSettingOriginMaxHTTPVersionUpdateResponseResultIDOriginMaxHTTPVersion:
		return true
	}
	return false
}

type ZoneSettingOriginMaxHTTPVersionUpdateParams struct {
	// Value of the Origin Max HTTP Version Setting.
	Value param.Field[CacheRulesOriginMaxHTTPVersionValue] `json:"value,required"`
}

func (r ZoneSettingOriginMaxHTTPVersionUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
