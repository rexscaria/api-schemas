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

// AccountAIRunCfTiiuaeService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountAIRunCfTiiuaeService] method instead.
type AccountAIRunCfTiiuaeService struct {
	Options []option.RequestOption
}

// NewAccountAIRunCfTiiuaeService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountAIRunCfTiiuaeService(opts ...option.RequestOption) (r *AccountAIRunCfTiiuaeService) {
	r = &AccountAIRunCfTiiuaeService{}
	r.Options = opts
	return
}

// Execute @cf/tiiuae/falcon-7b-instruct model.
func (r *AccountAIRunCfTiiuaeService) ExecuteFalcon7bInstruct(ctx context.Context, accountID string, params AccountAIRunCfTiiuaeExecuteFalcon7bInstructParams, opts ...option.RequestOption) (res *AccountAIRunCfTiiuaeExecuteFalcon7bInstructResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai/run/@cf/tiiuae/falcon-7b-instruct", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type AccountAIRunCfTiiuaeExecuteFalcon7bInstructResponse struct {
	Result  AccountAIRunCfTiiuaeExecuteFalcon7bInstructResponseResultUnion `json:"result" format:"binary"`
	Success bool                                                           `json:"success"`
	JSON    accountAIRunCfTiiuaeExecuteFalcon7bInstructResponseJSON        `json:"-"`
}

// accountAIRunCfTiiuaeExecuteFalcon7bInstructResponseJSON contains the JSON
// metadata for the struct [AccountAIRunCfTiiuaeExecuteFalcon7bInstructResponse]
type accountAIRunCfTiiuaeExecuteFalcon7bInstructResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIRunCfTiiuaeExecuteFalcon7bInstructResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunCfTiiuaeExecuteFalcon7bInstructResponseJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by
// [AccountAIRunCfTiiuaeExecuteFalcon7bInstructResponseResultObject] or
// [shared.UnionString].
type AccountAIRunCfTiiuaeExecuteFalcon7bInstructResponseResultUnion interface {
	ImplementsAccountAIRunCfTiiuaeExecuteFalcon7bInstructResponseResultUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccountAIRunCfTiiuaeExecuteFalcon7bInstructResponseResultUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccountAIRunCfTiiuaeExecuteFalcon7bInstructResponseResultObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type AccountAIRunCfTiiuaeExecuteFalcon7bInstructResponseResultObject struct {
	// The generated text response from the model
	Response string `json:"response,required"`
	// An array of tool calls requests made during the response generation
	ToolCalls []AccountAIRunCfTiiuaeExecuteFalcon7bInstructResponseResultObjectToolCall `json:"tool_calls"`
	// Usage statistics for the inference request
	Usage AccountAIRunCfTiiuaeExecuteFalcon7bInstructResponseResultObjectUsage `json:"usage"`
	JSON  accountAIRunCfTiiuaeExecuteFalcon7bInstructResponseResultObjectJSON  `json:"-"`
}

// accountAIRunCfTiiuaeExecuteFalcon7bInstructResponseResultObjectJSON contains the
// JSON metadata for the struct
// [AccountAIRunCfTiiuaeExecuteFalcon7bInstructResponseResultObject]
type accountAIRunCfTiiuaeExecuteFalcon7bInstructResponseResultObjectJSON struct {
	Response    apijson.Field
	ToolCalls   apijson.Field
	Usage       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIRunCfTiiuaeExecuteFalcon7bInstructResponseResultObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunCfTiiuaeExecuteFalcon7bInstructResponseResultObjectJSON) RawJSON() string {
	return r.raw
}

func (r AccountAIRunCfTiiuaeExecuteFalcon7bInstructResponseResultObject) ImplementsAccountAIRunCfTiiuaeExecuteFalcon7bInstructResponseResultUnion() {
}

type AccountAIRunCfTiiuaeExecuteFalcon7bInstructResponseResultObjectToolCall struct {
	// The arguments passed to be passed to the tool call request
	Arguments interface{} `json:"arguments"`
	// The name of the tool to be called
	Name string                                                                      `json:"name"`
	JSON accountAIRunCfTiiuaeExecuteFalcon7bInstructResponseResultObjectToolCallJSON `json:"-"`
}

// accountAIRunCfTiiuaeExecuteFalcon7bInstructResponseResultObjectToolCallJSON
// contains the JSON metadata for the struct
// [AccountAIRunCfTiiuaeExecuteFalcon7bInstructResponseResultObjectToolCall]
type accountAIRunCfTiiuaeExecuteFalcon7bInstructResponseResultObjectToolCallJSON struct {
	Arguments   apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIRunCfTiiuaeExecuteFalcon7bInstructResponseResultObjectToolCall) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunCfTiiuaeExecuteFalcon7bInstructResponseResultObjectToolCallJSON) RawJSON() string {
	return r.raw
}

// Usage statistics for the inference request
type AccountAIRunCfTiiuaeExecuteFalcon7bInstructResponseResultObjectUsage struct {
	// Total number of tokens in output
	CompletionTokens float64 `json:"completion_tokens"`
	// Total number of tokens in input
	PromptTokens float64 `json:"prompt_tokens"`
	// Total number of input and output tokens
	TotalTokens float64                                                                  `json:"total_tokens"`
	JSON        accountAIRunCfTiiuaeExecuteFalcon7bInstructResponseResultObjectUsageJSON `json:"-"`
}

// accountAIRunCfTiiuaeExecuteFalcon7bInstructResponseResultObjectUsageJSON
// contains the JSON metadata for the struct
// [AccountAIRunCfTiiuaeExecuteFalcon7bInstructResponseResultObjectUsage]
type accountAIRunCfTiiuaeExecuteFalcon7bInstructResponseResultObjectUsageJSON struct {
	CompletionTokens apijson.Field
	PromptTokens     apijson.Field
	TotalTokens      apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountAIRunCfTiiuaeExecuteFalcon7bInstructResponseResultObjectUsage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunCfTiiuaeExecuteFalcon7bInstructResponseResultObjectUsageJSON) RawJSON() string {
	return r.raw
}

type AccountAIRunCfTiiuaeExecuteFalcon7bInstructParams struct {
	QueueRequest param.Field[string]                                        `query:"queueRequest"`
	Body         AccountAIRunCfTiiuaeExecuteFalcon7bInstructParamsBodyUnion `json:"body"`
}

func (r AccountAIRunCfTiiuaeExecuteFalcon7bInstructParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

// URLQuery serializes [AccountAIRunCfTiiuaeExecuteFalcon7bInstructParams]'s query
// parameters as `url.Values`.
func (r AccountAIRunCfTiiuaeExecuteFalcon7bInstructParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountAIRunCfTiiuaeExecuteFalcon7bInstructParamsBody struct {
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

func (r AccountAIRunCfTiiuaeExecuteFalcon7bInstructParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfTiiuaeExecuteFalcon7bInstructParamsBody) implementsAccountAIRunCfTiiuaeExecuteFalcon7bInstructParamsBodyUnion() {
}

// Satisfied by [AccountAIRunCfTiiuaeExecuteFalcon7bInstructParamsBodyPrompt],
// [AccountAIRunCfTiiuaeExecuteFalcon7bInstructParamsBodyMessages],
// [AccountAIRunCfTiiuaeExecuteFalcon7bInstructParamsBody].
type AccountAIRunCfTiiuaeExecuteFalcon7bInstructParamsBodyUnion interface {
	implementsAccountAIRunCfTiiuaeExecuteFalcon7bInstructParamsBodyUnion()
}

type AccountAIRunCfTiiuaeExecuteFalcon7bInstructParamsBodyPrompt struct {
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
	RepetitionPenalty param.Field[float64]                                                                   `json:"repetition_penalty"`
	ResponseFormat    param.Field[AccountAIRunCfTiiuaeExecuteFalcon7bInstructParamsBodyPromptResponseFormat] `json:"response_format"`
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

func (r AccountAIRunCfTiiuaeExecuteFalcon7bInstructParamsBodyPrompt) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfTiiuaeExecuteFalcon7bInstructParamsBodyPrompt) implementsAccountAIRunCfTiiuaeExecuteFalcon7bInstructParamsBodyUnion() {
}

type AccountAIRunCfTiiuaeExecuteFalcon7bInstructParamsBodyPromptResponseFormat struct {
	JsonSchema param.Field[interface{}]                                                                   `json:"json_schema"`
	Type       param.Field[AccountAIRunCfTiiuaeExecuteFalcon7bInstructParamsBodyPromptResponseFormatType] `json:"type"`
}

func (r AccountAIRunCfTiiuaeExecuteFalcon7bInstructParamsBodyPromptResponseFormat) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfTiiuaeExecuteFalcon7bInstructParamsBodyPromptResponseFormatType string

const (
	AccountAIRunCfTiiuaeExecuteFalcon7bInstructParamsBodyPromptResponseFormatTypeJsonObject AccountAIRunCfTiiuaeExecuteFalcon7bInstructParamsBodyPromptResponseFormatType = "json_object"
	AccountAIRunCfTiiuaeExecuteFalcon7bInstructParamsBodyPromptResponseFormatTypeJsonSchema AccountAIRunCfTiiuaeExecuteFalcon7bInstructParamsBodyPromptResponseFormatType = "json_schema"
)

func (r AccountAIRunCfTiiuaeExecuteFalcon7bInstructParamsBodyPromptResponseFormatType) IsKnown() bool {
	switch r {
	case AccountAIRunCfTiiuaeExecuteFalcon7bInstructParamsBodyPromptResponseFormatTypeJsonObject, AccountAIRunCfTiiuaeExecuteFalcon7bInstructParamsBodyPromptResponseFormatTypeJsonSchema:
		return true
	}
	return false
}

type AccountAIRunCfTiiuaeExecuteFalcon7bInstructParamsBodyMessages struct {
	// An array of message objects representing the conversation history.
	Messages param.Field[[]AccountAIRunCfTiiuaeExecuteFalcon7bInstructParamsBodyMessagesMessage] `json:"messages,required"`
	// Decreases the likelihood of the model repeating the same lines verbatim.
	FrequencyPenalty param.Field[float64]                                                                 `json:"frequency_penalty"`
	Functions        param.Field[[]AccountAIRunCfTiiuaeExecuteFalcon7bInstructParamsBodyMessagesFunction] `json:"functions"`
	// The maximum number of tokens to generate in the response.
	MaxTokens param.Field[int64] `json:"max_tokens"`
	// Increases the likelihood of the model introducing new topics.
	PresencePenalty param.Field[float64] `json:"presence_penalty"`
	// If true, a chat template is not applied and you must adhere to the specific
	// model's expected formatting.
	Raw param.Field[bool] `json:"raw"`
	// Penalty for repeated tokens; higher values discourage repetition.
	RepetitionPenalty param.Field[float64]                                                                     `json:"repetition_penalty"`
	ResponseFormat    param.Field[AccountAIRunCfTiiuaeExecuteFalcon7bInstructParamsBodyMessagesResponseFormat] `json:"response_format"`
	// Random seed for reproducibility of the generation.
	Seed param.Field[int64] `json:"seed"`
	// If true, the response will be streamed back incrementally using SSE, Server Sent
	// Events.
	Stream param.Field[bool] `json:"stream"`
	// Controls the randomness of the output; higher values produce more random
	// results.
	Temperature param.Field[float64] `json:"temperature"`
	// A list of tools available for the assistant to use.
	Tools param.Field[[]AccountAIRunCfTiiuaeExecuteFalcon7bInstructParamsBodyMessagesToolUnion] `json:"tools"`
	// Limits the AI to choose from the top 'k' most probable words. Lower values make
	// responses more focused; higher values introduce more variety and potential
	// surprises.
	TopK param.Field[int64] `json:"top_k"`
	// Adjusts the creativity of the AI's responses by controlling how many possible
	// words it considers. Lower values make outputs more predictable; higher values
	// allow for more varied and creative responses.
	TopP param.Field[float64] `json:"top_p"`
}

func (r AccountAIRunCfTiiuaeExecuteFalcon7bInstructParamsBodyMessages) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfTiiuaeExecuteFalcon7bInstructParamsBodyMessages) implementsAccountAIRunCfTiiuaeExecuteFalcon7bInstructParamsBodyUnion() {
}

type AccountAIRunCfTiiuaeExecuteFalcon7bInstructParamsBodyMessagesMessage struct {
	// The content of the message as a string.
	Content param.Field[string] `json:"content,required"`
	// The role of the message sender (e.g., 'user', 'assistant', 'system', 'tool').
	Role param.Field[string] `json:"role,required"`
}

func (r AccountAIRunCfTiiuaeExecuteFalcon7bInstructParamsBodyMessagesMessage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfTiiuaeExecuteFalcon7bInstructParamsBodyMessagesFunction struct {
	Code param.Field[string] `json:"code,required"`
	Name param.Field[string] `json:"name,required"`
}

func (r AccountAIRunCfTiiuaeExecuteFalcon7bInstructParamsBodyMessagesFunction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfTiiuaeExecuteFalcon7bInstructParamsBodyMessagesResponseFormat struct {
	JsonSchema param.Field[interface{}]                                                                     `json:"json_schema"`
	Type       param.Field[AccountAIRunCfTiiuaeExecuteFalcon7bInstructParamsBodyMessagesResponseFormatType] `json:"type"`
}

func (r AccountAIRunCfTiiuaeExecuteFalcon7bInstructParamsBodyMessagesResponseFormat) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfTiiuaeExecuteFalcon7bInstructParamsBodyMessagesResponseFormatType string

const (
	AccountAIRunCfTiiuaeExecuteFalcon7bInstructParamsBodyMessagesResponseFormatTypeJsonObject AccountAIRunCfTiiuaeExecuteFalcon7bInstructParamsBodyMessagesResponseFormatType = "json_object"
	AccountAIRunCfTiiuaeExecuteFalcon7bInstructParamsBodyMessagesResponseFormatTypeJsonSchema AccountAIRunCfTiiuaeExecuteFalcon7bInstructParamsBodyMessagesResponseFormatType = "json_schema"
)

func (r AccountAIRunCfTiiuaeExecuteFalcon7bInstructParamsBodyMessagesResponseFormatType) IsKnown() bool {
	switch r {
	case AccountAIRunCfTiiuaeExecuteFalcon7bInstructParamsBodyMessagesResponseFormatTypeJsonObject, AccountAIRunCfTiiuaeExecuteFalcon7bInstructParamsBodyMessagesResponseFormatTypeJsonSchema:
		return true
	}
	return false
}

type AccountAIRunCfTiiuaeExecuteFalcon7bInstructParamsBodyMessagesTool struct {
	// A brief description of what the tool does.
	Description param.Field[string]      `json:"description"`
	Function    param.Field[interface{}] `json:"function"`
	// The name of the tool. More descriptive the better.
	Name       param.Field[string]      `json:"name"`
	Parameters param.Field[interface{}] `json:"parameters"`
	// Specifies the type of tool (e.g., 'function').
	Type param.Field[string] `json:"type"`
}

func (r AccountAIRunCfTiiuaeExecuteFalcon7bInstructParamsBodyMessagesTool) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfTiiuaeExecuteFalcon7bInstructParamsBodyMessagesTool) implementsAccountAIRunCfTiiuaeExecuteFalcon7bInstructParamsBodyMessagesToolUnion() {
}

// Satisfied by
// [AccountAIRunCfTiiuaeExecuteFalcon7bInstructParamsBodyMessagesToolsObject],
// [AccountAIRunCfTiiuaeExecuteFalcon7bInstructParamsBodyMessagesToolsObject],
// [AccountAIRunCfTiiuaeExecuteFalcon7bInstructParamsBodyMessagesTool].
type AccountAIRunCfTiiuaeExecuteFalcon7bInstructParamsBodyMessagesToolUnion interface {
	implementsAccountAIRunCfTiiuaeExecuteFalcon7bInstructParamsBodyMessagesToolUnion()
}

type AccountAIRunCfTiiuaeExecuteFalcon7bInstructParamsBodyMessagesToolsObject struct {
	// A brief description of what the tool does.
	Description param.Field[string] `json:"description,required"`
	// The name of the tool. More descriptive the better.
	Name param.Field[string] `json:"name,required"`
	// Schema defining the parameters accepted by the tool.
	Parameters param.Field[AccountAIRunCfTiiuaeExecuteFalcon7bInstructParamsBodyMessagesToolsObjectParameters] `json:"parameters,required"`
}

func (r AccountAIRunCfTiiuaeExecuteFalcon7bInstructParamsBodyMessagesToolsObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfTiiuaeExecuteFalcon7bInstructParamsBodyMessagesToolsObject) implementsAccountAIRunCfTiiuaeExecuteFalcon7bInstructParamsBodyMessagesToolUnion() {
}

// Schema defining the parameters accepted by the tool.
type AccountAIRunCfTiiuaeExecuteFalcon7bInstructParamsBodyMessagesToolsObjectParameters struct {
	// Definitions of each parameter.
	Properties param.Field[map[string]AccountAIRunCfTiiuaeExecuteFalcon7bInstructParamsBodyMessagesToolsObjectParametersProperties] `json:"properties,required"`
	// The type of the parameters object (usually 'object').
	Type param.Field[string] `json:"type,required"`
	// List of required parameter names.
	Required param.Field[[]string] `json:"required"`
}

func (r AccountAIRunCfTiiuaeExecuteFalcon7bInstructParamsBodyMessagesToolsObjectParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfTiiuaeExecuteFalcon7bInstructParamsBodyMessagesToolsObjectParametersProperties struct {
	// A description of the expected parameter.
	Description param.Field[string] `json:"description,required"`
	// The data type of the parameter.
	Type param.Field[string] `json:"type,required"`
}

func (r AccountAIRunCfTiiuaeExecuteFalcon7bInstructParamsBodyMessagesToolsObjectParametersProperties) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
