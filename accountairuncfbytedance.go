// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/apiquery"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
)

// AccountAIRunCfBytedanceService contains methods and other services that help
// with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountAIRunCfBytedanceService] method instead.
type AccountAIRunCfBytedanceService struct {
	Options []option.RequestOption
}

// NewAccountAIRunCfBytedanceService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountAIRunCfBytedanceService(opts ...option.RequestOption) (r *AccountAIRunCfBytedanceService) {
	r = &AccountAIRunCfBytedanceService{}
	r.Options = opts
	return
}

// Execute @cf/bytedance/stable-diffusion-xl-lightning model.
func (r *AccountAIRunCfBytedanceService) ExecuteStableDiffusionXlLightning(ctx context.Context, accountID string, params AccountAIRunCfBytedanceExecuteStableDiffusionXlLightningParams, opts ...option.RequestOption) (res *AccountAIRunCfBytedanceExecuteStableDiffusionXlLightningResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai/run/@cf/bytedance/stable-diffusion-xl-lightning", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type AccountAIRunCfBytedanceExecuteStableDiffusionXlLightningResponse = interface{}

type AccountAIRunCfBytedanceExecuteStableDiffusionXlLightningParams struct {
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

func (r AccountAIRunCfBytedanceExecuteStableDiffusionXlLightningParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes
// [AccountAIRunCfBytedanceExecuteStableDiffusionXlLightningParams]'s query
// parameters as `url.Values`.
func (r AccountAIRunCfBytedanceExecuteStableDiffusionXlLightningParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
