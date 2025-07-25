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

// ZoneSecondaryDNSService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneSecondaryDNSService] method instead.
type ZoneSecondaryDNSService struct {
	Options  []option.RequestOption
	Incoming *ZoneSecondaryDNSIncomingService
	Outgoing *ZoneSecondaryDNSOutgoingService
}

// NewZoneSecondaryDNSService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneSecondaryDNSService(opts ...option.RequestOption) (r *ZoneSecondaryDNSService) {
	r = &ZoneSecondaryDNSService{}
	r.Options = opts
	r.Incoming = NewZoneSecondaryDNSIncomingService(opts...)
	r.Outgoing = NewZoneSecondaryDNSOutgoingService(opts...)
	return
}

// Sends AXFR zone transfer request to primary nameserver(s).
func (r *ZoneSecondaryDNSService) ForceAxfr(ctx context.Context, zoneID string, body ZoneSecondaryDNSForceAxfrParams, opts ...option.RequestOption) (res *ZoneSecondaryDNSForceAxfrResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/secondary_dns/force_axfr", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type ZoneSecondaryDNSForceAxfrResponse struct {
	Errors   []SecondaryDNSMessages `json:"errors,required"`
	Messages []SecondaryDNSMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success ZoneSecondaryDNSForceAxfrResponseSuccess `json:"success,required"`
	// When force_axfr query parameter is set to true, the response is a simple string
	Result string                                `json:"result"`
	JSON   zoneSecondaryDNSForceAxfrResponseJSON `json:"-"`
}

// zoneSecondaryDNSForceAxfrResponseJSON contains the JSON metadata for the struct
// [ZoneSecondaryDNSForceAxfrResponse]
type zoneSecondaryDNSForceAxfrResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSecondaryDNSForceAxfrResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSecondaryDNSForceAxfrResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ZoneSecondaryDNSForceAxfrResponseSuccess bool

const (
	ZoneSecondaryDNSForceAxfrResponseSuccessTrue ZoneSecondaryDNSForceAxfrResponseSuccess = true
)

func (r ZoneSecondaryDNSForceAxfrResponseSuccess) IsKnown() bool {
	switch r {
	case ZoneSecondaryDNSForceAxfrResponseSuccessTrue:
		return true
	}
	return false
}

type ZoneSecondaryDNSForceAxfrParams struct {
	Body interface{} `json:"body,required"`
}

func (r ZoneSecondaryDNSForceAxfrParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}
