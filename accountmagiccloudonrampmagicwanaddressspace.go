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

// AccountMagicCloudOnrampMagicWanAddressSpaceService contains methods and other
// services that help with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountMagicCloudOnrampMagicWanAddressSpaceService] method instead.
type AccountMagicCloudOnrampMagicWanAddressSpaceService struct {
	Options []option.RequestOption
}

// NewAccountMagicCloudOnrampMagicWanAddressSpaceService generates a new service
// that applies the given options to each request. These options are applied after
// the parent client's options (if there is one), and before any request-specific
// options.
func NewAccountMagicCloudOnrampMagicWanAddressSpaceService(opts ...option.RequestOption) (r *AccountMagicCloudOnrampMagicWanAddressSpaceService) {
	r = &AccountMagicCloudOnrampMagicWanAddressSpaceService{}
	r.Options = opts
	return
}

// Read the Magic WAN Address Space (Closed Beta)
func (r *AccountMagicCloudOnrampMagicWanAddressSpaceService) Get(ctx context.Context, accountID string, opts ...option.RequestOption) (res *AccountMagicCloudOnrampMagicWanAddressSpaceGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/cloud/onramps/magic_wan_address_space", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update the Magic WAN Address Space (Closed Beta)
func (r *AccountMagicCloudOnrampMagicWanAddressSpaceService) Update(ctx context.Context, accountID string, body AccountMagicCloudOnrampMagicWanAddressSpaceUpdateParams, opts ...option.RequestOption) (res *McnUpdateMagicWanAddressSpaceResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/cloud/onramps/magic_wan_address_space", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Update the Magic WAN Address Space (Closed Beta)
func (r *AccountMagicCloudOnrampMagicWanAddressSpaceService) Patch(ctx context.Context, accountID string, body AccountMagicCloudOnrampMagicWanAddressSpacePatchParams, opts ...option.RequestOption) (res *McnUpdateMagicWanAddressSpaceResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/cloud/onramps/magic_wan_address_space", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

type McnMagicWanAddressSpace struct {
	Prefixes []string                    `json:"prefixes,required"`
	JSON     mcnMagicWanAddressSpaceJSON `json:"-"`
}

// mcnMagicWanAddressSpaceJSON contains the JSON metadata for the struct
// [McnMagicWanAddressSpace]
type mcnMagicWanAddressSpaceJSON struct {
	Prefixes    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *McnMagicWanAddressSpace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r mcnMagicWanAddressSpaceJSON) RawJSON() string {
	return r.raw
}

type McnUpdateMagicWanAddressSpaceRequestParam struct {
	Prefixes param.Field[[]string] `json:"prefixes,required"`
}

func (r McnUpdateMagicWanAddressSpaceRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type McnUpdateMagicWanAddressSpaceResponse struct {
	Result McnMagicWanAddressSpace                   `json:"result"`
	JSON   mcnUpdateMagicWanAddressSpaceResponseJSON `json:"-"`
	McnGoodResponse
}

// mcnUpdateMagicWanAddressSpaceResponseJSON contains the JSON metadata for the
// struct [McnUpdateMagicWanAddressSpaceResponse]
type mcnUpdateMagicWanAddressSpaceResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *McnUpdateMagicWanAddressSpaceResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r mcnUpdateMagicWanAddressSpaceResponseJSON) RawJSON() string {
	return r.raw
}

type AccountMagicCloudOnrampMagicWanAddressSpaceGetResponse struct {
	Result McnMagicWanAddressSpace                                    `json:"result"`
	JSON   accountMagicCloudOnrampMagicWanAddressSpaceGetResponseJSON `json:"-"`
	McnGoodResponse
}

// accountMagicCloudOnrampMagicWanAddressSpaceGetResponseJSON contains the JSON
// metadata for the struct [AccountMagicCloudOnrampMagicWanAddressSpaceGetResponse]
type accountMagicCloudOnrampMagicWanAddressSpaceGetResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicCloudOnrampMagicWanAddressSpaceGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMagicCloudOnrampMagicWanAddressSpaceGetResponseJSON) RawJSON() string {
	return r.raw
}

type AccountMagicCloudOnrampMagicWanAddressSpaceUpdateParams struct {
	McnUpdateMagicWanAddressSpaceRequest McnUpdateMagicWanAddressSpaceRequestParam `json:"mcn_update_magic_wan_address_space_request,required"`
}

func (r AccountMagicCloudOnrampMagicWanAddressSpaceUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.McnUpdateMagicWanAddressSpaceRequest)
}

type AccountMagicCloudOnrampMagicWanAddressSpacePatchParams struct {
	McnUpdateMagicWanAddressSpaceRequest McnUpdateMagicWanAddressSpaceRequestParam `json:"mcn_update_magic_wan_address_space_request,required"`
}

func (r AccountMagicCloudOnrampMagicWanAddressSpacePatchParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.McnUpdateMagicWanAddressSpaceRequest)
}
