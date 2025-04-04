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

// AccountAIRunHfNousresearchService contains methods and other services that help
// with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountAIRunHfNousresearchService] method instead.
type AccountAIRunHfNousresearchService struct {
	Options []option.RequestOption
}

// NewAccountAIRunHfNousresearchService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountAIRunHfNousresearchService(opts ...option.RequestOption) (r *AccountAIRunHfNousresearchService) {
	r = &AccountAIRunHfNousresearchService{}
	r.Options = opts
	return
}

// Execute @hf/nousresearch/hermes-2-pro-mistral-7b model.
func (r *AccountAIRunHfNousresearchService) ExecuteHermes2ProMistral7b(ctx context.Context, accountID string, params AccountAIRunHfNousresearchExecuteHermes2ProMistral7bParams, opts ...option.RequestOption) (res *AccountAIRunHfNousresearchExecuteHermes2ProMistral7bResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai/run/@hf/nousresearch/hermes-2-pro-mistral-7b", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type AccountAIRunHfNousresearchExecuteHermes2ProMistral7bResponse struct {
	Result  AccountAIRunHfNousresearchExecuteHermes2ProMistral7bResponseResultUnion `json:"result" format:"binary"`
	Success bool                                                                    `json:"success"`
	JSON    accountAIRunHfNousresearchExecuteHermes2ProMistral7bResponseJSON        `json:"-"`
}

// accountAIRunHfNousresearchExecuteHermes2ProMistral7bResponseJSON contains the
// JSON metadata for the struct
// [AccountAIRunHfNousresearchExecuteHermes2ProMistral7bResponse]
type accountAIRunHfNousresearchExecuteHermes2ProMistral7bResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIRunHfNousresearchExecuteHermes2ProMistral7bResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunHfNousresearchExecuteHermes2ProMistral7bResponseJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by
// [AccountAIRunHfNousresearchExecuteHermes2ProMistral7bResponseResultObject] or
// [shared.UnionString].
type AccountAIRunHfNousresearchExecuteHermes2ProMistral7bResponseResultUnion interface {
	ImplementsAccountAIRunHfNousresearchExecuteHermes2ProMistral7bResponseResultUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccountAIRunHfNousresearchExecuteHermes2ProMistral7bResponseResultUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccountAIRunHfNousresearchExecuteHermes2ProMistral7bResponseResultObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type AccountAIRunHfNousresearchExecuteHermes2ProMistral7bResponseResultObject struct {
	// The generated text response from the model
	Response string `json:"response,required"`
	// An array of tool calls requests made during the response generation
	ToolCalls []AccountAIRunHfNousresearchExecuteHermes2ProMistral7bResponseResultObjectToolCall `json:"tool_calls"`
	// Usage statistics for the inference request
	Usage AccountAIRunHfNousresearchExecuteHermes2ProMistral7bResponseResultObjectUsage `json:"usage"`
	JSON  accountAIRunHfNousresearchExecuteHermes2ProMistral7bResponseResultObjectJSON  `json:"-"`
}

// accountAIRunHfNousresearchExecuteHermes2ProMistral7bResponseResultObjectJSON
// contains the JSON metadata for the struct
// [AccountAIRunHfNousresearchExecuteHermes2ProMistral7bResponseResultObject]
type accountAIRunHfNousresearchExecuteHermes2ProMistral7bResponseResultObjectJSON struct {
	Response    apijson.Field
	ToolCalls   apijson.Field
	Usage       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIRunHfNousresearchExecuteHermes2ProMistral7bResponseResultObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunHfNousresearchExecuteHermes2ProMistral7bResponseResultObjectJSON) RawJSON() string {
	return r.raw
}

func (r AccountAIRunHfNousresearchExecuteHermes2ProMistral7bResponseResultObject) ImplementsAccountAIRunHfNousresearchExecuteHermes2ProMistral7bResponseResultUnion() {
}

type AccountAIRunHfNousresearchExecuteHermes2ProMistral7bResponseResultObjectToolCall struct {
	// The arguments passed to be passed to the tool call request
	Arguments interface{} `json:"arguments"`
	// The name of the tool to be called
	Name string                                                                               `json:"name"`
	JSON accountAIRunHfNousresearchExecuteHermes2ProMistral7bResponseResultObjectToolCallJSON `json:"-"`
}

// accountAIRunHfNousresearchExecuteHermes2ProMistral7bResponseResultObjectToolCallJSON
// contains the JSON metadata for the struct
// [AccountAIRunHfNousresearchExecuteHermes2ProMistral7bResponseResultObjectToolCall]
type accountAIRunHfNousresearchExecuteHermes2ProMistral7bResponseResultObjectToolCallJSON struct {
	Arguments   apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIRunHfNousresearchExecuteHermes2ProMistral7bResponseResultObjectToolCall) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunHfNousresearchExecuteHermes2ProMistral7bResponseResultObjectToolCallJSON) RawJSON() string {
	return r.raw
}

// Usage statistics for the inference request
type AccountAIRunHfNousresearchExecuteHermes2ProMistral7bResponseResultObjectUsage struct {
	// Total number of tokens in output
	CompletionTokens float64 `json:"completion_tokens"`
	// Total number of tokens in input
	PromptTokens float64 `json:"prompt_tokens"`
	// Total number of input and output tokens
	TotalTokens float64                                                                           `json:"total_tokens"`
	JSON        accountAIRunHfNousresearchExecuteHermes2ProMistral7bResponseResultObjectUsageJSON `json:"-"`
}

// accountAIRunHfNousresearchExecuteHermes2ProMistral7bResponseResultObjectUsageJSON
// contains the JSON metadata for the struct
// [AccountAIRunHfNousresearchExecuteHermes2ProMistral7bResponseResultObjectUsage]
type accountAIRunHfNousresearchExecuteHermes2ProMistral7bResponseResultObjectUsageJSON struct {
	CompletionTokens apijson.Field
	PromptTokens     apijson.Field
	TotalTokens      apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountAIRunHfNousresearchExecuteHermes2ProMistral7bResponseResultObjectUsage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunHfNousresearchExecuteHermes2ProMistral7bResponseResultObjectUsageJSON) RawJSON() string {
	return r.raw
}

type AccountAIRunHfNousresearchExecuteHermes2ProMistral7bParams struct {
	QueueRequest param.Field[string]                                                 `query:"queueRequest"`
	Body         AccountAIRunHfNousresearchExecuteHermes2ProMistral7bParamsBodyUnion `json:"body"`
}

func (r AccountAIRunHfNousresearchExecuteHermes2ProMistral7bParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

// URLQuery serializes
// [AccountAIRunHfNousresearchExecuteHermes2ProMistral7bParams]'s query parameters
// as `url.Values`.
func (r AccountAIRunHfNousresearchExecuteHermes2ProMistral7bParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountAIRunHfNousresearchExecuteHermes2ProMistral7bParamsBody struct {
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

func (r AccountAIRunHfNousresearchExecuteHermes2ProMistral7bParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunHfNousresearchExecuteHermes2ProMistral7bParamsBody) implementsAccountAIRunHfNousresearchExecuteHermes2ProMistral7bParamsBodyUnion() {
}

// Satisfied by
// [AccountAIRunHfNousresearchExecuteHermes2ProMistral7bParamsBodyPrompt],
// [AccountAIRunHfNousresearchExecuteHermes2ProMistral7bParamsBodyMessages],
// [AccountAIRunHfNousresearchExecuteHermes2ProMistral7bParamsBody].
type AccountAIRunHfNousresearchExecuteHermes2ProMistral7bParamsBodyUnion interface {
	implementsAccountAIRunHfNousresearchExecuteHermes2ProMistral7bParamsBodyUnion()
}

type AccountAIRunHfNousresearchExecuteHermes2ProMistral7bParamsBodyPrompt struct {
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
	RepetitionPenalty param.Field[float64]                                                                            `json:"repetition_penalty"`
	ResponseFormat    param.Field[AccountAIRunHfNousresearchExecuteHermes2ProMistral7bParamsBodyPromptResponseFormat] `json:"response_format"`
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

func (r AccountAIRunHfNousresearchExecuteHermes2ProMistral7bParamsBodyPrompt) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunHfNousresearchExecuteHermes2ProMistral7bParamsBodyPrompt) implementsAccountAIRunHfNousresearchExecuteHermes2ProMistral7bParamsBodyUnion() {
}

type AccountAIRunHfNousresearchExecuteHermes2ProMistral7bParamsBodyPromptResponseFormat struct {
	JsonSchema param.Field[interface{}]                                                                            `json:"json_schema"`
	Type       param.Field[AccountAIRunHfNousresearchExecuteHermes2ProMistral7bParamsBodyPromptResponseFormatType] `json:"type"`
}

func (r AccountAIRunHfNousresearchExecuteHermes2ProMistral7bParamsBodyPromptResponseFormat) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunHfNousresearchExecuteHermes2ProMistral7bParamsBodyPromptResponseFormatType string

const (
	AccountAIRunHfNousresearchExecuteHermes2ProMistral7bParamsBodyPromptResponseFormatTypeJsonObject AccountAIRunHfNousresearchExecuteHermes2ProMistral7bParamsBodyPromptResponseFormatType = "json_object"
	AccountAIRunHfNousresearchExecuteHermes2ProMistral7bParamsBodyPromptResponseFormatTypeJsonSchema AccountAIRunHfNousresearchExecuteHermes2ProMistral7bParamsBodyPromptResponseFormatType = "json_schema"
)

func (r AccountAIRunHfNousresearchExecuteHermes2ProMistral7bParamsBodyPromptResponseFormatType) IsKnown() bool {
	switch r {
	case AccountAIRunHfNousresearchExecuteHermes2ProMistral7bParamsBodyPromptResponseFormatTypeJsonObject, AccountAIRunHfNousresearchExecuteHermes2ProMistral7bParamsBodyPromptResponseFormatTypeJsonSchema:
		return true
	}
	return false
}

type AccountAIRunHfNousresearchExecuteHermes2ProMistral7bParamsBodyMessages struct {
	// An array of message objects representing the conversation history.
	Messages param.Field[[]AccountAIRunHfNousresearchExecuteHermes2ProMistral7bParamsBodyMessagesMessage] `json:"messages,required"`
	// Decreases the likelihood of the model repeating the same lines verbatim.
	FrequencyPenalty param.Field[float64]                                                                          `json:"frequency_penalty"`
	Functions        param.Field[[]AccountAIRunHfNousresearchExecuteHermes2ProMistral7bParamsBodyMessagesFunction] `json:"functions"`
	// The maximum number of tokens to generate in the response.
	MaxTokens param.Field[int64] `json:"max_tokens"`
	// Increases the likelihood of the model introducing new topics.
	PresencePenalty param.Field[float64] `json:"presence_penalty"`
	// If true, a chat template is not applied and you must adhere to the specific
	// model's expected formatting.
	Raw param.Field[bool] `json:"raw"`
	// Penalty for repeated tokens; higher values discourage repetition.
	RepetitionPenalty param.Field[float64]                                                                              `json:"repetition_penalty"`
	ResponseFormat    param.Field[AccountAIRunHfNousresearchExecuteHermes2ProMistral7bParamsBodyMessagesResponseFormat] `json:"response_format"`
	// Random seed for reproducibility of the generation.
	Seed param.Field[int64] `json:"seed"`
	// If true, the response will be streamed back incrementally using SSE, Server Sent
	// Events.
	Stream param.Field[bool] `json:"stream"`
	// Controls the randomness of the output; higher values produce more random
	// results.
	Temperature param.Field[float64] `json:"temperature"`
	// A list of tools available for the assistant to use.
	Tools param.Field[[]AccountAIRunHfNousresearchExecuteHermes2ProMistral7bParamsBodyMessagesToolUnion] `json:"tools"`
	// Limits the AI to choose from the top 'k' most probable words. Lower values make
	// responses more focused; higher values introduce more variety and potential
	// surprises.
	TopK param.Field[int64] `json:"top_k"`
	// Adjusts the creativity of the AI's responses by controlling how many possible
	// words it considers. Lower values make outputs more predictable; higher values
	// allow for more varied and creative responses.
	TopP param.Field[float64] `json:"top_p"`
}

func (r AccountAIRunHfNousresearchExecuteHermes2ProMistral7bParamsBodyMessages) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunHfNousresearchExecuteHermes2ProMistral7bParamsBodyMessages) implementsAccountAIRunHfNousresearchExecuteHermes2ProMistral7bParamsBodyUnion() {
}

type AccountAIRunHfNousresearchExecuteHermes2ProMistral7bParamsBodyMessagesMessage struct {
	// The content of the message as a string.
	Content param.Field[string] `json:"content,required"`
	// The role of the message sender (e.g., 'user', 'assistant', 'system', 'tool').
	Role param.Field[string] `json:"role,required"`
}

func (r AccountAIRunHfNousresearchExecuteHermes2ProMistral7bParamsBodyMessagesMessage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunHfNousresearchExecuteHermes2ProMistral7bParamsBodyMessagesFunction struct {
	Code param.Field[string] `json:"code,required"`
	Name param.Field[string] `json:"name,required"`
}

func (r AccountAIRunHfNousresearchExecuteHermes2ProMistral7bParamsBodyMessagesFunction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunHfNousresearchExecuteHermes2ProMistral7bParamsBodyMessagesResponseFormat struct {
	JsonSchema param.Field[interface{}]                                                                              `json:"json_schema"`
	Type       param.Field[AccountAIRunHfNousresearchExecuteHermes2ProMistral7bParamsBodyMessagesResponseFormatType] `json:"type"`
}

func (r AccountAIRunHfNousresearchExecuteHermes2ProMistral7bParamsBodyMessagesResponseFormat) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunHfNousresearchExecuteHermes2ProMistral7bParamsBodyMessagesResponseFormatType string

const (
	AccountAIRunHfNousresearchExecuteHermes2ProMistral7bParamsBodyMessagesResponseFormatTypeJsonObject AccountAIRunHfNousresearchExecuteHermes2ProMistral7bParamsBodyMessagesResponseFormatType = "json_object"
	AccountAIRunHfNousresearchExecuteHermes2ProMistral7bParamsBodyMessagesResponseFormatTypeJsonSchema AccountAIRunHfNousresearchExecuteHermes2ProMistral7bParamsBodyMessagesResponseFormatType = "json_schema"
)

func (r AccountAIRunHfNousresearchExecuteHermes2ProMistral7bParamsBodyMessagesResponseFormatType) IsKnown() bool {
	switch r {
	case AccountAIRunHfNousresearchExecuteHermes2ProMistral7bParamsBodyMessagesResponseFormatTypeJsonObject, AccountAIRunHfNousresearchExecuteHermes2ProMistral7bParamsBodyMessagesResponseFormatTypeJsonSchema:
		return true
	}
	return false
}

type AccountAIRunHfNousresearchExecuteHermes2ProMistral7bParamsBodyMessagesTool struct {
	// A brief description of what the tool does.
	Description param.Field[string]      `json:"description"`
	Function    param.Field[interface{}] `json:"function"`
	// The name of the tool. More descriptive the better.
	Name       param.Field[string]      `json:"name"`
	Parameters param.Field[interface{}] `json:"parameters"`
	// Specifies the type of tool (e.g., 'function').
	Type param.Field[string] `json:"type"`
}

func (r AccountAIRunHfNousresearchExecuteHermes2ProMistral7bParamsBodyMessagesTool) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunHfNousresearchExecuteHermes2ProMistral7bParamsBodyMessagesTool) implementsAccountAIRunHfNousresearchExecuteHermes2ProMistral7bParamsBodyMessagesToolUnion() {
}

// Satisfied by
// [AccountAIRunHfNousresearchExecuteHermes2ProMistral7bParamsBodyMessagesToolsObject],
// [AccountAIRunHfNousresearchExecuteHermes2ProMistral7bParamsBodyMessagesToolsObject],
// [AccountAIRunHfNousresearchExecuteHermes2ProMistral7bParamsBodyMessagesTool].
type AccountAIRunHfNousresearchExecuteHermes2ProMistral7bParamsBodyMessagesToolUnion interface {
	implementsAccountAIRunHfNousresearchExecuteHermes2ProMistral7bParamsBodyMessagesToolUnion()
}

type AccountAIRunHfNousresearchExecuteHermes2ProMistral7bParamsBodyMessagesToolsObject struct {
	// A brief description of what the tool does.
	Description param.Field[string] `json:"description,required"`
	// The name of the tool. More descriptive the better.
	Name param.Field[string] `json:"name,required"`
	// Schema defining the parameters accepted by the tool.
	Parameters param.Field[AccountAIRunHfNousresearchExecuteHermes2ProMistral7bParamsBodyMessagesToolsObjectParameters] `json:"parameters,required"`
}

func (r AccountAIRunHfNousresearchExecuteHermes2ProMistral7bParamsBodyMessagesToolsObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunHfNousresearchExecuteHermes2ProMistral7bParamsBodyMessagesToolsObject) implementsAccountAIRunHfNousresearchExecuteHermes2ProMistral7bParamsBodyMessagesToolUnion() {
}

// Schema defining the parameters accepted by the tool.
type AccountAIRunHfNousresearchExecuteHermes2ProMistral7bParamsBodyMessagesToolsObjectParameters struct {
	// Definitions of each parameter.
	Properties param.Field[map[string]AccountAIRunHfNousresearchExecuteHermes2ProMistral7bParamsBodyMessagesToolsObjectParametersProperties] `json:"properties,required"`
	// The type of the parameters object (usually 'object').
	Type param.Field[string] `json:"type,required"`
	// List of required parameter names.
	Required param.Field[[]string] `json:"required"`
}

func (r AccountAIRunHfNousresearchExecuteHermes2ProMistral7bParamsBodyMessagesToolsObjectParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunHfNousresearchExecuteHermes2ProMistral7bParamsBodyMessagesToolsObjectParametersProperties struct {
	// A description of the expected parameter.
	Description param.Field[string] `json:"description,required"`
	// The data type of the parameter.
	Type param.Field[string] `json:"type,required"`
}

func (r AccountAIRunHfNousresearchExecuteHermes2ProMistral7bParamsBodyMessagesToolsObjectParametersProperties) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
