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

// ZoneSettingFontService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneSettingFontService] method instead.
type ZoneSettingFontService struct {
	Options []option.RequestOption
}

// NewZoneSettingFontService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewZoneSettingFontService(opts ...option.RequestOption) (r *ZoneSettingFontService) {
	r = &ZoneSettingFontService{}
	r.Options = opts
	return
}

// Enhance your website's font delivery with Cloudflare Fonts. Deliver Google
// Hosted fonts from your own domain, boost performance, and enhance user privacy.
// Refer to the Cloudflare Fonts documentation for more information.
func (r *ZoneSettingFontService) Get(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *ZoneSettingFontGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/settings/fonts", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Enhance your website's font delivery with Cloudflare Fonts. Deliver Google
// Hosted fonts from your own domain, boost performance, and enhance user privacy.
// Refer to the Cloudflare Fonts documentation for more information.
func (r *ZoneSettingFontService) Update(ctx context.Context, zoneID string, body ZoneSettingFontUpdateParams, opts ...option.RequestOption) (res *ZoneSettingFontUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/settings/fonts", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

type SpeedBase struct {
	// Identifier of the zone setting.
	ID string `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SpeedBaseEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Current value of the zone setting.
	Value SpeedBaseValue `json:"value"`
	JSON  speedBaseJSON  `json:"-"`
}

// speedBaseJSON contains the JSON metadata for the struct [SpeedBase]
type speedBaseJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpeedBase) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r speedBaseJSON) RawJSON() string {
	return r.raw
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SpeedBaseEditable bool

const (
	SpeedBaseEditableTrue  SpeedBaseEditable = true
	SpeedBaseEditableFalse SpeedBaseEditable = false
)

func (r SpeedBaseEditable) IsKnown() bool {
	switch r {
	case SpeedBaseEditableTrue, SpeedBaseEditableFalse:
		return true
	}
	return false
}

// Current value of the zone setting.
type SpeedBaseValue string

const (
	SpeedBaseValueOn  SpeedBaseValue = "on"
	SpeedBaseValueOff SpeedBaseValue = "off"
)

func (r SpeedBaseValue) IsKnown() bool {
	switch r {
	case SpeedBaseValueOn, SpeedBaseValueOff:
		return true
	}
	return false
}

// Enhance your website's font delivery with Cloudflare Fonts. Deliver Google
// Hosted fonts from your own domain, boost performance, and enhance user privacy.
// Refer to the Cloudflare Fonts documentation for more information.
type SpeedCloudflareFonts struct {
	// ID of the zone setting.
	ID SpeedCloudflareFontsID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable SpeedCloudflareFontsEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Current value of the zone setting.
	Value SpeedCloudflareFontsValue `json:"value"`
	JSON  speedCloudflareFontsJSON  `json:"-"`
}

// speedCloudflareFontsJSON contains the JSON metadata for the struct
// [SpeedCloudflareFonts]
type speedCloudflareFontsJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpeedCloudflareFonts) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r speedCloudflareFontsJSON) RawJSON() string {
	return r.raw
}

// ID of the zone setting.
type SpeedCloudflareFontsID string

const (
	SpeedCloudflareFontsIDFonts SpeedCloudflareFontsID = "fonts"
)

func (r SpeedCloudflareFontsID) IsKnown() bool {
	switch r {
	case SpeedCloudflareFontsIDFonts:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type SpeedCloudflareFontsEditable bool

const (
	SpeedCloudflareFontsEditableTrue  SpeedCloudflareFontsEditable = true
	SpeedCloudflareFontsEditableFalse SpeedCloudflareFontsEditable = false
)

func (r SpeedCloudflareFontsEditable) IsKnown() bool {
	switch r {
	case SpeedCloudflareFontsEditableTrue, SpeedCloudflareFontsEditableFalse:
		return true
	}
	return false
}

// Current value of the zone setting.
type SpeedCloudflareFontsValue string

const (
	SpeedCloudflareFontsValueOn  SpeedCloudflareFontsValue = "on"
	SpeedCloudflareFontsValueOff SpeedCloudflareFontsValue = "off"
)

func (r SpeedCloudflareFontsValue) IsKnown() bool {
	switch r {
	case SpeedCloudflareFontsValueOn, SpeedCloudflareFontsValueOff:
		return true
	}
	return false
}

type SpeedMessagesItems struct {
	Code             int64                    `json:"code,required"`
	Message          string                   `json:"message,required"`
	DocumentationURL string                   `json:"documentation_url"`
	Source           SpeedMessagesItemsSource `json:"source"`
	JSON             speedMessagesItemsJSON   `json:"-"`
}

// speedMessagesItemsJSON contains the JSON metadata for the struct
// [SpeedMessagesItems]
type speedMessagesItemsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SpeedMessagesItems) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r speedMessagesItemsJSON) RawJSON() string {
	return r.raw
}

type SpeedMessagesItemsSource struct {
	Pointer string                       `json:"pointer"`
	JSON    speedMessagesItemsSourceJSON `json:"-"`
}

// speedMessagesItemsSourceJSON contains the JSON metadata for the struct
// [SpeedMessagesItemsSource]
type speedMessagesItemsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpeedMessagesItemsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r speedMessagesItemsSourceJSON) RawJSON() string {
	return r.raw
}

type ZoneSettingFontGetResponse struct {
	Errors   []SpeedMessagesItems `json:"errors,required"`
	Messages []SpeedMessagesItems `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Enhance your website's font delivery with Cloudflare Fonts. Deliver Google
	// Hosted fonts from your own domain, boost performance, and enhance user privacy.
	// Refer to the Cloudflare Fonts documentation for more information.
	Result SpeedCloudflareFonts           `json:"result"`
	JSON   zoneSettingFontGetResponseJSON `json:"-"`
}

// zoneSettingFontGetResponseJSON contains the JSON metadata for the struct
// [ZoneSettingFontGetResponse]
type zoneSettingFontGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingFontGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingFontGetResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneSettingFontUpdateResponse struct {
	Errors   []SpeedMessagesItems `json:"errors,required"`
	Messages []SpeedMessagesItems `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Enhance your website's font delivery with Cloudflare Fonts. Deliver Google
	// Hosted fonts from your own domain, boost performance, and enhance user privacy.
	// Refer to the Cloudflare Fonts documentation for more information.
	Result SpeedCloudflareFonts              `json:"result"`
	JSON   zoneSettingFontUpdateResponseJSON `json:"-"`
}

// zoneSettingFontUpdateResponseJSON contains the JSON metadata for the struct
// [ZoneSettingFontUpdateResponse]
type zoneSettingFontUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingFontUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSettingFontUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneSettingFontUpdateParams struct {
	// Whether the feature is enabled or disabled.
	Value param.Field[SpeedCloudflareFontsValue] `json:"value,required"`
}

func (r ZoneSettingFontUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
