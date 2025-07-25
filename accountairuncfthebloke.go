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

// AccountAIRunCfTheblokeService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountAIRunCfTheblokeService] method instead.
type AccountAIRunCfTheblokeService struct {
	Options []option.RequestOption
}

// NewAccountAIRunCfTheblokeService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountAIRunCfTheblokeService(opts ...option.RequestOption) (r *AccountAIRunCfTheblokeService) {
	r = &AccountAIRunCfTheblokeService{}
	r.Options = opts
	return
}

// Execute @cf/thebloke/discolm-german-7b-v1-awq model.
func (r *AccountAIRunCfTheblokeService) ExecuteDiscolmGerman7bV1Awq(ctx context.Context, accountID string, params AccountAIRunCfTheblokeExecuteDiscolmGerman7bV1AwqParams, opts ...option.RequestOption) (res *AccountAIRunCfTheblokeExecuteDiscolmGerman7bV1AwqResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai/run/@cf/thebloke/discolm-german-7b-v1-awq", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type AccountAIRunCfTheblokeExecuteDiscolmGerman7bV1AwqResponse = interface{}

type AccountAIRunCfTheblokeExecuteDiscolmGerman7bV1AwqParams struct {
	QueueRequest param.Field[string]                                              `query:"queueRequest"`
	Body         AccountAIRunCfTheblokeExecuteDiscolmGerman7bV1AwqParamsBodyUnion `json:"body"`
}

func (r AccountAIRunCfTheblokeExecuteDiscolmGerman7bV1AwqParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

// URLQuery serializes [AccountAIRunCfTheblokeExecuteDiscolmGerman7bV1AwqParams]'s
// query parameters as `url.Values`.
func (r AccountAIRunCfTheblokeExecuteDiscolmGerman7bV1AwqParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountAIRunCfTheblokeExecuteDiscolmGerman7bV1AwqParamsBody struct {
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

func (r AccountAIRunCfTheblokeExecuteDiscolmGerman7bV1AwqParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfTheblokeExecuteDiscolmGerman7bV1AwqParamsBody) implementsAccountAIRunCfTheblokeExecuteDiscolmGerman7bV1AwqParamsBodyUnion() {
}

// Satisfied by
// [AccountAIRunCfTheblokeExecuteDiscolmGerman7bV1AwqParamsBodyPrompt],
// [AccountAIRunCfTheblokeExecuteDiscolmGerman7bV1AwqParamsBodyMessages],
// [AccountAIRunCfTheblokeExecuteDiscolmGerman7bV1AwqParamsBody].
type AccountAIRunCfTheblokeExecuteDiscolmGerman7bV1AwqParamsBodyUnion interface {
	implementsAccountAIRunCfTheblokeExecuteDiscolmGerman7bV1AwqParamsBodyUnion()
}

type AccountAIRunCfTheblokeExecuteDiscolmGerman7bV1AwqParamsBodyPrompt struct {
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
	ResponseFormat    param.Field[AccountAIRunCfTheblokeExecuteDiscolmGerman7bV1AwqParamsBodyPromptResponseFormat] `json:"response_format"`
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

func (r AccountAIRunCfTheblokeExecuteDiscolmGerman7bV1AwqParamsBodyPrompt) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfTheblokeExecuteDiscolmGerman7bV1AwqParamsBodyPrompt) implementsAccountAIRunCfTheblokeExecuteDiscolmGerman7bV1AwqParamsBodyUnion() {
}

type AccountAIRunCfTheblokeExecuteDiscolmGerman7bV1AwqParamsBodyPromptResponseFormat struct {
	JsonSchema param.Field[interface{}]                                                                         `json:"json_schema"`
	Type       param.Field[AccountAIRunCfTheblokeExecuteDiscolmGerman7bV1AwqParamsBodyPromptResponseFormatType] `json:"type"`
}

func (r AccountAIRunCfTheblokeExecuteDiscolmGerman7bV1AwqParamsBodyPromptResponseFormat) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfTheblokeExecuteDiscolmGerman7bV1AwqParamsBodyPromptResponseFormatType string

const (
	AccountAIRunCfTheblokeExecuteDiscolmGerman7bV1AwqParamsBodyPromptResponseFormatTypeJsonObject AccountAIRunCfTheblokeExecuteDiscolmGerman7bV1AwqParamsBodyPromptResponseFormatType = "json_object"
	AccountAIRunCfTheblokeExecuteDiscolmGerman7bV1AwqParamsBodyPromptResponseFormatTypeJsonSchema AccountAIRunCfTheblokeExecuteDiscolmGerman7bV1AwqParamsBodyPromptResponseFormatType = "json_schema"
)

func (r AccountAIRunCfTheblokeExecuteDiscolmGerman7bV1AwqParamsBodyPromptResponseFormatType) IsKnown() bool {
	switch r {
	case AccountAIRunCfTheblokeExecuteDiscolmGerman7bV1AwqParamsBodyPromptResponseFormatTypeJsonObject, AccountAIRunCfTheblokeExecuteDiscolmGerman7bV1AwqParamsBodyPromptResponseFormatTypeJsonSchema:
		return true
	}
	return false
}

type AccountAIRunCfTheblokeExecuteDiscolmGerman7bV1AwqParamsBodyMessages struct {
	// An array of message objects representing the conversation history.
	Messages param.Field[[]AccountAIRunCfTheblokeExecuteDiscolmGerman7bV1AwqParamsBodyMessagesMessage] `json:"messages,required"`
	// Decreases the likelihood of the model repeating the same lines verbatim.
	FrequencyPenalty param.Field[float64]                                                                       `json:"frequency_penalty"`
	Functions        param.Field[[]AccountAIRunCfTheblokeExecuteDiscolmGerman7bV1AwqParamsBodyMessagesFunction] `json:"functions"`
	// The maximum number of tokens to generate in the response.
	MaxTokens param.Field[int64] `json:"max_tokens"`
	// Increases the likelihood of the model introducing new topics.
	PresencePenalty param.Field[float64] `json:"presence_penalty"`
	// If true, a chat template is not applied and you must adhere to the specific
	// model's expected formatting.
	Raw param.Field[bool] `json:"raw"`
	// Penalty for repeated tokens; higher values discourage repetition.
	RepetitionPenalty param.Field[float64]                                                                           `json:"repetition_penalty"`
	ResponseFormat    param.Field[AccountAIRunCfTheblokeExecuteDiscolmGerman7bV1AwqParamsBodyMessagesResponseFormat] `json:"response_format"`
	// Random seed for reproducibility of the generation.
	Seed param.Field[int64] `json:"seed"`
	// If true, the response will be streamed back incrementally using SSE, Server Sent
	// Events.
	Stream param.Field[bool] `json:"stream"`
	// Controls the randomness of the output; higher values produce more random
	// results.
	Temperature param.Field[float64] `json:"temperature"`
	// A list of tools available for the assistant to use.
	Tools param.Field[[]AccountAIRunCfTheblokeExecuteDiscolmGerman7bV1AwqParamsBodyMessagesToolUnion] `json:"tools"`
	// Limits the AI to choose from the top 'k' most probable words. Lower values make
	// responses more focused; higher values introduce more variety and potential
	// surprises.
	TopK param.Field[int64] `json:"top_k"`
	// Adjusts the creativity of the AI's responses by controlling how many possible
	// words it considers. Lower values make outputs more predictable; higher values
	// allow for more varied and creative responses.
	TopP param.Field[float64] `json:"top_p"`
}

func (r AccountAIRunCfTheblokeExecuteDiscolmGerman7bV1AwqParamsBodyMessages) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfTheblokeExecuteDiscolmGerman7bV1AwqParamsBodyMessages) implementsAccountAIRunCfTheblokeExecuteDiscolmGerman7bV1AwqParamsBodyUnion() {
}

type AccountAIRunCfTheblokeExecuteDiscolmGerman7bV1AwqParamsBodyMessagesMessage struct {
	// The content of the message as a string.
	Content param.Field[string] `json:"content,required"`
	// The role of the message sender (e.g., 'user', 'assistant', 'system', 'tool').
	Role param.Field[string] `json:"role,required"`
}

func (r AccountAIRunCfTheblokeExecuteDiscolmGerman7bV1AwqParamsBodyMessagesMessage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfTheblokeExecuteDiscolmGerman7bV1AwqParamsBodyMessagesFunction struct {
	Code param.Field[string] `json:"code,required"`
	Name param.Field[string] `json:"name,required"`
}

func (r AccountAIRunCfTheblokeExecuteDiscolmGerman7bV1AwqParamsBodyMessagesFunction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfTheblokeExecuteDiscolmGerman7bV1AwqParamsBodyMessagesResponseFormat struct {
	JsonSchema param.Field[interface{}]                                                                           `json:"json_schema"`
	Type       param.Field[AccountAIRunCfTheblokeExecuteDiscolmGerman7bV1AwqParamsBodyMessagesResponseFormatType] `json:"type"`
}

func (r AccountAIRunCfTheblokeExecuteDiscolmGerman7bV1AwqParamsBodyMessagesResponseFormat) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfTheblokeExecuteDiscolmGerman7bV1AwqParamsBodyMessagesResponseFormatType string

const (
	AccountAIRunCfTheblokeExecuteDiscolmGerman7bV1AwqParamsBodyMessagesResponseFormatTypeJsonObject AccountAIRunCfTheblokeExecuteDiscolmGerman7bV1AwqParamsBodyMessagesResponseFormatType = "json_object"
	AccountAIRunCfTheblokeExecuteDiscolmGerman7bV1AwqParamsBodyMessagesResponseFormatTypeJsonSchema AccountAIRunCfTheblokeExecuteDiscolmGerman7bV1AwqParamsBodyMessagesResponseFormatType = "json_schema"
)

func (r AccountAIRunCfTheblokeExecuteDiscolmGerman7bV1AwqParamsBodyMessagesResponseFormatType) IsKnown() bool {
	switch r {
	case AccountAIRunCfTheblokeExecuteDiscolmGerman7bV1AwqParamsBodyMessagesResponseFormatTypeJsonObject, AccountAIRunCfTheblokeExecuteDiscolmGerman7bV1AwqParamsBodyMessagesResponseFormatTypeJsonSchema:
		return true
	}
	return false
}

type AccountAIRunCfTheblokeExecuteDiscolmGerman7bV1AwqParamsBodyMessagesTool struct {
	// A brief description of what the tool does.
	Description param.Field[string]      `json:"description"`
	Function    param.Field[interface{}] `json:"function"`
	// The name of the tool. More descriptive the better.
	Name       param.Field[string]      `json:"name"`
	Parameters param.Field[interface{}] `json:"parameters"`
	// Specifies the type of tool (e.g., 'function').
	Type param.Field[string] `json:"type"`
}

func (r AccountAIRunCfTheblokeExecuteDiscolmGerman7bV1AwqParamsBodyMessagesTool) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfTheblokeExecuteDiscolmGerman7bV1AwqParamsBodyMessagesTool) implementsAccountAIRunCfTheblokeExecuteDiscolmGerman7bV1AwqParamsBodyMessagesToolUnion() {
}

// Satisfied by
// [AccountAIRunCfTheblokeExecuteDiscolmGerman7bV1AwqParamsBodyMessagesToolsObject],
// [AccountAIRunCfTheblokeExecuteDiscolmGerman7bV1AwqParamsBodyMessagesToolsObject],
// [AccountAIRunCfTheblokeExecuteDiscolmGerman7bV1AwqParamsBodyMessagesTool].
type AccountAIRunCfTheblokeExecuteDiscolmGerman7bV1AwqParamsBodyMessagesToolUnion interface {
	implementsAccountAIRunCfTheblokeExecuteDiscolmGerman7bV1AwqParamsBodyMessagesToolUnion()
}

type AccountAIRunCfTheblokeExecuteDiscolmGerman7bV1AwqParamsBodyMessagesToolsObject struct {
	// A brief description of what the tool does.
	Description param.Field[string] `json:"description,required"`
	// The name of the tool. More descriptive the better.
	Name param.Field[string] `json:"name,required"`
	// Schema defining the parameters accepted by the tool.
	Parameters param.Field[AccountAIRunCfTheblokeExecuteDiscolmGerman7bV1AwqParamsBodyMessagesToolsObjectParameters] `json:"parameters,required"`
}

func (r AccountAIRunCfTheblokeExecuteDiscolmGerman7bV1AwqParamsBodyMessagesToolsObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfTheblokeExecuteDiscolmGerman7bV1AwqParamsBodyMessagesToolsObject) implementsAccountAIRunCfTheblokeExecuteDiscolmGerman7bV1AwqParamsBodyMessagesToolUnion() {
}

// Schema defining the parameters accepted by the tool.
type AccountAIRunCfTheblokeExecuteDiscolmGerman7bV1AwqParamsBodyMessagesToolsObjectParameters struct {
	// Definitions of each parameter.
	Properties param.Field[map[string]AccountAIRunCfTheblokeExecuteDiscolmGerman7bV1AwqParamsBodyMessagesToolsObjectParametersProperties] `json:"properties,required"`
	// The type of the parameters object (usually 'object').
	Type param.Field[string] `json:"type,required"`
	// List of required parameter names.
	Required param.Field[[]string] `json:"required"`
}

func (r AccountAIRunCfTheblokeExecuteDiscolmGerman7bV1AwqParamsBodyMessagesToolsObjectParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfTheblokeExecuteDiscolmGerman7bV1AwqParamsBodyMessagesToolsObjectParametersProperties struct {
	// A description of the expected parameter.
	Description param.Field[string] `json:"description,required"`
	// The data type of the parameter.
	Type param.Field[string] `json:"type,required"`
}

func (r AccountAIRunCfTheblokeExecuteDiscolmGerman7bV1AwqParamsBodyMessagesToolsObjectParametersProperties) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
