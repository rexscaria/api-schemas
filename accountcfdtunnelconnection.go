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

// AccountCfdTunnelConnectionService contains methods and other services that help
// with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountCfdTunnelConnectionService] method instead.
type AccountCfdTunnelConnectionService struct {
	Options []option.RequestOption
}

// NewAccountCfdTunnelConnectionService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountCfdTunnelConnectionService(opts ...option.RequestOption) (r *AccountCfdTunnelConnectionService) {
	r = &AccountCfdTunnelConnectionService{}
	r.Options = opts
	return
}

// Fetches connection details for a Cloudflare Tunnel.
func (r *AccountCfdTunnelConnectionService) List(ctx context.Context, accountID string, tunnelID string, opts ...option.RequestOption) (res *AccountCfdTunnelConnectionListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if tunnelID == "" {
		err = errors.New("missing required tunnel_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cfd_tunnel/%s/connections", accountID, tunnelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Removes a connection (aka Cloudflare Tunnel Connector) from a Cloudflare Tunnel
// independently of its current state. If no connector id (client_id) is provided
// all connectors will be removed. We recommend running this command after rotating
// tokens.
func (r *AccountCfdTunnelConnectionService) Cleanup(ctx context.Context, accountID string, tunnelID string, params AccountCfdTunnelConnectionCleanupParams, opts ...option.RequestOption) (res *AccountCfdTunnelConnectionCleanupResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if tunnelID == "" {
		err = errors.New("missing required tunnel_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cfd_tunnel/%s/connections", accountID, tunnelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, params, &res, opts...)
	return
}

type AccountCfdTunnelConnectionListResponse struct {
	Result []TunnelClient                             `json:"result"`
	JSON   accountCfdTunnelConnectionListResponseJSON `json:"-"`
	APIResponseCollectionTunnel
}

// accountCfdTunnelConnectionListResponseJSON contains the JSON metadata for the
// struct [AccountCfdTunnelConnectionListResponse]
type accountCfdTunnelConnectionListResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCfdTunnelConnectionListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountCfdTunnelConnectionListResponseJSON) RawJSON() string {
	return r.raw
}

type AccountCfdTunnelConnectionCleanupResponse struct {
	Result interface{}                                   `json:"result,nullable"`
	JSON   accountCfdTunnelConnectionCleanupResponseJSON `json:"-"`
	APIResponseTunnel
}

// accountCfdTunnelConnectionCleanupResponseJSON contains the JSON metadata for the
// struct [AccountCfdTunnelConnectionCleanupResponse]
type accountCfdTunnelConnectionCleanupResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCfdTunnelConnectionCleanupResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountCfdTunnelConnectionCleanupResponseJSON) RawJSON() string {
	return r.raw
}

type AccountCfdTunnelConnectionCleanupParams struct {
	Body interface{} `json:"body,required"`
	// UUID of the Cloudflare Tunnel connector.
	ClientID param.Field[string] `query:"client_id" format:"uuid"`
}

func (r AccountCfdTunnelConnectionCleanupParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

// URLQuery serializes [AccountCfdTunnelConnectionCleanupParams]'s query parameters
// as `url.Values`.
func (r AccountCfdTunnelConnectionCleanupParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
