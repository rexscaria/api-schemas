// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/cf-rex-go/internal/apijson"
	"github.com/stainless-sdks/cf-rex-go/internal/requestconfig"
	"github.com/stainless-sdks/cf-rex-go/option"
)

// AccountAddressingAddressMapZoneService contains methods and other services that
// help with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountAddressingAddressMapZoneService] method instead.
type AccountAddressingAddressMapZoneService struct {
	Options []option.RequestOption
}

// NewAccountAddressingAddressMapZoneService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountAddressingAddressMapZoneService(opts ...option.RequestOption) (r *AccountAddressingAddressMapZoneService) {
	r = &AccountAddressingAddressMapZoneService{}
	r.Options = opts
	return
}

// Add a zone as a member of a particular address map.
func (r *AccountAddressingAddressMapZoneService) Add(ctx context.Context, accountID string, addressMapID string, zoneID string, body AccountAddressingAddressMapZoneAddParams, opts ...option.RequestOption) (res *APIResponseCollectionAddressing, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if addressMapID == "" {
		err = errors.New("missing required address_map_id parameter")
		return
	}
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/addressing/address_maps/%s/zones/%s", accountID, addressMapID, zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Remove a zone as a member of a particular address map.
func (r *AccountAddressingAddressMapZoneService) Remove(ctx context.Context, accountID string, addressMapID string, zoneID string, body AccountAddressingAddressMapZoneRemoveParams, opts ...option.RequestOption) (res *APIResponseCollectionAddressing, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if addressMapID == "" {
		err = errors.New("missing required address_map_id parameter")
		return
	}
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/addressing/address_maps/%s/zones/%s", accountID, addressMapID, zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

type AccountAddressingAddressMapZoneAddParams struct {
	Body interface{} `json:"body,required"`
}

func (r AccountAddressingAddressMapZoneAddParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type AccountAddressingAddressMapZoneRemoveParams struct {
	Body interface{} `json:"body,required"`
}

func (r AccountAddressingAddressMapZoneRemoveParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}
