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

// AccountRuleListService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountRuleListService] method instead.
type AccountRuleListService struct {
	Options []option.RequestOption
	Items   *AccountRuleListItemService
}

// NewAccountRuleListService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccountRuleListService(opts ...option.RequestOption) (r *AccountRuleListService) {
	r = &AccountRuleListService{}
	r.Options = opts
	r.Items = NewAccountRuleListItemService(opts...)
	return
}

// Creates a new list of the specified type.
func (r *AccountRuleListService) New(ctx context.Context, accountID string, body AccountRuleListNewParams, opts ...option.RequestOption) (res *ListResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/rules/lists", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Fetches the details of a list.
func (r *AccountRuleListService) Get(ctx context.Context, accountID string, listID string, opts ...option.RequestOption) (res *ListResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if listID == "" {
		err = errors.New("missing required list_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/rules/lists/%s", accountID, listID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates the description of a list.
func (r *AccountRuleListService) Update(ctx context.Context, accountID string, listID string, body AccountRuleListUpdateParams, opts ...option.RequestOption) (res *ListResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if listID == "" {
		err = errors.New("missing required list_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/rules/lists/%s", accountID, listID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Fetches all lists in the account.
func (r *AccountRuleListService) List(ctx context.Context, accountID string, opts ...option.RequestOption) (res *AccountRuleListListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/rules/lists", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Deletes a specific list and all its items.
func (r *AccountRuleListService) Delete(ctx context.Context, accountID string, listID string, body AccountRuleListDeleteParams, opts ...option.RequestOption) (res *AccountRuleListDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if listID == "" {
		err = errors.New("missing required list_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/rules/lists/%s", accountID, listID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

// Gets the current status of an asynchronous operation on a list.
//
// The `status` property can have one of the following values: `pending`,
// `running`, `completed`, or `failed`. If the status is `failed`, the `error`
// property will contain a message describing the error.
func (r *AccountRuleListService) GetBulkOperationStatus(ctx context.Context, accountIdentifier string, operationID string, opts ...option.RequestOption) (res *AccountRuleListGetBulkOperationStatusResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountIdentifier == "" {
		err = errors.New("missing required account_identifier parameter")
		return
	}
	if operationID == "" {
		err = errors.New("missing required operation_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/rules/lists/bulk_operations/%s", accountIdentifier, operationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type APIResponseCollectionLists struct {
	Result []interface{}                  `json:"result,nullable"`
	JSON   apiResponseCollectionListsJSON `json:"-"`
	APIResponseLists
}

// apiResponseCollectionListsJSON contains the JSON metadata for the struct
// [APIResponseCollectionLists]
type apiResponseCollectionListsJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIResponseCollectionLists) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiResponseCollectionListsJSON) RawJSON() string {
	return r.raw
}

type APIResponseLists struct {
	Errors   []ListMessagesItem `json:"errors,required"`
	Messages []ListMessagesItem `json:"messages,required"`
	Result   []interface{}      `json:"result,required"`
	// Whether the API call was successful
	Success APIResponseListsSuccess `json:"success,required"`
	JSON    apiResponseListsJSON    `json:"-"`
}

// apiResponseListsJSON contains the JSON metadata for the struct
// [APIResponseLists]
type apiResponseListsJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIResponseLists) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiResponseListsJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type APIResponseListsSuccess bool

const (
	APIResponseListsSuccessTrue APIResponseListsSuccess = true
)

func (r APIResponseListsSuccess) IsKnown() bool {
	switch r {
	case APIResponseListsSuccessTrue:
		return true
	}
	return false
}

// The type of the list. Each type supports specific list items (IP addresses,
// ASNs, hostnames or redirects).
type ListKind string

const (
	ListKindIP       ListKind = "ip"
	ListKindRedirect ListKind = "redirect"
	ListKindHostname ListKind = "hostname"
	ListKindAsn      ListKind = "asn"
)

func (r ListKind) IsKnown() bool {
	switch r {
	case ListKindIP, ListKindRedirect, ListKindHostname, ListKindAsn:
		return true
	}
	return false
}

type ListMessagesItem struct {
	Code    int64                `json:"code,required"`
	Message string               `json:"message,required"`
	JSON    listMessagesItemJSON `json:"-"`
}

// listMessagesItemJSON contains the JSON metadata for the struct
// [ListMessagesItem]
type listMessagesItemJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ListMessagesItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listMessagesItemJSON) RawJSON() string {
	return r.raw
}

type ListResponseCollection struct {
	Result ListRules                  `json:"result"`
	JSON   listResponseCollectionJSON `json:"-"`
	APIResponseLists
}

// listResponseCollectionJSON contains the JSON metadata for the struct
// [ListResponseCollection]
type listResponseCollectionJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ListResponseCollection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listResponseCollectionJSON) RawJSON() string {
	return r.raw
}

type ListRules struct {
	// The unique ID of the list.
	ID string `json:"id"`
	// The RFC 3339 timestamp of when the list was created.
	CreatedOn string `json:"created_on"`
	// An informative summary of the list.
	Description string `json:"description"`
	// The type of the list. Each type supports specific list items (IP addresses,
	// ASNs, hostnames or redirects).
	Kind ListKind `json:"kind"`
	// The RFC 3339 timestamp of when the list was last modified.
	ModifiedOn string `json:"modified_on"`
	// An informative name for the list. Use this name in filter and rule expressions.
	Name string `json:"name"`
	// The number of items in the list.
	NumItems float64 `json:"num_items"`
	// The number of [filters](/operations/filters-list-filters) referencing the list.
	NumReferencingFilters float64       `json:"num_referencing_filters"`
	JSON                  listRulesJSON `json:"-"`
}

// listRulesJSON contains the JSON metadata for the struct [ListRules]
type listRulesJSON struct {
	ID                    apijson.Field
	CreatedOn             apijson.Field
	Description           apijson.Field
	Kind                  apijson.Field
	ModifiedOn            apijson.Field
	Name                  apijson.Field
	NumItems              apijson.Field
	NumReferencingFilters apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *ListRules) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listRulesJSON) RawJSON() string {
	return r.raw
}

type AccountRuleListListResponse struct {
	Result []AccountRuleListListResponseResult `json:"result"`
	JSON   accountRuleListListResponseJSON     `json:"-"`
	APIResponseCollectionLists
}

// accountRuleListListResponseJSON contains the JSON metadata for the struct
// [AccountRuleListListResponse]
type accountRuleListListResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRuleListListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountRuleListListResponseJSON) RawJSON() string {
	return r.raw
}

type AccountRuleListListResponseResult struct {
	JSON accountRuleListListResponseResultJSON `json:"-"`
	ListRules
}

// accountRuleListListResponseResultJSON contains the JSON metadata for the struct
// [AccountRuleListListResponseResult]
type accountRuleListListResponseResultJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRuleListListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountRuleListListResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountRuleListDeleteResponse struct {
	Result AccountRuleListDeleteResponseResult `json:"result"`
	JSON   accountRuleListDeleteResponseJSON   `json:"-"`
	APIResponseLists
}

// accountRuleListDeleteResponseJSON contains the JSON metadata for the struct
// [AccountRuleListDeleteResponse]
type accountRuleListDeleteResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRuleListDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountRuleListDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type AccountRuleListDeleteResponseResult struct {
	// The unique ID of the item in the List.
	ID   string                                  `json:"id"`
	JSON accountRuleListDeleteResponseResultJSON `json:"-"`
}

// accountRuleListDeleteResponseResultJSON contains the JSON metadata for the
// struct [AccountRuleListDeleteResponseResult]
type accountRuleListDeleteResponseResultJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRuleListDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountRuleListDeleteResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountRuleListGetBulkOperationStatusResponse struct {
	Result AccountRuleListGetBulkOperationStatusResponseResult `json:"result"`
	JSON   accountRuleListGetBulkOperationStatusResponseJSON   `json:"-"`
	APIResponseCollectionLists
}

// accountRuleListGetBulkOperationStatusResponseJSON contains the JSON metadata for
// the struct [AccountRuleListGetBulkOperationStatusResponse]
type accountRuleListGetBulkOperationStatusResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRuleListGetBulkOperationStatusResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountRuleListGetBulkOperationStatusResponseJSON) RawJSON() string {
	return r.raw
}

type AccountRuleListGetBulkOperationStatusResponseResult struct {
	// The unique operation ID of the asynchronous action.
	ID string `json:"id,required"`
	// The current status of the asynchronous operation.
	Status AccountRuleListGetBulkOperationStatusResponseResultStatus `json:"status,required"`
	// The RFC 3339 timestamp of when the operation was completed.
	Completed string `json:"completed"`
	// A message describing the error when the status is `failed`.
	Error string                                                  `json:"error"`
	JSON  accountRuleListGetBulkOperationStatusResponseResultJSON `json:"-"`
}

// accountRuleListGetBulkOperationStatusResponseResultJSON contains the JSON
// metadata for the struct [AccountRuleListGetBulkOperationStatusResponseResult]
type accountRuleListGetBulkOperationStatusResponseResultJSON struct {
	ID          apijson.Field
	Status      apijson.Field
	Completed   apijson.Field
	Error       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRuleListGetBulkOperationStatusResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountRuleListGetBulkOperationStatusResponseResultJSON) RawJSON() string {
	return r.raw
}

// The current status of the asynchronous operation.
type AccountRuleListGetBulkOperationStatusResponseResultStatus string

const (
	AccountRuleListGetBulkOperationStatusResponseResultStatusPending   AccountRuleListGetBulkOperationStatusResponseResultStatus = "pending"
	AccountRuleListGetBulkOperationStatusResponseResultStatusRunning   AccountRuleListGetBulkOperationStatusResponseResultStatus = "running"
	AccountRuleListGetBulkOperationStatusResponseResultStatusCompleted AccountRuleListGetBulkOperationStatusResponseResultStatus = "completed"
	AccountRuleListGetBulkOperationStatusResponseResultStatusFailed    AccountRuleListGetBulkOperationStatusResponseResultStatus = "failed"
)

func (r AccountRuleListGetBulkOperationStatusResponseResultStatus) IsKnown() bool {
	switch r {
	case AccountRuleListGetBulkOperationStatusResponseResultStatusPending, AccountRuleListGetBulkOperationStatusResponseResultStatusRunning, AccountRuleListGetBulkOperationStatusResponseResultStatusCompleted, AccountRuleListGetBulkOperationStatusResponseResultStatusFailed:
		return true
	}
	return false
}

type AccountRuleListNewParams struct {
	// The type of the list. Each type supports specific list items (IP addresses,
	// ASNs, hostnames or redirects).
	Kind param.Field[ListKind] `json:"kind,required"`
	// An informative name for the list. Use this name in filter and rule expressions.
	Name param.Field[string] `json:"name,required"`
	// An informative summary of the list.
	Description param.Field[string] `json:"description"`
}

func (r AccountRuleListNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountRuleListUpdateParams struct {
	// An informative summary of the list.
	Description param.Field[string] `json:"description"`
}

func (r AccountRuleListUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountRuleListDeleteParams struct {
	Body interface{} `json:"body,required"`
}

func (r AccountRuleListDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}
