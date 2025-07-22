// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
	"github.com/rexscaria/api-schemas/shared"
	"github.com/tidwall/gjson"
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
	Code    int64                        `json:"code,required"`
	Message string                       `json:"message,required"`
	JSON    messagesMagicTransitItemJSON `json:"-"`
}

// messagesMagicTransitItemJSON contains the JSON metadata for the struct
// [MessagesMagicTransitItem]
type messagesMagicTransitItemJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessagesMagicTransitItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messagesMagicTransitItemJSON) RawJSON() string {
	return r.raw
}

type AccountDiagnosticRunTracerouteResponse struct {
	Errors   []MessagesMagicTransitItem                        `json:"errors,required"`
	Messages []MessagesMagicTransitItem                        `json:"messages,required"`
	Result   AccountDiagnosticRunTracerouteResponseResultUnion `json:"result,required,nullable"`
	// Whether the API call was successful
	Success AccountDiagnosticRunTracerouteResponseSuccess `json:"success,required"`
	JSON    accountDiagnosticRunTracerouteResponseJSON    `json:"-"`
}

// accountDiagnosticRunTracerouteResponseJSON contains the JSON metadata for the
// struct [AccountDiagnosticRunTracerouteResponse]
type accountDiagnosticRunTracerouteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDiagnosticRunTracerouteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountDiagnosticRunTracerouteResponseJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [AccountDiagnosticRunTracerouteResponseResultArray] or
// [shared.UnionString].
type AccountDiagnosticRunTracerouteResponseResultUnion interface {
	ImplementsAccountDiagnosticRunTracerouteResponseResultUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccountDiagnosticRunTracerouteResponseResultUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccountDiagnosticRunTracerouteResponseResultArray{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type AccountDiagnosticRunTracerouteResponseResultArray []interface{}

func (r AccountDiagnosticRunTracerouteResponseResultArray) ImplementsAccountDiagnosticRunTracerouteResponseResultUnion() {
}

// Whether the API call was successful
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
