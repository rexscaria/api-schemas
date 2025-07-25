// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"time"

	"github.com/rexscaria/api-schemas/internal/apiform"
	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/apiquery"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
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

// Fetches a snippet belonging to the zone.
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

// Creates or updates a snippet belonging to the zone.
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

// Fetches all snippets belonging to the zone.
func (r *ZoneSnippetService) List(ctx context.Context, zoneID string, query ZoneSnippetListParams, opts ...option.RequestOption) (res *ZoneSnippetListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/snippets", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Deletes a snippet belonging to the zone.
func (r *ZoneSnippetService) Delete(ctx context.Context, zoneID string, snippetName string, opts ...option.RequestOption) (res *ZoneSnippetDeleteResponse, err error) {
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

// Fetches the content of a snippet belonging to the zone.
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

// A response object.
type ZoneSnippetGetResponse struct {
	// A list of error messages.
	Errors []ZoneSnippetGetResponseError `json:"errors,required"`
	// A list of warning messages.
	Messages []ZoneSnippetGetResponseMessage `json:"messages,required"`
	// A result.
	Result ZoneSnippetGetResponseResult `json:"result,required"`
	// Whether the API call was successful.
	Success ZoneSnippetGetResponseSuccess `json:"success,required"`
	JSON    zoneSnippetGetResponseJSON    `json:"-"`
}

// zoneSnippetGetResponseJSON contains the JSON metadata for the struct
// [ZoneSnippetGetResponse]
type zoneSnippetGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSnippetGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSnippetGetResponseJSON) RawJSON() string {
	return r.raw
}

// A message.
type ZoneSnippetGetResponseError struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64                           `json:"code"`
	JSON zoneSnippetGetResponseErrorJSON `json:"-"`
}

// zoneSnippetGetResponseErrorJSON contains the JSON metadata for the struct
// [ZoneSnippetGetResponseError]
type zoneSnippetGetResponseErrorJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSnippetGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSnippetGetResponseErrorJSON) RawJSON() string {
	return r.raw
}

// A message.
type ZoneSnippetGetResponseMessage struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64                             `json:"code"`
	JSON zoneSnippetGetResponseMessageJSON `json:"-"`
}

// zoneSnippetGetResponseMessageJSON contains the JSON metadata for the struct
// [ZoneSnippetGetResponseMessage]
type zoneSnippetGetResponseMessageJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSnippetGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSnippetGetResponseMessageJSON) RawJSON() string {
	return r.raw
}

// A result.
type ZoneSnippetGetResponseResult struct {
	// The timestamp of when the snippet was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// The identifying name of the snippet.
	SnippetName string `json:"snippet_name,required"`
	// The timestamp of when the snippet was last modified.
	ModifiedOn time.Time                        `json:"modified_on" format:"date-time"`
	JSON       zoneSnippetGetResponseResultJSON `json:"-"`
}

// zoneSnippetGetResponseResultJSON contains the JSON metadata for the struct
// [ZoneSnippetGetResponseResult]
type zoneSnippetGetResponseResultJSON struct {
	CreatedOn   apijson.Field
	SnippetName apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSnippetGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSnippetGetResponseResultJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ZoneSnippetGetResponseSuccess bool

const (
	ZoneSnippetGetResponseSuccessTrue ZoneSnippetGetResponseSuccess = true
)

func (r ZoneSnippetGetResponseSuccess) IsKnown() bool {
	switch r {
	case ZoneSnippetGetResponseSuccessTrue:
		return true
	}
	return false
}

// A response object.
type ZoneSnippetUpdateResponse struct {
	// A list of error messages.
	Errors []ZoneSnippetUpdateResponseError `json:"errors,required"`
	// A list of warning messages.
	Messages []ZoneSnippetUpdateResponseMessage `json:"messages,required"`
	// A result.
	Result ZoneSnippetUpdateResponseResult `json:"result,required"`
	// Whether the API call was successful.
	Success ZoneSnippetUpdateResponseSuccess `json:"success,required"`
	JSON    zoneSnippetUpdateResponseJSON    `json:"-"`
}

// zoneSnippetUpdateResponseJSON contains the JSON metadata for the struct
// [ZoneSnippetUpdateResponse]
type zoneSnippetUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSnippetUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSnippetUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// A message.
type ZoneSnippetUpdateResponseError struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64                              `json:"code"`
	JSON zoneSnippetUpdateResponseErrorJSON `json:"-"`
}

// zoneSnippetUpdateResponseErrorJSON contains the JSON metadata for the struct
// [ZoneSnippetUpdateResponseError]
type zoneSnippetUpdateResponseErrorJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSnippetUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSnippetUpdateResponseErrorJSON) RawJSON() string {
	return r.raw
}

// A message.
type ZoneSnippetUpdateResponseMessage struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64                                `json:"code"`
	JSON zoneSnippetUpdateResponseMessageJSON `json:"-"`
}

// zoneSnippetUpdateResponseMessageJSON contains the JSON metadata for the struct
// [ZoneSnippetUpdateResponseMessage]
type zoneSnippetUpdateResponseMessageJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSnippetUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSnippetUpdateResponseMessageJSON) RawJSON() string {
	return r.raw
}

// A result.
type ZoneSnippetUpdateResponseResult struct {
	// The timestamp of when the snippet was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// The identifying name of the snippet.
	SnippetName string `json:"snippet_name,required"`
	// The timestamp of when the snippet was last modified.
	ModifiedOn time.Time                           `json:"modified_on" format:"date-time"`
	JSON       zoneSnippetUpdateResponseResultJSON `json:"-"`
}

// zoneSnippetUpdateResponseResultJSON contains the JSON metadata for the struct
// [ZoneSnippetUpdateResponseResult]
type zoneSnippetUpdateResponseResultJSON struct {
	CreatedOn   apijson.Field
	SnippetName apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSnippetUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSnippetUpdateResponseResultJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ZoneSnippetUpdateResponseSuccess bool

const (
	ZoneSnippetUpdateResponseSuccessTrue ZoneSnippetUpdateResponseSuccess = true
)

func (r ZoneSnippetUpdateResponseSuccess) IsKnown() bool {
	switch r {
	case ZoneSnippetUpdateResponseSuccessTrue:
		return true
	}
	return false
}

// A response object.
type ZoneSnippetListResponse struct {
	// A list of error messages.
	Errors []ZoneSnippetListResponseError `json:"errors,required"`
	// A list of warning messages.
	Messages []ZoneSnippetListResponseMessage `json:"messages,required"`
	// A list of snippets.
	Result []ZoneSnippetListResponseResult `json:"result,required"`
	// Whether the API call was successful.
	Success ZoneSnippetListResponseSuccess `json:"success,required"`
	// Additional information to navigate the results.
	ResultInfo ZoneSnippetListResponseResultInfo `json:"result_info"`
	JSON       zoneSnippetListResponseJSON       `json:"-"`
}

// zoneSnippetListResponseJSON contains the JSON metadata for the struct
// [ZoneSnippetListResponse]
type zoneSnippetListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSnippetListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSnippetListResponseJSON) RawJSON() string {
	return r.raw
}

// A message.
type ZoneSnippetListResponseError struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64                            `json:"code"`
	JSON zoneSnippetListResponseErrorJSON `json:"-"`
}

// zoneSnippetListResponseErrorJSON contains the JSON metadata for the struct
// [ZoneSnippetListResponseError]
type zoneSnippetListResponseErrorJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSnippetListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSnippetListResponseErrorJSON) RawJSON() string {
	return r.raw
}

// A message.
type ZoneSnippetListResponseMessage struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64                              `json:"code"`
	JSON zoneSnippetListResponseMessageJSON `json:"-"`
}

// zoneSnippetListResponseMessageJSON contains the JSON metadata for the struct
// [ZoneSnippetListResponseMessage]
type zoneSnippetListResponseMessageJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSnippetListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSnippetListResponseMessageJSON) RawJSON() string {
	return r.raw
}

// A snippet object.
type ZoneSnippetListResponseResult struct {
	// The timestamp of when the snippet was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// The identifying name of the snippet.
	SnippetName string `json:"snippet_name,required"`
	// The timestamp of when the snippet was last modified.
	ModifiedOn time.Time                         `json:"modified_on" format:"date-time"`
	JSON       zoneSnippetListResponseResultJSON `json:"-"`
}

// zoneSnippetListResponseResultJSON contains the JSON metadata for the struct
// [ZoneSnippetListResponseResult]
type zoneSnippetListResponseResultJSON struct {
	CreatedOn   apijson.Field
	SnippetName apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSnippetListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSnippetListResponseResultJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ZoneSnippetListResponseSuccess bool

const (
	ZoneSnippetListResponseSuccessTrue ZoneSnippetListResponseSuccess = true
)

func (r ZoneSnippetListResponseSuccess) IsKnown() bool {
	switch r {
	case ZoneSnippetListResponseSuccessTrue:
		return true
	}
	return false
}

// Additional information to navigate the results.
type ZoneSnippetListResponseResultInfo struct {
	// The number of results in the current page.
	Count int64 `json:"count,required"`
	// The current page number.
	Page int64 `json:"page,required"`
	// The number of results to return per page.
	PerPage int64 `json:"per_page,required"`
	// The total number of results.
	TotalCount int64 `json:"total_count,required"`
	// The total number of pages.
	TotalPages int64                                 `json:"total_pages,required"`
	JSON       zoneSnippetListResponseResultInfoJSON `json:"-"`
}

// zoneSnippetListResponseResultInfoJSON contains the JSON metadata for the struct
// [ZoneSnippetListResponseResultInfo]
type zoneSnippetListResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	TotalPages  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSnippetListResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSnippetListResponseResultInfoJSON) RawJSON() string {
	return r.raw
}

// A response object.
type ZoneSnippetDeleteResponse struct {
	// A list of error messages.
	Errors []ZoneSnippetDeleteResponseError `json:"errors,required"`
	// A list of warning messages.
	Messages []ZoneSnippetDeleteResponseMessage `json:"messages,required"`
	// A result.
	Result string `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success ZoneSnippetDeleteResponseSuccess `json:"success,required"`
	JSON    zoneSnippetDeleteResponseJSON    `json:"-"`
}

// zoneSnippetDeleteResponseJSON contains the JSON metadata for the struct
// [ZoneSnippetDeleteResponse]
type zoneSnippetDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSnippetDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSnippetDeleteResponseJSON) RawJSON() string {
	return r.raw
}

// A message.
type ZoneSnippetDeleteResponseError struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64                              `json:"code"`
	JSON zoneSnippetDeleteResponseErrorJSON `json:"-"`
}

// zoneSnippetDeleteResponseErrorJSON contains the JSON metadata for the struct
// [ZoneSnippetDeleteResponseError]
type zoneSnippetDeleteResponseErrorJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSnippetDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSnippetDeleteResponseErrorJSON) RawJSON() string {
	return r.raw
}

// A message.
type ZoneSnippetDeleteResponseMessage struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64                                `json:"code"`
	JSON zoneSnippetDeleteResponseMessageJSON `json:"-"`
}

// zoneSnippetDeleteResponseMessageJSON contains the JSON metadata for the struct
// [ZoneSnippetDeleteResponseMessage]
type zoneSnippetDeleteResponseMessageJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSnippetDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSnippetDeleteResponseMessageJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ZoneSnippetDeleteResponseSuccess bool

const (
	ZoneSnippetDeleteResponseSuccessTrue ZoneSnippetDeleteResponseSuccess = true
)

func (r ZoneSnippetDeleteResponseSuccess) IsKnown() bool {
	switch r {
	case ZoneSnippetDeleteResponseSuccessTrue:
		return true
	}
	return false
}

type ZoneSnippetUpdateParams struct {
	// The list of files belonging to the snippet.
	Files param.Field[[]io.Reader] `json:"files,required" format:"binary"`
	// Metadata about the snippet.
	Metadata param.Field[ZoneSnippetUpdateParamsMetadata] `json:"metadata,required"`
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

// Metadata about the snippet.
type ZoneSnippetUpdateParamsMetadata struct {
	// Name of the file that contains the main module of the snippet.
	MainModule param.Field[string] `json:"main_module,required"`
}

func (r ZoneSnippetUpdateParamsMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneSnippetListParams struct {
	// The current page number.
	Page param.Field[int64] `query:"page"`
	// The number of results to return per page.
	PerPage param.Field[int64] `query:"per_page"`
}

// URLQuery serializes [ZoneSnippetListParams]'s query parameters as `url.Values`.
func (r ZoneSnippetListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
