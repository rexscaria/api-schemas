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
	"github.com/tidwall/gjson"
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

type AccountAIRunCfBaaiExecuteBgeBaseEnV1_5Response struct {
	Result  AccountAIRunCfBaaiExecuteBgeBaseEnV1_5ResponseResult `json:"result"`
	Success bool                                                 `json:"success"`
	JSON    accountAIRunCfBaaiExecuteBgeBaseEnV1_5ResponseJSON   `json:"-"`
}

// accountAIRunCfBaaiExecuteBgeBaseEnV1_5ResponseJSON contains the JSON metadata
// for the struct [AccountAIRunCfBaaiExecuteBgeBaseEnV1_5Response]
type accountAIRunCfBaaiExecuteBgeBaseEnV1_5ResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIRunCfBaaiExecuteBgeBaseEnV1_5Response) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunCfBaaiExecuteBgeBaseEnV1_5ResponseJSON) RawJSON() string {
	return r.raw
}

type AccountAIRunCfBaaiExecuteBgeBaseEnV1_5ResponseResult struct {
	// Embeddings of the requested text values
	Data  [][]float64                                              `json:"data"`
	Shape []float64                                                `json:"shape"`
	JSON  accountAIRunCfBaaiExecuteBgeBaseEnV1_5ResponseResultJSON `json:"-"`
}

// accountAIRunCfBaaiExecuteBgeBaseEnV1_5ResponseResultJSON contains the JSON
// metadata for the struct [AccountAIRunCfBaaiExecuteBgeBaseEnV1_5ResponseResult]
type accountAIRunCfBaaiExecuteBgeBaseEnV1_5ResponseResultJSON struct {
	Data        apijson.Field
	Shape       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIRunCfBaaiExecuteBgeBaseEnV1_5ResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunCfBaaiExecuteBgeBaseEnV1_5ResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountAIRunCfBaaiExecuteBgeLargeEnV1_5Response struct {
	Result  AccountAIRunCfBaaiExecuteBgeLargeEnV1_5ResponseResult `json:"result"`
	Success bool                                                  `json:"success"`
	JSON    accountAIRunCfBaaiExecuteBgeLargeEnV1_5ResponseJSON   `json:"-"`
}

// accountAIRunCfBaaiExecuteBgeLargeEnV1_5ResponseJSON contains the JSON metadata
// for the struct [AccountAIRunCfBaaiExecuteBgeLargeEnV1_5Response]
type accountAIRunCfBaaiExecuteBgeLargeEnV1_5ResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIRunCfBaaiExecuteBgeLargeEnV1_5Response) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunCfBaaiExecuteBgeLargeEnV1_5ResponseJSON) RawJSON() string {
	return r.raw
}

type AccountAIRunCfBaaiExecuteBgeLargeEnV1_5ResponseResult struct {
	// Embeddings of the requested text values
	Data  [][]float64                                               `json:"data"`
	Shape []float64                                                 `json:"shape"`
	JSON  accountAIRunCfBaaiExecuteBgeLargeEnV1_5ResponseResultJSON `json:"-"`
}

// accountAIRunCfBaaiExecuteBgeLargeEnV1_5ResponseResultJSON contains the JSON
// metadata for the struct [AccountAIRunCfBaaiExecuteBgeLargeEnV1_5ResponseResult]
type accountAIRunCfBaaiExecuteBgeLargeEnV1_5ResponseResultJSON struct {
	Data        apijson.Field
	Shape       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIRunCfBaaiExecuteBgeLargeEnV1_5ResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunCfBaaiExecuteBgeLargeEnV1_5ResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountAIRunCfBaaiExecuteBgeM3Response struct {
	Result  AccountAIRunCfBaaiExecuteBgeM3ResponseResult `json:"result"`
	Success bool                                         `json:"success"`
	JSON    accountAIRunCfBaaiExecuteBgeM3ResponseJSON   `json:"-"`
}

// accountAIRunCfBaaiExecuteBgeM3ResponseJSON contains the JSON metadata for the
// struct [AccountAIRunCfBaaiExecuteBgeM3Response]
type accountAIRunCfBaaiExecuteBgeM3ResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIRunCfBaaiExecuteBgeM3Response) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunCfBaaiExecuteBgeM3ResponseJSON) RawJSON() string {
	return r.raw
}

type AccountAIRunCfBaaiExecuteBgeM3ResponseResult struct {
	// This field can have the runtime type of [[][]float64].
	Data interface{} `json:"data"`
	// The pooling method used in the embedding process.
	Pooling AccountAIRunCfBaaiExecuteBgeM3ResponseResultPooling `json:"pooling"`
	// This field can have the runtime type of
	// [[]AccountAIRunCfBaaiExecuteBgeM3ResponseResultBgeM3OuputQueryResponse],
	// [[][]float64].
	Response interface{} `json:"response"`
	// This field can have the runtime type of [[]float64].
	Shape interface{}                                      `json:"shape"`
	JSON  accountAIRunCfBaaiExecuteBgeM3ResponseResultJSON `json:"-"`
	union AccountAIRunCfBaaiExecuteBgeM3ResponseResultUnion
}

// accountAIRunCfBaaiExecuteBgeM3ResponseResultJSON contains the JSON metadata for
// the struct [AccountAIRunCfBaaiExecuteBgeM3ResponseResult]
type accountAIRunCfBaaiExecuteBgeM3ResponseResultJSON struct {
	Data        apijson.Field
	Pooling     apijson.Field
	Response    apijson.Field
	Shape       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r accountAIRunCfBaaiExecuteBgeM3ResponseResultJSON) RawJSON() string {
	return r.raw
}

func (r *AccountAIRunCfBaaiExecuteBgeM3ResponseResult) UnmarshalJSON(data []byte) (err error) {
	*r = AccountAIRunCfBaaiExecuteBgeM3ResponseResult{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [AccountAIRunCfBaaiExecuteBgeM3ResponseResultUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [AccountAIRunCfBaaiExecuteBgeM3ResponseResultBgeM3OuputQuery],
// [AccountAIRunCfBaaiExecuteBgeM3ResponseResultBgeM3OutputEmbeddingForContexts],
// [AccountAIRunCfBaaiExecuteBgeM3ResponseResultBgeM3OuputEmbedding].
func (r AccountAIRunCfBaaiExecuteBgeM3ResponseResult) AsUnion() AccountAIRunCfBaaiExecuteBgeM3ResponseResultUnion {
	return r.union
}

// Union satisfied by
// [AccountAIRunCfBaaiExecuteBgeM3ResponseResultBgeM3OuputQuery],
// [AccountAIRunCfBaaiExecuteBgeM3ResponseResultBgeM3OutputEmbeddingForContexts] or
// [AccountAIRunCfBaaiExecuteBgeM3ResponseResultBgeM3OuputEmbedding].
type AccountAIRunCfBaaiExecuteBgeM3ResponseResultUnion interface {
	implementsAccountAIRunCfBaaiExecuteBgeM3ResponseResult()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccountAIRunCfBaaiExecuteBgeM3ResponseResultUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccountAIRunCfBaaiExecuteBgeM3ResponseResultBgeM3OuputQuery{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccountAIRunCfBaaiExecuteBgeM3ResponseResultBgeM3OutputEmbeddingForContexts{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccountAIRunCfBaaiExecuteBgeM3ResponseResultBgeM3OuputEmbedding{}),
		},
	)
}

type AccountAIRunCfBaaiExecuteBgeM3ResponseResultBgeM3OuputQuery struct {
	Response []AccountAIRunCfBaaiExecuteBgeM3ResponseResultBgeM3OuputQueryResponse `json:"response"`
	JSON     accountAIRunCfBaaiExecuteBgeM3ResponseResultBgeM3OuputQueryJSON       `json:"-"`
}

// accountAIRunCfBaaiExecuteBgeM3ResponseResultBgeM3OuputQueryJSON contains the
// JSON metadata for the struct
// [AccountAIRunCfBaaiExecuteBgeM3ResponseResultBgeM3OuputQuery]
type accountAIRunCfBaaiExecuteBgeM3ResponseResultBgeM3OuputQueryJSON struct {
	Response    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIRunCfBaaiExecuteBgeM3ResponseResultBgeM3OuputQuery) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunCfBaaiExecuteBgeM3ResponseResultBgeM3OuputQueryJSON) RawJSON() string {
	return r.raw
}

func (r AccountAIRunCfBaaiExecuteBgeM3ResponseResultBgeM3OuputQuery) implementsAccountAIRunCfBaaiExecuteBgeM3ResponseResult() {
}

type AccountAIRunCfBaaiExecuteBgeM3ResponseResultBgeM3OuputQueryResponse struct {
	// Index of the context in the request
	ID int64 `json:"id"`
	// Score of the context under the index.
	Score float64                                                                 `json:"score"`
	JSON  accountAIRunCfBaaiExecuteBgeM3ResponseResultBgeM3OuputQueryResponseJSON `json:"-"`
}

// accountAIRunCfBaaiExecuteBgeM3ResponseResultBgeM3OuputQueryResponseJSON contains
// the JSON metadata for the struct
// [AccountAIRunCfBaaiExecuteBgeM3ResponseResultBgeM3OuputQueryResponse]
type accountAIRunCfBaaiExecuteBgeM3ResponseResultBgeM3OuputQueryResponseJSON struct {
	ID          apijson.Field
	Score       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIRunCfBaaiExecuteBgeM3ResponseResultBgeM3OuputQueryResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunCfBaaiExecuteBgeM3ResponseResultBgeM3OuputQueryResponseJSON) RawJSON() string {
	return r.raw
}

type AccountAIRunCfBaaiExecuteBgeM3ResponseResultBgeM3OutputEmbeddingForContexts struct {
	// The pooling method used in the embedding process.
	Pooling  AccountAIRunCfBaaiExecuteBgeM3ResponseResultBgeM3OutputEmbeddingForContextsPooling `json:"pooling"`
	Response [][]float64                                                                        `json:"response"`
	Shape    []float64                                                                          `json:"shape"`
	JSON     accountAIRunCfBaaiExecuteBgeM3ResponseResultBgeM3OutputEmbeddingForContextsJSON    `json:"-"`
}

// accountAIRunCfBaaiExecuteBgeM3ResponseResultBgeM3OutputEmbeddingForContextsJSON
// contains the JSON metadata for the struct
// [AccountAIRunCfBaaiExecuteBgeM3ResponseResultBgeM3OutputEmbeddingForContexts]
type accountAIRunCfBaaiExecuteBgeM3ResponseResultBgeM3OutputEmbeddingForContextsJSON struct {
	Pooling     apijson.Field
	Response    apijson.Field
	Shape       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIRunCfBaaiExecuteBgeM3ResponseResultBgeM3OutputEmbeddingForContexts) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunCfBaaiExecuteBgeM3ResponseResultBgeM3OutputEmbeddingForContextsJSON) RawJSON() string {
	return r.raw
}

func (r AccountAIRunCfBaaiExecuteBgeM3ResponseResultBgeM3OutputEmbeddingForContexts) implementsAccountAIRunCfBaaiExecuteBgeM3ResponseResult() {
}

// The pooling method used in the embedding process.
type AccountAIRunCfBaaiExecuteBgeM3ResponseResultBgeM3OutputEmbeddingForContextsPooling string

const (
	AccountAIRunCfBaaiExecuteBgeM3ResponseResultBgeM3OutputEmbeddingForContextsPoolingMean AccountAIRunCfBaaiExecuteBgeM3ResponseResultBgeM3OutputEmbeddingForContextsPooling = "mean"
	AccountAIRunCfBaaiExecuteBgeM3ResponseResultBgeM3OutputEmbeddingForContextsPoolingCls  AccountAIRunCfBaaiExecuteBgeM3ResponseResultBgeM3OutputEmbeddingForContextsPooling = "cls"
)

func (r AccountAIRunCfBaaiExecuteBgeM3ResponseResultBgeM3OutputEmbeddingForContextsPooling) IsKnown() bool {
	switch r {
	case AccountAIRunCfBaaiExecuteBgeM3ResponseResultBgeM3OutputEmbeddingForContextsPoolingMean, AccountAIRunCfBaaiExecuteBgeM3ResponseResultBgeM3OutputEmbeddingForContextsPoolingCls:
		return true
	}
	return false
}

type AccountAIRunCfBaaiExecuteBgeM3ResponseResultBgeM3OuputEmbedding struct {
	// Embeddings of the requested text values
	Data [][]float64 `json:"data"`
	// The pooling method used in the embedding process.
	Pooling AccountAIRunCfBaaiExecuteBgeM3ResponseResultBgeM3OuputEmbeddingPooling `json:"pooling"`
	Shape   []float64                                                              `json:"shape"`
	JSON    accountAIRunCfBaaiExecuteBgeM3ResponseResultBgeM3OuputEmbeddingJSON    `json:"-"`
}

// accountAIRunCfBaaiExecuteBgeM3ResponseResultBgeM3OuputEmbeddingJSON contains the
// JSON metadata for the struct
// [AccountAIRunCfBaaiExecuteBgeM3ResponseResultBgeM3OuputEmbedding]
type accountAIRunCfBaaiExecuteBgeM3ResponseResultBgeM3OuputEmbeddingJSON struct {
	Data        apijson.Field
	Pooling     apijson.Field
	Shape       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIRunCfBaaiExecuteBgeM3ResponseResultBgeM3OuputEmbedding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunCfBaaiExecuteBgeM3ResponseResultBgeM3OuputEmbeddingJSON) RawJSON() string {
	return r.raw
}

func (r AccountAIRunCfBaaiExecuteBgeM3ResponseResultBgeM3OuputEmbedding) implementsAccountAIRunCfBaaiExecuteBgeM3ResponseResult() {
}

// The pooling method used in the embedding process.
type AccountAIRunCfBaaiExecuteBgeM3ResponseResultBgeM3OuputEmbeddingPooling string

const (
	AccountAIRunCfBaaiExecuteBgeM3ResponseResultBgeM3OuputEmbeddingPoolingMean AccountAIRunCfBaaiExecuteBgeM3ResponseResultBgeM3OuputEmbeddingPooling = "mean"
	AccountAIRunCfBaaiExecuteBgeM3ResponseResultBgeM3OuputEmbeddingPoolingCls  AccountAIRunCfBaaiExecuteBgeM3ResponseResultBgeM3OuputEmbeddingPooling = "cls"
)

func (r AccountAIRunCfBaaiExecuteBgeM3ResponseResultBgeM3OuputEmbeddingPooling) IsKnown() bool {
	switch r {
	case AccountAIRunCfBaaiExecuteBgeM3ResponseResultBgeM3OuputEmbeddingPoolingMean, AccountAIRunCfBaaiExecuteBgeM3ResponseResultBgeM3OuputEmbeddingPoolingCls:
		return true
	}
	return false
}

// The pooling method used in the embedding process.
type AccountAIRunCfBaaiExecuteBgeM3ResponseResultPooling string

const (
	AccountAIRunCfBaaiExecuteBgeM3ResponseResultPoolingMean AccountAIRunCfBaaiExecuteBgeM3ResponseResultPooling = "mean"
	AccountAIRunCfBaaiExecuteBgeM3ResponseResultPoolingCls  AccountAIRunCfBaaiExecuteBgeM3ResponseResultPooling = "cls"
)

func (r AccountAIRunCfBaaiExecuteBgeM3ResponseResultPooling) IsKnown() bool {
	switch r {
	case AccountAIRunCfBaaiExecuteBgeM3ResponseResultPoolingMean, AccountAIRunCfBaaiExecuteBgeM3ResponseResultPoolingCls:
		return true
	}
	return false
}

type AccountAIRunCfBaaiExecuteBgeRerankerBaseResponse struct {
	Result  AccountAIRunCfBaaiExecuteBgeRerankerBaseResponseResult `json:"result"`
	Success bool                                                   `json:"success"`
	JSON    accountAIRunCfBaaiExecuteBgeRerankerBaseResponseJSON   `json:"-"`
}

// accountAIRunCfBaaiExecuteBgeRerankerBaseResponseJSON contains the JSON metadata
// for the struct [AccountAIRunCfBaaiExecuteBgeRerankerBaseResponse]
type accountAIRunCfBaaiExecuteBgeRerankerBaseResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIRunCfBaaiExecuteBgeRerankerBaseResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunCfBaaiExecuteBgeRerankerBaseResponseJSON) RawJSON() string {
	return r.raw
}

type AccountAIRunCfBaaiExecuteBgeRerankerBaseResponseResult struct {
	Response []AccountAIRunCfBaaiExecuteBgeRerankerBaseResponseResultResponse `json:"response"`
	JSON     accountAIRunCfBaaiExecuteBgeRerankerBaseResponseResultJSON       `json:"-"`
}

// accountAIRunCfBaaiExecuteBgeRerankerBaseResponseResultJSON contains the JSON
// metadata for the struct [AccountAIRunCfBaaiExecuteBgeRerankerBaseResponseResult]
type accountAIRunCfBaaiExecuteBgeRerankerBaseResponseResultJSON struct {
	Response    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIRunCfBaaiExecuteBgeRerankerBaseResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunCfBaaiExecuteBgeRerankerBaseResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountAIRunCfBaaiExecuteBgeRerankerBaseResponseResultResponse struct {
	// Index of the context in the request
	ID int64 `json:"id"`
	// Score of the context under the index.
	Score float64                                                            `json:"score"`
	JSON  accountAIRunCfBaaiExecuteBgeRerankerBaseResponseResultResponseJSON `json:"-"`
}

// accountAIRunCfBaaiExecuteBgeRerankerBaseResponseResultResponseJSON contains the
// JSON metadata for the struct
// [AccountAIRunCfBaaiExecuteBgeRerankerBaseResponseResultResponse]
type accountAIRunCfBaaiExecuteBgeRerankerBaseResponseResultResponseJSON struct {
	ID          apijson.Field
	Score       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIRunCfBaaiExecuteBgeRerankerBaseResponseResultResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunCfBaaiExecuteBgeRerankerBaseResponseResultResponseJSON) RawJSON() string {
	return r.raw
}

type AccountAIRunCfBaaiExecuteBgeSmallEnV1_5Response struct {
	Result  AccountAIRunCfBaaiExecuteBgeSmallEnV1_5ResponseResult `json:"result"`
	Success bool                                                  `json:"success"`
	JSON    accountAIRunCfBaaiExecuteBgeSmallEnV1_5ResponseJSON   `json:"-"`
}

// accountAIRunCfBaaiExecuteBgeSmallEnV1_5ResponseJSON contains the JSON metadata
// for the struct [AccountAIRunCfBaaiExecuteBgeSmallEnV1_5Response]
type accountAIRunCfBaaiExecuteBgeSmallEnV1_5ResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIRunCfBaaiExecuteBgeSmallEnV1_5Response) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunCfBaaiExecuteBgeSmallEnV1_5ResponseJSON) RawJSON() string {
	return r.raw
}

type AccountAIRunCfBaaiExecuteBgeSmallEnV1_5ResponseResult struct {
	// Embeddings of the requested text values
	Data  [][]float64                                               `json:"data"`
	Shape []float64                                                 `json:"shape"`
	JSON  accountAIRunCfBaaiExecuteBgeSmallEnV1_5ResponseResultJSON `json:"-"`
}

// accountAIRunCfBaaiExecuteBgeSmallEnV1_5ResponseResultJSON contains the JSON
// metadata for the struct [AccountAIRunCfBaaiExecuteBgeSmallEnV1_5ResponseResult]
type accountAIRunCfBaaiExecuteBgeSmallEnV1_5ResponseResultJSON struct {
	Data        apijson.Field
	Shape       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAIRunCfBaaiExecuteBgeSmallEnV1_5ResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAIRunCfBaaiExecuteBgeSmallEnV1_5ResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountAIRunCfBaaiExecuteBgeBaseEnV1_5Params struct {
	// The text to embed
	Text         param.Field[AccountAIRunCfBaaiExecuteBgeBaseEnV1_5ParamsTextUnion] `json:"text,required"`
	QueueRequest param.Field[string]                                                `query:"queueRequest"`
}

func (r AccountAIRunCfBaaiExecuteBgeBaseEnV1_5Params) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [AccountAIRunCfBaaiExecuteBgeBaseEnV1_5Params]'s query
// parameters as `url.Values`.
func (r AccountAIRunCfBaaiExecuteBgeBaseEnV1_5Params) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The text to embed
//
// Satisfied by [shared.UnionString],
// [AccountAIRunCfBaaiExecuteBgeBaseEnV1_5ParamsTextArray].
type AccountAIRunCfBaaiExecuteBgeBaseEnV1_5ParamsTextUnion interface {
	ImplementsAccountAIRunCfBaaiExecuteBgeBaseEnV1_5ParamsTextUnion()
}

type AccountAIRunCfBaaiExecuteBgeBaseEnV1_5ParamsTextArray []string

func (r AccountAIRunCfBaaiExecuteBgeBaseEnV1_5ParamsTextArray) ImplementsAccountAIRunCfBaaiExecuteBgeBaseEnV1_5ParamsTextUnion() {
}

type AccountAIRunCfBaaiExecuteBgeLargeEnV1_5Params struct {
	// The text to embed
	Text         param.Field[AccountAIRunCfBaaiExecuteBgeLargeEnV1_5ParamsTextUnion] `json:"text,required"`
	QueueRequest param.Field[string]                                                 `query:"queueRequest"`
}

func (r AccountAIRunCfBaaiExecuteBgeLargeEnV1_5Params) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [AccountAIRunCfBaaiExecuteBgeLargeEnV1_5Params]'s query
// parameters as `url.Values`.
func (r AccountAIRunCfBaaiExecuteBgeLargeEnV1_5Params) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The text to embed
//
// Satisfied by [shared.UnionString],
// [AccountAIRunCfBaaiExecuteBgeLargeEnV1_5ParamsTextArray].
type AccountAIRunCfBaaiExecuteBgeLargeEnV1_5ParamsTextUnion interface {
	ImplementsAccountAIRunCfBaaiExecuteBgeLargeEnV1_5ParamsTextUnion()
}

type AccountAIRunCfBaaiExecuteBgeLargeEnV1_5ParamsTextArray []string

func (r AccountAIRunCfBaaiExecuteBgeLargeEnV1_5ParamsTextArray) ImplementsAccountAIRunCfBaaiExecuteBgeLargeEnV1_5ParamsTextUnion() {
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
	Query param.Field[string]      `json:"query"`
	Text  param.Field[interface{}] `json:"text"`
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
	// The text to embed
	Text         param.Field[AccountAIRunCfBaaiExecuteBgeSmallEnV1_5ParamsTextUnion] `json:"text,required"`
	QueueRequest param.Field[string]                                                 `query:"queueRequest"`
}

func (r AccountAIRunCfBaaiExecuteBgeSmallEnV1_5Params) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [AccountAIRunCfBaaiExecuteBgeSmallEnV1_5Params]'s query
// parameters as `url.Values`.
func (r AccountAIRunCfBaaiExecuteBgeSmallEnV1_5Params) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The text to embed
//
// Satisfied by [shared.UnionString],
// [AccountAIRunCfBaaiExecuteBgeSmallEnV1_5ParamsTextArray].
type AccountAIRunCfBaaiExecuteBgeSmallEnV1_5ParamsTextUnion interface {
	ImplementsAccountAIRunCfBaaiExecuteBgeSmallEnV1_5ParamsTextUnion()
}

type AccountAIRunCfBaaiExecuteBgeSmallEnV1_5ParamsTextArray []string

func (r AccountAIRunCfBaaiExecuteBgeSmallEnV1_5ParamsTextArray) ImplementsAccountAIRunCfBaaiExecuteBgeSmallEnV1_5ParamsTextUnion() {
}
