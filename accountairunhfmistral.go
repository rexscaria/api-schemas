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

// AccountAIRunHfMistralService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountAIRunHfMistralService] method instead.
type AccountAIRunHfMistralService struct {
	Options []option.RequestOption
}

// NewAccountAIRunHfMistralService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountAIRunHfMistralService(opts ...option.RequestOption) (r *AccountAIRunHfMistralService) {
	r = &AccountAIRunHfMistralService{}
	r.Options = opts
	return
}

// Execute @hf/mistral/mistral-7b-instruct-v0.2 model.
func (r *AccountAIRunHfMistralService) ExecuteMistral7bInstructV0_2(ctx context.Context, accountID string, params AccountAIRunHfMistralExecuteMistral7bInstructV0_2Params, opts ...option.RequestOption) (res *AccountAIRunHfMistralExecuteMistral7bInstructV0_2Response, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai/run/@hf/mistral/mistral-7b-instruct-v0.2", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type AccountAIRunHfMistralExecuteMistral7bInstructV0_2Response = interface{}

type AccountAIRunHfMistralExecuteMistral7bInstructV0_2Params struct {
	QueueRequest param.Field[string]                                              `query:"queueRequest"`
	Body         AccountAIRunHfMistralExecuteMistral7bInstructV0_2ParamsBodyUnion `json:"body"`
}

func (r AccountAIRunHfMistralExecuteMistral7bInstructV0_2Params) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

// URLQuery serializes [AccountAIRunHfMistralExecuteMistral7bInstructV0_2Params]'s
// query parameters as `url.Values`.
func (r AccountAIRunHfMistralExecuteMistral7bInstructV0_2Params) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountAIRunHfMistralExecuteMistral7bInstructV0_2ParamsBody struct {
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

func (r AccountAIRunHfMistralExecuteMistral7bInstructV0_2ParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunHfMistralExecuteMistral7bInstructV0_2ParamsBody) implementsAccountAIRunHfMistralExecuteMistral7bInstructV0_2ParamsBodyUnion() {
}

// Satisfied by
// [AccountAIRunHfMistralExecuteMistral7bInstructV0_2ParamsBodyPrompt],
// [AccountAIRunHfMistralExecuteMistral7bInstructV0_2ParamsBodyMessages],
// [AccountAIRunHfMistralExecuteMistral7bInstructV0_2ParamsBody].
type AccountAIRunHfMistralExecuteMistral7bInstructV0_2ParamsBodyUnion interface {
	implementsAccountAIRunHfMistralExecuteMistral7bInstructV0_2ParamsBodyUnion()
}

type AccountAIRunHfMistralExecuteMistral7bInstructV0_2ParamsBodyPrompt struct {
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
	ResponseFormat    param.Field[AccountAIRunHfMistralExecuteMistral7bInstructV0_2ParamsBodyPromptResponseFormat] `json:"response_format"`
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

func (r AccountAIRunHfMistralExecuteMistral7bInstructV0_2ParamsBodyPrompt) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunHfMistralExecuteMistral7bInstructV0_2ParamsBodyPrompt) implementsAccountAIRunHfMistralExecuteMistral7bInstructV0_2ParamsBodyUnion() {
}

type AccountAIRunHfMistralExecuteMistral7bInstructV0_2ParamsBodyPromptResponseFormat struct {
	JsonSchema param.Field[interface{}]                                                                         `json:"json_schema"`
	Type       param.Field[AccountAIRunHfMistralExecuteMistral7bInstructV0_2ParamsBodyPromptResponseFormatType] `json:"type"`
}

func (r AccountAIRunHfMistralExecuteMistral7bInstructV0_2ParamsBodyPromptResponseFormat) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunHfMistralExecuteMistral7bInstructV0_2ParamsBodyPromptResponseFormatType string

const (
	AccountAIRunHfMistralExecuteMistral7bInstructV0_2ParamsBodyPromptResponseFormatTypeJsonObject AccountAIRunHfMistralExecuteMistral7bInstructV0_2ParamsBodyPromptResponseFormatType = "json_object"
	AccountAIRunHfMistralExecuteMistral7bInstructV0_2ParamsBodyPromptResponseFormatTypeJsonSchema AccountAIRunHfMistralExecuteMistral7bInstructV0_2ParamsBodyPromptResponseFormatType = "json_schema"
)

func (r AccountAIRunHfMistralExecuteMistral7bInstructV0_2ParamsBodyPromptResponseFormatType) IsKnown() bool {
	switch r {
	case AccountAIRunHfMistralExecuteMistral7bInstructV0_2ParamsBodyPromptResponseFormatTypeJsonObject, AccountAIRunHfMistralExecuteMistral7bInstructV0_2ParamsBodyPromptResponseFormatTypeJsonSchema:
		return true
	}
	return false
}

type AccountAIRunHfMistralExecuteMistral7bInstructV0_2ParamsBodyMessages struct {
	// An array of message objects representing the conversation history.
	Messages param.Field[[]AccountAIRunHfMistralExecuteMistral7bInstructV0_2ParamsBodyMessagesMessage] `json:"messages,required"`
	// Decreases the likelihood of the model repeating the same lines verbatim.
	FrequencyPenalty param.Field[float64]                                                                       `json:"frequency_penalty"`
	Functions        param.Field[[]AccountAIRunHfMistralExecuteMistral7bInstructV0_2ParamsBodyMessagesFunction] `json:"functions"`
	// The maximum number of tokens to generate in the response.
	MaxTokens param.Field[int64] `json:"max_tokens"`
	// Increases the likelihood of the model introducing new topics.
	PresencePenalty param.Field[float64] `json:"presence_penalty"`
	// If true, a chat template is not applied and you must adhere to the specific
	// model's expected formatting.
	Raw param.Field[bool] `json:"raw"`
	// Penalty for repeated tokens; higher values discourage repetition.
	RepetitionPenalty param.Field[float64]                                                                           `json:"repetition_penalty"`
	ResponseFormat    param.Field[AccountAIRunHfMistralExecuteMistral7bInstructV0_2ParamsBodyMessagesResponseFormat] `json:"response_format"`
	// Random seed for reproducibility of the generation.
	Seed param.Field[int64] `json:"seed"`
	// If true, the response will be streamed back incrementally using SSE, Server Sent
	// Events.
	Stream param.Field[bool] `json:"stream"`
	// Controls the randomness of the output; higher values produce more random
	// results.
	Temperature param.Field[float64] `json:"temperature"`
	// A list of tools available for the assistant to use.
	Tools param.Field[[]AccountAIRunHfMistralExecuteMistral7bInstructV0_2ParamsBodyMessagesToolUnion] `json:"tools"`
	// Limits the AI to choose from the top 'k' most probable words. Lower values make
	// responses more focused; higher values introduce more variety and potential
	// surprises.
	TopK param.Field[int64] `json:"top_k"`
	// Adjusts the creativity of the AI's responses by controlling how many possible
	// words it considers. Lower values make outputs more predictable; higher values
	// allow for more varied and creative responses.
	TopP param.Field[float64] `json:"top_p"`
}

func (r AccountAIRunHfMistralExecuteMistral7bInstructV0_2ParamsBodyMessages) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunHfMistralExecuteMistral7bInstructV0_2ParamsBodyMessages) implementsAccountAIRunHfMistralExecuteMistral7bInstructV0_2ParamsBodyUnion() {
}

type AccountAIRunHfMistralExecuteMistral7bInstructV0_2ParamsBodyMessagesMessage struct {
	// The content of the message as a string.
	Content param.Field[string] `json:"content,required"`
	// The role of the message sender (e.g., 'user', 'assistant', 'system', 'tool').
	Role param.Field[string] `json:"role,required"`
}

func (r AccountAIRunHfMistralExecuteMistral7bInstructV0_2ParamsBodyMessagesMessage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunHfMistralExecuteMistral7bInstructV0_2ParamsBodyMessagesFunction struct {
	Code param.Field[string] `json:"code,required"`
	Name param.Field[string] `json:"name,required"`
}

func (r AccountAIRunHfMistralExecuteMistral7bInstructV0_2ParamsBodyMessagesFunction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunHfMistralExecuteMistral7bInstructV0_2ParamsBodyMessagesResponseFormat struct {
	JsonSchema param.Field[interface{}]                                                                           `json:"json_schema"`
	Type       param.Field[AccountAIRunHfMistralExecuteMistral7bInstructV0_2ParamsBodyMessagesResponseFormatType] `json:"type"`
}

func (r AccountAIRunHfMistralExecuteMistral7bInstructV0_2ParamsBodyMessagesResponseFormat) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunHfMistralExecuteMistral7bInstructV0_2ParamsBodyMessagesResponseFormatType string

const (
	AccountAIRunHfMistralExecuteMistral7bInstructV0_2ParamsBodyMessagesResponseFormatTypeJsonObject AccountAIRunHfMistralExecuteMistral7bInstructV0_2ParamsBodyMessagesResponseFormatType = "json_object"
	AccountAIRunHfMistralExecuteMistral7bInstructV0_2ParamsBodyMessagesResponseFormatTypeJsonSchema AccountAIRunHfMistralExecuteMistral7bInstructV0_2ParamsBodyMessagesResponseFormatType = "json_schema"
)

func (r AccountAIRunHfMistralExecuteMistral7bInstructV0_2ParamsBodyMessagesResponseFormatType) IsKnown() bool {
	switch r {
	case AccountAIRunHfMistralExecuteMistral7bInstructV0_2ParamsBodyMessagesResponseFormatTypeJsonObject, AccountAIRunHfMistralExecuteMistral7bInstructV0_2ParamsBodyMessagesResponseFormatTypeJsonSchema:
		return true
	}
	return false
}

type AccountAIRunHfMistralExecuteMistral7bInstructV0_2ParamsBodyMessagesTool struct {
	// A brief description of what the tool does.
	Description param.Field[string]      `json:"description"`
	Function    param.Field[interface{}] `json:"function"`
	// The name of the tool. More descriptive the better.
	Name       param.Field[string]      `json:"name"`
	Parameters param.Field[interface{}] `json:"parameters"`
	// Specifies the type of tool (e.g., 'function').
	Type param.Field[string] `json:"type"`
}

func (r AccountAIRunHfMistralExecuteMistral7bInstructV0_2ParamsBodyMessagesTool) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunHfMistralExecuteMistral7bInstructV0_2ParamsBodyMessagesTool) implementsAccountAIRunHfMistralExecuteMistral7bInstructV0_2ParamsBodyMessagesToolUnion() {
}

// Satisfied by
// [AccountAIRunHfMistralExecuteMistral7bInstructV0_2ParamsBodyMessagesToolsObject],
// [AccountAIRunHfMistralExecuteMistral7bInstructV0_2ParamsBodyMessagesToolsObject],
// [AccountAIRunHfMistralExecuteMistral7bInstructV0_2ParamsBodyMessagesTool].
type AccountAIRunHfMistralExecuteMistral7bInstructV0_2ParamsBodyMessagesToolUnion interface {
	implementsAccountAIRunHfMistralExecuteMistral7bInstructV0_2ParamsBodyMessagesToolUnion()
}

type AccountAIRunHfMistralExecuteMistral7bInstructV0_2ParamsBodyMessagesToolsObject struct {
	// A brief description of what the tool does.
	Description param.Field[string] `json:"description,required"`
	// The name of the tool. More descriptive the better.
	Name param.Field[string] `json:"name,required"`
	// Schema defining the parameters accepted by the tool.
	Parameters param.Field[AccountAIRunHfMistralExecuteMistral7bInstructV0_2ParamsBodyMessagesToolsObjectParameters] `json:"parameters,required"`
}

func (r AccountAIRunHfMistralExecuteMistral7bInstructV0_2ParamsBodyMessagesToolsObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunHfMistralExecuteMistral7bInstructV0_2ParamsBodyMessagesToolsObject) implementsAccountAIRunHfMistralExecuteMistral7bInstructV0_2ParamsBodyMessagesToolUnion() {
}

// Schema defining the parameters accepted by the tool.
type AccountAIRunHfMistralExecuteMistral7bInstructV0_2ParamsBodyMessagesToolsObjectParameters struct {
	// Definitions of each parameter.
	Properties param.Field[map[string]AccountAIRunHfMistralExecuteMistral7bInstructV0_2ParamsBodyMessagesToolsObjectParametersProperties] `json:"properties,required"`
	// The type of the parameters object (usually 'object').
	Type param.Field[string] `json:"type,required"`
	// List of required parameter names.
	Required param.Field[[]string] `json:"required"`
}

func (r AccountAIRunHfMistralExecuteMistral7bInstructV0_2ParamsBodyMessagesToolsObjectParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunHfMistralExecuteMistral7bInstructV0_2ParamsBodyMessagesToolsObjectParametersProperties struct {
	// A description of the expected parameter.
	Description param.Field[string] `json:"description,required"`
	// The data type of the parameter.
	Type param.Field[string] `json:"type,required"`
}

func (r AccountAIRunHfMistralExecuteMistral7bInstructV0_2ParamsBodyMessagesToolsObjectParametersProperties) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
