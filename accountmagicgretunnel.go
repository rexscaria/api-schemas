// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
)

// AccountMagicGreTunnelService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountMagicGreTunnelService] method instead.
type AccountMagicGreTunnelService struct {
	Options []option.RequestOption
}

// NewAccountMagicGreTunnelService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountMagicGreTunnelService(opts ...option.RequestOption) (r *AccountMagicGreTunnelService) {
	r = &AccountMagicGreTunnelService{}
	r.Options = opts
	return
}

// Creates new GRE tunnels. Use `?validate_only=true` as an optional query
// parameter to only run validation without persisting changes.
func (r *AccountMagicGreTunnelService) New(ctx context.Context, accountID string, params AccountMagicGreTunnelNewParams, opts ...option.RequestOption) (res *MagicTunnelsCollectionResponse, err error) {
	if params.XMagicNewHcTarget.Present {
		opts = append(opts, option.WithHeader("x-magic-new-hc-target", fmt.Sprintf("%s", params.XMagicNewHcTarget)))
	}
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/gre_tunnels", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Lists informtion for a specific GRE tunnel.
func (r *AccountMagicGreTunnelService) Get(ctx context.Context, accountID string, greTunnelID string, query AccountMagicGreTunnelGetParams, opts ...option.RequestOption) (res *AccountMagicGreTunnelGetResponse, err error) {
	if query.XMagicNewHcTarget.Present {
		opts = append(opts, option.WithHeader("x-magic-new-hc-target", fmt.Sprintf("%s", query.XMagicNewHcTarget)))
	}
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if greTunnelID == "" {
		err = errors.New("missing required gre_tunnel_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/gre_tunnels/%s", accountID, greTunnelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates a specific GRE tunnel. Use `?validate_only=true` as an optional query
// parameter to only run validation without persisting changes.
func (r *AccountMagicGreTunnelService) Update(ctx context.Context, accountID string, greTunnelID string, params AccountMagicGreTunnelUpdateParams, opts ...option.RequestOption) (res *AccountMagicGreTunnelUpdateResponse, err error) {
	if params.XMagicNewHcTarget.Present {
		opts = append(opts, option.WithHeader("x-magic-new-hc-target", fmt.Sprintf("%s", params.XMagicNewHcTarget)))
	}
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if greTunnelID == "" {
		err = errors.New("missing required gre_tunnel_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/gre_tunnels/%s", accountID, greTunnelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

// Lists GRE tunnels associated with an account.
func (r *AccountMagicGreTunnelService) List(ctx context.Context, accountID string, query AccountMagicGreTunnelListParams, opts ...option.RequestOption) (res *MagicTunnelsCollectionResponse, err error) {
	if query.XMagicNewHcTarget.Present {
		opts = append(opts, option.WithHeader("x-magic-new-hc-target", fmt.Sprintf("%s", query.XMagicNewHcTarget)))
	}
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/gre_tunnels", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Disables and removes a specific static GRE tunnel. Use `?validate_only=true` as
// an optional query parameter to only run validation without persisting changes.
func (r *AccountMagicGreTunnelService) Delete(ctx context.Context, accountID string, greTunnelID string, params AccountMagicGreTunnelDeleteParams, opts ...option.RequestOption) (res *AccountMagicGreTunnelDeleteResponse, err error) {
	if params.XMagicNewHcTarget.Present {
		opts = append(opts, option.WithHeader("x-magic-new-hc-target", fmt.Sprintf("%s", params.XMagicNewHcTarget)))
	}
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if greTunnelID == "" {
		err = errors.New("missing required gre_tunnel_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/gre_tunnels/%s", accountID, greTunnelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, params, &res, opts...)
	return
}

type MagicGreTunnel struct {
	// The IP address assigned to the Cloudflare side of the GRE tunnel.
	CloudflareGreEndpoint string `json:"cloudflare_gre_endpoint,required"`
	// The IP address assigned to the customer side of the GRE tunnel.
	CustomerGreEndpoint string `json:"customer_gre_endpoint,required"`
	// A 31-bit prefix (/31 in CIDR notation) supporting two hosts, one for each side
	// of the tunnel. Select the subnet from the following private IP space:
	// 10.0.0.0–10.255.255.255, 172.16.0.0–172.31.255.255, 192.168.0.0–192.168.255.255.
	InterfaceAddress string `json:"interface_address,required"`
	// The name of the tunnel. The name cannot contain spaces or special characters,
	// must be 15 characters or less, and cannot share a name with another GRE tunnel.
	Name string `json:"name,required"`
	// Tunnel identifier tag.
	ID string `json:"id"`
	// The date and time the tunnel was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// An optional description of the GRE tunnel.
	Description string                 `json:"description"`
	HealthCheck MagicTunnelHealthCheck `json:"health_check"`
	// The date and time the tunnel was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Maximum Transmission Unit (MTU) in bytes for the GRE tunnel. The minimum value
	// is 576.
	Mtu int64 `json:"mtu"`
	// Time To Live (TTL) in number of hops of the GRE tunnel.
	Ttl  int64              `json:"ttl"`
	JSON magicGreTunnelJSON `json:"-"`
}

// magicGreTunnelJSON contains the JSON metadata for the struct [MagicGreTunnel]
type magicGreTunnelJSON struct {
	CloudflareGreEndpoint apijson.Field
	CustomerGreEndpoint   apijson.Field
	InterfaceAddress      apijson.Field
	Name                  apijson.Field
	ID                    apijson.Field
	CreatedOn             apijson.Field
	Description           apijson.Field
	HealthCheck           apijson.Field
	ModifiedOn            apijson.Field
	Mtu                   apijson.Field
	Ttl                   apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *MagicGreTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicGreTunnelJSON) RawJSON() string {
	return r.raw
}

type MagicTunnelsCollectionResponse struct {
	Result MagicTunnelsCollectionResponseResult `json:"result"`
	JSON   magicTunnelsCollectionResponseJSON   `json:"-"`
	MagicAPIResponseSingle
}

// magicTunnelsCollectionResponseJSON contains the JSON metadata for the struct
// [MagicTunnelsCollectionResponse]
type magicTunnelsCollectionResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicTunnelsCollectionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicTunnelsCollectionResponseJSON) RawJSON() string {
	return r.raw
}

type MagicTunnelsCollectionResponseResult struct {
	GreTunnels []MagicGreTunnel                         `json:"gre_tunnels"`
	JSON       magicTunnelsCollectionResponseResultJSON `json:"-"`
}

// magicTunnelsCollectionResponseResultJSON contains the JSON metadata for the
// struct [MagicTunnelsCollectionResponseResult]
type magicTunnelsCollectionResponseResultJSON struct {
	GreTunnels  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicTunnelsCollectionResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicTunnelsCollectionResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountMagicGreTunnelGetResponse struct {
	Result AccountMagicGreTunnelGetResponseResult `json:"result"`
	JSON   accountMagicGreTunnelGetResponseJSON   `json:"-"`
	MagicAPIResponseSingle
}

// accountMagicGreTunnelGetResponseJSON contains the JSON metadata for the struct
// [AccountMagicGreTunnelGetResponse]
type accountMagicGreTunnelGetResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicGreTunnelGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMagicGreTunnelGetResponseJSON) RawJSON() string {
	return r.raw
}

type AccountMagicGreTunnelGetResponseResult struct {
	GreTunnel MagicGreTunnel                             `json:"gre_tunnel"`
	JSON      accountMagicGreTunnelGetResponseResultJSON `json:"-"`
}

// accountMagicGreTunnelGetResponseResultJSON contains the JSON metadata for the
// struct [AccountMagicGreTunnelGetResponseResult]
type accountMagicGreTunnelGetResponseResultJSON struct {
	GreTunnel   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicGreTunnelGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMagicGreTunnelGetResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountMagicGreTunnelUpdateResponse struct {
	Result AccountMagicGreTunnelUpdateResponseResult `json:"result"`
	JSON   accountMagicGreTunnelUpdateResponseJSON   `json:"-"`
	MagicAPIResponseSingle
}

// accountMagicGreTunnelUpdateResponseJSON contains the JSON metadata for the
// struct [AccountMagicGreTunnelUpdateResponse]
type accountMagicGreTunnelUpdateResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicGreTunnelUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMagicGreTunnelUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type AccountMagicGreTunnelUpdateResponseResult struct {
	Modified          bool                                          `json:"modified"`
	ModifiedGreTunnel MagicGreTunnel                                `json:"modified_gre_tunnel"`
	JSON              accountMagicGreTunnelUpdateResponseResultJSON `json:"-"`
}

// accountMagicGreTunnelUpdateResponseResultJSON contains the JSON metadata for the
// struct [AccountMagicGreTunnelUpdateResponseResult]
type accountMagicGreTunnelUpdateResponseResultJSON struct {
	Modified          apijson.Field
	ModifiedGreTunnel apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *AccountMagicGreTunnelUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMagicGreTunnelUpdateResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountMagicGreTunnelDeleteResponse struct {
	Result AccountMagicGreTunnelDeleteResponseResult `json:"result"`
	JSON   accountMagicGreTunnelDeleteResponseJSON   `json:"-"`
	MagicAPIResponseSingle
}

// accountMagicGreTunnelDeleteResponseJSON contains the JSON metadata for the
// struct [AccountMagicGreTunnelDeleteResponse]
type accountMagicGreTunnelDeleteResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicGreTunnelDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMagicGreTunnelDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type AccountMagicGreTunnelDeleteResponseResult struct {
	Deleted          bool                                          `json:"deleted"`
	DeletedGreTunnel MagicGreTunnel                                `json:"deleted_gre_tunnel"`
	JSON             accountMagicGreTunnelDeleteResponseResultJSON `json:"-"`
}

// accountMagicGreTunnelDeleteResponseResultJSON contains the JSON metadata for the
// struct [AccountMagicGreTunnelDeleteResponseResult]
type accountMagicGreTunnelDeleteResponseResultJSON struct {
	Deleted          apijson.Field
	DeletedGreTunnel apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountMagicGreTunnelDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMagicGreTunnelDeleteResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountMagicGreTunnelNewParams struct {
	Body              interface{}       `json:"body,required"`
	XMagicNewHcTarget param.Field[bool] `header:"x-magic-new-hc-target"`
}

func (r AccountMagicGreTunnelNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type AccountMagicGreTunnelGetParams struct {
	XMagicNewHcTarget param.Field[bool] `header:"x-magic-new-hc-target"`
}

type AccountMagicGreTunnelUpdateParams struct {
	// The IP address assigned to the Cloudflare side of the GRE tunnel.
	CloudflareGreEndpoint param.Field[string] `json:"cloudflare_gre_endpoint,required"`
	// The IP address assigned to the customer side of the GRE tunnel.
	CustomerGreEndpoint param.Field[string] `json:"customer_gre_endpoint,required"`
	// A 31-bit prefix (/31 in CIDR notation) supporting two hosts, one for each side
	// of the tunnel. Select the subnet from the following private IP space:
	// 10.0.0.0–10.255.255.255, 172.16.0.0–172.31.255.255, 192.168.0.0–192.168.255.255.
	InterfaceAddress param.Field[string] `json:"interface_address,required"`
	// The name of the tunnel. The name cannot contain spaces or special characters,
	// must be 15 characters or less, and cannot share a name with another GRE tunnel.
	Name param.Field[string] `json:"name,required"`
	// An optional description of the GRE tunnel.
	Description param.Field[string]                      `json:"description"`
	HealthCheck param.Field[MagicTunnelHealthCheckParam] `json:"health_check"`
	// Maximum Transmission Unit (MTU) in bytes for the GRE tunnel. The minimum value
	// is 576.
	Mtu param.Field[int64] `json:"mtu"`
	// Time To Live (TTL) in number of hops of the GRE tunnel.
	Ttl               param.Field[int64] `json:"ttl"`
	XMagicNewHcTarget param.Field[bool]  `header:"x-magic-new-hc-target"`
}

func (r AccountMagicGreTunnelUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountMagicGreTunnelListParams struct {
	XMagicNewHcTarget param.Field[bool] `header:"x-magic-new-hc-target"`
}

type AccountMagicGreTunnelDeleteParams struct {
	Body              interface{}       `json:"body,required"`
	XMagicNewHcTarget param.Field[bool] `header:"x-magic-new-hc-target"`
}

func (r AccountMagicGreTunnelDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}
