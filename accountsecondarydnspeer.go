// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/cf-rex-go/internal/apijson"
	"github.com/stainless-sdks/cf-rex-go/internal/param"
	"github.com/stainless-sdks/cf-rex-go/internal/requestconfig"
	"github.com/stainless-sdks/cf-rex-go/option"
)

// AccountSecondaryDNSPeerService contains methods and other services that help
// with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountSecondaryDNSPeerService] method instead.
type AccountSecondaryDNSPeerService struct {
	Options []option.RequestOption
}

// NewAccountSecondaryDNSPeerService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountSecondaryDNSPeerService(opts ...option.RequestOption) (r *AccountSecondaryDNSPeerService) {
	r = &AccountSecondaryDNSPeerService{}
	r.Options = opts
	return
}

// Create Peer.
func (r *AccountSecondaryDNSPeerService) New(ctx context.Context, accountID string, body AccountSecondaryDNSPeerNewParams, opts ...option.RequestOption) (res *SchemasSecondaryDNSPeersSingleResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/secondary_dns/peers", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get Peer.
func (r *AccountSecondaryDNSPeerService) Get(ctx context.Context, accountID string, peerID string, opts ...option.RequestOption) (res *SchemasSecondaryDNSPeersSingleResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if peerID == "" {
		err = errors.New("missing required peer_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/secondary_dns/peers/%s", accountID, peerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Modify Peer.
func (r *AccountSecondaryDNSPeerService) Update(ctx context.Context, accountID string, peerID string, body AccountSecondaryDNSPeerUpdateParams, opts ...option.RequestOption) (res *SchemasSecondaryDNSPeersSingleResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if peerID == "" {
		err = errors.New("missing required peer_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/secondary_dns/peers/%s", accountID, peerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// List Peers.
func (r *AccountSecondaryDNSPeerService) List(ctx context.Context, accountID string, opts ...option.RequestOption) (res *AccountSecondaryDNSPeerListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/secondary_dns/peers", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete Peer.
func (r *AccountSecondaryDNSPeerService) Delete(ctx context.Context, accountID string, peerID string, body AccountSecondaryDNSPeerDeleteParams, opts ...option.RequestOption) (res *SchemasIDResponseSecondaryDNS, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if peerID == "" {
		err = errors.New("missing required peer_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/secondary_dns/peers/%s", accountID, peerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

type Peer struct {
	ID string `json:"id,required"`
	// The name of the peer.
	Name string `json:"name,required"`
	// IPv4/IPv6 address of primary or secondary nameserver, depending on what zone
	// this peer is linked to. For primary zones this IP defines the IP of the
	// secondary nameserver Cloudflare will NOTIFY upon zone changes. For secondary
	// zones this IP defines the IP of the primary nameserver Cloudflare will send
	// AXFR/IXFR requests to.
	IP string `json:"ip"`
	// Enable IXFR transfer protocol, default is AXFR. Only applicable to secondary
	// zones.
	IxfrEnable bool `json:"ixfr_enable"`
	// DNS port of primary or secondary nameserver, depending on what zone this peer is
	// linked to.
	Port float64 `json:"port"`
	// TSIG authentication will be used for zone transfer if configured.
	TsigID string   `json:"tsig_id"`
	JSON   peerJSON `json:"-"`
}

// peerJSON contains the JSON metadata for the struct [Peer]
type peerJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	IP          apijson.Field
	IxfrEnable  apijson.Field
	Port        apijson.Field
	TsigID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Peer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r peerJSON) RawJSON() string {
	return r.raw
}

type PeerParam struct {
	// The name of the peer.
	Name param.Field[string] `json:"name,required"`
	// IPv4/IPv6 address of primary or secondary nameserver, depending on what zone
	// this peer is linked to. For primary zones this IP defines the IP of the
	// secondary nameserver Cloudflare will NOTIFY upon zone changes. For secondary
	// zones this IP defines the IP of the primary nameserver Cloudflare will send
	// AXFR/IXFR requests to.
	IP param.Field[string] `json:"ip"`
	// Enable IXFR transfer protocol, default is AXFR. Only applicable to secondary
	// zones.
	IxfrEnable param.Field[bool] `json:"ixfr_enable"`
	// DNS port of primary or secondary nameserver, depending on what zone this peer is
	// linked to.
	Port param.Field[float64] `json:"port"`
	// TSIG authentication will be used for zone transfer if configured.
	TsigID param.Field[string] `json:"tsig_id"`
}

func (r PeerParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SchemasSecondaryDNSPeersSingleResponse struct {
	Result Peer                                       `json:"result"`
	JSON   schemasSecondaryDNSPeersSingleResponseJSON `json:"-"`
	ResponseSingleACL
}

// schemasSecondaryDNSPeersSingleResponseJSON contains the JSON metadata for the
// struct [SchemasSecondaryDNSPeersSingleResponse]
type schemasSecondaryDNSPeersSingleResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasSecondaryDNSPeersSingleResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r schemasSecondaryDNSPeersSingleResponseJSON) RawJSON() string {
	return r.raw
}

type AccountSecondaryDNSPeerListResponse struct {
	Result []Peer                                  `json:"result"`
	JSON   accountSecondaryDNSPeerListResponseJSON `json:"-"`
	ResponseCollectionDNSACLs
}

// accountSecondaryDNSPeerListResponseJSON contains the JSON metadata for the
// struct [AccountSecondaryDNSPeerListResponse]
type accountSecondaryDNSPeerListResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSecondaryDNSPeerListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountSecondaryDNSPeerListResponseJSON) RawJSON() string {
	return r.raw
}

type AccountSecondaryDNSPeerNewParams struct {
	// The name of the peer.
	Name param.Field[string] `json:"name,required"`
}

func (r AccountSecondaryDNSPeerNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountSecondaryDNSPeerUpdateParams struct {
	Peer PeerParam `json:"peer,required"`
}

func (r AccountSecondaryDNSPeerUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Peer)
}

type AccountSecondaryDNSPeerDeleteParams struct {
	Body interface{} `json:"body,required"`
}

func (r AccountSecondaryDNSPeerDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}
