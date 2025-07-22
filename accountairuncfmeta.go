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
	"reflect"

	"github.com/rexscaria/api-schemas/internal/apiform"
	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/apiquery"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
	"github.com/rexscaria/api-schemas/shared"
	"github.com/tidwall/gjson"
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

// Execute @cf/meta/llama-3.1-8b-instruct model.
func (r *AccountAIRunCfMetaService) ExecuteLlama3_1_8bInstruct(ctx context.Context, accountID string, params AccountAIRunCfMetaExecuteLlama3_1_8bInstructParams, opts ...option.RequestOption) (res *AccountAIRunCfMetaExecuteLlama3_1_8bInstructResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai/run/@cf/meta/llama-3.1-8b-instruct", accountID)
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

// Execute @cf/meta/llama-3.3-70b-instruct-fp8-fast model.
func (r *AccountAIRunCfMetaService) ExecuteLlama3_3_70bInstructFp8Fast(ctx context.Context, accountID string, params AccountAIRunCfMetaExecuteLlama3_3_70bInstructFp8FastParams, opts ...option.RequestOption) (res *AccountAIRunCfMetaExecuteLlama3_3_70bInstructFp8FastResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai/run/@cf/meta/llama-3.3-70b-instruct-fp8-fast", accountID)
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

type AccountAIRunCfMetaExecuteLlama2_7bChatFp16Response struct {
	Result  AccountAIRunCfMetaExecuteLlama2_7bChatFp16ResponseResultUnion `json:"result" format:"binary"`
	Success bool                                                          `json:"success"`
	JSON    accountAIRunCfMetaExecuteLlama2_7bChatFp16ResponseJSON        `json:"-"`
}

// accountAIRunCfMetaExecuteLlama2_7bChatFp16ResponseJSON contains the JSON
// metadata for the struct [AccountAIRunCfMetaExecuteLlama2_7bChatFp16Response]
type accountAIRunCfMetaExecuteLlama2_7bChatFp16ResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIRunCfMetaExecuteLlama2_7bChatFp16Response) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunCfMetaExecuteLlama2_7bChatFp16ResponseJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by
// [AccountAIRunCfMetaExecuteLlama2_7bChatFp16ResponseResultObject] or
// [shared.UnionString].
type AccountAIRunCfMetaExecuteLlama2_7bChatFp16ResponseResultUnion interface {
	ImplementsAccountAIRunCfMetaExecuteLlama2_7bChatFp16ResponseResultUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccountAIRunCfMetaExecuteLlama2_7bChatFp16ResponseResultUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccountAIRunCfMetaExecuteLlama2_7bChatFp16ResponseResultObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type AccountAIRunCfMetaExecuteLlama2_7bChatFp16ResponseResultObject struct {
	// The generated text response from the model
	Response string `json:"response,required"`
	// An array of tool calls requests made during the response generation
	ToolCalls []AccountAIRunCfMetaExecuteLlama2_7bChatFp16ResponseResultObjectToolCall `json:"tool_calls"`
	// Usage statistics for the inference request
	Usage AccountAIRunCfMetaExecuteLlama2_7bChatFp16ResponseResultObjectUsage `json:"usage"`
	JSON  accountAIRunCfMetaExecuteLlama2_7bChatFp16ResponseResultObjectJSON  `json:"-"`
}

// accountAIRunCfMetaExecuteLlama2_7bChatFp16ResponseResultObjectJSON contains the
// JSON metadata for the struct
// [AccountAIRunCfMetaExecuteLlama2_7bChatFp16ResponseResultObject]
type accountAIRunCfMetaExecuteLlama2_7bChatFp16ResponseResultObjectJSON struct {
	Response    apijson.Field
	ToolCalls   apijson.Field
	Usage       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIRunCfMetaExecuteLlama2_7bChatFp16ResponseResultObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunCfMetaExecuteLlama2_7bChatFp16ResponseResultObjectJSON) RawJSON() string {
	return r.raw
}

func (r AccountAIRunCfMetaExecuteLlama2_7bChatFp16ResponseResultObject) ImplementsAccountAIRunCfMetaExecuteLlama2_7bChatFp16ResponseResultUnion() {
}

type AccountAIRunCfMetaExecuteLlama2_7bChatFp16ResponseResultObjectToolCall struct {
	// The arguments passed to be passed to the tool call request
	Arguments interface{} `json:"arguments"`
	// The name of the tool to be called
	Name string                                                                     `json:"name"`
	JSON accountAIRunCfMetaExecuteLlama2_7bChatFp16ResponseResultObjectToolCallJSON `json:"-"`
}

// accountAIRunCfMetaExecuteLlama2_7bChatFp16ResponseResultObjectToolCallJSON
// contains the JSON metadata for the struct
// [AccountAIRunCfMetaExecuteLlama2_7bChatFp16ResponseResultObjectToolCall]
type accountAIRunCfMetaExecuteLlama2_7bChatFp16ResponseResultObjectToolCallJSON struct {
	Arguments   apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIRunCfMetaExecuteLlama2_7bChatFp16ResponseResultObjectToolCall) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunCfMetaExecuteLlama2_7bChatFp16ResponseResultObjectToolCallJSON) RawJSON() string {
	return r.raw
}

// Usage statistics for the inference request
type AccountAIRunCfMetaExecuteLlama2_7bChatFp16ResponseResultObjectUsage struct {
	// Total number of tokens in output
	CompletionTokens float64 `json:"completion_tokens"`
	// Total number of tokens in input
	PromptTokens float64 `json:"prompt_tokens"`
	// Total number of input and output tokens
	TotalTokens float64                                                                 `json:"total_tokens"`
	JSON        accountAIRunCfMetaExecuteLlama2_7bChatFp16ResponseResultObjectUsageJSON `json:"-"`
}

// accountAIRunCfMetaExecuteLlama2_7bChatFp16ResponseResultObjectUsageJSON contains
// the JSON metadata for the struct
// [AccountAIRunCfMetaExecuteLlama2_7bChatFp16ResponseResultObjectUsage]
type accountAIRunCfMetaExecuteLlama2_7bChatFp16ResponseResultObjectUsageJSON struct {
	CompletionTokens apijson.Field
	PromptTokens     apijson.Field
	TotalTokens      apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountAIRunCfMetaExecuteLlama2_7bChatFp16ResponseResultObjectUsage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunCfMetaExecuteLlama2_7bChatFp16ResponseResultObjectUsageJSON) RawJSON() string {
	return r.raw
}

type AccountAIRunCfMetaExecuteLlama2_7bChatInt8Response struct {
	Result  AccountAIRunCfMetaExecuteLlama2_7bChatInt8ResponseResultUnion `json:"result" format:"binary"`
	Success bool                                                          `json:"success"`
	JSON    accountAIRunCfMetaExecuteLlama2_7bChatInt8ResponseJSON        `json:"-"`
}

// accountAIRunCfMetaExecuteLlama2_7bChatInt8ResponseJSON contains the JSON
// metadata for the struct [AccountAIRunCfMetaExecuteLlama2_7bChatInt8Response]
type accountAIRunCfMetaExecuteLlama2_7bChatInt8ResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIRunCfMetaExecuteLlama2_7bChatInt8Response) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunCfMetaExecuteLlama2_7bChatInt8ResponseJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by
// [AccountAIRunCfMetaExecuteLlama2_7bChatInt8ResponseResultObject] or
// [shared.UnionString].
type AccountAIRunCfMetaExecuteLlama2_7bChatInt8ResponseResultUnion interface {
	ImplementsAccountAIRunCfMetaExecuteLlama2_7bChatInt8ResponseResultUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccountAIRunCfMetaExecuteLlama2_7bChatInt8ResponseResultUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccountAIRunCfMetaExecuteLlama2_7bChatInt8ResponseResultObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type AccountAIRunCfMetaExecuteLlama2_7bChatInt8ResponseResultObject struct {
	// The generated text response from the model
	Response string `json:"response,required"`
	// An array of tool calls requests made during the response generation
	ToolCalls []AccountAIRunCfMetaExecuteLlama2_7bChatInt8ResponseResultObjectToolCall `json:"tool_calls"`
	// Usage statistics for the inference request
	Usage AccountAIRunCfMetaExecuteLlama2_7bChatInt8ResponseResultObjectUsage `json:"usage"`
	JSON  accountAIRunCfMetaExecuteLlama2_7bChatInt8ResponseResultObjectJSON  `json:"-"`
}

// accountAIRunCfMetaExecuteLlama2_7bChatInt8ResponseResultObjectJSON contains the
// JSON metadata for the struct
// [AccountAIRunCfMetaExecuteLlama2_7bChatInt8ResponseResultObject]
type accountAIRunCfMetaExecuteLlama2_7bChatInt8ResponseResultObjectJSON struct {
	Response    apijson.Field
	ToolCalls   apijson.Field
	Usage       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIRunCfMetaExecuteLlama2_7bChatInt8ResponseResultObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunCfMetaExecuteLlama2_7bChatInt8ResponseResultObjectJSON) RawJSON() string {
	return r.raw
}

func (r AccountAIRunCfMetaExecuteLlama2_7bChatInt8ResponseResultObject) ImplementsAccountAIRunCfMetaExecuteLlama2_7bChatInt8ResponseResultUnion() {
}

type AccountAIRunCfMetaExecuteLlama2_7bChatInt8ResponseResultObjectToolCall struct {
	// The arguments passed to be passed to the tool call request
	Arguments interface{} `json:"arguments"`
	// The name of the tool to be called
	Name string                                                                     `json:"name"`
	JSON accountAIRunCfMetaExecuteLlama2_7bChatInt8ResponseResultObjectToolCallJSON `json:"-"`
}

// accountAIRunCfMetaExecuteLlama2_7bChatInt8ResponseResultObjectToolCallJSON
// contains the JSON metadata for the struct
// [AccountAIRunCfMetaExecuteLlama2_7bChatInt8ResponseResultObjectToolCall]
type accountAIRunCfMetaExecuteLlama2_7bChatInt8ResponseResultObjectToolCallJSON struct {
	Arguments   apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIRunCfMetaExecuteLlama2_7bChatInt8ResponseResultObjectToolCall) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunCfMetaExecuteLlama2_7bChatInt8ResponseResultObjectToolCallJSON) RawJSON() string {
	return r.raw
}

// Usage statistics for the inference request
type AccountAIRunCfMetaExecuteLlama2_7bChatInt8ResponseResultObjectUsage struct {
	// Total number of tokens in output
	CompletionTokens float64 `json:"completion_tokens"`
	// Total number of tokens in input
	PromptTokens float64 `json:"prompt_tokens"`
	// Total number of input and output tokens
	TotalTokens float64                                                                 `json:"total_tokens"`
	JSON        accountAIRunCfMetaExecuteLlama2_7bChatInt8ResponseResultObjectUsageJSON `json:"-"`
}

// accountAIRunCfMetaExecuteLlama2_7bChatInt8ResponseResultObjectUsageJSON contains
// the JSON metadata for the struct
// [AccountAIRunCfMetaExecuteLlama2_7bChatInt8ResponseResultObjectUsage]
type accountAIRunCfMetaExecuteLlama2_7bChatInt8ResponseResultObjectUsageJSON struct {
	CompletionTokens apijson.Field
	PromptTokens     apijson.Field
	TotalTokens      apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountAIRunCfMetaExecuteLlama2_7bChatInt8ResponseResultObjectUsage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunCfMetaExecuteLlama2_7bChatInt8ResponseResultObjectUsageJSON) RawJSON() string {
	return r.raw
}

type AccountAIRunCfMetaExecuteLlama3_1_70bInstructResponse struct {
	Result  AccountAIRunCfMetaExecuteLlama3_1_70bInstructResponseResultUnion `json:"result" format:"binary"`
	Success bool                                                             `json:"success"`
	JSON    accountAIRunCfMetaExecuteLlama3_1_70bInstructResponseJSON        `json:"-"`
}

// accountAIRunCfMetaExecuteLlama3_1_70bInstructResponseJSON contains the JSON
// metadata for the struct [AccountAIRunCfMetaExecuteLlama3_1_70bInstructResponse]
type accountAIRunCfMetaExecuteLlama3_1_70bInstructResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIRunCfMetaExecuteLlama3_1_70bInstructResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunCfMetaExecuteLlama3_1_70bInstructResponseJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by
// [AccountAIRunCfMetaExecuteLlama3_1_70bInstructResponseResultObject] or
// [shared.UnionString].
type AccountAIRunCfMetaExecuteLlama3_1_70bInstructResponseResultUnion interface {
	ImplementsAccountAIRunCfMetaExecuteLlama3_1_70bInstructResponseResultUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccountAIRunCfMetaExecuteLlama3_1_70bInstructResponseResultUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccountAIRunCfMetaExecuteLlama3_1_70bInstructResponseResultObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type AccountAIRunCfMetaExecuteLlama3_1_70bInstructResponseResultObject struct {
	// The generated text response from the model
	Response string `json:"response,required"`
	// An array of tool calls requests made during the response generation
	ToolCalls []AccountAIRunCfMetaExecuteLlama3_1_70bInstructResponseResultObjectToolCall `json:"tool_calls"`
	// Usage statistics for the inference request
	Usage AccountAIRunCfMetaExecuteLlama3_1_70bInstructResponseResultObjectUsage `json:"usage"`
	JSON  accountAIRunCfMetaExecuteLlama3_1_70bInstructResponseResultObjectJSON  `json:"-"`
}

// accountAIRunCfMetaExecuteLlama3_1_70bInstructResponseResultObjectJSON contains
// the JSON metadata for the struct
// [AccountAIRunCfMetaExecuteLlama3_1_70bInstructResponseResultObject]
type accountAIRunCfMetaExecuteLlama3_1_70bInstructResponseResultObjectJSON struct {
	Response    apijson.Field
	ToolCalls   apijson.Field
	Usage       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIRunCfMetaExecuteLlama3_1_70bInstructResponseResultObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunCfMetaExecuteLlama3_1_70bInstructResponseResultObjectJSON) RawJSON() string {
	return r.raw
}

func (r AccountAIRunCfMetaExecuteLlama3_1_70bInstructResponseResultObject) ImplementsAccountAIRunCfMetaExecuteLlama3_1_70bInstructResponseResultUnion() {
}

type AccountAIRunCfMetaExecuteLlama3_1_70bInstructResponseResultObjectToolCall struct {
	// The arguments passed to be passed to the tool call request
	Arguments interface{} `json:"arguments"`
	// The name of the tool to be called
	Name string                                                                        `json:"name"`
	JSON accountAIRunCfMetaExecuteLlama3_1_70bInstructResponseResultObjectToolCallJSON `json:"-"`
}

// accountAIRunCfMetaExecuteLlama3_1_70bInstructResponseResultObjectToolCallJSON
// contains the JSON metadata for the struct
// [AccountAIRunCfMetaExecuteLlama3_1_70bInstructResponseResultObjectToolCall]
type accountAIRunCfMetaExecuteLlama3_1_70bInstructResponseResultObjectToolCallJSON struct {
	Arguments   apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIRunCfMetaExecuteLlama3_1_70bInstructResponseResultObjectToolCall) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunCfMetaExecuteLlama3_1_70bInstructResponseResultObjectToolCallJSON) RawJSON() string {
	return r.raw
}

// Usage statistics for the inference request
type AccountAIRunCfMetaExecuteLlama3_1_70bInstructResponseResultObjectUsage struct {
	// Total number of tokens in output
	CompletionTokens float64 `json:"completion_tokens"`
	// Total number of tokens in input
	PromptTokens float64 `json:"prompt_tokens"`
	// Total number of input and output tokens
	TotalTokens float64                                                                    `json:"total_tokens"`
	JSON        accountAIRunCfMetaExecuteLlama3_1_70bInstructResponseResultObjectUsageJSON `json:"-"`
}

// accountAIRunCfMetaExecuteLlama3_1_70bInstructResponseResultObjectUsageJSON
// contains the JSON metadata for the struct
// [AccountAIRunCfMetaExecuteLlama3_1_70bInstructResponseResultObjectUsage]
type accountAIRunCfMetaExecuteLlama3_1_70bInstructResponseResultObjectUsageJSON struct {
	CompletionTokens apijson.Field
	PromptTokens     apijson.Field
	TotalTokens      apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountAIRunCfMetaExecuteLlama3_1_70bInstructResponseResultObjectUsage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunCfMetaExecuteLlama3_1_70bInstructResponseResultObjectUsageJSON) RawJSON() string {
	return r.raw
}

type AccountAIRunCfMetaExecuteLlama3_1_70bInstructPreviewResponse struct {
	Result  AccountAIRunCfMetaExecuteLlama3_1_70bInstructPreviewResponseResultUnion `json:"result" format:"binary"`
	Success bool                                                                    `json:"success"`
	JSON    accountAIRunCfMetaExecuteLlama3_1_70bInstructPreviewResponseJSON        `json:"-"`
}

// accountAIRunCfMetaExecuteLlama3_1_70bInstructPreviewResponseJSON contains the
// JSON metadata for the struct
// [AccountAIRunCfMetaExecuteLlama3_1_70bInstructPreviewResponse]
type accountAIRunCfMetaExecuteLlama3_1_70bInstructPreviewResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIRunCfMetaExecuteLlama3_1_70bInstructPreviewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunCfMetaExecuteLlama3_1_70bInstructPreviewResponseJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by
// [AccountAIRunCfMetaExecuteLlama3_1_70bInstructPreviewResponseResultObject] or
// [shared.UnionString].
type AccountAIRunCfMetaExecuteLlama3_1_70bInstructPreviewResponseResultUnion interface {
	ImplementsAccountAIRunCfMetaExecuteLlama3_1_70bInstructPreviewResponseResultUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccountAIRunCfMetaExecuteLlama3_1_70bInstructPreviewResponseResultUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccountAIRunCfMetaExecuteLlama3_1_70bInstructPreviewResponseResultObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type AccountAIRunCfMetaExecuteLlama3_1_70bInstructPreviewResponseResultObject struct {
	// The generated text response from the model
	Response string `json:"response,required"`
	// An array of tool calls requests made during the response generation
	ToolCalls []AccountAIRunCfMetaExecuteLlama3_1_70bInstructPreviewResponseResultObjectToolCall `json:"tool_calls"`
	// Usage statistics for the inference request
	Usage AccountAIRunCfMetaExecuteLlama3_1_70bInstructPreviewResponseResultObjectUsage `json:"usage"`
	JSON  accountAIRunCfMetaExecuteLlama3_1_70bInstructPreviewResponseResultObjectJSON  `json:"-"`
}

// accountAIRunCfMetaExecuteLlama3_1_70bInstructPreviewResponseResultObjectJSON
// contains the JSON metadata for the struct
// [AccountAIRunCfMetaExecuteLlama3_1_70bInstructPreviewResponseResultObject]
type accountAIRunCfMetaExecuteLlama3_1_70bInstructPreviewResponseResultObjectJSON struct {
	Response    apijson.Field
	ToolCalls   apijson.Field
	Usage       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIRunCfMetaExecuteLlama3_1_70bInstructPreviewResponseResultObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunCfMetaExecuteLlama3_1_70bInstructPreviewResponseResultObjectJSON) RawJSON() string {
	return r.raw
}

func (r AccountAIRunCfMetaExecuteLlama3_1_70bInstructPreviewResponseResultObject) ImplementsAccountAIRunCfMetaExecuteLlama3_1_70bInstructPreviewResponseResultUnion() {
}

type AccountAIRunCfMetaExecuteLlama3_1_70bInstructPreviewResponseResultObjectToolCall struct {
	// The arguments passed to be passed to the tool call request
	Arguments interface{} `json:"arguments"`
	// The name of the tool to be called
	Name string                                                                               `json:"name"`
	JSON accountAIRunCfMetaExecuteLlama3_1_70bInstructPreviewResponseResultObjectToolCallJSON `json:"-"`
}

// accountAIRunCfMetaExecuteLlama3_1_70bInstructPreviewResponseResultObjectToolCallJSON
// contains the JSON metadata for the struct
// [AccountAIRunCfMetaExecuteLlama3_1_70bInstructPreviewResponseResultObjectToolCall]
type accountAIRunCfMetaExecuteLlama3_1_70bInstructPreviewResponseResultObjectToolCallJSON struct {
	Arguments   apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIRunCfMetaExecuteLlama3_1_70bInstructPreviewResponseResultObjectToolCall) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunCfMetaExecuteLlama3_1_70bInstructPreviewResponseResultObjectToolCallJSON) RawJSON() string {
	return r.raw
}

// Usage statistics for the inference request
type AccountAIRunCfMetaExecuteLlama3_1_70bInstructPreviewResponseResultObjectUsage struct {
	// Total number of tokens in output
	CompletionTokens float64 `json:"completion_tokens"`
	// Total number of tokens in input
	PromptTokens float64 `json:"prompt_tokens"`
	// Total number of input and output tokens
	TotalTokens float64                                                                           `json:"total_tokens"`
	JSON        accountAIRunCfMetaExecuteLlama3_1_70bInstructPreviewResponseResultObjectUsageJSON `json:"-"`
}

// accountAIRunCfMetaExecuteLlama3_1_70bInstructPreviewResponseResultObjectUsageJSON
// contains the JSON metadata for the struct
// [AccountAIRunCfMetaExecuteLlama3_1_70bInstructPreviewResponseResultObjectUsage]
type accountAIRunCfMetaExecuteLlama3_1_70bInstructPreviewResponseResultObjectUsageJSON struct {
	CompletionTokens apijson.Field
	PromptTokens     apijson.Field
	TotalTokens      apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountAIRunCfMetaExecuteLlama3_1_70bInstructPreviewResponseResultObjectUsage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunCfMetaExecuteLlama3_1_70bInstructPreviewResponseResultObjectUsageJSON) RawJSON() string {
	return r.raw
}

type AccountAIRunCfMetaExecuteLlama3_1_70bPreviewResponse struct {
	Result  AccountAIRunCfMetaExecuteLlama3_1_70bPreviewResponseResultUnion `json:"result" format:"binary"`
	Success bool                                                            `json:"success"`
	JSON    accountAIRunCfMetaExecuteLlama3_1_70bPreviewResponseJSON        `json:"-"`
}

// accountAIRunCfMetaExecuteLlama3_1_70bPreviewResponseJSON contains the JSON
// metadata for the struct [AccountAIRunCfMetaExecuteLlama3_1_70bPreviewResponse]
type accountAIRunCfMetaExecuteLlama3_1_70bPreviewResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIRunCfMetaExecuteLlama3_1_70bPreviewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunCfMetaExecuteLlama3_1_70bPreviewResponseJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by
// [AccountAIRunCfMetaExecuteLlama3_1_70bPreviewResponseResultObject] or
// [shared.UnionString].
type AccountAIRunCfMetaExecuteLlama3_1_70bPreviewResponseResultUnion interface {
	ImplementsAccountAIRunCfMetaExecuteLlama3_1_70bPreviewResponseResultUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccountAIRunCfMetaExecuteLlama3_1_70bPreviewResponseResultUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccountAIRunCfMetaExecuteLlama3_1_70bPreviewResponseResultObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type AccountAIRunCfMetaExecuteLlama3_1_70bPreviewResponseResultObject struct {
	// The generated text response from the model
	Response string `json:"response,required"`
	// An array of tool calls requests made during the response generation
	ToolCalls []AccountAIRunCfMetaExecuteLlama3_1_70bPreviewResponseResultObjectToolCall `json:"tool_calls"`
	// Usage statistics for the inference request
	Usage AccountAIRunCfMetaExecuteLlama3_1_70bPreviewResponseResultObjectUsage `json:"usage"`
	JSON  accountAIRunCfMetaExecuteLlama3_1_70bPreviewResponseResultObjectJSON  `json:"-"`
}

// accountAIRunCfMetaExecuteLlama3_1_70bPreviewResponseResultObjectJSON contains
// the JSON metadata for the struct
// [AccountAIRunCfMetaExecuteLlama3_1_70bPreviewResponseResultObject]
type accountAIRunCfMetaExecuteLlama3_1_70bPreviewResponseResultObjectJSON struct {
	Response    apijson.Field
	ToolCalls   apijson.Field
	Usage       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIRunCfMetaExecuteLlama3_1_70bPreviewResponseResultObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunCfMetaExecuteLlama3_1_70bPreviewResponseResultObjectJSON) RawJSON() string {
	return r.raw
}

func (r AccountAIRunCfMetaExecuteLlama3_1_70bPreviewResponseResultObject) ImplementsAccountAIRunCfMetaExecuteLlama3_1_70bPreviewResponseResultUnion() {
}

type AccountAIRunCfMetaExecuteLlama3_1_70bPreviewResponseResultObjectToolCall struct {
	// The arguments passed to be passed to the tool call request
	Arguments interface{} `json:"arguments"`
	// The name of the tool to be called
	Name string                                                                       `json:"name"`
	JSON accountAIRunCfMetaExecuteLlama3_1_70bPreviewResponseResultObjectToolCallJSON `json:"-"`
}

// accountAIRunCfMetaExecuteLlama3_1_70bPreviewResponseResultObjectToolCallJSON
// contains the JSON metadata for the struct
// [AccountAIRunCfMetaExecuteLlama3_1_70bPreviewResponseResultObjectToolCall]
type accountAIRunCfMetaExecuteLlama3_1_70bPreviewResponseResultObjectToolCallJSON struct {
	Arguments   apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIRunCfMetaExecuteLlama3_1_70bPreviewResponseResultObjectToolCall) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunCfMetaExecuteLlama3_1_70bPreviewResponseResultObjectToolCallJSON) RawJSON() string {
	return r.raw
}

// Usage statistics for the inference request
type AccountAIRunCfMetaExecuteLlama3_1_70bPreviewResponseResultObjectUsage struct {
	// Total number of tokens in output
	CompletionTokens float64 `json:"completion_tokens"`
	// Total number of tokens in input
	PromptTokens float64 `json:"prompt_tokens"`
	// Total number of input and output tokens
	TotalTokens float64                                                                   `json:"total_tokens"`
	JSON        accountAIRunCfMetaExecuteLlama3_1_70bPreviewResponseResultObjectUsageJSON `json:"-"`
}

// accountAIRunCfMetaExecuteLlama3_1_70bPreviewResponseResultObjectUsageJSON
// contains the JSON metadata for the struct
// [AccountAIRunCfMetaExecuteLlama3_1_70bPreviewResponseResultObjectUsage]
type accountAIRunCfMetaExecuteLlama3_1_70bPreviewResponseResultObjectUsageJSON struct {
	CompletionTokens apijson.Field
	PromptTokens     apijson.Field
	TotalTokens      apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountAIRunCfMetaExecuteLlama3_1_70bPreviewResponseResultObjectUsage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunCfMetaExecuteLlama3_1_70bPreviewResponseResultObjectUsageJSON) RawJSON() string {
	return r.raw
}

type AccountAIRunCfMetaExecuteLlama3_1_8bInstructResponse struct {
	Result  AccountAIRunCfMetaExecuteLlama3_1_8bInstructResponseResultUnion `json:"result" format:"binary"`
	Success bool                                                            `json:"success"`
	JSON    accountAIRunCfMetaExecuteLlama3_1_8bInstructResponseJSON        `json:"-"`
}

// accountAIRunCfMetaExecuteLlama3_1_8bInstructResponseJSON contains the JSON
// metadata for the struct [AccountAIRunCfMetaExecuteLlama3_1_8bInstructResponse]
type accountAIRunCfMetaExecuteLlama3_1_8bInstructResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIRunCfMetaExecuteLlama3_1_8bInstructResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunCfMetaExecuteLlama3_1_8bInstructResponseJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by
// [AccountAIRunCfMetaExecuteLlama3_1_8bInstructResponseResultObject] or
// [shared.UnionString].
type AccountAIRunCfMetaExecuteLlama3_1_8bInstructResponseResultUnion interface {
	ImplementsAccountAIRunCfMetaExecuteLlama3_1_8bInstructResponseResultUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccountAIRunCfMetaExecuteLlama3_1_8bInstructResponseResultUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccountAIRunCfMetaExecuteLlama3_1_8bInstructResponseResultObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type AccountAIRunCfMetaExecuteLlama3_1_8bInstructResponseResultObject struct {
	// The generated text response from the model
	Response string `json:"response,required"`
	// An array of tool calls requests made during the response generation
	ToolCalls []AccountAIRunCfMetaExecuteLlama3_1_8bInstructResponseResultObjectToolCall `json:"tool_calls"`
	// Usage statistics for the inference request
	Usage AccountAIRunCfMetaExecuteLlama3_1_8bInstructResponseResultObjectUsage `json:"usage"`
	JSON  accountAIRunCfMetaExecuteLlama3_1_8bInstructResponseResultObjectJSON  `json:"-"`
}

// accountAIRunCfMetaExecuteLlama3_1_8bInstructResponseResultObjectJSON contains
// the JSON metadata for the struct
// [AccountAIRunCfMetaExecuteLlama3_1_8bInstructResponseResultObject]
type accountAIRunCfMetaExecuteLlama3_1_8bInstructResponseResultObjectJSON struct {
	Response    apijson.Field
	ToolCalls   apijson.Field
	Usage       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIRunCfMetaExecuteLlama3_1_8bInstructResponseResultObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunCfMetaExecuteLlama3_1_8bInstructResponseResultObjectJSON) RawJSON() string {
	return r.raw
}

func (r AccountAIRunCfMetaExecuteLlama3_1_8bInstructResponseResultObject) ImplementsAccountAIRunCfMetaExecuteLlama3_1_8bInstructResponseResultUnion() {
}

type AccountAIRunCfMetaExecuteLlama3_1_8bInstructResponseResultObjectToolCall struct {
	// The arguments passed to be passed to the tool call request
	Arguments interface{} `json:"arguments"`
	// The name of the tool to be called
	Name string                                                                       `json:"name"`
	JSON accountAIRunCfMetaExecuteLlama3_1_8bInstructResponseResultObjectToolCallJSON `json:"-"`
}

// accountAIRunCfMetaExecuteLlama3_1_8bInstructResponseResultObjectToolCallJSON
// contains the JSON metadata for the struct
// [AccountAIRunCfMetaExecuteLlama3_1_8bInstructResponseResultObjectToolCall]
type accountAIRunCfMetaExecuteLlama3_1_8bInstructResponseResultObjectToolCallJSON struct {
	Arguments   apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIRunCfMetaExecuteLlama3_1_8bInstructResponseResultObjectToolCall) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunCfMetaExecuteLlama3_1_8bInstructResponseResultObjectToolCallJSON) RawJSON() string {
	return r.raw
}

// Usage statistics for the inference request
type AccountAIRunCfMetaExecuteLlama3_1_8bInstructResponseResultObjectUsage struct {
	// Total number of tokens in output
	CompletionTokens float64 `json:"completion_tokens"`
	// Total number of tokens in input
	PromptTokens float64 `json:"prompt_tokens"`
	// Total number of input and output tokens
	TotalTokens float64                                                                   `json:"total_tokens"`
	JSON        accountAIRunCfMetaExecuteLlama3_1_8bInstructResponseResultObjectUsageJSON `json:"-"`
}

// accountAIRunCfMetaExecuteLlama3_1_8bInstructResponseResultObjectUsageJSON
// contains the JSON metadata for the struct
// [AccountAIRunCfMetaExecuteLlama3_1_8bInstructResponseResultObjectUsage]
type accountAIRunCfMetaExecuteLlama3_1_8bInstructResponseResultObjectUsageJSON struct {
	CompletionTokens apijson.Field
	PromptTokens     apijson.Field
	TotalTokens      apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountAIRunCfMetaExecuteLlama3_1_8bInstructResponseResultObjectUsage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunCfMetaExecuteLlama3_1_8bInstructResponseResultObjectUsageJSON) RawJSON() string {
	return r.raw
}

type AccountAIRunCfMetaExecuteLlama3_1_8bInstructAwqResponse struct {
	Result  AccountAIRunCfMetaExecuteLlama3_1_8bInstructAwqResponseResultUnion `json:"result" format:"binary"`
	Success bool                                                               `json:"success"`
	JSON    accountAIRunCfMetaExecuteLlama3_1_8bInstructAwqResponseJSON        `json:"-"`
}

// accountAIRunCfMetaExecuteLlama3_1_8bInstructAwqResponseJSON contains the JSON
// metadata for the struct
// [AccountAIRunCfMetaExecuteLlama3_1_8bInstructAwqResponse]
type accountAIRunCfMetaExecuteLlama3_1_8bInstructAwqResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIRunCfMetaExecuteLlama3_1_8bInstructAwqResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunCfMetaExecuteLlama3_1_8bInstructAwqResponseJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by
// [AccountAIRunCfMetaExecuteLlama3_1_8bInstructAwqResponseResultObject] or
// [shared.UnionString].
type AccountAIRunCfMetaExecuteLlama3_1_8bInstructAwqResponseResultUnion interface {
	ImplementsAccountAIRunCfMetaExecuteLlama3_1_8bInstructAwqResponseResultUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccountAIRunCfMetaExecuteLlama3_1_8bInstructAwqResponseResultUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccountAIRunCfMetaExecuteLlama3_1_8bInstructAwqResponseResultObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type AccountAIRunCfMetaExecuteLlama3_1_8bInstructAwqResponseResultObject struct {
	// The generated text response from the model
	Response string `json:"response,required"`
	// An array of tool calls requests made during the response generation
	ToolCalls []AccountAIRunCfMetaExecuteLlama3_1_8bInstructAwqResponseResultObjectToolCall `json:"tool_calls"`
	// Usage statistics for the inference request
	Usage AccountAIRunCfMetaExecuteLlama3_1_8bInstructAwqResponseResultObjectUsage `json:"usage"`
	JSON  accountAIRunCfMetaExecuteLlama3_1_8bInstructAwqResponseResultObjectJSON  `json:"-"`
}

// accountAIRunCfMetaExecuteLlama3_1_8bInstructAwqResponseResultObjectJSON contains
// the JSON metadata for the struct
// [AccountAIRunCfMetaExecuteLlama3_1_8bInstructAwqResponseResultObject]
type accountAIRunCfMetaExecuteLlama3_1_8bInstructAwqResponseResultObjectJSON struct {
	Response    apijson.Field
	ToolCalls   apijson.Field
	Usage       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIRunCfMetaExecuteLlama3_1_8bInstructAwqResponseResultObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunCfMetaExecuteLlama3_1_8bInstructAwqResponseResultObjectJSON) RawJSON() string {
	return r.raw
}

func (r AccountAIRunCfMetaExecuteLlama3_1_8bInstructAwqResponseResultObject) ImplementsAccountAIRunCfMetaExecuteLlama3_1_8bInstructAwqResponseResultUnion() {
}

type AccountAIRunCfMetaExecuteLlama3_1_8bInstructAwqResponseResultObjectToolCall struct {
	// The arguments passed to be passed to the tool call request
	Arguments interface{} `json:"arguments"`
	// The name of the tool to be called
	Name string                                                                          `json:"name"`
	JSON accountAIRunCfMetaExecuteLlama3_1_8bInstructAwqResponseResultObjectToolCallJSON `json:"-"`
}

// accountAIRunCfMetaExecuteLlama3_1_8bInstructAwqResponseResultObjectToolCallJSON
// contains the JSON metadata for the struct
// [AccountAIRunCfMetaExecuteLlama3_1_8bInstructAwqResponseResultObjectToolCall]
type accountAIRunCfMetaExecuteLlama3_1_8bInstructAwqResponseResultObjectToolCallJSON struct {
	Arguments   apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIRunCfMetaExecuteLlama3_1_8bInstructAwqResponseResultObjectToolCall) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunCfMetaExecuteLlama3_1_8bInstructAwqResponseResultObjectToolCallJSON) RawJSON() string {
	return r.raw
}

// Usage statistics for the inference request
type AccountAIRunCfMetaExecuteLlama3_1_8bInstructAwqResponseResultObjectUsage struct {
	// Total number of tokens in output
	CompletionTokens float64 `json:"completion_tokens"`
	// Total number of tokens in input
	PromptTokens float64 `json:"prompt_tokens"`
	// Total number of input and output tokens
	TotalTokens float64                                                                      `json:"total_tokens"`
	JSON        accountAIRunCfMetaExecuteLlama3_1_8bInstructAwqResponseResultObjectUsageJSON `json:"-"`
}

// accountAIRunCfMetaExecuteLlama3_1_8bInstructAwqResponseResultObjectUsageJSON
// contains the JSON metadata for the struct
// [AccountAIRunCfMetaExecuteLlama3_1_8bInstructAwqResponseResultObjectUsage]
type accountAIRunCfMetaExecuteLlama3_1_8bInstructAwqResponseResultObjectUsageJSON struct {
	CompletionTokens apijson.Field
	PromptTokens     apijson.Field
	TotalTokens      apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountAIRunCfMetaExecuteLlama3_1_8bInstructAwqResponseResultObjectUsage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunCfMetaExecuteLlama3_1_8bInstructAwqResponseResultObjectUsageJSON) RawJSON() string {
	return r.raw
}

type AccountAIRunCfMetaExecuteLlama3_1_8bInstructFastResponse struct {
	Result  AccountAIRunCfMetaExecuteLlama3_1_8bInstructFastResponseResultUnion `json:"result" format:"binary"`
	Success bool                                                                `json:"success"`
	JSON    accountAIRunCfMetaExecuteLlama3_1_8bInstructFastResponseJSON        `json:"-"`
}

// accountAIRunCfMetaExecuteLlama3_1_8bInstructFastResponseJSON contains the JSON
// metadata for the struct
// [AccountAIRunCfMetaExecuteLlama3_1_8bInstructFastResponse]
type accountAIRunCfMetaExecuteLlama3_1_8bInstructFastResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIRunCfMetaExecuteLlama3_1_8bInstructFastResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunCfMetaExecuteLlama3_1_8bInstructFastResponseJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by
// [AccountAIRunCfMetaExecuteLlama3_1_8bInstructFastResponseResultObject] or
// [shared.UnionString].
type AccountAIRunCfMetaExecuteLlama3_1_8bInstructFastResponseResultUnion interface {
	ImplementsAccountAIRunCfMetaExecuteLlama3_1_8bInstructFastResponseResultUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccountAIRunCfMetaExecuteLlama3_1_8bInstructFastResponseResultUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccountAIRunCfMetaExecuteLlama3_1_8bInstructFastResponseResultObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type AccountAIRunCfMetaExecuteLlama3_1_8bInstructFastResponseResultObject struct {
	// The generated text response from the model
	Response string `json:"response,required"`
	// An array of tool calls requests made during the response generation
	ToolCalls []AccountAIRunCfMetaExecuteLlama3_1_8bInstructFastResponseResultObjectToolCall `json:"tool_calls"`
	// Usage statistics for the inference request
	Usage AccountAIRunCfMetaExecuteLlama3_1_8bInstructFastResponseResultObjectUsage `json:"usage"`
	JSON  accountAIRunCfMetaExecuteLlama3_1_8bInstructFastResponseResultObjectJSON  `json:"-"`
}

// accountAIRunCfMetaExecuteLlama3_1_8bInstructFastResponseResultObjectJSON
// contains the JSON metadata for the struct
// [AccountAIRunCfMetaExecuteLlama3_1_8bInstructFastResponseResultObject]
type accountAIRunCfMetaExecuteLlama3_1_8bInstructFastResponseResultObjectJSON struct {
	Response    apijson.Field
	ToolCalls   apijson.Field
	Usage       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIRunCfMetaExecuteLlama3_1_8bInstructFastResponseResultObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunCfMetaExecuteLlama3_1_8bInstructFastResponseResultObjectJSON) RawJSON() string {
	return r.raw
}

func (r AccountAIRunCfMetaExecuteLlama3_1_8bInstructFastResponseResultObject) ImplementsAccountAIRunCfMetaExecuteLlama3_1_8bInstructFastResponseResultUnion() {
}

type AccountAIRunCfMetaExecuteLlama3_1_8bInstructFastResponseResultObjectToolCall struct {
	// The arguments passed to be passed to the tool call request
	Arguments interface{} `json:"arguments"`
	// The name of the tool to be called
	Name string                                                                           `json:"name"`
	JSON accountAIRunCfMetaExecuteLlama3_1_8bInstructFastResponseResultObjectToolCallJSON `json:"-"`
}

// accountAIRunCfMetaExecuteLlama3_1_8bInstructFastResponseResultObjectToolCallJSON
// contains the JSON metadata for the struct
// [AccountAIRunCfMetaExecuteLlama3_1_8bInstructFastResponseResultObjectToolCall]
type accountAIRunCfMetaExecuteLlama3_1_8bInstructFastResponseResultObjectToolCallJSON struct {
	Arguments   apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIRunCfMetaExecuteLlama3_1_8bInstructFastResponseResultObjectToolCall) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunCfMetaExecuteLlama3_1_8bInstructFastResponseResultObjectToolCallJSON) RawJSON() string {
	return r.raw
}

// Usage statistics for the inference request
type AccountAIRunCfMetaExecuteLlama3_1_8bInstructFastResponseResultObjectUsage struct {
	// Total number of tokens in output
	CompletionTokens float64 `json:"completion_tokens"`
	// Total number of tokens in input
	PromptTokens float64 `json:"prompt_tokens"`
	// Total number of input and output tokens
	TotalTokens float64                                                                       `json:"total_tokens"`
	JSON        accountAIRunCfMetaExecuteLlama3_1_8bInstructFastResponseResultObjectUsageJSON `json:"-"`
}

// accountAIRunCfMetaExecuteLlama3_1_8bInstructFastResponseResultObjectUsageJSON
// contains the JSON metadata for the struct
// [AccountAIRunCfMetaExecuteLlama3_1_8bInstructFastResponseResultObjectUsage]
type accountAIRunCfMetaExecuteLlama3_1_8bInstructFastResponseResultObjectUsageJSON struct {
	CompletionTokens apijson.Field
	PromptTokens     apijson.Field
	TotalTokens      apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountAIRunCfMetaExecuteLlama3_1_8bInstructFastResponseResultObjectUsage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunCfMetaExecuteLlama3_1_8bInstructFastResponseResultObjectUsageJSON) RawJSON() string {
	return r.raw
}

type AccountAIRunCfMetaExecuteLlama3_1_8bInstructFp8Response struct {
	Result  AccountAIRunCfMetaExecuteLlama3_1_8bInstructFp8ResponseResultUnion `json:"result" format:"binary"`
	Success bool                                                               `json:"success"`
	JSON    accountAIRunCfMetaExecuteLlama3_1_8bInstructFp8ResponseJSON        `json:"-"`
}

// accountAIRunCfMetaExecuteLlama3_1_8bInstructFp8ResponseJSON contains the JSON
// metadata for the struct
// [AccountAIRunCfMetaExecuteLlama3_1_8bInstructFp8Response]
type accountAIRunCfMetaExecuteLlama3_1_8bInstructFp8ResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIRunCfMetaExecuteLlama3_1_8bInstructFp8Response) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunCfMetaExecuteLlama3_1_8bInstructFp8ResponseJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by
// [AccountAIRunCfMetaExecuteLlama3_1_8bInstructFp8ResponseResultObject] or
// [shared.UnionString].
type AccountAIRunCfMetaExecuteLlama3_1_8bInstructFp8ResponseResultUnion interface {
	ImplementsAccountAIRunCfMetaExecuteLlama3_1_8bInstructFp8ResponseResultUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccountAIRunCfMetaExecuteLlama3_1_8bInstructFp8ResponseResultUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccountAIRunCfMetaExecuteLlama3_1_8bInstructFp8ResponseResultObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type AccountAIRunCfMetaExecuteLlama3_1_8bInstructFp8ResponseResultObject struct {
	// The generated text response from the model
	Response string `json:"response,required"`
	// An array of tool calls requests made during the response generation
	ToolCalls []AccountAIRunCfMetaExecuteLlama3_1_8bInstructFp8ResponseResultObjectToolCall `json:"tool_calls"`
	// Usage statistics for the inference request
	Usage AccountAIRunCfMetaExecuteLlama3_1_8bInstructFp8ResponseResultObjectUsage `json:"usage"`
	JSON  accountAIRunCfMetaExecuteLlama3_1_8bInstructFp8ResponseResultObjectJSON  `json:"-"`
}

// accountAIRunCfMetaExecuteLlama3_1_8bInstructFp8ResponseResultObjectJSON contains
// the JSON metadata for the struct
// [AccountAIRunCfMetaExecuteLlama3_1_8bInstructFp8ResponseResultObject]
type accountAIRunCfMetaExecuteLlama3_1_8bInstructFp8ResponseResultObjectJSON struct {
	Response    apijson.Field
	ToolCalls   apijson.Field
	Usage       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIRunCfMetaExecuteLlama3_1_8bInstructFp8ResponseResultObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunCfMetaExecuteLlama3_1_8bInstructFp8ResponseResultObjectJSON) RawJSON() string {
	return r.raw
}

func (r AccountAIRunCfMetaExecuteLlama3_1_8bInstructFp8ResponseResultObject) ImplementsAccountAIRunCfMetaExecuteLlama3_1_8bInstructFp8ResponseResultUnion() {
}

type AccountAIRunCfMetaExecuteLlama3_1_8bInstructFp8ResponseResultObjectToolCall struct {
	// The arguments passed to be passed to the tool call request
	Arguments interface{} `json:"arguments"`
	// The name of the tool to be called
	Name string                                                                          `json:"name"`
	JSON accountAIRunCfMetaExecuteLlama3_1_8bInstructFp8ResponseResultObjectToolCallJSON `json:"-"`
}

// accountAIRunCfMetaExecuteLlama3_1_8bInstructFp8ResponseResultObjectToolCallJSON
// contains the JSON metadata for the struct
// [AccountAIRunCfMetaExecuteLlama3_1_8bInstructFp8ResponseResultObjectToolCall]
type accountAIRunCfMetaExecuteLlama3_1_8bInstructFp8ResponseResultObjectToolCallJSON struct {
	Arguments   apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIRunCfMetaExecuteLlama3_1_8bInstructFp8ResponseResultObjectToolCall) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunCfMetaExecuteLlama3_1_8bInstructFp8ResponseResultObjectToolCallJSON) RawJSON() string {
	return r.raw
}

// Usage statistics for the inference request
type AccountAIRunCfMetaExecuteLlama3_1_8bInstructFp8ResponseResultObjectUsage struct {
	// Total number of tokens in output
	CompletionTokens float64 `json:"completion_tokens"`
	// Total number of tokens in input
	PromptTokens float64 `json:"prompt_tokens"`
	// Total number of input and output tokens
	TotalTokens float64                                                                      `json:"total_tokens"`
	JSON        accountAIRunCfMetaExecuteLlama3_1_8bInstructFp8ResponseResultObjectUsageJSON `json:"-"`
}

// accountAIRunCfMetaExecuteLlama3_1_8bInstructFp8ResponseResultObjectUsageJSON
// contains the JSON metadata for the struct
// [AccountAIRunCfMetaExecuteLlama3_1_8bInstructFp8ResponseResultObjectUsage]
type accountAIRunCfMetaExecuteLlama3_1_8bInstructFp8ResponseResultObjectUsageJSON struct {
	CompletionTokens apijson.Field
	PromptTokens     apijson.Field
	TotalTokens      apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountAIRunCfMetaExecuteLlama3_1_8bInstructFp8ResponseResultObjectUsage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunCfMetaExecuteLlama3_1_8bInstructFp8ResponseResultObjectUsageJSON) RawJSON() string {
	return r.raw
}

type AccountAIRunCfMetaExecuteLlama3_1_8bPreviewResponse struct {
	Result  AccountAIRunCfMetaExecuteLlama3_1_8bPreviewResponseResultUnion `json:"result" format:"binary"`
	Success bool                                                           `json:"success"`
	JSON    accountAIRunCfMetaExecuteLlama3_1_8bPreviewResponseJSON        `json:"-"`
}

// accountAIRunCfMetaExecuteLlama3_1_8bPreviewResponseJSON contains the JSON
// metadata for the struct [AccountAIRunCfMetaExecuteLlama3_1_8bPreviewResponse]
type accountAIRunCfMetaExecuteLlama3_1_8bPreviewResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIRunCfMetaExecuteLlama3_1_8bPreviewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunCfMetaExecuteLlama3_1_8bPreviewResponseJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by
// [AccountAIRunCfMetaExecuteLlama3_1_8bPreviewResponseResultObject] or
// [shared.UnionString].
type AccountAIRunCfMetaExecuteLlama3_1_8bPreviewResponseResultUnion interface {
	ImplementsAccountAIRunCfMetaExecuteLlama3_1_8bPreviewResponseResultUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccountAIRunCfMetaExecuteLlama3_1_8bPreviewResponseResultUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccountAIRunCfMetaExecuteLlama3_1_8bPreviewResponseResultObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type AccountAIRunCfMetaExecuteLlama3_1_8bPreviewResponseResultObject struct {
	// The generated text response from the model
	Response string `json:"response,required"`
	// An array of tool calls requests made during the response generation
	ToolCalls []AccountAIRunCfMetaExecuteLlama3_1_8bPreviewResponseResultObjectToolCall `json:"tool_calls"`
	// Usage statistics for the inference request
	Usage AccountAIRunCfMetaExecuteLlama3_1_8bPreviewResponseResultObjectUsage `json:"usage"`
	JSON  accountAIRunCfMetaExecuteLlama3_1_8bPreviewResponseResultObjectJSON  `json:"-"`
}

// accountAIRunCfMetaExecuteLlama3_1_8bPreviewResponseResultObjectJSON contains the
// JSON metadata for the struct
// [AccountAIRunCfMetaExecuteLlama3_1_8bPreviewResponseResultObject]
type accountAIRunCfMetaExecuteLlama3_1_8bPreviewResponseResultObjectJSON struct {
	Response    apijson.Field
	ToolCalls   apijson.Field
	Usage       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIRunCfMetaExecuteLlama3_1_8bPreviewResponseResultObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunCfMetaExecuteLlama3_1_8bPreviewResponseResultObjectJSON) RawJSON() string {
	return r.raw
}

func (r AccountAIRunCfMetaExecuteLlama3_1_8bPreviewResponseResultObject) ImplementsAccountAIRunCfMetaExecuteLlama3_1_8bPreviewResponseResultUnion() {
}

type AccountAIRunCfMetaExecuteLlama3_1_8bPreviewResponseResultObjectToolCall struct {
	// The arguments passed to be passed to the tool call request
	Arguments interface{} `json:"arguments"`
	// The name of the tool to be called
	Name string                                                                      `json:"name"`
	JSON accountAIRunCfMetaExecuteLlama3_1_8bPreviewResponseResultObjectToolCallJSON `json:"-"`
}

// accountAIRunCfMetaExecuteLlama3_1_8bPreviewResponseResultObjectToolCallJSON
// contains the JSON metadata for the struct
// [AccountAIRunCfMetaExecuteLlama3_1_8bPreviewResponseResultObjectToolCall]
type accountAIRunCfMetaExecuteLlama3_1_8bPreviewResponseResultObjectToolCallJSON struct {
	Arguments   apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIRunCfMetaExecuteLlama3_1_8bPreviewResponseResultObjectToolCall) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunCfMetaExecuteLlama3_1_8bPreviewResponseResultObjectToolCallJSON) RawJSON() string {
	return r.raw
}

// Usage statistics for the inference request
type AccountAIRunCfMetaExecuteLlama3_1_8bPreviewResponseResultObjectUsage struct {
	// Total number of tokens in output
	CompletionTokens float64 `json:"completion_tokens"`
	// Total number of tokens in input
	PromptTokens float64 `json:"prompt_tokens"`
	// Total number of input and output tokens
	TotalTokens float64                                                                  `json:"total_tokens"`
	JSON        accountAIRunCfMetaExecuteLlama3_1_8bPreviewResponseResultObjectUsageJSON `json:"-"`
}

// accountAIRunCfMetaExecuteLlama3_1_8bPreviewResponseResultObjectUsageJSON
// contains the JSON metadata for the struct
// [AccountAIRunCfMetaExecuteLlama3_1_8bPreviewResponseResultObjectUsage]
type accountAIRunCfMetaExecuteLlama3_1_8bPreviewResponseResultObjectUsageJSON struct {
	CompletionTokens apijson.Field
	PromptTokens     apijson.Field
	TotalTokens      apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountAIRunCfMetaExecuteLlama3_1_8bPreviewResponseResultObjectUsage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunCfMetaExecuteLlama3_1_8bPreviewResponseResultObjectUsageJSON) RawJSON() string {
	return r.raw
}

type AccountAIRunCfMetaExecuteLlama3_2_11bVisionInstructResponse struct {
	Result  AccountAIRunCfMetaExecuteLlama3_2_11bVisionInstructResponseResultUnion `json:"result" format:"binary"`
	Success bool                                                                   `json:"success"`
	JSON    accountAIRunCfMetaExecuteLlama3_2_11bVisionInstructResponseJSON        `json:"-"`
}

// accountAIRunCfMetaExecuteLlama3_2_11bVisionInstructResponseJSON contains the
// JSON metadata for the struct
// [AccountAIRunCfMetaExecuteLlama3_2_11bVisionInstructResponse]
type accountAIRunCfMetaExecuteLlama3_2_11bVisionInstructResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIRunCfMetaExecuteLlama3_2_11bVisionInstructResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunCfMetaExecuteLlama3_2_11bVisionInstructResponseJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by
// [AccountAIRunCfMetaExecuteLlama3_2_11bVisionInstructResponseResultObject] or
// [shared.UnionString].
type AccountAIRunCfMetaExecuteLlama3_2_11bVisionInstructResponseResultUnion interface {
	ImplementsAccountAIRunCfMetaExecuteLlama3_2_11bVisionInstructResponseResultUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccountAIRunCfMetaExecuteLlama3_2_11bVisionInstructResponseResultUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccountAIRunCfMetaExecuteLlama3_2_11bVisionInstructResponseResultObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type AccountAIRunCfMetaExecuteLlama3_2_11bVisionInstructResponseResultObject struct {
	// The generated text response from the model
	Response string `json:"response"`
	// An array of tool calls requests made during the response generation
	ToolCalls []AccountAIRunCfMetaExecuteLlama3_2_11bVisionInstructResponseResultObjectToolCall `json:"tool_calls"`
	JSON      accountAIRunCfMetaExecuteLlama3_2_11bVisionInstructResponseResultObjectJSON       `json:"-"`
}

// accountAIRunCfMetaExecuteLlama3_2_11bVisionInstructResponseResultObjectJSON
// contains the JSON metadata for the struct
// [AccountAIRunCfMetaExecuteLlama3_2_11bVisionInstructResponseResultObject]
type accountAIRunCfMetaExecuteLlama3_2_11bVisionInstructResponseResultObjectJSON struct {
	Response    apijson.Field
	ToolCalls   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIRunCfMetaExecuteLlama3_2_11bVisionInstructResponseResultObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunCfMetaExecuteLlama3_2_11bVisionInstructResponseResultObjectJSON) RawJSON() string {
	return r.raw
}

func (r AccountAIRunCfMetaExecuteLlama3_2_11bVisionInstructResponseResultObject) ImplementsAccountAIRunCfMetaExecuteLlama3_2_11bVisionInstructResponseResultUnion() {
}

type AccountAIRunCfMetaExecuteLlama3_2_11bVisionInstructResponseResultObjectToolCall struct {
	// The arguments passed to be passed to the tool call request
	Arguments interface{} `json:"arguments"`
	// The name of the tool to be called
	Name string                                                                              `json:"name"`
	JSON accountAIRunCfMetaExecuteLlama3_2_11bVisionInstructResponseResultObjectToolCallJSON `json:"-"`
}

// accountAIRunCfMetaExecuteLlama3_2_11bVisionInstructResponseResultObjectToolCallJSON
// contains the JSON metadata for the struct
// [AccountAIRunCfMetaExecuteLlama3_2_11bVisionInstructResponseResultObjectToolCall]
type accountAIRunCfMetaExecuteLlama3_2_11bVisionInstructResponseResultObjectToolCallJSON struct {
	Arguments   apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIRunCfMetaExecuteLlama3_2_11bVisionInstructResponseResultObjectToolCall) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunCfMetaExecuteLlama3_2_11bVisionInstructResponseResultObjectToolCallJSON) RawJSON() string {
	return r.raw
}

type AccountAIRunCfMetaExecuteLlama3_2_1bInstructResponse struct {
	Result  AccountAIRunCfMetaExecuteLlama3_2_1bInstructResponseResultUnion `json:"result" format:"binary"`
	Success bool                                                            `json:"success"`
	JSON    accountAIRunCfMetaExecuteLlama3_2_1bInstructResponseJSON        `json:"-"`
}

// accountAIRunCfMetaExecuteLlama3_2_1bInstructResponseJSON contains the JSON
// metadata for the struct [AccountAIRunCfMetaExecuteLlama3_2_1bInstructResponse]
type accountAIRunCfMetaExecuteLlama3_2_1bInstructResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIRunCfMetaExecuteLlama3_2_1bInstructResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunCfMetaExecuteLlama3_2_1bInstructResponseJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by
// [AccountAIRunCfMetaExecuteLlama3_2_1bInstructResponseResultObject] or
// [shared.UnionString].
type AccountAIRunCfMetaExecuteLlama3_2_1bInstructResponseResultUnion interface {
	ImplementsAccountAIRunCfMetaExecuteLlama3_2_1bInstructResponseResultUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccountAIRunCfMetaExecuteLlama3_2_1bInstructResponseResultUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccountAIRunCfMetaExecuteLlama3_2_1bInstructResponseResultObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type AccountAIRunCfMetaExecuteLlama3_2_1bInstructResponseResultObject struct {
	// The generated text response from the model
	Response string `json:"response,required"`
	// An array of tool calls requests made during the response generation
	ToolCalls []AccountAIRunCfMetaExecuteLlama3_2_1bInstructResponseResultObjectToolCall `json:"tool_calls"`
	// Usage statistics for the inference request
	Usage AccountAIRunCfMetaExecuteLlama3_2_1bInstructResponseResultObjectUsage `json:"usage"`
	JSON  accountAIRunCfMetaExecuteLlama3_2_1bInstructResponseResultObjectJSON  `json:"-"`
}

// accountAIRunCfMetaExecuteLlama3_2_1bInstructResponseResultObjectJSON contains
// the JSON metadata for the struct
// [AccountAIRunCfMetaExecuteLlama3_2_1bInstructResponseResultObject]
type accountAIRunCfMetaExecuteLlama3_2_1bInstructResponseResultObjectJSON struct {
	Response    apijson.Field
	ToolCalls   apijson.Field
	Usage       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIRunCfMetaExecuteLlama3_2_1bInstructResponseResultObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunCfMetaExecuteLlama3_2_1bInstructResponseResultObjectJSON) RawJSON() string {
	return r.raw
}

func (r AccountAIRunCfMetaExecuteLlama3_2_1bInstructResponseResultObject) ImplementsAccountAIRunCfMetaExecuteLlama3_2_1bInstructResponseResultUnion() {
}

type AccountAIRunCfMetaExecuteLlama3_2_1bInstructResponseResultObjectToolCall struct {
	// The arguments passed to be passed to the tool call request
	Arguments interface{} `json:"arguments"`
	// The name of the tool to be called
	Name string                                                                       `json:"name"`
	JSON accountAIRunCfMetaExecuteLlama3_2_1bInstructResponseResultObjectToolCallJSON `json:"-"`
}

// accountAIRunCfMetaExecuteLlama3_2_1bInstructResponseResultObjectToolCallJSON
// contains the JSON metadata for the struct
// [AccountAIRunCfMetaExecuteLlama3_2_1bInstructResponseResultObjectToolCall]
type accountAIRunCfMetaExecuteLlama3_2_1bInstructResponseResultObjectToolCallJSON struct {
	Arguments   apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIRunCfMetaExecuteLlama3_2_1bInstructResponseResultObjectToolCall) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunCfMetaExecuteLlama3_2_1bInstructResponseResultObjectToolCallJSON) RawJSON() string {
	return r.raw
}

// Usage statistics for the inference request
type AccountAIRunCfMetaExecuteLlama3_2_1bInstructResponseResultObjectUsage struct {
	// Total number of tokens in output
	CompletionTokens float64 `json:"completion_tokens"`
	// Total number of tokens in input
	PromptTokens float64 `json:"prompt_tokens"`
	// Total number of input and output tokens
	TotalTokens float64                                                                   `json:"total_tokens"`
	JSON        accountAIRunCfMetaExecuteLlama3_2_1bInstructResponseResultObjectUsageJSON `json:"-"`
}

// accountAIRunCfMetaExecuteLlama3_2_1bInstructResponseResultObjectUsageJSON
// contains the JSON metadata for the struct
// [AccountAIRunCfMetaExecuteLlama3_2_1bInstructResponseResultObjectUsage]
type accountAIRunCfMetaExecuteLlama3_2_1bInstructResponseResultObjectUsageJSON struct {
	CompletionTokens apijson.Field
	PromptTokens     apijson.Field
	TotalTokens      apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountAIRunCfMetaExecuteLlama3_2_1bInstructResponseResultObjectUsage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunCfMetaExecuteLlama3_2_1bInstructResponseResultObjectUsageJSON) RawJSON() string {
	return r.raw
}

type AccountAIRunCfMetaExecuteLlama3_2_3bInstructResponse struct {
	Result  AccountAIRunCfMetaExecuteLlama3_2_3bInstructResponseResultUnion `json:"result" format:"binary"`
	Success bool                                                            `json:"success"`
	JSON    accountAIRunCfMetaExecuteLlama3_2_3bInstructResponseJSON        `json:"-"`
}

// accountAIRunCfMetaExecuteLlama3_2_3bInstructResponseJSON contains the JSON
// metadata for the struct [AccountAIRunCfMetaExecuteLlama3_2_3bInstructResponse]
type accountAIRunCfMetaExecuteLlama3_2_3bInstructResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIRunCfMetaExecuteLlama3_2_3bInstructResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunCfMetaExecuteLlama3_2_3bInstructResponseJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by
// [AccountAIRunCfMetaExecuteLlama3_2_3bInstructResponseResultObject] or
// [shared.UnionString].
type AccountAIRunCfMetaExecuteLlama3_2_3bInstructResponseResultUnion interface {
	ImplementsAccountAIRunCfMetaExecuteLlama3_2_3bInstructResponseResultUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccountAIRunCfMetaExecuteLlama3_2_3bInstructResponseResultUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccountAIRunCfMetaExecuteLlama3_2_3bInstructResponseResultObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type AccountAIRunCfMetaExecuteLlama3_2_3bInstructResponseResultObject struct {
	// The generated text response from the model
	Response string `json:"response,required"`
	// An array of tool calls requests made during the response generation
	ToolCalls []AccountAIRunCfMetaExecuteLlama3_2_3bInstructResponseResultObjectToolCall `json:"tool_calls"`
	// Usage statistics for the inference request
	Usage AccountAIRunCfMetaExecuteLlama3_2_3bInstructResponseResultObjectUsage `json:"usage"`
	JSON  accountAIRunCfMetaExecuteLlama3_2_3bInstructResponseResultObjectJSON  `json:"-"`
}

// accountAIRunCfMetaExecuteLlama3_2_3bInstructResponseResultObjectJSON contains
// the JSON metadata for the struct
// [AccountAIRunCfMetaExecuteLlama3_2_3bInstructResponseResultObject]
type accountAIRunCfMetaExecuteLlama3_2_3bInstructResponseResultObjectJSON struct {
	Response    apijson.Field
	ToolCalls   apijson.Field
	Usage       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIRunCfMetaExecuteLlama3_2_3bInstructResponseResultObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunCfMetaExecuteLlama3_2_3bInstructResponseResultObjectJSON) RawJSON() string {
	return r.raw
}

func (r AccountAIRunCfMetaExecuteLlama3_2_3bInstructResponseResultObject) ImplementsAccountAIRunCfMetaExecuteLlama3_2_3bInstructResponseResultUnion() {
}

type AccountAIRunCfMetaExecuteLlama3_2_3bInstructResponseResultObjectToolCall struct {
	// The arguments passed to be passed to the tool call request
	Arguments interface{} `json:"arguments"`
	// The name of the tool to be called
	Name string                                                                       `json:"name"`
	JSON accountAIRunCfMetaExecuteLlama3_2_3bInstructResponseResultObjectToolCallJSON `json:"-"`
}

// accountAIRunCfMetaExecuteLlama3_2_3bInstructResponseResultObjectToolCallJSON
// contains the JSON metadata for the struct
// [AccountAIRunCfMetaExecuteLlama3_2_3bInstructResponseResultObjectToolCall]
type accountAIRunCfMetaExecuteLlama3_2_3bInstructResponseResultObjectToolCallJSON struct {
	Arguments   apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIRunCfMetaExecuteLlama3_2_3bInstructResponseResultObjectToolCall) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunCfMetaExecuteLlama3_2_3bInstructResponseResultObjectToolCallJSON) RawJSON() string {
	return r.raw
}

// Usage statistics for the inference request
type AccountAIRunCfMetaExecuteLlama3_2_3bInstructResponseResultObjectUsage struct {
	// Total number of tokens in output
	CompletionTokens float64 `json:"completion_tokens"`
	// Total number of tokens in input
	PromptTokens float64 `json:"prompt_tokens"`
	// Total number of input and output tokens
	TotalTokens float64                                                                   `json:"total_tokens"`
	JSON        accountAIRunCfMetaExecuteLlama3_2_3bInstructResponseResultObjectUsageJSON `json:"-"`
}

// accountAIRunCfMetaExecuteLlama3_2_3bInstructResponseResultObjectUsageJSON
// contains the JSON metadata for the struct
// [AccountAIRunCfMetaExecuteLlama3_2_3bInstructResponseResultObjectUsage]
type accountAIRunCfMetaExecuteLlama3_2_3bInstructResponseResultObjectUsageJSON struct {
	CompletionTokens apijson.Field
	PromptTokens     apijson.Field
	TotalTokens      apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountAIRunCfMetaExecuteLlama3_2_3bInstructResponseResultObjectUsage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunCfMetaExecuteLlama3_2_3bInstructResponseResultObjectUsageJSON) RawJSON() string {
	return r.raw
}

type AccountAIRunCfMetaExecuteLlama3_3_70bInstructFp8FastResponse struct {
	Result  AccountAIRunCfMetaExecuteLlama3_3_70bInstructFp8FastResponseResultUnion `json:"result" format:"binary"`
	Success bool                                                                    `json:"success"`
	JSON    accountAIRunCfMetaExecuteLlama3_3_70bInstructFp8FastResponseJSON        `json:"-"`
}

// accountAIRunCfMetaExecuteLlama3_3_70bInstructFp8FastResponseJSON contains the
// JSON metadata for the struct
// [AccountAIRunCfMetaExecuteLlama3_3_70bInstructFp8FastResponse]
type accountAIRunCfMetaExecuteLlama3_3_70bInstructFp8FastResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIRunCfMetaExecuteLlama3_3_70bInstructFp8FastResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunCfMetaExecuteLlama3_3_70bInstructFp8FastResponseJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by
// [AccountAIRunCfMetaExecuteLlama3_3_70bInstructFp8FastResponseResultObject] or
// [shared.UnionString].
type AccountAIRunCfMetaExecuteLlama3_3_70bInstructFp8FastResponseResultUnion interface {
	ImplementsAccountAIRunCfMetaExecuteLlama3_3_70bInstructFp8FastResponseResultUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccountAIRunCfMetaExecuteLlama3_3_70bInstructFp8FastResponseResultUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccountAIRunCfMetaExecuteLlama3_3_70bInstructFp8FastResponseResultObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type AccountAIRunCfMetaExecuteLlama3_3_70bInstructFp8FastResponseResultObject struct {
	// The generated text response from the model
	Response string `json:"response,required"`
	// An array of tool calls requests made during the response generation
	ToolCalls []AccountAIRunCfMetaExecuteLlama3_3_70bInstructFp8FastResponseResultObjectToolCall `json:"tool_calls"`
	// Usage statistics for the inference request
	Usage AccountAIRunCfMetaExecuteLlama3_3_70bInstructFp8FastResponseResultObjectUsage `json:"usage"`
	JSON  accountAIRunCfMetaExecuteLlama3_3_70bInstructFp8FastResponseResultObjectJSON  `json:"-"`
}

// accountAIRunCfMetaExecuteLlama3_3_70bInstructFp8FastResponseResultObjectJSON
// contains the JSON metadata for the struct
// [AccountAIRunCfMetaExecuteLlama3_3_70bInstructFp8FastResponseResultObject]
type accountAIRunCfMetaExecuteLlama3_3_70bInstructFp8FastResponseResultObjectJSON struct {
	Response    apijson.Field
	ToolCalls   apijson.Field
	Usage       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIRunCfMetaExecuteLlama3_3_70bInstructFp8FastResponseResultObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunCfMetaExecuteLlama3_3_70bInstructFp8FastResponseResultObjectJSON) RawJSON() string {
	return r.raw
}

func (r AccountAIRunCfMetaExecuteLlama3_3_70bInstructFp8FastResponseResultObject) ImplementsAccountAIRunCfMetaExecuteLlama3_3_70bInstructFp8FastResponseResultUnion() {
}

type AccountAIRunCfMetaExecuteLlama3_3_70bInstructFp8FastResponseResultObjectToolCall struct {
	// The arguments passed to be passed to the tool call request
	Arguments interface{} `json:"arguments"`
	// The name of the tool to be called
	Name string                                                                               `json:"name"`
	JSON accountAIRunCfMetaExecuteLlama3_3_70bInstructFp8FastResponseResultObjectToolCallJSON `json:"-"`
}

// accountAIRunCfMetaExecuteLlama3_3_70bInstructFp8FastResponseResultObjectToolCallJSON
// contains the JSON metadata for the struct
// [AccountAIRunCfMetaExecuteLlama3_3_70bInstructFp8FastResponseResultObjectToolCall]
type accountAIRunCfMetaExecuteLlama3_3_70bInstructFp8FastResponseResultObjectToolCallJSON struct {
	Arguments   apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIRunCfMetaExecuteLlama3_3_70bInstructFp8FastResponseResultObjectToolCall) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunCfMetaExecuteLlama3_3_70bInstructFp8FastResponseResultObjectToolCallJSON) RawJSON() string {
	return r.raw
}

// Usage statistics for the inference request
type AccountAIRunCfMetaExecuteLlama3_3_70bInstructFp8FastResponseResultObjectUsage struct {
	// Total number of tokens in output
	CompletionTokens float64 `json:"completion_tokens"`
	// Total number of tokens in input
	PromptTokens float64 `json:"prompt_tokens"`
	// Total number of input and output tokens
	TotalTokens float64                                                                           `json:"total_tokens"`
	JSON        accountAIRunCfMetaExecuteLlama3_3_70bInstructFp8FastResponseResultObjectUsageJSON `json:"-"`
}

// accountAIRunCfMetaExecuteLlama3_3_70bInstructFp8FastResponseResultObjectUsageJSON
// contains the JSON metadata for the struct
// [AccountAIRunCfMetaExecuteLlama3_3_70bInstructFp8FastResponseResultObjectUsage]
type accountAIRunCfMetaExecuteLlama3_3_70bInstructFp8FastResponseResultObjectUsageJSON struct {
	CompletionTokens apijson.Field
	PromptTokens     apijson.Field
	TotalTokens      apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountAIRunCfMetaExecuteLlama3_3_70bInstructFp8FastResponseResultObjectUsage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunCfMetaExecuteLlama3_3_70bInstructFp8FastResponseResultObjectUsageJSON) RawJSON() string {
	return r.raw
}

type AccountAIRunCfMetaExecuteLlama3_8bInstructResponse struct {
	Result  AccountAIRunCfMetaExecuteLlama3_8bInstructResponseResultUnion `json:"result" format:"binary"`
	Success bool                                                          `json:"success"`
	JSON    accountAIRunCfMetaExecuteLlama3_8bInstructResponseJSON        `json:"-"`
}

// accountAIRunCfMetaExecuteLlama3_8bInstructResponseJSON contains the JSON
// metadata for the struct [AccountAIRunCfMetaExecuteLlama3_8bInstructResponse]
type accountAIRunCfMetaExecuteLlama3_8bInstructResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIRunCfMetaExecuteLlama3_8bInstructResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunCfMetaExecuteLlama3_8bInstructResponseJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by
// [AccountAIRunCfMetaExecuteLlama3_8bInstructResponseResultObject] or
// [shared.UnionString].
type AccountAIRunCfMetaExecuteLlama3_8bInstructResponseResultUnion interface {
	ImplementsAccountAIRunCfMetaExecuteLlama3_8bInstructResponseResultUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccountAIRunCfMetaExecuteLlama3_8bInstructResponseResultUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccountAIRunCfMetaExecuteLlama3_8bInstructResponseResultObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type AccountAIRunCfMetaExecuteLlama3_8bInstructResponseResultObject struct {
	// The generated text response from the model
	Response string `json:"response,required"`
	// An array of tool calls requests made during the response generation
	ToolCalls []AccountAIRunCfMetaExecuteLlama3_8bInstructResponseResultObjectToolCall `json:"tool_calls"`
	// Usage statistics for the inference request
	Usage AccountAIRunCfMetaExecuteLlama3_8bInstructResponseResultObjectUsage `json:"usage"`
	JSON  accountAIRunCfMetaExecuteLlama3_8bInstructResponseResultObjectJSON  `json:"-"`
}

// accountAIRunCfMetaExecuteLlama3_8bInstructResponseResultObjectJSON contains the
// JSON metadata for the struct
// [AccountAIRunCfMetaExecuteLlama3_8bInstructResponseResultObject]
type accountAIRunCfMetaExecuteLlama3_8bInstructResponseResultObjectJSON struct {
	Response    apijson.Field
	ToolCalls   apijson.Field
	Usage       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIRunCfMetaExecuteLlama3_8bInstructResponseResultObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunCfMetaExecuteLlama3_8bInstructResponseResultObjectJSON) RawJSON() string {
	return r.raw
}

func (r AccountAIRunCfMetaExecuteLlama3_8bInstructResponseResultObject) ImplementsAccountAIRunCfMetaExecuteLlama3_8bInstructResponseResultUnion() {
}

type AccountAIRunCfMetaExecuteLlama3_8bInstructResponseResultObjectToolCall struct {
	// The arguments passed to be passed to the tool call request
	Arguments interface{} `json:"arguments"`
	// The name of the tool to be called
	Name string                                                                     `json:"name"`
	JSON accountAIRunCfMetaExecuteLlama3_8bInstructResponseResultObjectToolCallJSON `json:"-"`
}

// accountAIRunCfMetaExecuteLlama3_8bInstructResponseResultObjectToolCallJSON
// contains the JSON metadata for the struct
// [AccountAIRunCfMetaExecuteLlama3_8bInstructResponseResultObjectToolCall]
type accountAIRunCfMetaExecuteLlama3_8bInstructResponseResultObjectToolCallJSON struct {
	Arguments   apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIRunCfMetaExecuteLlama3_8bInstructResponseResultObjectToolCall) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunCfMetaExecuteLlama3_8bInstructResponseResultObjectToolCallJSON) RawJSON() string {
	return r.raw
}

// Usage statistics for the inference request
type AccountAIRunCfMetaExecuteLlama3_8bInstructResponseResultObjectUsage struct {
	// Total number of tokens in output
	CompletionTokens float64 `json:"completion_tokens"`
	// Total number of tokens in input
	PromptTokens float64 `json:"prompt_tokens"`
	// Total number of input and output tokens
	TotalTokens float64                                                                 `json:"total_tokens"`
	JSON        accountAIRunCfMetaExecuteLlama3_8bInstructResponseResultObjectUsageJSON `json:"-"`
}

// accountAIRunCfMetaExecuteLlama3_8bInstructResponseResultObjectUsageJSON contains
// the JSON metadata for the struct
// [AccountAIRunCfMetaExecuteLlama3_8bInstructResponseResultObjectUsage]
type accountAIRunCfMetaExecuteLlama3_8bInstructResponseResultObjectUsageJSON struct {
	CompletionTokens apijson.Field
	PromptTokens     apijson.Field
	TotalTokens      apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountAIRunCfMetaExecuteLlama3_8bInstructResponseResultObjectUsage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunCfMetaExecuteLlama3_8bInstructResponseResultObjectUsageJSON) RawJSON() string {
	return r.raw
}

type AccountAIRunCfMetaExecuteLlama3_8bInstructAwqResponse struct {
	Result  AccountAIRunCfMetaExecuteLlama3_8bInstructAwqResponseResultUnion `json:"result" format:"binary"`
	Success bool                                                             `json:"success"`
	JSON    accountAIRunCfMetaExecuteLlama3_8bInstructAwqResponseJSON        `json:"-"`
}

// accountAIRunCfMetaExecuteLlama3_8bInstructAwqResponseJSON contains the JSON
// metadata for the struct [AccountAIRunCfMetaExecuteLlama3_8bInstructAwqResponse]
type accountAIRunCfMetaExecuteLlama3_8bInstructAwqResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIRunCfMetaExecuteLlama3_8bInstructAwqResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunCfMetaExecuteLlama3_8bInstructAwqResponseJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by
// [AccountAIRunCfMetaExecuteLlama3_8bInstructAwqResponseResultObject] or
// [shared.UnionString].
type AccountAIRunCfMetaExecuteLlama3_8bInstructAwqResponseResultUnion interface {
	ImplementsAccountAIRunCfMetaExecuteLlama3_8bInstructAwqResponseResultUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccountAIRunCfMetaExecuteLlama3_8bInstructAwqResponseResultUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccountAIRunCfMetaExecuteLlama3_8bInstructAwqResponseResultObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type AccountAIRunCfMetaExecuteLlama3_8bInstructAwqResponseResultObject struct {
	// The generated text response from the model
	Response string `json:"response,required"`
	// An array of tool calls requests made during the response generation
	ToolCalls []AccountAIRunCfMetaExecuteLlama3_8bInstructAwqResponseResultObjectToolCall `json:"tool_calls"`
	// Usage statistics for the inference request
	Usage AccountAIRunCfMetaExecuteLlama3_8bInstructAwqResponseResultObjectUsage `json:"usage"`
	JSON  accountAIRunCfMetaExecuteLlama3_8bInstructAwqResponseResultObjectJSON  `json:"-"`
}

// accountAIRunCfMetaExecuteLlama3_8bInstructAwqResponseResultObjectJSON contains
// the JSON metadata for the struct
// [AccountAIRunCfMetaExecuteLlama3_8bInstructAwqResponseResultObject]
type accountAIRunCfMetaExecuteLlama3_8bInstructAwqResponseResultObjectJSON struct {
	Response    apijson.Field
	ToolCalls   apijson.Field
	Usage       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIRunCfMetaExecuteLlama3_8bInstructAwqResponseResultObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunCfMetaExecuteLlama3_8bInstructAwqResponseResultObjectJSON) RawJSON() string {
	return r.raw
}

func (r AccountAIRunCfMetaExecuteLlama3_8bInstructAwqResponseResultObject) ImplementsAccountAIRunCfMetaExecuteLlama3_8bInstructAwqResponseResultUnion() {
}

type AccountAIRunCfMetaExecuteLlama3_8bInstructAwqResponseResultObjectToolCall struct {
	// The arguments passed to be passed to the tool call request
	Arguments interface{} `json:"arguments"`
	// The name of the tool to be called
	Name string                                                                        `json:"name"`
	JSON accountAIRunCfMetaExecuteLlama3_8bInstructAwqResponseResultObjectToolCallJSON `json:"-"`
}

// accountAIRunCfMetaExecuteLlama3_8bInstructAwqResponseResultObjectToolCallJSON
// contains the JSON metadata for the struct
// [AccountAIRunCfMetaExecuteLlama3_8bInstructAwqResponseResultObjectToolCall]
type accountAIRunCfMetaExecuteLlama3_8bInstructAwqResponseResultObjectToolCallJSON struct {
	Arguments   apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIRunCfMetaExecuteLlama3_8bInstructAwqResponseResultObjectToolCall) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunCfMetaExecuteLlama3_8bInstructAwqResponseResultObjectToolCallJSON) RawJSON() string {
	return r.raw
}

// Usage statistics for the inference request
type AccountAIRunCfMetaExecuteLlama3_8bInstructAwqResponseResultObjectUsage struct {
	// Total number of tokens in output
	CompletionTokens float64 `json:"completion_tokens"`
	// Total number of tokens in input
	PromptTokens float64 `json:"prompt_tokens"`
	// Total number of input and output tokens
	TotalTokens float64                                                                    `json:"total_tokens"`
	JSON        accountAIRunCfMetaExecuteLlama3_8bInstructAwqResponseResultObjectUsageJSON `json:"-"`
}

// accountAIRunCfMetaExecuteLlama3_8bInstructAwqResponseResultObjectUsageJSON
// contains the JSON metadata for the struct
// [AccountAIRunCfMetaExecuteLlama3_8bInstructAwqResponseResultObjectUsage]
type accountAIRunCfMetaExecuteLlama3_8bInstructAwqResponseResultObjectUsageJSON struct {
	CompletionTokens apijson.Field
	PromptTokens     apijson.Field
	TotalTokens      apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountAIRunCfMetaExecuteLlama3_8bInstructAwqResponseResultObjectUsage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunCfMetaExecuteLlama3_8bInstructAwqResponseResultObjectUsageJSON) RawJSON() string {
	return r.raw
}

type AccountAIRunCfMetaExecuteLlamaGuard3_8bResponse struct {
	Result  AccountAIRunCfMetaExecuteLlamaGuard3_8bResponseResult `json:"result"`
	Success bool                                                  `json:"success"`
	JSON    accountAIRunCfMetaExecuteLlamaGuard3_8bResponseJSON   `json:"-"`
}

// accountAIRunCfMetaExecuteLlamaGuard3_8bResponseJSON contains the JSON metadata
// for the struct [AccountAIRunCfMetaExecuteLlamaGuard3_8bResponse]
type accountAIRunCfMetaExecuteLlamaGuard3_8bResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIRunCfMetaExecuteLlamaGuard3_8bResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunCfMetaExecuteLlamaGuard3_8bResponseJSON) RawJSON() string {
	return r.raw
}

type AccountAIRunCfMetaExecuteLlamaGuard3_8bResponseResult struct {
	// The generated text response from the model.
	Response AccountAIRunCfMetaExecuteLlamaGuard3_8bResponseResultResponseUnion `json:"response"`
	// Usage statistics for the inference request
	Usage AccountAIRunCfMetaExecuteLlamaGuard3_8bResponseResultUsage `json:"usage"`
	JSON  accountAIRunCfMetaExecuteLlamaGuard3_8bResponseResultJSON  `json:"-"`
}

// accountAIRunCfMetaExecuteLlamaGuard3_8bResponseResultJSON contains the JSON
// metadata for the struct [AccountAIRunCfMetaExecuteLlamaGuard3_8bResponseResult]
type accountAIRunCfMetaExecuteLlamaGuard3_8bResponseResultJSON struct {
	Response    apijson.Field
	Usage       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIRunCfMetaExecuteLlamaGuard3_8bResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunCfMetaExecuteLlamaGuard3_8bResponseResultJSON) RawJSON() string {
	return r.raw
}

// The generated text response from the model.
//
// Union satisfied by [shared.UnionString] or
// [AccountAIRunCfMetaExecuteLlamaGuard3_8bResponseResultResponseObject].
type AccountAIRunCfMetaExecuteLlamaGuard3_8bResponseResultResponseUnion interface {
	ImplementsAccountAIRunCfMetaExecuteLlamaGuard3_8bResponseResultResponseUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccountAIRunCfMetaExecuteLlamaGuard3_8bResponseResultResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccountAIRunCfMetaExecuteLlamaGuard3_8bResponseResultResponseObject{}),
		},
	)
}

// The json response parsed from the generated text response from the model.
type AccountAIRunCfMetaExecuteLlamaGuard3_8bResponseResultResponseObject struct {
	// A list of what hazard categories predicted for the conversation, if the
	// conversation is deemed unsafe.
	Categories []string `json:"categories"`
	// Whether the conversation is safe or not.
	Safe bool                                                                    `json:"safe"`
	JSON accountAIRunCfMetaExecuteLlamaGuard3_8bResponseResultResponseObjectJSON `json:"-"`
}

// accountAIRunCfMetaExecuteLlamaGuard3_8bResponseResultResponseObjectJSON contains
// the JSON metadata for the struct
// [AccountAIRunCfMetaExecuteLlamaGuard3_8bResponseResultResponseObject]
type accountAIRunCfMetaExecuteLlamaGuard3_8bResponseResultResponseObjectJSON struct {
	Categories  apijson.Field
	Safe        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIRunCfMetaExecuteLlamaGuard3_8bResponseResultResponseObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunCfMetaExecuteLlamaGuard3_8bResponseResultResponseObjectJSON) RawJSON() string {
	return r.raw
}

func (r AccountAIRunCfMetaExecuteLlamaGuard3_8bResponseResultResponseObject) ImplementsAccountAIRunCfMetaExecuteLlamaGuard3_8bResponseResultResponseUnion() {
}

// Usage statistics for the inference request
type AccountAIRunCfMetaExecuteLlamaGuard3_8bResponseResultUsage struct {
	// Total number of tokens in output
	CompletionTokens float64 `json:"completion_tokens"`
	// Total number of tokens in input
	PromptTokens float64 `json:"prompt_tokens"`
	// Total number of input and output tokens
	TotalTokens float64                                                        `json:"total_tokens"`
	JSON        accountAIRunCfMetaExecuteLlamaGuard3_8bResponseResultUsageJSON `json:"-"`
}

// accountAIRunCfMetaExecuteLlamaGuard3_8bResponseResultUsageJSON contains the JSON
// metadata for the struct
// [AccountAIRunCfMetaExecuteLlamaGuard3_8bResponseResultUsage]
type accountAIRunCfMetaExecuteLlamaGuard3_8bResponseResultUsageJSON struct {
	CompletionTokens apijson.Field
	PromptTokens     apijson.Field
	TotalTokens      apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountAIRunCfMetaExecuteLlamaGuard3_8bResponseResultUsage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunCfMetaExecuteLlamaGuard3_8bResponseResultUsageJSON) RawJSON() string {
	return r.raw
}

type AccountAIRunCfMetaExecuteM2m100_1_2bResponse struct {
	Result  AccountAIRunCfMetaExecuteM2m100_1_2bResponseResult `json:"result"`
	Success bool                                               `json:"success"`
	JSON    accountAIRunCfMetaExecuteM2m100_1_2bResponseJSON   `json:"-"`
}

// accountAIRunCfMetaExecuteM2m100_1_2bResponseJSON contains the JSON metadata for
// the struct [AccountAIRunCfMetaExecuteM2m100_1_2bResponse]
type accountAIRunCfMetaExecuteM2m100_1_2bResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIRunCfMetaExecuteM2m100_1_2bResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunCfMetaExecuteM2m100_1_2bResponseJSON) RawJSON() string {
	return r.raw
}

type AccountAIRunCfMetaExecuteM2m100_1_2bResponseResult struct {
	// The translated text in the target language
	TranslatedText string                                                 `json:"translated_text"`
	JSON           accountAIRunCfMetaExecuteM2m100_1_2bResponseResultJSON `json:"-"`
}

// accountAIRunCfMetaExecuteM2m100_1_2bResponseResultJSON contains the JSON
// metadata for the struct [AccountAIRunCfMetaExecuteM2m100_1_2bResponseResult]
type accountAIRunCfMetaExecuteM2m100_1_2bResponseResultJSON struct {
	TranslatedText apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccountAIRunCfMetaExecuteM2m100_1_2bResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunCfMetaExecuteM2m100_1_2bResponseResultJSON) RawJSON() string {
	return r.raw
}

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

type AccountAIRunCfMetaExecuteLlama3_1_8bInstructParams struct {
	QueueRequest param.Field[string]                                         `query:"queueRequest"`
	Body         AccountAIRunCfMetaExecuteLlama3_1_8bInstructParamsBodyUnion `json:"body"`
}

func (r AccountAIRunCfMetaExecuteLlama3_1_8bInstructParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

// URLQuery serializes [AccountAIRunCfMetaExecuteLlama3_1_8bInstructParams]'s query
// parameters as `url.Values`.
func (r AccountAIRunCfMetaExecuteLlama3_1_8bInstructParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountAIRunCfMetaExecuteLlama3_1_8bInstructParamsBody struct {
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

func (r AccountAIRunCfMetaExecuteLlama3_1_8bInstructParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfMetaExecuteLlama3_1_8bInstructParamsBody) implementsAccountAIRunCfMetaExecuteLlama3_1_8bInstructParamsBodyUnion() {
}

// Satisfied by [AccountAIRunCfMetaExecuteLlama3_1_8bInstructParamsBodyPrompt],
// [AccountAIRunCfMetaExecuteLlama3_1_8bInstructParamsBodyMessages],
// [AccountAIRunCfMetaExecuteLlama3_1_8bInstructParamsBody].
type AccountAIRunCfMetaExecuteLlama3_1_8bInstructParamsBodyUnion interface {
	implementsAccountAIRunCfMetaExecuteLlama3_1_8bInstructParamsBodyUnion()
}

type AccountAIRunCfMetaExecuteLlama3_1_8bInstructParamsBodyPrompt struct {
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
	ResponseFormat    param.Field[AccountAIRunCfMetaExecuteLlama3_1_8bInstructParamsBodyPromptResponseFormat] `json:"response_format"`
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

func (r AccountAIRunCfMetaExecuteLlama3_1_8bInstructParamsBodyPrompt) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfMetaExecuteLlama3_1_8bInstructParamsBodyPrompt) implementsAccountAIRunCfMetaExecuteLlama3_1_8bInstructParamsBodyUnion() {
}

type AccountAIRunCfMetaExecuteLlama3_1_8bInstructParamsBodyPromptResponseFormat struct {
	JsonSchema param.Field[interface{}]                                                                    `json:"json_schema"`
	Type       param.Field[AccountAIRunCfMetaExecuteLlama3_1_8bInstructParamsBodyPromptResponseFormatType] `json:"type"`
}

func (r AccountAIRunCfMetaExecuteLlama3_1_8bInstructParamsBodyPromptResponseFormat) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfMetaExecuteLlama3_1_8bInstructParamsBodyPromptResponseFormatType string

const (
	AccountAIRunCfMetaExecuteLlama3_1_8bInstructParamsBodyPromptResponseFormatTypeJsonObject AccountAIRunCfMetaExecuteLlama3_1_8bInstructParamsBodyPromptResponseFormatType = "json_object"
	AccountAIRunCfMetaExecuteLlama3_1_8bInstructParamsBodyPromptResponseFormatTypeJsonSchema AccountAIRunCfMetaExecuteLlama3_1_8bInstructParamsBodyPromptResponseFormatType = "json_schema"
)

func (r AccountAIRunCfMetaExecuteLlama3_1_8bInstructParamsBodyPromptResponseFormatType) IsKnown() bool {
	switch r {
	case AccountAIRunCfMetaExecuteLlama3_1_8bInstructParamsBodyPromptResponseFormatTypeJsonObject, AccountAIRunCfMetaExecuteLlama3_1_8bInstructParamsBodyPromptResponseFormatTypeJsonSchema:
		return true
	}
	return false
}

type AccountAIRunCfMetaExecuteLlama3_1_8bInstructParamsBodyMessages struct {
	// An array of message objects representing the conversation history.
	Messages param.Field[[]AccountAIRunCfMetaExecuteLlama3_1_8bInstructParamsBodyMessagesMessage] `json:"messages,required"`
	// Decreases the likelihood of the model repeating the same lines verbatim.
	FrequencyPenalty param.Field[float64]                                                                  `json:"frequency_penalty"`
	Functions        param.Field[[]AccountAIRunCfMetaExecuteLlama3_1_8bInstructParamsBodyMessagesFunction] `json:"functions"`
	// The maximum number of tokens to generate in the response.
	MaxTokens param.Field[int64] `json:"max_tokens"`
	// Increases the likelihood of the model introducing new topics.
	PresencePenalty param.Field[float64] `json:"presence_penalty"`
	// If true, a chat template is not applied and you must adhere to the specific
	// model's expected formatting.
	Raw param.Field[bool] `json:"raw"`
	// Penalty for repeated tokens; higher values discourage repetition.
	RepetitionPenalty param.Field[float64]                                                                      `json:"repetition_penalty"`
	ResponseFormat    param.Field[AccountAIRunCfMetaExecuteLlama3_1_8bInstructParamsBodyMessagesResponseFormat] `json:"response_format"`
	// Random seed for reproducibility of the generation.
	Seed param.Field[int64] `json:"seed"`
	// If true, the response will be streamed back incrementally using SSE, Server Sent
	// Events.
	Stream param.Field[bool] `json:"stream"`
	// Controls the randomness of the output; higher values produce more random
	// results.
	Temperature param.Field[float64] `json:"temperature"`
	// A list of tools available for the assistant to use.
	Tools param.Field[[]AccountAIRunCfMetaExecuteLlama3_1_8bInstructParamsBodyMessagesToolUnion] `json:"tools"`
	// Limits the AI to choose from the top 'k' most probable words. Lower values make
	// responses more focused; higher values introduce more variety and potential
	// surprises.
	TopK param.Field[int64] `json:"top_k"`
	// Adjusts the creativity of the AI's responses by controlling how many possible
	// words it considers. Lower values make outputs more predictable; higher values
	// allow for more varied and creative responses.
	TopP param.Field[float64] `json:"top_p"`
}

func (r AccountAIRunCfMetaExecuteLlama3_1_8bInstructParamsBodyMessages) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfMetaExecuteLlama3_1_8bInstructParamsBodyMessages) implementsAccountAIRunCfMetaExecuteLlama3_1_8bInstructParamsBodyUnion() {
}

type AccountAIRunCfMetaExecuteLlama3_1_8bInstructParamsBodyMessagesMessage struct {
	// The content of the message as a string.
	Content param.Field[string] `json:"content,required"`
	// The role of the message sender (e.g., 'user', 'assistant', 'system', 'tool').
	Role param.Field[string] `json:"role,required"`
}

func (r AccountAIRunCfMetaExecuteLlama3_1_8bInstructParamsBodyMessagesMessage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfMetaExecuteLlama3_1_8bInstructParamsBodyMessagesFunction struct {
	Code param.Field[string] `json:"code,required"`
	Name param.Field[string] `json:"name,required"`
}

func (r AccountAIRunCfMetaExecuteLlama3_1_8bInstructParamsBodyMessagesFunction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfMetaExecuteLlama3_1_8bInstructParamsBodyMessagesResponseFormat struct {
	JsonSchema param.Field[interface{}]                                                                      `json:"json_schema"`
	Type       param.Field[AccountAIRunCfMetaExecuteLlama3_1_8bInstructParamsBodyMessagesResponseFormatType] `json:"type"`
}

func (r AccountAIRunCfMetaExecuteLlama3_1_8bInstructParamsBodyMessagesResponseFormat) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfMetaExecuteLlama3_1_8bInstructParamsBodyMessagesResponseFormatType string

const (
	AccountAIRunCfMetaExecuteLlama3_1_8bInstructParamsBodyMessagesResponseFormatTypeJsonObject AccountAIRunCfMetaExecuteLlama3_1_8bInstructParamsBodyMessagesResponseFormatType = "json_object"
	AccountAIRunCfMetaExecuteLlama3_1_8bInstructParamsBodyMessagesResponseFormatTypeJsonSchema AccountAIRunCfMetaExecuteLlama3_1_8bInstructParamsBodyMessagesResponseFormatType = "json_schema"
)

func (r AccountAIRunCfMetaExecuteLlama3_1_8bInstructParamsBodyMessagesResponseFormatType) IsKnown() bool {
	switch r {
	case AccountAIRunCfMetaExecuteLlama3_1_8bInstructParamsBodyMessagesResponseFormatTypeJsonObject, AccountAIRunCfMetaExecuteLlama3_1_8bInstructParamsBodyMessagesResponseFormatTypeJsonSchema:
		return true
	}
	return false
}

type AccountAIRunCfMetaExecuteLlama3_1_8bInstructParamsBodyMessagesTool struct {
	// A brief description of what the tool does.
	Description param.Field[string]      `json:"description"`
	Function    param.Field[interface{}] `json:"function"`
	// The name of the tool. More descriptive the better.
	Name       param.Field[string]      `json:"name"`
	Parameters param.Field[interface{}] `json:"parameters"`
	// Specifies the type of tool (e.g., 'function').
	Type param.Field[string] `json:"type"`
}

func (r AccountAIRunCfMetaExecuteLlama3_1_8bInstructParamsBodyMessagesTool) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfMetaExecuteLlama3_1_8bInstructParamsBodyMessagesTool) implementsAccountAIRunCfMetaExecuteLlama3_1_8bInstructParamsBodyMessagesToolUnion() {
}

// Satisfied by
// [AccountAIRunCfMetaExecuteLlama3_1_8bInstructParamsBodyMessagesToolsObject],
// [AccountAIRunCfMetaExecuteLlama3_1_8bInstructParamsBodyMessagesToolsObject],
// [AccountAIRunCfMetaExecuteLlama3_1_8bInstructParamsBodyMessagesTool].
type AccountAIRunCfMetaExecuteLlama3_1_8bInstructParamsBodyMessagesToolUnion interface {
	implementsAccountAIRunCfMetaExecuteLlama3_1_8bInstructParamsBodyMessagesToolUnion()
}

type AccountAIRunCfMetaExecuteLlama3_1_8bInstructParamsBodyMessagesToolsObject struct {
	// A brief description of what the tool does.
	Description param.Field[string] `json:"description,required"`
	// The name of the tool. More descriptive the better.
	Name param.Field[string] `json:"name,required"`
	// Schema defining the parameters accepted by the tool.
	Parameters param.Field[AccountAIRunCfMetaExecuteLlama3_1_8bInstructParamsBodyMessagesToolsObjectParameters] `json:"parameters,required"`
}

func (r AccountAIRunCfMetaExecuteLlama3_1_8bInstructParamsBodyMessagesToolsObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfMetaExecuteLlama3_1_8bInstructParamsBodyMessagesToolsObject) implementsAccountAIRunCfMetaExecuteLlama3_1_8bInstructParamsBodyMessagesToolUnion() {
}

// Schema defining the parameters accepted by the tool.
type AccountAIRunCfMetaExecuteLlama3_1_8bInstructParamsBodyMessagesToolsObjectParameters struct {
	// Definitions of each parameter.
	Properties param.Field[map[string]AccountAIRunCfMetaExecuteLlama3_1_8bInstructParamsBodyMessagesToolsObjectParametersProperties] `json:"properties,required"`
	// The type of the parameters object (usually 'object').
	Type param.Field[string] `json:"type,required"`
	// List of required parameter names.
	Required param.Field[[]string] `json:"required"`
}

func (r AccountAIRunCfMetaExecuteLlama3_1_8bInstructParamsBodyMessagesToolsObjectParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfMetaExecuteLlama3_1_8bInstructParamsBodyMessagesToolsObjectParametersProperties struct {
	// A description of the expected parameter.
	Description param.Field[string] `json:"description,required"`
	// The data type of the parameter.
	Type param.Field[string] `json:"type,required"`
}

func (r AccountAIRunCfMetaExecuteLlama3_1_8bInstructParamsBodyMessagesToolsObjectParametersProperties) MarshalJSON() (data []byte, err error) {
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
	// integer values
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
// integer values
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
	// integer values
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
	Content param.Field[string] `json:"content,required"`
	// The role of the message sender (e.g., 'user', 'assistant', 'system', 'tool').
	Role param.Field[string] `json:"role,required"`
}

func (r AccountAIRunCfMetaExecuteLlama3_2_11bVisionInstructParamsBodyMessagesMessage) MarshalJSON() (data []byte, err error) {
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
// integer values
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

type AccountAIRunCfMetaExecuteLlama3_3_70bInstructFp8FastParams struct {
	QueueRequest param.Field[string]                                                 `query:"queueRequest"`
	Body         AccountAIRunCfMetaExecuteLlama3_3_70bInstructFp8FastParamsBodyUnion `json:"body"`
}

func (r AccountAIRunCfMetaExecuteLlama3_3_70bInstructFp8FastParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

// URLQuery serializes
// [AccountAIRunCfMetaExecuteLlama3_3_70bInstructFp8FastParams]'s query parameters
// as `url.Values`.
func (r AccountAIRunCfMetaExecuteLlama3_3_70bInstructFp8FastParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountAIRunCfMetaExecuteLlama3_3_70bInstructFp8FastParamsBody struct {
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

func (r AccountAIRunCfMetaExecuteLlama3_3_70bInstructFp8FastParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfMetaExecuteLlama3_3_70bInstructFp8FastParamsBody) implementsAccountAIRunCfMetaExecuteLlama3_3_70bInstructFp8FastParamsBodyUnion() {
}

// Satisfied by
// [AccountAIRunCfMetaExecuteLlama3_3_70bInstructFp8FastParamsBodyPrompt],
// [AccountAIRunCfMetaExecuteLlama3_3_70bInstructFp8FastParamsBodyMessages],
// [AccountAIRunCfMetaExecuteLlama3_3_70bInstructFp8FastParamsBody].
type AccountAIRunCfMetaExecuteLlama3_3_70bInstructFp8FastParamsBodyUnion interface {
	implementsAccountAIRunCfMetaExecuteLlama3_3_70bInstructFp8FastParamsBodyUnion()
}

type AccountAIRunCfMetaExecuteLlama3_3_70bInstructFp8FastParamsBodyPrompt struct {
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
	ResponseFormat    param.Field[AccountAIRunCfMetaExecuteLlama3_3_70bInstructFp8FastParamsBodyPromptResponseFormat] `json:"response_format"`
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

func (r AccountAIRunCfMetaExecuteLlama3_3_70bInstructFp8FastParamsBodyPrompt) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfMetaExecuteLlama3_3_70bInstructFp8FastParamsBodyPrompt) implementsAccountAIRunCfMetaExecuteLlama3_3_70bInstructFp8FastParamsBodyUnion() {
}

type AccountAIRunCfMetaExecuteLlama3_3_70bInstructFp8FastParamsBodyPromptResponseFormat struct {
	JsonSchema param.Field[interface{}]                                                                            `json:"json_schema"`
	Type       param.Field[AccountAIRunCfMetaExecuteLlama3_3_70bInstructFp8FastParamsBodyPromptResponseFormatType] `json:"type"`
}

func (r AccountAIRunCfMetaExecuteLlama3_3_70bInstructFp8FastParamsBodyPromptResponseFormat) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfMetaExecuteLlama3_3_70bInstructFp8FastParamsBodyPromptResponseFormatType string

const (
	AccountAIRunCfMetaExecuteLlama3_3_70bInstructFp8FastParamsBodyPromptResponseFormatTypeJsonObject AccountAIRunCfMetaExecuteLlama3_3_70bInstructFp8FastParamsBodyPromptResponseFormatType = "json_object"
	AccountAIRunCfMetaExecuteLlama3_3_70bInstructFp8FastParamsBodyPromptResponseFormatTypeJsonSchema AccountAIRunCfMetaExecuteLlama3_3_70bInstructFp8FastParamsBodyPromptResponseFormatType = "json_schema"
)

func (r AccountAIRunCfMetaExecuteLlama3_3_70bInstructFp8FastParamsBodyPromptResponseFormatType) IsKnown() bool {
	switch r {
	case AccountAIRunCfMetaExecuteLlama3_3_70bInstructFp8FastParamsBodyPromptResponseFormatTypeJsonObject, AccountAIRunCfMetaExecuteLlama3_3_70bInstructFp8FastParamsBodyPromptResponseFormatTypeJsonSchema:
		return true
	}
	return false
}

type AccountAIRunCfMetaExecuteLlama3_3_70bInstructFp8FastParamsBodyMessages struct {
	// An array of message objects representing the conversation history.
	Messages param.Field[[]AccountAIRunCfMetaExecuteLlama3_3_70bInstructFp8FastParamsBodyMessagesMessage] `json:"messages,required"`
	// Decreases the likelihood of the model repeating the same lines verbatim.
	FrequencyPenalty param.Field[float64]                                                                          `json:"frequency_penalty"`
	Functions        param.Field[[]AccountAIRunCfMetaExecuteLlama3_3_70bInstructFp8FastParamsBodyMessagesFunction] `json:"functions"`
	// The maximum number of tokens to generate in the response.
	MaxTokens param.Field[int64] `json:"max_tokens"`
	// Increases the likelihood of the model introducing new topics.
	PresencePenalty param.Field[float64] `json:"presence_penalty"`
	// If true, a chat template is not applied and you must adhere to the specific
	// model's expected formatting.
	Raw param.Field[bool] `json:"raw"`
	// Penalty for repeated tokens; higher values discourage repetition.
	RepetitionPenalty param.Field[float64]                                                                              `json:"repetition_penalty"`
	ResponseFormat    param.Field[AccountAIRunCfMetaExecuteLlama3_3_70bInstructFp8FastParamsBodyMessagesResponseFormat] `json:"response_format"`
	// Random seed for reproducibility of the generation.
	Seed param.Field[int64] `json:"seed"`
	// If true, the response will be streamed back incrementally using SSE, Server Sent
	// Events.
	Stream param.Field[bool] `json:"stream"`
	// Controls the randomness of the output; higher values produce more random
	// results.
	Temperature param.Field[float64] `json:"temperature"`
	// A list of tools available for the assistant to use.
	Tools param.Field[[]AccountAIRunCfMetaExecuteLlama3_3_70bInstructFp8FastParamsBodyMessagesToolUnion] `json:"tools"`
	// Limits the AI to choose from the top 'k' most probable words. Lower values make
	// responses more focused; higher values introduce more variety and potential
	// surprises.
	TopK param.Field[int64] `json:"top_k"`
	// Adjusts the creativity of the AI's responses by controlling how many possible
	// words it considers. Lower values make outputs more predictable; higher values
	// allow for more varied and creative responses.
	TopP param.Field[float64] `json:"top_p"`
}

func (r AccountAIRunCfMetaExecuteLlama3_3_70bInstructFp8FastParamsBodyMessages) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfMetaExecuteLlama3_3_70bInstructFp8FastParamsBodyMessages) implementsAccountAIRunCfMetaExecuteLlama3_3_70bInstructFp8FastParamsBodyUnion() {
}

type AccountAIRunCfMetaExecuteLlama3_3_70bInstructFp8FastParamsBodyMessagesMessage struct {
	// The content of the message as a string.
	Content param.Field[string] `json:"content,required"`
	// The role of the message sender (e.g., 'user', 'assistant', 'system', 'tool').
	Role param.Field[string] `json:"role,required"`
}

func (r AccountAIRunCfMetaExecuteLlama3_3_70bInstructFp8FastParamsBodyMessagesMessage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfMetaExecuteLlama3_3_70bInstructFp8FastParamsBodyMessagesFunction struct {
	Code param.Field[string] `json:"code,required"`
	Name param.Field[string] `json:"name,required"`
}

func (r AccountAIRunCfMetaExecuteLlama3_3_70bInstructFp8FastParamsBodyMessagesFunction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfMetaExecuteLlama3_3_70bInstructFp8FastParamsBodyMessagesResponseFormat struct {
	JsonSchema param.Field[interface{}]                                                                              `json:"json_schema"`
	Type       param.Field[AccountAIRunCfMetaExecuteLlama3_3_70bInstructFp8FastParamsBodyMessagesResponseFormatType] `json:"type"`
}

func (r AccountAIRunCfMetaExecuteLlama3_3_70bInstructFp8FastParamsBodyMessagesResponseFormat) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfMetaExecuteLlama3_3_70bInstructFp8FastParamsBodyMessagesResponseFormatType string

const (
	AccountAIRunCfMetaExecuteLlama3_3_70bInstructFp8FastParamsBodyMessagesResponseFormatTypeJsonObject AccountAIRunCfMetaExecuteLlama3_3_70bInstructFp8FastParamsBodyMessagesResponseFormatType = "json_object"
	AccountAIRunCfMetaExecuteLlama3_3_70bInstructFp8FastParamsBodyMessagesResponseFormatTypeJsonSchema AccountAIRunCfMetaExecuteLlama3_3_70bInstructFp8FastParamsBodyMessagesResponseFormatType = "json_schema"
)

func (r AccountAIRunCfMetaExecuteLlama3_3_70bInstructFp8FastParamsBodyMessagesResponseFormatType) IsKnown() bool {
	switch r {
	case AccountAIRunCfMetaExecuteLlama3_3_70bInstructFp8FastParamsBodyMessagesResponseFormatTypeJsonObject, AccountAIRunCfMetaExecuteLlama3_3_70bInstructFp8FastParamsBodyMessagesResponseFormatTypeJsonSchema:
		return true
	}
	return false
}

type AccountAIRunCfMetaExecuteLlama3_3_70bInstructFp8FastParamsBodyMessagesTool struct {
	// A brief description of what the tool does.
	Description param.Field[string]      `json:"description"`
	Function    param.Field[interface{}] `json:"function"`
	// The name of the tool. More descriptive the better.
	Name       param.Field[string]      `json:"name"`
	Parameters param.Field[interface{}] `json:"parameters"`
	// Specifies the type of tool (e.g., 'function').
	Type param.Field[string] `json:"type"`
}

func (r AccountAIRunCfMetaExecuteLlama3_3_70bInstructFp8FastParamsBodyMessagesTool) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfMetaExecuteLlama3_3_70bInstructFp8FastParamsBodyMessagesTool) implementsAccountAIRunCfMetaExecuteLlama3_3_70bInstructFp8FastParamsBodyMessagesToolUnion() {
}

// Satisfied by
// [AccountAIRunCfMetaExecuteLlama3_3_70bInstructFp8FastParamsBodyMessagesToolsObject],
// [AccountAIRunCfMetaExecuteLlama3_3_70bInstructFp8FastParamsBodyMessagesToolsObject],
// [AccountAIRunCfMetaExecuteLlama3_3_70bInstructFp8FastParamsBodyMessagesTool].
type AccountAIRunCfMetaExecuteLlama3_3_70bInstructFp8FastParamsBodyMessagesToolUnion interface {
	implementsAccountAIRunCfMetaExecuteLlama3_3_70bInstructFp8FastParamsBodyMessagesToolUnion()
}

type AccountAIRunCfMetaExecuteLlama3_3_70bInstructFp8FastParamsBodyMessagesToolsObject struct {
	// A brief description of what the tool does.
	Description param.Field[string] `json:"description,required"`
	// The name of the tool. More descriptive the better.
	Name param.Field[string] `json:"name,required"`
	// Schema defining the parameters accepted by the tool.
	Parameters param.Field[AccountAIRunCfMetaExecuteLlama3_3_70bInstructFp8FastParamsBodyMessagesToolsObjectParameters] `json:"parameters,required"`
}

func (r AccountAIRunCfMetaExecuteLlama3_3_70bInstructFp8FastParamsBodyMessagesToolsObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfMetaExecuteLlama3_3_70bInstructFp8FastParamsBodyMessagesToolsObject) implementsAccountAIRunCfMetaExecuteLlama3_3_70bInstructFp8FastParamsBodyMessagesToolUnion() {
}

// Schema defining the parameters accepted by the tool.
type AccountAIRunCfMetaExecuteLlama3_3_70bInstructFp8FastParamsBodyMessagesToolsObjectParameters struct {
	// Definitions of each parameter.
	Properties param.Field[map[string]AccountAIRunCfMetaExecuteLlama3_3_70bInstructFp8FastParamsBodyMessagesToolsObjectParametersProperties] `json:"properties,required"`
	// The type of the parameters object (usually 'object').
	Type param.Field[string] `json:"type,required"`
	// List of required parameter names.
	Required param.Field[[]string] `json:"required"`
}

func (r AccountAIRunCfMetaExecuteLlama3_3_70bInstructFp8FastParamsBodyMessagesToolsObjectParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfMetaExecuteLlama3_3_70bInstructFp8FastParamsBodyMessagesToolsObjectParametersProperties struct {
	// A description of the expected parameter.
	Description param.Field[string] `json:"description,required"`
	// The data type of the parameter.
	Type param.Field[string] `json:"type,required"`
}

func (r AccountAIRunCfMetaExecuteLlama3_3_70bInstructFp8FastParamsBodyMessagesToolsObjectParametersProperties) MarshalJSON() (data []byte, err error) {
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
	// The language code to translate the text into (e.g., 'es' for Spanish)
	TargetLang param.Field[string] `json:"target_lang,required"`
	// The text to be translated
	Text         param.Field[string] `json:"text,required"`
	QueueRequest param.Field[string] `query:"queueRequest"`
	// The language code of the source text (e.g., 'en' for English). Defaults to 'en'
	// if not specified
	SourceLang param.Field[string] `json:"source_lang"`
}

func (r AccountAIRunCfMetaExecuteM2m100_1_2bParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [AccountAIRunCfMetaExecuteM2m100_1_2bParams]'s query
// parameters as `url.Values`.
func (r AccountAIRunCfMetaExecuteM2m100_1_2bParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
