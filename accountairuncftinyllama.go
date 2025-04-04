// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"

	"github.com/stainless-sdks/cf-rex-go/internal/apijson"
	"github.com/stainless-sdks/cf-rex-go/internal/apiquery"
	"github.com/stainless-sdks/cf-rex-go/internal/param"
	"github.com/stainless-sdks/cf-rex-go/internal/requestconfig"
	"github.com/stainless-sdks/cf-rex-go/option"
	"github.com/stainless-sdks/cf-rex-go/shared"
	"github.com/tidwall/gjson"
)

// AccountAIRunCfTinyllamaService contains methods and other services that help
// with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountAIRunCfTinyllamaService] method instead.
type AccountAIRunCfTinyllamaService struct {
	Options []option.RequestOption
}

// NewAccountAIRunCfTinyllamaService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountAIRunCfTinyllamaService(opts ...option.RequestOption) (r *AccountAIRunCfTinyllamaService) {
	r = &AccountAIRunCfTinyllamaService{}
	r.Options = opts
	return
}

// Execute @cf/tinyllama/tinyllama-1.1b-chat-v1.0 model.
func (r *AccountAIRunCfTinyllamaService) ExecuteTinyllama1_1bChatV1_0(ctx context.Context, accountID string, params AccountAIRunCfTinyllamaExecuteTinyllama1_1bChatV1_0Params, opts ...option.RequestOption) (res *AccountAIRunCfTinyllamaExecuteTinyllama1_1bChatV1_0Response, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai/run/@cf/tinyllama/tinyllama-1.1b-chat-v1.0", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type AccountAIRunCfTinyllamaExecuteTinyllama1_1bChatV1_0Response struct {
	Result  AccountAIRunCfTinyllamaExecuteTinyllama1_1bChatV1_0ResponseResultUnion `json:"result" format:"binary"`
	Success bool                                                                   `json:"success"`
	JSON    accountAIRunCfTinyllamaExecuteTinyllama1_1bChatV1_0ResponseJSON        `json:"-"`
}

// accountAIRunCfTinyllamaExecuteTinyllama1_1bChatV1_0ResponseJSON contains the
// JSON metadata for the struct
// [AccountAIRunCfTinyllamaExecuteTinyllama1_1bChatV1_0Response]
type accountAIRunCfTinyllamaExecuteTinyllama1_1bChatV1_0ResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIRunCfTinyllamaExecuteTinyllama1_1bChatV1_0Response) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunCfTinyllamaExecuteTinyllama1_1bChatV1_0ResponseJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by
// [AccountAIRunCfTinyllamaExecuteTinyllama1_1bChatV1_0ResponseResultObject] or
// [shared.UnionString].
type AccountAIRunCfTinyllamaExecuteTinyllama1_1bChatV1_0ResponseResultUnion interface {
	ImplementsAccountAIRunCfTinyllamaExecuteTinyllama1_1bChatV1_0ResponseResultUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccountAIRunCfTinyllamaExecuteTinyllama1_1bChatV1_0ResponseResultUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccountAIRunCfTinyllamaExecuteTinyllama1_1bChatV1_0ResponseResultObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type AccountAIRunCfTinyllamaExecuteTinyllama1_1bChatV1_0ResponseResultObject struct {
	// The generated text response from the model
	Response string `json:"response,required"`
	// An array of tool calls requests made during the response generation
	ToolCalls []AccountAIRunCfTinyllamaExecuteTinyllama1_1bChatV1_0ResponseResultObjectToolCall `json:"tool_calls"`
	// Usage statistics for the inference request
	Usage AccountAIRunCfTinyllamaExecuteTinyllama1_1bChatV1_0ResponseResultObjectUsage `json:"usage"`
	JSON  accountAIRunCfTinyllamaExecuteTinyllama1_1bChatV1_0ResponseResultObjectJSON  `json:"-"`
}

// accountAIRunCfTinyllamaExecuteTinyllama1_1bChatV1_0ResponseResultObjectJSON
// contains the JSON metadata for the struct
// [AccountAIRunCfTinyllamaExecuteTinyllama1_1bChatV1_0ResponseResultObject]
type accountAIRunCfTinyllamaExecuteTinyllama1_1bChatV1_0ResponseResultObjectJSON struct {
	Response    apijson.Field
	ToolCalls   apijson.Field
	Usage       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIRunCfTinyllamaExecuteTinyllama1_1bChatV1_0ResponseResultObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunCfTinyllamaExecuteTinyllama1_1bChatV1_0ResponseResultObjectJSON) RawJSON() string {
	return r.raw
}

func (r AccountAIRunCfTinyllamaExecuteTinyllama1_1bChatV1_0ResponseResultObject) ImplementsAccountAIRunCfTinyllamaExecuteTinyllama1_1bChatV1_0ResponseResultUnion() {
}

type AccountAIRunCfTinyllamaExecuteTinyllama1_1bChatV1_0ResponseResultObjectToolCall struct {
	// The arguments passed to be passed to the tool call request
	Arguments interface{} `json:"arguments"`
	// The name of the tool to be called
	Name string                                                                              `json:"name"`
	JSON accountAIRunCfTinyllamaExecuteTinyllama1_1bChatV1_0ResponseResultObjectToolCallJSON `json:"-"`
}

// accountAIRunCfTinyllamaExecuteTinyllama1_1bChatV1_0ResponseResultObjectToolCallJSON
// contains the JSON metadata for the struct
// [AccountAIRunCfTinyllamaExecuteTinyllama1_1bChatV1_0ResponseResultObjectToolCall]
type accountAIRunCfTinyllamaExecuteTinyllama1_1bChatV1_0ResponseResultObjectToolCallJSON struct {
	Arguments   apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIRunCfTinyllamaExecuteTinyllama1_1bChatV1_0ResponseResultObjectToolCall) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunCfTinyllamaExecuteTinyllama1_1bChatV1_0ResponseResultObjectToolCallJSON) RawJSON() string {
	return r.raw
}

// Usage statistics for the inference request
type AccountAIRunCfTinyllamaExecuteTinyllama1_1bChatV1_0ResponseResultObjectUsage struct {
	// Total number of tokens in output
	CompletionTokens float64 `json:"completion_tokens"`
	// Total number of tokens in input
	PromptTokens float64 `json:"prompt_tokens"`
	// Total number of input and output tokens
	TotalTokens float64                                                                          `json:"total_tokens"`
	JSON        accountAIRunCfTinyllamaExecuteTinyllama1_1bChatV1_0ResponseResultObjectUsageJSON `json:"-"`
}

// accountAIRunCfTinyllamaExecuteTinyllama1_1bChatV1_0ResponseResultObjectUsageJSON
// contains the JSON metadata for the struct
// [AccountAIRunCfTinyllamaExecuteTinyllama1_1bChatV1_0ResponseResultObjectUsage]
type accountAIRunCfTinyllamaExecuteTinyllama1_1bChatV1_0ResponseResultObjectUsageJSON struct {
	CompletionTokens apijson.Field
	PromptTokens     apijson.Field
	TotalTokens      apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountAIRunCfTinyllamaExecuteTinyllama1_1bChatV1_0ResponseResultObjectUsage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunCfTinyllamaExecuteTinyllama1_1bChatV1_0ResponseResultObjectUsageJSON) RawJSON() string {
	return r.raw
}

type AccountAIRunCfTinyllamaExecuteTinyllama1_1bChatV1_0Params struct {
	QueueRequest param.Field[string]                                                `query:"queueRequest"`
	Body         AccountAIRunCfTinyllamaExecuteTinyllama1_1bChatV1_0ParamsBodyUnion `json:"body"`
}

func (r AccountAIRunCfTinyllamaExecuteTinyllama1_1bChatV1_0Params) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

// URLQuery serializes
// [AccountAIRunCfTinyllamaExecuteTinyllama1_1bChatV1_0Params]'s query parameters
// as `url.Values`.
func (r AccountAIRunCfTinyllamaExecuteTinyllama1_1bChatV1_0Params) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountAIRunCfTinyllamaExecuteTinyllama1_1bChatV1_0ParamsBody struct {
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

func (r AccountAIRunCfTinyllamaExecuteTinyllama1_1bChatV1_0ParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfTinyllamaExecuteTinyllama1_1bChatV1_0ParamsBody) implementsAccountAIRunCfTinyllamaExecuteTinyllama1_1bChatV1_0ParamsBodyUnion() {
}

// Satisfied by
// [AccountAIRunCfTinyllamaExecuteTinyllama1_1bChatV1_0ParamsBodyPrompt],
// [AccountAIRunCfTinyllamaExecuteTinyllama1_1bChatV1_0ParamsBodyMessages],
// [AccountAIRunCfTinyllamaExecuteTinyllama1_1bChatV1_0ParamsBody].
type AccountAIRunCfTinyllamaExecuteTinyllama1_1bChatV1_0ParamsBodyUnion interface {
	implementsAccountAIRunCfTinyllamaExecuteTinyllama1_1bChatV1_0ParamsBodyUnion()
}

type AccountAIRunCfTinyllamaExecuteTinyllama1_1bChatV1_0ParamsBodyPrompt struct {
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
	RepetitionPenalty param.Field[float64]                                                                           `json:"repetition_penalty"`
	ResponseFormat    param.Field[AccountAIRunCfTinyllamaExecuteTinyllama1_1bChatV1_0ParamsBodyPromptResponseFormat] `json:"response_format"`
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

func (r AccountAIRunCfTinyllamaExecuteTinyllama1_1bChatV1_0ParamsBodyPrompt) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfTinyllamaExecuteTinyllama1_1bChatV1_0ParamsBodyPrompt) implementsAccountAIRunCfTinyllamaExecuteTinyllama1_1bChatV1_0ParamsBodyUnion() {
}

type AccountAIRunCfTinyllamaExecuteTinyllama1_1bChatV1_0ParamsBodyPromptResponseFormat struct {
	JsonSchema param.Field[interface{}]                                                                           `json:"json_schema"`
	Type       param.Field[AccountAIRunCfTinyllamaExecuteTinyllama1_1bChatV1_0ParamsBodyPromptResponseFormatType] `json:"type"`
}

func (r AccountAIRunCfTinyllamaExecuteTinyllama1_1bChatV1_0ParamsBodyPromptResponseFormat) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfTinyllamaExecuteTinyllama1_1bChatV1_0ParamsBodyPromptResponseFormatType string

const (
	AccountAIRunCfTinyllamaExecuteTinyllama1_1bChatV1_0ParamsBodyPromptResponseFormatTypeJsonObject AccountAIRunCfTinyllamaExecuteTinyllama1_1bChatV1_0ParamsBodyPromptResponseFormatType = "json_object"
	AccountAIRunCfTinyllamaExecuteTinyllama1_1bChatV1_0ParamsBodyPromptResponseFormatTypeJsonSchema AccountAIRunCfTinyllamaExecuteTinyllama1_1bChatV1_0ParamsBodyPromptResponseFormatType = "json_schema"
)

func (r AccountAIRunCfTinyllamaExecuteTinyllama1_1bChatV1_0ParamsBodyPromptResponseFormatType) IsKnown() bool {
	switch r {
	case AccountAIRunCfTinyllamaExecuteTinyllama1_1bChatV1_0ParamsBodyPromptResponseFormatTypeJsonObject, AccountAIRunCfTinyllamaExecuteTinyllama1_1bChatV1_0ParamsBodyPromptResponseFormatTypeJsonSchema:
		return true
	}
	return false
}

type AccountAIRunCfTinyllamaExecuteTinyllama1_1bChatV1_0ParamsBodyMessages struct {
	// An array of message objects representing the conversation history.
	Messages param.Field[[]AccountAIRunCfTinyllamaExecuteTinyllama1_1bChatV1_0ParamsBodyMessagesMessage] `json:"messages,required"`
	// Decreases the likelihood of the model repeating the same lines verbatim.
	FrequencyPenalty param.Field[float64]                                                                         `json:"frequency_penalty"`
	Functions        param.Field[[]AccountAIRunCfTinyllamaExecuteTinyllama1_1bChatV1_0ParamsBodyMessagesFunction] `json:"functions"`
	// The maximum number of tokens to generate in the response.
	MaxTokens param.Field[int64] `json:"max_tokens"`
	// Increases the likelihood of the model introducing new topics.
	PresencePenalty param.Field[float64] `json:"presence_penalty"`
	// If true, a chat template is not applied and you must adhere to the specific
	// model's expected formatting.
	Raw param.Field[bool] `json:"raw"`
	// Penalty for repeated tokens; higher values discourage repetition.
	RepetitionPenalty param.Field[float64]                                                                             `json:"repetition_penalty"`
	ResponseFormat    param.Field[AccountAIRunCfTinyllamaExecuteTinyllama1_1bChatV1_0ParamsBodyMessagesResponseFormat] `json:"response_format"`
	// Random seed for reproducibility of the generation.
	Seed param.Field[int64] `json:"seed"`
	// If true, the response will be streamed back incrementally using SSE, Server Sent
	// Events.
	Stream param.Field[bool] `json:"stream"`
	// Controls the randomness of the output; higher values produce more random
	// results.
	Temperature param.Field[float64] `json:"temperature"`
	// A list of tools available for the assistant to use.
	Tools param.Field[[]AccountAIRunCfTinyllamaExecuteTinyllama1_1bChatV1_0ParamsBodyMessagesToolUnion] `json:"tools"`
	// Limits the AI to choose from the top 'k' most probable words. Lower values make
	// responses more focused; higher values introduce more variety and potential
	// surprises.
	TopK param.Field[int64] `json:"top_k"`
	// Adjusts the creativity of the AI's responses by controlling how many possible
	// words it considers. Lower values make outputs more predictable; higher values
	// allow for more varied and creative responses.
	TopP param.Field[float64] `json:"top_p"`
}

func (r AccountAIRunCfTinyllamaExecuteTinyllama1_1bChatV1_0ParamsBodyMessages) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfTinyllamaExecuteTinyllama1_1bChatV1_0ParamsBodyMessages) implementsAccountAIRunCfTinyllamaExecuteTinyllama1_1bChatV1_0ParamsBodyUnion() {
}

type AccountAIRunCfTinyllamaExecuteTinyllama1_1bChatV1_0ParamsBodyMessagesMessage struct {
	// The content of the message as a string.
	Content param.Field[string] `json:"content,required"`
	// The role of the message sender (e.g., 'user', 'assistant', 'system', 'tool').
	Role param.Field[string] `json:"role,required"`
}

func (r AccountAIRunCfTinyllamaExecuteTinyllama1_1bChatV1_0ParamsBodyMessagesMessage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfTinyllamaExecuteTinyllama1_1bChatV1_0ParamsBodyMessagesFunction struct {
	Code param.Field[string] `json:"code,required"`
	Name param.Field[string] `json:"name,required"`
}

func (r AccountAIRunCfTinyllamaExecuteTinyllama1_1bChatV1_0ParamsBodyMessagesFunction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfTinyllamaExecuteTinyllama1_1bChatV1_0ParamsBodyMessagesResponseFormat struct {
	JsonSchema param.Field[interface{}]                                                                             `json:"json_schema"`
	Type       param.Field[AccountAIRunCfTinyllamaExecuteTinyllama1_1bChatV1_0ParamsBodyMessagesResponseFormatType] `json:"type"`
}

func (r AccountAIRunCfTinyllamaExecuteTinyllama1_1bChatV1_0ParamsBodyMessagesResponseFormat) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfTinyllamaExecuteTinyllama1_1bChatV1_0ParamsBodyMessagesResponseFormatType string

const (
	AccountAIRunCfTinyllamaExecuteTinyllama1_1bChatV1_0ParamsBodyMessagesResponseFormatTypeJsonObject AccountAIRunCfTinyllamaExecuteTinyllama1_1bChatV1_0ParamsBodyMessagesResponseFormatType = "json_object"
	AccountAIRunCfTinyllamaExecuteTinyllama1_1bChatV1_0ParamsBodyMessagesResponseFormatTypeJsonSchema AccountAIRunCfTinyllamaExecuteTinyllama1_1bChatV1_0ParamsBodyMessagesResponseFormatType = "json_schema"
)

func (r AccountAIRunCfTinyllamaExecuteTinyllama1_1bChatV1_0ParamsBodyMessagesResponseFormatType) IsKnown() bool {
	switch r {
	case AccountAIRunCfTinyllamaExecuteTinyllama1_1bChatV1_0ParamsBodyMessagesResponseFormatTypeJsonObject, AccountAIRunCfTinyllamaExecuteTinyllama1_1bChatV1_0ParamsBodyMessagesResponseFormatTypeJsonSchema:
		return true
	}
	return false
}

type AccountAIRunCfTinyllamaExecuteTinyllama1_1bChatV1_0ParamsBodyMessagesTool struct {
	// A brief description of what the tool does.
	Description param.Field[string]      `json:"description"`
	Function    param.Field[interface{}] `json:"function"`
	// The name of the tool. More descriptive the better.
	Name       param.Field[string]      `json:"name"`
	Parameters param.Field[interface{}] `json:"parameters"`
	// Specifies the type of tool (e.g., 'function').
	Type param.Field[string] `json:"type"`
}

func (r AccountAIRunCfTinyllamaExecuteTinyllama1_1bChatV1_0ParamsBodyMessagesTool) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfTinyllamaExecuteTinyllama1_1bChatV1_0ParamsBodyMessagesTool) implementsAccountAIRunCfTinyllamaExecuteTinyllama1_1bChatV1_0ParamsBodyMessagesToolUnion() {
}

// Satisfied by
// [AccountAIRunCfTinyllamaExecuteTinyllama1_1bChatV1_0ParamsBodyMessagesToolsObject],
// [AccountAIRunCfTinyllamaExecuteTinyllama1_1bChatV1_0ParamsBodyMessagesToolsObject],
// [AccountAIRunCfTinyllamaExecuteTinyllama1_1bChatV1_0ParamsBodyMessagesTool].
type AccountAIRunCfTinyllamaExecuteTinyllama1_1bChatV1_0ParamsBodyMessagesToolUnion interface {
	implementsAccountAIRunCfTinyllamaExecuteTinyllama1_1bChatV1_0ParamsBodyMessagesToolUnion()
}

type AccountAIRunCfTinyllamaExecuteTinyllama1_1bChatV1_0ParamsBodyMessagesToolsObject struct {
	// A brief description of what the tool does.
	Description param.Field[string] `json:"description,required"`
	// The name of the tool. More descriptive the better.
	Name param.Field[string] `json:"name,required"`
	// Schema defining the parameters accepted by the tool.
	Parameters param.Field[AccountAIRunCfTinyllamaExecuteTinyllama1_1bChatV1_0ParamsBodyMessagesToolsObjectParameters] `json:"parameters,required"`
}

func (r AccountAIRunCfTinyllamaExecuteTinyllama1_1bChatV1_0ParamsBodyMessagesToolsObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfTinyllamaExecuteTinyllama1_1bChatV1_0ParamsBodyMessagesToolsObject) implementsAccountAIRunCfTinyllamaExecuteTinyllama1_1bChatV1_0ParamsBodyMessagesToolUnion() {
}

// Schema defining the parameters accepted by the tool.
type AccountAIRunCfTinyllamaExecuteTinyllama1_1bChatV1_0ParamsBodyMessagesToolsObjectParameters struct {
	// Definitions of each parameter.
	Properties param.Field[map[string]AccountAIRunCfTinyllamaExecuteTinyllama1_1bChatV1_0ParamsBodyMessagesToolsObjectParametersProperties] `json:"properties,required"`
	// The type of the parameters object (usually 'object').
	Type param.Field[string] `json:"type,required"`
	// List of required parameter names.
	Required param.Field[[]string] `json:"required"`
}

func (r AccountAIRunCfTinyllamaExecuteTinyllama1_1bChatV1_0ParamsBodyMessagesToolsObjectParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfTinyllamaExecuteTinyllama1_1bChatV1_0ParamsBodyMessagesToolsObjectParametersProperties struct {
	// A description of the expected parameter.
	Description param.Field[string] `json:"description,required"`
	// The data type of the parameter.
	Type param.Field[string] `json:"type,required"`
}

func (r AccountAIRunCfTinyllamaExecuteTinyllama1_1bChatV1_0ParamsBodyMessagesToolsObjectParametersProperties) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
