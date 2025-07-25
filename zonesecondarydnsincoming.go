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

// ZoneSecondaryDNSIncomingService contains methods and other services that help
// with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneSecondaryDNSIncomingService] method instead.
type ZoneSecondaryDNSIncomingService struct {
	Options []option.RequestOption
}

// NewZoneSecondaryDNSIncomingService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZoneSecondaryDNSIncomingService(opts ...option.RequestOption) (r *ZoneSecondaryDNSIncomingService) {
	r = &ZoneSecondaryDNSIncomingService{}
	r.Options = opts
	return
}

// Create secondary zone configuration for incoming zone transfers.
func (r *ZoneSecondaryDNSIncomingService) New(ctx context.Context, zoneID string, body ZoneSecondaryDNSIncomingNewParams, opts ...option.RequestOption) (res *SingleResponseIncoming, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/secondary_dns/incoming", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get secondary zone configuration for incoming zone transfers.
func (r *ZoneSecondaryDNSIncomingService) Get(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *SingleResponseIncoming, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/secondary_dns/incoming", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update secondary zone configuration for incoming zone transfers.
func (r *ZoneSecondaryDNSIncomingService) Update(ctx context.Context, zoneID string, body ZoneSecondaryDNSIncomingUpdateParams, opts ...option.RequestOption) (res *SingleResponseIncoming, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/secondary_dns/incoming", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Delete secondary zone configuration for incoming zone transfers.
func (r *ZoneSecondaryDNSIncomingService) Delete(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *IDResponseSecondaryDNS, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/secondary_dns/incoming", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type IDResponseSecondaryDNS struct {
	Errors   []SecondaryDNSMessages `json:"errors,required"`
	Messages []SecondaryDNSMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success IDResponseSecondaryDNSSuccess `json:"success,required"`
	Result  IDResponseSecondaryDNSResult  `json:"result"`
	JSON    idResponseSecondaryDNSJSON    `json:"-"`
}

// idResponseSecondaryDNSJSON contains the JSON metadata for the struct
// [IDResponseSecondaryDNS]
type idResponseSecondaryDNSJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IDResponseSecondaryDNS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r idResponseSecondaryDNSJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type IDResponseSecondaryDNSSuccess bool

const (
	IDResponseSecondaryDNSSuccessTrue IDResponseSecondaryDNSSuccess = true
)

func (r IDResponseSecondaryDNSSuccess) IsKnown() bool {
	switch r {
	case IDResponseSecondaryDNSSuccessTrue:
		return true
	}
	return false
}

type IDResponseSecondaryDNSResult struct {
	ID   string                           `json:"id"`
	JSON idResponseSecondaryDNSResultJSON `json:"-"`
}

// idResponseSecondaryDNSResultJSON contains the JSON metadata for the struct
// [IDResponseSecondaryDNSResult]
type idResponseSecondaryDNSResultJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IDResponseSecondaryDNSResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r idResponseSecondaryDNSResultJSON) RawJSON() string {
	return r.raw
}

type SecondaryZoneParam struct {
	// How often should a secondary zone auto refresh regardless of DNS NOTIFY. Not
	// applicable for primary zones.
	AutoRefreshSeconds param.Field[float64] `json:"auto_refresh_seconds,required"`
	// Zone name.
	Name param.Field[string] `json:"name,required"`
	// A list of peer tags.
	Peers param.Field[[]string] `json:"peers,required"`
}

func (r SecondaryZoneParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SingleResponseIncoming struct {
	Errors   []SecondaryDNSMessages `json:"errors,required"`
	Messages []SecondaryDNSMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success SingleResponseIncomingSuccess `json:"success,required"`
	Result  SingleResponseIncomingResult  `json:"result"`
	JSON    singleResponseIncomingJSON    `json:"-"`
}

// singleResponseIncomingJSON contains the JSON metadata for the struct
// [SingleResponseIncoming]
type singleResponseIncomingJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponseIncoming) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r singleResponseIncomingJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type SingleResponseIncomingSuccess bool

const (
	SingleResponseIncomingSuccessTrue SingleResponseIncomingSuccess = true
)

func (r SingleResponseIncomingSuccess) IsKnown() bool {
	switch r {
	case SingleResponseIncomingSuccessTrue:
		return true
	}
	return false
}

type SingleResponseIncomingResult struct {
	ID string `json:"id"`
	// How often should a secondary zone auto refresh regardless of DNS NOTIFY. Not
	// applicable for primary zones.
	AutoRefreshSeconds float64 `json:"auto_refresh_seconds"`
	// The time for a specific event.
	CheckedTime string `json:"checked_time"`
	// The time for a specific event.
	CreatedTime string `json:"created_time"`
	// The time for a specific event.
	ModifiedTime string `json:"modified_time"`
	// Zone name.
	Name string `json:"name"`
	// A list of peer tags.
	Peers []string `json:"peers"`
	// The serial number of the SOA for the given zone.
	SoaSerial float64                          `json:"soa_serial"`
	JSON      singleResponseIncomingResultJSON `json:"-"`
}

// singleResponseIncomingResultJSON contains the JSON metadata for the struct
// [SingleResponseIncomingResult]
type singleResponseIncomingResultJSON struct {
	ID                 apijson.Field
	AutoRefreshSeconds apijson.Field
	CheckedTime        apijson.Field
	CreatedTime        apijson.Field
	ModifiedTime       apijson.Field
	Name               apijson.Field
	Peers              apijson.Field
	SoaSerial          apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *SingleResponseIncomingResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r singleResponseIncomingResultJSON) RawJSON() string {
	return r.raw
}

type ZoneSecondaryDNSIncomingNewParams struct {
	SecondaryZone SecondaryZoneParam `json:"secondary_zone,required"`
}

func (r ZoneSecondaryDNSIncomingNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.SecondaryZone)
}

type ZoneSecondaryDNSIncomingUpdateParams struct {
	SecondaryZone SecondaryZoneParam `json:"secondary_zone,required"`
}

func (r ZoneSecondaryDNSIncomingUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.SecondaryZone)
}
