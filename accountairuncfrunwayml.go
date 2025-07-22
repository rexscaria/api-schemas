// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/stainless-sdks/cf-rex-go/internal/apijson"
	"github.com/stainless-sdks/cf-rex-go/internal/apiquery"
	"github.com/stainless-sdks/cf-rex-go/internal/param"
	"github.com/stainless-sdks/cf-rex-go/internal/requestconfig"
	"github.com/stainless-sdks/cf-rex-go/option"
)

// AccountAIRunCfRunwaymlService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountAIRunCfRunwaymlService] method instead.
type AccountAIRunCfRunwaymlService struct {
	Options []option.RequestOption
}

// NewAccountAIRunCfRunwaymlService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountAIRunCfRunwaymlService(opts ...option.RequestOption) (r *AccountAIRunCfRunwaymlService) {
	r = &AccountAIRunCfRunwaymlService{}
	r.Options = opts
	return
}

// Execute @cf/runwayml/stable-diffusion-v1-5-img2img model.
func (r *AccountAIRunCfRunwaymlService) ExecuteStableDiffusionV1_5Img2img(ctx context.Context, accountID string, params AccountAIRunCfRunwaymlExecuteStableDiffusionV1_5Img2imgParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "image/png")}, opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai/run/@cf/runwayml/stable-diffusion-v1-5-img2img", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Execute @cf/runwayml/stable-diffusion-v1-5-inpainting model.
func (r *AccountAIRunCfRunwaymlService) ExecuteStableDiffusionV1_5Inpainting(ctx context.Context, accountID string, params AccountAIRunCfRunwaymlExecuteStableDiffusionV1_5InpaintingParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "image/png")}, opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai/run/@cf/runwayml/stable-diffusion-v1-5-inpainting", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type AccountAIRunCfRunwaymlExecuteStableDiffusionV1_5Img2imgParams struct {
	// A text description of the image you want to generate
	Prompt       param.Field[string] `json:"prompt,required"`
	QueueRequest param.Field[string] `query:"queueRequest"`
	// Controls how closely the generated image should adhere to the prompt; higher
	// values make the image more aligned with the prompt
	Guidance param.Field[float64] `json:"guidance"`
	// The height of the generated image in pixels
	Height param.Field[int64] `json:"height"`
	// For use with img2img tasks. An array of integers that represent the image data
	// constrained to 8-bit unsigned integer values
	Image param.Field[[]float64] `json:"image"`
	// For use with img2img tasks. A base64-encoded string of the input image
	ImageB64 param.Field[string] `json:"image_b64"`
	// An array representing An array of integers that represent mask image data for
	// inpainting constrained to 8-bit unsigned integer values
	Mask param.Field[[]float64] `json:"mask"`
	// Text describing elements to avoid in the generated image
	NegativePrompt param.Field[string] `json:"negative_prompt"`
	// The number of diffusion steps; higher values can improve quality but take longer
	NumSteps param.Field[int64] `json:"num_steps"`
	// Random seed for reproducibility of the image generation
	Seed param.Field[int64] `json:"seed"`
	// A value between 0 and 1 indicating how strongly to apply the transformation
	// during img2img tasks; lower values make the output closer to the input image
	Strength param.Field[float64] `json:"strength"`
	// The width of the generated image in pixels
	Width param.Field[int64] `json:"width"`
}

func (r AccountAIRunCfRunwaymlExecuteStableDiffusionV1_5Img2imgParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes
// [AccountAIRunCfRunwaymlExecuteStableDiffusionV1_5Img2imgParams]'s query
// parameters as `url.Values`.
func (r AccountAIRunCfRunwaymlExecuteStableDiffusionV1_5Img2imgParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountAIRunCfRunwaymlExecuteStableDiffusionV1_5InpaintingParams struct {
	// A text description of the image you want to generate
	Prompt       param.Field[string] `json:"prompt,required"`
	QueueRequest param.Field[string] `query:"queueRequest"`
	// Controls how closely the generated image should adhere to the prompt; higher
	// values make the image more aligned with the prompt
	Guidance param.Field[float64] `json:"guidance"`
	// The height of the generated image in pixels
	Height param.Field[int64] `json:"height"`
	// For use with img2img tasks. An array of integers that represent the image data
	// constrained to 8-bit unsigned integer values
	Image param.Field[[]float64] `json:"image"`
	// For use with img2img tasks. A base64-encoded string of the input image
	ImageB64 param.Field[string] `json:"image_b64"`
	// An array representing An array of integers that represent mask image data for
	// inpainting constrained to 8-bit unsigned integer values
	Mask param.Field[[]float64] `json:"mask"`
	// Text describing elements to avoid in the generated image
	NegativePrompt param.Field[string] `json:"negative_prompt"`
	// The number of diffusion steps; higher values can improve quality but take longer
	NumSteps param.Field[int64] `json:"num_steps"`
	// Random seed for reproducibility of the image generation
	Seed param.Field[int64] `json:"seed"`
	// A value between 0 and 1 indicating how strongly to apply the transformation
	// during img2img tasks; lower values make the output closer to the input image
	Strength param.Field[float64] `json:"strength"`
	// The width of the generated image in pixels
	Width param.Field[int64] `json:"width"`
}

func (r AccountAIRunCfRunwaymlExecuteStableDiffusionV1_5InpaintingParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes
// [AccountAIRunCfRunwaymlExecuteStableDiffusionV1_5InpaintingParams]'s query
// parameters as `url.Values`.
func (r AccountAIRunCfRunwaymlExecuteStableDiffusionV1_5InpaintingParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
