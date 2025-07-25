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
func (r *AccountGatewayListService) Get(ctx context.Context, accountID string, listID string, opts ...option.RequestOption) (res *AccountGatewayListGetResponse, err error) {
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
func (r *AccountGatewayListService) Update(ctx context.Context, accountID string, listID string, body AccountGatewayListUpdateParams, opts ...option.RequestOption) (res *AccountGatewayListUpdateResponse, err error) {
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
func (r *AccountGatewayListService) Delete(ctx context.Context, accountID string, listID string, opts ...option.RequestOption) (res *ZeroTrustGatewayEmptyResponse, err error) {
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
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
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
func (r *AccountGatewayListService) Patch(ctx context.Context, accountID string, listID string, body AccountGatewayListPatchParams, opts ...option.RequestOption) (res *AccountGatewayListPatchResponse, err error) {
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

type ListGateway struct {
	// API Resource UUID tag.
	ID string `json:"id"`
	// The number of items in the list.
	Count     float64   `json:"count"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The description of the list.
	Description string `json:"description"`
	// The items in the list.
	Items []Items `json:"items"`
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
	Items       apijson.Field
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
	Errors   []ZeroTrustGatewayMessages `json:"errors,required"`
	Messages []ZeroTrustGatewayMessages `json:"messages,required"`
	// Whether the API call was successful
	Success SchemasZeroTrustGatewaySingleResponseSuccess `json:"success,required"`
	Result  Location                                     `json:"result"`
	JSON    schemasZeroTrustGatewaySingleResponseJSON    `json:"-"`
}

// schemasZeroTrustGatewaySingleResponseJSON contains the JSON metadata for the
// struct [SchemasZeroTrustGatewaySingleResponse]
type schemasZeroTrustGatewaySingleResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
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

// Whether the API call was successful
type SchemasZeroTrustGatewaySingleResponseSuccess bool

const (
	SchemasZeroTrustGatewaySingleResponseSuccessTrue SchemasZeroTrustGatewaySingleResponseSuccess = true
)

func (r SchemasZeroTrustGatewaySingleResponseSuccess) IsKnown() bool {
	switch r {
	case SchemasZeroTrustGatewaySingleResponseSuccessTrue:
		return true
	}
	return false
}

type ZeroTrustGatewayEmptyResponse struct {
	Errors   []ZeroTrustGatewayEmptyResponseError   `json:"errors,required"`
	Messages []ZeroTrustGatewayEmptyResponseMessage `json:"messages,required"`
	// Whether the API call was successful
	Success ZeroTrustGatewayEmptyResponseSuccess `json:"success,required"`
	Result  interface{}                          `json:"result"`
	JSON    zeroTrustGatewayEmptyResponseJSON    `json:"-"`
}

// zeroTrustGatewayEmptyResponseJSON contains the JSON metadata for the struct
// [ZeroTrustGatewayEmptyResponse]
type zeroTrustGatewayEmptyResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
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

type ZeroTrustGatewayEmptyResponseError struct {
	Code             int64                                     `json:"code,required"`
	Message          string                                    `json:"message,required"`
	DocumentationURL string                                    `json:"documentation_url"`
	Source           ZeroTrustGatewayEmptyResponseErrorsSource `json:"source"`
	JSON             zeroTrustGatewayEmptyResponseErrorJSON    `json:"-"`
}

// zeroTrustGatewayEmptyResponseErrorJSON contains the JSON metadata for the struct
// [ZeroTrustGatewayEmptyResponseError]
type zeroTrustGatewayEmptyResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZeroTrustGatewayEmptyResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGatewayEmptyResponseErrorJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustGatewayEmptyResponseErrorsSource struct {
	Pointer string                                        `json:"pointer"`
	JSON    zeroTrustGatewayEmptyResponseErrorsSourceJSON `json:"-"`
}

// zeroTrustGatewayEmptyResponseErrorsSourceJSON contains the JSON metadata for the
// struct [ZeroTrustGatewayEmptyResponseErrorsSource]
type zeroTrustGatewayEmptyResponseErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayEmptyResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGatewayEmptyResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustGatewayEmptyResponseMessage struct {
	Code             int64                                       `json:"code,required"`
	Message          string                                      `json:"message,required"`
	DocumentationURL string                                      `json:"documentation_url"`
	Source           ZeroTrustGatewayEmptyResponseMessagesSource `json:"source"`
	JSON             zeroTrustGatewayEmptyResponseMessageJSON    `json:"-"`
}

// zeroTrustGatewayEmptyResponseMessageJSON contains the JSON metadata for the
// struct [ZeroTrustGatewayEmptyResponseMessage]
type zeroTrustGatewayEmptyResponseMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZeroTrustGatewayEmptyResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGatewayEmptyResponseMessageJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustGatewayEmptyResponseMessagesSource struct {
	Pointer string                                          `json:"pointer"`
	JSON    zeroTrustGatewayEmptyResponseMessagesSourceJSON `json:"-"`
}

// zeroTrustGatewayEmptyResponseMessagesSourceJSON contains the JSON metadata for
// the struct [ZeroTrustGatewayEmptyResponseMessagesSource]
type zeroTrustGatewayEmptyResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayEmptyResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGatewayEmptyResponseMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ZeroTrustGatewayEmptyResponseSuccess bool

const (
	ZeroTrustGatewayEmptyResponseSuccessTrue ZeroTrustGatewayEmptyResponseSuccess = true
)

func (r ZeroTrustGatewayEmptyResponseSuccess) IsKnown() bool {
	switch r {
	case ZeroTrustGatewayEmptyResponseSuccessTrue:
		return true
	}
	return false
}

type AccountGatewayListNewResponse struct {
	Errors   []ZeroTrustGatewayMessages `json:"errors,required"`
	Messages []ZeroTrustGatewayMessages `json:"messages,required"`
	// Whether the API call was successful
	Success AccountGatewayListNewResponseSuccess `json:"success,required"`
	Result  AccountGatewayListNewResponseResult  `json:"result"`
	JSON    accountGatewayListNewResponseJSON    `json:"-"`
}

// accountGatewayListNewResponseJSON contains the JSON metadata for the struct
// [AccountGatewayListNewResponse]
type accountGatewayListNewResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
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

// Whether the API call was successful
type AccountGatewayListNewResponseSuccess bool

const (
	AccountGatewayListNewResponseSuccessTrue AccountGatewayListNewResponseSuccess = true
)

func (r AccountGatewayListNewResponseSuccess) IsKnown() bool {
	switch r {
	case AccountGatewayListNewResponseSuccessTrue:
		return true
	}
	return false
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

type AccountGatewayListGetResponse struct {
	Errors   []AccountGatewayListGetResponseError   `json:"errors,required"`
	Messages []AccountGatewayListGetResponseMessage `json:"messages,required"`
	// Whether the API call was successful
	Success AccountGatewayListGetResponseSuccess `json:"success,required"`
	Result  ListGateway                          `json:"result"`
	JSON    accountGatewayListGetResponseJSON    `json:"-"`
}

// accountGatewayListGetResponseJSON contains the JSON metadata for the struct
// [AccountGatewayListGetResponse]
type accountGatewayListGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayListGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountGatewayListGetResponseJSON) RawJSON() string {
	return r.raw
}

type AccountGatewayListGetResponseError struct {
	Code             int64                                     `json:"code,required"`
	Message          string                                    `json:"message,required"`
	DocumentationURL string                                    `json:"documentation_url"`
	Source           AccountGatewayListGetResponseErrorsSource `json:"source"`
	JSON             accountGatewayListGetResponseErrorJSON    `json:"-"`
}

// accountGatewayListGetResponseErrorJSON contains the JSON metadata for the struct
// [AccountGatewayListGetResponseError]
type accountGatewayListGetResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountGatewayListGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountGatewayListGetResponseErrorJSON) RawJSON() string {
	return r.raw
}

type AccountGatewayListGetResponseErrorsSource struct {
	Pointer string                                        `json:"pointer"`
	JSON    accountGatewayListGetResponseErrorsSourceJSON `json:"-"`
}

// accountGatewayListGetResponseErrorsSourceJSON contains the JSON metadata for the
// struct [AccountGatewayListGetResponseErrorsSource]
type accountGatewayListGetResponseErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayListGetResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountGatewayListGetResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AccountGatewayListGetResponseMessage struct {
	Code             int64                                       `json:"code,required"`
	Message          string                                      `json:"message,required"`
	DocumentationURL string                                      `json:"documentation_url"`
	Source           AccountGatewayListGetResponseMessagesSource `json:"source"`
	JSON             accountGatewayListGetResponseMessageJSON    `json:"-"`
}

// accountGatewayListGetResponseMessageJSON contains the JSON metadata for the
// struct [AccountGatewayListGetResponseMessage]
type accountGatewayListGetResponseMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountGatewayListGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountGatewayListGetResponseMessageJSON) RawJSON() string {
	return r.raw
}

type AccountGatewayListGetResponseMessagesSource struct {
	Pointer string                                          `json:"pointer"`
	JSON    accountGatewayListGetResponseMessagesSourceJSON `json:"-"`
}

// accountGatewayListGetResponseMessagesSourceJSON contains the JSON metadata for
// the struct [AccountGatewayListGetResponseMessagesSource]
type accountGatewayListGetResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayListGetResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountGatewayListGetResponseMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AccountGatewayListGetResponseSuccess bool

const (
	AccountGatewayListGetResponseSuccessTrue AccountGatewayListGetResponseSuccess = true
)

func (r AccountGatewayListGetResponseSuccess) IsKnown() bool {
	switch r {
	case AccountGatewayListGetResponseSuccessTrue:
		return true
	}
	return false
}

type AccountGatewayListUpdateResponse struct {
	Errors   []AccountGatewayListUpdateResponseError   `json:"errors,required"`
	Messages []AccountGatewayListUpdateResponseMessage `json:"messages,required"`
	// Whether the API call was successful
	Success AccountGatewayListUpdateResponseSuccess `json:"success,required"`
	Result  ListGateway                             `json:"result"`
	JSON    accountGatewayListUpdateResponseJSON    `json:"-"`
}

// accountGatewayListUpdateResponseJSON contains the JSON metadata for the struct
// [AccountGatewayListUpdateResponse]
type accountGatewayListUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayListUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountGatewayListUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type AccountGatewayListUpdateResponseError struct {
	Code             int64                                        `json:"code,required"`
	Message          string                                       `json:"message,required"`
	DocumentationURL string                                       `json:"documentation_url"`
	Source           AccountGatewayListUpdateResponseErrorsSource `json:"source"`
	JSON             accountGatewayListUpdateResponseErrorJSON    `json:"-"`
}

// accountGatewayListUpdateResponseErrorJSON contains the JSON metadata for the
// struct [AccountGatewayListUpdateResponseError]
type accountGatewayListUpdateResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountGatewayListUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountGatewayListUpdateResponseErrorJSON) RawJSON() string {
	return r.raw
}

type AccountGatewayListUpdateResponseErrorsSource struct {
	Pointer string                                           `json:"pointer"`
	JSON    accountGatewayListUpdateResponseErrorsSourceJSON `json:"-"`
}

// accountGatewayListUpdateResponseErrorsSourceJSON contains the JSON metadata for
// the struct [AccountGatewayListUpdateResponseErrorsSource]
type accountGatewayListUpdateResponseErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayListUpdateResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountGatewayListUpdateResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AccountGatewayListUpdateResponseMessage struct {
	Code             int64                                          `json:"code,required"`
	Message          string                                         `json:"message,required"`
	DocumentationURL string                                         `json:"documentation_url"`
	Source           AccountGatewayListUpdateResponseMessagesSource `json:"source"`
	JSON             accountGatewayListUpdateResponseMessageJSON    `json:"-"`
}

// accountGatewayListUpdateResponseMessageJSON contains the JSON metadata for the
// struct [AccountGatewayListUpdateResponseMessage]
type accountGatewayListUpdateResponseMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountGatewayListUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountGatewayListUpdateResponseMessageJSON) RawJSON() string {
	return r.raw
}

type AccountGatewayListUpdateResponseMessagesSource struct {
	Pointer string                                             `json:"pointer"`
	JSON    accountGatewayListUpdateResponseMessagesSourceJSON `json:"-"`
}

// accountGatewayListUpdateResponseMessagesSourceJSON contains the JSON metadata
// for the struct [AccountGatewayListUpdateResponseMessagesSource]
type accountGatewayListUpdateResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayListUpdateResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountGatewayListUpdateResponseMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AccountGatewayListUpdateResponseSuccess bool

const (
	AccountGatewayListUpdateResponseSuccessTrue AccountGatewayListUpdateResponseSuccess = true
)

func (r AccountGatewayListUpdateResponseSuccess) IsKnown() bool {
	switch r {
	case AccountGatewayListUpdateResponseSuccessTrue:
		return true
	}
	return false
}

type AccountGatewayListListResponse struct {
	Errors   []ZeroTrustGatewayMessages `json:"errors,required"`
	Messages []ZeroTrustGatewayMessages `json:"messages,required"`
	// Whether the API call was successful
	Success    AccountGatewayListListResponseSuccess    `json:"success,required"`
	Result     []ListGateway                            `json:"result"`
	ResultInfo AccountGatewayListListResponseResultInfo `json:"result_info"`
	JSON       accountGatewayListListResponseJSON       `json:"-"`
}

// accountGatewayListListResponseJSON contains the JSON metadata for the struct
// [AccountGatewayListListResponse]
type accountGatewayListListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayListListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountGatewayListListResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AccountGatewayListListResponseSuccess bool

const (
	AccountGatewayListListResponseSuccessTrue AccountGatewayListListResponseSuccess = true
)

func (r AccountGatewayListListResponseSuccess) IsKnown() bool {
	switch r {
	case AccountGatewayListListResponseSuccessTrue:
		return true
	}
	return false
}

type AccountGatewayListListResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                      `json:"total_count"`
	JSON       accountGatewayListListResponseResultInfoJSON `json:"-"`
}

// accountGatewayListListResponseResultInfoJSON contains the JSON metadata for the
// struct [AccountGatewayListListResponseResultInfo]
type accountGatewayListListResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayListListResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountGatewayListListResponseResultInfoJSON) RawJSON() string {
	return r.raw
}

type AccountGatewayListListItemsResponse struct {
	Errors   []AccountGatewayListListItemsResponseError   `json:"errors,required"`
	Messages []AccountGatewayListListItemsResponseMessage `json:"messages,required"`
	// Whether the API call was successful
	Success    AccountGatewayListListItemsResponseSuccess    `json:"success,required"`
	Result     [][]Items                                     `json:"result"`
	ResultInfo AccountGatewayListListItemsResponseResultInfo `json:"result_info"`
	JSON       accountGatewayListListItemsResponseJSON       `json:"-"`
}

// accountGatewayListListItemsResponseJSON contains the JSON metadata for the
// struct [AccountGatewayListListItemsResponse]
type accountGatewayListListItemsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
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

type AccountGatewayListListItemsResponseError struct {
	Code             int64                                           `json:"code,required"`
	Message          string                                          `json:"message,required"`
	DocumentationURL string                                          `json:"documentation_url"`
	Source           AccountGatewayListListItemsResponseErrorsSource `json:"source"`
	JSON             accountGatewayListListItemsResponseErrorJSON    `json:"-"`
}

// accountGatewayListListItemsResponseErrorJSON contains the JSON metadata for the
// struct [AccountGatewayListListItemsResponseError]
type accountGatewayListListItemsResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountGatewayListListItemsResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountGatewayListListItemsResponseErrorJSON) RawJSON() string {
	return r.raw
}

type AccountGatewayListListItemsResponseErrorsSource struct {
	Pointer string                                              `json:"pointer"`
	JSON    accountGatewayListListItemsResponseErrorsSourceJSON `json:"-"`
}

// accountGatewayListListItemsResponseErrorsSourceJSON contains the JSON metadata
// for the struct [AccountGatewayListListItemsResponseErrorsSource]
type accountGatewayListListItemsResponseErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayListListItemsResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountGatewayListListItemsResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AccountGatewayListListItemsResponseMessage struct {
	Code             int64                                             `json:"code,required"`
	Message          string                                            `json:"message,required"`
	DocumentationURL string                                            `json:"documentation_url"`
	Source           AccountGatewayListListItemsResponseMessagesSource `json:"source"`
	JSON             accountGatewayListListItemsResponseMessageJSON    `json:"-"`
}

// accountGatewayListListItemsResponseMessageJSON contains the JSON metadata for
// the struct [AccountGatewayListListItemsResponseMessage]
type accountGatewayListListItemsResponseMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountGatewayListListItemsResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountGatewayListListItemsResponseMessageJSON) RawJSON() string {
	return r.raw
}

type AccountGatewayListListItemsResponseMessagesSource struct {
	Pointer string                                                `json:"pointer"`
	JSON    accountGatewayListListItemsResponseMessagesSourceJSON `json:"-"`
}

// accountGatewayListListItemsResponseMessagesSourceJSON contains the JSON metadata
// for the struct [AccountGatewayListListItemsResponseMessagesSource]
type accountGatewayListListItemsResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayListListItemsResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountGatewayListListItemsResponseMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AccountGatewayListListItemsResponseSuccess bool

const (
	AccountGatewayListListItemsResponseSuccessTrue AccountGatewayListListItemsResponseSuccess = true
)

func (r AccountGatewayListListItemsResponseSuccess) IsKnown() bool {
	switch r {
	case AccountGatewayListListItemsResponseSuccessTrue:
		return true
	}
	return false
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

type AccountGatewayListPatchResponse struct {
	Errors   []AccountGatewayListPatchResponseError   `json:"errors,required"`
	Messages []AccountGatewayListPatchResponseMessage `json:"messages,required"`
	// Whether the API call was successful
	Success AccountGatewayListPatchResponseSuccess `json:"success,required"`
	Result  ListGateway                            `json:"result"`
	JSON    accountGatewayListPatchResponseJSON    `json:"-"`
}

// accountGatewayListPatchResponseJSON contains the JSON metadata for the struct
// [AccountGatewayListPatchResponse]
type accountGatewayListPatchResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayListPatchResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountGatewayListPatchResponseJSON) RawJSON() string {
	return r.raw
}

type AccountGatewayListPatchResponseError struct {
	Code             int64                                       `json:"code,required"`
	Message          string                                      `json:"message,required"`
	DocumentationURL string                                      `json:"documentation_url"`
	Source           AccountGatewayListPatchResponseErrorsSource `json:"source"`
	JSON             accountGatewayListPatchResponseErrorJSON    `json:"-"`
}

// accountGatewayListPatchResponseErrorJSON contains the JSON metadata for the
// struct [AccountGatewayListPatchResponseError]
type accountGatewayListPatchResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountGatewayListPatchResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountGatewayListPatchResponseErrorJSON) RawJSON() string {
	return r.raw
}

type AccountGatewayListPatchResponseErrorsSource struct {
	Pointer string                                          `json:"pointer"`
	JSON    accountGatewayListPatchResponseErrorsSourceJSON `json:"-"`
}

// accountGatewayListPatchResponseErrorsSourceJSON contains the JSON metadata for
// the struct [AccountGatewayListPatchResponseErrorsSource]
type accountGatewayListPatchResponseErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayListPatchResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountGatewayListPatchResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AccountGatewayListPatchResponseMessage struct {
	Code             int64                                         `json:"code,required"`
	Message          string                                        `json:"message,required"`
	DocumentationURL string                                        `json:"documentation_url"`
	Source           AccountGatewayListPatchResponseMessagesSource `json:"source"`
	JSON             accountGatewayListPatchResponseMessageJSON    `json:"-"`
}

// accountGatewayListPatchResponseMessageJSON contains the JSON metadata for the
// struct [AccountGatewayListPatchResponseMessage]
type accountGatewayListPatchResponseMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountGatewayListPatchResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountGatewayListPatchResponseMessageJSON) RawJSON() string {
	return r.raw
}

type AccountGatewayListPatchResponseMessagesSource struct {
	Pointer string                                            `json:"pointer"`
	JSON    accountGatewayListPatchResponseMessagesSourceJSON `json:"-"`
}

// accountGatewayListPatchResponseMessagesSourceJSON contains the JSON metadata for
// the struct [AccountGatewayListPatchResponseMessagesSource]
type accountGatewayListPatchResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayListPatchResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountGatewayListPatchResponseMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AccountGatewayListPatchResponseSuccess bool

const (
	AccountGatewayListPatchResponseSuccessTrue AccountGatewayListPatchResponseSuccess = true
)

func (r AccountGatewayListPatchResponseSuccess) IsKnown() bool {
	switch r {
	case AccountGatewayListPatchResponseSuccessTrue:
		return true
	}
	return false
}

type AccountGatewayListNewParams struct {
	// The name of the list.
	Name param.Field[string] `json:"name,required"`
	// The type of list.
	Type param.Field[SchemasListType] `json:"type,required"`
	// The description of the list.
	Description param.Field[string] `json:"description"`
	// items to add to the list.
	Items param.Field[[]AccountGatewayListNewParamsItem] `json:"items"`
}

func (r AccountGatewayListNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountGatewayListNewParamsItem struct {
	// The description of the list item, if present
	Description param.Field[string] `json:"description"`
	// The value of the item in a list.
	Value param.Field[string] `json:"value"`
}

func (r AccountGatewayListNewParamsItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountGatewayListUpdateParams struct {
	// The name of the list.
	Name param.Field[string] `json:"name,required"`
	// The description of the list.
	Description param.Field[string] `json:"description"`
	// items to add to the list.
	Items param.Field[[]AccountGatewayListUpdateParamsItem] `json:"items"`
}

func (r AccountGatewayListUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountGatewayListUpdateParamsItem struct {
	// The description of the list item, if present
	Description param.Field[string] `json:"description"`
	// The value of the item in a list.
	Value param.Field[string] `json:"value"`
}

func (r AccountGatewayListUpdateParamsItem) MarshalJSON() (data []byte, err error) {
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

type AccountGatewayListPatchParams struct {
	// items to add to the list.
	Append param.Field[[]AccountGatewayListPatchParamsAppend] `json:"append"`
	// A list of the item values you want to remove.
	Remove param.Field[[]string] `json:"remove"`
}

func (r AccountGatewayListPatchParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountGatewayListPatchParamsAppend struct {
	// The description of the list item, if present
	Description param.Field[string] `json:"description"`
	// The value of the item in a list.
	Value param.Field[string] `json:"value"`
}

func (r AccountGatewayListPatchParamsAppend) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
