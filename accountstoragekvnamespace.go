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

// AccountStorageKvNamespaceService contains methods and other services that help
// with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountStorageKvNamespaceService] method instead.
type AccountStorageKvNamespaceService struct {
	Options []option.RequestOption
	Bulk    *AccountStorageKvNamespaceBulkService
	Values  *AccountStorageKvNamespaceValueService
}

// NewAccountStorageKvNamespaceService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountStorageKvNamespaceService(opts ...option.RequestOption) (r *AccountStorageKvNamespaceService) {
	r = &AccountStorageKvNamespaceService{}
	r.Options = opts
	r.Bulk = NewAccountStorageKvNamespaceBulkService(opts...)
	r.Values = NewAccountStorageKvNamespaceValueService(opts...)
	return
}

// Creates a namespace under the given title. A `400` is returned if the account
// already owns a namespace with this title. A namespace must be explicitly deleted
// to be replaced.
func (r *AccountStorageKvNamespaceService) New(ctx context.Context, accountID string, body AccountStorageKvNamespaceNewParams, opts ...option.RequestOption) (res *AccountStorageKvNamespaceNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/storage/kv/namespaces", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get the namespace corresponding to the given ID.
func (r *AccountStorageKvNamespaceService) Get(ctx context.Context, accountID string, namespaceID string, opts ...option.RequestOption) (res *AccountStorageKvNamespaceGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if namespaceID == "" {
		err = errors.New("missing required namespace_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/storage/kv/namespaces/%s", accountID, namespaceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Modifies a namespace's title.
func (r *AccountStorageKvNamespaceService) Update(ctx context.Context, accountID string, namespaceID string, body AccountStorageKvNamespaceUpdateParams, opts ...option.RequestOption) (res *AccountStorageKvNamespaceUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if namespaceID == "" {
		err = errors.New("missing required namespace_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/storage/kv/namespaces/%s", accountID, namespaceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Returns the namespaces owned by an account.
func (r *AccountStorageKvNamespaceService) List(ctx context.Context, accountID string, query AccountStorageKvNamespaceListParams, opts ...option.RequestOption) (res *AccountStorageKvNamespaceListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/storage/kv/namespaces", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Deletes the namespace corresponding to the given ID.
func (r *AccountStorageKvNamespaceService) Delete(ctx context.Context, accountID string, namespaceID string, opts ...option.RequestOption) (res *APIResponseCommonNoResult, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if namespaceID == "" {
		err = errors.New("missing required namespace_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/storage/kv/namespaces/%s", accountID, namespaceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Returns the metadata associated with the given key in the given namespace. Use
// URL-encoding to use special characters (for example, `:`, `!`, `%`) in the key
// name.
func (r *AccountStorageKvNamespaceService) GetMetadata(ctx context.Context, accountID string, namespaceID string, keyName string, opts ...option.RequestOption) (res *AccountStorageKvNamespaceGetMetadataResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if namespaceID == "" {
		err = errors.New("missing required namespace_id parameter")
		return
	}
	if keyName == "" {
		err = errors.New("missing required key_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/storage/kv/namespaces/%s/metadata/%s", accountID, namespaceID, keyName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Lists a namespace's keys.
func (r *AccountStorageKvNamespaceService) ListKeys(ctx context.Context, accountID string, namespaceID string, query AccountStorageKvNamespaceListKeysParams, opts ...option.RequestOption) (res *AccountStorageKvNamespaceListKeysResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if namespaceID == "" {
		err = errors.New("missing required namespace_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/storage/kv/namespaces/%s/keys", accountID, namespaceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type APIResponseCommonNoResult struct {
	Errors   []APIResponseCommonNoResultError   `json:"errors,required"`
	Messages []APIResponseCommonNoResultMessage `json:"messages,required"`
	// Whether the API call was successful.
	Success APIResponseCommonNoResultSuccess `json:"success,required"`
	Result  APIResponseCommonNoResultResult  `json:"result,nullable"`
	JSON    apiResponseCommonNoResultJSON    `json:"-"`
}

// apiResponseCommonNoResultJSON contains the JSON metadata for the struct
// [APIResponseCommonNoResult]
type apiResponseCommonNoResultJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIResponseCommonNoResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiResponseCommonNoResultJSON) RawJSON() string {
	return r.raw
}

type APIResponseCommonNoResultError struct {
	Code             int64                                 `json:"code,required"`
	Message          string                                `json:"message,required"`
	DocumentationURL string                                `json:"documentation_url"`
	Source           APIResponseCommonNoResultErrorsSource `json:"source"`
	JSON             apiResponseCommonNoResultErrorJSON    `json:"-"`
}

// apiResponseCommonNoResultErrorJSON contains the JSON metadata for the struct
// [APIResponseCommonNoResultError]
type apiResponseCommonNoResultErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *APIResponseCommonNoResultError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiResponseCommonNoResultErrorJSON) RawJSON() string {
	return r.raw
}

type APIResponseCommonNoResultErrorsSource struct {
	Pointer string                                    `json:"pointer"`
	JSON    apiResponseCommonNoResultErrorsSourceJSON `json:"-"`
}

// apiResponseCommonNoResultErrorsSourceJSON contains the JSON metadata for the
// struct [APIResponseCommonNoResultErrorsSource]
type apiResponseCommonNoResultErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIResponseCommonNoResultErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiResponseCommonNoResultErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type APIResponseCommonNoResultMessage struct {
	Code             int64                                   `json:"code,required"`
	Message          string                                  `json:"message,required"`
	DocumentationURL string                                  `json:"documentation_url"`
	Source           APIResponseCommonNoResultMessagesSource `json:"source"`
	JSON             apiResponseCommonNoResultMessageJSON    `json:"-"`
}

// apiResponseCommonNoResultMessageJSON contains the JSON metadata for the struct
// [APIResponseCommonNoResultMessage]
type apiResponseCommonNoResultMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *APIResponseCommonNoResultMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiResponseCommonNoResultMessageJSON) RawJSON() string {
	return r.raw
}

type APIResponseCommonNoResultMessagesSource struct {
	Pointer string                                      `json:"pointer"`
	JSON    apiResponseCommonNoResultMessagesSourceJSON `json:"-"`
}

// apiResponseCommonNoResultMessagesSourceJSON contains the JSON metadata for the
// struct [APIResponseCommonNoResultMessagesSource]
type apiResponseCommonNoResultMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIResponseCommonNoResultMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiResponseCommonNoResultMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type APIResponseCommonNoResultSuccess bool

const (
	APIResponseCommonNoResultSuccessTrue APIResponseCommonNoResultSuccess = true
)

func (r APIResponseCommonNoResultSuccess) IsKnown() bool {
	switch r {
	case APIResponseCommonNoResultSuccessTrue:
		return true
	}
	return false
}

type APIResponseCommonNoResultResult struct {
	JSON apiResponseCommonNoResultResultJSON `json:"-"`
}

// apiResponseCommonNoResultResultJSON contains the JSON metadata for the struct
// [APIResponseCommonNoResultResult]
type apiResponseCommonNoResultResultJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIResponseCommonNoResultResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiResponseCommonNoResultResultJSON) RawJSON() string {
	return r.raw
}

type CreateRenameNamespaceBodyParam struct {
	// A human-readable string name for a Namespace.
	Title param.Field[string] `json:"title,required"`
}

func (r CreateRenameNamespaceBodyParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MessagesWorkersKvItem struct {
	Code             int64                       `json:"code,required"`
	Message          string                      `json:"message,required"`
	DocumentationURL string                      `json:"documentation_url"`
	Source           MessagesWorkersKvItemSource `json:"source"`
	JSON             messagesWorkersKvItemJSON   `json:"-"`
}

// messagesWorkersKvItemJSON contains the JSON metadata for the struct
// [MessagesWorkersKvItem]
type messagesWorkersKvItemJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *MessagesWorkersKvItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messagesWorkersKvItemJSON) RawJSON() string {
	return r.raw
}

type MessagesWorkersKvItemSource struct {
	Pointer string                          `json:"pointer"`
	JSON    messagesWorkersKvItemSourceJSON `json:"-"`
}

// messagesWorkersKvItemSourceJSON contains the JSON metadata for the struct
// [MessagesWorkersKvItemSource]
type messagesWorkersKvItemSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessagesWorkersKvItemSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messagesWorkersKvItemSourceJSON) RawJSON() string {
	return r.raw
}

type Namespace struct {
	// Namespace identifier tag.
	ID string `json:"id,required"`
	// A human-readable string name for a Namespace.
	Title string `json:"title,required"`
	// True if new beta namespace, with additional preview features.
	Beta bool `json:"beta"`
	// True if keys written on the URL will be URL-decoded before storing. For example,
	// if set to "true", a key written on the URL as "%3F" will be stored as "?".
	SupportsURLEncoding bool          `json:"supports_url_encoding"`
	JSON                namespaceJSON `json:"-"`
}

// namespaceJSON contains the JSON metadata for the struct [Namespace]
type namespaceJSON struct {
	ID                  apijson.Field
	Title               apijson.Field
	Beta                apijson.Field
	SupportsURLEncoding apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *Namespace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceJSON) RawJSON() string {
	return r.raw
}

type AccountStorageKvNamespaceNewResponse struct {
	Errors   []MessagesWorkersKvItem `json:"errors,required"`
	Messages []MessagesWorkersKvItem `json:"messages,required"`
	// Whether the API call was successful.
	Success AccountStorageKvNamespaceNewResponseSuccess `json:"success,required"`
	Result  Namespace                                   `json:"result"`
	JSON    accountStorageKvNamespaceNewResponseJSON    `json:"-"`
}

// accountStorageKvNamespaceNewResponseJSON contains the JSON metadata for the
// struct [AccountStorageKvNamespaceNewResponse]
type accountStorageKvNamespaceNewResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStorageKvNamespaceNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountStorageKvNamespaceNewResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccountStorageKvNamespaceNewResponseSuccess bool

const (
	AccountStorageKvNamespaceNewResponseSuccessTrue AccountStorageKvNamespaceNewResponseSuccess = true
)

func (r AccountStorageKvNamespaceNewResponseSuccess) IsKnown() bool {
	switch r {
	case AccountStorageKvNamespaceNewResponseSuccessTrue:
		return true
	}
	return false
}

type AccountStorageKvNamespaceGetResponse struct {
	Errors   []MessagesWorkersKvItem `json:"errors,required"`
	Messages []MessagesWorkersKvItem `json:"messages,required"`
	// Whether the API call was successful.
	Success AccountStorageKvNamespaceGetResponseSuccess `json:"success,required"`
	Result  Namespace                                   `json:"result"`
	JSON    accountStorageKvNamespaceGetResponseJSON    `json:"-"`
}

// accountStorageKvNamespaceGetResponseJSON contains the JSON metadata for the
// struct [AccountStorageKvNamespaceGetResponse]
type accountStorageKvNamespaceGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStorageKvNamespaceGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountStorageKvNamespaceGetResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccountStorageKvNamespaceGetResponseSuccess bool

const (
	AccountStorageKvNamespaceGetResponseSuccessTrue AccountStorageKvNamespaceGetResponseSuccess = true
)

func (r AccountStorageKvNamespaceGetResponseSuccess) IsKnown() bool {
	switch r {
	case AccountStorageKvNamespaceGetResponseSuccessTrue:
		return true
	}
	return false
}

type AccountStorageKvNamespaceUpdateResponse struct {
	Errors   []MessagesWorkersKvItem `json:"errors,required"`
	Messages []MessagesWorkersKvItem `json:"messages,required"`
	Result   Namespace               `json:"result,required"`
	// Whether the API call was successful.
	Success AccountStorageKvNamespaceUpdateResponseSuccess `json:"success,required"`
	JSON    accountStorageKvNamespaceUpdateResponseJSON    `json:"-"`
}

// accountStorageKvNamespaceUpdateResponseJSON contains the JSON metadata for the
// struct [AccountStorageKvNamespaceUpdateResponse]
type accountStorageKvNamespaceUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStorageKvNamespaceUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountStorageKvNamespaceUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccountStorageKvNamespaceUpdateResponseSuccess bool

const (
	AccountStorageKvNamespaceUpdateResponseSuccessTrue AccountStorageKvNamespaceUpdateResponseSuccess = true
)

func (r AccountStorageKvNamespaceUpdateResponseSuccess) IsKnown() bool {
	switch r {
	case AccountStorageKvNamespaceUpdateResponseSuccessTrue:
		return true
	}
	return false
}

type AccountStorageKvNamespaceListResponse struct {
	Errors   []MessagesWorkersKvItem `json:"errors,required"`
	Messages []MessagesWorkersKvItem `json:"messages,required"`
	// Whether the API call was successful.
	Success    AccountStorageKvNamespaceListResponseSuccess    `json:"success,required"`
	Result     []Namespace                                     `json:"result"`
	ResultInfo AccountStorageKvNamespaceListResponseResultInfo `json:"result_info"`
	JSON       accountStorageKvNamespaceListResponseJSON       `json:"-"`
}

// accountStorageKvNamespaceListResponseJSON contains the JSON metadata for the
// struct [AccountStorageKvNamespaceListResponse]
type accountStorageKvNamespaceListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStorageKvNamespaceListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountStorageKvNamespaceListResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccountStorageKvNamespaceListResponseSuccess bool

const (
	AccountStorageKvNamespaceListResponseSuccessTrue AccountStorageKvNamespaceListResponseSuccess = true
)

func (r AccountStorageKvNamespaceListResponseSuccess) IsKnown() bool {
	switch r {
	case AccountStorageKvNamespaceListResponseSuccessTrue:
		return true
	}
	return false
}

type AccountStorageKvNamespaceListResponseResultInfo struct {
	// Total number of results for the requested service.
	Count float64 `json:"count"`
	// Current page within paginated list of results.
	Page float64 `json:"page"`
	// Number of results per page of results.
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters.
	TotalCount float64                                             `json:"total_count"`
	JSON       accountStorageKvNamespaceListResponseResultInfoJSON `json:"-"`
}

// accountStorageKvNamespaceListResponseResultInfoJSON contains the JSON metadata
// for the struct [AccountStorageKvNamespaceListResponseResultInfo]
type accountStorageKvNamespaceListResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStorageKvNamespaceListResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountStorageKvNamespaceListResponseResultInfoJSON) RawJSON() string {
	return r.raw
}

type AccountStorageKvNamespaceGetMetadataResponse struct {
	Errors   []MessagesWorkersKvItem `json:"errors,required"`
	Messages []MessagesWorkersKvItem `json:"messages,required"`
	// Whether the API call was successful.
	Success AccountStorageKvNamespaceGetMetadataResponseSuccess `json:"success,required"`
	Result  interface{}                                         `json:"result"`
	JSON    accountStorageKvNamespaceGetMetadataResponseJSON    `json:"-"`
}

// accountStorageKvNamespaceGetMetadataResponseJSON contains the JSON metadata for
// the struct [AccountStorageKvNamespaceGetMetadataResponse]
type accountStorageKvNamespaceGetMetadataResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStorageKvNamespaceGetMetadataResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountStorageKvNamespaceGetMetadataResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccountStorageKvNamespaceGetMetadataResponseSuccess bool

const (
	AccountStorageKvNamespaceGetMetadataResponseSuccessTrue AccountStorageKvNamespaceGetMetadataResponseSuccess = true
)

func (r AccountStorageKvNamespaceGetMetadataResponseSuccess) IsKnown() bool {
	switch r {
	case AccountStorageKvNamespaceGetMetadataResponseSuccessTrue:
		return true
	}
	return false
}

type AccountStorageKvNamespaceListKeysResponse struct {
	Errors   []MessagesWorkersKvItem `json:"errors,required"`
	Messages []MessagesWorkersKvItem `json:"messages,required"`
	// Whether the API call was successful.
	Success    AccountStorageKvNamespaceListKeysResponseSuccess    `json:"success,required"`
	Result     []AccountStorageKvNamespaceListKeysResponseResult   `json:"result"`
	ResultInfo AccountStorageKvNamespaceListKeysResponseResultInfo `json:"result_info"`
	JSON       accountStorageKvNamespaceListKeysResponseJSON       `json:"-"`
}

// accountStorageKvNamespaceListKeysResponseJSON contains the JSON metadata for the
// struct [AccountStorageKvNamespaceListKeysResponse]
type accountStorageKvNamespaceListKeysResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStorageKvNamespaceListKeysResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountStorageKvNamespaceListKeysResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccountStorageKvNamespaceListKeysResponseSuccess bool

const (
	AccountStorageKvNamespaceListKeysResponseSuccessTrue AccountStorageKvNamespaceListKeysResponseSuccess = true
)

func (r AccountStorageKvNamespaceListKeysResponseSuccess) IsKnown() bool {
	switch r {
	case AccountStorageKvNamespaceListKeysResponseSuccessTrue:
		return true
	}
	return false
}

// A name for a value. A value stored under a given key may be retrieved via the
// same key.
type AccountStorageKvNamespaceListKeysResponseResult struct {
	// A key's name. The name may be at most 512 bytes. All printable, non-whitespace
	// characters are valid. Use percent-encoding to define key names as part of a URL.
	Name string `json:"name,required"`
	// The time, measured in number of seconds since the UNIX epoch, at which the key
	// will expire. This property is omitted for keys that will not expire.
	Expiration float64                                             `json:"expiration"`
	Metadata   interface{}                                         `json:"metadata"`
	JSON       accountStorageKvNamespaceListKeysResponseResultJSON `json:"-"`
}

// accountStorageKvNamespaceListKeysResponseResultJSON contains the JSON metadata
// for the struct [AccountStorageKvNamespaceListKeysResponseResult]
type accountStorageKvNamespaceListKeysResponseResultJSON struct {
	Name        apijson.Field
	Expiration  apijson.Field
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStorageKvNamespaceListKeysResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountStorageKvNamespaceListKeysResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountStorageKvNamespaceListKeysResponseResultInfo struct {
	// Total results returned based on your list parameters.
	Count float64 `json:"count"`
	// Opaque token indicating the position from which to continue when requesting the
	// next set of records if the amount of list results was limited by the limit
	// parameter. A valid value for the cursor can be obtained from the cursors object
	// in the result_info structure.
	Cursor string                                                  `json:"cursor"`
	JSON   accountStorageKvNamespaceListKeysResponseResultInfoJSON `json:"-"`
}

// accountStorageKvNamespaceListKeysResponseResultInfoJSON contains the JSON
// metadata for the struct [AccountStorageKvNamespaceListKeysResponseResultInfo]
type accountStorageKvNamespaceListKeysResponseResultInfoJSON struct {
	Count       apijson.Field
	Cursor      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStorageKvNamespaceListKeysResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountStorageKvNamespaceListKeysResponseResultInfoJSON) RawJSON() string {
	return r.raw
}

type AccountStorageKvNamespaceNewParams struct {
	CreateRenameNamespaceBody CreateRenameNamespaceBodyParam `json:"create_rename_namespace_body,required"`
}

func (r AccountStorageKvNamespaceNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.CreateRenameNamespaceBody)
}

type AccountStorageKvNamespaceUpdateParams struct {
	CreateRenameNamespaceBody CreateRenameNamespaceBodyParam `json:"create_rename_namespace_body,required"`
}

func (r AccountStorageKvNamespaceUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.CreateRenameNamespaceBody)
}

type AccountStorageKvNamespaceListParams struct {
	// Direction to order namespaces.
	Direction param.Field[AccountStorageKvNamespaceListParamsDirection] `query:"direction"`
	// Field to order results by.
	Order param.Field[AccountStorageKvNamespaceListParamsOrder] `query:"order"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Maximum number of results per page.
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes [AccountStorageKvNamespaceListParams]'s query parameters as
// `url.Values`.
func (r AccountStorageKvNamespaceListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Direction to order namespaces.
type AccountStorageKvNamespaceListParamsDirection string

const (
	AccountStorageKvNamespaceListParamsDirectionAsc  AccountStorageKvNamespaceListParamsDirection = "asc"
	AccountStorageKvNamespaceListParamsDirectionDesc AccountStorageKvNamespaceListParamsDirection = "desc"
)

func (r AccountStorageKvNamespaceListParamsDirection) IsKnown() bool {
	switch r {
	case AccountStorageKvNamespaceListParamsDirectionAsc, AccountStorageKvNamespaceListParamsDirectionDesc:
		return true
	}
	return false
}

// Field to order results by.
type AccountStorageKvNamespaceListParamsOrder string

const (
	AccountStorageKvNamespaceListParamsOrderID    AccountStorageKvNamespaceListParamsOrder = "id"
	AccountStorageKvNamespaceListParamsOrderTitle AccountStorageKvNamespaceListParamsOrder = "title"
)

func (r AccountStorageKvNamespaceListParamsOrder) IsKnown() bool {
	switch r {
	case AccountStorageKvNamespaceListParamsOrderID, AccountStorageKvNamespaceListParamsOrderTitle:
		return true
	}
	return false
}

type AccountStorageKvNamespaceListKeysParams struct {
	// Opaque token indicating the position from which to continue when requesting the
	// next set of records if the amount of list results was limited by the limit
	// parameter. A valid value for the cursor can be obtained from the `cursors`
	// object in the `result_info` structure.
	Cursor param.Field[string] `query:"cursor"`
	// Limits the number of keys returned in the response. The cursor attribute may be
	// used to iterate over the next batch of keys if there are more than the limit.
	Limit param.Field[float64] `query:"limit"`
	// Filters returned keys by a name prefix. Exact matches and any key names that
	// begin with the prefix will be returned.
	Prefix param.Field[string] `query:"prefix"`
}

// URLQuery serializes [AccountStorageKvNamespaceListKeysParams]'s query parameters
// as `url.Values`.
func (r AccountStorageKvNamespaceListKeysParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
