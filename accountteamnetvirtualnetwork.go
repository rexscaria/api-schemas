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

// AccountTeamnetVirtualNetworkService contains methods and other services that
// help with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountTeamnetVirtualNetworkService] method instead.
type AccountTeamnetVirtualNetworkService struct {
	Options []option.RequestOption
}

// NewAccountTeamnetVirtualNetworkService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountTeamnetVirtualNetworkService(opts ...option.RequestOption) (r *AccountTeamnetVirtualNetworkService) {
	r = &AccountTeamnetVirtualNetworkService{}
	r.Options = opts
	return
}

// Adds a new virtual network to an account.
func (r *AccountTeamnetVirtualNetworkService) New(ctx context.Context, accountID string, body AccountTeamnetVirtualNetworkNewParams, opts ...option.RequestOption) (res *VnetResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/teamnet/virtual_networks", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get a virtual network.
func (r *AccountTeamnetVirtualNetworkService) Get(ctx context.Context, accountID string, virtualNetworkID string, query AccountTeamnetVirtualNetworkGetParams, opts ...option.RequestOption) (res *VnetResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if virtualNetworkID == "" {
		err = errors.New("missing required virtual_network_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/teamnet/virtual_networks/%s", accountID, virtualNetworkID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Updates an existing virtual network.
func (r *AccountTeamnetVirtualNetworkService) Update(ctx context.Context, accountID string, virtualNetworkID string, body AccountTeamnetVirtualNetworkUpdateParams, opts ...option.RequestOption) (res *VnetResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if virtualNetworkID == "" {
		err = errors.New("missing required virtual_network_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/teamnet/virtual_networks/%s", accountID, virtualNetworkID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Lists and filters virtual networks in an account.
func (r *AccountTeamnetVirtualNetworkService) List(ctx context.Context, accountID string, query AccountTeamnetVirtualNetworkListParams, opts ...option.RequestOption) (res *AccountTeamnetVirtualNetworkListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/teamnet/virtual_networks", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Deletes an existing virtual network.
func (r *AccountTeamnetVirtualNetworkService) Delete(ctx context.Context, accountID string, virtualNetworkID string, body AccountTeamnetVirtualNetworkDeleteParams, opts ...option.RequestOption) (res *VnetResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if virtualNetworkID == "" {
		err = errors.New("missing required virtual_network_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/teamnet/virtual_networks/%s", accountID, virtualNetworkID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

type VirtualNetwork struct {
	// UUID of the virtual network.
	ID string `json:"id,required" format:"uuid"`
	// Optional remark describing the virtual network.
	Comment string `json:"comment,required"`
	// Timestamp of when the resource was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// If `true`, this virtual network is the default for the account.
	IsDefaultNetwork bool `json:"is_default_network,required"`
	// A user-friendly name for the virtual network.
	Name string `json:"name,required"`
	// Timestamp of when the resource was deleted. If `null`, the resource has not been
	// deleted.
	DeletedAt time.Time          `json:"deleted_at" format:"date-time"`
	JSON      virtualNetworkJSON `json:"-"`
}

// virtualNetworkJSON contains the JSON metadata for the struct [VirtualNetwork]
type virtualNetworkJSON struct {
	ID               apijson.Field
	Comment          apijson.Field
	CreatedAt        apijson.Field
	IsDefaultNetwork apijson.Field
	Name             apijson.Field
	DeletedAt        apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *VirtualNetwork) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r virtualNetworkJSON) RawJSON() string {
	return r.raw
}

type VnetResponseSingle struct {
	Result VirtualNetwork         `json:"result"`
	JSON   vnetResponseSingleJSON `json:"-"`
	APIResponseTunnel
}

// vnetResponseSingleJSON contains the JSON metadata for the struct
// [VnetResponseSingle]
type vnetResponseSingleJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VnetResponseSingle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r vnetResponseSingleJSON) RawJSON() string {
	return r.raw
}

type AccountTeamnetVirtualNetworkListResponse struct {
	Result []VirtualNetwork                             `json:"result"`
	JSON   accountTeamnetVirtualNetworkListResponseJSON `json:"-"`
	APIResponseCollectionTunnel
}

// accountTeamnetVirtualNetworkListResponseJSON contains the JSON metadata for the
// struct [AccountTeamnetVirtualNetworkListResponse]
type accountTeamnetVirtualNetworkListResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountTeamnetVirtualNetworkListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountTeamnetVirtualNetworkListResponseJSON) RawJSON() string {
	return r.raw
}

type AccountTeamnetVirtualNetworkNewParams struct {
	// A user-friendly name for the virtual network.
	Name param.Field[string] `json:"name,required"`
	// Optional remark describing the virtual network.
	Comment param.Field[string] `json:"comment"`
	// If `true`, this virtual network is the default for the account.
	IsDefault param.Field[bool] `json:"is_default"`
}

func (r AccountTeamnetVirtualNetworkNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountTeamnetVirtualNetworkGetParams struct {
}

type AccountTeamnetVirtualNetworkUpdateParams struct {
	// Optional remark describing the virtual network.
	Comment param.Field[string] `json:"comment"`
	// If `true`, this virtual network is the default for the account.
	IsDefaultNetwork param.Field[bool] `json:"is_default_network"`
	// A user-friendly name for the virtual network.
	Name param.Field[string] `json:"name"`
}

func (r AccountTeamnetVirtualNetworkUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountTeamnetVirtualNetworkListParams struct {
	// UUID of the virtual network.
	ID param.Field[string] `query:"id" format:"uuid"`
	// If `true`, only include the default virtual network. If `false`, exclude the
	// default virtual network. If empty, all virtual networks will be included.
	IsDefault param.Field[bool] `query:"is_default"`
	// If `true`, only include deleted virtual networks. If `false`, exclude deleted
	// virtual networks. If empty, all virtual networks will be included.
	IsDeleted param.Field[bool] `query:"is_deleted"`
	// A user-friendly name for the virtual network.
	Name param.Field[string] `query:"name"`
}

// URLQuery serializes [AccountTeamnetVirtualNetworkListParams]'s query parameters
// as `url.Values`.
func (r AccountTeamnetVirtualNetworkListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountTeamnetVirtualNetworkDeleteParams struct {
	Body interface{} `json:"body,required"`
}

func (r AccountTeamnetVirtualNetworkDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}
