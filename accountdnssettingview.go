// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/apiquery"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
)

// AccountDNSSettingViewService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountDNSSettingViewService] method instead.
type AccountDNSSettingViewService struct {
	Options []option.RequestOption
}

// NewAccountDNSSettingViewService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountDNSSettingViewService(opts ...option.RequestOption) (r *AccountDNSSettingViewService) {
	r = &AccountDNSSettingViewService{}
	r.Options = opts
	return
}

// Create Internal DNS View for an account
func (r *AccountDNSSettingViewService) New(ctx context.Context, accountID string, body AccountDNSSettingViewNewParams, opts ...option.RequestOption) (res *DNSViewResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dns_settings/views", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get DNS Internal View
func (r *AccountDNSSettingViewService) Get(ctx context.Context, accountID string, viewID string, opts ...option.RequestOption) (res *DNSViewResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if viewID == "" {
		err = errors.New("missing required view_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dns_settings/views/%s", accountID, viewID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update an existing Internal DNS View
func (r *AccountDNSSettingViewService) Update(ctx context.Context, accountID string, viewID string, body AccountDNSSettingViewUpdateParams, opts ...option.RequestOption) (res *DNSViewResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if viewID == "" {
		err = errors.New("missing required view_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dns_settings/views/%s", accountID, viewID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List DNS Internal Views for an Account
func (r *AccountDNSSettingViewService) List(ctx context.Context, accountID string, query AccountDNSSettingViewListParams, opts ...option.RequestOption) (res *AccountDNSSettingViewListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dns_settings/views", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete an existing Internal DNS View
func (r *AccountDNSSettingViewService) Delete(ctx context.Context, accountID string, viewID string, opts ...option.RequestOption) (res *AccountDNSSettingViewDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if viewID == "" {
		err = errors.New("missing required view_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dns_settings/views/%s", accountID, viewID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type DNSViewParam struct {
	// The name of the view.
	Name param.Field[string] `json:"name"`
	// The list of zones linked to this view.
	Zones param.Field[[]string] `json:"zones"`
}

func (r DNSViewParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DNSViewResponse struct {
	// Identifier.
	ID string `json:"id,required"`
	// When the view was created.
	CreatedTime time.Time `json:"created_time,required" format:"date-time"`
	// When the view was last modified.
	ModifiedTime time.Time `json:"modified_time,required" format:"date-time"`
	// The name of the view.
	Name string `json:"name,required"`
	// The list of zones linked to this view.
	Zones []string            `json:"zones,required"`
	JSON  dnsViewResponseJSON `json:"-"`
}

// dnsViewResponseJSON contains the JSON metadata for the struct [DNSViewResponse]
type dnsViewResponseJSON struct {
	ID           apijson.Field
	CreatedTime  apijson.Field
	ModifiedTime apijson.Field
	Name         apijson.Field
	Zones        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *DNSViewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsViewResponseJSON) RawJSON() string {
	return r.raw
}

type DNSViewResponseSingle struct {
	Errors   []DNSSettingsMessages `json:"errors,required"`
	Messages []DNSSettingsMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success DNSViewResponseSingleSuccess `json:"success,required"`
	Result  DNSViewResponse              `json:"result"`
	JSON    dnsViewResponseSingleJSON    `json:"-"`
}

// dnsViewResponseSingleJSON contains the JSON metadata for the struct
// [DNSViewResponseSingle]
type dnsViewResponseSingleJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSViewResponseSingle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsViewResponseSingleJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DNSViewResponseSingleSuccess bool

const (
	DNSViewResponseSingleSuccessTrue DNSViewResponseSingleSuccess = true
)

func (r DNSViewResponseSingleSuccess) IsKnown() bool {
	switch r {
	case DNSViewResponseSingleSuccessTrue:
		return true
	}
	return false
}

type AccountDNSSettingViewListResponse struct {
	Errors   []DNSSettingsMessages `json:"errors,required"`
	Messages []DNSSettingsMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success    AccountDNSSettingViewListResponseSuccess    `json:"success,required"`
	Result     []DNSViewResponse                           `json:"result"`
	ResultInfo AccountDNSSettingViewListResponseResultInfo `json:"result_info"`
	JSON       accountDNSSettingViewListResponseJSON       `json:"-"`
}

// accountDNSSettingViewListResponseJSON contains the JSON metadata for the struct
// [AccountDNSSettingViewListResponse]
type accountDNSSettingViewListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDNSSettingViewListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountDNSSettingViewListResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccountDNSSettingViewListResponseSuccess bool

const (
	AccountDNSSettingViewListResponseSuccessTrue AccountDNSSettingViewListResponseSuccess = true
)

func (r AccountDNSSettingViewListResponseSuccess) IsKnown() bool {
	switch r {
	case AccountDNSSettingViewListResponseSuccessTrue:
		return true
	}
	return false
}

type AccountDNSSettingViewListResponseResultInfo struct {
	// Total number of results for the requested service.
	Count float64 `json:"count"`
	// Current page within paginated list of results.
	Page float64 `json:"page"`
	// Number of results per page of results.
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters.
	TotalCount float64                                         `json:"total_count"`
	JSON       accountDNSSettingViewListResponseResultInfoJSON `json:"-"`
}

// accountDNSSettingViewListResponseResultInfoJSON contains the JSON metadata for
// the struct [AccountDNSSettingViewListResponseResultInfo]
type accountDNSSettingViewListResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDNSSettingViewListResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountDNSSettingViewListResponseResultInfoJSON) RawJSON() string {
	return r.raw
}

type AccountDNSSettingViewDeleteResponse struct {
	Result AccountDNSSettingViewDeleteResponseResult `json:"result"`
	JSON   accountDNSSettingViewDeleteResponseJSON   `json:"-"`
}

// accountDNSSettingViewDeleteResponseJSON contains the JSON metadata for the
// struct [AccountDNSSettingViewDeleteResponse]
type accountDNSSettingViewDeleteResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDNSSettingViewDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountDNSSettingViewDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type AccountDNSSettingViewDeleteResponseResult struct {
	// Identifier.
	ID   string                                        `json:"id"`
	JSON accountDNSSettingViewDeleteResponseResultJSON `json:"-"`
}

// accountDNSSettingViewDeleteResponseResultJSON contains the JSON metadata for the
// struct [AccountDNSSettingViewDeleteResponseResult]
type accountDNSSettingViewDeleteResponseResultJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDNSSettingViewDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountDNSSettingViewDeleteResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountDNSSettingViewNewParams struct {
	DNSView DNSViewParam `json:"dns_view,required"`
}

func (r AccountDNSSettingViewNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.DNSView)
}

type AccountDNSSettingViewUpdateParams struct {
	DNSView DNSViewParam `json:"dns_view,required"`
}

func (r AccountDNSSettingViewUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.DNSView)
}

type AccountDNSSettingViewListParams struct {
	// Direction to order DNS views in.
	Direction param.Field[AccountDNSSettingViewListParamsDirection] `query:"direction"`
	// Whether to match all search requirements or at least one (any). If set to `all`,
	// acts like a logical AND between filters. If set to `any`, acts like a logical OR
	// instead.
	Match param.Field[AccountDNSSettingViewListParamsMatch] `query:"match"`
	Name  param.Field[AccountDNSSettingViewListParamsName]  `query:"name"`
	// Field to order DNS views by.
	Order param.Field[AccountDNSSettingViewListParamsOrder] `query:"order"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Number of DNS views per page.
	PerPage param.Field[float64] `query:"per_page"`
	// A zone ID that exists in the zones list for the view.
	ZoneID param.Field[string] `query:"zone_id"`
	// A zone name that exists in the zones list for the view.
	ZoneName param.Field[string] `query:"zone_name"`
}

// URLQuery serializes [AccountDNSSettingViewListParams]'s query parameters as
// `url.Values`.
func (r AccountDNSSettingViewListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Direction to order DNS views in.
type AccountDNSSettingViewListParamsDirection string

const (
	AccountDNSSettingViewListParamsDirectionAsc  AccountDNSSettingViewListParamsDirection = "asc"
	AccountDNSSettingViewListParamsDirectionDesc AccountDNSSettingViewListParamsDirection = "desc"
)

func (r AccountDNSSettingViewListParamsDirection) IsKnown() bool {
	switch r {
	case AccountDNSSettingViewListParamsDirectionAsc, AccountDNSSettingViewListParamsDirectionDesc:
		return true
	}
	return false
}

// Whether to match all search requirements or at least one (any). If set to `all`,
// acts like a logical AND between filters. If set to `any`, acts like a logical OR
// instead.
type AccountDNSSettingViewListParamsMatch string

const (
	AccountDNSSettingViewListParamsMatchAny AccountDNSSettingViewListParamsMatch = "any"
	AccountDNSSettingViewListParamsMatchAll AccountDNSSettingViewListParamsMatch = "all"
)

func (r AccountDNSSettingViewListParamsMatch) IsKnown() bool {
	switch r {
	case AccountDNSSettingViewListParamsMatchAny, AccountDNSSettingViewListParamsMatchAll:
		return true
	}
	return false
}

type AccountDNSSettingViewListParamsName struct {
	// Substring of the DNS view name.
	Contains param.Field[string] `query:"contains"`
	// Suffix of the DNS view name.
	Endswith param.Field[string] `query:"endswith"`
	// Exact value of the DNS view name.
	Exact param.Field[string] `query:"exact"`
	// Prefix of the DNS view name.
	Startswith param.Field[string] `query:"startswith"`
}

// URLQuery serializes [AccountDNSSettingViewListParamsName]'s query parameters as
// `url.Values`.
func (r AccountDNSSettingViewListParamsName) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Field to order DNS views by.
type AccountDNSSettingViewListParamsOrder string

const (
	AccountDNSSettingViewListParamsOrderName       AccountDNSSettingViewListParamsOrder = "name"
	AccountDNSSettingViewListParamsOrderCreatedOn  AccountDNSSettingViewListParamsOrder = "created_on"
	AccountDNSSettingViewListParamsOrderModifiedOn AccountDNSSettingViewListParamsOrder = "modified_on"
)

func (r AccountDNSSettingViewListParamsOrder) IsKnown() bool {
	switch r {
	case AccountDNSSettingViewListParamsOrderName, AccountDNSSettingViewListParamsOrderCreatedOn, AccountDNSSettingViewListParamsOrderModifiedOn:
		return true
	}
	return false
}
