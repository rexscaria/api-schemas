// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/stainless-sdks/cf-rex-go/internal/apijson"
	"github.com/stainless-sdks/cf-rex-go/internal/param"
	"github.com/stainless-sdks/cf-rex-go/internal/requestconfig"
	"github.com/stainless-sdks/cf-rex-go/option"
)

// AccountMagicIpsecTunnelService contains methods and other services that help
// with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountMagicIpsecTunnelService] method instead.
type AccountMagicIpsecTunnelService struct {
	Options []option.RequestOption
}

// NewAccountMagicIpsecTunnelService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountMagicIpsecTunnelService(opts ...option.RequestOption) (r *AccountMagicIpsecTunnelService) {
	r = &AccountMagicIpsecTunnelService{}
	r.Options = opts
	return
}

// Creates new IPsec tunnels associated with an account. Use `?validate_only=true`
// as an optional query parameter to only run validation without persisting
// changes.
func (r *AccountMagicIpsecTunnelService) New(ctx context.Context, accountID string, params AccountMagicIpsecTunnelNewParams, opts ...option.RequestOption) (res *MagicSchemasTunnelsCollectionResponse, err error) {
	if params.XMagicNewHcTarget.Present {
		opts = append(opts, option.WithHeader("x-magic-new-hc-target", fmt.Sprintf("%s", params.XMagicNewHcTarget)))
	}
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/ipsec_tunnels", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Lists details for a specific IPsec tunnel.
func (r *AccountMagicIpsecTunnelService) Get(ctx context.Context, accountID string, ipsecTunnelID string, query AccountMagicIpsecTunnelGetParams, opts ...option.RequestOption) (res *AccountMagicIpsecTunnelGetResponse, err error) {
	if query.XMagicNewHcTarget.Present {
		opts = append(opts, option.WithHeader("x-magic-new-hc-target", fmt.Sprintf("%s", query.XMagicNewHcTarget)))
	}
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if ipsecTunnelID == "" {
		err = errors.New("missing required ipsec_tunnel_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/ipsec_tunnels/%s", accountID, ipsecTunnelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates a specific IPsec tunnel associated with an account. Use
// `?validate_only=true` as an optional query parameter to only run validation
// without persisting changes.
func (r *AccountMagicIpsecTunnelService) Update(ctx context.Context, accountID string, ipsecTunnelID string, params AccountMagicIpsecTunnelUpdateParams, opts ...option.RequestOption) (res *AccountMagicIpsecTunnelUpdateResponse, err error) {
	if params.XMagicNewHcTarget.Present {
		opts = append(opts, option.WithHeader("x-magic-new-hc-target", fmt.Sprintf("%s", params.XMagicNewHcTarget)))
	}
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if ipsecTunnelID == "" {
		err = errors.New("missing required ipsec_tunnel_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/ipsec_tunnels/%s", accountID, ipsecTunnelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

// Lists IPsec tunnels associated with an account.
func (r *AccountMagicIpsecTunnelService) List(ctx context.Context, accountID string, query AccountMagicIpsecTunnelListParams, opts ...option.RequestOption) (res *MagicSchemasTunnelsCollectionResponse, err error) {
	if query.XMagicNewHcTarget.Present {
		opts = append(opts, option.WithHeader("x-magic-new-hc-target", fmt.Sprintf("%s", query.XMagicNewHcTarget)))
	}
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/ipsec_tunnels", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Disables and removes a specific static IPsec Tunnel associated with an account.
// Use `?validate_only=true` as an optional query parameter to only run validation
// without persisting changes.
func (r *AccountMagicIpsecTunnelService) Delete(ctx context.Context, accountID string, ipsecTunnelID string, params AccountMagicIpsecTunnelDeleteParams, opts ...option.RequestOption) (res *AccountMagicIpsecTunnelDeleteResponse, err error) {
	if params.XMagicNewHcTarget.Present {
		opts = append(opts, option.WithHeader("x-magic-new-hc-target", fmt.Sprintf("%s", params.XMagicNewHcTarget)))
	}
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if ipsecTunnelID == "" {
		err = errors.New("missing required ipsec_tunnel_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/ipsec_tunnels/%s", accountID, ipsecTunnelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, params, &res, opts...)
	return
}

// Generates a Pre Shared Key for a specific IPsec tunnel used in the IKE session.
// Use `?validate_only=true` as an optional query parameter to only run validation
// without persisting changes. After a PSK is generated, the PSK is immediately
// persisted to Cloudflare's edge and cannot be retrieved later. Note the PSK in a
// safe place.
func (r *AccountMagicIpsecTunnelService) GeneratePsk(ctx context.Context, accountID string, ipsecTunnelID string, body AccountMagicIpsecTunnelGeneratePskParams, opts ...option.RequestOption) (res *AccountMagicIpsecTunnelGeneratePskResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if ipsecTunnelID == "" {
		err = errors.New("missing required ipsec_tunnel_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/ipsec_tunnels/%s/psk_generate", accountID, ipsecTunnelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type MagicIpsecTunnel struct {
	// The IP address assigned to the Cloudflare side of the IPsec tunnel.
	CloudflareEndpoint string `json:"cloudflare_endpoint,required"`
	// A 31-bit prefix (/31 in CIDR notation) supporting two hosts, one for each side
	// of the tunnel. Select the subnet from the following private IP space:
	// 10.0.0.0–10.255.255.255, 172.16.0.0–172.31.255.255, 192.168.0.0–192.168.255.255.
	InterfaceAddress string `json:"interface_address,required"`
	// The name of the IPsec tunnel. The name cannot share a name with other tunnels.
	Name string `json:"name,required"`
	// Tunnel identifier tag.
	ID string `json:"id"`
	// When `true`, the tunnel can use a null-cipher (`ENCR_NULL`) in the ESP tunnel
	// (Phase 2).
	AllowNullCipher bool `json:"allow_null_cipher"`
	// The date and time the tunnel was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// The IP address assigned to the customer side of the IPsec tunnel. Not required,
	// but must be set for proactive traceroutes to work.
	CustomerEndpoint string `json:"customer_endpoint"`
	// An optional description forthe IPsec tunnel.
	Description string                 `json:"description"`
	HealthCheck MagicTunnelHealthCheck `json:"health_check"`
	// The date and time the tunnel was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// The PSK metadata that includes when the PSK was generated.
	PskMetadata MagicPskMetadata `json:"psk_metadata"`
	// If `true`, then IPsec replay protection will be supported in the
	// Cloudflare-to-customer direction.
	ReplayProtection bool                 `json:"replay_protection"`
	JSON             magicIpsecTunnelJSON `json:"-"`
}

// magicIpsecTunnelJSON contains the JSON metadata for the struct
// [MagicIpsecTunnel]
type magicIpsecTunnelJSON struct {
	CloudflareEndpoint apijson.Field
	InterfaceAddress   apijson.Field
	Name               apijson.Field
	ID                 apijson.Field
	AllowNullCipher    apijson.Field
	CreatedOn          apijson.Field
	CustomerEndpoint   apijson.Field
	Description        apijson.Field
	HealthCheck        apijson.Field
	ModifiedOn         apijson.Field
	PskMetadata        apijson.Field
	ReplayProtection   apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *MagicIpsecTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicIpsecTunnelJSON) RawJSON() string {
	return r.raw
}

type MagicIpsecTunnelAddSingleRequestParam struct {
	// The IP address assigned to the Cloudflare side of the IPsec tunnel.
	CloudflareEndpoint param.Field[string] `json:"cloudflare_endpoint,required"`
	// A 31-bit prefix (/31 in CIDR notation) supporting two hosts, one for each side
	// of the tunnel. Select the subnet from the following private IP space:
	// 10.0.0.0–10.255.255.255, 172.16.0.0–172.31.255.255, 192.168.0.0–192.168.255.255.
	InterfaceAddress param.Field[string] `json:"interface_address,required"`
	// The name of the IPsec tunnel. The name cannot share a name with other tunnels.
	Name param.Field[string] `json:"name,required"`
	// The IP address assigned to the customer side of the IPsec tunnel. Not required,
	// but must be set for proactive traceroutes to work.
	CustomerEndpoint param.Field[string] `json:"customer_endpoint"`
	// An optional description forthe IPsec tunnel.
	Description param.Field[string]                      `json:"description"`
	HealthCheck param.Field[MagicTunnelHealthCheckParam] `json:"health_check"`
	// A randomly generated or provided string for use in the IPsec tunnel.
	Psk param.Field[string] `json:"psk"`
	// If `true`, then IPsec replay protection will be supported in the
	// Cloudflare-to-customer direction.
	ReplayProtection param.Field[bool] `json:"replay_protection"`
}

func (r MagicIpsecTunnelAddSingleRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The PSK metadata that includes when the PSK was generated.
type MagicPskMetadata struct {
	// The date and time the tunnel was last modified.
	LastGeneratedOn time.Time            `json:"last_generated_on" format:"date-time"`
	JSON            magicPskMetadataJSON `json:"-"`
}

// magicPskMetadataJSON contains the JSON metadata for the struct
// [MagicPskMetadata]
type magicPskMetadataJSON struct {
	LastGeneratedOn apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *MagicPskMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicPskMetadataJSON) RawJSON() string {
	return r.raw
}

type MagicSchemasTunnelsCollectionResponse struct {
	Result MagicSchemasTunnelsCollectionResponseResult `json:"result"`
	JSON   magicSchemasTunnelsCollectionResponseJSON   `json:"-"`
	MagicAPIResponseSingle
}

// magicSchemasTunnelsCollectionResponseJSON contains the JSON metadata for the
// struct [MagicSchemasTunnelsCollectionResponse]
type magicSchemasTunnelsCollectionResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicSchemasTunnelsCollectionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicSchemasTunnelsCollectionResponseJSON) RawJSON() string {
	return r.raw
}

type MagicSchemasTunnelsCollectionResponseResult struct {
	IpsecTunnels []MagicIpsecTunnel                              `json:"ipsec_tunnels"`
	JSON         magicSchemasTunnelsCollectionResponseResultJSON `json:"-"`
}

// magicSchemasTunnelsCollectionResponseResultJSON contains the JSON metadata for
// the struct [MagicSchemasTunnelsCollectionResponseResult]
type magicSchemasTunnelsCollectionResponseResultJSON struct {
	IpsecTunnels apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *MagicSchemasTunnelsCollectionResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicSchemasTunnelsCollectionResponseResultJSON) RawJSON() string {
	return r.raw
}

type MagicTunnelHealthCheck struct {
	// The direction of the flow of the healthcheck. Either unidirectional, where the
	// probe comes to you via the tunnel and the result comes back to Cloudflare via
	// the open Internet, or bidirectional where both the probe and result come and go
	// via the tunnel.
	Direction MagicTunnelHealthCheckDirection `json:"direction"`
	JSON      magicTunnelHealthCheckJSON      `json:"-"`
	MagicHealthCheckBase
}

// magicTunnelHealthCheckJSON contains the JSON metadata for the struct
// [MagicTunnelHealthCheck]
type magicTunnelHealthCheckJSON struct {
	Direction   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicTunnelHealthCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicTunnelHealthCheckJSON) RawJSON() string {
	return r.raw
}

// The direction of the flow of the healthcheck. Either unidirectional, where the
// probe comes to you via the tunnel and the result comes back to Cloudflare via
// the open Internet, or bidirectional where both the probe and result come and go
// via the tunnel.
type MagicTunnelHealthCheckDirection string

const (
	MagicTunnelHealthCheckDirectionUnidirectional MagicTunnelHealthCheckDirection = "unidirectional"
	MagicTunnelHealthCheckDirectionBidirectional  MagicTunnelHealthCheckDirection = "bidirectional"
)

func (r MagicTunnelHealthCheckDirection) IsKnown() bool {
	switch r {
	case MagicTunnelHealthCheckDirectionUnidirectional, MagicTunnelHealthCheckDirectionBidirectional:
		return true
	}
	return false
}

type MagicTunnelHealthCheckParam struct {
	// The direction of the flow of the healthcheck. Either unidirectional, where the
	// probe comes to you via the tunnel and the result comes back to Cloudflare via
	// the open Internet, or bidirectional where both the probe and result come and go
	// via the tunnel.
	Direction param.Field[MagicTunnelHealthCheckDirection] `json:"direction"`
	MagicHealthCheckBaseParam
}

func (r MagicTunnelHealthCheckParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountMagicIpsecTunnelGetResponse struct {
	Result AccountMagicIpsecTunnelGetResponseResult `json:"result"`
	JSON   accountMagicIpsecTunnelGetResponseJSON   `json:"-"`
	MagicAPIResponseSingle
}

// accountMagicIpsecTunnelGetResponseJSON contains the JSON metadata for the struct
// [AccountMagicIpsecTunnelGetResponse]
type accountMagicIpsecTunnelGetResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicIpsecTunnelGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMagicIpsecTunnelGetResponseJSON) RawJSON() string {
	return r.raw
}

type AccountMagicIpsecTunnelGetResponseResult struct {
	IpsecTunnel MagicIpsecTunnel                             `json:"ipsec_tunnel"`
	JSON        accountMagicIpsecTunnelGetResponseResultJSON `json:"-"`
}

// accountMagicIpsecTunnelGetResponseResultJSON contains the JSON metadata for the
// struct [AccountMagicIpsecTunnelGetResponseResult]
type accountMagicIpsecTunnelGetResponseResultJSON struct {
	IpsecTunnel apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicIpsecTunnelGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMagicIpsecTunnelGetResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountMagicIpsecTunnelUpdateResponse struct {
	Result AccountMagicIpsecTunnelUpdateResponseResult `json:"result"`
	JSON   accountMagicIpsecTunnelUpdateResponseJSON   `json:"-"`
	MagicAPIResponseSingle
}

// accountMagicIpsecTunnelUpdateResponseJSON contains the JSON metadata for the
// struct [AccountMagicIpsecTunnelUpdateResponse]
type accountMagicIpsecTunnelUpdateResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicIpsecTunnelUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMagicIpsecTunnelUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type AccountMagicIpsecTunnelUpdateResponseResult struct {
	Modified            bool                                            `json:"modified"`
	ModifiedIpsecTunnel MagicIpsecTunnel                                `json:"modified_ipsec_tunnel"`
	JSON                accountMagicIpsecTunnelUpdateResponseResultJSON `json:"-"`
}

// accountMagicIpsecTunnelUpdateResponseResultJSON contains the JSON metadata for
// the struct [AccountMagicIpsecTunnelUpdateResponseResult]
type accountMagicIpsecTunnelUpdateResponseResultJSON struct {
	Modified            apijson.Field
	ModifiedIpsecTunnel apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *AccountMagicIpsecTunnelUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMagicIpsecTunnelUpdateResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountMagicIpsecTunnelDeleteResponse struct {
	Result AccountMagicIpsecTunnelDeleteResponseResult `json:"result"`
	JSON   accountMagicIpsecTunnelDeleteResponseJSON   `json:"-"`
	MagicAPIResponseSingle
}

// accountMagicIpsecTunnelDeleteResponseJSON contains the JSON metadata for the
// struct [AccountMagicIpsecTunnelDeleteResponse]
type accountMagicIpsecTunnelDeleteResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicIpsecTunnelDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMagicIpsecTunnelDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type AccountMagicIpsecTunnelDeleteResponseResult struct {
	Deleted            bool                                            `json:"deleted"`
	DeletedIpsecTunnel MagicIpsecTunnel                                `json:"deleted_ipsec_tunnel"`
	JSON               accountMagicIpsecTunnelDeleteResponseResultJSON `json:"-"`
}

// accountMagicIpsecTunnelDeleteResponseResultJSON contains the JSON metadata for
// the struct [AccountMagicIpsecTunnelDeleteResponseResult]
type accountMagicIpsecTunnelDeleteResponseResultJSON struct {
	Deleted            apijson.Field
	DeletedIpsecTunnel apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccountMagicIpsecTunnelDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMagicIpsecTunnelDeleteResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountMagicIpsecTunnelGeneratePskResponse struct {
	Result AccountMagicIpsecTunnelGeneratePskResponseResult `json:"result"`
	JSON   accountMagicIpsecTunnelGeneratePskResponseJSON   `json:"-"`
	MagicAPIResponseSingle
}

// accountMagicIpsecTunnelGeneratePskResponseJSON contains the JSON metadata for
// the struct [AccountMagicIpsecTunnelGeneratePskResponse]
type accountMagicIpsecTunnelGeneratePskResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicIpsecTunnelGeneratePskResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMagicIpsecTunnelGeneratePskResponseJSON) RawJSON() string {
	return r.raw
}

type AccountMagicIpsecTunnelGeneratePskResponseResult struct {
	// Identifier
	IpsecTunnelID string `json:"ipsec_tunnel_id"`
	// A randomly generated or provided string for use in the IPsec tunnel.
	Psk string `json:"psk"`
	// The PSK metadata that includes when the PSK was generated.
	PskMetadata MagicPskMetadata                                     `json:"psk_metadata"`
	JSON        accountMagicIpsecTunnelGeneratePskResponseResultJSON `json:"-"`
}

// accountMagicIpsecTunnelGeneratePskResponseResultJSON contains the JSON metadata
// for the struct [AccountMagicIpsecTunnelGeneratePskResponseResult]
type accountMagicIpsecTunnelGeneratePskResponseResultJSON struct {
	IpsecTunnelID apijson.Field
	Psk           apijson.Field
	PskMetadata   apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccountMagicIpsecTunnelGeneratePskResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMagicIpsecTunnelGeneratePskResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountMagicIpsecTunnelNewParams struct {
	MagicIpsecTunnelAddSingleRequest MagicIpsecTunnelAddSingleRequestParam `json:"magic_ipsec_tunnel_add_single_request,required"`
	XMagicNewHcTarget                param.Field[bool]                     `header:"x-magic-new-hc-target"`
}

func (r AccountMagicIpsecTunnelNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.MagicIpsecTunnelAddSingleRequest)
}

type AccountMagicIpsecTunnelGetParams struct {
	XMagicNewHcTarget param.Field[bool] `header:"x-magic-new-hc-target"`
}

type AccountMagicIpsecTunnelUpdateParams struct {
	MagicIpsecTunnelAddSingleRequest MagicIpsecTunnelAddSingleRequestParam `json:"magic_ipsec_tunnel_add_single_request,required"`
	XMagicNewHcTarget                param.Field[bool]                     `header:"x-magic-new-hc-target"`
}

func (r AccountMagicIpsecTunnelUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.MagicIpsecTunnelAddSingleRequest)
}

type AccountMagicIpsecTunnelListParams struct {
	XMagicNewHcTarget param.Field[bool] `header:"x-magic-new-hc-target"`
}

type AccountMagicIpsecTunnelDeleteParams struct {
	Body              interface{}       `json:"body,required"`
	XMagicNewHcTarget param.Field[bool] `header:"x-magic-new-hc-target"`
}

func (r AccountMagicIpsecTunnelDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type AccountMagicIpsecTunnelGeneratePskParams struct {
	Body interface{} `json:"body,required"`
}

func (r AccountMagicIpsecTunnelGeneratePskParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}
