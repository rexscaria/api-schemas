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

// AccountMagicSiteService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountMagicSiteService] method instead.
type AccountMagicSiteService struct {
	Options    []option.RequestOption
	ACLs       *AccountMagicSiteACLService
	AppConfigs *AccountMagicSiteAppConfigService
	Lans       *AccountMagicSiteLanService
	Wans       *AccountMagicSiteWanService
}

// NewAccountMagicSiteService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountMagicSiteService(opts ...option.RequestOption) (r *AccountMagicSiteService) {
	r = &AccountMagicSiteService{}
	r.Options = opts
	r.ACLs = NewAccountMagicSiteACLService(opts...)
	r.AppConfigs = NewAccountMagicSiteAppConfigService(opts...)
	r.Lans = NewAccountMagicSiteLanService(opts...)
	r.Wans = NewAccountMagicSiteWanService(opts...)
	return
}

// Creates a new Site
func (r *AccountMagicSiteService) New(ctx context.Context, accountID string, body AccountMagicSiteNewParams, opts ...option.RequestOption) (res *MagicSiteSingleResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/sites", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get a specific Site.
func (r *AccountMagicSiteService) Get(ctx context.Context, accountID string, siteID string, query AccountMagicSiteGetParams, opts ...option.RequestOption) (res *MagicSiteSingleResponse, err error) {
	if query.XMagicNewHcTarget.Present {
		opts = append(opts, option.WithHeader("x-magic-new-hc-target", fmt.Sprintf("%s", query.XMagicNewHcTarget)))
	}
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if siteID == "" {
		err = errors.New("missing required site_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/sites/%s", accountID, siteID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update a specific Site.
func (r *AccountMagicSiteService) Update(ctx context.Context, accountID string, siteID string, body AccountMagicSiteUpdateParams, opts ...option.RequestOption) (res *MagicSiteModifiedResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if siteID == "" {
		err = errors.New("missing required site_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/sites/%s", accountID, siteID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Lists Sites associated with an account. Use connectorid query param to return
// sites where connectorid matches either site.ConnectorID or
// site.SecondaryConnectorID.
func (r *AccountMagicSiteService) List(ctx context.Context, accountID string, query AccountMagicSiteListParams, opts ...option.RequestOption) (res *AccountMagicSiteListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/sites", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Remove a specific Site.
func (r *AccountMagicSiteService) Delete(ctx context.Context, accountID string, siteID string, opts ...option.RequestOption) (res *AccountMagicSiteDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if siteID == "" {
		err = errors.New("missing required site_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/sites/%s", accountID, siteID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Patch a specific Site.
func (r *AccountMagicSiteService) Patch(ctx context.Context, accountID string, siteID string, body AccountMagicSitePatchParams, opts ...option.RequestOption) (res *MagicSiteModifiedResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if siteID == "" {
		err = errors.New("missing required site_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/sites/%s", accountID, siteID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

type MagicSite struct {
	// Identifier
	ID string `json:"id"`
	// Magic Connector identifier tag.
	ConnectorID string `json:"connector_id"`
	Description string `json:"description"`
	// Site high availability mode. If set to true, the site can have two connectors
	// and runs in high availability mode.
	HaMode bool `json:"ha_mode"`
	// Location of site in latitude and longitude.
	Location MagicSiteLocation `json:"location"`
	// The name of the site.
	Name string `json:"name"`
	// Magic Connector identifier tag. Used when high availability mode is on.
	SecondaryConnectorID string        `json:"secondary_connector_id"`
	JSON                 magicSiteJSON `json:"-"`
}

// magicSiteJSON contains the JSON metadata for the struct [MagicSite]
type magicSiteJSON struct {
	ID                   apijson.Field
	ConnectorID          apijson.Field
	Description          apijson.Field
	HaMode               apijson.Field
	Location             apijson.Field
	Name                 apijson.Field
	SecondaryConnectorID apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *MagicSite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicSiteJSON) RawJSON() string {
	return r.raw
}

// Location of site in latitude and longitude.
type MagicSiteLocation struct {
	// Latitude
	Lat string `json:"lat"`
	// Longitude
	Lon  string                `json:"lon"`
	JSON magicSiteLocationJSON `json:"-"`
}

// magicSiteLocationJSON contains the JSON metadata for the struct
// [MagicSiteLocation]
type magicSiteLocationJSON struct {
	Lat         apijson.Field
	Lon         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicSiteLocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicSiteLocationJSON) RawJSON() string {
	return r.raw
}

// Location of site in latitude and longitude.
type MagicSiteLocationParam struct {
	// Latitude
	Lat param.Field[string] `json:"lat"`
	// Longitude
	Lon param.Field[string] `json:"lon"`
}

func (r MagicSiteLocationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MagicSiteModifiedResponse struct {
	Errors   []MagicMessageItem `json:"errors,required"`
	Messages []MagicMessageItem `json:"messages,required"`
	Result   MagicSite          `json:"result,required"`
	// Whether the API call was successful
	Success MagicSiteModifiedResponseSuccess `json:"success,required"`
	JSON    magicSiteModifiedResponseJSON    `json:"-"`
}

// magicSiteModifiedResponseJSON contains the JSON metadata for the struct
// [MagicSiteModifiedResponse]
type magicSiteModifiedResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicSiteModifiedResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicSiteModifiedResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type MagicSiteModifiedResponseSuccess bool

const (
	MagicSiteModifiedResponseSuccessTrue MagicSiteModifiedResponseSuccess = true
)

func (r MagicSiteModifiedResponseSuccess) IsKnown() bool {
	switch r {
	case MagicSiteModifiedResponseSuccessTrue:
		return true
	}
	return false
}

type MagicSiteSingleResponse struct {
	Errors   []MagicMessageItem `json:"errors,required"`
	Messages []MagicMessageItem `json:"messages,required"`
	Result   MagicSite          `json:"result,required"`
	// Whether the API call was successful
	Success MagicSiteSingleResponseSuccess `json:"success,required"`
	JSON    magicSiteSingleResponseJSON    `json:"-"`
}

// magicSiteSingleResponseJSON contains the JSON metadata for the struct
// [MagicSiteSingleResponse]
type magicSiteSingleResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicSiteSingleResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicSiteSingleResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type MagicSiteSingleResponseSuccess bool

const (
	MagicSiteSingleResponseSuccessTrue MagicSiteSingleResponseSuccess = true
)

func (r MagicSiteSingleResponseSuccess) IsKnown() bool {
	switch r {
	case MagicSiteSingleResponseSuccessTrue:
		return true
	}
	return false
}

type MagicSiteUpdateRequestParam struct {
	// Magic Connector identifier tag.
	ConnectorID param.Field[string] `json:"connector_id"`
	Description param.Field[string] `json:"description"`
	// Location of site in latitude and longitude.
	Location param.Field[MagicSiteLocationParam] `json:"location"`
	// The name of the site.
	Name param.Field[string] `json:"name"`
	// Magic Connector identifier tag. Used when high availability mode is on.
	SecondaryConnectorID param.Field[string] `json:"secondary_connector_id"`
}

func (r MagicSiteUpdateRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountMagicSiteListResponse struct {
	Errors   []MagicMessageItem `json:"errors,required"`
	Messages []MagicMessageItem `json:"messages,required"`
	Result   []MagicSite        `json:"result,required"`
	// Whether the API call was successful
	Success AccountMagicSiteListResponseSuccess `json:"success,required"`
	JSON    accountMagicSiteListResponseJSON    `json:"-"`
}

// accountMagicSiteListResponseJSON contains the JSON metadata for the struct
// [AccountMagicSiteListResponse]
type accountMagicSiteListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicSiteListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMagicSiteListResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AccountMagicSiteListResponseSuccess bool

const (
	AccountMagicSiteListResponseSuccessTrue AccountMagicSiteListResponseSuccess = true
)

func (r AccountMagicSiteListResponseSuccess) IsKnown() bool {
	switch r {
	case AccountMagicSiteListResponseSuccessTrue:
		return true
	}
	return false
}

type AccountMagicSiteDeleteResponse struct {
	Errors   []MagicMessageItem `json:"errors,required"`
	Messages []MagicMessageItem `json:"messages,required"`
	Result   MagicSite          `json:"result,required"`
	// Whether the API call was successful
	Success AccountMagicSiteDeleteResponseSuccess `json:"success,required"`
	JSON    accountMagicSiteDeleteResponseJSON    `json:"-"`
}

// accountMagicSiteDeleteResponseJSON contains the JSON metadata for the struct
// [AccountMagicSiteDeleteResponse]
type accountMagicSiteDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicSiteDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMagicSiteDeleteResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AccountMagicSiteDeleteResponseSuccess bool

const (
	AccountMagicSiteDeleteResponseSuccessTrue AccountMagicSiteDeleteResponseSuccess = true
)

func (r AccountMagicSiteDeleteResponseSuccess) IsKnown() bool {
	switch r {
	case AccountMagicSiteDeleteResponseSuccessTrue:
		return true
	}
	return false
}

type AccountMagicSiteNewParams struct {
	// The name of the site.
	Name param.Field[string] `json:"name,required"`
	// Magic Connector identifier tag.
	ConnectorID param.Field[string] `json:"connector_id"`
	Description param.Field[string] `json:"description"`
	// Site high availability mode. If set to true, the site can have two connectors
	// and runs in high availability mode.
	HaMode param.Field[bool] `json:"ha_mode"`
	// Location of site in latitude and longitude.
	Location param.Field[MagicSiteLocationParam] `json:"location"`
	// Magic Connector identifier tag. Used when high availability mode is on.
	SecondaryConnectorID param.Field[string] `json:"secondary_connector_id"`
}

func (r AccountMagicSiteNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountMagicSiteGetParams struct {
	XMagicNewHcTarget param.Field[bool] `header:"x-magic-new-hc-target"`
}

type AccountMagicSiteUpdateParams struct {
	MagicSiteUpdateRequest MagicSiteUpdateRequestParam `json:"magic_site_update_request,required"`
}

func (r AccountMagicSiteUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.MagicSiteUpdateRequest)
}

type AccountMagicSiteListParams struct {
	// Identifier
	Connectorid param.Field[string] `query:"connectorid"`
}

// URLQuery serializes [AccountMagicSiteListParams]'s query parameters as
// `url.Values`.
func (r AccountMagicSiteListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountMagicSitePatchParams struct {
	MagicSiteUpdateRequest MagicSiteUpdateRequestParam `json:"magic_site_update_request,required"`
}

func (r AccountMagicSitePatchParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.MagicSiteUpdateRequest)
}
