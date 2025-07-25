// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
	"github.com/rexscaria/api-schemas/shared"
	"github.com/tidwall/gjson"
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

// Creates a new IPsec tunnel associated with an account. Use `?validate_only=true`
// as an optional query parameter to only run validation without persisting
// changes.
func (r *AccountMagicIpsecTunnelService) New(ctx context.Context, accountID string, params AccountMagicIpsecTunnelNewParams, opts ...option.RequestOption) (res *AccountMagicIpsecTunnelNewResponse, err error) {
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
func (r *AccountMagicIpsecTunnelService) Delete(ctx context.Context, accountID string, ipsecTunnelID string, body AccountMagicIpsecTunnelDeleteParams, opts ...option.RequestOption) (res *AccountMagicIpsecTunnelDeleteResponse, err error) {
	if body.XMagicNewHcTarget.Present {
		opts = append(opts, option.WithHeader("x-magic-new-hc-target", fmt.Sprintf("%s", body.XMagicNewHcTarget)))
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
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
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
	// Identifier
	ID string `json:"id,required"`
	// The IP address assigned to the Cloudflare side of the IPsec tunnel.
	CloudflareEndpoint string `json:"cloudflare_endpoint,required"`
	// A 31-bit prefix (/31 in CIDR notation) supporting two hosts, one for each side
	// of the tunnel. Select the subnet from the following private IP space:
	// 10.0.0.0–10.255.255.255, 172.16.0.0–172.31.255.255, 192.168.0.0–192.168.255.255.
	InterfaceAddress string `json:"interface_address,required"`
	// The name of the IPsec tunnel. The name cannot share a name with other tunnels.
	Name string `json:"name,required"`
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
	ID                 apijson.Field
	CloudflareEndpoint apijson.Field
	InterfaceAddress   apijson.Field
	Name               apijson.Field
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
	Errors   []MagicMessageItem                          `json:"errors,required"`
	Messages []MagicMessageItem                          `json:"messages,required"`
	Result   MagicSchemasTunnelsCollectionResponseResult `json:"result,required"`
	// Whether the API call was successful
	Success MagicSchemasTunnelsCollectionResponseSuccess `json:"success,required"`
	JSON    magicSchemasTunnelsCollectionResponseJSON    `json:"-"`
}

// magicSchemasTunnelsCollectionResponseJSON contains the JSON metadata for the
// struct [MagicSchemasTunnelsCollectionResponse]
type magicSchemasTunnelsCollectionResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
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

// Whether the API call was successful
type MagicSchemasTunnelsCollectionResponseSuccess bool

const (
	MagicSchemasTunnelsCollectionResponseSuccessTrue MagicSchemasTunnelsCollectionResponseSuccess = true
)

func (r MagicSchemasTunnelsCollectionResponseSuccess) IsKnown() bool {
	switch r {
	case MagicSchemasTunnelsCollectionResponseSuccessTrue:
		return true
	}
	return false
}

type MagicTunnelHealthCheck struct {
	// The direction of the flow of the healthcheck. Either unidirectional, where the
	// probe comes to you via the tunnel and the result comes back to Cloudflare via
	// the open Internet, or bidirectional where both the probe and result come and go
	// via the tunnel.
	Direction MagicTunnelHealthCheckDirection `json:"direction"`
	// Determines whether to run healthchecks for a tunnel.
	Enabled bool `json:"enabled"`
	// How frequent the health check is run. The default value is `mid`.
	Rate MagicTunnelHealthCheckRate `json:"rate"`
	// The destination address in a request type health check. After the healthcheck is
	// decapsulated at the customer end of the tunnel, the ICMP echo will be forwarded
	// to this address. This field defaults to `customer_gre_endpoint address`. This
	// field is ignored for bidirectional healthchecks as the interface_address (not
	// assigned to the Cloudflare side of the tunnel) is used as the target. Must be in
	// object form if the x-magic-new-hc-target header is set to true and string form
	// if x-magic-new-hc-target is absent or set to false.
	Target MagicTunnelHealthCheckTargetUnion `json:"target"`
	// The type of healthcheck to run, reply or request. The default value is `reply`.
	Type MagicTunnelHealthCheckType `json:"type"`
	JSON magicTunnelHealthCheckJSON `json:"-"`
}

// magicTunnelHealthCheckJSON contains the JSON metadata for the struct
// [MagicTunnelHealthCheck]
type magicTunnelHealthCheckJSON struct {
	Direction   apijson.Field
	Enabled     apijson.Field
	Rate        apijson.Field
	Target      apijson.Field
	Type        apijson.Field
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

// How frequent the health check is run. The default value is `mid`.
type MagicTunnelHealthCheckRate string

const (
	MagicTunnelHealthCheckRateLow  MagicTunnelHealthCheckRate = "low"
	MagicTunnelHealthCheckRateMid  MagicTunnelHealthCheckRate = "mid"
	MagicTunnelHealthCheckRateHigh MagicTunnelHealthCheckRate = "high"
)

func (r MagicTunnelHealthCheckRate) IsKnown() bool {
	switch r {
	case MagicTunnelHealthCheckRateLow, MagicTunnelHealthCheckRateMid, MagicTunnelHealthCheckRateHigh:
		return true
	}
	return false
}

// The destination address in a request type health check. After the healthcheck is
// decapsulated at the customer end of the tunnel, the ICMP echo will be forwarded
// to this address. This field defaults to `customer_gre_endpoint address`. This
// field is ignored for bidirectional healthchecks as the interface_address (not
// assigned to the Cloudflare side of the tunnel) is used as the target. Must be in
// object form if the x-magic-new-hc-target header is set to true and string form
// if x-magic-new-hc-target is absent or set to false.
//
// Union satisfied by [MagicTunnelHealthCheckTargetMagicHealthCheckTarget] or
// [shared.UnionString].
type MagicTunnelHealthCheckTargetUnion interface {
	ImplementsMagicTunnelHealthCheckTargetUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*MagicTunnelHealthCheckTargetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(MagicTunnelHealthCheckTargetMagicHealthCheckTarget{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// The destination address in a request type health check. After the healthcheck is
// decapsulated at the customer end of the tunnel, the ICMP echo will be forwarded
// to this address. This field defaults to `customer_gre_endpoint address`. This
// field is ignored for bidirectional healthchecks as the interface_address (not
// assigned to the Cloudflare side of the tunnel) is used as the target.
type MagicTunnelHealthCheckTargetMagicHealthCheckTarget struct {
	// The effective health check target. If 'saved' is empty, then this field will be
	// populated with the calculated default value on GET requests. Ignored in POST,
	// PUT, and PATCH requests.
	Effective string `json:"effective"`
	// The saved health check target. Setting the value to the empty string indicates
	// that the calculated default value will be used.
	Saved string                                                 `json:"saved"`
	JSON  magicTunnelHealthCheckTargetMagicHealthCheckTargetJSON `json:"-"`
}

// magicTunnelHealthCheckTargetMagicHealthCheckTargetJSON contains the JSON
// metadata for the struct [MagicTunnelHealthCheckTargetMagicHealthCheckTarget]
type magicTunnelHealthCheckTargetMagicHealthCheckTargetJSON struct {
	Effective   apijson.Field
	Saved       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicTunnelHealthCheckTargetMagicHealthCheckTarget) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicTunnelHealthCheckTargetMagicHealthCheckTargetJSON) RawJSON() string {
	return r.raw
}

func (r MagicTunnelHealthCheckTargetMagicHealthCheckTarget) ImplementsMagicTunnelHealthCheckTargetUnion() {
}

// The type of healthcheck to run, reply or request. The default value is `reply`.
type MagicTunnelHealthCheckType string

const (
	MagicTunnelHealthCheckTypeReply   MagicTunnelHealthCheckType = "reply"
	MagicTunnelHealthCheckTypeRequest MagicTunnelHealthCheckType = "request"
)

func (r MagicTunnelHealthCheckType) IsKnown() bool {
	switch r {
	case MagicTunnelHealthCheckTypeReply, MagicTunnelHealthCheckTypeRequest:
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
	// Determines whether to run healthchecks for a tunnel.
	Enabled param.Field[bool] `json:"enabled"`
	// How frequent the health check is run. The default value is `mid`.
	Rate param.Field[MagicTunnelHealthCheckRate] `json:"rate"`
	// The destination address in a request type health check. After the healthcheck is
	// decapsulated at the customer end of the tunnel, the ICMP echo will be forwarded
	// to this address. This field defaults to `customer_gre_endpoint address`. This
	// field is ignored for bidirectional healthchecks as the interface_address (not
	// assigned to the Cloudflare side of the tunnel) is used as the target. Must be in
	// object form if the x-magic-new-hc-target header is set to true and string form
	// if x-magic-new-hc-target is absent or set to false.
	Target param.Field[MagicTunnelHealthCheckTargetUnionParam] `json:"target"`
	// The type of healthcheck to run, reply or request. The default value is `reply`.
	Type param.Field[MagicTunnelHealthCheckType] `json:"type"`
}

func (r MagicTunnelHealthCheckParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The destination address in a request type health check. After the healthcheck is
// decapsulated at the customer end of the tunnel, the ICMP echo will be forwarded
// to this address. This field defaults to `customer_gre_endpoint address`. This
// field is ignored for bidirectional healthchecks as the interface_address (not
// assigned to the Cloudflare side of the tunnel) is used as the target. Must be in
// object form if the x-magic-new-hc-target header is set to true and string form
// if x-magic-new-hc-target is absent or set to false.
//
// Satisfied by [MagicTunnelHealthCheckTargetMagicHealthCheckTargetParam],
// [shared.UnionString].
type MagicTunnelHealthCheckTargetUnionParam interface {
	ImplementsMagicTunnelHealthCheckTargetUnionParam()
}

// The destination address in a request type health check. After the healthcheck is
// decapsulated at the customer end of the tunnel, the ICMP echo will be forwarded
// to this address. This field defaults to `customer_gre_endpoint address`. This
// field is ignored for bidirectional healthchecks as the interface_address (not
// assigned to the Cloudflare side of the tunnel) is used as the target.
type MagicTunnelHealthCheckTargetMagicHealthCheckTargetParam struct {
	// The saved health check target. Setting the value to the empty string indicates
	// that the calculated default value will be used.
	Saved param.Field[string] `json:"saved"`
}

func (r MagicTunnelHealthCheckTargetMagicHealthCheckTargetParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MagicTunnelHealthCheckTargetMagicHealthCheckTargetParam) ImplementsMagicTunnelHealthCheckTargetUnionParam() {
}

type AccountMagicIpsecTunnelNewResponse struct {
	Errors   []MagicMessageItem `json:"errors,required"`
	Messages []MagicMessageItem `json:"messages,required"`
	Result   MagicIpsecTunnel   `json:"result,required"`
	// Whether the API call was successful
	Success AccountMagicIpsecTunnelNewResponseSuccess `json:"success,required"`
	JSON    accountMagicIpsecTunnelNewResponseJSON    `json:"-"`
}

// accountMagicIpsecTunnelNewResponseJSON contains the JSON metadata for the struct
// [AccountMagicIpsecTunnelNewResponse]
type accountMagicIpsecTunnelNewResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicIpsecTunnelNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMagicIpsecTunnelNewResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AccountMagicIpsecTunnelNewResponseSuccess bool

const (
	AccountMagicIpsecTunnelNewResponseSuccessTrue AccountMagicIpsecTunnelNewResponseSuccess = true
)

func (r AccountMagicIpsecTunnelNewResponseSuccess) IsKnown() bool {
	switch r {
	case AccountMagicIpsecTunnelNewResponseSuccessTrue:
		return true
	}
	return false
}

type AccountMagicIpsecTunnelGetResponse struct {
	Errors   []MagicMessageItem                       `json:"errors,required"`
	Messages []MagicMessageItem                       `json:"messages,required"`
	Result   AccountMagicIpsecTunnelGetResponseResult `json:"result,required"`
	// Whether the API call was successful
	Success AccountMagicIpsecTunnelGetResponseSuccess `json:"success,required"`
	JSON    accountMagicIpsecTunnelGetResponseJSON    `json:"-"`
}

// accountMagicIpsecTunnelGetResponseJSON contains the JSON metadata for the struct
// [AccountMagicIpsecTunnelGetResponse]
type accountMagicIpsecTunnelGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
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

// Whether the API call was successful
type AccountMagicIpsecTunnelGetResponseSuccess bool

const (
	AccountMagicIpsecTunnelGetResponseSuccessTrue AccountMagicIpsecTunnelGetResponseSuccess = true
)

func (r AccountMagicIpsecTunnelGetResponseSuccess) IsKnown() bool {
	switch r {
	case AccountMagicIpsecTunnelGetResponseSuccessTrue:
		return true
	}
	return false
}

type AccountMagicIpsecTunnelUpdateResponse struct {
	Errors   []MagicMessageItem                          `json:"errors,required"`
	Messages []MagicMessageItem                          `json:"messages,required"`
	Result   AccountMagicIpsecTunnelUpdateResponseResult `json:"result,required"`
	// Whether the API call was successful
	Success AccountMagicIpsecTunnelUpdateResponseSuccess `json:"success,required"`
	JSON    accountMagicIpsecTunnelUpdateResponseJSON    `json:"-"`
}

// accountMagicIpsecTunnelUpdateResponseJSON contains the JSON metadata for the
// struct [AccountMagicIpsecTunnelUpdateResponse]
type accountMagicIpsecTunnelUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
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

// Whether the API call was successful
type AccountMagicIpsecTunnelUpdateResponseSuccess bool

const (
	AccountMagicIpsecTunnelUpdateResponseSuccessTrue AccountMagicIpsecTunnelUpdateResponseSuccess = true
)

func (r AccountMagicIpsecTunnelUpdateResponseSuccess) IsKnown() bool {
	switch r {
	case AccountMagicIpsecTunnelUpdateResponseSuccessTrue:
		return true
	}
	return false
}

type AccountMagicIpsecTunnelDeleteResponse struct {
	Errors   []MagicMessageItem                          `json:"errors,required"`
	Messages []MagicMessageItem                          `json:"messages,required"`
	Result   AccountMagicIpsecTunnelDeleteResponseResult `json:"result,required"`
	// Whether the API call was successful
	Success AccountMagicIpsecTunnelDeleteResponseSuccess `json:"success,required"`
	JSON    accountMagicIpsecTunnelDeleteResponseJSON    `json:"-"`
}

// accountMagicIpsecTunnelDeleteResponseJSON contains the JSON metadata for the
// struct [AccountMagicIpsecTunnelDeleteResponse]
type accountMagicIpsecTunnelDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
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

// Whether the API call was successful
type AccountMagicIpsecTunnelDeleteResponseSuccess bool

const (
	AccountMagicIpsecTunnelDeleteResponseSuccessTrue AccountMagicIpsecTunnelDeleteResponseSuccess = true
)

func (r AccountMagicIpsecTunnelDeleteResponseSuccess) IsKnown() bool {
	switch r {
	case AccountMagicIpsecTunnelDeleteResponseSuccessTrue:
		return true
	}
	return false
}

type AccountMagicIpsecTunnelGeneratePskResponse struct {
	Errors   []MagicMessageItem                               `json:"errors,required"`
	Messages []MagicMessageItem                               `json:"messages,required"`
	Result   AccountMagicIpsecTunnelGeneratePskResponseResult `json:"result,required"`
	// Whether the API call was successful
	Success AccountMagicIpsecTunnelGeneratePskResponseSuccess `json:"success,required"`
	JSON    accountMagicIpsecTunnelGeneratePskResponseJSON    `json:"-"`
}

// accountMagicIpsecTunnelGeneratePskResponseJSON contains the JSON metadata for
// the struct [AccountMagicIpsecTunnelGeneratePskResponse]
type accountMagicIpsecTunnelGeneratePskResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
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

// Whether the API call was successful
type AccountMagicIpsecTunnelGeneratePskResponseSuccess bool

const (
	AccountMagicIpsecTunnelGeneratePskResponseSuccessTrue AccountMagicIpsecTunnelGeneratePskResponseSuccess = true
)

func (r AccountMagicIpsecTunnelGeneratePskResponseSuccess) IsKnown() bool {
	switch r {
	case AccountMagicIpsecTunnelGeneratePskResponseSuccessTrue:
		return true
	}
	return false
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
	XMagicNewHcTarget param.Field[bool] `header:"x-magic-new-hc-target"`
}

type AccountMagicIpsecTunnelGeneratePskParams struct {
	Body interface{} `json:"body,required"`
}

func (r AccountMagicIpsecTunnelGeneratePskParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}
