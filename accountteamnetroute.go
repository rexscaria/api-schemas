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

// AccountTeamnetRouteService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountTeamnetRouteService] method instead.
type AccountTeamnetRouteService struct {
	Options []option.RequestOption
	Network *AccountTeamnetRouteNetworkService
}

// NewAccountTeamnetRouteService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountTeamnetRouteService(opts ...option.RequestOption) (r *AccountTeamnetRouteService) {
	r = &AccountTeamnetRouteService{}
	r.Options = opts
	r.Network = NewAccountTeamnetRouteNetworkService(opts...)
	return
}

// Routes a private network through a Cloudflare Tunnel.
func (r *AccountTeamnetRouteService) New(ctx context.Context, accountID string, body AccountTeamnetRouteNewParams, opts ...option.RequestOption) (res *TunnelRouteResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/teamnet/routes", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get a private network route in an account.
func (r *AccountTeamnetRouteService) Get(ctx context.Context, accountID string, routeID string, opts ...option.RequestOption) (res *TunnelRouteResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if routeID == "" {
		err = errors.New("missing required route_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/teamnet/routes/%s", accountID, routeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates an existing private network route in an account. The fields that are
// meant to be updated should be provided in the body of the request.
func (r *AccountTeamnetRouteService) Update(ctx context.Context, accountID string, routeID string, body AccountTeamnetRouteUpdateParams, opts ...option.RequestOption) (res *TunnelRouteResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if routeID == "" {
		err = errors.New("missing required route_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/teamnet/routes/%s", accountID, routeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Lists and filters private network routes in an account.
func (r *AccountTeamnetRouteService) List(ctx context.Context, accountID string, query AccountTeamnetRouteListParams, opts ...option.RequestOption) (res *AccountTeamnetRouteListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/teamnet/routes", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Deletes a private network route from an account.
func (r *AccountTeamnetRouteService) Delete(ctx context.Context, accountID string, routeID string, opts ...option.RequestOption) (res *TunnelRouteResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if routeID == "" {
		err = errors.New("missing required route_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/teamnet/routes/%s", accountID, routeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Fetches routes that contain the given IP address.
func (r *AccountTeamnetRouteService) GetByIP(ctx context.Context, accountID string, ip string, query AccountTeamnetRouteGetByIPParams, opts ...option.RequestOption) (res *AccountTeamnetRouteGetByIPResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if ip == "" {
		err = errors.New("missing required ip parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/teamnet/routes/ip/%s", accountID, ip)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type Teamnet struct {
	// UUID of the route.
	ID string `json:"id"`
	// Optional remark describing the route.
	Comment string `json:"comment"`
	// Timestamp of when the resource was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Timestamp of when the resource was deleted. If `null`, the resource has not been
	// deleted.
	DeletedAt time.Time `json:"deleted_at" format:"date-time"`
	// The private IPv4 or IPv6 range connected by the route, in CIDR notation.
	Network string `json:"network"`
	// The type of tunnel.
	TunType TunnelType `json:"tun_type"`
	// UUID of the tunnel.
	TunnelID string `json:"tunnel_id" format:"uuid"`
	// A user-friendly name for a tunnel.
	TunnelName string `json:"tunnel_name"`
	// UUID of the virtual network.
	VirtualNetworkID string `json:"virtual_network_id" format:"uuid"`
	// A user-friendly name for the virtual network.
	VirtualNetworkName string      `json:"virtual_network_name"`
	JSON               teamnetJSON `json:"-"`
}

// teamnetJSON contains the JSON metadata for the struct [Teamnet]
type teamnetJSON struct {
	ID                 apijson.Field
	Comment            apijson.Field
	CreatedAt          apijson.Field
	DeletedAt          apijson.Field
	Network            apijson.Field
	TunType            apijson.Field
	TunnelID           apijson.Field
	TunnelName         apijson.Field
	VirtualNetworkID   apijson.Field
	VirtualNetworkName apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *Teamnet) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r teamnetJSON) RawJSON() string {
	return r.raw
}

type TunnelRouteResponseSingle struct {
	Result TunnelRouteResponseSingleResult `json:"result"`
	JSON   tunnelRouteResponseSingleJSON   `json:"-"`
	APIResponseTunnel
}

// tunnelRouteResponseSingleJSON contains the JSON metadata for the struct
// [TunnelRouteResponseSingle]
type tunnelRouteResponseSingleJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelRouteResponseSingle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tunnelRouteResponseSingleJSON) RawJSON() string {
	return r.raw
}

type TunnelRouteResponseSingleResult struct {
	// UUID of the route.
	ID string `json:"id"`
	// Optional remark describing the route.
	Comment string `json:"comment"`
	// Timestamp of when the resource was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Timestamp of when the resource was deleted. If `null`, the resource has not been
	// deleted.
	DeletedAt time.Time `json:"deleted_at" format:"date-time"`
	// The private IPv4 or IPv6 range connected by the route, in CIDR notation.
	Network string `json:"network"`
	// UUID of the tunnel.
	TunnelID string `json:"tunnel_id" format:"uuid"`
	// UUID of the virtual network.
	VirtualNetworkID string                              `json:"virtual_network_id" format:"uuid"`
	JSON             tunnelRouteResponseSingleResultJSON `json:"-"`
}

// tunnelRouteResponseSingleResultJSON contains the JSON metadata for the struct
// [TunnelRouteResponseSingleResult]
type tunnelRouteResponseSingleResultJSON struct {
	ID               apijson.Field
	Comment          apijson.Field
	CreatedAt        apijson.Field
	DeletedAt        apijson.Field
	Network          apijson.Field
	TunnelID         apijson.Field
	VirtualNetworkID apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *TunnelRouteResponseSingleResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tunnelRouteResponseSingleResultJSON) RawJSON() string {
	return r.raw
}

type AccountTeamnetRouteListResponse struct {
	Result []Teamnet                           `json:"result"`
	JSON   accountTeamnetRouteListResponseJSON `json:"-"`
	APIResponseCollectionTunnel
}

// accountTeamnetRouteListResponseJSON contains the JSON metadata for the struct
// [AccountTeamnetRouteListResponse]
type accountTeamnetRouteListResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountTeamnetRouteListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountTeamnetRouteListResponseJSON) RawJSON() string {
	return r.raw
}

type AccountTeamnetRouteGetByIPResponse struct {
	Result Teamnet                                `json:"result"`
	JSON   accountTeamnetRouteGetByIPResponseJSON `json:"-"`
	APIResponseTunnel
}

// accountTeamnetRouteGetByIPResponseJSON contains the JSON metadata for the struct
// [AccountTeamnetRouteGetByIPResponse]
type accountTeamnetRouteGetByIPResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountTeamnetRouteGetByIPResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountTeamnetRouteGetByIPResponseJSON) RawJSON() string {
	return r.raw
}

type AccountTeamnetRouteNewParams struct {
	// The private IPv4 or IPv6 range connected by the route, in CIDR notation.
	Network param.Field[string] `json:"network,required"`
	// Optional remark describing the route.
	Comment param.Field[string] `json:"comment"`
	// UUID of the virtual network.
	VirtualNetworkID param.Field[string] `json:"virtual_network_id" format:"uuid"`
}

func (r AccountTeamnetRouteNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountTeamnetRouteUpdateParams struct {
	// Optional remark describing the route.
	Comment param.Field[string] `json:"comment"`
	// The private IPv4 or IPv6 range connected by the route, in CIDR notation.
	Network param.Field[string] `json:"network"`
	// UUID of the virtual network.
	VirtualNetworkID param.Field[string] `json:"virtual_network_id" format:"uuid"`
}

func (r AccountTeamnetRouteUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountTeamnetRouteListParams struct {
	// Optional remark describing the route.
	Comment param.Field[string] `query:"comment"`
	// If provided, include only resources that were created (and not deleted) before
	// this time. URL encoded.
	ExistedAt param.Field[string] `query:"existed_at" format:"url-encoded-date-time"`
	// If `true`, only include deleted routes. If `false`, exclude deleted routes. If
	// empty, all routes will be included.
	IsDeleted param.Field[bool] `query:"is_deleted"`
	// The private IPv4 or IPv6 range connected by the route, in CIDR notation.
	NetworkSubset param.Field[string] `query:"network_subset"`
	// The private IPv4 or IPv6 range connected by the route, in CIDR notation.
	NetworkSuperset param.Field[string] `query:"network_superset"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Number of results to display.
	PerPage param.Field[float64] `query:"per_page"`
	// UUID of the route.
	RouteID param.Field[string] `query:"route_id"`
	// The types of tunnels to filter by, separated by commas.
	TunTypes param.Field[[]TunnelType] `query:"tun_types"`
	// UUID of the tunnel.
	TunnelID param.Field[string] `query:"tunnel_id" format:"uuid"`
	// UUID of the virtual network.
	VirtualNetworkID param.Field[string] `query:"virtual_network_id" format:"uuid"`
}

// URLQuery serializes [AccountTeamnetRouteListParams]'s query parameters as
// `url.Values`.
func (r AccountTeamnetRouteListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountTeamnetRouteGetByIPParams struct {
	// When the virtual_network_id parameter is not provided the request filter will
	// default search routes that are in the default virtual network for the account.
	// If this parameter is set to false, the search will include routes that do not
	// have a virtual network.
	DefaultVirtualNetworkFallback param.Field[bool] `query:"default_virtual_network_fallback"`
	// UUID of the virtual network.
	VirtualNetworkID param.Field[string] `query:"virtual_network_id" format:"uuid"`
}

// URLQuery serializes [AccountTeamnetRouteGetByIPParams]'s query parameters as
// `url.Values`.
func (r AccountTeamnetRouteGetByIPParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
