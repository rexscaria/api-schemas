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

// AccountAIRunCfFblgitService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountAIRunCfFblgitService] method instead.
type AccountAIRunCfFblgitService struct {
	Options []option.RequestOption
}

// NewAccountAIRunCfFblgitService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountAIRunCfFblgitService(opts ...option.RequestOption) (r *AccountAIRunCfFblgitService) {
	r = &AccountAIRunCfFblgitService{}
	r.Options = opts
	return
}

// Execute @cf/fblgit/una-cybertron-7b-v2-bf16 model.
func (r *AccountAIRunCfFblgitService) ExecuteUnaCybertron7bV2Bf16(ctx context.Context, accountID string, params AccountAIRunCfFblgitExecuteUnaCybertron7bV2Bf16Params, opts ...option.RequestOption) (res *AccountAIRunCfFblgitExecuteUnaCybertron7bV2Bf16Response, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai/run/@cf/fblgit/una-cybertron-7b-v2-bf16", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type AccountAIRunCfFblgitExecuteUnaCybertron7bV2Bf16Response = interface{}

type AccountAIRunCfFblgitExecuteUnaCybertron7bV2Bf16Params struct {
	QueueRequest param.Field[string]                                            `query:"queueRequest"`
	Body         AccountAIRunCfFblgitExecuteUnaCybertron7bV2Bf16ParamsBodyUnion `json:"body"`
}

func (r AccountAIRunCfFblgitExecuteUnaCybertron7bV2Bf16Params) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

// URLQuery serializes [AccountAIRunCfFblgitExecuteUnaCybertron7bV2Bf16Params]'s
// query parameters as `url.Values`.
func (r AccountAIRunCfFblgitExecuteUnaCybertron7bV2Bf16Params) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountAIRunCfFblgitExecuteUnaCybertron7bV2Bf16ParamsBody struct {
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

func (r AccountAIRunCfFblgitExecuteUnaCybertron7bV2Bf16ParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfFblgitExecuteUnaCybertron7bV2Bf16ParamsBody) implementsAccountAIRunCfFblgitExecuteUnaCybertron7bV2Bf16ParamsBodyUnion() {
}

// Satisfied by [AccountAIRunCfFblgitExecuteUnaCybertron7bV2Bf16ParamsBodyPrompt],
// [AccountAIRunCfFblgitExecuteUnaCybertron7bV2Bf16ParamsBodyMessages],
// [AccountAIRunCfFblgitExecuteUnaCybertron7bV2Bf16ParamsBody].
type AccountAIRunCfFblgitExecuteUnaCybertron7bV2Bf16ParamsBodyUnion interface {
	implementsAccountAIRunCfFblgitExecuteUnaCybertron7bV2Bf16ParamsBodyUnion()
}

type AccountAIRunCfFblgitExecuteUnaCybertron7bV2Bf16ParamsBodyPrompt struct {
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
	RepetitionPenalty param.Field[float64]                                                                       `json:"repetition_penalty"`
	ResponseFormat    param.Field[AccountAIRunCfFblgitExecuteUnaCybertron7bV2Bf16ParamsBodyPromptResponseFormat] `json:"response_format"`
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

func (r AccountAIRunCfFblgitExecuteUnaCybertron7bV2Bf16ParamsBodyPrompt) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfFblgitExecuteUnaCybertron7bV2Bf16ParamsBodyPrompt) implementsAccountAIRunCfFblgitExecuteUnaCybertron7bV2Bf16ParamsBodyUnion() {
}

type AccountAIRunCfFblgitExecuteUnaCybertron7bV2Bf16ParamsBodyPromptResponseFormat struct {
	JsonSchema param.Field[interface{}]                                                                       `json:"json_schema"`
	Type       param.Field[AccountAIRunCfFblgitExecuteUnaCybertron7bV2Bf16ParamsBodyPromptResponseFormatType] `json:"type"`
}

func (r AccountAIRunCfFblgitExecuteUnaCybertron7bV2Bf16ParamsBodyPromptResponseFormat) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfFblgitExecuteUnaCybertron7bV2Bf16ParamsBodyPromptResponseFormatType string

const (
	AccountAIRunCfFblgitExecuteUnaCybertron7bV2Bf16ParamsBodyPromptResponseFormatTypeJsonObject AccountAIRunCfFblgitExecuteUnaCybertron7bV2Bf16ParamsBodyPromptResponseFormatType = "json_object"
	AccountAIRunCfFblgitExecuteUnaCybertron7bV2Bf16ParamsBodyPromptResponseFormatTypeJsonSchema AccountAIRunCfFblgitExecuteUnaCybertron7bV2Bf16ParamsBodyPromptResponseFormatType = "json_schema"
)

func (r AccountAIRunCfFblgitExecuteUnaCybertron7bV2Bf16ParamsBodyPromptResponseFormatType) IsKnown() bool {
	switch r {
	case AccountAIRunCfFblgitExecuteUnaCybertron7bV2Bf16ParamsBodyPromptResponseFormatTypeJsonObject, AccountAIRunCfFblgitExecuteUnaCybertron7bV2Bf16ParamsBodyPromptResponseFormatTypeJsonSchema:
		return true
	}
	return false
}

type AccountAIRunCfFblgitExecuteUnaCybertron7bV2Bf16ParamsBodyMessages struct {
	// An array of message objects representing the conversation history.
	Messages param.Field[[]AccountAIRunCfFblgitExecuteUnaCybertron7bV2Bf16ParamsBodyMessagesMessage] `json:"messages,required"`
	// Decreases the likelihood of the model repeating the same lines verbatim.
	FrequencyPenalty param.Field[float64]                                                                     `json:"frequency_penalty"`
	Functions        param.Field[[]AccountAIRunCfFblgitExecuteUnaCybertron7bV2Bf16ParamsBodyMessagesFunction] `json:"functions"`
	// The maximum number of tokens to generate in the response.
	MaxTokens param.Field[int64] `json:"max_tokens"`
	// Increases the likelihood of the model introducing new topics.
	PresencePenalty param.Field[float64] `json:"presence_penalty"`
	// If true, a chat template is not applied and you must adhere to the specific
	// model's expected formatting.
	Raw param.Field[bool] `json:"raw"`
	// Penalty for repeated tokens; higher values discourage repetition.
	RepetitionPenalty param.Field[float64]                                                                         `json:"repetition_penalty"`
	ResponseFormat    param.Field[AccountAIRunCfFblgitExecuteUnaCybertron7bV2Bf16ParamsBodyMessagesResponseFormat] `json:"response_format"`
	// Random seed for reproducibility of the generation.
	Seed param.Field[int64] `json:"seed"`
	// If true, the response will be streamed back incrementally using SSE, Server Sent
	// Events.
	Stream param.Field[bool] `json:"stream"`
	// Controls the randomness of the output; higher values produce more random
	// results.
	Temperature param.Field[float64] `json:"temperature"`
	// A list of tools available for the assistant to use.
	Tools param.Field[[]AccountAIRunCfFblgitExecuteUnaCybertron7bV2Bf16ParamsBodyMessagesToolUnion] `json:"tools"`
	// Limits the AI to choose from the top 'k' most probable words. Lower values make
	// responses more focused; higher values introduce more variety and potential
	// surprises.
	TopK param.Field[int64] `json:"top_k"`
	// Adjusts the creativity of the AI's responses by controlling how many possible
	// words it considers. Lower values make outputs more predictable; higher values
	// allow for more varied and creative responses.
	TopP param.Field[float64] `json:"top_p"`
}

func (r AccountAIRunCfFblgitExecuteUnaCybertron7bV2Bf16ParamsBodyMessages) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfFblgitExecuteUnaCybertron7bV2Bf16ParamsBodyMessages) implementsAccountAIRunCfFblgitExecuteUnaCybertron7bV2Bf16ParamsBodyUnion() {
}

type AccountAIRunCfFblgitExecuteUnaCybertron7bV2Bf16ParamsBodyMessagesMessage struct {
	// The content of the message as a string.
	Content param.Field[string] `json:"content,required"`
	// The role of the message sender (e.g., 'user', 'assistant', 'system', 'tool').
	Role param.Field[string] `json:"role,required"`
}

func (r AccountAIRunCfFblgitExecuteUnaCybertron7bV2Bf16ParamsBodyMessagesMessage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfFblgitExecuteUnaCybertron7bV2Bf16ParamsBodyMessagesFunction struct {
	Code param.Field[string] `json:"code,required"`
	Name param.Field[string] `json:"name,required"`
}

func (r AccountAIRunCfFblgitExecuteUnaCybertron7bV2Bf16ParamsBodyMessagesFunction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfFblgitExecuteUnaCybertron7bV2Bf16ParamsBodyMessagesResponseFormat struct {
	JsonSchema param.Field[interface{}]                                                                         `json:"json_schema"`
	Type       param.Field[AccountAIRunCfFblgitExecuteUnaCybertron7bV2Bf16ParamsBodyMessagesResponseFormatType] `json:"type"`
}

func (r AccountAIRunCfFblgitExecuteUnaCybertron7bV2Bf16ParamsBodyMessagesResponseFormat) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfFblgitExecuteUnaCybertron7bV2Bf16ParamsBodyMessagesResponseFormatType string

const (
	AccountAIRunCfFblgitExecuteUnaCybertron7bV2Bf16ParamsBodyMessagesResponseFormatTypeJsonObject AccountAIRunCfFblgitExecuteUnaCybertron7bV2Bf16ParamsBodyMessagesResponseFormatType = "json_object"
	AccountAIRunCfFblgitExecuteUnaCybertron7bV2Bf16ParamsBodyMessagesResponseFormatTypeJsonSchema AccountAIRunCfFblgitExecuteUnaCybertron7bV2Bf16ParamsBodyMessagesResponseFormatType = "json_schema"
)

func (r AccountAIRunCfFblgitExecuteUnaCybertron7bV2Bf16ParamsBodyMessagesResponseFormatType) IsKnown() bool {
	switch r {
	case AccountAIRunCfFblgitExecuteUnaCybertron7bV2Bf16ParamsBodyMessagesResponseFormatTypeJsonObject, AccountAIRunCfFblgitExecuteUnaCybertron7bV2Bf16ParamsBodyMessagesResponseFormatTypeJsonSchema:
		return true
	}
	return false
}

type AccountAIRunCfFblgitExecuteUnaCybertron7bV2Bf16ParamsBodyMessagesTool struct {
	// A brief description of what the tool does.
	Description param.Field[string]      `json:"description"`
	Function    param.Field[interface{}] `json:"function"`
	// The name of the tool. More descriptive the better.
	Name       param.Field[string]      `json:"name"`
	Parameters param.Field[interface{}] `json:"parameters"`
	// Specifies the type of tool (e.g., 'function').
	Type param.Field[string] `json:"type"`
}

func (r AccountAIRunCfFblgitExecuteUnaCybertron7bV2Bf16ParamsBodyMessagesTool) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfFblgitExecuteUnaCybertron7bV2Bf16ParamsBodyMessagesTool) implementsAccountAIRunCfFblgitExecuteUnaCybertron7bV2Bf16ParamsBodyMessagesToolUnion() {
}

// Satisfied by
// [AccountAIRunCfFblgitExecuteUnaCybertron7bV2Bf16ParamsBodyMessagesToolsObject],
// [AccountAIRunCfFblgitExecuteUnaCybertron7bV2Bf16ParamsBodyMessagesToolsObject],
// [AccountAIRunCfFblgitExecuteUnaCybertron7bV2Bf16ParamsBodyMessagesTool].
type AccountAIRunCfFblgitExecuteUnaCybertron7bV2Bf16ParamsBodyMessagesToolUnion interface {
	implementsAccountAIRunCfFblgitExecuteUnaCybertron7bV2Bf16ParamsBodyMessagesToolUnion()
}

type AccountAIRunCfFblgitExecuteUnaCybertron7bV2Bf16ParamsBodyMessagesToolsObject struct {
	// A brief description of what the tool does.
	Description param.Field[string] `json:"description,required"`
	// The name of the tool. More descriptive the better.
	Name param.Field[string] `json:"name,required"`
	// Schema defining the parameters accepted by the tool.
	Parameters param.Field[AccountAIRunCfFblgitExecuteUnaCybertron7bV2Bf16ParamsBodyMessagesToolsObjectParameters] `json:"parameters,required"`
}

func (r AccountAIRunCfFblgitExecuteUnaCybertron7bV2Bf16ParamsBodyMessagesToolsObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfFblgitExecuteUnaCybertron7bV2Bf16ParamsBodyMessagesToolsObject) implementsAccountAIRunCfFblgitExecuteUnaCybertron7bV2Bf16ParamsBodyMessagesToolUnion() {
}

// Schema defining the parameters accepted by the tool.
type AccountAIRunCfFblgitExecuteUnaCybertron7bV2Bf16ParamsBodyMessagesToolsObjectParameters struct {
	// Definitions of each parameter.
	Properties param.Field[map[string]AccountAIRunCfFblgitExecuteUnaCybertron7bV2Bf16ParamsBodyMessagesToolsObjectParametersProperties] `json:"properties,required"`
	// The type of the parameters object (usually 'object').
	Type param.Field[string] `json:"type,required"`
	// List of required parameter names.
	Required param.Field[[]string] `json:"required"`
}

func (r AccountAIRunCfFblgitExecuteUnaCybertron7bV2Bf16ParamsBodyMessagesToolsObjectParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfFblgitExecuteUnaCybertron7bV2Bf16ParamsBodyMessagesToolsObjectParametersProperties struct {
	// A description of the expected parameter.
	Description param.Field[string] `json:"description,required"`
	// The data type of the parameter.
	Type param.Field[string] `json:"type,required"`
}

func (r AccountAIRunCfFblgitExecuteUnaCybertron7bV2Bf16ParamsBodyMessagesToolsObjectParametersProperties) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
