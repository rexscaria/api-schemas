// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/apiquery"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
	"github.com/rexscaria/api-schemas/shared"
	"github.com/tidwall/gjson"
)

// AccountCfdTunnelService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountCfdTunnelService] method instead.
type AccountCfdTunnelService struct {
	Options        []option.RequestOption
	Configurations *AccountCfdTunnelConfigurationService
	Connections    *AccountCfdTunnelConnectionService
}

// NewAccountCfdTunnelService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountCfdTunnelService(opts ...option.RequestOption) (r *AccountCfdTunnelService) {
	r = &AccountCfdTunnelService{}
	r.Options = opts
	r.Configurations = NewAccountCfdTunnelConfigurationService(opts...)
	r.Connections = NewAccountCfdTunnelConnectionService(opts...)
	return
}

// Creates a new Cloudflare Tunnel in an account.
func (r *AccountCfdTunnelService) New(ctx context.Context, accountID string, body AccountCfdTunnelNewParams, opts ...option.RequestOption) (res *TunnelResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cfd_tunnel", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Fetches a single Cloudflare Tunnel.
func (r *AccountCfdTunnelService) Get(ctx context.Context, accountID string, tunnelID string, opts ...option.RequestOption) (res *TunnelResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if tunnelID == "" {
		err = errors.New("missing required tunnel_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cfd_tunnel/%s", accountID, tunnelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates an existing Cloudflare Tunnel.
func (r *AccountCfdTunnelService) Update(ctx context.Context, accountID string, tunnelID string, body AccountCfdTunnelUpdateParams, opts ...option.RequestOption) (res *TunnelResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if tunnelID == "" {
		err = errors.New("missing required tunnel_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cfd_tunnel/%s", accountID, tunnelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Lists and filters Cloudflare Tunnels in an account.
func (r *AccountCfdTunnelService) List(ctx context.Context, accountID string, query AccountCfdTunnelListParams, opts ...option.RequestOption) (res *TunnelResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cfd_tunnel", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Deletes a Cloudflare Tunnel from an account.
func (r *AccountCfdTunnelService) Delete(ctx context.Context, accountID string, tunnelID string, body AccountCfdTunnelDeleteParams, opts ...option.RequestOption) (res *TunnelResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if tunnelID == "" {
		err = errors.New("missing required tunnel_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cfd_tunnel/%s", accountID, tunnelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

// Fetches connector and connection details for a Cloudflare Tunnel.
func (r *AccountCfdTunnelService) GetConnector(ctx context.Context, accountID string, tunnelID string, connectorID string, opts ...option.RequestOption) (res *AccountCfdTunnelGetConnectorResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if tunnelID == "" {
		err = errors.New("missing required tunnel_id parameter")
		return
	}
	if connectorID == "" {
		err = errors.New("missing required connector_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cfd_tunnel/%s/connectors/%s", accountID, tunnelID, connectorID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Gets a management token used to access the management resources (i.e. Streaming
// Logs) of a tunnel.
func (r *AccountCfdTunnelService) GetManagementToken(ctx context.Context, accountID string, tunnelID string, body AccountCfdTunnelGetManagementTokenParams, opts ...option.RequestOption) (res *TunnelResponseToken, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if tunnelID == "" {
		err = errors.New("missing required tunnel_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cfd_tunnel/%s/management", accountID, tunnelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Gets the token used to associate cloudflared with a specific tunnel.
func (r *AccountCfdTunnelService) GetToken(ctx context.Context, accountID string, tunnelID string, opts ...option.RequestOption) (res *TunnelResponseToken, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if tunnelID == "" {
		err = errors.New("missing required tunnel_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cfd_tunnel/%s/token", accountID, tunnelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type APIResponseCollectionTunnel struct {
	Result     []interface{}                         `json:"result,nullable"`
	ResultInfo APIResponseCollectionTunnelResultInfo `json:"result_info"`
	JSON       apiResponseCollectionTunnelJSON       `json:"-"`
	APIResponseTunnel
}

// apiResponseCollectionTunnelJSON contains the JSON metadata for the struct
// [APIResponseCollectionTunnel]
type apiResponseCollectionTunnelJSON struct {
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIResponseCollectionTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiResponseCollectionTunnelJSON) RawJSON() string {
	return r.raw
}

type APIResponseCollectionTunnelResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                   `json:"total_count"`
	JSON       apiResponseCollectionTunnelResultInfoJSON `json:"-"`
}

// apiResponseCollectionTunnelResultInfoJSON contains the JSON metadata for the
// struct [APIResponseCollectionTunnelResultInfo]
type apiResponseCollectionTunnelResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIResponseCollectionTunnelResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiResponseCollectionTunnelResultInfoJSON) RawJSON() string {
	return r.raw
}

type APIResponseTunnel struct {
	Errors   []MessagesTunnelItem         `json:"errors,required"`
	Messages []MessagesTunnelItem         `json:"messages,required"`
	Result   APIResponseTunnelResultUnion `json:"result,required"`
	// Whether the API call was successful
	Success APIResponseTunnelSuccess `json:"success,required"`
	JSON    apiResponseTunnelJSON    `json:"-"`
}

// apiResponseTunnelJSON contains the JSON metadata for the struct
// [APIResponseTunnel]
type apiResponseTunnelJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIResponseTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiResponseTunnelJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [APIResponseTunnelResultArray] or [shared.UnionString].
type APIResponseTunnelResultUnion interface {
	ImplementsAPIResponseTunnelResultUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*APIResponseTunnelResultUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(APIResponseTunnelResultArray{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type APIResponseTunnelResultArray []interface{}

func (r APIResponseTunnelResultArray) ImplementsAPIResponseTunnelResultUnion() {}

// Whether the API call was successful
type APIResponseTunnelSuccess bool

const (
	APIResponseTunnelSuccessTrue APIResponseTunnelSuccess = true
)

func (r APIResponseTunnelSuccess) IsKnown() bool {
	switch r {
	case APIResponseTunnelSuccessTrue:
		return true
	}
	return false
}

// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
type CfdTunnel struct {
	// UUID of the tunnel.
	ID string `json:"id" format:"uuid"`
	// Cloudflare account ID
	AccountTag string `json:"account_tag"`
	// The Cloudflare Tunnel connections between your origin and Cloudflare's edge.
	Connections []SchemasConnection `json:"connections"`
	// Timestamp of when the tunnel established at least one connection to Cloudflare's
	// edge. If `null`, the tunnel is inactive.
	ConnsActiveAt time.Time `json:"conns_active_at" format:"date-time"`
	// Timestamp of when the tunnel became inactive (no connections to Cloudflare's
	// edge). If `null`, the tunnel is active.
	ConnsInactiveAt time.Time `json:"conns_inactive_at" format:"date-time"`
	// Timestamp of when the resource was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Timestamp of when the resource was deleted. If `null`, the resource has not been
	// deleted.
	DeletedAt time.Time `json:"deleted_at" format:"date-time"`
	// Metadata associated with the tunnel.
	Metadata interface{} `json:"metadata"`
	// A user-friendly name for a tunnel.
	Name string `json:"name"`
	// If `true`, the tunnel can be configured remotely from the Zero Trust dashboard.
	// If `false`, the tunnel must be configured locally on the origin machine.
	RemoteConfig bool `json:"remote_config"`
	// The status of the tunnel. Valid values are `inactive` (tunnel has never been
	// run), `degraded` (tunnel is active and able to serve traffic but in an unhealthy
	// state), `healthy` (tunnel is active and able to serve traffic), or `down`
	// (tunnel can not serve traffic as it has no connections to the Cloudflare Edge).
	Status StatusTunnel `json:"status"`
	// The type of tunnel.
	TunType TunnelType    `json:"tun_type"`
	JSON    cfdTunnelJSON `json:"-"`
}

// cfdTunnelJSON contains the JSON metadata for the struct [CfdTunnel]
type cfdTunnelJSON struct {
	ID              apijson.Field
	AccountTag      apijson.Field
	Connections     apijson.Field
	ConnsActiveAt   apijson.Field
	ConnsInactiveAt apijson.Field
	CreatedAt       apijson.Field
	DeletedAt       apijson.Field
	Metadata        apijson.Field
	Name            apijson.Field
	RemoteConfig    apijson.Field
	Status          apijson.Field
	TunType         apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *CfdTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cfdTunnelJSON) RawJSON() string {
	return r.raw
}

func (r CfdTunnel) implementsTunnelResponseCollectionResult() {}

func (r CfdTunnel) implementsTunnelResponseSingleResult() {}

type MessagesTunnelItem struct {
	Code    int64                  `json:"code,required"`
	Message string                 `json:"message,required"`
	JSON    messagesTunnelItemJSON `json:"-"`
}

// messagesTunnelItemJSON contains the JSON metadata for the struct
// [MessagesTunnelItem]
type messagesTunnelItemJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessagesTunnelItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messagesTunnelItemJSON) RawJSON() string {
	return r.raw
}

type SchemasConnection struct {
	// UUID of the Cloudflare Tunnel connection.
	ID string `json:"id" format:"uuid"`
	// UUID of the Cloudflare Tunnel connector.
	ClientID string `json:"client_id" format:"uuid"`
	// The cloudflared version used to establish this connection.
	ClientVersion string `json:"client_version"`
	// The Cloudflare data center used for this connection.
	ColoName string `json:"colo_name"`
	// Cloudflare continues to track connections for several minutes after they
	// disconnect. This is an optimization to improve latency and reliability of
	// reconnecting. If `true`, the connection has disconnected but is still being
	// tracked. If `false`, the connection is actively serving traffic.
	IsPendingReconnect bool `json:"is_pending_reconnect"`
	// Timestamp of when the connection was established.
	OpenedAt time.Time `json:"opened_at" format:"date-time"`
	// The public IP address of the host running cloudflared.
	OriginIP string `json:"origin_ip"`
	// UUID of the Cloudflare Tunnel connection.
	Uuid string                `json:"uuid" format:"uuid"`
	JSON schemasConnectionJSON `json:"-"`
}

// schemasConnectionJSON contains the JSON metadata for the struct
// [SchemasConnection]
type schemasConnectionJSON struct {
	ID                 apijson.Field
	ClientID           apijson.Field
	ClientVersion      apijson.Field
	ColoName           apijson.Field
	IsPendingReconnect apijson.Field
	OpenedAt           apijson.Field
	OriginIP           apijson.Field
	Uuid               apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *SchemasConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r schemasConnectionJSON) RawJSON() string {
	return r.raw
}

// The status of the tunnel. Valid values are `inactive` (tunnel has never been
// run), `degraded` (tunnel is active and able to serve traffic but in an unhealthy
// state), `healthy` (tunnel is active and able to serve traffic), or `down`
// (tunnel can not serve traffic as it has no connections to the Cloudflare Edge).
type StatusTunnel string

const (
	StatusTunnelInactive StatusTunnel = "inactive"
	StatusTunnelDegraded StatusTunnel = "degraded"
	StatusTunnelHealthy  StatusTunnel = "healthy"
	StatusTunnelDown     StatusTunnel = "down"
)

func (r StatusTunnel) IsKnown() bool {
	switch r {
	case StatusTunnelInactive, StatusTunnelDegraded, StatusTunnelHealthy, StatusTunnelDown:
		return true
	}
	return false
}

// A client (typically cloudflared) that maintains connections to a Cloudflare data
// center.
type TunnelClient struct {
	// UUID of the Cloudflare Tunnel connection.
	ID string `json:"id" format:"uuid"`
	// The cloudflared OS architecture used to establish this connection.
	Arch string `json:"arch"`
	// The version of the remote tunnel configuration. Used internally to sync
	// cloudflared with the Zero Trust dashboard.
	ConfigVersion int64 `json:"config_version"`
	// The Cloudflare Tunnel connections between your origin and Cloudflare's edge.
	Conns []SchemasConnection `json:"conns"`
	// Features enabled for the Cloudflare Tunnel.
	Features []string `json:"features"`
	// Timestamp of when the tunnel connection was started.
	RunAt time.Time `json:"run_at" format:"date-time"`
	// The cloudflared version used to establish this connection.
	Version string           `json:"version"`
	JSON    tunnelClientJSON `json:"-"`
}

// tunnelClientJSON contains the JSON metadata for the struct [TunnelClient]
type tunnelClientJSON struct {
	ID            apijson.Field
	Arch          apijson.Field
	ConfigVersion apijson.Field
	Conns         apijson.Field
	Features      apijson.Field
	RunAt         apijson.Field
	Version       apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *TunnelClient) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tunnelClientJSON) RawJSON() string {
	return r.raw
}

type TunnelResponseCollection struct {
	Result []TunnelResponseCollectionResult `json:"result"`
	JSON   tunnelResponseCollectionJSON     `json:"-"`
	APIResponseCollectionTunnel
}

// tunnelResponseCollectionJSON contains the JSON metadata for the struct
// [TunnelResponseCollection]
type tunnelResponseCollectionJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelResponseCollection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tunnelResponseCollectionJSON) RawJSON() string {
	return r.raw
}

// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
type TunnelResponseCollectionResult struct {
	// UUID of the tunnel.
	ID string `json:"id" format:"uuid"`
	// Cloudflare account ID
	AccountTag string `json:"account_tag"`
	// This field can have the runtime type of [[]SchemasConnection].
	Connections interface{} `json:"connections"`
	// Timestamp of when the tunnel established at least one connection to Cloudflare's
	// edge. If `null`, the tunnel is inactive.
	ConnsActiveAt time.Time `json:"conns_active_at" format:"date-time"`
	// Timestamp of when the tunnel became inactive (no connections to Cloudflare's
	// edge). If `null`, the tunnel is active.
	ConnsInactiveAt time.Time `json:"conns_inactive_at" format:"date-time"`
	// Timestamp of when the resource was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Timestamp of when the resource was deleted. If `null`, the resource has not been
	// deleted.
	DeletedAt time.Time `json:"deleted_at" format:"date-time"`
	// This field can have the runtime type of [interface{}].
	Metadata interface{} `json:"metadata"`
	// A user-friendly name for a tunnel.
	Name string `json:"name"`
	// If `true`, the tunnel can be configured remotely from the Zero Trust dashboard.
	// If `false`, the tunnel must be configured locally on the origin machine.
	RemoteConfig bool `json:"remote_config"`
	// The status of the tunnel. Valid values are `inactive` (tunnel has never been
	// run), `degraded` (tunnel is active and able to serve traffic but in an unhealthy
	// state), `healthy` (tunnel is active and able to serve traffic), or `down`
	// (tunnel can not serve traffic as it has no connections to the Cloudflare Edge).
	Status StatusTunnel `json:"status"`
	// The type of tunnel.
	TunType TunnelType                         `json:"tun_type"`
	JSON    tunnelResponseCollectionResultJSON `json:"-"`
	union   TunnelResponseCollectionResultUnion
}

// tunnelResponseCollectionResultJSON contains the JSON metadata for the struct
// [TunnelResponseCollectionResult]
type tunnelResponseCollectionResultJSON struct {
	ID              apijson.Field
	AccountTag      apijson.Field
	Connections     apijson.Field
	ConnsActiveAt   apijson.Field
	ConnsInactiveAt apijson.Field
	CreatedAt       apijson.Field
	DeletedAt       apijson.Field
	Metadata        apijson.Field
	Name            apijson.Field
	RemoteConfig    apijson.Field
	Status          apijson.Field
	TunType         apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r tunnelResponseCollectionResultJSON) RawJSON() string {
	return r.raw
}

func (r *TunnelResponseCollectionResult) UnmarshalJSON(data []byte) (err error) {
	*r = TunnelResponseCollectionResult{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [TunnelResponseCollectionResultUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are [CfdTunnel], [WarpConnectorTunnel].
func (r TunnelResponseCollectionResult) AsUnion() TunnelResponseCollectionResultUnion {
	return r.union
}

// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
//
// Union satisfied by [CfdTunnel] or [WarpConnectorTunnel].
type TunnelResponseCollectionResultUnion interface {
	implementsTunnelResponseCollectionResult()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TunnelResponseCollectionResultUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CfdTunnel{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(WarpConnectorTunnel{}),
		},
	)
}

type TunnelResponseSingle struct {
	// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
	Result TunnelResponseSingleResult `json:"result"`
	JSON   tunnelResponseSingleJSON   `json:"-"`
	APIResponseTunnel
}

// tunnelResponseSingleJSON contains the JSON metadata for the struct
// [TunnelResponseSingle]
type tunnelResponseSingleJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelResponseSingle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tunnelResponseSingleJSON) RawJSON() string {
	return r.raw
}

// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
type TunnelResponseSingleResult struct {
	// UUID of the tunnel.
	ID string `json:"id" format:"uuid"`
	// Cloudflare account ID
	AccountTag string `json:"account_tag"`
	// This field can have the runtime type of [[]SchemasConnection].
	Connections interface{} `json:"connections"`
	// Timestamp of when the tunnel established at least one connection to Cloudflare's
	// edge. If `null`, the tunnel is inactive.
	ConnsActiveAt time.Time `json:"conns_active_at" format:"date-time"`
	// Timestamp of when the tunnel became inactive (no connections to Cloudflare's
	// edge). If `null`, the tunnel is active.
	ConnsInactiveAt time.Time `json:"conns_inactive_at" format:"date-time"`
	// Timestamp of when the resource was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Timestamp of when the resource was deleted. If `null`, the resource has not been
	// deleted.
	DeletedAt time.Time `json:"deleted_at" format:"date-time"`
	// This field can have the runtime type of [interface{}].
	Metadata interface{} `json:"metadata"`
	// A user-friendly name for a tunnel.
	Name string `json:"name"`
	// If `true`, the tunnel can be configured remotely from the Zero Trust dashboard.
	// If `false`, the tunnel must be configured locally on the origin machine.
	RemoteConfig bool `json:"remote_config"`
	// The status of the tunnel. Valid values are `inactive` (tunnel has never been
	// run), `degraded` (tunnel is active and able to serve traffic but in an unhealthy
	// state), `healthy` (tunnel is active and able to serve traffic), or `down`
	// (tunnel can not serve traffic as it has no connections to the Cloudflare Edge).
	Status StatusTunnel `json:"status"`
	// The type of tunnel.
	TunType TunnelType                     `json:"tun_type"`
	JSON    tunnelResponseSingleResultJSON `json:"-"`
	union   TunnelResponseSingleResultUnion
}

// tunnelResponseSingleResultJSON contains the JSON metadata for the struct
// [TunnelResponseSingleResult]
type tunnelResponseSingleResultJSON struct {
	ID              apijson.Field
	AccountTag      apijson.Field
	Connections     apijson.Field
	ConnsActiveAt   apijson.Field
	ConnsInactiveAt apijson.Field
	CreatedAt       apijson.Field
	DeletedAt       apijson.Field
	Metadata        apijson.Field
	Name            apijson.Field
	RemoteConfig    apijson.Field
	Status          apijson.Field
	TunType         apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r tunnelResponseSingleResultJSON) RawJSON() string {
	return r.raw
}

func (r *TunnelResponseSingleResult) UnmarshalJSON(data []byte) (err error) {
	*r = TunnelResponseSingleResult{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [TunnelResponseSingleResultUnion] interface which you can cast
// to the specific types for more type safety.
//
// Possible runtime types of the union are [CfdTunnel], [WarpConnectorTunnel].
func (r TunnelResponseSingleResult) AsUnion() TunnelResponseSingleResultUnion {
	return r.union
}

// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
//
// Union satisfied by [CfdTunnel] or [WarpConnectorTunnel].
type TunnelResponseSingleResultUnion interface {
	implementsTunnelResponseSingleResult()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TunnelResponseSingleResultUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CfdTunnel{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(WarpConnectorTunnel{}),
		},
	)
}

type TunnelResponseToken struct {
	// The Tunnel Token is used as a mechanism to authenticate the operation of a
	// tunnel.
	Result string                  `json:"result"`
	JSON   tunnelResponseTokenJSON `json:"-"`
	APIResponseTunnel
}

// tunnelResponseTokenJSON contains the JSON metadata for the struct
// [TunnelResponseToken]
type tunnelResponseTokenJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelResponseToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tunnelResponseTokenJSON) RawJSON() string {
	return r.raw
}

// A Warp Connector Tunnel that connects your origin to Cloudflare's edge.
type WarpConnectorTunnel struct {
	// UUID of the tunnel.
	ID string `json:"id" format:"uuid"`
	// Cloudflare account ID
	AccountTag string `json:"account_tag"`
	// The Cloudflare Tunnel connections between your origin and Cloudflare's edge.
	Connections []SchemasConnection `json:"connections"`
	// Timestamp of when the tunnel established at least one connection to Cloudflare's
	// edge. If `null`, the tunnel is inactive.
	ConnsActiveAt time.Time `json:"conns_active_at" format:"date-time"`
	// Timestamp of when the tunnel became inactive (no connections to Cloudflare's
	// edge). If `null`, the tunnel is active.
	ConnsInactiveAt time.Time `json:"conns_inactive_at" format:"date-time"`
	// Timestamp of when the resource was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Timestamp of when the resource was deleted. If `null`, the resource has not been
	// deleted.
	DeletedAt time.Time `json:"deleted_at" format:"date-time"`
	// Metadata associated with the tunnel.
	Metadata interface{} `json:"metadata"`
	// A user-friendly name for a tunnel.
	Name string `json:"name"`
	// The status of the tunnel. Valid values are `inactive` (tunnel has never been
	// run), `degraded` (tunnel is active and able to serve traffic but in an unhealthy
	// state), `healthy` (tunnel is active and able to serve traffic), or `down`
	// (tunnel can not serve traffic as it has no connections to the Cloudflare Edge).
	Status StatusTunnel `json:"status"`
	// The type of tunnel.
	TunType TunnelType              `json:"tun_type"`
	JSON    warpConnectorTunnelJSON `json:"-"`
}

// warpConnectorTunnelJSON contains the JSON metadata for the struct
// [WarpConnectorTunnel]
type warpConnectorTunnelJSON struct {
	ID              apijson.Field
	AccountTag      apijson.Field
	Connections     apijson.Field
	ConnsActiveAt   apijson.Field
	ConnsInactiveAt apijson.Field
	CreatedAt       apijson.Field
	DeletedAt       apijson.Field
	Metadata        apijson.Field
	Name            apijson.Field
	Status          apijson.Field
	TunType         apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *WarpConnectorTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r warpConnectorTunnelJSON) RawJSON() string {
	return r.raw
}

func (r WarpConnectorTunnel) implementsTunnelResponseCollectionResult() {}

func (r WarpConnectorTunnel) implementsTunnelResponseSingleResult() {}

type AccountCfdTunnelGetConnectorResponse struct {
	// A client (typically cloudflared) that maintains connections to a Cloudflare data
	// center.
	Result TunnelClient                             `json:"result"`
	JSON   accountCfdTunnelGetConnectorResponseJSON `json:"-"`
	APIResponseTunnel
}

// accountCfdTunnelGetConnectorResponseJSON contains the JSON metadata for the
// struct [AccountCfdTunnelGetConnectorResponse]
type accountCfdTunnelGetConnectorResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCfdTunnelGetConnectorResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountCfdTunnelGetConnectorResponseJSON) RawJSON() string {
	return r.raw
}

type AccountCfdTunnelNewParams struct {
	// A user-friendly name for a tunnel.
	Name param.Field[string] `json:"name,required"`
	// Indicates if this is a locally or remotely configured tunnel. If `local`, manage
	// the tunnel using a YAML file on the origin machine. If `cloudflare`, manage the
	// tunnel on the Zero Trust dashboard.
	ConfigSrc param.Field[AccountCfdTunnelNewParamsConfigSrc] `json:"config_src"`
	// Sets the password required to run a locally-managed tunnel. Must be at least 32
	// bytes and encoded as a base64 string.
	TunnelSecret param.Field[string] `json:"tunnel_secret"`
}

func (r AccountCfdTunnelNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Indicates if this is a locally or remotely configured tunnel. If `local`, manage
// the tunnel using a YAML file on the origin machine. If `cloudflare`, manage the
// tunnel on the Zero Trust dashboard.
type AccountCfdTunnelNewParamsConfigSrc string

const (
	AccountCfdTunnelNewParamsConfigSrcLocal      AccountCfdTunnelNewParamsConfigSrc = "local"
	AccountCfdTunnelNewParamsConfigSrcCloudflare AccountCfdTunnelNewParamsConfigSrc = "cloudflare"
)

func (r AccountCfdTunnelNewParamsConfigSrc) IsKnown() bool {
	switch r {
	case AccountCfdTunnelNewParamsConfigSrcLocal, AccountCfdTunnelNewParamsConfigSrcCloudflare:
		return true
	}
	return false
}

type AccountCfdTunnelUpdateParams struct {
	// A user-friendly name for a tunnel.
	Name param.Field[string] `json:"name"`
	// Sets the password required to run a locally-managed tunnel. Must be at least 32
	// bytes and encoded as a base64 string.
	TunnelSecret param.Field[string] `json:"tunnel_secret"`
}

func (r AccountCfdTunnelUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountCfdTunnelListParams struct {
	ExcludePrefix param.Field[string] `query:"exclude_prefix"`
	// If provided, include only resources that were created (and not deleted) before
	// this time. URL encoded.
	ExistedAt     param.Field[string] `query:"existed_at" format:"url-encoded-date-time"`
	IncludePrefix param.Field[string] `query:"include_prefix"`
	// If `true`, only include deleted tunnels. If `false`, exclude deleted tunnels. If
	// empty, all tunnels will be included.
	IsDeleted param.Field[bool] `query:"is_deleted"`
	// A user-friendly name for a tunnel.
	Name param.Field[string] `query:"name"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Number of results to display.
	PerPage param.Field[float64] `query:"per_page"`
	// The status of the tunnel. Valid values are `inactive` (tunnel has never been
	// run), `degraded` (tunnel is active and able to serve traffic but in an unhealthy
	// state), `healthy` (tunnel is active and able to serve traffic), or `down`
	// (tunnel can not serve traffic as it has no connections to the Cloudflare Edge).
	Status param.Field[StatusTunnel] `query:"status"`
	// UUID of the tunnel.
	Uuid          param.Field[string]    `query:"uuid" format:"uuid"`
	WasActiveAt   param.Field[time.Time] `query:"was_active_at" format:"date-time"`
	WasInactiveAt param.Field[time.Time] `query:"was_inactive_at" format:"date-time"`
}

// URLQuery serializes [AccountCfdTunnelListParams]'s query parameters as
// `url.Values`.
func (r AccountCfdTunnelListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountCfdTunnelDeleteParams struct {
	Body interface{} `json:"body,required"`
}

func (r AccountCfdTunnelDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type AccountCfdTunnelGetManagementTokenParams struct {
	Resources param.Field[[]AccountCfdTunnelGetManagementTokenParamsResource] `json:"resources,required"`
}

func (r AccountCfdTunnelGetManagementTokenParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Management resources the token will have access to.
type AccountCfdTunnelGetManagementTokenParamsResource string

const (
	AccountCfdTunnelGetManagementTokenParamsResourceLogs AccountCfdTunnelGetManagementTokenParamsResource = "logs"
)

func (r AccountCfdTunnelGetManagementTokenParamsResource) IsKnown() bool {
	switch r {
	case AccountCfdTunnelGetManagementTokenParamsResourceLogs:
		return true
	}
	return false
}
