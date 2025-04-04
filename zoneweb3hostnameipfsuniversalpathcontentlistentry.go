// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/stainless-sdks/cf-rex-go/internal/apijson"
	"github.com/stainless-sdks/cf-rex-go/internal/param"
	"github.com/stainless-sdks/cf-rex-go/internal/requestconfig"
	"github.com/stainless-sdks/cf-rex-go/option"
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
func (r *ZoneWeb3HostnameIpfsUniversalPathContentListEntryService) Delete(ctx context.Context, zoneID string, identifier string, contentListEntryIdentifier string, body ZoneWeb3HostnameIpfsUniversalPathContentListEntryDeleteParams, opts ...option.RequestOption) (res *APIResponseSingleID, err error) {
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
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

// Content list entry to be blocked.
type ContentListEntry struct {
	// Identifier
	ID string `json:"id"`
	// CID or content path of content to block.
	Content   string    `json:"content"`
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// An optional description of the content list entry.
	Description string    `json:"description"`
	ModifiedOn  time.Time `json:"modified_on" format:"date-time"`
	// Type of content list entry to block.
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

// Content list entry to be blocked.
type ContentListEntryParam struct {
	// CID or content path of content to block.
	Content param.Field[string] `json:"content"`
	// An optional description of the content list entry.
	Description param.Field[string] `json:"description"`
	// Type of content list entry to block.
	Type param.Field[EntryType] `json:"type"`
}

func (r ContentListEntryParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EntryCreateRequestParam struct {
	// CID or content path of content to block.
	Content param.Field[string] `json:"content,required"`
	// Type of content list entry to block.
	Type param.Field[EntryType] `json:"type,required"`
	// An optional description of the content list entry.
	Description param.Field[string] `json:"description"`
}

func (r EntryCreateRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EntrySingleResponse struct {
	// Content list entry to be blocked.
	Result ContentListEntry        `json:"result"`
	JSON   entrySingleResponseJSON `json:"-"`
	APIResponseSingleWeb3
}

// entrySingleResponseJSON contains the JSON metadata for the struct
// [EntrySingleResponse]
type entrySingleResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EntrySingleResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r entrySingleResponseJSON) RawJSON() string {
	return r.raw
}

// Type of content list entry to block.
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
	Result ZoneWeb3HostnameIpfsUniversalPathContentListEntryListResponseResult `json:"result"`
	JSON   zoneWeb3HostnameIpfsUniversalPathContentListEntryListResponseJSON   `json:"-"`
	APIResponseCollectionWeb3
}

// zoneWeb3HostnameIpfsUniversalPathContentListEntryListResponseJSON contains the
// JSON metadata for the struct
// [ZoneWeb3HostnameIpfsUniversalPathContentListEntryListResponse]
type zoneWeb3HostnameIpfsUniversalPathContentListEntryListResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWeb3HostnameIpfsUniversalPathContentListEntryListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneWeb3HostnameIpfsUniversalPathContentListEntryListResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneWeb3HostnameIpfsUniversalPathContentListEntryListResponseResult struct {
	// Content list entries.
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

type ZoneWeb3HostnameIpfsUniversalPathContentListEntryDeleteParams struct {
	Body interface{} `json:"body,required"`
}

func (r ZoneWeb3HostnameIpfsUniversalPathContentListEntryDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}
