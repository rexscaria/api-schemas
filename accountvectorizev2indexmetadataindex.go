// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
)

// AccountVectorizeV2IndexMetadataIndexService contains methods and other services
// that help with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountVectorizeV2IndexMetadataIndexService] method instead.
type AccountVectorizeV2IndexMetadataIndexService struct {
	Options []option.RequestOption
}

// NewAccountVectorizeV2IndexMetadataIndexService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewAccountVectorizeV2IndexMetadataIndexService(opts ...option.RequestOption) (r *AccountVectorizeV2IndexMetadataIndexService) {
	r = &AccountVectorizeV2IndexMetadataIndexService{}
	r.Options = opts
	return
}

// Enable metadata filtering based on metadata property. Limited to 10 properties.
func (r *AccountVectorizeV2IndexMetadataIndexService) New(ctx context.Context, accountID string, indexName string, body AccountVectorizeV2IndexMetadataIndexNewParams, opts ...option.RequestOption) (res *AccountVectorizeV2IndexMetadataIndexNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if indexName == "" {
		err = errors.New("missing required index_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/vectorize/v2/indexes/%s/metadata_index/create", accountID, indexName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List Metadata Indexes for the specified Vectorize Index.
func (r *AccountVectorizeV2IndexMetadataIndexService) List(ctx context.Context, accountID string, indexName string, opts ...option.RequestOption) (res *AccountVectorizeV2IndexMetadataIndexListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if indexName == "" {
		err = errors.New("missing required index_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/vectorize/v2/indexes/%s/metadata_index/list", accountID, indexName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Allow Vectorize to delete the specified metadata index.
func (r *AccountVectorizeV2IndexMetadataIndexService) Delete(ctx context.Context, accountID string, indexName string, body AccountVectorizeV2IndexMetadataIndexDeleteParams, opts ...option.RequestOption) (res *AccountVectorizeV2IndexMetadataIndexDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if indexName == "" {
		err = errors.New("missing required index_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/vectorize/v2/indexes/%s/metadata_index/delete", accountID, indexName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type AccountVectorizeV2IndexMetadataIndexNewResponse struct {
	Errors   []VectorizeMessages                                   `json:"errors,required"`
	Messages []VectorizeMessages                                   `json:"messages,required"`
	Result   AccountVectorizeV2IndexMetadataIndexNewResponseResult `json:"result,required,nullable"`
	// Whether the API call was successful
	Success AccountVectorizeV2IndexMetadataIndexNewResponseSuccess `json:"success,required"`
	JSON    accountVectorizeV2IndexMetadataIndexNewResponseJSON    `json:"-"`
}

// accountVectorizeV2IndexMetadataIndexNewResponseJSON contains the JSON metadata
// for the struct [AccountVectorizeV2IndexMetadataIndexNewResponse]
type accountVectorizeV2IndexMetadataIndexNewResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountVectorizeV2IndexMetadataIndexNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountVectorizeV2IndexMetadataIndexNewResponseJSON) RawJSON() string {
	return r.raw
}

type AccountVectorizeV2IndexMetadataIndexNewResponseResult struct {
	// The unique identifier for the async mutation operation containing the changeset.
	MutationID string                                                    `json:"mutationId"`
	JSON       accountVectorizeV2IndexMetadataIndexNewResponseResultJSON `json:"-"`
}

// accountVectorizeV2IndexMetadataIndexNewResponseResultJSON contains the JSON
// metadata for the struct [AccountVectorizeV2IndexMetadataIndexNewResponseResult]
type accountVectorizeV2IndexMetadataIndexNewResponseResultJSON struct {
	MutationID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountVectorizeV2IndexMetadataIndexNewResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountVectorizeV2IndexMetadataIndexNewResponseResultJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AccountVectorizeV2IndexMetadataIndexNewResponseSuccess bool

const (
	AccountVectorizeV2IndexMetadataIndexNewResponseSuccessTrue AccountVectorizeV2IndexMetadataIndexNewResponseSuccess = true
)

func (r AccountVectorizeV2IndexMetadataIndexNewResponseSuccess) IsKnown() bool {
	switch r {
	case AccountVectorizeV2IndexMetadataIndexNewResponseSuccessTrue:
		return true
	}
	return false
}

type AccountVectorizeV2IndexMetadataIndexListResponse struct {
	Errors   []VectorizeMessages                                    `json:"errors,required"`
	Messages []VectorizeMessages                                    `json:"messages,required"`
	Result   AccountVectorizeV2IndexMetadataIndexListResponseResult `json:"result,required,nullable"`
	// Whether the API call was successful
	Success AccountVectorizeV2IndexMetadataIndexListResponseSuccess `json:"success,required"`
	JSON    accountVectorizeV2IndexMetadataIndexListResponseJSON    `json:"-"`
}

// accountVectorizeV2IndexMetadataIndexListResponseJSON contains the JSON metadata
// for the struct [AccountVectorizeV2IndexMetadataIndexListResponse]
type accountVectorizeV2IndexMetadataIndexListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountVectorizeV2IndexMetadataIndexListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountVectorizeV2IndexMetadataIndexListResponseJSON) RawJSON() string {
	return r.raw
}

type AccountVectorizeV2IndexMetadataIndexListResponseResult struct {
	// Array of indexed metadata properties.
	MetadataIndexes []AccountVectorizeV2IndexMetadataIndexListResponseResultMetadataIndex `json:"metadataIndexes"`
	JSON            accountVectorizeV2IndexMetadataIndexListResponseResultJSON            `json:"-"`
}

// accountVectorizeV2IndexMetadataIndexListResponseResultJSON contains the JSON
// metadata for the struct [AccountVectorizeV2IndexMetadataIndexListResponseResult]
type accountVectorizeV2IndexMetadataIndexListResponseResultJSON struct {
	MetadataIndexes apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccountVectorizeV2IndexMetadataIndexListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountVectorizeV2IndexMetadataIndexListResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountVectorizeV2IndexMetadataIndexListResponseResultMetadataIndex struct {
	// Specifies the type of indexed metadata property.
	IndexType AccountVectorizeV2IndexMetadataIndexListResponseResultMetadataIndexesIndexType `json:"indexType"`
	// Specifies the indexed metadata property.
	PropertyName string                                                                  `json:"propertyName"`
	JSON         accountVectorizeV2IndexMetadataIndexListResponseResultMetadataIndexJSON `json:"-"`
}

// accountVectorizeV2IndexMetadataIndexListResponseResultMetadataIndexJSON contains
// the JSON metadata for the struct
// [AccountVectorizeV2IndexMetadataIndexListResponseResultMetadataIndex]
type accountVectorizeV2IndexMetadataIndexListResponseResultMetadataIndexJSON struct {
	IndexType    apijson.Field
	PropertyName apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountVectorizeV2IndexMetadataIndexListResponseResultMetadataIndex) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountVectorizeV2IndexMetadataIndexListResponseResultMetadataIndexJSON) RawJSON() string {
	return r.raw
}

// Specifies the type of indexed metadata property.
type AccountVectorizeV2IndexMetadataIndexListResponseResultMetadataIndexesIndexType string

const (
	AccountVectorizeV2IndexMetadataIndexListResponseResultMetadataIndexesIndexTypeString  AccountVectorizeV2IndexMetadataIndexListResponseResultMetadataIndexesIndexType = "string"
	AccountVectorizeV2IndexMetadataIndexListResponseResultMetadataIndexesIndexTypeNumber  AccountVectorizeV2IndexMetadataIndexListResponseResultMetadataIndexesIndexType = "number"
	AccountVectorizeV2IndexMetadataIndexListResponseResultMetadataIndexesIndexTypeBoolean AccountVectorizeV2IndexMetadataIndexListResponseResultMetadataIndexesIndexType = "boolean"
)

func (r AccountVectorizeV2IndexMetadataIndexListResponseResultMetadataIndexesIndexType) IsKnown() bool {
	switch r {
	case AccountVectorizeV2IndexMetadataIndexListResponseResultMetadataIndexesIndexTypeString, AccountVectorizeV2IndexMetadataIndexListResponseResultMetadataIndexesIndexTypeNumber, AccountVectorizeV2IndexMetadataIndexListResponseResultMetadataIndexesIndexTypeBoolean:
		return true
	}
	return false
}

// Whether the API call was successful
type AccountVectorizeV2IndexMetadataIndexListResponseSuccess bool

const (
	AccountVectorizeV2IndexMetadataIndexListResponseSuccessTrue AccountVectorizeV2IndexMetadataIndexListResponseSuccess = true
)

func (r AccountVectorizeV2IndexMetadataIndexListResponseSuccess) IsKnown() bool {
	switch r {
	case AccountVectorizeV2IndexMetadataIndexListResponseSuccessTrue:
		return true
	}
	return false
}

type AccountVectorizeV2IndexMetadataIndexDeleteResponse struct {
	Errors   []VectorizeMessages                                      `json:"errors,required"`
	Messages []VectorizeMessages                                      `json:"messages,required"`
	Result   AccountVectorizeV2IndexMetadataIndexDeleteResponseResult `json:"result,required,nullable"`
	// Whether the API call was successful
	Success AccountVectorizeV2IndexMetadataIndexDeleteResponseSuccess `json:"success,required"`
	JSON    accountVectorizeV2IndexMetadataIndexDeleteResponseJSON    `json:"-"`
}

// accountVectorizeV2IndexMetadataIndexDeleteResponseJSON contains the JSON
// metadata for the struct [AccountVectorizeV2IndexMetadataIndexDeleteResponse]
type accountVectorizeV2IndexMetadataIndexDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountVectorizeV2IndexMetadataIndexDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountVectorizeV2IndexMetadataIndexDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type AccountVectorizeV2IndexMetadataIndexDeleteResponseResult struct {
	// The unique identifier for the async mutation operation containing the changeset.
	MutationID string                                                       `json:"mutationId"`
	JSON       accountVectorizeV2IndexMetadataIndexDeleteResponseResultJSON `json:"-"`
}

// accountVectorizeV2IndexMetadataIndexDeleteResponseResultJSON contains the JSON
// metadata for the struct
// [AccountVectorizeV2IndexMetadataIndexDeleteResponseResult]
type accountVectorizeV2IndexMetadataIndexDeleteResponseResultJSON struct {
	MutationID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountVectorizeV2IndexMetadataIndexDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountVectorizeV2IndexMetadataIndexDeleteResponseResultJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AccountVectorizeV2IndexMetadataIndexDeleteResponseSuccess bool

const (
	AccountVectorizeV2IndexMetadataIndexDeleteResponseSuccessTrue AccountVectorizeV2IndexMetadataIndexDeleteResponseSuccess = true
)

func (r AccountVectorizeV2IndexMetadataIndexDeleteResponseSuccess) IsKnown() bool {
	switch r {
	case AccountVectorizeV2IndexMetadataIndexDeleteResponseSuccessTrue:
		return true
	}
	return false
}

type AccountVectorizeV2IndexMetadataIndexNewParams struct {
	// Specifies the type of metadata property to index.
	IndexType param.Field[AccountVectorizeV2IndexMetadataIndexNewParamsIndexType] `json:"indexType,required"`
	// Specifies the metadata property to index.
	PropertyName param.Field[string] `json:"propertyName,required"`
}

func (r AccountVectorizeV2IndexMetadataIndexNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Specifies the type of metadata property to index.
type AccountVectorizeV2IndexMetadataIndexNewParamsIndexType string

const (
	AccountVectorizeV2IndexMetadataIndexNewParamsIndexTypeString  AccountVectorizeV2IndexMetadataIndexNewParamsIndexType = "string"
	AccountVectorizeV2IndexMetadataIndexNewParamsIndexTypeNumber  AccountVectorizeV2IndexMetadataIndexNewParamsIndexType = "number"
	AccountVectorizeV2IndexMetadataIndexNewParamsIndexTypeBoolean AccountVectorizeV2IndexMetadataIndexNewParamsIndexType = "boolean"
)

func (r AccountVectorizeV2IndexMetadataIndexNewParamsIndexType) IsKnown() bool {
	switch r {
	case AccountVectorizeV2IndexMetadataIndexNewParamsIndexTypeString, AccountVectorizeV2IndexMetadataIndexNewParamsIndexTypeNumber, AccountVectorizeV2IndexMetadataIndexNewParamsIndexTypeBoolean:
		return true
	}
	return false
}

type AccountVectorizeV2IndexMetadataIndexDeleteParams struct {
	// Specifies the metadata property for which the index must be deleted.
	PropertyName param.Field[string] `json:"propertyName,required"`
}

func (r AccountVectorizeV2IndexMetadataIndexDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
