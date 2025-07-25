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

// AccountTeamnetRouteNetworkService contains methods and other services that help
// with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountTeamnetRouteNetworkService] method instead.
type AccountTeamnetRouteNetworkService struct {
	Options []option.RequestOption
}

// NewAccountTeamnetRouteNetworkService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountTeamnetRouteNetworkService(opts ...option.RequestOption) (r *AccountTeamnetRouteNetworkService) {
	r = &AccountTeamnetRouteNetworkService{}
	r.Options = opts
	return
}

// Routes a private network through a Cloudflare Tunnel. The CIDR in
// `ip_network_encoded` must be written in URL-encoded format.
//
// Deprecated: This endpoint and its related APIs are deprecated in favor of the
// equivalent Tunnel Route (without CIDR) APIs.
func (r *AccountTeamnetRouteNetworkService) New(ctx context.Context, accountID string, ipNetworkEncoded string, body AccountTeamnetRouteNetworkNewParams, opts ...option.RequestOption) (res *TunnelRouteResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if ipNetworkEncoded == "" {
		err = errors.New("missing required ip_network_encoded parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/teamnet/routes/network/%s", accountID, ipNetworkEncoded)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Updates an existing private network route in an account. The CIDR in
// `ip_network_encoded` must be written in URL-encoded format.
//
// Deprecated: This endpoint and its related APIs are deprecated in favor of the
// equivalent Tunnel Route (without CIDR) APIs.
func (r *AccountTeamnetRouteNetworkService) Update(ctx context.Context, accountID string, ipNetworkEncoded string, opts ...option.RequestOption) (res *TunnelRouteResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if ipNetworkEncoded == "" {
		err = errors.New("missing required ip_network_encoded parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/teamnet/routes/network/%s", accountID, ipNetworkEncoded)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, nil, &res, opts...)
	return
}

// Deletes a private network route from an account. The CIDR in
// `ip_network_encoded` must be written in URL-encoded format. If no
// virtual_network_id is provided it will delete the route from the default vnet.
// If no tun_type is provided it will fetch the type from the tunnel_id or if that
// is missing it will assume Cloudflare Tunnel as default. If tunnel_id is provided
// it will delete the route from that tunnel, otherwise it will delete the route
// based on the vnet and tun_type.
//
// Deprecated: This endpoint and its related APIs are deprecated in favor of the
// equivalent Tunnel Route (without CIDR) APIs.
func (r *AccountTeamnetRouteNetworkService) Delete(ctx context.Context, accountID string, ipNetworkEncoded string, body AccountTeamnetRouteNetworkDeleteParams, opts ...option.RequestOption) (res *TunnelRouteResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if ipNetworkEncoded == "" {
		err = errors.New("missing required ip_network_encoded parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/teamnet/routes/network/%s", accountID, ipNetworkEncoded)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

// The type of tunnel.
type TunnelType string

const (
	TunnelTypeCfdTunnel     TunnelType = "cfd_tunnel"
	TunnelTypeWarpConnector TunnelType = "warp_connector"
	TunnelTypeWarp          TunnelType = "warp"
	TunnelTypeMagic         TunnelType = "magic"
	TunnelTypeIPSec         TunnelType = "ip_sec"
	TunnelTypeGre           TunnelType = "gre"
	TunnelTypeCni           TunnelType = "cni"
)

func (r TunnelType) IsKnown() bool {
	switch r {
	case TunnelTypeCfdTunnel, TunnelTypeWarpConnector, TunnelTypeWarp, TunnelTypeMagic, TunnelTypeIPSec, TunnelTypeGre, TunnelTypeCni:
		return true
	}
	return false
}

type AccountTeamnetRouteNetworkNewParams struct {
	// UUID of the tunnel.
	TunnelID param.Field[string] `json:"tunnel_id,required" format:"uuid"`
	// Optional remark describing the route.
	Comment param.Field[string] `json:"comment"`
	// UUID of the virtual network.
	VirtualNetworkID param.Field[string] `json:"virtual_network_id" format:"uuid"`
}

func (r AccountTeamnetRouteNetworkNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountTeamnetRouteNetworkDeleteParams struct {
	// The type of tunnel.
	TunType param.Field[TunnelType] `query:"tun_type"`
	// UUID of the tunnel.
	TunnelID param.Field[string] `query:"tunnel_id" format:"uuid"`
	// UUID of the virtual network.
	VirtualNetworkID param.Field[string] `query:"virtual_network_id" format:"uuid"`
}

// URLQuery serializes [AccountTeamnetRouteNetworkDeleteParams]'s query parameters
// as `url.Values`.
func (r AccountTeamnetRouteNetworkDeleteParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
