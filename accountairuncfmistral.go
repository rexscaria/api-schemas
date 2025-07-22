// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/apiquery"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
	"github.com/rexscaria/api-schemas/shared"
	"github.com/tidwall/gjson"
)

// AccountAIRunCfMistralService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountAIRunCfMistralService] method instead.
type AccountAIRunCfMistralService struct {
	Options []option.RequestOption
}

// NewAccountAIRunCfMistralService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountAIRunCfMistralService(opts ...option.RequestOption) (r *AccountAIRunCfMistralService) {
	r = &AccountAIRunCfMistralService{}
	r.Options = opts
	return
}

// Execute @cf/mistral/mistral-7b-instruct-v0.1 model.
func (r *AccountAIRunCfMistralService) ExecuteMistral7bInstructV0_1(ctx context.Context, accountID string, params AccountAIRunCfMistralExecuteMistral7bInstructV0_1Params, opts ...option.RequestOption) (res *AccountAIRunCfMistralExecuteMistral7bInstructV0_1Response, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai/run/@cf/mistral/mistral-7b-instruct-v0.1", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Execute @cf/mistral/mistral-7b-instruct-v0.2-lora model.
func (r *AccountAIRunCfMistralService) ExecuteMistral7bInstructV0_2Lora(ctx context.Context, accountID string, params AccountAIRunCfMistralExecuteMistral7bInstructV0_2LoraParams, opts ...option.RequestOption) (res *AccountAIRunCfMistralExecuteMistral7bInstructV0_2LoraResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai/run/@cf/mistral/mistral-7b-instruct-v0.2-lora", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type AccountAIRunCfMistralExecuteMistral7bInstructV0_1Response struct {
	Result  AccountAIRunCfMistralExecuteMistral7bInstructV0_1ResponseResultUnion `json:"result" format:"binary"`
	Success bool                                                                 `json:"success"`
	JSON    accountAIRunCfMistralExecuteMistral7bInstructV0_1ResponseJSON        `json:"-"`
}

// accountAIRunCfMistralExecuteMistral7bInstructV0_1ResponseJSON contains the JSON
// metadata for the struct
// [AccountAIRunCfMistralExecuteMistral7bInstructV0_1Response]
type accountAIRunCfMistralExecuteMistral7bInstructV0_1ResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIRunCfMistralExecuteMistral7bInstructV0_1Response) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunCfMistralExecuteMistral7bInstructV0_1ResponseJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by
// [AccountAIRunCfMistralExecuteMistral7bInstructV0_1ResponseResultObject] or
// [shared.UnionString].
type AccountAIRunCfMistralExecuteMistral7bInstructV0_1ResponseResultUnion interface {
	ImplementsAccountAIRunCfMistralExecuteMistral7bInstructV0_1ResponseResultUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccountAIRunCfMistralExecuteMistral7bInstructV0_1ResponseResultUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccountAIRunCfMistralExecuteMistral7bInstructV0_1ResponseResultObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type AccountAIRunCfMistralExecuteMistral7bInstructV0_1ResponseResultObject struct {
	// The generated text response from the model
	Response string `json:"response,required"`
	// An array of tool calls requests made during the response generation
	ToolCalls []AccountAIRunCfMistralExecuteMistral7bInstructV0_1ResponseResultObjectToolCall `json:"tool_calls"`
	// Usage statistics for the inference request
	Usage AccountAIRunCfMistralExecuteMistral7bInstructV0_1ResponseResultObjectUsage `json:"usage"`
	JSON  accountAIRunCfMistralExecuteMistral7bInstructV0_1ResponseResultObjectJSON  `json:"-"`
}

// accountAIRunCfMistralExecuteMistral7bInstructV0_1ResponseResultObjectJSON
// contains the JSON metadata for the struct
// [AccountAIRunCfMistralExecuteMistral7bInstructV0_1ResponseResultObject]
type accountAIRunCfMistralExecuteMistral7bInstructV0_1ResponseResultObjectJSON struct {
	Response    apijson.Field
	ToolCalls   apijson.Field
	Usage       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIRunCfMistralExecuteMistral7bInstructV0_1ResponseResultObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunCfMistralExecuteMistral7bInstructV0_1ResponseResultObjectJSON) RawJSON() string {
	return r.raw
}

func (r AccountAIRunCfMistralExecuteMistral7bInstructV0_1ResponseResultObject) ImplementsAccountAIRunCfMistralExecuteMistral7bInstructV0_1ResponseResultUnion() {
}

type AccountAIRunCfMistralExecuteMistral7bInstructV0_1ResponseResultObjectToolCall struct {
	// The arguments passed to be passed to the tool call request
	Arguments interface{} `json:"arguments"`
	// The name of the tool to be called
	Name string                                                                            `json:"name"`
	JSON accountAIRunCfMistralExecuteMistral7bInstructV0_1ResponseResultObjectToolCallJSON `json:"-"`
}

// accountAIRunCfMistralExecuteMistral7bInstructV0_1ResponseResultObjectToolCallJSON
// contains the JSON metadata for the struct
// [AccountAIRunCfMistralExecuteMistral7bInstructV0_1ResponseResultObjectToolCall]
type accountAIRunCfMistralExecuteMistral7bInstructV0_1ResponseResultObjectToolCallJSON struct {
	Arguments   apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIRunCfMistralExecuteMistral7bInstructV0_1ResponseResultObjectToolCall) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunCfMistralExecuteMistral7bInstructV0_1ResponseResultObjectToolCallJSON) RawJSON() string {
	return r.raw
}

// Usage statistics for the inference request
type AccountAIRunCfMistralExecuteMistral7bInstructV0_1ResponseResultObjectUsage struct {
	// Total number of tokens in output
	CompletionTokens float64 `json:"completion_tokens"`
	// Total number of tokens in input
	PromptTokens float64 `json:"prompt_tokens"`
	// Total number of input and output tokens
	TotalTokens float64                                                                        `json:"total_tokens"`
	JSON        accountAIRunCfMistralExecuteMistral7bInstructV0_1ResponseResultObjectUsageJSON `json:"-"`
}

// accountAIRunCfMistralExecuteMistral7bInstructV0_1ResponseResultObjectUsageJSON
// contains the JSON metadata for the struct
// [AccountAIRunCfMistralExecuteMistral7bInstructV0_1ResponseResultObjectUsage]
type accountAIRunCfMistralExecuteMistral7bInstructV0_1ResponseResultObjectUsageJSON struct {
	CompletionTokens apijson.Field
	PromptTokens     apijson.Field
	TotalTokens      apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountAIRunCfMistralExecuteMistral7bInstructV0_1ResponseResultObjectUsage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunCfMistralExecuteMistral7bInstructV0_1ResponseResultObjectUsageJSON) RawJSON() string {
	return r.raw
}

type AccountAIRunCfMistralExecuteMistral7bInstructV0_2LoraResponse struct {
	Result  AccountAIRunCfMistralExecuteMistral7bInstructV0_2LoraResponseResultUnion `json:"result" format:"binary"`
	Success bool                                                                     `json:"success"`
	JSON    accountAIRunCfMistralExecuteMistral7bInstructV0_2LoraResponseJSON        `json:"-"`
}

// accountAIRunCfMistralExecuteMistral7bInstructV0_2LoraResponseJSON contains the
// JSON metadata for the struct
// [AccountAIRunCfMistralExecuteMistral7bInstructV0_2LoraResponse]
type accountAIRunCfMistralExecuteMistral7bInstructV0_2LoraResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIRunCfMistralExecuteMistral7bInstructV0_2LoraResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunCfMistralExecuteMistral7bInstructV0_2LoraResponseJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by
// [AccountAIRunCfMistralExecuteMistral7bInstructV0_2LoraResponseResultObject] or
// [shared.UnionString].
type AccountAIRunCfMistralExecuteMistral7bInstructV0_2LoraResponseResultUnion interface {
	ImplementsAccountAIRunCfMistralExecuteMistral7bInstructV0_2LoraResponseResultUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccountAIRunCfMistralExecuteMistral7bInstructV0_2LoraResponseResultUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccountAIRunCfMistralExecuteMistral7bInstructV0_2LoraResponseResultObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type AccountAIRunCfMistralExecuteMistral7bInstructV0_2LoraResponseResultObject struct {
	// The generated text response from the model
	Response string `json:"response,required"`
	// An array of tool calls requests made during the response generation
	ToolCalls []AccountAIRunCfMistralExecuteMistral7bInstructV0_2LoraResponseResultObjectToolCall `json:"tool_calls"`
	// Usage statistics for the inference request
	Usage AccountAIRunCfMistralExecuteMistral7bInstructV0_2LoraResponseResultObjectUsage `json:"usage"`
	JSON  accountAIRunCfMistralExecuteMistral7bInstructV0_2LoraResponseResultObjectJSON  `json:"-"`
}

// accountAIRunCfMistralExecuteMistral7bInstructV0_2LoraResponseResultObjectJSON
// contains the JSON metadata for the struct
// [AccountAIRunCfMistralExecuteMistral7bInstructV0_2LoraResponseResultObject]
type accountAIRunCfMistralExecuteMistral7bInstructV0_2LoraResponseResultObjectJSON struct {
	Response    apijson.Field
	ToolCalls   apijson.Field
	Usage       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIRunCfMistralExecuteMistral7bInstructV0_2LoraResponseResultObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunCfMistralExecuteMistral7bInstructV0_2LoraResponseResultObjectJSON) RawJSON() string {
	return r.raw
}

func (r AccountAIRunCfMistralExecuteMistral7bInstructV0_2LoraResponseResultObject) ImplementsAccountAIRunCfMistralExecuteMistral7bInstructV0_2LoraResponseResultUnion() {
}

type AccountAIRunCfMistralExecuteMistral7bInstructV0_2LoraResponseResultObjectToolCall struct {
	// The arguments passed to be passed to the tool call request
	Arguments interface{} `json:"arguments"`
	// The name of the tool to be called
	Name string                                                                                `json:"name"`
	JSON accountAIRunCfMistralExecuteMistral7bInstructV0_2LoraResponseResultObjectToolCallJSON `json:"-"`
}

// accountAIRunCfMistralExecuteMistral7bInstructV0_2LoraResponseResultObjectToolCallJSON
// contains the JSON metadata for the struct
// [AccountAIRunCfMistralExecuteMistral7bInstructV0_2LoraResponseResultObjectToolCall]
type accountAIRunCfMistralExecuteMistral7bInstructV0_2LoraResponseResultObjectToolCallJSON struct {
	Arguments   apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIRunCfMistralExecuteMistral7bInstructV0_2LoraResponseResultObjectToolCall) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunCfMistralExecuteMistral7bInstructV0_2LoraResponseResultObjectToolCallJSON) RawJSON() string {
	return r.raw
}

// Usage statistics for the inference request
type AccountAIRunCfMistralExecuteMistral7bInstructV0_2LoraResponseResultObjectUsage struct {
	// Total number of tokens in output
	CompletionTokens float64 `json:"completion_tokens"`
	// Total number of tokens in input
	PromptTokens float64 `json:"prompt_tokens"`
	// Total number of input and output tokens
	TotalTokens float64                                                                            `json:"total_tokens"`
	JSON        accountAIRunCfMistralExecuteMistral7bInstructV0_2LoraResponseResultObjectUsageJSON `json:"-"`
}

// accountAIRunCfMistralExecuteMistral7bInstructV0_2LoraResponseResultObjectUsageJSON
// contains the JSON metadata for the struct
// [AccountAIRunCfMistralExecuteMistral7bInstructV0_2LoraResponseResultObjectUsage]
type accountAIRunCfMistralExecuteMistral7bInstructV0_2LoraResponseResultObjectUsageJSON struct {
	CompletionTokens apijson.Field
	PromptTokens     apijson.Field
	TotalTokens      apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountAIRunCfMistralExecuteMistral7bInstructV0_2LoraResponseResultObjectUsage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunCfMistralExecuteMistral7bInstructV0_2LoraResponseResultObjectUsageJSON) RawJSON() string {
	return r.raw
}

type AccountAIRunCfMistralExecuteMistral7bInstructV0_1Params struct {
	QueueRequest param.Field[string]                                              `query:"queueRequest"`
	Body         AccountAIRunCfMistralExecuteMistral7bInstructV0_1ParamsBodyUnion `json:"body"`
}

func (r AccountAIRunCfMistralExecuteMistral7bInstructV0_1Params) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

// URLQuery serializes [AccountAIRunCfMistralExecuteMistral7bInstructV0_1Params]'s
// query parameters as `url.Values`.
func (r AccountAIRunCfMistralExecuteMistral7bInstructV0_1Params) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountAIRunCfMistralExecuteMistral7bInstructV0_1ParamsBody struct {
	// Decreases the likelihood of the model repeating the same lines verbatim.
	FrequencyPenalty param.Field[float64]     `json:"frequency_penalty"`
	Functions        param.Field[interface{}] `json:"functions"`
	// Name of the LoRA (Low-Rank Adaptation) model to fine-tune the base model.
	Lora param.Field[string] `json:"lora"`
	// The maximum number of tokens to generate in the response.
	MaxTokens param.Field[int64]       `json:"max_tokens"`
	Messages  param.Field[interface{}] `json:"messages"`
	// Increases the likelihood of the model introducing new topics.
	PresencePenalty param.Field[float64] `json:"presence_penalty"`
	// The input text prompt for the model to generate a response.
	Prompt param.Field[string] `json:"prompt"`
	// If true, a chat template is not applied and you must adhere to the specific
	// model's expected formatting.
	Raw param.Field[bool] `json:"raw"`
	// Penalty for repeated tokens; higher values discourage repetition.
	RepetitionPenalty param.Field[float64]     `json:"repetition_penalty"`
	ResponseFormat    param.Field[interface{}] `json:"response_format"`
	// Random seed for reproducibility of the generation.
	Seed param.Field[int64] `json:"seed"`
	// If true, the response will be streamed back incrementally using SSE, Server Sent
	// Events.
	Stream param.Field[bool] `json:"stream"`
	// Controls the randomness of the output; higher values produce more random
	// results.
	Temperature param.Field[float64]     `json:"temperature"`
	Tools       param.Field[interface{}] `json:"tools"`
	// Limits the AI to choose from the top 'k' most probable words. Lower values make
	// responses more focused; higher values introduce more variety and potential
	// surprises.
	TopK param.Field[int64] `json:"top_k"`
	// Adjusts the creativity of the AI's responses by controlling how many possible
	// words it considers. Lower values make outputs more predictable; higher values
	// allow for more varied and creative responses.
	TopP param.Field[float64] `json:"top_p"`
}

func (r AccountAIRunCfMistralExecuteMistral7bInstructV0_1ParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfMistralExecuteMistral7bInstructV0_1ParamsBody) implementsAccountAIRunCfMistralExecuteMistral7bInstructV0_1ParamsBodyUnion() {
}

// Satisfied by
// [AccountAIRunCfMistralExecuteMistral7bInstructV0_1ParamsBodyPrompt],
// [AccountAIRunCfMistralExecuteMistral7bInstructV0_1ParamsBodyMessages],
// [AccountAIRunCfMistralExecuteMistral7bInstructV0_1ParamsBody].
type AccountAIRunCfMistralExecuteMistral7bInstructV0_1ParamsBodyUnion interface {
	implementsAccountAIRunCfMistralExecuteMistral7bInstructV0_1ParamsBodyUnion()
}

type AccountAIRunCfMistralExecuteMistral7bInstructV0_1ParamsBodyPrompt struct {
	// The input text prompt for the model to generate a response.
	Prompt param.Field[string] `json:"prompt,required"`
	// Decreases the likelihood of the model repeating the same lines verbatim.
	FrequencyPenalty param.Field[float64] `json:"frequency_penalty"`
	// Name of the LoRA (Low-Rank Adaptation) model to fine-tune the base model.
	Lora param.Field[string] `json:"lora"`
	// The maximum number of tokens to generate in the response.
	MaxTokens param.Field[int64] `json:"max_tokens"`
	// Increases the likelihood of the model introducing new topics.
	PresencePenalty param.Field[float64] `json:"presence_penalty"`
	// If true, a chat template is not applied and you must adhere to the specific
	// model's expected formatting.
	Raw param.Field[bool] `json:"raw"`
	// Penalty for repeated tokens; higher values discourage repetition.
	RepetitionPenalty param.Field[float64]                                                                         `json:"repetition_penalty"`
	ResponseFormat    param.Field[AccountAIRunCfMistralExecuteMistral7bInstructV0_1ParamsBodyPromptResponseFormat] `json:"response_format"`
	// Random seed for reproducibility of the generation.
	Seed param.Field[int64] `json:"seed"`
	// If true, the response will be streamed back incrementally using SSE, Server Sent
	// Events.
	Stream param.Field[bool] `json:"stream"`
	// Controls the randomness of the output; higher values produce more random
	// results.
	Temperature param.Field[float64] `json:"temperature"`
	// Limits the AI to choose from the top 'k' most probable words. Lower values make
	// responses more focused; higher values introduce more variety and potential
	// surprises.
	TopK param.Field[int64] `json:"top_k"`
	// Adjusts the creativity of the AI's responses by controlling how many possible
	// words it considers. Lower values make outputs more predictable; higher values
	// allow for more varied and creative responses.
	TopP param.Field[float64] `json:"top_p"`
}

func (r AccountAIRunCfMistralExecuteMistral7bInstructV0_1ParamsBodyPrompt) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfMistralExecuteMistral7bInstructV0_1ParamsBodyPrompt) implementsAccountAIRunCfMistralExecuteMistral7bInstructV0_1ParamsBodyUnion() {
}

type AccountAIRunCfMistralExecuteMistral7bInstructV0_1ParamsBodyPromptResponseFormat struct {
	JsonSchema param.Field[interface{}]                                                                         `json:"json_schema"`
	Type       param.Field[AccountAIRunCfMistralExecuteMistral7bInstructV0_1ParamsBodyPromptResponseFormatType] `json:"type"`
}

func (r AccountAIRunCfMistralExecuteMistral7bInstructV0_1ParamsBodyPromptResponseFormat) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfMistralExecuteMistral7bInstructV0_1ParamsBodyPromptResponseFormatType string

const (
	AccountAIRunCfMistralExecuteMistral7bInstructV0_1ParamsBodyPromptResponseFormatTypeJsonObject AccountAIRunCfMistralExecuteMistral7bInstructV0_1ParamsBodyPromptResponseFormatType = "json_object"
	AccountAIRunCfMistralExecuteMistral7bInstructV0_1ParamsBodyPromptResponseFormatTypeJsonSchema AccountAIRunCfMistralExecuteMistral7bInstructV0_1ParamsBodyPromptResponseFormatType = "json_schema"
)

func (r AccountAIRunCfMistralExecuteMistral7bInstructV0_1ParamsBodyPromptResponseFormatType) IsKnown() bool {
	switch r {
	case AccountAIRunCfMistralExecuteMistral7bInstructV0_1ParamsBodyPromptResponseFormatTypeJsonObject, AccountAIRunCfMistralExecuteMistral7bInstructV0_1ParamsBodyPromptResponseFormatTypeJsonSchema:
		return true
	}
	return false
}

type AccountAIRunCfMistralExecuteMistral7bInstructV0_1ParamsBodyMessages struct {
	// An array of message objects representing the conversation history.
	Messages param.Field[[]AccountAIRunCfMistralExecuteMistral7bInstructV0_1ParamsBodyMessagesMessage] `json:"messages,required"`
	// Decreases the likelihood of the model repeating the same lines verbatim.
	FrequencyPenalty param.Field[float64]                                                                       `json:"frequency_penalty"`
	Functions        param.Field[[]AccountAIRunCfMistralExecuteMistral7bInstructV0_1ParamsBodyMessagesFunction] `json:"functions"`
	// The maximum number of tokens to generate in the response.
	MaxTokens param.Field[int64] `json:"max_tokens"`
	// Increases the likelihood of the model introducing new topics.
	PresencePenalty param.Field[float64] `json:"presence_penalty"`
	// If true, a chat template is not applied and you must adhere to the specific
	// model's expected formatting.
	Raw param.Field[bool] `json:"raw"`
	// Penalty for repeated tokens; higher values discourage repetition.
	RepetitionPenalty param.Field[float64]                                                                           `json:"repetition_penalty"`
	ResponseFormat    param.Field[AccountAIRunCfMistralExecuteMistral7bInstructV0_1ParamsBodyMessagesResponseFormat] `json:"response_format"`
	// Random seed for reproducibility of the generation.
	Seed param.Field[int64] `json:"seed"`
	// If true, the response will be streamed back incrementally using SSE, Server Sent
	// Events.
	Stream param.Field[bool] `json:"stream"`
	// Controls the randomness of the output; higher values produce more random
	// results.
	Temperature param.Field[float64] `json:"temperature"`
	// A list of tools available for the assistant to use.
	Tools param.Field[[]AccountAIRunCfMistralExecuteMistral7bInstructV0_1ParamsBodyMessagesToolUnion] `json:"tools"`
	// Limits the AI to choose from the top 'k' most probable words. Lower values make
	// responses more focused; higher values introduce more variety and potential
	// surprises.
	TopK param.Field[int64] `json:"top_k"`
	// Adjusts the creativity of the AI's responses by controlling how many possible
	// words it considers. Lower values make outputs more predictable; higher values
	// allow for more varied and creative responses.
	TopP param.Field[float64] `json:"top_p"`
}

func (r AccountAIRunCfMistralExecuteMistral7bInstructV0_1ParamsBodyMessages) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfMistralExecuteMistral7bInstructV0_1ParamsBodyMessages) implementsAccountAIRunCfMistralExecuteMistral7bInstructV0_1ParamsBodyUnion() {
}

type AccountAIRunCfMistralExecuteMistral7bInstructV0_1ParamsBodyMessagesMessage struct {
	// The content of the message as a string.
	Content param.Field[string] `json:"content,required"`
	// The role of the message sender (e.g., 'user', 'assistant', 'system', 'tool').
	Role param.Field[string] `json:"role,required"`
}

func (r AccountAIRunCfMistralExecuteMistral7bInstructV0_1ParamsBodyMessagesMessage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfMistralExecuteMistral7bInstructV0_1ParamsBodyMessagesFunction struct {
	Code param.Field[string] `json:"code,required"`
	Name param.Field[string] `json:"name,required"`
}

func (r AccountAIRunCfMistralExecuteMistral7bInstructV0_1ParamsBodyMessagesFunction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfMistralExecuteMistral7bInstructV0_1ParamsBodyMessagesResponseFormat struct {
	JsonSchema param.Field[interface{}]                                                                           `json:"json_schema"`
	Type       param.Field[AccountAIRunCfMistralExecuteMistral7bInstructV0_1ParamsBodyMessagesResponseFormatType] `json:"type"`
}

func (r AccountAIRunCfMistralExecuteMistral7bInstructV0_1ParamsBodyMessagesResponseFormat) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfMistralExecuteMistral7bInstructV0_1ParamsBodyMessagesResponseFormatType string

const (
	AccountAIRunCfMistralExecuteMistral7bInstructV0_1ParamsBodyMessagesResponseFormatTypeJsonObject AccountAIRunCfMistralExecuteMistral7bInstructV0_1ParamsBodyMessagesResponseFormatType = "json_object"
	AccountAIRunCfMistralExecuteMistral7bInstructV0_1ParamsBodyMessagesResponseFormatTypeJsonSchema AccountAIRunCfMistralExecuteMistral7bInstructV0_1ParamsBodyMessagesResponseFormatType = "json_schema"
)

func (r AccountAIRunCfMistralExecuteMistral7bInstructV0_1ParamsBodyMessagesResponseFormatType) IsKnown() bool {
	switch r {
	case AccountAIRunCfMistralExecuteMistral7bInstructV0_1ParamsBodyMessagesResponseFormatTypeJsonObject, AccountAIRunCfMistralExecuteMistral7bInstructV0_1ParamsBodyMessagesResponseFormatTypeJsonSchema:
		return true
	}
	return false
}

type AccountAIRunCfMistralExecuteMistral7bInstructV0_1ParamsBodyMessagesTool struct {
	// A brief description of what the tool does.
	Description param.Field[string]      `json:"description"`
	Function    param.Field[interface{}] `json:"function"`
	// The name of the tool. More descriptive the better.
	Name       param.Field[string]      `json:"name"`
	Parameters param.Field[interface{}] `json:"parameters"`
	// Specifies the type of tool (e.g., 'function').
	Type param.Field[string] `json:"type"`
}

func (r AccountAIRunCfMistralExecuteMistral7bInstructV0_1ParamsBodyMessagesTool) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfMistralExecuteMistral7bInstructV0_1ParamsBodyMessagesTool) implementsAccountAIRunCfMistralExecuteMistral7bInstructV0_1ParamsBodyMessagesToolUnion() {
}

// Satisfied by
// [AccountAIRunCfMistralExecuteMistral7bInstructV0_1ParamsBodyMessagesToolsObject],
// [AccountAIRunCfMistralExecuteMistral7bInstructV0_1ParamsBodyMessagesToolsObject],
// [AccountAIRunCfMistralExecuteMistral7bInstructV0_1ParamsBodyMessagesTool].
type AccountAIRunCfMistralExecuteMistral7bInstructV0_1ParamsBodyMessagesToolUnion interface {
	implementsAccountAIRunCfMistralExecuteMistral7bInstructV0_1ParamsBodyMessagesToolUnion()
}

type AccountAIRunCfMistralExecuteMistral7bInstructV0_1ParamsBodyMessagesToolsObject struct {
	// A brief description of what the tool does.
	Description param.Field[string] `json:"description,required"`
	// The name of the tool. More descriptive the better.
	Name param.Field[string] `json:"name,required"`
	// Schema defining the parameters accepted by the tool.
	Parameters param.Field[AccountAIRunCfMistralExecuteMistral7bInstructV0_1ParamsBodyMessagesToolsObjectParameters] `json:"parameters,required"`
}

func (r AccountAIRunCfMistralExecuteMistral7bInstructV0_1ParamsBodyMessagesToolsObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfMistralExecuteMistral7bInstructV0_1ParamsBodyMessagesToolsObject) implementsAccountAIRunCfMistralExecuteMistral7bInstructV0_1ParamsBodyMessagesToolUnion() {
}

// Schema defining the parameters accepted by the tool.
type AccountAIRunCfMistralExecuteMistral7bInstructV0_1ParamsBodyMessagesToolsObjectParameters struct {
	// Definitions of each parameter.
	Properties param.Field[map[string]AccountAIRunCfMistralExecuteMistral7bInstructV0_1ParamsBodyMessagesToolsObjectParametersProperties] `json:"properties,required"`
	// The type of the parameters object (usually 'object').
	Type param.Field[string] `json:"type,required"`
	// List of required parameter names.
	Required param.Field[[]string] `json:"required"`
}

func (r AccountAIRunCfMistralExecuteMistral7bInstructV0_1ParamsBodyMessagesToolsObjectParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfMistralExecuteMistral7bInstructV0_1ParamsBodyMessagesToolsObjectParametersProperties struct {
	// A description of the expected parameter.
	Description param.Field[string] `json:"description,required"`
	// The data type of the parameter.
	Type param.Field[string] `json:"type,required"`
}

func (r AccountAIRunCfMistralExecuteMistral7bInstructV0_1ParamsBodyMessagesToolsObjectParametersProperties) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfMistralExecuteMistral7bInstructV0_2LoraParams struct {
	QueueRequest param.Field[string]                                                  `query:"queueRequest"`
	Body         AccountAIRunCfMistralExecuteMistral7bInstructV0_2LoraParamsBodyUnion `json:"body"`
}

func (r AccountAIRunCfMistralExecuteMistral7bInstructV0_2LoraParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

// URLQuery serializes
// [AccountAIRunCfMistralExecuteMistral7bInstructV0_2LoraParams]'s query parameters
// as `url.Values`.
func (r AccountAIRunCfMistralExecuteMistral7bInstructV0_2LoraParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountAIRunCfMistralExecuteMistral7bInstructV0_2LoraParamsBody struct {
	// Decreases the likelihood of the model repeating the same lines verbatim.
	FrequencyPenalty param.Field[float64]     `json:"frequency_penalty"`
	Functions        param.Field[interface{}] `json:"functions"`
	// Name of the LoRA (Low-Rank Adaptation) model to fine-tune the base model.
	Lora param.Field[string] `json:"lora"`
	// The maximum number of tokens to generate in the response.
	MaxTokens param.Field[int64]       `json:"max_tokens"`
	Messages  param.Field[interface{}] `json:"messages"`
	// Increases the likelihood of the model introducing new topics.
	PresencePenalty param.Field[float64] `json:"presence_penalty"`
	// The input text prompt for the model to generate a response.
	Prompt param.Field[string] `json:"prompt"`
	// If true, a chat template is not applied and you must adhere to the specific
	// model's expected formatting.
	Raw param.Field[bool] `json:"raw"`
	// Penalty for repeated tokens; higher values discourage repetition.
	RepetitionPenalty param.Field[float64]     `json:"repetition_penalty"`
	ResponseFormat    param.Field[interface{}] `json:"response_format"`
	// Random seed for reproducibility of the generation.
	Seed param.Field[int64] `json:"seed"`
	// If true, the response will be streamed back incrementally using SSE, Server Sent
	// Events.
	Stream param.Field[bool] `json:"stream"`
	// Controls the randomness of the output; higher values produce more random
	// results.
	Temperature param.Field[float64]     `json:"temperature"`
	Tools       param.Field[interface{}] `json:"tools"`
	// Limits the AI to choose from the top 'k' most probable words. Lower values make
	// responses more focused; higher values introduce more variety and potential
	// surprises.
	TopK param.Field[int64] `json:"top_k"`
	// Adjusts the creativity of the AI's responses by controlling how many possible
	// words it considers. Lower values make outputs more predictable; higher values
	// allow for more varied and creative responses.
	TopP param.Field[float64] `json:"top_p"`
}

func (r AccountAIRunCfMistralExecuteMistral7bInstructV0_2LoraParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfMistralExecuteMistral7bInstructV0_2LoraParamsBody) implementsAccountAIRunCfMistralExecuteMistral7bInstructV0_2LoraParamsBodyUnion() {
}

// Satisfied by
// [AccountAIRunCfMistralExecuteMistral7bInstructV0_2LoraParamsBodyPrompt],
// [AccountAIRunCfMistralExecuteMistral7bInstructV0_2LoraParamsBodyMessages],
// [AccountAIRunCfMistralExecuteMistral7bInstructV0_2LoraParamsBody].
type AccountAIRunCfMistralExecuteMistral7bInstructV0_2LoraParamsBodyUnion interface {
	implementsAccountAIRunCfMistralExecuteMistral7bInstructV0_2LoraParamsBodyUnion()
}

type AccountAIRunCfMistralExecuteMistral7bInstructV0_2LoraParamsBodyPrompt struct {
	// The input text prompt for the model to generate a response.
	Prompt param.Field[string] `json:"prompt,required"`
	// Decreases the likelihood of the model repeating the same lines verbatim.
	FrequencyPenalty param.Field[float64] `json:"frequency_penalty"`
	// Name of the LoRA (Low-Rank Adaptation) model to fine-tune the base model.
	Lora param.Field[string] `json:"lora"`
	// The maximum number of tokens to generate in the response.
	MaxTokens param.Field[int64] `json:"max_tokens"`
	// Increases the likelihood of the model introducing new topics.
	PresencePenalty param.Field[float64] `json:"presence_penalty"`
	// If true, a chat template is not applied and you must adhere to the specific
	// model's expected formatting.
	Raw param.Field[bool] `json:"raw"`
	// Penalty for repeated tokens; higher values discourage repetition.
	RepetitionPenalty param.Field[float64]                                                                             `json:"repetition_penalty"`
	ResponseFormat    param.Field[AccountAIRunCfMistralExecuteMistral7bInstructV0_2LoraParamsBodyPromptResponseFormat] `json:"response_format"`
	// Random seed for reproducibility of the generation.
	Seed param.Field[int64] `json:"seed"`
	// If true, the response will be streamed back incrementally using SSE, Server Sent
	// Events.
	Stream param.Field[bool] `json:"stream"`
	// Controls the randomness of the output; higher values produce more random
	// results.
	Temperature param.Field[float64] `json:"temperature"`
	// Limits the AI to choose from the top 'k' most probable words. Lower values make
	// responses more focused; higher values introduce more variety and potential
	// surprises.
	TopK param.Field[int64] `json:"top_k"`
	// Adjusts the creativity of the AI's responses by controlling how many possible
	// words it considers. Lower values make outputs more predictable; higher values
	// allow for more varied and creative responses.
	TopP param.Field[float64] `json:"top_p"`
}

func (r AccountAIRunCfMistralExecuteMistral7bInstructV0_2LoraParamsBodyPrompt) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfMistralExecuteMistral7bInstructV0_2LoraParamsBodyPrompt) implementsAccountAIRunCfMistralExecuteMistral7bInstructV0_2LoraParamsBodyUnion() {
}

type AccountAIRunCfMistralExecuteMistral7bInstructV0_2LoraParamsBodyPromptResponseFormat struct {
	JsonSchema param.Field[interface{}]                                                                             `json:"json_schema"`
	Type       param.Field[AccountAIRunCfMistralExecuteMistral7bInstructV0_2LoraParamsBodyPromptResponseFormatType] `json:"type"`
}

func (r AccountAIRunCfMistralExecuteMistral7bInstructV0_2LoraParamsBodyPromptResponseFormat) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfMistralExecuteMistral7bInstructV0_2LoraParamsBodyPromptResponseFormatType string

const (
	AccountAIRunCfMistralExecuteMistral7bInstructV0_2LoraParamsBodyPromptResponseFormatTypeJsonObject AccountAIRunCfMistralExecuteMistral7bInstructV0_2LoraParamsBodyPromptResponseFormatType = "json_object"
	AccountAIRunCfMistralExecuteMistral7bInstructV0_2LoraParamsBodyPromptResponseFormatTypeJsonSchema AccountAIRunCfMistralExecuteMistral7bInstructV0_2LoraParamsBodyPromptResponseFormatType = "json_schema"
)

func (r AccountAIRunCfMistralExecuteMistral7bInstructV0_2LoraParamsBodyPromptResponseFormatType) IsKnown() bool {
	switch r {
	case AccountAIRunCfMistralExecuteMistral7bInstructV0_2LoraParamsBodyPromptResponseFormatTypeJsonObject, AccountAIRunCfMistralExecuteMistral7bInstructV0_2LoraParamsBodyPromptResponseFormatTypeJsonSchema:
		return true
	}
	return false
}

type AccountAIRunCfMistralExecuteMistral7bInstructV0_2LoraParamsBodyMessages struct {
	// An array of message objects representing the conversation history.
	Messages param.Field[[]AccountAIRunCfMistralExecuteMistral7bInstructV0_2LoraParamsBodyMessagesMessage] `json:"messages,required"`
	// Decreases the likelihood of the model repeating the same lines verbatim.
	FrequencyPenalty param.Field[float64]                                                                           `json:"frequency_penalty"`
	Functions        param.Field[[]AccountAIRunCfMistralExecuteMistral7bInstructV0_2LoraParamsBodyMessagesFunction] `json:"functions"`
	// The maximum number of tokens to generate in the response.
	MaxTokens param.Field[int64] `json:"max_tokens"`
	// Increases the likelihood of the model introducing new topics.
	PresencePenalty param.Field[float64] `json:"presence_penalty"`
	// If true, a chat template is not applied and you must adhere to the specific
	// model's expected formatting.
	Raw param.Field[bool] `json:"raw"`
	// Penalty for repeated tokens; higher values discourage repetition.
	RepetitionPenalty param.Field[float64]                                                                               `json:"repetition_penalty"`
	ResponseFormat    param.Field[AccountAIRunCfMistralExecuteMistral7bInstructV0_2LoraParamsBodyMessagesResponseFormat] `json:"response_format"`
	// Random seed for reproducibility of the generation.
	Seed param.Field[int64] `json:"seed"`
	// If true, the response will be streamed back incrementally using SSE, Server Sent
	// Events.
	Stream param.Field[bool] `json:"stream"`
	// Controls the randomness of the output; higher values produce more random
	// results.
	Temperature param.Field[float64] `json:"temperature"`
	// A list of tools available for the assistant to use.
	Tools param.Field[[]AccountAIRunCfMistralExecuteMistral7bInstructV0_2LoraParamsBodyMessagesToolUnion] `json:"tools"`
	// Limits the AI to choose from the top 'k' most probable words. Lower values make
	// responses more focused; higher values introduce more variety and potential
	// surprises.
	TopK param.Field[int64] `json:"top_k"`
	// Adjusts the creativity of the AI's responses by controlling how many possible
	// words it considers. Lower values make outputs more predictable; higher values
	// allow for more varied and creative responses.
	TopP param.Field[float64] `json:"top_p"`
}

func (r AccountAIRunCfMistralExecuteMistral7bInstructV0_2LoraParamsBodyMessages) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfMistralExecuteMistral7bInstructV0_2LoraParamsBodyMessages) implementsAccountAIRunCfMistralExecuteMistral7bInstructV0_2LoraParamsBodyUnion() {
}

type AccountAIRunCfMistralExecuteMistral7bInstructV0_2LoraParamsBodyMessagesMessage struct {
	// The content of the message as a string.
	Content param.Field[string] `json:"content,required"`
	// The role of the message sender (e.g., 'user', 'assistant', 'system', 'tool').
	Role param.Field[string] `json:"role,required"`
}

func (r AccountAIRunCfMistralExecuteMistral7bInstructV0_2LoraParamsBodyMessagesMessage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfMistralExecuteMistral7bInstructV0_2LoraParamsBodyMessagesFunction struct {
	Code param.Field[string] `json:"code,required"`
	Name param.Field[string] `json:"name,required"`
}

func (r AccountAIRunCfMistralExecuteMistral7bInstructV0_2LoraParamsBodyMessagesFunction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfMistralExecuteMistral7bInstructV0_2LoraParamsBodyMessagesResponseFormat struct {
	JsonSchema param.Field[interface{}]                                                                               `json:"json_schema"`
	Type       param.Field[AccountAIRunCfMistralExecuteMistral7bInstructV0_2LoraParamsBodyMessagesResponseFormatType] `json:"type"`
}

func (r AccountAIRunCfMistralExecuteMistral7bInstructV0_2LoraParamsBodyMessagesResponseFormat) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfMistralExecuteMistral7bInstructV0_2LoraParamsBodyMessagesResponseFormatType string

const (
	AccountAIRunCfMistralExecuteMistral7bInstructV0_2LoraParamsBodyMessagesResponseFormatTypeJsonObject AccountAIRunCfMistralExecuteMistral7bInstructV0_2LoraParamsBodyMessagesResponseFormatType = "json_object"
	AccountAIRunCfMistralExecuteMistral7bInstructV0_2LoraParamsBodyMessagesResponseFormatTypeJsonSchema AccountAIRunCfMistralExecuteMistral7bInstructV0_2LoraParamsBodyMessagesResponseFormatType = "json_schema"
)

func (r AccountAIRunCfMistralExecuteMistral7bInstructV0_2LoraParamsBodyMessagesResponseFormatType) IsKnown() bool {
	switch r {
	case AccountAIRunCfMistralExecuteMistral7bInstructV0_2LoraParamsBodyMessagesResponseFormatTypeJsonObject, AccountAIRunCfMistralExecuteMistral7bInstructV0_2LoraParamsBodyMessagesResponseFormatTypeJsonSchema:
		return true
	}
	return false
}

type AccountAIRunCfMistralExecuteMistral7bInstructV0_2LoraParamsBodyMessagesTool struct {
	// A brief description of what the tool does.
	Description param.Field[string]      `json:"description"`
	Function    param.Field[interface{}] `json:"function"`
	// The name of the tool. More descriptive the better.
	Name       param.Field[string]      `json:"name"`
	Parameters param.Field[interface{}] `json:"parameters"`
	// Specifies the type of tool (e.g., 'function').
	Type param.Field[string] `json:"type"`
}

func (r AccountAIRunCfMistralExecuteMistral7bInstructV0_2LoraParamsBodyMessagesTool) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfMistralExecuteMistral7bInstructV0_2LoraParamsBodyMessagesTool) implementsAccountAIRunCfMistralExecuteMistral7bInstructV0_2LoraParamsBodyMessagesToolUnion() {
}

// Satisfied by
// [AccountAIRunCfMistralExecuteMistral7bInstructV0_2LoraParamsBodyMessagesToolsObject],
// [AccountAIRunCfMistralExecuteMistral7bInstructV0_2LoraParamsBodyMessagesToolsObject],
// [AccountAIRunCfMistralExecuteMistral7bInstructV0_2LoraParamsBodyMessagesTool].
type AccountAIRunCfMistralExecuteMistral7bInstructV0_2LoraParamsBodyMessagesToolUnion interface {
	implementsAccountAIRunCfMistralExecuteMistral7bInstructV0_2LoraParamsBodyMessagesToolUnion()
}

type AccountAIRunCfMistralExecuteMistral7bInstructV0_2LoraParamsBodyMessagesToolsObject struct {
	// A brief description of what the tool does.
	Description param.Field[string] `json:"description,required"`
	// The name of the tool. More descriptive the better.
	Name param.Field[string] `json:"name,required"`
	// Schema defining the parameters accepted by the tool.
	Parameters param.Field[AccountAIRunCfMistralExecuteMistral7bInstructV0_2LoraParamsBodyMessagesToolsObjectParameters] `json:"parameters,required"`
}

func (r AccountAIRunCfMistralExecuteMistral7bInstructV0_2LoraParamsBodyMessagesToolsObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfMistralExecuteMistral7bInstructV0_2LoraParamsBodyMessagesToolsObject) implementsAccountAIRunCfMistralExecuteMistral7bInstructV0_2LoraParamsBodyMessagesToolUnion() {
}

// Schema defining the parameters accepted by the tool.
type AccountAIRunCfMistralExecuteMistral7bInstructV0_2LoraParamsBodyMessagesToolsObjectParameters struct {
	// Definitions of each parameter.
	Properties param.Field[map[string]AccountAIRunCfMistralExecuteMistral7bInstructV0_2LoraParamsBodyMessagesToolsObjectParametersProperties] `json:"properties,required"`
	// The type of the parameters object (usually 'object').
	Type param.Field[string] `json:"type,required"`
	// List of required parameter names.
	Required param.Field[[]string] `json:"required"`
}

func (r AccountAIRunCfMistralExecuteMistral7bInstructV0_2LoraParamsBodyMessagesToolsObjectParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfMistralExecuteMistral7bInstructV0_2LoraParamsBodyMessagesToolsObjectParametersProperties struct {
	// A description of the expected parameter.
	Description param.Field[string] `json:"description,required"`
	// The data type of the parameter.
	Type param.Field[string] `json:"type,required"`
}

func (r AccountAIRunCfMistralExecuteMistral7bInstructV0_2LoraParamsBodyMessagesToolsObjectParametersProperties) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
