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

// AccountAddressingAddressMapIPService contains methods and other services that
// help with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountAddressingAddressMapIPService] method instead.
type AccountAddressingAddressMapIPService struct {
	Options []option.RequestOption
}

// NewAccountAddressingAddressMapIPService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountAddressingAddressMapIPService(opts ...option.RequestOption) (r *AccountAddressingAddressMapIPService) {
	r = &AccountAddressingAddressMapIPService{}
	r.Options = opts
	return
}

// Add an IP from a prefix owned by the account to a particular address map.
func (r *AccountAddressingAddressMapIPService) Add(ctx context.Context, accountID string, addressMapID string, ipAddress string, body AccountAddressingAddressMapIPAddParams, opts ...option.RequestOption) (res *APIResponseCollectionAddressing, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if addressMapID == "" {
		err = errors.New("missing required address_map_id parameter")
		return
	}
	if ipAddress == "" {
		err = errors.New("missing required ip_address parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/addressing/address_maps/%s/ips/%s", accountID, addressMapID, ipAddress)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Remove an IP from a particular address map.
func (r *AccountAddressingAddressMapIPService) Remove(ctx context.Context, accountID string, addressMapID string, ipAddress string, body AccountAddressingAddressMapIPRemoveParams, opts ...option.RequestOption) (res *APIResponseCollectionAddressing, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if addressMapID == "" {
		err = errors.New("missing required address_map_id parameter")
		return
	}
	if ipAddress == "" {
		err = errors.New("missing required ip_address parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/addressing/address_maps/%s/ips/%s", accountID, addressMapID, ipAddress)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

type AccountAddressingAddressMapIPAddParams struct {
	Body interface{} `json:"body,required"`
}

func (r AccountAddressingAddressMapIPAddParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type AccountAddressingAddressMapIPRemoveParams struct {
	Body interface{} `json:"body,required"`
}

func (r AccountAddressingAddressMapIPRemoveParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}
