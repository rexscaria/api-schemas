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

// AccountCfdTunnelConfigurationService contains methods and other services that
// help with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountCfdTunnelConfigurationService] method instead.
type AccountCfdTunnelConfigurationService struct {
	Options []option.RequestOption
}

// NewAccountCfdTunnelConfigurationService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountCfdTunnelConfigurationService(opts ...option.RequestOption) (r *AccountCfdTunnelConfigurationService) {
	r = &AccountCfdTunnelConfigurationService{}
	r.Options = opts
	return
}

// Gets the configuration for a remotely-managed tunnel
func (r *AccountCfdTunnelConfigurationService) Get(ctx context.Context, accountID string, tunnelID string, opts ...option.RequestOption) (res *ConfigurationResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if tunnelID == "" {
		err = errors.New("missing required tunnel_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cfd_tunnel/%s/configurations", accountID, tunnelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Adds or updates the configuration for a remotely-managed tunnel.
func (r *AccountCfdTunnelConfigurationService) Update(ctx context.Context, accountID string, tunnelID string, body AccountCfdTunnelConfigurationUpdateParams, opts ...option.RequestOption) (res *ConfigurationResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if tunnelID == "" {
		err = errors.New("missing required tunnel_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cfd_tunnel/%s/configurations", accountID, tunnelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type ConfigurationResponse struct {
	Errors   []ConfigurationResponseError   `json:"errors,required"`
	Messages []ConfigurationResponseMessage `json:"messages,required"`
	// Whether the API call was successful.
	Success ConfigurationResponseSuccess `json:"success,required"`
	// Cloudflare Tunnel configuration
	Result ConfigurationResponseResult `json:"result"`
	JSON   configurationResponseJSON   `json:"-"`
}

// configurationResponseJSON contains the JSON metadata for the struct
// [ConfigurationResponse]
type configurationResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigurationResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configurationResponseJSON) RawJSON() string {
	return r.raw
}

type ConfigurationResponseError struct {
	Code             int64                             `json:"code,required"`
	Message          string                            `json:"message,required"`
	DocumentationURL string                            `json:"documentation_url"`
	Source           ConfigurationResponseErrorsSource `json:"source"`
	JSON             configurationResponseErrorJSON    `json:"-"`
}

// configurationResponseErrorJSON contains the JSON metadata for the struct
// [ConfigurationResponseError]
type configurationResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ConfigurationResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configurationResponseErrorJSON) RawJSON() string {
	return r.raw
}

type ConfigurationResponseErrorsSource struct {
	Pointer string                                `json:"pointer"`
	JSON    configurationResponseErrorsSourceJSON `json:"-"`
}

// configurationResponseErrorsSourceJSON contains the JSON metadata for the struct
// [ConfigurationResponseErrorsSource]
type configurationResponseErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigurationResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configurationResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type ConfigurationResponseMessage struct {
	Code             int64                               `json:"code,required"`
	Message          string                              `json:"message,required"`
	DocumentationURL string                              `json:"documentation_url"`
	Source           ConfigurationResponseMessagesSource `json:"source"`
	JSON             configurationResponseMessageJSON    `json:"-"`
}

// configurationResponseMessageJSON contains the JSON metadata for the struct
// [ConfigurationResponseMessage]
type configurationResponseMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ConfigurationResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configurationResponseMessageJSON) RawJSON() string {
	return r.raw
}

type ConfigurationResponseMessagesSource struct {
	Pointer string                                  `json:"pointer"`
	JSON    configurationResponseMessagesSourceJSON `json:"-"`
}

// configurationResponseMessagesSourceJSON contains the JSON metadata for the
// struct [ConfigurationResponseMessagesSource]
type configurationResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigurationResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configurationResponseMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ConfigurationResponseSuccess bool

const (
	ConfigurationResponseSuccessTrue ConfigurationResponseSuccess = true
)

func (r ConfigurationResponseSuccess) IsKnown() bool {
	switch r {
	case ConfigurationResponseSuccessTrue:
		return true
	}
	return false
}

// Cloudflare Tunnel configuration
type ConfigurationResponseResult struct {
	// Identifier.
	AccountID string `json:"account_id"`
	// The tunnel configuration and ingress rules.
	Config    TunnelConfig `json:"config"`
	CreatedAt time.Time    `json:"created_at" format:"date-time"`
	// Indicates if this is a locally or remotely configured tunnel. If `local`, manage
	// the tunnel using a YAML file on the origin machine. If `cloudflare`, manage the
	// tunnel's configuration on the Zero Trust dashboard.
	Source ConfigurationResponseResultSource `json:"source"`
	// UUID of the tunnel.
	TunnelID string `json:"tunnel_id" format:"uuid"`
	// The version of the Tunnel Configuration.
	Version int64                           `json:"version"`
	JSON    configurationResponseResultJSON `json:"-"`
}

// configurationResponseResultJSON contains the JSON metadata for the struct
// [ConfigurationResponseResult]
type configurationResponseResultJSON struct {
	AccountID   apijson.Field
	Config      apijson.Field
	CreatedAt   apijson.Field
	Source      apijson.Field
	TunnelID    apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigurationResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configurationResponseResultJSON) RawJSON() string {
	return r.raw
}

// Indicates if this is a locally or remotely configured tunnel. If `local`, manage
// the tunnel using a YAML file on the origin machine. If `cloudflare`, manage the
// tunnel's configuration on the Zero Trust dashboard.
type ConfigurationResponseResultSource string

const (
	ConfigurationResponseResultSourceLocal      ConfigurationResponseResultSource = "local"
	ConfigurationResponseResultSourceCloudflare ConfigurationResponseResultSource = "cloudflare"
)

func (r ConfigurationResponseResultSource) IsKnown() bool {
	switch r {
	case ConfigurationResponseResultSourceLocal, ConfigurationResponseResultSourceCloudflare:
		return true
	}
	return false
}

// Configuration parameters for the public hostname specific connection settings
// between cloudflared and origin server.
type OriginRequest struct {
	// For all L7 requests to this hostname, cloudflared will validate each request's
	// Cf-Access-Jwt-Assertion request header.
	Access OriginRequestAccess `json:"access"`
	// Path to the certificate authority (CA) for the certificate of your origin. This
	// option should be used only if your certificate is not signed by Cloudflare.
	CaPool string `json:"caPool"`
	// Timeout for establishing a new TCP connection to your origin server. This
	// excludes the time taken to establish TLS, which is controlled by tlsTimeout.
	ConnectTimeout int64 `json:"connectTimeout"`
	// Disables chunked transfer encoding. Useful if you are running a WSGI server.
	DisableChunkedEncoding bool `json:"disableChunkedEncoding"`
	// Attempt to connect to origin using HTTP2. Origin must be configured as https.
	Http2Origin bool `json:"http2Origin"`
	// Sets the HTTP Host header on requests sent to the local service.
	HTTPHostHeader string `json:"httpHostHeader"`
	// Maximum number of idle keepalive connections between Tunnel and your origin.
	// This does not restrict the total number of concurrent connections.
	KeepAliveConnections int64 `json:"keepAliveConnections"`
	// Timeout after which an idle keepalive connection can be discarded.
	KeepAliveTimeout int64 `json:"keepAliveTimeout"`
	// Disable the “happy eyeballs” algorithm for IPv4/IPv6 fallback if your local
	// network has misconfigured one of the protocols.
	NoHappyEyeballs bool `json:"noHappyEyeballs"`
	// Disables TLS verification of the certificate presented by your origin. Will
	// allow any certificate from the origin to be accepted.
	NoTlsVerify bool `json:"noTLSVerify"`
	// Hostname that cloudflared should expect from your origin server certificate.
	OriginServerName string `json:"originServerName"`
	// cloudflared starts a proxy server to translate HTTP traffic into TCP when
	// proxying, for example, SSH or RDP. This configures what type of proxy will be
	// started. Valid options are: "" for the regular proxy and "socks" for a SOCKS5
	// proxy.
	ProxyType string `json:"proxyType"`
	// The timeout after which a TCP keepalive packet is sent on a connection between
	// Tunnel and the origin server.
	TcpKeepAlive int64 `json:"tcpKeepAlive"`
	// Timeout for completing a TLS handshake to your origin server, if you have chosen
	// to connect Tunnel to an HTTPS server.
	TlsTimeout int64             `json:"tlsTimeout"`
	JSON       originRequestJSON `json:"-"`
}

// originRequestJSON contains the JSON metadata for the struct [OriginRequest]
type originRequestJSON struct {
	Access                 apijson.Field
	CaPool                 apijson.Field
	ConnectTimeout         apijson.Field
	DisableChunkedEncoding apijson.Field
	Http2Origin            apijson.Field
	HTTPHostHeader         apijson.Field
	KeepAliveConnections   apijson.Field
	KeepAliveTimeout       apijson.Field
	NoHappyEyeballs        apijson.Field
	NoTlsVerify            apijson.Field
	OriginServerName       apijson.Field
	ProxyType              apijson.Field
	TcpKeepAlive           apijson.Field
	TlsTimeout             apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *OriginRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r originRequestJSON) RawJSON() string {
	return r.raw
}

// For all L7 requests to this hostname, cloudflared will validate each request's
// Cf-Access-Jwt-Assertion request header.
type OriginRequestAccess struct {
	// Access applications that are allowed to reach this hostname for this Tunnel.
	// Audience tags can be identified in the dashboard or via the List Access policies
	// API.
	AudTag   []string `json:"audTag,required"`
	TeamName string   `json:"teamName,required"`
	// Deny traffic that has not fulfilled Access authorization.
	Required bool                    `json:"required"`
	JSON     originRequestAccessJSON `json:"-"`
}

// originRequestAccessJSON contains the JSON metadata for the struct
// [OriginRequestAccess]
type originRequestAccessJSON struct {
	AudTag      apijson.Field
	TeamName    apijson.Field
	Required    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginRequestAccess) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r originRequestAccessJSON) RawJSON() string {
	return r.raw
}

// Configuration parameters for the public hostname specific connection settings
// between cloudflared and origin server.
type OriginRequestParam struct {
	// For all L7 requests to this hostname, cloudflared will validate each request's
	// Cf-Access-Jwt-Assertion request header.
	Access param.Field[OriginRequestAccessParam] `json:"access"`
	// Path to the certificate authority (CA) for the certificate of your origin. This
	// option should be used only if your certificate is not signed by Cloudflare.
	CaPool param.Field[string] `json:"caPool"`
	// Timeout for establishing a new TCP connection to your origin server. This
	// excludes the time taken to establish TLS, which is controlled by tlsTimeout.
	ConnectTimeout param.Field[int64] `json:"connectTimeout"`
	// Disables chunked transfer encoding. Useful if you are running a WSGI server.
	DisableChunkedEncoding param.Field[bool] `json:"disableChunkedEncoding"`
	// Attempt to connect to origin using HTTP2. Origin must be configured as https.
	Http2Origin param.Field[bool] `json:"http2Origin"`
	// Sets the HTTP Host header on requests sent to the local service.
	HTTPHostHeader param.Field[string] `json:"httpHostHeader"`
	// Maximum number of idle keepalive connections between Tunnel and your origin.
	// This does not restrict the total number of concurrent connections.
	KeepAliveConnections param.Field[int64] `json:"keepAliveConnections"`
	// Timeout after which an idle keepalive connection can be discarded.
	KeepAliveTimeout param.Field[int64] `json:"keepAliveTimeout"`
	// Disable the “happy eyeballs” algorithm for IPv4/IPv6 fallback if your local
	// network has misconfigured one of the protocols.
	NoHappyEyeballs param.Field[bool] `json:"noHappyEyeballs"`
	// Disables TLS verification of the certificate presented by your origin. Will
	// allow any certificate from the origin to be accepted.
	NoTlsVerify param.Field[bool] `json:"noTLSVerify"`
	// Hostname that cloudflared should expect from your origin server certificate.
	OriginServerName param.Field[string] `json:"originServerName"`
	// cloudflared starts a proxy server to translate HTTP traffic into TCP when
	// proxying, for example, SSH or RDP. This configures what type of proxy will be
	// started. Valid options are: "" for the regular proxy and "socks" for a SOCKS5
	// proxy.
	ProxyType param.Field[string] `json:"proxyType"`
	// The timeout after which a TCP keepalive packet is sent on a connection between
	// Tunnel and the origin server.
	TcpKeepAlive param.Field[int64] `json:"tcpKeepAlive"`
	// Timeout for completing a TLS handshake to your origin server, if you have chosen
	// to connect Tunnel to an HTTPS server.
	TlsTimeout param.Field[int64] `json:"tlsTimeout"`
}

func (r OriginRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// For all L7 requests to this hostname, cloudflared will validate each request's
// Cf-Access-Jwt-Assertion request header.
type OriginRequestAccessParam struct {
	// Access applications that are allowed to reach this hostname for this Tunnel.
	// Audience tags can be identified in the dashboard or via the List Access policies
	// API.
	AudTag   param.Field[[]string] `json:"audTag,required"`
	TeamName param.Field[string]   `json:"teamName,required"`
	// Deny traffic that has not fulfilled Access authorization.
	Required param.Field[bool] `json:"required"`
}

func (r OriginRequestAccessParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The tunnel configuration and ingress rules.
type TunnelConfig struct {
	// List of public hostname definitions. At least one ingress rule needs to be
	// defined for the tunnel.
	Ingress []TunnelConfigIngress `json:"ingress"`
	// Configuration parameters for the public hostname specific connection settings
	// between cloudflared and origin server.
	OriginRequest OriginRequest `json:"originRequest"`
	// Enable private network access from WARP users to private network routes. This is
	// enabled if the tunnel has an assigned route.
	WarpRouting TunnelConfigWarpRouting `json:"warp-routing"`
	JSON        tunnelConfigJSON        `json:"-"`
}

// tunnelConfigJSON contains the JSON metadata for the struct [TunnelConfig]
type tunnelConfigJSON struct {
	Ingress       apijson.Field
	OriginRequest apijson.Field
	WarpRouting   apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *TunnelConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tunnelConfigJSON) RawJSON() string {
	return r.raw
}

// Public hostname
type TunnelConfigIngress struct {
	// Public hostname for this service.
	Hostname string `json:"hostname,required"`
	// Protocol and address of destination server. Supported protocols: http://,
	// https://, unix://, tcp://, ssh://, rdp://, unix+tls://, smb://. Alternatively
	// can return a HTTP status code http_status:[code] e.g. 'http_status:404'.
	Service string `json:"service,required"`
	// Configuration parameters for the public hostname specific connection settings
	// between cloudflared and origin server.
	OriginRequest OriginRequest `json:"originRequest"`
	// Requests with this path route to this public hostname.
	Path string                  `json:"path"`
	JSON tunnelConfigIngressJSON `json:"-"`
}

// tunnelConfigIngressJSON contains the JSON metadata for the struct
// [TunnelConfigIngress]
type tunnelConfigIngressJSON struct {
	Hostname      apijson.Field
	Service       apijson.Field
	OriginRequest apijson.Field
	Path          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *TunnelConfigIngress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tunnelConfigIngressJSON) RawJSON() string {
	return r.raw
}

// Enable private network access from WARP users to private network routes. This is
// enabled if the tunnel has an assigned route.
type TunnelConfigWarpRouting struct {
	Enabled bool                        `json:"enabled"`
	JSON    tunnelConfigWarpRoutingJSON `json:"-"`
}

// tunnelConfigWarpRoutingJSON contains the JSON metadata for the struct
// [TunnelConfigWarpRouting]
type tunnelConfigWarpRoutingJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelConfigWarpRouting) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tunnelConfigWarpRoutingJSON) RawJSON() string {
	return r.raw
}

// The tunnel configuration and ingress rules.
type TunnelConfigParam struct {
	// List of public hostname definitions. At least one ingress rule needs to be
	// defined for the tunnel.
	Ingress param.Field[[]TunnelConfigIngressParam] `json:"ingress"`
	// Configuration parameters for the public hostname specific connection settings
	// between cloudflared and origin server.
	OriginRequest param.Field[OriginRequestParam] `json:"originRequest"`
}

func (r TunnelConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Public hostname
type TunnelConfigIngressParam struct {
	// Public hostname for this service.
	Hostname param.Field[string] `json:"hostname,required"`
	// Protocol and address of destination server. Supported protocols: http://,
	// https://, unix://, tcp://, ssh://, rdp://, unix+tls://, smb://. Alternatively
	// can return a HTTP status code http_status:[code] e.g. 'http_status:404'.
	Service param.Field[string] `json:"service,required"`
	// Configuration parameters for the public hostname specific connection settings
	// between cloudflared and origin server.
	OriginRequest param.Field[OriginRequestParam] `json:"originRequest"`
	// Requests with this path route to this public hostname.
	Path param.Field[string] `json:"path"`
}

func (r TunnelConfigIngressParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enable private network access from WARP users to private network routes. This is
// enabled if the tunnel has an assigned route.
type TunnelConfigWarpRoutingParam struct {
	Enabled param.Field[bool] `json:"enabled"`
}

func (r TunnelConfigWarpRoutingParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountCfdTunnelConfigurationUpdateParams struct {
	// The tunnel configuration and ingress rules.
	Config param.Field[TunnelConfigParam] `json:"config"`
}

func (r AccountCfdTunnelConfigurationUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
