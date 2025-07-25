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

// AccountDiagnosticService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountDiagnosticService] method instead.
type AccountDiagnosticService struct {
	Options []option.RequestOption
}

// NewAccountDiagnosticService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountDiagnosticService(opts ...option.RequestOption) (r *AccountDiagnosticService) {
	r = &AccountDiagnosticService{}
	r.Options = opts
	return
}

// Run traceroutes from Cloudflare colos.
func (r *AccountDiagnosticService) RunTraceroute(ctx context.Context, accountID string, body AccountDiagnosticRunTracerouteParams, opts ...option.RequestOption) (res *AccountDiagnosticRunTracerouteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/diagnostics/traceroute", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type MessagesMagicTransitItem struct {
	Code             int64                          `json:"code,required"`
	Message          string                         `json:"message,required"`
	DocumentationURL string                         `json:"documentation_url"`
	Source           MessagesMagicTransitItemSource `json:"source"`
	JSON             messagesMagicTransitItemJSON   `json:"-"`
}

// messagesMagicTransitItemJSON contains the JSON metadata for the struct
// [MessagesMagicTransitItem]
type messagesMagicTransitItemJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *MessagesMagicTransitItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messagesMagicTransitItemJSON) RawJSON() string {
	return r.raw
}

type MessagesMagicTransitItemSource struct {
	Pointer string                             `json:"pointer"`
	JSON    messagesMagicTransitItemSourceJSON `json:"-"`
}

// messagesMagicTransitItemSourceJSON contains the JSON metadata for the struct
// [MessagesMagicTransitItemSource]
type messagesMagicTransitItemSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessagesMagicTransitItemSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messagesMagicTransitItemSourceJSON) RawJSON() string {
	return r.raw
}

type AccountDiagnosticRunTracerouteResponse struct {
	Errors   []MessagesMagicTransitItem `json:"errors,required"`
	Messages []MessagesMagicTransitItem `json:"messages,required"`
	// Whether the API call was successful.
	Success AccountDiagnosticRunTracerouteResponseSuccess  `json:"success,required"`
	Result  []AccountDiagnosticRunTracerouteResponseResult `json:"result"`
	JSON    accountDiagnosticRunTracerouteResponseJSON     `json:"-"`
}

// accountDiagnosticRunTracerouteResponseJSON contains the JSON metadata for the
// struct [AccountDiagnosticRunTracerouteResponse]
type accountDiagnosticRunTracerouteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDiagnosticRunTracerouteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountDiagnosticRunTracerouteResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccountDiagnosticRunTracerouteResponseSuccess bool

const (
	AccountDiagnosticRunTracerouteResponseSuccessTrue AccountDiagnosticRunTracerouteResponseSuccess = true
)

func (r AccountDiagnosticRunTracerouteResponseSuccess) IsKnown() bool {
	switch r {
	case AccountDiagnosticRunTracerouteResponseSuccessTrue:
		return true
	}
	return false
}

type AccountDiagnosticRunTracerouteResponseResult struct {
	Colos []AccountDiagnosticRunTracerouteResponseResultColo `json:"colos"`
	// The target hostname, IPv6, or IPv6 address.
	Target string                                           `json:"target"`
	JSON   accountDiagnosticRunTracerouteResponseResultJSON `json:"-"`
}

// accountDiagnosticRunTracerouteResponseResultJSON contains the JSON metadata for
// the struct [AccountDiagnosticRunTracerouteResponseResult]
type accountDiagnosticRunTracerouteResponseResultJSON struct {
	Colos       apijson.Field
	Target      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDiagnosticRunTracerouteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountDiagnosticRunTracerouteResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountDiagnosticRunTracerouteResponseResultColo struct {
	Colo AccountDiagnosticRunTracerouteResponseResultColosColo `json:"colo"`
	// Errors resulting from collecting traceroute from colo to target.
	Error AccountDiagnosticRunTracerouteResponseResultColosError `json:"error"`
	Hops  []AccountDiagnosticRunTracerouteResponseResultColosHop `json:"hops"`
	// Aggregated statistics from all hops about the target.
	TargetSummary interface{} `json:"target_summary"`
	// Total time of traceroute in ms.
	TracerouteTimeMs int64                                                `json:"traceroute_time_ms"`
	JSON             accountDiagnosticRunTracerouteResponseResultColoJSON `json:"-"`
}

// accountDiagnosticRunTracerouteResponseResultColoJSON contains the JSON metadata
// for the struct [AccountDiagnosticRunTracerouteResponseResultColo]
type accountDiagnosticRunTracerouteResponseResultColoJSON struct {
	Colo             apijson.Field
	Error            apijson.Field
	Hops             apijson.Field
	TargetSummary    apijson.Field
	TracerouteTimeMs apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountDiagnosticRunTracerouteResponseResultColo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountDiagnosticRunTracerouteResponseResultColoJSON) RawJSON() string {
	return r.raw
}

type AccountDiagnosticRunTracerouteResponseResultColosColo struct {
	// Source colo city.
	City string `json:"city"`
	// Source colo name.
	Name string                                                    `json:"name"`
	JSON accountDiagnosticRunTracerouteResponseResultColosColoJSON `json:"-"`
}

// accountDiagnosticRunTracerouteResponseResultColosColoJSON contains the JSON
// metadata for the struct [AccountDiagnosticRunTracerouteResponseResultColosColo]
type accountDiagnosticRunTracerouteResponseResultColosColoJSON struct {
	City        apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDiagnosticRunTracerouteResponseResultColosColo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountDiagnosticRunTracerouteResponseResultColosColoJSON) RawJSON() string {
	return r.raw
}

// Errors resulting from collecting traceroute from colo to target.
type AccountDiagnosticRunTracerouteResponseResultColosError string

const (
	AccountDiagnosticRunTracerouteResponseResultColosErrorEmpty                             AccountDiagnosticRunTracerouteResponseResultColosError = ""
	AccountDiagnosticRunTracerouteResponseResultColosErrorCouldNotGatherTracerouteDataCode1 AccountDiagnosticRunTracerouteResponseResultColosError = "Could not gather traceroute data: Code 1"
	AccountDiagnosticRunTracerouteResponseResultColosErrorCouldNotGatherTracerouteDataCode2 AccountDiagnosticRunTracerouteResponseResultColosError = "Could not gather traceroute data: Code 2"
	AccountDiagnosticRunTracerouteResponseResultColosErrorCouldNotGatherTracerouteDataCode3 AccountDiagnosticRunTracerouteResponseResultColosError = "Could not gather traceroute data: Code 3"
	AccountDiagnosticRunTracerouteResponseResultColosErrorCouldNotGatherTracerouteDataCode4 AccountDiagnosticRunTracerouteResponseResultColosError = "Could not gather traceroute data: Code 4"
)

func (r AccountDiagnosticRunTracerouteResponseResultColosError) IsKnown() bool {
	switch r {
	case AccountDiagnosticRunTracerouteResponseResultColosErrorEmpty, AccountDiagnosticRunTracerouteResponseResultColosErrorCouldNotGatherTracerouteDataCode1, AccountDiagnosticRunTracerouteResponseResultColosErrorCouldNotGatherTracerouteDataCode2, AccountDiagnosticRunTracerouteResponseResultColosErrorCouldNotGatherTracerouteDataCode3, AccountDiagnosticRunTracerouteResponseResultColosErrorCouldNotGatherTracerouteDataCode4:
		return true
	}
	return false
}

type AccountDiagnosticRunTracerouteResponseResultColosHop struct {
	// An array of node objects.
	Nodes []AccountDiagnosticRunTracerouteResponseResultColosHopsNode `json:"nodes"`
	// Number of packets where no response was received.
	PacketsLost int64 `json:"packets_lost"`
	// Number of packets sent with specified TTL.
	PacketsSent int64 `json:"packets_sent"`
	// The time to live (TTL).
	PacketsTtl int64                                                    `json:"packets_ttl"`
	JSON       accountDiagnosticRunTracerouteResponseResultColosHopJSON `json:"-"`
}

// accountDiagnosticRunTracerouteResponseResultColosHopJSON contains the JSON
// metadata for the struct [AccountDiagnosticRunTracerouteResponseResultColosHop]
type accountDiagnosticRunTracerouteResponseResultColosHopJSON struct {
	Nodes       apijson.Field
	PacketsLost apijson.Field
	PacketsSent apijson.Field
	PacketsTtl  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDiagnosticRunTracerouteResponseResultColosHop) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountDiagnosticRunTracerouteResponseResultColosHopJSON) RawJSON() string {
	return r.raw
}

type AccountDiagnosticRunTracerouteResponseResultColosHopsNode struct {
	// AS number associated with the node object.
	Asn string `json:"asn"`
	// IP address of the node.
	IP string `json:"ip"`
	// Field appears if there is an additional annotation printed when the probe
	// returns. Field also appears when running a GRE+ICMP traceroute to denote which
	// traceroute a node comes from.
	Labels []string `json:"labels"`
	// Maximum RTT in ms.
	MaxRttMs float64 `json:"max_rtt_ms"`
	// Mean RTT in ms.
	MeanRttMs float64 `json:"mean_rtt_ms"`
	// Minimum RTT in ms.
	MinRttMs float64 `json:"min_rtt_ms"`
	// Host name of the address, this may be the same as the IP address.
	Name string `json:"name"`
	// Number of packets with a response from this node.
	PacketCount int64 `json:"packet_count"`
	// Standard deviation of the RTTs in ms.
	StdDevRttMs float64                                                       `json:"std_dev_rtt_ms"`
	JSON        accountDiagnosticRunTracerouteResponseResultColosHopsNodeJSON `json:"-"`
}

// accountDiagnosticRunTracerouteResponseResultColosHopsNodeJSON contains the JSON
// metadata for the struct
// [AccountDiagnosticRunTracerouteResponseResultColosHopsNode]
type accountDiagnosticRunTracerouteResponseResultColosHopsNodeJSON struct {
	Asn         apijson.Field
	IP          apijson.Field
	Labels      apijson.Field
	MaxRttMs    apijson.Field
	MeanRttMs   apijson.Field
	MinRttMs    apijson.Field
	Name        apijson.Field
	PacketCount apijson.Field
	StdDevRttMs apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDiagnosticRunTracerouteResponseResultColosHopsNode) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountDiagnosticRunTracerouteResponseResultColosHopsNodeJSON) RawJSON() string {
	return r.raw
}

type AccountDiagnosticRunTracerouteParams struct {
	Targets param.Field[[]string] `json:"targets,required"`
	// If no source colo names specified, all colos will be used. China colos are
	// unavailable for traceroutes.
	Colos   param.Field[[]string]                                    `json:"colos"`
	Options param.Field[AccountDiagnosticRunTracerouteParamsOptions] `json:"options"`
}

func (r AccountDiagnosticRunTracerouteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountDiagnosticRunTracerouteParamsOptions struct {
	// Max TTL.
	MaxTtl param.Field[int64] `json:"max_ttl"`
	// Type of packet sent.
	PacketType param.Field[AccountDiagnosticRunTracerouteParamsOptionsPacketType] `json:"packet_type"`
	// Number of packets sent at each TTL.
	PacketsPerTtl param.Field[int64] `json:"packets_per_ttl"`
	// For UDP and TCP, specifies the destination port. For ICMP, specifies the initial
	// ICMP sequence value. Default value 0 will choose the best value to use for each
	// protocol.
	Port param.Field[int64] `json:"port"`
	// Set the time (in seconds) to wait for a response to a probe.
	WaitTime param.Field[int64] `json:"wait_time"`
}

func (r AccountDiagnosticRunTracerouteParamsOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Type of packet sent.
type AccountDiagnosticRunTracerouteParamsOptionsPacketType string

const (
	AccountDiagnosticRunTracerouteParamsOptionsPacketTypeIcmp    AccountDiagnosticRunTracerouteParamsOptionsPacketType = "icmp"
	AccountDiagnosticRunTracerouteParamsOptionsPacketTypeTcp     AccountDiagnosticRunTracerouteParamsOptionsPacketType = "tcp"
	AccountDiagnosticRunTracerouteParamsOptionsPacketTypeUdp     AccountDiagnosticRunTracerouteParamsOptionsPacketType = "udp"
	AccountDiagnosticRunTracerouteParamsOptionsPacketTypeGre     AccountDiagnosticRunTracerouteParamsOptionsPacketType = "gre"
	AccountDiagnosticRunTracerouteParamsOptionsPacketTypeGreIcmp AccountDiagnosticRunTracerouteParamsOptionsPacketType = "gre+icmp"
)

func (r AccountDiagnosticRunTracerouteParamsOptionsPacketType) IsKnown() bool {
	switch r {
	case AccountDiagnosticRunTracerouteParamsOptionsPacketTypeIcmp, AccountDiagnosticRunTracerouteParamsOptionsPacketTypeTcp, AccountDiagnosticRunTracerouteParamsOptionsPacketTypeUdp, AccountDiagnosticRunTracerouteParamsOptionsPacketTypeGre, AccountDiagnosticRunTracerouteParamsOptionsPacketTypeGreIcmp:
		return true
	}
	return false
}
