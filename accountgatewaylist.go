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

// AccountGatewayListService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountGatewayListService] method instead.
type AccountGatewayListService struct {
	Options []option.RequestOption
}

// NewAccountGatewayListService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountGatewayListService(opts ...option.RequestOption) (r *AccountGatewayListService) {
	r = &AccountGatewayListService{}
	r.Options = opts
	return
}

// Creates a new Zero Trust list.
func (r *AccountGatewayListService) New(ctx context.Context, accountID string, body AccountGatewayListNewParams, opts ...option.RequestOption) (res *AccountGatewayListNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/gateway/lists", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Fetches a single Zero Trust list.
func (r *AccountGatewayListService) Get(ctx context.Context, accountID string, listID string, opts ...option.RequestOption) (res *SchemasZeroTrustGatewaySingleResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if listID == "" {
		err = errors.New("missing required list_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/gateway/lists/%s", accountID, listID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates a configured Zero Trust list. Skips updating list items if not included
// in the payload.
func (r *AccountGatewayListService) Update(ctx context.Context, accountID string, listID string, body AccountGatewayListUpdateParams, opts ...option.RequestOption) (res *SchemasZeroTrustGatewaySingleResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if listID == "" {
		err = errors.New("missing required list_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/gateway/lists/%s", accountID, listID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Fetches all Zero Trust lists for an account.
func (r *AccountGatewayListService) List(ctx context.Context, accountID string, query AccountGatewayListListParams, opts ...option.RequestOption) (res *AccountGatewayListListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/gateway/lists", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Deletes a Zero Trust list.
func (r *AccountGatewayListService) Delete(ctx context.Context, accountID string, listID string, body AccountGatewayListDeleteParams, opts ...option.RequestOption) (res *ZeroTrustGatewayEmptyResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if listID == "" {
		err = errors.New("missing required list_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/gateway/lists/%s", accountID, listID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

// Fetches all items in a single Zero Trust list.
func (r *AccountGatewayListService) ListItems(ctx context.Context, accountID string, listID string, opts ...option.RequestOption) (res *AccountGatewayListListItemsResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if listID == "" {
		err = errors.New("missing required list_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/gateway/lists/%s/items", accountID, listID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Appends or removes an item from a configured Zero Trust list.
func (r *AccountGatewayListService) Patch(ctx context.Context, accountID string, listID string, body AccountGatewayListPatchParams, opts ...option.RequestOption) (res *SchemasZeroTrustGatewaySingleResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if listID == "" {
		err = errors.New("missing required list_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/gateway/lists/%s", accountID, listID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

type Items struct {
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The description of the list item, if present
	Description string `json:"description"`
	// The value of the item in a list.
	Value string    `json:"value"`
	JSON  itemsJSON `json:"-"`
}

// itemsJSON contains the JSON metadata for the struct [Items]
type itemsJSON struct {
	CreatedAt   apijson.Field
	Description apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Items) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r itemsJSON) RawJSON() string {
	return r.raw
}

type ItemsParam struct {
	CreatedAt param.Field[time.Time] `json:"created_at" format:"date-time"`
	// The description of the list item, if present
	Description param.Field[string] `json:"description"`
	// The value of the item in a list.
	Value param.Field[string] `json:"value"`
}

func (r ItemsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ListGateway struct {
	// API Resource UUID tag.
	ID string `json:"id"`
	// The number of items in the list.
	Count     float64   `json:"count"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The description of the list.
	Description string `json:"description"`
	// The name of the list.
	Name string `json:"name"`
	// The type of list.
	Type      SchemasListType `json:"type"`
	UpdatedAt time.Time       `json:"updated_at" format:"date-time"`
	JSON      listGatewayJSON `json:"-"`
}

// listGatewayJSON contains the JSON metadata for the struct [ListGateway]
type listGatewayJSON struct {
	ID          apijson.Field
	Count       apijson.Field
	CreatedAt   apijson.Field
	Description apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ListGateway) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listGatewayJSON) RawJSON() string {
	return r.raw
}

// The type of list.
type SchemasListType string

const (
	SchemasListTypeSerial SchemasListType = "SERIAL"
	SchemasListTypeURL    SchemasListType = "URL"
	SchemasListTypeDomain SchemasListType = "DOMAIN"
	SchemasListTypeEmail  SchemasListType = "EMAIL"
	SchemasListTypeIP     SchemasListType = "IP"
)

func (r SchemasListType) IsKnown() bool {
	switch r {
	case SchemasListTypeSerial, SchemasListTypeURL, SchemasListTypeDomain, SchemasListTypeEmail, SchemasListTypeIP:
		return true
	}
	return false
}

type SchemasZeroTrustGatewaySingleResponse struct {
	Result ListGateway                               `json:"result"`
	JSON   schemasZeroTrustGatewaySingleResponseJSON `json:"-"`
	APIResponseSingleZeroTrustGateway
}

// schemasZeroTrustGatewaySingleResponseJSON contains the JSON metadata for the
// struct [SchemasZeroTrustGatewaySingleResponse]
type schemasZeroTrustGatewaySingleResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasZeroTrustGatewaySingleResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r schemasZeroTrustGatewaySingleResponseJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustGatewayEmptyResponse struct {
	Result interface{}                       `json:"result"`
	JSON   zeroTrustGatewayEmptyResponseJSON `json:"-"`
	APIResponseSingleZeroTrustGateway
}

// zeroTrustGatewayEmptyResponseJSON contains the JSON metadata for the struct
// [ZeroTrustGatewayEmptyResponse]
type zeroTrustGatewayEmptyResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayEmptyResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGatewayEmptyResponseJSON) RawJSON() string {
	return r.raw
}

type AccountGatewayListNewResponse struct {
	Result AccountGatewayListNewResponseResult `json:"result"`
	JSON   accountGatewayListNewResponseJSON   `json:"-"`
	APIResponseSingleZeroTrustGateway
}

// accountGatewayListNewResponseJSON contains the JSON metadata for the struct
// [AccountGatewayListNewResponse]
type accountGatewayListNewResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayListNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountGatewayListNewResponseJSON) RawJSON() string {
	return r.raw
}

type AccountGatewayListNewResponseResult struct {
	// API Resource UUID tag.
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The description of the list.
	Description string `json:"description"`
	// The items in the list.
	Items []Items `json:"items"`
	// The name of the list.
	Name string `json:"name"`
	// The type of list.
	Type      SchemasListType                         `json:"type"`
	UpdatedAt time.Time                               `json:"updated_at" format:"date-time"`
	JSON      accountGatewayListNewResponseResultJSON `json:"-"`
}

// accountGatewayListNewResponseResultJSON contains the JSON metadata for the
// struct [AccountGatewayListNewResponseResult]
type accountGatewayListNewResponseResultJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Description apijson.Field
	Items       apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayListNewResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountGatewayListNewResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountGatewayListListResponse struct {
	Result []ListGateway                      `json:"result"`
	JSON   accountGatewayListListResponseJSON `json:"-"`
	APIResponseCollectionZeroTrustGateway
}

// accountGatewayListListResponseJSON contains the JSON metadata for the struct
// [AccountGatewayListListResponse]
type accountGatewayListListResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayListListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountGatewayListListResponseJSON) RawJSON() string {
	return r.raw
}

type AccountGatewayListListItemsResponse struct {
	Result     [][]Items                                     `json:"result"`
	ResultInfo AccountGatewayListListItemsResponseResultInfo `json:"result_info"`
	JSON       accountGatewayListListItemsResponseJSON       `json:"-"`
	APIResponseCollectionZeroTrustGateway
}

// accountGatewayListListItemsResponseJSON contains the JSON metadata for the
// struct [AccountGatewayListListItemsResponse]
type accountGatewayListListItemsResponseJSON struct {
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayListListItemsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountGatewayListListItemsResponseJSON) RawJSON() string {
	return r.raw
}

type AccountGatewayListListItemsResponseResultInfo struct {
	// Total results returned based on your search parameters.
	Count float64 `json:"count"`
	// Current page within paginated list of results.
	Page float64 `json:"page"`
	// Number of results per page of results.
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters.
	TotalCount float64                                           `json:"total_count"`
	JSON       accountGatewayListListItemsResponseResultInfoJSON `json:"-"`
}

// accountGatewayListListItemsResponseResultInfoJSON contains the JSON metadata for
// the struct [AccountGatewayListListItemsResponseResultInfo]
type accountGatewayListListItemsResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayListListItemsResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountGatewayListListItemsResponseResultInfoJSON) RawJSON() string {
	return r.raw
}

type AccountGatewayListNewParams struct {
	// The name of the list.
	Name param.Field[string] `json:"name,required"`
	// The type of list.
	Type param.Field[SchemasListType] `json:"type,required"`
	// The description of the list.
	Description param.Field[string] `json:"description"`
	// The items in the list.
	Items param.Field[[]ItemsParam] `json:"items"`
}

func (r AccountGatewayListNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountGatewayListUpdateParams struct {
	// The name of the list.
	Name param.Field[string] `json:"name,required"`
	// The description of the list.
	Description param.Field[string] `json:"description"`
	// The items in the list.
	Items param.Field[[]ItemsParam] `json:"items"`
}

func (r AccountGatewayListUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountGatewayListListParams struct {
	// The type of list.
	Type param.Field[SchemasListType] `query:"type"`
}

// URLQuery serializes [AccountGatewayListListParams]'s query parameters as
// `url.Values`.
func (r AccountGatewayListListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountGatewayListDeleteParams struct {
	Body interface{} `json:"body,required"`
}

func (r AccountGatewayListDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type AccountGatewayListPatchParams struct {
	// The items in the list.
	Append param.Field[[]ItemsParam] `json:"append"`
	// A list of the item values you want to remove.
	Remove param.Field[[]string] `json:"remove"`
}

func (r AccountGatewayListPatchParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
