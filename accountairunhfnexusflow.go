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

// AccountAIRunHfNexusflowService contains methods and other services that help
// with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountAIRunHfNexusflowService] method instead.
type AccountAIRunHfNexusflowService struct {
	Options []option.RequestOption
}

// NewAccountAIRunHfNexusflowService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountAIRunHfNexusflowService(opts ...option.RequestOption) (r *AccountAIRunHfNexusflowService) {
	r = &AccountAIRunHfNexusflowService{}
	r.Options = opts
	return
}

// Execute @hf/nexusflow/starling-lm-7b-beta model.
func (r *AccountAIRunHfNexusflowService) ExecuteStarlingLm7bBeta(ctx context.Context, accountID string, params AccountAIRunHfNexusflowExecuteStarlingLm7bBetaParams, opts ...option.RequestOption) (res *AccountAIRunHfNexusflowExecuteStarlingLm7bBetaResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai/run/@hf/nexusflow/starling-lm-7b-beta", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type AccountAIRunHfNexusflowExecuteStarlingLm7bBetaResponse struct {
	Result  AccountAIRunHfNexusflowExecuteStarlingLm7bBetaResponseResultUnion `json:"result" format:"binary"`
	Success bool                                                              `json:"success"`
	JSON    accountAIRunHfNexusflowExecuteStarlingLm7bBetaResponseJSON        `json:"-"`
}

// accountAIRunHfNexusflowExecuteStarlingLm7bBetaResponseJSON contains the JSON
// metadata for the struct [AccountAIRunHfNexusflowExecuteStarlingLm7bBetaResponse]
type accountAIRunHfNexusflowExecuteStarlingLm7bBetaResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIRunHfNexusflowExecuteStarlingLm7bBetaResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunHfNexusflowExecuteStarlingLm7bBetaResponseJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by
// [AccountAIRunHfNexusflowExecuteStarlingLm7bBetaResponseResultObject] or
// [shared.UnionString].
type AccountAIRunHfNexusflowExecuteStarlingLm7bBetaResponseResultUnion interface {
	ImplementsAccountAIRunHfNexusflowExecuteStarlingLm7bBetaResponseResultUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccountAIRunHfNexusflowExecuteStarlingLm7bBetaResponseResultUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccountAIRunHfNexusflowExecuteStarlingLm7bBetaResponseResultObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type AccountAIRunHfNexusflowExecuteStarlingLm7bBetaResponseResultObject struct {
	// The generated text response from the model
	Response string `json:"response,required"`
	// An array of tool calls requests made during the response generation
	ToolCalls []AccountAIRunHfNexusflowExecuteStarlingLm7bBetaResponseResultObjectToolCall `json:"tool_calls"`
	// Usage statistics for the inference request
	Usage AccountAIRunHfNexusflowExecuteStarlingLm7bBetaResponseResultObjectUsage `json:"usage"`
	JSON  accountAIRunHfNexusflowExecuteStarlingLm7bBetaResponseResultObjectJSON  `json:"-"`
}

// accountAIRunHfNexusflowExecuteStarlingLm7bBetaResponseResultObjectJSON contains
// the JSON metadata for the struct
// [AccountAIRunHfNexusflowExecuteStarlingLm7bBetaResponseResultObject]
type accountAIRunHfNexusflowExecuteStarlingLm7bBetaResponseResultObjectJSON struct {
	Response    apijson.Field
	ToolCalls   apijson.Field
	Usage       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIRunHfNexusflowExecuteStarlingLm7bBetaResponseResultObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunHfNexusflowExecuteStarlingLm7bBetaResponseResultObjectJSON) RawJSON() string {
	return r.raw
}

func (r AccountAIRunHfNexusflowExecuteStarlingLm7bBetaResponseResultObject) ImplementsAccountAIRunHfNexusflowExecuteStarlingLm7bBetaResponseResultUnion() {
}

type AccountAIRunHfNexusflowExecuteStarlingLm7bBetaResponseResultObjectToolCall struct {
	// The arguments passed to be passed to the tool call request
	Arguments interface{} `json:"arguments"`
	// The name of the tool to be called
	Name string                                                                         `json:"name"`
	JSON accountAIRunHfNexusflowExecuteStarlingLm7bBetaResponseResultObjectToolCallJSON `json:"-"`
}

// accountAIRunHfNexusflowExecuteStarlingLm7bBetaResponseResultObjectToolCallJSON
// contains the JSON metadata for the struct
// [AccountAIRunHfNexusflowExecuteStarlingLm7bBetaResponseResultObjectToolCall]
type accountAIRunHfNexusflowExecuteStarlingLm7bBetaResponseResultObjectToolCallJSON struct {
	Arguments   apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIRunHfNexusflowExecuteStarlingLm7bBetaResponseResultObjectToolCall) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunHfNexusflowExecuteStarlingLm7bBetaResponseResultObjectToolCallJSON) RawJSON() string {
	return r.raw
}

// Usage statistics for the inference request
type AccountAIRunHfNexusflowExecuteStarlingLm7bBetaResponseResultObjectUsage struct {
	// Total number of tokens in output
	CompletionTokens float64 `json:"completion_tokens"`
	// Total number of tokens in input
	PromptTokens float64 `json:"prompt_tokens"`
	// Total number of input and output tokens
	TotalTokens float64                                                                     `json:"total_tokens"`
	JSON        accountAIRunHfNexusflowExecuteStarlingLm7bBetaResponseResultObjectUsageJSON `json:"-"`
}

// accountAIRunHfNexusflowExecuteStarlingLm7bBetaResponseResultObjectUsageJSON
// contains the JSON metadata for the struct
// [AccountAIRunHfNexusflowExecuteStarlingLm7bBetaResponseResultObjectUsage]
type accountAIRunHfNexusflowExecuteStarlingLm7bBetaResponseResultObjectUsageJSON struct {
	CompletionTokens apijson.Field
	PromptTokens     apijson.Field
	TotalTokens      apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountAIRunHfNexusflowExecuteStarlingLm7bBetaResponseResultObjectUsage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunHfNexusflowExecuteStarlingLm7bBetaResponseResultObjectUsageJSON) RawJSON() string {
	return r.raw
}

type AccountAIRunHfNexusflowExecuteStarlingLm7bBetaParams struct {
	QueueRequest param.Field[string]                                           `query:"queueRequest"`
	Body         AccountAIRunHfNexusflowExecuteStarlingLm7bBetaParamsBodyUnion `json:"body"`
}

func (r AccountAIRunHfNexusflowExecuteStarlingLm7bBetaParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

// URLQuery serializes [AccountAIRunHfNexusflowExecuteStarlingLm7bBetaParams]'s
// query parameters as `url.Values`.
func (r AccountAIRunHfNexusflowExecuteStarlingLm7bBetaParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountAIRunHfNexusflowExecuteStarlingLm7bBetaParamsBody struct {
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

func (r AccountAIRunHfNexusflowExecuteStarlingLm7bBetaParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunHfNexusflowExecuteStarlingLm7bBetaParamsBody) implementsAccountAIRunHfNexusflowExecuteStarlingLm7bBetaParamsBodyUnion() {
}

// Satisfied by [AccountAIRunHfNexusflowExecuteStarlingLm7bBetaParamsBodyPrompt],
// [AccountAIRunHfNexusflowExecuteStarlingLm7bBetaParamsBodyMessages],
// [AccountAIRunHfNexusflowExecuteStarlingLm7bBetaParamsBody].
type AccountAIRunHfNexusflowExecuteStarlingLm7bBetaParamsBodyUnion interface {
	implementsAccountAIRunHfNexusflowExecuteStarlingLm7bBetaParamsBodyUnion()
}

type AccountAIRunHfNexusflowExecuteStarlingLm7bBetaParamsBodyPrompt struct {
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
	ResponseFormat    param.Field[AccountAIRunHfNexusflowExecuteStarlingLm7bBetaParamsBodyPromptResponseFormat] `json:"response_format"`
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

func (r AccountAIRunHfNexusflowExecuteStarlingLm7bBetaParamsBodyPrompt) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunHfNexusflowExecuteStarlingLm7bBetaParamsBodyPrompt) implementsAccountAIRunHfNexusflowExecuteStarlingLm7bBetaParamsBodyUnion() {
}

type AccountAIRunHfNexusflowExecuteStarlingLm7bBetaParamsBodyPromptResponseFormat struct {
	JsonSchema param.Field[interface{}]                                                                      `json:"json_schema"`
	Type       param.Field[AccountAIRunHfNexusflowExecuteStarlingLm7bBetaParamsBodyPromptResponseFormatType] `json:"type"`
}

func (r AccountAIRunHfNexusflowExecuteStarlingLm7bBetaParamsBodyPromptResponseFormat) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunHfNexusflowExecuteStarlingLm7bBetaParamsBodyPromptResponseFormatType string

const (
	AccountAIRunHfNexusflowExecuteStarlingLm7bBetaParamsBodyPromptResponseFormatTypeJsonObject AccountAIRunHfNexusflowExecuteStarlingLm7bBetaParamsBodyPromptResponseFormatType = "json_object"
	AccountAIRunHfNexusflowExecuteStarlingLm7bBetaParamsBodyPromptResponseFormatTypeJsonSchema AccountAIRunHfNexusflowExecuteStarlingLm7bBetaParamsBodyPromptResponseFormatType = "json_schema"
)

func (r AccountAIRunHfNexusflowExecuteStarlingLm7bBetaParamsBodyPromptResponseFormatType) IsKnown() bool {
	switch r {
	case AccountAIRunHfNexusflowExecuteStarlingLm7bBetaParamsBodyPromptResponseFormatTypeJsonObject, AccountAIRunHfNexusflowExecuteStarlingLm7bBetaParamsBodyPromptResponseFormatTypeJsonSchema:
		return true
	}
	return false
}

type AccountAIRunHfNexusflowExecuteStarlingLm7bBetaParamsBodyMessages struct {
	// An array of message objects representing the conversation history.
	Messages param.Field[[]AccountAIRunHfNexusflowExecuteStarlingLm7bBetaParamsBodyMessagesMessage] `json:"messages,required"`
	// Decreases the likelihood of the model repeating the same lines verbatim.
	FrequencyPenalty param.Field[float64]                                                                    `json:"frequency_penalty"`
	Functions        param.Field[[]AccountAIRunHfNexusflowExecuteStarlingLm7bBetaParamsBodyMessagesFunction] `json:"functions"`
	// The maximum number of tokens to generate in the response.
	MaxTokens param.Field[int64] `json:"max_tokens"`
	// Increases the likelihood of the model introducing new topics.
	PresencePenalty param.Field[float64] `json:"presence_penalty"`
	// If true, a chat template is not applied and you must adhere to the specific
	// model's expected formatting.
	Raw param.Field[bool] `json:"raw"`
	// Penalty for repeated tokens; higher values discourage repetition.
	RepetitionPenalty param.Field[float64]                                                                        `json:"repetition_penalty"`
	ResponseFormat    param.Field[AccountAIRunHfNexusflowExecuteStarlingLm7bBetaParamsBodyMessagesResponseFormat] `json:"response_format"`
	// Random seed for reproducibility of the generation.
	Seed param.Field[int64] `json:"seed"`
	// If true, the response will be streamed back incrementally using SSE, Server Sent
	// Events.
	Stream param.Field[bool] `json:"stream"`
	// Controls the randomness of the output; higher values produce more random
	// results.
	Temperature param.Field[float64] `json:"temperature"`
	// A list of tools available for the assistant to use.
	Tools param.Field[[]AccountAIRunHfNexusflowExecuteStarlingLm7bBetaParamsBodyMessagesToolUnion] `json:"tools"`
	// Limits the AI to choose from the top 'k' most probable words. Lower values make
	// responses more focused; higher values introduce more variety and potential
	// surprises.
	TopK param.Field[int64] `json:"top_k"`
	// Adjusts the creativity of the AI's responses by controlling how many possible
	// words it considers. Lower values make outputs more predictable; higher values
	// allow for more varied and creative responses.
	TopP param.Field[float64] `json:"top_p"`
}

func (r AccountAIRunHfNexusflowExecuteStarlingLm7bBetaParamsBodyMessages) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunHfNexusflowExecuteStarlingLm7bBetaParamsBodyMessages) implementsAccountAIRunHfNexusflowExecuteStarlingLm7bBetaParamsBodyUnion() {
}

type AccountAIRunHfNexusflowExecuteStarlingLm7bBetaParamsBodyMessagesMessage struct {
	// The content of the message as a string.
	Content param.Field[string] `json:"content,required"`
	// The role of the message sender (e.g., 'user', 'assistant', 'system', 'tool').
	Role param.Field[string] `json:"role,required"`
}

func (r AccountAIRunHfNexusflowExecuteStarlingLm7bBetaParamsBodyMessagesMessage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunHfNexusflowExecuteStarlingLm7bBetaParamsBodyMessagesFunction struct {
	Code param.Field[string] `json:"code,required"`
	Name param.Field[string] `json:"name,required"`
}

func (r AccountAIRunHfNexusflowExecuteStarlingLm7bBetaParamsBodyMessagesFunction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunHfNexusflowExecuteStarlingLm7bBetaParamsBodyMessagesResponseFormat struct {
	JsonSchema param.Field[interface{}]                                                                        `json:"json_schema"`
	Type       param.Field[AccountAIRunHfNexusflowExecuteStarlingLm7bBetaParamsBodyMessagesResponseFormatType] `json:"type"`
}

func (r AccountAIRunHfNexusflowExecuteStarlingLm7bBetaParamsBodyMessagesResponseFormat) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunHfNexusflowExecuteStarlingLm7bBetaParamsBodyMessagesResponseFormatType string

const (
	AccountAIRunHfNexusflowExecuteStarlingLm7bBetaParamsBodyMessagesResponseFormatTypeJsonObject AccountAIRunHfNexusflowExecuteStarlingLm7bBetaParamsBodyMessagesResponseFormatType = "json_object"
	AccountAIRunHfNexusflowExecuteStarlingLm7bBetaParamsBodyMessagesResponseFormatTypeJsonSchema AccountAIRunHfNexusflowExecuteStarlingLm7bBetaParamsBodyMessagesResponseFormatType = "json_schema"
)

func (r AccountAIRunHfNexusflowExecuteStarlingLm7bBetaParamsBodyMessagesResponseFormatType) IsKnown() bool {
	switch r {
	case AccountAIRunHfNexusflowExecuteStarlingLm7bBetaParamsBodyMessagesResponseFormatTypeJsonObject, AccountAIRunHfNexusflowExecuteStarlingLm7bBetaParamsBodyMessagesResponseFormatTypeJsonSchema:
		return true
	}
	return false
}

type AccountAIRunHfNexusflowExecuteStarlingLm7bBetaParamsBodyMessagesTool struct {
	// A brief description of what the tool does.
	Description param.Field[string]      `json:"description"`
	Function    param.Field[interface{}] `json:"function"`
	// The name of the tool. More descriptive the better.
	Name       param.Field[string]      `json:"name"`
	Parameters param.Field[interface{}] `json:"parameters"`
	// Specifies the type of tool (e.g., 'function').
	Type param.Field[string] `json:"type"`
}

func (r AccountAIRunHfNexusflowExecuteStarlingLm7bBetaParamsBodyMessagesTool) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunHfNexusflowExecuteStarlingLm7bBetaParamsBodyMessagesTool) implementsAccountAIRunHfNexusflowExecuteStarlingLm7bBetaParamsBodyMessagesToolUnion() {
}

// Satisfied by
// [AccountAIRunHfNexusflowExecuteStarlingLm7bBetaParamsBodyMessagesToolsObject],
// [AccountAIRunHfNexusflowExecuteStarlingLm7bBetaParamsBodyMessagesToolsObject],
// [AccountAIRunHfNexusflowExecuteStarlingLm7bBetaParamsBodyMessagesTool].
type AccountAIRunHfNexusflowExecuteStarlingLm7bBetaParamsBodyMessagesToolUnion interface {
	implementsAccountAIRunHfNexusflowExecuteStarlingLm7bBetaParamsBodyMessagesToolUnion()
}

type AccountAIRunHfNexusflowExecuteStarlingLm7bBetaParamsBodyMessagesToolsObject struct {
	// A brief description of what the tool does.
	Description param.Field[string] `json:"description,required"`
	// The name of the tool. More descriptive the better.
	Name param.Field[string] `json:"name,required"`
	// Schema defining the parameters accepted by the tool.
	Parameters param.Field[AccountAIRunHfNexusflowExecuteStarlingLm7bBetaParamsBodyMessagesToolsObjectParameters] `json:"parameters,required"`
}

func (r AccountAIRunHfNexusflowExecuteStarlingLm7bBetaParamsBodyMessagesToolsObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunHfNexusflowExecuteStarlingLm7bBetaParamsBodyMessagesToolsObject) implementsAccountAIRunHfNexusflowExecuteStarlingLm7bBetaParamsBodyMessagesToolUnion() {
}

// Schema defining the parameters accepted by the tool.
type AccountAIRunHfNexusflowExecuteStarlingLm7bBetaParamsBodyMessagesToolsObjectParameters struct {
	// Definitions of each parameter.
	Properties param.Field[map[string]AccountAIRunHfNexusflowExecuteStarlingLm7bBetaParamsBodyMessagesToolsObjectParametersProperties] `json:"properties,required"`
	// The type of the parameters object (usually 'object').
	Type param.Field[string] `json:"type,required"`
	// List of required parameter names.
	Required param.Field[[]string] `json:"required"`
}

func (r AccountAIRunHfNexusflowExecuteStarlingLm7bBetaParamsBodyMessagesToolsObjectParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunHfNexusflowExecuteStarlingLm7bBetaParamsBodyMessagesToolsObjectParametersProperties struct {
	// A description of the expected parameter.
	Description param.Field[string] `json:"description,required"`
	// The data type of the parameter.
	Type param.Field[string] `json:"type,required"`
}

func (r AccountAIRunHfNexusflowExecuteStarlingLm7bBetaParamsBodyMessagesToolsObjectParametersProperties) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
