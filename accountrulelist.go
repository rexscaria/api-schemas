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

// Creates a new list of the specified kind.
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
func (r *AccountRuleListService) Delete(ctx context.Context, accountID string, listID string, opts ...option.RequestOption) (res *AccountRuleListDeleteResponse, err error) {
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
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
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

type ListResponseCollection struct {
	Errors   []ListResponseCollectionError   `json:"errors,required"`
	Messages []ListResponseCollectionMessage `json:"messages,required"`
	Result   ListResponseCollectionResult    `json:"result,required"`
	// Defines whether the API call was successful.
	Success ListResponseCollectionSuccess `json:"success,required"`
	JSON    listResponseCollectionJSON    `json:"-"`
}

// listResponseCollectionJSON contains the JSON metadata for the struct
// [ListResponseCollection]
type listResponseCollectionJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ListResponseCollection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listResponseCollectionJSON) RawJSON() string {
	return r.raw
}

type ListResponseCollectionError struct {
	Code             int64                              `json:"code,required"`
	Message          string                             `json:"message,required"`
	DocumentationURL string                             `json:"documentation_url"`
	Source           ListResponseCollectionErrorsSource `json:"source"`
	JSON             listResponseCollectionErrorJSON    `json:"-"`
}

// listResponseCollectionErrorJSON contains the JSON metadata for the struct
// [ListResponseCollectionError]
type listResponseCollectionErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ListResponseCollectionError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listResponseCollectionErrorJSON) RawJSON() string {
	return r.raw
}

type ListResponseCollectionErrorsSource struct {
	Pointer string                                 `json:"pointer"`
	JSON    listResponseCollectionErrorsSourceJSON `json:"-"`
}

// listResponseCollectionErrorsSourceJSON contains the JSON metadata for the struct
// [ListResponseCollectionErrorsSource]
type listResponseCollectionErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ListResponseCollectionErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listResponseCollectionErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type ListResponseCollectionMessage struct {
	Code             int64                                `json:"code,required"`
	Message          string                               `json:"message,required"`
	DocumentationURL string                               `json:"documentation_url"`
	Source           ListResponseCollectionMessagesSource `json:"source"`
	JSON             listResponseCollectionMessageJSON    `json:"-"`
}

// listResponseCollectionMessageJSON contains the JSON metadata for the struct
// [ListResponseCollectionMessage]
type listResponseCollectionMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ListResponseCollectionMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listResponseCollectionMessageJSON) RawJSON() string {
	return r.raw
}

type ListResponseCollectionMessagesSource struct {
	Pointer string                                   `json:"pointer"`
	JSON    listResponseCollectionMessagesSourceJSON `json:"-"`
}

// listResponseCollectionMessagesSourceJSON contains the JSON metadata for the
// struct [ListResponseCollectionMessagesSource]
type listResponseCollectionMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ListResponseCollectionMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listResponseCollectionMessagesSourceJSON) RawJSON() string {
	return r.raw
}

type ListResponseCollectionResult struct {
	// The unique ID of the list.
	ID string `json:"id,required"`
	// The RFC 3339 timestamp of when the list was created.
	CreatedOn string `json:"created_on,required"`
	// The type of the list. Each type supports specific list items (IP addresses,
	// ASNs, hostnames or redirects).
	Kind ListKind `json:"kind,required"`
	// The RFC 3339 timestamp of when the list was last modified.
	ModifiedOn string `json:"modified_on,required"`
	// An informative name for the list. Use this name in filter and rule expressions.
	Name string `json:"name,required"`
	// The number of items in the list.
	NumItems float64 `json:"num_items,required"`
	// The number of [filters](/api/resources/filters/) referencing the list.
	NumReferencingFilters float64 `json:"num_referencing_filters,required"`
	// An informative summary of the list.
	Description string                           `json:"description"`
	JSON        listResponseCollectionResultJSON `json:"-"`
}

// listResponseCollectionResultJSON contains the JSON metadata for the struct
// [ListResponseCollectionResult]
type listResponseCollectionResultJSON struct {
	ID                    apijson.Field
	CreatedOn             apijson.Field
	Kind                  apijson.Field
	ModifiedOn            apijson.Field
	Name                  apijson.Field
	NumItems              apijson.Field
	NumReferencingFilters apijson.Field
	Description           apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *ListResponseCollectionResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listResponseCollectionResultJSON) RawJSON() string {
	return r.raw
}

// Defines whether the API call was successful.
type ListResponseCollectionSuccess bool

const (
	ListResponseCollectionSuccessTrue ListResponseCollectionSuccess = true
)

func (r ListResponseCollectionSuccess) IsKnown() bool {
	switch r {
	case ListResponseCollectionSuccessTrue:
		return true
	}
	return false
}

type ListRules struct {
	// The unique ID of the list.
	ID string `json:"id,required"`
	// The RFC 3339 timestamp of when the list was created.
	CreatedOn string `json:"created_on,required"`
	// The type of the list. Each type supports specific list items (IP addresses,
	// ASNs, hostnames or redirects).
	Kind ListKind `json:"kind,required"`
	// The RFC 3339 timestamp of when the list was last modified.
	ModifiedOn string `json:"modified_on,required"`
	// An informative name for the list. Use this name in filter and rule expressions.
	Name string `json:"name,required"`
	// The number of items in the list.
	NumItems float64 `json:"num_items,required"`
	// The number of [filters](/api/resources/filters/) referencing the list.
	NumReferencingFilters float64 `json:"num_referencing_filters,required"`
	// An informative summary of the list.
	Description string        `json:"description"`
	JSON        listRulesJSON `json:"-"`
}

// listRulesJSON contains the JSON metadata for the struct [ListRules]
type listRulesJSON struct {
	ID                    apijson.Field
	CreatedOn             apijson.Field
	Kind                  apijson.Field
	ModifiedOn            apijson.Field
	Name                  apijson.Field
	NumItems              apijson.Field
	NumReferencingFilters apijson.Field
	Description           apijson.Field
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
	Errors   []AccountRuleListListResponseError   `json:"errors,required"`
	Messages []AccountRuleListListResponseMessage `json:"messages,required"`
	Result   []ListRules                          `json:"result,required"`
	// Defines whether the API call was successful.
	Success AccountRuleListListResponseSuccess `json:"success,required"`
	JSON    accountRuleListListResponseJSON    `json:"-"`
}

// accountRuleListListResponseJSON contains the JSON metadata for the struct
// [AccountRuleListListResponse]
type accountRuleListListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRuleListListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountRuleListListResponseJSON) RawJSON() string {
	return r.raw
}

type AccountRuleListListResponseError struct {
	Code             int64                                   `json:"code,required"`
	Message          string                                  `json:"message,required"`
	DocumentationURL string                                  `json:"documentation_url"`
	Source           AccountRuleListListResponseErrorsSource `json:"source"`
	JSON             accountRuleListListResponseErrorJSON    `json:"-"`
}

// accountRuleListListResponseErrorJSON contains the JSON metadata for the struct
// [AccountRuleListListResponseError]
type accountRuleListListResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountRuleListListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountRuleListListResponseErrorJSON) RawJSON() string {
	return r.raw
}

type AccountRuleListListResponseErrorsSource struct {
	Pointer string                                      `json:"pointer"`
	JSON    accountRuleListListResponseErrorsSourceJSON `json:"-"`
}

// accountRuleListListResponseErrorsSourceJSON contains the JSON metadata for the
// struct [AccountRuleListListResponseErrorsSource]
type accountRuleListListResponseErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRuleListListResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountRuleListListResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AccountRuleListListResponseMessage struct {
	Code             int64                                     `json:"code,required"`
	Message          string                                    `json:"message,required"`
	DocumentationURL string                                    `json:"documentation_url"`
	Source           AccountRuleListListResponseMessagesSource `json:"source"`
	JSON             accountRuleListListResponseMessageJSON    `json:"-"`
}

// accountRuleListListResponseMessageJSON contains the JSON metadata for the struct
// [AccountRuleListListResponseMessage]
type accountRuleListListResponseMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountRuleListListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountRuleListListResponseMessageJSON) RawJSON() string {
	return r.raw
}

type AccountRuleListListResponseMessagesSource struct {
	Pointer string                                        `json:"pointer"`
	JSON    accountRuleListListResponseMessagesSourceJSON `json:"-"`
}

// accountRuleListListResponseMessagesSourceJSON contains the JSON metadata for the
// struct [AccountRuleListListResponseMessagesSource]
type accountRuleListListResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRuleListListResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountRuleListListResponseMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Defines whether the API call was successful.
type AccountRuleListListResponseSuccess bool

const (
	AccountRuleListListResponseSuccessTrue AccountRuleListListResponseSuccess = true
)

func (r AccountRuleListListResponseSuccess) IsKnown() bool {
	switch r {
	case AccountRuleListListResponseSuccessTrue:
		return true
	}
	return false
}

type AccountRuleListDeleteResponse struct {
	Errors   []AccountRuleListDeleteResponseError   `json:"errors,required"`
	Messages []AccountRuleListDeleteResponseMessage `json:"messages,required"`
	Result   AccountRuleListDeleteResponseResult    `json:"result,required"`
	// Defines whether the API call was successful.
	Success AccountRuleListDeleteResponseSuccess `json:"success,required"`
	JSON    accountRuleListDeleteResponseJSON    `json:"-"`
}

// accountRuleListDeleteResponseJSON contains the JSON metadata for the struct
// [AccountRuleListDeleteResponse]
type accountRuleListDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRuleListDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountRuleListDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type AccountRuleListDeleteResponseError struct {
	Code             int64                                     `json:"code,required"`
	Message          string                                    `json:"message,required"`
	DocumentationURL string                                    `json:"documentation_url"`
	Source           AccountRuleListDeleteResponseErrorsSource `json:"source"`
	JSON             accountRuleListDeleteResponseErrorJSON    `json:"-"`
}

// accountRuleListDeleteResponseErrorJSON contains the JSON metadata for the struct
// [AccountRuleListDeleteResponseError]
type accountRuleListDeleteResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountRuleListDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountRuleListDeleteResponseErrorJSON) RawJSON() string {
	return r.raw
}

type AccountRuleListDeleteResponseErrorsSource struct {
	Pointer string                                        `json:"pointer"`
	JSON    accountRuleListDeleteResponseErrorsSourceJSON `json:"-"`
}

// accountRuleListDeleteResponseErrorsSourceJSON contains the JSON metadata for the
// struct [AccountRuleListDeleteResponseErrorsSource]
type accountRuleListDeleteResponseErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRuleListDeleteResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountRuleListDeleteResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AccountRuleListDeleteResponseMessage struct {
	Code             int64                                       `json:"code,required"`
	Message          string                                      `json:"message,required"`
	DocumentationURL string                                      `json:"documentation_url"`
	Source           AccountRuleListDeleteResponseMessagesSource `json:"source"`
	JSON             accountRuleListDeleteResponseMessageJSON    `json:"-"`
}

// accountRuleListDeleteResponseMessageJSON contains the JSON metadata for the
// struct [AccountRuleListDeleteResponseMessage]
type accountRuleListDeleteResponseMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountRuleListDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountRuleListDeleteResponseMessageJSON) RawJSON() string {
	return r.raw
}

type AccountRuleListDeleteResponseMessagesSource struct {
	Pointer string                                          `json:"pointer"`
	JSON    accountRuleListDeleteResponseMessagesSourceJSON `json:"-"`
}

// accountRuleListDeleteResponseMessagesSourceJSON contains the JSON metadata for
// the struct [AccountRuleListDeleteResponseMessagesSource]
type accountRuleListDeleteResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRuleListDeleteResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountRuleListDeleteResponseMessagesSourceJSON) RawJSON() string {
	return r.raw
}

type AccountRuleListDeleteResponseResult struct {
	// The unique ID of the list.
	ID   string                                  `json:"id,required"`
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

// Defines whether the API call was successful.
type AccountRuleListDeleteResponseSuccess bool

const (
	AccountRuleListDeleteResponseSuccessTrue AccountRuleListDeleteResponseSuccess = true
)

func (r AccountRuleListDeleteResponseSuccess) IsKnown() bool {
	switch r {
	case AccountRuleListDeleteResponseSuccessTrue:
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
