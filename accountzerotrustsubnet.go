// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/stainless-sdks/cf-rex-go/internal/apijson"
	"github.com/stainless-sdks/cf-rex-go/internal/apiquery"
	"github.com/stainless-sdks/cf-rex-go/internal/param"
	"github.com/stainless-sdks/cf-rex-go/internal/requestconfig"
	"github.com/stainless-sdks/cf-rex-go/option"
)

// AccountZerotrustSubnetService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountZerotrustSubnetService] method instead.
type AccountZerotrustSubnetService struct {
	Options []option.RequestOption
}

// NewAccountZerotrustSubnetService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountZerotrustSubnetService(opts ...option.RequestOption) (r *AccountZerotrustSubnetService) {
	r = &AccountZerotrustSubnetService{}
	r.Options = opts
	return
}

// Lists and filters subnets in an account.
func (r *AccountZerotrustSubnetService) List(ctx context.Context, accountID string, query AccountZerotrustSubnetListParams, opts ...option.RequestOption) (res *AccountZerotrustSubnetListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/zerotrust/subnets", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Updates the Cloudflare Source subnet of the given address family
func (r *AccountZerotrustSubnetService) UpdateCloudflareSource(ctx context.Context, accountID string, addressFamily AddressFamily, body AccountZerotrustSubnetUpdateCloudflareSourceParams, opts ...option.RequestOption) (res *AccountZerotrustSubnetUpdateCloudflareSourceResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/zerotrust/subnets/cloudflare_source/%v", accountID, addressFamily)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// IP address family, either `v4` (IPv4) or `v6` (IPv6)
type AddressFamily string

const (
	AddressFamilyV4 AddressFamily = "v4"
	AddressFamilyV6 AddressFamily = "v6"
)

func (r AddressFamily) IsKnown() bool {
	switch r {
	case AddressFamilyV4, AddressFamilyV6:
		return true
	}
	return false
}

type Subnet struct {
	// The UUID of the subnet.
	ID string `json:"id" format:"uuid"`
	// An optional description of the subnet.
	Comment string `json:"comment"`
	// Timestamp of when the resource was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Timestamp of when the resource was deleted. If `null`, the resource has not been
	// deleted.
	DeletedAt time.Time `json:"deleted_at" format:"date-time"`
	// If `true`, this is the default subnet for the account. There can only be one
	// default subnet per account.
	IsDefaultNetwork bool `json:"is_default_network"`
	// A user-friendly name for the subnet.
	Name string `json:"name"`
	// The private IPv4 or IPv6 range defining the subnet, in CIDR notation.
	Network string `json:"network"`
	// The type of subnet.
	SubnetType SubnetSubnetType `json:"subnet_type"`
	JSON       subnetJSON       `json:"-"`
}

// subnetJSON contains the JSON metadata for the struct [Subnet]
type subnetJSON struct {
	ID               apijson.Field
	Comment          apijson.Field
	CreatedAt        apijson.Field
	DeletedAt        apijson.Field
	IsDefaultNetwork apijson.Field
	Name             apijson.Field
	Network          apijson.Field
	SubnetType       apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *Subnet) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subnetJSON) RawJSON() string {
	return r.raw
}

// The type of subnet.
type SubnetSubnetType string

const (
	SubnetSubnetTypeCloudflareSource SubnetSubnetType = "cloudflare_source"
)

func (r SubnetSubnetType) IsKnown() bool {
	switch r {
	case SubnetSubnetTypeCloudflareSource:
		return true
	}
	return false
}

type AccountZerotrustSubnetListResponse struct {
	Result []Subnet                               `json:"result"`
	JSON   accountZerotrustSubnetListResponseJSON `json:"-"`
	APIResponseCollectionTunnel
}

// accountZerotrustSubnetListResponseJSON contains the JSON metadata for the struct
// [AccountZerotrustSubnetListResponse]
type accountZerotrustSubnetListResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountZerotrustSubnetListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountZerotrustSubnetListResponseJSON) RawJSON() string {
	return r.raw
}

type AccountZerotrustSubnetUpdateCloudflareSourceResponse struct {
	Result Subnet                                                   `json:"result"`
	JSON   accountZerotrustSubnetUpdateCloudflareSourceResponseJSON `json:"-"`
	APIResponseTunnel
}

// accountZerotrustSubnetUpdateCloudflareSourceResponseJSON contains the JSON
// metadata for the struct [AccountZerotrustSubnetUpdateCloudflareSourceResponse]
type accountZerotrustSubnetUpdateCloudflareSourceResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountZerotrustSubnetUpdateCloudflareSourceResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountZerotrustSubnetUpdateCloudflareSourceResponseJSON) RawJSON() string {
	return r.raw
}

type AccountZerotrustSubnetListParams struct {
	// If set, only include subnets in the given address family - `v4` or `v6`
	AddressFamily param.Field[AddressFamily] `query:"address_family"`
	// If set, only list subnets with the given comment.
	Comment param.Field[string] `query:"comment"`
	// If provided, include only resources that were created (and not deleted) before
	// this time. URL encoded.
	ExistedAt param.Field[string] `query:"existed_at" format:"url-encoded-date-time"`
	// If `true`, only include default subnets. If `false`, exclude default subnets
	// subnets. If not set, all subnets will be included.
	IsDefaultNetwork param.Field[bool] `query:"is_default_network"`
	// If `true`, only include deleted subnets. If `false`, exclude deleted subnets. If
	// not set, all subnets will be included.
	IsDeleted param.Field[bool] `query:"is_deleted"`
	// If set, only list subnets with the given name
	Name param.Field[string] `query:"name"`
	// IP/CIDR range in URL-encoded format
	Network param.Field[string] `query:"network"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Number of results to display.
	PerPage param.Field[float64] `query:"per_page"`
	// Sort order of the results. `asc` means oldest to newest, `desc` means newest to
	// oldest. If not set, they will not be in any particular order.
	SortOrder param.Field[AccountZerotrustSubnetListParamsSortOrder] `query:"sort_order"`
	// If set, the types of subnets to include, separated by comma.
	SubnetTypes param.Field[AccountZerotrustSubnetListParamsSubnetTypes] `query:"subnet_types"`
}

// URLQuery serializes [AccountZerotrustSubnetListParams]'s query parameters as
// `url.Values`.
func (r AccountZerotrustSubnetListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Sort order of the results. `asc` means oldest to newest, `desc` means newest to
// oldest. If not set, they will not be in any particular order.
type AccountZerotrustSubnetListParamsSortOrder string

const (
	AccountZerotrustSubnetListParamsSortOrderAsc  AccountZerotrustSubnetListParamsSortOrder = "asc"
	AccountZerotrustSubnetListParamsSortOrderDesc AccountZerotrustSubnetListParamsSortOrder = "desc"
)

func (r AccountZerotrustSubnetListParamsSortOrder) IsKnown() bool {
	switch r {
	case AccountZerotrustSubnetListParamsSortOrderAsc, AccountZerotrustSubnetListParamsSortOrderDesc:
		return true
	}
	return false
}

// If set, the types of subnets to include, separated by comma.
type AccountZerotrustSubnetListParamsSubnetTypes string

const (
	AccountZerotrustSubnetListParamsSubnetTypesCloudflareSource AccountZerotrustSubnetListParamsSubnetTypes = "cloudflare_source"
	AccountZerotrustSubnetListParamsSubnetTypesWarp             AccountZerotrustSubnetListParamsSubnetTypes = "warp"
)

func (r AccountZerotrustSubnetListParamsSubnetTypes) IsKnown() bool {
	switch r {
	case AccountZerotrustSubnetListParamsSubnetTypesCloudflareSource, AccountZerotrustSubnetListParamsSubnetTypesWarp:
		return true
	}
	return false
}

type AccountZerotrustSubnetUpdateCloudflareSourceParams struct {
	// An optional description of the subnet.
	Comment param.Field[string] `json:"comment"`
	// A user-friendly name for the subnet.
	Name param.Field[string] `json:"name"`
	// The private IPv4 or IPv6 range defining the subnet, in CIDR notation.
	Network param.Field[string] `json:"network"`
}

func (r AccountZerotrustSubnetUpdateCloudflareSourceParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
