// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
)

// AccountDNSFirewallReverseDNSService contains methods and other services that
// help with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountDNSFirewallReverseDNSService] method instead.
type AccountDNSFirewallReverseDNSService struct {
	Options []option.RequestOption
}

// NewAccountDNSFirewallReverseDNSService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountDNSFirewallReverseDNSService(opts ...option.RequestOption) (r *AccountDNSFirewallReverseDNSService) {
	r = &AccountDNSFirewallReverseDNSService{}
	r.Options = opts
	return
}

// Show reverse DNS configuration (PTR records) for a DNS Firewall cluster
func (r *AccountDNSFirewallReverseDNSService) Get(ctx context.Context, accountID string, dnsFirewallID string, opts ...option.RequestOption) (res *DNSFirewallReverseDNSResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if dnsFirewallID == "" {
		err = errors.New("missing required dns_firewall_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dns_firewall/%s/reverse_dns", accountID, dnsFirewallID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update reverse DNS configuration (PTR records) for a DNS Firewall cluster
func (r *AccountDNSFirewallReverseDNSService) Update(ctx context.Context, accountID string, dnsFirewallID string, body AccountDNSFirewallReverseDNSUpdateParams, opts ...option.RequestOption) (res *DNSFirewallReverseDNSResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if dnsFirewallID == "" {
		err = errors.New("missing required dns_firewall_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dns_firewall/%s/reverse_dns", accountID, dnsFirewallID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

type DNSFirewallReverseDNS struct {
	// Map of cluster IP addresses to PTR record contents
	Ptr  map[string]string         `json:"ptr"`
	JSON dnsFirewallReverseDNSJSON `json:"-"`
}

// dnsFirewallReverseDNSJSON contains the JSON metadata for the struct
// [DNSFirewallReverseDNS]
type dnsFirewallReverseDNSJSON struct {
	Ptr         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSFirewallReverseDNS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsFirewallReverseDNSJSON) RawJSON() string {
	return r.raw
}

type DNSFirewallReverseDNSParam struct {
	// Map of cluster IP addresses to PTR record contents
	Ptr param.Field[map[string]string] `json:"ptr"`
}

func (r DNSFirewallReverseDNSParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DNSFirewallReverseDNSResponse struct {
	Result DNSFirewallReverseDNS             `json:"result"`
	JSON   dnsFirewallReverseDNSResponseJSON `json:"-"`
	APIResponseSingleDNSFirewall
}

// dnsFirewallReverseDNSResponseJSON contains the JSON metadata for the struct
// [DNSFirewallReverseDNSResponse]
type dnsFirewallReverseDNSResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSFirewallReverseDNSResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsFirewallReverseDNSResponseJSON) RawJSON() string {
	return r.raw
}

type AccountDNSFirewallReverseDNSUpdateParams struct {
	DNSFirewallReverseDNS DNSFirewallReverseDNSParam `json:"dns_firewall_reverse_dns,required"`
}

func (r AccountDNSFirewallReverseDNSUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.DNSFirewallReverseDNS)
}
