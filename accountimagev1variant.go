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

// AccountImageV1VariantService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountImageV1VariantService] method instead.
type AccountImageV1VariantService struct {
	Options []option.RequestOption
}

// NewAccountImageV1VariantService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountImageV1VariantService(opts ...option.RequestOption) (r *AccountImageV1VariantService) {
	r = &AccountImageV1VariantService{}
	r.Options = opts
	return
}

// Specify variants that allow you to resize images for different use cases.
func (r *AccountImageV1VariantService) New(ctx context.Context, accountID string, body AccountImageV1VariantNewParams, opts ...option.RequestOption) (res *ImageVariantSimpleResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/images/v1/variants", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Fetch details for a single variant.
func (r *AccountImageV1VariantService) Get(ctx context.Context, accountID string, variantID string, opts ...option.RequestOption) (res *ImageVariantSimpleResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if variantID == "" {
		err = errors.New("missing required variant_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/images/v1/variants/%s", accountID, variantID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updating a variant purges the cache for all images associated with the variant.
func (r *AccountImageV1VariantService) Update(ctx context.Context, accountID string, variantID string, body AccountImageV1VariantUpdateParams, opts ...option.RequestOption) (res *ImageVariantSimpleResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if variantID == "" {
		err = errors.New("missing required variant_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/images/v1/variants/%s", accountID, variantID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Lists existing variants.
func (r *AccountImageV1VariantService) List(ctx context.Context, accountID string, opts ...option.RequestOption) (res *AccountImageV1VariantListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/images/v1/variants", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Deleting a variant purges the cache for all images associated with the variant.
func (r *AccountImageV1VariantService) Delete(ctx context.Context, accountID string, variantID string, body AccountImageV1VariantDeleteParams, opts ...option.RequestOption) (res *DeletedImagesResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if variantID == "" {
		err = errors.New("missing required variant_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/images/v1/variants/%s", accountID, variantID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

type ImageVariantDefinition struct {
	ID string `json:"id,required"`
	// Allows you to define image resizing sizes for different use cases.
	Options ImageVariantOptions `json:"options,required"`
	// Indicates whether the variant can access an image without a signature,
	// regardless of image access control.
	NeverRequireSignedURLs bool                       `json:"neverRequireSignedURLs"`
	JSON                   imageVariantDefinitionJSON `json:"-"`
}

// imageVariantDefinitionJSON contains the JSON metadata for the struct
// [ImageVariantDefinition]
type imageVariantDefinitionJSON struct {
	ID                     apijson.Field
	Options                apijson.Field
	NeverRequireSignedURLs apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ImageVariantDefinition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r imageVariantDefinitionJSON) RawJSON() string {
	return r.raw
}

type ImageVariantDefinitionParam struct {
	ID param.Field[string] `json:"id,required"`
	// Allows you to define image resizing sizes for different use cases.
	Options param.Field[ImageVariantOptionsParam] `json:"options,required"`
	// Indicates whether the variant can access an image without a signature,
	// regardless of image access control.
	NeverRequireSignedURLs param.Field[bool] `json:"neverRequireSignedURLs"`
}

func (r ImageVariantDefinitionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Allows you to define image resizing sizes for different use cases.
type ImageVariantOptions struct {
	// The fit property describes how the width and height dimensions should be
	// interpreted.
	Fit ImageVariantOptionsFit `json:"fit,required"`
	// Maximum height in image pixels.
	Height float64 `json:"height,required"`
	// What EXIF data should be preserved in the output image.
	Metadata ImageVariantOptionsMetadata `json:"metadata,required"`
	// Maximum width in image pixels.
	Width float64                 `json:"width,required"`
	JSON  imageVariantOptionsJSON `json:"-"`
}

// imageVariantOptionsJSON contains the JSON metadata for the struct
// [ImageVariantOptions]
type imageVariantOptionsJSON struct {
	Fit         apijson.Field
	Height      apijson.Field
	Metadata    apijson.Field
	Width       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ImageVariantOptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r imageVariantOptionsJSON) RawJSON() string {
	return r.raw
}

// The fit property describes how the width and height dimensions should be
// interpreted.
type ImageVariantOptionsFit string

const (
	ImageVariantOptionsFitScaleDown ImageVariantOptionsFit = "scale-down"
	ImageVariantOptionsFitContain   ImageVariantOptionsFit = "contain"
	ImageVariantOptionsFitCover     ImageVariantOptionsFit = "cover"
	ImageVariantOptionsFitCrop      ImageVariantOptionsFit = "crop"
	ImageVariantOptionsFitPad       ImageVariantOptionsFit = "pad"
)

func (r ImageVariantOptionsFit) IsKnown() bool {
	switch r {
	case ImageVariantOptionsFitScaleDown, ImageVariantOptionsFitContain, ImageVariantOptionsFitCover, ImageVariantOptionsFitCrop, ImageVariantOptionsFitPad:
		return true
	}
	return false
}

// What EXIF data should be preserved in the output image.
type ImageVariantOptionsMetadata string

const (
	ImageVariantOptionsMetadataKeep      ImageVariantOptionsMetadata = "keep"
	ImageVariantOptionsMetadataCopyright ImageVariantOptionsMetadata = "copyright"
	ImageVariantOptionsMetadataNone      ImageVariantOptionsMetadata = "none"
)

func (r ImageVariantOptionsMetadata) IsKnown() bool {
	switch r {
	case ImageVariantOptionsMetadataKeep, ImageVariantOptionsMetadataCopyright, ImageVariantOptionsMetadataNone:
		return true
	}
	return false
}

// Allows you to define image resizing sizes for different use cases.
type ImageVariantOptionsParam struct {
	// The fit property describes how the width and height dimensions should be
	// interpreted.
	Fit param.Field[ImageVariantOptionsFit] `json:"fit,required"`
	// Maximum height in image pixels.
	Height param.Field[float64] `json:"height,required"`
	// What EXIF data should be preserved in the output image.
	Metadata param.Field[ImageVariantOptionsMetadata] `json:"metadata,required"`
	// Maximum width in image pixels.
	Width param.Field[float64] `json:"width,required"`
}

func (r ImageVariantOptionsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ImageVariantSimpleResponse struct {
	Result ImageVariantSimpleResponseResult `json:"result"`
	JSON   imageVariantSimpleResponseJSON   `json:"-"`
	APIResponseSingleImages
}

// imageVariantSimpleResponseJSON contains the JSON metadata for the struct
// [ImageVariantSimpleResponse]
type imageVariantSimpleResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ImageVariantSimpleResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r imageVariantSimpleResponseJSON) RawJSON() string {
	return r.raw
}

type ImageVariantSimpleResponseResult struct {
	Variant ImageVariantDefinition               `json:"variant"`
	JSON    imageVariantSimpleResponseResultJSON `json:"-"`
}

// imageVariantSimpleResponseResultJSON contains the JSON metadata for the struct
// [ImageVariantSimpleResponseResult]
type imageVariantSimpleResponseResultJSON struct {
	Variant     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ImageVariantSimpleResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r imageVariantSimpleResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountImageV1VariantListResponse struct {
	Result AccountImageV1VariantListResponseResult `json:"result"`
	JSON   accountImageV1VariantListResponseJSON   `json:"-"`
	APIResponseImages
}

// accountImageV1VariantListResponseJSON contains the JSON metadata for the struct
// [AccountImageV1VariantListResponse]
type accountImageV1VariantListResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountImageV1VariantListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountImageV1VariantListResponseJSON) RawJSON() string {
	return r.raw
}

type AccountImageV1VariantListResponseResult struct {
	Variants AccountImageV1VariantListResponseResultVariants `json:"variants"`
	JSON     accountImageV1VariantListResponseResultJSON     `json:"-"`
}

// accountImageV1VariantListResponseResultJSON contains the JSON metadata for the
// struct [AccountImageV1VariantListResponseResult]
type accountImageV1VariantListResponseResultJSON struct {
	Variants    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountImageV1VariantListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountImageV1VariantListResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountImageV1VariantListResponseResultVariants struct {
	Hero AccountImageV1VariantListResponseResultVariantsHero `json:"hero"`
	JSON accountImageV1VariantListResponseResultVariantsJSON `json:"-"`
}

// accountImageV1VariantListResponseResultVariantsJSON contains the JSON metadata
// for the struct [AccountImageV1VariantListResponseResultVariants]
type accountImageV1VariantListResponseResultVariantsJSON struct {
	Hero        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountImageV1VariantListResponseResultVariants) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountImageV1VariantListResponseResultVariantsJSON) RawJSON() string {
	return r.raw
}

type AccountImageV1VariantListResponseResultVariantsHero struct {
	ID string `json:"id,required"`
	// Allows you to define image resizing sizes for different use cases.
	Options ImageVariantOptions `json:"options,required"`
	// Indicates whether the variant can access an image without a signature,
	// regardless of image access control.
	NeverRequireSignedURLs bool                                                    `json:"neverRequireSignedURLs"`
	JSON                   accountImageV1VariantListResponseResultVariantsHeroJSON `json:"-"`
}

// accountImageV1VariantListResponseResultVariantsHeroJSON contains the JSON
// metadata for the struct [AccountImageV1VariantListResponseResultVariantsHero]
type accountImageV1VariantListResponseResultVariantsHeroJSON struct {
	ID                     apijson.Field
	Options                apijson.Field
	NeverRequireSignedURLs apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccountImageV1VariantListResponseResultVariantsHero) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountImageV1VariantListResponseResultVariantsHeroJSON) RawJSON() string {
	return r.raw
}

type AccountImageV1VariantNewParams struct {
	ImageVariantDefinition ImageVariantDefinitionParam `json:"image_variant_definition,required"`
}

func (r AccountImageV1VariantNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.ImageVariantDefinition)
}

type AccountImageV1VariantUpdateParams struct {
	// Allows you to define image resizing sizes for different use cases.
	Options param.Field[ImageVariantOptionsParam] `json:"options,required"`
	// Indicates whether the variant can access an image without a signature,
	// regardless of image access control.
	NeverRequireSignedURLs param.Field[bool] `json:"neverRequireSignedURLs"`
}

func (r AccountImageV1VariantUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountImageV1VariantDeleteParams struct {
	Body interface{} `json:"body,required"`
}

func (r AccountImageV1VariantDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}
