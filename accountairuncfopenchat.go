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

// AccountAIRunCfOpenchatService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountAIRunCfOpenchatService] method instead.
type AccountAIRunCfOpenchatService struct {
	Options []option.RequestOption
}

// NewAccountAIRunCfOpenchatService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountAIRunCfOpenchatService(opts ...option.RequestOption) (r *AccountAIRunCfOpenchatService) {
	r = &AccountAIRunCfOpenchatService{}
	r.Options = opts
	return
}

// Execute @cf/openchat/openchat-3.5-0106 model.
func (r *AccountAIRunCfOpenchatService) ExecuteOpenchat3_5_0106(ctx context.Context, accountID string, params AccountAIRunCfOpenchatExecuteOpenchat3_5_0106Params, opts ...option.RequestOption) (res *AccountAIRunCfOpenchatExecuteOpenchat3_5_0106Response, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai/run/@cf/openchat/openchat-3.5-0106", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type AccountAIRunCfOpenchatExecuteOpenchat3_5_0106Response struct {
	Result  AccountAIRunCfOpenchatExecuteOpenchat3_5_0106ResponseResultUnion `json:"result" format:"binary"`
	Success bool                                                             `json:"success"`
	JSON    accountAIRunCfOpenchatExecuteOpenchat3_5_0106ResponseJSON        `json:"-"`
}

// accountAIRunCfOpenchatExecuteOpenchat3_5_0106ResponseJSON contains the JSON
// metadata for the struct [AccountAIRunCfOpenchatExecuteOpenchat3_5_0106Response]
type accountAIRunCfOpenchatExecuteOpenchat3_5_0106ResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIRunCfOpenchatExecuteOpenchat3_5_0106Response) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunCfOpenchatExecuteOpenchat3_5_0106ResponseJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by
// [AccountAIRunCfOpenchatExecuteOpenchat3_5_0106ResponseResultObject] or
// [shared.UnionString].
type AccountAIRunCfOpenchatExecuteOpenchat3_5_0106ResponseResultUnion interface {
	ImplementsAccountAIRunCfOpenchatExecuteOpenchat3_5_0106ResponseResultUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccountAIRunCfOpenchatExecuteOpenchat3_5_0106ResponseResultUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccountAIRunCfOpenchatExecuteOpenchat3_5_0106ResponseResultObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type AccountAIRunCfOpenchatExecuteOpenchat3_5_0106ResponseResultObject struct {
	// The generated text response from the model
	Response string `json:"response,required"`
	// An array of tool calls requests made during the response generation
	ToolCalls []AccountAIRunCfOpenchatExecuteOpenchat3_5_0106ResponseResultObjectToolCall `json:"tool_calls"`
	// Usage statistics for the inference request
	Usage AccountAIRunCfOpenchatExecuteOpenchat3_5_0106ResponseResultObjectUsage `json:"usage"`
	JSON  accountAIRunCfOpenchatExecuteOpenchat3_5_0106ResponseResultObjectJSON  `json:"-"`
}

// accountAIRunCfOpenchatExecuteOpenchat3_5_0106ResponseResultObjectJSON contains
// the JSON metadata for the struct
// [AccountAIRunCfOpenchatExecuteOpenchat3_5_0106ResponseResultObject]
type accountAIRunCfOpenchatExecuteOpenchat3_5_0106ResponseResultObjectJSON struct {
	Response    apijson.Field
	ToolCalls   apijson.Field
	Usage       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIRunCfOpenchatExecuteOpenchat3_5_0106ResponseResultObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunCfOpenchatExecuteOpenchat3_5_0106ResponseResultObjectJSON) RawJSON() string {
	return r.raw
}

func (r AccountAIRunCfOpenchatExecuteOpenchat3_5_0106ResponseResultObject) ImplementsAccountAIRunCfOpenchatExecuteOpenchat3_5_0106ResponseResultUnion() {
}

type AccountAIRunCfOpenchatExecuteOpenchat3_5_0106ResponseResultObjectToolCall struct {
	// The arguments passed to be passed to the tool call request
	Arguments interface{} `json:"arguments"`
	// The name of the tool to be called
	Name string                                                                        `json:"name"`
	JSON accountAIRunCfOpenchatExecuteOpenchat3_5_0106ResponseResultObjectToolCallJSON `json:"-"`
}

// accountAIRunCfOpenchatExecuteOpenchat3_5_0106ResponseResultObjectToolCallJSON
// contains the JSON metadata for the struct
// [AccountAIRunCfOpenchatExecuteOpenchat3_5_0106ResponseResultObjectToolCall]
type accountAIRunCfOpenchatExecuteOpenchat3_5_0106ResponseResultObjectToolCallJSON struct {
	Arguments   apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIRunCfOpenchatExecuteOpenchat3_5_0106ResponseResultObjectToolCall) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunCfOpenchatExecuteOpenchat3_5_0106ResponseResultObjectToolCallJSON) RawJSON() string {
	return r.raw
}

// Usage statistics for the inference request
type AccountAIRunCfOpenchatExecuteOpenchat3_5_0106ResponseResultObjectUsage struct {
	// Total number of tokens in output
	CompletionTokens float64 `json:"completion_tokens"`
	// Total number of tokens in input
	PromptTokens float64 `json:"prompt_tokens"`
	// Total number of input and output tokens
	TotalTokens float64                                                                    `json:"total_tokens"`
	JSON        accountAIRunCfOpenchatExecuteOpenchat3_5_0106ResponseResultObjectUsageJSON `json:"-"`
}

// accountAIRunCfOpenchatExecuteOpenchat3_5_0106ResponseResultObjectUsageJSON
// contains the JSON metadata for the struct
// [AccountAIRunCfOpenchatExecuteOpenchat3_5_0106ResponseResultObjectUsage]
type accountAIRunCfOpenchatExecuteOpenchat3_5_0106ResponseResultObjectUsageJSON struct {
	CompletionTokens apijson.Field
	PromptTokens     apijson.Field
	TotalTokens      apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountAIRunCfOpenchatExecuteOpenchat3_5_0106ResponseResultObjectUsage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunCfOpenchatExecuteOpenchat3_5_0106ResponseResultObjectUsageJSON) RawJSON() string {
	return r.raw
}

type AccountAIRunCfOpenchatExecuteOpenchat3_5_0106Params struct {
	QueueRequest param.Field[string]                                          `query:"queueRequest"`
	Body         AccountAIRunCfOpenchatExecuteOpenchat3_5_0106ParamsBodyUnion `json:"body"`
}

func (r AccountAIRunCfOpenchatExecuteOpenchat3_5_0106Params) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

// URLQuery serializes [AccountAIRunCfOpenchatExecuteOpenchat3_5_0106Params]'s
// query parameters as `url.Values`.
func (r AccountAIRunCfOpenchatExecuteOpenchat3_5_0106Params) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountAIRunCfOpenchatExecuteOpenchat3_5_0106ParamsBody struct {
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

func (r AccountAIRunCfOpenchatExecuteOpenchat3_5_0106ParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfOpenchatExecuteOpenchat3_5_0106ParamsBody) implementsAccountAIRunCfOpenchatExecuteOpenchat3_5_0106ParamsBodyUnion() {
}

// Satisfied by [AccountAIRunCfOpenchatExecuteOpenchat3_5_0106ParamsBodyPrompt],
// [AccountAIRunCfOpenchatExecuteOpenchat3_5_0106ParamsBodyMessages],
// [AccountAIRunCfOpenchatExecuteOpenchat3_5_0106ParamsBody].
type AccountAIRunCfOpenchatExecuteOpenchat3_5_0106ParamsBodyUnion interface {
	implementsAccountAIRunCfOpenchatExecuteOpenchat3_5_0106ParamsBodyUnion()
}

type AccountAIRunCfOpenchatExecuteOpenchat3_5_0106ParamsBodyPrompt struct {
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
	RepetitionPenalty param.Field[float64]                                                                     `json:"repetition_penalty"`
	ResponseFormat    param.Field[AccountAIRunCfOpenchatExecuteOpenchat3_5_0106ParamsBodyPromptResponseFormat] `json:"response_format"`
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

func (r AccountAIRunCfOpenchatExecuteOpenchat3_5_0106ParamsBodyPrompt) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfOpenchatExecuteOpenchat3_5_0106ParamsBodyPrompt) implementsAccountAIRunCfOpenchatExecuteOpenchat3_5_0106ParamsBodyUnion() {
}

type AccountAIRunCfOpenchatExecuteOpenchat3_5_0106ParamsBodyPromptResponseFormat struct {
	JsonSchema param.Field[interface{}]                                                                     `json:"json_schema"`
	Type       param.Field[AccountAIRunCfOpenchatExecuteOpenchat3_5_0106ParamsBodyPromptResponseFormatType] `json:"type"`
}

func (r AccountAIRunCfOpenchatExecuteOpenchat3_5_0106ParamsBodyPromptResponseFormat) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfOpenchatExecuteOpenchat3_5_0106ParamsBodyPromptResponseFormatType string

const (
	AccountAIRunCfOpenchatExecuteOpenchat3_5_0106ParamsBodyPromptResponseFormatTypeJsonObject AccountAIRunCfOpenchatExecuteOpenchat3_5_0106ParamsBodyPromptResponseFormatType = "json_object"
	AccountAIRunCfOpenchatExecuteOpenchat3_5_0106ParamsBodyPromptResponseFormatTypeJsonSchema AccountAIRunCfOpenchatExecuteOpenchat3_5_0106ParamsBodyPromptResponseFormatType = "json_schema"
)

func (r AccountAIRunCfOpenchatExecuteOpenchat3_5_0106ParamsBodyPromptResponseFormatType) IsKnown() bool {
	switch r {
	case AccountAIRunCfOpenchatExecuteOpenchat3_5_0106ParamsBodyPromptResponseFormatTypeJsonObject, AccountAIRunCfOpenchatExecuteOpenchat3_5_0106ParamsBodyPromptResponseFormatTypeJsonSchema:
		return true
	}
	return false
}

type AccountAIRunCfOpenchatExecuteOpenchat3_5_0106ParamsBodyMessages struct {
	// An array of message objects representing the conversation history.
	Messages param.Field[[]AccountAIRunCfOpenchatExecuteOpenchat3_5_0106ParamsBodyMessagesMessage] `json:"messages,required"`
	// Decreases the likelihood of the model repeating the same lines verbatim.
	FrequencyPenalty param.Field[float64]                                                                   `json:"frequency_penalty"`
	Functions        param.Field[[]AccountAIRunCfOpenchatExecuteOpenchat3_5_0106ParamsBodyMessagesFunction] `json:"functions"`
	// The maximum number of tokens to generate in the response.
	MaxTokens param.Field[int64] `json:"max_tokens"`
	// Increases the likelihood of the model introducing new topics.
	PresencePenalty param.Field[float64] `json:"presence_penalty"`
	// If true, a chat template is not applied and you must adhere to the specific
	// model's expected formatting.
	Raw param.Field[bool] `json:"raw"`
	// Penalty for repeated tokens; higher values discourage repetition.
	RepetitionPenalty param.Field[float64]                                                                       `json:"repetition_penalty"`
	ResponseFormat    param.Field[AccountAIRunCfOpenchatExecuteOpenchat3_5_0106ParamsBodyMessagesResponseFormat] `json:"response_format"`
	// Random seed for reproducibility of the generation.
	Seed param.Field[int64] `json:"seed"`
	// If true, the response will be streamed back incrementally using SSE, Server Sent
	// Events.
	Stream param.Field[bool] `json:"stream"`
	// Controls the randomness of the output; higher values produce more random
	// results.
	Temperature param.Field[float64] `json:"temperature"`
	// A list of tools available for the assistant to use.
	Tools param.Field[[]AccountAIRunCfOpenchatExecuteOpenchat3_5_0106ParamsBodyMessagesToolUnion] `json:"tools"`
	// Limits the AI to choose from the top 'k' most probable words. Lower values make
	// responses more focused; higher values introduce more variety and potential
	// surprises.
	TopK param.Field[int64] `json:"top_k"`
	// Adjusts the creativity of the AI's responses by controlling how many possible
	// words it considers. Lower values make outputs more predictable; higher values
	// allow for more varied and creative responses.
	TopP param.Field[float64] `json:"top_p"`
}

func (r AccountAIRunCfOpenchatExecuteOpenchat3_5_0106ParamsBodyMessages) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfOpenchatExecuteOpenchat3_5_0106ParamsBodyMessages) implementsAccountAIRunCfOpenchatExecuteOpenchat3_5_0106ParamsBodyUnion() {
}

type AccountAIRunCfOpenchatExecuteOpenchat3_5_0106ParamsBodyMessagesMessage struct {
	// The content of the message as a string.
	Content param.Field[string] `json:"content,required"`
	// The role of the message sender (e.g., 'user', 'assistant', 'system', 'tool').
	Role param.Field[string] `json:"role,required"`
}

func (r AccountAIRunCfOpenchatExecuteOpenchat3_5_0106ParamsBodyMessagesMessage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfOpenchatExecuteOpenchat3_5_0106ParamsBodyMessagesFunction struct {
	Code param.Field[string] `json:"code,required"`
	Name param.Field[string] `json:"name,required"`
}

func (r AccountAIRunCfOpenchatExecuteOpenchat3_5_0106ParamsBodyMessagesFunction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfOpenchatExecuteOpenchat3_5_0106ParamsBodyMessagesResponseFormat struct {
	JsonSchema param.Field[interface{}]                                                                       `json:"json_schema"`
	Type       param.Field[AccountAIRunCfOpenchatExecuteOpenchat3_5_0106ParamsBodyMessagesResponseFormatType] `json:"type"`
}

func (r AccountAIRunCfOpenchatExecuteOpenchat3_5_0106ParamsBodyMessagesResponseFormat) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfOpenchatExecuteOpenchat3_5_0106ParamsBodyMessagesResponseFormatType string

const (
	AccountAIRunCfOpenchatExecuteOpenchat3_5_0106ParamsBodyMessagesResponseFormatTypeJsonObject AccountAIRunCfOpenchatExecuteOpenchat3_5_0106ParamsBodyMessagesResponseFormatType = "json_object"
	AccountAIRunCfOpenchatExecuteOpenchat3_5_0106ParamsBodyMessagesResponseFormatTypeJsonSchema AccountAIRunCfOpenchatExecuteOpenchat3_5_0106ParamsBodyMessagesResponseFormatType = "json_schema"
)

func (r AccountAIRunCfOpenchatExecuteOpenchat3_5_0106ParamsBodyMessagesResponseFormatType) IsKnown() bool {
	switch r {
	case AccountAIRunCfOpenchatExecuteOpenchat3_5_0106ParamsBodyMessagesResponseFormatTypeJsonObject, AccountAIRunCfOpenchatExecuteOpenchat3_5_0106ParamsBodyMessagesResponseFormatTypeJsonSchema:
		return true
	}
	return false
}

type AccountAIRunCfOpenchatExecuteOpenchat3_5_0106ParamsBodyMessagesTool struct {
	// A brief description of what the tool does.
	Description param.Field[string]      `json:"description"`
	Function    param.Field[interface{}] `json:"function"`
	// The name of the tool. More descriptive the better.
	Name       param.Field[string]      `json:"name"`
	Parameters param.Field[interface{}] `json:"parameters"`
	// Specifies the type of tool (e.g., 'function').
	Type param.Field[string] `json:"type"`
}

func (r AccountAIRunCfOpenchatExecuteOpenchat3_5_0106ParamsBodyMessagesTool) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfOpenchatExecuteOpenchat3_5_0106ParamsBodyMessagesTool) implementsAccountAIRunCfOpenchatExecuteOpenchat3_5_0106ParamsBodyMessagesToolUnion() {
}

// Satisfied by
// [AccountAIRunCfOpenchatExecuteOpenchat3_5_0106ParamsBodyMessagesToolsObject],
// [AccountAIRunCfOpenchatExecuteOpenchat3_5_0106ParamsBodyMessagesToolsObject],
// [AccountAIRunCfOpenchatExecuteOpenchat3_5_0106ParamsBodyMessagesTool].
type AccountAIRunCfOpenchatExecuteOpenchat3_5_0106ParamsBodyMessagesToolUnion interface {
	implementsAccountAIRunCfOpenchatExecuteOpenchat3_5_0106ParamsBodyMessagesToolUnion()
}

type AccountAIRunCfOpenchatExecuteOpenchat3_5_0106ParamsBodyMessagesToolsObject struct {
	// A brief description of what the tool does.
	Description param.Field[string] `json:"description,required"`
	// The name of the tool. More descriptive the better.
	Name param.Field[string] `json:"name,required"`
	// Schema defining the parameters accepted by the tool.
	Parameters param.Field[AccountAIRunCfOpenchatExecuteOpenchat3_5_0106ParamsBodyMessagesToolsObjectParameters] `json:"parameters,required"`
}

func (r AccountAIRunCfOpenchatExecuteOpenchat3_5_0106ParamsBodyMessagesToolsObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfOpenchatExecuteOpenchat3_5_0106ParamsBodyMessagesToolsObject) implementsAccountAIRunCfOpenchatExecuteOpenchat3_5_0106ParamsBodyMessagesToolUnion() {
}

// Schema defining the parameters accepted by the tool.
type AccountAIRunCfOpenchatExecuteOpenchat3_5_0106ParamsBodyMessagesToolsObjectParameters struct {
	// Definitions of each parameter.
	Properties param.Field[map[string]AccountAIRunCfOpenchatExecuteOpenchat3_5_0106ParamsBodyMessagesToolsObjectParametersProperties] `json:"properties,required"`
	// The type of the parameters object (usually 'object').
	Type param.Field[string] `json:"type,required"`
	// List of required parameter names.
	Required param.Field[[]string] `json:"required"`
}

func (r AccountAIRunCfOpenchatExecuteOpenchat3_5_0106ParamsBodyMessagesToolsObjectParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfOpenchatExecuteOpenchat3_5_0106ParamsBodyMessagesToolsObjectParametersProperties struct {
	// A description of the expected parameter.
	Description param.Field[string] `json:"description,required"`
	// The data type of the parameter.
	Type param.Field[string] `json:"type,required"`
}

func (r AccountAIRunCfOpenchatExecuteOpenchat3_5_0106ParamsBodyMessagesToolsObjectParametersProperties) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
