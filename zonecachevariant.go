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

// ZoneCacheVariantService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneCacheVariantService] method instead.
type ZoneCacheVariantService struct {
	Options []option.RequestOption
}

// NewZoneCacheVariantService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneCacheVariantService(opts ...option.RequestOption) (r *ZoneCacheVariantService) {
	r = &ZoneCacheVariantService{}
	r.Options = opts
	return
}

// Variant support enables caching variants of images with certain file extensions
// in addition to the original. This only applies when the origin server sends the
// 'Vary: Accept' response header. If the origin server sends 'Vary: Accept' but
// does not serve the variant requested, the response will not be cached. This will
// be indicated with BYPASS cache status in the response headers.
func (r *ZoneCacheVariantService) Get(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *ZoneCacheVariantGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/cache/variants", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Variant support enables caching variants of images with certain file extensions
// in addition to the original. This only applies when the origin server sends the
// 'Vary: Accept' response header. If the origin server sends 'Vary: Accept' but
// does not serve the variant requested, the response will not be cached. This will
// be indicated with BYPASS cache status in the response headers.
func (r *ZoneCacheVariantService) Update(ctx context.Context, zoneID string, body ZoneCacheVariantUpdateParams, opts ...option.RequestOption) (res *ZoneCacheVariantUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/cache/variants", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Variant support enables caching variants of images with certain file extensions
// in addition to the original. This only applies when the origin server sends the
// 'Vary: Accept' response header. If the origin server sends 'Vary: Accept' but
// does not serve the variant requested, the response will not be cached. This will
// be indicated with BYPASS cache status in the response headers.
func (r *ZoneCacheVariantService) Delete(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *ZoneCacheVariantDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/cache/variants", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Value of the zone setting.
type VariantsValue struct {
	// List of strings with the MIME types of all the variants that should be served
	// for avif.
	Avif []string `json:"avif"`
	// List of strings with the MIME types of all the variants that should be served
	// for bmp.
	Bmp []string `json:"bmp"`
	// List of strings with the MIME types of all the variants that should be served
	// for gif.
	Gif []string `json:"gif"`
	// List of strings with the MIME types of all the variants that should be served
	// for jp2.
	Jp2 []string `json:"jp2"`
	// List of strings with the MIME types of all the variants that should be served
	// for jpeg.
	Jpeg []string `json:"jpeg"`
	// List of strings with the MIME types of all the variants that should be served
	// for jpg.
	Jpg []string `json:"jpg"`
	// List of strings with the MIME types of all the variants that should be served
	// for jpg2.
	Jpg2 []string `json:"jpg2"`
	// List of strings with the MIME types of all the variants that should be served
	// for png.
	Png []string `json:"png"`
	// List of strings with the MIME types of all the variants that should be served
	// for tif.
	Tif []string `json:"tif"`
	// List of strings with the MIME types of all the variants that should be served
	// for tiff.
	Tiff []string `json:"tiff"`
	// List of strings with the MIME types of all the variants that should be served
	// for webp.
	Webp []string          `json:"webp"`
	JSON variantsValueJSON `json:"-"`
}

// variantsValueJSON contains the JSON metadata for the struct [VariantsValue]
type variantsValueJSON struct {
	Avif        apijson.Field
	Bmp         apijson.Field
	Gif         apijson.Field
	Jp2         apijson.Field
	Jpeg        apijson.Field
	Jpg         apijson.Field
	Jpg2        apijson.Field
	Png         apijson.Field
	Tif         apijson.Field
	Tiff        apijson.Field
	Webp        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VariantsValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r variantsValueJSON) RawJSON() string {
	return r.raw
}

// Value of the zone setting.
type VariantsValueParam struct {
	// List of strings with the MIME types of all the variants that should be served
	// for avif.
	Avif param.Field[[]string] `json:"avif"`
	// List of strings with the MIME types of all the variants that should be served
	// for bmp.
	Bmp param.Field[[]string] `json:"bmp"`
	// List of strings with the MIME types of all the variants that should be served
	// for gif.
	Gif param.Field[[]string] `json:"gif"`
	// List of strings with the MIME types of all the variants that should be served
	// for jp2.
	Jp2 param.Field[[]string] `json:"jp2"`
	// List of strings with the MIME types of all the variants that should be served
	// for jpeg.
	Jpeg param.Field[[]string] `json:"jpeg"`
	// List of strings with the MIME types of all the variants that should be served
	// for jpg.
	Jpg param.Field[[]string] `json:"jpg"`
	// List of strings with the MIME types of all the variants that should be served
	// for jpg2.
	Jpg2 param.Field[[]string] `json:"jpg2"`
	// List of strings with the MIME types of all the variants that should be served
	// for png.
	Png param.Field[[]string] `json:"png"`
	// List of strings with the MIME types of all the variants that should be served
	// for tif.
	Tif param.Field[[]string] `json:"tif"`
	// List of strings with the MIME types of all the variants that should be served
	// for tiff.
	Tiff param.Field[[]string] `json:"tiff"`
	// List of strings with the MIME types of all the variants that should be served
	// for webp.
	Webp param.Field[[]string] `json:"webp"`
}

func (r VariantsValueParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneCacheVariantGetResponse struct {
	Errors   []MessagesCacheRulesItem `json:"errors,required"`
	Messages []MessagesCacheRulesItem `json:"messages,required"`
	// Whether the API call was successful
	Success ZoneCacheVariantGetResponseSuccess `json:"success,required"`
	Result  ZoneCacheVariantGetResponseResult  `json:"result"`
	JSON    zoneCacheVariantGetResponseJSON    `json:"-"`
}

// zoneCacheVariantGetResponseJSON contains the JSON metadata for the struct
// [ZoneCacheVariantGetResponse]
type zoneCacheVariantGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCacheVariantGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneCacheVariantGetResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ZoneCacheVariantGetResponseSuccess bool

const (
	ZoneCacheVariantGetResponseSuccessTrue ZoneCacheVariantGetResponseSuccess = true
)

func (r ZoneCacheVariantGetResponseSuccess) IsKnown() bool {
	switch r {
	case ZoneCacheVariantGetResponseSuccessTrue:
		return true
	}
	return false
}

type ZoneCacheVariantGetResponseResult struct {
	// ID of the zone setting.
	ID ZoneCacheVariantGetResponseResultID `json:"id,required"`
	// Whether the setting is editable
	Editable bool `json:"editable,required"`
	// The value of the feature
	Value VariantsValue `json:"value,required"`
	// Last time this setting was modified.
	ModifiedOn time.Time                             `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneCacheVariantGetResponseResultJSON `json:"-"`
}

// zoneCacheVariantGetResponseResultJSON contains the JSON metadata for the struct
// [ZoneCacheVariantGetResponseResult]
type zoneCacheVariantGetResponseResultJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	Value       apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCacheVariantGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneCacheVariantGetResponseResultJSON) RawJSON() string {
	return r.raw
}

// ID of the zone setting.
type ZoneCacheVariantGetResponseResultID string

const (
	ZoneCacheVariantGetResponseResultIDVariants ZoneCacheVariantGetResponseResultID = "variants"
)

func (r ZoneCacheVariantGetResponseResultID) IsKnown() bool {
	switch r {
	case ZoneCacheVariantGetResponseResultIDVariants:
		return true
	}
	return false
}

type ZoneCacheVariantUpdateResponse struct {
	Errors   []MessagesCacheRulesItem `json:"errors,required"`
	Messages []MessagesCacheRulesItem `json:"messages,required"`
	// Whether the API call was successful
	Success ZoneCacheVariantUpdateResponseSuccess `json:"success,required"`
	Result  ZoneCacheVariantUpdateResponseResult  `json:"result"`
	JSON    zoneCacheVariantUpdateResponseJSON    `json:"-"`
}

// zoneCacheVariantUpdateResponseJSON contains the JSON metadata for the struct
// [ZoneCacheVariantUpdateResponse]
type zoneCacheVariantUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCacheVariantUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneCacheVariantUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ZoneCacheVariantUpdateResponseSuccess bool

const (
	ZoneCacheVariantUpdateResponseSuccessTrue ZoneCacheVariantUpdateResponseSuccess = true
)

func (r ZoneCacheVariantUpdateResponseSuccess) IsKnown() bool {
	switch r {
	case ZoneCacheVariantUpdateResponseSuccessTrue:
		return true
	}
	return false
}

type ZoneCacheVariantUpdateResponseResult struct {
	// ID of the zone setting.
	ID ZoneCacheVariantUpdateResponseResultID `json:"id,required"`
	// Whether the setting is editable
	Editable bool `json:"editable,required"`
	// The value of the feature
	Value VariantsValue `json:"value,required"`
	// Last time this setting was modified.
	ModifiedOn time.Time                                `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneCacheVariantUpdateResponseResultJSON `json:"-"`
}

// zoneCacheVariantUpdateResponseResultJSON contains the JSON metadata for the
// struct [ZoneCacheVariantUpdateResponseResult]
type zoneCacheVariantUpdateResponseResultJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	Value       apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCacheVariantUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneCacheVariantUpdateResponseResultJSON) RawJSON() string {
	return r.raw
}

// ID of the zone setting.
type ZoneCacheVariantUpdateResponseResultID string

const (
	ZoneCacheVariantUpdateResponseResultIDVariants ZoneCacheVariantUpdateResponseResultID = "variants"
)

func (r ZoneCacheVariantUpdateResponseResultID) IsKnown() bool {
	switch r {
	case ZoneCacheVariantUpdateResponseResultIDVariants:
		return true
	}
	return false
}

type ZoneCacheVariantDeleteResponse struct {
	Errors   []MessagesCacheRulesItem `json:"errors,required"`
	Messages []MessagesCacheRulesItem `json:"messages,required"`
	// Whether the API call was successful
	Success ZoneCacheVariantDeleteResponseSuccess `json:"success,required"`
	Result  ZoneCacheVariantDeleteResponseResult  `json:"result"`
	JSON    zoneCacheVariantDeleteResponseJSON    `json:"-"`
}

// zoneCacheVariantDeleteResponseJSON contains the JSON metadata for the struct
// [ZoneCacheVariantDeleteResponse]
type zoneCacheVariantDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCacheVariantDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneCacheVariantDeleteResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ZoneCacheVariantDeleteResponseSuccess bool

const (
	ZoneCacheVariantDeleteResponseSuccessTrue ZoneCacheVariantDeleteResponseSuccess = true
)

func (r ZoneCacheVariantDeleteResponseSuccess) IsKnown() bool {
	switch r {
	case ZoneCacheVariantDeleteResponseSuccessTrue:
		return true
	}
	return false
}

type ZoneCacheVariantDeleteResponseResult struct {
	// ID of the zone setting.
	ID ZoneCacheVariantDeleteResponseResultID `json:"id,required"`
	// Whether the setting is editable
	Editable bool `json:"editable,required"`
	// Last time this setting was modified.
	ModifiedOn time.Time                                `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneCacheVariantDeleteResponseResultJSON `json:"-"`
}

// zoneCacheVariantDeleteResponseResultJSON contains the JSON metadata for the
// struct [ZoneCacheVariantDeleteResponseResult]
type zoneCacheVariantDeleteResponseResultJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCacheVariantDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneCacheVariantDeleteResponseResultJSON) RawJSON() string {
	return r.raw
}

// ID of the zone setting.
type ZoneCacheVariantDeleteResponseResultID string

const (
	ZoneCacheVariantDeleteResponseResultIDVariants ZoneCacheVariantDeleteResponseResultID = "variants"
)

func (r ZoneCacheVariantDeleteResponseResultID) IsKnown() bool {
	switch r {
	case ZoneCacheVariantDeleteResponseResultIDVariants:
		return true
	}
	return false
}

type ZoneCacheVariantUpdateParams struct {
	// Value of the zone setting.
	Value param.Field[VariantsValueParam] `json:"value,required"`
}

func (r ZoneCacheVariantUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
