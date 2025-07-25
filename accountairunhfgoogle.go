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

// AccountAIRunHfGoogleService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountAIRunHfGoogleService] method instead.
type AccountAIRunHfGoogleService struct {
	Options []option.RequestOption
}

// NewAccountAIRunHfGoogleService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountAIRunHfGoogleService(opts ...option.RequestOption) (r *AccountAIRunHfGoogleService) {
	r = &AccountAIRunHfGoogleService{}
	r.Options = opts
	return
}

// Execute @hf/google/gemma-7b-it model.
func (r *AccountAIRunHfGoogleService) ExecuteGemma7bIt(ctx context.Context, accountID string, params AccountAIRunHfGoogleExecuteGemma7bItParams, opts ...option.RequestOption) (res *AccountAIRunHfGoogleExecuteGemma7bItResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai/run/@hf/google/gemma-7b-it", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type AccountAIRunHfGoogleExecuteGemma7bItResponse = interface{}

type AccountAIRunHfGoogleExecuteGemma7bItParams struct {
	QueueRequest param.Field[string]                                 `query:"queueRequest"`
	Body         AccountAIRunHfGoogleExecuteGemma7bItParamsBodyUnion `json:"body"`
}

func (r AccountAIRunHfGoogleExecuteGemma7bItParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

// URLQuery serializes [AccountAIRunHfGoogleExecuteGemma7bItParams]'s query
// parameters as `url.Values`.
func (r AccountAIRunHfGoogleExecuteGemma7bItParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountAIRunHfGoogleExecuteGemma7bItParamsBody struct {
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

func (r AccountAIRunHfGoogleExecuteGemma7bItParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunHfGoogleExecuteGemma7bItParamsBody) implementsAccountAIRunHfGoogleExecuteGemma7bItParamsBodyUnion() {
}

// Satisfied by [AccountAIRunHfGoogleExecuteGemma7bItParamsBodyPrompt],
// [AccountAIRunHfGoogleExecuteGemma7bItParamsBodyMessages],
// [AccountAIRunHfGoogleExecuteGemma7bItParamsBody].
type AccountAIRunHfGoogleExecuteGemma7bItParamsBodyUnion interface {
	implementsAccountAIRunHfGoogleExecuteGemma7bItParamsBodyUnion()
}

type AccountAIRunHfGoogleExecuteGemma7bItParamsBodyPrompt struct {
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
	RepetitionPenalty param.Field[float64]                                                            `json:"repetition_penalty"`
	ResponseFormat    param.Field[AccountAIRunHfGoogleExecuteGemma7bItParamsBodyPromptResponseFormat] `json:"response_format"`
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

func (r AccountAIRunHfGoogleExecuteGemma7bItParamsBodyPrompt) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunHfGoogleExecuteGemma7bItParamsBodyPrompt) implementsAccountAIRunHfGoogleExecuteGemma7bItParamsBodyUnion() {
}

type AccountAIRunHfGoogleExecuteGemma7bItParamsBodyPromptResponseFormat struct {
	JsonSchema param.Field[interface{}]                                                            `json:"json_schema"`
	Type       param.Field[AccountAIRunHfGoogleExecuteGemma7bItParamsBodyPromptResponseFormatType] `json:"type"`
}

func (r AccountAIRunHfGoogleExecuteGemma7bItParamsBodyPromptResponseFormat) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunHfGoogleExecuteGemma7bItParamsBodyPromptResponseFormatType string

const (
	AccountAIRunHfGoogleExecuteGemma7bItParamsBodyPromptResponseFormatTypeJsonObject AccountAIRunHfGoogleExecuteGemma7bItParamsBodyPromptResponseFormatType = "json_object"
	AccountAIRunHfGoogleExecuteGemma7bItParamsBodyPromptResponseFormatTypeJsonSchema AccountAIRunHfGoogleExecuteGemma7bItParamsBodyPromptResponseFormatType = "json_schema"
)

func (r AccountAIRunHfGoogleExecuteGemma7bItParamsBodyPromptResponseFormatType) IsKnown() bool {
	switch r {
	case AccountAIRunHfGoogleExecuteGemma7bItParamsBodyPromptResponseFormatTypeJsonObject, AccountAIRunHfGoogleExecuteGemma7bItParamsBodyPromptResponseFormatTypeJsonSchema:
		return true
	}
	return false
}

type AccountAIRunHfGoogleExecuteGemma7bItParamsBodyMessages struct {
	// An array of message objects representing the conversation history.
	Messages param.Field[[]AccountAIRunHfGoogleExecuteGemma7bItParamsBodyMessagesMessage] `json:"messages,required"`
	// Decreases the likelihood of the model repeating the same lines verbatim.
	FrequencyPenalty param.Field[float64]                                                          `json:"frequency_penalty"`
	Functions        param.Field[[]AccountAIRunHfGoogleExecuteGemma7bItParamsBodyMessagesFunction] `json:"functions"`
	// The maximum number of tokens to generate in the response.
	MaxTokens param.Field[int64] `json:"max_tokens"`
	// Increases the likelihood of the model introducing new topics.
	PresencePenalty param.Field[float64] `json:"presence_penalty"`
	// If true, a chat template is not applied and you must adhere to the specific
	// model's expected formatting.
	Raw param.Field[bool] `json:"raw"`
	// Penalty for repeated tokens; higher values discourage repetition.
	RepetitionPenalty param.Field[float64]                                                              `json:"repetition_penalty"`
	ResponseFormat    param.Field[AccountAIRunHfGoogleExecuteGemma7bItParamsBodyMessagesResponseFormat] `json:"response_format"`
	// Random seed for reproducibility of the generation.
	Seed param.Field[int64] `json:"seed"`
	// If true, the response will be streamed back incrementally using SSE, Server Sent
	// Events.
	Stream param.Field[bool] `json:"stream"`
	// Controls the randomness of the output; higher values produce more random
	// results.
	Temperature param.Field[float64] `json:"temperature"`
	// A list of tools available for the assistant to use.
	Tools param.Field[[]AccountAIRunHfGoogleExecuteGemma7bItParamsBodyMessagesToolUnion] `json:"tools"`
	// Limits the AI to choose from the top 'k' most probable words. Lower values make
	// responses more focused; higher values introduce more variety and potential
	// surprises.
	TopK param.Field[int64] `json:"top_k"`
	// Adjusts the creativity of the AI's responses by controlling how many possible
	// words it considers. Lower values make outputs more predictable; higher values
	// allow for more varied and creative responses.
	TopP param.Field[float64] `json:"top_p"`
}

func (r AccountAIRunHfGoogleExecuteGemma7bItParamsBodyMessages) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunHfGoogleExecuteGemma7bItParamsBodyMessages) implementsAccountAIRunHfGoogleExecuteGemma7bItParamsBodyUnion() {
}

type AccountAIRunHfGoogleExecuteGemma7bItParamsBodyMessagesMessage struct {
	// The content of the message as a string.
	Content param.Field[string] `json:"content,required"`
	// The role of the message sender (e.g., 'user', 'assistant', 'system', 'tool').
	Role param.Field[string] `json:"role,required"`
}

func (r AccountAIRunHfGoogleExecuteGemma7bItParamsBodyMessagesMessage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunHfGoogleExecuteGemma7bItParamsBodyMessagesFunction struct {
	Code param.Field[string] `json:"code,required"`
	Name param.Field[string] `json:"name,required"`
}

func (r AccountAIRunHfGoogleExecuteGemma7bItParamsBodyMessagesFunction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunHfGoogleExecuteGemma7bItParamsBodyMessagesResponseFormat struct {
	JsonSchema param.Field[interface{}]                                                              `json:"json_schema"`
	Type       param.Field[AccountAIRunHfGoogleExecuteGemma7bItParamsBodyMessagesResponseFormatType] `json:"type"`
}

func (r AccountAIRunHfGoogleExecuteGemma7bItParamsBodyMessagesResponseFormat) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunHfGoogleExecuteGemma7bItParamsBodyMessagesResponseFormatType string

const (
	AccountAIRunHfGoogleExecuteGemma7bItParamsBodyMessagesResponseFormatTypeJsonObject AccountAIRunHfGoogleExecuteGemma7bItParamsBodyMessagesResponseFormatType = "json_object"
	AccountAIRunHfGoogleExecuteGemma7bItParamsBodyMessagesResponseFormatTypeJsonSchema AccountAIRunHfGoogleExecuteGemma7bItParamsBodyMessagesResponseFormatType = "json_schema"
)

func (r AccountAIRunHfGoogleExecuteGemma7bItParamsBodyMessagesResponseFormatType) IsKnown() bool {
	switch r {
	case AccountAIRunHfGoogleExecuteGemma7bItParamsBodyMessagesResponseFormatTypeJsonObject, AccountAIRunHfGoogleExecuteGemma7bItParamsBodyMessagesResponseFormatTypeJsonSchema:
		return true
	}
	return false
}

type AccountAIRunHfGoogleExecuteGemma7bItParamsBodyMessagesTool struct {
	// A brief description of what the tool does.
	Description param.Field[string]      `json:"description"`
	Function    param.Field[interface{}] `json:"function"`
	// The name of the tool. More descriptive the better.
	Name       param.Field[string]      `json:"name"`
	Parameters param.Field[interface{}] `json:"parameters"`
	// Specifies the type of tool (e.g., 'function').
	Type param.Field[string] `json:"type"`
}

func (r AccountAIRunHfGoogleExecuteGemma7bItParamsBodyMessagesTool) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunHfGoogleExecuteGemma7bItParamsBodyMessagesTool) implementsAccountAIRunHfGoogleExecuteGemma7bItParamsBodyMessagesToolUnion() {
}

// Satisfied by
// [AccountAIRunHfGoogleExecuteGemma7bItParamsBodyMessagesToolsObject],
// [AccountAIRunHfGoogleExecuteGemma7bItParamsBodyMessagesToolsObject],
// [AccountAIRunHfGoogleExecuteGemma7bItParamsBodyMessagesTool].
type AccountAIRunHfGoogleExecuteGemma7bItParamsBodyMessagesToolUnion interface {
	implementsAccountAIRunHfGoogleExecuteGemma7bItParamsBodyMessagesToolUnion()
}

type AccountAIRunHfGoogleExecuteGemma7bItParamsBodyMessagesToolsObject struct {
	// A brief description of what the tool does.
	Description param.Field[string] `json:"description,required"`
	// The name of the tool. More descriptive the better.
	Name param.Field[string] `json:"name,required"`
	// Schema defining the parameters accepted by the tool.
	Parameters param.Field[AccountAIRunHfGoogleExecuteGemma7bItParamsBodyMessagesToolsObjectParameters] `json:"parameters,required"`
}

func (r AccountAIRunHfGoogleExecuteGemma7bItParamsBodyMessagesToolsObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunHfGoogleExecuteGemma7bItParamsBodyMessagesToolsObject) implementsAccountAIRunHfGoogleExecuteGemma7bItParamsBodyMessagesToolUnion() {
}

// Schema defining the parameters accepted by the tool.
type AccountAIRunHfGoogleExecuteGemma7bItParamsBodyMessagesToolsObjectParameters struct {
	// Definitions of each parameter.
	Properties param.Field[map[string]AccountAIRunHfGoogleExecuteGemma7bItParamsBodyMessagesToolsObjectParametersProperties] `json:"properties,required"`
	// The type of the parameters object (usually 'object').
	Type param.Field[string] `json:"type,required"`
	// List of required parameter names.
	Required param.Field[[]string] `json:"required"`
}

func (r AccountAIRunHfGoogleExecuteGemma7bItParamsBodyMessagesToolsObjectParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunHfGoogleExecuteGemma7bItParamsBodyMessagesToolsObjectParametersProperties struct {
	// A description of the expected parameter.
	Description param.Field[string] `json:"description,required"`
	// The data type of the parameter.
	Type param.Field[string] `json:"type,required"`
}

func (r AccountAIRunHfGoogleExecuteGemma7bItParamsBodyMessagesToolsObjectParametersProperties) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
