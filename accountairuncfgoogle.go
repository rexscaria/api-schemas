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

// AccountAIRunCfGoogleService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountAIRunCfGoogleService] method instead.
type AccountAIRunCfGoogleService struct {
	Options []option.RequestOption
}

// NewAccountAIRunCfGoogleService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountAIRunCfGoogleService(opts ...option.RequestOption) (r *AccountAIRunCfGoogleService) {
	r = &AccountAIRunCfGoogleService{}
	r.Options = opts
	return
}

// Execute @cf/google/gemma-2b-it-lora model.
func (r *AccountAIRunCfGoogleService) ExecuteGemma2bItLora(ctx context.Context, accountID string, params AccountAIRunCfGoogleExecuteGemma2bItLoraParams, opts ...option.RequestOption) (res *AccountAIRunCfGoogleExecuteGemma2bItLoraResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai/run/@cf/google/gemma-2b-it-lora", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Execute @cf/google/gemma-7b-it-lora model.
func (r *AccountAIRunCfGoogleService) ExecuteGemma7bItLora(ctx context.Context, accountID string, params AccountAIRunCfGoogleExecuteGemma7bItLoraParams, opts ...option.RequestOption) (res *AccountAIRunCfGoogleExecuteGemma7bItLoraResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai/run/@cf/google/gemma-7b-it-lora", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type AccountAIRunCfGoogleExecuteGemma2bItLoraResponse = interface{}

type AccountAIRunCfGoogleExecuteGemma7bItLoraResponse = interface{}

type AccountAIRunCfGoogleExecuteGemma2bItLoraParams struct {
	QueueRequest param.Field[string]                                     `query:"queueRequest"`
	Body         AccountAIRunCfGoogleExecuteGemma2bItLoraParamsBodyUnion `json:"body"`
}

func (r AccountAIRunCfGoogleExecuteGemma2bItLoraParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

// URLQuery serializes [AccountAIRunCfGoogleExecuteGemma2bItLoraParams]'s query
// parameters as `url.Values`.
func (r AccountAIRunCfGoogleExecuteGemma2bItLoraParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountAIRunCfGoogleExecuteGemma2bItLoraParamsBody struct {
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

func (r AccountAIRunCfGoogleExecuteGemma2bItLoraParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfGoogleExecuteGemma2bItLoraParamsBody) implementsAccountAIRunCfGoogleExecuteGemma2bItLoraParamsBodyUnion() {
}

// Satisfied by [AccountAIRunCfGoogleExecuteGemma2bItLoraParamsBodyPrompt],
// [AccountAIRunCfGoogleExecuteGemma2bItLoraParamsBodyMessages],
// [AccountAIRunCfGoogleExecuteGemma2bItLoraParamsBody].
type AccountAIRunCfGoogleExecuteGemma2bItLoraParamsBodyUnion interface {
	implementsAccountAIRunCfGoogleExecuteGemma2bItLoraParamsBodyUnion()
}

type AccountAIRunCfGoogleExecuteGemma2bItLoraParamsBodyPrompt struct {
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
	RepetitionPenalty param.Field[float64]                                                                `json:"repetition_penalty"`
	ResponseFormat    param.Field[AccountAIRunCfGoogleExecuteGemma2bItLoraParamsBodyPromptResponseFormat] `json:"response_format"`
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

func (r AccountAIRunCfGoogleExecuteGemma2bItLoraParamsBodyPrompt) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfGoogleExecuteGemma2bItLoraParamsBodyPrompt) implementsAccountAIRunCfGoogleExecuteGemma2bItLoraParamsBodyUnion() {
}

type AccountAIRunCfGoogleExecuteGemma2bItLoraParamsBodyPromptResponseFormat struct {
	JsonSchema param.Field[interface{}]                                                                `json:"json_schema"`
	Type       param.Field[AccountAIRunCfGoogleExecuteGemma2bItLoraParamsBodyPromptResponseFormatType] `json:"type"`
}

func (r AccountAIRunCfGoogleExecuteGemma2bItLoraParamsBodyPromptResponseFormat) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfGoogleExecuteGemma2bItLoraParamsBodyPromptResponseFormatType string

const (
	AccountAIRunCfGoogleExecuteGemma2bItLoraParamsBodyPromptResponseFormatTypeJsonObject AccountAIRunCfGoogleExecuteGemma2bItLoraParamsBodyPromptResponseFormatType = "json_object"
	AccountAIRunCfGoogleExecuteGemma2bItLoraParamsBodyPromptResponseFormatTypeJsonSchema AccountAIRunCfGoogleExecuteGemma2bItLoraParamsBodyPromptResponseFormatType = "json_schema"
)

func (r AccountAIRunCfGoogleExecuteGemma2bItLoraParamsBodyPromptResponseFormatType) IsKnown() bool {
	switch r {
	case AccountAIRunCfGoogleExecuteGemma2bItLoraParamsBodyPromptResponseFormatTypeJsonObject, AccountAIRunCfGoogleExecuteGemma2bItLoraParamsBodyPromptResponseFormatTypeJsonSchema:
		return true
	}
	return false
}

type AccountAIRunCfGoogleExecuteGemma2bItLoraParamsBodyMessages struct {
	// An array of message objects representing the conversation history.
	Messages param.Field[[]AccountAIRunCfGoogleExecuteGemma2bItLoraParamsBodyMessagesMessage] `json:"messages,required"`
	// Decreases the likelihood of the model repeating the same lines verbatim.
	FrequencyPenalty param.Field[float64]                                                              `json:"frequency_penalty"`
	Functions        param.Field[[]AccountAIRunCfGoogleExecuteGemma2bItLoraParamsBodyMessagesFunction] `json:"functions"`
	// The maximum number of tokens to generate in the response.
	MaxTokens param.Field[int64] `json:"max_tokens"`
	// Increases the likelihood of the model introducing new topics.
	PresencePenalty param.Field[float64] `json:"presence_penalty"`
	// If true, a chat template is not applied and you must adhere to the specific
	// model's expected formatting.
	Raw param.Field[bool] `json:"raw"`
	// Penalty for repeated tokens; higher values discourage repetition.
	RepetitionPenalty param.Field[float64]                                                                  `json:"repetition_penalty"`
	ResponseFormat    param.Field[AccountAIRunCfGoogleExecuteGemma2bItLoraParamsBodyMessagesResponseFormat] `json:"response_format"`
	// Random seed for reproducibility of the generation.
	Seed param.Field[int64] `json:"seed"`
	// If true, the response will be streamed back incrementally using SSE, Server Sent
	// Events.
	Stream param.Field[bool] `json:"stream"`
	// Controls the randomness of the output; higher values produce more random
	// results.
	Temperature param.Field[float64] `json:"temperature"`
	// A list of tools available for the assistant to use.
	Tools param.Field[[]AccountAIRunCfGoogleExecuteGemma2bItLoraParamsBodyMessagesToolUnion] `json:"tools"`
	// Limits the AI to choose from the top 'k' most probable words. Lower values make
	// responses more focused; higher values introduce more variety and potential
	// surprises.
	TopK param.Field[int64] `json:"top_k"`
	// Adjusts the creativity of the AI's responses by controlling how many possible
	// words it considers. Lower values make outputs more predictable; higher values
	// allow for more varied and creative responses.
	TopP param.Field[float64] `json:"top_p"`
}

func (r AccountAIRunCfGoogleExecuteGemma2bItLoraParamsBodyMessages) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfGoogleExecuteGemma2bItLoraParamsBodyMessages) implementsAccountAIRunCfGoogleExecuteGemma2bItLoraParamsBodyUnion() {
}

type AccountAIRunCfGoogleExecuteGemma2bItLoraParamsBodyMessagesMessage struct {
	// The content of the message as a string.
	Content param.Field[string] `json:"content,required"`
	// The role of the message sender (e.g., 'user', 'assistant', 'system', 'tool').
	Role param.Field[string] `json:"role,required"`
}

func (r AccountAIRunCfGoogleExecuteGemma2bItLoraParamsBodyMessagesMessage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfGoogleExecuteGemma2bItLoraParamsBodyMessagesFunction struct {
	Code param.Field[string] `json:"code,required"`
	Name param.Field[string] `json:"name,required"`
}

func (r AccountAIRunCfGoogleExecuteGemma2bItLoraParamsBodyMessagesFunction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfGoogleExecuteGemma2bItLoraParamsBodyMessagesResponseFormat struct {
	JsonSchema param.Field[interface{}]                                                                  `json:"json_schema"`
	Type       param.Field[AccountAIRunCfGoogleExecuteGemma2bItLoraParamsBodyMessagesResponseFormatType] `json:"type"`
}

func (r AccountAIRunCfGoogleExecuteGemma2bItLoraParamsBodyMessagesResponseFormat) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfGoogleExecuteGemma2bItLoraParamsBodyMessagesResponseFormatType string

const (
	AccountAIRunCfGoogleExecuteGemma2bItLoraParamsBodyMessagesResponseFormatTypeJsonObject AccountAIRunCfGoogleExecuteGemma2bItLoraParamsBodyMessagesResponseFormatType = "json_object"
	AccountAIRunCfGoogleExecuteGemma2bItLoraParamsBodyMessagesResponseFormatTypeJsonSchema AccountAIRunCfGoogleExecuteGemma2bItLoraParamsBodyMessagesResponseFormatType = "json_schema"
)

func (r AccountAIRunCfGoogleExecuteGemma2bItLoraParamsBodyMessagesResponseFormatType) IsKnown() bool {
	switch r {
	case AccountAIRunCfGoogleExecuteGemma2bItLoraParamsBodyMessagesResponseFormatTypeJsonObject, AccountAIRunCfGoogleExecuteGemma2bItLoraParamsBodyMessagesResponseFormatTypeJsonSchema:
		return true
	}
	return false
}

type AccountAIRunCfGoogleExecuteGemma2bItLoraParamsBodyMessagesTool struct {
	// A brief description of what the tool does.
	Description param.Field[string]      `json:"description"`
	Function    param.Field[interface{}] `json:"function"`
	// The name of the tool. More descriptive the better.
	Name       param.Field[string]      `json:"name"`
	Parameters param.Field[interface{}] `json:"parameters"`
	// Specifies the type of tool (e.g., 'function').
	Type param.Field[string] `json:"type"`
}

func (r AccountAIRunCfGoogleExecuteGemma2bItLoraParamsBodyMessagesTool) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfGoogleExecuteGemma2bItLoraParamsBodyMessagesTool) implementsAccountAIRunCfGoogleExecuteGemma2bItLoraParamsBodyMessagesToolUnion() {
}

// Satisfied by
// [AccountAIRunCfGoogleExecuteGemma2bItLoraParamsBodyMessagesToolsObject],
// [AccountAIRunCfGoogleExecuteGemma2bItLoraParamsBodyMessagesToolsObject],
// [AccountAIRunCfGoogleExecuteGemma2bItLoraParamsBodyMessagesTool].
type AccountAIRunCfGoogleExecuteGemma2bItLoraParamsBodyMessagesToolUnion interface {
	implementsAccountAIRunCfGoogleExecuteGemma2bItLoraParamsBodyMessagesToolUnion()
}

type AccountAIRunCfGoogleExecuteGemma2bItLoraParamsBodyMessagesToolsObject struct {
	// A brief description of what the tool does.
	Description param.Field[string] `json:"description,required"`
	// The name of the tool. More descriptive the better.
	Name param.Field[string] `json:"name,required"`
	// Schema defining the parameters accepted by the tool.
	Parameters param.Field[AccountAIRunCfGoogleExecuteGemma2bItLoraParamsBodyMessagesToolsObjectParameters] `json:"parameters,required"`
}

func (r AccountAIRunCfGoogleExecuteGemma2bItLoraParamsBodyMessagesToolsObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfGoogleExecuteGemma2bItLoraParamsBodyMessagesToolsObject) implementsAccountAIRunCfGoogleExecuteGemma2bItLoraParamsBodyMessagesToolUnion() {
}

// Schema defining the parameters accepted by the tool.
type AccountAIRunCfGoogleExecuteGemma2bItLoraParamsBodyMessagesToolsObjectParameters struct {
	// Definitions of each parameter.
	Properties param.Field[map[string]AccountAIRunCfGoogleExecuteGemma2bItLoraParamsBodyMessagesToolsObjectParametersProperties] `json:"properties,required"`
	// The type of the parameters object (usually 'object').
	Type param.Field[string] `json:"type,required"`
	// List of required parameter names.
	Required param.Field[[]string] `json:"required"`
}

func (r AccountAIRunCfGoogleExecuteGemma2bItLoraParamsBodyMessagesToolsObjectParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfGoogleExecuteGemma2bItLoraParamsBodyMessagesToolsObjectParametersProperties struct {
	// A description of the expected parameter.
	Description param.Field[string] `json:"description,required"`
	// The data type of the parameter.
	Type param.Field[string] `json:"type,required"`
}

func (r AccountAIRunCfGoogleExecuteGemma2bItLoraParamsBodyMessagesToolsObjectParametersProperties) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfGoogleExecuteGemma7bItLoraParams struct {
	QueueRequest param.Field[string]                                     `query:"queueRequest"`
	Body         AccountAIRunCfGoogleExecuteGemma7bItLoraParamsBodyUnion `json:"body"`
}

func (r AccountAIRunCfGoogleExecuteGemma7bItLoraParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

// URLQuery serializes [AccountAIRunCfGoogleExecuteGemma7bItLoraParams]'s query
// parameters as `url.Values`.
func (r AccountAIRunCfGoogleExecuteGemma7bItLoraParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountAIRunCfGoogleExecuteGemma7bItLoraParamsBody struct {
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

func (r AccountAIRunCfGoogleExecuteGemma7bItLoraParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfGoogleExecuteGemma7bItLoraParamsBody) implementsAccountAIRunCfGoogleExecuteGemma7bItLoraParamsBodyUnion() {
}

// Satisfied by [AccountAIRunCfGoogleExecuteGemma7bItLoraParamsBodyPrompt],
// [AccountAIRunCfGoogleExecuteGemma7bItLoraParamsBodyMessages],
// [AccountAIRunCfGoogleExecuteGemma7bItLoraParamsBody].
type AccountAIRunCfGoogleExecuteGemma7bItLoraParamsBodyUnion interface {
	implementsAccountAIRunCfGoogleExecuteGemma7bItLoraParamsBodyUnion()
}

type AccountAIRunCfGoogleExecuteGemma7bItLoraParamsBodyPrompt struct {
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
	RepetitionPenalty param.Field[float64]                                                                `json:"repetition_penalty"`
	ResponseFormat    param.Field[AccountAIRunCfGoogleExecuteGemma7bItLoraParamsBodyPromptResponseFormat] `json:"response_format"`
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

func (r AccountAIRunCfGoogleExecuteGemma7bItLoraParamsBodyPrompt) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfGoogleExecuteGemma7bItLoraParamsBodyPrompt) implementsAccountAIRunCfGoogleExecuteGemma7bItLoraParamsBodyUnion() {
}

type AccountAIRunCfGoogleExecuteGemma7bItLoraParamsBodyPromptResponseFormat struct {
	JsonSchema param.Field[interface{}]                                                                `json:"json_schema"`
	Type       param.Field[AccountAIRunCfGoogleExecuteGemma7bItLoraParamsBodyPromptResponseFormatType] `json:"type"`
}

func (r AccountAIRunCfGoogleExecuteGemma7bItLoraParamsBodyPromptResponseFormat) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfGoogleExecuteGemma7bItLoraParamsBodyPromptResponseFormatType string

const (
	AccountAIRunCfGoogleExecuteGemma7bItLoraParamsBodyPromptResponseFormatTypeJsonObject AccountAIRunCfGoogleExecuteGemma7bItLoraParamsBodyPromptResponseFormatType = "json_object"
	AccountAIRunCfGoogleExecuteGemma7bItLoraParamsBodyPromptResponseFormatTypeJsonSchema AccountAIRunCfGoogleExecuteGemma7bItLoraParamsBodyPromptResponseFormatType = "json_schema"
)

func (r AccountAIRunCfGoogleExecuteGemma7bItLoraParamsBodyPromptResponseFormatType) IsKnown() bool {
	switch r {
	case AccountAIRunCfGoogleExecuteGemma7bItLoraParamsBodyPromptResponseFormatTypeJsonObject, AccountAIRunCfGoogleExecuteGemma7bItLoraParamsBodyPromptResponseFormatTypeJsonSchema:
		return true
	}
	return false
}

type AccountAIRunCfGoogleExecuteGemma7bItLoraParamsBodyMessages struct {
	// An array of message objects representing the conversation history.
	Messages param.Field[[]AccountAIRunCfGoogleExecuteGemma7bItLoraParamsBodyMessagesMessage] `json:"messages,required"`
	// Decreases the likelihood of the model repeating the same lines verbatim.
	FrequencyPenalty param.Field[float64]                                                              `json:"frequency_penalty"`
	Functions        param.Field[[]AccountAIRunCfGoogleExecuteGemma7bItLoraParamsBodyMessagesFunction] `json:"functions"`
	// The maximum number of tokens to generate in the response.
	MaxTokens param.Field[int64] `json:"max_tokens"`
	// Increases the likelihood of the model introducing new topics.
	PresencePenalty param.Field[float64] `json:"presence_penalty"`
	// If true, a chat template is not applied and you must adhere to the specific
	// model's expected formatting.
	Raw param.Field[bool] `json:"raw"`
	// Penalty for repeated tokens; higher values discourage repetition.
	RepetitionPenalty param.Field[float64]                                                                  `json:"repetition_penalty"`
	ResponseFormat    param.Field[AccountAIRunCfGoogleExecuteGemma7bItLoraParamsBodyMessagesResponseFormat] `json:"response_format"`
	// Random seed for reproducibility of the generation.
	Seed param.Field[int64] `json:"seed"`
	// If true, the response will be streamed back incrementally using SSE, Server Sent
	// Events.
	Stream param.Field[bool] `json:"stream"`
	// Controls the randomness of the output; higher values produce more random
	// results.
	Temperature param.Field[float64] `json:"temperature"`
	// A list of tools available for the assistant to use.
	Tools param.Field[[]AccountAIRunCfGoogleExecuteGemma7bItLoraParamsBodyMessagesToolUnion] `json:"tools"`
	// Limits the AI to choose from the top 'k' most probable words. Lower values make
	// responses more focused; higher values introduce more variety and potential
	// surprises.
	TopK param.Field[int64] `json:"top_k"`
	// Adjusts the creativity of the AI's responses by controlling how many possible
	// words it considers. Lower values make outputs more predictable; higher values
	// allow for more varied and creative responses.
	TopP param.Field[float64] `json:"top_p"`
}

func (r AccountAIRunCfGoogleExecuteGemma7bItLoraParamsBodyMessages) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfGoogleExecuteGemma7bItLoraParamsBodyMessages) implementsAccountAIRunCfGoogleExecuteGemma7bItLoraParamsBodyUnion() {
}

type AccountAIRunCfGoogleExecuteGemma7bItLoraParamsBodyMessagesMessage struct {
	// The content of the message as a string.
	Content param.Field[string] `json:"content,required"`
	// The role of the message sender (e.g., 'user', 'assistant', 'system', 'tool').
	Role param.Field[string] `json:"role,required"`
}

func (r AccountAIRunCfGoogleExecuteGemma7bItLoraParamsBodyMessagesMessage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfGoogleExecuteGemma7bItLoraParamsBodyMessagesFunction struct {
	Code param.Field[string] `json:"code,required"`
	Name param.Field[string] `json:"name,required"`
}

func (r AccountAIRunCfGoogleExecuteGemma7bItLoraParamsBodyMessagesFunction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfGoogleExecuteGemma7bItLoraParamsBodyMessagesResponseFormat struct {
	JsonSchema param.Field[interface{}]                                                                  `json:"json_schema"`
	Type       param.Field[AccountAIRunCfGoogleExecuteGemma7bItLoraParamsBodyMessagesResponseFormatType] `json:"type"`
}

func (r AccountAIRunCfGoogleExecuteGemma7bItLoraParamsBodyMessagesResponseFormat) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfGoogleExecuteGemma7bItLoraParamsBodyMessagesResponseFormatType string

const (
	AccountAIRunCfGoogleExecuteGemma7bItLoraParamsBodyMessagesResponseFormatTypeJsonObject AccountAIRunCfGoogleExecuteGemma7bItLoraParamsBodyMessagesResponseFormatType = "json_object"
	AccountAIRunCfGoogleExecuteGemma7bItLoraParamsBodyMessagesResponseFormatTypeJsonSchema AccountAIRunCfGoogleExecuteGemma7bItLoraParamsBodyMessagesResponseFormatType = "json_schema"
)

func (r AccountAIRunCfGoogleExecuteGemma7bItLoraParamsBodyMessagesResponseFormatType) IsKnown() bool {
	switch r {
	case AccountAIRunCfGoogleExecuteGemma7bItLoraParamsBodyMessagesResponseFormatTypeJsonObject, AccountAIRunCfGoogleExecuteGemma7bItLoraParamsBodyMessagesResponseFormatTypeJsonSchema:
		return true
	}
	return false
}

type AccountAIRunCfGoogleExecuteGemma7bItLoraParamsBodyMessagesTool struct {
	// A brief description of what the tool does.
	Description param.Field[string]      `json:"description"`
	Function    param.Field[interface{}] `json:"function"`
	// The name of the tool. More descriptive the better.
	Name       param.Field[string]      `json:"name"`
	Parameters param.Field[interface{}] `json:"parameters"`
	// Specifies the type of tool (e.g., 'function').
	Type param.Field[string] `json:"type"`
}

func (r AccountAIRunCfGoogleExecuteGemma7bItLoraParamsBodyMessagesTool) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfGoogleExecuteGemma7bItLoraParamsBodyMessagesTool) implementsAccountAIRunCfGoogleExecuteGemma7bItLoraParamsBodyMessagesToolUnion() {
}

// Satisfied by
// [AccountAIRunCfGoogleExecuteGemma7bItLoraParamsBodyMessagesToolsObject],
// [AccountAIRunCfGoogleExecuteGemma7bItLoraParamsBodyMessagesToolsObject],
// [AccountAIRunCfGoogleExecuteGemma7bItLoraParamsBodyMessagesTool].
type AccountAIRunCfGoogleExecuteGemma7bItLoraParamsBodyMessagesToolUnion interface {
	implementsAccountAIRunCfGoogleExecuteGemma7bItLoraParamsBodyMessagesToolUnion()
}

type AccountAIRunCfGoogleExecuteGemma7bItLoraParamsBodyMessagesToolsObject struct {
	// A brief description of what the tool does.
	Description param.Field[string] `json:"description,required"`
	// The name of the tool. More descriptive the better.
	Name param.Field[string] `json:"name,required"`
	// Schema defining the parameters accepted by the tool.
	Parameters param.Field[AccountAIRunCfGoogleExecuteGemma7bItLoraParamsBodyMessagesToolsObjectParameters] `json:"parameters,required"`
}

func (r AccountAIRunCfGoogleExecuteGemma7bItLoraParamsBodyMessagesToolsObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfGoogleExecuteGemma7bItLoraParamsBodyMessagesToolsObject) implementsAccountAIRunCfGoogleExecuteGemma7bItLoraParamsBodyMessagesToolUnion() {
}

// Schema defining the parameters accepted by the tool.
type AccountAIRunCfGoogleExecuteGemma7bItLoraParamsBodyMessagesToolsObjectParameters struct {
	// Definitions of each parameter.
	Properties param.Field[map[string]AccountAIRunCfGoogleExecuteGemma7bItLoraParamsBodyMessagesToolsObjectParametersProperties] `json:"properties,required"`
	// The type of the parameters object (usually 'object').
	Type param.Field[string] `json:"type,required"`
	// List of required parameter names.
	Required param.Field[[]string] `json:"required"`
}

func (r AccountAIRunCfGoogleExecuteGemma7bItLoraParamsBodyMessagesToolsObjectParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfGoogleExecuteGemma7bItLoraParamsBodyMessagesToolsObjectParametersProperties struct {
	// A description of the expected parameter.
	Description param.Field[string] `json:"description,required"`
	// The data type of the parameter.
	Type param.Field[string] `json:"type,required"`
}

func (r AccountAIRunCfGoogleExecuteGemma7bItLoraParamsBodyMessagesToolsObjectParametersProperties) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
