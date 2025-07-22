// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
)

// AccountCloudforceOneRequestService contains methods and other services that help
// with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountCloudforceOneRequestService] method instead.
type AccountCloudforceOneRequestService struct {
	Options  []option.RequestOption
	Asset    *AccountCloudforceOneRequestAssetService
	Message  *AccountCloudforceOneRequestMessageService
	Priority *AccountCloudforceOneRequestPriorityService
}

// NewAccountCloudforceOneRequestService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountCloudforceOneRequestService(opts ...option.RequestOption) (r *AccountCloudforceOneRequestService) {
	r = &AccountCloudforceOneRequestService{}
	r.Options = opts
	r.Asset = NewAccountCloudforceOneRequestAssetService(opts...)
	r.Message = NewAccountCloudforceOneRequestMessageService(opts...)
	r.Priority = NewAccountCloudforceOneRequestPriorityService(opts...)
	return
}

// Creating a request adds the request into the Cloudforce One queue for analysis.
// In addition to the content, a short title, type, priority, and releasability
// should be provided. If one is not provided, a default will be assigned.
func (r *AccountCloudforceOneRequestService) New(ctx context.Context, accountIdentifier string, body AccountCloudforceOneRequestNewParams, opts ...option.RequestOption) (res *AccountCloudforceOneRequestNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountIdentifier == "" {
		err = errors.New("missing required account_identifier parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/requests/new", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get a Request
func (r *AccountCloudforceOneRequestService) Get(ctx context.Context, accountIdentifier string, requestIdentifier string, opts ...option.RequestOption) (res *AccountCloudforceOneRequestGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountIdentifier == "" {
		err = errors.New("missing required account_identifier parameter")
		return
	}
	if requestIdentifier == "" {
		err = errors.New("missing required request_identifier parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/requests/%s", accountIdentifier, requestIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updating a request alters the request in the Cloudforce One queue. This API may
// be used to update any attributes of the request after the initial submission.
// Only fields that you choose to update need to be add to the request body.
func (r *AccountCloudforceOneRequestService) Update(ctx context.Context, accountIdentifier string, requestIdentifier string, body AccountCloudforceOneRequestUpdateParams, opts ...option.RequestOption) (res *AccountCloudforceOneRequestUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountIdentifier == "" {
		err = errors.New("missing required account_identifier parameter")
		return
	}
	if requestIdentifier == "" {
		err = errors.New("missing required request_identifier parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/requests/%s", accountIdentifier, requestIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// List Requests
func (r *AccountCloudforceOneRequestService) List(ctx context.Context, accountIdentifier string, body AccountCloudforceOneRequestListParams, opts ...option.RequestOption) (res *AccountCloudforceOneRequestListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountIdentifier == "" {
		err = errors.New("missing required account_identifier parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/requests", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Delete a Request
func (r *AccountCloudforceOneRequestService) Delete(ctx context.Context, accountIdentifier string, requestIdentifier string, opts ...option.RequestOption) (res *APIResponseRequests, err error) {
	opts = append(r.Options[:], opts...)
	if accountIdentifier == "" {
		err = errors.New("missing required account_identifier parameter")
		return
	}
	if requestIdentifier == "" {
		err = errors.New("missing required request_identifier parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/requests/%s", accountIdentifier, requestIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Get Request Priority, Status, and TLP constants
func (r *AccountCloudforceOneRequestService) GetConstants(ctx context.Context, accountIdentifier string, opts ...option.RequestOption) (res *AccountCloudforceOneRequestGetConstantsResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountIdentifier == "" {
		err = errors.New("missing required account_identifier parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/requests/constants", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get Priority Intelligence Requirement Quota
func (r *AccountCloudforceOneRequestService) GetPriorityQuota(ctx context.Context, accountIdentifier string, opts ...option.RequestOption) (res *AccountCloudforceOneRequestGetPriorityQuotaResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountIdentifier == "" {
		err = errors.New("missing required account_identifier parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/requests/priority/quota", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get Request Quota
func (r *AccountCloudforceOneRequestService) GetQuota(ctx context.Context, accountIdentifier string, opts ...option.RequestOption) (res *AccountCloudforceOneRequestGetQuotaResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountIdentifier == "" {
		err = errors.New("missing required account_identifier parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/requests/quota", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get Request Types
func (r *AccountCloudforceOneRequestService) GetTypes(ctx context.Context, accountIdentifier string, opts ...option.RequestOption) (res *AccountCloudforceOneRequestGetTypesResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountIdentifier == "" {
		err = errors.New("missing required account_identifier parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/requests/types", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type APIResponseRequests struct {
	Errors   []MessagesRequestsItems `json:"errors,required"`
	Messages []MessagesRequestsItems `json:"messages,required"`
	// Whether the API call was successful
	Success APIResponseRequestsSuccess `json:"success,required"`
	JSON    apiResponseRequestsJSON    `json:"-"`
}

// apiResponseRequestsJSON contains the JSON metadata for the struct
// [APIResponseRequests]
type apiResponseRequestsJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIResponseRequests) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiResponseRequestsJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type APIResponseRequestsSuccess bool

const (
	APIResponseRequestsSuccessTrue APIResponseRequestsSuccess = true
)

func (r APIResponseRequestsSuccess) IsKnown() bool {
	switch r {
	case APIResponseRequestsSuccessTrue:
		return true
	}
	return false
}

type MessagesRequestsItems struct {
	Code    int64                     `json:"code,required"`
	Message string                    `json:"message,required"`
	JSON    messagesRequestsItemsJSON `json:"-"`
}

// messagesRequestsItemsJSON contains the JSON metadata for the struct
// [MessagesRequestsItems]
type messagesRequestsItemsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessagesRequestsItems) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messagesRequestsItemsJSON) RawJSON() string {
	return r.raw
}

type Priority string

const (
	PriorityRoutine Priority = "routine"
	PriorityHigh    Priority = "high"
	PriorityUrgent  Priority = "urgent"
)

func (r Priority) IsKnown() bool {
	switch r {
	case PriorityRoutine, PriorityHigh, PriorityUrgent:
		return true
	}
	return false
}

type RequestEditParam struct {
	// Request content
	Content param.Field[string] `json:"content"`
	// Priority for analyzing the request
	Priority param.Field[string] `json:"priority"`
	// Requested information from request
	RequestType param.Field[string] `json:"request_type"`
	// Brief description of the request
	Summary param.Field[string] `json:"summary"`
	// The CISA defined Traffic Light Protocol (TLP)
	Tlp param.Field[Tlp] `json:"tlp"`
}

func (r RequestEditParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RequestItem struct {
	// UUID
	ID string `json:"id,required"`
	// Request content
	Content  string    `json:"content,required"`
	Created  time.Time `json:"created,required" format:"date-time"`
	Priority time.Time `json:"priority,required" format:"date-time"`
	// Requested information from request
	Request string `json:"request,required"`
	// Brief description of the request
	Summary string `json:"summary,required"`
	// The CISA defined Traffic Light Protocol (TLP)
	Tlp       Tlp       `json:"tlp,required"`
	Updated   time.Time `json:"updated,required" format:"date-time"`
	Completed time.Time `json:"completed" format:"date-time"`
	// Tokens for the request messages
	MessageTokens int64 `json:"message_tokens"`
	// Readable Request ID
	ReadableID string `json:"readable_id"`
	// Request Status
	Status RequestStatus `json:"status"`
	// Tokens for the request
	Tokens int64           `json:"tokens"`
	JSON   requestItemJSON `json:"-"`
}

// requestItemJSON contains the JSON metadata for the struct [RequestItem]
type requestItemJSON struct {
	ID            apijson.Field
	Content       apijson.Field
	Created       apijson.Field
	Priority      apijson.Field
	Request       apijson.Field
	Summary       apijson.Field
	Tlp           apijson.Field
	Updated       apijson.Field
	Completed     apijson.Field
	MessageTokens apijson.Field
	ReadableID    apijson.Field
	Status        apijson.Field
	Tokens        apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *RequestItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r requestItemJSON) RawJSON() string {
	return r.raw
}

// Request Status
type RequestStatus string

const (
	RequestStatusOpen      RequestStatus = "open"
	RequestStatusAccepted  RequestStatus = "accepted"
	RequestStatusReported  RequestStatus = "reported"
	RequestStatusApproved  RequestStatus = "approved"
	RequestStatusCompleted RequestStatus = "completed"
	RequestStatusDeclined  RequestStatus = "declined"
)

func (r RequestStatus) IsKnown() bool {
	switch r {
	case RequestStatusOpen, RequestStatusAccepted, RequestStatusReported, RequestStatusApproved, RequestStatusCompleted, RequestStatusDeclined:
		return true
	}
	return false
}

// The CISA defined Traffic Light Protocol (TLP)
type Tlp string

const (
	TlpClear       Tlp = "clear"
	TlpAmber       Tlp = "amber"
	TlpAmberStrict Tlp = "amber-strict"
	TlpGreen       Tlp = "green"
	TlpRed         Tlp = "red"
)

func (r Tlp) IsKnown() bool {
	switch r {
	case TlpClear, TlpAmber, TlpAmberStrict, TlpGreen, TlpRed:
		return true
	}
	return false
}

type AccountCloudforceOneRequestNewResponse struct {
	Result RequestItem                                `json:"result"`
	JSON   accountCloudforceOneRequestNewResponseJSON `json:"-"`
	APIResponseRequests
}

// accountCloudforceOneRequestNewResponseJSON contains the JSON metadata for the
// struct [AccountCloudforceOneRequestNewResponse]
type accountCloudforceOneRequestNewResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCloudforceOneRequestNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountCloudforceOneRequestNewResponseJSON) RawJSON() string {
	return r.raw
}

type AccountCloudforceOneRequestGetResponse struct {
	Result RequestItem                                `json:"result"`
	JSON   accountCloudforceOneRequestGetResponseJSON `json:"-"`
	APIResponseRequests
}

// accountCloudforceOneRequestGetResponseJSON contains the JSON metadata for the
// struct [AccountCloudforceOneRequestGetResponse]
type accountCloudforceOneRequestGetResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCloudforceOneRequestGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountCloudforceOneRequestGetResponseJSON) RawJSON() string {
	return r.raw
}

type AccountCloudforceOneRequestUpdateResponse struct {
	Result RequestItem                                   `json:"result"`
	JSON   accountCloudforceOneRequestUpdateResponseJSON `json:"-"`
	APIResponseRequests
}

// accountCloudforceOneRequestUpdateResponseJSON contains the JSON metadata for the
// struct [AccountCloudforceOneRequestUpdateResponse]
type accountCloudforceOneRequestUpdateResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCloudforceOneRequestUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountCloudforceOneRequestUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type AccountCloudforceOneRequestListResponse struct {
	Result []AccountCloudforceOneRequestListResponseResult `json:"result"`
	JSON   accountCloudforceOneRequestListResponseJSON     `json:"-"`
	APIResponseRequests
}

// accountCloudforceOneRequestListResponseJSON contains the JSON metadata for the
// struct [AccountCloudforceOneRequestListResponse]
type accountCloudforceOneRequestListResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCloudforceOneRequestListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountCloudforceOneRequestListResponseJSON) RawJSON() string {
	return r.raw
}

type AccountCloudforceOneRequestListResponseResult struct {
	// UUID
	ID string `json:"id,required"`
	// Request creation time
	Created  AccountCloudforceOneRequestListResponseResultCreated `json:"created,required"`
	Priority Priority                                             `json:"priority,required"`
	// Requested information from request
	Request string `json:"request,required"`
	// Brief description of the request
	Summary string `json:"summary,required"`
	// The CISA defined Traffic Light Protocol (TLP)
	Tlp Tlp `json:"tlp,required"`
	// Request last updated time
	Updated AccountCloudforceOneRequestListResponseResultUpdated `json:"updated,required"`
	// Request completion time
	Completed AccountCloudforceOneRequestListResponseResultCompleted `json:"completed"`
	// Tokens for the request messages
	MessageTokens int64 `json:"message_tokens"`
	// Readable Request ID
	ReadableID string `json:"readable_id"`
	// Request Status
	Status RequestStatus `json:"status"`
	// Tokens for the request
	Tokens int64                                             `json:"tokens"`
	JSON   accountCloudforceOneRequestListResponseResultJSON `json:"-"`
}

// accountCloudforceOneRequestListResponseResultJSON contains the JSON metadata for
// the struct [AccountCloudforceOneRequestListResponseResult]
type accountCloudforceOneRequestListResponseResultJSON struct {
	ID            apijson.Field
	Created       apijson.Field
	Priority      apijson.Field
	Request       apijson.Field
	Summary       apijson.Field
	Tlp           apijson.Field
	Updated       apijson.Field
	Completed     apijson.Field
	MessageTokens apijson.Field
	ReadableID    apijson.Field
	Status        apijson.Field
	Tokens        apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccountCloudforceOneRequestListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountCloudforceOneRequestListResponseResultJSON) RawJSON() string {
	return r.raw
}

// Request creation time
type AccountCloudforceOneRequestListResponseResultCreated struct {
	JSON accountCloudforceOneRequestListResponseResultCreatedJSON `json:"-"`
}

// accountCloudforceOneRequestListResponseResultCreatedJSON contains the JSON
// metadata for the struct [AccountCloudforceOneRequestListResponseResultCreated]
type accountCloudforceOneRequestListResponseResultCreatedJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCloudforceOneRequestListResponseResultCreated) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountCloudforceOneRequestListResponseResultCreatedJSON) RawJSON() string {
	return r.raw
}

// Request last updated time
type AccountCloudforceOneRequestListResponseResultUpdated struct {
	JSON accountCloudforceOneRequestListResponseResultUpdatedJSON `json:"-"`
}

// accountCloudforceOneRequestListResponseResultUpdatedJSON contains the JSON
// metadata for the struct [AccountCloudforceOneRequestListResponseResultUpdated]
type accountCloudforceOneRequestListResponseResultUpdatedJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCloudforceOneRequestListResponseResultUpdated) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountCloudforceOneRequestListResponseResultUpdatedJSON) RawJSON() string {
	return r.raw
}

// Request completion time
type AccountCloudforceOneRequestListResponseResultCompleted struct {
	JSON accountCloudforceOneRequestListResponseResultCompletedJSON `json:"-"`
}

// accountCloudforceOneRequestListResponseResultCompletedJSON contains the JSON
// metadata for the struct [AccountCloudforceOneRequestListResponseResultCompleted]
type accountCloudforceOneRequestListResponseResultCompletedJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCloudforceOneRequestListResponseResultCompleted) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountCloudforceOneRequestListResponseResultCompletedJSON) RawJSON() string {
	return r.raw
}

type AccountCloudforceOneRequestGetConstantsResponse struct {
	Result AccountCloudforceOneRequestGetConstantsResponseResult `json:"result"`
	JSON   accountCloudforceOneRequestGetConstantsResponseJSON   `json:"-"`
	APIResponseRequests
}

// accountCloudforceOneRequestGetConstantsResponseJSON contains the JSON metadata
// for the struct [AccountCloudforceOneRequestGetConstantsResponse]
type accountCloudforceOneRequestGetConstantsResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCloudforceOneRequestGetConstantsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountCloudforceOneRequestGetConstantsResponseJSON) RawJSON() string {
	return r.raw
}

type AccountCloudforceOneRequestGetConstantsResponseResult struct {
	Priority []Priority                                                `json:"priority"`
	Status   []RequestStatus                                           `json:"status"`
	Tlp      []Tlp                                                     `json:"tlp"`
	JSON     accountCloudforceOneRequestGetConstantsResponseResultJSON `json:"-"`
}

// accountCloudforceOneRequestGetConstantsResponseResultJSON contains the JSON
// metadata for the struct [AccountCloudforceOneRequestGetConstantsResponseResult]
type accountCloudforceOneRequestGetConstantsResponseResultJSON struct {
	Priority    apijson.Field
	Status      apijson.Field
	Tlp         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCloudforceOneRequestGetConstantsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountCloudforceOneRequestGetConstantsResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountCloudforceOneRequestGetPriorityQuotaResponse struct {
	Result Quota                                                   `json:"result"`
	JSON   accountCloudforceOneRequestGetPriorityQuotaResponseJSON `json:"-"`
	APIResponseRequests
}

// accountCloudforceOneRequestGetPriorityQuotaResponseJSON contains the JSON
// metadata for the struct [AccountCloudforceOneRequestGetPriorityQuotaResponse]
type accountCloudforceOneRequestGetPriorityQuotaResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCloudforceOneRequestGetPriorityQuotaResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountCloudforceOneRequestGetPriorityQuotaResponseJSON) RawJSON() string {
	return r.raw
}

type AccountCloudforceOneRequestGetQuotaResponse struct {
	Result Quota                                           `json:"result"`
	JSON   accountCloudforceOneRequestGetQuotaResponseJSON `json:"-"`
	APIResponseRequests
}

// accountCloudforceOneRequestGetQuotaResponseJSON contains the JSON metadata for
// the struct [AccountCloudforceOneRequestGetQuotaResponse]
type accountCloudforceOneRequestGetQuotaResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCloudforceOneRequestGetQuotaResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountCloudforceOneRequestGetQuotaResponseJSON) RawJSON() string {
	return r.raw
}

type AccountCloudforceOneRequestGetTypesResponse struct {
	Result []string                                        `json:"result"`
	JSON   accountCloudforceOneRequestGetTypesResponseJSON `json:"-"`
	APIResponseRequests
}

// accountCloudforceOneRequestGetTypesResponseJSON contains the JSON metadata for
// the struct [AccountCloudforceOneRequestGetTypesResponse]
type accountCloudforceOneRequestGetTypesResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCloudforceOneRequestGetTypesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountCloudforceOneRequestGetTypesResponseJSON) RawJSON() string {
	return r.raw
}

type AccountCloudforceOneRequestNewParams struct {
	RequestEdit RequestEditParam `json:"request_edit,required"`
}

func (r AccountCloudforceOneRequestNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.RequestEdit)
}

type AccountCloudforceOneRequestUpdateParams struct {
	RequestEdit RequestEditParam `json:"request_edit,required"`
}

func (r AccountCloudforceOneRequestUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.RequestEdit)
}

type AccountCloudforceOneRequestListParams struct {
	// Page number of results
	Page param.Field[int64] `json:"page,required"`
	// Number of results per page
	PerPage param.Field[int64] `json:"per_page,required"`
	// Retrieve requests completed after this time
	CompletedAfter param.Field[AccountCloudforceOneRequestListParamsCompletedAfter] `json:"completed_after"`
	// Retrieve requests completed before this time
	CompletedBefore param.Field[AccountCloudforceOneRequestListParamsCompletedBefore] `json:"completed_before"`
	// Retrieve requests created after this time
	CreatedAfter param.Field[AccountCloudforceOneRequestListParamsCreatedAfter] `json:"created_after"`
	// Retrieve requests created before this time
	CreatedBefore param.Field[AccountCloudforceOneRequestListParamsCreatedBefore] `json:"created_before"`
	// Requested information from request
	RequestType param.Field[string] `json:"request_type"`
	// Field to sort results by
	SortBy param.Field[string] `json:"sort_by"`
	// Sort order (asc or desc)
	SortOrder param.Field[AccountCloudforceOneRequestListParamsSortOrder] `json:"sort_order"`
	// Request Status
	Status param.Field[RequestStatus] `json:"status"`
}

func (r AccountCloudforceOneRequestListParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Retrieve requests completed after this time
type AccountCloudforceOneRequestListParamsCompletedAfter struct {
}

func (r AccountCloudforceOneRequestListParamsCompletedAfter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Retrieve requests completed before this time
type AccountCloudforceOneRequestListParamsCompletedBefore struct {
}

func (r AccountCloudforceOneRequestListParamsCompletedBefore) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Retrieve requests created after this time
type AccountCloudforceOneRequestListParamsCreatedAfter struct {
}

func (r AccountCloudforceOneRequestListParamsCreatedAfter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Retrieve requests created before this time
type AccountCloudforceOneRequestListParamsCreatedBefore struct {
}

func (r AccountCloudforceOneRequestListParamsCreatedBefore) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Sort order (asc or desc)
type AccountCloudforceOneRequestListParamsSortOrder string

const (
	AccountCloudforceOneRequestListParamsSortOrderAsc  AccountCloudforceOneRequestListParamsSortOrder = "asc"
	AccountCloudforceOneRequestListParamsSortOrderDesc AccountCloudforceOneRequestListParamsSortOrder = "desc"
)

func (r AccountCloudforceOneRequestListParamsSortOrder) IsKnown() bool {
	switch r {
	case AccountCloudforceOneRequestListParamsSortOrderAsc, AccountCloudforceOneRequestListParamsSortOrderDesc:
		return true
	}
	return false
}
