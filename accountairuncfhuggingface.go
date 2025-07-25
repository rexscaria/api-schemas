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

// AccountAIRunCfHuggingfaceService contains methods and other services that help
// with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountAIRunCfHuggingfaceService] method instead.
type AccountAIRunCfHuggingfaceService struct {
	Options []option.RequestOption
}

// NewAccountAIRunCfHuggingfaceService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountAIRunCfHuggingfaceService(opts ...option.RequestOption) (r *AccountAIRunCfHuggingfaceService) {
	r = &AccountAIRunCfHuggingfaceService{}
	r.Options = opts
	return
}

// Execute @cf/huggingface/distilbert-sst-2-int8 model.
func (r *AccountAIRunCfHuggingfaceService) ExecuteDistilbertSst2Int8(ctx context.Context, accountID string, params AccountAIRunCfHuggingfaceExecuteDistilbertSst2Int8Params, opts ...option.RequestOption) (res *AccountAIRunCfHuggingfaceExecuteDistilbertSst2Int8Response, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai/run/@cf/huggingface/distilbert-sst-2-int8", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type AccountAIRunCfHuggingfaceExecuteDistilbertSst2Int8Response = interface{}

type AccountAIRunCfHuggingfaceExecuteDistilbertSst2Int8Params struct {
	// The text that you want to classify
	Text         param.Field[string] `json:"text,required"`
	QueueRequest param.Field[string] `query:"queueRequest"`
}

func (r AccountAIRunCfHuggingfaceExecuteDistilbertSst2Int8Params) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [AccountAIRunCfHuggingfaceExecuteDistilbertSst2Int8Params]'s
// query parameters as `url.Values`.
func (r AccountAIRunCfHuggingfaceExecuteDistilbertSst2Int8Params) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
