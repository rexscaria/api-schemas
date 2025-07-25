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

// Creates a new GRE tunnel. Use `?validate_only=true` as an optional query
// parameter to only run validation without persisting changes.
func (r *AccountMagicGreTunnelService) New(ctx context.Context, accountID string, params AccountMagicGreTunnelNewParams, opts ...option.RequestOption) (res *AccountMagicGreTunnelNewResponse, err error) {
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
func (r *AccountMagicGreTunnelService) Delete(ctx context.Context, accountID string, greTunnelID string, body AccountMagicGreTunnelDeleteParams, opts ...option.RequestOption) (res *AccountMagicGreTunnelDeleteResponse, err error) {
	if body.XMagicNewHcTarget.Present {
		opts = append(opts, option.WithHeader("x-magic-new-hc-target", fmt.Sprintf("%s", body.XMagicNewHcTarget)))
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
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type MagicGreTunnel struct {
	// Identifier
	ID string `json:"id,required"`
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
	ID                    apijson.Field
	CloudflareGreEndpoint apijson.Field
	CustomerGreEndpoint   apijson.Field
	InterfaceAddress      apijson.Field
	Name                  apijson.Field
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
	Errors   []MagicMessageItem                   `json:"errors,required"`
	Messages []MagicMessageItem                   `json:"messages,required"`
	Result   MagicTunnelsCollectionResponseResult `json:"result,required"`
	// Whether the API call was successful
	Success MagicTunnelsCollectionResponseSuccess `json:"success,required"`
	JSON    magicTunnelsCollectionResponseJSON    `json:"-"`
}

// magicTunnelsCollectionResponseJSON contains the JSON metadata for the struct
// [MagicTunnelsCollectionResponse]
type magicTunnelsCollectionResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
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

// Whether the API call was successful
type MagicTunnelsCollectionResponseSuccess bool

const (
	MagicTunnelsCollectionResponseSuccessTrue MagicTunnelsCollectionResponseSuccess = true
)

func (r MagicTunnelsCollectionResponseSuccess) IsKnown() bool {
	switch r {
	case MagicTunnelsCollectionResponseSuccessTrue:
		return true
	}
	return false
}

type AccountMagicGreTunnelNewResponse struct {
	Errors   []AccountMagicGreTunnelNewResponseError   `json:"errors,required"`
	Messages []AccountMagicGreTunnelNewResponseMessage `json:"messages,required"`
	Result   MagicGreTunnel                            `json:"result,required"`
	// Whether the API call was successful
	Success AccountMagicGreTunnelNewResponseSuccess `json:"success,required"`
	JSON    accountMagicGreTunnelNewResponseJSON    `json:"-"`
}

// accountMagicGreTunnelNewResponseJSON contains the JSON metadata for the struct
// [AccountMagicGreTunnelNewResponse]
type accountMagicGreTunnelNewResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicGreTunnelNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMagicGreTunnelNewResponseJSON) RawJSON() string {
	return r.raw
}

type AccountMagicGreTunnelNewResponseError struct {
	Code             int64                                        `json:"code,required"`
	Message          string                                       `json:"message,required"`
	DocumentationURL string                                       `json:"documentation_url"`
	Source           AccountMagicGreTunnelNewResponseErrorsSource `json:"source"`
	JSON             accountMagicGreTunnelNewResponseErrorJSON    `json:"-"`
}

// accountMagicGreTunnelNewResponseErrorJSON contains the JSON metadata for the
// struct [AccountMagicGreTunnelNewResponseError]
type accountMagicGreTunnelNewResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountMagicGreTunnelNewResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMagicGreTunnelNewResponseErrorJSON) RawJSON() string {
	return r.raw
}

type AccountMagicGreTunnelNewResponseErrorsSource struct {
	Pointer string                                           `json:"pointer"`
	JSON    accountMagicGreTunnelNewResponseErrorsSourceJSON `json:"-"`
}

// accountMagicGreTunnelNewResponseErrorsSourceJSON contains the JSON metadata for
// the struct [AccountMagicGreTunnelNewResponseErrorsSource]
type accountMagicGreTunnelNewResponseErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicGreTunnelNewResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMagicGreTunnelNewResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AccountMagicGreTunnelNewResponseMessage struct {
	Code             int64                                          `json:"code,required"`
	Message          string                                         `json:"message,required"`
	DocumentationURL string                                         `json:"documentation_url"`
	Source           AccountMagicGreTunnelNewResponseMessagesSource `json:"source"`
	JSON             accountMagicGreTunnelNewResponseMessageJSON    `json:"-"`
}

// accountMagicGreTunnelNewResponseMessageJSON contains the JSON metadata for the
// struct [AccountMagicGreTunnelNewResponseMessage]
type accountMagicGreTunnelNewResponseMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountMagicGreTunnelNewResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMagicGreTunnelNewResponseMessageJSON) RawJSON() string {
	return r.raw
}

type AccountMagicGreTunnelNewResponseMessagesSource struct {
	Pointer string                                             `json:"pointer"`
	JSON    accountMagicGreTunnelNewResponseMessagesSourceJSON `json:"-"`
}

// accountMagicGreTunnelNewResponseMessagesSourceJSON contains the JSON metadata
// for the struct [AccountMagicGreTunnelNewResponseMessagesSource]
type accountMagicGreTunnelNewResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicGreTunnelNewResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMagicGreTunnelNewResponseMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AccountMagicGreTunnelNewResponseSuccess bool

const (
	AccountMagicGreTunnelNewResponseSuccessTrue AccountMagicGreTunnelNewResponseSuccess = true
)

func (r AccountMagicGreTunnelNewResponseSuccess) IsKnown() bool {
	switch r {
	case AccountMagicGreTunnelNewResponseSuccessTrue:
		return true
	}
	return false
}

type AccountMagicGreTunnelGetResponse struct {
	Errors   []MagicMessageItem                     `json:"errors,required"`
	Messages []MagicMessageItem                     `json:"messages,required"`
	Result   AccountMagicGreTunnelGetResponseResult `json:"result,required"`
	// Whether the API call was successful
	Success AccountMagicGreTunnelGetResponseSuccess `json:"success,required"`
	JSON    accountMagicGreTunnelGetResponseJSON    `json:"-"`
}

// accountMagicGreTunnelGetResponseJSON contains the JSON metadata for the struct
// [AccountMagicGreTunnelGetResponse]
type accountMagicGreTunnelGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
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

// Whether the API call was successful
type AccountMagicGreTunnelGetResponseSuccess bool

const (
	AccountMagicGreTunnelGetResponseSuccessTrue AccountMagicGreTunnelGetResponseSuccess = true
)

func (r AccountMagicGreTunnelGetResponseSuccess) IsKnown() bool {
	switch r {
	case AccountMagicGreTunnelGetResponseSuccessTrue:
		return true
	}
	return false
}

type AccountMagicGreTunnelUpdateResponse struct {
	Errors   []MagicMessageItem                        `json:"errors,required"`
	Messages []MagicMessageItem                        `json:"messages,required"`
	Result   AccountMagicGreTunnelUpdateResponseResult `json:"result,required"`
	// Whether the API call was successful
	Success AccountMagicGreTunnelUpdateResponseSuccess `json:"success,required"`
	JSON    accountMagicGreTunnelUpdateResponseJSON    `json:"-"`
}

// accountMagicGreTunnelUpdateResponseJSON contains the JSON metadata for the
// struct [AccountMagicGreTunnelUpdateResponse]
type accountMagicGreTunnelUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
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

// Whether the API call was successful
type AccountMagicGreTunnelUpdateResponseSuccess bool

const (
	AccountMagicGreTunnelUpdateResponseSuccessTrue AccountMagicGreTunnelUpdateResponseSuccess = true
)

func (r AccountMagicGreTunnelUpdateResponseSuccess) IsKnown() bool {
	switch r {
	case AccountMagicGreTunnelUpdateResponseSuccessTrue:
		return true
	}
	return false
}

type AccountMagicGreTunnelDeleteResponse struct {
	Errors   []MagicMessageItem                        `json:"errors,required"`
	Messages []MagicMessageItem                        `json:"messages,required"`
	Result   AccountMagicGreTunnelDeleteResponseResult `json:"result,required"`
	// Whether the API call was successful
	Success AccountMagicGreTunnelDeleteResponseSuccess `json:"success,required"`
	JSON    accountMagicGreTunnelDeleteResponseJSON    `json:"-"`
}

// accountMagicGreTunnelDeleteResponseJSON contains the JSON metadata for the
// struct [AccountMagicGreTunnelDeleteResponse]
type accountMagicGreTunnelDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
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

// Whether the API call was successful
type AccountMagicGreTunnelDeleteResponseSuccess bool

const (
	AccountMagicGreTunnelDeleteResponseSuccessTrue AccountMagicGreTunnelDeleteResponseSuccess = true
)

func (r AccountMagicGreTunnelDeleteResponseSuccess) IsKnown() bool {
	switch r {
	case AccountMagicGreTunnelDeleteResponseSuccessTrue:
		return true
	}
	return false
}

type AccountMagicGreTunnelNewParams struct {
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

func (r AccountMagicGreTunnelNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
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
	XMagicNewHcTarget param.Field[bool] `header:"x-magic-new-hc-target"`
}
