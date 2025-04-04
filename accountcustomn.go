// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/cf-rex-go/internal/apijson"
	"github.com/stainless-sdks/cf-rex-go/internal/param"
	"github.com/stainless-sdks/cf-rex-go/internal/requestconfig"
	"github.com/stainless-sdks/cf-rex-go/option"
)

// AccountCustomNService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountCustomNService] method instead.
type AccountCustomNService struct {
	Options []option.RequestOption
}

// NewAccountCustomNService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccountCustomNService(opts ...option.RequestOption) (r *AccountCustomNService) {
	r = &AccountCustomNService{}
	r.Options = opts
	return
}

// Add Account Custom Nameserver
func (r *AccountCustomNService) New(ctx context.Context, accountID string, body AccountCustomNNewParams, opts ...option.RequestOption) (res *AccountCustomNNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/custom_ns", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List an account's custom nameservers.
func (r *AccountCustomNService) List(ctx context.Context, accountID string, opts ...option.RequestOption) (res *AccountCustomNListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/custom_ns", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete Account Custom Nameserver
func (r *AccountCustomNService) Delete(ctx context.Context, accountID string, customNsID string, body AccountCustomNDeleteParams, opts ...option.RequestOption) (res *AccountCustomNDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if customNsID == "" {
		err = errors.New("missing required custom_ns_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/custom_ns/%s", accountID, customNsID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

type CommonResponseCustomNs struct {
	Errors   []CustomNsMessages `json:"errors,required"`
	Messages []CustomNsMessages `json:"messages,required"`
	// Whether the API call was successful
	Success CommonResponseCustomNsSuccess `json:"success,required"`
	JSON    commonResponseCustomNsJSON    `json:"-"`
}

// commonResponseCustomNsJSON contains the JSON metadata for the struct
// [CommonResponseCustomNs]
type commonResponseCustomNsJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CommonResponseCustomNs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r commonResponseCustomNsJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type CommonResponseCustomNsSuccess bool

const (
	CommonResponseCustomNsSuccessTrue CommonResponseCustomNsSuccess = true
)

func (r CommonResponseCustomNsSuccess) IsKnown() bool {
	switch r {
	case CommonResponseCustomNsSuccessTrue:
		return true
	}
	return false
}

// A single account custom nameserver.
type CustomNs struct {
	// A and AAAA records associated with the nameserver.
	DNSRecords []CustomNsDNSRecord `json:"dns_records,required"`
	// The FQDN of the name server.
	NsName string `json:"ns_name,required" format:"hostname"`
	// Verification status of the nameserver.
	//
	// Deprecated: deprecated
	Status CustomNsStatus `json:"status,required"`
	// Identifier
	ZoneTag string `json:"zone_tag,required"`
	// The number of the set that this name server belongs to.
	NsSet float64      `json:"ns_set"`
	JSON  customNsJSON `json:"-"`
}

// customNsJSON contains the JSON metadata for the struct [CustomNs]
type customNsJSON struct {
	DNSRecords  apijson.Field
	NsName      apijson.Field
	Status      apijson.Field
	ZoneTag     apijson.Field
	NsSet       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomNs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customNsJSON) RawJSON() string {
	return r.raw
}

type CustomNsDNSRecord struct {
	// DNS record type.
	Type CustomNsDNSRecordsType `json:"type"`
	// DNS record contents (an IPv4 or IPv6 address).
	Value string                `json:"value"`
	JSON  customNsDNSRecordJSON `json:"-"`
}

// customNsDNSRecordJSON contains the JSON metadata for the struct
// [CustomNsDNSRecord]
type customNsDNSRecordJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomNsDNSRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customNsDNSRecordJSON) RawJSON() string {
	return r.raw
}

// DNS record type.
type CustomNsDNSRecordsType string

const (
	CustomNsDNSRecordsTypeA    CustomNsDNSRecordsType = "A"
	CustomNsDNSRecordsTypeAaaa CustomNsDNSRecordsType = "AAAA"
)

func (r CustomNsDNSRecordsType) IsKnown() bool {
	switch r {
	case CustomNsDNSRecordsTypeA, CustomNsDNSRecordsTypeAaaa:
		return true
	}
	return false
}

// Verification status of the nameserver.
type CustomNsStatus string

const (
	CustomNsStatusMoved    CustomNsStatus = "moved"
	CustomNsStatusPending  CustomNsStatus = "pending"
	CustomNsStatusVerified CustomNsStatus = "verified"
)

func (r CustomNsStatus) IsKnown() bool {
	switch r {
	case CustomNsStatusMoved, CustomNsStatusPending, CustomNsStatusVerified:
		return true
	}
	return false
}

type CustomNsCollection struct {
	ResultInfo CustomNsCollectionResultInfo `json:"result_info"`
	JSON       customNsCollectionJSON       `json:"-"`
	CommonResponseCustomNs
}

// customNsCollectionJSON contains the JSON metadata for the struct
// [CustomNsCollection]
type customNsCollectionJSON struct {
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomNsCollection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customNsCollectionJSON) RawJSON() string {
	return r.raw
}

type CustomNsCollectionResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                          `json:"total_count"`
	JSON       customNsCollectionResultInfoJSON `json:"-"`
}

// customNsCollectionResultInfoJSON contains the JSON metadata for the struct
// [CustomNsCollectionResultInfo]
type customNsCollectionResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomNsCollectionResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customNsCollectionResultInfoJSON) RawJSON() string {
	return r.raw
}

type CustomNsMessages struct {
	Code    int64                `json:"code,required"`
	Message string               `json:"message,required"`
	JSON    customNsMessagesJSON `json:"-"`
}

// customNsMessagesJSON contains the JSON metadata for the struct
// [CustomNsMessages]
type customNsMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomNsMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customNsMessagesJSON) RawJSON() string {
	return r.raw
}

type AccountCustomNNewResponse struct {
	// A single account custom nameserver.
	Result CustomNs                      `json:"result"`
	JSON   accountCustomNNewResponseJSON `json:"-"`
	CommonResponseCustomNs
}

// accountCustomNNewResponseJSON contains the JSON metadata for the struct
// [AccountCustomNNewResponse]
type accountCustomNNewResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCustomNNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountCustomNNewResponseJSON) RawJSON() string {
	return r.raw
}

type AccountCustomNListResponse struct {
	Result []CustomNs                     `json:"result"`
	JSON   accountCustomNListResponseJSON `json:"-"`
	CustomNsCollection
}

// accountCustomNListResponseJSON contains the JSON metadata for the struct
// [AccountCustomNListResponse]
type accountCustomNListResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCustomNListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountCustomNListResponseJSON) RawJSON() string {
	return r.raw
}

type AccountCustomNDeleteResponse struct {
	Result []string                         `json:"result"`
	JSON   accountCustomNDeleteResponseJSON `json:"-"`
	CustomNsCollection
}

// accountCustomNDeleteResponseJSON contains the JSON metadata for the struct
// [AccountCustomNDeleteResponse]
type accountCustomNDeleteResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCustomNDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountCustomNDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type AccountCustomNNewParams struct {
	// The FQDN of the name server.
	NsName param.Field[string] `json:"ns_name,required" format:"hostname"`
	// The number of the set that this name server belongs to.
	NsSet param.Field[float64] `json:"ns_set"`
}

func (r AccountCustomNNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountCustomNDeleteParams struct {
	Body interface{} `json:"body,required"`
}

func (r AccountCustomNDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}
