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

// AccountAIRunCfMetaLlamaService contains methods and other services that help
// with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountAIRunCfMetaLlamaService] method instead.
type AccountAIRunCfMetaLlamaService struct {
	Options []option.RequestOption
}

// NewAccountAIRunCfMetaLlamaService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountAIRunCfMetaLlamaService(opts ...option.RequestOption) (r *AccountAIRunCfMetaLlamaService) {
	r = &AccountAIRunCfMetaLlamaService{}
	r.Options = opts
	return
}

// Execute @cf/meta-llama/llama-2-7b-chat-hf-lora model.
func (r *AccountAIRunCfMetaLlamaService) ExecuteLlama2_7bChatHfLora(ctx context.Context, accountID string, params AccountAIRunCfMetaLlamaExecuteLlama2_7bChatHfLoraParams, opts ...option.RequestOption) (res *AccountAIRunCfMetaLlamaExecuteLlama2_7bChatHfLoraResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai/run/@cf/meta-llama/llama-2-7b-chat-hf-lora", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type AccountAIRunCfMetaLlamaExecuteLlama2_7bChatHfLoraResponse = interface{}

type AccountAIRunCfMetaLlamaExecuteLlama2_7bChatHfLoraParams struct {
	QueueRequest param.Field[string]                                              `query:"queueRequest"`
	Body         AccountAIRunCfMetaLlamaExecuteLlama2_7bChatHfLoraParamsBodyUnion `json:"body"`
}

func (r AccountAIRunCfMetaLlamaExecuteLlama2_7bChatHfLoraParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

// URLQuery serializes [AccountAIRunCfMetaLlamaExecuteLlama2_7bChatHfLoraParams]'s
// query parameters as `url.Values`.
func (r AccountAIRunCfMetaLlamaExecuteLlama2_7bChatHfLoraParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountAIRunCfMetaLlamaExecuteLlama2_7bChatHfLoraParamsBody struct {
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

func (r AccountAIRunCfMetaLlamaExecuteLlama2_7bChatHfLoraParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfMetaLlamaExecuteLlama2_7bChatHfLoraParamsBody) implementsAccountAIRunCfMetaLlamaExecuteLlama2_7bChatHfLoraParamsBodyUnion() {
}

// Satisfied by
// [AccountAIRunCfMetaLlamaExecuteLlama2_7bChatHfLoraParamsBodyPrompt],
// [AccountAIRunCfMetaLlamaExecuteLlama2_7bChatHfLoraParamsBodyMessages],
// [AccountAIRunCfMetaLlamaExecuteLlama2_7bChatHfLoraParamsBody].
type AccountAIRunCfMetaLlamaExecuteLlama2_7bChatHfLoraParamsBodyUnion interface {
	implementsAccountAIRunCfMetaLlamaExecuteLlama2_7bChatHfLoraParamsBodyUnion()
}

type AccountAIRunCfMetaLlamaExecuteLlama2_7bChatHfLoraParamsBodyPrompt struct {
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
	ResponseFormat    param.Field[AccountAIRunCfMetaLlamaExecuteLlama2_7bChatHfLoraParamsBodyPromptResponseFormat] `json:"response_format"`
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

func (r AccountAIRunCfMetaLlamaExecuteLlama2_7bChatHfLoraParamsBodyPrompt) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfMetaLlamaExecuteLlama2_7bChatHfLoraParamsBodyPrompt) implementsAccountAIRunCfMetaLlamaExecuteLlama2_7bChatHfLoraParamsBodyUnion() {
}

type AccountAIRunCfMetaLlamaExecuteLlama2_7bChatHfLoraParamsBodyPromptResponseFormat struct {
	JsonSchema param.Field[interface{}]                                                                         `json:"json_schema"`
	Type       param.Field[AccountAIRunCfMetaLlamaExecuteLlama2_7bChatHfLoraParamsBodyPromptResponseFormatType] `json:"type"`
}

func (r AccountAIRunCfMetaLlamaExecuteLlama2_7bChatHfLoraParamsBodyPromptResponseFormat) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfMetaLlamaExecuteLlama2_7bChatHfLoraParamsBodyPromptResponseFormatType string

const (
	AccountAIRunCfMetaLlamaExecuteLlama2_7bChatHfLoraParamsBodyPromptResponseFormatTypeJsonObject AccountAIRunCfMetaLlamaExecuteLlama2_7bChatHfLoraParamsBodyPromptResponseFormatType = "json_object"
	AccountAIRunCfMetaLlamaExecuteLlama2_7bChatHfLoraParamsBodyPromptResponseFormatTypeJsonSchema AccountAIRunCfMetaLlamaExecuteLlama2_7bChatHfLoraParamsBodyPromptResponseFormatType = "json_schema"
)

func (r AccountAIRunCfMetaLlamaExecuteLlama2_7bChatHfLoraParamsBodyPromptResponseFormatType) IsKnown() bool {
	switch r {
	case AccountAIRunCfMetaLlamaExecuteLlama2_7bChatHfLoraParamsBodyPromptResponseFormatTypeJsonObject, AccountAIRunCfMetaLlamaExecuteLlama2_7bChatHfLoraParamsBodyPromptResponseFormatTypeJsonSchema:
		return true
	}
	return false
}

type AccountAIRunCfMetaLlamaExecuteLlama2_7bChatHfLoraParamsBodyMessages struct {
	// An array of message objects representing the conversation history.
	Messages param.Field[[]AccountAIRunCfMetaLlamaExecuteLlama2_7bChatHfLoraParamsBodyMessagesMessage] `json:"messages,required"`
	// Decreases the likelihood of the model repeating the same lines verbatim.
	FrequencyPenalty param.Field[float64]                                                                       `json:"frequency_penalty"`
	Functions        param.Field[[]AccountAIRunCfMetaLlamaExecuteLlama2_7bChatHfLoraParamsBodyMessagesFunction] `json:"functions"`
	// The maximum number of tokens to generate in the response.
	MaxTokens param.Field[int64] `json:"max_tokens"`
	// Increases the likelihood of the model introducing new topics.
	PresencePenalty param.Field[float64] `json:"presence_penalty"`
	// If true, a chat template is not applied and you must adhere to the specific
	// model's expected formatting.
	Raw param.Field[bool] `json:"raw"`
	// Penalty for repeated tokens; higher values discourage repetition.
	RepetitionPenalty param.Field[float64]                                                                           `json:"repetition_penalty"`
	ResponseFormat    param.Field[AccountAIRunCfMetaLlamaExecuteLlama2_7bChatHfLoraParamsBodyMessagesResponseFormat] `json:"response_format"`
	// Random seed for reproducibility of the generation.
	Seed param.Field[int64] `json:"seed"`
	// If true, the response will be streamed back incrementally using SSE, Server Sent
	// Events.
	Stream param.Field[bool] `json:"stream"`
	// Controls the randomness of the output; higher values produce more random
	// results.
	Temperature param.Field[float64] `json:"temperature"`
	// A list of tools available for the assistant to use.
	Tools param.Field[[]AccountAIRunCfMetaLlamaExecuteLlama2_7bChatHfLoraParamsBodyMessagesToolUnion] `json:"tools"`
	// Limits the AI to choose from the top 'k' most probable words. Lower values make
	// responses more focused; higher values introduce more variety and potential
	// surprises.
	TopK param.Field[int64] `json:"top_k"`
	// Adjusts the creativity of the AI's responses by controlling how many possible
	// words it considers. Lower values make outputs more predictable; higher values
	// allow for more varied and creative responses.
	TopP param.Field[float64] `json:"top_p"`
}

func (r AccountAIRunCfMetaLlamaExecuteLlama2_7bChatHfLoraParamsBodyMessages) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfMetaLlamaExecuteLlama2_7bChatHfLoraParamsBodyMessages) implementsAccountAIRunCfMetaLlamaExecuteLlama2_7bChatHfLoraParamsBodyUnion() {
}

type AccountAIRunCfMetaLlamaExecuteLlama2_7bChatHfLoraParamsBodyMessagesMessage struct {
	// The content of the message as a string.
	Content param.Field[string] `json:"content,required"`
	// The role of the message sender (e.g., 'user', 'assistant', 'system', 'tool').
	Role param.Field[string] `json:"role,required"`
}

func (r AccountAIRunCfMetaLlamaExecuteLlama2_7bChatHfLoraParamsBodyMessagesMessage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfMetaLlamaExecuteLlama2_7bChatHfLoraParamsBodyMessagesFunction struct {
	Code param.Field[string] `json:"code,required"`
	Name param.Field[string] `json:"name,required"`
}

func (r AccountAIRunCfMetaLlamaExecuteLlama2_7bChatHfLoraParamsBodyMessagesFunction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfMetaLlamaExecuteLlama2_7bChatHfLoraParamsBodyMessagesResponseFormat struct {
	JsonSchema param.Field[interface{}]                                                                           `json:"json_schema"`
	Type       param.Field[AccountAIRunCfMetaLlamaExecuteLlama2_7bChatHfLoraParamsBodyMessagesResponseFormatType] `json:"type"`
}

func (r AccountAIRunCfMetaLlamaExecuteLlama2_7bChatHfLoraParamsBodyMessagesResponseFormat) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfMetaLlamaExecuteLlama2_7bChatHfLoraParamsBodyMessagesResponseFormatType string

const (
	AccountAIRunCfMetaLlamaExecuteLlama2_7bChatHfLoraParamsBodyMessagesResponseFormatTypeJsonObject AccountAIRunCfMetaLlamaExecuteLlama2_7bChatHfLoraParamsBodyMessagesResponseFormatType = "json_object"
	AccountAIRunCfMetaLlamaExecuteLlama2_7bChatHfLoraParamsBodyMessagesResponseFormatTypeJsonSchema AccountAIRunCfMetaLlamaExecuteLlama2_7bChatHfLoraParamsBodyMessagesResponseFormatType = "json_schema"
)

func (r AccountAIRunCfMetaLlamaExecuteLlama2_7bChatHfLoraParamsBodyMessagesResponseFormatType) IsKnown() bool {
	switch r {
	case AccountAIRunCfMetaLlamaExecuteLlama2_7bChatHfLoraParamsBodyMessagesResponseFormatTypeJsonObject, AccountAIRunCfMetaLlamaExecuteLlama2_7bChatHfLoraParamsBodyMessagesResponseFormatTypeJsonSchema:
		return true
	}
	return false
}

type AccountAIRunCfMetaLlamaExecuteLlama2_7bChatHfLoraParamsBodyMessagesTool struct {
	// A brief description of what the tool does.
	Description param.Field[string]      `json:"description"`
	Function    param.Field[interface{}] `json:"function"`
	// The name of the tool. More descriptive the better.
	Name       param.Field[string]      `json:"name"`
	Parameters param.Field[interface{}] `json:"parameters"`
	// Specifies the type of tool (e.g., 'function').
	Type param.Field[string] `json:"type"`
}

func (r AccountAIRunCfMetaLlamaExecuteLlama2_7bChatHfLoraParamsBodyMessagesTool) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfMetaLlamaExecuteLlama2_7bChatHfLoraParamsBodyMessagesTool) implementsAccountAIRunCfMetaLlamaExecuteLlama2_7bChatHfLoraParamsBodyMessagesToolUnion() {
}

// Satisfied by
// [AccountAIRunCfMetaLlamaExecuteLlama2_7bChatHfLoraParamsBodyMessagesToolsObject],
// [AccountAIRunCfMetaLlamaExecuteLlama2_7bChatHfLoraParamsBodyMessagesToolsObject],
// [AccountAIRunCfMetaLlamaExecuteLlama2_7bChatHfLoraParamsBodyMessagesTool].
type AccountAIRunCfMetaLlamaExecuteLlama2_7bChatHfLoraParamsBodyMessagesToolUnion interface {
	implementsAccountAIRunCfMetaLlamaExecuteLlama2_7bChatHfLoraParamsBodyMessagesToolUnion()
}

type AccountAIRunCfMetaLlamaExecuteLlama2_7bChatHfLoraParamsBodyMessagesToolsObject struct {
	// A brief description of what the tool does.
	Description param.Field[string] `json:"description,required"`
	// The name of the tool. More descriptive the better.
	Name param.Field[string] `json:"name,required"`
	// Schema defining the parameters accepted by the tool.
	Parameters param.Field[AccountAIRunCfMetaLlamaExecuteLlama2_7bChatHfLoraParamsBodyMessagesToolsObjectParameters] `json:"parameters,required"`
}

func (r AccountAIRunCfMetaLlamaExecuteLlama2_7bChatHfLoraParamsBodyMessagesToolsObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfMetaLlamaExecuteLlama2_7bChatHfLoraParamsBodyMessagesToolsObject) implementsAccountAIRunCfMetaLlamaExecuteLlama2_7bChatHfLoraParamsBodyMessagesToolUnion() {
}

// Schema defining the parameters accepted by the tool.
type AccountAIRunCfMetaLlamaExecuteLlama2_7bChatHfLoraParamsBodyMessagesToolsObjectParameters struct {
	// Definitions of each parameter.
	Properties param.Field[map[string]AccountAIRunCfMetaLlamaExecuteLlama2_7bChatHfLoraParamsBodyMessagesToolsObjectParametersProperties] `json:"properties,required"`
	// The type of the parameters object (usually 'object').
	Type param.Field[string] `json:"type,required"`
	// List of required parameter names.
	Required param.Field[[]string] `json:"required"`
}

func (r AccountAIRunCfMetaLlamaExecuteLlama2_7bChatHfLoraParamsBodyMessagesToolsObjectParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfMetaLlamaExecuteLlama2_7bChatHfLoraParamsBodyMessagesToolsObjectParametersProperties struct {
	// A description of the expected parameter.
	Description param.Field[string] `json:"description,required"`
	// The data type of the parameter.
	Type param.Field[string] `json:"type,required"`
}

func (r AccountAIRunCfMetaLlamaExecuteLlama2_7bChatHfLoraParamsBodyMessagesToolsObjectParametersProperties) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
