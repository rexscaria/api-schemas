// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/apiquery"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
)

// AccountVectorizeV2IndexService contains methods and other services that help
// with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountVectorizeV2IndexService] method instead.
type AccountVectorizeV2IndexService struct {
	Options       []option.RequestOption
	MetadataIndex *AccountVectorizeV2IndexMetadataIndexService
}

// NewAccountVectorizeV2IndexService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountVectorizeV2IndexService(opts ...option.RequestOption) (r *AccountVectorizeV2IndexService) {
	r = &AccountVectorizeV2IndexService{}
	r.Options = opts
	r.MetadataIndex = NewAccountVectorizeV2IndexMetadataIndexService(opts...)
	return
}

// Creates and returns a new Vectorize Index.
func (r *AccountVectorizeV2IndexService) New(ctx context.Context, accountID string, body AccountVectorizeV2IndexNewParams, opts ...option.RequestOption) (res *AccountVectorizeV2IndexNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/vectorize/v2/indexes", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Returns the specified Vectorize Index.
func (r *AccountVectorizeV2IndexService) Get(ctx context.Context, accountID string, indexName string, opts ...option.RequestOption) (res *AccountVectorizeV2IndexGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if indexName == "" {
		err = errors.New("missing required index_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/vectorize/v2/indexes/%s", accountID, indexName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Returns a list of Vectorize Indexes
func (r *AccountVectorizeV2IndexService) List(ctx context.Context, accountID string, opts ...option.RequestOption) (res *AccountVectorizeV2IndexListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/vectorize/v2/indexes", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Deletes the specified Vectorize Index.
func (r *AccountVectorizeV2IndexService) Delete(ctx context.Context, accountID string, indexName string, opts ...option.RequestOption) (res *AccountVectorizeV2IndexDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if indexName == "" {
		err = errors.New("missing required index_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/vectorize/v2/indexes/%s", accountID, indexName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Delete a set of vectors from an index by their vector identifiers.
func (r *AccountVectorizeV2IndexService) DeleteByIDs(ctx context.Context, accountID string, indexName string, body AccountVectorizeV2IndexDeleteByIDsParams, opts ...option.RequestOption) (res *AccountVectorizeV2IndexDeleteByIDsResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if indexName == "" {
		err = errors.New("missing required index_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/vectorize/v2/indexes/%s/delete_by_ids", accountID, indexName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get a set of vectors from an index by their vector identifiers.
func (r *AccountVectorizeV2IndexService) GetByIDs(ctx context.Context, accountID string, indexName string, body AccountVectorizeV2IndexGetByIDsParams, opts ...option.RequestOption) (res *AccountVectorizeV2IndexGetByIDsResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if indexName == "" {
		err = errors.New("missing required index_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/vectorize/v2/indexes/%s/get_by_ids", accountID, indexName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get information about a vectorize index.
func (r *AccountVectorizeV2IndexService) GetInfo(ctx context.Context, accountID string, indexName string, opts ...option.RequestOption) (res *AccountVectorizeV2IndexGetInfoResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if indexName == "" {
		err = errors.New("missing required index_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/vectorize/v2/indexes/%s/info", accountID, indexName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Inserts vectors into the specified index and returns a mutation id corresponding
// to the vectors enqueued for insertion.
func (r *AccountVectorizeV2IndexService) Insert(ctx context.Context, accountID string, indexName string, params AccountVectorizeV2IndexInsertParams, opts ...option.RequestOption) (res *AccountVectorizeV2IndexInsertResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if indexName == "" {
		err = errors.New("missing required index_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/vectorize/v2/indexes/%s/insert", accountID, indexName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Finds vectors closest to a given vector in an index.
func (r *AccountVectorizeV2IndexService) Query(ctx context.Context, accountID string, indexName string, body AccountVectorizeV2IndexQueryParams, opts ...option.RequestOption) (res *AccountVectorizeV2IndexQueryResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if indexName == "" {
		err = errors.New("missing required index_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/vectorize/v2/indexes/%s/query", accountID, indexName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Upserts vectors into the specified index, creating them if they do not exist and
// returns a mutation id corresponding to the vectors enqueued for upsertion.
func (r *AccountVectorizeV2IndexService) Upsert(ctx context.Context, accountID string, indexName string, params AccountVectorizeV2IndexUpsertParams, opts ...option.RequestOption) (res *AccountVectorizeV2IndexUpsertResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if indexName == "" {
		err = errors.New("missing required index_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/vectorize/v2/indexes/%s/upsert", accountID, indexName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type AccountVectorizeV2IndexNewResponse struct {
	Result VectorizeCreateIndexResponse           `json:"result"`
	JSON   accountVectorizeV2IndexNewResponseJSON `json:"-"`
	VectorizeAPIResponseSingle
}

// accountVectorizeV2IndexNewResponseJSON contains the JSON metadata for the struct
// [AccountVectorizeV2IndexNewResponse]
type accountVectorizeV2IndexNewResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountVectorizeV2IndexNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountVectorizeV2IndexNewResponseJSON) RawJSON() string {
	return r.raw
}

type AccountVectorizeV2IndexGetResponse struct {
	Result VectorizeCreateIndexResponse           `json:"result"`
	JSON   accountVectorizeV2IndexGetResponseJSON `json:"-"`
	VectorizeAPIResponseSingle
}

// accountVectorizeV2IndexGetResponseJSON contains the JSON metadata for the struct
// [AccountVectorizeV2IndexGetResponse]
type accountVectorizeV2IndexGetResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountVectorizeV2IndexGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountVectorizeV2IndexGetResponseJSON) RawJSON() string {
	return r.raw
}

type AccountVectorizeV2IndexListResponse struct {
	Result []VectorizeCreateIndexResponse          `json:"result"`
	JSON   accountVectorizeV2IndexListResponseJSON `json:"-"`
	VectorizeAPIResponseCommon
}

// accountVectorizeV2IndexListResponseJSON contains the JSON metadata for the
// struct [AccountVectorizeV2IndexListResponse]
type accountVectorizeV2IndexListResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountVectorizeV2IndexListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountVectorizeV2IndexListResponseJSON) RawJSON() string {
	return r.raw
}

type AccountVectorizeV2IndexDeleteResponse struct {
	Result interface{}                               `json:"result,nullable"`
	JSON   accountVectorizeV2IndexDeleteResponseJSON `json:"-"`
	VectorizeAPIResponseSingle
}

// accountVectorizeV2IndexDeleteResponseJSON contains the JSON metadata for the
// struct [AccountVectorizeV2IndexDeleteResponse]
type accountVectorizeV2IndexDeleteResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountVectorizeV2IndexDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountVectorizeV2IndexDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type AccountVectorizeV2IndexDeleteByIDsResponse struct {
	Result AccountVectorizeV2IndexDeleteByIDsResponseResult `json:"result"`
	JSON   accountVectorizeV2IndexDeleteByIDsResponseJSON   `json:"-"`
	VectorizeAPIResponseSingle
}

// accountVectorizeV2IndexDeleteByIDsResponseJSON contains the JSON metadata for
// the struct [AccountVectorizeV2IndexDeleteByIDsResponse]
type accountVectorizeV2IndexDeleteByIDsResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountVectorizeV2IndexDeleteByIDsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountVectorizeV2IndexDeleteByIDsResponseJSON) RawJSON() string {
	return r.raw
}

type AccountVectorizeV2IndexDeleteByIDsResponseResult struct {
	// The unique identifier for the async mutation operation containing the changeset.
	MutationID interface{}                                          `json:"mutationId"`
	JSON       accountVectorizeV2IndexDeleteByIDsResponseResultJSON `json:"-"`
}

// accountVectorizeV2IndexDeleteByIDsResponseResultJSON contains the JSON metadata
// for the struct [AccountVectorizeV2IndexDeleteByIDsResponseResult]
type accountVectorizeV2IndexDeleteByIDsResponseResultJSON struct {
	MutationID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountVectorizeV2IndexDeleteByIDsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountVectorizeV2IndexDeleteByIDsResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountVectorizeV2IndexGetByIDsResponse struct {
	// Array of vectors with matching ids.
	Result []VectorizeIndexGetVectorsByIDResponse      `json:"result"`
	JSON   accountVectorizeV2IndexGetByIDsResponseJSON `json:"-"`
	VectorizeAPIResponseSingle
}

// accountVectorizeV2IndexGetByIDsResponseJSON contains the JSON metadata for the
// struct [AccountVectorizeV2IndexGetByIDsResponse]
type accountVectorizeV2IndexGetByIDsResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountVectorizeV2IndexGetByIDsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountVectorizeV2IndexGetByIDsResponseJSON) RawJSON() string {
	return r.raw
}

type AccountVectorizeV2IndexGetInfoResponse struct {
	Result AccountVectorizeV2IndexGetInfoResponseResult `json:"result"`
	JSON   accountVectorizeV2IndexGetInfoResponseJSON   `json:"-"`
	VectorizeAPIResponseSingle
}

// accountVectorizeV2IndexGetInfoResponseJSON contains the JSON metadata for the
// struct [AccountVectorizeV2IndexGetInfoResponse]
type accountVectorizeV2IndexGetInfoResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountVectorizeV2IndexGetInfoResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountVectorizeV2IndexGetInfoResponseJSON) RawJSON() string {
	return r.raw
}

type AccountVectorizeV2IndexGetInfoResponseResult struct {
	// Specifies the number of dimensions for the index
	Dimensions int64 `json:"dimensions"`
	// Specifies the timestamp the last mutation batch was processed as an ISO8601
	// string.
	ProcessedUpToDatetime time.Time `json:"processedUpToDatetime,nullable" format:"date-time"`
	// The unique identifier for the async mutation operation containing the changeset.
	ProcessedUpToMutation interface{} `json:"processedUpToMutation"`
	// Specifies the number of vectors present in the index
	VectorCount int64                                            `json:"vectorCount"`
	JSON        accountVectorizeV2IndexGetInfoResponseResultJSON `json:"-"`
}

// accountVectorizeV2IndexGetInfoResponseResultJSON contains the JSON metadata for
// the struct [AccountVectorizeV2IndexGetInfoResponseResult]
type accountVectorizeV2IndexGetInfoResponseResultJSON struct {
	Dimensions            apijson.Field
	ProcessedUpToDatetime apijson.Field
	ProcessedUpToMutation apijson.Field
	VectorCount           apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *AccountVectorizeV2IndexGetInfoResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountVectorizeV2IndexGetInfoResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountVectorizeV2IndexInsertResponse struct {
	Result AccountVectorizeV2IndexInsertResponseResult `json:"result"`
	JSON   accountVectorizeV2IndexInsertResponseJSON   `json:"-"`
	VectorizeAPIResponseSingle
}

// accountVectorizeV2IndexInsertResponseJSON contains the JSON metadata for the
// struct [AccountVectorizeV2IndexInsertResponse]
type accountVectorizeV2IndexInsertResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountVectorizeV2IndexInsertResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountVectorizeV2IndexInsertResponseJSON) RawJSON() string {
	return r.raw
}

type AccountVectorizeV2IndexInsertResponseResult struct {
	// The unique identifier for the async mutation operation containing the changeset.
	MutationID interface{}                                     `json:"mutationId"`
	JSON       accountVectorizeV2IndexInsertResponseResultJSON `json:"-"`
}

// accountVectorizeV2IndexInsertResponseResultJSON contains the JSON metadata for
// the struct [AccountVectorizeV2IndexInsertResponseResult]
type accountVectorizeV2IndexInsertResponseResultJSON struct {
	MutationID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountVectorizeV2IndexInsertResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountVectorizeV2IndexInsertResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountVectorizeV2IndexQueryResponse struct {
	Result AccountVectorizeV2IndexQueryResponseResult `json:"result"`
	JSON   accountVectorizeV2IndexQueryResponseJSON   `json:"-"`
	VectorizeAPIResponseSingle
}

// accountVectorizeV2IndexQueryResponseJSON contains the JSON metadata for the
// struct [AccountVectorizeV2IndexQueryResponse]
type accountVectorizeV2IndexQueryResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountVectorizeV2IndexQueryResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountVectorizeV2IndexQueryResponseJSON) RawJSON() string {
	return r.raw
}

type AccountVectorizeV2IndexQueryResponseResult struct {
	// Specifies the count of vectors returned by the search
	Count int64 `json:"count"`
	// Array of vectors matched by the search
	Matches []AccountVectorizeV2IndexQueryResponseResultMatch `json:"matches"`
	JSON    accountVectorizeV2IndexQueryResponseResultJSON    `json:"-"`
}

// accountVectorizeV2IndexQueryResponseResultJSON contains the JSON metadata for
// the struct [AccountVectorizeV2IndexQueryResponseResult]
type accountVectorizeV2IndexQueryResponseResultJSON struct {
	Count       apijson.Field
	Matches     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountVectorizeV2IndexQueryResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountVectorizeV2IndexQueryResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountVectorizeV2IndexQueryResponseResultMatch struct {
	// Identifier for a Vector
	ID        string      `json:"id"`
	Metadata  interface{} `json:"metadata,nullable"`
	Namespace string      `json:"namespace,nullable"`
	// The score of the vector according to the index's distance metric
	Score  float64                                             `json:"score"`
	Values []float64                                           `json:"values,nullable"`
	JSON   accountVectorizeV2IndexQueryResponseResultMatchJSON `json:"-"`
}

// accountVectorizeV2IndexQueryResponseResultMatchJSON contains the JSON metadata
// for the struct [AccountVectorizeV2IndexQueryResponseResultMatch]
type accountVectorizeV2IndexQueryResponseResultMatchJSON struct {
	ID          apijson.Field
	Metadata    apijson.Field
	Namespace   apijson.Field
	Score       apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountVectorizeV2IndexQueryResponseResultMatch) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountVectorizeV2IndexQueryResponseResultMatchJSON) RawJSON() string {
	return r.raw
}

type AccountVectorizeV2IndexUpsertResponse struct {
	Result AccountVectorizeV2IndexUpsertResponseResult `json:"result"`
	JSON   accountVectorizeV2IndexUpsertResponseJSON   `json:"-"`
	VectorizeAPIResponseSingle
}

// accountVectorizeV2IndexUpsertResponseJSON contains the JSON metadata for the
// struct [AccountVectorizeV2IndexUpsertResponse]
type accountVectorizeV2IndexUpsertResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountVectorizeV2IndexUpsertResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountVectorizeV2IndexUpsertResponseJSON) RawJSON() string {
	return r.raw
}

type AccountVectorizeV2IndexUpsertResponseResult struct {
	// The unique identifier for the async mutation operation containing the changeset.
	MutationID interface{}                                     `json:"mutationId"`
	JSON       accountVectorizeV2IndexUpsertResponseResultJSON `json:"-"`
}

// accountVectorizeV2IndexUpsertResponseResultJSON contains the JSON metadata for
// the struct [AccountVectorizeV2IndexUpsertResponseResult]
type accountVectorizeV2IndexUpsertResponseResultJSON struct {
	MutationID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountVectorizeV2IndexUpsertResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountVectorizeV2IndexUpsertResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountVectorizeV2IndexNewParams struct {
	VectorizeCreateIndexRequest VectorizeCreateIndexRequestParam `json:"vectorize_create_index_request,required"`
}

func (r AccountVectorizeV2IndexNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.VectorizeCreateIndexRequest)
}

type AccountVectorizeV2IndexDeleteByIDsParams struct {
	VectorizeIndexDeleteVectorsByIDRequest VectorizeIndexDeleteVectorsByIDRequestParam `json:"vectorize_index_delete_vectors_by_id_request,required"`
}

func (r AccountVectorizeV2IndexDeleteByIDsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.VectorizeIndexDeleteVectorsByIDRequest)
}

type AccountVectorizeV2IndexGetByIDsParams struct {
	VectorizeIndexGetVectorsByIDRequest VectorizeIndexGetVectorsByIDRequestParam `json:"vectorize_index_get_vectors_by_id_request,required"`
}

func (r AccountVectorizeV2IndexGetByIDsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.VectorizeIndexGetVectorsByIDRequest)
}

type AccountVectorizeV2IndexInsertParams struct {
	// ndjson file containing vectors to insert.
	Body string `json:"body,required"`
	// Behavior for ndjson parse failures.
	UnparsableBehavior param.Field[AccountVectorizeV2IndexInsertParamsUnparsableBehavior] `query:"unparsable-behavior"`
}

func (r AccountVectorizeV2IndexInsertParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

// URLQuery serializes [AccountVectorizeV2IndexInsertParams]'s query parameters as
// `url.Values`.
func (r AccountVectorizeV2IndexInsertParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Behavior for ndjson parse failures.
type AccountVectorizeV2IndexInsertParamsUnparsableBehavior string

const (
	AccountVectorizeV2IndexInsertParamsUnparsableBehaviorError   AccountVectorizeV2IndexInsertParamsUnparsableBehavior = "error"
	AccountVectorizeV2IndexInsertParamsUnparsableBehaviorDiscard AccountVectorizeV2IndexInsertParamsUnparsableBehavior = "discard"
)

func (r AccountVectorizeV2IndexInsertParamsUnparsableBehavior) IsKnown() bool {
	switch r {
	case AccountVectorizeV2IndexInsertParamsUnparsableBehaviorError, AccountVectorizeV2IndexInsertParamsUnparsableBehaviorDiscard:
		return true
	}
	return false
}

type AccountVectorizeV2IndexQueryParams struct {
	// The search vector that will be used to find the nearest neighbors.
	Vector param.Field[[]float64] `json:"vector,required"`
	// A metadata filter expression used to limit nearest neighbor results.
	Filter param.Field[interface{}] `json:"filter"`
	// Whether to return no metadata, indexed metadata or all metadata associated with
	// the closest vectors.
	ReturnMetadata param.Field[AccountVectorizeV2IndexQueryParamsReturnMetadata] `json:"returnMetadata"`
	// Whether to return the values associated with the closest vectors.
	ReturnValues param.Field[bool] `json:"returnValues"`
	// The number of nearest neighbors to find.
	TopK param.Field[float64] `json:"topK"`
}

func (r AccountVectorizeV2IndexQueryParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Whether to return no metadata, indexed metadata or all metadata associated with
// the closest vectors.
type AccountVectorizeV2IndexQueryParamsReturnMetadata string

const (
	AccountVectorizeV2IndexQueryParamsReturnMetadataNone    AccountVectorizeV2IndexQueryParamsReturnMetadata = "none"
	AccountVectorizeV2IndexQueryParamsReturnMetadataIndexed AccountVectorizeV2IndexQueryParamsReturnMetadata = "indexed"
	AccountVectorizeV2IndexQueryParamsReturnMetadataAll     AccountVectorizeV2IndexQueryParamsReturnMetadata = "all"
)

func (r AccountVectorizeV2IndexQueryParamsReturnMetadata) IsKnown() bool {
	switch r {
	case AccountVectorizeV2IndexQueryParamsReturnMetadataNone, AccountVectorizeV2IndexQueryParamsReturnMetadataIndexed, AccountVectorizeV2IndexQueryParamsReturnMetadataAll:
		return true
	}
	return false
}

type AccountVectorizeV2IndexUpsertParams struct {
	// ndjson file containing vectors to upsert.
	Body string `json:"body,required"`
	// Behavior for ndjson parse failures.
	UnparsableBehavior param.Field[AccountVectorizeV2IndexUpsertParamsUnparsableBehavior] `query:"unparsable-behavior"`
}

func (r AccountVectorizeV2IndexUpsertParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

// URLQuery serializes [AccountVectorizeV2IndexUpsertParams]'s query parameters as
// `url.Values`.
func (r AccountVectorizeV2IndexUpsertParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Behavior for ndjson parse failures.
type AccountVectorizeV2IndexUpsertParamsUnparsableBehavior string

const (
	AccountVectorizeV2IndexUpsertParamsUnparsableBehaviorError   AccountVectorizeV2IndexUpsertParamsUnparsableBehavior = "error"
	AccountVectorizeV2IndexUpsertParamsUnparsableBehaviorDiscard AccountVectorizeV2IndexUpsertParamsUnparsableBehavior = "discard"
)

func (r AccountVectorizeV2IndexUpsertParamsUnparsableBehavior) IsKnown() bool {
	switch r {
	case AccountVectorizeV2IndexUpsertParamsUnparsableBehaviorError, AccountVectorizeV2IndexUpsertParamsUnparsableBehaviorDiscard:
		return true
	}
	return false
}
