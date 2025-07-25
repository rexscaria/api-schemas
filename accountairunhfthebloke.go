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

// AccountAIRunHfTheblokeService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountAIRunHfTheblokeService] method instead.
type AccountAIRunHfTheblokeService struct {
	Options []option.RequestOption
}

// NewAccountAIRunHfTheblokeService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountAIRunHfTheblokeService(opts ...option.RequestOption) (r *AccountAIRunHfTheblokeService) {
	r = &AccountAIRunHfTheblokeService{}
	r.Options = opts
	return
}

// Execute @hf/thebloke/deepseek-coder-6.7b-base-awq model.
func (r *AccountAIRunHfTheblokeService) ExecuteDeepseekCoder6_7bBaseAwq(ctx context.Context, accountID string, params AccountAIRunHfTheblokeExecuteDeepseekCoder6_7bBaseAwqParams, opts ...option.RequestOption) (res *AccountAIRunHfTheblokeExecuteDeepseekCoder6_7bBaseAwqResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai/run/@hf/thebloke/deepseek-coder-6.7b-base-awq", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Execute @hf/thebloke/deepseek-coder-6.7b-instruct-awq model.
func (r *AccountAIRunHfTheblokeService) ExecuteDeepseekCoder6_7bInstructAwq(ctx context.Context, accountID string, params AccountAIRunHfTheblokeExecuteDeepseekCoder6_7bInstructAwqParams, opts ...option.RequestOption) (res *AccountAIRunHfTheblokeExecuteDeepseekCoder6_7bInstructAwqResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai/run/@hf/thebloke/deepseek-coder-6.7b-instruct-awq", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Execute @hf/thebloke/llama-2-13b-chat-awq model.
func (r *AccountAIRunHfTheblokeService) ExecuteLlama2_13bChatAwq(ctx context.Context, accountID string, params AccountAIRunHfTheblokeExecuteLlama2_13bChatAwqParams, opts ...option.RequestOption) (res *AccountAIRunHfTheblokeExecuteLlama2_13bChatAwqResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai/run/@hf/thebloke/llama-2-13b-chat-awq", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Execute @hf/thebloke/llamaguard-7b-awq model.
func (r *AccountAIRunHfTheblokeService) ExecuteLlamaguard7bAwq(ctx context.Context, accountID string, params AccountAIRunHfTheblokeExecuteLlamaguard7bAwqParams, opts ...option.RequestOption) (res *AccountAIRunHfTheblokeExecuteLlamaguard7bAwqResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai/run/@hf/thebloke/llamaguard-7b-awq", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Execute @hf/thebloke/mistral-7b-instruct-v0.1-awq model.
func (r *AccountAIRunHfTheblokeService) ExecuteMistral7bInstructV0_1Awq(ctx context.Context, accountID string, params AccountAIRunHfTheblokeExecuteMistral7bInstructV0_1AwqParams, opts ...option.RequestOption) (res *AccountAIRunHfTheblokeExecuteMistral7bInstructV0_1AwqResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai/run/@hf/thebloke/mistral-7b-instruct-v0.1-awq", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Execute @hf/thebloke/neural-chat-7b-v3-1-awq model.
func (r *AccountAIRunHfTheblokeService) ExecuteNeuralChat7bV3_1Awq(ctx context.Context, accountID string, params AccountAIRunHfTheblokeExecuteNeuralChat7bV3_1AwqParams, opts ...option.RequestOption) (res *AccountAIRunHfTheblokeExecuteNeuralChat7bV3_1AwqResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai/run/@hf/thebloke/neural-chat-7b-v3-1-awq", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Execute @hf/thebloke/openhermes-2.5-mistral-7b-awq model.
func (r *AccountAIRunHfTheblokeService) ExecuteOpenhermes2_5Mistral7bAwq(ctx context.Context, accountID string, params AccountAIRunHfTheblokeExecuteOpenhermes2_5Mistral7bAwqParams, opts ...option.RequestOption) (res *AccountAIRunHfTheblokeExecuteOpenhermes2_5Mistral7bAwqResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai/run/@hf/thebloke/openhermes-2.5-mistral-7b-awq", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Execute @hf/thebloke/zephyr-7b-beta-awq model.
func (r *AccountAIRunHfTheblokeService) ExecuteZephyr7bBetaAwq(ctx context.Context, accountID string, params AccountAIRunHfTheblokeExecuteZephyr7bBetaAwqParams, opts ...option.RequestOption) (res *AccountAIRunHfTheblokeExecuteZephyr7bBetaAwqResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai/run/@hf/thebloke/zephyr-7b-beta-awq", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type AccountAIRunHfTheblokeExecuteDeepseekCoder6_7bBaseAwqResponse = interface{}

type AccountAIRunHfTheblokeExecuteDeepseekCoder6_7bInstructAwqResponse = interface{}

type AccountAIRunHfTheblokeExecuteLlama2_13bChatAwqResponse = interface{}

type AccountAIRunHfTheblokeExecuteLlamaguard7bAwqResponse = interface{}

type AccountAIRunHfTheblokeExecuteMistral7bInstructV0_1AwqResponse = interface{}

type AccountAIRunHfTheblokeExecuteNeuralChat7bV3_1AwqResponse = interface{}

type AccountAIRunHfTheblokeExecuteOpenhermes2_5Mistral7bAwqResponse = interface{}

type AccountAIRunHfTheblokeExecuteZephyr7bBetaAwqResponse = interface{}

type AccountAIRunHfTheblokeExecuteDeepseekCoder6_7bBaseAwqParams struct {
	QueueRequest param.Field[string]                                                  `query:"queueRequest"`
	Body         AccountAIRunHfTheblokeExecuteDeepseekCoder6_7bBaseAwqParamsBodyUnion `json:"body"`
}

func (r AccountAIRunHfTheblokeExecuteDeepseekCoder6_7bBaseAwqParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

// URLQuery serializes
// [AccountAIRunHfTheblokeExecuteDeepseekCoder6_7bBaseAwqParams]'s query parameters
// as `url.Values`.
func (r AccountAIRunHfTheblokeExecuteDeepseekCoder6_7bBaseAwqParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountAIRunHfTheblokeExecuteDeepseekCoder6_7bBaseAwqParamsBody struct {
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

func (r AccountAIRunHfTheblokeExecuteDeepseekCoder6_7bBaseAwqParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunHfTheblokeExecuteDeepseekCoder6_7bBaseAwqParamsBody) implementsAccountAIRunHfTheblokeExecuteDeepseekCoder6_7bBaseAwqParamsBodyUnion() {
}

// Satisfied by
// [AccountAIRunHfTheblokeExecuteDeepseekCoder6_7bBaseAwqParamsBodyPrompt],
// [AccountAIRunHfTheblokeExecuteDeepseekCoder6_7bBaseAwqParamsBodyMessages],
// [AccountAIRunHfTheblokeExecuteDeepseekCoder6_7bBaseAwqParamsBody].
type AccountAIRunHfTheblokeExecuteDeepseekCoder6_7bBaseAwqParamsBodyUnion interface {
	implementsAccountAIRunHfTheblokeExecuteDeepseekCoder6_7bBaseAwqParamsBodyUnion()
}

type AccountAIRunHfTheblokeExecuteDeepseekCoder6_7bBaseAwqParamsBodyPrompt struct {
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
	ResponseFormat    param.Field[AccountAIRunHfTheblokeExecuteDeepseekCoder6_7bBaseAwqParamsBodyPromptResponseFormat] `json:"response_format"`
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

func (r AccountAIRunHfTheblokeExecuteDeepseekCoder6_7bBaseAwqParamsBodyPrompt) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunHfTheblokeExecuteDeepseekCoder6_7bBaseAwqParamsBodyPrompt) implementsAccountAIRunHfTheblokeExecuteDeepseekCoder6_7bBaseAwqParamsBodyUnion() {
}

type AccountAIRunHfTheblokeExecuteDeepseekCoder6_7bBaseAwqParamsBodyPromptResponseFormat struct {
	JsonSchema param.Field[interface{}]                                                                             `json:"json_schema"`
	Type       param.Field[AccountAIRunHfTheblokeExecuteDeepseekCoder6_7bBaseAwqParamsBodyPromptResponseFormatType] `json:"type"`
}

func (r AccountAIRunHfTheblokeExecuteDeepseekCoder6_7bBaseAwqParamsBodyPromptResponseFormat) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunHfTheblokeExecuteDeepseekCoder6_7bBaseAwqParamsBodyPromptResponseFormatType string

const (
	AccountAIRunHfTheblokeExecuteDeepseekCoder6_7bBaseAwqParamsBodyPromptResponseFormatTypeJsonObject AccountAIRunHfTheblokeExecuteDeepseekCoder6_7bBaseAwqParamsBodyPromptResponseFormatType = "json_object"
	AccountAIRunHfTheblokeExecuteDeepseekCoder6_7bBaseAwqParamsBodyPromptResponseFormatTypeJsonSchema AccountAIRunHfTheblokeExecuteDeepseekCoder6_7bBaseAwqParamsBodyPromptResponseFormatType = "json_schema"
)

func (r AccountAIRunHfTheblokeExecuteDeepseekCoder6_7bBaseAwqParamsBodyPromptResponseFormatType) IsKnown() bool {
	switch r {
	case AccountAIRunHfTheblokeExecuteDeepseekCoder6_7bBaseAwqParamsBodyPromptResponseFormatTypeJsonObject, AccountAIRunHfTheblokeExecuteDeepseekCoder6_7bBaseAwqParamsBodyPromptResponseFormatTypeJsonSchema:
		return true
	}
	return false
}

type AccountAIRunHfTheblokeExecuteDeepseekCoder6_7bBaseAwqParamsBodyMessages struct {
	// An array of message objects representing the conversation history.
	Messages param.Field[[]AccountAIRunHfTheblokeExecuteDeepseekCoder6_7bBaseAwqParamsBodyMessagesMessage] `json:"messages,required"`
	// Decreases the likelihood of the model repeating the same lines verbatim.
	FrequencyPenalty param.Field[float64]                                                                           `json:"frequency_penalty"`
	Functions        param.Field[[]AccountAIRunHfTheblokeExecuteDeepseekCoder6_7bBaseAwqParamsBodyMessagesFunction] `json:"functions"`
	// The maximum number of tokens to generate in the response.
	MaxTokens param.Field[int64] `json:"max_tokens"`
	// Increases the likelihood of the model introducing new topics.
	PresencePenalty param.Field[float64] `json:"presence_penalty"`
	// If true, a chat template is not applied and you must adhere to the specific
	// model's expected formatting.
	Raw param.Field[bool] `json:"raw"`
	// Penalty for repeated tokens; higher values discourage repetition.
	RepetitionPenalty param.Field[float64]                                                                               `json:"repetition_penalty"`
	ResponseFormat    param.Field[AccountAIRunHfTheblokeExecuteDeepseekCoder6_7bBaseAwqParamsBodyMessagesResponseFormat] `json:"response_format"`
	// Random seed for reproducibility of the generation.
	Seed param.Field[int64] `json:"seed"`
	// If true, the response will be streamed back incrementally using SSE, Server Sent
	// Events.
	Stream param.Field[bool] `json:"stream"`
	// Controls the randomness of the output; higher values produce more random
	// results.
	Temperature param.Field[float64] `json:"temperature"`
	// A list of tools available for the assistant to use.
	Tools param.Field[[]AccountAIRunHfTheblokeExecuteDeepseekCoder6_7bBaseAwqParamsBodyMessagesToolUnion] `json:"tools"`
	// Limits the AI to choose from the top 'k' most probable words. Lower values make
	// responses more focused; higher values introduce more variety and potential
	// surprises.
	TopK param.Field[int64] `json:"top_k"`
	// Adjusts the creativity of the AI's responses by controlling how many possible
	// words it considers. Lower values make outputs more predictable; higher values
	// allow for more varied and creative responses.
	TopP param.Field[float64] `json:"top_p"`
}

func (r AccountAIRunHfTheblokeExecuteDeepseekCoder6_7bBaseAwqParamsBodyMessages) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunHfTheblokeExecuteDeepseekCoder6_7bBaseAwqParamsBodyMessages) implementsAccountAIRunHfTheblokeExecuteDeepseekCoder6_7bBaseAwqParamsBodyUnion() {
}

type AccountAIRunHfTheblokeExecuteDeepseekCoder6_7bBaseAwqParamsBodyMessagesMessage struct {
	// The content of the message as a string.
	Content param.Field[string] `json:"content,required"`
	// The role of the message sender (e.g., 'user', 'assistant', 'system', 'tool').
	Role param.Field[string] `json:"role,required"`
}

func (r AccountAIRunHfTheblokeExecuteDeepseekCoder6_7bBaseAwqParamsBodyMessagesMessage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunHfTheblokeExecuteDeepseekCoder6_7bBaseAwqParamsBodyMessagesFunction struct {
	Code param.Field[string] `json:"code,required"`
	Name param.Field[string] `json:"name,required"`
}

func (r AccountAIRunHfTheblokeExecuteDeepseekCoder6_7bBaseAwqParamsBodyMessagesFunction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunHfTheblokeExecuteDeepseekCoder6_7bBaseAwqParamsBodyMessagesResponseFormat struct {
	JsonSchema param.Field[interface{}]                                                                               `json:"json_schema"`
	Type       param.Field[AccountAIRunHfTheblokeExecuteDeepseekCoder6_7bBaseAwqParamsBodyMessagesResponseFormatType] `json:"type"`
}

func (r AccountAIRunHfTheblokeExecuteDeepseekCoder6_7bBaseAwqParamsBodyMessagesResponseFormat) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunHfTheblokeExecuteDeepseekCoder6_7bBaseAwqParamsBodyMessagesResponseFormatType string

const (
	AccountAIRunHfTheblokeExecuteDeepseekCoder6_7bBaseAwqParamsBodyMessagesResponseFormatTypeJsonObject AccountAIRunHfTheblokeExecuteDeepseekCoder6_7bBaseAwqParamsBodyMessagesResponseFormatType = "json_object"
	AccountAIRunHfTheblokeExecuteDeepseekCoder6_7bBaseAwqParamsBodyMessagesResponseFormatTypeJsonSchema AccountAIRunHfTheblokeExecuteDeepseekCoder6_7bBaseAwqParamsBodyMessagesResponseFormatType = "json_schema"
)

func (r AccountAIRunHfTheblokeExecuteDeepseekCoder6_7bBaseAwqParamsBodyMessagesResponseFormatType) IsKnown() bool {
	switch r {
	case AccountAIRunHfTheblokeExecuteDeepseekCoder6_7bBaseAwqParamsBodyMessagesResponseFormatTypeJsonObject, AccountAIRunHfTheblokeExecuteDeepseekCoder6_7bBaseAwqParamsBodyMessagesResponseFormatTypeJsonSchema:
		return true
	}
	return false
}

type AccountAIRunHfTheblokeExecuteDeepseekCoder6_7bBaseAwqParamsBodyMessagesTool struct {
	// A brief description of what the tool does.
	Description param.Field[string]      `json:"description"`
	Function    param.Field[interface{}] `json:"function"`
	// The name of the tool. More descriptive the better.
	Name       param.Field[string]      `json:"name"`
	Parameters param.Field[interface{}] `json:"parameters"`
	// Specifies the type of tool (e.g., 'function').
	Type param.Field[string] `json:"type"`
}

func (r AccountAIRunHfTheblokeExecuteDeepseekCoder6_7bBaseAwqParamsBodyMessagesTool) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunHfTheblokeExecuteDeepseekCoder6_7bBaseAwqParamsBodyMessagesTool) implementsAccountAIRunHfTheblokeExecuteDeepseekCoder6_7bBaseAwqParamsBodyMessagesToolUnion() {
}

// Satisfied by
// [AccountAIRunHfTheblokeExecuteDeepseekCoder6_7bBaseAwqParamsBodyMessagesToolsObject],
// [AccountAIRunHfTheblokeExecuteDeepseekCoder6_7bBaseAwqParamsBodyMessagesToolsObject],
// [AccountAIRunHfTheblokeExecuteDeepseekCoder6_7bBaseAwqParamsBodyMessagesTool].
type AccountAIRunHfTheblokeExecuteDeepseekCoder6_7bBaseAwqParamsBodyMessagesToolUnion interface {
	implementsAccountAIRunHfTheblokeExecuteDeepseekCoder6_7bBaseAwqParamsBodyMessagesToolUnion()
}

type AccountAIRunHfTheblokeExecuteDeepseekCoder6_7bBaseAwqParamsBodyMessagesToolsObject struct {
	// A brief description of what the tool does.
	Description param.Field[string] `json:"description,required"`
	// The name of the tool. More descriptive the better.
	Name param.Field[string] `json:"name,required"`
	// Schema defining the parameters accepted by the tool.
	Parameters param.Field[AccountAIRunHfTheblokeExecuteDeepseekCoder6_7bBaseAwqParamsBodyMessagesToolsObjectParameters] `json:"parameters,required"`
}

func (r AccountAIRunHfTheblokeExecuteDeepseekCoder6_7bBaseAwqParamsBodyMessagesToolsObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunHfTheblokeExecuteDeepseekCoder6_7bBaseAwqParamsBodyMessagesToolsObject) implementsAccountAIRunHfTheblokeExecuteDeepseekCoder6_7bBaseAwqParamsBodyMessagesToolUnion() {
}

// Schema defining the parameters accepted by the tool.
type AccountAIRunHfTheblokeExecuteDeepseekCoder6_7bBaseAwqParamsBodyMessagesToolsObjectParameters struct {
	// Definitions of each parameter.
	Properties param.Field[map[string]AccountAIRunHfTheblokeExecuteDeepseekCoder6_7bBaseAwqParamsBodyMessagesToolsObjectParametersProperties] `json:"properties,required"`
	// The type of the parameters object (usually 'object').
	Type param.Field[string] `json:"type,required"`
	// List of required parameter names.
	Required param.Field[[]string] `json:"required"`
}

func (r AccountAIRunHfTheblokeExecuteDeepseekCoder6_7bBaseAwqParamsBodyMessagesToolsObjectParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunHfTheblokeExecuteDeepseekCoder6_7bBaseAwqParamsBodyMessagesToolsObjectParametersProperties struct {
	// A description of the expected parameter.
	Description param.Field[string] `json:"description,required"`
	// The data type of the parameter.
	Type param.Field[string] `json:"type,required"`
}

func (r AccountAIRunHfTheblokeExecuteDeepseekCoder6_7bBaseAwqParamsBodyMessagesToolsObjectParametersProperties) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunHfTheblokeExecuteDeepseekCoder6_7bInstructAwqParams struct {
	QueueRequest param.Field[string]                                                      `query:"queueRequest"`
	Body         AccountAIRunHfTheblokeExecuteDeepseekCoder6_7bInstructAwqParamsBodyUnion `json:"body"`
}

func (r AccountAIRunHfTheblokeExecuteDeepseekCoder6_7bInstructAwqParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

// URLQuery serializes
// [AccountAIRunHfTheblokeExecuteDeepseekCoder6_7bInstructAwqParams]'s query
// parameters as `url.Values`.
func (r AccountAIRunHfTheblokeExecuteDeepseekCoder6_7bInstructAwqParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountAIRunHfTheblokeExecuteDeepseekCoder6_7bInstructAwqParamsBody struct {
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

func (r AccountAIRunHfTheblokeExecuteDeepseekCoder6_7bInstructAwqParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunHfTheblokeExecuteDeepseekCoder6_7bInstructAwqParamsBody) implementsAccountAIRunHfTheblokeExecuteDeepseekCoder6_7bInstructAwqParamsBodyUnion() {
}

// Satisfied by
// [AccountAIRunHfTheblokeExecuteDeepseekCoder6_7bInstructAwqParamsBodyPrompt],
// [AccountAIRunHfTheblokeExecuteDeepseekCoder6_7bInstructAwqParamsBodyMessages],
// [AccountAIRunHfTheblokeExecuteDeepseekCoder6_7bInstructAwqParamsBody].
type AccountAIRunHfTheblokeExecuteDeepseekCoder6_7bInstructAwqParamsBodyUnion interface {
	implementsAccountAIRunHfTheblokeExecuteDeepseekCoder6_7bInstructAwqParamsBodyUnion()
}

type AccountAIRunHfTheblokeExecuteDeepseekCoder6_7bInstructAwqParamsBodyPrompt struct {
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
	RepetitionPenalty param.Field[float64]                                                                                 `json:"repetition_penalty"`
	ResponseFormat    param.Field[AccountAIRunHfTheblokeExecuteDeepseekCoder6_7bInstructAwqParamsBodyPromptResponseFormat] `json:"response_format"`
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

func (r AccountAIRunHfTheblokeExecuteDeepseekCoder6_7bInstructAwqParamsBodyPrompt) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunHfTheblokeExecuteDeepseekCoder6_7bInstructAwqParamsBodyPrompt) implementsAccountAIRunHfTheblokeExecuteDeepseekCoder6_7bInstructAwqParamsBodyUnion() {
}

type AccountAIRunHfTheblokeExecuteDeepseekCoder6_7bInstructAwqParamsBodyPromptResponseFormat struct {
	JsonSchema param.Field[interface{}]                                                                                 `json:"json_schema"`
	Type       param.Field[AccountAIRunHfTheblokeExecuteDeepseekCoder6_7bInstructAwqParamsBodyPromptResponseFormatType] `json:"type"`
}

func (r AccountAIRunHfTheblokeExecuteDeepseekCoder6_7bInstructAwqParamsBodyPromptResponseFormat) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunHfTheblokeExecuteDeepseekCoder6_7bInstructAwqParamsBodyPromptResponseFormatType string

const (
	AccountAIRunHfTheblokeExecuteDeepseekCoder6_7bInstructAwqParamsBodyPromptResponseFormatTypeJsonObject AccountAIRunHfTheblokeExecuteDeepseekCoder6_7bInstructAwqParamsBodyPromptResponseFormatType = "json_object"
	AccountAIRunHfTheblokeExecuteDeepseekCoder6_7bInstructAwqParamsBodyPromptResponseFormatTypeJsonSchema AccountAIRunHfTheblokeExecuteDeepseekCoder6_7bInstructAwqParamsBodyPromptResponseFormatType = "json_schema"
)

func (r AccountAIRunHfTheblokeExecuteDeepseekCoder6_7bInstructAwqParamsBodyPromptResponseFormatType) IsKnown() bool {
	switch r {
	case AccountAIRunHfTheblokeExecuteDeepseekCoder6_7bInstructAwqParamsBodyPromptResponseFormatTypeJsonObject, AccountAIRunHfTheblokeExecuteDeepseekCoder6_7bInstructAwqParamsBodyPromptResponseFormatTypeJsonSchema:
		return true
	}
	return false
}

type AccountAIRunHfTheblokeExecuteDeepseekCoder6_7bInstructAwqParamsBodyMessages struct {
	// An array of message objects representing the conversation history.
	Messages param.Field[[]AccountAIRunHfTheblokeExecuteDeepseekCoder6_7bInstructAwqParamsBodyMessagesMessage] `json:"messages,required"`
	// Decreases the likelihood of the model repeating the same lines verbatim.
	FrequencyPenalty param.Field[float64]                                                                               `json:"frequency_penalty"`
	Functions        param.Field[[]AccountAIRunHfTheblokeExecuteDeepseekCoder6_7bInstructAwqParamsBodyMessagesFunction] `json:"functions"`
	// The maximum number of tokens to generate in the response.
	MaxTokens param.Field[int64] `json:"max_tokens"`
	// Increases the likelihood of the model introducing new topics.
	PresencePenalty param.Field[float64] `json:"presence_penalty"`
	// If true, a chat template is not applied and you must adhere to the specific
	// model's expected formatting.
	Raw param.Field[bool] `json:"raw"`
	// Penalty for repeated tokens; higher values discourage repetition.
	RepetitionPenalty param.Field[float64]                                                                                   `json:"repetition_penalty"`
	ResponseFormat    param.Field[AccountAIRunHfTheblokeExecuteDeepseekCoder6_7bInstructAwqParamsBodyMessagesResponseFormat] `json:"response_format"`
	// Random seed for reproducibility of the generation.
	Seed param.Field[int64] `json:"seed"`
	// If true, the response will be streamed back incrementally using SSE, Server Sent
	// Events.
	Stream param.Field[bool] `json:"stream"`
	// Controls the randomness of the output; higher values produce more random
	// results.
	Temperature param.Field[float64] `json:"temperature"`
	// A list of tools available for the assistant to use.
	Tools param.Field[[]AccountAIRunHfTheblokeExecuteDeepseekCoder6_7bInstructAwqParamsBodyMessagesToolUnion] `json:"tools"`
	// Limits the AI to choose from the top 'k' most probable words. Lower values make
	// responses more focused; higher values introduce more variety and potential
	// surprises.
	TopK param.Field[int64] `json:"top_k"`
	// Adjusts the creativity of the AI's responses by controlling how many possible
	// words it considers. Lower values make outputs more predictable; higher values
	// allow for more varied and creative responses.
	TopP param.Field[float64] `json:"top_p"`
}

func (r AccountAIRunHfTheblokeExecuteDeepseekCoder6_7bInstructAwqParamsBodyMessages) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunHfTheblokeExecuteDeepseekCoder6_7bInstructAwqParamsBodyMessages) implementsAccountAIRunHfTheblokeExecuteDeepseekCoder6_7bInstructAwqParamsBodyUnion() {
}

type AccountAIRunHfTheblokeExecuteDeepseekCoder6_7bInstructAwqParamsBodyMessagesMessage struct {
	// The content of the message as a string.
	Content param.Field[string] `json:"content,required"`
	// The role of the message sender (e.g., 'user', 'assistant', 'system', 'tool').
	Role param.Field[string] `json:"role,required"`
}

func (r AccountAIRunHfTheblokeExecuteDeepseekCoder6_7bInstructAwqParamsBodyMessagesMessage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunHfTheblokeExecuteDeepseekCoder6_7bInstructAwqParamsBodyMessagesFunction struct {
	Code param.Field[string] `json:"code,required"`
	Name param.Field[string] `json:"name,required"`
}

func (r AccountAIRunHfTheblokeExecuteDeepseekCoder6_7bInstructAwqParamsBodyMessagesFunction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunHfTheblokeExecuteDeepseekCoder6_7bInstructAwqParamsBodyMessagesResponseFormat struct {
	JsonSchema param.Field[interface{}]                                                                                   `json:"json_schema"`
	Type       param.Field[AccountAIRunHfTheblokeExecuteDeepseekCoder6_7bInstructAwqParamsBodyMessagesResponseFormatType] `json:"type"`
}

func (r AccountAIRunHfTheblokeExecuteDeepseekCoder6_7bInstructAwqParamsBodyMessagesResponseFormat) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunHfTheblokeExecuteDeepseekCoder6_7bInstructAwqParamsBodyMessagesResponseFormatType string

const (
	AccountAIRunHfTheblokeExecuteDeepseekCoder6_7bInstructAwqParamsBodyMessagesResponseFormatTypeJsonObject AccountAIRunHfTheblokeExecuteDeepseekCoder6_7bInstructAwqParamsBodyMessagesResponseFormatType = "json_object"
	AccountAIRunHfTheblokeExecuteDeepseekCoder6_7bInstructAwqParamsBodyMessagesResponseFormatTypeJsonSchema AccountAIRunHfTheblokeExecuteDeepseekCoder6_7bInstructAwqParamsBodyMessagesResponseFormatType = "json_schema"
)

func (r AccountAIRunHfTheblokeExecuteDeepseekCoder6_7bInstructAwqParamsBodyMessagesResponseFormatType) IsKnown() bool {
	switch r {
	case AccountAIRunHfTheblokeExecuteDeepseekCoder6_7bInstructAwqParamsBodyMessagesResponseFormatTypeJsonObject, AccountAIRunHfTheblokeExecuteDeepseekCoder6_7bInstructAwqParamsBodyMessagesResponseFormatTypeJsonSchema:
		return true
	}
	return false
}

type AccountAIRunHfTheblokeExecuteDeepseekCoder6_7bInstructAwqParamsBodyMessagesTool struct {
	// A brief description of what the tool does.
	Description param.Field[string]      `json:"description"`
	Function    param.Field[interface{}] `json:"function"`
	// The name of the tool. More descriptive the better.
	Name       param.Field[string]      `json:"name"`
	Parameters param.Field[interface{}] `json:"parameters"`
	// Specifies the type of tool (e.g., 'function').
	Type param.Field[string] `json:"type"`
}

func (r AccountAIRunHfTheblokeExecuteDeepseekCoder6_7bInstructAwqParamsBodyMessagesTool) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunHfTheblokeExecuteDeepseekCoder6_7bInstructAwqParamsBodyMessagesTool) implementsAccountAIRunHfTheblokeExecuteDeepseekCoder6_7bInstructAwqParamsBodyMessagesToolUnion() {
}

// Satisfied by
// [AccountAIRunHfTheblokeExecuteDeepseekCoder6_7bInstructAwqParamsBodyMessagesToolsObject],
// [AccountAIRunHfTheblokeExecuteDeepseekCoder6_7bInstructAwqParamsBodyMessagesToolsObject],
// [AccountAIRunHfTheblokeExecuteDeepseekCoder6_7bInstructAwqParamsBodyMessagesTool].
type AccountAIRunHfTheblokeExecuteDeepseekCoder6_7bInstructAwqParamsBodyMessagesToolUnion interface {
	implementsAccountAIRunHfTheblokeExecuteDeepseekCoder6_7bInstructAwqParamsBodyMessagesToolUnion()
}

type AccountAIRunHfTheblokeExecuteDeepseekCoder6_7bInstructAwqParamsBodyMessagesToolsObject struct {
	// A brief description of what the tool does.
	Description param.Field[string] `json:"description,required"`
	// The name of the tool. More descriptive the better.
	Name param.Field[string] `json:"name,required"`
	// Schema defining the parameters accepted by the tool.
	Parameters param.Field[AccountAIRunHfTheblokeExecuteDeepseekCoder6_7bInstructAwqParamsBodyMessagesToolsObjectParameters] `json:"parameters,required"`
}

func (r AccountAIRunHfTheblokeExecuteDeepseekCoder6_7bInstructAwqParamsBodyMessagesToolsObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunHfTheblokeExecuteDeepseekCoder6_7bInstructAwqParamsBodyMessagesToolsObject) implementsAccountAIRunHfTheblokeExecuteDeepseekCoder6_7bInstructAwqParamsBodyMessagesToolUnion() {
}

// Schema defining the parameters accepted by the tool.
type AccountAIRunHfTheblokeExecuteDeepseekCoder6_7bInstructAwqParamsBodyMessagesToolsObjectParameters struct {
	// Definitions of each parameter.
	Properties param.Field[map[string]AccountAIRunHfTheblokeExecuteDeepseekCoder6_7bInstructAwqParamsBodyMessagesToolsObjectParametersProperties] `json:"properties,required"`
	// The type of the parameters object (usually 'object').
	Type param.Field[string] `json:"type,required"`
	// List of required parameter names.
	Required param.Field[[]string] `json:"required"`
}

func (r AccountAIRunHfTheblokeExecuteDeepseekCoder6_7bInstructAwqParamsBodyMessagesToolsObjectParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunHfTheblokeExecuteDeepseekCoder6_7bInstructAwqParamsBodyMessagesToolsObjectParametersProperties struct {
	// A description of the expected parameter.
	Description param.Field[string] `json:"description,required"`
	// The data type of the parameter.
	Type param.Field[string] `json:"type,required"`
}

func (r AccountAIRunHfTheblokeExecuteDeepseekCoder6_7bInstructAwqParamsBodyMessagesToolsObjectParametersProperties) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunHfTheblokeExecuteLlama2_13bChatAwqParams struct {
	QueueRequest param.Field[string]                                           `query:"queueRequest"`
	Body         AccountAIRunHfTheblokeExecuteLlama2_13bChatAwqParamsBodyUnion `json:"body"`
}

func (r AccountAIRunHfTheblokeExecuteLlama2_13bChatAwqParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

// URLQuery serializes [AccountAIRunHfTheblokeExecuteLlama2_13bChatAwqParams]'s
// query parameters as `url.Values`.
func (r AccountAIRunHfTheblokeExecuteLlama2_13bChatAwqParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountAIRunHfTheblokeExecuteLlama2_13bChatAwqParamsBody struct {
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

func (r AccountAIRunHfTheblokeExecuteLlama2_13bChatAwqParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunHfTheblokeExecuteLlama2_13bChatAwqParamsBody) implementsAccountAIRunHfTheblokeExecuteLlama2_13bChatAwqParamsBodyUnion() {
}

// Satisfied by [AccountAIRunHfTheblokeExecuteLlama2_13bChatAwqParamsBodyPrompt],
// [AccountAIRunHfTheblokeExecuteLlama2_13bChatAwqParamsBodyMessages],
// [AccountAIRunHfTheblokeExecuteLlama2_13bChatAwqParamsBody].
type AccountAIRunHfTheblokeExecuteLlama2_13bChatAwqParamsBodyUnion interface {
	implementsAccountAIRunHfTheblokeExecuteLlama2_13bChatAwqParamsBodyUnion()
}

type AccountAIRunHfTheblokeExecuteLlama2_13bChatAwqParamsBodyPrompt struct {
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
	RepetitionPenalty param.Field[float64]                                                                      `json:"repetition_penalty"`
	ResponseFormat    param.Field[AccountAIRunHfTheblokeExecuteLlama2_13bChatAwqParamsBodyPromptResponseFormat] `json:"response_format"`
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

func (r AccountAIRunHfTheblokeExecuteLlama2_13bChatAwqParamsBodyPrompt) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunHfTheblokeExecuteLlama2_13bChatAwqParamsBodyPrompt) implementsAccountAIRunHfTheblokeExecuteLlama2_13bChatAwqParamsBodyUnion() {
}

type AccountAIRunHfTheblokeExecuteLlama2_13bChatAwqParamsBodyPromptResponseFormat struct {
	JsonSchema param.Field[interface{}]                                                                      `json:"json_schema"`
	Type       param.Field[AccountAIRunHfTheblokeExecuteLlama2_13bChatAwqParamsBodyPromptResponseFormatType] `json:"type"`
}

func (r AccountAIRunHfTheblokeExecuteLlama2_13bChatAwqParamsBodyPromptResponseFormat) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunHfTheblokeExecuteLlama2_13bChatAwqParamsBodyPromptResponseFormatType string

const (
	AccountAIRunHfTheblokeExecuteLlama2_13bChatAwqParamsBodyPromptResponseFormatTypeJsonObject AccountAIRunHfTheblokeExecuteLlama2_13bChatAwqParamsBodyPromptResponseFormatType = "json_object"
	AccountAIRunHfTheblokeExecuteLlama2_13bChatAwqParamsBodyPromptResponseFormatTypeJsonSchema AccountAIRunHfTheblokeExecuteLlama2_13bChatAwqParamsBodyPromptResponseFormatType = "json_schema"
)

func (r AccountAIRunHfTheblokeExecuteLlama2_13bChatAwqParamsBodyPromptResponseFormatType) IsKnown() bool {
	switch r {
	case AccountAIRunHfTheblokeExecuteLlama2_13bChatAwqParamsBodyPromptResponseFormatTypeJsonObject, AccountAIRunHfTheblokeExecuteLlama2_13bChatAwqParamsBodyPromptResponseFormatTypeJsonSchema:
		return true
	}
	return false
}

type AccountAIRunHfTheblokeExecuteLlama2_13bChatAwqParamsBodyMessages struct {
	// An array of message objects representing the conversation history.
	Messages param.Field[[]AccountAIRunHfTheblokeExecuteLlama2_13bChatAwqParamsBodyMessagesMessage] `json:"messages,required"`
	// Decreases the likelihood of the model repeating the same lines verbatim.
	FrequencyPenalty param.Field[float64]                                                                    `json:"frequency_penalty"`
	Functions        param.Field[[]AccountAIRunHfTheblokeExecuteLlama2_13bChatAwqParamsBodyMessagesFunction] `json:"functions"`
	// The maximum number of tokens to generate in the response.
	MaxTokens param.Field[int64] `json:"max_tokens"`
	// Increases the likelihood of the model introducing new topics.
	PresencePenalty param.Field[float64] `json:"presence_penalty"`
	// If true, a chat template is not applied and you must adhere to the specific
	// model's expected formatting.
	Raw param.Field[bool] `json:"raw"`
	// Penalty for repeated tokens; higher values discourage repetition.
	RepetitionPenalty param.Field[float64]                                                                        `json:"repetition_penalty"`
	ResponseFormat    param.Field[AccountAIRunHfTheblokeExecuteLlama2_13bChatAwqParamsBodyMessagesResponseFormat] `json:"response_format"`
	// Random seed for reproducibility of the generation.
	Seed param.Field[int64] `json:"seed"`
	// If true, the response will be streamed back incrementally using SSE, Server Sent
	// Events.
	Stream param.Field[bool] `json:"stream"`
	// Controls the randomness of the output; higher values produce more random
	// results.
	Temperature param.Field[float64] `json:"temperature"`
	// A list of tools available for the assistant to use.
	Tools param.Field[[]AccountAIRunHfTheblokeExecuteLlama2_13bChatAwqParamsBodyMessagesToolUnion] `json:"tools"`
	// Limits the AI to choose from the top 'k' most probable words. Lower values make
	// responses more focused; higher values introduce more variety and potential
	// surprises.
	TopK param.Field[int64] `json:"top_k"`
	// Adjusts the creativity of the AI's responses by controlling how many possible
	// words it considers. Lower values make outputs more predictable; higher values
	// allow for more varied and creative responses.
	TopP param.Field[float64] `json:"top_p"`
}

func (r AccountAIRunHfTheblokeExecuteLlama2_13bChatAwqParamsBodyMessages) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunHfTheblokeExecuteLlama2_13bChatAwqParamsBodyMessages) implementsAccountAIRunHfTheblokeExecuteLlama2_13bChatAwqParamsBodyUnion() {
}

type AccountAIRunHfTheblokeExecuteLlama2_13bChatAwqParamsBodyMessagesMessage struct {
	// The content of the message as a string.
	Content param.Field[string] `json:"content,required"`
	// The role of the message sender (e.g., 'user', 'assistant', 'system', 'tool').
	Role param.Field[string] `json:"role,required"`
}

func (r AccountAIRunHfTheblokeExecuteLlama2_13bChatAwqParamsBodyMessagesMessage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunHfTheblokeExecuteLlama2_13bChatAwqParamsBodyMessagesFunction struct {
	Code param.Field[string] `json:"code,required"`
	Name param.Field[string] `json:"name,required"`
}

func (r AccountAIRunHfTheblokeExecuteLlama2_13bChatAwqParamsBodyMessagesFunction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunHfTheblokeExecuteLlama2_13bChatAwqParamsBodyMessagesResponseFormat struct {
	JsonSchema param.Field[interface{}]                                                                        `json:"json_schema"`
	Type       param.Field[AccountAIRunHfTheblokeExecuteLlama2_13bChatAwqParamsBodyMessagesResponseFormatType] `json:"type"`
}

func (r AccountAIRunHfTheblokeExecuteLlama2_13bChatAwqParamsBodyMessagesResponseFormat) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunHfTheblokeExecuteLlama2_13bChatAwqParamsBodyMessagesResponseFormatType string

const (
	AccountAIRunHfTheblokeExecuteLlama2_13bChatAwqParamsBodyMessagesResponseFormatTypeJsonObject AccountAIRunHfTheblokeExecuteLlama2_13bChatAwqParamsBodyMessagesResponseFormatType = "json_object"
	AccountAIRunHfTheblokeExecuteLlama2_13bChatAwqParamsBodyMessagesResponseFormatTypeJsonSchema AccountAIRunHfTheblokeExecuteLlama2_13bChatAwqParamsBodyMessagesResponseFormatType = "json_schema"
)

func (r AccountAIRunHfTheblokeExecuteLlama2_13bChatAwqParamsBodyMessagesResponseFormatType) IsKnown() bool {
	switch r {
	case AccountAIRunHfTheblokeExecuteLlama2_13bChatAwqParamsBodyMessagesResponseFormatTypeJsonObject, AccountAIRunHfTheblokeExecuteLlama2_13bChatAwqParamsBodyMessagesResponseFormatTypeJsonSchema:
		return true
	}
	return false
}

type AccountAIRunHfTheblokeExecuteLlama2_13bChatAwqParamsBodyMessagesTool struct {
	// A brief description of what the tool does.
	Description param.Field[string]      `json:"description"`
	Function    param.Field[interface{}] `json:"function"`
	// The name of the tool. More descriptive the better.
	Name       param.Field[string]      `json:"name"`
	Parameters param.Field[interface{}] `json:"parameters"`
	// Specifies the type of tool (e.g., 'function').
	Type param.Field[string] `json:"type"`
}

func (r AccountAIRunHfTheblokeExecuteLlama2_13bChatAwqParamsBodyMessagesTool) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunHfTheblokeExecuteLlama2_13bChatAwqParamsBodyMessagesTool) implementsAccountAIRunHfTheblokeExecuteLlama2_13bChatAwqParamsBodyMessagesToolUnion() {
}

// Satisfied by
// [AccountAIRunHfTheblokeExecuteLlama2_13bChatAwqParamsBodyMessagesToolsObject],
// [AccountAIRunHfTheblokeExecuteLlama2_13bChatAwqParamsBodyMessagesToolsObject],
// [AccountAIRunHfTheblokeExecuteLlama2_13bChatAwqParamsBodyMessagesTool].
type AccountAIRunHfTheblokeExecuteLlama2_13bChatAwqParamsBodyMessagesToolUnion interface {
	implementsAccountAIRunHfTheblokeExecuteLlama2_13bChatAwqParamsBodyMessagesToolUnion()
}

type AccountAIRunHfTheblokeExecuteLlama2_13bChatAwqParamsBodyMessagesToolsObject struct {
	// A brief description of what the tool does.
	Description param.Field[string] `json:"description,required"`
	// The name of the tool. More descriptive the better.
	Name param.Field[string] `json:"name,required"`
	// Schema defining the parameters accepted by the tool.
	Parameters param.Field[AccountAIRunHfTheblokeExecuteLlama2_13bChatAwqParamsBodyMessagesToolsObjectParameters] `json:"parameters,required"`
}

func (r AccountAIRunHfTheblokeExecuteLlama2_13bChatAwqParamsBodyMessagesToolsObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunHfTheblokeExecuteLlama2_13bChatAwqParamsBodyMessagesToolsObject) implementsAccountAIRunHfTheblokeExecuteLlama2_13bChatAwqParamsBodyMessagesToolUnion() {
}

// Schema defining the parameters accepted by the tool.
type AccountAIRunHfTheblokeExecuteLlama2_13bChatAwqParamsBodyMessagesToolsObjectParameters struct {
	// Definitions of each parameter.
	Properties param.Field[map[string]AccountAIRunHfTheblokeExecuteLlama2_13bChatAwqParamsBodyMessagesToolsObjectParametersProperties] `json:"properties,required"`
	// The type of the parameters object (usually 'object').
	Type param.Field[string] `json:"type,required"`
	// List of required parameter names.
	Required param.Field[[]string] `json:"required"`
}

func (r AccountAIRunHfTheblokeExecuteLlama2_13bChatAwqParamsBodyMessagesToolsObjectParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunHfTheblokeExecuteLlama2_13bChatAwqParamsBodyMessagesToolsObjectParametersProperties struct {
	// A description of the expected parameter.
	Description param.Field[string] `json:"description,required"`
	// The data type of the parameter.
	Type param.Field[string] `json:"type,required"`
}

func (r AccountAIRunHfTheblokeExecuteLlama2_13bChatAwqParamsBodyMessagesToolsObjectParametersProperties) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunHfTheblokeExecuteLlamaguard7bAwqParams struct {
	QueueRequest param.Field[string]                                         `query:"queueRequest"`
	Body         AccountAIRunHfTheblokeExecuteLlamaguard7bAwqParamsBodyUnion `json:"body"`
}

func (r AccountAIRunHfTheblokeExecuteLlamaguard7bAwqParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

// URLQuery serializes [AccountAIRunHfTheblokeExecuteLlamaguard7bAwqParams]'s query
// parameters as `url.Values`.
func (r AccountAIRunHfTheblokeExecuteLlamaguard7bAwqParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountAIRunHfTheblokeExecuteLlamaguard7bAwqParamsBody struct {
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

func (r AccountAIRunHfTheblokeExecuteLlamaguard7bAwqParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunHfTheblokeExecuteLlamaguard7bAwqParamsBody) implementsAccountAIRunHfTheblokeExecuteLlamaguard7bAwqParamsBodyUnion() {
}

// Satisfied by [AccountAIRunHfTheblokeExecuteLlamaguard7bAwqParamsBodyPrompt],
// [AccountAIRunHfTheblokeExecuteLlamaguard7bAwqParamsBodyMessages],
// [AccountAIRunHfTheblokeExecuteLlamaguard7bAwqParamsBody].
type AccountAIRunHfTheblokeExecuteLlamaguard7bAwqParamsBodyUnion interface {
	implementsAccountAIRunHfTheblokeExecuteLlamaguard7bAwqParamsBodyUnion()
}

type AccountAIRunHfTheblokeExecuteLlamaguard7bAwqParamsBodyPrompt struct {
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
	RepetitionPenalty param.Field[float64]                                                                    `json:"repetition_penalty"`
	ResponseFormat    param.Field[AccountAIRunHfTheblokeExecuteLlamaguard7bAwqParamsBodyPromptResponseFormat] `json:"response_format"`
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

func (r AccountAIRunHfTheblokeExecuteLlamaguard7bAwqParamsBodyPrompt) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunHfTheblokeExecuteLlamaguard7bAwqParamsBodyPrompt) implementsAccountAIRunHfTheblokeExecuteLlamaguard7bAwqParamsBodyUnion() {
}

type AccountAIRunHfTheblokeExecuteLlamaguard7bAwqParamsBodyPromptResponseFormat struct {
	JsonSchema param.Field[interface{}]                                                                    `json:"json_schema"`
	Type       param.Field[AccountAIRunHfTheblokeExecuteLlamaguard7bAwqParamsBodyPromptResponseFormatType] `json:"type"`
}

func (r AccountAIRunHfTheblokeExecuteLlamaguard7bAwqParamsBodyPromptResponseFormat) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunHfTheblokeExecuteLlamaguard7bAwqParamsBodyPromptResponseFormatType string

const (
	AccountAIRunHfTheblokeExecuteLlamaguard7bAwqParamsBodyPromptResponseFormatTypeJsonObject AccountAIRunHfTheblokeExecuteLlamaguard7bAwqParamsBodyPromptResponseFormatType = "json_object"
	AccountAIRunHfTheblokeExecuteLlamaguard7bAwqParamsBodyPromptResponseFormatTypeJsonSchema AccountAIRunHfTheblokeExecuteLlamaguard7bAwqParamsBodyPromptResponseFormatType = "json_schema"
)

func (r AccountAIRunHfTheblokeExecuteLlamaguard7bAwqParamsBodyPromptResponseFormatType) IsKnown() bool {
	switch r {
	case AccountAIRunHfTheblokeExecuteLlamaguard7bAwqParamsBodyPromptResponseFormatTypeJsonObject, AccountAIRunHfTheblokeExecuteLlamaguard7bAwqParamsBodyPromptResponseFormatTypeJsonSchema:
		return true
	}
	return false
}

type AccountAIRunHfTheblokeExecuteLlamaguard7bAwqParamsBodyMessages struct {
	// An array of message objects representing the conversation history.
	Messages param.Field[[]AccountAIRunHfTheblokeExecuteLlamaguard7bAwqParamsBodyMessagesMessage] `json:"messages,required"`
	// Decreases the likelihood of the model repeating the same lines verbatim.
	FrequencyPenalty param.Field[float64]                                                                  `json:"frequency_penalty"`
	Functions        param.Field[[]AccountAIRunHfTheblokeExecuteLlamaguard7bAwqParamsBodyMessagesFunction] `json:"functions"`
	// The maximum number of tokens to generate in the response.
	MaxTokens param.Field[int64] `json:"max_tokens"`
	// Increases the likelihood of the model introducing new topics.
	PresencePenalty param.Field[float64] `json:"presence_penalty"`
	// If true, a chat template is not applied and you must adhere to the specific
	// model's expected formatting.
	Raw param.Field[bool] `json:"raw"`
	// Penalty for repeated tokens; higher values discourage repetition.
	RepetitionPenalty param.Field[float64]                                                                      `json:"repetition_penalty"`
	ResponseFormat    param.Field[AccountAIRunHfTheblokeExecuteLlamaguard7bAwqParamsBodyMessagesResponseFormat] `json:"response_format"`
	// Random seed for reproducibility of the generation.
	Seed param.Field[int64] `json:"seed"`
	// If true, the response will be streamed back incrementally using SSE, Server Sent
	// Events.
	Stream param.Field[bool] `json:"stream"`
	// Controls the randomness of the output; higher values produce more random
	// results.
	Temperature param.Field[float64] `json:"temperature"`
	// A list of tools available for the assistant to use.
	Tools param.Field[[]AccountAIRunHfTheblokeExecuteLlamaguard7bAwqParamsBodyMessagesToolUnion] `json:"tools"`
	// Limits the AI to choose from the top 'k' most probable words. Lower values make
	// responses more focused; higher values introduce more variety and potential
	// surprises.
	TopK param.Field[int64] `json:"top_k"`
	// Adjusts the creativity of the AI's responses by controlling how many possible
	// words it considers. Lower values make outputs more predictable; higher values
	// allow for more varied and creative responses.
	TopP param.Field[float64] `json:"top_p"`
}

func (r AccountAIRunHfTheblokeExecuteLlamaguard7bAwqParamsBodyMessages) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunHfTheblokeExecuteLlamaguard7bAwqParamsBodyMessages) implementsAccountAIRunHfTheblokeExecuteLlamaguard7bAwqParamsBodyUnion() {
}

type AccountAIRunHfTheblokeExecuteLlamaguard7bAwqParamsBodyMessagesMessage struct {
	// The content of the message as a string.
	Content param.Field[string] `json:"content,required"`
	// The role of the message sender (e.g., 'user', 'assistant', 'system', 'tool').
	Role param.Field[string] `json:"role,required"`
}

func (r AccountAIRunHfTheblokeExecuteLlamaguard7bAwqParamsBodyMessagesMessage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunHfTheblokeExecuteLlamaguard7bAwqParamsBodyMessagesFunction struct {
	Code param.Field[string] `json:"code,required"`
	Name param.Field[string] `json:"name,required"`
}

func (r AccountAIRunHfTheblokeExecuteLlamaguard7bAwqParamsBodyMessagesFunction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunHfTheblokeExecuteLlamaguard7bAwqParamsBodyMessagesResponseFormat struct {
	JsonSchema param.Field[interface{}]                                                                      `json:"json_schema"`
	Type       param.Field[AccountAIRunHfTheblokeExecuteLlamaguard7bAwqParamsBodyMessagesResponseFormatType] `json:"type"`
}

func (r AccountAIRunHfTheblokeExecuteLlamaguard7bAwqParamsBodyMessagesResponseFormat) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunHfTheblokeExecuteLlamaguard7bAwqParamsBodyMessagesResponseFormatType string

const (
	AccountAIRunHfTheblokeExecuteLlamaguard7bAwqParamsBodyMessagesResponseFormatTypeJsonObject AccountAIRunHfTheblokeExecuteLlamaguard7bAwqParamsBodyMessagesResponseFormatType = "json_object"
	AccountAIRunHfTheblokeExecuteLlamaguard7bAwqParamsBodyMessagesResponseFormatTypeJsonSchema AccountAIRunHfTheblokeExecuteLlamaguard7bAwqParamsBodyMessagesResponseFormatType = "json_schema"
)

func (r AccountAIRunHfTheblokeExecuteLlamaguard7bAwqParamsBodyMessagesResponseFormatType) IsKnown() bool {
	switch r {
	case AccountAIRunHfTheblokeExecuteLlamaguard7bAwqParamsBodyMessagesResponseFormatTypeJsonObject, AccountAIRunHfTheblokeExecuteLlamaguard7bAwqParamsBodyMessagesResponseFormatTypeJsonSchema:
		return true
	}
	return false
}

type AccountAIRunHfTheblokeExecuteLlamaguard7bAwqParamsBodyMessagesTool struct {
	// A brief description of what the tool does.
	Description param.Field[string]      `json:"description"`
	Function    param.Field[interface{}] `json:"function"`
	// The name of the tool. More descriptive the better.
	Name       param.Field[string]      `json:"name"`
	Parameters param.Field[interface{}] `json:"parameters"`
	// Specifies the type of tool (e.g., 'function').
	Type param.Field[string] `json:"type"`
}

func (r AccountAIRunHfTheblokeExecuteLlamaguard7bAwqParamsBodyMessagesTool) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunHfTheblokeExecuteLlamaguard7bAwqParamsBodyMessagesTool) implementsAccountAIRunHfTheblokeExecuteLlamaguard7bAwqParamsBodyMessagesToolUnion() {
}

// Satisfied by
// [AccountAIRunHfTheblokeExecuteLlamaguard7bAwqParamsBodyMessagesToolsObject],
// [AccountAIRunHfTheblokeExecuteLlamaguard7bAwqParamsBodyMessagesToolsObject],
// [AccountAIRunHfTheblokeExecuteLlamaguard7bAwqParamsBodyMessagesTool].
type AccountAIRunHfTheblokeExecuteLlamaguard7bAwqParamsBodyMessagesToolUnion interface {
	implementsAccountAIRunHfTheblokeExecuteLlamaguard7bAwqParamsBodyMessagesToolUnion()
}

type AccountAIRunHfTheblokeExecuteLlamaguard7bAwqParamsBodyMessagesToolsObject struct {
	// A brief description of what the tool does.
	Description param.Field[string] `json:"description,required"`
	// The name of the tool. More descriptive the better.
	Name param.Field[string] `json:"name,required"`
	// Schema defining the parameters accepted by the tool.
	Parameters param.Field[AccountAIRunHfTheblokeExecuteLlamaguard7bAwqParamsBodyMessagesToolsObjectParameters] `json:"parameters,required"`
}

func (r AccountAIRunHfTheblokeExecuteLlamaguard7bAwqParamsBodyMessagesToolsObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunHfTheblokeExecuteLlamaguard7bAwqParamsBodyMessagesToolsObject) implementsAccountAIRunHfTheblokeExecuteLlamaguard7bAwqParamsBodyMessagesToolUnion() {
}

// Schema defining the parameters accepted by the tool.
type AccountAIRunHfTheblokeExecuteLlamaguard7bAwqParamsBodyMessagesToolsObjectParameters struct {
	// Definitions of each parameter.
	Properties param.Field[map[string]AccountAIRunHfTheblokeExecuteLlamaguard7bAwqParamsBodyMessagesToolsObjectParametersProperties] `json:"properties,required"`
	// The type of the parameters object (usually 'object').
	Type param.Field[string] `json:"type,required"`
	// List of required parameter names.
	Required param.Field[[]string] `json:"required"`
}

func (r AccountAIRunHfTheblokeExecuteLlamaguard7bAwqParamsBodyMessagesToolsObjectParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunHfTheblokeExecuteLlamaguard7bAwqParamsBodyMessagesToolsObjectParametersProperties struct {
	// A description of the expected parameter.
	Description param.Field[string] `json:"description,required"`
	// The data type of the parameter.
	Type param.Field[string] `json:"type,required"`
}

func (r AccountAIRunHfTheblokeExecuteLlamaguard7bAwqParamsBodyMessagesToolsObjectParametersProperties) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunHfTheblokeExecuteMistral7bInstructV0_1AwqParams struct {
	QueueRequest param.Field[string]                                                  `query:"queueRequest"`
	Body         AccountAIRunHfTheblokeExecuteMistral7bInstructV0_1AwqParamsBodyUnion `json:"body"`
}

func (r AccountAIRunHfTheblokeExecuteMistral7bInstructV0_1AwqParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

// URLQuery serializes
// [AccountAIRunHfTheblokeExecuteMistral7bInstructV0_1AwqParams]'s query parameters
// as `url.Values`.
func (r AccountAIRunHfTheblokeExecuteMistral7bInstructV0_1AwqParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountAIRunHfTheblokeExecuteMistral7bInstructV0_1AwqParamsBody struct {
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

func (r AccountAIRunHfTheblokeExecuteMistral7bInstructV0_1AwqParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunHfTheblokeExecuteMistral7bInstructV0_1AwqParamsBody) implementsAccountAIRunHfTheblokeExecuteMistral7bInstructV0_1AwqParamsBodyUnion() {
}

// Satisfied by
// [AccountAIRunHfTheblokeExecuteMistral7bInstructV0_1AwqParamsBodyPrompt],
// [AccountAIRunHfTheblokeExecuteMistral7bInstructV0_1AwqParamsBodyMessages],
// [AccountAIRunHfTheblokeExecuteMistral7bInstructV0_1AwqParamsBody].
type AccountAIRunHfTheblokeExecuteMistral7bInstructV0_1AwqParamsBodyUnion interface {
	implementsAccountAIRunHfTheblokeExecuteMistral7bInstructV0_1AwqParamsBodyUnion()
}

type AccountAIRunHfTheblokeExecuteMistral7bInstructV0_1AwqParamsBodyPrompt struct {
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
	ResponseFormat    param.Field[AccountAIRunHfTheblokeExecuteMistral7bInstructV0_1AwqParamsBodyPromptResponseFormat] `json:"response_format"`
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

func (r AccountAIRunHfTheblokeExecuteMistral7bInstructV0_1AwqParamsBodyPrompt) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunHfTheblokeExecuteMistral7bInstructV0_1AwqParamsBodyPrompt) implementsAccountAIRunHfTheblokeExecuteMistral7bInstructV0_1AwqParamsBodyUnion() {
}

type AccountAIRunHfTheblokeExecuteMistral7bInstructV0_1AwqParamsBodyPromptResponseFormat struct {
	JsonSchema param.Field[interface{}]                                                                             `json:"json_schema"`
	Type       param.Field[AccountAIRunHfTheblokeExecuteMistral7bInstructV0_1AwqParamsBodyPromptResponseFormatType] `json:"type"`
}

func (r AccountAIRunHfTheblokeExecuteMistral7bInstructV0_1AwqParamsBodyPromptResponseFormat) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunHfTheblokeExecuteMistral7bInstructV0_1AwqParamsBodyPromptResponseFormatType string

const (
	AccountAIRunHfTheblokeExecuteMistral7bInstructV0_1AwqParamsBodyPromptResponseFormatTypeJsonObject AccountAIRunHfTheblokeExecuteMistral7bInstructV0_1AwqParamsBodyPromptResponseFormatType = "json_object"
	AccountAIRunHfTheblokeExecuteMistral7bInstructV0_1AwqParamsBodyPromptResponseFormatTypeJsonSchema AccountAIRunHfTheblokeExecuteMistral7bInstructV0_1AwqParamsBodyPromptResponseFormatType = "json_schema"
)

func (r AccountAIRunHfTheblokeExecuteMistral7bInstructV0_1AwqParamsBodyPromptResponseFormatType) IsKnown() bool {
	switch r {
	case AccountAIRunHfTheblokeExecuteMistral7bInstructV0_1AwqParamsBodyPromptResponseFormatTypeJsonObject, AccountAIRunHfTheblokeExecuteMistral7bInstructV0_1AwqParamsBodyPromptResponseFormatTypeJsonSchema:
		return true
	}
	return false
}

type AccountAIRunHfTheblokeExecuteMistral7bInstructV0_1AwqParamsBodyMessages struct {
	// An array of message objects representing the conversation history.
	Messages param.Field[[]AccountAIRunHfTheblokeExecuteMistral7bInstructV0_1AwqParamsBodyMessagesMessage] `json:"messages,required"`
	// Decreases the likelihood of the model repeating the same lines verbatim.
	FrequencyPenalty param.Field[float64]                                                                           `json:"frequency_penalty"`
	Functions        param.Field[[]AccountAIRunHfTheblokeExecuteMistral7bInstructV0_1AwqParamsBodyMessagesFunction] `json:"functions"`
	// The maximum number of tokens to generate in the response.
	MaxTokens param.Field[int64] `json:"max_tokens"`
	// Increases the likelihood of the model introducing new topics.
	PresencePenalty param.Field[float64] `json:"presence_penalty"`
	// If true, a chat template is not applied and you must adhere to the specific
	// model's expected formatting.
	Raw param.Field[bool] `json:"raw"`
	// Penalty for repeated tokens; higher values discourage repetition.
	RepetitionPenalty param.Field[float64]                                                                               `json:"repetition_penalty"`
	ResponseFormat    param.Field[AccountAIRunHfTheblokeExecuteMistral7bInstructV0_1AwqParamsBodyMessagesResponseFormat] `json:"response_format"`
	// Random seed for reproducibility of the generation.
	Seed param.Field[int64] `json:"seed"`
	// If true, the response will be streamed back incrementally using SSE, Server Sent
	// Events.
	Stream param.Field[bool] `json:"stream"`
	// Controls the randomness of the output; higher values produce more random
	// results.
	Temperature param.Field[float64] `json:"temperature"`
	// A list of tools available for the assistant to use.
	Tools param.Field[[]AccountAIRunHfTheblokeExecuteMistral7bInstructV0_1AwqParamsBodyMessagesToolUnion] `json:"tools"`
	// Limits the AI to choose from the top 'k' most probable words. Lower values make
	// responses more focused; higher values introduce more variety and potential
	// surprises.
	TopK param.Field[int64] `json:"top_k"`
	// Adjusts the creativity of the AI's responses by controlling how many possible
	// words it considers. Lower values make outputs more predictable; higher values
	// allow for more varied and creative responses.
	TopP param.Field[float64] `json:"top_p"`
}

func (r AccountAIRunHfTheblokeExecuteMistral7bInstructV0_1AwqParamsBodyMessages) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunHfTheblokeExecuteMistral7bInstructV0_1AwqParamsBodyMessages) implementsAccountAIRunHfTheblokeExecuteMistral7bInstructV0_1AwqParamsBodyUnion() {
}

type AccountAIRunHfTheblokeExecuteMistral7bInstructV0_1AwqParamsBodyMessagesMessage struct {
	// The content of the message as a string.
	Content param.Field[string] `json:"content,required"`
	// The role of the message sender (e.g., 'user', 'assistant', 'system', 'tool').
	Role param.Field[string] `json:"role,required"`
}

func (r AccountAIRunHfTheblokeExecuteMistral7bInstructV0_1AwqParamsBodyMessagesMessage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunHfTheblokeExecuteMistral7bInstructV0_1AwqParamsBodyMessagesFunction struct {
	Code param.Field[string] `json:"code,required"`
	Name param.Field[string] `json:"name,required"`
}

func (r AccountAIRunHfTheblokeExecuteMistral7bInstructV0_1AwqParamsBodyMessagesFunction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunHfTheblokeExecuteMistral7bInstructV0_1AwqParamsBodyMessagesResponseFormat struct {
	JsonSchema param.Field[interface{}]                                                                               `json:"json_schema"`
	Type       param.Field[AccountAIRunHfTheblokeExecuteMistral7bInstructV0_1AwqParamsBodyMessagesResponseFormatType] `json:"type"`
}

func (r AccountAIRunHfTheblokeExecuteMistral7bInstructV0_1AwqParamsBodyMessagesResponseFormat) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunHfTheblokeExecuteMistral7bInstructV0_1AwqParamsBodyMessagesResponseFormatType string

const (
	AccountAIRunHfTheblokeExecuteMistral7bInstructV0_1AwqParamsBodyMessagesResponseFormatTypeJsonObject AccountAIRunHfTheblokeExecuteMistral7bInstructV0_1AwqParamsBodyMessagesResponseFormatType = "json_object"
	AccountAIRunHfTheblokeExecuteMistral7bInstructV0_1AwqParamsBodyMessagesResponseFormatTypeJsonSchema AccountAIRunHfTheblokeExecuteMistral7bInstructV0_1AwqParamsBodyMessagesResponseFormatType = "json_schema"
)

func (r AccountAIRunHfTheblokeExecuteMistral7bInstructV0_1AwqParamsBodyMessagesResponseFormatType) IsKnown() bool {
	switch r {
	case AccountAIRunHfTheblokeExecuteMistral7bInstructV0_1AwqParamsBodyMessagesResponseFormatTypeJsonObject, AccountAIRunHfTheblokeExecuteMistral7bInstructV0_1AwqParamsBodyMessagesResponseFormatTypeJsonSchema:
		return true
	}
	return false
}

type AccountAIRunHfTheblokeExecuteMistral7bInstructV0_1AwqParamsBodyMessagesTool struct {
	// A brief description of what the tool does.
	Description param.Field[string]      `json:"description"`
	Function    param.Field[interface{}] `json:"function"`
	// The name of the tool. More descriptive the better.
	Name       param.Field[string]      `json:"name"`
	Parameters param.Field[interface{}] `json:"parameters"`
	// Specifies the type of tool (e.g., 'function').
	Type param.Field[string] `json:"type"`
}

func (r AccountAIRunHfTheblokeExecuteMistral7bInstructV0_1AwqParamsBodyMessagesTool) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunHfTheblokeExecuteMistral7bInstructV0_1AwqParamsBodyMessagesTool) implementsAccountAIRunHfTheblokeExecuteMistral7bInstructV0_1AwqParamsBodyMessagesToolUnion() {
}

// Satisfied by
// [AccountAIRunHfTheblokeExecuteMistral7bInstructV0_1AwqParamsBodyMessagesToolsObject],
// [AccountAIRunHfTheblokeExecuteMistral7bInstructV0_1AwqParamsBodyMessagesToolsObject],
// [AccountAIRunHfTheblokeExecuteMistral7bInstructV0_1AwqParamsBodyMessagesTool].
type AccountAIRunHfTheblokeExecuteMistral7bInstructV0_1AwqParamsBodyMessagesToolUnion interface {
	implementsAccountAIRunHfTheblokeExecuteMistral7bInstructV0_1AwqParamsBodyMessagesToolUnion()
}

type AccountAIRunHfTheblokeExecuteMistral7bInstructV0_1AwqParamsBodyMessagesToolsObject struct {
	// A brief description of what the tool does.
	Description param.Field[string] `json:"description,required"`
	// The name of the tool. More descriptive the better.
	Name param.Field[string] `json:"name,required"`
	// Schema defining the parameters accepted by the tool.
	Parameters param.Field[AccountAIRunHfTheblokeExecuteMistral7bInstructV0_1AwqParamsBodyMessagesToolsObjectParameters] `json:"parameters,required"`
}

func (r AccountAIRunHfTheblokeExecuteMistral7bInstructV0_1AwqParamsBodyMessagesToolsObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunHfTheblokeExecuteMistral7bInstructV0_1AwqParamsBodyMessagesToolsObject) implementsAccountAIRunHfTheblokeExecuteMistral7bInstructV0_1AwqParamsBodyMessagesToolUnion() {
}

// Schema defining the parameters accepted by the tool.
type AccountAIRunHfTheblokeExecuteMistral7bInstructV0_1AwqParamsBodyMessagesToolsObjectParameters struct {
	// Definitions of each parameter.
	Properties param.Field[map[string]AccountAIRunHfTheblokeExecuteMistral7bInstructV0_1AwqParamsBodyMessagesToolsObjectParametersProperties] `json:"properties,required"`
	// The type of the parameters object (usually 'object').
	Type param.Field[string] `json:"type,required"`
	// List of required parameter names.
	Required param.Field[[]string] `json:"required"`
}

func (r AccountAIRunHfTheblokeExecuteMistral7bInstructV0_1AwqParamsBodyMessagesToolsObjectParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunHfTheblokeExecuteMistral7bInstructV0_1AwqParamsBodyMessagesToolsObjectParametersProperties struct {
	// A description of the expected parameter.
	Description param.Field[string] `json:"description,required"`
	// The data type of the parameter.
	Type param.Field[string] `json:"type,required"`
}

func (r AccountAIRunHfTheblokeExecuteMistral7bInstructV0_1AwqParamsBodyMessagesToolsObjectParametersProperties) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunHfTheblokeExecuteNeuralChat7bV3_1AwqParams struct {
	QueueRequest param.Field[string]                                             `query:"queueRequest"`
	Body         AccountAIRunHfTheblokeExecuteNeuralChat7bV3_1AwqParamsBodyUnion `json:"body"`
}

func (r AccountAIRunHfTheblokeExecuteNeuralChat7bV3_1AwqParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

// URLQuery serializes [AccountAIRunHfTheblokeExecuteNeuralChat7bV3_1AwqParams]'s
// query parameters as `url.Values`.
func (r AccountAIRunHfTheblokeExecuteNeuralChat7bV3_1AwqParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountAIRunHfTheblokeExecuteNeuralChat7bV3_1AwqParamsBody struct {
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

func (r AccountAIRunHfTheblokeExecuteNeuralChat7bV3_1AwqParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunHfTheblokeExecuteNeuralChat7bV3_1AwqParamsBody) implementsAccountAIRunHfTheblokeExecuteNeuralChat7bV3_1AwqParamsBodyUnion() {
}

// Satisfied by [AccountAIRunHfTheblokeExecuteNeuralChat7bV3_1AwqParamsBodyPrompt],
// [AccountAIRunHfTheblokeExecuteNeuralChat7bV3_1AwqParamsBodyMessages],
// [AccountAIRunHfTheblokeExecuteNeuralChat7bV3_1AwqParamsBody].
type AccountAIRunHfTheblokeExecuteNeuralChat7bV3_1AwqParamsBodyUnion interface {
	implementsAccountAIRunHfTheblokeExecuteNeuralChat7bV3_1AwqParamsBodyUnion()
}

type AccountAIRunHfTheblokeExecuteNeuralChat7bV3_1AwqParamsBodyPrompt struct {
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
	RepetitionPenalty param.Field[float64]                                                                        `json:"repetition_penalty"`
	ResponseFormat    param.Field[AccountAIRunHfTheblokeExecuteNeuralChat7bV3_1AwqParamsBodyPromptResponseFormat] `json:"response_format"`
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

func (r AccountAIRunHfTheblokeExecuteNeuralChat7bV3_1AwqParamsBodyPrompt) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunHfTheblokeExecuteNeuralChat7bV3_1AwqParamsBodyPrompt) implementsAccountAIRunHfTheblokeExecuteNeuralChat7bV3_1AwqParamsBodyUnion() {
}

type AccountAIRunHfTheblokeExecuteNeuralChat7bV3_1AwqParamsBodyPromptResponseFormat struct {
	JsonSchema param.Field[interface{}]                                                                        `json:"json_schema"`
	Type       param.Field[AccountAIRunHfTheblokeExecuteNeuralChat7bV3_1AwqParamsBodyPromptResponseFormatType] `json:"type"`
}

func (r AccountAIRunHfTheblokeExecuteNeuralChat7bV3_1AwqParamsBodyPromptResponseFormat) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunHfTheblokeExecuteNeuralChat7bV3_1AwqParamsBodyPromptResponseFormatType string

const (
	AccountAIRunHfTheblokeExecuteNeuralChat7bV3_1AwqParamsBodyPromptResponseFormatTypeJsonObject AccountAIRunHfTheblokeExecuteNeuralChat7bV3_1AwqParamsBodyPromptResponseFormatType = "json_object"
	AccountAIRunHfTheblokeExecuteNeuralChat7bV3_1AwqParamsBodyPromptResponseFormatTypeJsonSchema AccountAIRunHfTheblokeExecuteNeuralChat7bV3_1AwqParamsBodyPromptResponseFormatType = "json_schema"
)

func (r AccountAIRunHfTheblokeExecuteNeuralChat7bV3_1AwqParamsBodyPromptResponseFormatType) IsKnown() bool {
	switch r {
	case AccountAIRunHfTheblokeExecuteNeuralChat7bV3_1AwqParamsBodyPromptResponseFormatTypeJsonObject, AccountAIRunHfTheblokeExecuteNeuralChat7bV3_1AwqParamsBodyPromptResponseFormatTypeJsonSchema:
		return true
	}
	return false
}

type AccountAIRunHfTheblokeExecuteNeuralChat7bV3_1AwqParamsBodyMessages struct {
	// An array of message objects representing the conversation history.
	Messages param.Field[[]AccountAIRunHfTheblokeExecuteNeuralChat7bV3_1AwqParamsBodyMessagesMessage] `json:"messages,required"`
	// Decreases the likelihood of the model repeating the same lines verbatim.
	FrequencyPenalty param.Field[float64]                                                                      `json:"frequency_penalty"`
	Functions        param.Field[[]AccountAIRunHfTheblokeExecuteNeuralChat7bV3_1AwqParamsBodyMessagesFunction] `json:"functions"`
	// The maximum number of tokens to generate in the response.
	MaxTokens param.Field[int64] `json:"max_tokens"`
	// Increases the likelihood of the model introducing new topics.
	PresencePenalty param.Field[float64] `json:"presence_penalty"`
	// If true, a chat template is not applied and you must adhere to the specific
	// model's expected formatting.
	Raw param.Field[bool] `json:"raw"`
	// Penalty for repeated tokens; higher values discourage repetition.
	RepetitionPenalty param.Field[float64]                                                                          `json:"repetition_penalty"`
	ResponseFormat    param.Field[AccountAIRunHfTheblokeExecuteNeuralChat7bV3_1AwqParamsBodyMessagesResponseFormat] `json:"response_format"`
	// Random seed for reproducibility of the generation.
	Seed param.Field[int64] `json:"seed"`
	// If true, the response will be streamed back incrementally using SSE, Server Sent
	// Events.
	Stream param.Field[bool] `json:"stream"`
	// Controls the randomness of the output; higher values produce more random
	// results.
	Temperature param.Field[float64] `json:"temperature"`
	// A list of tools available for the assistant to use.
	Tools param.Field[[]AccountAIRunHfTheblokeExecuteNeuralChat7bV3_1AwqParamsBodyMessagesToolUnion] `json:"tools"`
	// Limits the AI to choose from the top 'k' most probable words. Lower values make
	// responses more focused; higher values introduce more variety and potential
	// surprises.
	TopK param.Field[int64] `json:"top_k"`
	// Adjusts the creativity of the AI's responses by controlling how many possible
	// words it considers. Lower values make outputs more predictable; higher values
	// allow for more varied and creative responses.
	TopP param.Field[float64] `json:"top_p"`
}

func (r AccountAIRunHfTheblokeExecuteNeuralChat7bV3_1AwqParamsBodyMessages) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunHfTheblokeExecuteNeuralChat7bV3_1AwqParamsBodyMessages) implementsAccountAIRunHfTheblokeExecuteNeuralChat7bV3_1AwqParamsBodyUnion() {
}

type AccountAIRunHfTheblokeExecuteNeuralChat7bV3_1AwqParamsBodyMessagesMessage struct {
	// The content of the message as a string.
	Content param.Field[string] `json:"content,required"`
	// The role of the message sender (e.g., 'user', 'assistant', 'system', 'tool').
	Role param.Field[string] `json:"role,required"`
}

func (r AccountAIRunHfTheblokeExecuteNeuralChat7bV3_1AwqParamsBodyMessagesMessage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunHfTheblokeExecuteNeuralChat7bV3_1AwqParamsBodyMessagesFunction struct {
	Code param.Field[string] `json:"code,required"`
	Name param.Field[string] `json:"name,required"`
}

func (r AccountAIRunHfTheblokeExecuteNeuralChat7bV3_1AwqParamsBodyMessagesFunction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunHfTheblokeExecuteNeuralChat7bV3_1AwqParamsBodyMessagesResponseFormat struct {
	JsonSchema param.Field[interface{}]                                                                          `json:"json_schema"`
	Type       param.Field[AccountAIRunHfTheblokeExecuteNeuralChat7bV3_1AwqParamsBodyMessagesResponseFormatType] `json:"type"`
}

func (r AccountAIRunHfTheblokeExecuteNeuralChat7bV3_1AwqParamsBodyMessagesResponseFormat) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunHfTheblokeExecuteNeuralChat7bV3_1AwqParamsBodyMessagesResponseFormatType string

const (
	AccountAIRunHfTheblokeExecuteNeuralChat7bV3_1AwqParamsBodyMessagesResponseFormatTypeJsonObject AccountAIRunHfTheblokeExecuteNeuralChat7bV3_1AwqParamsBodyMessagesResponseFormatType = "json_object"
	AccountAIRunHfTheblokeExecuteNeuralChat7bV3_1AwqParamsBodyMessagesResponseFormatTypeJsonSchema AccountAIRunHfTheblokeExecuteNeuralChat7bV3_1AwqParamsBodyMessagesResponseFormatType = "json_schema"
)

func (r AccountAIRunHfTheblokeExecuteNeuralChat7bV3_1AwqParamsBodyMessagesResponseFormatType) IsKnown() bool {
	switch r {
	case AccountAIRunHfTheblokeExecuteNeuralChat7bV3_1AwqParamsBodyMessagesResponseFormatTypeJsonObject, AccountAIRunHfTheblokeExecuteNeuralChat7bV3_1AwqParamsBodyMessagesResponseFormatTypeJsonSchema:
		return true
	}
	return false
}

type AccountAIRunHfTheblokeExecuteNeuralChat7bV3_1AwqParamsBodyMessagesTool struct {
	// A brief description of what the tool does.
	Description param.Field[string]      `json:"description"`
	Function    param.Field[interface{}] `json:"function"`
	// The name of the tool. More descriptive the better.
	Name       param.Field[string]      `json:"name"`
	Parameters param.Field[interface{}] `json:"parameters"`
	// Specifies the type of tool (e.g., 'function').
	Type param.Field[string] `json:"type"`
}

func (r AccountAIRunHfTheblokeExecuteNeuralChat7bV3_1AwqParamsBodyMessagesTool) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunHfTheblokeExecuteNeuralChat7bV3_1AwqParamsBodyMessagesTool) implementsAccountAIRunHfTheblokeExecuteNeuralChat7bV3_1AwqParamsBodyMessagesToolUnion() {
}

// Satisfied by
// [AccountAIRunHfTheblokeExecuteNeuralChat7bV3_1AwqParamsBodyMessagesToolsObject],
// [AccountAIRunHfTheblokeExecuteNeuralChat7bV3_1AwqParamsBodyMessagesToolsObject],
// [AccountAIRunHfTheblokeExecuteNeuralChat7bV3_1AwqParamsBodyMessagesTool].
type AccountAIRunHfTheblokeExecuteNeuralChat7bV3_1AwqParamsBodyMessagesToolUnion interface {
	implementsAccountAIRunHfTheblokeExecuteNeuralChat7bV3_1AwqParamsBodyMessagesToolUnion()
}

type AccountAIRunHfTheblokeExecuteNeuralChat7bV3_1AwqParamsBodyMessagesToolsObject struct {
	// A brief description of what the tool does.
	Description param.Field[string] `json:"description,required"`
	// The name of the tool. More descriptive the better.
	Name param.Field[string] `json:"name,required"`
	// Schema defining the parameters accepted by the tool.
	Parameters param.Field[AccountAIRunHfTheblokeExecuteNeuralChat7bV3_1AwqParamsBodyMessagesToolsObjectParameters] `json:"parameters,required"`
}

func (r AccountAIRunHfTheblokeExecuteNeuralChat7bV3_1AwqParamsBodyMessagesToolsObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunHfTheblokeExecuteNeuralChat7bV3_1AwqParamsBodyMessagesToolsObject) implementsAccountAIRunHfTheblokeExecuteNeuralChat7bV3_1AwqParamsBodyMessagesToolUnion() {
}

// Schema defining the parameters accepted by the tool.
type AccountAIRunHfTheblokeExecuteNeuralChat7bV3_1AwqParamsBodyMessagesToolsObjectParameters struct {
	// Definitions of each parameter.
	Properties param.Field[map[string]AccountAIRunHfTheblokeExecuteNeuralChat7bV3_1AwqParamsBodyMessagesToolsObjectParametersProperties] `json:"properties,required"`
	// The type of the parameters object (usually 'object').
	Type param.Field[string] `json:"type,required"`
	// List of required parameter names.
	Required param.Field[[]string] `json:"required"`
}

func (r AccountAIRunHfTheblokeExecuteNeuralChat7bV3_1AwqParamsBodyMessagesToolsObjectParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunHfTheblokeExecuteNeuralChat7bV3_1AwqParamsBodyMessagesToolsObjectParametersProperties struct {
	// A description of the expected parameter.
	Description param.Field[string] `json:"description,required"`
	// The data type of the parameter.
	Type param.Field[string] `json:"type,required"`
}

func (r AccountAIRunHfTheblokeExecuteNeuralChat7bV3_1AwqParamsBodyMessagesToolsObjectParametersProperties) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunHfTheblokeExecuteOpenhermes2_5Mistral7bAwqParams struct {
	QueueRequest param.Field[string]                                                   `query:"queueRequest"`
	Body         AccountAIRunHfTheblokeExecuteOpenhermes2_5Mistral7bAwqParamsBodyUnion `json:"body"`
}

func (r AccountAIRunHfTheblokeExecuteOpenhermes2_5Mistral7bAwqParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

// URLQuery serializes
// [AccountAIRunHfTheblokeExecuteOpenhermes2_5Mistral7bAwqParams]'s query
// parameters as `url.Values`.
func (r AccountAIRunHfTheblokeExecuteOpenhermes2_5Mistral7bAwqParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountAIRunHfTheblokeExecuteOpenhermes2_5Mistral7bAwqParamsBody struct {
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

func (r AccountAIRunHfTheblokeExecuteOpenhermes2_5Mistral7bAwqParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunHfTheblokeExecuteOpenhermes2_5Mistral7bAwqParamsBody) implementsAccountAIRunHfTheblokeExecuteOpenhermes2_5Mistral7bAwqParamsBodyUnion() {
}

// Satisfied by
// [AccountAIRunHfTheblokeExecuteOpenhermes2_5Mistral7bAwqParamsBodyPrompt],
// [AccountAIRunHfTheblokeExecuteOpenhermes2_5Mistral7bAwqParamsBodyMessages],
// [AccountAIRunHfTheblokeExecuteOpenhermes2_5Mistral7bAwqParamsBody].
type AccountAIRunHfTheblokeExecuteOpenhermes2_5Mistral7bAwqParamsBodyUnion interface {
	implementsAccountAIRunHfTheblokeExecuteOpenhermes2_5Mistral7bAwqParamsBodyUnion()
}

type AccountAIRunHfTheblokeExecuteOpenhermes2_5Mistral7bAwqParamsBodyPrompt struct {
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
	RepetitionPenalty param.Field[float64]                                                                              `json:"repetition_penalty"`
	ResponseFormat    param.Field[AccountAIRunHfTheblokeExecuteOpenhermes2_5Mistral7bAwqParamsBodyPromptResponseFormat] `json:"response_format"`
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

func (r AccountAIRunHfTheblokeExecuteOpenhermes2_5Mistral7bAwqParamsBodyPrompt) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunHfTheblokeExecuteOpenhermes2_5Mistral7bAwqParamsBodyPrompt) implementsAccountAIRunHfTheblokeExecuteOpenhermes2_5Mistral7bAwqParamsBodyUnion() {
}

type AccountAIRunHfTheblokeExecuteOpenhermes2_5Mistral7bAwqParamsBodyPromptResponseFormat struct {
	JsonSchema param.Field[interface{}]                                                                              `json:"json_schema"`
	Type       param.Field[AccountAIRunHfTheblokeExecuteOpenhermes2_5Mistral7bAwqParamsBodyPromptResponseFormatType] `json:"type"`
}

func (r AccountAIRunHfTheblokeExecuteOpenhermes2_5Mistral7bAwqParamsBodyPromptResponseFormat) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunHfTheblokeExecuteOpenhermes2_5Mistral7bAwqParamsBodyPromptResponseFormatType string

const (
	AccountAIRunHfTheblokeExecuteOpenhermes2_5Mistral7bAwqParamsBodyPromptResponseFormatTypeJsonObject AccountAIRunHfTheblokeExecuteOpenhermes2_5Mistral7bAwqParamsBodyPromptResponseFormatType = "json_object"
	AccountAIRunHfTheblokeExecuteOpenhermes2_5Mistral7bAwqParamsBodyPromptResponseFormatTypeJsonSchema AccountAIRunHfTheblokeExecuteOpenhermes2_5Mistral7bAwqParamsBodyPromptResponseFormatType = "json_schema"
)

func (r AccountAIRunHfTheblokeExecuteOpenhermes2_5Mistral7bAwqParamsBodyPromptResponseFormatType) IsKnown() bool {
	switch r {
	case AccountAIRunHfTheblokeExecuteOpenhermes2_5Mistral7bAwqParamsBodyPromptResponseFormatTypeJsonObject, AccountAIRunHfTheblokeExecuteOpenhermes2_5Mistral7bAwqParamsBodyPromptResponseFormatTypeJsonSchema:
		return true
	}
	return false
}

type AccountAIRunHfTheblokeExecuteOpenhermes2_5Mistral7bAwqParamsBodyMessages struct {
	// An array of message objects representing the conversation history.
	Messages param.Field[[]AccountAIRunHfTheblokeExecuteOpenhermes2_5Mistral7bAwqParamsBodyMessagesMessage] `json:"messages,required"`
	// Decreases the likelihood of the model repeating the same lines verbatim.
	FrequencyPenalty param.Field[float64]                                                                            `json:"frequency_penalty"`
	Functions        param.Field[[]AccountAIRunHfTheblokeExecuteOpenhermes2_5Mistral7bAwqParamsBodyMessagesFunction] `json:"functions"`
	// The maximum number of tokens to generate in the response.
	MaxTokens param.Field[int64] `json:"max_tokens"`
	// Increases the likelihood of the model introducing new topics.
	PresencePenalty param.Field[float64] `json:"presence_penalty"`
	// If true, a chat template is not applied and you must adhere to the specific
	// model's expected formatting.
	Raw param.Field[bool] `json:"raw"`
	// Penalty for repeated tokens; higher values discourage repetition.
	RepetitionPenalty param.Field[float64]                                                                                `json:"repetition_penalty"`
	ResponseFormat    param.Field[AccountAIRunHfTheblokeExecuteOpenhermes2_5Mistral7bAwqParamsBodyMessagesResponseFormat] `json:"response_format"`
	// Random seed for reproducibility of the generation.
	Seed param.Field[int64] `json:"seed"`
	// If true, the response will be streamed back incrementally using SSE, Server Sent
	// Events.
	Stream param.Field[bool] `json:"stream"`
	// Controls the randomness of the output; higher values produce more random
	// results.
	Temperature param.Field[float64] `json:"temperature"`
	// A list of tools available for the assistant to use.
	Tools param.Field[[]AccountAIRunHfTheblokeExecuteOpenhermes2_5Mistral7bAwqParamsBodyMessagesToolUnion] `json:"tools"`
	// Limits the AI to choose from the top 'k' most probable words. Lower values make
	// responses more focused; higher values introduce more variety and potential
	// surprises.
	TopK param.Field[int64] `json:"top_k"`
	// Adjusts the creativity of the AI's responses by controlling how many possible
	// words it considers. Lower values make outputs more predictable; higher values
	// allow for more varied and creative responses.
	TopP param.Field[float64] `json:"top_p"`
}

func (r AccountAIRunHfTheblokeExecuteOpenhermes2_5Mistral7bAwqParamsBodyMessages) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunHfTheblokeExecuteOpenhermes2_5Mistral7bAwqParamsBodyMessages) implementsAccountAIRunHfTheblokeExecuteOpenhermes2_5Mistral7bAwqParamsBodyUnion() {
}

type AccountAIRunHfTheblokeExecuteOpenhermes2_5Mistral7bAwqParamsBodyMessagesMessage struct {
	// The content of the message as a string.
	Content param.Field[string] `json:"content,required"`
	// The role of the message sender (e.g., 'user', 'assistant', 'system', 'tool').
	Role param.Field[string] `json:"role,required"`
}

func (r AccountAIRunHfTheblokeExecuteOpenhermes2_5Mistral7bAwqParamsBodyMessagesMessage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunHfTheblokeExecuteOpenhermes2_5Mistral7bAwqParamsBodyMessagesFunction struct {
	Code param.Field[string] `json:"code,required"`
	Name param.Field[string] `json:"name,required"`
}

func (r AccountAIRunHfTheblokeExecuteOpenhermes2_5Mistral7bAwqParamsBodyMessagesFunction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunHfTheblokeExecuteOpenhermes2_5Mistral7bAwqParamsBodyMessagesResponseFormat struct {
	JsonSchema param.Field[interface{}]                                                                                `json:"json_schema"`
	Type       param.Field[AccountAIRunHfTheblokeExecuteOpenhermes2_5Mistral7bAwqParamsBodyMessagesResponseFormatType] `json:"type"`
}

func (r AccountAIRunHfTheblokeExecuteOpenhermes2_5Mistral7bAwqParamsBodyMessagesResponseFormat) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunHfTheblokeExecuteOpenhermes2_5Mistral7bAwqParamsBodyMessagesResponseFormatType string

const (
	AccountAIRunHfTheblokeExecuteOpenhermes2_5Mistral7bAwqParamsBodyMessagesResponseFormatTypeJsonObject AccountAIRunHfTheblokeExecuteOpenhermes2_5Mistral7bAwqParamsBodyMessagesResponseFormatType = "json_object"
	AccountAIRunHfTheblokeExecuteOpenhermes2_5Mistral7bAwqParamsBodyMessagesResponseFormatTypeJsonSchema AccountAIRunHfTheblokeExecuteOpenhermes2_5Mistral7bAwqParamsBodyMessagesResponseFormatType = "json_schema"
)

func (r AccountAIRunHfTheblokeExecuteOpenhermes2_5Mistral7bAwqParamsBodyMessagesResponseFormatType) IsKnown() bool {
	switch r {
	case AccountAIRunHfTheblokeExecuteOpenhermes2_5Mistral7bAwqParamsBodyMessagesResponseFormatTypeJsonObject, AccountAIRunHfTheblokeExecuteOpenhermes2_5Mistral7bAwqParamsBodyMessagesResponseFormatTypeJsonSchema:
		return true
	}
	return false
}

type AccountAIRunHfTheblokeExecuteOpenhermes2_5Mistral7bAwqParamsBodyMessagesTool struct {
	// A brief description of what the tool does.
	Description param.Field[string]      `json:"description"`
	Function    param.Field[interface{}] `json:"function"`
	// The name of the tool. More descriptive the better.
	Name       param.Field[string]      `json:"name"`
	Parameters param.Field[interface{}] `json:"parameters"`
	// Specifies the type of tool (e.g., 'function').
	Type param.Field[string] `json:"type"`
}

func (r AccountAIRunHfTheblokeExecuteOpenhermes2_5Mistral7bAwqParamsBodyMessagesTool) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunHfTheblokeExecuteOpenhermes2_5Mistral7bAwqParamsBodyMessagesTool) implementsAccountAIRunHfTheblokeExecuteOpenhermes2_5Mistral7bAwqParamsBodyMessagesToolUnion() {
}

// Satisfied by
// [AccountAIRunHfTheblokeExecuteOpenhermes2_5Mistral7bAwqParamsBodyMessagesToolsObject],
// [AccountAIRunHfTheblokeExecuteOpenhermes2_5Mistral7bAwqParamsBodyMessagesToolsObject],
// [AccountAIRunHfTheblokeExecuteOpenhermes2_5Mistral7bAwqParamsBodyMessagesTool].
type AccountAIRunHfTheblokeExecuteOpenhermes2_5Mistral7bAwqParamsBodyMessagesToolUnion interface {
	implementsAccountAIRunHfTheblokeExecuteOpenhermes2_5Mistral7bAwqParamsBodyMessagesToolUnion()
}

type AccountAIRunHfTheblokeExecuteOpenhermes2_5Mistral7bAwqParamsBodyMessagesToolsObject struct {
	// A brief description of what the tool does.
	Description param.Field[string] `json:"description,required"`
	// The name of the tool. More descriptive the better.
	Name param.Field[string] `json:"name,required"`
	// Schema defining the parameters accepted by the tool.
	Parameters param.Field[AccountAIRunHfTheblokeExecuteOpenhermes2_5Mistral7bAwqParamsBodyMessagesToolsObjectParameters] `json:"parameters,required"`
}

func (r AccountAIRunHfTheblokeExecuteOpenhermes2_5Mistral7bAwqParamsBodyMessagesToolsObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunHfTheblokeExecuteOpenhermes2_5Mistral7bAwqParamsBodyMessagesToolsObject) implementsAccountAIRunHfTheblokeExecuteOpenhermes2_5Mistral7bAwqParamsBodyMessagesToolUnion() {
}

// Schema defining the parameters accepted by the tool.
type AccountAIRunHfTheblokeExecuteOpenhermes2_5Mistral7bAwqParamsBodyMessagesToolsObjectParameters struct {
	// Definitions of each parameter.
	Properties param.Field[map[string]AccountAIRunHfTheblokeExecuteOpenhermes2_5Mistral7bAwqParamsBodyMessagesToolsObjectParametersProperties] `json:"properties,required"`
	// The type of the parameters object (usually 'object').
	Type param.Field[string] `json:"type,required"`
	// List of required parameter names.
	Required param.Field[[]string] `json:"required"`
}

func (r AccountAIRunHfTheblokeExecuteOpenhermes2_5Mistral7bAwqParamsBodyMessagesToolsObjectParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunHfTheblokeExecuteOpenhermes2_5Mistral7bAwqParamsBodyMessagesToolsObjectParametersProperties struct {
	// A description of the expected parameter.
	Description param.Field[string] `json:"description,required"`
	// The data type of the parameter.
	Type param.Field[string] `json:"type,required"`
}

func (r AccountAIRunHfTheblokeExecuteOpenhermes2_5Mistral7bAwqParamsBodyMessagesToolsObjectParametersProperties) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunHfTheblokeExecuteZephyr7bBetaAwqParams struct {
	QueueRequest param.Field[string]                                         `query:"queueRequest"`
	Body         AccountAIRunHfTheblokeExecuteZephyr7bBetaAwqParamsBodyUnion `json:"body"`
}

func (r AccountAIRunHfTheblokeExecuteZephyr7bBetaAwqParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

// URLQuery serializes [AccountAIRunHfTheblokeExecuteZephyr7bBetaAwqParams]'s query
// parameters as `url.Values`.
func (r AccountAIRunHfTheblokeExecuteZephyr7bBetaAwqParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountAIRunHfTheblokeExecuteZephyr7bBetaAwqParamsBody struct {
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

func (r AccountAIRunHfTheblokeExecuteZephyr7bBetaAwqParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunHfTheblokeExecuteZephyr7bBetaAwqParamsBody) implementsAccountAIRunHfTheblokeExecuteZephyr7bBetaAwqParamsBodyUnion() {
}

// Satisfied by [AccountAIRunHfTheblokeExecuteZephyr7bBetaAwqParamsBodyPrompt],
// [AccountAIRunHfTheblokeExecuteZephyr7bBetaAwqParamsBodyMessages],
// [AccountAIRunHfTheblokeExecuteZephyr7bBetaAwqParamsBody].
type AccountAIRunHfTheblokeExecuteZephyr7bBetaAwqParamsBodyUnion interface {
	implementsAccountAIRunHfTheblokeExecuteZephyr7bBetaAwqParamsBodyUnion()
}

type AccountAIRunHfTheblokeExecuteZephyr7bBetaAwqParamsBodyPrompt struct {
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
	RepetitionPenalty param.Field[float64]                                                                    `json:"repetition_penalty"`
	ResponseFormat    param.Field[AccountAIRunHfTheblokeExecuteZephyr7bBetaAwqParamsBodyPromptResponseFormat] `json:"response_format"`
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

func (r AccountAIRunHfTheblokeExecuteZephyr7bBetaAwqParamsBodyPrompt) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunHfTheblokeExecuteZephyr7bBetaAwqParamsBodyPrompt) implementsAccountAIRunHfTheblokeExecuteZephyr7bBetaAwqParamsBodyUnion() {
}

type AccountAIRunHfTheblokeExecuteZephyr7bBetaAwqParamsBodyPromptResponseFormat struct {
	JsonSchema param.Field[interface{}]                                                                    `json:"json_schema"`
	Type       param.Field[AccountAIRunHfTheblokeExecuteZephyr7bBetaAwqParamsBodyPromptResponseFormatType] `json:"type"`
}

func (r AccountAIRunHfTheblokeExecuteZephyr7bBetaAwqParamsBodyPromptResponseFormat) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunHfTheblokeExecuteZephyr7bBetaAwqParamsBodyPromptResponseFormatType string

const (
	AccountAIRunHfTheblokeExecuteZephyr7bBetaAwqParamsBodyPromptResponseFormatTypeJsonObject AccountAIRunHfTheblokeExecuteZephyr7bBetaAwqParamsBodyPromptResponseFormatType = "json_object"
	AccountAIRunHfTheblokeExecuteZephyr7bBetaAwqParamsBodyPromptResponseFormatTypeJsonSchema AccountAIRunHfTheblokeExecuteZephyr7bBetaAwqParamsBodyPromptResponseFormatType = "json_schema"
)

func (r AccountAIRunHfTheblokeExecuteZephyr7bBetaAwqParamsBodyPromptResponseFormatType) IsKnown() bool {
	switch r {
	case AccountAIRunHfTheblokeExecuteZephyr7bBetaAwqParamsBodyPromptResponseFormatTypeJsonObject, AccountAIRunHfTheblokeExecuteZephyr7bBetaAwqParamsBodyPromptResponseFormatTypeJsonSchema:
		return true
	}
	return false
}

type AccountAIRunHfTheblokeExecuteZephyr7bBetaAwqParamsBodyMessages struct {
	// An array of message objects representing the conversation history.
	Messages param.Field[[]AccountAIRunHfTheblokeExecuteZephyr7bBetaAwqParamsBodyMessagesMessage] `json:"messages,required"`
	// Decreases the likelihood of the model repeating the same lines verbatim.
	FrequencyPenalty param.Field[float64]                                                                  `json:"frequency_penalty"`
	Functions        param.Field[[]AccountAIRunHfTheblokeExecuteZephyr7bBetaAwqParamsBodyMessagesFunction] `json:"functions"`
	// The maximum number of tokens to generate in the response.
	MaxTokens param.Field[int64] `json:"max_tokens"`
	// Increases the likelihood of the model introducing new topics.
	PresencePenalty param.Field[float64] `json:"presence_penalty"`
	// If true, a chat template is not applied and you must adhere to the specific
	// model's expected formatting.
	Raw param.Field[bool] `json:"raw"`
	// Penalty for repeated tokens; higher values discourage repetition.
	RepetitionPenalty param.Field[float64]                                                                      `json:"repetition_penalty"`
	ResponseFormat    param.Field[AccountAIRunHfTheblokeExecuteZephyr7bBetaAwqParamsBodyMessagesResponseFormat] `json:"response_format"`
	// Random seed for reproducibility of the generation.
	Seed param.Field[int64] `json:"seed"`
	// If true, the response will be streamed back incrementally using SSE, Server Sent
	// Events.
	Stream param.Field[bool] `json:"stream"`
	// Controls the randomness of the output; higher values produce more random
	// results.
	Temperature param.Field[float64] `json:"temperature"`
	// A list of tools available for the assistant to use.
	Tools param.Field[[]AccountAIRunHfTheblokeExecuteZephyr7bBetaAwqParamsBodyMessagesToolUnion] `json:"tools"`
	// Limits the AI to choose from the top 'k' most probable words. Lower values make
	// responses more focused; higher values introduce more variety and potential
	// surprises.
	TopK param.Field[int64] `json:"top_k"`
	// Adjusts the creativity of the AI's responses by controlling how many possible
	// words it considers. Lower values make outputs more predictable; higher values
	// allow for more varied and creative responses.
	TopP param.Field[float64] `json:"top_p"`
}

func (r AccountAIRunHfTheblokeExecuteZephyr7bBetaAwqParamsBodyMessages) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunHfTheblokeExecuteZephyr7bBetaAwqParamsBodyMessages) implementsAccountAIRunHfTheblokeExecuteZephyr7bBetaAwqParamsBodyUnion() {
}

type AccountAIRunHfTheblokeExecuteZephyr7bBetaAwqParamsBodyMessagesMessage struct {
	// The content of the message as a string.
	Content param.Field[string] `json:"content,required"`
	// The role of the message sender (e.g., 'user', 'assistant', 'system', 'tool').
	Role param.Field[string] `json:"role,required"`
}

func (r AccountAIRunHfTheblokeExecuteZephyr7bBetaAwqParamsBodyMessagesMessage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunHfTheblokeExecuteZephyr7bBetaAwqParamsBodyMessagesFunction struct {
	Code param.Field[string] `json:"code,required"`
	Name param.Field[string] `json:"name,required"`
}

func (r AccountAIRunHfTheblokeExecuteZephyr7bBetaAwqParamsBodyMessagesFunction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunHfTheblokeExecuteZephyr7bBetaAwqParamsBodyMessagesResponseFormat struct {
	JsonSchema param.Field[interface{}]                                                                      `json:"json_schema"`
	Type       param.Field[AccountAIRunHfTheblokeExecuteZephyr7bBetaAwqParamsBodyMessagesResponseFormatType] `json:"type"`
}

func (r AccountAIRunHfTheblokeExecuteZephyr7bBetaAwqParamsBodyMessagesResponseFormat) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunHfTheblokeExecuteZephyr7bBetaAwqParamsBodyMessagesResponseFormatType string

const (
	AccountAIRunHfTheblokeExecuteZephyr7bBetaAwqParamsBodyMessagesResponseFormatTypeJsonObject AccountAIRunHfTheblokeExecuteZephyr7bBetaAwqParamsBodyMessagesResponseFormatType = "json_object"
	AccountAIRunHfTheblokeExecuteZephyr7bBetaAwqParamsBodyMessagesResponseFormatTypeJsonSchema AccountAIRunHfTheblokeExecuteZephyr7bBetaAwqParamsBodyMessagesResponseFormatType = "json_schema"
)

func (r AccountAIRunHfTheblokeExecuteZephyr7bBetaAwqParamsBodyMessagesResponseFormatType) IsKnown() bool {
	switch r {
	case AccountAIRunHfTheblokeExecuteZephyr7bBetaAwqParamsBodyMessagesResponseFormatTypeJsonObject, AccountAIRunHfTheblokeExecuteZephyr7bBetaAwqParamsBodyMessagesResponseFormatTypeJsonSchema:
		return true
	}
	return false
}

type AccountAIRunHfTheblokeExecuteZephyr7bBetaAwqParamsBodyMessagesTool struct {
	// A brief description of what the tool does.
	Description param.Field[string]      `json:"description"`
	Function    param.Field[interface{}] `json:"function"`
	// The name of the tool. More descriptive the better.
	Name       param.Field[string]      `json:"name"`
	Parameters param.Field[interface{}] `json:"parameters"`
	// Specifies the type of tool (e.g., 'function').
	Type param.Field[string] `json:"type"`
}

func (r AccountAIRunHfTheblokeExecuteZephyr7bBetaAwqParamsBodyMessagesTool) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunHfTheblokeExecuteZephyr7bBetaAwqParamsBodyMessagesTool) implementsAccountAIRunHfTheblokeExecuteZephyr7bBetaAwqParamsBodyMessagesToolUnion() {
}

// Satisfied by
// [AccountAIRunHfTheblokeExecuteZephyr7bBetaAwqParamsBodyMessagesToolsObject],
// [AccountAIRunHfTheblokeExecuteZephyr7bBetaAwqParamsBodyMessagesToolsObject],
// [AccountAIRunHfTheblokeExecuteZephyr7bBetaAwqParamsBodyMessagesTool].
type AccountAIRunHfTheblokeExecuteZephyr7bBetaAwqParamsBodyMessagesToolUnion interface {
	implementsAccountAIRunHfTheblokeExecuteZephyr7bBetaAwqParamsBodyMessagesToolUnion()
}

type AccountAIRunHfTheblokeExecuteZephyr7bBetaAwqParamsBodyMessagesToolsObject struct {
	// A brief description of what the tool does.
	Description param.Field[string] `json:"description,required"`
	// The name of the tool. More descriptive the better.
	Name param.Field[string] `json:"name,required"`
	// Schema defining the parameters accepted by the tool.
	Parameters param.Field[AccountAIRunHfTheblokeExecuteZephyr7bBetaAwqParamsBodyMessagesToolsObjectParameters] `json:"parameters,required"`
}

func (r AccountAIRunHfTheblokeExecuteZephyr7bBetaAwqParamsBodyMessagesToolsObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunHfTheblokeExecuteZephyr7bBetaAwqParamsBodyMessagesToolsObject) implementsAccountAIRunHfTheblokeExecuteZephyr7bBetaAwqParamsBodyMessagesToolUnion() {
}

// Schema defining the parameters accepted by the tool.
type AccountAIRunHfTheblokeExecuteZephyr7bBetaAwqParamsBodyMessagesToolsObjectParameters struct {
	// Definitions of each parameter.
	Properties param.Field[map[string]AccountAIRunHfTheblokeExecuteZephyr7bBetaAwqParamsBodyMessagesToolsObjectParametersProperties] `json:"properties,required"`
	// The type of the parameters object (usually 'object').
	Type param.Field[string] `json:"type,required"`
	// List of required parameter names.
	Required param.Field[[]string] `json:"required"`
}

func (r AccountAIRunHfTheblokeExecuteZephyr7bBetaAwqParamsBodyMessagesToolsObjectParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunHfTheblokeExecuteZephyr7bBetaAwqParamsBodyMessagesToolsObjectParametersProperties struct {
	// A description of the expected parameter.
	Description param.Field[string] `json:"description,required"`
	// The data type of the parameter.
	Type param.Field[string] `json:"type,required"`
}

func (r AccountAIRunHfTheblokeExecuteZephyr7bBetaAwqParamsBodyMessagesToolsObjectParametersProperties) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
