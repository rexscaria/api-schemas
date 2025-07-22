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

// AccountAIRunCfQwenService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountAIRunCfQwenService] method instead.
type AccountAIRunCfQwenService struct {
	Options []option.RequestOption
}

// NewAccountAIRunCfQwenService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountAIRunCfQwenService(opts ...option.RequestOption) (r *AccountAIRunCfQwenService) {
	r = &AccountAIRunCfQwenService{}
	r.Options = opts
	return
}

// Execute @cf/qwen/qwen1.5-0.5b-chat model.
func (r *AccountAIRunCfQwenService) ExecuteQwen1_5_0_5bChat(ctx context.Context, accountID string, params AccountAIRunCfQwenExecuteQwen1_5_0_5bChatParams, opts ...option.RequestOption) (res *AccountAIRunCfQwenExecuteQwen1_5_0_5bChatResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai/run/@cf/qwen/qwen1.5-0.5b-chat", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Execute @cf/qwen/qwen1.5-1.8b-chat model.
func (r *AccountAIRunCfQwenService) ExecuteQwen1_5_1_8bChat(ctx context.Context, accountID string, params AccountAIRunCfQwenExecuteQwen1_5_1_8bChatParams, opts ...option.RequestOption) (res *AccountAIRunCfQwenExecuteQwen1_5_1_8bChatResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai/run/@cf/qwen/qwen1.5-1.8b-chat", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Execute @cf/qwen/qwen1.5-14b-chat-awq model.
func (r *AccountAIRunCfQwenService) ExecuteQwen1_5_14bChatAwq(ctx context.Context, accountID string, params AccountAIRunCfQwenExecuteQwen1_5_14bChatAwqParams, opts ...option.RequestOption) (res *AccountAIRunCfQwenExecuteQwen1_5_14bChatAwqResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai/run/@cf/qwen/qwen1.5-14b-chat-awq", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Execute @cf/qwen/qwen1.5-7b-chat-awq model.
func (r *AccountAIRunCfQwenService) ExecuteQwen1_5_7bChatAwq(ctx context.Context, accountID string, params AccountAIRunCfQwenExecuteQwen1_5_7bChatAwqParams, opts ...option.RequestOption) (res *AccountAIRunCfQwenExecuteQwen1_5_7bChatAwqResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai/run/@cf/qwen/qwen1.5-7b-chat-awq", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type AccountAIRunCfQwenExecuteQwen1_5_0_5bChatResponse struct {
	Result  AccountAIRunCfQwenExecuteQwen1_5_0_5bChatResponseResultUnion `json:"result" format:"binary"`
	Success bool                                                         `json:"success"`
	JSON    accountAIRunCfQwenExecuteQwen1_5_0_5bChatResponseJSON        `json:"-"`
}

// accountAIRunCfQwenExecuteQwen1_5_0_5bChatResponseJSON contains the JSON metadata
// for the struct [AccountAIRunCfQwenExecuteQwen1_5_0_5bChatResponse]
type accountAIRunCfQwenExecuteQwen1_5_0_5bChatResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIRunCfQwenExecuteQwen1_5_0_5bChatResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunCfQwenExecuteQwen1_5_0_5bChatResponseJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by
// [AccountAIRunCfQwenExecuteQwen1_5_0_5bChatResponseResultObject] or
// [shared.UnionString].
type AccountAIRunCfQwenExecuteQwen1_5_0_5bChatResponseResultUnion interface {
	ImplementsAccountAIRunCfQwenExecuteQwen1_5_0_5bChatResponseResultUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccountAIRunCfQwenExecuteQwen1_5_0_5bChatResponseResultUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccountAIRunCfQwenExecuteQwen1_5_0_5bChatResponseResultObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type AccountAIRunCfQwenExecuteQwen1_5_0_5bChatResponseResultObject struct {
	// The generated text response from the model
	Response string `json:"response,required"`
	// An array of tool calls requests made during the response generation
	ToolCalls []AccountAIRunCfQwenExecuteQwen1_5_0_5bChatResponseResultObjectToolCall `json:"tool_calls"`
	// Usage statistics for the inference request
	Usage AccountAIRunCfQwenExecuteQwen1_5_0_5bChatResponseResultObjectUsage `json:"usage"`
	JSON  accountAIRunCfQwenExecuteQwen1_5_0_5bChatResponseResultObjectJSON  `json:"-"`
}

// accountAIRunCfQwenExecuteQwen1_5_0_5bChatResponseResultObjectJSON contains the
// JSON metadata for the struct
// [AccountAIRunCfQwenExecuteQwen1_5_0_5bChatResponseResultObject]
type accountAIRunCfQwenExecuteQwen1_5_0_5bChatResponseResultObjectJSON struct {
	Response    apijson.Field
	ToolCalls   apijson.Field
	Usage       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIRunCfQwenExecuteQwen1_5_0_5bChatResponseResultObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunCfQwenExecuteQwen1_5_0_5bChatResponseResultObjectJSON) RawJSON() string {
	return r.raw
}

func (r AccountAIRunCfQwenExecuteQwen1_5_0_5bChatResponseResultObject) ImplementsAccountAIRunCfQwenExecuteQwen1_5_0_5bChatResponseResultUnion() {
}

type AccountAIRunCfQwenExecuteQwen1_5_0_5bChatResponseResultObjectToolCall struct {
	// The arguments passed to be passed to the tool call request
	Arguments interface{} `json:"arguments"`
	// The name of the tool to be called
	Name string                                                                    `json:"name"`
	JSON accountAIRunCfQwenExecuteQwen1_5_0_5bChatResponseResultObjectToolCallJSON `json:"-"`
}

// accountAIRunCfQwenExecuteQwen1_5_0_5bChatResponseResultObjectToolCallJSON
// contains the JSON metadata for the struct
// [AccountAIRunCfQwenExecuteQwen1_5_0_5bChatResponseResultObjectToolCall]
type accountAIRunCfQwenExecuteQwen1_5_0_5bChatResponseResultObjectToolCallJSON struct {
	Arguments   apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIRunCfQwenExecuteQwen1_5_0_5bChatResponseResultObjectToolCall) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunCfQwenExecuteQwen1_5_0_5bChatResponseResultObjectToolCallJSON) RawJSON() string {
	return r.raw
}

// Usage statistics for the inference request
type AccountAIRunCfQwenExecuteQwen1_5_0_5bChatResponseResultObjectUsage struct {
	// Total number of tokens in output
	CompletionTokens float64 `json:"completion_tokens"`
	// Total number of tokens in input
	PromptTokens float64 `json:"prompt_tokens"`
	// Total number of input and output tokens
	TotalTokens float64                                                                `json:"total_tokens"`
	JSON        accountAIRunCfQwenExecuteQwen1_5_0_5bChatResponseResultObjectUsageJSON `json:"-"`
}

// accountAIRunCfQwenExecuteQwen1_5_0_5bChatResponseResultObjectUsageJSON contains
// the JSON metadata for the struct
// [AccountAIRunCfQwenExecuteQwen1_5_0_5bChatResponseResultObjectUsage]
type accountAIRunCfQwenExecuteQwen1_5_0_5bChatResponseResultObjectUsageJSON struct {
	CompletionTokens apijson.Field
	PromptTokens     apijson.Field
	TotalTokens      apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountAIRunCfQwenExecuteQwen1_5_0_5bChatResponseResultObjectUsage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunCfQwenExecuteQwen1_5_0_5bChatResponseResultObjectUsageJSON) RawJSON() string {
	return r.raw
}

type AccountAIRunCfQwenExecuteQwen1_5_1_8bChatResponse struct {
	Result  AccountAIRunCfQwenExecuteQwen1_5_1_8bChatResponseResultUnion `json:"result" format:"binary"`
	Success bool                                                         `json:"success"`
	JSON    accountAIRunCfQwenExecuteQwen1_5_1_8bChatResponseJSON        `json:"-"`
}

// accountAIRunCfQwenExecuteQwen1_5_1_8bChatResponseJSON contains the JSON metadata
// for the struct [AccountAIRunCfQwenExecuteQwen1_5_1_8bChatResponse]
type accountAIRunCfQwenExecuteQwen1_5_1_8bChatResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIRunCfQwenExecuteQwen1_5_1_8bChatResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunCfQwenExecuteQwen1_5_1_8bChatResponseJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by
// [AccountAIRunCfQwenExecuteQwen1_5_1_8bChatResponseResultObject] or
// [shared.UnionString].
type AccountAIRunCfQwenExecuteQwen1_5_1_8bChatResponseResultUnion interface {
	ImplementsAccountAIRunCfQwenExecuteQwen1_5_1_8bChatResponseResultUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccountAIRunCfQwenExecuteQwen1_5_1_8bChatResponseResultUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccountAIRunCfQwenExecuteQwen1_5_1_8bChatResponseResultObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type AccountAIRunCfQwenExecuteQwen1_5_1_8bChatResponseResultObject struct {
	// The generated text response from the model
	Response string `json:"response,required"`
	// An array of tool calls requests made during the response generation
	ToolCalls []AccountAIRunCfQwenExecuteQwen1_5_1_8bChatResponseResultObjectToolCall `json:"tool_calls"`
	// Usage statistics for the inference request
	Usage AccountAIRunCfQwenExecuteQwen1_5_1_8bChatResponseResultObjectUsage `json:"usage"`
	JSON  accountAIRunCfQwenExecuteQwen1_5_1_8bChatResponseResultObjectJSON  `json:"-"`
}

// accountAIRunCfQwenExecuteQwen1_5_1_8bChatResponseResultObjectJSON contains the
// JSON metadata for the struct
// [AccountAIRunCfQwenExecuteQwen1_5_1_8bChatResponseResultObject]
type accountAIRunCfQwenExecuteQwen1_5_1_8bChatResponseResultObjectJSON struct {
	Response    apijson.Field
	ToolCalls   apijson.Field
	Usage       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIRunCfQwenExecuteQwen1_5_1_8bChatResponseResultObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunCfQwenExecuteQwen1_5_1_8bChatResponseResultObjectJSON) RawJSON() string {
	return r.raw
}

func (r AccountAIRunCfQwenExecuteQwen1_5_1_8bChatResponseResultObject) ImplementsAccountAIRunCfQwenExecuteQwen1_5_1_8bChatResponseResultUnion() {
}

type AccountAIRunCfQwenExecuteQwen1_5_1_8bChatResponseResultObjectToolCall struct {
	// The arguments passed to be passed to the tool call request
	Arguments interface{} `json:"arguments"`
	// The name of the tool to be called
	Name string                                                                    `json:"name"`
	JSON accountAIRunCfQwenExecuteQwen1_5_1_8bChatResponseResultObjectToolCallJSON `json:"-"`
}

// accountAIRunCfQwenExecuteQwen1_5_1_8bChatResponseResultObjectToolCallJSON
// contains the JSON metadata for the struct
// [AccountAIRunCfQwenExecuteQwen1_5_1_8bChatResponseResultObjectToolCall]
type accountAIRunCfQwenExecuteQwen1_5_1_8bChatResponseResultObjectToolCallJSON struct {
	Arguments   apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIRunCfQwenExecuteQwen1_5_1_8bChatResponseResultObjectToolCall) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunCfQwenExecuteQwen1_5_1_8bChatResponseResultObjectToolCallJSON) RawJSON() string {
	return r.raw
}

// Usage statistics for the inference request
type AccountAIRunCfQwenExecuteQwen1_5_1_8bChatResponseResultObjectUsage struct {
	// Total number of tokens in output
	CompletionTokens float64 `json:"completion_tokens"`
	// Total number of tokens in input
	PromptTokens float64 `json:"prompt_tokens"`
	// Total number of input and output tokens
	TotalTokens float64                                                                `json:"total_tokens"`
	JSON        accountAIRunCfQwenExecuteQwen1_5_1_8bChatResponseResultObjectUsageJSON `json:"-"`
}

// accountAIRunCfQwenExecuteQwen1_5_1_8bChatResponseResultObjectUsageJSON contains
// the JSON metadata for the struct
// [AccountAIRunCfQwenExecuteQwen1_5_1_8bChatResponseResultObjectUsage]
type accountAIRunCfQwenExecuteQwen1_5_1_8bChatResponseResultObjectUsageJSON struct {
	CompletionTokens apijson.Field
	PromptTokens     apijson.Field
	TotalTokens      apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountAIRunCfQwenExecuteQwen1_5_1_8bChatResponseResultObjectUsage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunCfQwenExecuteQwen1_5_1_8bChatResponseResultObjectUsageJSON) RawJSON() string {
	return r.raw
}

type AccountAIRunCfQwenExecuteQwen1_5_14bChatAwqResponse struct {
	Result  AccountAIRunCfQwenExecuteQwen1_5_14bChatAwqResponseResultUnion `json:"result" format:"binary"`
	Success bool                                                           `json:"success"`
	JSON    accountAIRunCfQwenExecuteQwen1_5_14bChatAwqResponseJSON        `json:"-"`
}

// accountAIRunCfQwenExecuteQwen1_5_14bChatAwqResponseJSON contains the JSON
// metadata for the struct [AccountAIRunCfQwenExecuteQwen1_5_14bChatAwqResponse]
type accountAIRunCfQwenExecuteQwen1_5_14bChatAwqResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIRunCfQwenExecuteQwen1_5_14bChatAwqResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunCfQwenExecuteQwen1_5_14bChatAwqResponseJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by
// [AccountAIRunCfQwenExecuteQwen1_5_14bChatAwqResponseResultObject] or
// [shared.UnionString].
type AccountAIRunCfQwenExecuteQwen1_5_14bChatAwqResponseResultUnion interface {
	ImplementsAccountAIRunCfQwenExecuteQwen1_5_14bChatAwqResponseResultUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccountAIRunCfQwenExecuteQwen1_5_14bChatAwqResponseResultUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccountAIRunCfQwenExecuteQwen1_5_14bChatAwqResponseResultObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type AccountAIRunCfQwenExecuteQwen1_5_14bChatAwqResponseResultObject struct {
	// The generated text response from the model
	Response string `json:"response,required"`
	// An array of tool calls requests made during the response generation
	ToolCalls []AccountAIRunCfQwenExecuteQwen1_5_14bChatAwqResponseResultObjectToolCall `json:"tool_calls"`
	// Usage statistics for the inference request
	Usage AccountAIRunCfQwenExecuteQwen1_5_14bChatAwqResponseResultObjectUsage `json:"usage"`
	JSON  accountAIRunCfQwenExecuteQwen1_5_14bChatAwqResponseResultObjectJSON  `json:"-"`
}

// accountAIRunCfQwenExecuteQwen1_5_14bChatAwqResponseResultObjectJSON contains the
// JSON metadata for the struct
// [AccountAIRunCfQwenExecuteQwen1_5_14bChatAwqResponseResultObject]
type accountAIRunCfQwenExecuteQwen1_5_14bChatAwqResponseResultObjectJSON struct {
	Response    apijson.Field
	ToolCalls   apijson.Field
	Usage       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIRunCfQwenExecuteQwen1_5_14bChatAwqResponseResultObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunCfQwenExecuteQwen1_5_14bChatAwqResponseResultObjectJSON) RawJSON() string {
	return r.raw
}

func (r AccountAIRunCfQwenExecuteQwen1_5_14bChatAwqResponseResultObject) ImplementsAccountAIRunCfQwenExecuteQwen1_5_14bChatAwqResponseResultUnion() {
}

type AccountAIRunCfQwenExecuteQwen1_5_14bChatAwqResponseResultObjectToolCall struct {
	// The arguments passed to be passed to the tool call request
	Arguments interface{} `json:"arguments"`
	// The name of the tool to be called
	Name string                                                                      `json:"name"`
	JSON accountAIRunCfQwenExecuteQwen1_5_14bChatAwqResponseResultObjectToolCallJSON `json:"-"`
}

// accountAIRunCfQwenExecuteQwen1_5_14bChatAwqResponseResultObjectToolCallJSON
// contains the JSON metadata for the struct
// [AccountAIRunCfQwenExecuteQwen1_5_14bChatAwqResponseResultObjectToolCall]
type accountAIRunCfQwenExecuteQwen1_5_14bChatAwqResponseResultObjectToolCallJSON struct {
	Arguments   apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIRunCfQwenExecuteQwen1_5_14bChatAwqResponseResultObjectToolCall) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunCfQwenExecuteQwen1_5_14bChatAwqResponseResultObjectToolCallJSON) RawJSON() string {
	return r.raw
}

// Usage statistics for the inference request
type AccountAIRunCfQwenExecuteQwen1_5_14bChatAwqResponseResultObjectUsage struct {
	// Total number of tokens in output
	CompletionTokens float64 `json:"completion_tokens"`
	// Total number of tokens in input
	PromptTokens float64 `json:"prompt_tokens"`
	// Total number of input and output tokens
	TotalTokens float64                                                                  `json:"total_tokens"`
	JSON        accountAIRunCfQwenExecuteQwen1_5_14bChatAwqResponseResultObjectUsageJSON `json:"-"`
}

// accountAIRunCfQwenExecuteQwen1_5_14bChatAwqResponseResultObjectUsageJSON
// contains the JSON metadata for the struct
// [AccountAIRunCfQwenExecuteQwen1_5_14bChatAwqResponseResultObjectUsage]
type accountAIRunCfQwenExecuteQwen1_5_14bChatAwqResponseResultObjectUsageJSON struct {
	CompletionTokens apijson.Field
	PromptTokens     apijson.Field
	TotalTokens      apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountAIRunCfQwenExecuteQwen1_5_14bChatAwqResponseResultObjectUsage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunCfQwenExecuteQwen1_5_14bChatAwqResponseResultObjectUsageJSON) RawJSON() string {
	return r.raw
}

type AccountAIRunCfQwenExecuteQwen1_5_7bChatAwqResponse struct {
	Result  AccountAIRunCfQwenExecuteQwen1_5_7bChatAwqResponseResultUnion `json:"result" format:"binary"`
	Success bool                                                          `json:"success"`
	JSON    accountAIRunCfQwenExecuteQwen1_5_7bChatAwqResponseJSON        `json:"-"`
}

// accountAIRunCfQwenExecuteQwen1_5_7bChatAwqResponseJSON contains the JSON
// metadata for the struct [AccountAIRunCfQwenExecuteQwen1_5_7bChatAwqResponse]
type accountAIRunCfQwenExecuteQwen1_5_7bChatAwqResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIRunCfQwenExecuteQwen1_5_7bChatAwqResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunCfQwenExecuteQwen1_5_7bChatAwqResponseJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by
// [AccountAIRunCfQwenExecuteQwen1_5_7bChatAwqResponseResultObject] or
// [shared.UnionString].
type AccountAIRunCfQwenExecuteQwen1_5_7bChatAwqResponseResultUnion interface {
	ImplementsAccountAIRunCfQwenExecuteQwen1_5_7bChatAwqResponseResultUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccountAIRunCfQwenExecuteQwen1_5_7bChatAwqResponseResultUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccountAIRunCfQwenExecuteQwen1_5_7bChatAwqResponseResultObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type AccountAIRunCfQwenExecuteQwen1_5_7bChatAwqResponseResultObject struct {
	// The generated text response from the model
	Response string `json:"response,required"`
	// An array of tool calls requests made during the response generation
	ToolCalls []AccountAIRunCfQwenExecuteQwen1_5_7bChatAwqResponseResultObjectToolCall `json:"tool_calls"`
	// Usage statistics for the inference request
	Usage AccountAIRunCfQwenExecuteQwen1_5_7bChatAwqResponseResultObjectUsage `json:"usage"`
	JSON  accountAIRunCfQwenExecuteQwen1_5_7bChatAwqResponseResultObjectJSON  `json:"-"`
}

// accountAIRunCfQwenExecuteQwen1_5_7bChatAwqResponseResultObjectJSON contains the
// JSON metadata for the struct
// [AccountAIRunCfQwenExecuteQwen1_5_7bChatAwqResponseResultObject]
type accountAIRunCfQwenExecuteQwen1_5_7bChatAwqResponseResultObjectJSON struct {
	Response    apijson.Field
	ToolCalls   apijson.Field
	Usage       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIRunCfQwenExecuteQwen1_5_7bChatAwqResponseResultObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunCfQwenExecuteQwen1_5_7bChatAwqResponseResultObjectJSON) RawJSON() string {
	return r.raw
}

func (r AccountAIRunCfQwenExecuteQwen1_5_7bChatAwqResponseResultObject) ImplementsAccountAIRunCfQwenExecuteQwen1_5_7bChatAwqResponseResultUnion() {
}

type AccountAIRunCfQwenExecuteQwen1_5_7bChatAwqResponseResultObjectToolCall struct {
	// The arguments passed to be passed to the tool call request
	Arguments interface{} `json:"arguments"`
	// The name of the tool to be called
	Name string                                                                     `json:"name"`
	JSON accountAIRunCfQwenExecuteQwen1_5_7bChatAwqResponseResultObjectToolCallJSON `json:"-"`
}

// accountAIRunCfQwenExecuteQwen1_5_7bChatAwqResponseResultObjectToolCallJSON
// contains the JSON metadata for the struct
// [AccountAIRunCfQwenExecuteQwen1_5_7bChatAwqResponseResultObjectToolCall]
type accountAIRunCfQwenExecuteQwen1_5_7bChatAwqResponseResultObjectToolCallJSON struct {
	Arguments   apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIRunCfQwenExecuteQwen1_5_7bChatAwqResponseResultObjectToolCall) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunCfQwenExecuteQwen1_5_7bChatAwqResponseResultObjectToolCallJSON) RawJSON() string {
	return r.raw
}

// Usage statistics for the inference request
type AccountAIRunCfQwenExecuteQwen1_5_7bChatAwqResponseResultObjectUsage struct {
	// Total number of tokens in output
	CompletionTokens float64 `json:"completion_tokens"`
	// Total number of tokens in input
	PromptTokens float64 `json:"prompt_tokens"`
	// Total number of input and output tokens
	TotalTokens float64                                                                 `json:"total_tokens"`
	JSON        accountAIRunCfQwenExecuteQwen1_5_7bChatAwqResponseResultObjectUsageJSON `json:"-"`
}

// accountAIRunCfQwenExecuteQwen1_5_7bChatAwqResponseResultObjectUsageJSON contains
// the JSON metadata for the struct
// [AccountAIRunCfQwenExecuteQwen1_5_7bChatAwqResponseResultObjectUsage]
type accountAIRunCfQwenExecuteQwen1_5_7bChatAwqResponseResultObjectUsageJSON struct {
	CompletionTokens apijson.Field
	PromptTokens     apijson.Field
	TotalTokens      apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountAIRunCfQwenExecuteQwen1_5_7bChatAwqResponseResultObjectUsage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunCfQwenExecuteQwen1_5_7bChatAwqResponseResultObjectUsageJSON) RawJSON() string {
	return r.raw
}

type AccountAIRunCfQwenExecuteQwen1_5_0_5bChatParams struct {
	QueueRequest param.Field[string]                                      `query:"queueRequest"`
	Body         AccountAIRunCfQwenExecuteQwen1_5_0_5bChatParamsBodyUnion `json:"body"`
}

func (r AccountAIRunCfQwenExecuteQwen1_5_0_5bChatParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

// URLQuery serializes [AccountAIRunCfQwenExecuteQwen1_5_0_5bChatParams]'s query
// parameters as `url.Values`.
func (r AccountAIRunCfQwenExecuteQwen1_5_0_5bChatParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountAIRunCfQwenExecuteQwen1_5_0_5bChatParamsBody struct {
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

func (r AccountAIRunCfQwenExecuteQwen1_5_0_5bChatParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfQwenExecuteQwen1_5_0_5bChatParamsBody) implementsAccountAIRunCfQwenExecuteQwen1_5_0_5bChatParamsBodyUnion() {
}

// Satisfied by [AccountAIRunCfQwenExecuteQwen1_5_0_5bChatParamsBodyPrompt],
// [AccountAIRunCfQwenExecuteQwen1_5_0_5bChatParamsBodyMessages],
// [AccountAIRunCfQwenExecuteQwen1_5_0_5bChatParamsBody].
type AccountAIRunCfQwenExecuteQwen1_5_0_5bChatParamsBodyUnion interface {
	implementsAccountAIRunCfQwenExecuteQwen1_5_0_5bChatParamsBodyUnion()
}

type AccountAIRunCfQwenExecuteQwen1_5_0_5bChatParamsBodyPrompt struct {
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
	RepetitionPenalty param.Field[float64]                                                                 `json:"repetition_penalty"`
	ResponseFormat    param.Field[AccountAIRunCfQwenExecuteQwen1_5_0_5bChatParamsBodyPromptResponseFormat] `json:"response_format"`
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

func (r AccountAIRunCfQwenExecuteQwen1_5_0_5bChatParamsBodyPrompt) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfQwenExecuteQwen1_5_0_5bChatParamsBodyPrompt) implementsAccountAIRunCfQwenExecuteQwen1_5_0_5bChatParamsBodyUnion() {
}

type AccountAIRunCfQwenExecuteQwen1_5_0_5bChatParamsBodyPromptResponseFormat struct {
	JsonSchema param.Field[interface{}]                                                                 `json:"json_schema"`
	Type       param.Field[AccountAIRunCfQwenExecuteQwen1_5_0_5bChatParamsBodyPromptResponseFormatType] `json:"type"`
}

func (r AccountAIRunCfQwenExecuteQwen1_5_0_5bChatParamsBodyPromptResponseFormat) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfQwenExecuteQwen1_5_0_5bChatParamsBodyPromptResponseFormatType string

const (
	AccountAIRunCfQwenExecuteQwen1_5_0_5bChatParamsBodyPromptResponseFormatTypeJsonObject AccountAIRunCfQwenExecuteQwen1_5_0_5bChatParamsBodyPromptResponseFormatType = "json_object"
	AccountAIRunCfQwenExecuteQwen1_5_0_5bChatParamsBodyPromptResponseFormatTypeJsonSchema AccountAIRunCfQwenExecuteQwen1_5_0_5bChatParamsBodyPromptResponseFormatType = "json_schema"
)

func (r AccountAIRunCfQwenExecuteQwen1_5_0_5bChatParamsBodyPromptResponseFormatType) IsKnown() bool {
	switch r {
	case AccountAIRunCfQwenExecuteQwen1_5_0_5bChatParamsBodyPromptResponseFormatTypeJsonObject, AccountAIRunCfQwenExecuteQwen1_5_0_5bChatParamsBodyPromptResponseFormatTypeJsonSchema:
		return true
	}
	return false
}

type AccountAIRunCfQwenExecuteQwen1_5_0_5bChatParamsBodyMessages struct {
	// An array of message objects representing the conversation history.
	Messages param.Field[[]AccountAIRunCfQwenExecuteQwen1_5_0_5bChatParamsBodyMessagesMessage] `json:"messages,required"`
	// Decreases the likelihood of the model repeating the same lines verbatim.
	FrequencyPenalty param.Field[float64]                                                               `json:"frequency_penalty"`
	Functions        param.Field[[]AccountAIRunCfQwenExecuteQwen1_5_0_5bChatParamsBodyMessagesFunction] `json:"functions"`
	// The maximum number of tokens to generate in the response.
	MaxTokens param.Field[int64] `json:"max_tokens"`
	// Increases the likelihood of the model introducing new topics.
	PresencePenalty param.Field[float64] `json:"presence_penalty"`
	// If true, a chat template is not applied and you must adhere to the specific
	// model's expected formatting.
	Raw param.Field[bool] `json:"raw"`
	// Penalty for repeated tokens; higher values discourage repetition.
	RepetitionPenalty param.Field[float64]                                                                   `json:"repetition_penalty"`
	ResponseFormat    param.Field[AccountAIRunCfQwenExecuteQwen1_5_0_5bChatParamsBodyMessagesResponseFormat] `json:"response_format"`
	// Random seed for reproducibility of the generation.
	Seed param.Field[int64] `json:"seed"`
	// If true, the response will be streamed back incrementally using SSE, Server Sent
	// Events.
	Stream param.Field[bool] `json:"stream"`
	// Controls the randomness of the output; higher values produce more random
	// results.
	Temperature param.Field[float64] `json:"temperature"`
	// A list of tools available for the assistant to use.
	Tools param.Field[[]AccountAIRunCfQwenExecuteQwen1_5_0_5bChatParamsBodyMessagesToolUnion] `json:"tools"`
	// Limits the AI to choose from the top 'k' most probable words. Lower values make
	// responses more focused; higher values introduce more variety and potential
	// surprises.
	TopK param.Field[int64] `json:"top_k"`
	// Adjusts the creativity of the AI's responses by controlling how many possible
	// words it considers. Lower values make outputs more predictable; higher values
	// allow for more varied and creative responses.
	TopP param.Field[float64] `json:"top_p"`
}

func (r AccountAIRunCfQwenExecuteQwen1_5_0_5bChatParamsBodyMessages) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfQwenExecuteQwen1_5_0_5bChatParamsBodyMessages) implementsAccountAIRunCfQwenExecuteQwen1_5_0_5bChatParamsBodyUnion() {
}

type AccountAIRunCfQwenExecuteQwen1_5_0_5bChatParamsBodyMessagesMessage struct {
	// The content of the message as a string.
	Content param.Field[string] `json:"content,required"`
	// The role of the message sender (e.g., 'user', 'assistant', 'system', 'tool').
	Role param.Field[string] `json:"role,required"`
}

func (r AccountAIRunCfQwenExecuteQwen1_5_0_5bChatParamsBodyMessagesMessage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfQwenExecuteQwen1_5_0_5bChatParamsBodyMessagesFunction struct {
	Code param.Field[string] `json:"code,required"`
	Name param.Field[string] `json:"name,required"`
}

func (r AccountAIRunCfQwenExecuteQwen1_5_0_5bChatParamsBodyMessagesFunction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfQwenExecuteQwen1_5_0_5bChatParamsBodyMessagesResponseFormat struct {
	JsonSchema param.Field[interface{}]                                                                   `json:"json_schema"`
	Type       param.Field[AccountAIRunCfQwenExecuteQwen1_5_0_5bChatParamsBodyMessagesResponseFormatType] `json:"type"`
}

func (r AccountAIRunCfQwenExecuteQwen1_5_0_5bChatParamsBodyMessagesResponseFormat) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfQwenExecuteQwen1_5_0_5bChatParamsBodyMessagesResponseFormatType string

const (
	AccountAIRunCfQwenExecuteQwen1_5_0_5bChatParamsBodyMessagesResponseFormatTypeJsonObject AccountAIRunCfQwenExecuteQwen1_5_0_5bChatParamsBodyMessagesResponseFormatType = "json_object"
	AccountAIRunCfQwenExecuteQwen1_5_0_5bChatParamsBodyMessagesResponseFormatTypeJsonSchema AccountAIRunCfQwenExecuteQwen1_5_0_5bChatParamsBodyMessagesResponseFormatType = "json_schema"
)

func (r AccountAIRunCfQwenExecuteQwen1_5_0_5bChatParamsBodyMessagesResponseFormatType) IsKnown() bool {
	switch r {
	case AccountAIRunCfQwenExecuteQwen1_5_0_5bChatParamsBodyMessagesResponseFormatTypeJsonObject, AccountAIRunCfQwenExecuteQwen1_5_0_5bChatParamsBodyMessagesResponseFormatTypeJsonSchema:
		return true
	}
	return false
}

type AccountAIRunCfQwenExecuteQwen1_5_0_5bChatParamsBodyMessagesTool struct {
	// A brief description of what the tool does.
	Description param.Field[string]      `json:"description"`
	Function    param.Field[interface{}] `json:"function"`
	// The name of the tool. More descriptive the better.
	Name       param.Field[string]      `json:"name"`
	Parameters param.Field[interface{}] `json:"parameters"`
	// Specifies the type of tool (e.g., 'function').
	Type param.Field[string] `json:"type"`
}

func (r AccountAIRunCfQwenExecuteQwen1_5_0_5bChatParamsBodyMessagesTool) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfQwenExecuteQwen1_5_0_5bChatParamsBodyMessagesTool) implementsAccountAIRunCfQwenExecuteQwen1_5_0_5bChatParamsBodyMessagesToolUnion() {
}

// Satisfied by
// [AccountAIRunCfQwenExecuteQwen1_5_0_5bChatParamsBodyMessagesToolsObject],
// [AccountAIRunCfQwenExecuteQwen1_5_0_5bChatParamsBodyMessagesToolsObject],
// [AccountAIRunCfQwenExecuteQwen1_5_0_5bChatParamsBodyMessagesTool].
type AccountAIRunCfQwenExecuteQwen1_5_0_5bChatParamsBodyMessagesToolUnion interface {
	implementsAccountAIRunCfQwenExecuteQwen1_5_0_5bChatParamsBodyMessagesToolUnion()
}

type AccountAIRunCfQwenExecuteQwen1_5_0_5bChatParamsBodyMessagesToolsObject struct {
	// A brief description of what the tool does.
	Description param.Field[string] `json:"description,required"`
	// The name of the tool. More descriptive the better.
	Name param.Field[string] `json:"name,required"`
	// Schema defining the parameters accepted by the tool.
	Parameters param.Field[AccountAIRunCfQwenExecuteQwen1_5_0_5bChatParamsBodyMessagesToolsObjectParameters] `json:"parameters,required"`
}

func (r AccountAIRunCfQwenExecuteQwen1_5_0_5bChatParamsBodyMessagesToolsObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfQwenExecuteQwen1_5_0_5bChatParamsBodyMessagesToolsObject) implementsAccountAIRunCfQwenExecuteQwen1_5_0_5bChatParamsBodyMessagesToolUnion() {
}

// Schema defining the parameters accepted by the tool.
type AccountAIRunCfQwenExecuteQwen1_5_0_5bChatParamsBodyMessagesToolsObjectParameters struct {
	// Definitions of each parameter.
	Properties param.Field[map[string]AccountAIRunCfQwenExecuteQwen1_5_0_5bChatParamsBodyMessagesToolsObjectParametersProperties] `json:"properties,required"`
	// The type of the parameters object (usually 'object').
	Type param.Field[string] `json:"type,required"`
	// List of required parameter names.
	Required param.Field[[]string] `json:"required"`
}

func (r AccountAIRunCfQwenExecuteQwen1_5_0_5bChatParamsBodyMessagesToolsObjectParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfQwenExecuteQwen1_5_0_5bChatParamsBodyMessagesToolsObjectParametersProperties struct {
	// A description of the expected parameter.
	Description param.Field[string] `json:"description,required"`
	// The data type of the parameter.
	Type param.Field[string] `json:"type,required"`
}

func (r AccountAIRunCfQwenExecuteQwen1_5_0_5bChatParamsBodyMessagesToolsObjectParametersProperties) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfQwenExecuteQwen1_5_1_8bChatParams struct {
	QueueRequest param.Field[string]                                      `query:"queueRequest"`
	Body         AccountAIRunCfQwenExecuteQwen1_5_1_8bChatParamsBodyUnion `json:"body"`
}

func (r AccountAIRunCfQwenExecuteQwen1_5_1_8bChatParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

// URLQuery serializes [AccountAIRunCfQwenExecuteQwen1_5_1_8bChatParams]'s query
// parameters as `url.Values`.
func (r AccountAIRunCfQwenExecuteQwen1_5_1_8bChatParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountAIRunCfQwenExecuteQwen1_5_1_8bChatParamsBody struct {
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

func (r AccountAIRunCfQwenExecuteQwen1_5_1_8bChatParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfQwenExecuteQwen1_5_1_8bChatParamsBody) implementsAccountAIRunCfQwenExecuteQwen1_5_1_8bChatParamsBodyUnion() {
}

// Satisfied by [AccountAIRunCfQwenExecuteQwen1_5_1_8bChatParamsBodyPrompt],
// [AccountAIRunCfQwenExecuteQwen1_5_1_8bChatParamsBodyMessages],
// [AccountAIRunCfQwenExecuteQwen1_5_1_8bChatParamsBody].
type AccountAIRunCfQwenExecuteQwen1_5_1_8bChatParamsBodyUnion interface {
	implementsAccountAIRunCfQwenExecuteQwen1_5_1_8bChatParamsBodyUnion()
}

type AccountAIRunCfQwenExecuteQwen1_5_1_8bChatParamsBodyPrompt struct {
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
	RepetitionPenalty param.Field[float64]                                                                 `json:"repetition_penalty"`
	ResponseFormat    param.Field[AccountAIRunCfQwenExecuteQwen1_5_1_8bChatParamsBodyPromptResponseFormat] `json:"response_format"`
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

func (r AccountAIRunCfQwenExecuteQwen1_5_1_8bChatParamsBodyPrompt) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfQwenExecuteQwen1_5_1_8bChatParamsBodyPrompt) implementsAccountAIRunCfQwenExecuteQwen1_5_1_8bChatParamsBodyUnion() {
}

type AccountAIRunCfQwenExecuteQwen1_5_1_8bChatParamsBodyPromptResponseFormat struct {
	JsonSchema param.Field[interface{}]                                                                 `json:"json_schema"`
	Type       param.Field[AccountAIRunCfQwenExecuteQwen1_5_1_8bChatParamsBodyPromptResponseFormatType] `json:"type"`
}

func (r AccountAIRunCfQwenExecuteQwen1_5_1_8bChatParamsBodyPromptResponseFormat) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfQwenExecuteQwen1_5_1_8bChatParamsBodyPromptResponseFormatType string

const (
	AccountAIRunCfQwenExecuteQwen1_5_1_8bChatParamsBodyPromptResponseFormatTypeJsonObject AccountAIRunCfQwenExecuteQwen1_5_1_8bChatParamsBodyPromptResponseFormatType = "json_object"
	AccountAIRunCfQwenExecuteQwen1_5_1_8bChatParamsBodyPromptResponseFormatTypeJsonSchema AccountAIRunCfQwenExecuteQwen1_5_1_8bChatParamsBodyPromptResponseFormatType = "json_schema"
)

func (r AccountAIRunCfQwenExecuteQwen1_5_1_8bChatParamsBodyPromptResponseFormatType) IsKnown() bool {
	switch r {
	case AccountAIRunCfQwenExecuteQwen1_5_1_8bChatParamsBodyPromptResponseFormatTypeJsonObject, AccountAIRunCfQwenExecuteQwen1_5_1_8bChatParamsBodyPromptResponseFormatTypeJsonSchema:
		return true
	}
	return false
}

type AccountAIRunCfQwenExecuteQwen1_5_1_8bChatParamsBodyMessages struct {
	// An array of message objects representing the conversation history.
	Messages param.Field[[]AccountAIRunCfQwenExecuteQwen1_5_1_8bChatParamsBodyMessagesMessage] `json:"messages,required"`
	// Decreases the likelihood of the model repeating the same lines verbatim.
	FrequencyPenalty param.Field[float64]                                                               `json:"frequency_penalty"`
	Functions        param.Field[[]AccountAIRunCfQwenExecuteQwen1_5_1_8bChatParamsBodyMessagesFunction] `json:"functions"`
	// The maximum number of tokens to generate in the response.
	MaxTokens param.Field[int64] `json:"max_tokens"`
	// Increases the likelihood of the model introducing new topics.
	PresencePenalty param.Field[float64] `json:"presence_penalty"`
	// If true, a chat template is not applied and you must adhere to the specific
	// model's expected formatting.
	Raw param.Field[bool] `json:"raw"`
	// Penalty for repeated tokens; higher values discourage repetition.
	RepetitionPenalty param.Field[float64]                                                                   `json:"repetition_penalty"`
	ResponseFormat    param.Field[AccountAIRunCfQwenExecuteQwen1_5_1_8bChatParamsBodyMessagesResponseFormat] `json:"response_format"`
	// Random seed for reproducibility of the generation.
	Seed param.Field[int64] `json:"seed"`
	// If true, the response will be streamed back incrementally using SSE, Server Sent
	// Events.
	Stream param.Field[bool] `json:"stream"`
	// Controls the randomness of the output; higher values produce more random
	// results.
	Temperature param.Field[float64] `json:"temperature"`
	// A list of tools available for the assistant to use.
	Tools param.Field[[]AccountAIRunCfQwenExecuteQwen1_5_1_8bChatParamsBodyMessagesToolUnion] `json:"tools"`
	// Limits the AI to choose from the top 'k' most probable words. Lower values make
	// responses more focused; higher values introduce more variety and potential
	// surprises.
	TopK param.Field[int64] `json:"top_k"`
	// Adjusts the creativity of the AI's responses by controlling how many possible
	// words it considers. Lower values make outputs more predictable; higher values
	// allow for more varied and creative responses.
	TopP param.Field[float64] `json:"top_p"`
}

func (r AccountAIRunCfQwenExecuteQwen1_5_1_8bChatParamsBodyMessages) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfQwenExecuteQwen1_5_1_8bChatParamsBodyMessages) implementsAccountAIRunCfQwenExecuteQwen1_5_1_8bChatParamsBodyUnion() {
}

type AccountAIRunCfQwenExecuteQwen1_5_1_8bChatParamsBodyMessagesMessage struct {
	// The content of the message as a string.
	Content param.Field[string] `json:"content,required"`
	// The role of the message sender (e.g., 'user', 'assistant', 'system', 'tool').
	Role param.Field[string] `json:"role,required"`
}

func (r AccountAIRunCfQwenExecuteQwen1_5_1_8bChatParamsBodyMessagesMessage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfQwenExecuteQwen1_5_1_8bChatParamsBodyMessagesFunction struct {
	Code param.Field[string] `json:"code,required"`
	Name param.Field[string] `json:"name,required"`
}

func (r AccountAIRunCfQwenExecuteQwen1_5_1_8bChatParamsBodyMessagesFunction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfQwenExecuteQwen1_5_1_8bChatParamsBodyMessagesResponseFormat struct {
	JsonSchema param.Field[interface{}]                                                                   `json:"json_schema"`
	Type       param.Field[AccountAIRunCfQwenExecuteQwen1_5_1_8bChatParamsBodyMessagesResponseFormatType] `json:"type"`
}

func (r AccountAIRunCfQwenExecuteQwen1_5_1_8bChatParamsBodyMessagesResponseFormat) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfQwenExecuteQwen1_5_1_8bChatParamsBodyMessagesResponseFormatType string

const (
	AccountAIRunCfQwenExecuteQwen1_5_1_8bChatParamsBodyMessagesResponseFormatTypeJsonObject AccountAIRunCfQwenExecuteQwen1_5_1_8bChatParamsBodyMessagesResponseFormatType = "json_object"
	AccountAIRunCfQwenExecuteQwen1_5_1_8bChatParamsBodyMessagesResponseFormatTypeJsonSchema AccountAIRunCfQwenExecuteQwen1_5_1_8bChatParamsBodyMessagesResponseFormatType = "json_schema"
)

func (r AccountAIRunCfQwenExecuteQwen1_5_1_8bChatParamsBodyMessagesResponseFormatType) IsKnown() bool {
	switch r {
	case AccountAIRunCfQwenExecuteQwen1_5_1_8bChatParamsBodyMessagesResponseFormatTypeJsonObject, AccountAIRunCfQwenExecuteQwen1_5_1_8bChatParamsBodyMessagesResponseFormatTypeJsonSchema:
		return true
	}
	return false
}

type AccountAIRunCfQwenExecuteQwen1_5_1_8bChatParamsBodyMessagesTool struct {
	// A brief description of what the tool does.
	Description param.Field[string]      `json:"description"`
	Function    param.Field[interface{}] `json:"function"`
	// The name of the tool. More descriptive the better.
	Name       param.Field[string]      `json:"name"`
	Parameters param.Field[interface{}] `json:"parameters"`
	// Specifies the type of tool (e.g., 'function').
	Type param.Field[string] `json:"type"`
}

func (r AccountAIRunCfQwenExecuteQwen1_5_1_8bChatParamsBodyMessagesTool) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfQwenExecuteQwen1_5_1_8bChatParamsBodyMessagesTool) implementsAccountAIRunCfQwenExecuteQwen1_5_1_8bChatParamsBodyMessagesToolUnion() {
}

// Satisfied by
// [AccountAIRunCfQwenExecuteQwen1_5_1_8bChatParamsBodyMessagesToolsObject],
// [AccountAIRunCfQwenExecuteQwen1_5_1_8bChatParamsBodyMessagesToolsObject],
// [AccountAIRunCfQwenExecuteQwen1_5_1_8bChatParamsBodyMessagesTool].
type AccountAIRunCfQwenExecuteQwen1_5_1_8bChatParamsBodyMessagesToolUnion interface {
	implementsAccountAIRunCfQwenExecuteQwen1_5_1_8bChatParamsBodyMessagesToolUnion()
}

type AccountAIRunCfQwenExecuteQwen1_5_1_8bChatParamsBodyMessagesToolsObject struct {
	// A brief description of what the tool does.
	Description param.Field[string] `json:"description,required"`
	// The name of the tool. More descriptive the better.
	Name param.Field[string] `json:"name,required"`
	// Schema defining the parameters accepted by the tool.
	Parameters param.Field[AccountAIRunCfQwenExecuteQwen1_5_1_8bChatParamsBodyMessagesToolsObjectParameters] `json:"parameters,required"`
}

func (r AccountAIRunCfQwenExecuteQwen1_5_1_8bChatParamsBodyMessagesToolsObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfQwenExecuteQwen1_5_1_8bChatParamsBodyMessagesToolsObject) implementsAccountAIRunCfQwenExecuteQwen1_5_1_8bChatParamsBodyMessagesToolUnion() {
}

// Schema defining the parameters accepted by the tool.
type AccountAIRunCfQwenExecuteQwen1_5_1_8bChatParamsBodyMessagesToolsObjectParameters struct {
	// Definitions of each parameter.
	Properties param.Field[map[string]AccountAIRunCfQwenExecuteQwen1_5_1_8bChatParamsBodyMessagesToolsObjectParametersProperties] `json:"properties,required"`
	// The type of the parameters object (usually 'object').
	Type param.Field[string] `json:"type,required"`
	// List of required parameter names.
	Required param.Field[[]string] `json:"required"`
}

func (r AccountAIRunCfQwenExecuteQwen1_5_1_8bChatParamsBodyMessagesToolsObjectParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfQwenExecuteQwen1_5_1_8bChatParamsBodyMessagesToolsObjectParametersProperties struct {
	// A description of the expected parameter.
	Description param.Field[string] `json:"description,required"`
	// The data type of the parameter.
	Type param.Field[string] `json:"type,required"`
}

func (r AccountAIRunCfQwenExecuteQwen1_5_1_8bChatParamsBodyMessagesToolsObjectParametersProperties) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfQwenExecuteQwen1_5_14bChatAwqParams struct {
	QueueRequest param.Field[string]                                        `query:"queueRequest"`
	Body         AccountAIRunCfQwenExecuteQwen1_5_14bChatAwqParamsBodyUnion `json:"body"`
}

func (r AccountAIRunCfQwenExecuteQwen1_5_14bChatAwqParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

// URLQuery serializes [AccountAIRunCfQwenExecuteQwen1_5_14bChatAwqParams]'s query
// parameters as `url.Values`.
func (r AccountAIRunCfQwenExecuteQwen1_5_14bChatAwqParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountAIRunCfQwenExecuteQwen1_5_14bChatAwqParamsBody struct {
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

func (r AccountAIRunCfQwenExecuteQwen1_5_14bChatAwqParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfQwenExecuteQwen1_5_14bChatAwqParamsBody) implementsAccountAIRunCfQwenExecuteQwen1_5_14bChatAwqParamsBodyUnion() {
}

// Satisfied by [AccountAIRunCfQwenExecuteQwen1_5_14bChatAwqParamsBodyPrompt],
// [AccountAIRunCfQwenExecuteQwen1_5_14bChatAwqParamsBodyMessages],
// [AccountAIRunCfQwenExecuteQwen1_5_14bChatAwqParamsBody].
type AccountAIRunCfQwenExecuteQwen1_5_14bChatAwqParamsBodyUnion interface {
	implementsAccountAIRunCfQwenExecuteQwen1_5_14bChatAwqParamsBodyUnion()
}

type AccountAIRunCfQwenExecuteQwen1_5_14bChatAwqParamsBodyPrompt struct {
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
	ResponseFormat    param.Field[AccountAIRunCfQwenExecuteQwen1_5_14bChatAwqParamsBodyPromptResponseFormat] `json:"response_format"`
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

func (r AccountAIRunCfQwenExecuteQwen1_5_14bChatAwqParamsBodyPrompt) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfQwenExecuteQwen1_5_14bChatAwqParamsBodyPrompt) implementsAccountAIRunCfQwenExecuteQwen1_5_14bChatAwqParamsBodyUnion() {
}

type AccountAIRunCfQwenExecuteQwen1_5_14bChatAwqParamsBodyPromptResponseFormat struct {
	JsonSchema param.Field[interface{}]                                                                   `json:"json_schema"`
	Type       param.Field[AccountAIRunCfQwenExecuteQwen1_5_14bChatAwqParamsBodyPromptResponseFormatType] `json:"type"`
}

func (r AccountAIRunCfQwenExecuteQwen1_5_14bChatAwqParamsBodyPromptResponseFormat) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfQwenExecuteQwen1_5_14bChatAwqParamsBodyPromptResponseFormatType string

const (
	AccountAIRunCfQwenExecuteQwen1_5_14bChatAwqParamsBodyPromptResponseFormatTypeJsonObject AccountAIRunCfQwenExecuteQwen1_5_14bChatAwqParamsBodyPromptResponseFormatType = "json_object"
	AccountAIRunCfQwenExecuteQwen1_5_14bChatAwqParamsBodyPromptResponseFormatTypeJsonSchema AccountAIRunCfQwenExecuteQwen1_5_14bChatAwqParamsBodyPromptResponseFormatType = "json_schema"
)

func (r AccountAIRunCfQwenExecuteQwen1_5_14bChatAwqParamsBodyPromptResponseFormatType) IsKnown() bool {
	switch r {
	case AccountAIRunCfQwenExecuteQwen1_5_14bChatAwqParamsBodyPromptResponseFormatTypeJsonObject, AccountAIRunCfQwenExecuteQwen1_5_14bChatAwqParamsBodyPromptResponseFormatTypeJsonSchema:
		return true
	}
	return false
}

type AccountAIRunCfQwenExecuteQwen1_5_14bChatAwqParamsBodyMessages struct {
	// An array of message objects representing the conversation history.
	Messages param.Field[[]AccountAIRunCfQwenExecuteQwen1_5_14bChatAwqParamsBodyMessagesMessage] `json:"messages,required"`
	// Decreases the likelihood of the model repeating the same lines verbatim.
	FrequencyPenalty param.Field[float64]                                                                 `json:"frequency_penalty"`
	Functions        param.Field[[]AccountAIRunCfQwenExecuteQwen1_5_14bChatAwqParamsBodyMessagesFunction] `json:"functions"`
	// The maximum number of tokens to generate in the response.
	MaxTokens param.Field[int64] `json:"max_tokens"`
	// Increases the likelihood of the model introducing new topics.
	PresencePenalty param.Field[float64] `json:"presence_penalty"`
	// If true, a chat template is not applied and you must adhere to the specific
	// model's expected formatting.
	Raw param.Field[bool] `json:"raw"`
	// Penalty for repeated tokens; higher values discourage repetition.
	RepetitionPenalty param.Field[float64]                                                                     `json:"repetition_penalty"`
	ResponseFormat    param.Field[AccountAIRunCfQwenExecuteQwen1_5_14bChatAwqParamsBodyMessagesResponseFormat] `json:"response_format"`
	// Random seed for reproducibility of the generation.
	Seed param.Field[int64] `json:"seed"`
	// If true, the response will be streamed back incrementally using SSE, Server Sent
	// Events.
	Stream param.Field[bool] `json:"stream"`
	// Controls the randomness of the output; higher values produce more random
	// results.
	Temperature param.Field[float64] `json:"temperature"`
	// A list of tools available for the assistant to use.
	Tools param.Field[[]AccountAIRunCfQwenExecuteQwen1_5_14bChatAwqParamsBodyMessagesToolUnion] `json:"tools"`
	// Limits the AI to choose from the top 'k' most probable words. Lower values make
	// responses more focused; higher values introduce more variety and potential
	// surprises.
	TopK param.Field[int64] `json:"top_k"`
	// Adjusts the creativity of the AI's responses by controlling how many possible
	// words it considers. Lower values make outputs more predictable; higher values
	// allow for more varied and creative responses.
	TopP param.Field[float64] `json:"top_p"`
}

func (r AccountAIRunCfQwenExecuteQwen1_5_14bChatAwqParamsBodyMessages) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfQwenExecuteQwen1_5_14bChatAwqParamsBodyMessages) implementsAccountAIRunCfQwenExecuteQwen1_5_14bChatAwqParamsBodyUnion() {
}

type AccountAIRunCfQwenExecuteQwen1_5_14bChatAwqParamsBodyMessagesMessage struct {
	// The content of the message as a string.
	Content param.Field[string] `json:"content,required"`
	// The role of the message sender (e.g., 'user', 'assistant', 'system', 'tool').
	Role param.Field[string] `json:"role,required"`
}

func (r AccountAIRunCfQwenExecuteQwen1_5_14bChatAwqParamsBodyMessagesMessage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfQwenExecuteQwen1_5_14bChatAwqParamsBodyMessagesFunction struct {
	Code param.Field[string] `json:"code,required"`
	Name param.Field[string] `json:"name,required"`
}

func (r AccountAIRunCfQwenExecuteQwen1_5_14bChatAwqParamsBodyMessagesFunction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfQwenExecuteQwen1_5_14bChatAwqParamsBodyMessagesResponseFormat struct {
	JsonSchema param.Field[interface{}]                                                                     `json:"json_schema"`
	Type       param.Field[AccountAIRunCfQwenExecuteQwen1_5_14bChatAwqParamsBodyMessagesResponseFormatType] `json:"type"`
}

func (r AccountAIRunCfQwenExecuteQwen1_5_14bChatAwqParamsBodyMessagesResponseFormat) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfQwenExecuteQwen1_5_14bChatAwqParamsBodyMessagesResponseFormatType string

const (
	AccountAIRunCfQwenExecuteQwen1_5_14bChatAwqParamsBodyMessagesResponseFormatTypeJsonObject AccountAIRunCfQwenExecuteQwen1_5_14bChatAwqParamsBodyMessagesResponseFormatType = "json_object"
	AccountAIRunCfQwenExecuteQwen1_5_14bChatAwqParamsBodyMessagesResponseFormatTypeJsonSchema AccountAIRunCfQwenExecuteQwen1_5_14bChatAwqParamsBodyMessagesResponseFormatType = "json_schema"
)

func (r AccountAIRunCfQwenExecuteQwen1_5_14bChatAwqParamsBodyMessagesResponseFormatType) IsKnown() bool {
	switch r {
	case AccountAIRunCfQwenExecuteQwen1_5_14bChatAwqParamsBodyMessagesResponseFormatTypeJsonObject, AccountAIRunCfQwenExecuteQwen1_5_14bChatAwqParamsBodyMessagesResponseFormatTypeJsonSchema:
		return true
	}
	return false
}

type AccountAIRunCfQwenExecuteQwen1_5_14bChatAwqParamsBodyMessagesTool struct {
	// A brief description of what the tool does.
	Description param.Field[string]      `json:"description"`
	Function    param.Field[interface{}] `json:"function"`
	// The name of the tool. More descriptive the better.
	Name       param.Field[string]      `json:"name"`
	Parameters param.Field[interface{}] `json:"parameters"`
	// Specifies the type of tool (e.g., 'function').
	Type param.Field[string] `json:"type"`
}

func (r AccountAIRunCfQwenExecuteQwen1_5_14bChatAwqParamsBodyMessagesTool) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfQwenExecuteQwen1_5_14bChatAwqParamsBodyMessagesTool) implementsAccountAIRunCfQwenExecuteQwen1_5_14bChatAwqParamsBodyMessagesToolUnion() {
}

// Satisfied by
// [AccountAIRunCfQwenExecuteQwen1_5_14bChatAwqParamsBodyMessagesToolsObject],
// [AccountAIRunCfQwenExecuteQwen1_5_14bChatAwqParamsBodyMessagesToolsObject],
// [AccountAIRunCfQwenExecuteQwen1_5_14bChatAwqParamsBodyMessagesTool].
type AccountAIRunCfQwenExecuteQwen1_5_14bChatAwqParamsBodyMessagesToolUnion interface {
	implementsAccountAIRunCfQwenExecuteQwen1_5_14bChatAwqParamsBodyMessagesToolUnion()
}

type AccountAIRunCfQwenExecuteQwen1_5_14bChatAwqParamsBodyMessagesToolsObject struct {
	// A brief description of what the tool does.
	Description param.Field[string] `json:"description,required"`
	// The name of the tool. More descriptive the better.
	Name param.Field[string] `json:"name,required"`
	// Schema defining the parameters accepted by the tool.
	Parameters param.Field[AccountAIRunCfQwenExecuteQwen1_5_14bChatAwqParamsBodyMessagesToolsObjectParameters] `json:"parameters,required"`
}

func (r AccountAIRunCfQwenExecuteQwen1_5_14bChatAwqParamsBodyMessagesToolsObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfQwenExecuteQwen1_5_14bChatAwqParamsBodyMessagesToolsObject) implementsAccountAIRunCfQwenExecuteQwen1_5_14bChatAwqParamsBodyMessagesToolUnion() {
}

// Schema defining the parameters accepted by the tool.
type AccountAIRunCfQwenExecuteQwen1_5_14bChatAwqParamsBodyMessagesToolsObjectParameters struct {
	// Definitions of each parameter.
	Properties param.Field[map[string]AccountAIRunCfQwenExecuteQwen1_5_14bChatAwqParamsBodyMessagesToolsObjectParametersProperties] `json:"properties,required"`
	// The type of the parameters object (usually 'object').
	Type param.Field[string] `json:"type,required"`
	// List of required parameter names.
	Required param.Field[[]string] `json:"required"`
}

func (r AccountAIRunCfQwenExecuteQwen1_5_14bChatAwqParamsBodyMessagesToolsObjectParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfQwenExecuteQwen1_5_14bChatAwqParamsBodyMessagesToolsObjectParametersProperties struct {
	// A description of the expected parameter.
	Description param.Field[string] `json:"description,required"`
	// The data type of the parameter.
	Type param.Field[string] `json:"type,required"`
}

func (r AccountAIRunCfQwenExecuteQwen1_5_14bChatAwqParamsBodyMessagesToolsObjectParametersProperties) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfQwenExecuteQwen1_5_7bChatAwqParams struct {
	QueueRequest param.Field[string]                                       `query:"queueRequest"`
	Body         AccountAIRunCfQwenExecuteQwen1_5_7bChatAwqParamsBodyUnion `json:"body"`
}

func (r AccountAIRunCfQwenExecuteQwen1_5_7bChatAwqParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

// URLQuery serializes [AccountAIRunCfQwenExecuteQwen1_5_7bChatAwqParams]'s query
// parameters as `url.Values`.
func (r AccountAIRunCfQwenExecuteQwen1_5_7bChatAwqParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountAIRunCfQwenExecuteQwen1_5_7bChatAwqParamsBody struct {
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

func (r AccountAIRunCfQwenExecuteQwen1_5_7bChatAwqParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfQwenExecuteQwen1_5_7bChatAwqParamsBody) implementsAccountAIRunCfQwenExecuteQwen1_5_7bChatAwqParamsBodyUnion() {
}

// Satisfied by [AccountAIRunCfQwenExecuteQwen1_5_7bChatAwqParamsBodyPrompt],
// [AccountAIRunCfQwenExecuteQwen1_5_7bChatAwqParamsBodyMessages],
// [AccountAIRunCfQwenExecuteQwen1_5_7bChatAwqParamsBody].
type AccountAIRunCfQwenExecuteQwen1_5_7bChatAwqParamsBodyUnion interface {
	implementsAccountAIRunCfQwenExecuteQwen1_5_7bChatAwqParamsBodyUnion()
}

type AccountAIRunCfQwenExecuteQwen1_5_7bChatAwqParamsBodyPrompt struct {
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
	ResponseFormat    param.Field[AccountAIRunCfQwenExecuteQwen1_5_7bChatAwqParamsBodyPromptResponseFormat] `json:"response_format"`
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

func (r AccountAIRunCfQwenExecuteQwen1_5_7bChatAwqParamsBodyPrompt) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfQwenExecuteQwen1_5_7bChatAwqParamsBodyPrompt) implementsAccountAIRunCfQwenExecuteQwen1_5_7bChatAwqParamsBodyUnion() {
}

type AccountAIRunCfQwenExecuteQwen1_5_7bChatAwqParamsBodyPromptResponseFormat struct {
	JsonSchema param.Field[interface{}]                                                                  `json:"json_schema"`
	Type       param.Field[AccountAIRunCfQwenExecuteQwen1_5_7bChatAwqParamsBodyPromptResponseFormatType] `json:"type"`
}

func (r AccountAIRunCfQwenExecuteQwen1_5_7bChatAwqParamsBodyPromptResponseFormat) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfQwenExecuteQwen1_5_7bChatAwqParamsBodyPromptResponseFormatType string

const (
	AccountAIRunCfQwenExecuteQwen1_5_7bChatAwqParamsBodyPromptResponseFormatTypeJsonObject AccountAIRunCfQwenExecuteQwen1_5_7bChatAwqParamsBodyPromptResponseFormatType = "json_object"
	AccountAIRunCfQwenExecuteQwen1_5_7bChatAwqParamsBodyPromptResponseFormatTypeJsonSchema AccountAIRunCfQwenExecuteQwen1_5_7bChatAwqParamsBodyPromptResponseFormatType = "json_schema"
)

func (r AccountAIRunCfQwenExecuteQwen1_5_7bChatAwqParamsBodyPromptResponseFormatType) IsKnown() bool {
	switch r {
	case AccountAIRunCfQwenExecuteQwen1_5_7bChatAwqParamsBodyPromptResponseFormatTypeJsonObject, AccountAIRunCfQwenExecuteQwen1_5_7bChatAwqParamsBodyPromptResponseFormatTypeJsonSchema:
		return true
	}
	return false
}

type AccountAIRunCfQwenExecuteQwen1_5_7bChatAwqParamsBodyMessages struct {
	// An array of message objects representing the conversation history.
	Messages param.Field[[]AccountAIRunCfQwenExecuteQwen1_5_7bChatAwqParamsBodyMessagesMessage] `json:"messages,required"`
	// Decreases the likelihood of the model repeating the same lines verbatim.
	FrequencyPenalty param.Field[float64]                                                                `json:"frequency_penalty"`
	Functions        param.Field[[]AccountAIRunCfQwenExecuteQwen1_5_7bChatAwqParamsBodyMessagesFunction] `json:"functions"`
	// The maximum number of tokens to generate in the response.
	MaxTokens param.Field[int64] `json:"max_tokens"`
	// Increases the likelihood of the model introducing new topics.
	PresencePenalty param.Field[float64] `json:"presence_penalty"`
	// If true, a chat template is not applied and you must adhere to the specific
	// model's expected formatting.
	Raw param.Field[bool] `json:"raw"`
	// Penalty for repeated tokens; higher values discourage repetition.
	RepetitionPenalty param.Field[float64]                                                                    `json:"repetition_penalty"`
	ResponseFormat    param.Field[AccountAIRunCfQwenExecuteQwen1_5_7bChatAwqParamsBodyMessagesResponseFormat] `json:"response_format"`
	// Random seed for reproducibility of the generation.
	Seed param.Field[int64] `json:"seed"`
	// If true, the response will be streamed back incrementally using SSE, Server Sent
	// Events.
	Stream param.Field[bool] `json:"stream"`
	// Controls the randomness of the output; higher values produce more random
	// results.
	Temperature param.Field[float64] `json:"temperature"`
	// A list of tools available for the assistant to use.
	Tools param.Field[[]AccountAIRunCfQwenExecuteQwen1_5_7bChatAwqParamsBodyMessagesToolUnion] `json:"tools"`
	// Limits the AI to choose from the top 'k' most probable words. Lower values make
	// responses more focused; higher values introduce more variety and potential
	// surprises.
	TopK param.Field[int64] `json:"top_k"`
	// Adjusts the creativity of the AI's responses by controlling how many possible
	// words it considers. Lower values make outputs more predictable; higher values
	// allow for more varied and creative responses.
	TopP param.Field[float64] `json:"top_p"`
}

func (r AccountAIRunCfQwenExecuteQwen1_5_7bChatAwqParamsBodyMessages) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfQwenExecuteQwen1_5_7bChatAwqParamsBodyMessages) implementsAccountAIRunCfQwenExecuteQwen1_5_7bChatAwqParamsBodyUnion() {
}

type AccountAIRunCfQwenExecuteQwen1_5_7bChatAwqParamsBodyMessagesMessage struct {
	// The content of the message as a string.
	Content param.Field[string] `json:"content,required"`
	// The role of the message sender (e.g., 'user', 'assistant', 'system', 'tool').
	Role param.Field[string] `json:"role,required"`
}

func (r AccountAIRunCfQwenExecuteQwen1_5_7bChatAwqParamsBodyMessagesMessage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfQwenExecuteQwen1_5_7bChatAwqParamsBodyMessagesFunction struct {
	Code param.Field[string] `json:"code,required"`
	Name param.Field[string] `json:"name,required"`
}

func (r AccountAIRunCfQwenExecuteQwen1_5_7bChatAwqParamsBodyMessagesFunction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfQwenExecuteQwen1_5_7bChatAwqParamsBodyMessagesResponseFormat struct {
	JsonSchema param.Field[interface{}]                                                                    `json:"json_schema"`
	Type       param.Field[AccountAIRunCfQwenExecuteQwen1_5_7bChatAwqParamsBodyMessagesResponseFormatType] `json:"type"`
}

func (r AccountAIRunCfQwenExecuteQwen1_5_7bChatAwqParamsBodyMessagesResponseFormat) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfQwenExecuteQwen1_5_7bChatAwqParamsBodyMessagesResponseFormatType string

const (
	AccountAIRunCfQwenExecuteQwen1_5_7bChatAwqParamsBodyMessagesResponseFormatTypeJsonObject AccountAIRunCfQwenExecuteQwen1_5_7bChatAwqParamsBodyMessagesResponseFormatType = "json_object"
	AccountAIRunCfQwenExecuteQwen1_5_7bChatAwqParamsBodyMessagesResponseFormatTypeJsonSchema AccountAIRunCfQwenExecuteQwen1_5_7bChatAwqParamsBodyMessagesResponseFormatType = "json_schema"
)

func (r AccountAIRunCfQwenExecuteQwen1_5_7bChatAwqParamsBodyMessagesResponseFormatType) IsKnown() bool {
	switch r {
	case AccountAIRunCfQwenExecuteQwen1_5_7bChatAwqParamsBodyMessagesResponseFormatTypeJsonObject, AccountAIRunCfQwenExecuteQwen1_5_7bChatAwqParamsBodyMessagesResponseFormatTypeJsonSchema:
		return true
	}
	return false
}

type AccountAIRunCfQwenExecuteQwen1_5_7bChatAwqParamsBodyMessagesTool struct {
	// A brief description of what the tool does.
	Description param.Field[string]      `json:"description"`
	Function    param.Field[interface{}] `json:"function"`
	// The name of the tool. More descriptive the better.
	Name       param.Field[string]      `json:"name"`
	Parameters param.Field[interface{}] `json:"parameters"`
	// Specifies the type of tool (e.g., 'function').
	Type param.Field[string] `json:"type"`
}

func (r AccountAIRunCfQwenExecuteQwen1_5_7bChatAwqParamsBodyMessagesTool) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfQwenExecuteQwen1_5_7bChatAwqParamsBodyMessagesTool) implementsAccountAIRunCfQwenExecuteQwen1_5_7bChatAwqParamsBodyMessagesToolUnion() {
}

// Satisfied by
// [AccountAIRunCfQwenExecuteQwen1_5_7bChatAwqParamsBodyMessagesToolsObject],
// [AccountAIRunCfQwenExecuteQwen1_5_7bChatAwqParamsBodyMessagesToolsObject],
// [AccountAIRunCfQwenExecuteQwen1_5_7bChatAwqParamsBodyMessagesTool].
type AccountAIRunCfQwenExecuteQwen1_5_7bChatAwqParamsBodyMessagesToolUnion interface {
	implementsAccountAIRunCfQwenExecuteQwen1_5_7bChatAwqParamsBodyMessagesToolUnion()
}

type AccountAIRunCfQwenExecuteQwen1_5_7bChatAwqParamsBodyMessagesToolsObject struct {
	// A brief description of what the tool does.
	Description param.Field[string] `json:"description,required"`
	// The name of the tool. More descriptive the better.
	Name param.Field[string] `json:"name,required"`
	// Schema defining the parameters accepted by the tool.
	Parameters param.Field[AccountAIRunCfQwenExecuteQwen1_5_7bChatAwqParamsBodyMessagesToolsObjectParameters] `json:"parameters,required"`
}

func (r AccountAIRunCfQwenExecuteQwen1_5_7bChatAwqParamsBodyMessagesToolsObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfQwenExecuteQwen1_5_7bChatAwqParamsBodyMessagesToolsObject) implementsAccountAIRunCfQwenExecuteQwen1_5_7bChatAwqParamsBodyMessagesToolUnion() {
}

// Schema defining the parameters accepted by the tool.
type AccountAIRunCfQwenExecuteQwen1_5_7bChatAwqParamsBodyMessagesToolsObjectParameters struct {
	// Definitions of each parameter.
	Properties param.Field[map[string]AccountAIRunCfQwenExecuteQwen1_5_7bChatAwqParamsBodyMessagesToolsObjectParametersProperties] `json:"properties,required"`
	// The type of the parameters object (usually 'object').
	Type param.Field[string] `json:"type,required"`
	// List of required parameter names.
	Required param.Field[[]string] `json:"required"`
}

func (r AccountAIRunCfQwenExecuteQwen1_5_7bChatAwqParamsBodyMessagesToolsObjectParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfQwenExecuteQwen1_5_7bChatAwqParamsBodyMessagesToolsObjectParametersProperties struct {
	// A description of the expected parameter.
	Description param.Field[string] `json:"description,required"`
	// The data type of the parameter.
	Type param.Field[string] `json:"type,required"`
}

func (r AccountAIRunCfQwenExecuteQwen1_5_7bChatAwqParamsBodyMessagesToolsObjectParametersProperties) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
