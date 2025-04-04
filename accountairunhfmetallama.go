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

// AccountAIRunHfMetaLlamaService contains methods and other services that help
// with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountAIRunHfMetaLlamaService] method instead.
type AccountAIRunHfMetaLlamaService struct {
	Options []option.RequestOption
}

// NewAccountAIRunHfMetaLlamaService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountAIRunHfMetaLlamaService(opts ...option.RequestOption) (r *AccountAIRunHfMetaLlamaService) {
	r = &AccountAIRunHfMetaLlamaService{}
	r.Options = opts
	return
}

// Execute @hf/meta-llama/meta-llama-3-8b-instruct model.
func (r *AccountAIRunHfMetaLlamaService) ExecuteMetaLlama3_8bInstruct(ctx context.Context, accountID string, params AccountAIRunHfMetaLlamaExecuteMetaLlama3_8bInstructParams, opts ...option.RequestOption) (res *AccountAIRunHfMetaLlamaExecuteMetaLlama3_8bInstructResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai/run/@hf/meta-llama/meta-llama-3-8b-instruct", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type AccountAIRunHfMetaLlamaExecuteMetaLlama3_8bInstructResponse struct {
	Result  AccountAIRunHfMetaLlamaExecuteMetaLlama3_8bInstructResponseResultUnion `json:"result" format:"binary"`
	Success bool                                                                   `json:"success"`
	JSON    accountAIRunHfMetaLlamaExecuteMetaLlama3_8bInstructResponseJSON        `json:"-"`
}

// accountAIRunHfMetaLlamaExecuteMetaLlama3_8bInstructResponseJSON contains the
// JSON metadata for the struct
// [AccountAIRunHfMetaLlamaExecuteMetaLlama3_8bInstructResponse]
type accountAIRunHfMetaLlamaExecuteMetaLlama3_8bInstructResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIRunHfMetaLlamaExecuteMetaLlama3_8bInstructResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunHfMetaLlamaExecuteMetaLlama3_8bInstructResponseJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by
// [AccountAIRunHfMetaLlamaExecuteMetaLlama3_8bInstructResponseResultObject] or
// [shared.UnionString].
type AccountAIRunHfMetaLlamaExecuteMetaLlama3_8bInstructResponseResultUnion interface {
	ImplementsAccountAIRunHfMetaLlamaExecuteMetaLlama3_8bInstructResponseResultUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccountAIRunHfMetaLlamaExecuteMetaLlama3_8bInstructResponseResultUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccountAIRunHfMetaLlamaExecuteMetaLlama3_8bInstructResponseResultObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type AccountAIRunHfMetaLlamaExecuteMetaLlama3_8bInstructResponseResultObject struct {
	// The generated text response from the model
	Response string `json:"response,required"`
	// An array of tool calls requests made during the response generation
	ToolCalls []AccountAIRunHfMetaLlamaExecuteMetaLlama3_8bInstructResponseResultObjectToolCall `json:"tool_calls"`
	// Usage statistics for the inference request
	Usage AccountAIRunHfMetaLlamaExecuteMetaLlama3_8bInstructResponseResultObjectUsage `json:"usage"`
	JSON  accountAIRunHfMetaLlamaExecuteMetaLlama3_8bInstructResponseResultObjectJSON  `json:"-"`
}

// accountAIRunHfMetaLlamaExecuteMetaLlama3_8bInstructResponseResultObjectJSON
// contains the JSON metadata for the struct
// [AccountAIRunHfMetaLlamaExecuteMetaLlama3_8bInstructResponseResultObject]
type accountAIRunHfMetaLlamaExecuteMetaLlama3_8bInstructResponseResultObjectJSON struct {
	Response    apijson.Field
	ToolCalls   apijson.Field
	Usage       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIRunHfMetaLlamaExecuteMetaLlama3_8bInstructResponseResultObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunHfMetaLlamaExecuteMetaLlama3_8bInstructResponseResultObjectJSON) RawJSON() string {
	return r.raw
}

func (r AccountAIRunHfMetaLlamaExecuteMetaLlama3_8bInstructResponseResultObject) ImplementsAccountAIRunHfMetaLlamaExecuteMetaLlama3_8bInstructResponseResultUnion() {
}

type AccountAIRunHfMetaLlamaExecuteMetaLlama3_8bInstructResponseResultObjectToolCall struct {
	// The arguments passed to be passed to the tool call request
	Arguments interface{} `json:"arguments"`
	// The name of the tool to be called
	Name string                                                                              `json:"name"`
	JSON accountAIRunHfMetaLlamaExecuteMetaLlama3_8bInstructResponseResultObjectToolCallJSON `json:"-"`
}

// accountAIRunHfMetaLlamaExecuteMetaLlama3_8bInstructResponseResultObjectToolCallJSON
// contains the JSON metadata for the struct
// [AccountAIRunHfMetaLlamaExecuteMetaLlama3_8bInstructResponseResultObjectToolCall]
type accountAIRunHfMetaLlamaExecuteMetaLlama3_8bInstructResponseResultObjectToolCallJSON struct {
	Arguments   apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIRunHfMetaLlamaExecuteMetaLlama3_8bInstructResponseResultObjectToolCall) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunHfMetaLlamaExecuteMetaLlama3_8bInstructResponseResultObjectToolCallJSON) RawJSON() string {
	return r.raw
}

// Usage statistics for the inference request
type AccountAIRunHfMetaLlamaExecuteMetaLlama3_8bInstructResponseResultObjectUsage struct {
	// Total number of tokens in output
	CompletionTokens float64 `json:"completion_tokens"`
	// Total number of tokens in input
	PromptTokens float64 `json:"prompt_tokens"`
	// Total number of input and output tokens
	TotalTokens float64                                                                          `json:"total_tokens"`
	JSON        accountAIRunHfMetaLlamaExecuteMetaLlama3_8bInstructResponseResultObjectUsageJSON `json:"-"`
}

// accountAIRunHfMetaLlamaExecuteMetaLlama3_8bInstructResponseResultObjectUsageJSON
// contains the JSON metadata for the struct
// [AccountAIRunHfMetaLlamaExecuteMetaLlama3_8bInstructResponseResultObjectUsage]
type accountAIRunHfMetaLlamaExecuteMetaLlama3_8bInstructResponseResultObjectUsageJSON struct {
	CompletionTokens apijson.Field
	PromptTokens     apijson.Field
	TotalTokens      apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountAIRunHfMetaLlamaExecuteMetaLlama3_8bInstructResponseResultObjectUsage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunHfMetaLlamaExecuteMetaLlama3_8bInstructResponseResultObjectUsageJSON) RawJSON() string {
	return r.raw
}

type AccountAIRunHfMetaLlamaExecuteMetaLlama3_8bInstructParams struct {
	QueueRequest param.Field[string]                                                `query:"queueRequest"`
	Body         AccountAIRunHfMetaLlamaExecuteMetaLlama3_8bInstructParamsBodyUnion `json:"body"`
}

func (r AccountAIRunHfMetaLlamaExecuteMetaLlama3_8bInstructParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

// URLQuery serializes
// [AccountAIRunHfMetaLlamaExecuteMetaLlama3_8bInstructParams]'s query parameters
// as `url.Values`.
func (r AccountAIRunHfMetaLlamaExecuteMetaLlama3_8bInstructParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountAIRunHfMetaLlamaExecuteMetaLlama3_8bInstructParamsBody struct {
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

func (r AccountAIRunHfMetaLlamaExecuteMetaLlama3_8bInstructParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunHfMetaLlamaExecuteMetaLlama3_8bInstructParamsBody) implementsAccountAIRunHfMetaLlamaExecuteMetaLlama3_8bInstructParamsBodyUnion() {
}

// Satisfied by
// [AccountAIRunHfMetaLlamaExecuteMetaLlama3_8bInstructParamsBodyPrompt],
// [AccountAIRunHfMetaLlamaExecuteMetaLlama3_8bInstructParamsBodyMessages],
// [AccountAIRunHfMetaLlamaExecuteMetaLlama3_8bInstructParamsBody].
type AccountAIRunHfMetaLlamaExecuteMetaLlama3_8bInstructParamsBodyUnion interface {
	implementsAccountAIRunHfMetaLlamaExecuteMetaLlama3_8bInstructParamsBodyUnion()
}

type AccountAIRunHfMetaLlamaExecuteMetaLlama3_8bInstructParamsBodyPrompt struct {
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
	ResponseFormat    param.Field[AccountAIRunHfMetaLlamaExecuteMetaLlama3_8bInstructParamsBodyPromptResponseFormat] `json:"response_format"`
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

func (r AccountAIRunHfMetaLlamaExecuteMetaLlama3_8bInstructParamsBodyPrompt) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunHfMetaLlamaExecuteMetaLlama3_8bInstructParamsBodyPrompt) implementsAccountAIRunHfMetaLlamaExecuteMetaLlama3_8bInstructParamsBodyUnion() {
}

type AccountAIRunHfMetaLlamaExecuteMetaLlama3_8bInstructParamsBodyPromptResponseFormat struct {
	JsonSchema param.Field[interface{}]                                                                           `json:"json_schema"`
	Type       param.Field[AccountAIRunHfMetaLlamaExecuteMetaLlama3_8bInstructParamsBodyPromptResponseFormatType] `json:"type"`
}

func (r AccountAIRunHfMetaLlamaExecuteMetaLlama3_8bInstructParamsBodyPromptResponseFormat) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunHfMetaLlamaExecuteMetaLlama3_8bInstructParamsBodyPromptResponseFormatType string

const (
	AccountAIRunHfMetaLlamaExecuteMetaLlama3_8bInstructParamsBodyPromptResponseFormatTypeJsonObject AccountAIRunHfMetaLlamaExecuteMetaLlama3_8bInstructParamsBodyPromptResponseFormatType = "json_object"
	AccountAIRunHfMetaLlamaExecuteMetaLlama3_8bInstructParamsBodyPromptResponseFormatTypeJsonSchema AccountAIRunHfMetaLlamaExecuteMetaLlama3_8bInstructParamsBodyPromptResponseFormatType = "json_schema"
)

func (r AccountAIRunHfMetaLlamaExecuteMetaLlama3_8bInstructParamsBodyPromptResponseFormatType) IsKnown() bool {
	switch r {
	case AccountAIRunHfMetaLlamaExecuteMetaLlama3_8bInstructParamsBodyPromptResponseFormatTypeJsonObject, AccountAIRunHfMetaLlamaExecuteMetaLlama3_8bInstructParamsBodyPromptResponseFormatTypeJsonSchema:
		return true
	}
	return false
}

type AccountAIRunHfMetaLlamaExecuteMetaLlama3_8bInstructParamsBodyMessages struct {
	// An array of message objects representing the conversation history.
	Messages param.Field[[]AccountAIRunHfMetaLlamaExecuteMetaLlama3_8bInstructParamsBodyMessagesMessage] `json:"messages,required"`
	// Decreases the likelihood of the model repeating the same lines verbatim.
	FrequencyPenalty param.Field[float64]                                                                         `json:"frequency_penalty"`
	Functions        param.Field[[]AccountAIRunHfMetaLlamaExecuteMetaLlama3_8bInstructParamsBodyMessagesFunction] `json:"functions"`
	// The maximum number of tokens to generate in the response.
	MaxTokens param.Field[int64] `json:"max_tokens"`
	// Increases the likelihood of the model introducing new topics.
	PresencePenalty param.Field[float64] `json:"presence_penalty"`
	// If true, a chat template is not applied and you must adhere to the specific
	// model's expected formatting.
	Raw param.Field[bool] `json:"raw"`
	// Penalty for repeated tokens; higher values discourage repetition.
	RepetitionPenalty param.Field[float64]                                                                             `json:"repetition_penalty"`
	ResponseFormat    param.Field[AccountAIRunHfMetaLlamaExecuteMetaLlama3_8bInstructParamsBodyMessagesResponseFormat] `json:"response_format"`
	// Random seed for reproducibility of the generation.
	Seed param.Field[int64] `json:"seed"`
	// If true, the response will be streamed back incrementally using SSE, Server Sent
	// Events.
	Stream param.Field[bool] `json:"stream"`
	// Controls the randomness of the output; higher values produce more random
	// results.
	Temperature param.Field[float64] `json:"temperature"`
	// A list of tools available for the assistant to use.
	Tools param.Field[[]AccountAIRunHfMetaLlamaExecuteMetaLlama3_8bInstructParamsBodyMessagesToolUnion] `json:"tools"`
	// Limits the AI to choose from the top 'k' most probable words. Lower values make
	// responses more focused; higher values introduce more variety and potential
	// surprises.
	TopK param.Field[int64] `json:"top_k"`
	// Adjusts the creativity of the AI's responses by controlling how many possible
	// words it considers. Lower values make outputs more predictable; higher values
	// allow for more varied and creative responses.
	TopP param.Field[float64] `json:"top_p"`
}

func (r AccountAIRunHfMetaLlamaExecuteMetaLlama3_8bInstructParamsBodyMessages) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunHfMetaLlamaExecuteMetaLlama3_8bInstructParamsBodyMessages) implementsAccountAIRunHfMetaLlamaExecuteMetaLlama3_8bInstructParamsBodyUnion() {
}

type AccountAIRunHfMetaLlamaExecuteMetaLlama3_8bInstructParamsBodyMessagesMessage struct {
	// The content of the message as a string.
	Content param.Field[string] `json:"content,required"`
	// The role of the message sender (e.g., 'user', 'assistant', 'system', 'tool').
	Role param.Field[string] `json:"role,required"`
}

func (r AccountAIRunHfMetaLlamaExecuteMetaLlama3_8bInstructParamsBodyMessagesMessage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunHfMetaLlamaExecuteMetaLlama3_8bInstructParamsBodyMessagesFunction struct {
	Code param.Field[string] `json:"code,required"`
	Name param.Field[string] `json:"name,required"`
}

func (r AccountAIRunHfMetaLlamaExecuteMetaLlama3_8bInstructParamsBodyMessagesFunction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunHfMetaLlamaExecuteMetaLlama3_8bInstructParamsBodyMessagesResponseFormat struct {
	JsonSchema param.Field[interface{}]                                                                             `json:"json_schema"`
	Type       param.Field[AccountAIRunHfMetaLlamaExecuteMetaLlama3_8bInstructParamsBodyMessagesResponseFormatType] `json:"type"`
}

func (r AccountAIRunHfMetaLlamaExecuteMetaLlama3_8bInstructParamsBodyMessagesResponseFormat) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunHfMetaLlamaExecuteMetaLlama3_8bInstructParamsBodyMessagesResponseFormatType string

const (
	AccountAIRunHfMetaLlamaExecuteMetaLlama3_8bInstructParamsBodyMessagesResponseFormatTypeJsonObject AccountAIRunHfMetaLlamaExecuteMetaLlama3_8bInstructParamsBodyMessagesResponseFormatType = "json_object"
	AccountAIRunHfMetaLlamaExecuteMetaLlama3_8bInstructParamsBodyMessagesResponseFormatTypeJsonSchema AccountAIRunHfMetaLlamaExecuteMetaLlama3_8bInstructParamsBodyMessagesResponseFormatType = "json_schema"
)

func (r AccountAIRunHfMetaLlamaExecuteMetaLlama3_8bInstructParamsBodyMessagesResponseFormatType) IsKnown() bool {
	switch r {
	case AccountAIRunHfMetaLlamaExecuteMetaLlama3_8bInstructParamsBodyMessagesResponseFormatTypeJsonObject, AccountAIRunHfMetaLlamaExecuteMetaLlama3_8bInstructParamsBodyMessagesResponseFormatTypeJsonSchema:
		return true
	}
	return false
}

type AccountAIRunHfMetaLlamaExecuteMetaLlama3_8bInstructParamsBodyMessagesTool struct {
	// A brief description of what the tool does.
	Description param.Field[string]      `json:"description"`
	Function    param.Field[interface{}] `json:"function"`
	// The name of the tool. More descriptive the better.
	Name       param.Field[string]      `json:"name"`
	Parameters param.Field[interface{}] `json:"parameters"`
	// Specifies the type of tool (e.g., 'function').
	Type param.Field[string] `json:"type"`
}

func (r AccountAIRunHfMetaLlamaExecuteMetaLlama3_8bInstructParamsBodyMessagesTool) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunHfMetaLlamaExecuteMetaLlama3_8bInstructParamsBodyMessagesTool) implementsAccountAIRunHfMetaLlamaExecuteMetaLlama3_8bInstructParamsBodyMessagesToolUnion() {
}

// Satisfied by
// [AccountAIRunHfMetaLlamaExecuteMetaLlama3_8bInstructParamsBodyMessagesToolsObject],
// [AccountAIRunHfMetaLlamaExecuteMetaLlama3_8bInstructParamsBodyMessagesToolsObject],
// [AccountAIRunHfMetaLlamaExecuteMetaLlama3_8bInstructParamsBodyMessagesTool].
type AccountAIRunHfMetaLlamaExecuteMetaLlama3_8bInstructParamsBodyMessagesToolUnion interface {
	implementsAccountAIRunHfMetaLlamaExecuteMetaLlama3_8bInstructParamsBodyMessagesToolUnion()
}

type AccountAIRunHfMetaLlamaExecuteMetaLlama3_8bInstructParamsBodyMessagesToolsObject struct {
	// A brief description of what the tool does.
	Description param.Field[string] `json:"description,required"`
	// The name of the tool. More descriptive the better.
	Name param.Field[string] `json:"name,required"`
	// Schema defining the parameters accepted by the tool.
	Parameters param.Field[AccountAIRunHfMetaLlamaExecuteMetaLlama3_8bInstructParamsBodyMessagesToolsObjectParameters] `json:"parameters,required"`
}

func (r AccountAIRunHfMetaLlamaExecuteMetaLlama3_8bInstructParamsBodyMessagesToolsObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunHfMetaLlamaExecuteMetaLlama3_8bInstructParamsBodyMessagesToolsObject) implementsAccountAIRunHfMetaLlamaExecuteMetaLlama3_8bInstructParamsBodyMessagesToolUnion() {
}

// Schema defining the parameters accepted by the tool.
type AccountAIRunHfMetaLlamaExecuteMetaLlama3_8bInstructParamsBodyMessagesToolsObjectParameters struct {
	// Definitions of each parameter.
	Properties param.Field[map[string]AccountAIRunHfMetaLlamaExecuteMetaLlama3_8bInstructParamsBodyMessagesToolsObjectParametersProperties] `json:"properties,required"`
	// The type of the parameters object (usually 'object').
	Type param.Field[string] `json:"type,required"`
	// List of required parameter names.
	Required param.Field[[]string] `json:"required"`
}

func (r AccountAIRunHfMetaLlamaExecuteMetaLlama3_8bInstructParamsBodyMessagesToolsObjectParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunHfMetaLlamaExecuteMetaLlama3_8bInstructParamsBodyMessagesToolsObjectParametersProperties struct {
	// A description of the expected parameter.
	Description param.Field[string] `json:"description,required"`
	// The data type of the parameter.
	Type param.Field[string] `json:"type,required"`
}

func (r AccountAIRunHfMetaLlamaExecuteMetaLlama3_8bInstructParamsBodyMessagesToolsObjectParametersProperties) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
