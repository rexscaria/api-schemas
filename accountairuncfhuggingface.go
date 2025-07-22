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

type AccountAIRunCfHuggingfaceExecuteDistilbertSst2Int8Response struct {
	// An array of classification results for the input text
	Result  []AccountAIRunCfHuggingfaceExecuteDistilbertSst2Int8ResponseResult `json:"result"`
	Success bool                                                               `json:"success"`
	JSON    accountAIRunCfHuggingfaceExecuteDistilbertSst2Int8ResponseJSON     `json:"-"`
}

// accountAIRunCfHuggingfaceExecuteDistilbertSst2Int8ResponseJSON contains the JSON
// metadata for the struct
// [AccountAIRunCfHuggingfaceExecuteDistilbertSst2Int8Response]
type accountAIRunCfHuggingfaceExecuteDistilbertSst2Int8ResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIRunCfHuggingfaceExecuteDistilbertSst2Int8Response) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunCfHuggingfaceExecuteDistilbertSst2Int8ResponseJSON) RawJSON() string {
	return r.raw
}

type AccountAIRunCfHuggingfaceExecuteDistilbertSst2Int8ResponseResult struct {
	// The classification label assigned to the text (e.g., 'POSITIVE' or 'NEGATIVE')
	Label string `json:"label"`
	// Confidence score indicating the likelihood that the text belongs to the
	// specified label
	Score float64                                                              `json:"score"`
	JSON  accountAIRunCfHuggingfaceExecuteDistilbertSst2Int8ResponseResultJSON `json:"-"`
}

// accountAIRunCfHuggingfaceExecuteDistilbertSst2Int8ResponseResultJSON contains
// the JSON metadata for the struct
// [AccountAIRunCfHuggingfaceExecuteDistilbertSst2Int8ResponseResult]
type accountAIRunCfHuggingfaceExecuteDistilbertSst2Int8ResponseResultJSON struct {
	Label       apijson.Field
	Score       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIRunCfHuggingfaceExecuteDistilbertSst2Int8ResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunCfHuggingfaceExecuteDistilbertSst2Int8ResponseResultJSON) RawJSON() string {
	return r.raw
}

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
