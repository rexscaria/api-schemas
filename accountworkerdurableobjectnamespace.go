// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/stainless-sdks/cf-rex-go/internal/apijson"
	"github.com/stainless-sdks/cf-rex-go/internal/apiquery"
	"github.com/stainless-sdks/cf-rex-go/internal/param"
	"github.com/stainless-sdks/cf-rex-go/internal/requestconfig"
	"github.com/stainless-sdks/cf-rex-go/option"
)

// AccountWorkerDurableObjectNamespaceService contains methods and other services
// that help with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountWorkerDurableObjectNamespaceService] method instead.
type AccountWorkerDurableObjectNamespaceService struct {
	Options []option.RequestOption
}

// NewAccountWorkerDurableObjectNamespaceService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewAccountWorkerDurableObjectNamespaceService(opts ...option.RequestOption) (r *AccountWorkerDurableObjectNamespaceService) {
	r = &AccountWorkerDurableObjectNamespaceService{}
	r.Options = opts
	return
}

// Returns the Durable Object namespaces owned by an account.
func (r *AccountWorkerDurableObjectNamespaceService) List(ctx context.Context, accountID string, opts ...option.RequestOption) (res *AccountWorkerDurableObjectNamespaceListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/durable_objects/namespaces", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Returns the Durable Objects in a given namespace.
func (r *AccountWorkerDurableObjectNamespaceService) ListObjects(ctx context.Context, accountID string, id string, query AccountWorkerDurableObjectNamespaceListObjectsParams, opts ...option.RequestOption) (res *AccountWorkerDurableObjectNamespaceListObjectsResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/durable_objects/namespaces/%s/objects", accountID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type Collection struct {
	ResultInfo CollectionResultInfo `json:"result_info"`
	JSON       collectionJSON       `json:"-"`
	CommonResponseWorkers
}

// collectionJSON contains the JSON metadata for the struct [Collection]
type collectionJSON struct {
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Collection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r collectionJSON) RawJSON() string {
	return r.raw
}

type CollectionResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                  `json:"total_count"`
	JSON       collectionResultInfoJSON `json:"-"`
}

// collectionResultInfoJSON contains the JSON metadata for the struct
// [CollectionResultInfo]
type collectionResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CollectionResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r collectionResultInfoJSON) RawJSON() string {
	return r.raw
}

type AccountWorkerDurableObjectNamespaceListResponse struct {
	Result []AccountWorkerDurableObjectNamespaceListResponseResult `json:"result"`
	JSON   accountWorkerDurableObjectNamespaceListResponseJSON     `json:"-"`
	Collection
}

// accountWorkerDurableObjectNamespaceListResponseJSON contains the JSON metadata
// for the struct [AccountWorkerDurableObjectNamespaceListResponse]
type accountWorkerDurableObjectNamespaceListResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerDurableObjectNamespaceListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountWorkerDurableObjectNamespaceListResponseJSON) RawJSON() string {
	return r.raw
}

type AccountWorkerDurableObjectNamespaceListResponseResult struct {
	ID        string                                                    `json:"id"`
	Class     string                                                    `json:"class"`
	Name      string                                                    `json:"name"`
	Script    string                                                    `json:"script"`
	UseSqlite bool                                                      `json:"use_sqlite"`
	JSON      accountWorkerDurableObjectNamespaceListResponseResultJSON `json:"-"`
}

// accountWorkerDurableObjectNamespaceListResponseResultJSON contains the JSON
// metadata for the struct [AccountWorkerDurableObjectNamespaceListResponseResult]
type accountWorkerDurableObjectNamespaceListResponseResultJSON struct {
	ID          apijson.Field
	Class       apijson.Field
	Name        apijson.Field
	Script      apijson.Field
	UseSqlite   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerDurableObjectNamespaceListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountWorkerDurableObjectNamespaceListResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountWorkerDurableObjectNamespaceListObjectsResponse struct {
	Result     []AccountWorkerDurableObjectNamespaceListObjectsResponseResult   `json:"result"`
	ResultInfo AccountWorkerDurableObjectNamespaceListObjectsResponseResultInfo `json:"result_info"`
	JSON       accountWorkerDurableObjectNamespaceListObjectsResponseJSON       `json:"-"`
	Collection
}

// accountWorkerDurableObjectNamespaceListObjectsResponseJSON contains the JSON
// metadata for the struct [AccountWorkerDurableObjectNamespaceListObjectsResponse]
type accountWorkerDurableObjectNamespaceListObjectsResponseJSON struct {
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerDurableObjectNamespaceListObjectsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountWorkerDurableObjectNamespaceListObjectsResponseJSON) RawJSON() string {
	return r.raw
}

type AccountWorkerDurableObjectNamespaceListObjectsResponseResult struct {
	// ID of the Durable Object.
	ID string `json:"id"`
	// Whether the Durable Object has stored data.
	HasStoredData bool                                                             `json:"hasStoredData"`
	JSON          accountWorkerDurableObjectNamespaceListObjectsResponseResultJSON `json:"-"`
}

// accountWorkerDurableObjectNamespaceListObjectsResponseResultJSON contains the
// JSON metadata for the struct
// [AccountWorkerDurableObjectNamespaceListObjectsResponseResult]
type accountWorkerDurableObjectNamespaceListObjectsResponseResultJSON struct {
	ID            apijson.Field
	HasStoredData apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccountWorkerDurableObjectNamespaceListObjectsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountWorkerDurableObjectNamespaceListObjectsResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountWorkerDurableObjectNamespaceListObjectsResponseResultInfo struct {
	// Total results returned based on your list parameters.
	Count float64 `json:"count"`
	// Opaque token indicating the position from which to continue when requesting the
	// next set of records. A valid value for the cursor can be obtained from the
	// cursors object in the result_info structure.
	Cursor string                                                               `json:"cursor"`
	JSON   accountWorkerDurableObjectNamespaceListObjectsResponseResultInfoJSON `json:"-"`
}

// accountWorkerDurableObjectNamespaceListObjectsResponseResultInfoJSON contains
// the JSON metadata for the struct
// [AccountWorkerDurableObjectNamespaceListObjectsResponseResultInfo]
type accountWorkerDurableObjectNamespaceListObjectsResponseResultInfoJSON struct {
	Count       apijson.Field
	Cursor      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerDurableObjectNamespaceListObjectsResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountWorkerDurableObjectNamespaceListObjectsResponseResultInfoJSON) RawJSON() string {
	return r.raw
}

type AccountWorkerDurableObjectNamespaceListObjectsParams struct {
	// Opaque token indicating the position from which to continue when requesting the
	// next set of records. A valid value for the cursor can be obtained from the
	// cursors object in the result_info structure.
	Cursor param.Field[string] `query:"cursor"`
	// The number of objects to return. The cursor attribute may be used to iterate
	// over the next batch of objects if there are more than the limit.
	Limit param.Field[float64] `query:"limit"`
}

// URLQuery serializes [AccountWorkerDurableObjectNamespaceListObjectsParams]'s
// query parameters as `url.Values`.
func (r AccountWorkerDurableObjectNamespaceListObjectsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
