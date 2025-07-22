// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"reflect"
	"time"

	"github.com/stainless-sdks/cf-rex-go/internal/apiform"
	"github.com/stainless-sdks/cf-rex-go/internal/apijson"
	"github.com/stainless-sdks/cf-rex-go/internal/param"
	"github.com/stainless-sdks/cf-rex-go/internal/requestconfig"
	"github.com/stainless-sdks/cf-rex-go/option"
	"github.com/stainless-sdks/cf-rex-go/shared"
	"github.com/tidwall/gjson"
)

// AccountVectorizeIndexService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountVectorizeIndexService] method instead.
type AccountVectorizeIndexService struct {
	Options []option.RequestOption
}

// NewAccountVectorizeIndexService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountVectorizeIndexService(opts ...option.RequestOption) (r *AccountVectorizeIndexService) {
	r = &AccountVectorizeIndexService{}
	r.Options = opts
	return
}

// Creates and returns a new Vectorize Index.
func (r *AccountVectorizeIndexService) New(ctx context.Context, accountID string, body AccountVectorizeIndexNewParams, opts ...option.RequestOption) (res *AccountVectorizeIndexNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/vectorize/indexes", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Returns the specified Vectorize Index.
func (r *AccountVectorizeIndexService) Get(ctx context.Context, accountID string, indexName string, opts ...option.RequestOption) (res *AccountVectorizeIndexGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if indexName == "" {
		err = errors.New("missing required index_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/vectorize/indexes/%s", accountID, indexName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates and returns the specified Vectorize Index.
func (r *AccountVectorizeIndexService) Update(ctx context.Context, accountID string, indexName string, body AccountVectorizeIndexUpdateParams, opts ...option.RequestOption) (res *AccountVectorizeIndexUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if indexName == "" {
		err = errors.New("missing required index_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/vectorize/indexes/%s", accountID, indexName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Returns a list of Vectorize Indexes
func (r *AccountVectorizeIndexService) List(ctx context.Context, accountID string, opts ...option.RequestOption) (res *AccountVectorizeIndexListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/vectorize/indexes", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Deletes the specified Vectorize Index.
func (r *AccountVectorizeIndexService) Delete(ctx context.Context, accountID string, indexName string, opts ...option.RequestOption) (res *AccountVectorizeIndexDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if indexName == "" {
		err = errors.New("missing required index_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/vectorize/indexes/%s", accountID, indexName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Delete a set of vectors from an index by their vector identifiers.
func (r *AccountVectorizeIndexService) DeleteByIDs(ctx context.Context, accountID string, indexName string, body AccountVectorizeIndexDeleteByIDsParams, opts ...option.RequestOption) (res *AccountVectorizeIndexDeleteByIDsResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if indexName == "" {
		err = errors.New("missing required index_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/vectorize/indexes/%s/delete-by-ids", accountID, indexName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get a set of vectors from an index by their vector identifiers.
func (r *AccountVectorizeIndexService) GetByIDs(ctx context.Context, accountID string, indexName string, body AccountVectorizeIndexGetByIDsParams, opts ...option.RequestOption) (res *AccountVectorizeIndexGetByIDsResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if indexName == "" {
		err = errors.New("missing required index_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/vectorize/indexes/%s/get-by-ids", accountID, indexName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Inserts vectors into the specified index and returns the count of the vectors
// successfully inserted.
func (r *AccountVectorizeIndexService) Insert(ctx context.Context, accountID string, indexName string, body AccountVectorizeIndexInsertParams, opts ...option.RequestOption) (res *AccountVectorizeIndexInsertResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if indexName == "" {
		err = errors.New("missing required index_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/vectorize/indexes/%s/insert", accountID, indexName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Finds vectors closest to a given vector in an index.
func (r *AccountVectorizeIndexService) Query(ctx context.Context, accountID string, indexName string, body AccountVectorizeIndexQueryParams, opts ...option.RequestOption) (res *AccountVectorizeIndexQueryResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if indexName == "" {
		err = errors.New("missing required index_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/vectorize/indexes/%s/query", accountID, indexName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Upserts vectors into the specified index, creating them if they do not exist and
// returns the count of values and ids successfully inserted.
func (r *AccountVectorizeIndexService) Upsert(ctx context.Context, accountID string, indexName string, body AccountVectorizeIndexUpsertParams, opts ...option.RequestOption) (res *AccountVectorizeIndexUpsertResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if indexName == "" {
		err = errors.New("missing required index_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/vectorize/indexes/%s/upsert", accountID, indexName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type VectorizeAPIResponseCommon struct {
	Errors   []VectorizeMessages                   `json:"errors,required"`
	Messages []VectorizeMessages                   `json:"messages,required"`
	Result   VectorizeAPIResponseCommonResultUnion `json:"result,required"`
	// Whether the API call was successful
	Success VectorizeAPIResponseCommonSuccess `json:"success,required"`
	JSON    vectorizeAPIResponseCommonJSON    `json:"-"`
}

// vectorizeAPIResponseCommonJSON contains the JSON metadata for the struct
// [VectorizeAPIResponseCommon]
type vectorizeAPIResponseCommonJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VectorizeAPIResponseCommon) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r vectorizeAPIResponseCommonJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [VectorizeAPIResponseCommonResultArray] or
// [shared.UnionString].
type VectorizeAPIResponseCommonResultUnion interface {
	ImplementsVectorizeAPIResponseCommonResultUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*VectorizeAPIResponseCommonResultUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(VectorizeAPIResponseCommonResultArray{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type VectorizeAPIResponseCommonResultArray []interface{}

func (r VectorizeAPIResponseCommonResultArray) ImplementsVectorizeAPIResponseCommonResultUnion() {}

// Whether the API call was successful
type VectorizeAPIResponseCommonSuccess bool

const (
	VectorizeAPIResponseCommonSuccessTrue VectorizeAPIResponseCommonSuccess = true
)

func (r VectorizeAPIResponseCommonSuccess) IsKnown() bool {
	switch r {
	case VectorizeAPIResponseCommonSuccessTrue:
		return true
	}
	return false
}

type VectorizeAPIResponseSingle struct {
	Result interface{}                    `json:"result,nullable"`
	JSON   vectorizeAPIResponseSingleJSON `json:"-"`
	VectorizeAPIResponseCommon
}

// vectorizeAPIResponseSingleJSON contains the JSON metadata for the struct
// [VectorizeAPIResponseSingle]
type vectorizeAPIResponseSingleJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VectorizeAPIResponseSingle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r vectorizeAPIResponseSingleJSON) RawJSON() string {
	return r.raw
}

type VectorizeCreateIndexRequestParam struct {
	// Specifies the type of configuration to use for the index.
	Config param.Field[VectorizeCreateIndexRequestConfigUnionParam] `json:"config,required"`
	Name   param.Field[string]                                      `json:"name,required"`
	// Specifies the description of the index.
	Description param.Field[string] `json:"description"`
}

func (r VectorizeCreateIndexRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Specifies the type of configuration to use for the index.
type VectorizeCreateIndexRequestConfigParam struct {
	// Specifies the number of dimensions for the index
	Dimensions param.Field[int64] `json:"dimensions"`
	// Specifies the type of metric to use calculating distance.
	Metric param.Field[VectorizeCreateIndexRequestConfigMetric] `json:"metric"`
	// Specifies the preset to use for the index.
	Preset param.Field[VectorizeCreateIndexRequestConfigPreset] `json:"preset"`
}

func (r VectorizeCreateIndexRequestConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r VectorizeCreateIndexRequestConfigParam) implementsVectorizeCreateIndexRequestConfigUnionParam() {
}

// Specifies the type of configuration to use for the index.
//
// Satisfied by [VectorizeIndexDimensionConfigurationParam],
// [VectorizeCreateIndexRequestConfigVectorizeIndexPresetConfigurationParam],
// [VectorizeCreateIndexRequestConfigParam].
type VectorizeCreateIndexRequestConfigUnionParam interface {
	implementsVectorizeCreateIndexRequestConfigUnionParam()
}

type VectorizeCreateIndexRequestConfigVectorizeIndexPresetConfigurationParam struct {
	// Specifies the preset to use for the index.
	Preset param.Field[VectorizeCreateIndexRequestConfigVectorizeIndexPresetConfigurationPreset] `json:"preset,required"`
}

func (r VectorizeCreateIndexRequestConfigVectorizeIndexPresetConfigurationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r VectorizeCreateIndexRequestConfigVectorizeIndexPresetConfigurationParam) implementsVectorizeCreateIndexRequestConfigUnionParam() {
}

// Specifies the preset to use for the index.
type VectorizeCreateIndexRequestConfigVectorizeIndexPresetConfigurationPreset string

const (
	VectorizeCreateIndexRequestConfigVectorizeIndexPresetConfigurationPresetCfBaaiBgeSmallEnV1_5        VectorizeCreateIndexRequestConfigVectorizeIndexPresetConfigurationPreset = "@cf/baai/bge-small-en-v1.5"
	VectorizeCreateIndexRequestConfigVectorizeIndexPresetConfigurationPresetCfBaaiBgeBaseEnV1_5         VectorizeCreateIndexRequestConfigVectorizeIndexPresetConfigurationPreset = "@cf/baai/bge-base-en-v1.5"
	VectorizeCreateIndexRequestConfigVectorizeIndexPresetConfigurationPresetCfBaaiBgeLargeEnV1_5        VectorizeCreateIndexRequestConfigVectorizeIndexPresetConfigurationPreset = "@cf/baai/bge-large-en-v1.5"
	VectorizeCreateIndexRequestConfigVectorizeIndexPresetConfigurationPresetOpenAITextEmbeddingAda002   VectorizeCreateIndexRequestConfigVectorizeIndexPresetConfigurationPreset = "openai/text-embedding-ada-002"
	VectorizeCreateIndexRequestConfigVectorizeIndexPresetConfigurationPresetCohereEmbedMultilingualV2_0 VectorizeCreateIndexRequestConfigVectorizeIndexPresetConfigurationPreset = "cohere/embed-multilingual-v2.0"
)

func (r VectorizeCreateIndexRequestConfigVectorizeIndexPresetConfigurationPreset) IsKnown() bool {
	switch r {
	case VectorizeCreateIndexRequestConfigVectorizeIndexPresetConfigurationPresetCfBaaiBgeSmallEnV1_5, VectorizeCreateIndexRequestConfigVectorizeIndexPresetConfigurationPresetCfBaaiBgeBaseEnV1_5, VectorizeCreateIndexRequestConfigVectorizeIndexPresetConfigurationPresetCfBaaiBgeLargeEnV1_5, VectorizeCreateIndexRequestConfigVectorizeIndexPresetConfigurationPresetOpenAITextEmbeddingAda002, VectorizeCreateIndexRequestConfigVectorizeIndexPresetConfigurationPresetCohereEmbedMultilingualV2_0:
		return true
	}
	return false
}

// Specifies the type of metric to use calculating distance.
type VectorizeCreateIndexRequestConfigMetric string

const (
	VectorizeCreateIndexRequestConfigMetricCosine     VectorizeCreateIndexRequestConfigMetric = "cosine"
	VectorizeCreateIndexRequestConfigMetricEuclidean  VectorizeCreateIndexRequestConfigMetric = "euclidean"
	VectorizeCreateIndexRequestConfigMetricDotProduct VectorizeCreateIndexRequestConfigMetric = "dot-product"
)

func (r VectorizeCreateIndexRequestConfigMetric) IsKnown() bool {
	switch r {
	case VectorizeCreateIndexRequestConfigMetricCosine, VectorizeCreateIndexRequestConfigMetricEuclidean, VectorizeCreateIndexRequestConfigMetricDotProduct:
		return true
	}
	return false
}

// Specifies the preset to use for the index.
type VectorizeCreateIndexRequestConfigPreset string

const (
	VectorizeCreateIndexRequestConfigPresetCfBaaiBgeSmallEnV1_5        VectorizeCreateIndexRequestConfigPreset = "@cf/baai/bge-small-en-v1.5"
	VectorizeCreateIndexRequestConfigPresetCfBaaiBgeBaseEnV1_5         VectorizeCreateIndexRequestConfigPreset = "@cf/baai/bge-base-en-v1.5"
	VectorizeCreateIndexRequestConfigPresetCfBaaiBgeLargeEnV1_5        VectorizeCreateIndexRequestConfigPreset = "@cf/baai/bge-large-en-v1.5"
	VectorizeCreateIndexRequestConfigPresetOpenAITextEmbeddingAda002   VectorizeCreateIndexRequestConfigPreset = "openai/text-embedding-ada-002"
	VectorizeCreateIndexRequestConfigPresetCohereEmbedMultilingualV2_0 VectorizeCreateIndexRequestConfigPreset = "cohere/embed-multilingual-v2.0"
)

func (r VectorizeCreateIndexRequestConfigPreset) IsKnown() bool {
	switch r {
	case VectorizeCreateIndexRequestConfigPresetCfBaaiBgeSmallEnV1_5, VectorizeCreateIndexRequestConfigPresetCfBaaiBgeBaseEnV1_5, VectorizeCreateIndexRequestConfigPresetCfBaaiBgeLargeEnV1_5, VectorizeCreateIndexRequestConfigPresetOpenAITextEmbeddingAda002, VectorizeCreateIndexRequestConfigPresetCohereEmbedMultilingualV2_0:
		return true
	}
	return false
}

type VectorizeCreateIndexResponse struct {
	Config VectorizeIndexDimensionConfiguration `json:"config"`
	// Specifies the timestamp the resource was created as an ISO8601 string.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// Specifies the description of the index.
	Description string `json:"description"`
	// Specifies the timestamp the resource was modified as an ISO8601 string.
	ModifiedOn time.Time                        `json:"modified_on" format:"date-time"`
	Name       string                           `json:"name"`
	JSON       vectorizeCreateIndexResponseJSON `json:"-"`
}

// vectorizeCreateIndexResponseJSON contains the JSON metadata for the struct
// [VectorizeCreateIndexResponse]
type vectorizeCreateIndexResponseJSON struct {
	Config      apijson.Field
	CreatedOn   apijson.Field
	Description apijson.Field
	ModifiedOn  apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VectorizeCreateIndexResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r vectorizeCreateIndexResponseJSON) RawJSON() string {
	return r.raw
}

type VectorizeIndexDeleteVectorsByIDRequestParam struct {
	// A list of vector identifiers to delete from the index indicated by the path.
	IDs param.Field[[]string] `json:"ids"`
}

func (r VectorizeIndexDeleteVectorsByIDRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type VectorizeIndexDimensionConfiguration struct {
	// Specifies the number of dimensions for the index
	Dimensions int64 `json:"dimensions,required"`
	// Specifies the type of metric to use calculating distance.
	Metric VectorizeIndexDimensionConfigurationMetric `json:"metric,required"`
	JSON   vectorizeIndexDimensionConfigurationJSON   `json:"-"`
}

// vectorizeIndexDimensionConfigurationJSON contains the JSON metadata for the
// struct [VectorizeIndexDimensionConfiguration]
type vectorizeIndexDimensionConfigurationJSON struct {
	Dimensions  apijson.Field
	Metric      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VectorizeIndexDimensionConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r vectorizeIndexDimensionConfigurationJSON) RawJSON() string {
	return r.raw
}

// Specifies the type of metric to use calculating distance.
type VectorizeIndexDimensionConfigurationMetric string

const (
	VectorizeIndexDimensionConfigurationMetricCosine     VectorizeIndexDimensionConfigurationMetric = "cosine"
	VectorizeIndexDimensionConfigurationMetricEuclidean  VectorizeIndexDimensionConfigurationMetric = "euclidean"
	VectorizeIndexDimensionConfigurationMetricDotProduct VectorizeIndexDimensionConfigurationMetric = "dot-product"
)

func (r VectorizeIndexDimensionConfigurationMetric) IsKnown() bool {
	switch r {
	case VectorizeIndexDimensionConfigurationMetricCosine, VectorizeIndexDimensionConfigurationMetricEuclidean, VectorizeIndexDimensionConfigurationMetricDotProduct:
		return true
	}
	return false
}

type VectorizeIndexDimensionConfigurationParam struct {
	// Specifies the number of dimensions for the index
	Dimensions param.Field[int64] `json:"dimensions,required"`
	// Specifies the type of metric to use calculating distance.
	Metric param.Field[VectorizeIndexDimensionConfigurationMetric] `json:"metric,required"`
}

func (r VectorizeIndexDimensionConfigurationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r VectorizeIndexDimensionConfigurationParam) implementsVectorizeCreateIndexRequestConfigUnionParam() {
}

type VectorizeIndexGetVectorsByIDRequestParam struct {
	// A list of vector identifiers to retrieve from the index indicated by the path.
	IDs param.Field[[]string] `json:"ids"`
}

func (r VectorizeIndexGetVectorsByIDRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type VectorizeIndexGetVectorsByIDResponse struct {
	// Identifier for a Vector
	ID        string                                   `json:"id"`
	Metadata  interface{}                              `json:"metadata"`
	Namespace string                                   `json:"namespace,nullable"`
	Values    []float64                                `json:"values"`
	JSON      vectorizeIndexGetVectorsByIDResponseJSON `json:"-"`
}

// vectorizeIndexGetVectorsByIDResponseJSON contains the JSON metadata for the
// struct [VectorizeIndexGetVectorsByIDResponse]
type vectorizeIndexGetVectorsByIDResponseJSON struct {
	ID          apijson.Field
	Metadata    apijson.Field
	Namespace   apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VectorizeIndexGetVectorsByIDResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r vectorizeIndexGetVectorsByIDResponseJSON) RawJSON() string {
	return r.raw
}

type VectorizeMessages struct {
	Code    int64                 `json:"code,required"`
	Message string                `json:"message,required"`
	JSON    vectorizeMessagesJSON `json:"-"`
}

// vectorizeMessagesJSON contains the JSON metadata for the struct
// [VectorizeMessages]
type vectorizeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VectorizeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r vectorizeMessagesJSON) RawJSON() string {
	return r.raw
}

type AccountVectorizeIndexNewResponse struct {
	Result VectorizeCreateIndexResponse         `json:"result"`
	JSON   accountVectorizeIndexNewResponseJSON `json:"-"`
	VectorizeAPIResponseSingle
}

// accountVectorizeIndexNewResponseJSON contains the JSON metadata for the struct
// [AccountVectorizeIndexNewResponse]
type accountVectorizeIndexNewResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountVectorizeIndexNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountVectorizeIndexNewResponseJSON) RawJSON() string {
	return r.raw
}

type AccountVectorizeIndexGetResponse struct {
	Result VectorizeCreateIndexResponse         `json:"result"`
	JSON   accountVectorizeIndexGetResponseJSON `json:"-"`
	VectorizeAPIResponseSingle
}

// accountVectorizeIndexGetResponseJSON contains the JSON metadata for the struct
// [AccountVectorizeIndexGetResponse]
type accountVectorizeIndexGetResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountVectorizeIndexGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountVectorizeIndexGetResponseJSON) RawJSON() string {
	return r.raw
}

type AccountVectorizeIndexUpdateResponse struct {
	Result VectorizeCreateIndexResponse            `json:"result"`
	JSON   accountVectorizeIndexUpdateResponseJSON `json:"-"`
	VectorizeAPIResponseSingle
}

// accountVectorizeIndexUpdateResponseJSON contains the JSON metadata for the
// struct [AccountVectorizeIndexUpdateResponse]
type accountVectorizeIndexUpdateResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountVectorizeIndexUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountVectorizeIndexUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type AccountVectorizeIndexListResponse struct {
	Result []VectorizeCreateIndexResponse        `json:"result"`
	JSON   accountVectorizeIndexListResponseJSON `json:"-"`
	VectorizeAPIResponseCommon
}

// accountVectorizeIndexListResponseJSON contains the JSON metadata for the struct
// [AccountVectorizeIndexListResponse]
type accountVectorizeIndexListResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountVectorizeIndexListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountVectorizeIndexListResponseJSON) RawJSON() string {
	return r.raw
}

type AccountVectorizeIndexDeleteResponse struct {
	Result interface{}                             `json:"result,nullable"`
	JSON   accountVectorizeIndexDeleteResponseJSON `json:"-"`
	VectorizeAPIResponseSingle
}

// accountVectorizeIndexDeleteResponseJSON contains the JSON metadata for the
// struct [AccountVectorizeIndexDeleteResponse]
type accountVectorizeIndexDeleteResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountVectorizeIndexDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountVectorizeIndexDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type AccountVectorizeIndexDeleteByIDsResponse struct {
	Result AccountVectorizeIndexDeleteByIDsResponseResult `json:"result"`
	JSON   accountVectorizeIndexDeleteByIDsResponseJSON   `json:"-"`
	VectorizeAPIResponseSingle
}

// accountVectorizeIndexDeleteByIDsResponseJSON contains the JSON metadata for the
// struct [AccountVectorizeIndexDeleteByIDsResponse]
type accountVectorizeIndexDeleteByIDsResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountVectorizeIndexDeleteByIDsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountVectorizeIndexDeleteByIDsResponseJSON) RawJSON() string {
	return r.raw
}

type AccountVectorizeIndexDeleteByIDsResponseResult struct {
	// The count of the vectors successfully deleted.
	Count int64 `json:"count"`
	// Array of vector identifiers of the vectors that were successfully processed for
	// deletion.
	IDs  []string                                           `json:"ids"`
	JSON accountVectorizeIndexDeleteByIDsResponseResultJSON `json:"-"`
}

// accountVectorizeIndexDeleteByIDsResponseResultJSON contains the JSON metadata
// for the struct [AccountVectorizeIndexDeleteByIDsResponseResult]
type accountVectorizeIndexDeleteByIDsResponseResultJSON struct {
	Count       apijson.Field
	IDs         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountVectorizeIndexDeleteByIDsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountVectorizeIndexDeleteByIDsResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountVectorizeIndexGetByIDsResponse struct {
	// Array of vectors with matching ids.
	Result []VectorizeIndexGetVectorsByIDResponse    `json:"result"`
	JSON   accountVectorizeIndexGetByIDsResponseJSON `json:"-"`
	VectorizeAPIResponseSingle
}

// accountVectorizeIndexGetByIDsResponseJSON contains the JSON metadata for the
// struct [AccountVectorizeIndexGetByIDsResponse]
type accountVectorizeIndexGetByIDsResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountVectorizeIndexGetByIDsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountVectorizeIndexGetByIDsResponseJSON) RawJSON() string {
	return r.raw
}

type AccountVectorizeIndexInsertResponse struct {
	Result AccountVectorizeIndexInsertResponseResult `json:"result"`
	JSON   accountVectorizeIndexInsertResponseJSON   `json:"-"`
	VectorizeAPIResponseSingle
}

// accountVectorizeIndexInsertResponseJSON contains the JSON metadata for the
// struct [AccountVectorizeIndexInsertResponse]
type accountVectorizeIndexInsertResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountVectorizeIndexInsertResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountVectorizeIndexInsertResponseJSON) RawJSON() string {
	return r.raw
}

type AccountVectorizeIndexInsertResponseResult struct {
	// Specifies the count of the vectors successfully inserted.
	Count int64 `json:"count"`
	// Array of vector identifiers of the vectors successfully inserted.
	IDs  []string                                      `json:"ids"`
	JSON accountVectorizeIndexInsertResponseResultJSON `json:"-"`
}

// accountVectorizeIndexInsertResponseResultJSON contains the JSON metadata for the
// struct [AccountVectorizeIndexInsertResponseResult]
type accountVectorizeIndexInsertResponseResultJSON struct {
	Count       apijson.Field
	IDs         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountVectorizeIndexInsertResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountVectorizeIndexInsertResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountVectorizeIndexQueryResponse struct {
	Result AccountVectorizeIndexQueryResponseResult `json:"result"`
	JSON   accountVectorizeIndexQueryResponseJSON   `json:"-"`
	VectorizeAPIResponseSingle
}

// accountVectorizeIndexQueryResponseJSON contains the JSON metadata for the struct
// [AccountVectorizeIndexQueryResponse]
type accountVectorizeIndexQueryResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountVectorizeIndexQueryResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountVectorizeIndexQueryResponseJSON) RawJSON() string {
	return r.raw
}

type AccountVectorizeIndexQueryResponseResult struct {
	// Specifies the count of vectors returned by the search
	Count int64 `json:"count"`
	// Array of vectors matched by the search
	Matches []AccountVectorizeIndexQueryResponseResultMatch `json:"matches"`
	JSON    accountVectorizeIndexQueryResponseResultJSON    `json:"-"`
}

// accountVectorizeIndexQueryResponseResultJSON contains the JSON metadata for the
// struct [AccountVectorizeIndexQueryResponseResult]
type accountVectorizeIndexQueryResponseResultJSON struct {
	Count       apijson.Field
	Matches     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountVectorizeIndexQueryResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountVectorizeIndexQueryResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountVectorizeIndexQueryResponseResultMatch struct {
	// Identifier for a Vector
	ID       string      `json:"id"`
	Metadata interface{} `json:"metadata,nullable"`
	// The score of the vector according to the index's distance metric
	Score  float64                                           `json:"score"`
	Values []float64                                         `json:"values,nullable"`
	JSON   accountVectorizeIndexQueryResponseResultMatchJSON `json:"-"`
}

// accountVectorizeIndexQueryResponseResultMatchJSON contains the JSON metadata for
// the struct [AccountVectorizeIndexQueryResponseResultMatch]
type accountVectorizeIndexQueryResponseResultMatchJSON struct {
	ID          apijson.Field
	Metadata    apijson.Field
	Score       apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountVectorizeIndexQueryResponseResultMatch) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountVectorizeIndexQueryResponseResultMatchJSON) RawJSON() string {
	return r.raw
}

type AccountVectorizeIndexUpsertResponse struct {
	Result AccountVectorizeIndexUpsertResponseResult `json:"result"`
	JSON   accountVectorizeIndexUpsertResponseJSON   `json:"-"`
	VectorizeAPIResponseSingle
}

// accountVectorizeIndexUpsertResponseJSON contains the JSON metadata for the
// struct [AccountVectorizeIndexUpsertResponse]
type accountVectorizeIndexUpsertResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountVectorizeIndexUpsertResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountVectorizeIndexUpsertResponseJSON) RawJSON() string {
	return r.raw
}

type AccountVectorizeIndexUpsertResponseResult struct {
	// Specifies the count of the vectors successfully inserted.
	Count int64 `json:"count"`
	// Array of vector identifiers of the vectors successfully inserted.
	IDs  []string                                      `json:"ids"`
	JSON accountVectorizeIndexUpsertResponseResultJSON `json:"-"`
}

// accountVectorizeIndexUpsertResponseResultJSON contains the JSON metadata for the
// struct [AccountVectorizeIndexUpsertResponseResult]
type accountVectorizeIndexUpsertResponseResultJSON struct {
	Count       apijson.Field
	IDs         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountVectorizeIndexUpsertResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountVectorizeIndexUpsertResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountVectorizeIndexNewParams struct {
	VectorizeCreateIndexRequest VectorizeCreateIndexRequestParam `json:"vectorize_create_index_request,required"`
}

func (r AccountVectorizeIndexNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.VectorizeCreateIndexRequest)
}

type AccountVectorizeIndexUpdateParams struct {
	// Specifies the description of the index.
	Description param.Field[string] `json:"description,required"`
}

func (r AccountVectorizeIndexUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountVectorizeIndexDeleteByIDsParams struct {
	VectorizeIndexDeleteVectorsByIDRequest VectorizeIndexDeleteVectorsByIDRequestParam `json:"vectorize_index_delete_vectors_by_id_request,required"`
}

func (r AccountVectorizeIndexDeleteByIDsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.VectorizeIndexDeleteVectorsByIDRequest)
}

type AccountVectorizeIndexGetByIDsParams struct {
	VectorizeIndexGetVectorsByIDRequest VectorizeIndexGetVectorsByIDRequestParam `json:"vectorize_index_get_vectors_by_id_request,required"`
}

func (r AccountVectorizeIndexGetByIDsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.VectorizeIndexGetVectorsByIDRequest)
}

type AccountVectorizeIndexInsertParams struct {
	// ndjson file containing vectors to insert.
	Body io.Reader `json:"body,required" format:"binary"`
}

func (r AccountVectorizeIndexInsertParams) MarshalMultipart() (data []byte, contentType string, err error) {
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

type AccountVectorizeIndexQueryParams struct {
	// The search vector that will be used to find the nearest neighbors.
	Vector param.Field[[]float64] `json:"vector,required"`
	// A metadata filter expression used to limit nearest neighbor results.
	Filter param.Field[interface{}] `json:"filter"`
	// Whether to return the metadata associated with the closest vectors.
	ReturnMetadata param.Field[bool] `json:"returnMetadata"`
	// Whether to return the values associated with the closest vectors.
	ReturnValues param.Field[bool] `json:"returnValues"`
	// The number of nearest neighbors to find.
	TopK param.Field[float64] `json:"topK"`
}

func (r AccountVectorizeIndexQueryParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountVectorizeIndexUpsertParams struct {
	// ndjson file containing vectors to upsert.
	Body io.Reader `json:"body,required" format:"binary"`
}

func (r AccountVectorizeIndexUpsertParams) MarshalMultipart() (data []byte, contentType string, err error) {
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
