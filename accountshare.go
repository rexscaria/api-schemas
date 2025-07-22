// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/apiquery"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
	"github.com/rexscaria/api-schemas/shared"
	"github.com/tidwall/gjson"
)

// AccountShareService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountShareService] method instead.
type AccountShareService struct {
	Options    []option.RequestOption
	Recipients *AccountShareRecipientService
	Resources  *AccountShareResourceService
}

// NewAccountShareService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccountShareService(opts ...option.RequestOption) (r *AccountShareService) {
	r = &AccountShareService{}
	r.Options = opts
	r.Recipients = NewAccountShareRecipientService(opts...)
	r.Resources = NewAccountShareResourceService(opts...)
	return
}

// Create a new share
func (r *AccountShareService) New(ctx context.Context, accountID string, body AccountShareNewParams, opts ...option.RequestOption) (res *ShareResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/shares", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Fetches share by ID.
func (r *AccountShareService) Get(ctx context.Context, accountID string, shareID string, opts ...option.RequestOption) (res *ShareResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if shareID == "" {
		err = errors.New("missing required share_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/shares/%s", accountID, shareID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updating is not immediate, an updated share object with a new status will be
// returned.
func (r *AccountShareService) Update(ctx context.Context, accountID string, shareID string, body AccountShareUpdateParams, opts ...option.RequestOption) (res *ShareResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if shareID == "" {
		err = errors.New("missing required share_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/shares/%s", accountID, shareID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Lists all account shares.
func (r *AccountShareService) List(ctx context.Context, accountID string, query AccountShareListParams, opts ...option.RequestOption) (res *ShareResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/shares", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Deletion is not immediate, an updated share object with a new status will be
// returned.
func (r *AccountShareService) Delete(ctx context.Context, accountID string, shareID string, opts ...option.RequestOption) (res *ShareResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if shareID == "" {
		err = errors.New("missing required share_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/shares/%s", accountID, shareID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type APIResponseCollectionResourceSharing struct {
	Result     []interface{}                                  `json:"result,nullable"`
	ResultInfo APIResponseCollectionResourceSharingResultInfo `json:"result_info"`
	JSON       apiResponseCollectionResourceSharingJSON       `json:"-"`
	APIResponseResourceSharing
}

// apiResponseCollectionResourceSharingJSON contains the JSON metadata for the
// struct [APIResponseCollectionResourceSharing]
type apiResponseCollectionResourceSharingJSON struct {
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIResponseCollectionResourceSharing) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiResponseCollectionResourceSharingJSON) RawJSON() string {
	return r.raw
}

type APIResponseCollectionResourceSharingResultInfo struct {
	// Total number of results for the requested service.
	Count float64 `json:"count"`
	// Current page within paginated list of results.
	Page float64 `json:"page"`
	// Number of results per page of results.
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters.
	TotalCount float64 `json:"total_count"`
	// Total number of pages using the given per page.
	TotalPages float64                                            `json:"total_pages"`
	JSON       apiResponseCollectionResourceSharingResultInfoJSON `json:"-"`
}

// apiResponseCollectionResourceSharingResultInfoJSON contains the JSON metadata
// for the struct [APIResponseCollectionResourceSharingResultInfo]
type apiResponseCollectionResourceSharingResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	TotalPages  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIResponseCollectionResourceSharingResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiResponseCollectionResourceSharingResultInfoJSON) RawJSON() string {
	return r.raw
}

type APIResponseResourceSharing struct {
	Errors []APIResponseResourceSharingError `json:"errors,required"`
	// Whether the API call was successful.
	Success bool                                  `json:"success,required"`
	Result  APIResponseResourceSharingResultUnion `json:"result"`
	JSON    apiResponseResourceSharingJSON        `json:"-"`
}

// apiResponseResourceSharingJSON contains the JSON metadata for the struct
// [APIResponseResourceSharing]
type apiResponseResourceSharingJSON struct {
	Errors      apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIResponseResourceSharing) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiResponseResourceSharingJSON) RawJSON() string {
	return r.raw
}

type APIResponseResourceSharingError struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    apiResponseResourceSharingErrorJSON `json:"-"`
}

// apiResponseResourceSharingErrorJSON contains the JSON metadata for the struct
// [APIResponseResourceSharingError]
type apiResponseResourceSharingErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIResponseResourceSharingError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiResponseResourceSharingErrorJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [APIResponseResourceSharingResultArray] or
// [shared.UnionString].
type APIResponseResourceSharingResultUnion interface {
	ImplementsAPIResponseResourceSharingResultUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*APIResponseResourceSharingResultUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(APIResponseResourceSharingResultArray{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type APIResponseResourceSharingResultArray []interface{}

func (r APIResponseResourceSharingResultArray) ImplementsAPIResponseResourceSharingResultUnion() {}

type ShareObject struct {
	// Share identifier tag.
	ID string `json:"id,required"`
	// Account identifier.
	AccountID string `json:"account_id,required"`
	// The display name of an account.
	AccountName string `json:"account_name,required"`
	// When the share was created.
	Created time.Time `json:"created,required" format:"date-time"`
	// When the share was modified.
	Modified time.Time `json:"modified,required" format:"date-time"`
	// The name of the share.
	Name string `json:"name,required"`
	// Organization identifier.
	OrganizationID string                `json:"organization_id,required"`
	Status         ShareObjectStatus     `json:"status,required"`
	TargetType     ShareObjectTargetType `json:"target_type,required"`
	Kind           ShareObjectKind       `json:"kind"`
	JSON           shareObjectJSON       `json:"-"`
}

// shareObjectJSON contains the JSON metadata for the struct [ShareObject]
type shareObjectJSON struct {
	ID             apijson.Field
	AccountID      apijson.Field
	AccountName    apijson.Field
	Created        apijson.Field
	Modified       apijson.Field
	Name           apijson.Field
	OrganizationID apijson.Field
	Status         apijson.Field
	TargetType     apijson.Field
	Kind           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ShareObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r shareObjectJSON) RawJSON() string {
	return r.raw
}

type ShareObjectStatus string

const (
	ShareObjectStatusActive   ShareObjectStatus = "active"
	ShareObjectStatusDeleting ShareObjectStatus = "deleting"
	ShareObjectStatusDeleted  ShareObjectStatus = "deleted"
)

func (r ShareObjectStatus) IsKnown() bool {
	switch r {
	case ShareObjectStatusActive, ShareObjectStatusDeleting, ShareObjectStatusDeleted:
		return true
	}
	return false
}

type ShareObjectTargetType string

const (
	ShareObjectTargetTypeAccount      ShareObjectTargetType = "account"
	ShareObjectTargetTypeOrganization ShareObjectTargetType = "organization"
)

func (r ShareObjectTargetType) IsKnown() bool {
	switch r {
	case ShareObjectTargetTypeAccount, ShareObjectTargetTypeOrganization:
		return true
	}
	return false
}

type ShareObjectKind string

const (
	ShareObjectKindSent     ShareObjectKind = "sent"
	ShareObjectKindReceived ShareObjectKind = "received"
)

func (r ShareObjectKind) IsKnown() bool {
	switch r {
	case ShareObjectKindSent, ShareObjectKindReceived:
		return true
	}
	return false
}

type ShareResponseCollection struct {
	Result []ShareObject               `json:"result"`
	JSON   shareResponseCollectionJSON `json:"-"`
	APIResponseCollectionResourceSharing
}

// shareResponseCollectionJSON contains the JSON metadata for the struct
// [ShareResponseCollection]
type shareResponseCollectionJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ShareResponseCollection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r shareResponseCollectionJSON) RawJSON() string {
	return r.raw
}

type ShareResponseSingle struct {
	Result ShareObject             `json:"result"`
	JSON   shareResponseSingleJSON `json:"-"`
	APIResponseResourceSharing
}

// shareResponseSingleJSON contains the JSON metadata for the struct
// [ShareResponseSingle]
type shareResponseSingleJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ShareResponseSingle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r shareResponseSingleJSON) RawJSON() string {
	return r.raw
}

type AccountShareNewParams struct {
	// The name of the share.
	Name       param.Field[string]                             `json:"name,required"`
	Recipients param.Field[[]CreateShareRecipientRequestParam] `json:"recipients,required"`
	Resources  param.Field[[]CreateShareResourceRequestParam]  `json:"resources,required"`
}

func (r AccountShareNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountShareUpdateParams struct {
	// The name of the share.
	Name param.Field[string] `json:"name,required"`
}

func (r AccountShareUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountShareListParams struct {
	// Direction to sort objects.
	Direction param.Field[AccountShareListParamsDirection] `query:"direction"`
	// Filter shares by kind.
	Kind param.Field[AccountShareListParamsKind] `query:"kind"`
	// Order shares by values in the given field.
	Order param.Field[AccountShareListParamsOrder] `query:"order"`
	// Page number.
	Page param.Field[int64] `query:"page"`
	// Number of objects to return per page.
	PerPage param.Field[int64] `query:"per_page"`
	// Filter shares by status.
	Status param.Field[AccountShareListParamsStatus] `query:"status"`
	// Filter shares by target_type.
	TargetType param.Field[AccountShareListParamsTargetType] `query:"target_type"`
}

// URLQuery serializes [AccountShareListParams]'s query parameters as `url.Values`.
func (r AccountShareListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Direction to sort objects.
type AccountShareListParamsDirection string

const (
	AccountShareListParamsDirectionAsc  AccountShareListParamsDirection = "asc"
	AccountShareListParamsDirectionDesc AccountShareListParamsDirection = "desc"
)

func (r AccountShareListParamsDirection) IsKnown() bool {
	switch r {
	case AccountShareListParamsDirectionAsc, AccountShareListParamsDirectionDesc:
		return true
	}
	return false
}

// Filter shares by kind.
type AccountShareListParamsKind string

const (
	AccountShareListParamsKindSent     AccountShareListParamsKind = "sent"
	AccountShareListParamsKindReceived AccountShareListParamsKind = "received"
)

func (r AccountShareListParamsKind) IsKnown() bool {
	switch r {
	case AccountShareListParamsKindSent, AccountShareListParamsKindReceived:
		return true
	}
	return false
}

// Order shares by values in the given field.
type AccountShareListParamsOrder string

const (
	AccountShareListParamsOrderName    AccountShareListParamsOrder = "name"
	AccountShareListParamsOrderCreated AccountShareListParamsOrder = "created"
)

func (r AccountShareListParamsOrder) IsKnown() bool {
	switch r {
	case AccountShareListParamsOrderName, AccountShareListParamsOrderCreated:
		return true
	}
	return false
}

// Filter shares by status.
type AccountShareListParamsStatus string

const (
	AccountShareListParamsStatusActive   AccountShareListParamsStatus = "active"
	AccountShareListParamsStatusDeleting AccountShareListParamsStatus = "deleting"
	AccountShareListParamsStatusDeleted  AccountShareListParamsStatus = "deleted"
)

func (r AccountShareListParamsStatus) IsKnown() bool {
	switch r {
	case AccountShareListParamsStatusActive, AccountShareListParamsStatusDeleting, AccountShareListParamsStatusDeleted:
		return true
	}
	return false
}

// Filter shares by target_type.
type AccountShareListParamsTargetType string

const (
	AccountShareListParamsTargetTypeAccount      AccountShareListParamsTargetType = "account"
	AccountShareListParamsTargetTypeOrganization AccountShareListParamsTargetType = "organization"
)

func (r AccountShareListParamsTargetType) IsKnown() bool {
	switch r {
	case AccountShareListParamsTargetTypeAccount, AccountShareListParamsTargetTypeOrganization:
		return true
	}
	return false
}
