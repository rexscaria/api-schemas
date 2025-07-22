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

type AccountAIRunCfBlackForestLabExecuteFlux1SchnellResponse struct {
	Result  AccountAIRunCfBlackForestLabExecuteFlux1SchnellResponseResult `json:"result"`
	Success bool                                                          `json:"success"`
	JSON    accountAIRunCfBlackForestLabExecuteFlux1SchnellResponseJSON   `json:"-"`
}

// accountAIRunCfBlackForestLabExecuteFlux1SchnellResponseJSON contains the JSON
// metadata for the struct
// [AccountAIRunCfBlackForestLabExecuteFlux1SchnellResponse]
type accountAIRunCfBlackForestLabExecuteFlux1SchnellResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIRunCfBlackForestLabExecuteFlux1SchnellResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunCfBlackForestLabExecuteFlux1SchnellResponseJSON) RawJSON() string {
	return r.raw
}

type AccountAIRunCfBlackForestLabExecuteFlux1SchnellResponseResult struct {
	// The generated image in Base64 format.
	Image string                                                            `json:"image"`
	JSON  accountAIRunCfBlackForestLabExecuteFlux1SchnellResponseResultJSON `json:"-"`
}

// accountAIRunCfBlackForestLabExecuteFlux1SchnellResponseResultJSON contains the
// JSON metadata for the struct
// [AccountAIRunCfBlackForestLabExecuteFlux1SchnellResponseResult]
type accountAIRunCfBlackForestLabExecuteFlux1SchnellResponseResultJSON struct {
	Image       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIRunCfBlackForestLabExecuteFlux1SchnellResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunCfBlackForestLabExecuteFlux1SchnellResponseResultJSON) RawJSON() string {
	return r.raw
}

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
