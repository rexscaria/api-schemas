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

// ZoneAPIGatewayDiscoveryService contains methods and other services that help
// with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneAPIGatewayDiscoveryService] method instead.
type ZoneAPIGatewayDiscoveryService struct {
	Options    []option.RequestOption
	Operations *ZoneAPIGatewayDiscoveryOperationService
}

// NewZoneAPIGatewayDiscoveryService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneAPIGatewayDiscoveryService(opts ...option.RequestOption) (r *ZoneAPIGatewayDiscoveryService) {
	r = &ZoneAPIGatewayDiscoveryService{}
	r.Options = opts
	r.Operations = NewZoneAPIGatewayDiscoveryOperationService(opts...)
	return
}

// Retrieve the most up to date view of discovered operations, rendered as OpenAPI
// schemas
func (r *ZoneAPIGatewayDiscoveryService) Get(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *ZoneAPIGatewayDiscoveryGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/api_gateway/discovery", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ZoneAPIGatewayDiscoveryGetResponse struct {
	Errors   []MessagesAPIShieldItem                  `json:"errors,required"`
	Messages []MessagesAPIShieldItem                  `json:"messages,required"`
	Result   ZoneAPIGatewayDiscoveryGetResponseResult `json:"result,required"`
	// Whether the API call was successful.
	Success ZoneAPIGatewayDiscoveryGetResponseSuccess `json:"success,required"`
	JSON    zoneAPIGatewayDiscoveryGetResponseJSON    `json:"-"`
}

// zoneAPIGatewayDiscoveryGetResponseJSON contains the JSON metadata for the struct
// [ZoneAPIGatewayDiscoveryGetResponse]
type zoneAPIGatewayDiscoveryGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAPIGatewayDiscoveryGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneAPIGatewayDiscoveryGetResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneAPIGatewayDiscoveryGetResponseResult struct {
	Schemas   []interface{}                                `json:"schemas,required"`
	Timestamp SchemasTimestamp                             `json:"timestamp,required" format:"date-time"`
	JSON      zoneAPIGatewayDiscoveryGetResponseResultJSON `json:"-"`
}

// zoneAPIGatewayDiscoveryGetResponseResultJSON contains the JSON metadata for the
// struct [ZoneAPIGatewayDiscoveryGetResponseResult]
type zoneAPIGatewayDiscoveryGetResponseResultJSON struct {
	Schemas     apijson.Field
	Timestamp   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAPIGatewayDiscoveryGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneAPIGatewayDiscoveryGetResponseResultJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ZoneAPIGatewayDiscoveryGetResponseSuccess bool

const (
	ZoneAPIGatewayDiscoveryGetResponseSuccessTrue ZoneAPIGatewayDiscoveryGetResponseSuccess = true
)

func (r ZoneAPIGatewayDiscoveryGetResponseSuccess) IsKnown() bool {
	switch r {
	case ZoneAPIGatewayDiscoveryGetResponseSuccessTrue:
		return true
	}
	return false
}
