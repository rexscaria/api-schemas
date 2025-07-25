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

// AccountAIRunCfBaaiService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountAIRunCfBaaiService] method instead.
type AccountAIRunCfBaaiService struct {
	Options []option.RequestOption
}

// NewAccountAIRunCfBaaiService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountAIRunCfBaaiService(opts ...option.RequestOption) (r *AccountAIRunCfBaaiService) {
	r = &AccountAIRunCfBaaiService{}
	r.Options = opts
	return
}

// Execute @cf/baai/bge-base-en-v1.5 model.
func (r *AccountAIRunCfBaaiService) ExecuteBgeBaseEnV1_5(ctx context.Context, accountID string, params AccountAIRunCfBaaiExecuteBgeBaseEnV1_5Params, opts ...option.RequestOption) (res *AccountAIRunCfBaaiExecuteBgeBaseEnV1_5Response, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai/run/@cf/baai/bge-base-en-v1.5", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Execute @cf/baai/bge-large-en-v1.5 model.
func (r *AccountAIRunCfBaaiService) ExecuteBgeLargeEnV1_5(ctx context.Context, accountID string, params AccountAIRunCfBaaiExecuteBgeLargeEnV1_5Params, opts ...option.RequestOption) (res *AccountAIRunCfBaaiExecuteBgeLargeEnV1_5Response, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai/run/@cf/baai/bge-large-en-v1.5", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Execute @cf/baai/bge-m3 model.
func (r *AccountAIRunCfBaaiService) ExecuteBgeM3(ctx context.Context, accountID string, params AccountAIRunCfBaaiExecuteBgeM3Params, opts ...option.RequestOption) (res *AccountAIRunCfBaaiExecuteBgeM3Response, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai/run/@cf/baai/bge-m3", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Execute @cf/baai/bge-reranker-base model.
func (r *AccountAIRunCfBaaiService) ExecuteBgeRerankerBase(ctx context.Context, accountID string, params AccountAIRunCfBaaiExecuteBgeRerankerBaseParams, opts ...option.RequestOption) (res *AccountAIRunCfBaaiExecuteBgeRerankerBaseResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai/run/@cf/baai/bge-reranker-base", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Execute @cf/baai/bge-small-en-v1.5 model.
func (r *AccountAIRunCfBaaiService) ExecuteBgeSmallEnV1_5(ctx context.Context, accountID string, params AccountAIRunCfBaaiExecuteBgeSmallEnV1_5Params, opts ...option.RequestOption) (res *AccountAIRunCfBaaiExecuteBgeSmallEnV1_5Response, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai/run/@cf/baai/bge-small-en-v1.5", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type AccountAIRunCfBaaiExecuteBgeBaseEnV1_5Response = interface{}

type AccountAIRunCfBaaiExecuteBgeLargeEnV1_5Response = interface{}

type AccountAIRunCfBaaiExecuteBgeM3Response = interface{}

type AccountAIRunCfBaaiExecuteBgeRerankerBaseResponse = interface{}

type AccountAIRunCfBaaiExecuteBgeSmallEnV1_5Response = interface{}

type AccountAIRunCfBaaiExecuteBgeBaseEnV1_5Params struct {
	QueueRequest param.Field[string]                                   `query:"queueRequest"`
	Body         AccountAIRunCfBaaiExecuteBgeBaseEnV1_5ParamsBodyUnion `json:"body"`
}

func (r AccountAIRunCfBaaiExecuteBgeBaseEnV1_5Params) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

// URLQuery serializes [AccountAIRunCfBaaiExecuteBgeBaseEnV1_5Params]'s query
// parameters as `url.Values`.
func (r AccountAIRunCfBaaiExecuteBgeBaseEnV1_5Params) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountAIRunCfBaaiExecuteBgeBaseEnV1_5ParamsBody struct {
	// The pooling method used in the embedding process. `cls` pooling will generate
	// more accurate embeddings on larger inputs - however, embeddings created with cls
	// pooling are not compatible with embeddings generated with mean pooling. The
	// default pooling method is `mean` in order for this to not be a breaking change,
	// but we highly suggest using the new `cls` pooling for better accuracy.
	Pooling  param.Field[AccountAIRunCfBaaiExecuteBgeBaseEnV1_5ParamsBodyPooling] `json:"pooling"`
	Requests param.Field[interface{}]                                             `json:"requests"`
	Text     param.Field[interface{}]                                             `json:"text"`
}

func (r AccountAIRunCfBaaiExecuteBgeBaseEnV1_5ParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfBaaiExecuteBgeBaseEnV1_5ParamsBody) implementsAccountAIRunCfBaaiExecuteBgeBaseEnV1_5ParamsBodyUnion() {
}

// Satisfied by [AccountAIRunCfBaaiExecuteBgeBaseEnV1_5ParamsBodyObject],
// [AccountAIRunCfBaaiExecuteBgeBaseEnV1_5ParamsBodyRequests],
// [AccountAIRunCfBaaiExecuteBgeBaseEnV1_5ParamsBody].
type AccountAIRunCfBaaiExecuteBgeBaseEnV1_5ParamsBodyUnion interface {
	implementsAccountAIRunCfBaaiExecuteBgeBaseEnV1_5ParamsBodyUnion()
}

type AccountAIRunCfBaaiExecuteBgeBaseEnV1_5ParamsBodyObject struct {
	// The text to embed
	Text param.Field[AccountAIRunCfBaaiExecuteBgeBaseEnV1_5ParamsBodyObjectTextUnion] `json:"text,required"`
	// The pooling method used in the embedding process. `cls` pooling will generate
	// more accurate embeddings on larger inputs - however, embeddings created with cls
	// pooling are not compatible with embeddings generated with mean pooling. The
	// default pooling method is `mean` in order for this to not be a breaking change,
	// but we highly suggest using the new `cls` pooling for better accuracy.
	Pooling param.Field[AccountAIRunCfBaaiExecuteBgeBaseEnV1_5ParamsBodyObjectPooling] `json:"pooling"`
}

func (r AccountAIRunCfBaaiExecuteBgeBaseEnV1_5ParamsBodyObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfBaaiExecuteBgeBaseEnV1_5ParamsBodyObject) implementsAccountAIRunCfBaaiExecuteBgeBaseEnV1_5ParamsBodyUnion() {
}

// The text to embed
//
// Satisfied by [shared.UnionString],
// [AccountAIRunCfBaaiExecuteBgeBaseEnV1_5ParamsBodyObjectTextArray].
type AccountAIRunCfBaaiExecuteBgeBaseEnV1_5ParamsBodyObjectTextUnion interface {
	ImplementsAccountAIRunCfBaaiExecuteBgeBaseEnV1_5ParamsBodyObjectTextUnion()
}

type AccountAIRunCfBaaiExecuteBgeBaseEnV1_5ParamsBodyObjectTextArray []string

func (r AccountAIRunCfBaaiExecuteBgeBaseEnV1_5ParamsBodyObjectTextArray) ImplementsAccountAIRunCfBaaiExecuteBgeBaseEnV1_5ParamsBodyObjectTextUnion() {
}

// The pooling method used in the embedding process. `cls` pooling will generate
// more accurate embeddings on larger inputs - however, embeddings created with cls
// pooling are not compatible with embeddings generated with mean pooling. The
// default pooling method is `mean` in order for this to not be a breaking change,
// but we highly suggest using the new `cls` pooling for better accuracy.
type AccountAIRunCfBaaiExecuteBgeBaseEnV1_5ParamsBodyObjectPooling string

const (
	AccountAIRunCfBaaiExecuteBgeBaseEnV1_5ParamsBodyObjectPoolingMean AccountAIRunCfBaaiExecuteBgeBaseEnV1_5ParamsBodyObjectPooling = "mean"
	AccountAIRunCfBaaiExecuteBgeBaseEnV1_5ParamsBodyObjectPoolingCls  AccountAIRunCfBaaiExecuteBgeBaseEnV1_5ParamsBodyObjectPooling = "cls"
)

func (r AccountAIRunCfBaaiExecuteBgeBaseEnV1_5ParamsBodyObjectPooling) IsKnown() bool {
	switch r {
	case AccountAIRunCfBaaiExecuteBgeBaseEnV1_5ParamsBodyObjectPoolingMean, AccountAIRunCfBaaiExecuteBgeBaseEnV1_5ParamsBodyObjectPoolingCls:
		return true
	}
	return false
}

type AccountAIRunCfBaaiExecuteBgeBaseEnV1_5ParamsBodyRequests struct {
	// Batch of the embeddings requests to run using async-queue
	Requests param.Field[[]AccountAIRunCfBaaiExecuteBgeBaseEnV1_5ParamsBodyRequestsRequest] `json:"requests,required"`
}

func (r AccountAIRunCfBaaiExecuteBgeBaseEnV1_5ParamsBodyRequests) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfBaaiExecuteBgeBaseEnV1_5ParamsBodyRequests) implementsAccountAIRunCfBaaiExecuteBgeBaseEnV1_5ParamsBodyUnion() {
}

type AccountAIRunCfBaaiExecuteBgeBaseEnV1_5ParamsBodyRequestsRequest struct {
	// The text to embed
	Text param.Field[AccountAIRunCfBaaiExecuteBgeBaseEnV1_5ParamsBodyRequestsRequestsTextUnion] `json:"text,required"`
	// The pooling method used in the embedding process. `cls` pooling will generate
	// more accurate embeddings on larger inputs - however, embeddings created with cls
	// pooling are not compatible with embeddings generated with mean pooling. The
	// default pooling method is `mean` in order for this to not be a breaking change,
	// but we highly suggest using the new `cls` pooling for better accuracy.
	Pooling param.Field[AccountAIRunCfBaaiExecuteBgeBaseEnV1_5ParamsBodyRequestsRequestsPooling] `json:"pooling"`
}

func (r AccountAIRunCfBaaiExecuteBgeBaseEnV1_5ParamsBodyRequestsRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The text to embed
//
// Satisfied by [shared.UnionString],
// [AccountAIRunCfBaaiExecuteBgeBaseEnV1_5ParamsBodyRequestsRequestsTextArray].
type AccountAIRunCfBaaiExecuteBgeBaseEnV1_5ParamsBodyRequestsRequestsTextUnion interface {
	ImplementsAccountAIRunCfBaaiExecuteBgeBaseEnV1_5ParamsBodyRequestsRequestsTextUnion()
}

type AccountAIRunCfBaaiExecuteBgeBaseEnV1_5ParamsBodyRequestsRequestsTextArray []string

func (r AccountAIRunCfBaaiExecuteBgeBaseEnV1_5ParamsBodyRequestsRequestsTextArray) ImplementsAccountAIRunCfBaaiExecuteBgeBaseEnV1_5ParamsBodyRequestsRequestsTextUnion() {
}

// The pooling method used in the embedding process. `cls` pooling will generate
// more accurate embeddings on larger inputs - however, embeddings created with cls
// pooling are not compatible with embeddings generated with mean pooling. The
// default pooling method is `mean` in order for this to not be a breaking change,
// but we highly suggest using the new `cls` pooling for better accuracy.
type AccountAIRunCfBaaiExecuteBgeBaseEnV1_5ParamsBodyRequestsRequestsPooling string

const (
	AccountAIRunCfBaaiExecuteBgeBaseEnV1_5ParamsBodyRequestsRequestsPoolingMean AccountAIRunCfBaaiExecuteBgeBaseEnV1_5ParamsBodyRequestsRequestsPooling = "mean"
	AccountAIRunCfBaaiExecuteBgeBaseEnV1_5ParamsBodyRequestsRequestsPoolingCls  AccountAIRunCfBaaiExecuteBgeBaseEnV1_5ParamsBodyRequestsRequestsPooling = "cls"
)

func (r AccountAIRunCfBaaiExecuteBgeBaseEnV1_5ParamsBodyRequestsRequestsPooling) IsKnown() bool {
	switch r {
	case AccountAIRunCfBaaiExecuteBgeBaseEnV1_5ParamsBodyRequestsRequestsPoolingMean, AccountAIRunCfBaaiExecuteBgeBaseEnV1_5ParamsBodyRequestsRequestsPoolingCls:
		return true
	}
	return false
}

// The pooling method used in the embedding process. `cls` pooling will generate
// more accurate embeddings on larger inputs - however, embeddings created with cls
// pooling are not compatible with embeddings generated with mean pooling. The
// default pooling method is `mean` in order for this to not be a breaking change,
// but we highly suggest using the new `cls` pooling for better accuracy.
type AccountAIRunCfBaaiExecuteBgeBaseEnV1_5ParamsBodyPooling string

const (
	AccountAIRunCfBaaiExecuteBgeBaseEnV1_5ParamsBodyPoolingMean AccountAIRunCfBaaiExecuteBgeBaseEnV1_5ParamsBodyPooling = "mean"
	AccountAIRunCfBaaiExecuteBgeBaseEnV1_5ParamsBodyPoolingCls  AccountAIRunCfBaaiExecuteBgeBaseEnV1_5ParamsBodyPooling = "cls"
)

func (r AccountAIRunCfBaaiExecuteBgeBaseEnV1_5ParamsBodyPooling) IsKnown() bool {
	switch r {
	case AccountAIRunCfBaaiExecuteBgeBaseEnV1_5ParamsBodyPoolingMean, AccountAIRunCfBaaiExecuteBgeBaseEnV1_5ParamsBodyPoolingCls:
		return true
	}
	return false
}

type AccountAIRunCfBaaiExecuteBgeLargeEnV1_5Params struct {
	QueueRequest param.Field[string]                                    `query:"queueRequest"`
	Body         AccountAIRunCfBaaiExecuteBgeLargeEnV1_5ParamsBodyUnion `json:"body"`
}

func (r AccountAIRunCfBaaiExecuteBgeLargeEnV1_5Params) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

// URLQuery serializes [AccountAIRunCfBaaiExecuteBgeLargeEnV1_5Params]'s query
// parameters as `url.Values`.
func (r AccountAIRunCfBaaiExecuteBgeLargeEnV1_5Params) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountAIRunCfBaaiExecuteBgeLargeEnV1_5ParamsBody struct {
	// The pooling method used in the embedding process. `cls` pooling will generate
	// more accurate embeddings on larger inputs - however, embeddings created with cls
	// pooling are not compatible with embeddings generated with mean pooling. The
	// default pooling method is `mean` in order for this to not be a breaking change,
	// but we highly suggest using the new `cls` pooling for better accuracy.
	Pooling  param.Field[AccountAIRunCfBaaiExecuteBgeLargeEnV1_5ParamsBodyPooling] `json:"pooling"`
	Requests param.Field[interface{}]                                              `json:"requests"`
	Text     param.Field[interface{}]                                              `json:"text"`
}

func (r AccountAIRunCfBaaiExecuteBgeLargeEnV1_5ParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfBaaiExecuteBgeLargeEnV1_5ParamsBody) implementsAccountAIRunCfBaaiExecuteBgeLargeEnV1_5ParamsBodyUnion() {
}

// Satisfied by [AccountAIRunCfBaaiExecuteBgeLargeEnV1_5ParamsBodyObject],
// [AccountAIRunCfBaaiExecuteBgeLargeEnV1_5ParamsBodyRequests],
// [AccountAIRunCfBaaiExecuteBgeLargeEnV1_5ParamsBody].
type AccountAIRunCfBaaiExecuteBgeLargeEnV1_5ParamsBodyUnion interface {
	implementsAccountAIRunCfBaaiExecuteBgeLargeEnV1_5ParamsBodyUnion()
}

type AccountAIRunCfBaaiExecuteBgeLargeEnV1_5ParamsBodyObject struct {
	// The text to embed
	Text param.Field[AccountAIRunCfBaaiExecuteBgeLargeEnV1_5ParamsBodyObjectTextUnion] `json:"text,required"`
	// The pooling method used in the embedding process. `cls` pooling will generate
	// more accurate embeddings on larger inputs - however, embeddings created with cls
	// pooling are not compatible with embeddings generated with mean pooling. The
	// default pooling method is `mean` in order for this to not be a breaking change,
	// but we highly suggest using the new `cls` pooling for better accuracy.
	Pooling param.Field[AccountAIRunCfBaaiExecuteBgeLargeEnV1_5ParamsBodyObjectPooling] `json:"pooling"`
}

func (r AccountAIRunCfBaaiExecuteBgeLargeEnV1_5ParamsBodyObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfBaaiExecuteBgeLargeEnV1_5ParamsBodyObject) implementsAccountAIRunCfBaaiExecuteBgeLargeEnV1_5ParamsBodyUnion() {
}

// The text to embed
//
// Satisfied by [shared.UnionString],
// [AccountAIRunCfBaaiExecuteBgeLargeEnV1_5ParamsBodyObjectTextArray].
type AccountAIRunCfBaaiExecuteBgeLargeEnV1_5ParamsBodyObjectTextUnion interface {
	ImplementsAccountAIRunCfBaaiExecuteBgeLargeEnV1_5ParamsBodyObjectTextUnion()
}

type AccountAIRunCfBaaiExecuteBgeLargeEnV1_5ParamsBodyObjectTextArray []string

func (r AccountAIRunCfBaaiExecuteBgeLargeEnV1_5ParamsBodyObjectTextArray) ImplementsAccountAIRunCfBaaiExecuteBgeLargeEnV1_5ParamsBodyObjectTextUnion() {
}

// The pooling method used in the embedding process. `cls` pooling will generate
// more accurate embeddings on larger inputs - however, embeddings created with cls
// pooling are not compatible with embeddings generated with mean pooling. The
// default pooling method is `mean` in order for this to not be a breaking change,
// but we highly suggest using the new `cls` pooling for better accuracy.
type AccountAIRunCfBaaiExecuteBgeLargeEnV1_5ParamsBodyObjectPooling string

const (
	AccountAIRunCfBaaiExecuteBgeLargeEnV1_5ParamsBodyObjectPoolingMean AccountAIRunCfBaaiExecuteBgeLargeEnV1_5ParamsBodyObjectPooling = "mean"
	AccountAIRunCfBaaiExecuteBgeLargeEnV1_5ParamsBodyObjectPoolingCls  AccountAIRunCfBaaiExecuteBgeLargeEnV1_5ParamsBodyObjectPooling = "cls"
)

func (r AccountAIRunCfBaaiExecuteBgeLargeEnV1_5ParamsBodyObjectPooling) IsKnown() bool {
	switch r {
	case AccountAIRunCfBaaiExecuteBgeLargeEnV1_5ParamsBodyObjectPoolingMean, AccountAIRunCfBaaiExecuteBgeLargeEnV1_5ParamsBodyObjectPoolingCls:
		return true
	}
	return false
}

type AccountAIRunCfBaaiExecuteBgeLargeEnV1_5ParamsBodyRequests struct {
	// Batch of the embeddings requests to run using async-queue
	Requests param.Field[[]AccountAIRunCfBaaiExecuteBgeLargeEnV1_5ParamsBodyRequestsRequest] `json:"requests,required"`
}

func (r AccountAIRunCfBaaiExecuteBgeLargeEnV1_5ParamsBodyRequests) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfBaaiExecuteBgeLargeEnV1_5ParamsBodyRequests) implementsAccountAIRunCfBaaiExecuteBgeLargeEnV1_5ParamsBodyUnion() {
}

type AccountAIRunCfBaaiExecuteBgeLargeEnV1_5ParamsBodyRequestsRequest struct {
	// The text to embed
	Text param.Field[AccountAIRunCfBaaiExecuteBgeLargeEnV1_5ParamsBodyRequestsRequestsTextUnion] `json:"text,required"`
	// The pooling method used in the embedding process. `cls` pooling will generate
	// more accurate embeddings on larger inputs - however, embeddings created with cls
	// pooling are not compatible with embeddings generated with mean pooling. The
	// default pooling method is `mean` in order for this to not be a breaking change,
	// but we highly suggest using the new `cls` pooling for better accuracy.
	Pooling param.Field[AccountAIRunCfBaaiExecuteBgeLargeEnV1_5ParamsBodyRequestsRequestsPooling] `json:"pooling"`
}

func (r AccountAIRunCfBaaiExecuteBgeLargeEnV1_5ParamsBodyRequestsRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The text to embed
//
// Satisfied by [shared.UnionString],
// [AccountAIRunCfBaaiExecuteBgeLargeEnV1_5ParamsBodyRequestsRequestsTextArray].
type AccountAIRunCfBaaiExecuteBgeLargeEnV1_5ParamsBodyRequestsRequestsTextUnion interface {
	ImplementsAccountAIRunCfBaaiExecuteBgeLargeEnV1_5ParamsBodyRequestsRequestsTextUnion()
}

type AccountAIRunCfBaaiExecuteBgeLargeEnV1_5ParamsBodyRequestsRequestsTextArray []string

func (r AccountAIRunCfBaaiExecuteBgeLargeEnV1_5ParamsBodyRequestsRequestsTextArray) ImplementsAccountAIRunCfBaaiExecuteBgeLargeEnV1_5ParamsBodyRequestsRequestsTextUnion() {
}

// The pooling method used in the embedding process. `cls` pooling will generate
// more accurate embeddings on larger inputs - however, embeddings created with cls
// pooling are not compatible with embeddings generated with mean pooling. The
// default pooling method is `mean` in order for this to not be a breaking change,
// but we highly suggest using the new `cls` pooling for better accuracy.
type AccountAIRunCfBaaiExecuteBgeLargeEnV1_5ParamsBodyRequestsRequestsPooling string

const (
	AccountAIRunCfBaaiExecuteBgeLargeEnV1_5ParamsBodyRequestsRequestsPoolingMean AccountAIRunCfBaaiExecuteBgeLargeEnV1_5ParamsBodyRequestsRequestsPooling = "mean"
	AccountAIRunCfBaaiExecuteBgeLargeEnV1_5ParamsBodyRequestsRequestsPoolingCls  AccountAIRunCfBaaiExecuteBgeLargeEnV1_5ParamsBodyRequestsRequestsPooling = "cls"
)

func (r AccountAIRunCfBaaiExecuteBgeLargeEnV1_5ParamsBodyRequestsRequestsPooling) IsKnown() bool {
	switch r {
	case AccountAIRunCfBaaiExecuteBgeLargeEnV1_5ParamsBodyRequestsRequestsPoolingMean, AccountAIRunCfBaaiExecuteBgeLargeEnV1_5ParamsBodyRequestsRequestsPoolingCls:
		return true
	}
	return false
}

// The pooling method used in the embedding process. `cls` pooling will generate
// more accurate embeddings on larger inputs - however, embeddings created with cls
// pooling are not compatible with embeddings generated with mean pooling. The
// default pooling method is `mean` in order for this to not be a breaking change,
// but we highly suggest using the new `cls` pooling for better accuracy.
type AccountAIRunCfBaaiExecuteBgeLargeEnV1_5ParamsBodyPooling string

const (
	AccountAIRunCfBaaiExecuteBgeLargeEnV1_5ParamsBodyPoolingMean AccountAIRunCfBaaiExecuteBgeLargeEnV1_5ParamsBodyPooling = "mean"
	AccountAIRunCfBaaiExecuteBgeLargeEnV1_5ParamsBodyPoolingCls  AccountAIRunCfBaaiExecuteBgeLargeEnV1_5ParamsBodyPooling = "cls"
)

func (r AccountAIRunCfBaaiExecuteBgeLargeEnV1_5ParamsBodyPooling) IsKnown() bool {
	switch r {
	case AccountAIRunCfBaaiExecuteBgeLargeEnV1_5ParamsBodyPoolingMean, AccountAIRunCfBaaiExecuteBgeLargeEnV1_5ParamsBodyPoolingCls:
		return true
	}
	return false
}

type AccountAIRunCfBaaiExecuteBgeM3Params struct {
	QueueRequest param.Field[string]                           `query:"queueRequest"`
	Body         AccountAIRunCfBaaiExecuteBgeM3ParamsBodyUnion `json:"body"`
}

func (r AccountAIRunCfBaaiExecuteBgeM3Params) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

// URLQuery serializes [AccountAIRunCfBaaiExecuteBgeM3Params]'s query parameters as
// `url.Values`.
func (r AccountAIRunCfBaaiExecuteBgeM3Params) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountAIRunCfBaaiExecuteBgeM3ParamsBody struct {
	Contexts param.Field[interface{}] `json:"contexts"`
	// A query you wish to perform against the provided contexts. If no query is
	// provided the model with respond with embeddings for contexts
	Query    param.Field[string]      `json:"query"`
	Requests param.Field[interface{}] `json:"requests"`
	Text     param.Field[interface{}] `json:"text"`
	// When provided with too long context should the model error out or truncate the
	// context to fit?
	TruncateInputs param.Field[bool] `json:"truncate_inputs"`
}

func (r AccountAIRunCfBaaiExecuteBgeM3ParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfBaaiExecuteBgeM3ParamsBody) implementsAccountAIRunCfBaaiExecuteBgeM3ParamsBodyUnion() {
}

// Satisfied by
// [AccountAIRunCfBaaiExecuteBgeM3ParamsBodyBgeM3InputQueryAndContexts],
// [AccountAIRunCfBaaiExecuteBgeM3ParamsBodyBgeM3InputEmbedding],
// [AccountAIRunCfBaaiExecuteBgeM3ParamsBodyRequests],
// [AccountAIRunCfBaaiExecuteBgeM3ParamsBody].
type AccountAIRunCfBaaiExecuteBgeM3ParamsBodyUnion interface {
	implementsAccountAIRunCfBaaiExecuteBgeM3ParamsBodyUnion()
}

type AccountAIRunCfBaaiExecuteBgeM3ParamsBodyBgeM3InputQueryAndContexts struct {
	// List of provided contexts. Note that the index in this array is important, as
	// the response will refer to it.
	Contexts param.Field[[]AccountAIRunCfBaaiExecuteBgeM3ParamsBodyBgeM3InputQueryAndContextsContext] `json:"contexts,required"`
	// A query you wish to perform against the provided contexts. If no query is
	// provided the model with respond with embeddings for contexts
	Query param.Field[string] `json:"query"`
	// When provided with too long context should the model error out or truncate the
	// context to fit?
	TruncateInputs param.Field[bool] `json:"truncate_inputs"`
}

func (r AccountAIRunCfBaaiExecuteBgeM3ParamsBodyBgeM3InputQueryAndContexts) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfBaaiExecuteBgeM3ParamsBodyBgeM3InputQueryAndContexts) implementsAccountAIRunCfBaaiExecuteBgeM3ParamsBodyUnion() {
}

type AccountAIRunCfBaaiExecuteBgeM3ParamsBodyBgeM3InputQueryAndContextsContext struct {
	// One of the provided context content
	Text param.Field[string] `json:"text"`
}

func (r AccountAIRunCfBaaiExecuteBgeM3ParamsBodyBgeM3InputQueryAndContextsContext) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfBaaiExecuteBgeM3ParamsBodyBgeM3InputEmbedding struct {
	// The text to embed
	Text param.Field[AccountAIRunCfBaaiExecuteBgeM3ParamsBodyBgeM3InputEmbeddingTextUnion] `json:"text,required"`
	// When provided with too long context should the model error out or truncate the
	// context to fit?
	TruncateInputs param.Field[bool] `json:"truncate_inputs"`
}

func (r AccountAIRunCfBaaiExecuteBgeM3ParamsBodyBgeM3InputEmbedding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfBaaiExecuteBgeM3ParamsBodyBgeM3InputEmbedding) implementsAccountAIRunCfBaaiExecuteBgeM3ParamsBodyUnion() {
}

// The text to embed
//
// Satisfied by [shared.UnionString],
// [AccountAIRunCfBaaiExecuteBgeM3ParamsBodyBgeM3InputEmbeddingTextArray].
type AccountAIRunCfBaaiExecuteBgeM3ParamsBodyBgeM3InputEmbeddingTextUnion interface {
	ImplementsAccountAIRunCfBaaiExecuteBgeM3ParamsBodyBgeM3InputEmbeddingTextUnion()
}

type AccountAIRunCfBaaiExecuteBgeM3ParamsBodyBgeM3InputEmbeddingTextArray []string

func (r AccountAIRunCfBaaiExecuteBgeM3ParamsBodyBgeM3InputEmbeddingTextArray) ImplementsAccountAIRunCfBaaiExecuteBgeM3ParamsBodyBgeM3InputEmbeddingTextUnion() {
}

type AccountAIRunCfBaaiExecuteBgeM3ParamsBodyRequests struct {
	// Batch of the embeddings requests to run using async-queue
	Requests param.Field[[]AccountAIRunCfBaaiExecuteBgeM3ParamsBodyRequestsRequestUnion] `json:"requests,required"`
}

func (r AccountAIRunCfBaaiExecuteBgeM3ParamsBodyRequests) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfBaaiExecuteBgeM3ParamsBodyRequests) implementsAccountAIRunCfBaaiExecuteBgeM3ParamsBodyUnion() {
}

type AccountAIRunCfBaaiExecuteBgeM3ParamsBodyRequestsRequest struct {
	Contexts param.Field[interface{}] `json:"contexts"`
	// A query you wish to perform against the provided contexts. If no query is
	// provided the model with respond with embeddings for contexts
	Query param.Field[string]      `json:"query"`
	Text  param.Field[interface{}] `json:"text"`
	// When provided with too long context should the model error out or truncate the
	// context to fit?
	TruncateInputs param.Field[bool] `json:"truncate_inputs"`
}

func (r AccountAIRunCfBaaiExecuteBgeM3ParamsBodyRequestsRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfBaaiExecuteBgeM3ParamsBodyRequestsRequest) implementsAccountAIRunCfBaaiExecuteBgeM3ParamsBodyRequestsRequestUnion() {
}

// Satisfied by
// [AccountAIRunCfBaaiExecuteBgeM3ParamsBodyRequestsRequestsBgeM3InputQueryAndContexts],
// [AccountAIRunCfBaaiExecuteBgeM3ParamsBodyRequestsRequestsBgeM3InputEmbedding],
// [AccountAIRunCfBaaiExecuteBgeM3ParamsBodyRequestsRequest].
type AccountAIRunCfBaaiExecuteBgeM3ParamsBodyRequestsRequestUnion interface {
	implementsAccountAIRunCfBaaiExecuteBgeM3ParamsBodyRequestsRequestUnion()
}

type AccountAIRunCfBaaiExecuteBgeM3ParamsBodyRequestsRequestsBgeM3InputQueryAndContexts struct {
	// List of provided contexts. Note that the index in this array is important, as
	// the response will refer to it.
	Contexts param.Field[[]AccountAIRunCfBaaiExecuteBgeM3ParamsBodyRequestsRequestsBgeM3InputQueryAndContextsContext] `json:"contexts,required"`
	// A query you wish to perform against the provided contexts. If no query is
	// provided the model with respond with embeddings for contexts
	Query param.Field[string] `json:"query"`
	// When provided with too long context should the model error out or truncate the
	// context to fit?
	TruncateInputs param.Field[bool] `json:"truncate_inputs"`
}

func (r AccountAIRunCfBaaiExecuteBgeM3ParamsBodyRequestsRequestsBgeM3InputQueryAndContexts) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfBaaiExecuteBgeM3ParamsBodyRequestsRequestsBgeM3InputQueryAndContexts) implementsAccountAIRunCfBaaiExecuteBgeM3ParamsBodyRequestsRequestUnion() {
}

type AccountAIRunCfBaaiExecuteBgeM3ParamsBodyRequestsRequestsBgeM3InputQueryAndContextsContext struct {
	// One of the provided context content
	Text param.Field[string] `json:"text"`
}

func (r AccountAIRunCfBaaiExecuteBgeM3ParamsBodyRequestsRequestsBgeM3InputQueryAndContextsContext) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfBaaiExecuteBgeM3ParamsBodyRequestsRequestsBgeM3InputEmbedding struct {
	// The text to embed
	Text param.Field[AccountAIRunCfBaaiExecuteBgeM3ParamsBodyRequestsRequestsBgeM3InputEmbeddingTextUnion] `json:"text,required"`
	// When provided with too long context should the model error out or truncate the
	// context to fit?
	TruncateInputs param.Field[bool] `json:"truncate_inputs"`
}

func (r AccountAIRunCfBaaiExecuteBgeM3ParamsBodyRequestsRequestsBgeM3InputEmbedding) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfBaaiExecuteBgeM3ParamsBodyRequestsRequestsBgeM3InputEmbedding) implementsAccountAIRunCfBaaiExecuteBgeM3ParamsBodyRequestsRequestUnion() {
}

// The text to embed
//
// Satisfied by [shared.UnionString],
// [AccountAIRunCfBaaiExecuteBgeM3ParamsBodyRequestsRequestsBgeM3InputEmbeddingTextArray].
type AccountAIRunCfBaaiExecuteBgeM3ParamsBodyRequestsRequestsBgeM3InputEmbeddingTextUnion interface {
	ImplementsAccountAIRunCfBaaiExecuteBgeM3ParamsBodyRequestsRequestsBgeM3InputEmbeddingTextUnion()
}

type AccountAIRunCfBaaiExecuteBgeM3ParamsBodyRequestsRequestsBgeM3InputEmbeddingTextArray []string

func (r AccountAIRunCfBaaiExecuteBgeM3ParamsBodyRequestsRequestsBgeM3InputEmbeddingTextArray) ImplementsAccountAIRunCfBaaiExecuteBgeM3ParamsBodyRequestsRequestsBgeM3InputEmbeddingTextUnion() {
}

type AccountAIRunCfBaaiExecuteBgeRerankerBaseParams struct {
	// List of provided contexts. Note that the index in this array is important, as
	// the response will refer to it.
	Contexts param.Field[[]AccountAIRunCfBaaiExecuteBgeRerankerBaseParamsContext] `json:"contexts,required"`
	// A query you wish to perform against the provided contexts.
	Query        param.Field[string] `json:"query,required"`
	QueueRequest param.Field[string] `query:"queueRequest"`
	// Number of returned results starting with the best score.
	TopK param.Field[int64] `json:"top_k"`
}

func (r AccountAIRunCfBaaiExecuteBgeRerankerBaseParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [AccountAIRunCfBaaiExecuteBgeRerankerBaseParams]'s query
// parameters as `url.Values`.
func (r AccountAIRunCfBaaiExecuteBgeRerankerBaseParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountAIRunCfBaaiExecuteBgeRerankerBaseParamsContext struct {
	// One of the provided context content
	Text param.Field[string] `json:"text"`
}

func (r AccountAIRunCfBaaiExecuteBgeRerankerBaseParamsContext) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAIRunCfBaaiExecuteBgeSmallEnV1_5Params struct {
	QueueRequest param.Field[string]                                    `query:"queueRequest"`
	Body         AccountAIRunCfBaaiExecuteBgeSmallEnV1_5ParamsBodyUnion `json:"body"`
}

func (r AccountAIRunCfBaaiExecuteBgeSmallEnV1_5Params) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

// URLQuery serializes [AccountAIRunCfBaaiExecuteBgeSmallEnV1_5Params]'s query
// parameters as `url.Values`.
func (r AccountAIRunCfBaaiExecuteBgeSmallEnV1_5Params) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountAIRunCfBaaiExecuteBgeSmallEnV1_5ParamsBody struct {
	// The pooling method used in the embedding process. `cls` pooling will generate
	// more accurate embeddings on larger inputs - however, embeddings created with cls
	// pooling are not compatible with embeddings generated with mean pooling. The
	// default pooling method is `mean` in order for this to not be a breaking change,
	// but we highly suggest using the new `cls` pooling for better accuracy.
	Pooling  param.Field[AccountAIRunCfBaaiExecuteBgeSmallEnV1_5ParamsBodyPooling] `json:"pooling"`
	Requests param.Field[interface{}]                                              `json:"requests"`
	Text     param.Field[interface{}]                                              `json:"text"`
}

func (r AccountAIRunCfBaaiExecuteBgeSmallEnV1_5ParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfBaaiExecuteBgeSmallEnV1_5ParamsBody) implementsAccountAIRunCfBaaiExecuteBgeSmallEnV1_5ParamsBodyUnion() {
}

// Satisfied by [AccountAIRunCfBaaiExecuteBgeSmallEnV1_5ParamsBodyObject],
// [AccountAIRunCfBaaiExecuteBgeSmallEnV1_5ParamsBodyRequests],
// [AccountAIRunCfBaaiExecuteBgeSmallEnV1_5ParamsBody].
type AccountAIRunCfBaaiExecuteBgeSmallEnV1_5ParamsBodyUnion interface {
	implementsAccountAIRunCfBaaiExecuteBgeSmallEnV1_5ParamsBodyUnion()
}

type AccountAIRunCfBaaiExecuteBgeSmallEnV1_5ParamsBodyObject struct {
	// The text to embed
	Text param.Field[AccountAIRunCfBaaiExecuteBgeSmallEnV1_5ParamsBodyObjectTextUnion] `json:"text,required"`
	// The pooling method used in the embedding process. `cls` pooling will generate
	// more accurate embeddings on larger inputs - however, embeddings created with cls
	// pooling are not compatible with embeddings generated with mean pooling. The
	// default pooling method is `mean` in order for this to not be a breaking change,
	// but we highly suggest using the new `cls` pooling for better accuracy.
	Pooling param.Field[AccountAIRunCfBaaiExecuteBgeSmallEnV1_5ParamsBodyObjectPooling] `json:"pooling"`
}

func (r AccountAIRunCfBaaiExecuteBgeSmallEnV1_5ParamsBodyObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfBaaiExecuteBgeSmallEnV1_5ParamsBodyObject) implementsAccountAIRunCfBaaiExecuteBgeSmallEnV1_5ParamsBodyUnion() {
}

// The text to embed
//
// Satisfied by [shared.UnionString],
// [AccountAIRunCfBaaiExecuteBgeSmallEnV1_5ParamsBodyObjectTextArray].
type AccountAIRunCfBaaiExecuteBgeSmallEnV1_5ParamsBodyObjectTextUnion interface {
	ImplementsAccountAIRunCfBaaiExecuteBgeSmallEnV1_5ParamsBodyObjectTextUnion()
}

type AccountAIRunCfBaaiExecuteBgeSmallEnV1_5ParamsBodyObjectTextArray []string

func (r AccountAIRunCfBaaiExecuteBgeSmallEnV1_5ParamsBodyObjectTextArray) ImplementsAccountAIRunCfBaaiExecuteBgeSmallEnV1_5ParamsBodyObjectTextUnion() {
}

// The pooling method used in the embedding process. `cls` pooling will generate
// more accurate embeddings on larger inputs - however, embeddings created with cls
// pooling are not compatible with embeddings generated with mean pooling. The
// default pooling method is `mean` in order for this to not be a breaking change,
// but we highly suggest using the new `cls` pooling for better accuracy.
type AccountAIRunCfBaaiExecuteBgeSmallEnV1_5ParamsBodyObjectPooling string

const (
	AccountAIRunCfBaaiExecuteBgeSmallEnV1_5ParamsBodyObjectPoolingMean AccountAIRunCfBaaiExecuteBgeSmallEnV1_5ParamsBodyObjectPooling = "mean"
	AccountAIRunCfBaaiExecuteBgeSmallEnV1_5ParamsBodyObjectPoolingCls  AccountAIRunCfBaaiExecuteBgeSmallEnV1_5ParamsBodyObjectPooling = "cls"
)

func (r AccountAIRunCfBaaiExecuteBgeSmallEnV1_5ParamsBodyObjectPooling) IsKnown() bool {
	switch r {
	case AccountAIRunCfBaaiExecuteBgeSmallEnV1_5ParamsBodyObjectPoolingMean, AccountAIRunCfBaaiExecuteBgeSmallEnV1_5ParamsBodyObjectPoolingCls:
		return true
	}
	return false
}

type AccountAIRunCfBaaiExecuteBgeSmallEnV1_5ParamsBodyRequests struct {
	// Batch of the embeddings requests to run using async-queue
	Requests param.Field[[]AccountAIRunCfBaaiExecuteBgeSmallEnV1_5ParamsBodyRequestsRequest] `json:"requests,required"`
}

func (r AccountAIRunCfBaaiExecuteBgeSmallEnV1_5ParamsBodyRequests) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountAIRunCfBaaiExecuteBgeSmallEnV1_5ParamsBodyRequests) implementsAccountAIRunCfBaaiExecuteBgeSmallEnV1_5ParamsBodyUnion() {
}

type AccountAIRunCfBaaiExecuteBgeSmallEnV1_5ParamsBodyRequestsRequest struct {
	// The text to embed
	Text param.Field[AccountAIRunCfBaaiExecuteBgeSmallEnV1_5ParamsBodyRequestsRequestsTextUnion] `json:"text,required"`
	// The pooling method used in the embedding process. `cls` pooling will generate
	// more accurate embeddings on larger inputs - however, embeddings created with cls
	// pooling are not compatible with embeddings generated with mean pooling. The
	// default pooling method is `mean` in order for this to not be a breaking change,
	// but we highly suggest using the new `cls` pooling for better accuracy.
	Pooling param.Field[AccountAIRunCfBaaiExecuteBgeSmallEnV1_5ParamsBodyRequestsRequestsPooling] `json:"pooling"`
}

func (r AccountAIRunCfBaaiExecuteBgeSmallEnV1_5ParamsBodyRequestsRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The text to embed
//
// Satisfied by [shared.UnionString],
// [AccountAIRunCfBaaiExecuteBgeSmallEnV1_5ParamsBodyRequestsRequestsTextArray].
type AccountAIRunCfBaaiExecuteBgeSmallEnV1_5ParamsBodyRequestsRequestsTextUnion interface {
	ImplementsAccountAIRunCfBaaiExecuteBgeSmallEnV1_5ParamsBodyRequestsRequestsTextUnion()
}

type AccountAIRunCfBaaiExecuteBgeSmallEnV1_5ParamsBodyRequestsRequestsTextArray []string

func (r AccountAIRunCfBaaiExecuteBgeSmallEnV1_5ParamsBodyRequestsRequestsTextArray) ImplementsAccountAIRunCfBaaiExecuteBgeSmallEnV1_5ParamsBodyRequestsRequestsTextUnion() {
}

// The pooling method used in the embedding process. `cls` pooling will generate
// more accurate embeddings on larger inputs - however, embeddings created with cls
// pooling are not compatible with embeddings generated with mean pooling. The
// default pooling method is `mean` in order for this to not be a breaking change,
// but we highly suggest using the new `cls` pooling for better accuracy.
type AccountAIRunCfBaaiExecuteBgeSmallEnV1_5ParamsBodyRequestsRequestsPooling string

const (
	AccountAIRunCfBaaiExecuteBgeSmallEnV1_5ParamsBodyRequestsRequestsPoolingMean AccountAIRunCfBaaiExecuteBgeSmallEnV1_5ParamsBodyRequestsRequestsPooling = "mean"
	AccountAIRunCfBaaiExecuteBgeSmallEnV1_5ParamsBodyRequestsRequestsPoolingCls  AccountAIRunCfBaaiExecuteBgeSmallEnV1_5ParamsBodyRequestsRequestsPooling = "cls"
)

func (r AccountAIRunCfBaaiExecuteBgeSmallEnV1_5ParamsBodyRequestsRequestsPooling) IsKnown() bool {
	switch r {
	case AccountAIRunCfBaaiExecuteBgeSmallEnV1_5ParamsBodyRequestsRequestsPoolingMean, AccountAIRunCfBaaiExecuteBgeSmallEnV1_5ParamsBodyRequestsRequestsPoolingCls:
		return true
	}
	return false
}

// The pooling method used in the embedding process. `cls` pooling will generate
// more accurate embeddings on larger inputs - however, embeddings created with cls
// pooling are not compatible with embeddings generated with mean pooling. The
// default pooling method is `mean` in order for this to not be a breaking change,
// but we highly suggest using the new `cls` pooling for better accuracy.
type AccountAIRunCfBaaiExecuteBgeSmallEnV1_5ParamsBodyPooling string

const (
	AccountAIRunCfBaaiExecuteBgeSmallEnV1_5ParamsBodyPoolingMean AccountAIRunCfBaaiExecuteBgeSmallEnV1_5ParamsBodyPooling = "mean"
	AccountAIRunCfBaaiExecuteBgeSmallEnV1_5ParamsBodyPoolingCls  AccountAIRunCfBaaiExecuteBgeSmallEnV1_5ParamsBodyPooling = "cls"
)

func (r AccountAIRunCfBaaiExecuteBgeSmallEnV1_5ParamsBodyPooling) IsKnown() bool {
	switch r {
	case AccountAIRunCfBaaiExecuteBgeSmallEnV1_5ParamsBodyPoolingMean, AccountAIRunCfBaaiExecuteBgeSmallEnV1_5ParamsBodyPoolingCls:
		return true
	}
	return false
}
