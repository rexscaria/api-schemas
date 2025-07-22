// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/apiquery"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
	"github.com/rexscaria/api-schemas/shared"
	"github.com/tidwall/gjson"
)

// ZoneHoldService contains methods and other services that help with interacting
// with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneHoldService] method instead.
type ZoneHoldService struct {
	Options []option.RequestOption
}

// NewZoneHoldService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewZoneHoldService(opts ...option.RequestOption) (r *ZoneHoldService) {
	r = &ZoneHoldService{}
	r.Options = opts
	return
}

// Enforce a zone hold on the zone, blocking the creation and activation of zones
// with this zone's hostname.
func (r *ZoneHoldService) New(ctx context.Context, zoneID string, body ZoneHoldNewParams, opts ...option.RequestOption) (res *ZoneHoldNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/hold", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve whether the zone is subject to a zone hold, and metadata about the
// hold.
func (r *ZoneHoldService) Get(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *ZoneHoldGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/hold", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update the `hold_after` and/or `include_subdomains` values on an existing zone
// hold. The hold is enabled if the `hold_after` date-time value is in the past.
func (r *ZoneHoldService) Update(ctx context.Context, zoneID string, body ZoneHoldUpdateParams, opts ...option.RequestOption) (res *ZoneHoldUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/hold", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Stop enforcement of a zone hold on the zone, permanently or temporarily,
// allowing the creation and activation of zones with this zone's hostname.
func (r *ZoneHoldService) Remove(ctx context.Context, zoneID string, body ZoneHoldRemoveParams, opts ...option.RequestOption) (res *ZoneHoldRemoveResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/hold", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

type APIResponseSingleZonesSchemas struct {
	Errors   []MessagesZonesItem                      `json:"errors,required"`
	Messages []MessagesZonesItem                      `json:"messages,required"`
	Result   APIResponseSingleZonesSchemasResultUnion `json:"result,required"`
	// Whether the API call was successful
	Success APIResponseSingleZonesSchemasSuccess `json:"success,required"`
	JSON    apiResponseSingleZonesSchemasJSON    `json:"-"`
}

// apiResponseSingleZonesSchemasJSON contains the JSON metadata for the struct
// [APIResponseSingleZonesSchemas]
type apiResponseSingleZonesSchemasJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIResponseSingleZonesSchemas) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiResponseSingleZonesSchemasJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [APIResponseSingleZonesSchemasResultArray] or
// [shared.UnionString].
type APIResponseSingleZonesSchemasResultUnion interface {
	ImplementsAPIResponseSingleZonesSchemasResultUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*APIResponseSingleZonesSchemasResultUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(APIResponseSingleZonesSchemasResultArray{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type APIResponseSingleZonesSchemasResultArray []interface{}

func (r APIResponseSingleZonesSchemasResultArray) ImplementsAPIResponseSingleZonesSchemasResultUnion() {
}

// Whether the API call was successful
type APIResponseSingleZonesSchemasSuccess bool

const (
	APIResponseSingleZonesSchemasSuccessTrue APIResponseSingleZonesSchemasSuccess = true
)

func (r APIResponseSingleZonesSchemasSuccess) IsKnown() bool {
	switch r {
	case APIResponseSingleZonesSchemasSuccessTrue:
		return true
	}
	return false
}

type ZoneHoldNewResponse struct {
	Result ZoneHoldNewResponseResult `json:"result"`
	JSON   zoneHoldNewResponseJSON   `json:"-"`
	APIResponseSingleZonesSchemas
}

// zoneHoldNewResponseJSON contains the JSON metadata for the struct
// [ZoneHoldNewResponse]
type zoneHoldNewResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneHoldNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneHoldNewResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneHoldNewResponseResult struct {
	Hold              bool                          `json:"hold"`
	HoldAfter         string                        `json:"hold_after"`
	IncludeSubdomains string                        `json:"include_subdomains"`
	JSON              zoneHoldNewResponseResultJSON `json:"-"`
}

// zoneHoldNewResponseResultJSON contains the JSON metadata for the struct
// [ZoneHoldNewResponseResult]
type zoneHoldNewResponseResultJSON struct {
	Hold              apijson.Field
	HoldAfter         apijson.Field
	IncludeSubdomains apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ZoneHoldNewResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneHoldNewResponseResultJSON) RawJSON() string {
	return r.raw
}

type ZoneHoldGetResponse struct {
	Result ZoneHoldGetResponseResult `json:"result"`
	JSON   zoneHoldGetResponseJSON   `json:"-"`
	APIResponseSingleZonesSchemas
}

// zoneHoldGetResponseJSON contains the JSON metadata for the struct
// [ZoneHoldGetResponse]
type zoneHoldGetResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneHoldGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneHoldGetResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneHoldGetResponseResult struct {
	Hold              bool                          `json:"hold"`
	HoldAfter         string                        `json:"hold_after"`
	IncludeSubdomains string                        `json:"include_subdomains"`
	JSON              zoneHoldGetResponseResultJSON `json:"-"`
}

// zoneHoldGetResponseResultJSON contains the JSON metadata for the struct
// [ZoneHoldGetResponseResult]
type zoneHoldGetResponseResultJSON struct {
	Hold              apijson.Field
	HoldAfter         apijson.Field
	IncludeSubdomains apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ZoneHoldGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneHoldGetResponseResultJSON) RawJSON() string {
	return r.raw
}

type ZoneHoldUpdateResponse struct {
	Result ZoneHoldUpdateResponseResult `json:"result"`
	JSON   zoneHoldUpdateResponseJSON   `json:"-"`
	APIResponseSingleZonesSchemas
}

// zoneHoldUpdateResponseJSON contains the JSON metadata for the struct
// [ZoneHoldUpdateResponse]
type zoneHoldUpdateResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneHoldUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneHoldUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneHoldUpdateResponseResult struct {
	Hold              bool                             `json:"hold"`
	HoldAfter         string                           `json:"hold_after"`
	IncludeSubdomains string                           `json:"include_subdomains"`
	JSON              zoneHoldUpdateResponseResultJSON `json:"-"`
}

// zoneHoldUpdateResponseResultJSON contains the JSON metadata for the struct
// [ZoneHoldUpdateResponseResult]
type zoneHoldUpdateResponseResultJSON struct {
	Hold              apijson.Field
	HoldAfter         apijson.Field
	IncludeSubdomains apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ZoneHoldUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneHoldUpdateResponseResultJSON) RawJSON() string {
	return r.raw
}

type ZoneHoldRemoveResponse struct {
	Result ZoneHoldRemoveResponseResult `json:"result"`
	JSON   zoneHoldRemoveResponseJSON   `json:"-"`
	APIResponseSingleZonesSchemas
}

// zoneHoldRemoveResponseJSON contains the JSON metadata for the struct
// [ZoneHoldRemoveResponse]
type zoneHoldRemoveResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneHoldRemoveResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneHoldRemoveResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneHoldRemoveResponseResult struct {
	Hold              bool                             `json:"hold"`
	HoldAfter         string                           `json:"hold_after"`
	IncludeSubdomains string                           `json:"include_subdomains"`
	JSON              zoneHoldRemoveResponseResultJSON `json:"-"`
}

// zoneHoldRemoveResponseResultJSON contains the JSON metadata for the struct
// [ZoneHoldRemoveResponseResult]
type zoneHoldRemoveResponseResultJSON struct {
	Hold              apijson.Field
	HoldAfter         apijson.Field
	IncludeSubdomains apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ZoneHoldRemoveResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneHoldRemoveResponseResultJSON) RawJSON() string {
	return r.raw
}

type ZoneHoldNewParams struct {
	// If provided, the zone hold will extend to block any subdomain of the given zone,
	// as well as SSL4SaaS Custom Hostnames. For example, a zone hold on a zone with
	// the hostname 'example.com' and include_subdomains=true will block 'example.com',
	// 'staging.example.com', 'api.staging.example.com', etc.
	IncludeSubdomains param.Field[bool] `query:"include_subdomains"`
}

// URLQuery serializes [ZoneHoldNewParams]'s query parameters as `url.Values`.
func (r ZoneHoldNewParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ZoneHoldUpdateParams struct {
	// If `hold_after` is provided and future-dated, the hold will be temporarily
	// disabled, then automatically re-enabled by the system at the time specified in
	// this RFC3339-formatted timestamp. A past-dated `hold_after` value will have no
	// effect on an existing, enabled hold. Providing an empty string will set its
	// value to the current time.
	HoldAfter param.Field[string] `json:"hold_after"`
	// If `true`, the zone hold will extend to block any subdomain of the given zone,
	// as well as SSL4SaaS Custom Hostnames. For example, a zone hold on a zone with
	// the hostname 'example.com' and include_subdomains=true will block 'example.com',
	// 'staging.example.com', 'api.staging.example.com', etc.
	IncludeSubdomains param.Field[bool] `json:"include_subdomains"`
}

func (r ZoneHoldUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneHoldRemoveParams struct {
	// If `hold_after` is provided, the hold will be temporarily disabled, then
	// automatically re-enabled by the system at the time specified in this
	// RFC3339-formatted timestamp. Otherwise, the hold will be disabled indefinitely.
	HoldAfter param.Field[string] `query:"hold_after"`
}

// URLQuery serializes [ZoneHoldRemoveParams]'s query parameters as `url.Values`.
func (r ZoneHoldRemoveParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
