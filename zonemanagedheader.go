// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
)

// ZoneManagedHeaderService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneManagedHeaderService] method instead.
type ZoneManagedHeaderService struct {
	Options []option.RequestOption
}

// NewZoneManagedHeaderService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneManagedHeaderService(opts ...option.RequestOption) (r *ZoneManagedHeaderService) {
	r = &ZoneManagedHeaderService{}
	r.Options = opts
	return
}

// Updates the status of one or more Managed Transforms.
func (r *ZoneManagedHeaderService) Update(ctx context.Context, zoneID string, body ZoneManagedHeaderUpdateParams, opts ...option.RequestOption) (res *ZoneManagedHeaderUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/managed_headers", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Fetches a list of all Managed Transforms.
func (r *ZoneManagedHeaderService) List(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *ZoneManagedHeaderListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/managed_headers", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Disables all Managed Transforms.
func (r *ZoneManagedHeaderService) Delete(ctx context.Context, zoneID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return err
	}
	path := fmt.Sprintf("zones/%s/managed_headers", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// A response object.
type ZoneManagedHeaderUpdateResponse struct {
	// A list of error messages.
	Errors []ZoneManagedHeaderUpdateResponseError `json:"errors" api:"required"`
	// A list of warning messages.
	Messages []ZoneManagedHeaderUpdateResponseMessage `json:"messages" api:"required"`
	// A result.
	Result ZoneManagedHeaderUpdateResponseResult `json:"result" api:"required"`
	// Whether the API call was successful.
	Success ZoneManagedHeaderUpdateResponseSuccess `json:"success" api:"required"`
	JSON    zoneManagedHeaderUpdateResponseJSON    `json:"-"`
}

// zoneManagedHeaderUpdateResponseJSON contains the JSON metadata for the struct
// [ZoneManagedHeaderUpdateResponse]
type zoneManagedHeaderUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneManagedHeaderUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneManagedHeaderUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// A message.
type ZoneManagedHeaderUpdateResponseError struct {
	// A text description of this message.
	Message string `json:"message" api:"required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source ZoneManagedHeaderUpdateResponseErrorsSource `json:"source"`
	JSON   zoneManagedHeaderUpdateResponseErrorJSON    `json:"-"`
}

// zoneManagedHeaderUpdateResponseErrorJSON contains the JSON metadata for the
// struct [ZoneManagedHeaderUpdateResponseError]
type zoneManagedHeaderUpdateResponseErrorJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneManagedHeaderUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneManagedHeaderUpdateResponseErrorJSON) RawJSON() string {
	return r.raw
}

// The source of this message.
type ZoneManagedHeaderUpdateResponseErrorsSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                          `json:"pointer" api:"required"`
	JSON    zoneManagedHeaderUpdateResponseErrorsSourceJSON `json:"-"`
}

// zoneManagedHeaderUpdateResponseErrorsSourceJSON contains the JSON metadata for
// the struct [ZoneManagedHeaderUpdateResponseErrorsSource]
type zoneManagedHeaderUpdateResponseErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneManagedHeaderUpdateResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneManagedHeaderUpdateResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

// A message.
type ZoneManagedHeaderUpdateResponseMessage struct {
	// A text description of this message.
	Message string `json:"message" api:"required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source ZoneManagedHeaderUpdateResponseMessagesSource `json:"source"`
	JSON   zoneManagedHeaderUpdateResponseMessageJSON    `json:"-"`
}

// zoneManagedHeaderUpdateResponseMessageJSON contains the JSON metadata for the
// struct [ZoneManagedHeaderUpdateResponseMessage]
type zoneManagedHeaderUpdateResponseMessageJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneManagedHeaderUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneManagedHeaderUpdateResponseMessageJSON) RawJSON() string {
	return r.raw
}

// The source of this message.
type ZoneManagedHeaderUpdateResponseMessagesSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                            `json:"pointer" api:"required"`
	JSON    zoneManagedHeaderUpdateResponseMessagesSourceJSON `json:"-"`
}

// zoneManagedHeaderUpdateResponseMessagesSourceJSON contains the JSON metadata for
// the struct [ZoneManagedHeaderUpdateResponseMessagesSource]
type zoneManagedHeaderUpdateResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneManagedHeaderUpdateResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneManagedHeaderUpdateResponseMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// A result.
type ZoneManagedHeaderUpdateResponseResult struct {
	// The list of Managed Request Transforms.
	ManagedRequestHeaders []ZoneManagedHeaderUpdateResponseResultManagedRequestHeader `json:"managed_request_headers" api:"required"`
	// The list of Managed Response Transforms.
	ManagedResponseHeaders []ZoneManagedHeaderUpdateResponseResultManagedResponseHeader `json:"managed_response_headers" api:"required"`
	JSON                   zoneManagedHeaderUpdateResponseResultJSON                    `json:"-"`
}

// zoneManagedHeaderUpdateResponseResultJSON contains the JSON metadata for the
// struct [ZoneManagedHeaderUpdateResponseResult]
type zoneManagedHeaderUpdateResponseResultJSON struct {
	ManagedRequestHeaders  apijson.Field
	ManagedResponseHeaders apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZoneManagedHeaderUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneManagedHeaderUpdateResponseResultJSON) RawJSON() string {
	return r.raw
}

// A Managed Transform object.
type ZoneManagedHeaderUpdateResponseResultManagedRequestHeader struct {
	// The human-readable identifier of the Managed Transform.
	ID string `json:"id" api:"required"`
	// Whether the Managed Transform is enabled.
	Enabled bool `json:"enabled" api:"required"`
	// Whether the Managed Transform conflicts with the currently-enabled Managed
	// Transforms.
	HasConflict bool `json:"has_conflict" api:"required"`
	// The Managed Transforms that this Managed Transform conflicts with.
	ConflictsWith []string                                                      `json:"conflicts_with"`
	JSON          zoneManagedHeaderUpdateResponseResultManagedRequestHeaderJSON `json:"-"`
}

// zoneManagedHeaderUpdateResponseResultManagedRequestHeaderJSON contains the JSON
// metadata for the struct
// [ZoneManagedHeaderUpdateResponseResultManagedRequestHeader]
type zoneManagedHeaderUpdateResponseResultManagedRequestHeaderJSON struct {
	ID            apijson.Field
	Enabled       apijson.Field
	HasConflict   apijson.Field
	ConflictsWith apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ZoneManagedHeaderUpdateResponseResultManagedRequestHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneManagedHeaderUpdateResponseResultManagedRequestHeaderJSON) RawJSON() string {
	return r.raw
}

// A Managed Transform object.
type ZoneManagedHeaderUpdateResponseResultManagedResponseHeader struct {
	// The human-readable identifier of the Managed Transform.
	ID string `json:"id" api:"required"`
	// Whether the Managed Transform is enabled.
	Enabled bool `json:"enabled" api:"required"`
	// Whether the Managed Transform conflicts with the currently-enabled Managed
	// Transforms.
	HasConflict bool `json:"has_conflict" api:"required"`
	// The Managed Transforms that this Managed Transform conflicts with.
	ConflictsWith []string                                                       `json:"conflicts_with"`
	JSON          zoneManagedHeaderUpdateResponseResultManagedResponseHeaderJSON `json:"-"`
}

// zoneManagedHeaderUpdateResponseResultManagedResponseHeaderJSON contains the JSON
// metadata for the struct
// [ZoneManagedHeaderUpdateResponseResultManagedResponseHeader]
type zoneManagedHeaderUpdateResponseResultManagedResponseHeaderJSON struct {
	ID            apijson.Field
	Enabled       apijson.Field
	HasConflict   apijson.Field
	ConflictsWith apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ZoneManagedHeaderUpdateResponseResultManagedResponseHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneManagedHeaderUpdateResponseResultManagedResponseHeaderJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ZoneManagedHeaderUpdateResponseSuccess bool

const (
	ZoneManagedHeaderUpdateResponseSuccessTrue ZoneManagedHeaderUpdateResponseSuccess = true
)

func (r ZoneManagedHeaderUpdateResponseSuccess) IsKnown() bool {
	switch r {
	case ZoneManagedHeaderUpdateResponseSuccessTrue:
		return true
	}
	return false
}

// A response object.
type ZoneManagedHeaderListResponse struct {
	// A list of error messages.
	Errors []ZoneManagedHeaderListResponseError `json:"errors" api:"required"`
	// A list of warning messages.
	Messages []ZoneManagedHeaderListResponseMessage `json:"messages" api:"required"`
	// A result.
	Result ZoneManagedHeaderListResponseResult `json:"result" api:"required"`
	// Whether the API call was successful.
	Success ZoneManagedHeaderListResponseSuccess `json:"success" api:"required"`
	JSON    zoneManagedHeaderListResponseJSON    `json:"-"`
}

// zoneManagedHeaderListResponseJSON contains the JSON metadata for the struct
// [ZoneManagedHeaderListResponse]
type zoneManagedHeaderListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneManagedHeaderListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneManagedHeaderListResponseJSON) RawJSON() string {
	return r.raw
}

// A message.
type ZoneManagedHeaderListResponseError struct {
	// A text description of this message.
	Message string `json:"message" api:"required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source ZoneManagedHeaderListResponseErrorsSource `json:"source"`
	JSON   zoneManagedHeaderListResponseErrorJSON    `json:"-"`
}

// zoneManagedHeaderListResponseErrorJSON contains the JSON metadata for the struct
// [ZoneManagedHeaderListResponseError]
type zoneManagedHeaderListResponseErrorJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneManagedHeaderListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneManagedHeaderListResponseErrorJSON) RawJSON() string {
	return r.raw
}

// The source of this message.
type ZoneManagedHeaderListResponseErrorsSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                        `json:"pointer" api:"required"`
	JSON    zoneManagedHeaderListResponseErrorsSourceJSON `json:"-"`
}

// zoneManagedHeaderListResponseErrorsSourceJSON contains the JSON metadata for the
// struct [ZoneManagedHeaderListResponseErrorsSource]
type zoneManagedHeaderListResponseErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneManagedHeaderListResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneManagedHeaderListResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

// A message.
type ZoneManagedHeaderListResponseMessage struct {
	// A text description of this message.
	Message string `json:"message" api:"required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source ZoneManagedHeaderListResponseMessagesSource `json:"source"`
	JSON   zoneManagedHeaderListResponseMessageJSON    `json:"-"`
}

// zoneManagedHeaderListResponseMessageJSON contains the JSON metadata for the
// struct [ZoneManagedHeaderListResponseMessage]
type zoneManagedHeaderListResponseMessageJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneManagedHeaderListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneManagedHeaderListResponseMessageJSON) RawJSON() string {
	return r.raw
}

// The source of this message.
type ZoneManagedHeaderListResponseMessagesSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                          `json:"pointer" api:"required"`
	JSON    zoneManagedHeaderListResponseMessagesSourceJSON `json:"-"`
}

// zoneManagedHeaderListResponseMessagesSourceJSON contains the JSON metadata for
// the struct [ZoneManagedHeaderListResponseMessagesSource]
type zoneManagedHeaderListResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneManagedHeaderListResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneManagedHeaderListResponseMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// A result.
type ZoneManagedHeaderListResponseResult struct {
	// The list of Managed Request Transforms.
	ManagedRequestHeaders []ZoneManagedHeaderListResponseResultManagedRequestHeader `json:"managed_request_headers" api:"required"`
	// The list of Managed Response Transforms.
	ManagedResponseHeaders []ZoneManagedHeaderListResponseResultManagedResponseHeader `json:"managed_response_headers" api:"required"`
	JSON                   zoneManagedHeaderListResponseResultJSON                    `json:"-"`
}

// zoneManagedHeaderListResponseResultJSON contains the JSON metadata for the
// struct [ZoneManagedHeaderListResponseResult]
type zoneManagedHeaderListResponseResultJSON struct {
	ManagedRequestHeaders  apijson.Field
	ManagedResponseHeaders apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZoneManagedHeaderListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneManagedHeaderListResponseResultJSON) RawJSON() string {
	return r.raw
}

// A Managed Transform object.
type ZoneManagedHeaderListResponseResultManagedRequestHeader struct {
	// The human-readable identifier of the Managed Transform.
	ID string `json:"id" api:"required"`
	// Whether the Managed Transform is enabled.
	Enabled bool `json:"enabled" api:"required"`
	// Whether the Managed Transform conflicts with the currently-enabled Managed
	// Transforms.
	HasConflict bool `json:"has_conflict" api:"required"`
	// The Managed Transforms that this Managed Transform conflicts with.
	ConflictsWith []string                                                    `json:"conflicts_with"`
	JSON          zoneManagedHeaderListResponseResultManagedRequestHeaderJSON `json:"-"`
}

// zoneManagedHeaderListResponseResultManagedRequestHeaderJSON contains the JSON
// metadata for the struct
// [ZoneManagedHeaderListResponseResultManagedRequestHeader]
type zoneManagedHeaderListResponseResultManagedRequestHeaderJSON struct {
	ID            apijson.Field
	Enabled       apijson.Field
	HasConflict   apijson.Field
	ConflictsWith apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ZoneManagedHeaderListResponseResultManagedRequestHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneManagedHeaderListResponseResultManagedRequestHeaderJSON) RawJSON() string {
	return r.raw
}

// A Managed Transform object.
type ZoneManagedHeaderListResponseResultManagedResponseHeader struct {
	// The human-readable identifier of the Managed Transform.
	ID string `json:"id" api:"required"`
	// Whether the Managed Transform is enabled.
	Enabled bool `json:"enabled" api:"required"`
	// Whether the Managed Transform conflicts with the currently-enabled Managed
	// Transforms.
	HasConflict bool `json:"has_conflict" api:"required"`
	// The Managed Transforms that this Managed Transform conflicts with.
	ConflictsWith []string                                                     `json:"conflicts_with"`
	JSON          zoneManagedHeaderListResponseResultManagedResponseHeaderJSON `json:"-"`
}

// zoneManagedHeaderListResponseResultManagedResponseHeaderJSON contains the JSON
// metadata for the struct
// [ZoneManagedHeaderListResponseResultManagedResponseHeader]
type zoneManagedHeaderListResponseResultManagedResponseHeaderJSON struct {
	ID            apijson.Field
	Enabled       apijson.Field
	HasConflict   apijson.Field
	ConflictsWith apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ZoneManagedHeaderListResponseResultManagedResponseHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneManagedHeaderListResponseResultManagedResponseHeaderJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ZoneManagedHeaderListResponseSuccess bool

const (
	ZoneManagedHeaderListResponseSuccessTrue ZoneManagedHeaderListResponseSuccess = true
)

func (r ZoneManagedHeaderListResponseSuccess) IsKnown() bool {
	switch r {
	case ZoneManagedHeaderListResponseSuccessTrue:
		return true
	}
	return false
}

type ZoneManagedHeaderUpdateParams struct {
	// The list of Managed Request Transforms.
	ManagedRequestHeaders param.Field[[]ZoneManagedHeaderUpdateParamsManagedRequestHeader] `json:"managed_request_headers" api:"required"`
	// The list of Managed Response Transforms.
	ManagedResponseHeaders param.Field[[]ZoneManagedHeaderUpdateParamsManagedResponseHeader] `json:"managed_response_headers" api:"required"`
}

func (r ZoneManagedHeaderUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A Managed Transform object.
type ZoneManagedHeaderUpdateParamsManagedRequestHeader struct {
	// The human-readable identifier of the Managed Transform.
	ID param.Field[string] `json:"id" api:"required"`
	// Whether the Managed Transform is enabled.
	Enabled param.Field[bool] `json:"enabled" api:"required"`
}

func (r ZoneManagedHeaderUpdateParamsManagedRequestHeader) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A Managed Transform object.
type ZoneManagedHeaderUpdateParamsManagedResponseHeader struct {
	// The human-readable identifier of the Managed Transform.
	ID param.Field[string] `json:"id" api:"required"`
	// Whether the Managed Transform is enabled.
	Enabled param.Field[bool] `json:"enabled" api:"required"`
}

func (r ZoneManagedHeaderUpdateParamsManagedResponseHeader) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
