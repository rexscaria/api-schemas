// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"

	"github.com/stainless-sdks/cf-rex-go/internal/apijson"
	"github.com/stainless-sdks/cf-rex-go/internal/apiquery"
	"github.com/stainless-sdks/cf-rex-go/internal/param"
	"github.com/stainless-sdks/cf-rex-go/internal/requestconfig"
	"github.com/stainless-sdks/cf-rex-go/option"
	"github.com/tidwall/gjson"
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
// [Get bulk operation status](/operations/lists-get-bulk-operation-status)
// endpoint with the returned `operation_id`.
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

// Fetches a list item in the list.
func (r *AccountRuleListItemService) Get(ctx context.Context, accountIdentifier string, listID string, itemID string, opts ...option.RequestOption) (res *AccountRuleListItemGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountIdentifier == "" {
		err = errors.New("missing required account_identifier parameter")
		return
	}
	if listID == "" {
		err = errors.New("missing required list_id parameter")
		return
	}
	if itemID == "" {
		err = errors.New("missing required item_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/rules/lists/%s/items/%s", accountIdentifier, listID, itemID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Removes all existing items from the list and adds the provided items to the
// list.
//
// This operation is asynchronous. To get current the operation status, invoke the
// [Get bulk operation status](/operations/lists-get-bulk-operation-status)
// endpoint with the returned `operation_id`.
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
// [Get bulk operation status](/operations/lists-get-bulk-operation-status)
// endpoint with the returned `operation_id`.
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
	Result AsyncResponseResult `json:"result"`
	JSON   asyncResponseJSON   `json:"-"`
	APIResponseCollectionLists
}

// asyncResponseJSON contains the JSON metadata for the struct [AsyncResponse]
type asyncResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AsyncResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r asyncResponseJSON) RawJSON() string {
	return r.raw
}

type AsyncResponseResult struct {
	// The unique operation ID of the asynchronous action.
	OperationID string                  `json:"operation_id"`
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

// An IPv4 address, an IPv4 CIDR, or an IPv6 CIDR. IPv6 CIDRs are limited to a
// maximum of /64.
type ListItem struct {
	// The unique ID of the list.
	ID string `json:"id"`
	// A non-negative 32 bit integer
	Asn int64 `json:"asn"`
	// An informative summary of the list item.
	Comment string `json:"comment"`
	// The RFC 3339 timestamp of when the item was created.
	CreatedOn string `json:"created_on"`
	// Valid characters for hostnames are ASCII(7) letters from a to z, the digits from
	// 0 to 9, wildcards (\*), and the hyphen (-).
	Hostname          ListItemHostname `json:"hostname"`
	IncludeSubdomains bool             `json:"include_subdomains"`
	// An IPv4 address, an IPv4 CIDR, or an IPv6 CIDR. IPv6 CIDRs are limited to a
	// maximum of /64.
	IP string `json:"ip"`
	// The RFC 3339 timestamp of when the item was last modified.
	ModifiedOn          string `json:"modified_on"`
	PreservePathSuffix  bool   `json:"preserve_path_suffix"`
	PreserveQueryString bool   `json:"preserve_query_string"`
	// The definition of the redirect.
	Redirect        ListItemRedirect   `json:"redirect"`
	SourceURL       string             `json:"source_url"`
	StatusCode      ListItemStatusCode `json:"status_code"`
	SubpathMatching bool               `json:"subpath_matching"`
	TargetURL       string             `json:"target_url"`
	URLHostname     string             `json:"url_hostname"`
	JSON            listItemJSON       `json:"-"`
	union           ListItemUnion
}

// listItemJSON contains the JSON metadata for the struct [ListItem]
type listItemJSON struct {
	ID                  apijson.Field
	Asn                 apijson.Field
	Comment             apijson.Field
	CreatedOn           apijson.Field
	Hostname            apijson.Field
	IncludeSubdomains   apijson.Field
	IP                  apijson.Field
	ModifiedOn          apijson.Field
	PreservePathSuffix  apijson.Field
	PreserveQueryString apijson.Field
	Redirect            apijson.Field
	SourceURL           apijson.Field
	StatusCode          apijson.Field
	SubpathMatching     apijson.Field
	TargetURL           apijson.Field
	URLHostname         apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r listItemJSON) RawJSON() string {
	return r.raw
}

func (r *ListItem) UnmarshalJSON(data []byte) (err error) {
	*r = ListItem{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ListItemUnion] interface which you can cast to the specific
// types for more type safety.
//
// Possible runtime types of the union are [ListItemIntersection],
// [ListItemAccountsRulesListsItemsListItemRedirect],
// [ListItemAccountsRulesListsItemsListItemHostname], [ListItemIntersection].
func (r ListItem) AsUnion() ListItemUnion {
	return r.union
}

// An IPv4 address, an IPv4 CIDR, or an IPv6 CIDR. IPv6 CIDRs are limited to a
// maximum of /64.
//
// Union satisfied by [ListItemIntersection],
// [ListItemAccountsRulesListsItemsListItemRedirect],
// [ListItemAccountsRulesListsItemsListItemHostname] or [ListItemIntersection].
type ListItemUnion interface {
	implementsListItem()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ListItemUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ListItemIntersection{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ListItemAccountsRulesListsItemsListItemRedirect{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ListItemAccountsRulesListsItemsListItemHostname{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ListItemIntersection{}),
		},
	)
}

// An IPv4 address, an IPv4 CIDR, or an IPv6 CIDR. IPv6 CIDRs are limited to a
// maximum of /64.
type ListItemIntersection struct {
	JSON listItemIntersectionJSON `json:"-"`
}

// listItemIntersectionJSON contains the JSON metadata for the struct
// [ListItemIntersection]
type listItemIntersectionJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ListItemIntersection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listItemIntersectionJSON) RawJSON() string {
	return r.raw
}

func (r ListItemIntersection) implementsListItem() {}

// The definition of the redirect.
type ListItemAccountsRulesListsItemsListItemRedirect struct {
	// The unique ID of the list.
	ID string `json:"id"`
	// A non-negative 32 bit integer
	Asn int64 `json:"asn"`
	// An informative summary of the list item.
	Comment string `json:"comment"`
	// The RFC 3339 timestamp of when the item was created.
	CreatedOn string `json:"created_on"`
	// Valid characters for hostnames are ASCII(7) letters from a to z, the digits from
	// 0 to 9, wildcards (\*), and the hyphen (-).
	Hostname ListItemHostname `json:"hostname"`
	// An IPv4 address, an IPv4 CIDR, or an IPv6 CIDR. IPv6 CIDRs are limited to a
	// maximum of /64.
	IP string `json:"ip"`
	// The RFC 3339 timestamp of when the item was last modified.
	ModifiedOn string `json:"modified_on"`
	// The definition of the redirect.
	Redirect ListItemRedirect                                    `json:"redirect"`
	JSON     listItemAccountsRulesListsItemsListItemRedirectJSON `json:"-"`
	ListItemRedirect
}

// listItemAccountsRulesListsItemsListItemRedirectJSON contains the JSON metadata
// for the struct [ListItemAccountsRulesListsItemsListItemRedirect]
type listItemAccountsRulesListsItemsListItemRedirectJSON struct {
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

func (r *ListItemAccountsRulesListsItemsListItemRedirect) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listItemAccountsRulesListsItemsListItemRedirectJSON) RawJSON() string {
	return r.raw
}

func (r ListItemAccountsRulesListsItemsListItemRedirect) implementsListItem() {}

// Valid characters for hostnames are ASCII(7) letters from a to z, the digits from
// 0 to 9, wildcards (\*), and the hyphen (-).
type ListItemAccountsRulesListsItemsListItemHostname struct {
	// The unique ID of the list.
	ID string `json:"id"`
	// A non-negative 32 bit integer
	Asn int64 `json:"asn"`
	// An informative summary of the list item.
	Comment string `json:"comment"`
	// The RFC 3339 timestamp of when the item was created.
	CreatedOn string `json:"created_on"`
	// Valid characters for hostnames are ASCII(7) letters from a to z, the digits from
	// 0 to 9, wildcards (\*), and the hyphen (-).
	Hostname ListItemHostname `json:"hostname"`
	// An IPv4 address, an IPv4 CIDR, or an IPv6 CIDR. IPv6 CIDRs are limited to a
	// maximum of /64.
	IP string `json:"ip"`
	// The RFC 3339 timestamp of when the item was last modified.
	ModifiedOn string `json:"modified_on"`
	// The definition of the redirect.
	Redirect ListItemRedirect                                    `json:"redirect"`
	JSON     listItemAccountsRulesListsItemsListItemHostnameJSON `json:"-"`
	ListItemHostname
}

// listItemAccountsRulesListsItemsListItemHostnameJSON contains the JSON metadata
// for the struct [ListItemAccountsRulesListsItemsListItemHostname]
type listItemAccountsRulesListsItemsListItemHostnameJSON struct {
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

func (r *ListItemAccountsRulesListsItemsListItemHostname) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listItemAccountsRulesListsItemsListItemHostnameJSON) RawJSON() string {
	return r.raw
}

func (r ListItemAccountsRulesListsItemsListItemHostname) implementsListItem() {}

type ListItemStatusCode int64

const (
	ListItemStatusCode301 ListItemStatusCode = 301
	ListItemStatusCode302 ListItemStatusCode = 302
	ListItemStatusCode307 ListItemStatusCode = 307
	ListItemStatusCode308 ListItemStatusCode = 308
)

func (r ListItemStatusCode) IsKnown() bool {
	switch r {
	case ListItemStatusCode301, ListItemStatusCode302, ListItemStatusCode307, ListItemStatusCode308:
		return true
	}
	return false
}

// Valid characters for hostnames are ASCII(7) letters from a to z, the digits from
// 0 to 9, wildcards (\*), and the hyphen (-).
type ListItemHostname struct {
	URLHostname string               `json:"url_hostname,required"`
	JSON        listItemHostnameJSON `json:"-"`
}

// listItemHostnameJSON contains the JSON metadata for the struct
// [ListItemHostname]
type listItemHostnameJSON struct {
	URLHostname apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
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
	// A non-negative 32 bit integer
	Asn param.Field[int64] `json:"asn"`
	// An informative summary of the list item.
	Comment param.Field[string] `json:"comment"`
	// Valid characters for hostnames are ASCII(7) letters from a to z, the digits from
	// 0 to 9, wildcards (\*), and the hyphen (-).
	Hostname param.Field[ListItemHostnameParam] `json:"hostname"`
	// An IPv4 address, an IPv4 CIDR, or an IPv6 CIDR. IPv6 CIDRs are limited to a
	// maximum of /64.
	IP param.Field[string] `json:"ip"`
	// The definition of the redirect.
	Redirect param.Field[ListItemRedirectParam] `json:"redirect"`
}

func (r UpdateRequestCollectionItemParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountRuleListItemGetResponse struct {
	// An IPv4 address, an IPv4 CIDR, or an IPv6 CIDR. IPv6 CIDRs are limited to a
	// maximum of /64.
	Result ListItem                           `json:"result"`
	JSON   accountRuleListItemGetResponseJSON `json:"-"`
	APIResponseCollectionLists
}

// accountRuleListItemGetResponseJSON contains the JSON metadata for the struct
// [AccountRuleListItemGetResponse]
type accountRuleListItemGetResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRuleListItemGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountRuleListItemGetResponseJSON) RawJSON() string {
	return r.raw
}

type AccountRuleListItemListResponse struct {
	Result     []ListItem                                `json:"result"`
	ResultInfo AccountRuleListItemListResponseResultInfo `json:"result_info"`
	JSON       accountRuleListItemListResponseJSON       `json:"-"`
	APIResponseCollectionLists
}

// accountRuleListItemListResponseJSON contains the JSON metadata for the struct
// [AccountRuleListItemListResponse]
type accountRuleListItemListResponseJSON struct {
	Result      apijson.Field
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
	Body []UpdateRequestCollectionItemParam `json:"body,required"`
}

func (r AccountRuleListItemNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type AccountRuleListItemUpdateParams struct {
	Body []UpdateRequestCollectionItemParam `json:"body,required"`
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
	// The unique ID of the item in the List.
	ID param.Field[string] `json:"id"`
}

func (r AccountRuleListItemDeleteParamsItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
