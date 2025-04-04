// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/stainless-sdks/cf-rex-go/internal/apijson"
	"github.com/stainless-sdks/cf-rex-go/internal/param"
	"github.com/stainless-sdks/cf-rex-go/internal/requestconfig"
	"github.com/stainless-sdks/cf-rex-go/option"
)

// AccountAddressingPrefixBgpStatusService contains methods and other services that
// help with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountAddressingPrefixBgpStatusService] method instead.
type AccountAddressingPrefixBgpStatusService struct {
	Options []option.RequestOption
}

// NewAccountAddressingPrefixBgpStatusService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountAddressingPrefixBgpStatusService(opts ...option.RequestOption) (r *AccountAddressingPrefixBgpStatusService) {
	r = &AccountAddressingPrefixBgpStatusService{}
	r.Options = opts
	return
}

// Advertise or withdraw the BGP route for a prefix.
//
// **Deprecated:** Prefer the BGP Prefixes endpoints, which additionally allow for
// advertising and withdrawing subnets of an IP prefix.
func (r *AccountAddressingPrefixBgpStatusService) Update(ctx context.Context, accountID string, prefixID string, body AccountAddressingPrefixBgpStatusUpdateParams, opts ...option.RequestOption) (res *AdvertisedResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if prefixID == "" {
		err = errors.New("missing required prefix_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/addressing/prefixes/%s/bgp/status", accountID, prefixID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// View the current advertisement state for a prefix.
//
// **Deprecated:** Prefer the BGP Prefixes endpoints, which additionally allow for
// advertising and withdrawing subnets of an IP prefix.
func (r *AccountAddressingPrefixBgpStatusService) Get(ctx context.Context, accountID string, prefixID string, opts ...option.RequestOption) (res *AdvertisedResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if prefixID == "" {
		err = errors.New("missing required prefix_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/addressing/prefixes/%s/bgp/status", accountID, prefixID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AdvertisedResponse struct {
	Result AdvertisedResponseResult `json:"result"`
	JSON   advertisedResponseJSON   `json:"-"`
	AddressingAPIResponseSingle
}

// advertisedResponseJSON contains the JSON metadata for the struct
// [AdvertisedResponse]
type advertisedResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AdvertisedResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advertisedResponseJSON) RawJSON() string {
	return r.raw
}

type AdvertisedResponseResult struct {
	// Advertisement status of the prefix. If `true`, the BGP route for the prefix is
	// advertised to the Internet. If `false`, the BGP route is withdrawn.
	Advertised bool `json:"advertised"`
	// Last time the advertisement status was changed. This field is only not 'null' if
	// on demand is enabled.
	AdvertisedModifiedAt time.Time                    `json:"advertised_modified_at,nullable" format:"date-time"`
	JSON                 advertisedResponseResultJSON `json:"-"`
}

// advertisedResponseResultJSON contains the JSON metadata for the struct
// [AdvertisedResponseResult]
type advertisedResponseResultJSON struct {
	Advertised           apijson.Field
	AdvertisedModifiedAt apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AdvertisedResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r advertisedResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountAddressingPrefixBgpStatusUpdateParams struct {
	// Advertisement status of the prefix. If `true`, the BGP route for the prefix is
	// advertised to the Internet. If `false`, the BGP route is withdrawn.
	Advertised param.Field[bool] `json:"advertised,required"`
}

func (r AccountAddressingPrefixBgpStatusUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
