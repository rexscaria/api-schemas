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

// AccountZerotrustConnectivitySettingService contains methods and other services
// that help with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountZerotrustConnectivitySettingService] method instead.
type AccountZerotrustConnectivitySettingService struct {
	Options []option.RequestOption
}

// NewAccountZerotrustConnectivitySettingService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewAccountZerotrustConnectivitySettingService(opts ...option.RequestOption) (r *AccountZerotrustConnectivitySettingService) {
	r = &AccountZerotrustConnectivitySettingService{}
	r.Options = opts
	return
}

// Gets the Zero Trust Connectivity Settings for the given account.
func (r *AccountZerotrustConnectivitySettingService) Get(ctx context.Context, accountID string, opts ...option.RequestOption) (res *ConnectivitySettingsResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/zerotrust/connectivity_settings", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates the Zero Trust Connectivity Settings for the given account.
func (r *AccountZerotrustConnectivitySettingService) Update(ctx context.Context, accountID string, body AccountZerotrustConnectivitySettingUpdateParams, opts ...option.RequestOption) (res *ConnectivitySettingsResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/zerotrust/connectivity_settings", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

type ConnectivitySettingsResponse struct {
	Result ConnectivitySettingsResponseResult `json:"result"`
	JSON   connectivitySettingsResponseJSON   `json:"-"`
	APIResponseTunnel
}

// connectivitySettingsResponseJSON contains the JSON metadata for the struct
// [ConnectivitySettingsResponse]
type connectivitySettingsResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectivitySettingsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectivitySettingsResponseJSON) RawJSON() string {
	return r.raw
}

type ConnectivitySettingsResponseResult struct {
	// A flag to enable the ICMP proxy for the account network.
	IcmpProxyEnabled bool `json:"icmp_proxy_enabled"`
	// A flag to enable WARP to WARP traffic.
	OfframpWarpEnabled bool                                   `json:"offramp_warp_enabled"`
	JSON               connectivitySettingsResponseResultJSON `json:"-"`
}

// connectivitySettingsResponseResultJSON contains the JSON metadata for the struct
// [ConnectivitySettingsResponseResult]
type connectivitySettingsResponseResultJSON struct {
	IcmpProxyEnabled   apijson.Field
	OfframpWarpEnabled apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ConnectivitySettingsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectivitySettingsResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountZerotrustConnectivitySettingUpdateParams struct {
	// A flag to enable the ICMP proxy for the account network.
	IcmpProxyEnabled param.Field[bool] `json:"icmp_proxy_enabled"`
	// A flag to enable WARP to WARP traffic.
	OfframpWarpEnabled param.Field[bool] `json:"offramp_warp_enabled"`
}

func (r AccountZerotrustConnectivitySettingUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
