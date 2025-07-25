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

// AccountAIRunCfMyshellAIService contains methods and other services that help
// with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountAIRunCfMyshellAIService] method instead.
type AccountAIRunCfMyshellAIService struct {
	Options []option.RequestOption
}

// NewAccountAIRunCfMyshellAIService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountAIRunCfMyshellAIService(opts ...option.RequestOption) (r *AccountAIRunCfMyshellAIService) {
	r = &AccountAIRunCfMyshellAIService{}
	r.Options = opts
	return
}

// Execute @cf/myshell-ai/melotts model.
func (r *AccountAIRunCfMyshellAIService) ExecuteMelotts(ctx context.Context, accountID string, params AccountAIRunCfMyshellAIExecuteMelottsParams, opts ...option.RequestOption) (res *AccountAIRunCfMyshellAIExecuteMelottsResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai/run/@cf/myshell-ai/melotts", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type AccountAIRunCfMyshellAIExecuteMelottsResponse = interface{}

type AccountAIRunCfMyshellAIExecuteMelottsParams struct {
	// A text description of the audio you want to generate
	Prompt       param.Field[string] `json:"prompt,required"`
	QueueRequest param.Field[string] `query:"queueRequest"`
	// The speech language (e.g., 'en' for English, 'fr' for French). Defaults to 'en'
	// if not specified
	Lang param.Field[string] `json:"lang"`
}

func (r AccountAIRunCfMyshellAIExecuteMelottsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [AccountAIRunCfMyshellAIExecuteMelottsParams]'s query
// parameters as `url.Values`.
func (r AccountAIRunCfMyshellAIExecuteMelottsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
