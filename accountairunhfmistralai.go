// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/apiquery"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
	"github.com/rexscaria/api-schemas/shared"
	"github.com/tidwall/gjson"
)

// AccountAIRunHfMistralaiService contains methods and other services that help
// with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountAIRunHfMistralaiService] method instead.
type AccountAIRunHfMistralaiService struct {
	Options []option.RequestOption
}

// NewAccountAIRunHfMistralaiService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountAIRunHfMistralaiService(opts ...option.RequestOption) (r *AccountAIRunHfMistralaiService) {
	r = &AccountAIRunHfMistralaiService{}
	r.Options = opts
	return
}

// Execute @hf/mistralai/mistral-7b-instruct-v0.2 model.
func (r *AccountAIRunHfMistralaiService) ExecuteMistral7bInstructV0_2(ctx context.Context, accountID string, params AccountAIRunHfMistralaiExecuteMistral7bInstructV0_2Params, opts ...option.RequestOption) (res *AccountAIRunHfMistralaiExecuteMistral7bInstructV0_2Response, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai/run/@hf/mistralai/mistral-7b-instruct-v0.2", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type AccountAIRunHfMistralaiExecuteMistral7bInstructV0_2Response struct {
	Result  AccountAIRunHfMistralaiExecuteMistral7bInstructV0_2ResponseResultUnion `json:"result" format:"binary"`
	Success bool                                                                   `json:"success"`
	JSON    accountAIRunHfMistralaiExecuteMistral7bInstructV0_2ResponseJSON        `json:"-"`
}

// accountAIRunHfMistralaiExecuteMistral7bInstructV0_2ResponseJSON contains the
// JSON metadata for the struct
// [AccountAIRunHfMistralaiExecuteMistral7bInstructV0_2Response]
type accountAIRunHfMistralaiExecuteMistral7bInstructV0_2ResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIRunHfMistralaiExecuteMistral7bInstructV0_2Response) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunHfMistralaiExecuteMistral7bInstructV0_2ResponseJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by
// [AccountAIRunHfMistralaiExecuteMistral7bInstructV0_2ResponseResultObject] or
// [shared.UnionString].
type AccountAIRunHfMistralaiExecuteMistral7bInstructV0_2ResponseResultUnion interface {
	ImplementsAccountAIRunHfMistralaiExecuteMistral7bInstructV0_2ResponseResultUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccountAIRunHfMistralaiExecuteMistral7bInstructV0_2ResponseResultUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccountAIRunHfMistralaiExecuteMistral7bInstructV0_2ResponseResultObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type AccountAIRunHfMistralaiExecuteMistral7bInstructV0_2ResponseResultObject struct {
	// The generated text response from the model
	Response string `json:"response,required"`
	// An array of tool calls requests made during the response generation
	ToolCalls []AccountAIRunHfMistralaiExecuteMistral7bInstructV0_2ResponseResultObjectToolCall `json:"tool_calls"`
	// Usage statistics for the inference request
	Usage AccountAIRunHfMistralaiExecuteMistral7bInstructV0_2ResponseResultObjectUsage `json:"usage"`
	JSON  accountAIRunHfMistralaiExecuteMistral7bInstructV0_2ResponseResultObjectJSON  `json:"-"`
}

// accountAIRunHfMistralaiExecuteMistral7bInstructV0_2ResponseResultObjectJSON
// contains the JSON metadata for the struct
// [AccountAIRunHfMistralaiExecuteMistral7bInstructV0_2ResponseResultObject]
type accountAIRunHfMistralaiExecuteMistral7bInstructV0_2ResponseResultObjectJSON struct {
	Response    apijson.Field
	ToolCalls   apijson.Field
	Usage       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIRunHfMistralaiExecuteMistral7bInstructV0_2ResponseResultObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunHfMistralaiExecuteMistral7bInstructV0_2ResponseResultObjectJSON) RawJSON() string {
	return r.raw
}

func (r AccountAIRunHfMistralaiExecuteMistral7bInstructV0_2ResponseResultObject) ImplementsAccountAIRunHfMistralaiExecuteMistral7bInstructV0_2ResponseResultUnion() {
}

type AccountAIRunHfMistralaiExecuteMistral7bInstructV0_2ResponseResultObjectToolCall struct {
	// The arguments passed to be passed to the tool call request
	Arguments interface{} `json:"arguments"`
	// The name of the tool to be called
	Name string                                                                              `json:"name"`
	JSON accountAIRunHfMistralaiExecuteMistral7bInstructV0_2ResponseResultObjectToolCallJSON `json:"-"`
}

// accountAIRunHfMistralaiExecuteMistral7bInstructV0_2ResponseResultObjectToolCallJSON
// contains the JSON metadata for the struct
// [AccountAIRunHfMistralaiExecuteMistral7bInstructV0_2ResponseResultObjectToolCall]
type accountAIRunHfMistralaiExecuteMistral7bInstructV0_2ResponseResultObjectToolCallJSON struct {
	Arguments   apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIRunHfMistralaiExecuteMistral7bInstructV0_2ResponseResultObjectToolCall) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunHfMistralaiExecuteMistral7bInstructV0_2ResponseResultObjectToolCallJSON) RawJSON() string {
	return r.raw
}

// Usage statistics for the inference request
type AccountAIRunHfMistralaiExecuteMistral7bInstructV0_2ResponseResultObjectUsage struct {
	// Total number of tokens in output
	CompletionTokens float64 `json:"completion_tokens"`
	// Total number of tokens in input
	PromptTokens float64 `json:"prompt_tokens"`
	// Total number of input and output tokens
	TotalTokens float64                                                                          `json:"total_tokens"`
	JSON        accountAIRunHfMistralaiExecuteMistral7bInstructV0_2ResponseResultObjectUsageJSON `json:"-"`
}

// accountAIRunHfMistralaiExecuteMistral7bInstructV0_2ResponseResultObjectUsageJSON
// contains the JSON metadata for the struct
// [AccountAIRunHfMistralaiExecuteMistral7bInstructV0_2ResponseResultObjectUsage]
type accountAIRunHfMistralaiExecuteMistral7bInstructV0_2ResponseResultObjectUsageJSON struct {
	CompletionTokens apijson.Field
	PromptTokens     apijson.Field
	TotalTokens      apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountAIRunHfMistralaiExecuteMistral7bInstructV0_2ResponseResultObjectUsage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunHfMistralaiExecuteMistral7bInstructV0_2ResponseResultObjectUsageJSON) RawJSON() string {
	return r.raw
}

type AccountAIRunHfMistralaiExecuteMistral7bInstructV0_2Params struct {
	QueueRequest param.Field[string]                                                `query:"queueRequest"`
	Body         AccountAIRunHfMistralaiExecuteMistral7bInstructV0_2ParamsBodyUnion `json:"body"`
}

func (r AccountAIRunHfMistralaiExecuteMistral7bInstructV0_2Params) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

// URLQuery serializes
// [AccountAIRunHfMistralaiExecuteMistral7bInstructV0_2Params]'s query parameters
// as `url.Values`.
func (r AccountAIRunHfMistralaiExecuteMistral7bInstructV0_2Params) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountAIRunHfMistralaiExecuteMistral7bInstructV0_2ParamsBody struct {
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

func (r AccountAIRunHfMistralaiExecuteMistral7bInstructV0_2ParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunHfMistralaiExecuteMistral7bInstructV0_2ParamsBody) implementsAccountAIRunHfMistralaiExecuteMistral7bInstructV0_2ParamsBodyUnion() {
}

// Satisfied by
// [AccountAIRunHfMistralaiExecuteMistral7bInstructV0_2ParamsBodyPrompt],
// [AccountAIRunHfMistralaiExecuteMistral7bInstructV0_2ParamsBodyMessages],
// [AccountAIRunHfMistralaiExecuteMistral7bInstructV0_2ParamsBody].
type AccountAIRunHfMistralaiExecuteMistral7bInstructV0_2ParamsBodyUnion interface {
	implementsAccountAIRunHfMistralaiExecuteMistral7bInstructV0_2ParamsBodyUnion()
}

type AccountAIRunHfMistralaiExecuteMistral7bInstructV0_2ParamsBodyPrompt struct {
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
	RepetitionPenalty param.Field[float64]                                                                           `json:"repetition_penalty"`
	ResponseFormat    param.Field[AccountAIRunHfMistralaiExecuteMistral7bInstructV0_2ParamsBodyPromptResponseFormat] `json:"response_format"`
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

func (r AccountAIRunHfMistralaiExecuteMistral7bInstructV0_2ParamsBodyPrompt) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunHfMistralaiExecuteMistral7bInstructV0_2ParamsBodyPrompt) implementsAccountAIRunHfMistralaiExecuteMistral7bInstructV0_2ParamsBodyUnion() {
}

type AccountAIRunHfMistralaiExecuteMistral7bInstructV0_2ParamsBodyPromptResponseFormat struct {
	JsonSchema param.Field[interface{}]                                                                           `json:"json_schema"`
	Type       param.Field[AccountAIRunHfMistralaiExecuteMistral7bInstructV0_2ParamsBodyPromptResponseFormatType] `json:"type"`
}

func (r AccountAIRunHfMistralaiExecuteMistral7bInstructV0_2ParamsBodyPromptResponseFormat) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunHfMistralaiExecuteMistral7bInstructV0_2ParamsBodyPromptResponseFormatType string

const (
	AccountAIRunHfMistralaiExecuteMistral7bInstructV0_2ParamsBodyPromptResponseFormatTypeJsonObject AccountAIRunHfMistralaiExecuteMistral7bInstructV0_2ParamsBodyPromptResponseFormatType = "json_object"
	AccountAIRunHfMistralaiExecuteMistral7bInstructV0_2ParamsBodyPromptResponseFormatTypeJsonSchema AccountAIRunHfMistralaiExecuteMistral7bInstructV0_2ParamsBodyPromptResponseFormatType = "json_schema"
)

func (r AccountAIRunHfMistralaiExecuteMistral7bInstructV0_2ParamsBodyPromptResponseFormatType) IsKnown() bool {
	switch r {
	case AccountAIRunHfMistralaiExecuteMistral7bInstructV0_2ParamsBodyPromptResponseFormatTypeJsonObject, AccountAIRunHfMistralaiExecuteMistral7bInstructV0_2ParamsBodyPromptResponseFormatTypeJsonSchema:
		return true
	}
	return false
}

type AccountAIRunHfMistralaiExecuteMistral7bInstructV0_2ParamsBodyMessages struct {
	// An array of message objects representing the conversation history.
	Messages param.Field[[]AccountAIRunHfMistralaiExecuteMistral7bInstructV0_2ParamsBodyMessagesMessage] `json:"messages,required"`
	// Decreases the likelihood of the model repeating the same lines verbatim.
	FrequencyPenalty param.Field[float64]                                                                         `json:"frequency_penalty"`
	Functions        param.Field[[]AccountAIRunHfMistralaiExecuteMistral7bInstructV0_2ParamsBodyMessagesFunction] `json:"functions"`
	// The maximum number of tokens to generate in the response.
	MaxTokens param.Field[int64] `json:"max_tokens"`
	// Increases the likelihood of the model introducing new topics.
	PresencePenalty param.Field[float64] `json:"presence_penalty"`
	// If true, a chat template is not applied and you must adhere to the specific
	// model's expected formatting.
	Raw param.Field[bool] `json:"raw"`
	// Penalty for repeated tokens; higher values discourage repetition.
	RepetitionPenalty param.Field[float64]                                                                             `json:"repetition_penalty"`
	ResponseFormat    param.Field[AccountAIRunHfMistralaiExecuteMistral7bInstructV0_2ParamsBodyMessagesResponseFormat] `json:"response_format"`
	// Random seed for reproducibility of the generation.
	Seed param.Field[int64] `json:"seed"`
	// If true, the response will be streamed back incrementally using SSE, Server Sent
	// Events.
	Stream param.Field[bool] `json:"stream"`
	// Controls the randomness of the output; higher values produce more random
	// results.
	Temperature param.Field[float64] `json:"temperature"`
	// A list of tools available for the assistant to use.
	Tools param.Field[[]AccountAIRunHfMistralaiExecuteMistral7bInstructV0_2ParamsBodyMessagesToolUnion] `json:"tools"`
	// Limits the AI to choose from the top 'k' most probable words. Lower values make
	// responses more focused; higher values introduce more variety and potential
	// surprises.
	TopK param.Field[int64] `json:"top_k"`
	// Adjusts the creativity of the AI's responses by controlling how many possible
	// words it considers. Lower values make outputs more predictable; higher values
	// allow for more varied and creative responses.
	TopP param.Field[float64] `json:"top_p"`
}

func (r AccountAIRunHfMistralaiExecuteMistral7bInstructV0_2ParamsBodyMessages) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunHfMistralaiExecuteMistral7bInstructV0_2ParamsBodyMessages) implementsAccountAIRunHfMistralaiExecuteMistral7bInstructV0_2ParamsBodyUnion() {
}

type AccountAIRunHfMistralaiExecuteMistral7bInstructV0_2ParamsBodyMessagesMessage struct {
	// The content of the message as a string.
	Content param.Field[string] `json:"content,required"`
	// The role of the message sender (e.g., 'user', 'assistant', 'system', 'tool').
	Role param.Field[string] `json:"role,required"`
}

func (r AccountAIRunHfMistralaiExecuteMistral7bInstructV0_2ParamsBodyMessagesMessage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunHfMistralaiExecuteMistral7bInstructV0_2ParamsBodyMessagesFunction struct {
	Code param.Field[string] `json:"code,required"`
	Name param.Field[string] `json:"name,required"`
}

func (r AccountAIRunHfMistralaiExecuteMistral7bInstructV0_2ParamsBodyMessagesFunction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunHfMistralaiExecuteMistral7bInstructV0_2ParamsBodyMessagesResponseFormat struct {
	JsonSchema param.Field[interface{}]                                                                             `json:"json_schema"`
	Type       param.Field[AccountAIRunHfMistralaiExecuteMistral7bInstructV0_2ParamsBodyMessagesResponseFormatType] `json:"type"`
}

func (r AccountAIRunHfMistralaiExecuteMistral7bInstructV0_2ParamsBodyMessagesResponseFormat) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunHfMistralaiExecuteMistral7bInstructV0_2ParamsBodyMessagesResponseFormatType string

const (
	AccountAIRunHfMistralaiExecuteMistral7bInstructV0_2ParamsBodyMessagesResponseFormatTypeJsonObject AccountAIRunHfMistralaiExecuteMistral7bInstructV0_2ParamsBodyMessagesResponseFormatType = "json_object"
	AccountAIRunHfMistralaiExecuteMistral7bInstructV0_2ParamsBodyMessagesResponseFormatTypeJsonSchema AccountAIRunHfMistralaiExecuteMistral7bInstructV0_2ParamsBodyMessagesResponseFormatType = "json_schema"
)

func (r AccountAIRunHfMistralaiExecuteMistral7bInstructV0_2ParamsBodyMessagesResponseFormatType) IsKnown() bool {
	switch r {
	case AccountAIRunHfMistralaiExecuteMistral7bInstructV0_2ParamsBodyMessagesResponseFormatTypeJsonObject, AccountAIRunHfMistralaiExecuteMistral7bInstructV0_2ParamsBodyMessagesResponseFormatTypeJsonSchema:
		return true
	}
	return false
}

type AccountAIRunHfMistralaiExecuteMistral7bInstructV0_2ParamsBodyMessagesTool struct {
	// A brief description of what the tool does.
	Description param.Field[string]      `json:"description"`
	Function    param.Field[interface{}] `json:"function"`
	// The name of the tool. More descriptive the better.
	Name       param.Field[string]      `json:"name"`
	Parameters param.Field[interface{}] `json:"parameters"`
	// Specifies the type of tool (e.g., 'function').
	Type param.Field[string] `json:"type"`
}

func (r AccountAIRunHfMistralaiExecuteMistral7bInstructV0_2ParamsBodyMessagesTool) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunHfMistralaiExecuteMistral7bInstructV0_2ParamsBodyMessagesTool) implementsAccountAIRunHfMistralaiExecuteMistral7bInstructV0_2ParamsBodyMessagesToolUnion() {
}

// Satisfied by
// [AccountAIRunHfMistralaiExecuteMistral7bInstructV0_2ParamsBodyMessagesToolsObject],
// [AccountAIRunHfMistralaiExecuteMistral7bInstructV0_2ParamsBodyMessagesToolsObject],
// [AccountAIRunHfMistralaiExecuteMistral7bInstructV0_2ParamsBodyMessagesTool].
type AccountAIRunHfMistralaiExecuteMistral7bInstructV0_2ParamsBodyMessagesToolUnion interface {
	implementsAccountAIRunHfMistralaiExecuteMistral7bInstructV0_2ParamsBodyMessagesToolUnion()
}

type AccountAIRunHfMistralaiExecuteMistral7bInstructV0_2ParamsBodyMessagesToolsObject struct {
	// A brief description of what the tool does.
	Description param.Field[string] `json:"description,required"`
	// The name of the tool. More descriptive the better.
	Name param.Field[string] `json:"name,required"`
	// Schema defining the parameters accepted by the tool.
	Parameters param.Field[AccountAIRunHfMistralaiExecuteMistral7bInstructV0_2ParamsBodyMessagesToolsObjectParameters] `json:"parameters,required"`
}

func (r AccountAIRunHfMistralaiExecuteMistral7bInstructV0_2ParamsBodyMessagesToolsObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunHfMistralaiExecuteMistral7bInstructV0_2ParamsBodyMessagesToolsObject) implementsAccountAIRunHfMistralaiExecuteMistral7bInstructV0_2ParamsBodyMessagesToolUnion() {
}

// Schema defining the parameters accepted by the tool.
type AccountAIRunHfMistralaiExecuteMistral7bInstructV0_2ParamsBodyMessagesToolsObjectParameters struct {
	// Definitions of each parameter.
	Properties param.Field[map[string]AccountAIRunHfMistralaiExecuteMistral7bInstructV0_2ParamsBodyMessagesToolsObjectParametersProperties] `json:"properties,required"`
	// The type of the parameters object (usually 'object').
	Type param.Field[string] `json:"type,required"`
	// List of required parameter names.
	Required param.Field[[]string] `json:"required"`
}

func (r AccountAIRunHfMistralaiExecuteMistral7bInstructV0_2ParamsBodyMessagesToolsObjectParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunHfMistralaiExecuteMistral7bInstructV0_2ParamsBodyMessagesToolsObjectParametersProperties struct {
	// A description of the expected parameter.
	Description param.Field[string] `json:"description,required"`
	// The data type of the parameter.
	Type param.Field[string] `json:"type,required"`
}

func (r AccountAIRunHfMistralaiExecuteMistral7bInstructV0_2ParamsBodyMessagesToolsObjectParametersProperties) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
