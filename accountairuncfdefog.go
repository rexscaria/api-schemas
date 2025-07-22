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

// AccountAIRunCfDefogService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountAIRunCfDefogService] method instead.
type AccountAIRunCfDefogService struct {
	Options []option.RequestOption
}

// NewAccountAIRunCfDefogService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountAIRunCfDefogService(opts ...option.RequestOption) (r *AccountAIRunCfDefogService) {
	r = &AccountAIRunCfDefogService{}
	r.Options = opts
	return
}

// Execute @cf/defog/sqlcoder-7b-2 model.
func (r *AccountAIRunCfDefogService) ExecuteSqlcoder7b2(ctx context.Context, accountID string, params AccountAIRunCfDefogExecuteSqlcoder7b2Params, opts ...option.RequestOption) (res *AccountAIRunCfDefogExecuteSqlcoder7b2Response, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai/run/@cf/defog/sqlcoder-7b-2", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type AccountAIRunCfDefogExecuteSqlcoder7b2Response struct {
	Result  AccountAIRunCfDefogExecuteSqlcoder7b2ResponseResultUnion `json:"result" format:"binary"`
	Success bool                                                     `json:"success"`
	JSON    accountAIRunCfDefogExecuteSqlcoder7b2ResponseJSON        `json:"-"`
}

// accountAIRunCfDefogExecuteSqlcoder7b2ResponseJSON contains the JSON metadata for
// the struct [AccountAIRunCfDefogExecuteSqlcoder7b2Response]
type accountAIRunCfDefogExecuteSqlcoder7b2ResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIRunCfDefogExecuteSqlcoder7b2Response) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunCfDefogExecuteSqlcoder7b2ResponseJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [AccountAIRunCfDefogExecuteSqlcoder7b2ResponseResultObject]
// or [shared.UnionString].
type AccountAIRunCfDefogExecuteSqlcoder7b2ResponseResultUnion interface {
	ImplementsAccountAIRunCfDefogExecuteSqlcoder7b2ResponseResultUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccountAIRunCfDefogExecuteSqlcoder7b2ResponseResultUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccountAIRunCfDefogExecuteSqlcoder7b2ResponseResultObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type AccountAIRunCfDefogExecuteSqlcoder7b2ResponseResultObject struct {
	// The generated text response from the model
	Response string `json:"response,required"`
	// An array of tool calls requests made during the response generation
	ToolCalls []AccountAIRunCfDefogExecuteSqlcoder7b2ResponseResultObjectToolCall `json:"tool_calls"`
	// Usage statistics for the inference request
	Usage AccountAIRunCfDefogExecuteSqlcoder7b2ResponseResultObjectUsage `json:"usage"`
	JSON  accountAIRunCfDefogExecuteSqlcoder7b2ResponseResultObjectJSON  `json:"-"`
}

// accountAIRunCfDefogExecuteSqlcoder7b2ResponseResultObjectJSON contains the JSON
// metadata for the struct
// [AccountAIRunCfDefogExecuteSqlcoder7b2ResponseResultObject]
type accountAIRunCfDefogExecuteSqlcoder7b2ResponseResultObjectJSON struct {
	Response    apijson.Field
	ToolCalls   apijson.Field
	Usage       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIRunCfDefogExecuteSqlcoder7b2ResponseResultObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunCfDefogExecuteSqlcoder7b2ResponseResultObjectJSON) RawJSON() string {
	return r.raw
}

func (r AccountAIRunCfDefogExecuteSqlcoder7b2ResponseResultObject) ImplementsAccountAIRunCfDefogExecuteSqlcoder7b2ResponseResultUnion() {
}

type AccountAIRunCfDefogExecuteSqlcoder7b2ResponseResultObjectToolCall struct {
	// The arguments passed to be passed to the tool call request
	Arguments interface{} `json:"arguments"`
	// The name of the tool to be called
	Name string                                                                `json:"name"`
	JSON accountAIRunCfDefogExecuteSqlcoder7b2ResponseResultObjectToolCallJSON `json:"-"`
}

// accountAIRunCfDefogExecuteSqlcoder7b2ResponseResultObjectToolCallJSON contains
// the JSON metadata for the struct
// [AccountAIRunCfDefogExecuteSqlcoder7b2ResponseResultObjectToolCall]
type accountAIRunCfDefogExecuteSqlcoder7b2ResponseResultObjectToolCallJSON struct {
	Arguments   apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIRunCfDefogExecuteSqlcoder7b2ResponseResultObjectToolCall) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunCfDefogExecuteSqlcoder7b2ResponseResultObjectToolCallJSON) RawJSON() string {
	return r.raw
}

// Usage statistics for the inference request
type AccountAIRunCfDefogExecuteSqlcoder7b2ResponseResultObjectUsage struct {
	// Total number of tokens in output
	CompletionTokens float64 `json:"completion_tokens"`
	// Total number of tokens in input
	PromptTokens float64 `json:"prompt_tokens"`
	// Total number of input and output tokens
	TotalTokens float64                                                            `json:"total_tokens"`
	JSON        accountAIRunCfDefogExecuteSqlcoder7b2ResponseResultObjectUsageJSON `json:"-"`
}

// accountAIRunCfDefogExecuteSqlcoder7b2ResponseResultObjectUsageJSON contains the
// JSON metadata for the struct
// [AccountAIRunCfDefogExecuteSqlcoder7b2ResponseResultObjectUsage]
type accountAIRunCfDefogExecuteSqlcoder7b2ResponseResultObjectUsageJSON struct {
	CompletionTokens apijson.Field
	PromptTokens     apijson.Field
	TotalTokens      apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountAIRunCfDefogExecuteSqlcoder7b2ResponseResultObjectUsage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunCfDefogExecuteSqlcoder7b2ResponseResultObjectUsageJSON) RawJSON() string {
	return r.raw
}

type AccountAIRunCfDefogExecuteSqlcoder7b2Params struct {
	QueueRequest param.Field[string]                                  `query:"queueRequest"`
	Body         AccountAIRunCfDefogExecuteSqlcoder7b2ParamsBodyUnion `json:"body"`
}

func (r AccountAIRunCfDefogExecuteSqlcoder7b2Params) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

// URLQuery serializes [AccountAIRunCfDefogExecuteSqlcoder7b2Params]'s query
// parameters as `url.Values`.
func (r AccountAIRunCfDefogExecuteSqlcoder7b2Params) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountAIRunCfDefogExecuteSqlcoder7b2ParamsBody struct {
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

func (r AccountAIRunCfDefogExecuteSqlcoder7b2ParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfDefogExecuteSqlcoder7b2ParamsBody) implementsAccountAIRunCfDefogExecuteSqlcoder7b2ParamsBodyUnion() {
}

// Satisfied by [AccountAIRunCfDefogExecuteSqlcoder7b2ParamsBodyPrompt],
// [AccountAIRunCfDefogExecuteSqlcoder7b2ParamsBodyMessages],
// [AccountAIRunCfDefogExecuteSqlcoder7b2ParamsBody].
type AccountAIRunCfDefogExecuteSqlcoder7b2ParamsBodyUnion interface {
	implementsAccountAIRunCfDefogExecuteSqlcoder7b2ParamsBodyUnion()
}

type AccountAIRunCfDefogExecuteSqlcoder7b2ParamsBodyPrompt struct {
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
	RepetitionPenalty param.Field[float64]                                                             `json:"repetition_penalty"`
	ResponseFormat    param.Field[AccountAIRunCfDefogExecuteSqlcoder7b2ParamsBodyPromptResponseFormat] `json:"response_format"`
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

func (r AccountAIRunCfDefogExecuteSqlcoder7b2ParamsBodyPrompt) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfDefogExecuteSqlcoder7b2ParamsBodyPrompt) implementsAccountAIRunCfDefogExecuteSqlcoder7b2ParamsBodyUnion() {
}

type AccountAIRunCfDefogExecuteSqlcoder7b2ParamsBodyPromptResponseFormat struct {
	JsonSchema param.Field[interface{}]                                                             `json:"json_schema"`
	Type       param.Field[AccountAIRunCfDefogExecuteSqlcoder7b2ParamsBodyPromptResponseFormatType] `json:"type"`
}

func (r AccountAIRunCfDefogExecuteSqlcoder7b2ParamsBodyPromptResponseFormat) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfDefogExecuteSqlcoder7b2ParamsBodyPromptResponseFormatType string

const (
	AccountAIRunCfDefogExecuteSqlcoder7b2ParamsBodyPromptResponseFormatTypeJsonObject AccountAIRunCfDefogExecuteSqlcoder7b2ParamsBodyPromptResponseFormatType = "json_object"
	AccountAIRunCfDefogExecuteSqlcoder7b2ParamsBodyPromptResponseFormatTypeJsonSchema AccountAIRunCfDefogExecuteSqlcoder7b2ParamsBodyPromptResponseFormatType = "json_schema"
)

func (r AccountAIRunCfDefogExecuteSqlcoder7b2ParamsBodyPromptResponseFormatType) IsKnown() bool {
	switch r {
	case AccountAIRunCfDefogExecuteSqlcoder7b2ParamsBodyPromptResponseFormatTypeJsonObject, AccountAIRunCfDefogExecuteSqlcoder7b2ParamsBodyPromptResponseFormatTypeJsonSchema:
		return true
	}
	return false
}

type AccountAIRunCfDefogExecuteSqlcoder7b2ParamsBodyMessages struct {
	// An array of message objects representing the conversation history.
	Messages param.Field[[]AccountAIRunCfDefogExecuteSqlcoder7b2ParamsBodyMessagesMessage] `json:"messages,required"`
	// Decreases the likelihood of the model repeating the same lines verbatim.
	FrequencyPenalty param.Field[float64]                                                           `json:"frequency_penalty"`
	Functions        param.Field[[]AccountAIRunCfDefogExecuteSqlcoder7b2ParamsBodyMessagesFunction] `json:"functions"`
	// The maximum number of tokens to generate in the response.
	MaxTokens param.Field[int64] `json:"max_tokens"`
	// Increases the likelihood of the model introducing new topics.
	PresencePenalty param.Field[float64] `json:"presence_penalty"`
	// If true, a chat template is not applied and you must adhere to the specific
	// model's expected formatting.
	Raw param.Field[bool] `json:"raw"`
	// Penalty for repeated tokens; higher values discourage repetition.
	RepetitionPenalty param.Field[float64]                                                               `json:"repetition_penalty"`
	ResponseFormat    param.Field[AccountAIRunCfDefogExecuteSqlcoder7b2ParamsBodyMessagesResponseFormat] `json:"response_format"`
	// Random seed for reproducibility of the generation.
	Seed param.Field[int64] `json:"seed"`
	// If true, the response will be streamed back incrementally using SSE, Server Sent
	// Events.
	Stream param.Field[bool] `json:"stream"`
	// Controls the randomness of the output; higher values produce more random
	// results.
	Temperature param.Field[float64] `json:"temperature"`
	// A list of tools available for the assistant to use.
	Tools param.Field[[]AccountAIRunCfDefogExecuteSqlcoder7b2ParamsBodyMessagesToolUnion] `json:"tools"`
	// Limits the AI to choose from the top 'k' most probable words. Lower values make
	// responses more focused; higher values introduce more variety and potential
	// surprises.
	TopK param.Field[int64] `json:"top_k"`
	// Adjusts the creativity of the AI's responses by controlling how many possible
	// words it considers. Lower values make outputs more predictable; higher values
	// allow for more varied and creative responses.
	TopP param.Field[float64] `json:"top_p"`
}

func (r AccountAIRunCfDefogExecuteSqlcoder7b2ParamsBodyMessages) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfDefogExecuteSqlcoder7b2ParamsBodyMessages) implementsAccountAIRunCfDefogExecuteSqlcoder7b2ParamsBodyUnion() {
}

type AccountAIRunCfDefogExecuteSqlcoder7b2ParamsBodyMessagesMessage struct {
	// The content of the message as a string.
	Content param.Field[string] `json:"content,required"`
	// The role of the message sender (e.g., 'user', 'assistant', 'system', 'tool').
	Role param.Field[string] `json:"role,required"`
}

func (r AccountAIRunCfDefogExecuteSqlcoder7b2ParamsBodyMessagesMessage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfDefogExecuteSqlcoder7b2ParamsBodyMessagesFunction struct {
	Code param.Field[string] `json:"code,required"`
	Name param.Field[string] `json:"name,required"`
}

func (r AccountAIRunCfDefogExecuteSqlcoder7b2ParamsBodyMessagesFunction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfDefogExecuteSqlcoder7b2ParamsBodyMessagesResponseFormat struct {
	JsonSchema param.Field[interface{}]                                                               `json:"json_schema"`
	Type       param.Field[AccountAIRunCfDefogExecuteSqlcoder7b2ParamsBodyMessagesResponseFormatType] `json:"type"`
}

func (r AccountAIRunCfDefogExecuteSqlcoder7b2ParamsBodyMessagesResponseFormat) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfDefogExecuteSqlcoder7b2ParamsBodyMessagesResponseFormatType string

const (
	AccountAIRunCfDefogExecuteSqlcoder7b2ParamsBodyMessagesResponseFormatTypeJsonObject AccountAIRunCfDefogExecuteSqlcoder7b2ParamsBodyMessagesResponseFormatType = "json_object"
	AccountAIRunCfDefogExecuteSqlcoder7b2ParamsBodyMessagesResponseFormatTypeJsonSchema AccountAIRunCfDefogExecuteSqlcoder7b2ParamsBodyMessagesResponseFormatType = "json_schema"
)

func (r AccountAIRunCfDefogExecuteSqlcoder7b2ParamsBodyMessagesResponseFormatType) IsKnown() bool {
	switch r {
	case AccountAIRunCfDefogExecuteSqlcoder7b2ParamsBodyMessagesResponseFormatTypeJsonObject, AccountAIRunCfDefogExecuteSqlcoder7b2ParamsBodyMessagesResponseFormatTypeJsonSchema:
		return true
	}
	return false
}

type AccountAIRunCfDefogExecuteSqlcoder7b2ParamsBodyMessagesTool struct {
	// A brief description of what the tool does.
	Description param.Field[string]      `json:"description"`
	Function    param.Field[interface{}] `json:"function"`
	// The name of the tool. More descriptive the better.
	Name       param.Field[string]      `json:"name"`
	Parameters param.Field[interface{}] `json:"parameters"`
	// Specifies the type of tool (e.g., 'function').
	Type param.Field[string] `json:"type"`
}

func (r AccountAIRunCfDefogExecuteSqlcoder7b2ParamsBodyMessagesTool) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfDefogExecuteSqlcoder7b2ParamsBodyMessagesTool) implementsAccountAIRunCfDefogExecuteSqlcoder7b2ParamsBodyMessagesToolUnion() {
}

// Satisfied by
// [AccountAIRunCfDefogExecuteSqlcoder7b2ParamsBodyMessagesToolsObject],
// [AccountAIRunCfDefogExecuteSqlcoder7b2ParamsBodyMessagesToolsObject],
// [AccountAIRunCfDefogExecuteSqlcoder7b2ParamsBodyMessagesTool].
type AccountAIRunCfDefogExecuteSqlcoder7b2ParamsBodyMessagesToolUnion interface {
	implementsAccountAIRunCfDefogExecuteSqlcoder7b2ParamsBodyMessagesToolUnion()
}

type AccountAIRunCfDefogExecuteSqlcoder7b2ParamsBodyMessagesToolsObject struct {
	// A brief description of what the tool does.
	Description param.Field[string] `json:"description,required"`
	// The name of the tool. More descriptive the better.
	Name param.Field[string] `json:"name,required"`
	// Schema defining the parameters accepted by the tool.
	Parameters param.Field[AccountAIRunCfDefogExecuteSqlcoder7b2ParamsBodyMessagesToolsObjectParameters] `json:"parameters,required"`
}

func (r AccountAIRunCfDefogExecuteSqlcoder7b2ParamsBodyMessagesToolsObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfDefogExecuteSqlcoder7b2ParamsBodyMessagesToolsObject) implementsAccountAIRunCfDefogExecuteSqlcoder7b2ParamsBodyMessagesToolUnion() {
}

// Schema defining the parameters accepted by the tool.
type AccountAIRunCfDefogExecuteSqlcoder7b2ParamsBodyMessagesToolsObjectParameters struct {
	// Definitions of each parameter.
	Properties param.Field[map[string]AccountAIRunCfDefogExecuteSqlcoder7b2ParamsBodyMessagesToolsObjectParametersProperties] `json:"properties,required"`
	// The type of the parameters object (usually 'object').
	Type param.Field[string] `json:"type,required"`
	// List of required parameter names.
	Required param.Field[[]string] `json:"required"`
}

func (r AccountAIRunCfDefogExecuteSqlcoder7b2ParamsBodyMessagesToolsObjectParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfDefogExecuteSqlcoder7b2ParamsBodyMessagesToolsObjectParametersProperties struct {
	// A description of the expected parameter.
	Description param.Field[string] `json:"description,required"`
	// The data type of the parameter.
	Type param.Field[string] `json:"type,required"`
}

func (r AccountAIRunCfDefogExecuteSqlcoder7b2ParamsBodyMessagesToolsObjectParametersProperties) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
