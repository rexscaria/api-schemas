// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/apiquery"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
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
//
// Deprecated: deprecated
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
//
// Deprecated: deprecated
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
//
// Deprecated: deprecated
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
//
// Deprecated: deprecated
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
//
// Deprecated: deprecated
func (r *ZoneFilterService) Delete(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *ZoneFilterDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/filters", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Deletes an existing filter.
//
// Deprecated: deprecated
func (r *ZoneFilterService) DeleteSingle(ctx context.Context, zoneID string, filterID string, opts ...option.RequestOption) (res *ZoneFilterDeleteSingleResponse, err error) {
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
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Updates an existing filter.
//
// Deprecated: deprecated
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
	Errors   []FirewallFilterCollectionError   `json:"errors,required"`
	Messages []FirewallFilterCollectionMessage `json:"messages,required"`
	Result   []FirewallFilter                  `json:"result,required,nullable"`
	// Defines whether the API call was successful.
	Success    FirewallFilterCollectionSuccess    `json:"success,required"`
	ResultInfo FirewallFilterCollectionResultInfo `json:"result_info"`
	JSON       firewallFilterCollectionJSON       `json:"-"`
}

// firewallFilterCollectionJSON contains the JSON metadata for the struct
// [FirewallFilterCollection]
type firewallFilterCollectionJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallFilterCollection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallFilterCollectionJSON) RawJSON() string {
	return r.raw
}

type FirewallFilterCollectionError struct {
	Code             int64                                `json:"code,required"`
	Message          string                               `json:"message,required"`
	DocumentationURL string                               `json:"documentation_url"`
	Source           FirewallFilterCollectionErrorsSource `json:"source"`
	JSON             firewallFilterCollectionErrorJSON    `json:"-"`
}

// firewallFilterCollectionErrorJSON contains the JSON metadata for the struct
// [FirewallFilterCollectionError]
type firewallFilterCollectionErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *FirewallFilterCollectionError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallFilterCollectionErrorJSON) RawJSON() string {
	return r.raw
}

type FirewallFilterCollectionErrorsSource struct {
	Pointer string                                   `json:"pointer"`
	JSON    firewallFilterCollectionErrorsSourceJSON `json:"-"`
}

// firewallFilterCollectionErrorsSourceJSON contains the JSON metadata for the
// struct [FirewallFilterCollectionErrorsSource]
type firewallFilterCollectionErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallFilterCollectionErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallFilterCollectionErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type FirewallFilterCollectionMessage struct {
	Code             int64                                  `json:"code,required"`
	Message          string                                 `json:"message,required"`
	DocumentationURL string                                 `json:"documentation_url"`
	Source           FirewallFilterCollectionMessagesSource `json:"source"`
	JSON             firewallFilterCollectionMessageJSON    `json:"-"`
}

// firewallFilterCollectionMessageJSON contains the JSON metadata for the struct
// [FirewallFilterCollectionMessage]
type firewallFilterCollectionMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *FirewallFilterCollectionMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallFilterCollectionMessageJSON) RawJSON() string {
	return r.raw
}

type FirewallFilterCollectionMessagesSource struct {
	Pointer string                                     `json:"pointer"`
	JSON    firewallFilterCollectionMessagesSourceJSON `json:"-"`
}

// firewallFilterCollectionMessagesSourceJSON contains the JSON metadata for the
// struct [FirewallFilterCollectionMessagesSource]
type firewallFilterCollectionMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallFilterCollectionMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallFilterCollectionMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Defines whether the API call was successful.
type FirewallFilterCollectionSuccess bool

const (
	FirewallFilterCollectionSuccessTrue FirewallFilterCollectionSuccess = true
)

func (r FirewallFilterCollectionSuccess) IsKnown() bool {
	switch r {
	case FirewallFilterCollectionSuccessTrue:
		return true
	}
	return false
}

type FirewallFilterCollectionResultInfo struct {
	// Defines the total number of results for the requested service.
	Count float64 `json:"count"`
	// Defines the current page within paginated list of results.
	Page float64 `json:"page"`
	// Defines the number of results per page of results.
	PerPage float64 `json:"per_page"`
	// Defines the total results available without any search parameters.
	TotalCount float64                                `json:"total_count"`
	JSON       firewallFilterCollectionResultInfoJSON `json:"-"`
}

// firewallFilterCollectionResultInfoJSON contains the JSON metadata for the struct
// [FirewallFilterCollectionResultInfo]
type firewallFilterCollectionResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallFilterCollectionResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallFilterCollectionResultInfoJSON) RawJSON() string {
	return r.raw
}

type FirewallFilterSingle struct {
	Errors   []FirewallFilterSingleError   `json:"errors,required"`
	Messages []FirewallFilterSingleMessage `json:"messages,required"`
	Result   FirewallFilter                `json:"result,required"`
	// Defines whether the API call was successful.
	Success FirewallFilterSingleSuccess `json:"success,required"`
	JSON    firewallFilterSingleJSON    `json:"-"`
}

// firewallFilterSingleJSON contains the JSON metadata for the struct
// [FirewallFilterSingle]
type firewallFilterSingleJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallFilterSingle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallFilterSingleJSON) RawJSON() string {
	return r.raw
}

type FirewallFilterSingleError struct {
	Code             int64                            `json:"code,required"`
	Message          string                           `json:"message,required"`
	DocumentationURL string                           `json:"documentation_url"`
	Source           FirewallFilterSingleErrorsSource `json:"source"`
	JSON             firewallFilterSingleErrorJSON    `json:"-"`
}

// firewallFilterSingleErrorJSON contains the JSON metadata for the struct
// [FirewallFilterSingleError]
type firewallFilterSingleErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *FirewallFilterSingleError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallFilterSingleErrorJSON) RawJSON() string {
	return r.raw
}

type FirewallFilterSingleErrorsSource struct {
	Pointer string                               `json:"pointer"`
	JSON    firewallFilterSingleErrorsSourceJSON `json:"-"`
}

// firewallFilterSingleErrorsSourceJSON contains the JSON metadata for the struct
// [FirewallFilterSingleErrorsSource]
type firewallFilterSingleErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallFilterSingleErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallFilterSingleErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type FirewallFilterSingleMessage struct {
	Code             int64                              `json:"code,required"`
	Message          string                             `json:"message,required"`
	DocumentationURL string                             `json:"documentation_url"`
	Source           FirewallFilterSingleMessagesSource `json:"source"`
	JSON             firewallFilterSingleMessageJSON    `json:"-"`
}

// firewallFilterSingleMessageJSON contains the JSON metadata for the struct
// [FirewallFilterSingleMessage]
type firewallFilterSingleMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *FirewallFilterSingleMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallFilterSingleMessageJSON) RawJSON() string {
	return r.raw
}

type FirewallFilterSingleMessagesSource struct {
	Pointer string                                 `json:"pointer"`
	JSON    firewallFilterSingleMessagesSourceJSON `json:"-"`
}

// firewallFilterSingleMessagesSourceJSON contains the JSON metadata for the struct
// [FirewallFilterSingleMessagesSource]
type firewallFilterSingleMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallFilterSingleMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallFilterSingleMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Defines whether the API call was successful.
type FirewallFilterSingleSuccess bool

const (
	FirewallFilterSingleSuccessTrue FirewallFilterSingleSuccess = true
)

func (r FirewallFilterSingleSuccess) IsKnown() bool {
	switch r {
	case FirewallFilterSingleSuccessTrue:
		return true
	}
	return false
}

type ZoneFilterDeleteResponse struct {
	Errors   []ZoneFilterDeleteResponseError   `json:"errors,required"`
	Messages []ZoneFilterDeleteResponseMessage `json:"messages,required"`
	Result   []FirewallFilter                  `json:"result,required,nullable"`
	// Defines whether the API call was successful.
	Success    ZoneFilterDeleteResponseSuccess    `json:"success,required"`
	ResultInfo ZoneFilterDeleteResponseResultInfo `json:"result_info"`
	JSON       zoneFilterDeleteResponseJSON       `json:"-"`
}

// zoneFilterDeleteResponseJSON contains the JSON metadata for the struct
// [ZoneFilterDeleteResponse]
type zoneFilterDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFilterDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneFilterDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneFilterDeleteResponseError struct {
	Code             int64                                `json:"code,required"`
	Message          string                               `json:"message,required"`
	DocumentationURL string                               `json:"documentation_url"`
	Source           ZoneFilterDeleteResponseErrorsSource `json:"source"`
	JSON             zoneFilterDeleteResponseErrorJSON    `json:"-"`
}

// zoneFilterDeleteResponseErrorJSON contains the JSON metadata for the struct
// [ZoneFilterDeleteResponseError]
type zoneFilterDeleteResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZoneFilterDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneFilterDeleteResponseErrorJSON) RawJSON() string {
	return r.raw
}

type ZoneFilterDeleteResponseErrorsSource struct {
	Pointer string                                   `json:"pointer"`
	JSON    zoneFilterDeleteResponseErrorsSourceJSON `json:"-"`
}

// zoneFilterDeleteResponseErrorsSourceJSON contains the JSON metadata for the
// struct [ZoneFilterDeleteResponseErrorsSource]
type zoneFilterDeleteResponseErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFilterDeleteResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneFilterDeleteResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type ZoneFilterDeleteResponseMessage struct {
	Code             int64                                  `json:"code,required"`
	Message          string                                 `json:"message,required"`
	DocumentationURL string                                 `json:"documentation_url"`
	Source           ZoneFilterDeleteResponseMessagesSource `json:"source"`
	JSON             zoneFilterDeleteResponseMessageJSON    `json:"-"`
}

// zoneFilterDeleteResponseMessageJSON contains the JSON metadata for the struct
// [ZoneFilterDeleteResponseMessage]
type zoneFilterDeleteResponseMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZoneFilterDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneFilterDeleteResponseMessageJSON) RawJSON() string {
	return r.raw
}

type ZoneFilterDeleteResponseMessagesSource struct {
	Pointer string                                     `json:"pointer"`
	JSON    zoneFilterDeleteResponseMessagesSourceJSON `json:"-"`
}

// zoneFilterDeleteResponseMessagesSourceJSON contains the JSON metadata for the
// struct [ZoneFilterDeleteResponseMessagesSource]
type zoneFilterDeleteResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFilterDeleteResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneFilterDeleteResponseMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Defines whether the API call was successful.
type ZoneFilterDeleteResponseSuccess bool

const (
	ZoneFilterDeleteResponseSuccessTrue ZoneFilterDeleteResponseSuccess = true
)

func (r ZoneFilterDeleteResponseSuccess) IsKnown() bool {
	switch r {
	case ZoneFilterDeleteResponseSuccessTrue:
		return true
	}
	return false
}

type ZoneFilterDeleteResponseResultInfo struct {
	// Defines the total number of results for the requested service.
	Count float64 `json:"count"`
	// Defines the current page within paginated list of results.
	Page float64 `json:"page"`
	// Defines the number of results per page of results.
	PerPage float64 `json:"per_page"`
	// Defines the total results available without any search parameters.
	TotalCount float64                                `json:"total_count"`
	JSON       zoneFilterDeleteResponseResultInfoJSON `json:"-"`
}

// zoneFilterDeleteResponseResultInfoJSON contains the JSON metadata for the struct
// [ZoneFilterDeleteResponseResultInfo]
type zoneFilterDeleteResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFilterDeleteResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneFilterDeleteResponseResultInfoJSON) RawJSON() string {
	return r.raw
}

type ZoneFilterDeleteSingleResponse struct {
	Errors   []ZoneFilterDeleteSingleResponseError   `json:"errors,required"`
	Messages []ZoneFilterDeleteSingleResponseMessage `json:"messages,required"`
	Result   FirewallFilter                          `json:"result,required"`
	// Defines whether the API call was successful.
	Success ZoneFilterDeleteSingleResponseSuccess `json:"success,required"`
	JSON    zoneFilterDeleteSingleResponseJSON    `json:"-"`
}

// zoneFilterDeleteSingleResponseJSON contains the JSON metadata for the struct
// [ZoneFilterDeleteSingleResponse]
type zoneFilterDeleteSingleResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFilterDeleteSingleResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneFilterDeleteSingleResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneFilterDeleteSingleResponseError struct {
	Code             int64                                      `json:"code,required"`
	Message          string                                     `json:"message,required"`
	DocumentationURL string                                     `json:"documentation_url"`
	Source           ZoneFilterDeleteSingleResponseErrorsSource `json:"source"`
	JSON             zoneFilterDeleteSingleResponseErrorJSON    `json:"-"`
}

// zoneFilterDeleteSingleResponseErrorJSON contains the JSON metadata for the
// struct [ZoneFilterDeleteSingleResponseError]
type zoneFilterDeleteSingleResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZoneFilterDeleteSingleResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneFilterDeleteSingleResponseErrorJSON) RawJSON() string {
	return r.raw
}

type ZoneFilterDeleteSingleResponseErrorsSource struct {
	Pointer string                                         `json:"pointer"`
	JSON    zoneFilterDeleteSingleResponseErrorsSourceJSON `json:"-"`
}

// zoneFilterDeleteSingleResponseErrorsSourceJSON contains the JSON metadata for
// the struct [ZoneFilterDeleteSingleResponseErrorsSource]
type zoneFilterDeleteSingleResponseErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFilterDeleteSingleResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneFilterDeleteSingleResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type ZoneFilterDeleteSingleResponseMessage struct {
	Code             int64                                        `json:"code,required"`
	Message          string                                       `json:"message,required"`
	DocumentationURL string                                       `json:"documentation_url"`
	Source           ZoneFilterDeleteSingleResponseMessagesSource `json:"source"`
	JSON             zoneFilterDeleteSingleResponseMessageJSON    `json:"-"`
}

// zoneFilterDeleteSingleResponseMessageJSON contains the JSON metadata for the
// struct [ZoneFilterDeleteSingleResponseMessage]
type zoneFilterDeleteSingleResponseMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZoneFilterDeleteSingleResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneFilterDeleteSingleResponseMessageJSON) RawJSON() string {
	return r.raw
}

type ZoneFilterDeleteSingleResponseMessagesSource struct {
	Pointer string                                           `json:"pointer"`
	JSON    zoneFilterDeleteSingleResponseMessagesSourceJSON `json:"-"`
}

// zoneFilterDeleteSingleResponseMessagesSourceJSON contains the JSON metadata for
// the struct [ZoneFilterDeleteSingleResponseMessagesSource]
type zoneFilterDeleteSingleResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFilterDeleteSingleResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneFilterDeleteSingleResponseMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Defines whether the API call was successful.
type ZoneFilterDeleteSingleResponseSuccess bool

const (
	ZoneFilterDeleteSingleResponseSuccessTrue ZoneFilterDeleteSingleResponseSuccess = true
)

func (r ZoneFilterDeleteSingleResponseSuccess) IsKnown() bool {
	switch r {
	case ZoneFilterDeleteSingleResponseSuccessTrue:
		return true
	}
	return false
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

type ZoneFilterUpdateSingleParams struct {
	Body interface{} `json:"body,required"`
}

func (r ZoneFilterUpdateSingleParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}
