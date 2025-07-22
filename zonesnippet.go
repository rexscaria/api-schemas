// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"mime/multipart"
	"net/http"

	"github.com/stainless-sdks/cf-rex-go/internal/apiform"
	"github.com/stainless-sdks/cf-rex-go/internal/apijson"
	"github.com/stainless-sdks/cf-rex-go/internal/param"
	"github.com/stainless-sdks/cf-rex-go/internal/requestconfig"
	"github.com/stainless-sdks/cf-rex-go/option"
)

// ZoneSnippetService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneSnippetService] method instead.
type ZoneSnippetService struct {
	Options      []option.RequestOption
	SnippetRules *ZoneSnippetSnippetRuleService
}

// NewZoneSnippetService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewZoneSnippetService(opts ...option.RequestOption) (r *ZoneSnippetService) {
	r = &ZoneSnippetService{}
	r.Options = opts
	r.SnippetRules = NewZoneSnippetSnippetRuleService(opts...)
	return
}

// Snippet
func (r *ZoneSnippetService) Get(ctx context.Context, zoneID string, snippetName string, opts ...option.RequestOption) (res *ZoneSnippetGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if snippetName == "" {
		err = errors.New("missing required snippet_name parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/snippets/%s", zoneID, snippetName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Put Snippet
func (r *ZoneSnippetService) Update(ctx context.Context, zoneID string, snippetName string, body ZoneSnippetUpdateParams, opts ...option.RequestOption) (res *ZoneSnippetUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if snippetName == "" {
		err = errors.New("missing required snippet_name parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/snippets/%s", zoneID, snippetName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// All Snippets
func (r *ZoneSnippetService) List(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *ZoneSnippetListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/snippets", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete Snippet
func (r *ZoneSnippetService) Delete(ctx context.Context, zoneID string, snippetName string, opts ...option.RequestOption) (res *APIResponseSnippets, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if snippetName == "" {
		err = errors.New("missing required snippet_name parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/snippets/%s", zoneID, snippetName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Snippet Content
func (r *ZoneSnippetService) GetContent(ctx context.Context, zoneID string, snippetName string, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "multipart/form-data")}, opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if snippetName == "" {
		err = errors.New("missing required snippet_name parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/snippets/%s/content", zoneID, snippetName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type APIResponseSnippets struct {
	Errors   []SnippetMessageItem `json:"errors,required"`
	Messages []SnippetMessageItem `json:"messages,required"`
	// Whether the API call was successful
	Success APIResponseSnippetsSuccess `json:"success,required"`
	JSON    apiResponseSnippetsJSON    `json:"-"`
}

// apiResponseSnippetsJSON contains the JSON metadata for the struct
// [APIResponseSnippets]
type apiResponseSnippetsJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIResponseSnippets) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiResponseSnippetsJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type APIResponseSnippetsSuccess bool

const (
	APIResponseSnippetsSuccessTrue APIResponseSnippetsSuccess = true
)

func (r APIResponseSnippetsSuccess) IsKnown() bool {
	switch r {
	case APIResponseSnippetsSuccessTrue:
		return true
	}
	return false
}

// Snippet Information
type Snippet struct {
	// Creation time of the snippet
	CreatedOn string `json:"created_on"`
	// Modification time of the snippet
	ModifiedOn string `json:"modified_on"`
	// Snippet identifying name
	SnippetName string      `json:"snippet_name"`
	JSON        snippetJSON `json:"-"`
}

// snippetJSON contains the JSON metadata for the struct [Snippet]
type snippetJSON struct {
	CreatedOn   apijson.Field
	ModifiedOn  apijson.Field
	SnippetName apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Snippet) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r snippetJSON) RawJSON() string {
	return r.raw
}

type SnippetMessageItem struct {
	Code    int64                  `json:"code,required"`
	Message string                 `json:"message,required"`
	JSON    snippetMessageItemJSON `json:"-"`
}

// snippetMessageItemJSON contains the JSON metadata for the struct
// [SnippetMessageItem]
type snippetMessageItemJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SnippetMessageItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r snippetMessageItemJSON) RawJSON() string {
	return r.raw
}

type ZoneSnippetGetResponse struct {
	// Snippet Information
	Result Snippet                    `json:"result"`
	JSON   zoneSnippetGetResponseJSON `json:"-"`
	APIResponseSnippets
}

// zoneSnippetGetResponseJSON contains the JSON metadata for the struct
// [ZoneSnippetGetResponse]
type zoneSnippetGetResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSnippetGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSnippetGetResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneSnippetUpdateResponse struct {
	// Snippet Information
	Result Snippet                       `json:"result"`
	JSON   zoneSnippetUpdateResponseJSON `json:"-"`
	APIResponseSnippets
}

// zoneSnippetUpdateResponseJSON contains the JSON metadata for the struct
// [ZoneSnippetUpdateResponse]
type zoneSnippetUpdateResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSnippetUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSnippetUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneSnippetListResponse struct {
	// List of all zone snippets
	Result []Snippet                   `json:"result"`
	JSON   zoneSnippetListResponseJSON `json:"-"`
	APIResponseSnippets
}

// zoneSnippetListResponseJSON contains the JSON metadata for the struct
// [ZoneSnippetListResponse]
type zoneSnippetListResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSnippetListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSnippetListResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneSnippetUpdateParams struct {
	// Content files of uploaded snippet
	Files    param.Field[string]                          `json:"files"`
	Metadata param.Field[ZoneSnippetUpdateParamsMetadata] `json:"metadata"`
}

func (r ZoneSnippetUpdateParams) MarshalMultipart() (data []byte, contentType string, err error) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	err = apiform.MarshalRoot(r, writer)
	if err != nil {
		writer.Close()
		return nil, "", err
	}
	err = writer.Close()
	if err != nil {
		return nil, "", err
	}
	return buf.Bytes(), writer.FormDataContentType(), nil
}

type ZoneSnippetUpdateParamsMetadata struct {
	// Main module name of uploaded snippet
	MainModule param.Field[string] `json:"main_module"`
}

func (r ZoneSnippetUpdateParamsMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
