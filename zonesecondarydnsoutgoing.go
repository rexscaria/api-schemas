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

// ZoneSecondaryDNSOutgoingService contains methods and other services that help
// with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneSecondaryDNSOutgoingService] method instead.
type ZoneSecondaryDNSOutgoingService struct {
	Options []option.RequestOption
}

// NewZoneSecondaryDNSOutgoingService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZoneSecondaryDNSOutgoingService(opts ...option.RequestOption) (r *ZoneSecondaryDNSOutgoingService) {
	r = &ZoneSecondaryDNSOutgoingService{}
	r.Options = opts
	return
}

// Create primary zone configuration for outgoing zone transfers.
func (r *ZoneSecondaryDNSOutgoingService) New(ctx context.Context, zoneID string, body ZoneSecondaryDNSOutgoingNewParams, opts ...option.RequestOption) (res *SingleResponseOutgoing, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/secondary_dns/outgoing", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get primary zone configuration for outgoing zone transfers.
func (r *ZoneSecondaryDNSOutgoingService) Get(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *SingleResponseOutgoing, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/secondary_dns/outgoing", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update primary zone configuration for outgoing zone transfers.
func (r *ZoneSecondaryDNSOutgoingService) Update(ctx context.Context, zoneID string, body ZoneSecondaryDNSOutgoingUpdateParams, opts ...option.RequestOption) (res *SingleResponseOutgoing, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/secondary_dns/outgoing", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Delete primary zone configuration for outgoing zone transfers.
func (r *ZoneSecondaryDNSOutgoingService) Delete(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *IDResponseSecondaryDNS, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/secondary_dns/outgoing", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Disable outgoing zone transfers for primary zone and clears IXFR backlog of
// primary zone.
func (r *ZoneSecondaryDNSOutgoingService) Disable(ctx context.Context, zoneID string, body ZoneSecondaryDNSOutgoingDisableParams, opts ...option.RequestOption) (res *ZoneSecondaryDNSOutgoingDisableResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/secondary_dns/outgoing/disable", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Enable outgoing zone transfers for primary zone.
func (r *ZoneSecondaryDNSOutgoingService) Enable(ctx context.Context, zoneID string, body ZoneSecondaryDNSOutgoingEnableParams, opts ...option.RequestOption) (res *EnableTransferResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/secondary_dns/outgoing/enable", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Notifies the secondary nameserver(s) and clears IXFR backlog of primary zone.
func (r *ZoneSecondaryDNSOutgoingService) ForceNotify(ctx context.Context, zoneID string, body ZoneSecondaryDNSOutgoingForceNotifyParams, opts ...option.RequestOption) (res *ZoneSecondaryDNSOutgoingForceNotifyResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/secondary_dns/outgoing/force_notify", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get primary zone transfer status.
func (r *ZoneSecondaryDNSOutgoingService) Status(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *EnableTransferResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/secondary_dns/outgoing/status", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type EnableTransferResponse struct {
	Errors   []SecondaryDNSMessages `json:"errors,required"`
	Messages []SecondaryDNSMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success EnableTransferResponseSuccess `json:"success,required"`
	// The zone transfer status of a primary zone
	Result string                     `json:"result"`
	JSON   enableTransferResponseJSON `json:"-"`
}

// enableTransferResponseJSON contains the JSON metadata for the struct
// [EnableTransferResponse]
type enableTransferResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EnableTransferResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r enableTransferResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type EnableTransferResponseSuccess bool

const (
	EnableTransferResponseSuccessTrue EnableTransferResponseSuccess = true
)

func (r EnableTransferResponseSuccess) IsKnown() bool {
	switch r {
	case EnableTransferResponseSuccessTrue:
		return true
	}
	return false
}

type SingleRequestParam struct {
	// Zone name.
	Name param.Field[string] `json:"name,required"`
	// A list of peer tags.
	Peers param.Field[[]string] `json:"peers,required"`
}

func (r SingleRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SingleResponseOutgoing struct {
	Errors   []SecondaryDNSMessages `json:"errors,required"`
	Messages []SecondaryDNSMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success SingleResponseOutgoingSuccess `json:"success,required"`
	Result  SingleResponseOutgoingResult  `json:"result"`
	JSON    singleResponseOutgoingJSON    `json:"-"`
}

// singleResponseOutgoingJSON contains the JSON metadata for the struct
// [SingleResponseOutgoing]
type singleResponseOutgoingJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponseOutgoing) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r singleResponseOutgoingJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type SingleResponseOutgoingSuccess bool

const (
	SingleResponseOutgoingSuccessTrue SingleResponseOutgoingSuccess = true
)

func (r SingleResponseOutgoingSuccess) IsKnown() bool {
	switch r {
	case SingleResponseOutgoingSuccessTrue:
		return true
	}
	return false
}

type SingleResponseOutgoingResult struct {
	ID string `json:"id"`
	// The time for a specific event.
	CheckedTime string `json:"checked_time"`
	// The time for a specific event.
	CreatedTime string `json:"created_time"`
	// The time for a specific event.
	LastTransferredTime string `json:"last_transferred_time"`
	// Zone name.
	Name string `json:"name"`
	// A list of peer tags.
	Peers []string `json:"peers"`
	// The serial number of the SOA for the given zone.
	SoaSerial float64                          `json:"soa_serial"`
	JSON      singleResponseOutgoingResultJSON `json:"-"`
}

// singleResponseOutgoingResultJSON contains the JSON metadata for the struct
// [SingleResponseOutgoingResult]
type singleResponseOutgoingResultJSON struct {
	ID                  apijson.Field
	CheckedTime         apijson.Field
	CreatedTime         apijson.Field
	LastTransferredTime apijson.Field
	Name                apijson.Field
	Peers               apijson.Field
	SoaSerial           apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *SingleResponseOutgoingResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r singleResponseOutgoingResultJSON) RawJSON() string {
	return r.raw
}

type ZoneSecondaryDNSOutgoingDisableResponse struct {
	Errors   []SecondaryDNSMessages `json:"errors,required"`
	Messages []SecondaryDNSMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success ZoneSecondaryDNSOutgoingDisableResponseSuccess `json:"success,required"`
	// The zone transfer status of a primary zone
	Result string                                      `json:"result"`
	JSON   zoneSecondaryDNSOutgoingDisableResponseJSON `json:"-"`
}

// zoneSecondaryDNSOutgoingDisableResponseJSON contains the JSON metadata for the
// struct [ZoneSecondaryDNSOutgoingDisableResponse]
type zoneSecondaryDNSOutgoingDisableResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSecondaryDNSOutgoingDisableResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSecondaryDNSOutgoingDisableResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ZoneSecondaryDNSOutgoingDisableResponseSuccess bool

const (
	ZoneSecondaryDNSOutgoingDisableResponseSuccessTrue ZoneSecondaryDNSOutgoingDisableResponseSuccess = true
)

func (r ZoneSecondaryDNSOutgoingDisableResponseSuccess) IsKnown() bool {
	switch r {
	case ZoneSecondaryDNSOutgoingDisableResponseSuccessTrue:
		return true
	}
	return false
}

type ZoneSecondaryDNSOutgoingForceNotifyResponse struct {
	Errors   []SecondaryDNSMessages `json:"errors,required"`
	Messages []SecondaryDNSMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success ZoneSecondaryDNSOutgoingForceNotifyResponseSuccess `json:"success,required"`
	// When force_notify query parameter is set to true, the response is a simple
	// string
	Result string                                          `json:"result"`
	JSON   zoneSecondaryDNSOutgoingForceNotifyResponseJSON `json:"-"`
}

// zoneSecondaryDNSOutgoingForceNotifyResponseJSON contains the JSON metadata for
// the struct [ZoneSecondaryDNSOutgoingForceNotifyResponse]
type zoneSecondaryDNSOutgoingForceNotifyResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSecondaryDNSOutgoingForceNotifyResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSecondaryDNSOutgoingForceNotifyResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ZoneSecondaryDNSOutgoingForceNotifyResponseSuccess bool

const (
	ZoneSecondaryDNSOutgoingForceNotifyResponseSuccessTrue ZoneSecondaryDNSOutgoingForceNotifyResponseSuccess = true
)

func (r ZoneSecondaryDNSOutgoingForceNotifyResponseSuccess) IsKnown() bool {
	switch r {
	case ZoneSecondaryDNSOutgoingForceNotifyResponseSuccessTrue:
		return true
	}
	return false
}

type ZoneSecondaryDNSOutgoingNewParams struct {
	SingleRequest SingleRequestParam `json:"single_request,required"`
}

func (r ZoneSecondaryDNSOutgoingNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.SingleRequest)
}

type ZoneSecondaryDNSOutgoingUpdateParams struct {
	SingleRequest SingleRequestParam `json:"single_request,required"`
}

func (r ZoneSecondaryDNSOutgoingUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.SingleRequest)
}

type ZoneSecondaryDNSOutgoingDisableParams struct {
	Body interface{} `json:"body,required"`
}

func (r ZoneSecondaryDNSOutgoingDisableParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type ZoneSecondaryDNSOutgoingEnableParams struct {
	Body interface{} `json:"body,required"`
}

func (r ZoneSecondaryDNSOutgoingEnableParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type ZoneSecondaryDNSOutgoingForceNotifyParams struct {
	Body interface{} `json:"body,required"`
}

func (r ZoneSecondaryDNSOutgoingForceNotifyParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}
