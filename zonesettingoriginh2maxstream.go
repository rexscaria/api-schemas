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

type ZoneSettingOriginH2MaxStreamGetResponse struct {
	Errors   []MessagesCacheRulesItem `json:"errors,required"`
	Messages []MessagesCacheRulesItem `json:"messages,required"`
	// Whether the API call was successful
	Success ZoneSettingOriginH2MaxStreamGetResponseSuccess `json:"success,required"`
	Result  ZoneSettingOriginH2MaxStreamGetResponseResult  `json:"result"`
	JSON    zoneSettingOriginH2MaxStreamGetResponseJSON    `json:"-"`
}

// zoneSettingOriginH2MaxStreamGetResponseJSON contains the JSON metadata for the
// struct [ZoneSettingOriginH2MaxStreamGetResponse]
type zoneSettingOriginH2MaxStreamGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingOriginH2MaxStreamGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingOriginH2MaxStreamGetResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ZoneSettingOriginH2MaxStreamGetResponseSuccess bool

const (
	ZoneSettingOriginH2MaxStreamGetResponseSuccessTrue ZoneSettingOriginH2MaxStreamGetResponseSuccess = true
)

func (r ZoneSettingOriginH2MaxStreamGetResponseSuccess) IsKnown() bool {
	switch r {
	case ZoneSettingOriginH2MaxStreamGetResponseSuccessTrue:
		return true
	}
	return false
}

type ZoneSettingOriginH2MaxStreamGetResponseResult struct {
	// Value of the zone setting.
	ID ZoneSettingOriginH2MaxStreamGetResponseResultID `json:"id,required"`
	// Whether the setting is editable
	Editable bool `json:"editable,required"`
	// The value of the feature
	Value string `json:"value,required"`
	// Last time this setting was modified.
	ModifiedOn time.Time                                         `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingOriginH2MaxStreamGetResponseResultJSON `json:"-"`
}

// zoneSettingOriginH2MaxStreamGetResponseResultJSON contains the JSON metadata for
// the struct [ZoneSettingOriginH2MaxStreamGetResponseResult]
type zoneSettingOriginH2MaxStreamGetResponseResultJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	Value       apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingOriginH2MaxStreamGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingOriginH2MaxStreamGetResponseResultJSON) RawJSON() string {
	return r.raw
}

// Value of the zone setting.
type ZoneSettingOriginH2MaxStreamGetResponseResultID string

const (
	ZoneSettingOriginH2MaxStreamGetResponseResultIDOriginH2MaxStreams ZoneSettingOriginH2MaxStreamGetResponseResultID = "origin_h2_max_streams"
)

func (r ZoneSettingOriginH2MaxStreamGetResponseResultID) IsKnown() bool {
	switch r {
	case ZoneSettingOriginH2MaxStreamGetResponseResultIDOriginH2MaxStreams:
		return true
	}
	return false
}

type ZoneSettingOriginH2MaxStreamUpdateResponse struct {
	Errors   []MessagesCacheRulesItem `json:"errors,required"`
	Messages []MessagesCacheRulesItem `json:"messages,required"`
	// Whether the API call was successful
	Success ZoneSettingOriginH2MaxStreamUpdateResponseSuccess `json:"success,required"`
	Result  ZoneSettingOriginH2MaxStreamUpdateResponseResult  `json:"result"`
	JSON    zoneSettingOriginH2MaxStreamUpdateResponseJSON    `json:"-"`
}

// zoneSettingOriginH2MaxStreamUpdateResponseJSON contains the JSON metadata for
// the struct [ZoneSettingOriginH2MaxStreamUpdateResponse]
type zoneSettingOriginH2MaxStreamUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingOriginH2MaxStreamUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingOriginH2MaxStreamUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ZoneSettingOriginH2MaxStreamUpdateResponseSuccess bool

const (
	ZoneSettingOriginH2MaxStreamUpdateResponseSuccessTrue ZoneSettingOriginH2MaxStreamUpdateResponseSuccess = true
)

func (r ZoneSettingOriginH2MaxStreamUpdateResponseSuccess) IsKnown() bool {
	switch r {
	case ZoneSettingOriginH2MaxStreamUpdateResponseSuccessTrue:
		return true
	}
	return false
}

type ZoneSettingOriginH2MaxStreamUpdateResponseResult struct {
	// Value of the zone setting.
	ID ZoneSettingOriginH2MaxStreamUpdateResponseResultID `json:"id,required"`
	// Whether the setting is editable
	Editable bool `json:"editable,required"`
	// The value of the feature
	Value string `json:"value,required"`
	// Last time this setting was modified.
	ModifiedOn time.Time                                            `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingOriginH2MaxStreamUpdateResponseResultJSON `json:"-"`
}

// zoneSettingOriginH2MaxStreamUpdateResponseResultJSON contains the JSON metadata
// for the struct [ZoneSettingOriginH2MaxStreamUpdateResponseResult]
type zoneSettingOriginH2MaxStreamUpdateResponseResultJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	Value       apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingOriginH2MaxStreamUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingOriginH2MaxStreamUpdateResponseResultJSON) RawJSON() string {
	return r.raw
}

// Value of the zone setting.
type ZoneSettingOriginH2MaxStreamUpdateResponseResultID string

const (
	ZoneSettingOriginH2MaxStreamUpdateResponseResultIDOriginH2MaxStreams ZoneSettingOriginH2MaxStreamUpdateResponseResultID = "origin_h2_max_streams"
)

func (r ZoneSettingOriginH2MaxStreamUpdateResponseResultID) IsKnown() bool {
	switch r {
	case ZoneSettingOriginH2MaxStreamUpdateResponseResultIDOriginH2MaxStreams:
		return true
	}
	return false
}

type ZoneSettingOriginH2MaxStreamUpdateParams struct {
	// Value of the Origin H2 Max Streams Setting.
	Value param.Field[int64] `json:"value,required"`
}

func (r ZoneSettingOriginH2MaxStreamUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
