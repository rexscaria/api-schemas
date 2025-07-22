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

// AccountAddressingAddressMapAccountService contains methods and other services
// that help with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountAddressingAddressMapAccountService] method instead.
type AccountAddressingAddressMapAccountService struct {
	Options []option.RequestOption
}

// NewAccountAddressingAddressMapAccountService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewAccountAddressingAddressMapAccountService(opts ...option.RequestOption) (r *AccountAddressingAddressMapAccountService) {
	r = &AccountAddressingAddressMapAccountService{}
	r.Options = opts
	return
}

// Add an account as a member of a particular address map.
func (r *AccountAddressingAddressMapAccountService) Add(ctx context.Context, accountID string, addressMapID string, body AccountAddressingAddressMapAccountAddParams, opts ...option.RequestOption) (res *APIResponseCollectionAddressing, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if addressMapID == "" {
		err = errors.New("missing required address_map_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/addressing/address_maps/%s/accounts/%s", accountID, addressMapID, accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Remove an account as a member of a particular address map.
func (r *AccountAddressingAddressMapAccountService) Remove(ctx context.Context, accountID string, addressMapID string, body AccountAddressingAddressMapAccountRemoveParams, opts ...option.RequestOption) (res *APIResponseCollectionAddressing, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if addressMapID == "" {
		err = errors.New("missing required address_map_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/addressing/address_maps/%s/accounts/%s", accountID, addressMapID, accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

type AccountAddressingAddressMapAccountAddParams struct {
	Body interface{} `json:"body,required"`
}

func (r AccountAddressingAddressMapAccountAddParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type AccountAddressingAddressMapAccountRemoveParams struct {
	Body interface{} `json:"body,required"`
}

func (r AccountAddressingAddressMapAccountRemoveParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}
