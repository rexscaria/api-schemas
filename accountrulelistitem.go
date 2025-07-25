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

// AccountRuleListItemService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountRuleListItemService] method instead.
type AccountRuleListItemService struct {
	Options []option.RequestOption
}

// NewAccountRuleListItemService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountRuleListItemService(opts ...option.RequestOption) (r *AccountRuleListItemService) {
	r = &AccountRuleListItemService{}
	r.Options = opts
	return
}

// Appends new items to the list.
//
// This operation is asynchronous. To get current the operation status, invoke the
// `Get bulk operation status` endpoint with the returned `operation_id`.
func (r *AccountRuleListItemService) New(ctx context.Context, accountID string, listID string, body AccountRuleListItemNewParams, opts ...option.RequestOption) (res *AsyncResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if listID == "" {
		err = errors.New("missing required list_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/rules/lists/%s/items", accountID, listID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Removes all existing items from the list and adds the provided items to the
// list.
//
// This operation is asynchronous. To get current the operation status, invoke the
// `Get bulk operation status` endpoint with the returned `operation_id`.
func (r *AccountRuleListItemService) Update(ctx context.Context, accountID string, listID string, body AccountRuleListItemUpdateParams, opts ...option.RequestOption) (res *AsyncResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if listID == "" {
		err = errors.New("missing required list_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/rules/lists/%s/items", accountID, listID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Fetches all the items in the list.
func (r *AccountRuleListItemService) List(ctx context.Context, accountID string, listID string, query AccountRuleListItemListParams, opts ...option.RequestOption) (res *AccountRuleListItemListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if listID == "" {
		err = errors.New("missing required list_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/rules/lists/%s/items", accountID, listID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Removes one or more items from a list.
//
// This operation is asynchronous. To get current the operation status, invoke the
// `Get bulk operation status` endpoint with the returned `operation_id`.
func (r *AccountRuleListItemService) Delete(ctx context.Context, accountID string, listID string, body AccountRuleListItemDeleteParams, opts ...option.RequestOption) (res *AsyncResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if listID == "" {
		err = errors.New("missing required list_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/rules/lists/%s/items", accountID, listID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

type AsyncResponse struct {
	Errors   []AsyncResponseError   `json:"errors,required"`
	Messages []AsyncResponseMessage `json:"messages,required"`
	Result   AsyncResponseResult    `json:"result,required"`
	// Defines whether the API call was successful.
	Success AsyncResponseSuccess `json:"success,required"`
	JSON    asyncResponseJSON    `json:"-"`
}

// asyncResponseJSON contains the JSON metadata for the struct [AsyncResponse]
type asyncResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AsyncResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r asyncResponseJSON) RawJSON() string {
	return r.raw
}

type AsyncResponseError struct {
	Code             int64                     `json:"code,required"`
	Message          string                    `json:"message,required"`
	DocumentationURL string                    `json:"documentation_url"`
	Source           AsyncResponseErrorsSource `json:"source"`
	JSON             asyncResponseErrorJSON    `json:"-"`
}

// asyncResponseErrorJSON contains the JSON metadata for the struct
// [AsyncResponseError]
type asyncResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AsyncResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r asyncResponseErrorJSON) RawJSON() string {
	return r.raw
}

type AsyncResponseErrorsSource struct {
	Pointer string                        `json:"pointer"`
	JSON    asyncResponseErrorsSourceJSON `json:"-"`
}

// asyncResponseErrorsSourceJSON contains the JSON metadata for the struct
// [AsyncResponseErrorsSource]
type asyncResponseErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AsyncResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r asyncResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AsyncResponseMessage struct {
	Code             int64                       `json:"code,required"`
	Message          string                      `json:"message,required"`
	DocumentationURL string                      `json:"documentation_url"`
	Source           AsyncResponseMessagesSource `json:"source"`
	JSON             asyncResponseMessageJSON    `json:"-"`
}

// asyncResponseMessageJSON contains the JSON metadata for the struct
// [AsyncResponseMessage]
type asyncResponseMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AsyncResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r asyncResponseMessageJSON) RawJSON() string {
	return r.raw
}

type AsyncResponseMessagesSource struct {
	Pointer string                          `json:"pointer"`
	JSON    asyncResponseMessagesSourceJSON `json:"-"`
}

// asyncResponseMessagesSourceJSON contains the JSON metadata for the struct
// [AsyncResponseMessagesSource]
type asyncResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AsyncResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r asyncResponseMessagesSourceJSON) RawJSON() string {
	return r.raw
}

type AsyncResponseResult struct {
	// The unique operation ID of the asynchronous action.
	OperationID string                  `json:"operation_id,required"`
	JSON        asyncResponseResultJSON `json:"-"`
}

// asyncResponseResultJSON contains the JSON metadata for the struct
// [AsyncResponseResult]
type asyncResponseResultJSON struct {
	OperationID apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AsyncResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r asyncResponseResultJSON) RawJSON() string {
	return r.raw
}

// Defines whether the API call was successful.
type AsyncResponseSuccess bool

const (
	AsyncResponseSuccessTrue AsyncResponseSuccess = true
)

func (r AsyncResponseSuccess) IsKnown() bool {
	switch r {
	case AsyncResponseSuccessTrue:
		return true
	}
	return false
}

type ListItem struct {
	// The unique ID of the list.
	ID string `json:"id"`
	// Defines a non-negative 32 bit integer.
	Asn int64 `json:"asn"`
	// Defines an informative summary of the list item.
	Comment string `json:"comment"`
	// The RFC 3339 timestamp of when the item was created.
	CreatedOn string `json:"created_on"`
	// Valid characters for hostnames are ASCII(7) letters from a to z, the digits from
	// 0 to 9, wildcards (\*), and the hyphen (-).
	Hostname ListItemHostname `json:"hostname"`
	// An IPv4 address, an IPv4 CIDR, an IPv6 address, or an IPv6 CIDR.
	IP string `json:"ip"`
	// The RFC 3339 timestamp of when the item was last modified.
	ModifiedOn string `json:"modified_on"`
	// The definition of the redirect.
	Redirect ListItemRedirect `json:"redirect"`
	JSON     listItemJSON     `json:"-"`
}

// listItemJSON contains the JSON metadata for the struct [ListItem]
type listItemJSON struct {
	ID          apijson.Field
	Asn         apijson.Field
	Comment     apijson.Field
	CreatedOn   apijson.Field
	Hostname    apijson.Field
	IP          apijson.Field
	ModifiedOn  apijson.Field
	Redirect    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ListItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listItemJSON) RawJSON() string {
	return r.raw
}

// Valid characters for hostnames are ASCII(7) letters from a to z, the digits from
// 0 to 9, wildcards (\*), and the hyphen (-).
type ListItemHostname struct {
	URLHostname string `json:"url_hostname,required"`
	// Only applies to wildcard hostnames (e.g., \*.example.com). When true (default),
	// only subdomains are blocked. When false, both the root domain and subdomains are
	// blocked.
	ExcludeExactHostname bool                 `json:"exclude_exact_hostname"`
	JSON                 listItemHostnameJSON `json:"-"`
}

// listItemHostnameJSON contains the JSON metadata for the struct
// [ListItemHostname]
type listItemHostnameJSON struct {
	URLHostname          apijson.Field
	ExcludeExactHostname apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ListItemHostname) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listItemHostnameJSON) RawJSON() string {
	return r.raw
}

// Valid characters for hostnames are ASCII(7) letters from a to z, the digits from
// 0 to 9, wildcards (\*), and the hyphen (-).
type ListItemHostnameParam struct {
	URLHostname param.Field[string] `json:"url_hostname,required"`
	// Only applies to wildcard hostnames (e.g., \*.example.com). When true (default),
	// only subdomains are blocked. When false, both the root domain and subdomains are
	// blocked.
	ExcludeExactHostname param.Field[bool] `json:"exclude_exact_hostname"`
}

func (r ListItemHostnameParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The definition of the redirect.
type ListItemRedirect struct {
	SourceURL           string                     `json:"source_url,required"`
	TargetURL           string                     `json:"target_url,required"`
	IncludeSubdomains   bool                       `json:"include_subdomains"`
	PreservePathSuffix  bool                       `json:"preserve_path_suffix"`
	PreserveQueryString bool                       `json:"preserve_query_string"`
	StatusCode          ListItemRedirectStatusCode `json:"status_code"`
	SubpathMatching     bool                       `json:"subpath_matching"`
	JSON                listItemRedirectJSON       `json:"-"`
}

// listItemRedirectJSON contains the JSON metadata for the struct
// [ListItemRedirect]
type listItemRedirectJSON struct {
	SourceURL           apijson.Field
	TargetURL           apijson.Field
	IncludeSubdomains   apijson.Field
	PreservePathSuffix  apijson.Field
	PreserveQueryString apijson.Field
	StatusCode          apijson.Field
	SubpathMatching     apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *ListItemRedirect) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listItemRedirectJSON) RawJSON() string {
	return r.raw
}

type ListItemRedirectStatusCode int64

const (
	ListItemRedirectStatusCode301 ListItemRedirectStatusCode = 301
	ListItemRedirectStatusCode302 ListItemRedirectStatusCode = 302
	ListItemRedirectStatusCode307 ListItemRedirectStatusCode = 307
	ListItemRedirectStatusCode308 ListItemRedirectStatusCode = 308
)

func (r ListItemRedirectStatusCode) IsKnown() bool {
	switch r {
	case ListItemRedirectStatusCode301, ListItemRedirectStatusCode302, ListItemRedirectStatusCode307, ListItemRedirectStatusCode308:
		return true
	}
	return false
}

// The definition of the redirect.
type ListItemRedirectParam struct {
	SourceURL           param.Field[string]                     `json:"source_url,required"`
	TargetURL           param.Field[string]                     `json:"target_url,required"`
	IncludeSubdomains   param.Field[bool]                       `json:"include_subdomains"`
	PreservePathSuffix  param.Field[bool]                       `json:"preserve_path_suffix"`
	PreserveQueryString param.Field[bool]                       `json:"preserve_query_string"`
	StatusCode          param.Field[ListItemRedirectStatusCode] `json:"status_code"`
	SubpathMatching     param.Field[bool]                       `json:"subpath_matching"`
}

func (r ListItemRedirectParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UpdateRequestCollectionItemParam struct {
	// Defines a non-negative 32 bit integer.
	Asn param.Field[int64] `json:"asn"`
	// Defines an informative summary of the list item.
	Comment param.Field[string] `json:"comment"`
	// Valid characters for hostnames are ASCII(7) letters from a to z, the digits from
	// 0 to 9, wildcards (\*), and the hyphen (-).
	Hostname param.Field[ListItemHostnameParam] `json:"hostname"`
	// An IPv4 address, an IPv4 CIDR, an IPv6 address, or an IPv6 CIDR.
	IP param.Field[string] `json:"ip"`
	// The definition of the redirect.
	Redirect param.Field[ListItemRedirectParam] `json:"redirect"`
}

func (r UpdateRequestCollectionItemParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UpdateRequestCollectionItemParam) implementsUpdateRequestCollectionItemUnionParam() {}

// Satisfied by [UpdateRequestCollectionItemObjectParam],
// [UpdateRequestCollectionItemObjectParam],
// [UpdateRequestCollectionItemObjectParam],
// [UpdateRequestCollectionItemObjectParam], [UpdateRequestCollectionItemParam].
type UpdateRequestCollectionItemUnionParam interface {
	implementsUpdateRequestCollectionItemUnionParam()
}

type UpdateRequestCollectionItemObjectParam struct {
	// An IPv4 address, an IPv4 CIDR, an IPv6 address, or an IPv6 CIDR.
	IP param.Field[string] `json:"ip,required"`
	// Defines an informative summary of the list item.
	Comment param.Field[string] `json:"comment"`
}

func (r UpdateRequestCollectionItemObjectParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r UpdateRequestCollectionItemObjectParam) implementsUpdateRequestCollectionItemUnionParam() {}

type AccountRuleListItemListResponse struct {
	Errors   []AccountRuleListItemListResponseError   `json:"errors,required"`
	Messages []AccountRuleListItemListResponseMessage `json:"messages,required"`
	Result   []ListItem                               `json:"result,required"`
	// Defines whether the API call was successful.
	Success    AccountRuleListItemListResponseSuccess    `json:"success,required"`
	ResultInfo AccountRuleListItemListResponseResultInfo `json:"result_info"`
	JSON       accountRuleListItemListResponseJSON       `json:"-"`
}

// accountRuleListItemListResponseJSON contains the JSON metadata for the struct
// [AccountRuleListItemListResponse]
type accountRuleListItemListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRuleListItemListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountRuleListItemListResponseJSON) RawJSON() string {
	return r.raw
}

type AccountRuleListItemListResponseError struct {
	Code             int64                                       `json:"code,required"`
	Message          string                                      `json:"message,required"`
	DocumentationURL string                                      `json:"documentation_url"`
	Source           AccountRuleListItemListResponseErrorsSource `json:"source"`
	JSON             accountRuleListItemListResponseErrorJSON    `json:"-"`
}

// accountRuleListItemListResponseErrorJSON contains the JSON metadata for the
// struct [AccountRuleListItemListResponseError]
type accountRuleListItemListResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountRuleListItemListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountRuleListItemListResponseErrorJSON) RawJSON() string {
	return r.raw
}

type AccountRuleListItemListResponseErrorsSource struct {
	Pointer string                                          `json:"pointer"`
	JSON    accountRuleListItemListResponseErrorsSourceJSON `json:"-"`
}

// accountRuleListItemListResponseErrorsSourceJSON contains the JSON metadata for
// the struct [AccountRuleListItemListResponseErrorsSource]
type accountRuleListItemListResponseErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRuleListItemListResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountRuleListItemListResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AccountRuleListItemListResponseMessage struct {
	Code             int64                                         `json:"code,required"`
	Message          string                                        `json:"message,required"`
	DocumentationURL string                                        `json:"documentation_url"`
	Source           AccountRuleListItemListResponseMessagesSource `json:"source"`
	JSON             accountRuleListItemListResponseMessageJSON    `json:"-"`
}

// accountRuleListItemListResponseMessageJSON contains the JSON metadata for the
// struct [AccountRuleListItemListResponseMessage]
type accountRuleListItemListResponseMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountRuleListItemListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountRuleListItemListResponseMessageJSON) RawJSON() string {
	return r.raw
}

type AccountRuleListItemListResponseMessagesSource struct {
	Pointer string                                            `json:"pointer"`
	JSON    accountRuleListItemListResponseMessagesSourceJSON `json:"-"`
}

// accountRuleListItemListResponseMessagesSourceJSON contains the JSON metadata for
// the struct [AccountRuleListItemListResponseMessagesSource]
type accountRuleListItemListResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRuleListItemListResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountRuleListItemListResponseMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Defines whether the API call was successful.
type AccountRuleListItemListResponseSuccess bool

const (
	AccountRuleListItemListResponseSuccessTrue AccountRuleListItemListResponseSuccess = true
)

func (r AccountRuleListItemListResponseSuccess) IsKnown() bool {
	switch r {
	case AccountRuleListItemListResponseSuccessTrue:
		return true
	}
	return false
}

type AccountRuleListItemListResponseResultInfo struct {
	Cursors AccountRuleListItemListResponseResultInfoCursors `json:"cursors"`
	JSON    accountRuleListItemListResponseResultInfoJSON    `json:"-"`
}

// accountRuleListItemListResponseResultInfoJSON contains the JSON metadata for the
// struct [AccountRuleListItemListResponseResultInfo]
type accountRuleListItemListResponseResultInfoJSON struct {
	Cursors     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRuleListItemListResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountRuleListItemListResponseResultInfoJSON) RawJSON() string {
	return r.raw
}

type AccountRuleListItemListResponseResultInfoCursors struct {
	After  string                                               `json:"after"`
	Before string                                               `json:"before"`
	JSON   accountRuleListItemListResponseResultInfoCursorsJSON `json:"-"`
}

// accountRuleListItemListResponseResultInfoCursorsJSON contains the JSON metadata
// for the struct [AccountRuleListItemListResponseResultInfoCursors]
type accountRuleListItemListResponseResultInfoCursorsJSON struct {
	After       apijson.Field
	Before      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRuleListItemListResponseResultInfoCursors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountRuleListItemListResponseResultInfoCursorsJSON) RawJSON() string {
	return r.raw
}

type AccountRuleListItemNewParams struct {
	Body []UpdateRequestCollectionItemUnionParam `json:"body,required"`
}

func (r AccountRuleListItemNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type AccountRuleListItemUpdateParams struct {
	Body []UpdateRequestCollectionItemUnionParam `json:"body,required"`
}

func (r AccountRuleListItemUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type AccountRuleListItemListParams struct {
	// The pagination cursor. An opaque string token indicating the position from which
	// to continue when requesting the next/previous set of records. Cursor values are
	// provided under `result_info.cursors` in the response. You should make no
	// assumptions about a cursor's content or length.
	Cursor param.Field[string] `query:"cursor"`
	// Amount of results to include in each paginated response. A non-negative 32 bit
	// integer.
	PerPage param.Field[int64] `query:"per_page"`
	// A search query to filter returned items. Its meaning depends on the list type:
	// IP addresses must start with the provided string, hostnames and bulk redirects
	// must contain the string, and ASNs must match the string exactly.
	Search param.Field[string] `query:"search"`
}

// URLQuery serializes [AccountRuleListItemListParams]'s query parameters as
// `url.Values`.
func (r AccountRuleListItemListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountRuleListItemDeleteParams struct {
	Items param.Field[[]AccountRuleListItemDeleteParamsItem] `json:"items"`
}

func (r AccountRuleListItemDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountRuleListItemDeleteParamsItem struct {
}

func (r AccountRuleListItemDeleteParamsItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
