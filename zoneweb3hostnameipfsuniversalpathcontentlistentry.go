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

// ZoneWeb3HostnameIpfsUniversalPathContentListEntryService contains methods and
// other services that help with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneWeb3HostnameIpfsUniversalPathContentListEntryService] method
// instead.
type ZoneWeb3HostnameIpfsUniversalPathContentListEntryService struct {
	Options []option.RequestOption
}

// NewZoneWeb3HostnameIpfsUniversalPathContentListEntryService generates a new
// service that applies the given options to each request. These options are
// applied after the parent client's options (if there is one), and before any
// request-specific options.
func NewZoneWeb3HostnameIpfsUniversalPathContentListEntryService(opts ...option.RequestOption) (r *ZoneWeb3HostnameIpfsUniversalPathContentListEntryService) {
	r = &ZoneWeb3HostnameIpfsUniversalPathContentListEntryService{}
	r.Options = opts
	return
}

// Create IPFS Universal Path Gateway Content List Entry
func (r *ZoneWeb3HostnameIpfsUniversalPathContentListEntryService) New(ctx context.Context, zoneID string, identifier string, body ZoneWeb3HostnameIpfsUniversalPathContentListEntryNewParams, opts ...option.RequestOption) (res *EntrySingleResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if identifier == "" {
		err = errors.New("missing required identifier parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/web3/hostnames/%s/ipfs_universal_path/content_list/entries", zoneID, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// IPFS Universal Path Gateway Content List Entry Details
func (r *ZoneWeb3HostnameIpfsUniversalPathContentListEntryService) Get(ctx context.Context, zoneID string, identifier string, contentListEntryIdentifier string, opts ...option.RequestOption) (res *EntrySingleResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if identifier == "" {
		err = errors.New("missing required identifier parameter")
		return
	}
	if contentListEntryIdentifier == "" {
		err = errors.New("missing required content_list_entry_identifier parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/web3/hostnames/%s/ipfs_universal_path/content_list/entries/%s", zoneID, identifier, contentListEntryIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Edit IPFS Universal Path Gateway Content List Entry
func (r *ZoneWeb3HostnameIpfsUniversalPathContentListEntryService) Update(ctx context.Context, zoneID string, identifier string, contentListEntryIdentifier string, body ZoneWeb3HostnameIpfsUniversalPathContentListEntryUpdateParams, opts ...option.RequestOption) (res *EntrySingleResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if identifier == "" {
		err = errors.New("missing required identifier parameter")
		return
	}
	if contentListEntryIdentifier == "" {
		err = errors.New("missing required content_list_entry_identifier parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/web3/hostnames/%s/ipfs_universal_path/content_list/entries/%s", zoneID, identifier, contentListEntryIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// List IPFS Universal Path Gateway Content List Entries
func (r *ZoneWeb3HostnameIpfsUniversalPathContentListEntryService) List(ctx context.Context, zoneID string, identifier string, opts ...option.RequestOption) (res *ZoneWeb3HostnameIpfsUniversalPathContentListEntryListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if identifier == "" {
		err = errors.New("missing required identifier parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/web3/hostnames/%s/ipfs_universal_path/content_list/entries", zoneID, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete IPFS Universal Path Gateway Content List Entry
func (r *ZoneWeb3HostnameIpfsUniversalPathContentListEntryService) Delete(ctx context.Context, zoneID string, identifier string, contentListEntryIdentifier string, opts ...option.RequestOption) (res *APIResponseSingleID, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if identifier == "" {
		err = errors.New("missing required identifier parameter")
		return
	}
	if contentListEntryIdentifier == "" {
		err = errors.New("missing required content_list_entry_identifier parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/web3/hostnames/%s/ipfs_universal_path/content_list/entries/%s", zoneID, identifier, contentListEntryIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Specify a content list entry to block.
type ContentListEntry struct {
	// Specify the identifier of the hostname.
	ID string `json:"id"`
	// Specify the CID or content path of content to block.
	Content   string    `json:"content"`
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// Specify an optional description of the content list entry.
	Description string    `json:"description"`
	ModifiedOn  time.Time `json:"modified_on" format:"date-time"`
	// Specify the type of content list entry to block.
	Type EntryType            `json:"type"`
	JSON contentListEntryJSON `json:"-"`
}

// contentListEntryJSON contains the JSON metadata for the struct
// [ContentListEntry]
type contentListEntryJSON struct {
	ID          apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Description apijson.Field
	ModifiedOn  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContentListEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contentListEntryJSON) RawJSON() string {
	return r.raw
}

// Specify a content list entry to block.
type ContentListEntryParam struct {
	// Specify the CID or content path of content to block.
	Content param.Field[string] `json:"content"`
	// Specify an optional description of the content list entry.
	Description param.Field[string] `json:"description"`
	// Specify the type of content list entry to block.
	Type param.Field[EntryType] `json:"type"`
}

func (r ContentListEntryParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EntryCreateRequestParam struct {
	// Specify the CID or content path of content to block.
	Content param.Field[string] `json:"content,required"`
	// Specify the type of content list entry to block.
	Type param.Field[EntryType] `json:"type,required"`
	// Specify an optional description of the content list entry.
	Description param.Field[string] `json:"description"`
}

func (r EntryCreateRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EntrySingleResponse struct {
	Errors   []EntrySingleResponseError   `json:"errors,required"`
	Messages []EntrySingleResponseMessage `json:"messages,required"`
	// Specify a content list entry to block.
	Result ContentListEntry `json:"result,required"`
	// Specifies whether the API call was successful.
	Success EntrySingleResponseSuccess `json:"success,required"`
	// Provides the API response.
	ResultInfo interface{}             `json:"result_info"`
	JSON       entrySingleResponseJSON `json:"-"`
}

// entrySingleResponseJSON contains the JSON metadata for the struct
// [EntrySingleResponse]
type entrySingleResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EntrySingleResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r entrySingleResponseJSON) RawJSON() string {
	return r.raw
}

type EntrySingleResponseError struct {
	Code             int64                           `json:"code,required"`
	Message          string                          `json:"message,required"`
	DocumentationURL string                          `json:"documentation_url"`
	Source           EntrySingleResponseErrorsSource `json:"source"`
	JSON             entrySingleResponseErrorJSON    `json:"-"`
}

// entrySingleResponseErrorJSON contains the JSON metadata for the struct
// [EntrySingleResponseError]
type entrySingleResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *EntrySingleResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r entrySingleResponseErrorJSON) RawJSON() string {
	return r.raw
}

type EntrySingleResponseErrorsSource struct {
	Pointer string                              `json:"pointer"`
	JSON    entrySingleResponseErrorsSourceJSON `json:"-"`
}

// entrySingleResponseErrorsSourceJSON contains the JSON metadata for the struct
// [EntrySingleResponseErrorsSource]
type entrySingleResponseErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EntrySingleResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r entrySingleResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type EntrySingleResponseMessage struct {
	Code             int64                             `json:"code,required"`
	Message          string                            `json:"message,required"`
	DocumentationURL string                            `json:"documentation_url"`
	Source           EntrySingleResponseMessagesSource `json:"source"`
	JSON             entrySingleResponseMessageJSON    `json:"-"`
}

// entrySingleResponseMessageJSON contains the JSON metadata for the struct
// [EntrySingleResponseMessage]
type entrySingleResponseMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *EntrySingleResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r entrySingleResponseMessageJSON) RawJSON() string {
	return r.raw
}

type EntrySingleResponseMessagesSource struct {
	Pointer string                                `json:"pointer"`
	JSON    entrySingleResponseMessagesSourceJSON `json:"-"`
}

// entrySingleResponseMessagesSourceJSON contains the JSON metadata for the struct
// [EntrySingleResponseMessagesSource]
type entrySingleResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EntrySingleResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r entrySingleResponseMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Specifies whether the API call was successful.
type EntrySingleResponseSuccess bool

const (
	EntrySingleResponseSuccessTrue EntrySingleResponseSuccess = true
)

func (r EntrySingleResponseSuccess) IsKnown() bool {
	switch r {
	case EntrySingleResponseSuccessTrue:
		return true
	}
	return false
}

// Specify the type of content list entry to block.
type EntryType string

const (
	EntryTypeCid         EntryType = "cid"
	EntryTypeContentPath EntryType = "content_path"
)

func (r EntryType) IsKnown() bool {
	switch r {
	case EntryTypeCid, EntryTypeContentPath:
		return true
	}
	return false
}

type ZoneWeb3HostnameIpfsUniversalPathContentListEntryListResponse struct {
	Errors   []ZoneWeb3HostnameIpfsUniversalPathContentListEntryListResponseError   `json:"errors,required"`
	Messages []ZoneWeb3HostnameIpfsUniversalPathContentListEntryListResponseMessage `json:"messages,required"`
	Result   ZoneWeb3HostnameIpfsUniversalPathContentListEntryListResponseResult    `json:"result,required,nullable"`
	// Specifies whether the API call was successful.
	Success    ZoneWeb3HostnameIpfsUniversalPathContentListEntryListResponseSuccess    `json:"success,required"`
	ResultInfo ZoneWeb3HostnameIpfsUniversalPathContentListEntryListResponseResultInfo `json:"result_info"`
	JSON       zoneWeb3HostnameIpfsUniversalPathContentListEntryListResponseJSON       `json:"-"`
}

// zoneWeb3HostnameIpfsUniversalPathContentListEntryListResponseJSON contains the
// JSON metadata for the struct
// [ZoneWeb3HostnameIpfsUniversalPathContentListEntryListResponse]
type zoneWeb3HostnameIpfsUniversalPathContentListEntryListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWeb3HostnameIpfsUniversalPathContentListEntryListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneWeb3HostnameIpfsUniversalPathContentListEntryListResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneWeb3HostnameIpfsUniversalPathContentListEntryListResponseError struct {
	Code             int64                                                                     `json:"code,required"`
	Message          string                                                                    `json:"message,required"`
	DocumentationURL string                                                                    `json:"documentation_url"`
	Source           ZoneWeb3HostnameIpfsUniversalPathContentListEntryListResponseErrorsSource `json:"source"`
	JSON             zoneWeb3HostnameIpfsUniversalPathContentListEntryListResponseErrorJSON    `json:"-"`
}

// zoneWeb3HostnameIpfsUniversalPathContentListEntryListResponseErrorJSON contains
// the JSON metadata for the struct
// [ZoneWeb3HostnameIpfsUniversalPathContentListEntryListResponseError]
type zoneWeb3HostnameIpfsUniversalPathContentListEntryListResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZoneWeb3HostnameIpfsUniversalPathContentListEntryListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneWeb3HostnameIpfsUniversalPathContentListEntryListResponseErrorJSON) RawJSON() string {
	return r.raw
}

type ZoneWeb3HostnameIpfsUniversalPathContentListEntryListResponseErrorsSource struct {
	Pointer string                                                                        `json:"pointer"`
	JSON    zoneWeb3HostnameIpfsUniversalPathContentListEntryListResponseErrorsSourceJSON `json:"-"`
}

// zoneWeb3HostnameIpfsUniversalPathContentListEntryListResponseErrorsSourceJSON
// contains the JSON metadata for the struct
// [ZoneWeb3HostnameIpfsUniversalPathContentListEntryListResponseErrorsSource]
type zoneWeb3HostnameIpfsUniversalPathContentListEntryListResponseErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWeb3HostnameIpfsUniversalPathContentListEntryListResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneWeb3HostnameIpfsUniversalPathContentListEntryListResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type ZoneWeb3HostnameIpfsUniversalPathContentListEntryListResponseMessage struct {
	Code             int64                                                                       `json:"code,required"`
	Message          string                                                                      `json:"message,required"`
	DocumentationURL string                                                                      `json:"documentation_url"`
	Source           ZoneWeb3HostnameIpfsUniversalPathContentListEntryListResponseMessagesSource `json:"source"`
	JSON             zoneWeb3HostnameIpfsUniversalPathContentListEntryListResponseMessageJSON    `json:"-"`
}

// zoneWeb3HostnameIpfsUniversalPathContentListEntryListResponseMessageJSON
// contains the JSON metadata for the struct
// [ZoneWeb3HostnameIpfsUniversalPathContentListEntryListResponseMessage]
type zoneWeb3HostnameIpfsUniversalPathContentListEntryListResponseMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZoneWeb3HostnameIpfsUniversalPathContentListEntryListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneWeb3HostnameIpfsUniversalPathContentListEntryListResponseMessageJSON) RawJSON() string {
	return r.raw
}

type ZoneWeb3HostnameIpfsUniversalPathContentListEntryListResponseMessagesSource struct {
	Pointer string                                                                          `json:"pointer"`
	JSON    zoneWeb3HostnameIpfsUniversalPathContentListEntryListResponseMessagesSourceJSON `json:"-"`
}

// zoneWeb3HostnameIpfsUniversalPathContentListEntryListResponseMessagesSourceJSON
// contains the JSON metadata for the struct
// [ZoneWeb3HostnameIpfsUniversalPathContentListEntryListResponseMessagesSource]
type zoneWeb3HostnameIpfsUniversalPathContentListEntryListResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWeb3HostnameIpfsUniversalPathContentListEntryListResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneWeb3HostnameIpfsUniversalPathContentListEntryListResponseMessagesSourceJSON) RawJSON() string {
	return r.raw
}

type ZoneWeb3HostnameIpfsUniversalPathContentListEntryListResponseResult struct {
	// Provides content list entries.
	Entries []ContentListEntry                                                      `json:"entries"`
	JSON    zoneWeb3HostnameIpfsUniversalPathContentListEntryListResponseResultJSON `json:"-"`
}

// zoneWeb3HostnameIpfsUniversalPathContentListEntryListResponseResultJSON contains
// the JSON metadata for the struct
// [ZoneWeb3HostnameIpfsUniversalPathContentListEntryListResponseResult]
type zoneWeb3HostnameIpfsUniversalPathContentListEntryListResponseResultJSON struct {
	Entries     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWeb3HostnameIpfsUniversalPathContentListEntryListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneWeb3HostnameIpfsUniversalPathContentListEntryListResponseResultJSON) RawJSON() string {
	return r.raw
}

// Specifies whether the API call was successful.
type ZoneWeb3HostnameIpfsUniversalPathContentListEntryListResponseSuccess bool

const (
	ZoneWeb3HostnameIpfsUniversalPathContentListEntryListResponseSuccessTrue ZoneWeb3HostnameIpfsUniversalPathContentListEntryListResponseSuccess = true
)

func (r ZoneWeb3HostnameIpfsUniversalPathContentListEntryListResponseSuccess) IsKnown() bool {
	switch r {
	case ZoneWeb3HostnameIpfsUniversalPathContentListEntryListResponseSuccessTrue:
		return true
	}
	return false
}

type ZoneWeb3HostnameIpfsUniversalPathContentListEntryListResponseResultInfo struct {
	// Specifies the total number of results for the requested service.
	Count float64 `json:"count"`
	// Specifies the current page within paginated list of results.
	Page float64 `json:"page"`
	// Specifies the number of results per page of results.
	PerPage float64 `json:"per_page"`
	// Specifies the total results available without any search parameters.
	TotalCount float64                                                                     `json:"total_count"`
	JSON       zoneWeb3HostnameIpfsUniversalPathContentListEntryListResponseResultInfoJSON `json:"-"`
}

// zoneWeb3HostnameIpfsUniversalPathContentListEntryListResponseResultInfoJSON
// contains the JSON metadata for the struct
// [ZoneWeb3HostnameIpfsUniversalPathContentListEntryListResponseResultInfo]
type zoneWeb3HostnameIpfsUniversalPathContentListEntryListResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWeb3HostnameIpfsUniversalPathContentListEntryListResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneWeb3HostnameIpfsUniversalPathContentListEntryListResponseResultInfoJSON) RawJSON() string {
	return r.raw
}

type ZoneWeb3HostnameIpfsUniversalPathContentListEntryNewParams struct {
	EntryCreateRequest EntryCreateRequestParam `json:"entry_create_request,required"`
}

func (r ZoneWeb3HostnameIpfsUniversalPathContentListEntryNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.EntryCreateRequest)
}

type ZoneWeb3HostnameIpfsUniversalPathContentListEntryUpdateParams struct {
	EntryCreateRequest EntryCreateRequestParam `json:"entry_create_request,required"`
}

func (r ZoneWeb3HostnameIpfsUniversalPathContentListEntryUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.EntryCreateRequest)
}
