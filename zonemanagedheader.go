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
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/managed_headers", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Fetches a list of all Managed Transforms.
func (r *ZoneManagedHeaderService) List(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *ZoneManagedHeaderListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/managed_headers", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Disables all Managed Transforms.
func (r *ZoneManagedHeaderService) Delete(ctx context.Context, zoneID string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/managed_headers", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type ZoneManagedHeaderUpdateResponse struct {
	// A list of error messages.
	Errors interface{} `json:"errors,required"`
	// A list of warning messages.
	Messages []ZoneManagedHeaderUpdateResponseMessage `json:"messages,required"`
	// A result.
	Result interface{} `json:"result,required"`
	// Whether the API call was successful.
	Success ZoneManagedHeaderUpdateResponseSuccess `json:"success,required"`
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
type ZoneManagedHeaderUpdateResponseMessage struct {
	// A text description of this message.
	Message string `json:"message,required"`
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
	Pointer string                                            `json:"pointer,required"`
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

type ZoneManagedHeaderListResponse struct {
	// A list of error messages.
	Errors interface{} `json:"errors,required"`
	// A list of warning messages.
	Messages []ZoneManagedHeaderListResponseMessage `json:"messages,required"`
	// A result.
	Result interface{} `json:"result,required"`
	// Whether the API call was successful.
	Success ZoneManagedHeaderListResponseSuccess `json:"success,required"`
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
type ZoneManagedHeaderListResponseMessage struct {
	// A text description of this message.
	Message string `json:"message,required"`
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
	Pointer string                                          `json:"pointer,required"`
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
	ManagedRequestHeaders param.Field[[]ZoneManagedHeaderUpdateParamsManagedRequestHeader] `json:"managed_request_headers,required"`
	// The list of Managed Response Transforms.
	ManagedResponseHeaders param.Field[[]ZoneManagedHeaderUpdateParamsManagedResponseHeader] `json:"managed_response_headers,required"`
}

func (r ZoneManagedHeaderUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneManagedHeaderUpdateParamsManagedRequestHeader struct {
	// The human-readable identifier of the Managed Transform.
	ID param.Field[interface{}] `json:"id,required"`
	// Whether the Managed Transform is enabled.
	Enabled param.Field[bool] `json:"enabled,required"`
}

func (r ZoneManagedHeaderUpdateParamsManagedRequestHeader) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneManagedHeaderUpdateParamsManagedResponseHeader struct {
	// The human-readable identifier of the Managed Transform.
	ID param.Field[interface{}] `json:"id,required"`
	// Whether the Managed Transform is enabled.
	Enabled param.Field[bool] `json:"enabled,required"`
}

func (r ZoneManagedHeaderUpdateParamsManagedResponseHeader) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
