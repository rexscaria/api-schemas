// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"mime/multipart"
	"net/http"
	"net/url"

	"github.com/rexscaria/api-schemas/internal/apiform"
	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/apiquery"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
)

// AccountAIRunCfMetaService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountAIRunCfMetaService] method instead.
type AccountAIRunCfMetaService struct {
	Options []option.RequestOption
}

// NewAccountAIRunCfMetaService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountAIRunCfMetaService(opts ...option.RequestOption) (r *AccountAIRunCfMetaService) {
	r = &AccountAIRunCfMetaService{}
	r.Options = opts
	return
}

// Execute @cf/meta/llama-2-7b-chat-fp16 model.
func (r *AccountAIRunCfMetaService) ExecuteLlama2_7bChatFp16(ctx context.Context, accountID string, params AccountAIRunCfMetaExecuteLlama2_7bChatFp16Params, opts ...option.RequestOption) (res *AccountAIRunCfMetaExecuteLlama2_7bChatFp16Response, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai/run/@cf/meta/llama-2-7b-chat-fp16", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Execute @cf/meta/llama-2-7b-chat-int8 model.
func (r *AccountAIRunCfMetaService) ExecuteLlama2_7bChatInt8(ctx context.Context, accountID string, params AccountAIRunCfMetaExecuteLlama2_7bChatInt8Params, opts ...option.RequestOption) (res *AccountAIRunCfMetaExecuteLlama2_7bChatInt8Response, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai/run/@cf/meta/llama-2-7b-chat-int8", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Execute @cf/meta/llama-3.1-70b-instruct model.
func (r *AccountAIRunCfMetaService) ExecuteLlama3_1_70bInstruct(ctx context.Context, accountID string, params AccountAIRunCfMetaExecuteLlama3_1_70bInstructParams, opts ...option.RequestOption) (res *AccountAIRunCfMetaExecuteLlama3_1_70bInstructResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai/run/@cf/meta/llama-3.1-70b-instruct", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Execute @cf/meta/llama-3.1-70b-instruct-preview model.
func (r *AccountAIRunCfMetaService) ExecuteLlama3_1_70bInstructPreview(ctx context.Context, accountID string, params AccountAIRunCfMetaExecuteLlama3_1_70bInstructPreviewParams, opts ...option.RequestOption) (res *AccountAIRunCfMetaExecuteLlama3_1_70bInstructPreviewResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai/run/@cf/meta/llama-3.1-70b-instruct-preview", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Execute @cf/meta/llama-3.1-70b-preview model.
func (r *AccountAIRunCfMetaService) ExecuteLlama3_1_70bPreview(ctx context.Context, accountID string, params AccountAIRunCfMetaExecuteLlama3_1_70bPreviewParams, opts ...option.RequestOption) (res *AccountAIRunCfMetaExecuteLlama3_1_70bPreviewResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai/run/@cf/meta/llama-3.1-70b-preview", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Execute @cf/meta/llama-3.1-8b-instruct-awq model.
func (r *AccountAIRunCfMetaService) ExecuteLlama3_1_8bInstructAwq(ctx context.Context, accountID string, params AccountAIRunCfMetaExecuteLlama3_1_8bInstructAwqParams, opts ...option.RequestOption) (res *AccountAIRunCfMetaExecuteLlama3_1_8bInstructAwqResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai/run/@cf/meta/llama-3.1-8b-instruct-awq", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Execute @cf/meta/llama-3.1-8b-instruct-fast model.
func (r *AccountAIRunCfMetaService) ExecuteLlama3_1_8bInstructFast(ctx context.Context, accountID string, params AccountAIRunCfMetaExecuteLlama3_1_8bInstructFastParams, opts ...option.RequestOption) (res *AccountAIRunCfMetaExecuteLlama3_1_8bInstructFastResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai/run/@cf/meta/llama-3.1-8b-instruct-fast", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Execute @cf/meta/llama-3.1-8b-instruct-fp8 model.
func (r *AccountAIRunCfMetaService) ExecuteLlama3_1_8bInstructFp8(ctx context.Context, accountID string, params AccountAIRunCfMetaExecuteLlama3_1_8bInstructFp8Params, opts ...option.RequestOption) (res *AccountAIRunCfMetaExecuteLlama3_1_8bInstructFp8Response, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai/run/@cf/meta/llama-3.1-8b-instruct-fp8", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Execute @cf/meta/llama-3.1-8b-preview model.
func (r *AccountAIRunCfMetaService) ExecuteLlama3_1_8bPreview(ctx context.Context, accountID string, params AccountAIRunCfMetaExecuteLlama3_1_8bPreviewParams, opts ...option.RequestOption) (res *AccountAIRunCfMetaExecuteLlama3_1_8bPreviewResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai/run/@cf/meta/llama-3.1-8b-preview", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Execute @cf/meta/llama-3.2-11b-vision-instruct model.
func (r *AccountAIRunCfMetaService) ExecuteLlama3_2_11bVisionInstruct(ctx context.Context, accountID string, params AccountAIRunCfMetaExecuteLlama3_2_11bVisionInstructParams, opts ...option.RequestOption) (res *AccountAIRunCfMetaExecuteLlama3_2_11bVisionInstructResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai/run/@cf/meta/llama-3.2-11b-vision-instruct", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Execute @cf/meta/llama-3.2-1b-instruct model.
func (r *AccountAIRunCfMetaService) ExecuteLlama3_2_1bInstruct(ctx context.Context, accountID string, params AccountAIRunCfMetaExecuteLlama3_2_1bInstructParams, opts ...option.RequestOption) (res *AccountAIRunCfMetaExecuteLlama3_2_1bInstructResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai/run/@cf/meta/llama-3.2-1b-instruct", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Execute @cf/meta/llama-3.2-3b-instruct model.
func (r *AccountAIRunCfMetaService) ExecuteLlama3_2_3bInstruct(ctx context.Context, accountID string, params AccountAIRunCfMetaExecuteLlama3_2_3bInstructParams, opts ...option.RequestOption) (res *AccountAIRunCfMetaExecuteLlama3_2_3bInstructResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai/run/@cf/meta/llama-3.2-3b-instruct", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Execute @cf/meta/llama-3-8b-instruct model.
func (r *AccountAIRunCfMetaService) ExecuteLlama3_8bInstruct(ctx context.Context, accountID string, params AccountAIRunCfMetaExecuteLlama3_8bInstructParams, opts ...option.RequestOption) (res *AccountAIRunCfMetaExecuteLlama3_8bInstructResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai/run/@cf/meta/llama-3-8b-instruct", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Execute @cf/meta/llama-3-8b-instruct-awq model.
func (r *AccountAIRunCfMetaService) ExecuteLlama3_8bInstructAwq(ctx context.Context, accountID string, params AccountAIRunCfMetaExecuteLlama3_8bInstructAwqParams, opts ...option.RequestOption) (res *AccountAIRunCfMetaExecuteLlama3_8bInstructAwqResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai/run/@cf/meta/llama-3-8b-instruct-awq", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Execute @cf/meta/llama-guard-3-8b model.
func (r *AccountAIRunCfMetaService) ExecuteLlamaGuard3_8b(ctx context.Context, accountID string, params AccountAIRunCfMetaExecuteLlamaGuard3_8bParams, opts ...option.RequestOption) (res *AccountAIRunCfMetaExecuteLlamaGuard3_8bResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai/run/@cf/meta/llama-guard-3-8b", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Execute @cf/meta/m2m100-1.2b model.
func (r *AccountAIRunCfMetaService) ExecuteM2m100_1_2b(ctx context.Context, accountID string, params AccountAIRunCfMetaExecuteM2m100_1_2bParams, opts ...option.RequestOption) (res *AccountAIRunCfMetaExecuteM2m100_1_2bResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai/run/@cf/meta/m2m100-1.2b", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type AccountAIRunCfMetaExecuteLlama2_7bChatFp16Response = interface{}

type AccountAIRunCfMetaExecuteLlama2_7bChatInt8Response = interface{}

type AccountAIRunCfMetaExecuteLlama3_1_70bInstructResponse = interface{}

type AccountAIRunCfMetaExecuteLlama3_1_70bInstructPreviewResponse = interface{}

type AccountAIRunCfMetaExecuteLlama3_1_70bPreviewResponse = interface{}

type AccountAIRunCfMetaExecuteLlama3_1_8bInstructAwqResponse = interface{}

type AccountAIRunCfMetaExecuteLlama3_1_8bInstructFastResponse = interface{}

type AccountAIRunCfMetaExecuteLlama3_1_8bInstructFp8Response = interface{}

type AccountAIRunCfMetaExecuteLlama3_1_8bPreviewResponse = interface{}

type AccountAIRunCfMetaExecuteLlama3_2_11bVisionInstructResponse = interface{}

type AccountAIRunCfMetaExecuteLlama3_2_1bInstructResponse = interface{}

type AccountAIRunCfMetaExecuteLlama3_2_3bInstructResponse = interface{}

type AccountAIRunCfMetaExecuteLlama3_8bInstructResponse = interface{}

type AccountAIRunCfMetaExecuteLlama3_8bInstructAwqResponse = interface{}

type AccountAIRunCfMetaExecuteLlamaGuard3_8bResponse = interface{}

type AccountAIRunCfMetaExecuteM2m100_1_2bResponse = interface{}

type AccountAIRunCfMetaExecuteLlama2_7bChatFp16Params struct {
	QueueRequest param.Field[string]                                       `query:"queueRequest"`
	Body         AccountAIRunCfMetaExecuteLlama2_7bChatFp16ParamsBodyUnion `json:"body"`
}

func (r AccountAIRunCfMetaExecuteLlama2_7bChatFp16Params) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

// URLQuery serializes [AccountAIRunCfMetaExecuteLlama2_7bChatFp16Params]'s query
// parameters as `url.Values`.
func (r AccountAIRunCfMetaExecuteLlama2_7bChatFp16Params) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountAIRunCfMetaExecuteLlama2_7bChatFp16ParamsBody struct {
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

func (r AccountAIRunCfMetaExecuteLlama2_7bChatFp16ParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfMetaExecuteLlama2_7bChatFp16ParamsBody) implementsAccountAIRunCfMetaExecuteLlama2_7bChatFp16ParamsBodyUnion() {
}

// Satisfied by [AccountAIRunCfMetaExecuteLlama2_7bChatFp16ParamsBodyPrompt],
// [AccountAIRunCfMetaExecuteLlama2_7bChatFp16ParamsBodyMessages],
// [AccountAIRunCfMetaExecuteLlama2_7bChatFp16ParamsBody].
type AccountAIRunCfMetaExecuteLlama2_7bChatFp16ParamsBodyUnion interface {
	implementsAccountAIRunCfMetaExecuteLlama2_7bChatFp16ParamsBodyUnion()
}

type AccountAIRunCfMetaExecuteLlama2_7bChatFp16ParamsBodyPrompt struct {
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
	RepetitionPenalty param.Field[float64]                                                                  `json:"repetition_penalty"`
	ResponseFormat    param.Field[AccountAIRunCfMetaExecuteLlama2_7bChatFp16ParamsBodyPromptResponseFormat] `json:"response_format"`
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

func (r AccountAIRunCfMetaExecuteLlama2_7bChatFp16ParamsBodyPrompt) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfMetaExecuteLlama2_7bChatFp16ParamsBodyPrompt) implementsAccountAIRunCfMetaExecuteLlama2_7bChatFp16ParamsBodyUnion() {
}

type AccountAIRunCfMetaExecuteLlama2_7bChatFp16ParamsBodyPromptResponseFormat struct {
	JsonSchema param.Field[interface{}]                                                                  `json:"json_schema"`
	Type       param.Field[AccountAIRunCfMetaExecuteLlama2_7bChatFp16ParamsBodyPromptResponseFormatType] `json:"type"`
}

func (r AccountAIRunCfMetaExecuteLlama2_7bChatFp16ParamsBodyPromptResponseFormat) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfMetaExecuteLlama2_7bChatFp16ParamsBodyPromptResponseFormatType string

const (
	AccountAIRunCfMetaExecuteLlama2_7bChatFp16ParamsBodyPromptResponseFormatTypeJsonObject AccountAIRunCfMetaExecuteLlama2_7bChatFp16ParamsBodyPromptResponseFormatType = "json_object"
	AccountAIRunCfMetaExecuteLlama2_7bChatFp16ParamsBodyPromptResponseFormatTypeJsonSchema AccountAIRunCfMetaExecuteLlama2_7bChatFp16ParamsBodyPromptResponseFormatType = "json_schema"
)

func (r AccountAIRunCfMetaExecuteLlama2_7bChatFp16ParamsBodyPromptResponseFormatType) IsKnown() bool {
	switch r {
	case AccountAIRunCfMetaExecuteLlama2_7bChatFp16ParamsBodyPromptResponseFormatTypeJsonObject, AccountAIRunCfMetaExecuteLlama2_7bChatFp16ParamsBodyPromptResponseFormatTypeJsonSchema:
		return true
	}
	return false
}

type AccountAIRunCfMetaExecuteLlama2_7bChatFp16ParamsBodyMessages struct {
	// An array of message objects representing the conversation history.
	Messages param.Field[[]AccountAIRunCfMetaExecuteLlama2_7bChatFp16ParamsBodyMessagesMessage] `json:"messages,required"`
	// Decreases the likelihood of the model repeating the same lines verbatim.
	FrequencyPenalty param.Field[float64]                                                                `json:"frequency_penalty"`
	Functions        param.Field[[]AccountAIRunCfMetaExecuteLlama2_7bChatFp16ParamsBodyMessagesFunction] `json:"functions"`
	// The maximum number of tokens to generate in the response.
	MaxTokens param.Field[int64] `json:"max_tokens"`
	// Increases the likelihood of the model introducing new topics.
	PresencePenalty param.Field[float64] `json:"presence_penalty"`
	// If true, a chat template is not applied and you must adhere to the specific
	// model's expected formatting.
	Raw param.Field[bool] `json:"raw"`
	// Penalty for repeated tokens; higher values discourage repetition.
	RepetitionPenalty param.Field[float64]                                                                    `json:"repetition_penalty"`
	ResponseFormat    param.Field[AccountAIRunCfMetaExecuteLlama2_7bChatFp16ParamsBodyMessagesResponseFormat] `json:"response_format"`
	// Random seed for reproducibility of the generation.
	Seed param.Field[int64] `json:"seed"`
	// If true, the response will be streamed back incrementally using SSE, Server Sent
	// Events.
	Stream param.Field[bool] `json:"stream"`
	// Controls the randomness of the output; higher values produce more random
	// results.
	Temperature param.Field[float64] `json:"temperature"`
	// A list of tools available for the assistant to use.
	Tools param.Field[[]AccountAIRunCfMetaExecuteLlama2_7bChatFp16ParamsBodyMessagesToolUnion] `json:"tools"`
	// Limits the AI to choose from the top 'k' most probable words. Lower values make
	// responses more focused; higher values introduce more variety and potential
	// surprises.
	TopK param.Field[int64] `json:"top_k"`
	// Adjusts the creativity of the AI's responses by controlling how many possible
	// words it considers. Lower values make outputs more predictable; higher values
	// allow for more varied and creative responses.
	TopP param.Field[float64] `json:"top_p"`
}

func (r AccountAIRunCfMetaExecuteLlama2_7bChatFp16ParamsBodyMessages) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfMetaExecuteLlama2_7bChatFp16ParamsBodyMessages) implementsAccountAIRunCfMetaExecuteLlama2_7bChatFp16ParamsBodyUnion() {
}

type AccountAIRunCfMetaExecuteLlama2_7bChatFp16ParamsBodyMessagesMessage struct {
	// The content of the message as a string.
	Content param.Field[string] `json:"content,required"`
	// The role of the message sender (e.g., 'user', 'assistant', 'system', 'tool').
	Role param.Field[string] `json:"role,required"`
}

func (r AccountAIRunCfMetaExecuteLlama2_7bChatFp16ParamsBodyMessagesMessage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfMetaExecuteLlama2_7bChatFp16ParamsBodyMessagesFunction struct {
	Code param.Field[string] `json:"code,required"`
	Name param.Field[string] `json:"name,required"`
}

func (r AccountAIRunCfMetaExecuteLlama2_7bChatFp16ParamsBodyMessagesFunction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfMetaExecuteLlama2_7bChatFp16ParamsBodyMessagesResponseFormat struct {
	JsonSchema param.Field[interface{}]                                                                    `json:"json_schema"`
	Type       param.Field[AccountAIRunCfMetaExecuteLlama2_7bChatFp16ParamsBodyMessagesResponseFormatType] `json:"type"`
}

func (r AccountAIRunCfMetaExecuteLlama2_7bChatFp16ParamsBodyMessagesResponseFormat) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfMetaExecuteLlama2_7bChatFp16ParamsBodyMessagesResponseFormatType string

const (
	AccountAIRunCfMetaExecuteLlama2_7bChatFp16ParamsBodyMessagesResponseFormatTypeJsonObject AccountAIRunCfMetaExecuteLlama2_7bChatFp16ParamsBodyMessagesResponseFormatType = "json_object"
	AccountAIRunCfMetaExecuteLlama2_7bChatFp16ParamsBodyMessagesResponseFormatTypeJsonSchema AccountAIRunCfMetaExecuteLlama2_7bChatFp16ParamsBodyMessagesResponseFormatType = "json_schema"
)

func (r AccountAIRunCfMetaExecuteLlama2_7bChatFp16ParamsBodyMessagesResponseFormatType) IsKnown() bool {
	switch r {
	case AccountAIRunCfMetaExecuteLlama2_7bChatFp16ParamsBodyMessagesResponseFormatTypeJsonObject, AccountAIRunCfMetaExecuteLlama2_7bChatFp16ParamsBodyMessagesResponseFormatTypeJsonSchema:
		return true
	}
	return false
}

type AccountAIRunCfMetaExecuteLlama2_7bChatFp16ParamsBodyMessagesTool struct {
	// A brief description of what the tool does.
	Description param.Field[string]      `json:"description"`
	Function    param.Field[interface{}] `json:"function"`
	// The name of the tool. More descriptive the better.
	Name       param.Field[string]      `json:"name"`
	Parameters param.Field[interface{}] `json:"parameters"`
	// Specifies the type of tool (e.g., 'function').
	Type param.Field[string] `json:"type"`
}

func (r AccountAIRunCfMetaExecuteLlama2_7bChatFp16ParamsBodyMessagesTool) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfMetaExecuteLlama2_7bChatFp16ParamsBodyMessagesTool) implementsAccountAIRunCfMetaExecuteLlama2_7bChatFp16ParamsBodyMessagesToolUnion() {
}

// Satisfied by
// [AccountAIRunCfMetaExecuteLlama2_7bChatFp16ParamsBodyMessagesToolsObject],
// [AccountAIRunCfMetaExecuteLlama2_7bChatFp16ParamsBodyMessagesToolsObject],
// [AccountAIRunCfMetaExecuteLlama2_7bChatFp16ParamsBodyMessagesTool].
type AccountAIRunCfMetaExecuteLlama2_7bChatFp16ParamsBodyMessagesToolUnion interface {
	implementsAccountAIRunCfMetaExecuteLlama2_7bChatFp16ParamsBodyMessagesToolUnion()
}

type AccountAIRunCfMetaExecuteLlama2_7bChatFp16ParamsBodyMessagesToolsObject struct {
	// A brief description of what the tool does.
	Description param.Field[string] `json:"description,required"`
	// The name of the tool. More descriptive the better.
	Name param.Field[string] `json:"name,required"`
	// Schema defining the parameters accepted by the tool.
	Parameters param.Field[AccountAIRunCfMetaExecuteLlama2_7bChatFp16ParamsBodyMessagesToolsObjectParameters] `json:"parameters,required"`
}

func (r AccountAIRunCfMetaExecuteLlama2_7bChatFp16ParamsBodyMessagesToolsObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfMetaExecuteLlama2_7bChatFp16ParamsBodyMessagesToolsObject) implementsAccountAIRunCfMetaExecuteLlama2_7bChatFp16ParamsBodyMessagesToolUnion() {
}

// Schema defining the parameters accepted by the tool.
type AccountAIRunCfMetaExecuteLlama2_7bChatFp16ParamsBodyMessagesToolsObjectParameters struct {
	// Definitions of each parameter.
	Properties param.Field[map[string]AccountAIRunCfMetaExecuteLlama2_7bChatFp16ParamsBodyMessagesToolsObjectParametersProperties] `json:"properties,required"`
	// The type of the parameters object (usually 'object').
	Type param.Field[string] `json:"type,required"`
	// List of required parameter names.
	Required param.Field[[]string] `json:"required"`
}

func (r AccountAIRunCfMetaExecuteLlama2_7bChatFp16ParamsBodyMessagesToolsObjectParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfMetaExecuteLlama2_7bChatFp16ParamsBodyMessagesToolsObjectParametersProperties struct {
	// A description of the expected parameter.
	Description param.Field[string] `json:"description,required"`
	// The data type of the parameter.
	Type param.Field[string] `json:"type,required"`
}

func (r AccountAIRunCfMetaExecuteLlama2_7bChatFp16ParamsBodyMessagesToolsObjectParametersProperties) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfMetaExecuteLlama2_7bChatInt8Params struct {
	QueueRequest param.Field[string]                                       `query:"queueRequest"`
	Body         AccountAIRunCfMetaExecuteLlama2_7bChatInt8ParamsBodyUnion `json:"body"`
}

func (r AccountAIRunCfMetaExecuteLlama2_7bChatInt8Params) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

// URLQuery serializes [AccountAIRunCfMetaExecuteLlama2_7bChatInt8Params]'s query
// parameters as `url.Values`.
func (r AccountAIRunCfMetaExecuteLlama2_7bChatInt8Params) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountAIRunCfMetaExecuteLlama2_7bChatInt8ParamsBody struct {
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

func (r AccountAIRunCfMetaExecuteLlama2_7bChatInt8ParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfMetaExecuteLlama2_7bChatInt8ParamsBody) implementsAccountAIRunCfMetaExecuteLlama2_7bChatInt8ParamsBodyUnion() {
}

// Satisfied by [AccountAIRunCfMetaExecuteLlama2_7bChatInt8ParamsBodyPrompt],
// [AccountAIRunCfMetaExecuteLlama2_7bChatInt8ParamsBodyMessages],
// [AccountAIRunCfMetaExecuteLlama2_7bChatInt8ParamsBody].
type AccountAIRunCfMetaExecuteLlama2_7bChatInt8ParamsBodyUnion interface {
	implementsAccountAIRunCfMetaExecuteLlama2_7bChatInt8ParamsBodyUnion()
}

type AccountAIRunCfMetaExecuteLlama2_7bChatInt8ParamsBodyPrompt struct {
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
	RepetitionPenalty param.Field[float64]                                                                  `json:"repetition_penalty"`
	ResponseFormat    param.Field[AccountAIRunCfMetaExecuteLlama2_7bChatInt8ParamsBodyPromptResponseFormat] `json:"response_format"`
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

func (r AccountAIRunCfMetaExecuteLlama2_7bChatInt8ParamsBodyPrompt) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfMetaExecuteLlama2_7bChatInt8ParamsBodyPrompt) implementsAccountAIRunCfMetaExecuteLlama2_7bChatInt8ParamsBodyUnion() {
}

type AccountAIRunCfMetaExecuteLlama2_7bChatInt8ParamsBodyPromptResponseFormat struct {
	JsonSchema param.Field[interface{}]                                                                  `json:"json_schema"`
	Type       param.Field[AccountAIRunCfMetaExecuteLlama2_7bChatInt8ParamsBodyPromptResponseFormatType] `json:"type"`
}

func (r AccountAIRunCfMetaExecuteLlama2_7bChatInt8ParamsBodyPromptResponseFormat) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfMetaExecuteLlama2_7bChatInt8ParamsBodyPromptResponseFormatType string

const (
	AccountAIRunCfMetaExecuteLlama2_7bChatInt8ParamsBodyPromptResponseFormatTypeJsonObject AccountAIRunCfMetaExecuteLlama2_7bChatInt8ParamsBodyPromptResponseFormatType = "json_object"
	AccountAIRunCfMetaExecuteLlama2_7bChatInt8ParamsBodyPromptResponseFormatTypeJsonSchema AccountAIRunCfMetaExecuteLlama2_7bChatInt8ParamsBodyPromptResponseFormatType = "json_schema"
)

func (r AccountAIRunCfMetaExecuteLlama2_7bChatInt8ParamsBodyPromptResponseFormatType) IsKnown() bool {
	switch r {
	case AccountAIRunCfMetaExecuteLlama2_7bChatInt8ParamsBodyPromptResponseFormatTypeJsonObject, AccountAIRunCfMetaExecuteLlama2_7bChatInt8ParamsBodyPromptResponseFormatTypeJsonSchema:
		return true
	}
	return false
}

type AccountAIRunCfMetaExecuteLlama2_7bChatInt8ParamsBodyMessages struct {
	// An array of message objects representing the conversation history.
	Messages param.Field[[]AccountAIRunCfMetaExecuteLlama2_7bChatInt8ParamsBodyMessagesMessage] `json:"messages,required"`
	// Decreases the likelihood of the model repeating the same lines verbatim.
	FrequencyPenalty param.Field[float64]                                                                `json:"frequency_penalty"`
	Functions        param.Field[[]AccountAIRunCfMetaExecuteLlama2_7bChatInt8ParamsBodyMessagesFunction] `json:"functions"`
	// The maximum number of tokens to generate in the response.
	MaxTokens param.Field[int64] `json:"max_tokens"`
	// Increases the likelihood of the model introducing new topics.
	PresencePenalty param.Field[float64] `json:"presence_penalty"`
	// If true, a chat template is not applied and you must adhere to the specific
	// model's expected formatting.
	Raw param.Field[bool] `json:"raw"`
	// Penalty for repeated tokens; higher values discourage repetition.
	RepetitionPenalty param.Field[float64]                                                                    `json:"repetition_penalty"`
	ResponseFormat    param.Field[AccountAIRunCfMetaExecuteLlama2_7bChatInt8ParamsBodyMessagesResponseFormat] `json:"response_format"`
	// Random seed for reproducibility of the generation.
	Seed param.Field[int64] `json:"seed"`
	// If true, the response will be streamed back incrementally using SSE, Server Sent
	// Events.
	Stream param.Field[bool] `json:"stream"`
	// Controls the randomness of the output; higher values produce more random
	// results.
	Temperature param.Field[float64] `json:"temperature"`
	// A list of tools available for the assistant to use.
	Tools param.Field[[]AccountAIRunCfMetaExecuteLlama2_7bChatInt8ParamsBodyMessagesToolUnion] `json:"tools"`
	// Limits the AI to choose from the top 'k' most probable words. Lower values make
	// responses more focused; higher values introduce more variety and potential
	// surprises.
	TopK param.Field[int64] `json:"top_k"`
	// Adjusts the creativity of the AI's responses by controlling how many possible
	// words it considers. Lower values make outputs more predictable; higher values
	// allow for more varied and creative responses.
	TopP param.Field[float64] `json:"top_p"`
}

func (r AccountAIRunCfMetaExecuteLlama2_7bChatInt8ParamsBodyMessages) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfMetaExecuteLlama2_7bChatInt8ParamsBodyMessages) implementsAccountAIRunCfMetaExecuteLlama2_7bChatInt8ParamsBodyUnion() {
}

type AccountAIRunCfMetaExecuteLlama2_7bChatInt8ParamsBodyMessagesMessage struct {
	// The content of the message as a string.
	Content param.Field[string] `json:"content,required"`
	// The role of the message sender (e.g., 'user', 'assistant', 'system', 'tool').
	Role param.Field[string] `json:"role,required"`
}

func (r AccountAIRunCfMetaExecuteLlama2_7bChatInt8ParamsBodyMessagesMessage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfMetaExecuteLlama2_7bChatInt8ParamsBodyMessagesFunction struct {
	Code param.Field[string] `json:"code,required"`
	Name param.Field[string] `json:"name,required"`
}

func (r AccountAIRunCfMetaExecuteLlama2_7bChatInt8ParamsBodyMessagesFunction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfMetaExecuteLlama2_7bChatInt8ParamsBodyMessagesResponseFormat struct {
	JsonSchema param.Field[interface{}]                                                                    `json:"json_schema"`
	Type       param.Field[AccountAIRunCfMetaExecuteLlama2_7bChatInt8ParamsBodyMessagesResponseFormatType] `json:"type"`
}

func (r AccountAIRunCfMetaExecuteLlama2_7bChatInt8ParamsBodyMessagesResponseFormat) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfMetaExecuteLlama2_7bChatInt8ParamsBodyMessagesResponseFormatType string

const (
	AccountAIRunCfMetaExecuteLlama2_7bChatInt8ParamsBodyMessagesResponseFormatTypeJsonObject AccountAIRunCfMetaExecuteLlama2_7bChatInt8ParamsBodyMessagesResponseFormatType = "json_object"
	AccountAIRunCfMetaExecuteLlama2_7bChatInt8ParamsBodyMessagesResponseFormatTypeJsonSchema AccountAIRunCfMetaExecuteLlama2_7bChatInt8ParamsBodyMessagesResponseFormatType = "json_schema"
)

func (r AccountAIRunCfMetaExecuteLlama2_7bChatInt8ParamsBodyMessagesResponseFormatType) IsKnown() bool {
	switch r {
	case AccountAIRunCfMetaExecuteLlama2_7bChatInt8ParamsBodyMessagesResponseFormatTypeJsonObject, AccountAIRunCfMetaExecuteLlama2_7bChatInt8ParamsBodyMessagesResponseFormatTypeJsonSchema:
		return true
	}
	return false
}

type AccountAIRunCfMetaExecuteLlama2_7bChatInt8ParamsBodyMessagesTool struct {
	// A brief description of what the tool does.
	Description param.Field[string]      `json:"description"`
	Function    param.Field[interface{}] `json:"function"`
	// The name of the tool. More descriptive the better.
	Name       param.Field[string]      `json:"name"`
	Parameters param.Field[interface{}] `json:"parameters"`
	// Specifies the type of tool (e.g., 'function').
	Type param.Field[string] `json:"type"`
}

func (r AccountAIRunCfMetaExecuteLlama2_7bChatInt8ParamsBodyMessagesTool) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfMetaExecuteLlama2_7bChatInt8ParamsBodyMessagesTool) implementsAccountAIRunCfMetaExecuteLlama2_7bChatInt8ParamsBodyMessagesToolUnion() {
}

// Satisfied by
// [AccountAIRunCfMetaExecuteLlama2_7bChatInt8ParamsBodyMessagesToolsObject],
// [AccountAIRunCfMetaExecuteLlama2_7bChatInt8ParamsBodyMessagesToolsObject],
// [AccountAIRunCfMetaExecuteLlama2_7bChatInt8ParamsBodyMessagesTool].
type AccountAIRunCfMetaExecuteLlama2_7bChatInt8ParamsBodyMessagesToolUnion interface {
	implementsAccountAIRunCfMetaExecuteLlama2_7bChatInt8ParamsBodyMessagesToolUnion()
}

type AccountAIRunCfMetaExecuteLlama2_7bChatInt8ParamsBodyMessagesToolsObject struct {
	// A brief description of what the tool does.
	Description param.Field[string] `json:"description,required"`
	// The name of the tool. More descriptive the better.
	Name param.Field[string] `json:"name,required"`
	// Schema defining the parameters accepted by the tool.
	Parameters param.Field[AccountAIRunCfMetaExecuteLlama2_7bChatInt8ParamsBodyMessagesToolsObjectParameters] `json:"parameters,required"`
}

func (r AccountAIRunCfMetaExecuteLlama2_7bChatInt8ParamsBodyMessagesToolsObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfMetaExecuteLlama2_7bChatInt8ParamsBodyMessagesToolsObject) implementsAccountAIRunCfMetaExecuteLlama2_7bChatInt8ParamsBodyMessagesToolUnion() {
}

// Schema defining the parameters accepted by the tool.
type AccountAIRunCfMetaExecuteLlama2_7bChatInt8ParamsBodyMessagesToolsObjectParameters struct {
	// Definitions of each parameter.
	Properties param.Field[map[string]AccountAIRunCfMetaExecuteLlama2_7bChatInt8ParamsBodyMessagesToolsObjectParametersProperties] `json:"properties,required"`
	// The type of the parameters object (usually 'object').
	Type param.Field[string] `json:"type,required"`
	// List of required parameter names.
	Required param.Field[[]string] `json:"required"`
}

func (r AccountAIRunCfMetaExecuteLlama2_7bChatInt8ParamsBodyMessagesToolsObjectParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfMetaExecuteLlama2_7bChatInt8ParamsBodyMessagesToolsObjectParametersProperties struct {
	// A description of the expected parameter.
	Description param.Field[string] `json:"description,required"`
	// The data type of the parameter.
	Type param.Field[string] `json:"type,required"`
}

func (r AccountAIRunCfMetaExecuteLlama2_7bChatInt8ParamsBodyMessagesToolsObjectParametersProperties) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfMetaExecuteLlama3_1_70bInstructParams struct {
	QueueRequest param.Field[string]                                          `query:"queueRequest"`
	Body         AccountAIRunCfMetaExecuteLlama3_1_70bInstructParamsBodyUnion `json:"body"`
}

func (r AccountAIRunCfMetaExecuteLlama3_1_70bInstructParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

// URLQuery serializes [AccountAIRunCfMetaExecuteLlama3_1_70bInstructParams]'s
// query parameters as `url.Values`.
func (r AccountAIRunCfMetaExecuteLlama3_1_70bInstructParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountAIRunCfMetaExecuteLlama3_1_70bInstructParamsBody struct {
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

func (r AccountAIRunCfMetaExecuteLlama3_1_70bInstructParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfMetaExecuteLlama3_1_70bInstructParamsBody) implementsAccountAIRunCfMetaExecuteLlama3_1_70bInstructParamsBodyUnion() {
}

// Satisfied by [AccountAIRunCfMetaExecuteLlama3_1_70bInstructParamsBodyPrompt],
// [AccountAIRunCfMetaExecuteLlama3_1_70bInstructParamsBodyMessages],
// [AccountAIRunCfMetaExecuteLlama3_1_70bInstructParamsBody].
type AccountAIRunCfMetaExecuteLlama3_1_70bInstructParamsBodyUnion interface {
	implementsAccountAIRunCfMetaExecuteLlama3_1_70bInstructParamsBodyUnion()
}

type AccountAIRunCfMetaExecuteLlama3_1_70bInstructParamsBodyPrompt struct {
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
	RepetitionPenalty param.Field[float64]                                                                     `json:"repetition_penalty"`
	ResponseFormat    param.Field[AccountAIRunCfMetaExecuteLlama3_1_70bInstructParamsBodyPromptResponseFormat] `json:"response_format"`
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

func (r AccountAIRunCfMetaExecuteLlama3_1_70bInstructParamsBodyPrompt) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfMetaExecuteLlama3_1_70bInstructParamsBodyPrompt) implementsAccountAIRunCfMetaExecuteLlama3_1_70bInstructParamsBodyUnion() {
}

type AccountAIRunCfMetaExecuteLlama3_1_70bInstructParamsBodyPromptResponseFormat struct {
	JsonSchema param.Field[interface{}]                                                                     `json:"json_schema"`
	Type       param.Field[AccountAIRunCfMetaExecuteLlama3_1_70bInstructParamsBodyPromptResponseFormatType] `json:"type"`
}

func (r AccountAIRunCfMetaExecuteLlama3_1_70bInstructParamsBodyPromptResponseFormat) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfMetaExecuteLlama3_1_70bInstructParamsBodyPromptResponseFormatType string

const (
	AccountAIRunCfMetaExecuteLlama3_1_70bInstructParamsBodyPromptResponseFormatTypeJsonObject AccountAIRunCfMetaExecuteLlama3_1_70bInstructParamsBodyPromptResponseFormatType = "json_object"
	AccountAIRunCfMetaExecuteLlama3_1_70bInstructParamsBodyPromptResponseFormatTypeJsonSchema AccountAIRunCfMetaExecuteLlama3_1_70bInstructParamsBodyPromptResponseFormatType = "json_schema"
)

func (r AccountAIRunCfMetaExecuteLlama3_1_70bInstructParamsBodyPromptResponseFormatType) IsKnown() bool {
	switch r {
	case AccountAIRunCfMetaExecuteLlama3_1_70bInstructParamsBodyPromptResponseFormatTypeJsonObject, AccountAIRunCfMetaExecuteLlama3_1_70bInstructParamsBodyPromptResponseFormatTypeJsonSchema:
		return true
	}
	return false
}

type AccountAIRunCfMetaExecuteLlama3_1_70bInstructParamsBodyMessages struct {
	// An array of message objects representing the conversation history.
	Messages param.Field[[]AccountAIRunCfMetaExecuteLlama3_1_70bInstructParamsBodyMessagesMessage] `json:"messages,required"`
	// Decreases the likelihood of the model repeating the same lines verbatim.
	FrequencyPenalty param.Field[float64]                                                                   `json:"frequency_penalty"`
	Functions        param.Field[[]AccountAIRunCfMetaExecuteLlama3_1_70bInstructParamsBodyMessagesFunction] `json:"functions"`
	// The maximum number of tokens to generate in the response.
	MaxTokens param.Field[int64] `json:"max_tokens"`
	// Increases the likelihood of the model introducing new topics.
	PresencePenalty param.Field[float64] `json:"presence_penalty"`
	// If true, a chat template is not applied and you must adhere to the specific
	// model's expected formatting.
	Raw param.Field[bool] `json:"raw"`
	// Penalty for repeated tokens; higher values discourage repetition.
	RepetitionPenalty param.Field[float64]                                                                       `json:"repetition_penalty"`
	ResponseFormat    param.Field[AccountAIRunCfMetaExecuteLlama3_1_70bInstructParamsBodyMessagesResponseFormat] `json:"response_format"`
	// Random seed for reproducibility of the generation.
	Seed param.Field[int64] `json:"seed"`
	// If true, the response will be streamed back incrementally using SSE, Server Sent
	// Events.
	Stream param.Field[bool] `json:"stream"`
	// Controls the randomness of the output; higher values produce more random
	// results.
	Temperature param.Field[float64] `json:"temperature"`
	// A list of tools available for the assistant to use.
	Tools param.Field[[]AccountAIRunCfMetaExecuteLlama3_1_70bInstructParamsBodyMessagesToolUnion] `json:"tools"`
	// Limits the AI to choose from the top 'k' most probable words. Lower values make
	// responses more focused; higher values introduce more variety and potential
	// surprises.
	TopK param.Field[int64] `json:"top_k"`
	// Adjusts the creativity of the AI's responses by controlling how many possible
	// words it considers. Lower values make outputs more predictable; higher values
	// allow for more varied and creative responses.
	TopP param.Field[float64] `json:"top_p"`
}

func (r AccountAIRunCfMetaExecuteLlama3_1_70bInstructParamsBodyMessages) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfMetaExecuteLlama3_1_70bInstructParamsBodyMessages) implementsAccountAIRunCfMetaExecuteLlama3_1_70bInstructParamsBodyUnion() {
}

type AccountAIRunCfMetaExecuteLlama3_1_70bInstructParamsBodyMessagesMessage struct {
	// The content of the message as a string.
	Content param.Field[string] `json:"content,required"`
	// The role of the message sender (e.g., 'user', 'assistant', 'system', 'tool').
	Role param.Field[string] `json:"role,required"`
}

func (r AccountAIRunCfMetaExecuteLlama3_1_70bInstructParamsBodyMessagesMessage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfMetaExecuteLlama3_1_70bInstructParamsBodyMessagesFunction struct {
	Code param.Field[string] `json:"code,required"`
	Name param.Field[string] `json:"name,required"`
}

func (r AccountAIRunCfMetaExecuteLlama3_1_70bInstructParamsBodyMessagesFunction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfMetaExecuteLlama3_1_70bInstructParamsBodyMessagesResponseFormat struct {
	JsonSchema param.Field[interface{}]                                                                       `json:"json_schema"`
	Type       param.Field[AccountAIRunCfMetaExecuteLlama3_1_70bInstructParamsBodyMessagesResponseFormatType] `json:"type"`
}

func (r AccountAIRunCfMetaExecuteLlama3_1_70bInstructParamsBodyMessagesResponseFormat) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfMetaExecuteLlama3_1_70bInstructParamsBodyMessagesResponseFormatType string

const (
	AccountAIRunCfMetaExecuteLlama3_1_70bInstructParamsBodyMessagesResponseFormatTypeJsonObject AccountAIRunCfMetaExecuteLlama3_1_70bInstructParamsBodyMessagesResponseFormatType = "json_object"
	AccountAIRunCfMetaExecuteLlama3_1_70bInstructParamsBodyMessagesResponseFormatTypeJsonSchema AccountAIRunCfMetaExecuteLlama3_1_70bInstructParamsBodyMessagesResponseFormatType = "json_schema"
)

func (r AccountAIRunCfMetaExecuteLlama3_1_70bInstructParamsBodyMessagesResponseFormatType) IsKnown() bool {
	switch r {
	case AccountAIRunCfMetaExecuteLlama3_1_70bInstructParamsBodyMessagesResponseFormatTypeJsonObject, AccountAIRunCfMetaExecuteLlama3_1_70bInstructParamsBodyMessagesResponseFormatTypeJsonSchema:
		return true
	}
	return false
}

type AccountAIRunCfMetaExecuteLlama3_1_70bInstructParamsBodyMessagesTool struct {
	// A brief description of what the tool does.
	Description param.Field[string]      `json:"description"`
	Function    param.Field[interface{}] `json:"function"`
	// The name of the tool. More descriptive the better.
	Name       param.Field[string]      `json:"name"`
	Parameters param.Field[interface{}] `json:"parameters"`
	// Specifies the type of tool (e.g., 'function').
	Type param.Field[string] `json:"type"`
}

func (r AccountAIRunCfMetaExecuteLlama3_1_70bInstructParamsBodyMessagesTool) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfMetaExecuteLlama3_1_70bInstructParamsBodyMessagesTool) implementsAccountAIRunCfMetaExecuteLlama3_1_70bInstructParamsBodyMessagesToolUnion() {
}

// Satisfied by
// [AccountAIRunCfMetaExecuteLlama3_1_70bInstructParamsBodyMessagesToolsObject],
// [AccountAIRunCfMetaExecuteLlama3_1_70bInstructParamsBodyMessagesToolsObject],
// [AccountAIRunCfMetaExecuteLlama3_1_70bInstructParamsBodyMessagesTool].
type AccountAIRunCfMetaExecuteLlama3_1_70bInstructParamsBodyMessagesToolUnion interface {
	implementsAccountAIRunCfMetaExecuteLlama3_1_70bInstructParamsBodyMessagesToolUnion()
}

type AccountAIRunCfMetaExecuteLlama3_1_70bInstructParamsBodyMessagesToolsObject struct {
	// A brief description of what the tool does.
	Description param.Field[string] `json:"description,required"`
	// The name of the tool. More descriptive the better.
	Name param.Field[string] `json:"name,required"`
	// Schema defining the parameters accepted by the tool.
	Parameters param.Field[AccountAIRunCfMetaExecuteLlama3_1_70bInstructParamsBodyMessagesToolsObjectParameters] `json:"parameters,required"`
}

func (r AccountAIRunCfMetaExecuteLlama3_1_70bInstructParamsBodyMessagesToolsObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfMetaExecuteLlama3_1_70bInstructParamsBodyMessagesToolsObject) implementsAccountAIRunCfMetaExecuteLlama3_1_70bInstructParamsBodyMessagesToolUnion() {
}

// Schema defining the parameters accepted by the tool.
type AccountAIRunCfMetaExecuteLlama3_1_70bInstructParamsBodyMessagesToolsObjectParameters struct {
	// Definitions of each parameter.
	Properties param.Field[map[string]AccountAIRunCfMetaExecuteLlama3_1_70bInstructParamsBodyMessagesToolsObjectParametersProperties] `json:"properties,required"`
	// The type of the parameters object (usually 'object').
	Type param.Field[string] `json:"type,required"`
	// List of required parameter names.
	Required param.Field[[]string] `json:"required"`
}

func (r AccountAIRunCfMetaExecuteLlama3_1_70bInstructParamsBodyMessagesToolsObjectParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfMetaExecuteLlama3_1_70bInstructParamsBodyMessagesToolsObjectParametersProperties struct {
	// A description of the expected parameter.
	Description param.Field[string] `json:"description,required"`
	// The data type of the parameter.
	Type param.Field[string] `json:"type,required"`
}

func (r AccountAIRunCfMetaExecuteLlama3_1_70bInstructParamsBodyMessagesToolsObjectParametersProperties) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfMetaExecuteLlama3_1_70bInstructPreviewParams struct {
	QueueRequest param.Field[string]                                                 `query:"queueRequest"`
	Body         AccountAIRunCfMetaExecuteLlama3_1_70bInstructPreviewParamsBodyUnion `json:"body"`
}

func (r AccountAIRunCfMetaExecuteLlama3_1_70bInstructPreviewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

// URLQuery serializes
// [AccountAIRunCfMetaExecuteLlama3_1_70bInstructPreviewParams]'s query parameters
// as `url.Values`.
func (r AccountAIRunCfMetaExecuteLlama3_1_70bInstructPreviewParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountAIRunCfMetaExecuteLlama3_1_70bInstructPreviewParamsBody struct {
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

func (r AccountAIRunCfMetaExecuteLlama3_1_70bInstructPreviewParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfMetaExecuteLlama3_1_70bInstructPreviewParamsBody) implementsAccountAIRunCfMetaExecuteLlama3_1_70bInstructPreviewParamsBodyUnion() {
}

// Satisfied by
// [AccountAIRunCfMetaExecuteLlama3_1_70bInstructPreviewParamsBodyPrompt],
// [AccountAIRunCfMetaExecuteLlama3_1_70bInstructPreviewParamsBodyMessages],
// [AccountAIRunCfMetaExecuteLlama3_1_70bInstructPreviewParamsBody].
type AccountAIRunCfMetaExecuteLlama3_1_70bInstructPreviewParamsBodyUnion interface {
	implementsAccountAIRunCfMetaExecuteLlama3_1_70bInstructPreviewParamsBodyUnion()
}

type AccountAIRunCfMetaExecuteLlama3_1_70bInstructPreviewParamsBodyPrompt struct {
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
	ResponseFormat    param.Field[AccountAIRunCfMetaExecuteLlama3_1_70bInstructPreviewParamsBodyPromptResponseFormat] `json:"response_format"`
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

func (r AccountAIRunCfMetaExecuteLlama3_1_70bInstructPreviewParamsBodyPrompt) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfMetaExecuteLlama3_1_70bInstructPreviewParamsBodyPrompt) implementsAccountAIRunCfMetaExecuteLlama3_1_70bInstructPreviewParamsBodyUnion() {
}

type AccountAIRunCfMetaExecuteLlama3_1_70bInstructPreviewParamsBodyPromptResponseFormat struct {
	JsonSchema param.Field[interface{}]                                                                            `json:"json_schema"`
	Type       param.Field[AccountAIRunCfMetaExecuteLlama3_1_70bInstructPreviewParamsBodyPromptResponseFormatType] `json:"type"`
}

func (r AccountAIRunCfMetaExecuteLlama3_1_70bInstructPreviewParamsBodyPromptResponseFormat) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfMetaExecuteLlama3_1_70bInstructPreviewParamsBodyPromptResponseFormatType string

const (
	AccountAIRunCfMetaExecuteLlama3_1_70bInstructPreviewParamsBodyPromptResponseFormatTypeJsonObject AccountAIRunCfMetaExecuteLlama3_1_70bInstructPreviewParamsBodyPromptResponseFormatType = "json_object"
	AccountAIRunCfMetaExecuteLlama3_1_70bInstructPreviewParamsBodyPromptResponseFormatTypeJsonSchema AccountAIRunCfMetaExecuteLlama3_1_70bInstructPreviewParamsBodyPromptResponseFormatType = "json_schema"
)

func (r AccountAIRunCfMetaExecuteLlama3_1_70bInstructPreviewParamsBodyPromptResponseFormatType) IsKnown() bool {
	switch r {
	case AccountAIRunCfMetaExecuteLlama3_1_70bInstructPreviewParamsBodyPromptResponseFormatTypeJsonObject, AccountAIRunCfMetaExecuteLlama3_1_70bInstructPreviewParamsBodyPromptResponseFormatTypeJsonSchema:
		return true
	}
	return false
}

type AccountAIRunCfMetaExecuteLlama3_1_70bInstructPreviewParamsBodyMessages struct {
	// An array of message objects representing the conversation history.
	Messages param.Field[[]AccountAIRunCfMetaExecuteLlama3_1_70bInstructPreviewParamsBodyMessagesMessage] `json:"messages,required"`
	// Decreases the likelihood of the model repeating the same lines verbatim.
	FrequencyPenalty param.Field[float64]                                                                          `json:"frequency_penalty"`
	Functions        param.Field[[]AccountAIRunCfMetaExecuteLlama3_1_70bInstructPreviewParamsBodyMessagesFunction] `json:"functions"`
	// The maximum number of tokens to generate in the response.
	MaxTokens param.Field[int64] `json:"max_tokens"`
	// Increases the likelihood of the model introducing new topics.
	PresencePenalty param.Field[float64] `json:"presence_penalty"`
	// If true, a chat template is not applied and you must adhere to the specific
	// model's expected formatting.
	Raw param.Field[bool] `json:"raw"`
	// Penalty for repeated tokens; higher values discourage repetition.
	RepetitionPenalty param.Field[float64]                                                                              `json:"repetition_penalty"`
	ResponseFormat    param.Field[AccountAIRunCfMetaExecuteLlama3_1_70bInstructPreviewParamsBodyMessagesResponseFormat] `json:"response_format"`
	// Random seed for reproducibility of the generation.
	Seed param.Field[int64] `json:"seed"`
	// If true, the response will be streamed back incrementally using SSE, Server Sent
	// Events.
	Stream param.Field[bool] `json:"stream"`
	// Controls the randomness of the output; higher values produce more random
	// results.
	Temperature param.Field[float64] `json:"temperature"`
	// A list of tools available for the assistant to use.
	Tools param.Field[[]AccountAIRunCfMetaExecuteLlama3_1_70bInstructPreviewParamsBodyMessagesToolUnion] `json:"tools"`
	// Limits the AI to choose from the top 'k' most probable words. Lower values make
	// responses more focused; higher values introduce more variety and potential
	// surprises.
	TopK param.Field[int64] `json:"top_k"`
	// Adjusts the creativity of the AI's responses by controlling how many possible
	// words it considers. Lower values make outputs more predictable; higher values
	// allow for more varied and creative responses.
	TopP param.Field[float64] `json:"top_p"`
}

func (r AccountAIRunCfMetaExecuteLlama3_1_70bInstructPreviewParamsBodyMessages) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfMetaExecuteLlama3_1_70bInstructPreviewParamsBodyMessages) implementsAccountAIRunCfMetaExecuteLlama3_1_70bInstructPreviewParamsBodyUnion() {
}

type AccountAIRunCfMetaExecuteLlama3_1_70bInstructPreviewParamsBodyMessagesMessage struct {
	// The content of the message as a string.
	Content param.Field[string] `json:"content,required"`
	// The role of the message sender (e.g., 'user', 'assistant', 'system', 'tool').
	Role param.Field[string] `json:"role,required"`
}

func (r AccountAIRunCfMetaExecuteLlama3_1_70bInstructPreviewParamsBodyMessagesMessage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfMetaExecuteLlama3_1_70bInstructPreviewParamsBodyMessagesFunction struct {
	Code param.Field[string] `json:"code,required"`
	Name param.Field[string] `json:"name,required"`
}

func (r AccountAIRunCfMetaExecuteLlama3_1_70bInstructPreviewParamsBodyMessagesFunction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfMetaExecuteLlama3_1_70bInstructPreviewParamsBodyMessagesResponseFormat struct {
	JsonSchema param.Field[interface{}]                                                                              `json:"json_schema"`
	Type       param.Field[AccountAIRunCfMetaExecuteLlama3_1_70bInstructPreviewParamsBodyMessagesResponseFormatType] `json:"type"`
}

func (r AccountAIRunCfMetaExecuteLlama3_1_70bInstructPreviewParamsBodyMessagesResponseFormat) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfMetaExecuteLlama3_1_70bInstructPreviewParamsBodyMessagesResponseFormatType string

const (
	AccountAIRunCfMetaExecuteLlama3_1_70bInstructPreviewParamsBodyMessagesResponseFormatTypeJsonObject AccountAIRunCfMetaExecuteLlama3_1_70bInstructPreviewParamsBodyMessagesResponseFormatType = "json_object"
	AccountAIRunCfMetaExecuteLlama3_1_70bInstructPreviewParamsBodyMessagesResponseFormatTypeJsonSchema AccountAIRunCfMetaExecuteLlama3_1_70bInstructPreviewParamsBodyMessagesResponseFormatType = "json_schema"
)

func (r AccountAIRunCfMetaExecuteLlama3_1_70bInstructPreviewParamsBodyMessagesResponseFormatType) IsKnown() bool {
	switch r {
	case AccountAIRunCfMetaExecuteLlama3_1_70bInstructPreviewParamsBodyMessagesResponseFormatTypeJsonObject, AccountAIRunCfMetaExecuteLlama3_1_70bInstructPreviewParamsBodyMessagesResponseFormatTypeJsonSchema:
		return true
	}
	return false
}

type AccountAIRunCfMetaExecuteLlama3_1_70bInstructPreviewParamsBodyMessagesTool struct {
	// A brief description of what the tool does.
	Description param.Field[string]      `json:"description"`
	Function    param.Field[interface{}] `json:"function"`
	// The name of the tool. More descriptive the better.
	Name       param.Field[string]      `json:"name"`
	Parameters param.Field[interface{}] `json:"parameters"`
	// Specifies the type of tool (e.g., 'function').
	Type param.Field[string] `json:"type"`
}

func (r AccountAIRunCfMetaExecuteLlama3_1_70bInstructPreviewParamsBodyMessagesTool) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfMetaExecuteLlama3_1_70bInstructPreviewParamsBodyMessagesTool) implementsAccountAIRunCfMetaExecuteLlama3_1_70bInstructPreviewParamsBodyMessagesToolUnion() {
}

// Satisfied by
// [AccountAIRunCfMetaExecuteLlama3_1_70bInstructPreviewParamsBodyMessagesToolsObject],
// [AccountAIRunCfMetaExecuteLlama3_1_70bInstructPreviewParamsBodyMessagesToolsObject],
// [AccountAIRunCfMetaExecuteLlama3_1_70bInstructPreviewParamsBodyMessagesTool].
type AccountAIRunCfMetaExecuteLlama3_1_70bInstructPreviewParamsBodyMessagesToolUnion interface {
	implementsAccountAIRunCfMetaExecuteLlama3_1_70bInstructPreviewParamsBodyMessagesToolUnion()
}

type AccountAIRunCfMetaExecuteLlama3_1_70bInstructPreviewParamsBodyMessagesToolsObject struct {
	// A brief description of what the tool does.
	Description param.Field[string] `json:"description,required"`
	// The name of the tool. More descriptive the better.
	Name param.Field[string] `json:"name,required"`
	// Schema defining the parameters accepted by the tool.
	Parameters param.Field[AccountAIRunCfMetaExecuteLlama3_1_70bInstructPreviewParamsBodyMessagesToolsObjectParameters] `json:"parameters,required"`
}

func (r AccountAIRunCfMetaExecuteLlama3_1_70bInstructPreviewParamsBodyMessagesToolsObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfMetaExecuteLlama3_1_70bInstructPreviewParamsBodyMessagesToolsObject) implementsAccountAIRunCfMetaExecuteLlama3_1_70bInstructPreviewParamsBodyMessagesToolUnion() {
}

// Schema defining the parameters accepted by the tool.
type AccountAIRunCfMetaExecuteLlama3_1_70bInstructPreviewParamsBodyMessagesToolsObjectParameters struct {
	// Definitions of each parameter.
	Properties param.Field[map[string]AccountAIRunCfMetaExecuteLlama3_1_70bInstructPreviewParamsBodyMessagesToolsObjectParametersProperties] `json:"properties,required"`
	// The type of the parameters object (usually 'object').
	Type param.Field[string] `json:"type,required"`
	// List of required parameter names.
	Required param.Field[[]string] `json:"required"`
}

func (r AccountAIRunCfMetaExecuteLlama3_1_70bInstructPreviewParamsBodyMessagesToolsObjectParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfMetaExecuteLlama3_1_70bInstructPreviewParamsBodyMessagesToolsObjectParametersProperties struct {
	// A description of the expected parameter.
	Description param.Field[string] `json:"description,required"`
	// The data type of the parameter.
	Type param.Field[string] `json:"type,required"`
}

func (r AccountAIRunCfMetaExecuteLlama3_1_70bInstructPreviewParamsBodyMessagesToolsObjectParametersProperties) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfMetaExecuteLlama3_1_70bPreviewParams struct {
	QueueRequest param.Field[string]                                         `query:"queueRequest"`
	Body         AccountAIRunCfMetaExecuteLlama3_1_70bPreviewParamsBodyUnion `json:"body"`
}

func (r AccountAIRunCfMetaExecuteLlama3_1_70bPreviewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

// URLQuery serializes [AccountAIRunCfMetaExecuteLlama3_1_70bPreviewParams]'s query
// parameters as `url.Values`.
func (r AccountAIRunCfMetaExecuteLlama3_1_70bPreviewParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountAIRunCfMetaExecuteLlama3_1_70bPreviewParamsBody struct {
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

func (r AccountAIRunCfMetaExecuteLlama3_1_70bPreviewParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfMetaExecuteLlama3_1_70bPreviewParamsBody) implementsAccountAIRunCfMetaExecuteLlama3_1_70bPreviewParamsBodyUnion() {
}

// Satisfied by [AccountAIRunCfMetaExecuteLlama3_1_70bPreviewParamsBodyPrompt],
// [AccountAIRunCfMetaExecuteLlama3_1_70bPreviewParamsBodyMessages],
// [AccountAIRunCfMetaExecuteLlama3_1_70bPreviewParamsBody].
type AccountAIRunCfMetaExecuteLlama3_1_70bPreviewParamsBodyUnion interface {
	implementsAccountAIRunCfMetaExecuteLlama3_1_70bPreviewParamsBodyUnion()
}

type AccountAIRunCfMetaExecuteLlama3_1_70bPreviewParamsBodyPrompt struct {
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
	ResponseFormat    param.Field[AccountAIRunCfMetaExecuteLlama3_1_70bPreviewParamsBodyPromptResponseFormat] `json:"response_format"`
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

func (r AccountAIRunCfMetaExecuteLlama3_1_70bPreviewParamsBodyPrompt) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfMetaExecuteLlama3_1_70bPreviewParamsBodyPrompt) implementsAccountAIRunCfMetaExecuteLlama3_1_70bPreviewParamsBodyUnion() {
}

type AccountAIRunCfMetaExecuteLlama3_1_70bPreviewParamsBodyPromptResponseFormat struct {
	JsonSchema param.Field[interface{}]                                                                    `json:"json_schema"`
	Type       param.Field[AccountAIRunCfMetaExecuteLlama3_1_70bPreviewParamsBodyPromptResponseFormatType] `json:"type"`
}

func (r AccountAIRunCfMetaExecuteLlama3_1_70bPreviewParamsBodyPromptResponseFormat) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfMetaExecuteLlama3_1_70bPreviewParamsBodyPromptResponseFormatType string

const (
	AccountAIRunCfMetaExecuteLlama3_1_70bPreviewParamsBodyPromptResponseFormatTypeJsonObject AccountAIRunCfMetaExecuteLlama3_1_70bPreviewParamsBodyPromptResponseFormatType = "json_object"
	AccountAIRunCfMetaExecuteLlama3_1_70bPreviewParamsBodyPromptResponseFormatTypeJsonSchema AccountAIRunCfMetaExecuteLlama3_1_70bPreviewParamsBodyPromptResponseFormatType = "json_schema"
)

func (r AccountAIRunCfMetaExecuteLlama3_1_70bPreviewParamsBodyPromptResponseFormatType) IsKnown() bool {
	switch r {
	case AccountAIRunCfMetaExecuteLlama3_1_70bPreviewParamsBodyPromptResponseFormatTypeJsonObject, AccountAIRunCfMetaExecuteLlama3_1_70bPreviewParamsBodyPromptResponseFormatTypeJsonSchema:
		return true
	}
	return false
}

type AccountAIRunCfMetaExecuteLlama3_1_70bPreviewParamsBodyMessages struct {
	// An array of message objects representing the conversation history.
	Messages param.Field[[]AccountAIRunCfMetaExecuteLlama3_1_70bPreviewParamsBodyMessagesMessage] `json:"messages,required"`
	// Decreases the likelihood of the model repeating the same lines verbatim.
	FrequencyPenalty param.Field[float64]                                                                  `json:"frequency_penalty"`
	Functions        param.Field[[]AccountAIRunCfMetaExecuteLlama3_1_70bPreviewParamsBodyMessagesFunction] `json:"functions"`
	// The maximum number of tokens to generate in the response.
	MaxTokens param.Field[int64] `json:"max_tokens"`
	// Increases the likelihood of the model introducing new topics.
	PresencePenalty param.Field[float64] `json:"presence_penalty"`
	// If true, a chat template is not applied and you must adhere to the specific
	// model's expected formatting.
	Raw param.Field[bool] `json:"raw"`
	// Penalty for repeated tokens; higher values discourage repetition.
	RepetitionPenalty param.Field[float64]                                                                      `json:"repetition_penalty"`
	ResponseFormat    param.Field[AccountAIRunCfMetaExecuteLlama3_1_70bPreviewParamsBodyMessagesResponseFormat] `json:"response_format"`
	// Random seed for reproducibility of the generation.
	Seed param.Field[int64] `json:"seed"`
	// If true, the response will be streamed back incrementally using SSE, Server Sent
	// Events.
	Stream param.Field[bool] `json:"stream"`
	// Controls the randomness of the output; higher values produce more random
	// results.
	Temperature param.Field[float64] `json:"temperature"`
	// A list of tools available for the assistant to use.
	Tools param.Field[[]AccountAIRunCfMetaExecuteLlama3_1_70bPreviewParamsBodyMessagesToolUnion] `json:"tools"`
	// Limits the AI to choose from the top 'k' most probable words. Lower values make
	// responses more focused; higher values introduce more variety and potential
	// surprises.
	TopK param.Field[int64] `json:"top_k"`
	// Adjusts the creativity of the AI's responses by controlling how many possible
	// words it considers. Lower values make outputs more predictable; higher values
	// allow for more varied and creative responses.
	TopP param.Field[float64] `json:"top_p"`
}

func (r AccountAIRunCfMetaExecuteLlama3_1_70bPreviewParamsBodyMessages) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfMetaExecuteLlama3_1_70bPreviewParamsBodyMessages) implementsAccountAIRunCfMetaExecuteLlama3_1_70bPreviewParamsBodyUnion() {
}

type AccountAIRunCfMetaExecuteLlama3_1_70bPreviewParamsBodyMessagesMessage struct {
	// The content of the message as a string.
	Content param.Field[string] `json:"content,required"`
	// The role of the message sender (e.g., 'user', 'assistant', 'system', 'tool').
	Role param.Field[string] `json:"role,required"`
}

func (r AccountAIRunCfMetaExecuteLlama3_1_70bPreviewParamsBodyMessagesMessage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfMetaExecuteLlama3_1_70bPreviewParamsBodyMessagesFunction struct {
	Code param.Field[string] `json:"code,required"`
	Name param.Field[string] `json:"name,required"`
}

func (r AccountAIRunCfMetaExecuteLlama3_1_70bPreviewParamsBodyMessagesFunction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfMetaExecuteLlama3_1_70bPreviewParamsBodyMessagesResponseFormat struct {
	JsonSchema param.Field[interface{}]                                                                      `json:"json_schema"`
	Type       param.Field[AccountAIRunCfMetaExecuteLlama3_1_70bPreviewParamsBodyMessagesResponseFormatType] `json:"type"`
}

func (r AccountAIRunCfMetaExecuteLlama3_1_70bPreviewParamsBodyMessagesResponseFormat) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfMetaExecuteLlama3_1_70bPreviewParamsBodyMessagesResponseFormatType string

const (
	AccountAIRunCfMetaExecuteLlama3_1_70bPreviewParamsBodyMessagesResponseFormatTypeJsonObject AccountAIRunCfMetaExecuteLlama3_1_70bPreviewParamsBodyMessagesResponseFormatType = "json_object"
	AccountAIRunCfMetaExecuteLlama3_1_70bPreviewParamsBodyMessagesResponseFormatTypeJsonSchema AccountAIRunCfMetaExecuteLlama3_1_70bPreviewParamsBodyMessagesResponseFormatType = "json_schema"
)

func (r AccountAIRunCfMetaExecuteLlama3_1_70bPreviewParamsBodyMessagesResponseFormatType) IsKnown() bool {
	switch r {
	case AccountAIRunCfMetaExecuteLlama3_1_70bPreviewParamsBodyMessagesResponseFormatTypeJsonObject, AccountAIRunCfMetaExecuteLlama3_1_70bPreviewParamsBodyMessagesResponseFormatTypeJsonSchema:
		return true
	}
	return false
}

type AccountAIRunCfMetaExecuteLlama3_1_70bPreviewParamsBodyMessagesTool struct {
	// A brief description of what the tool does.
	Description param.Field[string]      `json:"description"`
	Function    param.Field[interface{}] `json:"function"`
	// The name of the tool. More descriptive the better.
	Name       param.Field[string]      `json:"name"`
	Parameters param.Field[interface{}] `json:"parameters"`
	// Specifies the type of tool (e.g., 'function').
	Type param.Field[string] `json:"type"`
}

func (r AccountAIRunCfMetaExecuteLlama3_1_70bPreviewParamsBodyMessagesTool) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfMetaExecuteLlama3_1_70bPreviewParamsBodyMessagesTool) implementsAccountAIRunCfMetaExecuteLlama3_1_70bPreviewParamsBodyMessagesToolUnion() {
}

// Satisfied by
// [AccountAIRunCfMetaExecuteLlama3_1_70bPreviewParamsBodyMessagesToolsObject],
// [AccountAIRunCfMetaExecuteLlama3_1_70bPreviewParamsBodyMessagesToolsObject],
// [AccountAIRunCfMetaExecuteLlama3_1_70bPreviewParamsBodyMessagesTool].
type AccountAIRunCfMetaExecuteLlama3_1_70bPreviewParamsBodyMessagesToolUnion interface {
	implementsAccountAIRunCfMetaExecuteLlama3_1_70bPreviewParamsBodyMessagesToolUnion()
}

type AccountAIRunCfMetaExecuteLlama3_1_70bPreviewParamsBodyMessagesToolsObject struct {
	// A brief description of what the tool does.
	Description param.Field[string] `json:"description,required"`
	// The name of the tool. More descriptive the better.
	Name param.Field[string] `json:"name,required"`
	// Schema defining the parameters accepted by the tool.
	Parameters param.Field[AccountAIRunCfMetaExecuteLlama3_1_70bPreviewParamsBodyMessagesToolsObjectParameters] `json:"parameters,required"`
}

func (r AccountAIRunCfMetaExecuteLlama3_1_70bPreviewParamsBodyMessagesToolsObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfMetaExecuteLlama3_1_70bPreviewParamsBodyMessagesToolsObject) implementsAccountAIRunCfMetaExecuteLlama3_1_70bPreviewParamsBodyMessagesToolUnion() {
}

// Schema defining the parameters accepted by the tool.
type AccountAIRunCfMetaExecuteLlama3_1_70bPreviewParamsBodyMessagesToolsObjectParameters struct {
	// Definitions of each parameter.
	Properties param.Field[map[string]AccountAIRunCfMetaExecuteLlama3_1_70bPreviewParamsBodyMessagesToolsObjectParametersProperties] `json:"properties,required"`
	// The type of the parameters object (usually 'object').
	Type param.Field[string] `json:"type,required"`
	// List of required parameter names.
	Required param.Field[[]string] `json:"required"`
}

func (r AccountAIRunCfMetaExecuteLlama3_1_70bPreviewParamsBodyMessagesToolsObjectParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfMetaExecuteLlama3_1_70bPreviewParamsBodyMessagesToolsObjectParametersProperties struct {
	// A description of the expected parameter.
	Description param.Field[string] `json:"description,required"`
	// The data type of the parameter.
	Type param.Field[string] `json:"type,required"`
}

func (r AccountAIRunCfMetaExecuteLlama3_1_70bPreviewParamsBodyMessagesToolsObjectParametersProperties) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfMetaExecuteLlama3_1_8bInstructAwqParams struct {
	QueueRequest param.Field[string]                                            `query:"queueRequest"`
	Body         AccountAIRunCfMetaExecuteLlama3_1_8bInstructAwqParamsBodyUnion `json:"body"`
}

func (r AccountAIRunCfMetaExecuteLlama3_1_8bInstructAwqParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

// URLQuery serializes [AccountAIRunCfMetaExecuteLlama3_1_8bInstructAwqParams]'s
// query parameters as `url.Values`.
func (r AccountAIRunCfMetaExecuteLlama3_1_8bInstructAwqParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountAIRunCfMetaExecuteLlama3_1_8bInstructAwqParamsBody struct {
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

func (r AccountAIRunCfMetaExecuteLlama3_1_8bInstructAwqParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfMetaExecuteLlama3_1_8bInstructAwqParamsBody) implementsAccountAIRunCfMetaExecuteLlama3_1_8bInstructAwqParamsBodyUnion() {
}

// Satisfied by [AccountAIRunCfMetaExecuteLlama3_1_8bInstructAwqParamsBodyPrompt],
// [AccountAIRunCfMetaExecuteLlama3_1_8bInstructAwqParamsBodyMessages],
// [AccountAIRunCfMetaExecuteLlama3_1_8bInstructAwqParamsBody].
type AccountAIRunCfMetaExecuteLlama3_1_8bInstructAwqParamsBodyUnion interface {
	implementsAccountAIRunCfMetaExecuteLlama3_1_8bInstructAwqParamsBodyUnion()
}

type AccountAIRunCfMetaExecuteLlama3_1_8bInstructAwqParamsBodyPrompt struct {
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
	ResponseFormat    param.Field[AccountAIRunCfMetaExecuteLlama3_1_8bInstructAwqParamsBodyPromptResponseFormat] `json:"response_format"`
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

func (r AccountAIRunCfMetaExecuteLlama3_1_8bInstructAwqParamsBodyPrompt) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfMetaExecuteLlama3_1_8bInstructAwqParamsBodyPrompt) implementsAccountAIRunCfMetaExecuteLlama3_1_8bInstructAwqParamsBodyUnion() {
}

type AccountAIRunCfMetaExecuteLlama3_1_8bInstructAwqParamsBodyPromptResponseFormat struct {
	JsonSchema param.Field[interface{}]                                                                       `json:"json_schema"`
	Type       param.Field[AccountAIRunCfMetaExecuteLlama3_1_8bInstructAwqParamsBodyPromptResponseFormatType] `json:"type"`
}

func (r AccountAIRunCfMetaExecuteLlama3_1_8bInstructAwqParamsBodyPromptResponseFormat) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfMetaExecuteLlama3_1_8bInstructAwqParamsBodyPromptResponseFormatType string

const (
	AccountAIRunCfMetaExecuteLlama3_1_8bInstructAwqParamsBodyPromptResponseFormatTypeJsonObject AccountAIRunCfMetaExecuteLlama3_1_8bInstructAwqParamsBodyPromptResponseFormatType = "json_object"
	AccountAIRunCfMetaExecuteLlama3_1_8bInstructAwqParamsBodyPromptResponseFormatTypeJsonSchema AccountAIRunCfMetaExecuteLlama3_1_8bInstructAwqParamsBodyPromptResponseFormatType = "json_schema"
)

func (r AccountAIRunCfMetaExecuteLlama3_1_8bInstructAwqParamsBodyPromptResponseFormatType) IsKnown() bool {
	switch r {
	case AccountAIRunCfMetaExecuteLlama3_1_8bInstructAwqParamsBodyPromptResponseFormatTypeJsonObject, AccountAIRunCfMetaExecuteLlama3_1_8bInstructAwqParamsBodyPromptResponseFormatTypeJsonSchema:
		return true
	}
	return false
}

type AccountAIRunCfMetaExecuteLlama3_1_8bInstructAwqParamsBodyMessages struct {
	// An array of message objects representing the conversation history.
	Messages param.Field[[]AccountAIRunCfMetaExecuteLlama3_1_8bInstructAwqParamsBodyMessagesMessage] `json:"messages,required"`
	// Decreases the likelihood of the model repeating the same lines verbatim.
	FrequencyPenalty param.Field[float64]                                                                     `json:"frequency_penalty"`
	Functions        param.Field[[]AccountAIRunCfMetaExecuteLlama3_1_8bInstructAwqParamsBodyMessagesFunction] `json:"functions"`
	// The maximum number of tokens to generate in the response.
	MaxTokens param.Field[int64] `json:"max_tokens"`
	// Increases the likelihood of the model introducing new topics.
	PresencePenalty param.Field[float64] `json:"presence_penalty"`
	// If true, a chat template is not applied and you must adhere to the specific
	// model's expected formatting.
	Raw param.Field[bool] `json:"raw"`
	// Penalty for repeated tokens; higher values discourage repetition.
	RepetitionPenalty param.Field[float64]                                                                         `json:"repetition_penalty"`
	ResponseFormat    param.Field[AccountAIRunCfMetaExecuteLlama3_1_8bInstructAwqParamsBodyMessagesResponseFormat] `json:"response_format"`
	// Random seed for reproducibility of the generation.
	Seed param.Field[int64] `json:"seed"`
	// If true, the response will be streamed back incrementally using SSE, Server Sent
	// Events.
	Stream param.Field[bool] `json:"stream"`
	// Controls the randomness of the output; higher values produce more random
	// results.
	Temperature param.Field[float64] `json:"temperature"`
	// A list of tools available for the assistant to use.
	Tools param.Field[[]AccountAIRunCfMetaExecuteLlama3_1_8bInstructAwqParamsBodyMessagesToolUnion] `json:"tools"`
	// Limits the AI to choose from the top 'k' most probable words. Lower values make
	// responses more focused; higher values introduce more variety and potential
	// surprises.
	TopK param.Field[int64] `json:"top_k"`
	// Adjusts the creativity of the AI's responses by controlling how many possible
	// words it considers. Lower values make outputs more predictable; higher values
	// allow for more varied and creative responses.
	TopP param.Field[float64] `json:"top_p"`
}

func (r AccountAIRunCfMetaExecuteLlama3_1_8bInstructAwqParamsBodyMessages) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfMetaExecuteLlama3_1_8bInstructAwqParamsBodyMessages) implementsAccountAIRunCfMetaExecuteLlama3_1_8bInstructAwqParamsBodyUnion() {
}

type AccountAIRunCfMetaExecuteLlama3_1_8bInstructAwqParamsBodyMessagesMessage struct {
	// The content of the message as a string.
	Content param.Field[string] `json:"content,required"`
	// The role of the message sender (e.g., 'user', 'assistant', 'system', 'tool').
	Role param.Field[string] `json:"role,required"`
}

func (r AccountAIRunCfMetaExecuteLlama3_1_8bInstructAwqParamsBodyMessagesMessage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfMetaExecuteLlama3_1_8bInstructAwqParamsBodyMessagesFunction struct {
	Code param.Field[string] `json:"code,required"`
	Name param.Field[string] `json:"name,required"`
}

func (r AccountAIRunCfMetaExecuteLlama3_1_8bInstructAwqParamsBodyMessagesFunction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfMetaExecuteLlama3_1_8bInstructAwqParamsBodyMessagesResponseFormat struct {
	JsonSchema param.Field[interface{}]                                                                         `json:"json_schema"`
	Type       param.Field[AccountAIRunCfMetaExecuteLlama3_1_8bInstructAwqParamsBodyMessagesResponseFormatType] `json:"type"`
}

func (r AccountAIRunCfMetaExecuteLlama3_1_8bInstructAwqParamsBodyMessagesResponseFormat) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfMetaExecuteLlama3_1_8bInstructAwqParamsBodyMessagesResponseFormatType string

const (
	AccountAIRunCfMetaExecuteLlama3_1_8bInstructAwqParamsBodyMessagesResponseFormatTypeJsonObject AccountAIRunCfMetaExecuteLlama3_1_8bInstructAwqParamsBodyMessagesResponseFormatType = "json_object"
	AccountAIRunCfMetaExecuteLlama3_1_8bInstructAwqParamsBodyMessagesResponseFormatTypeJsonSchema AccountAIRunCfMetaExecuteLlama3_1_8bInstructAwqParamsBodyMessagesResponseFormatType = "json_schema"
)

func (r AccountAIRunCfMetaExecuteLlama3_1_8bInstructAwqParamsBodyMessagesResponseFormatType) IsKnown() bool {
	switch r {
	case AccountAIRunCfMetaExecuteLlama3_1_8bInstructAwqParamsBodyMessagesResponseFormatTypeJsonObject, AccountAIRunCfMetaExecuteLlama3_1_8bInstructAwqParamsBodyMessagesResponseFormatTypeJsonSchema:
		return true
	}
	return false
}

type AccountAIRunCfMetaExecuteLlama3_1_8bInstructAwqParamsBodyMessagesTool struct {
	// A brief description of what the tool does.
	Description param.Field[string]      `json:"description"`
	Function    param.Field[interface{}] `json:"function"`
	// The name of the tool. More descriptive the better.
	Name       param.Field[string]      `json:"name"`
	Parameters param.Field[interface{}] `json:"parameters"`
	// Specifies the type of tool (e.g., 'function').
	Type param.Field[string] `json:"type"`
}

func (r AccountAIRunCfMetaExecuteLlama3_1_8bInstructAwqParamsBodyMessagesTool) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfMetaExecuteLlama3_1_8bInstructAwqParamsBodyMessagesTool) implementsAccountAIRunCfMetaExecuteLlama3_1_8bInstructAwqParamsBodyMessagesToolUnion() {
}

// Satisfied by
// [AccountAIRunCfMetaExecuteLlama3_1_8bInstructAwqParamsBodyMessagesToolsObject],
// [AccountAIRunCfMetaExecuteLlama3_1_8bInstructAwqParamsBodyMessagesToolsObject],
// [AccountAIRunCfMetaExecuteLlama3_1_8bInstructAwqParamsBodyMessagesTool].
type AccountAIRunCfMetaExecuteLlama3_1_8bInstructAwqParamsBodyMessagesToolUnion interface {
	implementsAccountAIRunCfMetaExecuteLlama3_1_8bInstructAwqParamsBodyMessagesToolUnion()
}

type AccountAIRunCfMetaExecuteLlama3_1_8bInstructAwqParamsBodyMessagesToolsObject struct {
	// A brief description of what the tool does.
	Description param.Field[string] `json:"description,required"`
	// The name of the tool. More descriptive the better.
	Name param.Field[string] `json:"name,required"`
	// Schema defining the parameters accepted by the tool.
	Parameters param.Field[AccountAIRunCfMetaExecuteLlama3_1_8bInstructAwqParamsBodyMessagesToolsObjectParameters] `json:"parameters,required"`
}

func (r AccountAIRunCfMetaExecuteLlama3_1_8bInstructAwqParamsBodyMessagesToolsObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfMetaExecuteLlama3_1_8bInstructAwqParamsBodyMessagesToolsObject) implementsAccountAIRunCfMetaExecuteLlama3_1_8bInstructAwqParamsBodyMessagesToolUnion() {
}

// Schema defining the parameters accepted by the tool.
type AccountAIRunCfMetaExecuteLlama3_1_8bInstructAwqParamsBodyMessagesToolsObjectParameters struct {
	// Definitions of each parameter.
	Properties param.Field[map[string]AccountAIRunCfMetaExecuteLlama3_1_8bInstructAwqParamsBodyMessagesToolsObjectParametersProperties] `json:"properties,required"`
	// The type of the parameters object (usually 'object').
	Type param.Field[string] `json:"type,required"`
	// List of required parameter names.
	Required param.Field[[]string] `json:"required"`
}

func (r AccountAIRunCfMetaExecuteLlama3_1_8bInstructAwqParamsBodyMessagesToolsObjectParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfMetaExecuteLlama3_1_8bInstructAwqParamsBodyMessagesToolsObjectParametersProperties struct {
	// A description of the expected parameter.
	Description param.Field[string] `json:"description,required"`
	// The data type of the parameter.
	Type param.Field[string] `json:"type,required"`
}

func (r AccountAIRunCfMetaExecuteLlama3_1_8bInstructAwqParamsBodyMessagesToolsObjectParametersProperties) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfMetaExecuteLlama3_1_8bInstructFastParams struct {
	QueueRequest param.Field[string]                                             `query:"queueRequest"`
	Body         AccountAIRunCfMetaExecuteLlama3_1_8bInstructFastParamsBodyUnion `json:"body"`
}

func (r AccountAIRunCfMetaExecuteLlama3_1_8bInstructFastParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

// URLQuery serializes [AccountAIRunCfMetaExecuteLlama3_1_8bInstructFastParams]'s
// query parameters as `url.Values`.
func (r AccountAIRunCfMetaExecuteLlama3_1_8bInstructFastParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountAIRunCfMetaExecuteLlama3_1_8bInstructFastParamsBody struct {
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

func (r AccountAIRunCfMetaExecuteLlama3_1_8bInstructFastParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfMetaExecuteLlama3_1_8bInstructFastParamsBody) implementsAccountAIRunCfMetaExecuteLlama3_1_8bInstructFastParamsBodyUnion() {
}

// Satisfied by [AccountAIRunCfMetaExecuteLlama3_1_8bInstructFastParamsBodyPrompt],
// [AccountAIRunCfMetaExecuteLlama3_1_8bInstructFastParamsBodyMessages],
// [AccountAIRunCfMetaExecuteLlama3_1_8bInstructFastParamsBody].
type AccountAIRunCfMetaExecuteLlama3_1_8bInstructFastParamsBodyUnion interface {
	implementsAccountAIRunCfMetaExecuteLlama3_1_8bInstructFastParamsBodyUnion()
}

type AccountAIRunCfMetaExecuteLlama3_1_8bInstructFastParamsBodyPrompt struct {
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
	ResponseFormat    param.Field[AccountAIRunCfMetaExecuteLlama3_1_8bInstructFastParamsBodyPromptResponseFormat] `json:"response_format"`
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

func (r AccountAIRunCfMetaExecuteLlama3_1_8bInstructFastParamsBodyPrompt) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfMetaExecuteLlama3_1_8bInstructFastParamsBodyPrompt) implementsAccountAIRunCfMetaExecuteLlama3_1_8bInstructFastParamsBodyUnion() {
}

type AccountAIRunCfMetaExecuteLlama3_1_8bInstructFastParamsBodyPromptResponseFormat struct {
	JsonSchema param.Field[interface{}]                                                                        `json:"json_schema"`
	Type       param.Field[AccountAIRunCfMetaExecuteLlama3_1_8bInstructFastParamsBodyPromptResponseFormatType] `json:"type"`
}

func (r AccountAIRunCfMetaExecuteLlama3_1_8bInstructFastParamsBodyPromptResponseFormat) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfMetaExecuteLlama3_1_8bInstructFastParamsBodyPromptResponseFormatType string

const (
	AccountAIRunCfMetaExecuteLlama3_1_8bInstructFastParamsBodyPromptResponseFormatTypeJsonObject AccountAIRunCfMetaExecuteLlama3_1_8bInstructFastParamsBodyPromptResponseFormatType = "json_object"
	AccountAIRunCfMetaExecuteLlama3_1_8bInstructFastParamsBodyPromptResponseFormatTypeJsonSchema AccountAIRunCfMetaExecuteLlama3_1_8bInstructFastParamsBodyPromptResponseFormatType = "json_schema"
)

func (r AccountAIRunCfMetaExecuteLlama3_1_8bInstructFastParamsBodyPromptResponseFormatType) IsKnown() bool {
	switch r {
	case AccountAIRunCfMetaExecuteLlama3_1_8bInstructFastParamsBodyPromptResponseFormatTypeJsonObject, AccountAIRunCfMetaExecuteLlama3_1_8bInstructFastParamsBodyPromptResponseFormatTypeJsonSchema:
		return true
	}
	return false
}

type AccountAIRunCfMetaExecuteLlama3_1_8bInstructFastParamsBodyMessages struct {
	// An array of message objects representing the conversation history.
	Messages param.Field[[]AccountAIRunCfMetaExecuteLlama3_1_8bInstructFastParamsBodyMessagesMessage] `json:"messages,required"`
	// Decreases the likelihood of the model repeating the same lines verbatim.
	FrequencyPenalty param.Field[float64]                                                                      `json:"frequency_penalty"`
	Functions        param.Field[[]AccountAIRunCfMetaExecuteLlama3_1_8bInstructFastParamsBodyMessagesFunction] `json:"functions"`
	// The maximum number of tokens to generate in the response.
	MaxTokens param.Field[int64] `json:"max_tokens"`
	// Increases the likelihood of the model introducing new topics.
	PresencePenalty param.Field[float64] `json:"presence_penalty"`
	// If true, a chat template is not applied and you must adhere to the specific
	// model's expected formatting.
	Raw param.Field[bool] `json:"raw"`
	// Penalty for repeated tokens; higher values discourage repetition.
	RepetitionPenalty param.Field[float64]                                                                          `json:"repetition_penalty"`
	ResponseFormat    param.Field[AccountAIRunCfMetaExecuteLlama3_1_8bInstructFastParamsBodyMessagesResponseFormat] `json:"response_format"`
	// Random seed for reproducibility of the generation.
	Seed param.Field[int64] `json:"seed"`
	// If true, the response will be streamed back incrementally using SSE, Server Sent
	// Events.
	Stream param.Field[bool] `json:"stream"`
	// Controls the randomness of the output; higher values produce more random
	// results.
	Temperature param.Field[float64] `json:"temperature"`
	// A list of tools available for the assistant to use.
	Tools param.Field[[]AccountAIRunCfMetaExecuteLlama3_1_8bInstructFastParamsBodyMessagesToolUnion] `json:"tools"`
	// Limits the AI to choose from the top 'k' most probable words. Lower values make
	// responses more focused; higher values introduce more variety and potential
	// surprises.
	TopK param.Field[int64] `json:"top_k"`
	// Adjusts the creativity of the AI's responses by controlling how many possible
	// words it considers. Lower values make outputs more predictable; higher values
	// allow for more varied and creative responses.
	TopP param.Field[float64] `json:"top_p"`
}

func (r AccountAIRunCfMetaExecuteLlama3_1_8bInstructFastParamsBodyMessages) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfMetaExecuteLlama3_1_8bInstructFastParamsBodyMessages) implementsAccountAIRunCfMetaExecuteLlama3_1_8bInstructFastParamsBodyUnion() {
}

type AccountAIRunCfMetaExecuteLlama3_1_8bInstructFastParamsBodyMessagesMessage struct {
	// The content of the message as a string.
	Content param.Field[string] `json:"content,required"`
	// The role of the message sender (e.g., 'user', 'assistant', 'system', 'tool').
	Role param.Field[string] `json:"role,required"`
}

func (r AccountAIRunCfMetaExecuteLlama3_1_8bInstructFastParamsBodyMessagesMessage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfMetaExecuteLlama3_1_8bInstructFastParamsBodyMessagesFunction struct {
	Code param.Field[string] `json:"code,required"`
	Name param.Field[string] `json:"name,required"`
}

func (r AccountAIRunCfMetaExecuteLlama3_1_8bInstructFastParamsBodyMessagesFunction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfMetaExecuteLlama3_1_8bInstructFastParamsBodyMessagesResponseFormat struct {
	JsonSchema param.Field[interface{}]                                                                          `json:"json_schema"`
	Type       param.Field[AccountAIRunCfMetaExecuteLlama3_1_8bInstructFastParamsBodyMessagesResponseFormatType] `json:"type"`
}

func (r AccountAIRunCfMetaExecuteLlama3_1_8bInstructFastParamsBodyMessagesResponseFormat) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfMetaExecuteLlama3_1_8bInstructFastParamsBodyMessagesResponseFormatType string

const (
	AccountAIRunCfMetaExecuteLlama3_1_8bInstructFastParamsBodyMessagesResponseFormatTypeJsonObject AccountAIRunCfMetaExecuteLlama3_1_8bInstructFastParamsBodyMessagesResponseFormatType = "json_object"
	AccountAIRunCfMetaExecuteLlama3_1_8bInstructFastParamsBodyMessagesResponseFormatTypeJsonSchema AccountAIRunCfMetaExecuteLlama3_1_8bInstructFastParamsBodyMessagesResponseFormatType = "json_schema"
)

func (r AccountAIRunCfMetaExecuteLlama3_1_8bInstructFastParamsBodyMessagesResponseFormatType) IsKnown() bool {
	switch r {
	case AccountAIRunCfMetaExecuteLlama3_1_8bInstructFastParamsBodyMessagesResponseFormatTypeJsonObject, AccountAIRunCfMetaExecuteLlama3_1_8bInstructFastParamsBodyMessagesResponseFormatTypeJsonSchema:
		return true
	}
	return false
}

type AccountAIRunCfMetaExecuteLlama3_1_8bInstructFastParamsBodyMessagesTool struct {
	// A brief description of what the tool does.
	Description param.Field[string]      `json:"description"`
	Function    param.Field[interface{}] `json:"function"`
	// The name of the tool. More descriptive the better.
	Name       param.Field[string]      `json:"name"`
	Parameters param.Field[interface{}] `json:"parameters"`
	// Specifies the type of tool (e.g., 'function').
	Type param.Field[string] `json:"type"`
}

func (r AccountAIRunCfMetaExecuteLlama3_1_8bInstructFastParamsBodyMessagesTool) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfMetaExecuteLlama3_1_8bInstructFastParamsBodyMessagesTool) implementsAccountAIRunCfMetaExecuteLlama3_1_8bInstructFastParamsBodyMessagesToolUnion() {
}

// Satisfied by
// [AccountAIRunCfMetaExecuteLlama3_1_8bInstructFastParamsBodyMessagesToolsObject],
// [AccountAIRunCfMetaExecuteLlama3_1_8bInstructFastParamsBodyMessagesToolsObject],
// [AccountAIRunCfMetaExecuteLlama3_1_8bInstructFastParamsBodyMessagesTool].
type AccountAIRunCfMetaExecuteLlama3_1_8bInstructFastParamsBodyMessagesToolUnion interface {
	implementsAccountAIRunCfMetaExecuteLlama3_1_8bInstructFastParamsBodyMessagesToolUnion()
}

type AccountAIRunCfMetaExecuteLlama3_1_8bInstructFastParamsBodyMessagesToolsObject struct {
	// A brief description of what the tool does.
	Description param.Field[string] `json:"description,required"`
	// The name of the tool. More descriptive the better.
	Name param.Field[string] `json:"name,required"`
	// Schema defining the parameters accepted by the tool.
	Parameters param.Field[AccountAIRunCfMetaExecuteLlama3_1_8bInstructFastParamsBodyMessagesToolsObjectParameters] `json:"parameters,required"`
}

func (r AccountAIRunCfMetaExecuteLlama3_1_8bInstructFastParamsBodyMessagesToolsObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfMetaExecuteLlama3_1_8bInstructFastParamsBodyMessagesToolsObject) implementsAccountAIRunCfMetaExecuteLlama3_1_8bInstructFastParamsBodyMessagesToolUnion() {
}

// Schema defining the parameters accepted by the tool.
type AccountAIRunCfMetaExecuteLlama3_1_8bInstructFastParamsBodyMessagesToolsObjectParameters struct {
	// Definitions of each parameter.
	Properties param.Field[map[string]AccountAIRunCfMetaExecuteLlama3_1_8bInstructFastParamsBodyMessagesToolsObjectParametersProperties] `json:"properties,required"`
	// The type of the parameters object (usually 'object').
	Type param.Field[string] `json:"type,required"`
	// List of required parameter names.
	Required param.Field[[]string] `json:"required"`
}

func (r AccountAIRunCfMetaExecuteLlama3_1_8bInstructFastParamsBodyMessagesToolsObjectParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfMetaExecuteLlama3_1_8bInstructFastParamsBodyMessagesToolsObjectParametersProperties struct {
	// A description of the expected parameter.
	Description param.Field[string] `json:"description,required"`
	// The data type of the parameter.
	Type param.Field[string] `json:"type,required"`
}

func (r AccountAIRunCfMetaExecuteLlama3_1_8bInstructFastParamsBodyMessagesToolsObjectParametersProperties) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfMetaExecuteLlama3_1_8bInstructFp8Params struct {
	QueueRequest param.Field[string]                                            `query:"queueRequest"`
	Body         AccountAIRunCfMetaExecuteLlama3_1_8bInstructFp8ParamsBodyUnion `json:"body"`
}

func (r AccountAIRunCfMetaExecuteLlama3_1_8bInstructFp8Params) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

// URLQuery serializes [AccountAIRunCfMetaExecuteLlama3_1_8bInstructFp8Params]'s
// query parameters as `url.Values`.
func (r AccountAIRunCfMetaExecuteLlama3_1_8bInstructFp8Params) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountAIRunCfMetaExecuteLlama3_1_8bInstructFp8ParamsBody struct {
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

func (r AccountAIRunCfMetaExecuteLlama3_1_8bInstructFp8ParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfMetaExecuteLlama3_1_8bInstructFp8ParamsBody) implementsAccountAIRunCfMetaExecuteLlama3_1_8bInstructFp8ParamsBodyUnion() {
}

// Satisfied by [AccountAIRunCfMetaExecuteLlama3_1_8bInstructFp8ParamsBodyPrompt],
// [AccountAIRunCfMetaExecuteLlama3_1_8bInstructFp8ParamsBodyMessages],
// [AccountAIRunCfMetaExecuteLlama3_1_8bInstructFp8ParamsBody].
type AccountAIRunCfMetaExecuteLlama3_1_8bInstructFp8ParamsBodyUnion interface {
	implementsAccountAIRunCfMetaExecuteLlama3_1_8bInstructFp8ParamsBodyUnion()
}

type AccountAIRunCfMetaExecuteLlama3_1_8bInstructFp8ParamsBodyPrompt struct {
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
	ResponseFormat    param.Field[AccountAIRunCfMetaExecuteLlama3_1_8bInstructFp8ParamsBodyPromptResponseFormat] `json:"response_format"`
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

func (r AccountAIRunCfMetaExecuteLlama3_1_8bInstructFp8ParamsBodyPrompt) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfMetaExecuteLlama3_1_8bInstructFp8ParamsBodyPrompt) implementsAccountAIRunCfMetaExecuteLlama3_1_8bInstructFp8ParamsBodyUnion() {
}

type AccountAIRunCfMetaExecuteLlama3_1_8bInstructFp8ParamsBodyPromptResponseFormat struct {
	JsonSchema param.Field[interface{}]                                                                       `json:"json_schema"`
	Type       param.Field[AccountAIRunCfMetaExecuteLlama3_1_8bInstructFp8ParamsBodyPromptResponseFormatType] `json:"type"`
}

func (r AccountAIRunCfMetaExecuteLlama3_1_8bInstructFp8ParamsBodyPromptResponseFormat) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfMetaExecuteLlama3_1_8bInstructFp8ParamsBodyPromptResponseFormatType string

const (
	AccountAIRunCfMetaExecuteLlama3_1_8bInstructFp8ParamsBodyPromptResponseFormatTypeJsonObject AccountAIRunCfMetaExecuteLlama3_1_8bInstructFp8ParamsBodyPromptResponseFormatType = "json_object"
	AccountAIRunCfMetaExecuteLlama3_1_8bInstructFp8ParamsBodyPromptResponseFormatTypeJsonSchema AccountAIRunCfMetaExecuteLlama3_1_8bInstructFp8ParamsBodyPromptResponseFormatType = "json_schema"
)

func (r AccountAIRunCfMetaExecuteLlama3_1_8bInstructFp8ParamsBodyPromptResponseFormatType) IsKnown() bool {
	switch r {
	case AccountAIRunCfMetaExecuteLlama3_1_8bInstructFp8ParamsBodyPromptResponseFormatTypeJsonObject, AccountAIRunCfMetaExecuteLlama3_1_8bInstructFp8ParamsBodyPromptResponseFormatTypeJsonSchema:
		return true
	}
	return false
}

type AccountAIRunCfMetaExecuteLlama3_1_8bInstructFp8ParamsBodyMessages struct {
	// An array of message objects representing the conversation history.
	Messages param.Field[[]AccountAIRunCfMetaExecuteLlama3_1_8bInstructFp8ParamsBodyMessagesMessage] `json:"messages,required"`
	// Decreases the likelihood of the model repeating the same lines verbatim.
	FrequencyPenalty param.Field[float64]                                                                     `json:"frequency_penalty"`
	Functions        param.Field[[]AccountAIRunCfMetaExecuteLlama3_1_8bInstructFp8ParamsBodyMessagesFunction] `json:"functions"`
	// The maximum number of tokens to generate in the response.
	MaxTokens param.Field[int64] `json:"max_tokens"`
	// Increases the likelihood of the model introducing new topics.
	PresencePenalty param.Field[float64] `json:"presence_penalty"`
	// If true, a chat template is not applied and you must adhere to the specific
	// model's expected formatting.
	Raw param.Field[bool] `json:"raw"`
	// Penalty for repeated tokens; higher values discourage repetition.
	RepetitionPenalty param.Field[float64]                                                                         `json:"repetition_penalty"`
	ResponseFormat    param.Field[AccountAIRunCfMetaExecuteLlama3_1_8bInstructFp8ParamsBodyMessagesResponseFormat] `json:"response_format"`
	// Random seed for reproducibility of the generation.
	Seed param.Field[int64] `json:"seed"`
	// If true, the response will be streamed back incrementally using SSE, Server Sent
	// Events.
	Stream param.Field[bool] `json:"stream"`
	// Controls the randomness of the output; higher values produce more random
	// results.
	Temperature param.Field[float64] `json:"temperature"`
	// A list of tools available for the assistant to use.
	Tools param.Field[[]AccountAIRunCfMetaExecuteLlama3_1_8bInstructFp8ParamsBodyMessagesToolUnion] `json:"tools"`
	// Limits the AI to choose from the top 'k' most probable words. Lower values make
	// responses more focused; higher values introduce more variety and potential
	// surprises.
	TopK param.Field[int64] `json:"top_k"`
	// Adjusts the creativity of the AI's responses by controlling how many possible
	// words it considers. Lower values make outputs more predictable; higher values
	// allow for more varied and creative responses.
	TopP param.Field[float64] `json:"top_p"`
}

func (r AccountAIRunCfMetaExecuteLlama3_1_8bInstructFp8ParamsBodyMessages) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfMetaExecuteLlama3_1_8bInstructFp8ParamsBodyMessages) implementsAccountAIRunCfMetaExecuteLlama3_1_8bInstructFp8ParamsBodyUnion() {
}

type AccountAIRunCfMetaExecuteLlama3_1_8bInstructFp8ParamsBodyMessagesMessage struct {
	// The content of the message as a string.
	Content param.Field[string] `json:"content,required"`
	// The role of the message sender (e.g., 'user', 'assistant', 'system', 'tool').
	Role param.Field[string] `json:"role,required"`
}

func (r AccountAIRunCfMetaExecuteLlama3_1_8bInstructFp8ParamsBodyMessagesMessage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfMetaExecuteLlama3_1_8bInstructFp8ParamsBodyMessagesFunction struct {
	Code param.Field[string] `json:"code,required"`
	Name param.Field[string] `json:"name,required"`
}

func (r AccountAIRunCfMetaExecuteLlama3_1_8bInstructFp8ParamsBodyMessagesFunction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfMetaExecuteLlama3_1_8bInstructFp8ParamsBodyMessagesResponseFormat struct {
	JsonSchema param.Field[interface{}]                                                                         `json:"json_schema"`
	Type       param.Field[AccountAIRunCfMetaExecuteLlama3_1_8bInstructFp8ParamsBodyMessagesResponseFormatType] `json:"type"`
}

func (r AccountAIRunCfMetaExecuteLlama3_1_8bInstructFp8ParamsBodyMessagesResponseFormat) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfMetaExecuteLlama3_1_8bInstructFp8ParamsBodyMessagesResponseFormatType string

const (
	AccountAIRunCfMetaExecuteLlama3_1_8bInstructFp8ParamsBodyMessagesResponseFormatTypeJsonObject AccountAIRunCfMetaExecuteLlama3_1_8bInstructFp8ParamsBodyMessagesResponseFormatType = "json_object"
	AccountAIRunCfMetaExecuteLlama3_1_8bInstructFp8ParamsBodyMessagesResponseFormatTypeJsonSchema AccountAIRunCfMetaExecuteLlama3_1_8bInstructFp8ParamsBodyMessagesResponseFormatType = "json_schema"
)

func (r AccountAIRunCfMetaExecuteLlama3_1_8bInstructFp8ParamsBodyMessagesResponseFormatType) IsKnown() bool {
	switch r {
	case AccountAIRunCfMetaExecuteLlama3_1_8bInstructFp8ParamsBodyMessagesResponseFormatTypeJsonObject, AccountAIRunCfMetaExecuteLlama3_1_8bInstructFp8ParamsBodyMessagesResponseFormatTypeJsonSchema:
		return true
	}
	return false
}

type AccountAIRunCfMetaExecuteLlama3_1_8bInstructFp8ParamsBodyMessagesTool struct {
	// A brief description of what the tool does.
	Description param.Field[string]      `json:"description"`
	Function    param.Field[interface{}] `json:"function"`
	// The name of the tool. More descriptive the better.
	Name       param.Field[string]      `json:"name"`
	Parameters param.Field[interface{}] `json:"parameters"`
	// Specifies the type of tool (e.g., 'function').
	Type param.Field[string] `json:"type"`
}

func (r AccountAIRunCfMetaExecuteLlama3_1_8bInstructFp8ParamsBodyMessagesTool) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfMetaExecuteLlama3_1_8bInstructFp8ParamsBodyMessagesTool) implementsAccountAIRunCfMetaExecuteLlama3_1_8bInstructFp8ParamsBodyMessagesToolUnion() {
}

// Satisfied by
// [AccountAIRunCfMetaExecuteLlama3_1_8bInstructFp8ParamsBodyMessagesToolsObject],
// [AccountAIRunCfMetaExecuteLlama3_1_8bInstructFp8ParamsBodyMessagesToolsObject],
// [AccountAIRunCfMetaExecuteLlama3_1_8bInstructFp8ParamsBodyMessagesTool].
type AccountAIRunCfMetaExecuteLlama3_1_8bInstructFp8ParamsBodyMessagesToolUnion interface {
	implementsAccountAIRunCfMetaExecuteLlama3_1_8bInstructFp8ParamsBodyMessagesToolUnion()
}

type AccountAIRunCfMetaExecuteLlama3_1_8bInstructFp8ParamsBodyMessagesToolsObject struct {
	// A brief description of what the tool does.
	Description param.Field[string] `json:"description,required"`
	// The name of the tool. More descriptive the better.
	Name param.Field[string] `json:"name,required"`
	// Schema defining the parameters accepted by the tool.
	Parameters param.Field[AccountAIRunCfMetaExecuteLlama3_1_8bInstructFp8ParamsBodyMessagesToolsObjectParameters] `json:"parameters,required"`
}

func (r AccountAIRunCfMetaExecuteLlama3_1_8bInstructFp8ParamsBodyMessagesToolsObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfMetaExecuteLlama3_1_8bInstructFp8ParamsBodyMessagesToolsObject) implementsAccountAIRunCfMetaExecuteLlama3_1_8bInstructFp8ParamsBodyMessagesToolUnion() {
}

// Schema defining the parameters accepted by the tool.
type AccountAIRunCfMetaExecuteLlama3_1_8bInstructFp8ParamsBodyMessagesToolsObjectParameters struct {
	// Definitions of each parameter.
	Properties param.Field[map[string]AccountAIRunCfMetaExecuteLlama3_1_8bInstructFp8ParamsBodyMessagesToolsObjectParametersProperties] `json:"properties,required"`
	// The type of the parameters object (usually 'object').
	Type param.Field[string] `json:"type,required"`
	// List of required parameter names.
	Required param.Field[[]string] `json:"required"`
}

func (r AccountAIRunCfMetaExecuteLlama3_1_8bInstructFp8ParamsBodyMessagesToolsObjectParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfMetaExecuteLlama3_1_8bInstructFp8ParamsBodyMessagesToolsObjectParametersProperties struct {
	// A description of the expected parameter.
	Description param.Field[string] `json:"description,required"`
	// The data type of the parameter.
	Type param.Field[string] `json:"type,required"`
}

func (r AccountAIRunCfMetaExecuteLlama3_1_8bInstructFp8ParamsBodyMessagesToolsObjectParametersProperties) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfMetaExecuteLlama3_1_8bPreviewParams struct {
	QueueRequest param.Field[string]                                        `query:"queueRequest"`
	Body         AccountAIRunCfMetaExecuteLlama3_1_8bPreviewParamsBodyUnion `json:"body"`
}

func (r AccountAIRunCfMetaExecuteLlama3_1_8bPreviewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

// URLQuery serializes [AccountAIRunCfMetaExecuteLlama3_1_8bPreviewParams]'s query
// parameters as `url.Values`.
func (r AccountAIRunCfMetaExecuteLlama3_1_8bPreviewParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountAIRunCfMetaExecuteLlama3_1_8bPreviewParamsBody struct {
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

func (r AccountAIRunCfMetaExecuteLlama3_1_8bPreviewParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfMetaExecuteLlama3_1_8bPreviewParamsBody) implementsAccountAIRunCfMetaExecuteLlama3_1_8bPreviewParamsBodyUnion() {
}

// Satisfied by [AccountAIRunCfMetaExecuteLlama3_1_8bPreviewParamsBodyPrompt],
// [AccountAIRunCfMetaExecuteLlama3_1_8bPreviewParamsBodyMessages],
// [AccountAIRunCfMetaExecuteLlama3_1_8bPreviewParamsBody].
type AccountAIRunCfMetaExecuteLlama3_1_8bPreviewParamsBodyUnion interface {
	implementsAccountAIRunCfMetaExecuteLlama3_1_8bPreviewParamsBodyUnion()
}

type AccountAIRunCfMetaExecuteLlama3_1_8bPreviewParamsBodyPrompt struct {
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
	ResponseFormat    param.Field[AccountAIRunCfMetaExecuteLlama3_1_8bPreviewParamsBodyPromptResponseFormat] `json:"response_format"`
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

func (r AccountAIRunCfMetaExecuteLlama3_1_8bPreviewParamsBodyPrompt) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfMetaExecuteLlama3_1_8bPreviewParamsBodyPrompt) implementsAccountAIRunCfMetaExecuteLlama3_1_8bPreviewParamsBodyUnion() {
}

type AccountAIRunCfMetaExecuteLlama3_1_8bPreviewParamsBodyPromptResponseFormat struct {
	JsonSchema param.Field[interface{}]                                                                   `json:"json_schema"`
	Type       param.Field[AccountAIRunCfMetaExecuteLlama3_1_8bPreviewParamsBodyPromptResponseFormatType] `json:"type"`
}

func (r AccountAIRunCfMetaExecuteLlama3_1_8bPreviewParamsBodyPromptResponseFormat) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfMetaExecuteLlama3_1_8bPreviewParamsBodyPromptResponseFormatType string

const (
	AccountAIRunCfMetaExecuteLlama3_1_8bPreviewParamsBodyPromptResponseFormatTypeJsonObject AccountAIRunCfMetaExecuteLlama3_1_8bPreviewParamsBodyPromptResponseFormatType = "json_object"
	AccountAIRunCfMetaExecuteLlama3_1_8bPreviewParamsBodyPromptResponseFormatTypeJsonSchema AccountAIRunCfMetaExecuteLlama3_1_8bPreviewParamsBodyPromptResponseFormatType = "json_schema"
)

func (r AccountAIRunCfMetaExecuteLlama3_1_8bPreviewParamsBodyPromptResponseFormatType) IsKnown() bool {
	switch r {
	case AccountAIRunCfMetaExecuteLlama3_1_8bPreviewParamsBodyPromptResponseFormatTypeJsonObject, AccountAIRunCfMetaExecuteLlama3_1_8bPreviewParamsBodyPromptResponseFormatTypeJsonSchema:
		return true
	}
	return false
}

type AccountAIRunCfMetaExecuteLlama3_1_8bPreviewParamsBodyMessages struct {
	// An array of message objects representing the conversation history.
	Messages param.Field[[]AccountAIRunCfMetaExecuteLlama3_1_8bPreviewParamsBodyMessagesMessage] `json:"messages,required"`
	// Decreases the likelihood of the model repeating the same lines verbatim.
	FrequencyPenalty param.Field[float64]                                                                 `json:"frequency_penalty"`
	Functions        param.Field[[]AccountAIRunCfMetaExecuteLlama3_1_8bPreviewParamsBodyMessagesFunction] `json:"functions"`
	// The maximum number of tokens to generate in the response.
	MaxTokens param.Field[int64] `json:"max_tokens"`
	// Increases the likelihood of the model introducing new topics.
	PresencePenalty param.Field[float64] `json:"presence_penalty"`
	// If true, a chat template is not applied and you must adhere to the specific
	// model's expected formatting.
	Raw param.Field[bool] `json:"raw"`
	// Penalty for repeated tokens; higher values discourage repetition.
	RepetitionPenalty param.Field[float64]                                                                     `json:"repetition_penalty"`
	ResponseFormat    param.Field[AccountAIRunCfMetaExecuteLlama3_1_8bPreviewParamsBodyMessagesResponseFormat] `json:"response_format"`
	// Random seed for reproducibility of the generation.
	Seed param.Field[int64] `json:"seed"`
	// If true, the response will be streamed back incrementally using SSE, Server Sent
	// Events.
	Stream param.Field[bool] `json:"stream"`
	// Controls the randomness of the output; higher values produce more random
	// results.
	Temperature param.Field[float64] `json:"temperature"`
	// A list of tools available for the assistant to use.
	Tools param.Field[[]AccountAIRunCfMetaExecuteLlama3_1_8bPreviewParamsBodyMessagesToolUnion] `json:"tools"`
	// Limits the AI to choose from the top 'k' most probable words. Lower values make
	// responses more focused; higher values introduce more variety and potential
	// surprises.
	TopK param.Field[int64] `json:"top_k"`
	// Adjusts the creativity of the AI's responses by controlling how many possible
	// words it considers. Lower values make outputs more predictable; higher values
	// allow for more varied and creative responses.
	TopP param.Field[float64] `json:"top_p"`
}

func (r AccountAIRunCfMetaExecuteLlama3_1_8bPreviewParamsBodyMessages) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfMetaExecuteLlama3_1_8bPreviewParamsBodyMessages) implementsAccountAIRunCfMetaExecuteLlama3_1_8bPreviewParamsBodyUnion() {
}

type AccountAIRunCfMetaExecuteLlama3_1_8bPreviewParamsBodyMessagesMessage struct {
	// The content of the message as a string.
	Content param.Field[string] `json:"content,required"`
	// The role of the message sender (e.g., 'user', 'assistant', 'system', 'tool').
	Role param.Field[string] `json:"role,required"`
}

func (r AccountAIRunCfMetaExecuteLlama3_1_8bPreviewParamsBodyMessagesMessage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfMetaExecuteLlama3_1_8bPreviewParamsBodyMessagesFunction struct {
	Code param.Field[string] `json:"code,required"`
	Name param.Field[string] `json:"name,required"`
}

func (r AccountAIRunCfMetaExecuteLlama3_1_8bPreviewParamsBodyMessagesFunction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfMetaExecuteLlama3_1_8bPreviewParamsBodyMessagesResponseFormat struct {
	JsonSchema param.Field[interface{}]                                                                     `json:"json_schema"`
	Type       param.Field[AccountAIRunCfMetaExecuteLlama3_1_8bPreviewParamsBodyMessagesResponseFormatType] `json:"type"`
}

func (r AccountAIRunCfMetaExecuteLlama3_1_8bPreviewParamsBodyMessagesResponseFormat) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfMetaExecuteLlama3_1_8bPreviewParamsBodyMessagesResponseFormatType string

const (
	AccountAIRunCfMetaExecuteLlama3_1_8bPreviewParamsBodyMessagesResponseFormatTypeJsonObject AccountAIRunCfMetaExecuteLlama3_1_8bPreviewParamsBodyMessagesResponseFormatType = "json_object"
	AccountAIRunCfMetaExecuteLlama3_1_8bPreviewParamsBodyMessagesResponseFormatTypeJsonSchema AccountAIRunCfMetaExecuteLlama3_1_8bPreviewParamsBodyMessagesResponseFormatType = "json_schema"
)

func (r AccountAIRunCfMetaExecuteLlama3_1_8bPreviewParamsBodyMessagesResponseFormatType) IsKnown() bool {
	switch r {
	case AccountAIRunCfMetaExecuteLlama3_1_8bPreviewParamsBodyMessagesResponseFormatTypeJsonObject, AccountAIRunCfMetaExecuteLlama3_1_8bPreviewParamsBodyMessagesResponseFormatTypeJsonSchema:
		return true
	}
	return false
}

type AccountAIRunCfMetaExecuteLlama3_1_8bPreviewParamsBodyMessagesTool struct {
	// A brief description of what the tool does.
	Description param.Field[string]      `json:"description"`
	Function    param.Field[interface{}] `json:"function"`
	// The name of the tool. More descriptive the better.
	Name       param.Field[string]      `json:"name"`
	Parameters param.Field[interface{}] `json:"parameters"`
	// Specifies the type of tool (e.g., 'function').
	Type param.Field[string] `json:"type"`
}

func (r AccountAIRunCfMetaExecuteLlama3_1_8bPreviewParamsBodyMessagesTool) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfMetaExecuteLlama3_1_8bPreviewParamsBodyMessagesTool) implementsAccountAIRunCfMetaExecuteLlama3_1_8bPreviewParamsBodyMessagesToolUnion() {
}

// Satisfied by
// [AccountAIRunCfMetaExecuteLlama3_1_8bPreviewParamsBodyMessagesToolsObject],
// [AccountAIRunCfMetaExecuteLlama3_1_8bPreviewParamsBodyMessagesToolsObject],
// [AccountAIRunCfMetaExecuteLlama3_1_8bPreviewParamsBodyMessagesTool].
type AccountAIRunCfMetaExecuteLlama3_1_8bPreviewParamsBodyMessagesToolUnion interface {
	implementsAccountAIRunCfMetaExecuteLlama3_1_8bPreviewParamsBodyMessagesToolUnion()
}

type AccountAIRunCfMetaExecuteLlama3_1_8bPreviewParamsBodyMessagesToolsObject struct {
	// A brief description of what the tool does.
	Description param.Field[string] `json:"description,required"`
	// The name of the tool. More descriptive the better.
	Name param.Field[string] `json:"name,required"`
	// Schema defining the parameters accepted by the tool.
	Parameters param.Field[AccountAIRunCfMetaExecuteLlama3_1_8bPreviewParamsBodyMessagesToolsObjectParameters] `json:"parameters,required"`
}

func (r AccountAIRunCfMetaExecuteLlama3_1_8bPreviewParamsBodyMessagesToolsObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfMetaExecuteLlama3_1_8bPreviewParamsBodyMessagesToolsObject) implementsAccountAIRunCfMetaExecuteLlama3_1_8bPreviewParamsBodyMessagesToolUnion() {
}

// Schema defining the parameters accepted by the tool.
type AccountAIRunCfMetaExecuteLlama3_1_8bPreviewParamsBodyMessagesToolsObjectParameters struct {
	// Definitions of each parameter.
	Properties param.Field[map[string]AccountAIRunCfMetaExecuteLlama3_1_8bPreviewParamsBodyMessagesToolsObjectParametersProperties] `json:"properties,required"`
	// The type of the parameters object (usually 'object').
	Type param.Field[string] `json:"type,required"`
	// List of required parameter names.
	Required param.Field[[]string] `json:"required"`
}

func (r AccountAIRunCfMetaExecuteLlama3_1_8bPreviewParamsBodyMessagesToolsObjectParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfMetaExecuteLlama3_1_8bPreviewParamsBodyMessagesToolsObjectParametersProperties struct {
	// A description of the expected parameter.
	Description param.Field[string] `json:"description,required"`
	// The data type of the parameter.
	Type param.Field[string] `json:"type,required"`
}

func (r AccountAIRunCfMetaExecuteLlama3_1_8bPreviewParamsBodyMessagesToolsObjectParametersProperties) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfMetaExecuteLlama3_2_11bVisionInstructParams struct {
	QueueRequest param.Field[string]                                                `query:"queueRequest"`
	Body         AccountAIRunCfMetaExecuteLlama3_2_11bVisionInstructParamsBodyUnion `json:"body"`
}

func (r AccountAIRunCfMetaExecuteLlama3_2_11bVisionInstructParams) MarshalMultipart() (data []byte, contentType string, err error) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	err = apiform.MarshalRoot(r, writer)
	if err != nil {
		writer.Close()
		return nil, "", err
	}
	err = writer.Close()
	if err != nil {
		return nil, "", err
	}
	return buf.Bytes(), writer.FormDataContentType(), nil
}

// URLQuery serializes
// [AccountAIRunCfMetaExecuteLlama3_2_11bVisionInstructParams]'s query parameters
// as `url.Values`.
func (r AccountAIRunCfMetaExecuteLlama3_2_11bVisionInstructParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountAIRunCfMetaExecuteLlama3_2_11bVisionInstructParamsBody struct {
	// Decreases the likelihood of the model repeating the same lines verbatim.
	FrequencyPenalty param.Field[float64]     `json:"frequency_penalty"`
	Functions        param.Field[interface{}] `json:"functions"`
	Image            param.Field[interface{}] `json:"image"`
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
	RepetitionPenalty param.Field[float64] `json:"repetition_penalty"`
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

func (r AccountAIRunCfMetaExecuteLlama3_2_11bVisionInstructParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfMetaExecuteLlama3_2_11bVisionInstructParamsBody) implementsAccountAIRunCfMetaExecuteLlama3_2_11bVisionInstructParamsBodyUnion() {
}

// Satisfied by
// [AccountAIRunCfMetaExecuteLlama3_2_11bVisionInstructParamsBodyPrompt],
// [AccountAIRunCfMetaExecuteLlama3_2_11bVisionInstructParamsBodyMessages],
// [AccountAIRunCfMetaExecuteLlama3_2_11bVisionInstructParamsBody].
type AccountAIRunCfMetaExecuteLlama3_2_11bVisionInstructParamsBodyUnion interface {
	implementsAccountAIRunCfMetaExecuteLlama3_2_11bVisionInstructParamsBodyUnion()
}

type AccountAIRunCfMetaExecuteLlama3_2_11bVisionInstructParamsBodyPrompt struct {
	// The input text prompt for the model to generate a response.
	Prompt param.Field[string] `json:"prompt,required"`
	// Decreases the likelihood of the model repeating the same lines verbatim.
	FrequencyPenalty param.Field[float64] `json:"frequency_penalty"`
	// An array of integers that represent the image data constrained to 8-bit unsigned
	// integer values. Deprecated, use image as a part of messages now.
	Image param.Field[AccountAIRunCfMetaExecuteLlama3_2_11bVisionInstructParamsBodyPromptImageUnion] `json:"image" format:"binary"`
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
	RepetitionPenalty param.Field[float64] `json:"repetition_penalty"`
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

func (r AccountAIRunCfMetaExecuteLlama3_2_11bVisionInstructParamsBodyPrompt) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfMetaExecuteLlama3_2_11bVisionInstructParamsBodyPrompt) implementsAccountAIRunCfMetaExecuteLlama3_2_11bVisionInstructParamsBodyUnion() {
}

// An array of integers that represent the image data constrained to 8-bit unsigned
// integer values. Deprecated, use image as a part of messages now.
//
// Satisfied by
// [AccountAIRunCfMetaExecuteLlama3_2_11bVisionInstructParamsBodyPromptImageArray],
// [shared.UnionString].
type AccountAIRunCfMetaExecuteLlama3_2_11bVisionInstructParamsBodyPromptImageUnion interface {
	ImplementsAccountAIRunCfMetaExecuteLlama3_2_11bVisionInstructParamsBodyPromptImageUnion()
}

type AccountAIRunCfMetaExecuteLlama3_2_11bVisionInstructParamsBodyPromptImageArray []float64

func (r AccountAIRunCfMetaExecuteLlama3_2_11bVisionInstructParamsBodyPromptImageArray) ImplementsAccountAIRunCfMetaExecuteLlama3_2_11bVisionInstructParamsBodyPromptImageUnion() {
}

type AccountAIRunCfMetaExecuteLlama3_2_11bVisionInstructParamsBodyMessages struct {
	// An array of message objects representing the conversation history.
	Messages param.Field[[]AccountAIRunCfMetaExecuteLlama3_2_11bVisionInstructParamsBodyMessagesMessage] `json:"messages,required"`
	// Decreases the likelihood of the model repeating the same lines verbatim.
	FrequencyPenalty param.Field[float64]                                                                         `json:"frequency_penalty"`
	Functions        param.Field[[]AccountAIRunCfMetaExecuteLlama3_2_11bVisionInstructParamsBodyMessagesFunction] `json:"functions"`
	// An array of integers that represent the image data constrained to 8-bit unsigned
	// integer values. Deprecated, use image as a part of messages now.
	Image param.Field[AccountAIRunCfMetaExecuteLlama3_2_11bVisionInstructParamsBodyMessagesImageUnion] `json:"image" format:"binary"`
	// The maximum number of tokens to generate in the response.
	MaxTokens param.Field[int64] `json:"max_tokens"`
	// Increases the likelihood of the model introducing new topics.
	PresencePenalty param.Field[float64] `json:"presence_penalty"`
	// Penalty for repeated tokens; higher values discourage repetition.
	RepetitionPenalty param.Field[float64] `json:"repetition_penalty"`
	// Random seed for reproducibility of the generation.
	Seed param.Field[int64] `json:"seed"`
	// If true, the response will be streamed back incrementally.
	Stream param.Field[bool] `json:"stream"`
	// Controls the randomness of the output; higher values produce more random
	// results.
	Temperature param.Field[float64] `json:"temperature"`
	// A list of tools available for the assistant to use.
	Tools param.Field[[]AccountAIRunCfMetaExecuteLlama3_2_11bVisionInstructParamsBodyMessagesToolUnion] `json:"tools"`
	// Limits the AI to choose from the top 'k' most probable words. Lower values make
	// responses more focused; higher values introduce more variety and potential
	// surprises.
	TopK param.Field[int64] `json:"top_k"`
	// Controls the creativity of the AI's responses by adjusting how many possible
	// words it considers. Lower values make outputs more predictable; higher values
	// allow for more varied and creative responses.
	TopP param.Field[float64] `json:"top_p"`
}

func (r AccountAIRunCfMetaExecuteLlama3_2_11bVisionInstructParamsBodyMessages) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfMetaExecuteLlama3_2_11bVisionInstructParamsBodyMessages) implementsAccountAIRunCfMetaExecuteLlama3_2_11bVisionInstructParamsBodyUnion() {
}

type AccountAIRunCfMetaExecuteLlama3_2_11bVisionInstructParamsBodyMessagesMessage struct {
	// The content of the message as a string.
	Content param.Field[AccountAIRunCfMetaExecuteLlama3_2_11bVisionInstructParamsBodyMessagesMessagesContentUnion] `json:"content"`
	// The role of the message sender (e.g., 'user', 'assistant', 'system', 'tool').
	Role param.Field[string] `json:"role"`
	// The tool call id. Must be supplied for tool calls for Mistral-3. If you don't
	// know what to put here you can fall back to 000000001
	ToolCallID param.Field[string] `json:"tool_call_id"`
}

func (r AccountAIRunCfMetaExecuteLlama3_2_11bVisionInstructParamsBodyMessagesMessage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The content of the message as a string.
//
// Satisfied by [shared.UnionString],
// [AccountAIRunCfMetaExecuteLlama3_2_11bVisionInstructParamsBodyMessagesMessagesContentArray],
// [AccountAIRunCfMetaExecuteLlama3_2_11bVisionInstructParamsBodyMessagesMessagesContentObject].
type AccountAIRunCfMetaExecuteLlama3_2_11bVisionInstructParamsBodyMessagesMessagesContentUnion interface {
	ImplementsAccountAIRunCfMetaExecuteLlama3_2_11bVisionInstructParamsBodyMessagesMessagesContentUnion()
}

type AccountAIRunCfMetaExecuteLlama3_2_11bVisionInstructParamsBodyMessagesMessagesContentArray []AccountAIRunCfMetaExecuteLlama3_2_11bVisionInstructParamsBodyMessagesMessagesContentArrayItem

func (r AccountAIRunCfMetaExecuteLlama3_2_11bVisionInstructParamsBodyMessagesMessagesContentArray) ImplementsAccountAIRunCfMetaExecuteLlama3_2_11bVisionInstructParamsBodyMessagesMessagesContentUnion() {
}

type AccountAIRunCfMetaExecuteLlama3_2_11bVisionInstructParamsBodyMessagesMessagesContentArrayItem struct {
	ImageURL param.Field[AccountAIRunCfMetaExecuteLlama3_2_11bVisionInstructParamsBodyMessagesMessagesContentArrayImageURL] `json:"image_url"`
	Text     param.Field[string]                                                                                            `json:"text"`
	// Type of the content provided
	Type param.Field[string] `json:"type"`
}

func (r AccountAIRunCfMetaExecuteLlama3_2_11bVisionInstructParamsBodyMessagesMessagesContentArrayItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfMetaExecuteLlama3_2_11bVisionInstructParamsBodyMessagesMessagesContentArrayImageURL struct {
	// image uri with data (e.g. data:image/jpeg;base64,/9j/...). HTTP URL will not be
	// accepted
	URL param.Field[string] `json:"url"`
}

func (r AccountAIRunCfMetaExecuteLlama3_2_11bVisionInstructParamsBodyMessagesMessagesContentArrayImageURL) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfMetaExecuteLlama3_2_11bVisionInstructParamsBodyMessagesMessagesContentObject struct {
	ImageURL param.Field[AccountAIRunCfMetaExecuteLlama3_2_11bVisionInstructParamsBodyMessagesMessagesContentObjectImageURL] `json:"image_url"`
	Text     param.Field[string]                                                                                             `json:"text"`
	// Type of the content provided
	Type param.Field[string] `json:"type"`
}

func (r AccountAIRunCfMetaExecuteLlama3_2_11bVisionInstructParamsBodyMessagesMessagesContentObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfMetaExecuteLlama3_2_11bVisionInstructParamsBodyMessagesMessagesContentObject) ImplementsAccountAIRunCfMetaExecuteLlama3_2_11bVisionInstructParamsBodyMessagesMessagesContentUnion() {
}

type AccountAIRunCfMetaExecuteLlama3_2_11bVisionInstructParamsBodyMessagesMessagesContentObjectImageURL struct {
	// image uri with data (e.g. data:image/jpeg;base64,/9j/...). HTTP URL will not be
	// accepted
	URL param.Field[string] `json:"url"`
}

func (r AccountAIRunCfMetaExecuteLlama3_2_11bVisionInstructParamsBodyMessagesMessagesContentObjectImageURL) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfMetaExecuteLlama3_2_11bVisionInstructParamsBodyMessagesFunction struct {
	Code param.Field[string] `json:"code,required"`
	Name param.Field[string] `json:"name,required"`
}

func (r AccountAIRunCfMetaExecuteLlama3_2_11bVisionInstructParamsBodyMessagesFunction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// An array of integers that represent the image data constrained to 8-bit unsigned
// integer values. Deprecated, use image as a part of messages now.
//
// Satisfied by
// [AccountAIRunCfMetaExecuteLlama3_2_11bVisionInstructParamsBodyMessagesImageArray],
// [shared.UnionString].
type AccountAIRunCfMetaExecuteLlama3_2_11bVisionInstructParamsBodyMessagesImageUnion interface {
	ImplementsAccountAIRunCfMetaExecuteLlama3_2_11bVisionInstructParamsBodyMessagesImageUnion()
}

type AccountAIRunCfMetaExecuteLlama3_2_11bVisionInstructParamsBodyMessagesImageArray []float64

func (r AccountAIRunCfMetaExecuteLlama3_2_11bVisionInstructParamsBodyMessagesImageArray) ImplementsAccountAIRunCfMetaExecuteLlama3_2_11bVisionInstructParamsBodyMessagesImageUnion() {
}

type AccountAIRunCfMetaExecuteLlama3_2_11bVisionInstructParamsBodyMessagesTool struct {
	// A brief description of what the tool does.
	Description param.Field[string]      `json:"description"`
	Function    param.Field[interface{}] `json:"function"`
	// The name of the tool. More descriptive the better.
	Name       param.Field[string]      `json:"name"`
	Parameters param.Field[interface{}] `json:"parameters"`
	// Specifies the type of tool (e.g., 'function').
	Type param.Field[string] `json:"type"`
}

func (r AccountAIRunCfMetaExecuteLlama3_2_11bVisionInstructParamsBodyMessagesTool) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfMetaExecuteLlama3_2_11bVisionInstructParamsBodyMessagesTool) implementsAccountAIRunCfMetaExecuteLlama3_2_11bVisionInstructParamsBodyMessagesToolUnion() {
}

// Satisfied by
// [AccountAIRunCfMetaExecuteLlama3_2_11bVisionInstructParamsBodyMessagesToolsObject],
// [AccountAIRunCfMetaExecuteLlama3_2_11bVisionInstructParamsBodyMessagesToolsObject],
// [AccountAIRunCfMetaExecuteLlama3_2_11bVisionInstructParamsBodyMessagesTool].
type AccountAIRunCfMetaExecuteLlama3_2_11bVisionInstructParamsBodyMessagesToolUnion interface {
	implementsAccountAIRunCfMetaExecuteLlama3_2_11bVisionInstructParamsBodyMessagesToolUnion()
}

type AccountAIRunCfMetaExecuteLlama3_2_11bVisionInstructParamsBodyMessagesToolsObject struct {
	// A brief description of what the tool does.
	Description param.Field[string] `json:"description,required"`
	// The name of the tool. More descriptive the better.
	Name param.Field[string] `json:"name,required"`
	// Schema defining the parameters accepted by the tool.
	Parameters param.Field[AccountAIRunCfMetaExecuteLlama3_2_11bVisionInstructParamsBodyMessagesToolsObjectParameters] `json:"parameters,required"`
}

func (r AccountAIRunCfMetaExecuteLlama3_2_11bVisionInstructParamsBodyMessagesToolsObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfMetaExecuteLlama3_2_11bVisionInstructParamsBodyMessagesToolsObject) implementsAccountAIRunCfMetaExecuteLlama3_2_11bVisionInstructParamsBodyMessagesToolUnion() {
}

// Schema defining the parameters accepted by the tool.
type AccountAIRunCfMetaExecuteLlama3_2_11bVisionInstructParamsBodyMessagesToolsObjectParameters struct {
	// Definitions of each parameter.
	Properties param.Field[map[string]AccountAIRunCfMetaExecuteLlama3_2_11bVisionInstructParamsBodyMessagesToolsObjectParametersProperties] `json:"properties,required"`
	// The type of the parameters object (usually 'object').
	Type param.Field[string] `json:"type,required"`
	// List of required parameter names.
	Required param.Field[[]string] `json:"required"`
}

func (r AccountAIRunCfMetaExecuteLlama3_2_11bVisionInstructParamsBodyMessagesToolsObjectParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfMetaExecuteLlama3_2_11bVisionInstructParamsBodyMessagesToolsObjectParametersProperties struct {
	// A description of the expected parameter.
	Description param.Field[string] `json:"description,required"`
	// The data type of the parameter.
	Type param.Field[string] `json:"type,required"`
}

func (r AccountAIRunCfMetaExecuteLlama3_2_11bVisionInstructParamsBodyMessagesToolsObjectParametersProperties) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfMetaExecuteLlama3_2_1bInstructParams struct {
	QueueRequest param.Field[string]                                         `query:"queueRequest"`
	Body         AccountAIRunCfMetaExecuteLlama3_2_1bInstructParamsBodyUnion `json:"body"`
}

func (r AccountAIRunCfMetaExecuteLlama3_2_1bInstructParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

// URLQuery serializes [AccountAIRunCfMetaExecuteLlama3_2_1bInstructParams]'s query
// parameters as `url.Values`.
func (r AccountAIRunCfMetaExecuteLlama3_2_1bInstructParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountAIRunCfMetaExecuteLlama3_2_1bInstructParamsBody struct {
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

func (r AccountAIRunCfMetaExecuteLlama3_2_1bInstructParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfMetaExecuteLlama3_2_1bInstructParamsBody) implementsAccountAIRunCfMetaExecuteLlama3_2_1bInstructParamsBodyUnion() {
}

// Satisfied by [AccountAIRunCfMetaExecuteLlama3_2_1bInstructParamsBodyPrompt],
// [AccountAIRunCfMetaExecuteLlama3_2_1bInstructParamsBodyMessages],
// [AccountAIRunCfMetaExecuteLlama3_2_1bInstructParamsBody].
type AccountAIRunCfMetaExecuteLlama3_2_1bInstructParamsBodyUnion interface {
	implementsAccountAIRunCfMetaExecuteLlama3_2_1bInstructParamsBodyUnion()
}

type AccountAIRunCfMetaExecuteLlama3_2_1bInstructParamsBodyPrompt struct {
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
	ResponseFormat    param.Field[AccountAIRunCfMetaExecuteLlama3_2_1bInstructParamsBodyPromptResponseFormat] `json:"response_format"`
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

func (r AccountAIRunCfMetaExecuteLlama3_2_1bInstructParamsBodyPrompt) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfMetaExecuteLlama3_2_1bInstructParamsBodyPrompt) implementsAccountAIRunCfMetaExecuteLlama3_2_1bInstructParamsBodyUnion() {
}

type AccountAIRunCfMetaExecuteLlama3_2_1bInstructParamsBodyPromptResponseFormat struct {
	JsonSchema param.Field[interface{}]                                                                    `json:"json_schema"`
	Type       param.Field[AccountAIRunCfMetaExecuteLlama3_2_1bInstructParamsBodyPromptResponseFormatType] `json:"type"`
}

func (r AccountAIRunCfMetaExecuteLlama3_2_1bInstructParamsBodyPromptResponseFormat) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfMetaExecuteLlama3_2_1bInstructParamsBodyPromptResponseFormatType string

const (
	AccountAIRunCfMetaExecuteLlama3_2_1bInstructParamsBodyPromptResponseFormatTypeJsonObject AccountAIRunCfMetaExecuteLlama3_2_1bInstructParamsBodyPromptResponseFormatType = "json_object"
	AccountAIRunCfMetaExecuteLlama3_2_1bInstructParamsBodyPromptResponseFormatTypeJsonSchema AccountAIRunCfMetaExecuteLlama3_2_1bInstructParamsBodyPromptResponseFormatType = "json_schema"
)

func (r AccountAIRunCfMetaExecuteLlama3_2_1bInstructParamsBodyPromptResponseFormatType) IsKnown() bool {
	switch r {
	case AccountAIRunCfMetaExecuteLlama3_2_1bInstructParamsBodyPromptResponseFormatTypeJsonObject, AccountAIRunCfMetaExecuteLlama3_2_1bInstructParamsBodyPromptResponseFormatTypeJsonSchema:
		return true
	}
	return false
}

type AccountAIRunCfMetaExecuteLlama3_2_1bInstructParamsBodyMessages struct {
	// An array of message objects representing the conversation history.
	Messages param.Field[[]AccountAIRunCfMetaExecuteLlama3_2_1bInstructParamsBodyMessagesMessage] `json:"messages,required"`
	// Decreases the likelihood of the model repeating the same lines verbatim.
	FrequencyPenalty param.Field[float64]                                                                  `json:"frequency_penalty"`
	Functions        param.Field[[]AccountAIRunCfMetaExecuteLlama3_2_1bInstructParamsBodyMessagesFunction] `json:"functions"`
	// The maximum number of tokens to generate in the response.
	MaxTokens param.Field[int64] `json:"max_tokens"`
	// Increases the likelihood of the model introducing new topics.
	PresencePenalty param.Field[float64] `json:"presence_penalty"`
	// If true, a chat template is not applied and you must adhere to the specific
	// model's expected formatting.
	Raw param.Field[bool] `json:"raw"`
	// Penalty for repeated tokens; higher values discourage repetition.
	RepetitionPenalty param.Field[float64]                                                                      `json:"repetition_penalty"`
	ResponseFormat    param.Field[AccountAIRunCfMetaExecuteLlama3_2_1bInstructParamsBodyMessagesResponseFormat] `json:"response_format"`
	// Random seed for reproducibility of the generation.
	Seed param.Field[int64] `json:"seed"`
	// If true, the response will be streamed back incrementally using SSE, Server Sent
	// Events.
	Stream param.Field[bool] `json:"stream"`
	// Controls the randomness of the output; higher values produce more random
	// results.
	Temperature param.Field[float64] `json:"temperature"`
	// A list of tools available for the assistant to use.
	Tools param.Field[[]AccountAIRunCfMetaExecuteLlama3_2_1bInstructParamsBodyMessagesToolUnion] `json:"tools"`
	// Limits the AI to choose from the top 'k' most probable words. Lower values make
	// responses more focused; higher values introduce more variety and potential
	// surprises.
	TopK param.Field[int64] `json:"top_k"`
	// Adjusts the creativity of the AI's responses by controlling how many possible
	// words it considers. Lower values make outputs more predictable; higher values
	// allow for more varied and creative responses.
	TopP param.Field[float64] `json:"top_p"`
}

func (r AccountAIRunCfMetaExecuteLlama3_2_1bInstructParamsBodyMessages) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfMetaExecuteLlama3_2_1bInstructParamsBodyMessages) implementsAccountAIRunCfMetaExecuteLlama3_2_1bInstructParamsBodyUnion() {
}

type AccountAIRunCfMetaExecuteLlama3_2_1bInstructParamsBodyMessagesMessage struct {
	// The content of the message as a string.
	Content param.Field[string] `json:"content,required"`
	// The role of the message sender (e.g., 'user', 'assistant', 'system', 'tool').
	Role param.Field[string] `json:"role,required"`
}

func (r AccountAIRunCfMetaExecuteLlama3_2_1bInstructParamsBodyMessagesMessage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfMetaExecuteLlama3_2_1bInstructParamsBodyMessagesFunction struct {
	Code param.Field[string] `json:"code,required"`
	Name param.Field[string] `json:"name,required"`
}

func (r AccountAIRunCfMetaExecuteLlama3_2_1bInstructParamsBodyMessagesFunction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfMetaExecuteLlama3_2_1bInstructParamsBodyMessagesResponseFormat struct {
	JsonSchema param.Field[interface{}]                                                                      `json:"json_schema"`
	Type       param.Field[AccountAIRunCfMetaExecuteLlama3_2_1bInstructParamsBodyMessagesResponseFormatType] `json:"type"`
}

func (r AccountAIRunCfMetaExecuteLlama3_2_1bInstructParamsBodyMessagesResponseFormat) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfMetaExecuteLlama3_2_1bInstructParamsBodyMessagesResponseFormatType string

const (
	AccountAIRunCfMetaExecuteLlama3_2_1bInstructParamsBodyMessagesResponseFormatTypeJsonObject AccountAIRunCfMetaExecuteLlama3_2_1bInstructParamsBodyMessagesResponseFormatType = "json_object"
	AccountAIRunCfMetaExecuteLlama3_2_1bInstructParamsBodyMessagesResponseFormatTypeJsonSchema AccountAIRunCfMetaExecuteLlama3_2_1bInstructParamsBodyMessagesResponseFormatType = "json_schema"
)

func (r AccountAIRunCfMetaExecuteLlama3_2_1bInstructParamsBodyMessagesResponseFormatType) IsKnown() bool {
	switch r {
	case AccountAIRunCfMetaExecuteLlama3_2_1bInstructParamsBodyMessagesResponseFormatTypeJsonObject, AccountAIRunCfMetaExecuteLlama3_2_1bInstructParamsBodyMessagesResponseFormatTypeJsonSchema:
		return true
	}
	return false
}

type AccountAIRunCfMetaExecuteLlama3_2_1bInstructParamsBodyMessagesTool struct {
	// A brief description of what the tool does.
	Description param.Field[string]      `json:"description"`
	Function    param.Field[interface{}] `json:"function"`
	// The name of the tool. More descriptive the better.
	Name       param.Field[string]      `json:"name"`
	Parameters param.Field[interface{}] `json:"parameters"`
	// Specifies the type of tool (e.g., 'function').
	Type param.Field[string] `json:"type"`
}

func (r AccountAIRunCfMetaExecuteLlama3_2_1bInstructParamsBodyMessagesTool) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfMetaExecuteLlama3_2_1bInstructParamsBodyMessagesTool) implementsAccountAIRunCfMetaExecuteLlama3_2_1bInstructParamsBodyMessagesToolUnion() {
}

// Satisfied by
// [AccountAIRunCfMetaExecuteLlama3_2_1bInstructParamsBodyMessagesToolsObject],
// [AccountAIRunCfMetaExecuteLlama3_2_1bInstructParamsBodyMessagesToolsObject],
// [AccountAIRunCfMetaExecuteLlama3_2_1bInstructParamsBodyMessagesTool].
type AccountAIRunCfMetaExecuteLlama3_2_1bInstructParamsBodyMessagesToolUnion interface {
	implementsAccountAIRunCfMetaExecuteLlama3_2_1bInstructParamsBodyMessagesToolUnion()
}

type AccountAIRunCfMetaExecuteLlama3_2_1bInstructParamsBodyMessagesToolsObject struct {
	// A brief description of what the tool does.
	Description param.Field[string] `json:"description,required"`
	// The name of the tool. More descriptive the better.
	Name param.Field[string] `json:"name,required"`
	// Schema defining the parameters accepted by the tool.
	Parameters param.Field[AccountAIRunCfMetaExecuteLlama3_2_1bInstructParamsBodyMessagesToolsObjectParameters] `json:"parameters,required"`
}

func (r AccountAIRunCfMetaExecuteLlama3_2_1bInstructParamsBodyMessagesToolsObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfMetaExecuteLlama3_2_1bInstructParamsBodyMessagesToolsObject) implementsAccountAIRunCfMetaExecuteLlama3_2_1bInstructParamsBodyMessagesToolUnion() {
}

// Schema defining the parameters accepted by the tool.
type AccountAIRunCfMetaExecuteLlama3_2_1bInstructParamsBodyMessagesToolsObjectParameters struct {
	// Definitions of each parameter.
	Properties param.Field[map[string]AccountAIRunCfMetaExecuteLlama3_2_1bInstructParamsBodyMessagesToolsObjectParametersProperties] `json:"properties,required"`
	// The type of the parameters object (usually 'object').
	Type param.Field[string] `json:"type,required"`
	// List of required parameter names.
	Required param.Field[[]string] `json:"required"`
}

func (r AccountAIRunCfMetaExecuteLlama3_2_1bInstructParamsBodyMessagesToolsObjectParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfMetaExecuteLlama3_2_1bInstructParamsBodyMessagesToolsObjectParametersProperties struct {
	// A description of the expected parameter.
	Description param.Field[string] `json:"description,required"`
	// The data type of the parameter.
	Type param.Field[string] `json:"type,required"`
}

func (r AccountAIRunCfMetaExecuteLlama3_2_1bInstructParamsBodyMessagesToolsObjectParametersProperties) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfMetaExecuteLlama3_2_3bInstructParams struct {
	QueueRequest param.Field[string]                                         `query:"queueRequest"`
	Body         AccountAIRunCfMetaExecuteLlama3_2_3bInstructParamsBodyUnion `json:"body"`
}

func (r AccountAIRunCfMetaExecuteLlama3_2_3bInstructParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

// URLQuery serializes [AccountAIRunCfMetaExecuteLlama3_2_3bInstructParams]'s query
// parameters as `url.Values`.
func (r AccountAIRunCfMetaExecuteLlama3_2_3bInstructParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountAIRunCfMetaExecuteLlama3_2_3bInstructParamsBody struct {
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

func (r AccountAIRunCfMetaExecuteLlama3_2_3bInstructParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfMetaExecuteLlama3_2_3bInstructParamsBody) implementsAccountAIRunCfMetaExecuteLlama3_2_3bInstructParamsBodyUnion() {
}

// Satisfied by [AccountAIRunCfMetaExecuteLlama3_2_3bInstructParamsBodyPrompt],
// [AccountAIRunCfMetaExecuteLlama3_2_3bInstructParamsBodyMessages],
// [AccountAIRunCfMetaExecuteLlama3_2_3bInstructParamsBody].
type AccountAIRunCfMetaExecuteLlama3_2_3bInstructParamsBodyUnion interface {
	implementsAccountAIRunCfMetaExecuteLlama3_2_3bInstructParamsBodyUnion()
}

type AccountAIRunCfMetaExecuteLlama3_2_3bInstructParamsBodyPrompt struct {
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
	ResponseFormat    param.Field[AccountAIRunCfMetaExecuteLlama3_2_3bInstructParamsBodyPromptResponseFormat] `json:"response_format"`
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

func (r AccountAIRunCfMetaExecuteLlama3_2_3bInstructParamsBodyPrompt) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfMetaExecuteLlama3_2_3bInstructParamsBodyPrompt) implementsAccountAIRunCfMetaExecuteLlama3_2_3bInstructParamsBodyUnion() {
}

type AccountAIRunCfMetaExecuteLlama3_2_3bInstructParamsBodyPromptResponseFormat struct {
	JsonSchema param.Field[interface{}]                                                                    `json:"json_schema"`
	Type       param.Field[AccountAIRunCfMetaExecuteLlama3_2_3bInstructParamsBodyPromptResponseFormatType] `json:"type"`
}

func (r AccountAIRunCfMetaExecuteLlama3_2_3bInstructParamsBodyPromptResponseFormat) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfMetaExecuteLlama3_2_3bInstructParamsBodyPromptResponseFormatType string

const (
	AccountAIRunCfMetaExecuteLlama3_2_3bInstructParamsBodyPromptResponseFormatTypeJsonObject AccountAIRunCfMetaExecuteLlama3_2_3bInstructParamsBodyPromptResponseFormatType = "json_object"
	AccountAIRunCfMetaExecuteLlama3_2_3bInstructParamsBodyPromptResponseFormatTypeJsonSchema AccountAIRunCfMetaExecuteLlama3_2_3bInstructParamsBodyPromptResponseFormatType = "json_schema"
)

func (r AccountAIRunCfMetaExecuteLlama3_2_3bInstructParamsBodyPromptResponseFormatType) IsKnown() bool {
	switch r {
	case AccountAIRunCfMetaExecuteLlama3_2_3bInstructParamsBodyPromptResponseFormatTypeJsonObject, AccountAIRunCfMetaExecuteLlama3_2_3bInstructParamsBodyPromptResponseFormatTypeJsonSchema:
		return true
	}
	return false
}

type AccountAIRunCfMetaExecuteLlama3_2_3bInstructParamsBodyMessages struct {
	// An array of message objects representing the conversation history.
	Messages param.Field[[]AccountAIRunCfMetaExecuteLlama3_2_3bInstructParamsBodyMessagesMessage] `json:"messages,required"`
	// Decreases the likelihood of the model repeating the same lines verbatim.
	FrequencyPenalty param.Field[float64]                                                                  `json:"frequency_penalty"`
	Functions        param.Field[[]AccountAIRunCfMetaExecuteLlama3_2_3bInstructParamsBodyMessagesFunction] `json:"functions"`
	// The maximum number of tokens to generate in the response.
	MaxTokens param.Field[int64] `json:"max_tokens"`
	// Increases the likelihood of the model introducing new topics.
	PresencePenalty param.Field[float64] `json:"presence_penalty"`
	// If true, a chat template is not applied and you must adhere to the specific
	// model's expected formatting.
	Raw param.Field[bool] `json:"raw"`
	// Penalty for repeated tokens; higher values discourage repetition.
	RepetitionPenalty param.Field[float64]                                                                      `json:"repetition_penalty"`
	ResponseFormat    param.Field[AccountAIRunCfMetaExecuteLlama3_2_3bInstructParamsBodyMessagesResponseFormat] `json:"response_format"`
	// Random seed for reproducibility of the generation.
	Seed param.Field[int64] `json:"seed"`
	// If true, the response will be streamed back incrementally using SSE, Server Sent
	// Events.
	Stream param.Field[bool] `json:"stream"`
	// Controls the randomness of the output; higher values produce more random
	// results.
	Temperature param.Field[float64] `json:"temperature"`
	// A list of tools available for the assistant to use.
	Tools param.Field[[]AccountAIRunCfMetaExecuteLlama3_2_3bInstructParamsBodyMessagesToolUnion] `json:"tools"`
	// Limits the AI to choose from the top 'k' most probable words. Lower values make
	// responses more focused; higher values introduce more variety and potential
	// surprises.
	TopK param.Field[int64] `json:"top_k"`
	// Adjusts the creativity of the AI's responses by controlling how many possible
	// words it considers. Lower values make outputs more predictable; higher values
	// allow for more varied and creative responses.
	TopP param.Field[float64] `json:"top_p"`
}

func (r AccountAIRunCfMetaExecuteLlama3_2_3bInstructParamsBodyMessages) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfMetaExecuteLlama3_2_3bInstructParamsBodyMessages) implementsAccountAIRunCfMetaExecuteLlama3_2_3bInstructParamsBodyUnion() {
}

type AccountAIRunCfMetaExecuteLlama3_2_3bInstructParamsBodyMessagesMessage struct {
	// The content of the message as a string.
	Content param.Field[string] `json:"content,required"`
	// The role of the message sender (e.g., 'user', 'assistant', 'system', 'tool').
	Role param.Field[string] `json:"role,required"`
}

func (r AccountAIRunCfMetaExecuteLlama3_2_3bInstructParamsBodyMessagesMessage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfMetaExecuteLlama3_2_3bInstructParamsBodyMessagesFunction struct {
	Code param.Field[string] `json:"code,required"`
	Name param.Field[string] `json:"name,required"`
}

func (r AccountAIRunCfMetaExecuteLlama3_2_3bInstructParamsBodyMessagesFunction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfMetaExecuteLlama3_2_3bInstructParamsBodyMessagesResponseFormat struct {
	JsonSchema param.Field[interface{}]                                                                      `json:"json_schema"`
	Type       param.Field[AccountAIRunCfMetaExecuteLlama3_2_3bInstructParamsBodyMessagesResponseFormatType] `json:"type"`
}

func (r AccountAIRunCfMetaExecuteLlama3_2_3bInstructParamsBodyMessagesResponseFormat) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfMetaExecuteLlama3_2_3bInstructParamsBodyMessagesResponseFormatType string

const (
	AccountAIRunCfMetaExecuteLlama3_2_3bInstructParamsBodyMessagesResponseFormatTypeJsonObject AccountAIRunCfMetaExecuteLlama3_2_3bInstructParamsBodyMessagesResponseFormatType = "json_object"
	AccountAIRunCfMetaExecuteLlama3_2_3bInstructParamsBodyMessagesResponseFormatTypeJsonSchema AccountAIRunCfMetaExecuteLlama3_2_3bInstructParamsBodyMessagesResponseFormatType = "json_schema"
)

func (r AccountAIRunCfMetaExecuteLlama3_2_3bInstructParamsBodyMessagesResponseFormatType) IsKnown() bool {
	switch r {
	case AccountAIRunCfMetaExecuteLlama3_2_3bInstructParamsBodyMessagesResponseFormatTypeJsonObject, AccountAIRunCfMetaExecuteLlama3_2_3bInstructParamsBodyMessagesResponseFormatTypeJsonSchema:
		return true
	}
	return false
}

type AccountAIRunCfMetaExecuteLlama3_2_3bInstructParamsBodyMessagesTool struct {
	// A brief description of what the tool does.
	Description param.Field[string]      `json:"description"`
	Function    param.Field[interface{}] `json:"function"`
	// The name of the tool. More descriptive the better.
	Name       param.Field[string]      `json:"name"`
	Parameters param.Field[interface{}] `json:"parameters"`
	// Specifies the type of tool (e.g., 'function').
	Type param.Field[string] `json:"type"`
}

func (r AccountAIRunCfMetaExecuteLlama3_2_3bInstructParamsBodyMessagesTool) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfMetaExecuteLlama3_2_3bInstructParamsBodyMessagesTool) implementsAccountAIRunCfMetaExecuteLlama3_2_3bInstructParamsBodyMessagesToolUnion() {
}

// Satisfied by
// [AccountAIRunCfMetaExecuteLlama3_2_3bInstructParamsBodyMessagesToolsObject],
// [AccountAIRunCfMetaExecuteLlama3_2_3bInstructParamsBodyMessagesToolsObject],
// [AccountAIRunCfMetaExecuteLlama3_2_3bInstructParamsBodyMessagesTool].
type AccountAIRunCfMetaExecuteLlama3_2_3bInstructParamsBodyMessagesToolUnion interface {
	implementsAccountAIRunCfMetaExecuteLlama3_2_3bInstructParamsBodyMessagesToolUnion()
}

type AccountAIRunCfMetaExecuteLlama3_2_3bInstructParamsBodyMessagesToolsObject struct {
	// A brief description of what the tool does.
	Description param.Field[string] `json:"description,required"`
	// The name of the tool. More descriptive the better.
	Name param.Field[string] `json:"name,required"`
	// Schema defining the parameters accepted by the tool.
	Parameters param.Field[AccountAIRunCfMetaExecuteLlama3_2_3bInstructParamsBodyMessagesToolsObjectParameters] `json:"parameters,required"`
}

func (r AccountAIRunCfMetaExecuteLlama3_2_3bInstructParamsBodyMessagesToolsObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfMetaExecuteLlama3_2_3bInstructParamsBodyMessagesToolsObject) implementsAccountAIRunCfMetaExecuteLlama3_2_3bInstructParamsBodyMessagesToolUnion() {
}

// Schema defining the parameters accepted by the tool.
type AccountAIRunCfMetaExecuteLlama3_2_3bInstructParamsBodyMessagesToolsObjectParameters struct {
	// Definitions of each parameter.
	Properties param.Field[map[string]AccountAIRunCfMetaExecuteLlama3_2_3bInstructParamsBodyMessagesToolsObjectParametersProperties] `json:"properties,required"`
	// The type of the parameters object (usually 'object').
	Type param.Field[string] `json:"type,required"`
	// List of required parameter names.
	Required param.Field[[]string] `json:"required"`
}

func (r AccountAIRunCfMetaExecuteLlama3_2_3bInstructParamsBodyMessagesToolsObjectParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfMetaExecuteLlama3_2_3bInstructParamsBodyMessagesToolsObjectParametersProperties struct {
	// A description of the expected parameter.
	Description param.Field[string] `json:"description,required"`
	// The data type of the parameter.
	Type param.Field[string] `json:"type,required"`
}

func (r AccountAIRunCfMetaExecuteLlama3_2_3bInstructParamsBodyMessagesToolsObjectParametersProperties) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfMetaExecuteLlama3_8bInstructParams struct {
	QueueRequest param.Field[string]                                       `query:"queueRequest"`
	Body         AccountAIRunCfMetaExecuteLlama3_8bInstructParamsBodyUnion `json:"body"`
}

func (r AccountAIRunCfMetaExecuteLlama3_8bInstructParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

// URLQuery serializes [AccountAIRunCfMetaExecuteLlama3_8bInstructParams]'s query
// parameters as `url.Values`.
func (r AccountAIRunCfMetaExecuteLlama3_8bInstructParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountAIRunCfMetaExecuteLlama3_8bInstructParamsBody struct {
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

func (r AccountAIRunCfMetaExecuteLlama3_8bInstructParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfMetaExecuteLlama3_8bInstructParamsBody) implementsAccountAIRunCfMetaExecuteLlama3_8bInstructParamsBodyUnion() {
}

// Satisfied by [AccountAIRunCfMetaExecuteLlama3_8bInstructParamsBodyPrompt],
// [AccountAIRunCfMetaExecuteLlama3_8bInstructParamsBodyMessages],
// [AccountAIRunCfMetaExecuteLlama3_8bInstructParamsBody].
type AccountAIRunCfMetaExecuteLlama3_8bInstructParamsBodyUnion interface {
	implementsAccountAIRunCfMetaExecuteLlama3_8bInstructParamsBodyUnion()
}

type AccountAIRunCfMetaExecuteLlama3_8bInstructParamsBodyPrompt struct {
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
	RepetitionPenalty param.Field[float64]                                                                  `json:"repetition_penalty"`
	ResponseFormat    param.Field[AccountAIRunCfMetaExecuteLlama3_8bInstructParamsBodyPromptResponseFormat] `json:"response_format"`
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

func (r AccountAIRunCfMetaExecuteLlama3_8bInstructParamsBodyPrompt) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfMetaExecuteLlama3_8bInstructParamsBodyPrompt) implementsAccountAIRunCfMetaExecuteLlama3_8bInstructParamsBodyUnion() {
}

type AccountAIRunCfMetaExecuteLlama3_8bInstructParamsBodyPromptResponseFormat struct {
	JsonSchema param.Field[interface{}]                                                                  `json:"json_schema"`
	Type       param.Field[AccountAIRunCfMetaExecuteLlama3_8bInstructParamsBodyPromptResponseFormatType] `json:"type"`
}

func (r AccountAIRunCfMetaExecuteLlama3_8bInstructParamsBodyPromptResponseFormat) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfMetaExecuteLlama3_8bInstructParamsBodyPromptResponseFormatType string

const (
	AccountAIRunCfMetaExecuteLlama3_8bInstructParamsBodyPromptResponseFormatTypeJsonObject AccountAIRunCfMetaExecuteLlama3_8bInstructParamsBodyPromptResponseFormatType = "json_object"
	AccountAIRunCfMetaExecuteLlama3_8bInstructParamsBodyPromptResponseFormatTypeJsonSchema AccountAIRunCfMetaExecuteLlama3_8bInstructParamsBodyPromptResponseFormatType = "json_schema"
)

func (r AccountAIRunCfMetaExecuteLlama3_8bInstructParamsBodyPromptResponseFormatType) IsKnown() bool {
	switch r {
	case AccountAIRunCfMetaExecuteLlama3_8bInstructParamsBodyPromptResponseFormatTypeJsonObject, AccountAIRunCfMetaExecuteLlama3_8bInstructParamsBodyPromptResponseFormatTypeJsonSchema:
		return true
	}
	return false
}

type AccountAIRunCfMetaExecuteLlama3_8bInstructParamsBodyMessages struct {
	// An array of message objects representing the conversation history.
	Messages param.Field[[]AccountAIRunCfMetaExecuteLlama3_8bInstructParamsBodyMessagesMessage] `json:"messages,required"`
	// Decreases the likelihood of the model repeating the same lines verbatim.
	FrequencyPenalty param.Field[float64]                                                                `json:"frequency_penalty"`
	Functions        param.Field[[]AccountAIRunCfMetaExecuteLlama3_8bInstructParamsBodyMessagesFunction] `json:"functions"`
	// The maximum number of tokens to generate in the response.
	MaxTokens param.Field[int64] `json:"max_tokens"`
	// Increases the likelihood of the model introducing new topics.
	PresencePenalty param.Field[float64] `json:"presence_penalty"`
	// If true, a chat template is not applied and you must adhere to the specific
	// model's expected formatting.
	Raw param.Field[bool] `json:"raw"`
	// Penalty for repeated tokens; higher values discourage repetition.
	RepetitionPenalty param.Field[float64]                                                                    `json:"repetition_penalty"`
	ResponseFormat    param.Field[AccountAIRunCfMetaExecuteLlama3_8bInstructParamsBodyMessagesResponseFormat] `json:"response_format"`
	// Random seed for reproducibility of the generation.
	Seed param.Field[int64] `json:"seed"`
	// If true, the response will be streamed back incrementally using SSE, Server Sent
	// Events.
	Stream param.Field[bool] `json:"stream"`
	// Controls the randomness of the output; higher values produce more random
	// results.
	Temperature param.Field[float64] `json:"temperature"`
	// A list of tools available for the assistant to use.
	Tools param.Field[[]AccountAIRunCfMetaExecuteLlama3_8bInstructParamsBodyMessagesToolUnion] `json:"tools"`
	// Limits the AI to choose from the top 'k' most probable words. Lower values make
	// responses more focused; higher values introduce more variety and potential
	// surprises.
	TopK param.Field[int64] `json:"top_k"`
	// Adjusts the creativity of the AI's responses by controlling how many possible
	// words it considers. Lower values make outputs more predictable; higher values
	// allow for more varied and creative responses.
	TopP param.Field[float64] `json:"top_p"`
}

func (r AccountAIRunCfMetaExecuteLlama3_8bInstructParamsBodyMessages) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfMetaExecuteLlama3_8bInstructParamsBodyMessages) implementsAccountAIRunCfMetaExecuteLlama3_8bInstructParamsBodyUnion() {
}

type AccountAIRunCfMetaExecuteLlama3_8bInstructParamsBodyMessagesMessage struct {
	// The content of the message as a string.
	Content param.Field[string] `json:"content,required"`
	// The role of the message sender (e.g., 'user', 'assistant', 'system', 'tool').
	Role param.Field[string] `json:"role,required"`
}

func (r AccountAIRunCfMetaExecuteLlama3_8bInstructParamsBodyMessagesMessage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfMetaExecuteLlama3_8bInstructParamsBodyMessagesFunction struct {
	Code param.Field[string] `json:"code,required"`
	Name param.Field[string] `json:"name,required"`
}

func (r AccountAIRunCfMetaExecuteLlama3_8bInstructParamsBodyMessagesFunction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfMetaExecuteLlama3_8bInstructParamsBodyMessagesResponseFormat struct {
	JsonSchema param.Field[interface{}]                                                                    `json:"json_schema"`
	Type       param.Field[AccountAIRunCfMetaExecuteLlama3_8bInstructParamsBodyMessagesResponseFormatType] `json:"type"`
}

func (r AccountAIRunCfMetaExecuteLlama3_8bInstructParamsBodyMessagesResponseFormat) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfMetaExecuteLlama3_8bInstructParamsBodyMessagesResponseFormatType string

const (
	AccountAIRunCfMetaExecuteLlama3_8bInstructParamsBodyMessagesResponseFormatTypeJsonObject AccountAIRunCfMetaExecuteLlama3_8bInstructParamsBodyMessagesResponseFormatType = "json_object"
	AccountAIRunCfMetaExecuteLlama3_8bInstructParamsBodyMessagesResponseFormatTypeJsonSchema AccountAIRunCfMetaExecuteLlama3_8bInstructParamsBodyMessagesResponseFormatType = "json_schema"
)

func (r AccountAIRunCfMetaExecuteLlama3_8bInstructParamsBodyMessagesResponseFormatType) IsKnown() bool {
	switch r {
	case AccountAIRunCfMetaExecuteLlama3_8bInstructParamsBodyMessagesResponseFormatTypeJsonObject, AccountAIRunCfMetaExecuteLlama3_8bInstructParamsBodyMessagesResponseFormatTypeJsonSchema:
		return true
	}
	return false
}

type AccountAIRunCfMetaExecuteLlama3_8bInstructParamsBodyMessagesTool struct {
	// A brief description of what the tool does.
	Description param.Field[string]      `json:"description"`
	Function    param.Field[interface{}] `json:"function"`
	// The name of the tool. More descriptive the better.
	Name       param.Field[string]      `json:"name"`
	Parameters param.Field[interface{}] `json:"parameters"`
	// Specifies the type of tool (e.g., 'function').
	Type param.Field[string] `json:"type"`
}

func (r AccountAIRunCfMetaExecuteLlama3_8bInstructParamsBodyMessagesTool) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfMetaExecuteLlama3_8bInstructParamsBodyMessagesTool) implementsAccountAIRunCfMetaExecuteLlama3_8bInstructParamsBodyMessagesToolUnion() {
}

// Satisfied by
// [AccountAIRunCfMetaExecuteLlama3_8bInstructParamsBodyMessagesToolsObject],
// [AccountAIRunCfMetaExecuteLlama3_8bInstructParamsBodyMessagesToolsObject],
// [AccountAIRunCfMetaExecuteLlama3_8bInstructParamsBodyMessagesTool].
type AccountAIRunCfMetaExecuteLlama3_8bInstructParamsBodyMessagesToolUnion interface {
	implementsAccountAIRunCfMetaExecuteLlama3_8bInstructParamsBodyMessagesToolUnion()
}

type AccountAIRunCfMetaExecuteLlama3_8bInstructParamsBodyMessagesToolsObject struct {
	// A brief description of what the tool does.
	Description param.Field[string] `json:"description,required"`
	// The name of the tool. More descriptive the better.
	Name param.Field[string] `json:"name,required"`
	// Schema defining the parameters accepted by the tool.
	Parameters param.Field[AccountAIRunCfMetaExecuteLlama3_8bInstructParamsBodyMessagesToolsObjectParameters] `json:"parameters,required"`
}

func (r AccountAIRunCfMetaExecuteLlama3_8bInstructParamsBodyMessagesToolsObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfMetaExecuteLlama3_8bInstructParamsBodyMessagesToolsObject) implementsAccountAIRunCfMetaExecuteLlama3_8bInstructParamsBodyMessagesToolUnion() {
}

// Schema defining the parameters accepted by the tool.
type AccountAIRunCfMetaExecuteLlama3_8bInstructParamsBodyMessagesToolsObjectParameters struct {
	// Definitions of each parameter.
	Properties param.Field[map[string]AccountAIRunCfMetaExecuteLlama3_8bInstructParamsBodyMessagesToolsObjectParametersProperties] `json:"properties,required"`
	// The type of the parameters object (usually 'object').
	Type param.Field[string] `json:"type,required"`
	// List of required parameter names.
	Required param.Field[[]string] `json:"required"`
}

func (r AccountAIRunCfMetaExecuteLlama3_8bInstructParamsBodyMessagesToolsObjectParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfMetaExecuteLlama3_8bInstructParamsBodyMessagesToolsObjectParametersProperties struct {
	// A description of the expected parameter.
	Description param.Field[string] `json:"description,required"`
	// The data type of the parameter.
	Type param.Field[string] `json:"type,required"`
}

func (r AccountAIRunCfMetaExecuteLlama3_8bInstructParamsBodyMessagesToolsObjectParametersProperties) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfMetaExecuteLlama3_8bInstructAwqParams struct {
	QueueRequest param.Field[string]                                          `query:"queueRequest"`
	Body         AccountAIRunCfMetaExecuteLlama3_8bInstructAwqParamsBodyUnion `json:"body"`
}

func (r AccountAIRunCfMetaExecuteLlama3_8bInstructAwqParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

// URLQuery serializes [AccountAIRunCfMetaExecuteLlama3_8bInstructAwqParams]'s
// query parameters as `url.Values`.
func (r AccountAIRunCfMetaExecuteLlama3_8bInstructAwqParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountAIRunCfMetaExecuteLlama3_8bInstructAwqParamsBody struct {
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

func (r AccountAIRunCfMetaExecuteLlama3_8bInstructAwqParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfMetaExecuteLlama3_8bInstructAwqParamsBody) implementsAccountAIRunCfMetaExecuteLlama3_8bInstructAwqParamsBodyUnion() {
}

// Satisfied by [AccountAIRunCfMetaExecuteLlama3_8bInstructAwqParamsBodyPrompt],
// [AccountAIRunCfMetaExecuteLlama3_8bInstructAwqParamsBodyMessages],
// [AccountAIRunCfMetaExecuteLlama3_8bInstructAwqParamsBody].
type AccountAIRunCfMetaExecuteLlama3_8bInstructAwqParamsBodyUnion interface {
	implementsAccountAIRunCfMetaExecuteLlama3_8bInstructAwqParamsBodyUnion()
}

type AccountAIRunCfMetaExecuteLlama3_8bInstructAwqParamsBodyPrompt struct {
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
	RepetitionPenalty param.Field[float64]                                                                     `json:"repetition_penalty"`
	ResponseFormat    param.Field[AccountAIRunCfMetaExecuteLlama3_8bInstructAwqParamsBodyPromptResponseFormat] `json:"response_format"`
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

func (r AccountAIRunCfMetaExecuteLlama3_8bInstructAwqParamsBodyPrompt) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfMetaExecuteLlama3_8bInstructAwqParamsBodyPrompt) implementsAccountAIRunCfMetaExecuteLlama3_8bInstructAwqParamsBodyUnion() {
}

type AccountAIRunCfMetaExecuteLlama3_8bInstructAwqParamsBodyPromptResponseFormat struct {
	JsonSchema param.Field[interface{}]                                                                     `json:"json_schema"`
	Type       param.Field[AccountAIRunCfMetaExecuteLlama3_8bInstructAwqParamsBodyPromptResponseFormatType] `json:"type"`
}

func (r AccountAIRunCfMetaExecuteLlama3_8bInstructAwqParamsBodyPromptResponseFormat) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfMetaExecuteLlama3_8bInstructAwqParamsBodyPromptResponseFormatType string

const (
	AccountAIRunCfMetaExecuteLlama3_8bInstructAwqParamsBodyPromptResponseFormatTypeJsonObject AccountAIRunCfMetaExecuteLlama3_8bInstructAwqParamsBodyPromptResponseFormatType = "json_object"
	AccountAIRunCfMetaExecuteLlama3_8bInstructAwqParamsBodyPromptResponseFormatTypeJsonSchema AccountAIRunCfMetaExecuteLlama3_8bInstructAwqParamsBodyPromptResponseFormatType = "json_schema"
)

func (r AccountAIRunCfMetaExecuteLlama3_8bInstructAwqParamsBodyPromptResponseFormatType) IsKnown() bool {
	switch r {
	case AccountAIRunCfMetaExecuteLlama3_8bInstructAwqParamsBodyPromptResponseFormatTypeJsonObject, AccountAIRunCfMetaExecuteLlama3_8bInstructAwqParamsBodyPromptResponseFormatTypeJsonSchema:
		return true
	}
	return false
}

type AccountAIRunCfMetaExecuteLlama3_8bInstructAwqParamsBodyMessages struct {
	// An array of message objects representing the conversation history.
	Messages param.Field[[]AccountAIRunCfMetaExecuteLlama3_8bInstructAwqParamsBodyMessagesMessage] `json:"messages,required"`
	// Decreases the likelihood of the model repeating the same lines verbatim.
	FrequencyPenalty param.Field[float64]                                                                   `json:"frequency_penalty"`
	Functions        param.Field[[]AccountAIRunCfMetaExecuteLlama3_8bInstructAwqParamsBodyMessagesFunction] `json:"functions"`
	// The maximum number of tokens to generate in the response.
	MaxTokens param.Field[int64] `json:"max_tokens"`
	// Increases the likelihood of the model introducing new topics.
	PresencePenalty param.Field[float64] `json:"presence_penalty"`
	// If true, a chat template is not applied and you must adhere to the specific
	// model's expected formatting.
	Raw param.Field[bool] `json:"raw"`
	// Penalty for repeated tokens; higher values discourage repetition.
	RepetitionPenalty param.Field[float64]                                                                       `json:"repetition_penalty"`
	ResponseFormat    param.Field[AccountAIRunCfMetaExecuteLlama3_8bInstructAwqParamsBodyMessagesResponseFormat] `json:"response_format"`
	// Random seed for reproducibility of the generation.
	Seed param.Field[int64] `json:"seed"`
	// If true, the response will be streamed back incrementally using SSE, Server Sent
	// Events.
	Stream param.Field[bool] `json:"stream"`
	// Controls the randomness of the output; higher values produce more random
	// results.
	Temperature param.Field[float64] `json:"temperature"`
	// A list of tools available for the assistant to use.
	Tools param.Field[[]AccountAIRunCfMetaExecuteLlama3_8bInstructAwqParamsBodyMessagesToolUnion] `json:"tools"`
	// Limits the AI to choose from the top 'k' most probable words. Lower values make
	// responses more focused; higher values introduce more variety and potential
	// surprises.
	TopK param.Field[int64] `json:"top_k"`
	// Adjusts the creativity of the AI's responses by controlling how many possible
	// words it considers. Lower values make outputs more predictable; higher values
	// allow for more varied and creative responses.
	TopP param.Field[float64] `json:"top_p"`
}

func (r AccountAIRunCfMetaExecuteLlama3_8bInstructAwqParamsBodyMessages) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfMetaExecuteLlama3_8bInstructAwqParamsBodyMessages) implementsAccountAIRunCfMetaExecuteLlama3_8bInstructAwqParamsBodyUnion() {
}

type AccountAIRunCfMetaExecuteLlama3_8bInstructAwqParamsBodyMessagesMessage struct {
	// The content of the message as a string.
	Content param.Field[string] `json:"content,required"`
	// The role of the message sender (e.g., 'user', 'assistant', 'system', 'tool').
	Role param.Field[string] `json:"role,required"`
}

func (r AccountAIRunCfMetaExecuteLlama3_8bInstructAwqParamsBodyMessagesMessage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfMetaExecuteLlama3_8bInstructAwqParamsBodyMessagesFunction struct {
	Code param.Field[string] `json:"code,required"`
	Name param.Field[string] `json:"name,required"`
}

func (r AccountAIRunCfMetaExecuteLlama3_8bInstructAwqParamsBodyMessagesFunction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfMetaExecuteLlama3_8bInstructAwqParamsBodyMessagesResponseFormat struct {
	JsonSchema param.Field[interface{}]                                                                       `json:"json_schema"`
	Type       param.Field[AccountAIRunCfMetaExecuteLlama3_8bInstructAwqParamsBodyMessagesResponseFormatType] `json:"type"`
}

func (r AccountAIRunCfMetaExecuteLlama3_8bInstructAwqParamsBodyMessagesResponseFormat) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfMetaExecuteLlama3_8bInstructAwqParamsBodyMessagesResponseFormatType string

const (
	AccountAIRunCfMetaExecuteLlama3_8bInstructAwqParamsBodyMessagesResponseFormatTypeJsonObject AccountAIRunCfMetaExecuteLlama3_8bInstructAwqParamsBodyMessagesResponseFormatType = "json_object"
	AccountAIRunCfMetaExecuteLlama3_8bInstructAwqParamsBodyMessagesResponseFormatTypeJsonSchema AccountAIRunCfMetaExecuteLlama3_8bInstructAwqParamsBodyMessagesResponseFormatType = "json_schema"
)

func (r AccountAIRunCfMetaExecuteLlama3_8bInstructAwqParamsBodyMessagesResponseFormatType) IsKnown() bool {
	switch r {
	case AccountAIRunCfMetaExecuteLlama3_8bInstructAwqParamsBodyMessagesResponseFormatTypeJsonObject, AccountAIRunCfMetaExecuteLlama3_8bInstructAwqParamsBodyMessagesResponseFormatTypeJsonSchema:
		return true
	}
	return false
}

type AccountAIRunCfMetaExecuteLlama3_8bInstructAwqParamsBodyMessagesTool struct {
	// A brief description of what the tool does.
	Description param.Field[string]      `json:"description"`
	Function    param.Field[interface{}] `json:"function"`
	// The name of the tool. More descriptive the better.
	Name       param.Field[string]      `json:"name"`
	Parameters param.Field[interface{}] `json:"parameters"`
	// Specifies the type of tool (e.g., 'function').
	Type param.Field[string] `json:"type"`
}

func (r AccountAIRunCfMetaExecuteLlama3_8bInstructAwqParamsBodyMessagesTool) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfMetaExecuteLlama3_8bInstructAwqParamsBodyMessagesTool) implementsAccountAIRunCfMetaExecuteLlama3_8bInstructAwqParamsBodyMessagesToolUnion() {
}

// Satisfied by
// [AccountAIRunCfMetaExecuteLlama3_8bInstructAwqParamsBodyMessagesToolsObject],
// [AccountAIRunCfMetaExecuteLlama3_8bInstructAwqParamsBodyMessagesToolsObject],
// [AccountAIRunCfMetaExecuteLlama3_8bInstructAwqParamsBodyMessagesTool].
type AccountAIRunCfMetaExecuteLlama3_8bInstructAwqParamsBodyMessagesToolUnion interface {
	implementsAccountAIRunCfMetaExecuteLlama3_8bInstructAwqParamsBodyMessagesToolUnion()
}

type AccountAIRunCfMetaExecuteLlama3_8bInstructAwqParamsBodyMessagesToolsObject struct {
	// A brief description of what the tool does.
	Description param.Field[string] `json:"description,required"`
	// The name of the tool. More descriptive the better.
	Name param.Field[string] `json:"name,required"`
	// Schema defining the parameters accepted by the tool.
	Parameters param.Field[AccountAIRunCfMetaExecuteLlama3_8bInstructAwqParamsBodyMessagesToolsObjectParameters] `json:"parameters,required"`
}

func (r AccountAIRunCfMetaExecuteLlama3_8bInstructAwqParamsBodyMessagesToolsObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfMetaExecuteLlama3_8bInstructAwqParamsBodyMessagesToolsObject) implementsAccountAIRunCfMetaExecuteLlama3_8bInstructAwqParamsBodyMessagesToolUnion() {
}

// Schema defining the parameters accepted by the tool.
type AccountAIRunCfMetaExecuteLlama3_8bInstructAwqParamsBodyMessagesToolsObjectParameters struct {
	// Definitions of each parameter.
	Properties param.Field[map[string]AccountAIRunCfMetaExecuteLlama3_8bInstructAwqParamsBodyMessagesToolsObjectParametersProperties] `json:"properties,required"`
	// The type of the parameters object (usually 'object').
	Type param.Field[string] `json:"type,required"`
	// List of required parameter names.
	Required param.Field[[]string] `json:"required"`
}

func (r AccountAIRunCfMetaExecuteLlama3_8bInstructAwqParamsBodyMessagesToolsObjectParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfMetaExecuteLlama3_8bInstructAwqParamsBodyMessagesToolsObjectParametersProperties struct {
	// A description of the expected parameter.
	Description param.Field[string] `json:"description,required"`
	// The data type of the parameter.
	Type param.Field[string] `json:"type,required"`
}

func (r AccountAIRunCfMetaExecuteLlama3_8bInstructAwqParamsBodyMessagesToolsObjectParametersProperties) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfMetaExecuteLlamaGuard3_8bParams struct {
	// An array of message objects representing the conversation history.
	Messages     param.Field[[]AccountAIRunCfMetaExecuteLlamaGuard3_8bParamsMessage] `json:"messages,required"`
	QueueRequest param.Field[string]                                                 `query:"queueRequest"`
	// The maximum number of tokens to generate in the response.
	MaxTokens param.Field[int64] `json:"max_tokens"`
	// Dictate the output format of the generated response.
	ResponseFormat param.Field[AccountAIRunCfMetaExecuteLlamaGuard3_8bParamsResponseFormat] `json:"response_format"`
	// Controls the randomness of the output; higher values produce more random
	// results.
	Temperature param.Field[float64] `json:"temperature"`
}

func (r AccountAIRunCfMetaExecuteLlamaGuard3_8bParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [AccountAIRunCfMetaExecuteLlamaGuard3_8bParams]'s query
// parameters as `url.Values`.
func (r AccountAIRunCfMetaExecuteLlamaGuard3_8bParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountAIRunCfMetaExecuteLlamaGuard3_8bParamsMessage struct {
	// The content of the message as a string.
	Content param.Field[string] `json:"content,required"`
	// The role of the message sender must alternate between 'user' and 'assistant'.
	Role param.Field[AccountAIRunCfMetaExecuteLlamaGuard3_8bParamsMessagesRole] `json:"role,required"`
}

func (r AccountAIRunCfMetaExecuteLlamaGuard3_8bParamsMessage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The role of the message sender must alternate between 'user' and 'assistant'.
type AccountAIRunCfMetaExecuteLlamaGuard3_8bParamsMessagesRole string

const (
	AccountAIRunCfMetaExecuteLlamaGuard3_8bParamsMessagesRoleUser      AccountAIRunCfMetaExecuteLlamaGuard3_8bParamsMessagesRole = "user"
	AccountAIRunCfMetaExecuteLlamaGuard3_8bParamsMessagesRoleAssistant AccountAIRunCfMetaExecuteLlamaGuard3_8bParamsMessagesRole = "assistant"
)

func (r AccountAIRunCfMetaExecuteLlamaGuard3_8bParamsMessagesRole) IsKnown() bool {
	switch r {
	case AccountAIRunCfMetaExecuteLlamaGuard3_8bParamsMessagesRoleUser, AccountAIRunCfMetaExecuteLlamaGuard3_8bParamsMessagesRoleAssistant:
		return true
	}
	return false
}

// Dictate the output format of the generated response.
type AccountAIRunCfMetaExecuteLlamaGuard3_8bParamsResponseFormat struct {
	// Set to json_object to process and output generated text as JSON.
	Type param.Field[string] `json:"type"`
}

func (r AccountAIRunCfMetaExecuteLlamaGuard3_8bParamsResponseFormat) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfMetaExecuteM2m100_1_2bParams struct {
	QueueRequest param.Field[string]                                 `query:"queueRequest"`
	Body         AccountAIRunCfMetaExecuteM2m100_1_2bParamsBodyUnion `json:"body"`
}

func (r AccountAIRunCfMetaExecuteM2m100_1_2bParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

// URLQuery serializes [AccountAIRunCfMetaExecuteM2m100_1_2bParams]'s query
// parameters as `url.Values`.
func (r AccountAIRunCfMetaExecuteM2m100_1_2bParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountAIRunCfMetaExecuteM2m100_1_2bParamsBody struct {
	Requests param.Field[interface{}] `json:"requests"`
	// The language code of the source text (e.g., 'en' for English). Defaults to 'en'
	// if not specified
	SourceLang param.Field[string] `json:"source_lang"`
	// The language code to translate the text into (e.g., 'es' for Spanish)
	TargetLang param.Field[string] `json:"target_lang"`
	// The text to be translated
	Text param.Field[string] `json:"text"`
}

func (r AccountAIRunCfMetaExecuteM2m100_1_2bParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfMetaExecuteM2m100_1_2bParamsBody) implementsAccountAIRunCfMetaExecuteM2m100_1_2bParamsBodyUnion() {
}

// Satisfied by [AccountAIRunCfMetaExecuteM2m100_1_2bParamsBodyObject],
// [AccountAIRunCfMetaExecuteM2m100_1_2bParamsBodyRequests],
// [AccountAIRunCfMetaExecuteM2m100_1_2bParamsBody].
type AccountAIRunCfMetaExecuteM2m100_1_2bParamsBodyUnion interface {
	implementsAccountAIRunCfMetaExecuteM2m100_1_2bParamsBodyUnion()
}

type AccountAIRunCfMetaExecuteM2m100_1_2bParamsBodyObject struct {
	// The language code to translate the text into (e.g., 'es' for Spanish)
	TargetLang param.Field[string] `json:"target_lang,required"`
	// The text to be translated
	Text param.Field[string] `json:"text,required"`
	// The language code of the source text (e.g., 'en' for English). Defaults to 'en'
	// if not specified
	SourceLang param.Field[string] `json:"source_lang"`
}

func (r AccountAIRunCfMetaExecuteM2m100_1_2bParamsBodyObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfMetaExecuteM2m100_1_2bParamsBodyObject) implementsAccountAIRunCfMetaExecuteM2m100_1_2bParamsBodyUnion() {
}

type AccountAIRunCfMetaExecuteM2m100_1_2bParamsBodyRequests struct {
	// Batch of the embeddings requests to run using async-queue
	Requests param.Field[[]AccountAIRunCfMetaExecuteM2m100_1_2bParamsBodyRequestsRequest] `json:"requests,required"`
}

func (r AccountAIRunCfMetaExecuteM2m100_1_2bParamsBodyRequests) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfMetaExecuteM2m100_1_2bParamsBodyRequests) implementsAccountAIRunCfMetaExecuteM2m100_1_2bParamsBodyUnion() {
}

type AccountAIRunCfMetaExecuteM2m100_1_2bParamsBodyRequestsRequest struct {
	// The language code to translate the text into (e.g., 'es' for Spanish)
	TargetLang param.Field[string] `json:"target_lang,required"`
	// The text to be translated
	Text param.Field[string] `json:"text,required"`
	// The language code of the source text (e.g., 'en' for English). Defaults to 'en'
	// if not specified
	SourceLang param.Field[string] `json:"source_lang"`
}

func (r AccountAIRunCfMetaExecuteM2m100_1_2bParamsBodyRequestsRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
