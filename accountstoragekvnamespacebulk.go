// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
	"github.com/rexscaria/api-schemas/shared"
	"github.com/tidwall/gjson"
)

// AccountStorageKvNamespaceBulkService contains methods and other services that
// help with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountStorageKvNamespaceBulkService] method instead.
type AccountStorageKvNamespaceBulkService struct {
	Options []option.RequestOption
}

// NewAccountStorageKvNamespaceBulkService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountStorageKvNamespaceBulkService(opts ...option.RequestOption) (r *AccountStorageKvNamespaceBulkService) {
	r = &AccountStorageKvNamespaceBulkService{}
	r.Options = opts
	return
}

// Remove multiple KV pairs from the namespace. Body should be an array of up to
// 10,000 keys to be removed.
//
// Deprecated: deprecated
func (r *AccountStorageKvNamespaceBulkService) Delete(ctx context.Context, accountID string, namespaceID string, opts ...option.RequestOption) (res *AccountStorageKvNamespaceBulkDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if namespaceID == "" {
		err = errors.New("missing required namespace_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/storage/kv/namespaces/%s/bulk", accountID, namespaceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Remove multiple KV pairs from the namespace. Body should be an array of up to
// 10,000 keys to be removed.
func (r *AccountStorageKvNamespaceBulkService) DeleteMultiple(ctx context.Context, accountID string, namespaceID string, body AccountStorageKvNamespaceBulkDeleteMultipleParams, opts ...option.RequestOption) (res *AccountStorageKvNamespaceBulkDeleteMultipleResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if namespaceID == "" {
		err = errors.New("missing required namespace_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/storage/kv/namespaces/%s/bulk/delete", accountID, namespaceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve up to 100 KV pairs from the namespace. Keys must contain text-based
// values. JSON values can optionally be parsed instead of being returned as a
// string value. Metadata can be included if `withMetadata` is true.
func (r *AccountStorageKvNamespaceBulkService) GetMultiple(ctx context.Context, accountID string, namespaceID string, body AccountStorageKvNamespaceBulkGetMultipleParams, opts ...option.RequestOption) (res *AccountStorageKvNamespaceBulkGetMultipleResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if namespaceID == "" {
		err = errors.New("missing required namespace_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/storage/kv/namespaces/%s/bulk/get", accountID, namespaceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Write multiple keys and values at once. Body should be an array of up to 10,000
// key-value pairs to be stored, along with optional expiration information.
// Existing values and expirations will be overwritten. If neither `expiration` nor
// `expiration_ttl` is specified, the key-value pair will never expire. If both are
// set, `expiration_ttl` is used and `expiration` is ignored. The entire request
// size must be 100 megabytes or less.
func (r *AccountStorageKvNamespaceBulkService) Write(ctx context.Context, accountID string, namespaceID string, body AccountStorageKvNamespaceBulkWriteParams, opts ...option.RequestOption) (res *AccountStorageKvNamespaceBulkWriteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if namespaceID == "" {
		err = errors.New("missing required namespace_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/storage/kv/namespaces/%s/bulk", accountID, namespaceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type BulkResult struct {
	// Number of keys successfully updated.
	SuccessfulKeyCount float64 `json:"successful_key_count"`
	// Name of the keys that failed to be fully updated. They should be retried.
	UnsuccessfulKeys []string       `json:"unsuccessful_keys"`
	JSON             bulkResultJSON `json:"-"`
}

// bulkResultJSON contains the JSON metadata for the struct [BulkResult]
type bulkResultJSON struct {
	SuccessfulKeyCount apijson.Field
	UnsuccessfulKeys   apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *BulkResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bulkResultJSON) RawJSON() string {
	return r.raw
}

type AccountStorageKvNamespaceBulkDeleteResponse struct {
	Errors   []MessagesWorkersKvItem `json:"errors,required"`
	Messages []MessagesWorkersKvItem `json:"messages,required"`
	// Whether the API call was successful.
	Success AccountStorageKvNamespaceBulkDeleteResponseSuccess `json:"success,required"`
	Result  BulkResult                                         `json:"result,nullable"`
	JSON    accountStorageKvNamespaceBulkDeleteResponseJSON    `json:"-"`
}

// accountStorageKvNamespaceBulkDeleteResponseJSON contains the JSON metadata for
// the struct [AccountStorageKvNamespaceBulkDeleteResponse]
type accountStorageKvNamespaceBulkDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStorageKvNamespaceBulkDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountStorageKvNamespaceBulkDeleteResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccountStorageKvNamespaceBulkDeleteResponseSuccess bool

const (
	AccountStorageKvNamespaceBulkDeleteResponseSuccessTrue AccountStorageKvNamespaceBulkDeleteResponseSuccess = true
)

func (r AccountStorageKvNamespaceBulkDeleteResponseSuccess) IsKnown() bool {
	switch r {
	case AccountStorageKvNamespaceBulkDeleteResponseSuccessTrue:
		return true
	}
	return false
}

type AccountStorageKvNamespaceBulkDeleteMultipleResponse struct {
	Errors   []MessagesWorkersKvItem `json:"errors,required"`
	Messages []MessagesWorkersKvItem `json:"messages,required"`
	// Whether the API call was successful.
	Success AccountStorageKvNamespaceBulkDeleteMultipleResponseSuccess `json:"success,required"`
	Result  BulkResult                                                 `json:"result,nullable"`
	JSON    accountStorageKvNamespaceBulkDeleteMultipleResponseJSON    `json:"-"`
}

// accountStorageKvNamespaceBulkDeleteMultipleResponseJSON contains the JSON
// metadata for the struct [AccountStorageKvNamespaceBulkDeleteMultipleResponse]
type accountStorageKvNamespaceBulkDeleteMultipleResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStorageKvNamespaceBulkDeleteMultipleResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountStorageKvNamespaceBulkDeleteMultipleResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccountStorageKvNamespaceBulkDeleteMultipleResponseSuccess bool

const (
	AccountStorageKvNamespaceBulkDeleteMultipleResponseSuccessTrue AccountStorageKvNamespaceBulkDeleteMultipleResponseSuccess = true
)

func (r AccountStorageKvNamespaceBulkDeleteMultipleResponseSuccess) IsKnown() bool {
	switch r {
	case AccountStorageKvNamespaceBulkDeleteMultipleResponseSuccessTrue:
		return true
	}
	return false
}

type AccountStorageKvNamespaceBulkGetMultipleResponse struct {
	Errors   []MessagesWorkersKvItem `json:"errors,required"`
	Messages []MessagesWorkersKvItem `json:"messages,required"`
	// Whether the API call was successful.
	Success AccountStorageKvNamespaceBulkGetMultipleResponseSuccess `json:"success,required"`
	Result  AccountStorageKvNamespaceBulkGetMultipleResponseResult  `json:"result,nullable"`
	JSON    accountStorageKvNamespaceBulkGetMultipleResponseJSON    `json:"-"`
}

// accountStorageKvNamespaceBulkGetMultipleResponseJSON contains the JSON metadata
// for the struct [AccountStorageKvNamespaceBulkGetMultipleResponse]
type accountStorageKvNamespaceBulkGetMultipleResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStorageKvNamespaceBulkGetMultipleResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountStorageKvNamespaceBulkGetMultipleResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccountStorageKvNamespaceBulkGetMultipleResponseSuccess bool

const (
	AccountStorageKvNamespaceBulkGetMultipleResponseSuccessTrue AccountStorageKvNamespaceBulkGetMultipleResponseSuccess = true
)

func (r AccountStorageKvNamespaceBulkGetMultipleResponseSuccess) IsKnown() bool {
	switch r {
	case AccountStorageKvNamespaceBulkGetMultipleResponseSuccessTrue:
		return true
	}
	return false
}

type AccountStorageKvNamespaceBulkGetMultipleResponseResult struct {
	// This field can have the runtime type of
	// [map[string]AccountStorageKvNamespaceBulkGetMultipleResponseResultWorkersKvBulkGetResultValuesUnion],
	// [map[string]AccountStorageKvNamespaceBulkGetMultipleResponseResultWorkersKvBulkGetResultWithMetadataValue].
	Values interface{}                                                `json:"values"`
	JSON   accountStorageKvNamespaceBulkGetMultipleResponseResultJSON `json:"-"`
	union  AccountStorageKvNamespaceBulkGetMultipleResponseResultUnion
}

// accountStorageKvNamespaceBulkGetMultipleResponseResultJSON contains the JSON
// metadata for the struct [AccountStorageKvNamespaceBulkGetMultipleResponseResult]
type accountStorageKvNamespaceBulkGetMultipleResponseResultJSON struct {
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r accountStorageKvNamespaceBulkGetMultipleResponseResultJSON) RawJSON() string {
	return r.raw
}

func (r *AccountStorageKvNamespaceBulkGetMultipleResponseResult) UnmarshalJSON(data []byte) (err error) {
	*r = AccountStorageKvNamespaceBulkGetMultipleResponseResult{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [AccountStorageKvNamespaceBulkGetMultipleResponseResultUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [AccountStorageKvNamespaceBulkGetMultipleResponseResultWorkersKvBulkGetResult],
// [AccountStorageKvNamespaceBulkGetMultipleResponseResultWorkersKvBulkGetResultWithMetadata].
func (r AccountStorageKvNamespaceBulkGetMultipleResponseResult) AsUnion() AccountStorageKvNamespaceBulkGetMultipleResponseResultUnion {
	return r.union
}

// Union satisfied by
// [AccountStorageKvNamespaceBulkGetMultipleResponseResultWorkersKvBulkGetResult]
// or
// [AccountStorageKvNamespaceBulkGetMultipleResponseResultWorkersKvBulkGetResultWithMetadata].
type AccountStorageKvNamespaceBulkGetMultipleResponseResultUnion interface {
	implementsAccountStorageKvNamespaceBulkGetMultipleResponseResult()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccountStorageKvNamespaceBulkGetMultipleResponseResultUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccountStorageKvNamespaceBulkGetMultipleResponseResultWorkersKvBulkGetResult{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccountStorageKvNamespaceBulkGetMultipleResponseResultWorkersKvBulkGetResultWithMetadata{}),
		},
	)
}

type AccountStorageKvNamespaceBulkGetMultipleResponseResultWorkersKvBulkGetResult struct {
	// Requested keys are paired with their values in an object.
	Values map[string]AccountStorageKvNamespaceBulkGetMultipleResponseResultWorkersKvBulkGetResultValuesUnion `json:"values"`
	JSON   accountStorageKvNamespaceBulkGetMultipleResponseResultWorkersKvBulkGetResultJSON                   `json:"-"`
}

// accountStorageKvNamespaceBulkGetMultipleResponseResultWorkersKvBulkGetResultJSON
// contains the JSON metadata for the struct
// [AccountStorageKvNamespaceBulkGetMultipleResponseResultWorkersKvBulkGetResult]
type accountStorageKvNamespaceBulkGetMultipleResponseResultWorkersKvBulkGetResultJSON struct {
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStorageKvNamespaceBulkGetMultipleResponseResultWorkersKvBulkGetResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountStorageKvNamespaceBulkGetMultipleResponseResultWorkersKvBulkGetResultJSON) RawJSON() string {
	return r.raw
}

func (r AccountStorageKvNamespaceBulkGetMultipleResponseResultWorkersKvBulkGetResult) implementsAccountStorageKvNamespaceBulkGetMultipleResponseResult() {
}

// The value associated with the key.
//
// Union satisfied by [shared.UnionString], [shared.UnionFloat], [shared.UnionBool]
// or
// [AccountStorageKvNamespaceBulkGetMultipleResponseResultWorkersKvBulkGetResultValuesMap].
type AccountStorageKvNamespaceBulkGetMultipleResponseResultWorkersKvBulkGetResultValuesUnion interface {
	ImplementsAccountStorageKvNamespaceBulkGetMultipleResponseResultWorkersKvBulkGetResultValuesUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccountStorageKvNamespaceBulkGetMultipleResponseResultWorkersKvBulkGetResultValuesUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccountStorageKvNamespaceBulkGetMultipleResponseResultWorkersKvBulkGetResultValuesMap{}),
		},
	)
}

type AccountStorageKvNamespaceBulkGetMultipleResponseResultWorkersKvBulkGetResultValuesMap map[string]interface{}

func (r AccountStorageKvNamespaceBulkGetMultipleResponseResultWorkersKvBulkGetResultValuesMap) ImplementsAccountStorageKvNamespaceBulkGetMultipleResponseResultWorkersKvBulkGetResultValuesUnion() {
}

type AccountStorageKvNamespaceBulkGetMultipleResponseResultWorkersKvBulkGetResultWithMetadata struct {
	// Requested keys are paired with their values and metadata in an object.
	Values map[string]AccountStorageKvNamespaceBulkGetMultipleResponseResultWorkersKvBulkGetResultWithMetadataValue `json:"values"`
	JSON   accountStorageKvNamespaceBulkGetMultipleResponseResultWorkersKvBulkGetResultWithMetadataJSON             `json:"-"`
}

// accountStorageKvNamespaceBulkGetMultipleResponseResultWorkersKvBulkGetResultWithMetadataJSON
// contains the JSON metadata for the struct
// [AccountStorageKvNamespaceBulkGetMultipleResponseResultWorkersKvBulkGetResultWithMetadata]
type accountStorageKvNamespaceBulkGetMultipleResponseResultWorkersKvBulkGetResultWithMetadataJSON struct {
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStorageKvNamespaceBulkGetMultipleResponseResultWorkersKvBulkGetResultWithMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountStorageKvNamespaceBulkGetMultipleResponseResultWorkersKvBulkGetResultWithMetadataJSON) RawJSON() string {
	return r.raw
}

func (r AccountStorageKvNamespaceBulkGetMultipleResponseResultWorkersKvBulkGetResultWithMetadata) implementsAccountStorageKvNamespaceBulkGetMultipleResponseResult() {
}

type AccountStorageKvNamespaceBulkGetMultipleResponseResultWorkersKvBulkGetResultWithMetadataValue struct {
	Metadata interface{} `json:"metadata,required"`
	Value    interface{} `json:"value,required"`
	// Expires the key at a certain time, measured in number of seconds since the UNIX
	// epoch.
	Expiration float64                                                                                           `json:"expiration"`
	JSON       accountStorageKvNamespaceBulkGetMultipleResponseResultWorkersKvBulkGetResultWithMetadataValueJSON `json:"-"`
}

// accountStorageKvNamespaceBulkGetMultipleResponseResultWorkersKvBulkGetResultWithMetadataValueJSON
// contains the JSON metadata for the struct
// [AccountStorageKvNamespaceBulkGetMultipleResponseResultWorkersKvBulkGetResultWithMetadataValue]
type accountStorageKvNamespaceBulkGetMultipleResponseResultWorkersKvBulkGetResultWithMetadataValueJSON struct {
	Metadata    apijson.Field
	Value       apijson.Field
	Expiration  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStorageKvNamespaceBulkGetMultipleResponseResultWorkersKvBulkGetResultWithMetadataValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountStorageKvNamespaceBulkGetMultipleResponseResultWorkersKvBulkGetResultWithMetadataValueJSON) RawJSON() string {
	return r.raw
}

type AccountStorageKvNamespaceBulkWriteResponse struct {
	Errors   []MessagesWorkersKvItem `json:"errors,required"`
	Messages []MessagesWorkersKvItem `json:"messages,required"`
	// Whether the API call was successful.
	Success AccountStorageKvNamespaceBulkWriteResponseSuccess `json:"success,required"`
	Result  BulkResult                                        `json:"result,nullable"`
	JSON    accountStorageKvNamespaceBulkWriteResponseJSON    `json:"-"`
}

// accountStorageKvNamespaceBulkWriteResponseJSON contains the JSON metadata for
// the struct [AccountStorageKvNamespaceBulkWriteResponse]
type accountStorageKvNamespaceBulkWriteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStorageKvNamespaceBulkWriteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountStorageKvNamespaceBulkWriteResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccountStorageKvNamespaceBulkWriteResponseSuccess bool

const (
	AccountStorageKvNamespaceBulkWriteResponseSuccessTrue AccountStorageKvNamespaceBulkWriteResponseSuccess = true
)

func (r AccountStorageKvNamespaceBulkWriteResponseSuccess) IsKnown() bool {
	switch r {
	case AccountStorageKvNamespaceBulkWriteResponseSuccessTrue:
		return true
	}
	return false
}

type AccountStorageKvNamespaceBulkDeleteMultipleParams struct {
	Body []string `json:"body,required"`
}

func (r AccountStorageKvNamespaceBulkDeleteMultipleParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type AccountStorageKvNamespaceBulkGetMultipleParams struct {
	// Array of keys to retrieve (maximum of 100).
	Keys param.Field[[]string] `json:"keys,required"`
	// Whether to parse JSON values in the response.
	Type param.Field[AccountStorageKvNamespaceBulkGetMultipleParamsType] `json:"type"`
	// Whether to include metadata in the response.
	WithMetadata param.Field[bool] `json:"withMetadata"`
}

func (r AccountStorageKvNamespaceBulkGetMultipleParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Whether to parse JSON values in the response.
type AccountStorageKvNamespaceBulkGetMultipleParamsType string

const (
	AccountStorageKvNamespaceBulkGetMultipleParamsTypeText AccountStorageKvNamespaceBulkGetMultipleParamsType = "text"
	AccountStorageKvNamespaceBulkGetMultipleParamsTypeJson AccountStorageKvNamespaceBulkGetMultipleParamsType = "json"
)

func (r AccountStorageKvNamespaceBulkGetMultipleParamsType) IsKnown() bool {
	switch r {
	case AccountStorageKvNamespaceBulkGetMultipleParamsTypeText, AccountStorageKvNamespaceBulkGetMultipleParamsTypeJson:
		return true
	}
	return false
}

type AccountStorageKvNamespaceBulkWriteParams struct {
	Body []AccountStorageKvNamespaceBulkWriteParamsBody `json:"body,required"`
}

func (r AccountStorageKvNamespaceBulkWriteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type AccountStorageKvNamespaceBulkWriteParamsBody struct {
	// A key's name. The name may be at most 512 bytes. All printable, non-whitespace
	// characters are valid.
	Key param.Field[string] `json:"key,required"`
	// A UTF-8 encoded string to be stored, up to 25 MiB in length.
	Value param.Field[string] `json:"value,required"`
	// Indicates whether or not the server should base64 decode the value before
	// storing it. Useful for writing values that wouldn't otherwise be valid JSON
	// strings, such as images.
	Base64 param.Field[bool] `json:"base64"`
	// Expires the key at a certain time, measured in number of seconds since the UNIX
	// epoch.
	Expiration param.Field[float64] `json:"expiration"`
	// Expires the key after a number of seconds. Must be at least 60.
	ExpirationTtl param.Field[float64]     `json:"expiration_ttl"`
	Metadata      param.Field[interface{}] `json:"metadata"`
}

func (r AccountStorageKvNamespaceBulkWriteParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
