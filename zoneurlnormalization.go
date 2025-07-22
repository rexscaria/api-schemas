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

// ZoneURLNormalizationService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneURLNormalizationService] method instead.
type ZoneURLNormalizationService struct {
	Options []option.RequestOption
}

// NewZoneURLNormalizationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneURLNormalizationService(opts ...option.RequestOption) (r *ZoneURLNormalizationService) {
	r = &ZoneURLNormalizationService{}
	r.Options = opts
	return
}

// Fetches the current URL Normalization settings.
func (r *ZoneURLNormalizationService) Get(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *ZoneURLNormalizationGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/url_normalization", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates the URL Normalization settings.
func (r *ZoneURLNormalizationService) Update(ctx context.Context, zoneID string, body ZoneURLNormalizationUpdateParams, opts ...option.RequestOption) (res *ZoneURLNormalizationUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/url_normalization", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Deletes the URL Normalization settings.
func (r *ZoneURLNormalizationService) Delete(ctx context.Context, zoneID string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/url_normalization", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type ZoneURLNormalizationGetResponse struct {
	// A list of error messages.
	Errors interface{} `json:"errors,required"`
	// A list of warning messages.
	Messages []ZoneURLNormalizationGetResponseMessage `json:"messages,required"`
	// A result.
	Result interface{} `json:"result,required"`
	// Whether the API call was successful.
	Success ZoneURLNormalizationGetResponseSuccess `json:"success,required"`
	JSON    zoneURLNormalizationGetResponseJSON    `json:"-"`
}

// zoneURLNormalizationGetResponseJSON contains the JSON metadata for the struct
// [ZoneURLNormalizationGetResponse]
type zoneURLNormalizationGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneURLNormalizationGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneURLNormalizationGetResponseJSON) RawJSON() string {
	return r.raw
}

// A message.
type ZoneURLNormalizationGetResponseMessage struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source ZoneURLNormalizationGetResponseMessagesSource `json:"source"`
	JSON   zoneURLNormalizationGetResponseMessageJSON    `json:"-"`
}

// zoneURLNormalizationGetResponseMessageJSON contains the JSON metadata for the
// struct [ZoneURLNormalizationGetResponseMessage]
type zoneURLNormalizationGetResponseMessageJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneURLNormalizationGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneURLNormalizationGetResponseMessageJSON) RawJSON() string {
	return r.raw
}

// The source of this message.
type ZoneURLNormalizationGetResponseMessagesSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                            `json:"pointer,required"`
	JSON    zoneURLNormalizationGetResponseMessagesSourceJSON `json:"-"`
}

// zoneURLNormalizationGetResponseMessagesSourceJSON contains the JSON metadata for
// the struct [ZoneURLNormalizationGetResponseMessagesSource]
type zoneURLNormalizationGetResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneURLNormalizationGetResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneURLNormalizationGetResponseMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ZoneURLNormalizationGetResponseSuccess bool

const (
	ZoneURLNormalizationGetResponseSuccessTrue ZoneURLNormalizationGetResponseSuccess = true
)

func (r ZoneURLNormalizationGetResponseSuccess) IsKnown() bool {
	switch r {
	case ZoneURLNormalizationGetResponseSuccessTrue:
		return true
	}
	return false
}

type ZoneURLNormalizationUpdateResponse struct {
	// A list of error messages.
	Errors interface{} `json:"errors,required"`
	// A list of warning messages.
	Messages []ZoneURLNormalizationUpdateResponseMessage `json:"messages,required"`
	// A result.
	Result interface{} `json:"result,required"`
	// Whether the API call was successful.
	Success ZoneURLNormalizationUpdateResponseSuccess `json:"success,required"`
	JSON    zoneURLNormalizationUpdateResponseJSON    `json:"-"`
}

// zoneURLNormalizationUpdateResponseJSON contains the JSON metadata for the struct
// [ZoneURLNormalizationUpdateResponse]
type zoneURLNormalizationUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneURLNormalizationUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneURLNormalizationUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// A message.
type ZoneURLNormalizationUpdateResponseMessage struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source ZoneURLNormalizationUpdateResponseMessagesSource `json:"source"`
	JSON   zoneURLNormalizationUpdateResponseMessageJSON    `json:"-"`
}

// zoneURLNormalizationUpdateResponseMessageJSON contains the JSON metadata for the
// struct [ZoneURLNormalizationUpdateResponseMessage]
type zoneURLNormalizationUpdateResponseMessageJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneURLNormalizationUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneURLNormalizationUpdateResponseMessageJSON) RawJSON() string {
	return r.raw
}

// The source of this message.
type ZoneURLNormalizationUpdateResponseMessagesSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                               `json:"pointer,required"`
	JSON    zoneURLNormalizationUpdateResponseMessagesSourceJSON `json:"-"`
}

// zoneURLNormalizationUpdateResponseMessagesSourceJSON contains the JSON metadata
// for the struct [ZoneURLNormalizationUpdateResponseMessagesSource]
type zoneURLNormalizationUpdateResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneURLNormalizationUpdateResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneURLNormalizationUpdateResponseMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ZoneURLNormalizationUpdateResponseSuccess bool

const (
	ZoneURLNormalizationUpdateResponseSuccessTrue ZoneURLNormalizationUpdateResponseSuccess = true
)

func (r ZoneURLNormalizationUpdateResponseSuccess) IsKnown() bool {
	switch r {
	case ZoneURLNormalizationUpdateResponseSuccessTrue:
		return true
	}
	return false
}

type ZoneURLNormalizationUpdateParams struct {
	// The scope of the URL normalization.
	Scope param.Field[ZoneURLNormalizationUpdateParamsScope] `json:"scope,required"`
	// The type of URL normalization performed by Cloudflare.
	Type param.Field[ZoneURLNormalizationUpdateParamsType] `json:"type,required"`
}

func (r ZoneURLNormalizationUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The scope of the URL normalization.
type ZoneURLNormalizationUpdateParamsScope string

const (
	ZoneURLNormalizationUpdateParamsScopeIncoming ZoneURLNormalizationUpdateParamsScope = "incoming"
	ZoneURLNormalizationUpdateParamsScopeBoth     ZoneURLNormalizationUpdateParamsScope = "both"
)

func (r ZoneURLNormalizationUpdateParamsScope) IsKnown() bool {
	switch r {
	case ZoneURLNormalizationUpdateParamsScopeIncoming, ZoneURLNormalizationUpdateParamsScopeBoth:
		return true
	}
	return false
}

// The type of URL normalization performed by Cloudflare.
type ZoneURLNormalizationUpdateParamsType string

const (
	ZoneURLNormalizationUpdateParamsTypeCloudflare ZoneURLNormalizationUpdateParamsType = "cloudflare"
	ZoneURLNormalizationUpdateParamsTypeRfc3986    ZoneURLNormalizationUpdateParamsType = "rfc3986"
)

func (r ZoneURLNormalizationUpdateParamsType) IsKnown() bool {
	switch r {
	case ZoneURLNormalizationUpdateParamsTypeCloudflare, ZoneURLNormalizationUpdateParamsTypeRfc3986:
		return true
	}
	return false
}
