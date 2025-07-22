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

// AccountCloudforceOneRequestMessageService contains methods and other services
// that help with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountCloudforceOneRequestMessageService] method instead.
type AccountCloudforceOneRequestMessageService struct {
	Options []option.RequestOption
}

// NewAccountCloudforceOneRequestMessageService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewAccountCloudforceOneRequestMessageService(opts ...option.RequestOption) (r *AccountCloudforceOneRequestMessageService) {
	r = &AccountCloudforceOneRequestMessageService{}
	r.Options = opts
	return
}

// Create a New Request Message
func (r *AccountCloudforceOneRequestMessageService) New(ctx context.Context, accountIdentifier string, requestIdentifier string, body AccountCloudforceOneRequestMessageNewParams, opts ...option.RequestOption) (res *AccountCloudforceOneRequestMessageNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountIdentifier == "" {
		err = errors.New("missing required account_identifier parameter")
		return
	}
	if requestIdentifier == "" {
		err = errors.New("missing required request_identifier parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/requests/%s/message/new", accountIdentifier, requestIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Update a Request Message
func (r *AccountCloudforceOneRequestMessageService) Update(ctx context.Context, accountIdentifier string, requestIdentifier string, messageIdentifer int64, body AccountCloudforceOneRequestMessageUpdateParams, opts ...option.RequestOption) (res *AccountCloudforceOneRequestMessageUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountIdentifier == "" {
		err = errors.New("missing required account_identifier parameter")
		return
	}
	if requestIdentifier == "" {
		err = errors.New("missing required request_identifier parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/requests/%s/message/%v", accountIdentifier, requestIdentifier, messageIdentifer)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// List Request Messages
func (r *AccountCloudforceOneRequestMessageService) List(ctx context.Context, accountIdentifier string, requestIdentifier string, body AccountCloudforceOneRequestMessageListParams, opts ...option.RequestOption) (res *AccountCloudforceOneRequestMessageListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountIdentifier == "" {
		err = errors.New("missing required account_identifier parameter")
		return
	}
	if requestIdentifier == "" {
		err = errors.New("missing required request_identifier parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/requests/%s/message", accountIdentifier, requestIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Delete a Request Message
func (r *AccountCloudforceOneRequestMessageService) Delete(ctx context.Context, accountIdentifier string, requestIdentifier string, messageIdentifer int64, opts ...option.RequestOption) (res *APIResponseRequests, err error) {
	opts = append(r.Options[:], opts...)
	if accountIdentifier == "" {
		err = errors.New("missing required account_identifier parameter")
		return
	}
	if requestIdentifier == "" {
		err = errors.New("missing required request_identifier parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/requests/%s/message/%v", accountIdentifier, requestIdentifier, messageIdentifer)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type MessageEditParam struct {
	// Content of message
	Content param.Field[string] `json:"content"`
}

func (r MessageEditParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MessageRequestItem struct {
	// Message ID
	ID int64 `json:"id,required"`
	// Author of message
	Author string `json:"author,required"`
	// Content of message
	Content string `json:"content,required"`
	// Whether the message is a follow-on request
	IsFollowOnRequest bool `json:"is_follow_on_request,required"`
	// Message last updated time
	Updated MessageRequestItemUpdated `json:"updated,required"`
	// Message creation time
	Created MessageRequestItemCreated `json:"created"`
	JSON    messageRequestItemJSON    `json:"-"`
}

// messageRequestItemJSON contains the JSON metadata for the struct
// [MessageRequestItem]
type messageRequestItemJSON struct {
	ID                apijson.Field
	Author            apijson.Field
	Content           apijson.Field
	IsFollowOnRequest apijson.Field
	Updated           apijson.Field
	Created           apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *MessageRequestItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageRequestItemJSON) RawJSON() string {
	return r.raw
}

// Message last updated time
type MessageRequestItemUpdated struct {
	JSON messageRequestItemUpdatedJSON `json:"-"`
}

// messageRequestItemUpdatedJSON contains the JSON metadata for the struct
// [MessageRequestItemUpdated]
type messageRequestItemUpdatedJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageRequestItemUpdated) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageRequestItemUpdatedJSON) RawJSON() string {
	return r.raw
}

// Message creation time
type MessageRequestItemCreated struct {
	JSON messageRequestItemCreatedJSON `json:"-"`
}

// messageRequestItemCreatedJSON contains the JSON metadata for the struct
// [MessageRequestItemCreated]
type messageRequestItemCreatedJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageRequestItemCreated) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageRequestItemCreatedJSON) RawJSON() string {
	return r.raw
}

type AccountCloudforceOneRequestMessageNewResponse struct {
	Result MessageRequestItem                                `json:"result"`
	JSON   accountCloudforceOneRequestMessageNewResponseJSON `json:"-"`
	APIResponseRequests
}

// accountCloudforceOneRequestMessageNewResponseJSON contains the JSON metadata for
// the struct [AccountCloudforceOneRequestMessageNewResponse]
type accountCloudforceOneRequestMessageNewResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCloudforceOneRequestMessageNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountCloudforceOneRequestMessageNewResponseJSON) RawJSON() string {
	return r.raw
}

type AccountCloudforceOneRequestMessageUpdateResponse struct {
	Result MessageRequestItem                                   `json:"result"`
	JSON   accountCloudforceOneRequestMessageUpdateResponseJSON `json:"-"`
	APIResponseRequests
}

// accountCloudforceOneRequestMessageUpdateResponseJSON contains the JSON metadata
// for the struct [AccountCloudforceOneRequestMessageUpdateResponse]
type accountCloudforceOneRequestMessageUpdateResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCloudforceOneRequestMessageUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountCloudforceOneRequestMessageUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type AccountCloudforceOneRequestMessageListResponse struct {
	Result []MessageRequestItem                               `json:"result"`
	JSON   accountCloudforceOneRequestMessageListResponseJSON `json:"-"`
	APIResponseRequests
}

// accountCloudforceOneRequestMessageListResponseJSON contains the JSON metadata
// for the struct [AccountCloudforceOneRequestMessageListResponse]
type accountCloudforceOneRequestMessageListResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCloudforceOneRequestMessageListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountCloudforceOneRequestMessageListResponseJSON) RawJSON() string {
	return r.raw
}

type AccountCloudforceOneRequestMessageNewParams struct {
	MessageEdit MessageEditParam `json:"message_edit,required"`
}

func (r AccountCloudforceOneRequestMessageNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.MessageEdit)
}

type AccountCloudforceOneRequestMessageUpdateParams struct {
	MessageEdit MessageEditParam `json:"message_edit,required"`
}

func (r AccountCloudforceOneRequestMessageUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.MessageEdit)
}

type AccountCloudforceOneRequestMessageListParams struct {
	// Page number of results
	Page param.Field[int64] `json:"page,required"`
	// Number of results per page
	PerPage param.Field[int64] `json:"per_page,required"`
	// Retrieve messages created after this time
	After param.Field[AccountCloudforceOneRequestMessageListParamsAfter] `json:"after"`
	// Retrieve messages created before this time
	Before param.Field[AccountCloudforceOneRequestMessageListParamsBefore] `json:"before"`
	// Field to sort results by
	SortBy param.Field[string] `json:"sort_by"`
	// Sort order (asc or desc)
	SortOrder param.Field[AccountCloudforceOneRequestMessageListParamsSortOrder] `json:"sort_order"`
}

func (r AccountCloudforceOneRequestMessageListParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Retrieve messages created after this time
type AccountCloudforceOneRequestMessageListParamsAfter struct {
}

func (r AccountCloudforceOneRequestMessageListParamsAfter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Retrieve messages created before this time
type AccountCloudforceOneRequestMessageListParamsBefore struct {
}

func (r AccountCloudforceOneRequestMessageListParamsBefore) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Sort order (asc or desc)
type AccountCloudforceOneRequestMessageListParamsSortOrder string

const (
	AccountCloudforceOneRequestMessageListParamsSortOrderAsc  AccountCloudforceOneRequestMessageListParamsSortOrder = "asc"
	AccountCloudforceOneRequestMessageListParamsSortOrderDesc AccountCloudforceOneRequestMessageListParamsSortOrder = "desc"
)

func (r AccountCloudforceOneRequestMessageListParamsSortOrder) IsKnown() bool {
	switch r {
	case AccountCloudforceOneRequestMessageListParamsSortOrderAsc, AccountCloudforceOneRequestMessageListParamsSortOrderDesc:
		return true
	}
	return false
}
