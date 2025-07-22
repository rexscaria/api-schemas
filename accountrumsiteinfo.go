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

// AccountRumSiteInfoService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountRumSiteInfoService] method instead.
type AccountRumSiteInfoService struct {
	Options []option.RequestOption
}

// NewAccountRumSiteInfoService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountRumSiteInfoService(opts ...option.RequestOption) (r *AccountRumSiteInfoService) {
	r = &AccountRumSiteInfoService{}
	r.Options = opts
	return
}

// Creates a new Web Analytics site.
func (r *AccountRumSiteInfoService) New(ctx context.Context, accountID string, body AccountRumSiteInfoNewParams, opts ...option.RequestOption) (res *ResponseSingleSite, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/rum/site_info", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieves a Web Analytics site.
func (r *AccountRumSiteInfoService) Get(ctx context.Context, accountID string, siteID string, opts ...option.RequestOption) (res *ResponseSingleSite, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if siteID == "" {
		err = errors.New("missing required site_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/rum/site_info/%s", accountID, siteID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates an existing Web Analytics site.
func (r *AccountRumSiteInfoService) Update(ctx context.Context, accountID string, siteID string, body AccountRumSiteInfoUpdateParams, opts ...option.RequestOption) (res *ResponseSingleSite, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if siteID == "" {
		err = errors.New("missing required site_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/rum/site_info/%s", accountID, siteID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Lists all Web Analytics sites of an account.
func (r *AccountRumSiteInfoService) List(ctx context.Context, accountID string, query AccountRumSiteInfoListParams, opts ...option.RequestOption) (res *AccountRumSiteInfoListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/rum/site_info/list", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Deletes an existing Web Analytics site.
func (r *AccountRumSiteInfoService) Delete(ctx context.Context, accountID string, siteID string, opts ...option.RequestOption) (res *AccountRumSiteInfoDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if siteID == "" {
		err = errors.New("missing required site_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/rum/site_info/%s", accountID, siteID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type Common struct {
	Errors   []RumMessages `json:"errors,required"`
	Messages []RumMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success bool       `json:"success,required"`
	JSON    commonJSON `json:"-"`
}

// commonJSON contains the JSON metadata for the struct [Common]
type commonJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Common) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r commonJSON) RawJSON() string {
	return r.raw
}

type ResponseSingleSite struct {
	Result Site                   `json:"result"`
	JSON   responseSingleSiteJSON `json:"-"`
	Common
}

// responseSingleSiteJSON contains the JSON metadata for the struct
// [ResponseSingleSite]
type responseSingleSiteJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResponseSingleSite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r responseSingleSiteJSON) RawJSON() string {
	return r.raw
}

type RumMessages struct {
	Code    int64           `json:"code,required"`
	Message string          `json:"message,required"`
	JSON    rumMessagesJSON `json:"-"`
}

// rumMessagesJSON contains the JSON metadata for the struct [RumMessages]
type rumMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RumMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rumMessagesJSON) RawJSON() string {
	return r.raw
}

type Site struct {
	// If enabled, the JavaScript snippet is automatically injected for orange-clouded
	// sites.
	AutoInstall bool      `json:"auto_install"`
	Created     time.Time `json:"created" format:"date-time"`
	// A list of rules.
	Rules   []RumRule `json:"rules"`
	Ruleset Ruleset   `json:"ruleset"`
	// The Web Analytics site identifier.
	SiteTag string `json:"site_tag"`
	// The Web Analytics site token.
	SiteToken string `json:"site_token"`
	// Encoded JavaScript snippet.
	Snippet string   `json:"snippet"`
	JSON    siteJSON `json:"-"`
}

// siteJSON contains the JSON metadata for the struct [Site]
type siteJSON struct {
	AutoInstall apijson.Field
	Created     apijson.Field
	Rules       apijson.Field
	Ruleset     apijson.Field
	SiteTag     apijson.Field
	SiteToken   apijson.Field
	Snippet     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Site) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteJSON) RawJSON() string {
	return r.raw
}

type AccountRumSiteInfoListResponse struct {
	Result     []Site                                   `json:"result"`
	ResultInfo AccountRumSiteInfoListResponseResultInfo `json:"result_info"`
	JSON       accountRumSiteInfoListResponseJSON       `json:"-"`
	Common
}

// accountRumSiteInfoListResponseJSON contains the JSON metadata for the struct
// [AccountRumSiteInfoListResponse]
type accountRumSiteInfoListResponseJSON struct {
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRumSiteInfoListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountRumSiteInfoListResponseJSON) RawJSON() string {
	return r.raw
}

type AccountRumSiteInfoListResponseResultInfo struct {
	// The total number of items on the current page.
	Count int64 `json:"count"`
	// Current page within the paginated list of results.
	Page int64 `json:"page"`
	// The maximum number of items to return per page of results.
	PerPage int64 `json:"per_page"`
	// The total number of items.
	TotalCount int64 `json:"total_count"`
	// The total number of pages.
	TotalPages int64                                        `json:"total_pages,nullable"`
	JSON       accountRumSiteInfoListResponseResultInfoJSON `json:"-"`
}

// accountRumSiteInfoListResponseResultInfoJSON contains the JSON metadata for the
// struct [AccountRumSiteInfoListResponseResultInfo]
type accountRumSiteInfoListResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	TotalPages  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRumSiteInfoListResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountRumSiteInfoListResponseResultInfoJSON) RawJSON() string {
	return r.raw
}

type AccountRumSiteInfoDeleteResponse struct {
	Result AccountRumSiteInfoDeleteResponseResult `json:"result"`
	JSON   accountRumSiteInfoDeleteResponseJSON   `json:"-"`
	Common
}

// accountRumSiteInfoDeleteResponseJSON contains the JSON metadata for the struct
// [AccountRumSiteInfoDeleteResponse]
type accountRumSiteInfoDeleteResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRumSiteInfoDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountRumSiteInfoDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type AccountRumSiteInfoDeleteResponseResult struct {
	// The Web Analytics site identifier.
	SiteTag string                                     `json:"site_tag"`
	JSON    accountRumSiteInfoDeleteResponseResultJSON `json:"-"`
}

// accountRumSiteInfoDeleteResponseResultJSON contains the JSON metadata for the
// struct [AccountRumSiteInfoDeleteResponseResult]
type accountRumSiteInfoDeleteResponseResultJSON struct {
	SiteTag     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRumSiteInfoDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountRumSiteInfoDeleteResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountRumSiteInfoNewParams struct {
	// If enabled, the JavaScript snippet is automatically injected for orange-clouded
	// sites.
	AutoInstall param.Field[bool] `json:"auto_install"`
	// The hostname to use for gray-clouded sites.
	Host param.Field[string] `json:"host"`
	// The zone identifier.
	ZoneTag param.Field[string] `json:"zone_tag"`
}

func (r AccountRumSiteInfoNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountRumSiteInfoUpdateParams struct {
	// If enabled, the JavaScript snippet is automatically injected for orange-clouded
	// sites.
	AutoInstall param.Field[bool] `json:"auto_install"`
	// Enables or disables RUM. This option can be used only when auto_install is set
	// to true.
	Enabled param.Field[bool] `json:"enabled"`
	// The hostname to use for gray-clouded sites.
	Host param.Field[string] `json:"host"`
	// If enabled, the JavaScript snippet will not be injected for visitors from the
	// EU.
	Lite param.Field[bool] `json:"lite"`
	// The zone identifier.
	ZoneTag param.Field[string] `json:"zone_tag"`
}

func (r AccountRumSiteInfoUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountRumSiteInfoListParams struct {
	// The property used to sort the list of results.
	OrderBy param.Field[AccountRumSiteInfoListParamsOrderBy] `query:"order_by"`
	// Current page within the paginated list of results.
	Page param.Field[float64] `query:"page"`
	// Number of items to return per page of results.
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes [AccountRumSiteInfoListParams]'s query parameters as
// `url.Values`.
func (r AccountRumSiteInfoListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The property used to sort the list of results.
type AccountRumSiteInfoListParamsOrderBy string

const (
	AccountRumSiteInfoListParamsOrderByHost    AccountRumSiteInfoListParamsOrderBy = "host"
	AccountRumSiteInfoListParamsOrderByCreated AccountRumSiteInfoListParamsOrderBy = "created"
)

func (r AccountRumSiteInfoListParamsOrderBy) IsKnown() bool {
	switch r {
	case AccountRumSiteInfoListParamsOrderByHost, AccountRumSiteInfoListParamsOrderByCreated:
		return true
	}
	return false
}
