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

// AccountAIRunCfBlackForestLabService contains methods and other services that
// help with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountAIRunCfBlackForestLabService] method instead.
type AccountAIRunCfBlackForestLabService struct {
	Options []option.RequestOption
}

// NewAccountAIRunCfBlackForestLabService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountAIRunCfBlackForestLabService(opts ...option.RequestOption) (r *AccountAIRunCfBlackForestLabService) {
	r = &AccountAIRunCfBlackForestLabService{}
	r.Options = opts
	return
}

// Execute @cf/black-forest-labs/flux-1-schnell model.
func (r *AccountAIRunCfBlackForestLabService) ExecuteFlux1Schnell(ctx context.Context, accountID string, params AccountAIRunCfBlackForestLabExecuteFlux1SchnellParams, opts ...option.RequestOption) (res *AccountAIRunCfBlackForestLabExecuteFlux1SchnellResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai/run/@cf/black-forest-labs/flux-1-schnell", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type AccountAIRunCfBlackForestLabExecuteFlux1SchnellResponse = interface{}

type AccountAIRunCfBlackForestLabExecuteFlux1SchnellParams struct {
	// A text description of the image you want to generate.
	Prompt       param.Field[string] `json:"prompt,required"`
	QueueRequest param.Field[string] `query:"queueRequest"`
	// The number of diffusion steps; higher values can improve quality but take
	// longer.
	Steps param.Field[int64] `json:"steps"`
}

func (r AccountAIRunCfBlackForestLabExecuteFlux1SchnellParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [AccountAIRunCfBlackForestLabExecuteFlux1SchnellParams]'s
// query parameters as `url.Values`.
func (r AccountAIRunCfBlackForestLabExecuteFlux1SchnellParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
