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

// AccountCustomPageService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountCustomPageService] method instead.
type AccountCustomPageService struct {
	Options []option.RequestOption
}

// NewAccountCustomPageService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountCustomPageService(opts ...option.RequestOption) (r *AccountCustomPageService) {
	r = &AccountCustomPageService{}
	r.Options = opts
	return
}

// Fetches all the custom pages at the account level.
func (r *AccountCustomPageService) ListCustomPages(ctx context.Context, accountIdentifier string, opts ...option.RequestOption) (res *CustomPagesResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	if accountIdentifier == "" {
		err = errors.New("missing required account_identifier parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/custom_pages", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Fetches the details of a custom page.
func (r *AccountCustomPageService) GetCustomPage(ctx context.Context, accountIdentifier string, identifier string, opts ...option.RequestOption) (res *CustomPagesResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	if accountIdentifier == "" {
		err = errors.New("missing required account_identifier parameter")
		return
	}
	if identifier == "" {
		err = errors.New("missing required identifier parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/custom_pages/%s", accountIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates the configuration of an existing custom page.
func (r *AccountCustomPageService) UpdateCustomPage(ctx context.Context, accountIdentifier string, identifier string, body AccountCustomPageUpdateCustomPageParams, opts ...option.RequestOption) (res *CustomPagesResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	if accountIdentifier == "" {
		err = errors.New("missing required account_identifier parameter")
		return
	}
	if identifier == "" {
		err = errors.New("missing required identifier parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/custom_pages/%s", accountIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type APIResponseCustomPages struct {
	Errors   []MessagesCustomPageItem          `json:"errors,required"`
	Messages []MessagesCustomPageItem          `json:"messages,required"`
	Result   APIResponseCustomPagesResultUnion `json:"result,required"`
	// Whether the API call was successful
	Success APIResponseCustomPagesSuccess `json:"success,required"`
	JSON    apiResponseCustomPagesJSON    `json:"-"`
}

// apiResponseCustomPagesJSON contains the JSON metadata for the struct
// [APIResponseCustomPages]
type apiResponseCustomPagesJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIResponseCustomPages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiResponseCustomPagesJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [APIResponseCustomPagesResultArray] or [shared.UnionString].
type APIResponseCustomPagesResultUnion interface {
	ImplementsAPIResponseCustomPagesResultUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*APIResponseCustomPagesResultUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(APIResponseCustomPagesResultArray{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type APIResponseCustomPagesResultArray []interface{}

func (r APIResponseCustomPagesResultArray) ImplementsAPIResponseCustomPagesResultUnion() {}

// Whether the API call was successful
type APIResponseCustomPagesSuccess bool

const (
	APIResponseCustomPagesSuccessTrue APIResponseCustomPagesSuccess = true
)

func (r APIResponseCustomPagesSuccess) IsKnown() bool {
	switch r {
	case APIResponseCustomPagesSuccessTrue:
		return true
	}
	return false
}

type CustomPagesResponseCollection struct {
	Result     []interface{}                           `json:"result,nullable"`
	ResultInfo CustomPagesResponseCollectionResultInfo `json:"result_info"`
	JSON       customPagesResponseCollectionJSON       `json:"-"`
	APIResponseCustomPages
}

// customPagesResponseCollectionJSON contains the JSON metadata for the struct
// [CustomPagesResponseCollection]
type customPagesResponseCollectionJSON struct {
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomPagesResponseCollection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customPagesResponseCollectionJSON) RawJSON() string {
	return r.raw
}

type CustomPagesResponseCollectionResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                     `json:"total_count"`
	JSON       customPagesResponseCollectionResultInfoJSON `json:"-"`
}

// customPagesResponseCollectionResultInfoJSON contains the JSON metadata for the
// struct [CustomPagesResponseCollectionResultInfo]
type customPagesResponseCollectionResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomPagesResponseCollectionResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customPagesResponseCollectionResultInfoJSON) RawJSON() string {
	return r.raw
}

type CustomPagesResponseSingle struct {
	Result interface{}                   `json:"result"`
	JSON   customPagesResponseSingleJSON `json:"-"`
	APIResponseCustomPages
}

// customPagesResponseSingleJSON contains the JSON metadata for the struct
// [CustomPagesResponseSingle]
type customPagesResponseSingleJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomPagesResponseSingle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customPagesResponseSingleJSON) RawJSON() string {
	return r.raw
}

type MessagesCustomPageItem struct {
	Code    int64                      `json:"code,required"`
	Message string                     `json:"message,required"`
	JSON    messagesCustomPageItemJSON `json:"-"`
}

// messagesCustomPageItemJSON contains the JSON metadata for the struct
// [MessagesCustomPageItem]
type messagesCustomPageItemJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessagesCustomPageItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messagesCustomPageItemJSON) RawJSON() string {
	return r.raw
}

// The custom page state.
type State string

const (
	StateDefault    State = "default"
	StateCustomized State = "customized"
)

func (r State) IsKnown() bool {
	switch r {
	case StateDefault, StateCustomized:
		return true
	}
	return false
}

type AccountCustomPageUpdateCustomPageParams struct {
	// The custom page state.
	State param.Field[State] `json:"state,required"`
	// The URL associated with the custom page.
	URL param.Field[string] `json:"url,required" format:"uri"`
}

func (r AccountCustomPageUpdateCustomPageParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
