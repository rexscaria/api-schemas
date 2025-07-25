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

// AccountAIRunCfDeepseekAIService contains methods and other services that help
// with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountAIRunCfDeepseekAIService] method instead.
type AccountAIRunCfDeepseekAIService struct {
	Options []option.RequestOption
}

// NewAccountAIRunCfDeepseekAIService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountAIRunCfDeepseekAIService(opts ...option.RequestOption) (r *AccountAIRunCfDeepseekAIService) {
	r = &AccountAIRunCfDeepseekAIService{}
	r.Options = opts
	return
}

// Execute @cf/deepseek-ai/deepseek-math-7b-instruct model.
func (r *AccountAIRunCfDeepseekAIService) ExecuteDeepseekMath7bInstruct(ctx context.Context, accountID string, params AccountAIRunCfDeepseekAIExecuteDeepseekMath7bInstructParams, opts ...option.RequestOption) (res *AccountAIRunCfDeepseekAIExecuteDeepseekMath7bInstructResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai/run/@cf/deepseek-ai/deepseek-math-7b-instruct", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Execute @cf/deepseek-ai/deepseek-r1-distill-qwen-32b model.
func (r *AccountAIRunCfDeepseekAIService) ExecuteDeepseekR1DistillQwen32b(ctx context.Context, accountID string, params AccountAIRunCfDeepseekAIExecuteDeepseekR1DistillQwen32bParams, opts ...option.RequestOption) (res *AccountAIRunCfDeepseekAIExecuteDeepseekR1DistillQwen32bResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai/run/@cf/deepseek-ai/deepseek-r1-distill-qwen-32b", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type AccountAIRunCfDeepseekAIExecuteDeepseekMath7bInstructResponse = interface{}

type AccountAIRunCfDeepseekAIExecuteDeepseekR1DistillQwen32bResponse = interface{}

type AccountAIRunCfDeepseekAIExecuteDeepseekMath7bInstructParams struct {
	QueueRequest param.Field[string]                                                  `query:"queueRequest"`
	Body         AccountAIRunCfDeepseekAIExecuteDeepseekMath7bInstructParamsBodyUnion `json:"body"`
}

func (r AccountAIRunCfDeepseekAIExecuteDeepseekMath7bInstructParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

// URLQuery serializes
// [AccountAIRunCfDeepseekAIExecuteDeepseekMath7bInstructParams]'s query parameters
// as `url.Values`.
func (r AccountAIRunCfDeepseekAIExecuteDeepseekMath7bInstructParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountAIRunCfDeepseekAIExecuteDeepseekMath7bInstructParamsBody struct {
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

func (r AccountAIRunCfDeepseekAIExecuteDeepseekMath7bInstructParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfDeepseekAIExecuteDeepseekMath7bInstructParamsBody) implementsAccountAIRunCfDeepseekAIExecuteDeepseekMath7bInstructParamsBodyUnion() {
}

// Satisfied by
// [AccountAIRunCfDeepseekAIExecuteDeepseekMath7bInstructParamsBodyPrompt],
// [AccountAIRunCfDeepseekAIExecuteDeepseekMath7bInstructParamsBodyMessages],
// [AccountAIRunCfDeepseekAIExecuteDeepseekMath7bInstructParamsBody].
type AccountAIRunCfDeepseekAIExecuteDeepseekMath7bInstructParamsBodyUnion interface {
	implementsAccountAIRunCfDeepseekAIExecuteDeepseekMath7bInstructParamsBodyUnion()
}

type AccountAIRunCfDeepseekAIExecuteDeepseekMath7bInstructParamsBodyPrompt struct {
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
	ResponseFormat    param.Field[AccountAIRunCfDeepseekAIExecuteDeepseekMath7bInstructParamsBodyPromptResponseFormat] `json:"response_format"`
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

func (r AccountAIRunCfDeepseekAIExecuteDeepseekMath7bInstructParamsBodyPrompt) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfDeepseekAIExecuteDeepseekMath7bInstructParamsBodyPrompt) implementsAccountAIRunCfDeepseekAIExecuteDeepseekMath7bInstructParamsBodyUnion() {
}

type AccountAIRunCfDeepseekAIExecuteDeepseekMath7bInstructParamsBodyPromptResponseFormat struct {
	JsonSchema param.Field[interface{}]                                                                             `json:"json_schema"`
	Type       param.Field[AccountAIRunCfDeepseekAIExecuteDeepseekMath7bInstructParamsBodyPromptResponseFormatType] `json:"type"`
}

func (r AccountAIRunCfDeepseekAIExecuteDeepseekMath7bInstructParamsBodyPromptResponseFormat) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfDeepseekAIExecuteDeepseekMath7bInstructParamsBodyPromptResponseFormatType string

const (
	AccountAIRunCfDeepseekAIExecuteDeepseekMath7bInstructParamsBodyPromptResponseFormatTypeJsonObject AccountAIRunCfDeepseekAIExecuteDeepseekMath7bInstructParamsBodyPromptResponseFormatType = "json_object"
	AccountAIRunCfDeepseekAIExecuteDeepseekMath7bInstructParamsBodyPromptResponseFormatTypeJsonSchema AccountAIRunCfDeepseekAIExecuteDeepseekMath7bInstructParamsBodyPromptResponseFormatType = "json_schema"
)

func (r AccountAIRunCfDeepseekAIExecuteDeepseekMath7bInstructParamsBodyPromptResponseFormatType) IsKnown() bool {
	switch r {
	case AccountAIRunCfDeepseekAIExecuteDeepseekMath7bInstructParamsBodyPromptResponseFormatTypeJsonObject, AccountAIRunCfDeepseekAIExecuteDeepseekMath7bInstructParamsBodyPromptResponseFormatTypeJsonSchema:
		return true
	}
	return false
}

type AccountAIRunCfDeepseekAIExecuteDeepseekMath7bInstructParamsBodyMessages struct {
	// An array of message objects representing the conversation history.
	Messages param.Field[[]AccountAIRunCfDeepseekAIExecuteDeepseekMath7bInstructParamsBodyMessagesMessage] `json:"messages,required"`
	// Decreases the likelihood of the model repeating the same lines verbatim.
	FrequencyPenalty param.Field[float64]                                                                           `json:"frequency_penalty"`
	Functions        param.Field[[]AccountAIRunCfDeepseekAIExecuteDeepseekMath7bInstructParamsBodyMessagesFunction] `json:"functions"`
	// The maximum number of tokens to generate in the response.
	MaxTokens param.Field[int64] `json:"max_tokens"`
	// Increases the likelihood of the model introducing new topics.
	PresencePenalty param.Field[float64] `json:"presence_penalty"`
	// If true, a chat template is not applied and you must adhere to the specific
	// model's expected formatting.
	Raw param.Field[bool] `json:"raw"`
	// Penalty for repeated tokens; higher values discourage repetition.
	RepetitionPenalty param.Field[float64]                                                                               `json:"repetition_penalty"`
	ResponseFormat    param.Field[AccountAIRunCfDeepseekAIExecuteDeepseekMath7bInstructParamsBodyMessagesResponseFormat] `json:"response_format"`
	// Random seed for reproducibility of the generation.
	Seed param.Field[int64] `json:"seed"`
	// If true, the response will be streamed back incrementally using SSE, Server Sent
	// Events.
	Stream param.Field[bool] `json:"stream"`
	// Controls the randomness of the output; higher values produce more random
	// results.
	Temperature param.Field[float64] `json:"temperature"`
	// A list of tools available for the assistant to use.
	Tools param.Field[[]AccountAIRunCfDeepseekAIExecuteDeepseekMath7bInstructParamsBodyMessagesToolUnion] `json:"tools"`
	// Limits the AI to choose from the top 'k' most probable words. Lower values make
	// responses more focused; higher values introduce more variety and potential
	// surprises.
	TopK param.Field[int64] `json:"top_k"`
	// Adjusts the creativity of the AI's responses by controlling how many possible
	// words it considers. Lower values make outputs more predictable; higher values
	// allow for more varied and creative responses.
	TopP param.Field[float64] `json:"top_p"`
}

func (r AccountAIRunCfDeepseekAIExecuteDeepseekMath7bInstructParamsBodyMessages) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfDeepseekAIExecuteDeepseekMath7bInstructParamsBodyMessages) implementsAccountAIRunCfDeepseekAIExecuteDeepseekMath7bInstructParamsBodyUnion() {
}

type AccountAIRunCfDeepseekAIExecuteDeepseekMath7bInstructParamsBodyMessagesMessage struct {
	// The content of the message as a string.
	Content param.Field[string] `json:"content,required"`
	// The role of the message sender (e.g., 'user', 'assistant', 'system', 'tool').
	Role param.Field[string] `json:"role,required"`
}

func (r AccountAIRunCfDeepseekAIExecuteDeepseekMath7bInstructParamsBodyMessagesMessage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfDeepseekAIExecuteDeepseekMath7bInstructParamsBodyMessagesFunction struct {
	Code param.Field[string] `json:"code,required"`
	Name param.Field[string] `json:"name,required"`
}

func (r AccountAIRunCfDeepseekAIExecuteDeepseekMath7bInstructParamsBodyMessagesFunction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfDeepseekAIExecuteDeepseekMath7bInstructParamsBodyMessagesResponseFormat struct {
	JsonSchema param.Field[interface{}]                                                                               `json:"json_schema"`
	Type       param.Field[AccountAIRunCfDeepseekAIExecuteDeepseekMath7bInstructParamsBodyMessagesResponseFormatType] `json:"type"`
}

func (r AccountAIRunCfDeepseekAIExecuteDeepseekMath7bInstructParamsBodyMessagesResponseFormat) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfDeepseekAIExecuteDeepseekMath7bInstructParamsBodyMessagesResponseFormatType string

const (
	AccountAIRunCfDeepseekAIExecuteDeepseekMath7bInstructParamsBodyMessagesResponseFormatTypeJsonObject AccountAIRunCfDeepseekAIExecuteDeepseekMath7bInstructParamsBodyMessagesResponseFormatType = "json_object"
	AccountAIRunCfDeepseekAIExecuteDeepseekMath7bInstructParamsBodyMessagesResponseFormatTypeJsonSchema AccountAIRunCfDeepseekAIExecuteDeepseekMath7bInstructParamsBodyMessagesResponseFormatType = "json_schema"
)

func (r AccountAIRunCfDeepseekAIExecuteDeepseekMath7bInstructParamsBodyMessagesResponseFormatType) IsKnown() bool {
	switch r {
	case AccountAIRunCfDeepseekAIExecuteDeepseekMath7bInstructParamsBodyMessagesResponseFormatTypeJsonObject, AccountAIRunCfDeepseekAIExecuteDeepseekMath7bInstructParamsBodyMessagesResponseFormatTypeJsonSchema:
		return true
	}
	return false
}

type AccountAIRunCfDeepseekAIExecuteDeepseekMath7bInstructParamsBodyMessagesTool struct {
	// A brief description of what the tool does.
	Description param.Field[string]      `json:"description"`
	Function    param.Field[interface{}] `json:"function"`
	// The name of the tool. More descriptive the better.
	Name       param.Field[string]      `json:"name"`
	Parameters param.Field[interface{}] `json:"parameters"`
	// Specifies the type of tool (e.g., 'function').
	Type param.Field[string] `json:"type"`
}

func (r AccountAIRunCfDeepseekAIExecuteDeepseekMath7bInstructParamsBodyMessagesTool) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfDeepseekAIExecuteDeepseekMath7bInstructParamsBodyMessagesTool) implementsAccountAIRunCfDeepseekAIExecuteDeepseekMath7bInstructParamsBodyMessagesToolUnion() {
}

// Satisfied by
// [AccountAIRunCfDeepseekAIExecuteDeepseekMath7bInstructParamsBodyMessagesToolsObject],
// [AccountAIRunCfDeepseekAIExecuteDeepseekMath7bInstructParamsBodyMessagesToolsObject],
// [AccountAIRunCfDeepseekAIExecuteDeepseekMath7bInstructParamsBodyMessagesTool].
type AccountAIRunCfDeepseekAIExecuteDeepseekMath7bInstructParamsBodyMessagesToolUnion interface {
	implementsAccountAIRunCfDeepseekAIExecuteDeepseekMath7bInstructParamsBodyMessagesToolUnion()
}

type AccountAIRunCfDeepseekAIExecuteDeepseekMath7bInstructParamsBodyMessagesToolsObject struct {
	// A brief description of what the tool does.
	Description param.Field[string] `json:"description,required"`
	// The name of the tool. More descriptive the better.
	Name param.Field[string] `json:"name,required"`
	// Schema defining the parameters accepted by the tool.
	Parameters param.Field[AccountAIRunCfDeepseekAIExecuteDeepseekMath7bInstructParamsBodyMessagesToolsObjectParameters] `json:"parameters,required"`
}

func (r AccountAIRunCfDeepseekAIExecuteDeepseekMath7bInstructParamsBodyMessagesToolsObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfDeepseekAIExecuteDeepseekMath7bInstructParamsBodyMessagesToolsObject) implementsAccountAIRunCfDeepseekAIExecuteDeepseekMath7bInstructParamsBodyMessagesToolUnion() {
}

// Schema defining the parameters accepted by the tool.
type AccountAIRunCfDeepseekAIExecuteDeepseekMath7bInstructParamsBodyMessagesToolsObjectParameters struct {
	// Definitions of each parameter.
	Properties param.Field[map[string]AccountAIRunCfDeepseekAIExecuteDeepseekMath7bInstructParamsBodyMessagesToolsObjectParametersProperties] `json:"properties,required"`
	// The type of the parameters object (usually 'object').
	Type param.Field[string] `json:"type,required"`
	// List of required parameter names.
	Required param.Field[[]string] `json:"required"`
}

func (r AccountAIRunCfDeepseekAIExecuteDeepseekMath7bInstructParamsBodyMessagesToolsObjectParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfDeepseekAIExecuteDeepseekMath7bInstructParamsBodyMessagesToolsObjectParametersProperties struct {
	// A description of the expected parameter.
	Description param.Field[string] `json:"description,required"`
	// The data type of the parameter.
	Type param.Field[string] `json:"type,required"`
}

func (r AccountAIRunCfDeepseekAIExecuteDeepseekMath7bInstructParamsBodyMessagesToolsObjectParametersProperties) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfDeepseekAIExecuteDeepseekR1DistillQwen32bParams struct {
	QueueRequest param.Field[string]                                                    `query:"queueRequest"`
	Body         AccountAIRunCfDeepseekAIExecuteDeepseekR1DistillQwen32bParamsBodyUnion `json:"body"`
}

func (r AccountAIRunCfDeepseekAIExecuteDeepseekR1DistillQwen32bParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

// URLQuery serializes
// [AccountAIRunCfDeepseekAIExecuteDeepseekR1DistillQwen32bParams]'s query
// parameters as `url.Values`.
func (r AccountAIRunCfDeepseekAIExecuteDeepseekR1DistillQwen32bParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountAIRunCfDeepseekAIExecuteDeepseekR1DistillQwen32bParamsBody struct {
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

func (r AccountAIRunCfDeepseekAIExecuteDeepseekR1DistillQwen32bParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfDeepseekAIExecuteDeepseekR1DistillQwen32bParamsBody) implementsAccountAIRunCfDeepseekAIExecuteDeepseekR1DistillQwen32bParamsBodyUnion() {
}

// Satisfied by
// [AccountAIRunCfDeepseekAIExecuteDeepseekR1DistillQwen32bParamsBodyPrompt],
// [AccountAIRunCfDeepseekAIExecuteDeepseekR1DistillQwen32bParamsBodyMessages],
// [AccountAIRunCfDeepseekAIExecuteDeepseekR1DistillQwen32bParamsBody].
type AccountAIRunCfDeepseekAIExecuteDeepseekR1DistillQwen32bParamsBodyUnion interface {
	implementsAccountAIRunCfDeepseekAIExecuteDeepseekR1DistillQwen32bParamsBodyUnion()
}

type AccountAIRunCfDeepseekAIExecuteDeepseekR1DistillQwen32bParamsBodyPrompt struct {
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
	RepetitionPenalty param.Field[float64]                                                                               `json:"repetition_penalty"`
	ResponseFormat    param.Field[AccountAIRunCfDeepseekAIExecuteDeepseekR1DistillQwen32bParamsBodyPromptResponseFormat] `json:"response_format"`
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

func (r AccountAIRunCfDeepseekAIExecuteDeepseekR1DistillQwen32bParamsBodyPrompt) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfDeepseekAIExecuteDeepseekR1DistillQwen32bParamsBodyPrompt) implementsAccountAIRunCfDeepseekAIExecuteDeepseekR1DistillQwen32bParamsBodyUnion() {
}

type AccountAIRunCfDeepseekAIExecuteDeepseekR1DistillQwen32bParamsBodyPromptResponseFormat struct {
	JsonSchema param.Field[interface{}]                                                                               `json:"json_schema"`
	Type       param.Field[AccountAIRunCfDeepseekAIExecuteDeepseekR1DistillQwen32bParamsBodyPromptResponseFormatType] `json:"type"`
}

func (r AccountAIRunCfDeepseekAIExecuteDeepseekR1DistillQwen32bParamsBodyPromptResponseFormat) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfDeepseekAIExecuteDeepseekR1DistillQwen32bParamsBodyPromptResponseFormatType string

const (
	AccountAIRunCfDeepseekAIExecuteDeepseekR1DistillQwen32bParamsBodyPromptResponseFormatTypeJsonObject AccountAIRunCfDeepseekAIExecuteDeepseekR1DistillQwen32bParamsBodyPromptResponseFormatType = "json_object"
	AccountAIRunCfDeepseekAIExecuteDeepseekR1DistillQwen32bParamsBodyPromptResponseFormatTypeJsonSchema AccountAIRunCfDeepseekAIExecuteDeepseekR1DistillQwen32bParamsBodyPromptResponseFormatType = "json_schema"
)

func (r AccountAIRunCfDeepseekAIExecuteDeepseekR1DistillQwen32bParamsBodyPromptResponseFormatType) IsKnown() bool {
	switch r {
	case AccountAIRunCfDeepseekAIExecuteDeepseekR1DistillQwen32bParamsBodyPromptResponseFormatTypeJsonObject, AccountAIRunCfDeepseekAIExecuteDeepseekR1DistillQwen32bParamsBodyPromptResponseFormatTypeJsonSchema:
		return true
	}
	return false
}

type AccountAIRunCfDeepseekAIExecuteDeepseekR1DistillQwen32bParamsBodyMessages struct {
	// An array of message objects representing the conversation history.
	Messages param.Field[[]AccountAIRunCfDeepseekAIExecuteDeepseekR1DistillQwen32bParamsBodyMessagesMessage] `json:"messages,required"`
	// Decreases the likelihood of the model repeating the same lines verbatim.
	FrequencyPenalty param.Field[float64]                                                                             `json:"frequency_penalty"`
	Functions        param.Field[[]AccountAIRunCfDeepseekAIExecuteDeepseekR1DistillQwen32bParamsBodyMessagesFunction] `json:"functions"`
	// The maximum number of tokens to generate in the response.
	MaxTokens param.Field[int64] `json:"max_tokens"`
	// Increases the likelihood of the model introducing new topics.
	PresencePenalty param.Field[float64] `json:"presence_penalty"`
	// If true, a chat template is not applied and you must adhere to the specific
	// model's expected formatting.
	Raw param.Field[bool] `json:"raw"`
	// Penalty for repeated tokens; higher values discourage repetition.
	RepetitionPenalty param.Field[float64]                                                                                 `json:"repetition_penalty"`
	ResponseFormat    param.Field[AccountAIRunCfDeepseekAIExecuteDeepseekR1DistillQwen32bParamsBodyMessagesResponseFormat] `json:"response_format"`
	// Random seed for reproducibility of the generation.
	Seed param.Field[int64] `json:"seed"`
	// If true, the response will be streamed back incrementally using SSE, Server Sent
	// Events.
	Stream param.Field[bool] `json:"stream"`
	// Controls the randomness of the output; higher values produce more random
	// results.
	Temperature param.Field[float64] `json:"temperature"`
	// A list of tools available for the assistant to use.
	Tools param.Field[[]AccountAIRunCfDeepseekAIExecuteDeepseekR1DistillQwen32bParamsBodyMessagesToolUnion] `json:"tools"`
	// Limits the AI to choose from the top 'k' most probable words. Lower values make
	// responses more focused; higher values introduce more variety and potential
	// surprises.
	TopK param.Field[int64] `json:"top_k"`
	// Adjusts the creativity of the AI's responses by controlling how many possible
	// words it considers. Lower values make outputs more predictable; higher values
	// allow for more varied and creative responses.
	TopP param.Field[float64] `json:"top_p"`
}

func (r AccountAIRunCfDeepseekAIExecuteDeepseekR1DistillQwen32bParamsBodyMessages) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfDeepseekAIExecuteDeepseekR1DistillQwen32bParamsBodyMessages) implementsAccountAIRunCfDeepseekAIExecuteDeepseekR1DistillQwen32bParamsBodyUnion() {
}

type AccountAIRunCfDeepseekAIExecuteDeepseekR1DistillQwen32bParamsBodyMessagesMessage struct {
	// The content of the message as a string.
	Content param.Field[string] `json:"content,required"`
	// The role of the message sender (e.g., 'user', 'assistant', 'system', 'tool').
	Role param.Field[string] `json:"role,required"`
}

func (r AccountAIRunCfDeepseekAIExecuteDeepseekR1DistillQwen32bParamsBodyMessagesMessage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfDeepseekAIExecuteDeepseekR1DistillQwen32bParamsBodyMessagesFunction struct {
	Code param.Field[string] `json:"code,required"`
	Name param.Field[string] `json:"name,required"`
}

func (r AccountAIRunCfDeepseekAIExecuteDeepseekR1DistillQwen32bParamsBodyMessagesFunction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfDeepseekAIExecuteDeepseekR1DistillQwen32bParamsBodyMessagesResponseFormat struct {
	JsonSchema param.Field[interface{}]                                                                                 `json:"json_schema"`
	Type       param.Field[AccountAIRunCfDeepseekAIExecuteDeepseekR1DistillQwen32bParamsBodyMessagesResponseFormatType] `json:"type"`
}

func (r AccountAIRunCfDeepseekAIExecuteDeepseekR1DistillQwen32bParamsBodyMessagesResponseFormat) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfDeepseekAIExecuteDeepseekR1DistillQwen32bParamsBodyMessagesResponseFormatType string

const (
	AccountAIRunCfDeepseekAIExecuteDeepseekR1DistillQwen32bParamsBodyMessagesResponseFormatTypeJsonObject AccountAIRunCfDeepseekAIExecuteDeepseekR1DistillQwen32bParamsBodyMessagesResponseFormatType = "json_object"
	AccountAIRunCfDeepseekAIExecuteDeepseekR1DistillQwen32bParamsBodyMessagesResponseFormatTypeJsonSchema AccountAIRunCfDeepseekAIExecuteDeepseekR1DistillQwen32bParamsBodyMessagesResponseFormatType = "json_schema"
)

func (r AccountAIRunCfDeepseekAIExecuteDeepseekR1DistillQwen32bParamsBodyMessagesResponseFormatType) IsKnown() bool {
	switch r {
	case AccountAIRunCfDeepseekAIExecuteDeepseekR1DistillQwen32bParamsBodyMessagesResponseFormatTypeJsonObject, AccountAIRunCfDeepseekAIExecuteDeepseekR1DistillQwen32bParamsBodyMessagesResponseFormatTypeJsonSchema:
		return true
	}
	return false
}

type AccountAIRunCfDeepseekAIExecuteDeepseekR1DistillQwen32bParamsBodyMessagesTool struct {
	// A brief description of what the tool does.
	Description param.Field[string]      `json:"description"`
	Function    param.Field[interface{}] `json:"function"`
	// The name of the tool. More descriptive the better.
	Name       param.Field[string]      `json:"name"`
	Parameters param.Field[interface{}] `json:"parameters"`
	// Specifies the type of tool (e.g., 'function').
	Type param.Field[string] `json:"type"`
}

func (r AccountAIRunCfDeepseekAIExecuteDeepseekR1DistillQwen32bParamsBodyMessagesTool) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfDeepseekAIExecuteDeepseekR1DistillQwen32bParamsBodyMessagesTool) implementsAccountAIRunCfDeepseekAIExecuteDeepseekR1DistillQwen32bParamsBodyMessagesToolUnion() {
}

// Satisfied by
// [AccountAIRunCfDeepseekAIExecuteDeepseekR1DistillQwen32bParamsBodyMessagesToolsObject],
// [AccountAIRunCfDeepseekAIExecuteDeepseekR1DistillQwen32bParamsBodyMessagesToolsObject],
// [AccountAIRunCfDeepseekAIExecuteDeepseekR1DistillQwen32bParamsBodyMessagesTool].
type AccountAIRunCfDeepseekAIExecuteDeepseekR1DistillQwen32bParamsBodyMessagesToolUnion interface {
	implementsAccountAIRunCfDeepseekAIExecuteDeepseekR1DistillQwen32bParamsBodyMessagesToolUnion()
}

type AccountAIRunCfDeepseekAIExecuteDeepseekR1DistillQwen32bParamsBodyMessagesToolsObject struct {
	// A brief description of what the tool does.
	Description param.Field[string] `json:"description,required"`
	// The name of the tool. More descriptive the better.
	Name param.Field[string] `json:"name,required"`
	// Schema defining the parameters accepted by the tool.
	Parameters param.Field[AccountAIRunCfDeepseekAIExecuteDeepseekR1DistillQwen32bParamsBodyMessagesToolsObjectParameters] `json:"parameters,required"`
}

func (r AccountAIRunCfDeepseekAIExecuteDeepseekR1DistillQwen32bParamsBodyMessagesToolsObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfDeepseekAIExecuteDeepseekR1DistillQwen32bParamsBodyMessagesToolsObject) implementsAccountAIRunCfDeepseekAIExecuteDeepseekR1DistillQwen32bParamsBodyMessagesToolUnion() {
}

// Schema defining the parameters accepted by the tool.
type AccountAIRunCfDeepseekAIExecuteDeepseekR1DistillQwen32bParamsBodyMessagesToolsObjectParameters struct {
	// Definitions of each parameter.
	Properties param.Field[map[string]AccountAIRunCfDeepseekAIExecuteDeepseekR1DistillQwen32bParamsBodyMessagesToolsObjectParametersProperties] `json:"properties,required"`
	// The type of the parameters object (usually 'object').
	Type param.Field[string] `json:"type,required"`
	// List of required parameter names.
	Required param.Field[[]string] `json:"required"`
}

func (r AccountAIRunCfDeepseekAIExecuteDeepseekR1DistillQwen32bParamsBodyMessagesToolsObjectParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfDeepseekAIExecuteDeepseekR1DistillQwen32bParamsBodyMessagesToolsObjectParametersProperties struct {
	// A description of the expected parameter.
	Description param.Field[string] `json:"description,required"`
	// The data type of the parameter.
	Type param.Field[string] `json:"type,required"`
}

func (r AccountAIRunCfDeepseekAIExecuteDeepseekR1DistillQwen32bParamsBodyMessagesToolsObjectParametersProperties) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
