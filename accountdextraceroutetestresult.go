// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
)

// AccountDexTracerouteTestResultService contains methods and other services that
// help with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountDexTracerouteTestResultService] method instead.
type AccountDexTracerouteTestResultService struct {
	Options []option.RequestOption
}

// NewAccountDexTracerouteTestResultService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountDexTracerouteTestResultService(opts ...option.RequestOption) (r *AccountDexTracerouteTestResultService) {
	r = &AccountDexTracerouteTestResultService{}
	r.Options = opts
	return
}

// Get a breakdown of hops and performance metrics for a specific traceroute test
// run
func (r *AccountDexTracerouteTestResultService) GetNetworkPath(ctx context.Context, accountID string, testResultID string, opts ...option.RequestOption) (res *AccountDexTracerouteTestResultGetNetworkPathResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if testResultID == "" {
		err = errors.New("missing required test_result_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dex/traceroute-test-results/%s/network-path", accountID, testResultID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AccountDexTracerouteTestResultGetNetworkPathResponse struct {
	Result AccountDexTracerouteTestResultGetNetworkPathResponseResult `json:"result"`
	JSON   accountDexTracerouteTestResultGetNetworkPathResponseJSON   `json:"-"`
	APIResponseSingleDigitalExperience
}

// accountDexTracerouteTestResultGetNetworkPathResponseJSON contains the JSON
// metadata for the struct [AccountDexTracerouteTestResultGetNetworkPathResponse]
type accountDexTracerouteTestResultGetNetworkPathResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDexTracerouteTestResultGetNetworkPathResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountDexTracerouteTestResultGetNetworkPathResponseJSON) RawJSON() string {
	return r.raw
}

type AccountDexTracerouteTestResultGetNetworkPathResponseResult struct {
	// an array of the hops taken by the device to reach the end destination
	Hops []AccountDexTracerouteTestResultGetNetworkPathResponseResultHop `json:"hops,required"`
	// API Resource UUID tag.
	ResultID string `json:"resultId,required"`
	// name of the device associated with this network path response
	DeviceName string `json:"deviceName"`
	// API Resource UUID tag.
	TestID string `json:"testId"`
	// name of the tracroute test
	TestName string                                                         `json:"testName"`
	JSON     accountDexTracerouteTestResultGetNetworkPathResponseResultJSON `json:"-"`
}

// accountDexTracerouteTestResultGetNetworkPathResponseResultJSON contains the JSON
// metadata for the struct
// [AccountDexTracerouteTestResultGetNetworkPathResponseResult]
type accountDexTracerouteTestResultGetNetworkPathResponseResultJSON struct {
	Hops        apijson.Field
	ResultID    apijson.Field
	DeviceName  apijson.Field
	TestID      apijson.Field
	TestName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDexTracerouteTestResultGetNetworkPathResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountDexTracerouteTestResultGetNetworkPathResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountDexTracerouteTestResultGetNetworkPathResponseResultHop struct {
	Ttl           int64                                                                  `json:"ttl,required"`
	Asn           int64                                                                  `json:"asn,nullable"`
	Aso           string                                                                 `json:"aso,nullable"`
	IPAddress     string                                                                 `json:"ipAddress,nullable"`
	Location      AccountDexTracerouteTestResultGetNetworkPathResponseResultHopsLocation `json:"location,nullable"`
	Mile          AccountDexTracerouteTestResultGetNetworkPathResponseResultHopsMile     `json:"mile,nullable"`
	Name          string                                                                 `json:"name,nullable"`
	PacketLossPct float64                                                                `json:"packetLossPct,nullable"`
	RttMs         int64                                                                  `json:"rttMs,nullable"`
	JSON          accountDexTracerouteTestResultGetNetworkPathResponseResultHopJSON      `json:"-"`
}

// accountDexTracerouteTestResultGetNetworkPathResponseResultHopJSON contains the
// JSON metadata for the struct
// [AccountDexTracerouteTestResultGetNetworkPathResponseResultHop]
type accountDexTracerouteTestResultGetNetworkPathResponseResultHopJSON struct {
	Ttl           apijson.Field
	Asn           apijson.Field
	Aso           apijson.Field
	IPAddress     apijson.Field
	Location      apijson.Field
	Mile          apijson.Field
	Name          apijson.Field
	PacketLossPct apijson.Field
	RttMs         apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccountDexTracerouteTestResultGetNetworkPathResponseResultHop) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountDexTracerouteTestResultGetNetworkPathResponseResultHopJSON) RawJSON() string {
	return r.raw
}

type AccountDexTracerouteTestResultGetNetworkPathResponseResultHopsLocation struct {
	City  string                                                                     `json:"city,nullable"`
	State string                                                                     `json:"state,nullable"`
	Zip   string                                                                     `json:"zip,nullable"`
	JSON  accountDexTracerouteTestResultGetNetworkPathResponseResultHopsLocationJSON `json:"-"`
}

// accountDexTracerouteTestResultGetNetworkPathResponseResultHopsLocationJSON
// contains the JSON metadata for the struct
// [AccountDexTracerouteTestResultGetNetworkPathResponseResultHopsLocation]
type accountDexTracerouteTestResultGetNetworkPathResponseResultHopsLocationJSON struct {
	City        apijson.Field
	State       apijson.Field
	Zip         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDexTracerouteTestResultGetNetworkPathResponseResultHopsLocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountDexTracerouteTestResultGetNetworkPathResponseResultHopsLocationJSON) RawJSON() string {
	return r.raw
}

type AccountDexTracerouteTestResultGetNetworkPathResponseResultHopsMile string

const (
	AccountDexTracerouteTestResultGetNetworkPathResponseResultHopsMileClientToApp       AccountDexTracerouteTestResultGetNetworkPathResponseResultHopsMile = "client-to-app"
	AccountDexTracerouteTestResultGetNetworkPathResponseResultHopsMileClientToCfEgress  AccountDexTracerouteTestResultGetNetworkPathResponseResultHopsMile = "client-to-cf-egress"
	AccountDexTracerouteTestResultGetNetworkPathResponseResultHopsMileClientToCfIngress AccountDexTracerouteTestResultGetNetworkPathResponseResultHopsMile = "client-to-cf-ingress"
	AccountDexTracerouteTestResultGetNetworkPathResponseResultHopsMileClientToIsp       AccountDexTracerouteTestResultGetNetworkPathResponseResultHopsMile = "client-to-isp"
)

func (r AccountDexTracerouteTestResultGetNetworkPathResponseResultHopsMile) IsKnown() bool {
	switch r {
	case AccountDexTracerouteTestResultGetNetworkPathResponseResultHopsMileClientToApp, AccountDexTracerouteTestResultGetNetworkPathResponseResultHopsMileClientToCfEgress, AccountDexTracerouteTestResultGetNetworkPathResponseResultHopsMileClientToCfIngress, AccountDexTracerouteTestResultGetNetworkPathResponseResultHopsMileClientToIsp:
		return true
	}
	return false
}
