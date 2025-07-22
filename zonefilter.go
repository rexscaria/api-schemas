// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/stainless-sdks/cf-rex-go/internal/apijson"
	"github.com/stainless-sdks/cf-rex-go/internal/apiquery"
	"github.com/stainless-sdks/cf-rex-go/internal/param"
	"github.com/stainless-sdks/cf-rex-go/internal/requestconfig"
	"github.com/stainless-sdks/cf-rex-go/option"
)

// ZoneFilterService contains methods and other services that help with interacting
// with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneFilterService] method instead.
type ZoneFilterService struct {
	Options []option.RequestOption
}

// NewZoneFilterService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewZoneFilterService(opts ...option.RequestOption) (r *ZoneFilterService) {
	r = &ZoneFilterService{}
	r.Options = opts
	return
}

// Creates one or more filters.
func (r *ZoneFilterService) New(ctx context.Context, zoneID string, body ZoneFilterNewParams, opts ...option.RequestOption) (res *FirewallFilterCollection, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/filters", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Fetches the details of a filter.
func (r *ZoneFilterService) Get(ctx context.Context, zoneID string, filterID string, opts ...option.RequestOption) (res *FirewallFilterSingle, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if filterID == "" {
		err = errors.New("missing required filter_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/filters/%s", zoneID, filterID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates one or more existing filters.
func (r *ZoneFilterService) Update(ctx context.Context, zoneID string, body ZoneFilterUpdateParams, opts ...option.RequestOption) (res *FirewallFilterCollection, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/filters", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Fetches filters in a zone. You can filter the results using several optional
// parameters.
func (r *ZoneFilterService) List(ctx context.Context, zoneID string, query ZoneFilterListParams, opts ...option.RequestOption) (res *FirewallFilterCollection, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/filters", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Deletes one or more existing filters.
func (r *ZoneFilterService) Delete(ctx context.Context, zoneID string, body ZoneFilterDeleteParams, opts ...option.RequestOption) (res *ZoneFilterDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/filters", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

// Deletes an existing filter.
func (r *ZoneFilterService) DeleteSingle(ctx context.Context, zoneID string, filterID string, body ZoneFilterDeleteSingleParams, opts ...option.RequestOption) (res *ZoneFilterDeleteSingleResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if filterID == "" {
		err = errors.New("missing required filter_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/filters/%s", zoneID, filterID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

// Updates an existing filter.
func (r *ZoneFilterService) UpdateSingle(ctx context.Context, zoneID string, filterID string, body ZoneFilterUpdateSingleParams, opts ...option.RequestOption) (res *FirewallFilterSingle, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if filterID == "" {
		err = errors.New("missing required filter_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/filters/%s", zoneID, filterID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type FirewallFilterCollection struct {
	Result []FirewallFilterCollectionResult `json:"result"`
	JSON   firewallFilterCollectionJSON     `json:"-"`
	FirewallAPIResponseCollection
}

// firewallFilterCollectionJSON contains the JSON metadata for the struct
// [FirewallFilterCollection]
type firewallFilterCollectionJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallFilterCollection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallFilterCollectionJSON) RawJSON() string {
	return r.raw
}

type FirewallFilterCollectionResult struct {
	JSON firewallFilterCollectionResultJSON `json:"-"`
	FirewallFilter
}

// firewallFilterCollectionResultJSON contains the JSON metadata for the struct
// [FirewallFilterCollectionResult]
type firewallFilterCollectionResultJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallFilterCollectionResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallFilterCollectionResultJSON) RawJSON() string {
	return r.raw
}

type FirewallFilterSingle struct {
	Result FirewallFilter           `json:"result,required"`
	JSON   firewallFilterSingleJSON `json:"-"`
	FirewallAPIResponseSingle
}

// firewallFilterSingleJSON contains the JSON metadata for the struct
// [FirewallFilterSingle]
type firewallFilterSingleJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallFilterSingle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallFilterSingleJSON) RawJSON() string {
	return r.raw
}

type ZoneFilterDeleteResponse struct {
	Result []ZoneFilterDeleteResponseResult `json:"result"`
	JSON   zoneFilterDeleteResponseJSON     `json:"-"`
	FirewallAPIResponseCollection
}

// zoneFilterDeleteResponseJSON contains the JSON metadata for the struct
// [ZoneFilterDeleteResponse]
type zoneFilterDeleteResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFilterDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneFilterDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneFilterDeleteResponseResult struct {
	JSON zoneFilterDeleteResponseResultJSON `json:"-"`
	FirewallFilter
}

// zoneFilterDeleteResponseResultJSON contains the JSON metadata for the struct
// [ZoneFilterDeleteResponseResult]
type zoneFilterDeleteResponseResultJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFilterDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneFilterDeleteResponseResultJSON) RawJSON() string {
	return r.raw
}

type ZoneFilterDeleteSingleResponse struct {
	Result ZoneFilterDeleteSingleResponseResult `json:"result,required"`
	JSON   zoneFilterDeleteSingleResponseJSON   `json:"-"`
	FirewallAPIResponseSingle
}

// zoneFilterDeleteSingleResponseJSON contains the JSON metadata for the struct
// [ZoneFilterDeleteSingleResponse]
type zoneFilterDeleteSingleResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFilterDeleteSingleResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneFilterDeleteSingleResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneFilterDeleteSingleResponseResult struct {
	JSON zoneFilterDeleteSingleResponseResultJSON `json:"-"`
	FirewallFilter
}

// zoneFilterDeleteSingleResponseResultJSON contains the JSON metadata for the
// struct [ZoneFilterDeleteSingleResponseResult]
type zoneFilterDeleteSingleResponseResultJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFilterDeleteSingleResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneFilterDeleteSingleResponseResultJSON) RawJSON() string {
	return r.raw
}

type ZoneFilterNewParams struct {
	// The filter expression. For more information, refer to
	// [Expressions](https://developers.cloudflare.com/ruleset-engine/rules-language/expressions/).
	Expression param.Field[string] `json:"expression,required"`
}

func (r ZoneFilterNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneFilterUpdateParams struct {
}

func (r ZoneFilterUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneFilterListParams struct {
	// The unique identifier of the filter.
	ID param.Field[string] `query:"id"`
	// A case-insensitive string to find in the description.
	Description param.Field[string] `query:"description"`
	// A case-insensitive string to find in the expression.
	Expression param.Field[string] `query:"expression"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// When true, indicates that the filter is currently paused.
	Paused param.Field[bool] `query:"paused"`
	// Number of filters per page.
	PerPage param.Field[float64] `query:"per_page"`
	// The filter ref (a short reference tag) to search for. Must be an exact match.
	Ref param.Field[string] `query:"ref"`
}

// URLQuery serializes [ZoneFilterListParams]'s query parameters as `url.Values`.
func (r ZoneFilterListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ZoneFilterDeleteParams struct {
}

func (r ZoneFilterDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneFilterDeleteSingleParams struct {
	Body interface{} `json:"body,required"`
}

func (r ZoneFilterDeleteSingleParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type ZoneFilterUpdateSingleParams struct {
	Body interface{} `json:"body,required"`
}

func (r ZoneFilterUpdateSingleParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}
